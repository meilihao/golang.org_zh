

# template
`import "html/template"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package template (html/template) implements data-driven templates for
generating HTML output safe against code injection. It provides the
same interface as package text/template and should be used instead of
text/template whenever the output is HTML.

The documentation here focuses on the security features of the package.
For information about how to program the templates themselves, see the
documentation for text/template.

### Introduction
This package wraps package text/template so you can share its template API
to parse and execute HTML templates safely.


	tmpl, err := template.New("name").Parse(...)
	// Error checking elided
	err = tmpl.Execute(out, data)

If successful, tmpl will now be injection-safe. Otherwise, err is an error
defined in the docs for ErrorCode.

HTML templates treat data values as plain text which should be encoded so they
can be safely embedded in an HTML document. The escaping is contextual, so
actions can appear within JavaScript, CSS, and URI contexts.

The security model used by this package assumes that template authors are
trusted, while Execute's data parameter is not. More details are
provided below.

Example


	import "text/template"
	...
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")

produces


	Hello, <script>alert('you have been pwned')</script>!

but the contextual autoescaping in html/template


	import "html/template"
	...
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")

produces safe, escaped HTML output


	Hello, &lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;!

### Contexts
This package understands HTML, CSS, JavaScript, and URIs. It adds sanitizing
functions to each simple action pipeline, so given the excerpt


	<a href="/search?q={{.}}">{{.}}</a>

At parse time each {{.}} is overwritten to add escaping functions as necessary.
In this case it becomes


	<a href="/search?q={{. | urlescaper | attrescaper}}">{{. | htmlescaper}}</a>

where urlescaper, attrescaper, and htmlescaper are aliases for internal escaping
functions.

For these internal escaping functions, if an action pipeline evaluates to
a nil interface value, it is treated as though it were an empty string.

### Errors
See the documentation of ErrorCode for details.

### A fuller picture
The rest of this package comment may be skipped on first reading; it includes
details necessary to understand escaping contexts and error messages. Most users
will not need to understand these details.

### Contexts
Assuming {{.}} is `O'Reilly: How are <i>you</i>?`, the table below shows
how {{.}} appears when used in the context to the left.


	Context                          {{.}} After
	{{.}}                            O'Reilly: How are &lt;i&gt;you&lt;/i&gt;?
	<a title='{{.}}'>                O&#39;Reilly: How are you?
	<a href="/{{.}}">                O&#39;Reilly: How are %3ci%3eyou%3c/i%3e?
	<a href="?q={{.}}">              O&#39;Reilly%3a%20How%20are%3ci%3e...%3f
	<a onx='f("{{.}}")'>             O\x27Reilly: How are \x3ci\x3eyou...?
	<a onx='f({{.}})'>               "O\x27Reilly: How are \x3ci\x3eyou...?"
	<a onx='pattern = /{{.}}/;'>     O\x27Reilly: How are \x3ci\x3eyou...\x3f

If used in an unsafe context, then the value might be filtered out:


	Context                          {{.}} After
	<a href="{{.}}">                 #ZgotmplZ

since "O'Reilly:" is not an allowed protocol like "http:".

If {{.}} is the innocuous word, `left`, then it can appear more widely,


	Context                              {{.}} After
	{{.}}                                left
	<a title='{{.}}'>                    left
	<a href='{{.}}'>                     left
	<a href='/{{.}}'>                    left
	<a href='?dir={{.}}'>                left
	<a style="border-{{.}}: 4px">        left
	<a style="align: {{.}}">             left
	<a style="background: '{{.}}'>       left
	<a style="background: url('{{.}}')>  left
	<style>p.{{.}} {color:red}</style>   left

Non-string values can be used in JavaScript contexts.
If {{.}} is


	struct{A,B string}{ "foo", "bar" }

in the escaped template


	<script>var pair = {{.}};</script>

then the template output is


	<script>var pair = {"A": "foo", "B": "bar"};</script>

See package json to understand how non-string content is marshaled for
embedding in JavaScript contexts.

### Typed Strings
By default, this package assumes that all pipelines produce a plain text string.
It adds escaping pipeline stages necessary to correctly and safely embed that
plain text string in the appropriate context.

When a data value is not plain text, you can make sure it is not over-escaped
by marking it with its type.

Types HTML, JS, URL, and others from content.go can carry safe content that is
exempted from escaping.

The template


	Hello, {{.}}!

can be invoked with


	tmpl.Execute(out, template.HTML(`<b>World</b>`))

to produce


	Hello, <b>World</b>!

instead of the


	Hello, &lt;b&gt;World&lt;b&gt;!

that would have been produced if {{.}} was a regular string.

### Security Model
<a href="https://rawgit.com/mikesamuel/sanitized-jquery-templates/trunk/safetemplate.html#problem_definition">https://rawgit.com/mikesamuel/sanitized-jquery-templates/trunk/safetemplate.html#problem_definition</a> defines "safe" as used by this package.

This package assumes that template authors are trusted, that Execute's data
parameter is not, and seeks to preserve the properties below in the face
of untrusted data:

Structure Preservation Property:
"... when a template author writes an HTML tag in a safe templating language,
the browser will interpret the corresponding portion of the output as a tag
regardless of the values of untrusted data, and similarly for other structures
such as attribute boundaries and JS and CSS string boundaries."

Code Effect Property:
"... only code specified by the template author should run as a result of
injecting the template output into a page and all code specified by the
template author should run as a result of the same."

