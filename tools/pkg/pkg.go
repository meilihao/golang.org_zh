// 提取pkgs
// https://github.com/golang/tools/blob/master/godoc/static/packageroot.html
// ```
// $ dlv debug /home/chen/git/go/src/golang.org/x/tools/cmd/godoc -- -http=:6060
// (dlv) b /home/chen/git/go/src/golang.org/x/tools/godoc/server.go:334
// (dlv) print info.Dirs.List
// ```

// 提取具体pkg
// https://github.com/golang/tools/blob/master/godoc/static/package.html
// ```
// (dlv) b /home/chen/git/go/src/golang.org/x/tools/godoc/server.go:336
// (dlv) print info.PDoc
// ```
// [md template](https://github.com/davecheney/godoc2md) by `./godoc2md archive/tar > tar.md`

package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"text/template"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"golang.org/x/tools/godoc"
	"golang.org/x/tools/godoc/vfs"
)

var (
	client *http.Client

	pkgUrl  string
	version string
	isDebug bool

	// test in https://www.debuggex.com/
	goVersionReg    = regexp.MustCompile(`\"(.+)\"`)
	goDevVersionReg = regexp.MustCompile(`\+(.*?)\s`)
	paddingLeftReg  = regexp.MustCompile(`: (.+)px`)
)

// from godoc
var (
	pres *godoc.Presentation
	fs   = vfs.NameSpace{}
)

const (
	goVersionFlagStr    = "goVersion"
	goDevVersionFlagStr = "devel"

	tableHeader = `
| name / 名称   |  synopsis / 概述    |progress / 进度|
|--------------|:-------------------|--------------|
`
)

func init() {
	flag.StringVar(&pkgUrl, "pkg", "https://golang.org/pkg/", "go pkg url")
	flag.BoolVar(&isDebug, "debug", false, "generate debug subdir")

	flag.Parse()
}

func main() {
	InitHttpClient()

	resp, err := client.Get(pkgUrl)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Panic(err)
	}

	FindGoVersion(doc)

	log.Println(version)

	if isDebug {
		version = "test"
	}

	InitPresentation()

	FindPkgs(doc)
}

func InitHttpClient() {
	client = &http.Client{
		Timeout: time.Second * 7, //超时时间
	}

	// 使用`https_proxy=https://127.0.0.1:8123`, 即scheme=https时会阻塞, 去掉`https://`或换成`http://`可恢复正常
	tr := &http.Transport{
		//Proxy: http.ProxyURL(proxy),
		Proxy:           http.ProxyFromEnvironment, // is best
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client.Transport = tr
}

// 提取go verison
func FindGoVersion(doc *goquery.Document) {
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), goVersionFlagStr) {
			version = s.Text()
		}
	})

	if version == "" {
		log.Panic("no version")
	}

	if strings.Contains(version, goDevVersionFlagStr) {
		if tmp := goDevVersionReg.FindStringSubmatch(version); len(tmp) > 0 {
			version = goDevVersionFlagStr + "_" + tmp[1]
		}
	} else {
		if tmp := goVersionReg.FindStringSubmatch(version); len(tmp) > 0 {
			version = tmp[1]
		}
	}

	if version != "" { //按照version存储pkgs
		os.RemoveAll(version)
		os.MkdirAll(version, 0755)
	}
}

type PkgInfo struct {
	Name, Synopsis string
	Process        int    // 进度
	Level          byte   // 缩进层级
	Href           string // pkg url
	HrefFull       string // pkg full url, with pkgUrl
	StoreName      string // pkg 保存时的名称, 不包括`.md`后缀
}

