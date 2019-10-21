{{with .PDoc}}
{{if $.IsMain}}
{{comment_md .Doc}}
{{else}}
# {{ .Name }}
`import "{{.ImportPath}}"`

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
{{range .}}[{{.|filename|html}}]({{.|srcLink|html|appendSrcUrl}}) {{end}}
{{end}}

{{with .Consts}}## <a id="pkg-constants">Constants</a>
{{range .}}{{comment_md .Doc}}
<pre>{{node_html $ .Decl true}}</pre>{{end}}{{end}}

{{with .Vars}}## <a id="pkg-variables">Variables</a>
{{range .}}{{comment_md .Doc}}
<pre>{{node_html $ .Decl true}}</pre>{{end}}{{end}}

{{range .Funcs}}{{$name_html := html .Name}}## <a id="{{$name_html}}">func</a> [{{$name_html}}]({{posLink_url $ .Decl|appendSrcUrl}})
<pre>{{node_html $ .Decl true}}</pre>
{{comment_md .Doc}}
{{example_html $ .Name}}
{{callgraph_html $ "" .Name}}{{end}}

{{range .Types}}{{$tname := .Name}}{{$tname_html := html .Name}}## <a id="{{$tname_html}}">type</a> [{{$tname_html}}]({{posLink_url $ .Decl|appendSrcUrl}})
{{comment_md .Doc}}
<pre>{{node_html $ .Decl true}}</pre>

{{range .Consts}}
{{comment_md .Doc}}
<pre>{{node_html $ .Decl true}}</pre>{{end}}

{{range .Vars}}
{{comment_md .Doc}}
<pre>{{node_html $ .Decl true}}</pre>{{end}}

{{example_html $ $tname}}
{{implements_html $ $tname}}
{{methodset_html $ $tname}}

{{range .Funcs}}{{$name_html := html .Name}}### <a id="{{$name_html}}">func</a> [{{$name_html}}]({{posLink_url $ .Decl|appendSrcUrl}})
<pre>{{node_html $ .Decl true}}</pre>
{{comment_md .Doc}}
{{example_html $ .Name}}
{{callgraph_html $ "" .Name}}
{{end}}

{{range .Methods}}{{$name_html := html .Name}}### <a id="{{$tname_html}}.{{$name_html}}">func</a> ({{md .Recv}}) [{{$name_html}}]({{posLink_url $ .Decl|appendSrcUrl}})
<pre>{{node_html $ .Decl true}}</pre>
{{comment_md .Doc}}
{{$name := printf "%s_%s" $tname .Name}}{{example_html $ $name}}
{{callgraph_html $ .Recv .Name}}
{{end}}{{end}}{{end}}

{{with $.Notes}}
{{range $marker, $content := .}}
## <a id="pkg-note-{{$marker}}">{{noteTitle $marker | html}}s
{{range .}}
- {{comment_md .Body}}({{posLink_url $ .}})
{{end}}
{{end}}
{{end}}
{{end}}