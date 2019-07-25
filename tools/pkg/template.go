// https://github.com/davecheney/godoc2md
package main

import (
	"bytes"
	"strings"
)

var (
	funcs = map[string]interface{}{
		"comment_md":  commentMdFunc,
		"md":          mdFunc,
		"pre":         preFunc,
		"kebab":       kebabFunc,
		"bitscape":    bitscapeFunc, //Escape [] for bitbucket confusion
		"trim_prefix": strings.TrimPrefix,
	}
)

func commentMdFunc(comment string) string {
	var buf bytes.Buffer
	ToMD(&buf, comment)
	return buf.String()
}

func mdFunc(text string) string {
	text = strings.Replace(text, "*", "\\*", -1)
	text = strings.Replace(text, "_", "\\_", -1)
	return text
}

func preFunc(text string) string {
	return "``` go\n" + text + "\n```"
}

func kebabFunc(text string) string {
	s := strings.Replace(strings.ToLower(text), " ", "-", -1)
	s = strings.Replace(s, ".", "-", -1)
	s = strings.Replace(s, "\\*", "42", -1)
	return s
}

func bitscapeFunc(text string) string {
	s := strings.Replace(text, "[", "\\[", -1)
	s = strings.Replace(s, "]", "\\]", -1)
	return s
}

var pkgTemplate = `{{with .PDoc}}
{{if $.IsMain}}
{{comment_md .Doc}}
{{else}}
# {{ .Name }}
` + "`" + `import "{{.ImportPath}}"` + "`" + `

* [Overview](#pkg-overview)
* [Index](#pkg-index){{if $.Examples}}
* [Examples](#pkg-examples){{- end}}{{if $.Dirs}}
* [Subdirectories](#pkg-subdirectories){{- end}}

## <a id="pkg-overview">Overview</a>
{{comment_md .Doc}}
{{example_html $ ""}}

## <a id="pkg-index">Index</a>{{if .Consts}}
* [Constants](#pkg-constants){{end}}{{if .Vars}}
* [Variables](#pkg-variables){{end}}{{- range .Funcs -}}{{$name_html := html .Name}}
* [{{node_html $ .Decl false | sanitize}}](#{{$name_html}}){{- end}}{{- range .Types}}{{$tname_html := html .Name}}
* [type {{$tname_html}}](#{{$tname_html}}){{- range .Funcs}}{{$name_html := html .Name}}
  * [{{node_html $ .Decl false | sanitize}}](#{{$name_html}}){{- end}}{{- range .Methods}}{{$name_html := html .Name}}
  * [{{node_html $ .Decl false | sanitize}}](#{{$tname_html}}.{{$name_html}}){{- end}}{{- end}}{{- if $.Notes}}{{- range $marker, $item := $.Notes}}
* [{{noteTitle $marker | html}}s](#pkg-note-{{$marker}}){{end}}{{end}}
{{if $.Examples}}
#### <a id="pkg-examples">Examples</a>{{- range $.Examples}}
* [{{example_name .Name}}](#example_{{.Name}}){{- end}}{{- end}}
{{with .Filenames}}
#### <a id="pkg-files">Package files</a>
{{range .}}[{{.|filename|html}}]({{.|srcLink|html}}) {{end}}
{{end}}

{{with .Consts}}## <a id="pkg-constants">Constants</a>
{{range .}}{{node $ .Decl | pre}}
{{comment_md .Doc}}{{end}}{{end}}
{{with .Vars}}## <a id="pkg-variables">Variables</a>
{{range .}}{{node $ .Decl | pre}}
{{comment_md .Doc}}{{end}}{{end}}

{{range .Funcs}}{{$name_html := html .Name}}## <a id="{{$name_html}}">func</a> [{{$name_html}}]({{posLink_url $ .Decl}})
{{node $ .Decl | pre}}
{{comment_md .Doc}}
{{example_html $ .Name}}
{{callgraph_html $ "" .Name}}{{end}}
{{range .Types}}{{$tname := .Name}}{{$tname_html := html .Name}}## <a id="{{$tname_html}}">type</a> [{{$tname_html}}]({{posLink_url $ .Decl}})
{{node $ .Decl | pre}}
{{comment_md .Doc}}{{range .Consts}}
{{node $ .Decl | pre }}
{{comment_md .Doc}}{{end}}{{range .Vars}}
{{node $ .Decl | pre }}
{{comment_md .Doc}}{{end}}

{{example_html $ $tname}}
{{implements_html $ $tname}}
{{methodset_html $ $tname}}

{{range .Funcs}}{{$name_html := html .Name}}### <a id="{{$name_html}}">func</a> [{{$name_html}}]({{posLink_url $ .Decl}})
{{node $ .Decl | pre}}
{{comment_md .Doc}}
{{example_html $ .Name}}{{end}}
{{callgraph_html $ "" .Name}}

{{range .Methods}}{{$name_html := html .Name}}### <a id="{{$tname_html}}.{{$name_html}}">func</a> ({{md .Recv}}) [{{$name_html}}]({{posLink_url $ .Decl}})
{{node $ .Decl | pre}}
{{comment_md .Doc}}
{{$name := printf "%s_%s" $tname .Name}}{{example_html $ $name}}
{{callgraph_html $ .Recv .Name}}
{{end}}{{end}}{{end}}

{{with $.Notes}}
{{range $marker, $content := .}}
## <a id="pkg-note-{{$marker}}">{{noteTitle $marker | html}}s
<ul style="list-style: none; padding: 0;">
{{range .}}
<li><a href="{{posLink_url $ .}}">&#x261e;</a> {{html .Body}}</li>
{{end}}
</ul>
{{end}}
{{end}}
{{end}}
`
