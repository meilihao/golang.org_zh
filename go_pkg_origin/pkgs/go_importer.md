

# importer
`import "go/importer"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package importer provides access to export data importers.




## <a id="pkg-index">Index</a>
* [func Default() types.Importer](#Default)
* [func For(compiler string, lookup Lookup) types.Importer](#For)
* [func ForCompiler(fset *token.FileSet, compiler string, lookup Lookup) types.Importer](#ForCompiler)
* [type Lookup](#Lookup)




#### <a id="pkg-files">Package files</a>
[importer.go](https://golang.org/src/go/importer/importer.go) 






## <a id="Default">func</a> [Default](https://golang.org/src/go/importer/importer.go?s=2726:2755#L72)
<pre>func Default() <a href="/pkg/go/types/">types</a>.<a href="/pkg/go/types/#Importer">Importer</a></pre>
Default returns an Importer for the compiler that built the running binary.
If available, the result implements types.ImporterFrom.



## <a id="For">func</a> [For](https://golang.org/src/go/importer/importer.go?s=2469:2524#L66)
<pre>func For(compiler <a href="/pkg/builtin/#string">string</a>, lookup <a href="#Lookup">Lookup</a>) <a href="/pkg/go/types/">types</a>.<a href="/pkg/go/types/#Importer">Importer</a></pre>
For calls ForCompiler with a new FileSet.

Deprecated: Use ForCompiler, which populates a FileSet
with the positions of objects created by the importer.



## <a id="ForCompiler">func</a> [ForCompiler](https://golang.org/src/go/importer/importer.go?s=1551:1635#L30)
<pre>func ForCompiler(fset *<a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#FileSet">FileSet</a>, compiler <a href="/pkg/builtin/#string">string</a>, lookup <a href="#Lookup">Lookup</a>) <a href="/pkg/go/types/">types</a>.<a href="/pkg/go/types/#Importer">Importer</a></pre>
ForCompiler returns an Importer for importing from installed packages
for the compilers "gc" and "gccgo", or for importing directly
from the source if the compiler argument is "source". In this
latter case, importing may fail under circumstances where the
exported API is not entirely defined in pure Go source code
(if the package API depends on cgo-defined entities, the type
checker won't have access to those).

The lookup function is called each time the resulting importer needs
to resolve an import path. In this mode the importer can only be
invoked with canonical import paths (not relative or absolute ones);
it is assumed that the translation to canonical import paths is being
done by the client of the importer.

A lookup function must be provided for correct module-aware operation.
Deprecated: If lookup is nil, for backwards-compatibility, the importer
will attempt to resolve imports in the $GOPATH workspace.





## <a id="Lookup">type</a> [Lookup](https://golang.org/src/go/importer/importer.go?s=521:573#L11)
A Lookup function returns a reader to access package data for
a given import path, or an error if no matching package is found.


<pre>type Lookup func(path <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadCloser">ReadCloser</a>, <a href="/pkg/builtin/#error">error</a>)</pre>