Least Surprise Property:
"A developer (or code reviewer) familiar with HTML, CSS, and JavaScript, who
knows that contextual autoescaping happens should be able to look at a {{.}}
and correctly infer what sanitization happens."


<a id="example_">Example</a>
<a id="example__autoescaping">Example (Autoescaping)</a>
<a id="example__escape">Example (Escape)</a>


## <a id="pkg-index">Index</a>
* [func HTMLEscape(w io.Writer, b []byte)](#HTMLEscape)
* [func HTMLEscapeString(s string) string](#HTMLEscapeString)
* [func HTMLEscaper(args ...interface{}) string](#HTMLEscaper)
* [func IsTrue(val interface{}) (truth, ok bool)](#IsTrue)
* [func JSEscape(w io.Writer, b []byte)](#JSEscape)
* [func JSEscapeString(s string) string](#JSEscapeString)
* [func JSEscaper(args ...interface{}) string](#JSEscaper)
* [func URLQueryEscaper(args ...interface{}) string](#URLQueryEscaper)
* [type CSS](#CSS)
* [type Error](#Error)
  * [func (e *Error) Error() string](#Error.Error)
* [type ErrorCode](#ErrorCode)
* [type FuncMap](#FuncMap)
* [type HTML](#HTML)
* [type HTMLAttr](#HTMLAttr)
* [type JS](#JS)
* [type JSStr](#JSStr)
* [type Srcset](#Srcset)
* [type Template](#Template)
  * [func Must(t *Template, err error) *Template](#Must)
  * [func New(name string) *Template](#New)
  * [func ParseFiles(filenames ...string) (*Template, error)](#ParseFiles)
  * [func ParseGlob(pattern string) (*Template, error)](#ParseGlob)
  * [func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error)](#Template.AddParseTree)
  * [func (t *Template) Clone() (*Template, error)](#Template.Clone)
  * [func (t *Template) DefinedTemplates() string](#Template.DefinedTemplates)
  * [func (t *Template) Delims(left, right string) *Template](#Template.Delims)
  * [func (t *Template) Execute(wr io.Writer, data interface{}) error](#Template.Execute)
  * [func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error](#Template.ExecuteTemplate)
  * [func (t *Template) Funcs(funcMap FuncMap) *Template](#Template.Funcs)
  * [func (t *Template) Lookup(name string) *Template](#Template.Lookup)
  * [func (t *Template) Name() string](#Template.Name)
  * [func (t *Template) New(name string) *Template](#Template.New)
  * [func (t *Template) Option(opt ...string) *Template](#Template.Option)
  * [func (t *Template) Parse(text string) (*Template, error)](#Template.Parse)
  * [func (t *Template) ParseFiles(filenames ...string) (*Template, error)](#Template.ParseFiles)
  * [func (t *Template) ParseGlob(pattern string) (*Template, error)](#Template.ParseGlob)
  * [func (t *Template) Templates() []*Template](#Template.Templates)
* [type URL](#URL)


#### <a id="pkg-examples">Examples</a>
* [Package](#example_)
* [Template.Delims](#example_Template_Delims)
* [Template (Block)](#example_Template_block)
* [Template (Glob)](#example_Template_glob)
* [Template (Helpers)](#example_Template_helpers)
* [Template (Parsefiles)](#example_Template_parsefiles)
* [Template (Share)](#example_Template_share)
* [Package (Autoescaping)](#example__autoescaping)
* [Package (Escape)](#example__escape)


#### <a id="pkg-files">Package files</a>
[attr.go](https://golang.org/src/html/template/attr.go) [attr_string.go](https://golang.org/src/html/template/attr_string.go) [content.go](https://golang.org/src/html/template/content.go) [context.go](https://golang.org/src/html/template/context.go) [css.go](https://golang.org/src/html/template/css.go) [delim_string.go](https://golang.org/src/html/template/delim_string.go) [doc.go](https://golang.org/src/html/template/doc.go) [element_string.go](https://golang.org/src/html/template/element_string.go) [error.go](https://golang.org/src/html/template/error.go) [escape.go](https://golang.org/src/html/template/escape.go) [html.go](https://golang.org/src/html/template/html.go) [js.go](https://golang.org/src/html/template/js.go) [jsctx_string.go](https://golang.org/src/html/template/jsctx_string.go) [state_string.go](https://golang.org/src/html/template/state_string.go) [template.go](https://golang.org/src/html/template/template.go) [transition.go](https://golang.org/src/html/template/transition.go) [url.go](https://golang.org/src/html/template/url.go) [urlpart_string.go](https://golang.org/src/html/template/urlpart_string.go) 






## <a id="HTMLEscape">func</a> [HTMLEscape](https://golang.org/src/html/template/escape.go?s=28583:28621#L845)
<pre>func HTMLEscape(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, b []<a href="/pkg/builtin/#byte">byte</a>)</pre>
HTMLEscape writes to w the escaped HTML equivalent of the plain text data b.



## <a id="HTMLEscapeString">func</a> [HTMLEscapeString](https://golang.org/src/html/template/escape.go?s=28736:28774#L850)
<pre>func HTMLEscapeString(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
HTMLEscapeString returns the escaped HTML equivalent of the plain text data s.



## <a id="HTMLEscaper">func</a> [HTMLEscaper](https://golang.org/src/html/template/escape.go?s=28919:28963#L856)
<pre>func HTMLEscaper(args ...interface{}) <a href="/pkg/builtin/#string">string</a></pre>
HTMLEscaper returns the escaped HTML equivalent of the textual
representation of its arguments.



## <a id="IsTrue">func</a> [IsTrue](https://golang.org/src/html/template/template.go?s=16151:16196#L478)
<pre>func IsTrue(val interface{}) (truth, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
IsTrue reports whether the value is 'true', in the sense of not the zero of its type,
and whether the value has a meaningful truth value. This is the definition of
truth used by if and other such actions.



## <a id="JSEscape">func</a> [JSEscape](https://golang.org/src/html/template/escape.go?s=29091:29127#L861)
<pre>func JSEscape(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, b []<a href="/pkg/builtin/#byte">byte</a>)</pre>
JSEscape writes to w the escaped JavaScript equivalent of the plain text data b.



## <a id="JSEscapeString">func</a> [JSEscapeString](https://golang.org/src/html/template/escape.go?s=29244:29280#L866)
<pre>func JSEscapeString(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
JSEscapeString returns the escaped JavaScript equivalent of the plain text data s.



## <a id="JSEscaper">func</a> [JSEscaper](https://golang.org/src/html/template/escape.go?s=29427:29469#L872)
<pre>func JSEscaper(args ...interface{}) <a href="/pkg/builtin/#string">string</a></pre>
JSEscaper returns the escaped JavaScript equivalent of the textual
representation of its arguments.



## <a id="URLQueryEscaper">func</a> [URLQueryEscaper](https://golang.org/src/html/template/escape.go?s=29655:29703#L878)
<pre>func URLQueryEscaper(args ...interface{}) <a href="/pkg/builtin/#string">string</a></pre>
URLQueryEscaper returns the escaped value of the textual representation of
its arguments in a form suitable for embedding in a URL query.





## <a id="CSS">type</a> [CSS](https://golang.org/src/html/template/content.go?s=920:930#L15)
CSS encapsulates known safe content that matches any of:


	1. The CSS3 stylesheet production, such as `p { color: purple }`.
	2. The CSS3 rule production, such as `a[href=~"https:"].foo#bar`.
	3. CSS3 declaration productions, such as `color: red; margin: 2px`.
	4. The CSS3 value production, such as `rgba(0, 0, 255, 127)`.

See <a href="https://www.w3.org/TR/css3-syntax/#parsing">https://www.w3.org/TR/css3-syntax/#parsing</a> and
<a href="https://web.archive.org/web/20090211114933/http://w3.org/TR/css3-syntax#style">https://web.archive.org/web/20090211114933/http://w3.org/TR/css3-syntax#style</a>

Use of this type presents a security risk:
the encapsulated content should come from a trusted source,
as it will be included verbatim in the template output.


<pre>type CSS <a href="/pkg/builtin/#string">string</a></pre>











## <a id="Error">type</a> [Error](https://golang.org/src/html/template/error.go?s=287:739#L3)
Error describes a problem encountered during template Escaping.


<pre>type Error struct {
<span id="Error.ErrorCode"></span>    <span class="comment">// ErrorCode describes the kind of error.</span>
    ErrorCode <a href="#ErrorCode">ErrorCode</a>
<span id="Error.Node"></span>    <span class="comment">// Node is the node that caused the problem, if known.</span>
    <span class="comment">// If not nil, it overrides Name and Line.</span>
    Node <a href="/pkg/text/template/parse/">parse</a>.<a href="/pkg/text/template/parse/#Node">Node</a>
<span id="Error.Name"></span>    <span class="comment">// Name is the name of the template in which the error was encountered.</span>
    Name <a href="/pkg/builtin/#string">string</a>
<span id="Error.Line"></span>    <span class="comment">// Line is the line number of the error in the template source or 0.</span>
    Line <a href="/pkg/builtin/#int">int</a>
<span id="Error.Description"></span>    <span class="comment">// Description is a human-readable description of the problem.</span>
    Description <a href="/pkg/builtin/#string">string</a>
}
</pre>











### <a id="Error.Error">func</a> (\*Error) [Error](https://golang.org/src/html/template/error.go?s=8400:8430#L206)
<pre>func (e *<a href="#Error">Error</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="ErrorCode">type</a> [ErrorCode](https://golang.org/src/html/template/error.go?s=785:803#L18)
ErrorCode is a code for a kind of error.


<pre>type ErrorCode <a href="/pkg/builtin/#int">int</a></pre>


We define codes for each error that manifests while escaping templates, but
escaped templates may also fail at runtime.

Output: "ZgotmplZ"
Example:


	<img src="{{.X}}">
	where {{.X}} evaluates to `javascript:...`

Discussion:


	"ZgotmplZ" is a special value that indicates that unsafe content reached a
	CSS or URL context at runtime. The output of the example will be
	  <img src="#ZgotmplZ">
	If the data comes from a trusted source, use content types to exempt it
	from filtering: URL(`javascript:...`).


<pre>const (
    <span class="comment">// OK indicates the lack of an error.</span>
    <span id="OK">OK</span> <a href="#ErrorCode">ErrorCode</a> = <a href="/pkg/builtin/#iota">iota</a>

    <span class="comment">// ErrAmbigContext: &#34;... appears in an ambiguous context within a URL&#34;</span>
    <span class="comment">// Example:</span>
    <span class="comment">//   &lt;a href=&#34;</span>
    <span class="comment">//      {{if .C}}</span>
    <span class="comment">//        /path/</span>
    <span class="comment">//      {{else}}</span>
    <span class="comment">//        /search?q=</span>
    <span class="comment">//      {{end}}</span>
    <span class="comment">//      {{.X}}</span>
    <span class="comment">//   &#34;&gt;</span>
    <span class="comment">// Discussion:</span>
    <span class="comment">//   {{.X}} is in an ambiguous URL context since, depending on {{.C}},</span>
    <span class="comment">//  it may be either a URL suffix or a query parameter.</span>
    <span class="comment">//   Moving {{.X}} into the condition removes the ambiguity:</span>
    <span class="comment">//   &lt;a href=&#34;{{if .C}}/path/{{.X}}{{else}}/search?q={{.X}}&#34;&gt;</span>
    <span id="ErrAmbigContext">ErrAmbigContext</span>

    <span class="comment">// ErrBadHTML: &#34;expected space, attr name, or end of tag, but got ...&#34;,</span>
    <span class="comment">//   &#34;... in unquoted attr&#34;, &#34;... in attribute name&#34;</span>
    <span class="comment">// Example:</span>
    <span class="comment">//   &lt;a href = /search?q=foo&gt;</span>
    <span class="comment">//   &lt;href=foo&gt;</span>
    <span class="comment">//   &lt;form na&lt;e=...&gt;</span>
    <span class="comment">//   &lt;option selected&lt;</span>
    <span class="comment">// Discussion:</span>
    <span class="comment">//   This is often due to a typo in an HTML element, but some runes</span>
    <span class="comment">//   are banned in tag names, attribute names, and unquoted attribute</span>
    <span class="comment">//   values because they can tickle parser ambiguities.</span>
    <span class="comment">//   Quoting all attributes is the best policy.</span>
    <span id="ErrBadHTML">ErrBadHTML</span>

    <span class="comment">// ErrBranchEnd: &#34;{{if}} branches end in different contexts&#34;</span>
    <span class="comment">// Example:</span>
    <span class="comment">//   {{if .C}}&lt;a href=&#34;{{end}}{{.X}}</span>
    <span class="comment">// Discussion:</span>
    <span class="comment">//   Package html/template statically examines each path through an</span>
    <span class="comment">//   {{if}}, {{range}}, or {{with}} to escape any following pipelines.</span>
    <span class="comment">//   The example is ambiguous since {{.X}} might be an HTML text node,</span>
    <span class="comment">//   or a URL prefix in an HTML attribute. The context of {{.X}} is</span>
    <span class="comment">//   used to figure out how to escape it, but that context depends on</span>
    <span class="comment">//   the run-time value of {{.C}} which is not statically known.</span>
    <span class="comment">//</span>
    <span class="comment">//   The problem is usually something like missing quotes or angle</span>
    <span class="comment">//   brackets, or can be avoided by refactoring to put the two contexts</span>
    <span class="comment">//   into different branches of an if, range or with. If the problem</span>
    <span class="comment">//   is in a {{range}} over a collection that should never be empty,</span>
    <span class="comment">//   adding a dummy {{else}} can help.</span>
    <span id="ErrBranchEnd">ErrBranchEnd</span>

    <span class="comment">// ErrEndContext: &#34;... ends in a non-text context: ...&#34;</span>
    <span class="comment">// Examples:</span>
    <span class="comment">//   &lt;div</span>
    <span class="comment">//   &lt;div title=&#34;no close quote&gt;</span>
    <span class="comment">//   &lt;script&gt;f()</span>
    <span class="comment">// Discussion:</span>
    <span class="comment">//   Executed templates should produce a DocumentFragment of HTML.</span>
    <span class="comment">//   Templates that end without closing tags will trigger this error.</span>
    <span class="comment">//   Templates that should not be used in an HTML context or that</span>
    <span class="comment">//   produce incomplete Fragments should not be executed directly.</span>
    <span class="comment">//</span>
    <span class="comment">//   {{define &#34;main&#34;}} &lt;script&gt;{{template &#34;helper&#34;}}&lt;/script&gt; {{end}}</span>
    <span class="comment">//   {{define &#34;helper&#34;}} document.write(&#39; &lt;div title=&#34; &#39;) {{end}}</span>
    <span class="comment">//</span>
    <span class="comment">//   &#34;helper&#34; does not produce a valid document fragment, so should</span>
    <span class="comment">//   not be Executed directly.</span>
    <span id="ErrEndContext">ErrEndContext</span>

    <span class="comment">// ErrNoSuchTemplate: &#34;no such template ...&#34;</span>
    <span class="comment">// Examples:</span>
    <span class="comment">//   {{define &#34;main&#34;}}&lt;div {{template &#34;attrs&#34;}}&gt;{{end}}</span>
    <span class="comment">//   {{define &#34;attrs&#34;}}href=&#34;{{.URL}}&#34;{{end}}</span>
    <span class="comment">// Discussion:</span>
    <span class="comment">//   Package html/template looks through template calls to compute the</span>
    <span class="comment">//   context.</span>
    <span class="comment">//   Here the {{.URL}} in &#34;attrs&#34; must be treated as a URL when called</span>
    <span class="comment">//   from &#34;main&#34;, but you will get this error if &#34;attrs&#34; is not defined</span>
    <span class="comment">//   when &#34;main&#34; is parsed.</span>
    <span id="ErrNoSuchTemplate">ErrNoSuchTemplate</span>

    <span class="comment">// ErrOutputContext: &#34;cannot compute output context for template ...&#34;</span>
    <span class="comment">// Examples:</span>
    <span class="comment">//   {{define &#34;t&#34;}}{{if .T}}{{template &#34;t&#34; .T}}{{end}}{{.H}}&#34;,{{end}}</span>
    <span class="comment">// Discussion:</span>
    <span class="comment">//   A recursive template does not end in the same context in which it</span>
    <span class="comment">//   starts, and a reliable output context cannot be computed.</span>
    <span class="comment">//   Look for typos in the named template.</span>
    <span class="comment">//   If the template should not be called in the named start context,</span>
    <span class="comment">//   look for calls to that template in unexpected contexts.</span>
    <span class="comment">//   Maybe refactor recursive templates to not be recursive.</span>
    <span id="ErrOutputContext">ErrOutputContext</span>

    <span class="comment">// ErrPartialCharset: &#34;unfinished JS regexp charset in ...&#34;</span>
    <span class="comment">// Example:</span>
    <span class="comment">//     &lt;script&gt;var pattern = /foo[{{.Chars}}]/&lt;/script&gt;</span>
    <span class="comment">// Discussion:</span>
    <span class="comment">//   Package html/template does not support interpolation into regular</span>
    <span class="comment">//   expression literal character sets.</span>
    <span id="ErrPartialCharset">ErrPartialCharset</span>

    <span class="comment">// ErrPartialEscape: &#34;unfinished escape sequence in ...&#34;</span>
    <span class="comment">// Example:</span>
    <span class="comment">//   &lt;script&gt;alert(&#34;\{{.X}}&#34;)&lt;/script&gt;</span>
    <span class="comment">// Discussion:</span>
    <span class="comment">//   Package html/template does not support actions following a</span>
    <span class="comment">//   backslash.</span>
    <span class="comment">//   This is usually an error and there are better solutions; for</span>
    <span class="comment">//   example</span>
    <span class="comment">//     &lt;script&gt;alert(&#34;{{.X}}&#34;)&lt;/script&gt;</span>
    <span class="comment">//   should work, and if {{.X}} is a partial escape sequence such as</span>
    <span class="comment">//   &#34;xA0&#34;, mark the whole sequence as safe content: JSStr(`\xA0`)</span>
    <span id="ErrPartialEscape">ErrPartialEscape</span>

    <span class="comment">// ErrRangeLoopReentry: &#34;on range loop re-entry: ...&#34;</span>
    <span class="comment">// Example:</span>
    <span class="comment">//   &lt;script&gt;var x = [{{range .}}&#39;{{.}},{{end}}]&lt;/script&gt;</span>
    <span class="comment">// Discussion:</span>
    <span class="comment">//   If an iteration through a range would cause it to end in a</span>
    <span class="comment">//   different context than an earlier pass, there is no single context.</span>
    <span class="comment">//   In the example, there is missing a quote, so it is not clear</span>
    <span class="comment">//   whether {{.}} is meant to be inside a JS string or in a JS value</span>
    <span class="comment">//   context. The second iteration would produce something like</span>
    <span class="comment">//</span>
    <span class="comment">//     &lt;script&gt;var x = [&#39;firstValue,&#39;secondValue]&lt;/script&gt;</span>
    <span id="ErrRangeLoopReentry">ErrRangeLoopReentry</span>

    <span class="comment">// ErrSlashAmbig: &#39;/&#39; could start a division or regexp.</span>
    <span class="comment">// Example:</span>
    <span class="comment">//   &lt;script&gt;</span>
    <span class="comment">//     {{if .C}}var x = 1{{end}}</span>
    <span class="comment">//     /-{{.N}}/i.test(x) ? doThis : doThat();</span>
    <span class="comment">//   &lt;/script&gt;</span>
    <span class="comment">// Discussion:</span>
    <span class="comment">//   The example above could produce `var x = 1/-2/i.test(s)...`</span>
    <span class="comment">//   in which the first &#39;/&#39; is a mathematical division operator or it</span>
    <span class="comment">//   could produce `/-2/i.test(s)` in which the first &#39;/&#39; starts a</span>
    <span class="comment">//   regexp literal.</span>
    <span class="comment">//   Look for missing semicolons inside branches, and maybe add</span>
    <span class="comment">//   parentheses to make it clear which interpretation you intend.</span>
    <span id="ErrSlashAmbig">ErrSlashAmbig</span>

    <span class="comment">// ErrPredefinedEscaper: &#34;predefined escaper ... disallowed in template&#34;</span>
    <span class="comment">// Example:</span>
    <span class="comment">//   &lt;div class={{. | html}}&gt;Hello&lt;div&gt;</span>
    <span class="comment">// Discussion:</span>
    <span class="comment">//   Package html/template already contextually escapes all pipelines to</span>
    <span class="comment">//   produce HTML output safe against code injection. Manually escaping</span>
    <span class="comment">//   pipeline output using the predefined escapers &#34;html&#34; or &#34;urlquery&#34; is</span>
    <span class="comment">//   unnecessary, and may affect the correctness or safety of the escaped</span>
    <span class="comment">//   pipeline output in Go 1.8 and earlier.</span>
    <span class="comment">//</span>
    <span class="comment">//   In most cases, such as the given example, this error can be resolved by</span>
    <span class="comment">//   simply removing the predefined escaper from the pipeline and letting the</span>
    <span class="comment">//   contextual autoescaper handle the escaping of the pipeline. In other</span>
    <span class="comment">//   instances, where the predefined escaper occurs in the middle of a</span>
    <span class="comment">//   pipeline where subsequent commands expect escaped input, e.g.</span>
    <span class="comment">//     {{.X | html | makeALink}}</span>
    <span class="comment">//   where makeALink does</span>
    <span class="comment">//     return `&lt;a href=&#34;`+input+`&#34;&gt;link&lt;/a&gt;`</span>
    <span class="comment">//   consider refactoring the surrounding template to make use of the</span>
    <span class="comment">//   contextual autoescaper, i.e.</span>
    <span class="comment">//     &lt;a href=&#34;{{.X}}&#34;&gt;link&lt;/a&gt;</span>
    <span class="comment">//</span>
    <span class="comment">//   To ease migration to Go 1.9 and beyond, &#34;html&#34; and &#34;urlquery&#34; will</span>
    <span class="comment">//   continue to be allowed as the last command in a pipeline. However, if the</span>
    <span class="comment">//   pipeline occurs in an unquoted attribute value context, &#34;html&#34; is</span>
    <span class="comment">//   disallowed. Avoid using &#34;html&#34; and &#34;urlquery&#34; entirely in new templates.</span>
    <span id="ErrPredefinedEscaper">ErrPredefinedEscaper</span>
)</pre>









## <a id="FuncMap">type</a> [FuncMap](https://golang.org/src/html/template/template.go?s=10293:10328#L326)
FuncMap is the type of the map defining the mapping from names to
functions. Each function must have either a single return value, or two
return values of which the second has type error. In that case, if the
second (error) argument evaluates to non-nil during execution, execution
terminates and Execute returns that error. FuncMap has the same base type
as FuncMap in "text/template", copied here so clients need not import
"text/template".


<pre>type FuncMap map[<a href="/pkg/builtin/#string">string</a>]interface{}</pre>











## <a id="HTML">type</a> [HTML](https://golang.org/src/html/template/content.go?s=1375:1386#L25)
HTML encapsulates a known safe HTML document fragment.
It should not be used for HTML from a third-party, or HTML with
unclosed tags or comments. The outputs of a sound HTML sanitizer
and a template escaped by this package are fine for use with HTML.

Use of this type presents a security risk:
the encapsulated content should come from a trusted source,
as it will be included verbatim in the template output.


<pre>type HTML <a href="/pkg/builtin/#string">string</a></pre>











## <a id="HTMLAttr">type</a> [HTMLAttr](https://golang.org/src/html/template/content.go?s=1662:1677#L33)
HTMLAttr encapsulates an HTML attribute from a trusted source,
for example, ` dir="ltr"`.

Use of this type presents a security risk:
the encapsulated content should come from a trusted source,
as it will be included verbatim in the template output.


<pre>type HTMLAttr <a href="/pkg/builtin/#string">string</a></pre>











## <a id="JS">type</a> [JS](https://golang.org/src/html/template/content.go?s=2549:2558#L51)
JS encapsulates a known safe EcmaScript5 Expression, for example,
`(x + y * z())`.
Template authors are responsible for ensuring that typed expressions
do not break the intended precedence and that there is no
statement/expression ambiguity as when passing an expression like
"{ foo: bar() }\n['foo']()", which is both a valid Expression and a
valid Program with a very different meaning.

Use of this type presents a security risk:
the encapsulated content should come from a trusted source,
as it will be included verbatim in the template output.

Using JS to include valid but untrusted JSON is not safe.
A safe alternative is to parse the JSON with json.Unmarshal and then
pass the resultant object into the template, where it will be
converted to sanitized JSON when presented in a JavaScript context.


<pre>type JS <a href="/pkg/builtin/#string">string</a></pre>











## <a id="JSStr">type</a> [JSStr](https://golang.org/src/html/template/content.go?s=3132:3144#L64)
JSStr encapsulates a sequence of characters meant to be embedded
between quotes in a JavaScript expression.
The string must match a series of StringCharacters:


	StringCharacter :: SourceCharacter but not `\` or LineTerminator
	                 | EscapeSequence

Note that LineContinuations are not allowed.
JSStr("foo\\nbar") is fine, but JSStr("foo\\\nbar") is not.

Use of this type presents a security risk:
the encapsulated content should come from a trusted source,
as it will be included verbatim in the template output.


<pre>type JSStr <a href="/pkg/builtin/#string">string</a></pre>











## <a id="Srcset">type</a> [Srcset](https://golang.org/src/html/template/content.go?s=3974:3987#L83)
Srcset encapsulates a known safe srcset attribute
(see <a href="https://w3c.github.io/html/semantics-embedded-content.html#element-attrdef-img-srcset">https://w3c.github.io/html/semantics-embedded-content.html#element-attrdef-img-srcset</a>).

Use of this type presents a security risk:
the encapsulated content should come from a trusted source,
as it will be included verbatim in the template output.


<pre>type Srcset <a href="/pkg/builtin/#string">string</a></pre>











## <a id="Template">type</a> [Template](https://golang.org/src/html/template/template.go?s=388:824#L9)
Template is a specialized Template from "text/template" that produces a safe
HTML document fragment.


<pre>type Template struct {

    <span class="comment">// The underlying template&#39;s parse tree, updated to be HTML-safe.</span>
<span id="Template.Tree"></span>    Tree *<a href="/pkg/text/template/parse/">parse</a>.<a href="/pkg/text/template/parse/#Tree">Tree</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Template_block">Example (Block)</a>
<a id="example_Template_glob">Example (Glob)</a>
<p>Here we demonstrate loading a set of templates from a directory.
</p><a id="example_Template_helpers">Example (Helpers)</a>
<p>This example demonstrates one way to share some templates
and use them in different contexts. In this variant we add multiple driver
templates by hand to an existing bundle of templates.
</p><a id="example_Template_parsefiles">Example (Parsefiles)</a>
<p>Here we demonstrate loading a set of templates from files in different directories
</p><a id="example_Template_share">Example (Share)</a>
<p>This example demonstrates how to use one group of driver
templates with distinct sets of helper templates.
</p>



### <a id="Must">func</a> [Must](https://golang.org/src/html/template/template.go?s=11683:11726#L360)
<pre>func Must(t *<a href="#Template">Template</a>, err <a href="/pkg/builtin/#error">error</a>) *<a href="#Template">Template</a></pre>
Must is a helper that wraps a call to a function returning (*Template, error)
and panics if the error is non-nil. It is intended for use in variable initializations
such as


	var t = template.Must(template.New("name").Parse("html"))




### <a id="New">func</a> [New](https://golang.org/src/html/template/template.go?s=8698:8729#L272)
<pre>func New(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Template">Template</a></pre>
New allocates a new HTML template with the given name.




### <a id="ParseFiles">func</a> [ParseFiles](https://golang.org/src/html/template/template.go?s=12335:12390#L376)
<pre>func ParseFiles(filenames ...<a href="/pkg/builtin/#string">string</a>) (*<a href="#Template">Template</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseFiles creates a new Template and parses the template definitions from
the named files. The returned template's name will have the (base) name and
(parsed) contents of the first file. There must be at least one file.
If an error occurs, parsing stops and the returned *Template is nil.

When parsing multiple files with the same name in different directories,
the last one mentioned will be the one that results.
For instance, ParseFiles("a/foo", "b/foo") stores "b/foo" as the template
named "foo", while "a/foo" is unavailable.




### <a id="ParseGlob">func</a> [ParseGlob](https://golang.org/src/html/template/template.go?s=14740:14789#L442)
<pre>func ParseGlob(pattern <a href="/pkg/builtin/#string">string</a>) (*<a href="#Template">Template</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseGlob creates a new Template and parses the template definitions from
the files identified by the pattern. The files are matched according to the
semantics of filepath.Match, and the pattern must match at least one file.
The returned template will have the (base) name and (parsed) contents of the
first file matched by the pattern. ParseGlob is equivalent to calling
ParseFiles with the list of files matched by the pattern.

When parsing multiple files with the same name in different directories,
the last one mentioned will be the one that results.






### <a id="Template.AddParseTree">func</a> (\*Template) [AddParseTree](https://golang.org/src/html/template/template.go?s=6869:6950#L205)
<pre>func (t *<a href="#Template">Template</a>) AddParseTree(name <a href="/pkg/builtin/#string">string</a>, tree *<a href="/pkg/text/template/parse/">parse</a>.<a href="/pkg/text/template/parse/#Tree">Tree</a>) (*<a href="#Template">Template</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
AddParseTree creates a new template with the name and parse tree
and associates it with t.

It returns an error if t or any associated template has already been executed.




### <a id="Template.Clone">func</a> (\*Template) [Clone](https://golang.org/src/html/template/template.go?s=7749:7794#L234)
<pre>func (t *<a href="#Template">Template</a>) Clone() (*<a href="#Template">Template</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Clone returns a duplicate of the template, including all associated
templates. The actual representation is not copied, but the name space of
associated templates is, so further calls to Parse in the copy will add
templates to the copy but not to the original. Clone can be used to prepare
common templates and use them with variant definitions for other templates
by adding the variants after the clone is made.

It returns an error if t has already been executed.




### <a id="Template.DefinedTemplates">func</a> (\*Template) [DefinedTemplates](https://golang.org/src/html/template/template.go?s=5364:5408#L159)
<pre>func (t *<a href="#Template">Template</a>) DefinedTemplates() <a href="/pkg/builtin/#string">string</a></pre>
DefinedTemplates returns a string listing the defined templates,
prefixed by the string "; defined templates are: ". If there are none,
it returns the empty string. Used to generate an error message.




### <a id="Template.Delims">func</a> (\*Template) [Delims](https://golang.org/src/html/template/template.go?s=11095:11150#L343)
<pre>func (t *<a href="#Template">Template</a>) Delims(left, right <a href="/pkg/builtin/#string">string</a>) *<a href="#Template">Template</a></pre>
Delims sets the action delimiters to the specified strings, to be used in
subsequent calls to Parse, ParseFiles, or ParseGlob. Nested template
definitions will inherit the settings. An empty delimiter stands for the
corresponding default: {{ or }}.
The return value is the template, so calls can be chained.


<a id="example_Template_Delims">Example</a>


### <a id="Template.Execute">func</a> (\*Template) [Execute](https://golang.org/src/html/template/template.go?s=3487:3551#L108)
<pre>func (t *<a href="#Template">Template</a>) Execute(wr <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, data interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Execute applies a parsed template to the specified data object,
writing the output to wr.
If an error occurs executing the template or writing its output,
execution stops, but partial results may already have been written to
the output writer.
A template may be executed safely in parallel, although if parallel
executions share a Writer the output may be interleaved.




### <a id="Template.ExecuteTemplate">func</a> (\*Template) [ExecuteTemplate](https://golang.org/src/html/template/template.go?s=4079:4164#L122)
<pre>func (t *<a href="#Template">Template</a>) ExecuteTemplate(wr <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, name <a href="/pkg/builtin/#string">string</a>, data interface{}) <a href="/pkg/builtin/#error">error</a></pre>
ExecuteTemplate applies the template associated with t that has the given
name to the specified data object and writes the output to wr.
If an error occurs executing the template or writing its output,
execution stops, but partial results may already have been written to
the output writer.
A template may be executed safely in parallel, although if parallel
executions share a Writer the output may be interleaved.




### <a id="Template.Funcs">func</a> (\*Template) [Funcs](https://golang.org/src/html/template/template.go?s=10664:10715#L333)
<pre>func (t *<a href="#Template">Template</a>) Funcs(funcMap <a href="#FuncMap">FuncMap</a>) *<a href="#Template">Template</a></pre>
Funcs adds the elements of the argument map to the template's function map.
It must be called before the template is parsed.
It panics if a value in the map is not a function with appropriate return
type. However, it is legal to overwrite elements of the map. The return
value is the template, so calls can be chained.




### <a id="Template.Lookup">func</a> (\*Template) [Lookup](https://golang.org/src/html/template/template.go?s=11312:11360#L350)
<pre>func (t *<a href="#Template">Template</a>) Lookup(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Template">Template</a></pre>
Lookup returns the template with the given name that is associated with t,
or nil if there is no such template.




### <a id="Template.Name">func</a> (\*Template) [Name](https://golang.org/src/html/template/template.go?s=9769:9801#L315)
<pre>func (t *<a href="#Template">Template</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the name of the template.




### <a id="Template.New">func</a> (\*Template) [New](https://golang.org/src/html/template/template.go?s=9283:9328#L292)
<pre>func (t *<a href="#Template">Template</a>) New(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Template">Template</a></pre>
New allocates a new HTML template associated with the given one
and with the same delimiters. The association, which is transitive,
allows one template to invoke another with a {{template}} action.

If a template with the given name already exists, the new HTML template
will replace it. The existing template will be reset and disassociated with
t.




### <a id="Template.Option">func</a> (\*Template) [Option](https://golang.org/src/html/template/template.go?s=2223:2273#L64)
<pre>func (t *<a href="#Template">Template</a>) Option(opt ...<a href="/pkg/builtin/#string">string</a>) *<a href="#Template">Template</a></pre>
Option sets options for the template. Options are described by
strings, either a simple string or "key=value". There can be at
most one equals sign in an option string. If the option string
is unrecognized or otherwise invalid, Option panics.

Known options:

missingkey: Control the behavior during execution if a map is
indexed with a key that is not present in the map.


	"missingkey=default" or "missingkey=invalid"
		The default behavior: Do nothing and continue execution.
		If printed, the result of the index operation is the string
		"<no value>".
	"missingkey=zero"
		The operation returns the zero value for the map type's element.
	"missingkey=error"
		Execution stops immediately with an error.




### <a id="Template.Parse">func</a> (\*Template) [Parse](https://golang.org/src/html/template/template.go?s=6073:6129#L174)
<pre>func (t *<a href="#Template">Template</a>) Parse(text <a href="/pkg/builtin/#string">string</a>) (*<a href="#Template">Template</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Parse parses text as a template body for t.
Named template definitions ({{define ...}} or {{block ...}} statements) in text
define additional templates associated with t and are removed from the
definition of t itself.

Templates can be redefined in successive calls to Parse,
before the first use of Execute on t or any associated template.
A template definition with a body containing only white space and comments
is considered empty and will not replace an existing template's body.
This allows using Parse to add new named template definitions without
overwriting the main template body.




### <a id="Template.ParseFiles">func</a> (\*Template) [ParseFiles](https://golang.org/src/html/template/template.go?s=12872:12941#L388)
<pre>func (t *<a href="#Template">Template</a>) ParseFiles(filenames ...<a href="/pkg/builtin/#string">string</a>) (*<a href="#Template">Template</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseFiles parses the named files and associates the resulting templates with
t. If an error occurs, parsing stops and the returned template is nil;
otherwise it is t. There must be at least one file.

When parsing multiple files with the same name in different directories,
the last one mentioned will be the one that results.

ParseFiles returns an error if t or any associated template has already been executed.




### <a id="Template.ParseGlob">func</a> (\*Template) [ParseGlob](https://golang.org/src/html/template/template.go?s=15406:15469#L456)
<pre>func (t *<a href="#Template">Template</a>) ParseGlob(pattern <a href="/pkg/builtin/#string">string</a>) (*<a href="#Template">Template</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseGlob parses the template definitions in the files identified by the
pattern and associates the resulting templates with t. The files are matched
according to the semantics of filepath.Match, and the pattern must match at
least one file. ParseGlob is equivalent to calling t.ParseFiles with the
list of files matched by the pattern.

When parsing multiple files with the same name in different directories,
the last one mentioned will be the one that results.

ParseGlob returns an error if t or any associated template has already been executed.




### <a id="Template.Templates">func</a> (\*Template) [Templates](https://golang.org/src/html/template/template.go?s=1222:1264#L34)
<pre>func (t *<a href="#Template">Template</a>) Templates() []*<a href="#Template">Template</a></pre>
Templates returns a slice of the templates associated with t, including t
itself.




## <a id="URL">type</a> [URL](https://golang.org/src/html/template/content.go?s=3635:3645#L75)
URL encapsulates a known safe URL or URL substring (see RFC 3986).
A URL like `javascript:checkThatFormNotEditedBeforeLeavingPage()`
from a trusted source should go in the page, but by default dynamic
`javascript:` URLs are filtered out since they are a frequently
exploited injection vector.

Use of this type presents a security risk:
the encapsulated content should come from a trusted source,
as it will be included verbatim in the template output.


<pre>type URL <a href="/pkg/builtin/#string">string</a></pre>