func FindPkgs(doc *goquery.Document) {
	var nodes []*html.Node
	var paddingLeftStr string // pkg名称的缩进
	var hrefStr string        // pkg的具体页面
	var pkgs []*PkgInfo

	dirPkgs := "pkgs"
	dirSubPkg := filepath.Join(version, dirPkgs) // 存放具体pkg的目录
	os.RemoveAll(dirSubPkg)
	os.MkdirAll(dirSubPkg, 0755)

	doc.Find(".pkg-dir tr").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if i == 0 { // 跳过table header
			return true
		}

		if isDebug && i == 7 { // debug dirSubPkg
			return false
		}

		nodes = s.Children().Nodes

		if len(nodes) != 2 {
			log.Panic("incorrect format")
		}

		paddingLeftStr, _ = s.Find("td.pkg-name").Attr("style")
		hrefStr, _ = s.Find("td.pkg-name a").Attr("href")

		tmp := &PkgInfo{
			Name:      strings.TrimSpace(s.Find("td.pkg-name a").Text()),
			Synopsis:  strings.TrimSpace(s.Find("td.pkg-synopsis").Text()),
			Level:     ParsePaddingLeft(paddingLeftStr),
			Href:      strings.TrimSuffix(hrefStr, "/"),
			HrefFull:  pkgUrl + hrefStr,
			StoreName: strings.Replace(strings.TrimSuffix(hrefStr, "/"), "/", "_", -1),
		}

		if tmp.Synopsis == "" { // 没摘要, 没具体页面
			tmp.Process = -1
		} else {
			FindSubPkg(tmp.StoreName, tmp.Href, dirSubPkg)
		}

		pkgs = append(pkgs, tmp)

		log.Printf("%03d - %s", i, tmp.StoreName) // 打印进度

		return true
	})

	var output string = tableHeader
	for _, v := range pkgs {
		if v.Process == -1 {
			output += fmt.Sprintf("|%s|%s||\n", v.Name, v.Synopsis)
		} else {
			output += fmt.Sprintf("|[%s](%s)|%s|%d%%|\n", v.Name, dirPkgs+"/"+v.StoreName+".md", v.Synopsis, v.Process)
		}
	}

	f, err := os.Create(filepath.Join(version, "pkg.md"))
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	f.WriteString(output)
	f.Sync()
}

func FindSubPkg(name, subUrl, dir string) {
	content := InjectInfo(subUrl)

	if content == "" {
		log.Panic(name, subUrl, dir)
	}

	f, err := os.Create(filepath.Join(dir, fmt.Sprintf("%s.md", name)))
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	f.WriteString(content)
	f.Sync()
}

func ParsePaddingLeft(s string) byte {
	var level byte

	if tmp := paddingLeftReg.FindStringSubmatch(s); len(tmp) > 0 {
		switch tmp[1] {
		case "0":
			level = 1
		case "20":
			level = 2
		case "40":
			level = 3
		default:
			log.Panic("unknown padding left")
		}
	}

	return level
}

// 参考:
// - https://github.com/davecheney/godoc2md
// - https://github.com/golang/tools/blob/master/cmd/godoc/main.go#L182
func InitPresentation() {
	fs.Bind("/", vfs.OS(runtime.GOROOT()), "/", vfs.BindReplace)

	// // Bind $GOPATH trees into Go root.
	// for _, p := range filepath.SplitList(build.Default.GOPATH) {
	// 	fs.Bind("/src", vfs.OS(p), "/src", vfs.BindAfter)
	// }

	corpus := godoc.NewCorpus(fs)
	corpus.Verbose = true

	pres = godoc.NewPresentation(corpus)

	// from https://github.com/golang/tools/blob/master///cmd/godoc/handlers.go#L125
	pres.PackageHTML = readTemplate("package.html", pkgTemplate)
	pres.ExampleHTML = readTemplate("example.html", exampleTemplate)
}

func InjectInfo(subPath string) string {
	info := pres.GetPkgPageInfo("/src/"+subPath, subPath, 0)

	return string(applyTemplate(pres.PackageHTML, "packageHTML", info))
}

func applyTemplate(t *template.Template, name string, data interface{}) []byte {
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		log.Panicf("%s.Execute: %s", name, err)
	}
	return buf.Bytes()
}

func readTemplate(name, content string) *template.Template {
	// be explicit with errors (for app engine use)
	t, err := template.New(name).Funcs(pres.FuncMap()).Funcs(funcs).Parse(content)
	if err != nil {
		log.Fatal("readTemplate: ", err)
	}
	return t
}
