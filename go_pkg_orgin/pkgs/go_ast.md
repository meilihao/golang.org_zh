

# ast
`import "go/ast"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package ast declares the types used to represent syntax trees for Go
packages.




## <a id="pkg-index">Index</a>
* [func FileExports(src *File) bool](#FileExports)
* [func FilterDecl(decl Decl, f Filter) bool](#FilterDecl)
* [func FilterFile(src *File, f Filter) bool](#FilterFile)
* [func FilterPackage(pkg *Package, f Filter) bool](#FilterPackage)
* [func Fprint(w io.Writer, fset *token.FileSet, x interface{}, f FieldFilter) error](#Fprint)
* [func Inspect(node Node, f func(Node) bool)](#Inspect)
* [func IsExported(name string) bool](#IsExported)
* [func NotNilFilter(_ string, v reflect.Value) bool](#NotNilFilter)
* [func PackageExports(pkg *Package) bool](#PackageExports)
* [func Print(fset *token.FileSet, x interface{}) error](#Print)
* [func SortImports(fset *token.FileSet, f *File)](#SortImports)
* [func Walk(v Visitor, node Node)](#Walk)
* [type ArrayType](#ArrayType)
  * [func (x *ArrayType) End() token.Pos](#ArrayType.End)
  * [func (x *ArrayType) Pos() token.Pos](#ArrayType.Pos)
* [type AssignStmt](#AssignStmt)
  * [func (s *AssignStmt) End() token.Pos](#AssignStmt.End)
  * [func (s *AssignStmt) Pos() token.Pos](#AssignStmt.Pos)
* [type BadDecl](#BadDecl)
  * [func (d *BadDecl) End() token.Pos](#BadDecl.End)
  * [func (d *BadDecl) Pos() token.Pos](#BadDecl.Pos)
* [type BadExpr](#BadExpr)
  * [func (x *BadExpr) End() token.Pos](#BadExpr.End)
  * [func (x *BadExpr) Pos() token.Pos](#BadExpr.Pos)
* [type BadStmt](#BadStmt)
  * [func (s *BadStmt) End() token.Pos](#BadStmt.End)
  * [func (s *BadStmt) Pos() token.Pos](#BadStmt.Pos)
* [type BasicLit](#BasicLit)
  * [func (x *BasicLit) End() token.Pos](#BasicLit.End)
  * [func (x *BasicLit) Pos() token.Pos](#BasicLit.Pos)
* [type BinaryExpr](#BinaryExpr)
  * [func (x *BinaryExpr) End() token.Pos](#BinaryExpr.End)
  * [func (x *BinaryExpr) Pos() token.Pos](#BinaryExpr.Pos)
* [type BlockStmt](#BlockStmt)
  * [func (s *BlockStmt) End() token.Pos](#BlockStmt.End)
  * [func (s *BlockStmt) Pos() token.Pos](#BlockStmt.Pos)
* [type BranchStmt](#BranchStmt)
  * [func (s *BranchStmt) End() token.Pos](#BranchStmt.End)
  * [func (s *BranchStmt) Pos() token.Pos](#BranchStmt.Pos)
* [type CallExpr](#CallExpr)
  * [func (x *CallExpr) End() token.Pos](#CallExpr.End)
  * [func (x *CallExpr) Pos() token.Pos](#CallExpr.Pos)
* [type CaseClause](#CaseClause)
  * [func (s *CaseClause) End() token.Pos](#CaseClause.End)
  * [func (s *CaseClause) Pos() token.Pos](#CaseClause.Pos)
* [type ChanDir](#ChanDir)
* [type ChanType](#ChanType)
  * [func (x *ChanType) End() token.Pos](#ChanType.End)
  * [func (x *ChanType) Pos() token.Pos](#ChanType.Pos)
* [type CommClause](#CommClause)
  * [func (s *CommClause) End() token.Pos](#CommClause.End)
  * [func (s *CommClause) Pos() token.Pos](#CommClause.Pos)
* [type Comment](#Comment)
  * [func (c *Comment) End() token.Pos](#Comment.End)
  * [func (c *Comment) Pos() token.Pos](#Comment.Pos)
* [type CommentGroup](#CommentGroup)
  * [func (g *CommentGroup) End() token.Pos](#CommentGroup.End)
  * [func (g *CommentGroup) Pos() token.Pos](#CommentGroup.Pos)
  * [func (g *CommentGroup) Text() string](#CommentGroup.Text)
* [type CommentMap](#CommentMap)
  * [func NewCommentMap(fset *token.FileSet, node Node, comments []*CommentGroup) CommentMap](#NewCommentMap)
  * [func (cmap CommentMap) Comments() []*CommentGroup](#CommentMap.Comments)
  * [func (cmap CommentMap) Filter(node Node) CommentMap](#CommentMap.Filter)
  * [func (cmap CommentMap) String() string](#CommentMap.String)
  * [func (cmap CommentMap) Update(old, new Node) Node](#CommentMap.Update)
* [type CompositeLit](#CompositeLit)
  * [func (x *CompositeLit) End() token.Pos](#CompositeLit.End)
  * [func (x *CompositeLit) Pos() token.Pos](#CompositeLit.Pos)
* [type Decl](#Decl)
* [type DeclStmt](#DeclStmt)
  * [func (s *DeclStmt) End() token.Pos](#DeclStmt.End)
  * [func (s *DeclStmt) Pos() token.Pos](#DeclStmt.Pos)
* [type DeferStmt](#DeferStmt)
  * [func (s *DeferStmt) End() token.Pos](#DeferStmt.End)
  * [func (s *DeferStmt) Pos() token.Pos](#DeferStmt.Pos)
* [type Ellipsis](#Ellipsis)
  * [func (x *Ellipsis) End() token.Pos](#Ellipsis.End)
  * [func (x *Ellipsis) Pos() token.Pos](#Ellipsis.Pos)
* [type EmptyStmt](#EmptyStmt)
  * [func (s *EmptyStmt) End() token.Pos](#EmptyStmt.End)
  * [func (s *EmptyStmt) Pos() token.Pos](#EmptyStmt.Pos)
* [type Expr](#Expr)
* [type ExprStmt](#ExprStmt)
  * [func (s *ExprStmt) End() token.Pos](#ExprStmt.End)
  * [func (s *ExprStmt) Pos() token.Pos](#ExprStmt.Pos)
* [type Field](#Field)
  * [func (f *Field) End() token.Pos](#Field.End)
  * [func (f *Field) Pos() token.Pos](#Field.Pos)
* [type FieldFilter](#FieldFilter)
* [type FieldList](#FieldList)
  * [func (f *FieldList) End() token.Pos](#FieldList.End)
  * [func (f *FieldList) NumFields() int](#FieldList.NumFields)
  * [func (f *FieldList) Pos() token.Pos](#FieldList.Pos)
* [type File](#File)
  * [func MergePackageFiles(pkg *Package, mode MergeMode) *File](#MergePackageFiles)
  * [func (f *File) End() token.Pos](#File.End)
  * [func (f *File) Pos() token.Pos](#File.Pos)
* [type Filter](#Filter)
* [type ForStmt](#ForStmt)
  * [func (s *ForStmt) End() token.Pos](#ForStmt.End)
  * [func (s *ForStmt) Pos() token.Pos](#ForStmt.Pos)
* [type FuncDecl](#FuncDecl)
  * [func (d *FuncDecl) End() token.Pos](#FuncDecl.End)
  * [func (d *FuncDecl) Pos() token.Pos](#FuncDecl.Pos)
* [type FuncLit](#FuncLit)
  * [func (x *FuncLit) End() token.Pos](#FuncLit.End)
  * [func (x *FuncLit) Pos() token.Pos](#FuncLit.Pos)
* [type FuncType](#FuncType)
  * [func (x *FuncType) End() token.Pos](#FuncType.End)
  * [func (x *FuncType) Pos() token.Pos](#FuncType.Pos)
* [type GenDecl](#GenDecl)
  * [func (d *GenDecl) End() token.Pos](#GenDecl.End)
  * [func (d *GenDecl) Pos() token.Pos](#GenDecl.Pos)
* [type GoStmt](#GoStmt)
  * [func (s *GoStmt) End() token.Pos](#GoStmt.End)
  * [func (s *GoStmt) Pos() token.Pos](#GoStmt.Pos)
* [type Ident](#Ident)
  * [func NewIdent(name string) *Ident](#NewIdent)
  * [func (x *Ident) End() token.Pos](#Ident.End)
  * [func (id *Ident) IsExported() bool](#Ident.IsExported)
  * [func (x *Ident) Pos() token.Pos](#Ident.Pos)
  * [func (id *Ident) String() string](#Ident.String)
* [type IfStmt](#IfStmt)
  * [func (s *IfStmt) End() token.Pos](#IfStmt.End)
  * [func (s *IfStmt) Pos() token.Pos](#IfStmt.Pos)
* [type ImportSpec](#ImportSpec)
  * [func (s *ImportSpec) End() token.Pos](#ImportSpec.End)
  * [func (s *ImportSpec) Pos() token.Pos](#ImportSpec.Pos)
* [type Importer](#Importer)
* [type IncDecStmt](#IncDecStmt)
  * [func (s *IncDecStmt) End() token.Pos](#IncDecStmt.End)
  * [func (s *IncDecStmt) Pos() token.Pos](#IncDecStmt.Pos)
* [type IndexExpr](#IndexExpr)
  * [func (x *IndexExpr) End() token.Pos](#IndexExpr.End)
  * [func (x *IndexExpr) Pos() token.Pos](#IndexExpr.Pos)
* [type InterfaceType](#InterfaceType)
  * [func (x *InterfaceType) End() token.Pos](#InterfaceType.End)
  * [func (x *InterfaceType) Pos() token.Pos](#InterfaceType.Pos)
* [type KeyValueExpr](#KeyValueExpr)
  * [func (x *KeyValueExpr) End() token.Pos](#KeyValueExpr.End)
  * [func (x *KeyValueExpr) Pos() token.Pos](#KeyValueExpr.Pos)
* [type LabeledStmt](#LabeledStmt)
  * [func (s *LabeledStmt) End() token.Pos](#LabeledStmt.End)
  * [func (s *LabeledStmt) Pos() token.Pos](#LabeledStmt.Pos)
* [type MapType](#MapType)
  * [func (x *MapType) End() token.Pos](#MapType.End)
  * [func (x *MapType) Pos() token.Pos](#MapType.Pos)
* [type MergeMode](#MergeMode)
* [type Node](#Node)
* [type ObjKind](#ObjKind)
  * [func (kind ObjKind) String() string](#ObjKind.String)
* [type Object](#Object)
  * [func NewObj(kind ObjKind, name string) *Object](#NewObj)
  * [func (obj *Object) Pos() token.Pos](#Object.Pos)
* [type Package](#Package)
  * [func NewPackage(fset *token.FileSet, files map[string]*File, importer Importer, universe *Scope) (*Package, error)](#NewPackage)
  * [func (p *Package) End() token.Pos](#Package.End)
  * [func (p *Package) Pos() token.Pos](#Package.Pos)
* [type ParenExpr](#ParenExpr)
  * [func (x *ParenExpr) End() token.Pos](#ParenExpr.End)
  * [func (x *ParenExpr) Pos() token.Pos](#ParenExpr.Pos)
* [type RangeStmt](#RangeStmt)
  * [func (s *RangeStmt) End() token.Pos](#RangeStmt.End)
  * [func (s *RangeStmt) Pos() token.Pos](#RangeStmt.Pos)
* [type ReturnStmt](#ReturnStmt)
  * [func (s *ReturnStmt) End() token.Pos](#ReturnStmt.End)
  * [func (s *ReturnStmt) Pos() token.Pos](#ReturnStmt.Pos)
* [type Scope](#Scope)
  * [func NewScope(outer *Scope) *Scope](#NewScope)
  * [func (s *Scope) Insert(obj *Object) (alt *Object)](#Scope.Insert)
  * [func (s *Scope) Lookup(name string) *Object](#Scope.Lookup)
  * [func (s *Scope) String() string](#Scope.String)
* [type SelectStmt](#SelectStmt)
  * [func (s *SelectStmt) End() token.Pos](#SelectStmt.End)
  * [func (s *SelectStmt) Pos() token.Pos](#SelectStmt.Pos)
* [type SelectorExpr](#SelectorExpr)
  * [func (x *SelectorExpr) End() token.Pos](#SelectorExpr.End)
  * [func (x *SelectorExpr) Pos() token.Pos](#SelectorExpr.Pos)
* [type SendStmt](#SendStmt)
  * [func (s *SendStmt) End() token.Pos](#SendStmt.End)
  * [func (s *SendStmt) Pos() token.Pos](#SendStmt.Pos)
* [type SliceExpr](#SliceExpr)
  * [func (x *SliceExpr) End() token.Pos](#SliceExpr.End)
  * [func (x *SliceExpr) Pos() token.Pos](#SliceExpr.Pos)
* [type Spec](#Spec)
* [type StarExpr](#StarExpr)
  * [func (x *StarExpr) End() token.Pos](#StarExpr.End)
  * [func (x *StarExpr) Pos() token.Pos](#StarExpr.Pos)
* [type Stmt](#Stmt)
* [type StructType](#StructType)
  * [func (x *StructType) End() token.Pos](#StructType.End)
  * [func (x *StructType) Pos() token.Pos](#StructType.Pos)
* [type SwitchStmt](#SwitchStmt)
  * [func (s *SwitchStmt) End() token.Pos](#SwitchStmt.End)
  * [func (s *SwitchStmt) Pos() token.Pos](#SwitchStmt.Pos)
* [type TypeAssertExpr](#TypeAssertExpr)
  * [func (x *TypeAssertExpr) End() token.Pos](#TypeAssertExpr.End)
  * [func (x *TypeAssertExpr) Pos() token.Pos](#TypeAssertExpr.Pos)
* [type TypeSpec](#TypeSpec)
  * [func (s *TypeSpec) End() token.Pos](#TypeSpec.End)
  * [func (s *TypeSpec) Pos() token.Pos](#TypeSpec.Pos)
* [type TypeSwitchStmt](#TypeSwitchStmt)
  * [func (s *TypeSwitchStmt) End() token.Pos](#TypeSwitchStmt.End)
  * [func (s *TypeSwitchStmt) Pos() token.Pos](#TypeSwitchStmt.Pos)
* [type UnaryExpr](#UnaryExpr)
  * [func (x *UnaryExpr) End() token.Pos](#UnaryExpr.End)
  * [func (x *UnaryExpr) Pos() token.Pos](#UnaryExpr.Pos)
* [type ValueSpec](#ValueSpec)
  * [func (s *ValueSpec) End() token.Pos](#ValueSpec.End)
  * [func (s *ValueSpec) Pos() token.Pos](#ValueSpec.Pos)
* [type Visitor](#Visitor)


#### <a id="pkg-examples">Examples</a>
* [CommentMap](#example_CommentMap)
* [Inspect](#example_Inspect)
* [Print](#example_Print)


#### <a id="pkg-files">Package files</a>
[ast.go](https://golang.org/src/go/ast/ast.go) [commentmap.go](https://golang.org/src/go/ast/commentmap.go) [filter.go](https://golang.org/src/go/ast/filter.go) [import.go](https://golang.org/src/go/ast/import.go) [print.go](https://golang.org/src/go/ast/print.go) [resolve.go](https://golang.org/src/go/ast/resolve.go) [scope.go](https://golang.org/src/go/ast/scope.go) [walk.go](https://golang.org/src/go/ast/walk.go) 






## <a id="FileExports">func</a> [FileExports](https://golang.org/src/go/ast/filter.go?s=869:901#L18)
<pre>func FileExports(src *<a href="#File">File</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
FileExports trims the AST for a Go source file in place such that
only exported nodes remain: all top-level identifiers which are not exported
and their associated information (such as type, initial value, or function
body) are removed. Non-exported fields and methods of exported types are
stripped. The File.Comments list is not changed.

FileExports reports whether there are exported declarations.



## <a id="FilterDecl">func</a> [FilterDecl](https://golang.org/src/go/ast/filter.go?s=5502:5543#L223)
<pre>func FilterDecl(decl <a href="#Decl">Decl</a>, f <a href="#Filter">Filter</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
FilterDecl trims the AST for a Go declaration in place by removing
all names (including struct field and interface method names, but
not from parameter lists) that don't pass through the filter f.

FilterDecl reports whether there are any declared names left after
filtering.



## <a id="FilterFile">func</a> [FilterFile](https://golang.org/src/go/ast/filter.go?s=6312:6353#L248)
<pre>func FilterFile(src *<a href="#File">File</a>, f <a href="#Filter">Filter</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
FilterFile trims the AST for a Go file in place by removing all
names from top-level declarations (including struct field and
interface method names, but not from parameter lists) that don't
pass through the filter f. If the declaration is empty afterwards,
the declaration is removed from the AST. Import declarations are
always removed. The File.Comments list is not changed.

FilterFile reports whether there are any top-level declarations
left after filtering.



## <a id="FilterPackage">func</a> [FilterPackage](https://golang.org/src/go/ast/filter.go?s=7128:7175#L275)
<pre>func FilterPackage(pkg *<a href="#Package">Package</a>, f <a href="#Filter">Filter</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
FilterPackage trims the AST for a Go package in place by removing
all names from top-level declarations (including struct field and
interface method names, but not from parameter lists) that don't
pass through the filter f. If the declaration is empty afterwards,
the declaration is removed from the AST. The pkg.Files list is not
changed, so that file names and top-level package comments don't get
lost.

FilterPackage reports whether there are any top-level declarations
left after filtering.



## <a id="Fprint">func</a> [Fprint](https://golang.org/src/go/ast/print.go?s=1165:1246#L29)
<pre>func Fprint(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, fset *<a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#FileSet">FileSet</a>, x interface{}, f <a href="#FieldFilter">FieldFilter</a>) <a href="/pkg/builtin/#error">error</a></pre>
Fprint prints the (sub-)tree starting at AST node x to w.
If fset != nil, position information is interpreted relative
to that file set. Otherwise positions are printed as integer
values (file set specific offsets).

A non-nil FieldFilter f may be provided to control the output:
struct fields for which f(fieldname, fieldvalue) is true are
printed; all others are filtered from the output. Unexported
struct fields are never printed.



## <a id="Inspect">func</a> [Inspect](https://golang.org/src/go/ast/walk.go?s=6291:6333#L374)
<pre>func Inspect(node <a href="#Node">Node</a>, f func(<a href="#Node">Node</a>) <a href="/pkg/builtin/#bool">bool</a>)</pre>
Inspect traverses an AST in depth-first order: It starts by calling
f(node); node must not be nil. If f returns true, Inspect invokes f
recursively for each of the non-nil children of node, followed by a
call of f(nil).


<a id="example_Inspect">Example</a>
```go
```

output:
```txt
```
<p>This example demonstrates how to inspect the AST of a Go program.
</p>
## <a id="IsExported">func</a> [IsExported](https://golang.org/src/go/ast/ast.go?s=16554:16587#L516)
<pre>func IsExported(name <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsExported reports whether name starts with an upper-case letter.



## <a id="NotNilFilter">func</a> [NotNilFilter](https://golang.org/src/go/ast/print.go?s=500:549#L12)
<pre>func NotNilFilter(_ <a href="/pkg/builtin/#string">string</a>, v <a href="/pkg/reflect/">reflect</a>.<a href="/pkg/reflect/#Value">Value</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
NotNilFilter returns true for field values that are not nil;
it returns false otherwise.



## <a id="PackageExports">func</a> [PackageExports](https://golang.org/src/go/ast/filter.go?s=1258:1296#L29)
<pre>func PackageExports(pkg *<a href="#Package">Package</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
PackageExports trims the AST for a Go package in place such that
only exported nodes remain. The pkg.Files list is not changed, so that
file names and top-level package comments don't get lost.

PackageExports reports whether there are exported declarations;
it returns false otherwise.



## <a id="Print">func</a> [Print](https://golang.org/src/go/ast/print.go?s=1951:2003#L63)
<pre>func Print(fset *<a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#FileSet">FileSet</a>, x interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Print prints x to standard output, skipping nil fields.
Print(fset, x) is the same as Fprint(os.Stdout, fset, x, NotNilFilter).


<a id="example_Print">Example</a>
```go
```

output:
```txt
```
<p>This example shows what an AST looks like when printed for debugging.
</p>
## <a id="SortImports">func</a> [SortImports](https://golang.org/src/go/ast/import.go?s=378:424#L5)
<pre>func SortImports(fset *<a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#FileSet">FileSet</a>, f *<a href="#File">File</a>)</pre>
SortImports sorts runs of consecutive import lines in import blocks in f.
It also removes duplicate imports when it is possible to do so without data loss.



## <a id="Walk">func</a> [Walk](https://golang.org/src/go/ast/walk.go?s=1311:1342#L41)
<pre>func Walk(v <a href="#Visitor">Visitor</a>, node <a href="#Node">Node</a>)</pre>
Walk traverses an AST in depth-first order: It starts by calling
v.Visit(node); node must not be nil. If the visitor w returned by
v.Visit(node) is not nil, Walk is invoked recursively with visitor
w for each of the non-nil children of node, followed by a call of
w.Visit(nil).





## <a id="ArrayType">type</a> [ArrayType](https://golang.org/src/go/ast/ast.go?s=10688:10862#L366)
An ArrayType node represents an array or slice type.


<pre>type ArrayType struct {
<span id="ArrayType.Lbrack"></span>    Lbrack <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;[&#34;</span>
<span id="ArrayType.Len"></span>    Len    <a href="#Expr">Expr</a>      <span class="comment">// Ellipsis node for [...]T array types, nil for slice types</span>
<span id="ArrayType.Elt"></span>    Elt    <a href="#Expr">Expr</a>      <span class="comment">// element type</span>
}
</pre>











### <a id="ArrayType.End">func</a> (\*ArrayType) [End](https://golang.org/src/go/ast/ast.go?s=14832:14867#L467)
<pre>func (x *<a href="#ArrayType">ArrayType</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="ArrayType.Pos">func</a> (\*ArrayType) [Pos](https://golang.org/src/go/ast/ast.go?s=13214:13249#L434)
<pre>func (x *<a href="#ArrayType">ArrayType</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="AssignStmt">type</a> [AssignStmt](https://golang.org/src/go/ast/ast.go?s=18447:18590#L589)
An AssignStmt node represents an assignment or
a short variable declaration.


<pre>type AssignStmt struct {
<span id="AssignStmt.Lhs"></span>    Lhs    []<a href="#Expr">Expr</a>
<span id="AssignStmt.TokPos"></span>    TokPos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>   <span class="comment">// position of Tok</span>
<span id="AssignStmt.Tok"></span>    Tok    <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Token">Token</a> <span class="comment">// assignment token, DEFINE</span>
<span id="AssignStmt.Rhs"></span>    Rhs    []<a href="#Expr">Expr</a>
}
</pre>











### <a id="AssignStmt.End">func</a> (\*AssignStmt) [End](https://golang.org/src/go/ast/ast.go?s=23604:23640#L735)
<pre>func (s *<a href="#AssignStmt">AssignStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="AssignStmt.Pos">func</a> (\*AssignStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=22270:22306#L706)
<pre>func (s *<a href="#AssignStmt">AssignStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="BadDecl">type</a> [BadDecl](https://golang.org/src/go/ast/ast.go?s=28196:28273#L886)
A BadDecl node is a placeholder for declarations containing
syntax errors for which no correct declaration nodes can be
created.


<pre>type BadDecl struct {
<span id="BadDecl.From"></span>    From, To <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position range of bad declaration</span>
}
</pre>











### <a id="BadDecl.End">func</a> (\*BadDecl) [End](https://golang.org/src/go/ast/ast.go?s=29587:29620#L926)
<pre>func (d *<a href="#BadDecl">BadDecl</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="BadDecl.Pos">func</a> (\*BadDecl) [Pos](https://golang.org/src/go/ast/ast.go?s=29419:29452#L922)
<pre>func (d *<a href="#BadDecl">BadDecl</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="BadExpr">type</a> [BadExpr](https://golang.org/src/go/ast/ast.go?s=6425:6501#L223)
A BadExpr node is a placeholder for expressions containing
syntax errors for which no correct expression nodes can be
created.


<pre>type BadExpr struct {
<span id="BadExpr.From"></span>    From, To <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position range of bad expression</span>
}
</pre>











### <a id="BadExpr.End">func</a> (\*BadExpr) [End](https://golang.org/src/go/ast/ast.go?s=13713:13746#L446)
<pre>func (x *<a href="#BadExpr">BadExpr</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="BadExpr.Pos">func</a> (\*BadExpr) [Pos](https://golang.org/src/go/ast/ast.go?s=12209:12242#L413)
<pre>func (x *<a href="#BadExpr">BadExpr</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="BadStmt">type</a> [BadStmt](https://golang.org/src/go/ast/ast.go?s=17215:17290#L540)
A BadStmt node is a placeholder for statements containing
syntax errors for which no correct statement nodes can be
created.


<pre>type BadStmt struct {
<span id="BadStmt.From"></span>    From, To <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position range of bad statement</span>
}
</pre>











### <a id="BadStmt.End">func</a> (\*BadStmt) [End](https://golang.org/src/go/ast/ast.go?s=23112:23145#L721)
<pre>func (s *<a href="#BadStmt">BadStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="BadStmt.Pos">func</a> (\*BadStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=21827:21860#L699)
<pre>func (s *<a href="#BadStmt">BadStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="BasicLit">type</a> [BasicLit](https://golang.org/src/go/ast/ast.go?s=7007:7266#L243)
A BasicLit node represents a literal of basic type.


<pre>type BasicLit struct {
<span id="BasicLit.ValuePos"></span>    ValuePos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>   <span class="comment">// literal position</span>
<span id="BasicLit.Kind"></span>    Kind     <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Token">Token</a> <span class="comment">// token.INT, token.FLOAT, token.IMAG, token.CHAR, or token.STRING</span>
<span id="BasicLit.Value"></span>    Value    <a href="/pkg/builtin/#string">string</a>      <span class="comment">// literal string; e.g. 42, 0x7f, 3.14, 1e-9, 2.4i, &#39;a&#39;, &#39;\x7f&#39;, &#34;foo&#34; or `\m\n\o`</span>
}
</pre>











### <a id="BasicLit.End">func</a> (\*BasicLit) [End](https://golang.org/src/go/ast/ast.go?s=13967:14001#L454)
<pre>func (x *<a href="#BasicLit">BasicLit</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="BasicLit.Pos">func</a> (\*BasicLit) [Pos](https://golang.org/src/go/ast/ast.go?s=12375:12409#L416)
<pre>func (x *<a href="#BasicLit">BasicLit</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="BinaryExpr">type</a> [BinaryExpr](https://golang.org/src/go/ast/ast.go?s=9980:10145#L333)
A BinaryExpr node represents a binary expression.


<pre>type BinaryExpr struct {
<span id="BinaryExpr.X"></span>    X     <a href="#Expr">Expr</a>        <span class="comment">// left operand</span>
<span id="BinaryExpr.OpPos"></span>    OpPos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>   <span class="comment">// position of Op</span>
<span id="BinaryExpr.Op"></span>    Op    <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Token">Token</a> <span class="comment">// operator</span>
<span id="BinaryExpr.Y"></span>    Y     <a href="#Expr">Expr</a>        <span class="comment">// right operand</span>
}
</pre>











### <a id="BinaryExpr.End">func</a> (\*BinaryExpr) [End](https://golang.org/src/go/ast/ast.go?s=14704:14740#L465)
<pre>func (x *<a href="#BinaryExpr">BinaryExpr</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="BinaryExpr.Pos">func</a> (\*BinaryExpr) [Pos](https://golang.org/src/go/ast/ast.go?s=13088:13124#L432)
<pre>func (x *<a href="#BinaryExpr">BinaryExpr</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="BlockStmt">type</a> [BlockStmt](https://golang.org/src/go/ast/ast.go?s=19375:19488#L624)
A BlockStmt node represents a braced statement list.


<pre>type BlockStmt struct {
<span id="BlockStmt.Lbrace"></span>    Lbrace <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;{&#34;</span>
<span id="BlockStmt.List"></span>    List   []<a href="#Stmt">Stmt</a>
<span id="BlockStmt.Rbrace"></span>    Rbrace <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;}&#34;</span>
}
</pre>











### <a id="BlockStmt.End">func</a> (\*BlockStmt) [End](https://golang.org/src/go/ast/ast.go?s=24088:24123#L750)
<pre>func (s *<a href="#BlockStmt">BlockStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="BlockStmt.Pos">func</a> (\*BlockStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=22576:22611#L711)
<pre>func (s *<a href="#BlockStmt">BlockStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="BranchStmt">type</a> [BranchStmt](https://golang.org/src/go/ast/ast.go?s=19135:19315#L617)
A BranchStmt node represents a break, continue, goto,
or fallthrough statement.


<pre>type BranchStmt struct {
<span id="BranchStmt.TokPos"></span>    TokPos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>   <span class="comment">// position of Tok</span>
<span id="BranchStmt.Tok"></span>    Tok    <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Token">Token</a> <span class="comment">// keyword token (BREAK, CONTINUE, GOTO, FALLTHROUGH)</span>
<span id="BranchStmt.Label"></span>    Label  *<a href="#Ident">Ident</a>      <span class="comment">// label name; or nil</span>
}
</pre>











### <a id="BranchStmt.End">func</a> (\*BranchStmt) [End](https://golang.org/src/go/ast/ast.go?s=23945:23981#L744)
<pre>func (s *<a href="#BranchStmt">BranchStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="BranchStmt.Pos">func</a> (\*BranchStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=22515:22551#L710)
<pre>func (s *<a href="#BranchStmt">BranchStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="CallExpr">type</a> [CallExpr](https://golang.org/src/go/ast/ast.go?s=9172:9444#L307)
A CallExpr node represents an expression followed by an argument list.


<pre>type CallExpr struct {
<span id="CallExpr.Fun"></span>    Fun      <a href="#Expr">Expr</a>      <span class="comment">// function expression</span>
<span id="CallExpr.Lparen"></span>    Lparen   <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;(&#34;</span>
<span id="CallExpr.Args"></span>    Args     []<a href="#Expr">Expr</a>    <span class="comment">// function arguments; or nil</span>
<span id="CallExpr.Ellipsis"></span>    Ellipsis <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;...&#34; (token.NoPos if there is no &#34;...&#34;)</span>
<span id="CallExpr.Rparen"></span>    Rparen   <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;)&#34;</span>
}
</pre>











### <a id="CallExpr.End">func</a> (\*CallExpr) [End](https://golang.org/src/go/ast/ast.go?s=14515:14549#L462)
<pre>func (x *<a href="#CallExpr">CallExpr</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="CallExpr.Pos">func</a> (\*CallExpr) [Pos](https://golang.org/src/go/ast/ast.go?s=12905:12939#L429)
<pre>func (x *<a href="#CallExpr">CallExpr</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="CaseClause">type</a> [CaseClause](https://golang.org/src/go/ast/ast.go?s=19818:20056#L640)
A CaseClause represents a case of an expression or type switch statement.


<pre>type CaseClause struct {
<span id="CaseClause.Case"></span>    Case  <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;case&#34; or &#34;default&#34; keyword</span>
<span id="CaseClause.List"></span>    List  []<a href="#Expr">Expr</a>    <span class="comment">// list of expressions or types; nil means default case</span>
<span id="CaseClause.Colon"></span>    Colon <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;:&#34;</span>
<span id="CaseClause.Body"></span>    Body  []<a href="#Stmt">Stmt</a>    <span class="comment">// statement list; or nil</span>
}
</pre>











### <a id="CaseClause.End">func</a> (\*CaseClause) [End](https://golang.org/src/go/ast/ast.go?s=24251:24287#L757)
<pre>func (s *<a href="#CaseClause">CaseClause</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="CaseClause.Pos">func</a> (\*CaseClause) [Pos](https://golang.org/src/go/ast/ast.go?s=22694:22730#L713)
<pre>func (s *<a href="#CaseClause">CaseClause</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="ChanDir">type</a> [ChanDir](https://golang.org/src/go/ast/ast.go?s=10443:10459#L353)
The direction of a channel type is indicated by a bit
mask including one or both of the following constants.


<pre>type ChanDir <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="SEND">SEND</span> <a href="#ChanDir">ChanDir</a> = 1 &lt;&lt; <a href="/pkg/builtin/#iota">iota</a>
    <span id="RECV">RECV</span>
)</pre>









## <a id="ChanType">type</a> [ChanType](https://golang.org/src/go/ast/ast.go?s=11903:12146#L403)
A ChanType node represents a channel type.


<pre>type ChanType struct {
<span id="ChanType.Begin"></span>    Begin <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;chan&#34; keyword or &#34;&lt;-&#34; (whichever comes first)</span>
<span id="ChanType.Arrow"></span>    Arrow <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;&lt;-&#34; (token.NoPos if there is no &#34;&lt;-&#34;)</span>
<span id="ChanType.Dir"></span>    Dir   <a href="#ChanDir">ChanDir</a>   <span class="comment">// channel direction</span>
<span id="ChanType.Value"></span>    Value <a href="#Expr">Expr</a>      <span class="comment">// value type</span>
}
</pre>











### <a id="ChanType.End">func</a> (\*ChanType) [End](https://golang.org/src/go/ast/ast.go?s=15208:15242#L477)
<pre>func (x *<a href="#ChanType">ChanType</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="ChanType.Pos">func</a> (\*ChanType) [Pos](https://golang.org/src/go/ast/ast.go?s=13653:13687#L444)
<pre>func (x *<a href="#ChanType">ChanType</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="CommClause">type</a> [CommClause](https://golang.org/src/go/ast/ast.go?s=20695:20930#L664)
A CommClause node represents a case of a select statement.


<pre>type CommClause struct {
<span id="CommClause.Case"></span>    Case  <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;case&#34; or &#34;default&#34; keyword</span>
<span id="CommClause.Comm"></span>    Comm  <a href="#Stmt">Stmt</a>      <span class="comment">// send or receive statement; nil means default case</span>
<span id="CommClause.Colon"></span>    Colon <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;:&#34;</span>
<span id="CommClause.Body"></span>    Body  []<a href="#Stmt">Stmt</a>    <span class="comment">// statement list; or nil</span>
}
</pre>











### <a id="CommClause.End">func</a> (\*CommClause) [End](https://golang.org/src/go/ast/ast.go?s=24502:24538#L765)
<pre>func (s *<a href="#CommClause">CommClause</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="CommClause.Pos">func</a> (\*CommClause) [Pos](https://golang.org/src/go/ast/ast.go?s=22875:22911#L716)
<pre>func (s *<a href="#CommClause">CommClause</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="Comment">type</a> [Comment](https://golang.org/src/go/ast/ast.go?s=1803:1955#L50)
A Comment node represents a single //-style or /*-style comment.


<pre>type Comment struct {
<span id="Comment.Slash"></span>    Slash <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;/&#34; starting the comment</span>
<span id="Comment.Text"></span>    Text  <a href="/pkg/builtin/#string">string</a>    <span class="comment">// comment text (excluding &#39;\n&#39; for //-style comments)</span>
}
</pre>











### <a id="Comment.End">func</a> (\*Comment) [End](https://golang.org/src/go/ast/ast.go?s=2010:2043#L56)
<pre>func (c *<a href="#Comment">Comment</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="Comment.Pos">func</a> (\*Comment) [Pos](https://golang.org/src/go/ast/ast.go?s=1957:1990#L55)
<pre>func (c *<a href="#Comment">Comment</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="CommentGroup">type</a> [CommentGroup](https://golang.org/src/go/ast/ast.go?s=2201:2263#L61)
A CommentGroup represents a sequence of comments
with no other tokens and no empty lines between.


<pre>type CommentGroup struct {
<span id="CommentGroup.List"></span>    List []*<a href="#Comment">Comment</a> <span class="comment">// len(List) &gt; 0</span>
}
</pre>











### <a id="CommentGroup.End">func</a> (\*CommentGroup) [End](https://golang.org/src/go/ast/ast.go?s=2331:2369#L66)
<pre>func (g *<a href="#CommentGroup">CommentGroup</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="CommentGroup.Pos">func</a> (\*CommentGroup) [Pos](https://golang.org/src/go/ast/ast.go?s=2265:2303#L65)
<pre>func (g *<a href="#CommentGroup">CommentGroup</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="CommentGroup.Text">func</a> (\*CommentGroup) [Text](https://golang.org/src/go/ast/ast.go?s=2941:2977#L84)
<pre>func (g *<a href="#CommentGroup">CommentGroup</a>) Text() <a href="/pkg/builtin/#string">string</a></pre>
Text returns the text of the comment.
Comment markers (//, /*, and */), the first space of a line comment, and
leading and trailing empty lines are removed. Multiple empty lines are
reduced to one, and trailing space on lines is trimmed. Unless the result
is empty, it is newline-terminated.




## <a id="CommentMap">type</a> [CommentMap](https://golang.org/src/go/ast/commentmap.go?s=932:972#L25)
A CommentMap maps an AST node to a list of comment groups
associated with it. See NewCommentMap for a description of
the association.


<pre>type CommentMap map[<a href="#Node">Node</a>][]*<a href="#CommentGroup">CommentGroup</a></pre>





<a id="example_CommentMap">Example</a>
```go
```

output:
```txt
```
<p>This example illustrates how to remove a variable declaration
in a Go program while maintaining correct comment association
using an ast.CommentMap.
</p>



### <a id="NewCommentMap">func</a> [NewCommentMap](https://golang.org/src/go/ast/commentmap.go?s=3964:4051#L133)
<pre>func NewCommentMap(fset *<a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#FileSet">FileSet</a>, node <a href="#Node">Node</a>, comments []*<a href="#CommentGroup">CommentGroup</a>) <a href="#CommentMap">CommentMap</a></pre>
NewCommentMap creates a new comment map by associating comment groups
of the comments list with the nodes of the AST specified by node.

A comment group g is associated with a node n if:


	- g starts on the same line as n ends
	- g starts on the line immediately following n, and there is
	  at least one empty line after g and before the next node
	- g starts before n and is not associated to the node before n
	  via the previous rules

NewCommentMap tries to associate a comment group to the "largest"
node possible: For instance, if the comment is a line comment
trailing an assignment, the comment is associated with the entire
assignment rather than just the last operand in the assignment.






### <a id="CommentMap.Comments">func</a> (CommentMap) [Comments](https://golang.org/src/go/ast/commentmap.go?s=7690:7739#L262)
<pre>func (cmap <a href="#CommentMap">CommentMap</a>) Comments() []*<a href="#CommentGroup">CommentGroup</a></pre>
Comments returns the list of comment groups in the comment map.
The result is sorted in source order.




### <a id="CommentMap.Filter">func</a> (CommentMap) [Filter](https://golang.org/src/go/ast/commentmap.go?s=7379:7430#L248)
<pre>func (cmap <a href="#CommentMap">CommentMap</a>) Filter(node <a href="#Node">Node</a>) <a href="#CommentMap">CommentMap</a></pre>
Filter returns a new comment map consisting of only those
entries of cmap for which a corresponding node exists in
the AST specified by node.




### <a id="CommentMap.String">func</a> (CommentMap) [String](https://golang.org/src/go/ast/commentmap.go?s=8595:8633#L307)
<pre>func (cmap <a href="#CommentMap">CommentMap</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="CommentMap.Update">func</a> (CommentMap) [Update](https://golang.org/src/go/ast/commentmap.go?s=7055:7104#L236)
<pre>func (cmap <a href="#CommentMap">CommentMap</a>) Update(old, new <a href="#Node">Node</a>) <a href="#Node">Node</a></pre>
Update replaces an old node in the comment map with the new node
and returns the new node. Comments that were associated with the
old node are associated with the new node.




## <a id="CompositeLit">type</a> [CompositeLit](https://golang.org/src/go/ast/ast.go?s=7467:7767#L256)
A CompositeLit node represents a composite literal.


<pre>type CompositeLit struct {
<span id="CompositeLit.Type"></span>    Type       <a href="#Expr">Expr</a>      <span class="comment">// literal type; or nil</span>
<span id="CompositeLit.Lbrace"></span>    Lbrace     <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;{&#34;</span>
<span id="CompositeLit.Elts"></span>    Elts       []<a href="#Expr">Expr</a>    <span class="comment">// list of composite elements; or nil</span>
<span id="CompositeLit.Rbrace"></span>    Rbrace     <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;}&#34;</span>
<span id="CompositeLit.Incomplete"></span>    Incomplete <a href="/pkg/builtin/#bool">bool</a>      <span class="comment">// true if (source) expressions are missing in the Elts list</span>
}
</pre>











### <a id="CompositeLit.End">func</a> (\*CompositeLit) [End](https://golang.org/src/go/ast/ast.go?s=14126:14164#L456)
<pre>func (x *<a href="#CompositeLit">CompositeLit</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="CompositeLit.Pos">func</a> (\*CompositeLit) [Pos](https://golang.org/src/go/ast/ast.go?s=12491:12529#L418)
<pre>func (x *<a href="#CompositeLit">CompositeLit</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="Decl">type</a> [Decl](https://golang.org/src/go/ast/ast.go?s=1599:1640#L41)
All declaration nodes implement the Decl interface.


<pre>type Decl interface {
    <a href="#Node">Node</a>
    <span class="comment">// contains filtered or unexported methods</span>
}</pre>











## <a id="DeclStmt">type</a> [DeclStmt](https://golang.org/src/go/ast/ast.go?s=17359:17434#L545)
A DeclStmt node represents a declaration in a statement list.


<pre>type DeclStmt struct {
<span id="DeclStmt.Decl"></span>    Decl <a href="#Decl">Decl</a> <span class="comment">// *GenDecl with CONST, TYPE, or VAR token</span>
}
</pre>











### <a id="DeclStmt.End">func</a> (\*DeclStmt) [End](https://golang.org/src/go/ast/ast.go?s=23163:23197#L722)
<pre>func (s *<a href="#DeclStmt">DeclStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="DeclStmt.Pos">func</a> (\*DeclStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=21886:21920#L700)
<pre>func (s *<a href="#DeclStmt">DeclStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="DeferStmt">type</a> [DeferStmt](https://golang.org/src/go/ast/ast.go?s=18772:18860#L603)
A DeferStmt node represents a defer statement.


<pre>type DeferStmt struct {
<span id="DeferStmt.Defer"></span>    Defer <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;defer&#34; keyword</span>
<span id="DeferStmt.Call"></span>    Call  *<a href="#CallExpr">CallExpr</a>
}
</pre>











### <a id="DeferStmt.End">func</a> (\*DeferStmt) [End](https://golang.org/src/go/ast/ast.go?s=23739:23774#L737)
<pre>func (s *<a href="#DeferStmt">DeferStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="DeferStmt.Pos">func</a> (\*DeferStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=22394:22429#L708)
<pre>func (s *<a href="#DeferStmt">DeferStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="Ellipsis">type</a> [Ellipsis](https://golang.org/src/go/ast/ast.go?s=6809:6948#L237)
An Ellipsis node stands for the "..." type in a
parameter list or the "..." length in an array type.


<pre>type Ellipsis struct {
<span id="Ellipsis.Ellipsis"></span>    Ellipsis <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;...&#34;</span>
<span id="Ellipsis.Elt"></span>    Elt      <a href="#Expr">Expr</a>      <span class="comment">// ellipsis element type (parameter lists only); or nil</span>
}
</pre>











### <a id="Ellipsis.End">func</a> (\*Ellipsis) [End](https://golang.org/src/go/ast/ast.go?s=13848:13882#L448)
<pre>func (x *<a href="#Ellipsis">Ellipsis</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="Ellipsis.Pos">func</a> (\*Ellipsis) [Pos](https://golang.org/src/go/ast/ast.go?s=12318:12352#L415)
<pre>func (x *<a href="#Ellipsis">Ellipsis</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="EmptyStmt">type</a> [EmptyStmt](https://golang.org/src/go/ast/ast.go?s=17619:17754#L553)
An EmptyStmt node represents an empty statement.
The "position" of the empty statement is the position
of the immediately following (explicit or implicit) semicolon.


<pre>type EmptyStmt struct {
<span id="EmptyStmt.Semicolon"></span>    Semicolon <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of following &#34;;&#34;</span>
<span id="EmptyStmt.Implicit"></span>    Implicit  <a href="/pkg/builtin/#bool">bool</a>      <span class="comment">// if set, &#34;;&#34; was omitted in the source</span>
}
</pre>











### <a id="EmptyStmt.End">func</a> (\*EmptyStmt) [End](https://golang.org/src/go/ast/ast.go?s=23222:23257#L723)
<pre>func (s *<a href="#EmptyStmt">EmptyStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="EmptyStmt.Pos">func</a> (\*EmptyStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=21951:21986#L701)
<pre>func (s *<a href="#EmptyStmt">EmptyStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="Expr">type</a> [Expr](https://golang.org/src/go/ast/ast.go?s=1405:1446#L29)
All expression nodes implement the Expr interface.


<pre>type Expr interface {
    <a href="#Node">Node</a>
    <span class="comment">// contains filtered or unexported methods</span>
}</pre>











## <a id="ExprStmt">type</a> [ExprStmt](https://golang.org/src/go/ast/ast.go?s=17991:18034#L568)
An ExprStmt node represents a (stand-alone) expression
in a statement list.


<pre>type ExprStmt struct {
<span id="ExprStmt.X"></span>    X <a href="#Expr">Expr</a> <span class="comment">// expression</span>
}
</pre>











### <a id="ExprStmt.End">func</a> (\*ExprStmt) [End](https://golang.org/src/go/ast/ast.go?s=23404:23438#L730)
<pre>func (s *<a href="#ExprStmt">ExprStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="ExprStmt.Pos">func</a> (\*ExprStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=22081:22115#L703)
<pre>func (s *<a href="#ExprStmt">ExprStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="Field">type</a> [Field](https://golang.org/src/go/ast/ast.go?s=4589:4878#L147)
A Field represents a Field declaration list in a struct type,
a method list in an interface type, or a parameter/result declaration
in a signature.
Field.Names is nil for unnamed parameters (parameter lists which only contain types)
and embedded struct fields. In the latter case, the field name is the type name.


<pre>type Field struct {
<span id="Field.Doc"></span>    Doc     *<a href="#CommentGroup">CommentGroup</a> <span class="comment">// associated documentation; or nil</span>
<span id="Field.Names"></span>    Names   []*<a href="#Ident">Ident</a>      <span class="comment">// field/method/parameter names; or nil</span>
<span id="Field.Type"></span>    Type    <a href="#Expr">Expr</a>          <span class="comment">// field/method/parameter type</span>
<span id="Field.Tag"></span>    Tag     *<a href="#BasicLit">BasicLit</a>     <span class="comment">// field tag; or nil</span>
<span id="Field.Comment"></span>    Comment *<a href="#CommentGroup">CommentGroup</a> <span class="comment">// line comments; or nil</span>
}
</pre>











### <a id="Field.End">func</a> (\*Field) [End](https://golang.org/src/go/ast/ast.go?s=4990:5021#L162)
<pre>func (f *<a href="#Field">Field</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="Field.Pos">func</a> (\*Field) [Pos](https://golang.org/src/go/ast/ast.go?s=4880:4911#L155)
<pre>func (f *<a href="#Field">Field</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="FieldFilter">type</a> [FieldFilter](https://golang.org/src/go/ast/print.go?s=343:403#L8)
A FieldFilter may be provided to Fprint to control the output.


<pre>type FieldFilter func(name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/reflect/">reflect</a>.<a href="/pkg/reflect/#Value">Value</a>) <a href="/pkg/builtin/#bool">bool</a></pre>











## <a id="FieldList">type</a> [FieldList](https://golang.org/src/go/ast/ast.go?s=5170:5372#L170)
A FieldList represents a list of Fields, enclosed by parentheses or braces.


<pre>type FieldList struct {
<span id="FieldList.Opening"></span>    Opening <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of opening parenthesis/brace, if any</span>
<span id="FieldList.List"></span>    List    []*<a href="#Field">Field</a>  <span class="comment">// field list; or nil</span>
<span id="FieldList.Closing"></span>    Closing <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of closing parenthesis/brace, if any</span>
}
</pre>











### <a id="FieldList.End">func</a> (\*FieldList) [End](https://golang.org/src/go/ast/ast.go?s=5627:5662#L188)
<pre>func (f *<a href="#FieldList">FieldList</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="FieldList.NumFields">func</a> (\*FieldList) [NumFields](https://golang.org/src/go/ast/ast.go?s=5985:6020#L201)
<pre>func (f *<a href="#FieldList">FieldList</a>) NumFields() <a href="/pkg/builtin/#int">int</a></pre>
NumFields returns the number of parameters or struct fields represented by a FieldList.




### <a id="FieldList.Pos">func</a> (\*FieldList) [Pos](https://golang.org/src/go/ast/ast.go?s=5374:5409#L176)
<pre>func (f *<a href="#FieldList">FieldList</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="File">type</a> [File](https://golang.org/src/go/ast/ast.go?s=31191:31694#L969)
A File node represents a Go source file.

The Comments list contains all comments in the source file in order of
appearance, including the comments that are pointed to from other nodes
via Doc and Comment fields.

For correct printing of source code containing comments (using packages
go/format and go/printer), special care must be taken to update comments
when a File's syntax tree is modified: For printing, comments are interspersed
between tokens based on their position. If syntax tree nodes are
removed or moved, relevant comments in their vicinity must also be removed
(from the File.Comments list) or moved accordingly (by updating their
positions). A CommentMap may be used to facilitate some of these operations.

Whether and how a comment is associated with a node depends on the
interpretation of the syntax tree by the manipulating program: Except for Doc
and Comment comments directly associated with nodes, the remaining comments
are "free-floating" (see also issues #18593, #20744).


<pre>type File struct {
<span id="File.Doc"></span>    Doc        *<a href="#CommentGroup">CommentGroup</a>   <span class="comment">// associated documentation; or nil</span>
<span id="File.Package"></span>    Package    <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>       <span class="comment">// position of &#34;package&#34; keyword</span>
<span id="File.Name"></span>    Name       *<a href="#Ident">Ident</a>          <span class="comment">// package name</span>
<span id="File.Decls"></span>    Decls      []<a href="#Decl">Decl</a>          <span class="comment">// top-level declarations; or nil</span>
<span id="File.Scope"></span>    Scope      *<a href="#Scope">Scope</a>          <span class="comment">// package scope (this file only)</span>
<span id="File.Imports"></span>    Imports    []*<a href="#ImportSpec">ImportSpec</a>   <span class="comment">// imports in this file</span>
<span id="File.Unresolved"></span>    Unresolved []*<a href="#Ident">Ident</a>        <span class="comment">// unresolved identifiers in this file</span>
<span id="File.Comments"></span>    Comments   []*<a href="#CommentGroup">CommentGroup</a> <span class="comment">// list of all comments in the source file</span>
}
</pre>









### <a id="MergePackageFiles">func</a> [MergePackageFiles](https://golang.org/src/go/ast/filter.go?s=8888:8946#L334)
<pre>func MergePackageFiles(pkg *<a href="#Package">Package</a>, mode <a href="#MergeMode">MergeMode</a>) *<a href="#File">File</a></pre>
MergePackageFiles creates a file AST by merging the ASTs of the
files belonging to a package. The mode flags control merging behavior.






### <a id="File.End">func</a> (\*File) [End](https://golang.org/src/go/ast/ast.go?s=31748:31778#L981)
<pre>func (f *<a href="#File">File</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="File.Pos">func</a> (\*File) [Pos](https://golang.org/src/go/ast/ast.go?s=31696:31726#L980)
<pre>func (f *<a href="#File">File</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="Filter">type</a> [Filter](https://golang.org/src/go/ast/filter.go?s=1451:1480#L36)

<pre>type Filter func(<a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>











## <a id="ForStmt">type</a> [ForStmt](https://golang.org/src/go/ast/ast.go?s=21146:21373#L678)
A ForStmt represents a for statement.


<pre>type ForStmt struct {
<span id="ForStmt.For"></span>    For  <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;for&#34; keyword</span>
<span id="ForStmt.Init"></span>    Init <a href="#Stmt">Stmt</a>      <span class="comment">// initialization statement; or nil</span>
<span id="ForStmt.Cond"></span>    Cond <a href="#Expr">Expr</a>      <span class="comment">// condition; or nil</span>
<span id="ForStmt.Post"></span>    Post <a href="#Stmt">Stmt</a>      <span class="comment">// post iteration statement; or nil</span>
<span id="ForStmt.Body"></span>    Body *<a href="#BlockStmt">BlockStmt</a>
}
</pre>











### <a id="ForStmt.End">func</a> (\*ForStmt) [End](https://golang.org/src/go/ast/ast.go?s=24684:24717#L772)
<pre>func (s *<a href="#ForStmt">ForStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="ForStmt.Pos">func</a> (\*ForStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=22995:23028#L718)
<pre>func (s *<a href="#ForStmt">ForStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="FuncDecl">type</a> [FuncDecl](https://golang.org/src/go/ast/ast.go?s=29002:29360#L911)
A FuncDecl node represents a function declaration.


<pre>type FuncDecl struct {
<span id="FuncDecl.Doc"></span>    Doc  *<a href="#CommentGroup">CommentGroup</a> <span class="comment">// associated documentation; or nil</span>
<span id="FuncDecl.Recv"></span>    Recv *<a href="#FieldList">FieldList</a>    <span class="comment">// receiver (methods); or nil (functions)</span>
<span id="FuncDecl.Name"></span>    Name *<a href="#Ident">Ident</a>        <span class="comment">// function/method name</span>
<span id="FuncDecl.Type"></span>    Type *<a href="#FuncType">FuncType</a>     <span class="comment">// function signature: parameters, results, and position of &#34;func&#34; keyword</span>
<span id="FuncDecl.Body"></span>    Body *<a href="#BlockStmt">BlockStmt</a>    <span class="comment">// function body; or nil for external (non-Go) function</span>
}
</pre>











### <a id="FuncDecl.End">func</a> (\*FuncDecl) [End](https://golang.org/src/go/ast/ast.go?s=29750:29784#L933)
<pre>func (d *<a href="#FuncDecl">FuncDecl</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="FuncDecl.Pos">func</a> (\*FuncDecl) [Pos](https://golang.org/src/go/ast/ast.go?s=29527:29561#L924)
<pre>func (d *<a href="#FuncDecl">FuncDecl</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="FuncLit">type</a> [FuncLit](https://golang.org/src/go/ast/ast.go?s=7319:7408#L250)
A FuncLit node represents a function literal.


<pre>type FuncLit struct {
<span id="FuncLit.Type"></span>    Type *<a href="#FuncType">FuncType</a>  <span class="comment">// function type</span>
<span id="FuncLit.Body"></span>    Body *<a href="#BlockStmt">BlockStmt</a> <span class="comment">// function body</span>
}
</pre>











### <a id="FuncLit.End">func</a> (\*FuncLit) [End](https://golang.org/src/go/ast/ast.go?s=14061:14094#L455)
<pre>func (x *<a href="#FuncLit">FuncLit</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="FuncLit.Pos">func</a> (\*FuncLit) [Pos](https://golang.org/src/go/ast/ast.go?s=12432:12465#L417)
<pre>func (x *<a href="#FuncLit">FuncLit</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="FuncType">type</a> [FuncType](https://golang.org/src/go/ast/ast.go?s=11233:11446#L382)
A FuncType node represents a function type.


<pre>type FuncType struct {
<span id="FuncType.Func"></span>    Func    <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>  <span class="comment">// position of &#34;func&#34; keyword (token.NoPos if there is no &#34;func&#34;)</span>
<span id="FuncType.Params"></span>    Params  *<a href="#FieldList">FieldList</a> <span class="comment">// (incoming) parameters; non-nil</span>
<span id="FuncType.Results"></span>    Results *<a href="#FieldList">FieldList</a> <span class="comment">// (outgoing) results; or nil</span>
}
</pre>











### <a id="FuncType.End">func</a> (\*FuncType) [End](https://golang.org/src/go/ast/ast.go?s=14963:14997#L469)
<pre>func (x *<a href="#FuncType">FuncType</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="FuncType.Pos">func</a> (\*FuncType) [Pos](https://golang.org/src/go/ast/ast.go?s=13336:13370#L436)
<pre>func (x *<a href="#FuncType">FuncType</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="GenDecl">type</a> [GenDecl](https://golang.org/src/go/ast/ast.go?s=28661:28944#L901)
A GenDecl node (generic declaration node) represents an import,
constant, type or variable declaration. A valid Lparen position
(Lparen.IsValid()) indicates a parenthesized declaration.

Relationship between Tok value and Specs element type:


	token.IMPORT  *ImportSpec
	token.CONST   *ValueSpec
	token.TYPE    *TypeSpec
	token.VAR     *ValueSpec


<pre>type GenDecl struct {
<span id="GenDecl.Doc"></span>    Doc    *<a href="#CommentGroup">CommentGroup</a> <span class="comment">// associated documentation; or nil</span>
<span id="GenDecl.TokPos"></span>    TokPos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>     <span class="comment">// position of Tok</span>
<span id="GenDecl.Tok"></span>    Tok    <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Token">Token</a>   <span class="comment">// IMPORT, CONST, TYPE, VAR</span>
<span id="GenDecl.Lparen"></span>    Lparen <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>     <span class="comment">// position of &#39;(&#39;, if any</span>
<span id="GenDecl.Specs"></span>    Specs  []<a href="#Spec">Spec</a>
<span id="GenDecl.Rparen"></span>    Rparen <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#39;)&#39;, if any</span>
}
</pre>











### <a id="GenDecl.End">func</a> (\*GenDecl) [End](https://golang.org/src/go/ast/ast.go?s=29637:29670#L927)
<pre>func (d *<a href="#GenDecl">GenDecl</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="GenDecl.Pos">func</a> (\*GenDecl) [Pos](https://golang.org/src/go/ast/ast.go?s=29472:29505#L923)
<pre>func (d *<a href="#GenDecl">GenDecl</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="GoStmt">type</a> [GoStmt](https://golang.org/src/go/ast/ast.go?s=18638:18718#L597)
A GoStmt node represents a go statement.


<pre>type GoStmt struct {
<span id="GoStmt.Go"></span>    Go   <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;go&#34; keyword</span>
<span id="GoStmt.Call"></span>    Call *<a href="#CallExpr">CallExpr</a>
}
</pre>











### <a id="GoStmt.End">func</a> (\*GoStmt) [End](https://golang.org/src/go/ast/ast.go?s=23678:23710#L736)
<pre>func (s *<a href="#GoStmt">GoStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="GoStmt.Pos">func</a> (\*GoStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=22337:22369#L707)
<pre>func (s *<a href="#GoStmt">GoStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="Ident">type</a> [Ident](https://golang.org/src/go/ast/ast.go?s=6548:6693#L228)
An Ident node represents an identifier.


<pre>type Ident struct {
<span id="Ident.NamePos"></span>    NamePos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// identifier position</span>
<span id="Ident.Name"></span>    Name    <a href="/pkg/builtin/#string">string</a>    <span class="comment">// identifier name</span>
<span id="Ident.Obj"></span>    Obj     *<a href="#Object">Object</a>   <span class="comment">// denoted object; or nil</span>
}
</pre>









### <a id="NewIdent">func</a> [NewIdent](https://golang.org/src/go/ast/ast.go?s=16405:16438#L512)
<pre>func NewIdent(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Ident">Ident</a></pre>
NewIdent creates a new Ident without position.
Useful for ASTs generated by code other than the Go parser.






### <a id="Ident.End">func</a> (\*Ident) [End](https://golang.org/src/go/ast/ast.go?s=13763:13794#L447)
<pre>func (x *<a href="#Ident">Ident</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="Ident.IsExported">func</a> (\*Ident) [IsExported](https://golang.org/src/go/ast/ast.go?s=16693:16727#L520)
<pre>func (id *<a href="#Ident">Ident</a>) IsExported() <a href="/pkg/builtin/#bool">bool</a></pre>
IsExported reports whether id starts with an upper-case letter.




### <a id="Ident.Pos">func</a> (\*Ident) [Pos](https://golang.org/src/go/ast/ast.go?s=12262:12293#L414)
<pre>func (x *<a href="#Ident">Ident</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="Ident.String">func</a> (\*Ident) [String](https://golang.org/src/go/ast/ast.go?s=16766:16798#L522)
<pre>func (id *<a href="#Ident">Ident</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="IfStmt">type</a> [IfStmt](https://golang.org/src/go/ast/ast.go?s=19538:19737#L631)
An IfStmt node represents an if statement.


<pre>type IfStmt struct {
<span id="IfStmt.If"></span>    If   <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;if&#34; keyword</span>
<span id="IfStmt.Init"></span>    Init <a href="#Stmt">Stmt</a>      <span class="comment">// initialization statement; or nil</span>
<span id="IfStmt.Cond"></span>    Cond <a href="#Expr">Expr</a>      <span class="comment">// condition</span>
<span id="IfStmt.Body"></span>    Body *<a href="#BlockStmt">BlockStmt</a>
<span id="IfStmt.Else"></span>    Else <a href="#Stmt">Stmt</a> <span class="comment">// else branch; or nil</span>
}
</pre>











### <a id="IfStmt.End">func</a> (\*IfStmt) [End](https://golang.org/src/go/ast/ast.go?s=24148:24180#L751)
<pre>func (s *<a href="#IfStmt">IfStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="IfStmt.Pos">func</a> (\*IfStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=22637:22669#L712)
<pre>func (s *<a href="#IfStmt">IfStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="ImportSpec">type</a> [ImportSpec](https://golang.org/src/go/ast/ast.go?s=26058:26369#L814)
An ImportSpec node represents a single package import.


<pre>type ImportSpec struct {
<span id="ImportSpec.Doc"></span>    Doc     *<a href="#CommentGroup">CommentGroup</a> <span class="comment">// associated documentation; or nil</span>
<span id="ImportSpec.Name"></span>    Name    *<a href="#Ident">Ident</a>        <span class="comment">// local package name (including &#34;.&#34;); or nil</span>
<span id="ImportSpec.Path"></span>    Path    *<a href="#BasicLit">BasicLit</a>     <span class="comment">// import path</span>
<span id="ImportSpec.Comment"></span>    Comment *<a href="#CommentGroup">CommentGroup</a> <span class="comment">// line comments; or nil</span>
<span id="ImportSpec.EndPos"></span>    EndPos  <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>     <span class="comment">// end of spec (overrides Path.Pos if nonzero)</span>
}
</pre>











### <a id="ImportSpec.End">func</a> (\*ImportSpec) [End](https://golang.org/src/go/ast/ast.go?s=27438:27474#L854)
<pre>func (s *<a href="#ImportSpec">ImportSpec</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="ImportSpec.Pos">func</a> (\*ImportSpec) [Pos](https://golang.org/src/go/ast/ast.go?s=27206:27242#L845)
<pre>func (s *<a href="#ImportSpec">ImportSpec</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="Importer">type</a> [Importer](https://golang.org/src/go/ast/resolve.go?s=1751:1835#L53)
An Importer resolves import paths to package Objects.
The imports map records the packages already imported,
indexed by package id (canonical import path).
An Importer must determine the canonical import path and
check the map to see if it is already present in the imports map.
If so, the Importer can return the map entry. Otherwise, the
Importer should load the package data for the given path into
a new *Object (pkg), record pkg in the imports map, and then
return pkg.


<pre>type Importer func(imports map[<a href="/pkg/builtin/#string">string</a>]*<a href="#Object">Object</a>, path <a href="/pkg/builtin/#string">string</a>) (pkg *<a href="#Object">Object</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>











## <a id="IncDecStmt">type</a> [IncDecStmt](https://golang.org/src/go/ast/ast.go?s=18244:18355#L580)
An IncDecStmt node represents an increment or decrement statement.


<pre>type IncDecStmt struct {
<span id="IncDecStmt.X"></span>    X      <a href="#Expr">Expr</a>
<span id="IncDecStmt.TokPos"></span>    TokPos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>   <span class="comment">// position of Tok</span>
<span id="IncDecStmt.Tok"></span>    Tok    <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Token">Token</a> <span class="comment">// INC or DEC</span>
}
</pre>











### <a id="IncDecStmt.End">func</a> (\*IncDecStmt) [End](https://golang.org/src/go/ast/ast.go?s=23526:23562#L732)
<pre>func (s *<a href="#IncDecStmt">IncDecStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="IncDecStmt.Pos">func</a> (\*IncDecStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=22208:22244#L705)
<pre>func (s *<a href="#IncDecStmt">IncDecStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="IndexExpr">type</a> [IndexExpr](https://golang.org/src/go/ast/ast.go?s=8204:8373#L278)
An IndexExpr node represents an expression followed by an index.


<pre>type IndexExpr struct {
<span id="IndexExpr.X"></span>    X      <a href="#Expr">Expr</a>      <span class="comment">// expression</span>
<span id="IndexExpr.Lbrack"></span>    Lbrack <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;[&#34;</span>
<span id="IndexExpr.Index"></span>    Index  <a href="#Expr">Expr</a>      <span class="comment">// index expression</span>
<span id="IndexExpr.Rbrack"></span>    Rbrack <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;]&#34;</span>
}
</pre>











### <a id="IndexExpr.End">func</a> (\*IndexExpr) [End](https://golang.org/src/go/ast/ast.go?s=14320:14355#L459)
<pre>func (x *<a href="#IndexExpr">IndexExpr</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="IndexExpr.Pos">func</a> (\*IndexExpr) [Pos](https://golang.org/src/go/ast/ast.go?s=12719:12754#L426)
<pre>func (x *<a href="#IndexExpr">IndexExpr</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="InterfaceType">type</a> [InterfaceType](https://golang.org/src/go/ast/ast.go?s=11505:11716#L389)
An InterfaceType node represents an interface type.


<pre>type InterfaceType struct {
<span id="InterfaceType.Interface"></span>    Interface  <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>  <span class="comment">// position of &#34;interface&#34; keyword</span>
<span id="InterfaceType.Methods"></span>    Methods    *<a href="#FieldList">FieldList</a> <span class="comment">// list of methods</span>
<span id="InterfaceType.Incomplete"></span>    Incomplete <a href="/pkg/builtin/#bool">bool</a>       <span class="comment">// true if (source) methods are missing in the Methods list</span>
}
</pre>











### <a id="InterfaceType.End">func</a> (\*InterfaceType) [End](https://golang.org/src/go/ast/ast.go?s=15076:15115#L475)
<pre>func (x *<a href="#InterfaceType">InterfaceType</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="InterfaceType.Pos">func</a> (\*InterfaceType) [Pos](https://golang.org/src/go/ast/ast.go?s=13533:13572#L442)
<pre>func (x *<a href="#InterfaceType">InterfaceType</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="KeyValueExpr">type</a> [KeyValueExpr](https://golang.org/src/go/ast/ast.go?s=10234:10321#L343)
A KeyValueExpr node represents (key : value) pairs
in composite literals.


<pre>type KeyValueExpr struct {
<span id="KeyValueExpr.Key"></span>    Key   <a href="#Expr">Expr</a>
<span id="KeyValueExpr.Colon"></span>    Colon <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;:&#34;</span>
<span id="KeyValueExpr.Value"></span>    Value <a href="#Expr">Expr</a>
}
</pre>











### <a id="KeyValueExpr.End">func</a> (\*KeyValueExpr) [End](https://golang.org/src/go/ast/ast.go?s=14766:14804#L466)
<pre>func (x *<a href="#KeyValueExpr">KeyValueExpr</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="KeyValueExpr.Pos">func</a> (\*KeyValueExpr) [Pos](https://golang.org/src/go/ast/ast.go?s=13150:13188#L433)
<pre>func (x *<a href="#KeyValueExpr">KeyValueExpr</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="LabeledStmt">type</a> [LabeledStmt](https://golang.org/src/go/ast/ast.go?s=17812:17900#L559)
A LabeledStmt node represents a labeled statement.


<pre>type LabeledStmt struct {
<span id="LabeledStmt.Label"></span>    Label *<a href="#Ident">Ident</a>
<span id="LabeledStmt.Colon"></span>    Colon <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;:&#34;</span>
<span id="LabeledStmt.Stmt"></span>    Stmt  <a href="#Stmt">Stmt</a>
}
</pre>











### <a id="LabeledStmt.End">func</a> (\*LabeledStmt) [End](https://golang.org/src/go/ast/ast.go?s=23342:23379#L729)
<pre>func (s *<a href="#LabeledStmt">LabeledStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="LabeledStmt.Pos">func</a> (\*LabeledStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=22015:22052#L702)
<pre>func (s *<a href="#LabeledStmt">LabeledStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="MapType">type</a> [MapType](https://golang.org/src/go/ast/ast.go?s=11761:11853#L396)
A MapType node represents a map type.


<pre>type MapType struct {
<span id="MapType.Map"></span>    Map   <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;map&#34; keyword</span>
<span id="MapType.Key"></span>    Key   <a href="#Expr">Expr</a>
<span id="MapType.Value"></span>    Value <a href="#Expr">Expr</a>
}
</pre>











### <a id="MapType.End">func</a> (\*MapType) [End](https://golang.org/src/go/ast/ast.go?s=15143:15176#L476)
<pre>func (x *<a href="#MapType">MapType</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="MapType.Pos">func</a> (\*MapType) [Pos](https://golang.org/src/go/ast/ast.go?s=13596:13629#L443)
<pre>func (x *<a href="#MapType">MapType</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="MergeMode">type</a> [MergeMode](https://golang.org/src/go/ast/filter.go?s=7588:7607#L293)
The MergeMode flags control the behavior of MergePackageFiles.


<pre>type MergeMode <a href="/pkg/builtin/#uint">uint</a></pre>



<pre>const (
    <span class="comment">// If set, duplicate function declarations are excluded.</span>
    <span id="FilterFuncDuplicates">FilterFuncDuplicates</span> <a href="#MergeMode">MergeMode</a> = 1 &lt;&lt; <a href="/pkg/builtin/#iota">iota</a>
    <span class="comment">// If set, comments that are not associated with a specific</span>
    <span class="comment">// AST node (as Doc or Comment) are excluded.</span>
    <span id="FilterUnassociatedComments">FilterUnassociatedComments</span>
    <span class="comment">// If set, duplicate import declarations are excluded.</span>
    <span id="FilterImportDuplicates">FilterImportDuplicates</span>
)</pre>









## <a id="Node">type</a> [Node](https://golang.org/src/go/ast/ast.go?s=1181:1349#L23)
All node types implement the Node interface.


<pre>type Node interface {
    Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of first character belonging to the node</span>
    End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of first character immediately after the node</span>
}</pre>











## <a id="ObjKind">type</a> [ObjKind](https://golang.org/src/go/ast/scope.go?s=3516:3532#L127)
ObjKind describes what an object represents.


<pre>type ObjKind <a href="/pkg/builtin/#int">int</a></pre>


The list of possible Object kinds.


<pre>const (
    <span id="Bad">Bad</span> <a href="#ObjKind">ObjKind</a> = <a href="/pkg/builtin/#iota">iota</a> <span class="comment">// for error handling</span>
    <span id="Pkg">Pkg</span>                <span class="comment">// package</span>
    <span id="Con">Con</span>                <span class="comment">// constant</span>
    <span id="Typ">Typ</span>                <span class="comment">// type</span>
    <span id="Var">Var</span>                <span class="comment">// variable</span>
    <span id="Fun">Fun</span>                <span class="comment">// function or method</span>
    <span id="Lbl">Lbl</span>                <span class="comment">// label</span>
)</pre>









### <a id="ObjKind.String">func</a> (ObjKind) [String](https://golang.org/src/go/ast/scope.go?s=3957:3992#L150)
<pre>func (kind <a href="#ObjKind">ObjKind</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Object">type</a> [Object](https://golang.org/src/go/ast/scope.go?s=2004:2291#L66)
An Object describes a named language entity such as a package,
constant, type, variable, function (incl. methods), or label.

The Data fields contains object-specific data:


	Kind    Data type         Data value
	Pkg     *Scope            package scope
	Con     int               iota for the respective declaration


<pre>type Object struct {
<span id="Object.Kind"></span>    Kind <a href="#ObjKind">ObjKind</a>
<span id="Object.Name"></span>    Name <a href="/pkg/builtin/#string">string</a>      <span class="comment">// declared name</span>
<span id="Object.Decl"></span>    Decl interface{} <span class="comment">// corresponding Field, XxxSpec, FuncDecl, LabeledStmt, AssignStmt, Scope; or nil</span>
<span id="Object.Data"></span>    Data interface{} <span class="comment">// object-specific data; or nil</span>
<span id="Object.Type"></span>    Type interface{} <span class="comment">// placeholder for type information; may be nil</span>
}
</pre>









### <a id="NewObj">func</a> [NewObj](https://golang.org/src/go/ast/scope.go?s=2350:2396#L75)
<pre>func NewObj(kind <a href="#ObjKind">ObjKind</a>, name <a href="/pkg/builtin/#string">string</a>) *<a href="#Object">Object</a></pre>
NewObj creates a new object of a given kind and name.






### <a id="Object.Pos">func</a> (\*Object) [Pos](https://golang.org/src/go/ast/scope.go?s=2623:2657#L82)
<pre>func (obj *<a href="#Object">Object</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>
Pos computes the source position of the declaration of an object name.
The result may be an invalid position if it cannot be computed
(obj.Decl may be nil or not correct).




## <a id="Package">type</a> [Package](https://golang.org/src/go/ast/ast.go?s=31960:32215#L991)
A Package node represents a set of source files
collectively building a Go package.


<pre>type Package struct {
<span id="Package.Name"></span>    Name    <a href="/pkg/builtin/#string">string</a>             <span class="comment">// package name</span>
<span id="Package.Scope"></span>    Scope   *<a href="#Scope">Scope</a>             <span class="comment">// package scope across all files</span>
<span id="Package.Imports"></span>    Imports map[<a href="/pkg/builtin/#string">string</a>]*<a href="#Object">Object</a> <span class="comment">// map of package id -&gt; package object</span>
<span id="Package.Files"></span>    Files   map[<a href="/pkg/builtin/#string">string</a>]*<a href="#File">File</a>   <span class="comment">// Go source files by filename</span>
}
</pre>









### <a id="NewPackage">func</a> [NewPackage](https://golang.org/src/go/ast/resolve.go?s=2445:2559#L64)
<pre>func NewPackage(fset *<a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#FileSet">FileSet</a>, files map[<a href="/pkg/builtin/#string">string</a>]*<a href="#File">File</a>, importer <a href="#Importer">Importer</a>, universe *<a href="#Scope">Scope</a>) (*<a href="#Package">Package</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewPackage creates a new Package node from a set of File nodes. It resolves
unresolved identifiers across files and updates each file's Unresolved list
accordingly. If a non-nil importer and universe scope are provided, they are
used to resolve identifiers not declared in any of the package files. Any
remaining unresolved identifiers are reported as undeclared. If the files
belong to different packages, one package name is selected and files with
different package names are reported and then ignored.
The result is a package node and a scanner.ErrorList if there were errors.






### <a id="Package.End">func</a> (\*Package) [End](https://golang.org/src/go/ast/ast.go?s=32274:32307#L999)
<pre>func (p *<a href="#Package">Package</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="Package.Pos">func</a> (\*Package) [Pos](https://golang.org/src/go/ast/ast.go?s=32217:32250#L998)
<pre>func (p *<a href="#Package">Package</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="ParenExpr">type</a> [ParenExpr](https://golang.org/src/go/ast/ast.go?s=7830:7974#L265)
A ParenExpr node represents a parenthesized expression.


<pre>type ParenExpr struct {
<span id="ParenExpr.Lparen"></span>    Lparen <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;(&#34;</span>
<span id="ParenExpr.X"></span>    X      <a href="#Expr">Expr</a>      <span class="comment">// parenthesized expression</span>
<span id="ParenExpr.Rparen"></span>    Rparen <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;)&#34;</span>
}
</pre>











### <a id="ParenExpr.End">func</a> (\*ParenExpr) [End](https://golang.org/src/go/ast/ast.go?s=14191:14226#L457)
<pre>func (x *<a href="#ParenExpr">ParenExpr</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="ParenExpr.Pos">func</a> (\*ParenExpr) [Pos](https://golang.org/src/go/ast/ast.go?s=12596:12631#L424)
<pre>func (x *<a href="#ParenExpr">ParenExpr</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="RangeStmt">type</a> [RangeStmt](https://golang.org/src/go/ast/ast.go?s=21440:21770#L687)
A RangeStmt represents a for statement with a range clause.


<pre>type RangeStmt struct {
<span id="RangeStmt.For"></span>    For        <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>   <span class="comment">// position of &#34;for&#34; keyword</span>
<span id="RangeStmt.Key"></span>    Key, Value <a href="#Expr">Expr</a>        <span class="comment">// Key, Value may be nil</span>
<span id="RangeStmt.TokPos"></span>    TokPos     <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>   <span class="comment">// position of Tok; invalid if Key == nil</span>
<span id="RangeStmt.Tok"></span>    Tok        <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Token">Token</a> <span class="comment">// ILLEGAL if Key == nil, ASSIGN, DEFINE</span>
<span id="RangeStmt.X"></span>    X          <a href="#Expr">Expr</a>        <span class="comment">// value to range over</span>
<span id="RangeStmt.Body"></span>    Body       *<a href="#BlockStmt">BlockStmt</a>
}
</pre>











### <a id="RangeStmt.End">func</a> (\*RangeStmt) [End](https://golang.org/src/go/ast/ast.go?s=24745:24780#L773)
<pre>func (s *<a href="#RangeStmt">RangeStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="RangeStmt.Pos">func</a> (\*RangeStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=23053:23088#L719)
<pre>func (s *<a href="#RangeStmt">RangeStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="ReturnStmt">type</a> [ReturnStmt](https://golang.org/src/go/ast/ast.go?s=18916:19040#L609)
A ReturnStmt node represents a return statement.


<pre>type ReturnStmt struct {
<span id="ReturnStmt.Return"></span>    Return  <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;return&#34; keyword</span>
<span id="ReturnStmt.Results"></span>    Results []<a href="#Expr">Expr</a>    <span class="comment">// result expressions; or nil</span>
}
</pre>











### <a id="ReturnStmt.End">func</a> (\*ReturnStmt) [End](https://golang.org/src/go/ast/ast.go?s=23800:23836#L738)
<pre>func (s *<a href="#ReturnStmt">ReturnStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="ReturnStmt.Pos">func</a> (\*ReturnStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=22454:22490#L709)
<pre>func (s *<a href="#ReturnStmt">ReturnStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="Scope">type</a> [Scope](https://golang.org/src/go/ast/scope.go?s=419:484#L9)
A Scope maintains the set of named language entities declared
in the scope and a link to the immediately surrounding (outer)
scope.


<pre>type Scope struct {
<span id="Scope.Outer"></span>    Outer   *<a href="#Scope">Scope</a>
<span id="Scope.Objects"></span>    Objects map[<a href="/pkg/builtin/#string">string</a>]*<a href="#Object">Object</a>
}
</pre>









### <a id="NewScope">func</a> [NewScope](https://golang.org/src/go/ast/scope.go?s=545:579#L15)
<pre>func NewScope(outer *<a href="#Scope">Scope</a>) *<a href="#Scope">Scope</a></pre>
NewScope creates a new scope nested in the outer scope.






### <a id="Scope.Insert">func</a> (\*Scope) [Insert](https://golang.org/src/go/ast/scope.go?s=1120:1169#L33)
<pre>func (s *<a href="#Scope">Scope</a>) Insert(obj *<a href="#Object">Object</a>) (alt *<a href="#Object">Object</a>)</pre>
Insert attempts to insert a named object obj into the scope s.
If the scope already contains an object alt with the same name,
Insert leaves the scope unchanged and returns alt. Otherwise
it inserts obj and returns nil.




### <a id="Scope.Lookup">func</a> (\*Scope) [Lookup](https://golang.org/src/go/ast/scope.go?s=812:855#L24)
<pre>func (s *<a href="#Scope">Scope</a>) Lookup(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Object">Object</a></pre>
Lookup returns the object with the given name if it is
found in scope s, otherwise it returns nil. Outer scopes
are ignored.




### <a id="Scope.String">func</a> (\*Scope) [String](https://golang.org/src/go/ast/scope.go?s=1279:1310#L41)
<pre>func (s *<a href="#Scope">Scope</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
Debugging support




## <a id="SelectStmt">type</a> [SelectStmt](https://golang.org/src/go/ast/ast.go?s=20987:21101#L672)
An SelectStmt node represents a select statement.


<pre>type SelectStmt struct {
<span id="SelectStmt.Select"></span>    Select <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>  <span class="comment">// position of &#34;select&#34; keyword</span>
<span id="SelectStmt.Body"></span>    Body   *<a href="#BlockStmt">BlockStmt</a> <span class="comment">// CommClauses only</span>
}
</pre>











### <a id="SelectStmt.End">func</a> (\*SelectStmt) [End](https://golang.org/src/go/ast/ast.go?s=24623:24659#L771)
<pre>func (s *<a href="#SelectStmt">SelectStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="SelectStmt.Pos">func</a> (\*SelectStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=22934:22970#L717)
<pre>func (s *<a href="#SelectStmt">SelectStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="SelectorExpr">type</a> [SelectorExpr](https://golang.org/src/go/ast/ast.go?s=8050:8132#L272)
A SelectorExpr node represents an expression followed by a selector.


<pre>type SelectorExpr struct {
<span id="SelectorExpr.X"></span>    X   <a href="#Expr">Expr</a>   <span class="comment">// expression</span>
<span id="SelectorExpr.Sel"></span>    Sel *<a href="#Ident">Ident</a> <span class="comment">// field selector</span>
}
</pre>











### <a id="SelectorExpr.End">func</a> (\*SelectorExpr) [End](https://golang.org/src/go/ast/ast.go?s=14256:14294#L458)
<pre>func (x *<a href="#SelectorExpr">SelectorExpr</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="SelectorExpr.Pos">func</a> (\*SelectorExpr) [Pos](https://golang.org/src/go/ast/ast.go?s=12657:12695#L425)
<pre>func (x *<a href="#SelectorExpr">SelectorExpr</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="SendStmt">type</a> [SendStmt](https://golang.org/src/go/ast/ast.go?s=18086:18170#L573)
A SendStmt node represents a send statement.


<pre>type SendStmt struct {
<span id="SendStmt.Chan"></span>    Chan  <a href="#Expr">Expr</a>
<span id="SendStmt.Arrow"></span>    Arrow <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;&lt;-&#34;</span>
<span id="SendStmt.Value"></span>    Value <a href="#Expr">Expr</a>
}
</pre>











### <a id="SendStmt.End">func</a> (\*SendStmt) [End](https://golang.org/src/go/ast/ast.go?s=23463:23497#L731)
<pre>func (s *<a href="#SendStmt">SendStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="SendStmt.Pos">func</a> (\*SendStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=22143:22177#L704)
<pre>func (s *<a href="#SendStmt">SendStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="SliceExpr">type</a> [SliceExpr](https://golang.org/src/go/ast/ast.go?s=8450:8799#L286)
An SliceExpr node represents an expression followed by slice indices.


<pre>type SliceExpr struct {
<span id="SliceExpr.X"></span>    X      <a href="#Expr">Expr</a>      <span class="comment">// expression</span>
<span id="SliceExpr.Lbrack"></span>    Lbrack <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;[&#34;</span>
<span id="SliceExpr.Low"></span>    Low    <a href="#Expr">Expr</a>      <span class="comment">// begin of slice range; or nil</span>
<span id="SliceExpr.High"></span>    High   <a href="#Expr">Expr</a>      <span class="comment">// end of slice range; or nil</span>
<span id="SliceExpr.Max"></span>    Max    <a href="#Expr">Expr</a>      <span class="comment">// maximum capacity of slice; or nil</span>
<span id="SliceExpr.Slice3"></span>    Slice3 <a href="/pkg/builtin/#bool">bool</a>      <span class="comment">// true if 3-index slice (2 colons present)</span>
<span id="SliceExpr.Rbrack"></span>    Rbrack <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;]&#34;</span>
}
</pre>











### <a id="SliceExpr.End">func</a> (\*SliceExpr) [End](https://golang.org/src/go/ast/ast.go?s=14385:14420#L460)
<pre>func (x *<a href="#SliceExpr">SliceExpr</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="SliceExpr.Pos">func</a> (\*SliceExpr) [Pos](https://golang.org/src/go/ast/ast.go?s=12781:12816#L427)
<pre>func (x *<a href="#SliceExpr">SliceExpr</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="Spec">type</a> [Spec](https://golang.org/src/go/ast/ast.go?s=25957:25996#L808)
The Spec type stands for any of *ImportSpec, *ValueSpec, and *TypeSpec.


<pre>type Spec interface {
    <a href="#Node">Node</a>
    <span class="comment">// contains filtered or unexported methods</span>
}</pre>











## <a id="StarExpr">type</a> [StarExpr](https://golang.org/src/go/ast/ast.go?s=9596:9680#L318)
A StarExpr node represents an expression of the form "*" Expression.
Semantically it could be a unary "*" expression, or a pointer type.


<pre>type StarExpr struct {
<span id="StarExpr.Star"></span>    Star <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;*&#34;</span>
<span id="StarExpr.X"></span>    X    <a href="#Expr">Expr</a>      <span class="comment">// operand</span>
}
</pre>











### <a id="StarExpr.End">func</a> (\*StarExpr) [End](https://golang.org/src/go/ast/ast.go?s=14580:14614#L463)
<pre>func (x *<a href="#StarExpr">StarExpr</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="StarExpr.Pos">func</a> (\*StarExpr) [Pos](https://golang.org/src/go/ast/ast.go?s=12969:13003#L430)
<pre>func (x *<a href="#StarExpr">StarExpr</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="Stmt">type</a> [Stmt](https://golang.org/src/go/ast/ast.go?s=1501:1542#L35)
All statement nodes implement the Stmt interface.


<pre>type Stmt interface {
    <a href="#Node">Node</a>
    <span class="comment">// contains filtered or unexported methods</span>
}</pre>











## <a id="StructType">type</a> [StructType](https://golang.org/src/go/ast/ast.go?s=10913:11127#L373)
A StructType node represents a struct type.


<pre>type StructType struct {
<span id="StructType.Struct"></span>    Struct     <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>  <span class="comment">// position of &#34;struct&#34; keyword</span>
<span id="StructType.Fields"></span>    Fields     *<a href="#FieldList">FieldList</a> <span class="comment">// list of field declarations</span>
<span id="StructType.Incomplete"></span>    Incomplete <a href="/pkg/builtin/#bool">bool</a>       <span class="comment">// true if (source) fields are missing in the Fields list</span>
}
</pre>











### <a id="StructType.End">func</a> (\*StructType) [End](https://golang.org/src/go/ast/ast.go?s=14896:14932#L468)
<pre>func (x *<a href="#StructType">StructType</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="StructType.Pos">func</a> (\*StructType) [Pos](https://golang.org/src/go/ast/ast.go?s=13275:13311#L435)
<pre>func (x *<a href="#StructType">StructType</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="SwitchStmt">type</a> [SwitchStmt](https://golang.org/src/go/ast/ast.go?s=20124:20340#L648)
A SwitchStmt node represents an expression switch statement.


<pre>type SwitchStmt struct {
<span id="SwitchStmt.Switch"></span>    Switch <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>  <span class="comment">// position of &#34;switch&#34; keyword</span>
<span id="SwitchStmt.Init"></span>    Init   <a href="#Stmt">Stmt</a>       <span class="comment">// initialization statement; or nil</span>
<span id="SwitchStmt.Tag"></span>    Tag    <a href="#Expr">Expr</a>       <span class="comment">// tag expression; or nil</span>
<span id="SwitchStmt.Body"></span>    Body   *<a href="#BlockStmt">BlockStmt</a> <span class="comment">// CaseClauses only</span>
}
</pre>











### <a id="SwitchStmt.End">func</a> (\*SwitchStmt) [End](https://golang.org/src/go/ast/ast.go?s=24372:24408#L763)
<pre>func (s *<a href="#SwitchStmt">SwitchStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="SwitchStmt.Pos">func</a> (\*SwitchStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=22753:22789#L714)
<pre>func (s *<a href="#SwitchStmt">SwitchStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="TypeAssertExpr">type</a> [TypeAssertExpr](https://golang.org/src/go/ast/ast.go?s=8891:9094#L299)
A TypeAssertExpr node represents an expression followed by a
type assertion.


<pre>type TypeAssertExpr struct {
<span id="TypeAssertExpr.X"></span>    X      <a href="#Expr">Expr</a>      <span class="comment">// expression</span>
<span id="TypeAssertExpr.Lparen"></span>    Lparen <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;(&#34;</span>
<span id="TypeAssertExpr.Type"></span>    Type   <a href="#Expr">Expr</a>      <span class="comment">// asserted type; nil means type switch X.(type)</span>
<span id="TypeAssertExpr.Rparen"></span>    Rparen <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of &#34;)&#34;</span>
}
</pre>











### <a id="TypeAssertExpr.End">func</a> (\*TypeAssertExpr) [End](https://golang.org/src/go/ast/ast.go?s=14450:14490#L461)
<pre>func (x *<a href="#TypeAssertExpr">TypeAssertExpr</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="TypeAssertExpr.Pos">func</a> (\*TypeAssertExpr) [Pos](https://golang.org/src/go/ast/ast.go?s=12843:12883#L428)
<pre>func (x *<a href="#TypeAssertExpr">TypeAssertExpr</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="TypeSpec">type</a> [TypeSpec](https://golang.org/src/go/ast/ast.go?s=26840:27154#L834)
A TypeSpec node represents a type declaration (TypeSpec production).


<pre>type TypeSpec struct {
<span id="TypeSpec.Doc"></span>    Doc     *<a href="#CommentGroup">CommentGroup</a> <span class="comment">// associated documentation; or nil</span>
<span id="TypeSpec.Name"></span>    Name    *<a href="#Ident">Ident</a>        <span class="comment">// type name</span>
<span id="TypeSpec.Assign"></span>    Assign  <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>     <span class="comment">// position of &#39;=&#39;, if any</span>
<span id="TypeSpec.Type"></span>    Type    <a href="#Expr">Expr</a>          <span class="comment">// *Ident, *ParenExpr, *SelectorExpr, *StarExpr, or any of the *XxxTypes</span>
<span id="TypeSpec.Comment"></span>    Comment *<a href="#CommentGroup">CommentGroup</a> <span class="comment">// line comments; or nil</span>
}
</pre>











### <a id="TypeSpec.End">func</a> (\*TypeSpec) [End](https://golang.org/src/go/ast/ast.go?s=27729:27763#L870)
<pre>func (s *<a href="#TypeSpec">TypeSpec</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="TypeSpec.Pos">func</a> (\*TypeSpec) [Pos](https://golang.org/src/go/ast/ast.go?s=27377:27411#L852)
<pre>func (s *<a href="#TypeSpec">TypeSpec</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="TypeSwitchStmt">type</a> [TypeSwitchStmt](https://golang.org/src/go/ast/ast.go?s=20406:20629#L656)
An TypeSwitchStmt node represents a type switch statement.


<pre>type TypeSwitchStmt struct {
<span id="TypeSwitchStmt.Switch"></span>    Switch <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>  <span class="comment">// position of &#34;switch&#34; keyword</span>
<span id="TypeSwitchStmt.Init"></span>    Init   <a href="#Stmt">Stmt</a>       <span class="comment">// initialization statement; or nil</span>
<span id="TypeSwitchStmt.Assign"></span>    Assign <a href="#Stmt">Stmt</a>       <span class="comment">// x := y.(type) or y.(type)</span>
<span id="TypeSwitchStmt.Body"></span>    Body   *<a href="#BlockStmt">BlockStmt</a> <span class="comment">// CaseClauses only</span>
}
</pre>











### <a id="TypeSwitchStmt.End">func</a> (\*TypeSwitchStmt) [End](https://golang.org/src/go/ast/ast.go?s=24437:24477#L764)
<pre>func (s *<a href="#TypeSwitchStmt">TypeSwitchStmt</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="TypeSwitchStmt.Pos">func</a> (\*TypeSwitchStmt) [Pos](https://golang.org/src/go/ast/ast.go?s=22814:22854#L715)
<pre>func (s *<a href="#TypeSwitchStmt">TypeSwitchStmt</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="UnaryExpr">type</a> [UnaryExpr](https://golang.org/src/go/ast/ast.go?s=9801:9923#L326)
A UnaryExpr node represents a unary expression.
Unary "*" expressions are represented via StarExpr nodes.


<pre>type UnaryExpr struct {
<span id="UnaryExpr.OpPos"></span>    OpPos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>   <span class="comment">// position of Op</span>
<span id="UnaryExpr.Op"></span>    Op    <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Token">Token</a> <span class="comment">// operator</span>
<span id="UnaryExpr.X"></span>    X     <a href="#Expr">Expr</a>        <span class="comment">// operand</span>
}
</pre>











### <a id="UnaryExpr.End">func</a> (\*UnaryExpr) [End](https://golang.org/src/go/ast/ast.go?s=14642:14677#L464)
<pre>func (x *<a href="#UnaryExpr">UnaryExpr</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="UnaryExpr.Pos">func</a> (\*UnaryExpr) [Pos](https://golang.org/src/go/ast/ast.go?s=13028:13063#L431)
<pre>func (x *<a href="#UnaryExpr">UnaryExpr</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="ValueSpec">type</a> [ValueSpec](https://golang.org/src/go/ast/ast.go?s=26482:26764#L825)
A ValueSpec node represents a constant or variable declaration
(ConstSpec or VarSpec production).


<pre>type ValueSpec struct {
<span id="ValueSpec.Doc"></span>    Doc     *<a href="#CommentGroup">CommentGroup</a> <span class="comment">// associated documentation; or nil</span>
<span id="ValueSpec.Names"></span>    Names   []*<a href="#Ident">Ident</a>      <span class="comment">// value names (len(Names) &gt; 0)</span>
<span id="ValueSpec.Type"></span>    Type    <a href="#Expr">Expr</a>          <span class="comment">// value type; or nil</span>
<span id="ValueSpec.Values"></span>    Values  []<a href="#Expr">Expr</a>        <span class="comment">// initial values; or nil</span>
<span id="ValueSpec.Comment"></span>    Comment *<a href="#CommentGroup">CommentGroup</a> <span class="comment">// line comments; or nil</span>
}
</pre>











### <a id="ValueSpec.End">func</a> (\*ValueSpec) [End](https://golang.org/src/go/ast/ast.go?s=27542:27577#L861)
<pre>func (s *<a href="#ValueSpec">ValueSpec</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="ValueSpec.Pos">func</a> (\*ValueSpec) [Pos](https://golang.org/src/go/ast/ast.go?s=27313:27348#L851)
<pre>func (s *<a href="#ValueSpec">ValueSpec</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



## <a id="Visitor">type</a> [Visitor](https://golang.org/src/go/ast/walk.go?s=400:456#L2)
A Visitor's Visit method is invoked for each node encountered by Walk.
If the result visitor w is not nil, Walk visits each of the children
of node with the visitor w, followed by a call of w.Visit(nil).


<pre>type Visitor interface {
    Visit(node <a href="#Node">Node</a>) (w <a href="#Visitor">Visitor</a>)
}</pre>














