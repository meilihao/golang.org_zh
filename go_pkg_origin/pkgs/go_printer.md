

# printer
`import "go/printer"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package printer implements printing of AST nodes.




## <a id="pkg-index">Index</a>
* [func Fprint(output io.Writer, fset *token.FileSet, node interface{}) error](#Fprint)
* [type CommentedNode](#CommentedNode)
* [type Config](#Config)
  * [func (cfg *Config) Fprint(output io.Writer, fset *token.FileSet, node interface{}) error](#Config.Fprint)
* [type Mode](#Mode)


#### <a id="pkg-examples">Examples</a>
* [Fprint](#example_Fprint)


#### <a id="pkg-files">Package files</a>
[nodes.go](https://golang.org/src/go/printer/nodes.go) [printer.go](https://golang.org/src/go/printer/printer.go) 






## <a id="Fprint">func</a> [Fprint](https://golang.org/src/go/printer/printer.go?s=39540:39614#L1349)
<pre>func Fprint(output <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, fset *<a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#FileSet">FileSet</a>, node interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Fprint "pretty-prints" an AST node to output.
It calls Config.Fprint with default settings.
Note that gofmt uses tabs for indentation but spaces for alignment;
use format.Node (package go/format) for output that matches gofmt.



<a id="example_Fprint">Example</a>


```go
```

output:
```txt
```



## <a id="CommentedNode">type</a> [CommentedNode](https://golang.org/src/go/printer/printer.go?s=38697:38837#L1330)
A CommentedNode bundles an AST node and corresponding comments.
It may be provided as argument to any of the Fprint functions.


<pre>type CommentedNode struct {
<span id="CommentedNode.Node"></span>    Node     interface{} <span class="comment">// *ast.File, or ast.Expr, ast.Decl, ast.Spec, or ast.Stmt</span>
<span id="CommentedNode.Comments"></span>    Comments []*<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#CommentGroup">CommentGroup</a>
}
</pre>











## <a id="Config">type</a> [Config](https://golang.org/src/go/printer/printer.go?s=36977:37131#L1272)
A Config node controls the output of Fprint.


<pre>type Config struct {
<span id="Config.Mode"></span>    Mode     <a href="#Mode">Mode</a> <span class="comment">// default: 0</span>
<span id="Config.Tabwidth"></span>    Tabwidth <a href="/pkg/builtin/#int">int</a>  <span class="comment">// default: 8</span>
<span id="Config.Indent"></span>    Indent   <a href="/pkg/builtin/#int">int</a>  <span class="comment">// default: 0 (all code is indented at least by this much)</span>
}
</pre>











### <a id="Config.Fprint">func</a> (\*Config) [Fprint](https://golang.org/src/go/printer/printer.go?s=39141:39229#L1340)
<pre>func (cfg *<a href="#Config">Config</a>) Fprint(output <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, fset *<a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#FileSet">FileSet</a>, node interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Fprint "pretty-prints" an AST node to output for a given configuration cfg.
Position information is interpreted relative to the file set fset.
The node type must be *ast.File, *CommentedNode, []ast.Decl, []ast.Stmt,
or assignment-compatible to ast.Expr, ast.Decl, ast.Spec, or ast.Stmt.




## <a id="Mode">type</a> [Mode](https://golang.org/src/go/printer/printer.go?s=36573:36587#L1262)
A Mode value is a set of flags (or 0). They control printing.


<pre>type Mode <a href="/pkg/builtin/#uint">uint</a></pre>



<pre>const (
    <span id="RawFormat">RawFormat</span> <a href="#Mode">Mode</a> = 1 &lt;&lt; <a href="/pkg/builtin/#iota">iota</a> <span class="comment">// do not use a tabwriter; if set, UseSpaces is ignored</span>
    <span id="TabIndent">TabIndent</span>                  <span class="comment">// use tabs for indentation independent of UseSpaces</span>
    <span id="UseSpaces">UseSpaces</span>                  <span class="comment">// use spaces instead of tabs for alignment</span>
    <span id="SourcePos">SourcePos</span>                  <span class="comment">// emit //line directives to preserve original source positions</span>
)</pre>












