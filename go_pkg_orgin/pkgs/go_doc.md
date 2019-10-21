

# doc
`import "go/doc"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package doc extracts source code documentation from a Go AST.




## <a id="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func IsPredeclared(s string) bool](#IsPredeclared)
* [func Synopsis(s string) string](#Synopsis)
* [func ToHTML(w io.Writer, text string, words map[string]string)](#ToHTML)
* [func ToText(w io.Writer, text string, indent, preIndent string, width int)](#ToText)
* [type Example](#Example)
  * [func Examples(files ...*ast.File) []*Example](#Examples)
* [type Filter](#Filter)
* [type Func](#Func)
* [type Mode](#Mode)
* [type Note](#Note)
* [type Package](#Package)
  * [func New(pkg *ast.Package, importPath string, mode Mode) *Package](#New)
  * [func (p *Package) Filter(f Filter)](#Package.Filter)
* [type Type](#Type)
* [type Value](#Value)




#### <a id="pkg-files">Package files</a>
[comment.go](https://golang.org/src/go/doc/comment.go) [doc.go](https://golang.org/src/go/doc/doc.go) [example.go](https://golang.org/src/go/doc/example.go) [exports.go](https://golang.org/src/go/doc/exports.go) [filter.go](https://golang.org/src/go/doc/filter.go) [reader.go](https://golang.org/src/go/doc/reader.go) [synopsis.go](https://golang.org/src/go/doc/synopsis.go) 




## <a id="pkg-variables">Variables</a>

<pre>var <span id="IllegalPrefixes">IllegalPrefixes</span> = []<a href="/pkg/builtin/#string">string</a>{
    &#34;copyright&#34;,
    &#34;all rights&#34;,
    &#34;author&#34;,
}</pre>

## <a id="IsPredeclared">func</a> [IsPredeclared](https://golang.org/src/go/doc/reader.go?s=24453:24486#L857)
<pre>func IsPredeclared(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsPredeclared reports whether s is a predeclared identifier.



## <a id="Synopsis">func</a> [Synopsis](https://golang.org/src/go/doc/synopsis.go?s=1726:1756#L58)
<pre>func Synopsis(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Synopsis returns a cleaned version of the first sentence in s.
That sentence ends after the first period followed by space and
not preceded by exactly one uppercase letter. The result string
has no \n, \r, or \t characters and uses only single spaces between
words. If s starts with any of the IllegalPrefixes, the result
is the empty string.



## <a id="ToHTML">func</a> [ToHTML](https://golang.org/src/go/doc/comment.go?s=8221:8283#L296)
<pre>func ToHTML(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, text <a href="/pkg/builtin/#string">string</a>, words map[<a href="/pkg/builtin/#string">string</a>]<a href="/pkg/builtin/#string">string</a>)</pre>
ToHTML converts comment text to formatted HTML.
The comment was prepared by DocReader,
so it is known not to have leading, trailing blank lines
nor to have trailing spaces at the end of lines.
The comment markers have already been removed.

Each span of unindented non-blank lines is converted into
a single paragraph. There is one exception to the rule: a span that
consists of a single line, is followed by another paragraph span,
begins with a capital letter, and contains no punctuation
other than parentheses and commas is formatted as a heading.

A span of indented lines is converted into a <pre> block,
with the common indent prefix removed.

URLs in the comment text are converted into links; if the URL also appears
in the words map, the link is taken from the map (if the corresponding map
value is the empty string, the URL is not converted into a link).

Go identifiers that appear in the words map are italicized; if the corresponding
map value is not the empty string, it is considered a URL and the word is converted
into a link.



## <a id="ToText">func</a> [ToText](https://golang.org/src/go/doc/comment.go?s=10646:10720#L410)
<pre>func ToText(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, text <a href="/pkg/builtin/#string">string</a>, indent, preIndent <a href="/pkg/builtin/#string">string</a>, width <a href="/pkg/builtin/#int">int</a>)</pre>
ToText prepares comment text for presentation in textual output.
It wraps paragraphs of text to width or fewer Unicode code points
and then prefixes each line with the indent. In preformatted sections
(such as program text), it prefixes each non-blank line with preIndent.





## <a id="Example">type</a> [Example](https://golang.org/src/go/doc/example.go?s=411:809#L12)
An Example represents an example function found in a source files.


<pre>type Example struct {
<span id="Example.Name"></span>    Name        <a href="/pkg/builtin/#string">string</a> <span class="comment">// name of the item being exemplified</span>
<span id="Example.Doc"></span>    Doc         <a href="/pkg/builtin/#string">string</a> <span class="comment">// example function doc string</span>
<span id="Example.Code"></span>    Code        <a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#Node">Node</a>
<span id="Example.Play"></span>    Play        *<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#File">File</a> <span class="comment">// a whole program version of the example</span>
<span id="Example.Comments"></span>    Comments    []*<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#CommentGroup">CommentGroup</a>
<span id="Example.Output"></span>    Output      <a href="/pkg/builtin/#string">string</a> <span class="comment">// expected output</span>
<span id="Example.Unordered"></span>    Unordered   <a href="/pkg/builtin/#bool">bool</a>
<span id="Example.EmptyOutput"></span>    EmptyOutput <a href="/pkg/builtin/#bool">bool</a> <span class="comment">// expect empty output</span>
<span id="Example.Order"></span>    Order       <a href="/pkg/builtin/#int">int</a>  <span class="comment">// original source code order</span>
}
</pre>









### <a id="Examples">func</a> [Examples](https://golang.org/src/go/doc/example.go?s=1601:1645#L37)
<pre>func Examples(files ...*<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#File">File</a>) []*<a href="#Example">Example</a></pre>
Examples returns the examples found in the files, sorted by Name field.
The Order fields record the order in which the examples were encountered.

Playable Examples must be in a package whose name ends in "_test".
An Example is "playable" (the Play field is non-nil) in either of these
circumstances:


	- The example function is self-contained: the function references only
	  identifiers from other packages (or predeclared identifiers, such as
	  "int") and the test file does not include a dot import.
	- The entire test file is the example: the file contains exactly one
	  example function, zero test or benchmark functions, and at least one
	  top-level function, type, variable, or constant declaration other
	  than the example function.






## <a id="Filter">type</a> [Filter](https://golang.org/src/go/doc/filter.go?s=190:219#L1)

<pre>type Filter func(<a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>











## <a id="Func">type</a> [Func](https://golang.org/src/go/doc/doc.go?s=1393:1686#L46)
Func is the documentation for a func declaration.


<pre>type Func struct {
<span id="Func.Doc"></span>    Doc  <a href="/pkg/builtin/#string">string</a>
<span id="Func.Name"></span>    Name <a href="/pkg/builtin/#string">string</a>
<span id="Func.Decl"></span>    Decl *<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#FuncDecl">FuncDecl</a>

    <span class="comment">// methods</span>
    <span class="comment">// (for functions, these fields have the respective zero value)</span>
<span id="Func.Recv"></span>    Recv  <a href="/pkg/builtin/#string">string</a> <span class="comment">// actual   receiver &#34;T&#34; or &#34;*T&#34;</span>
<span id="Func.Orig"></span>    Orig  <a href="/pkg/builtin/#string">string</a> <span class="comment">// original receiver &#34;T&#34; or &#34;*T&#34;</span>
<span id="Func.Level"></span>    Level <a href="/pkg/builtin/#int">int</a>    <span class="comment">// embedding level; 0 means not embedded</span>
}
</pre>











## <a id="Mode">type</a> [Mode](https://golang.org/src/go/doc/doc.go?s=2229:2242#L69)
Mode values control the operation of New.


<pre>type Mode <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span class="comment">// AllDecls says to extract documentation for all package-level</span>
    <span class="comment">// declarations, not just exported ones.</span>
    <span id="AllDecls">AllDecls</span> <a href="#Mode">Mode</a> = 1 &lt;&lt; <a href="/pkg/builtin/#iota">iota</a>

    <span class="comment">// AllMethods says to show all embedded methods, not just the ones of</span>
    <span class="comment">// invisible (unexported) anonymous fields.</span>
    <span id="AllMethods">AllMethods</span>

    <span class="comment">// PreserveAST says to leave the AST unmodified. Originally, pieces of</span>
    <span class="comment">// the AST such as function bodies were nil-ed out to save memory in</span>
    <span class="comment">// godoc, but not all programs want that behavior.</span>
    <span id="PreserveAST">PreserveAST</span>
)</pre>









## <a id="Note">type</a> [Note](https://golang.org/src/go/doc/doc.go?s=2000:2182#L62)
A Note represents a marked comment starting with "MARKER(uid): note body".
Any note with a marker of 2 or more upper case [A-Z] letters and a uid of
at least one character is recognized. The ":" following the uid is optional.
Notes are collected in the Package.Notes map indexed by the notes marker.


<pre>type Note struct {
<span id="Note.Pos"></span>    Pos, End <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position range of the comment containing the marker</span>
<span id="Note.UID"></span>    UID      <a href="/pkg/builtin/#string">string</a>    <span class="comment">// uid found with the marker</span>
<span id="Note.Body"></span>    Body     <a href="/pkg/builtin/#string">string</a>    <span class="comment">// note body text</span>
}
</pre>











## <a id="Package">type</a> [Package](https://golang.org/src/go/doc/doc.go?s=327:695#L4)
Package is the documentation for an entire package.


<pre>type Package struct {
<span id="Package.Doc"></span>    Doc        <a href="/pkg/builtin/#string">string</a>
<span id="Package.Name"></span>    Name       <a href="/pkg/builtin/#string">string</a>
<span id="Package.ImportPath"></span>    ImportPath <a href="/pkg/builtin/#string">string</a>
<span id="Package.Imports"></span>    Imports    []<a href="/pkg/builtin/#string">string</a>
<span id="Package.Filenames"></span>    Filenames  []<a href="/pkg/builtin/#string">string</a>
<span id="Package.Notes"></span>    Notes      map[<a href="/pkg/builtin/#string">string</a>][]*<a href="#Note">Note</a>

    <span class="comment">// Deprecated: For backward compatibility Bugs is still populated,</span>
    <span class="comment">// but all new code should use Notes instead.</span>
<span id="Package.Bugs"></span>    Bugs []<a href="/pkg/builtin/#string">string</a>

    <span class="comment">// declarations</span>
<span id="Package.Consts"></span>    Consts []*<a href="#Value">Value</a>
<span id="Package.Types"></span>    Types  []*<a href="#Type">Type</a>
<span id="Package.Vars"></span>    Vars   []*<a href="#Value">Value</a>
<span id="Package.Funcs"></span>    Funcs  []*<a href="#Func">Func</a>
}
</pre>









### <a id="New">func</a> [New](https://golang.org/src/go/doc/doc.go?s=2866:2931#L89)
<pre>func New(pkg *<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#Package">Package</a>, importPath <a href="/pkg/builtin/#string">string</a>, mode <a href="#Mode">Mode</a>) *<a href="#Package">Package</a></pre>
New computes the package documentation for the given package AST.
New takes ownership of the AST pkg and may edit or overwrite it.






### <a id="Package.Filter">func</a> (\*Package) [Filter](https://golang.org/src/go/doc/filter.go?s=1932:1966#L89)
<pre>func (p *<a href="#Package">Package</a>) Filter(f <a href="#Filter">Filter</a>)</pre>
Filter eliminates documentation for names that don't pass through the filter f.
TODO(gri): Recognize "Type.Method" as a name.




## <a id="Type">type</a> [Type](https://golang.org/src/go/doc/doc.go?s=959:1338#L33)
Type is the documentation for a type declaration.


<pre>type Type struct {
<span id="Type.Doc"></span>    Doc  <a href="/pkg/builtin/#string">string</a>
<span id="Type.Name"></span>    Name <a href="/pkg/builtin/#string">string</a>
<span id="Type.Decl"></span>    Decl *<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#GenDecl">GenDecl</a>

    <span class="comment">// associated declarations</span>
<span id="Type.Consts"></span>    Consts  []*<a href="#Value">Value</a> <span class="comment">// sorted list of constants of (mostly) this type</span>
<span id="Type.Vars"></span>    Vars    []*<a href="#Value">Value</a> <span class="comment">// sorted list of variables of (mostly) this type</span>
<span id="Type.Funcs"></span>    Funcs   []*<a href="#Func">Func</a>  <span class="comment">// sorted list of functions returning this type</span>
<span id="Type.Methods"></span>    Methods []*<a href="#Func">Func</a>  <span class="comment">// sorted list of methods (including embedded ones) of this type</span>
}
</pre>











## <a id="Value">type</a> [Value](https://golang.org/src/go/doc/doc.go?s=778:904#L24)
Value is the documentation for a (possibly grouped) var or const declaration.


<pre>type Value struct {
<span id="Value.Doc"></span>    Doc   <a href="/pkg/builtin/#string">string</a>
<span id="Value.Names"></span>    Names []<a href="/pkg/builtin/#string">string</a> <span class="comment">// var or const names in declaration order</span>
<span id="Value.Decl"></span>    Decl  *<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#GenDecl">GenDecl</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>














