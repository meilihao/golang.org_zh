

# parser
`import "go/parser"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package parser implements a parser for Go source files. Input may be
provided in a variety of forms (see the various Parse* functions); the
output is an abstract syntax tree (AST) representing the Go source. The
parser is invoked through one of the Parse* functions.

The parser accepts a larger language than is syntactically permitted by
the Go spec, for simplicity, and for improved robustness in the presence
of syntax errors. For instance, in method declarations, the receiver is
treated like an ordinary parameter list and thus may contain multiple
entries where the spec permits exactly one. Consequently, the corresponding
field in the AST (ast.FuncDecl.Recv) field is not restricted to one entry.




## <a id="pkg-index">Index</a>
* [func ParseDir(fset *token.FileSet, path string, filter func(os.FileInfo) bool, mode Mode) (pkgs map[string]*ast.Package, first error)](#ParseDir)
* [func ParseExpr(x string) (ast.Expr, error)](#ParseExpr)
* [func ParseExprFrom(fset *token.FileSet, filename string, src interface{}, mode Mode) (ast.Expr, error)](#ParseExprFrom)
* [func ParseFile(fset *token.FileSet, filename string, src interface{}, mode Mode) (f *ast.File, err error)](#ParseFile)
* [type Mode](#Mode)


#### <a id="pkg-examples">Examples</a>
* [ParseFile](#example_ParseFile)


#### <a id="pkg-files">Package files</a>
[interface.go](https://golang.org/src/go/parser/interface.go) [parser.go](https://golang.org/src/go/parser/parser.go) 






## <a id="ParseDir">func</a> [ParseDir](https://golang.org/src/go/parser/interface.go?s=4325:4458#L125)
<pre>func ParseDir(fset *<a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#FileSet">FileSet</a>, path <a href="/pkg/builtin/#string">string</a>, filter func(<a href="/pkg/os/">os</a>.<a href="/pkg/os/#FileInfo">FileInfo</a>) <a href="/pkg/builtin/#bool">bool</a>, mode <a href="#Mode">Mode</a>) (pkgs map[<a href="/pkg/builtin/#string">string</a>]*<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#Package">Package</a>, first <a href="/pkg/builtin/#error">error</a>)</pre>
ParseDir calls ParseFile for all files with names ending in ".go" in the
directory specified by path and returns a map of package name -> package
AST with all the packages found.

If filter != nil, only the files with os.FileInfo entries passing through
the filter (and ending in ".go") are considered. The mode bits are passed
to ParseFile unchanged. Position information is recorded in fset, which
must not be nil.

If the directory couldn't be read, a nil map and the respective error are
returned. If a parse error occurred, a non-nil but incomplete map and the
first error encountered are returned.



## <a id="ParseExpr">func</a> [ParseExpr](https://golang.org/src/go/parser/interface.go?s=6750:6792#L220)
<pre>func ParseExpr(x <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#Expr">Expr</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseExpr is a convenience function for obtaining the AST of an expression x.
The position information recorded in the AST is undefined. The filename used
in error messages is the empty string.



## <a id="ParseExprFrom">func</a> [ParseExprFrom](https://golang.org/src/go/parser/interface.go?s=5387:5489#L166)
<pre>func ParseExprFrom(fset *<a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#FileSet">FileSet</a>, filename <a href="/pkg/builtin/#string">string</a>, src interface{}, mode <a href="#Mode">Mode</a>) (<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#Expr">Expr</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseExprFrom is a convenience function for parsing an expression.
The arguments have the same meaning as for ParseFile, but the source must
be a valid Go (type or value) expression. Specifically, fset must not
be nil.



## <a id="ParseFile">func</a> [ParseFile](https://golang.org/src/go/parser/interface.go?s=2851:2956#L70)
<pre>func ParseFile(fset *<a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#FileSet">FileSet</a>, filename <a href="/pkg/builtin/#string">string</a>, src interface{}, mode <a href="#Mode">Mode</a>) (f *<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#File">File</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ParseFile parses the source code of a single Go source file and returns
the corresponding ast.File node. The source code may be provided via
the filename of the source file, or via the src parameter.

If src != nil, ParseFile parses the source from src and the filename is
only used when recording position information. The type of the argument
for the src parameter must be string, []byte, or io.Reader.
If src == nil, ParseFile parses the file specified by filename.

The mode parameter controls the amount of source text parsed and other
optional parser functionality. Position information is recorded in the
file set fset, which must not be nil.

If the source couldn't be read, the returned AST is nil and the error
indicates the specific failure. If the source was read but syntax
errors were found, the result is a partial AST (with ast.Bad* nodes
representing the fragments of erroneous source code). Multiple errors
are returned via a scanner.ErrorList which is sorted by file position.



<a id="example_ParseFile">Example</a>


```go
```

output:
```txt
```



## <a id="Mode">type</a> [Mode](https://golang.org/src/go/parser/interface.go?s=1140:1154#L39)
A Mode value is a set of flags (or 0).
They control the amount of source code parsed and other optional
parser functionality.


<pre>type Mode <a href="/pkg/builtin/#uint">uint</a></pre>



<pre>const (
    <span id="PackageClauseOnly">PackageClauseOnly</span> <a href="#Mode">Mode</a>             = 1 &lt;&lt; <a href="/pkg/builtin/#iota">iota</a> <span class="comment">// stop parsing after package clause</span>
    <span id="ImportsOnly">ImportsOnly</span>                                    <span class="comment">// stop parsing after import declarations</span>
    <span id="ParseComments">ParseComments</span>                                  <span class="comment">// parse comments and add them to AST</span>
    <span id="Trace">Trace</span>                                          <span class="comment">// print a trace of parsed productions</span>
    <span id="DeclarationErrors">DeclarationErrors</span>                              <span class="comment">// report declaration errors</span>
    <span id="SpuriousErrors">SpuriousErrors</span>                                 <span class="comment">// same as AllErrors, for backward-compatibility</span>
    <span id="AllErrors">AllErrors</span>         = <a href="#SpuriousErrors">SpuriousErrors</a>             <span class="comment">// report all errors (not just the first 10 on different lines)</span>
)</pre>












