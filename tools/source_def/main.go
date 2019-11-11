package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	regSourceRef = regexp.MustCompile(`(?m)^#+\s+<a id="(.+)">.+</a>.+\((.+)\)$`)
)

type SourceRef struct {
	Raw string // regSourceRef匹配到字符串
	ID  string // Raw中<a>的id
	Url string // Raw中引用go源码的网址
}

type PkgSourceRef struct {
	base string
	m    map[string]map[string]SourceRef
}

func main() {
	ls := []string{"../../go_pkg_latest/pkgs", "../../go_pkg_origin/pkgs", "../../go_pkg/pkgs"}

	ps := make([]*PkgSourceRef, 0, len(ls))
	for _, v := range ls {
		ps = append(ps, &PkgSourceRef{
			base: v,
			m:    LoadSourceRef(v),
		})
	}

	Sync(ps[0], ps[1]) // sync to go_pkg_origin
	Sync(ps[0], ps[2]) // sync to go_pkg
}

func LoadSourceRef(base string) map[string]map[string]SourceRef {
	files, err := ioutil.ReadDir(base)
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]map[string]SourceRef)

	for _, fi := range files {
		if fi.IsDir() || !strings.HasSuffix(fi.Name(), ".md") {
			log.Fatal("invalid file:", base, fi.Name())
		}

		ds, err := pkgSourceRef(base, fi)
		if err != nil {
			log.Fatal("cant source def:", base, fi.Name(), err)
		}

		m[fi.Name()] = ds // fi.Name() like "archive_tar.md"
	}

	return m
}

func pkgSourceRef(base string, fi os.FileInfo) (map[string]SourceRef, error) {
	data, err := ioutil.ReadFile(filepath.Join(base, fi.Name()))
	if err != nil {
		return nil, err
	}

	m := map[string]SourceRef{}

	all := regSourceRef.FindAllStringSubmatch(string(data), -1)
	for _, v := range all {
		m[v[1]] = SourceRef{
			Raw: v[0], // like "### <a id="Format.String">func</a> (Format) [String](https://golang.org/src/archive/tar/format.go?s=4648:4679#L107)"
			ID:  v[1], // like "Format.String"
			Url: v[2], // like "https://golang.org/src/archive/tar/format.go?s=4648:4679#L107"
		}
	}

	return m, nil
}

func Sync(latest, dst *PkgSourceRef) {
	for k := range dst.m {
		sync(latest.m[k], dst.m[k], filepath.Join(dst.base, k))
	}
}

func sync(latest, dst map[string]SourceRef, filepath string) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	var has bool
	var latestRef SourceRef

	for k, v := range dst {
		latestRef, has = latest[k]
		if !has {
			log.Printf("cant't find %s#%s\n", filepath, k)

			continue
		}

		if latestRef.Raw == v.Raw {
			continue
		}

		data = bytes.ReplaceAll(data, []byte(v.Raw), []byte(latestRef.Raw))
	}

	ioutil.WriteFile(filepath, data, 0644)
}
