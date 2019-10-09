

# types
`import "go/types"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package types declares the data types and implements
the algorithms for type-checking of Go packages. Use
Config.Check to invoke the type checker for a package.
Alternatively, create a new type checker with NewChecker
and invoke it incrementally by calling Checker.Files.

Type-checking consists of several interdependent phases:

Name resolution maps each identifier (ast.Ident) in the program to the
language object (Object) it denotes.
Use Info.{Defs,Uses,Implicits} for the results of name resolution.

Constant folding computes the exact constant value (constant.Value)
for every expression (ast.Expr) that is a compile-time constant.
Use Info.Types[expr].Value for the results of constant folding.

Type inference computes the type (Type) of every expression (ast.Expr)
and checks for compliance with the language specification.
Use Info.Types[expr].Type for the results of type inference.

For a tutorial, see <a href="https://golang.org/s/types-tutorial">https://golang.org/s/types-tutorial</a>.




## <a id="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func AssertableTo(V *Interface, T Type) bool](#AssertableTo)
* [func AssignableTo(V, T Type) bool](#AssignableTo)
* [func CheckExpr(fset *token.FileSet, pkg *Package, pos token.Pos, expr ast.Expr, info *Info) (err error)](#CheckExpr)
* [func Comparable(T Type) bool](#Comparable)
* [func ConvertibleTo(V, T Type) bool](#ConvertibleTo)
* [func DefPredeclaredTestFuncs()](#DefPredeclaredTestFuncs)
* [func ExprString(x ast.Expr) string](#ExprString)
* [func Id(pkg *Package, name string) string](#Id)
* [func Identical(x, y Type) bool](#Identical)
* [func IdenticalIgnoreTags(x, y Type) bool](#IdenticalIgnoreTags)
* [func Implements(V Type, T *Interface) bool](#Implements)
* [func IsInterface(typ Type) bool](#IsInterface)
* [func ObjectString(obj Object, qf Qualifier) string](#ObjectString)
* [func SelectionString(s *Selection, qf Qualifier) string](#SelectionString)
* [func TypeString(typ Type, qf Qualifier) string](#TypeString)
* [func WriteExpr(buf *bytes.Buffer, x ast.Expr)](#WriteExpr)
* [func WriteSignature(buf *bytes.Buffer, sig *Signature, qf Qualifier)](#WriteSignature)
* [func WriteType(buf *bytes.Buffer, typ Type, qf Qualifier)](#WriteType)
* [type Array](#Array)
  * [func NewArray(elem Type, len int64) *Array](#NewArray)
  * [func (a *Array) Elem() Type](#Array.Elem)
  * [func (a *Array) Len() int64](#Array.Len)
  * [func (a *Array) String() string](#Array.String)
  * [func (a *Array) Underlying() Type](#Array.Underlying)
* [type Basic](#Basic)
  * [func (b *Basic) Info() BasicInfo](#Basic.Info)
  * [func (b *Basic) Kind() BasicKind](#Basic.Kind)
  * [func (b *Basic) Name() string](#Basic.Name)
  * [func (b *Basic) String() string](#Basic.String)
  * [func (b *Basic) Underlying() Type](#Basic.Underlying)
* [type BasicInfo](#BasicInfo)
* [type BasicKind](#BasicKind)
* [type Builtin](#Builtin)
  * [func (obj *Builtin) Exported() bool](#Builtin.Exported)
  * [func (obj *Builtin) Id() string](#Builtin.Id)
  * [func (obj *Builtin) Name() string](#Builtin.Name)
  * [func (obj *Builtin) Parent() *Scope](#Builtin.Parent)
  * [func (obj *Builtin) Pkg() *Package](#Builtin.Pkg)
  * [func (obj *Builtin) Pos() token.Pos](#Builtin.Pos)
  * [func (obj *Builtin) String() string](#Builtin.String)
  * [func (obj *Builtin) Type() Type](#Builtin.Type)
* [type Chan](#Chan)
  * [func NewChan(dir ChanDir, elem Type) *Chan](#NewChan)
  * [func (c *Chan) Dir() ChanDir](#Chan.Dir)
  * [func (c *Chan) Elem() Type](#Chan.Elem)
  * [func (c *Chan) String() string](#Chan.String)
  * [func (c *Chan) Underlying() Type](#Chan.Underlying)
* [type ChanDir](#ChanDir)
* [type Checker](#Checker)
  * [func NewChecker(conf *Config, fset *token.FileSet, pkg *Package, info *Info) *Checker](#NewChecker)
  * [func (check *Checker) Files(files []*ast.File) error](#Checker.Files)
* [type Config](#Config)
  * [func (conf *Config) Check(path string, fset *token.FileSet, files []*ast.File, info *Info) (*Package, error)](#Config.Check)
* [type Const](#Const)
  * [func NewConst(pos token.Pos, pkg *Package, name string, typ Type, val constant.Value) *Const](#NewConst)
  * [func (obj *Const) Exported() bool](#Const.Exported)
  * [func (obj *Const) Id() string](#Const.Id)
  * [func (obj *Const) Name() string](#Const.Name)
  * [func (obj *Const) Parent() *Scope](#Const.Parent)
  * [func (obj *Const) Pkg() *Package](#Const.Pkg)
  * [func (obj *Const) Pos() token.Pos](#Const.Pos)
  * [func (obj *Const) String() string](#Const.String)
  * [func (obj *Const) Type() Type](#Const.Type)
  * [func (obj *Const) Val() constant.Value](#Const.Val)
* [type Error](#Error)
  * [func (err Error) Error() string](#Error.Error)
* [type Func](#Func)
  * [func MissingMethod(V Type, T *Interface, static bool) (method *Func, wrongType bool)](#MissingMethod)
  * [func NewFunc(pos token.Pos, pkg *Package, name string, sig *Signature) *Func](#NewFunc)
  * [func (obj *Func) Exported() bool](#Func.Exported)
  * [func (obj *Func) FullName() string](#Func.FullName)
  * [func (obj *Func) Id() string](#Func.Id)
  * [func (obj *Func) Name() string](#Func.Name)
  * [func (obj *Func) Parent() *Scope](#Func.Parent)
  * [func (obj *Func) Pkg() *Package](#Func.Pkg)
  * [func (obj *Func) Pos() token.Pos](#Func.Pos)
  * [func (obj *Func) Scope() *Scope](#Func.Scope)
  * [func (obj *Func) String() string](#Func.String)
  * [func (obj *Func) Type() Type](#Func.Type)
* [type ImportMode](#ImportMode)
* [type Importer](#Importer)
* [type ImporterFrom](#ImporterFrom)
* [type Info](#Info)
  * [func (info *Info) ObjectOf(id *ast.Ident) Object](#Info.ObjectOf)
  * [func (info *Info) TypeOf(e ast.Expr) Type](#Info.TypeOf)
* [type Initializer](#Initializer)
  * [func (init *Initializer) String() string](#Initializer.String)
* [type Interface](#Interface)
  * [func NewInterface(methods []*Func, embeddeds []*Named) *Interface](#NewInterface)
  * [func NewInterfaceType(methods []*Func, embeddeds []Type) *Interface](#NewInterfaceType)
  * [func (t *Interface) Complete() *Interface](#Interface.Complete)
  * [func (t *Interface) Embedded(i int) *Named](#Interface.Embedded)
  * [func (t *Interface) EmbeddedType(i int) Type](#Interface.EmbeddedType)
  * [func (t *Interface) Empty() bool](#Interface.Empty)
  * [func (t *Interface) ExplicitMethod(i int) *Func](#Interface.ExplicitMethod)
  * [func (t *Interface) Method(i int) *Func](#Interface.Method)
  * [func (t *Interface) NumEmbeddeds() int](#Interface.NumEmbeddeds)
  * [func (t *Interface) NumExplicitMethods() int](#Interface.NumExplicitMethods)
  * [func (t *Interface) NumMethods() int](#Interface.NumMethods)
  * [func (t *Interface) String() string](#Interface.String)
  * [func (t *Interface) Underlying() Type](#Interface.Underlying)
* [type Label](#Label)
  * [func NewLabel(pos token.Pos, pkg *Package, name string) *Label](#NewLabel)
  * [func (obj *Label) Exported() bool](#Label.Exported)
  * [func (obj *Label) Id() string](#Label.Id)
  * [func (obj *Label) Name() string](#Label.Name)
  * [func (obj *Label) Parent() *Scope](#Label.Parent)
  * [func (obj *Label) Pkg() *Package](#Label.Pkg)
  * [func (obj *Label) Pos() token.Pos](#Label.Pos)
  * [func (obj *Label) String() string](#Label.String)
  * [func (obj *Label) Type() Type](#Label.Type)
* [type Map](#Map)
  * [func NewMap(key, elem Type) *Map](#NewMap)
  * [func (m *Map) Elem() Type](#Map.Elem)
  * [func (m *Map) Key() Type](#Map.Key)
  * [func (m *Map) String() string](#Map.String)
  * [func (m *Map) Underlying() Type](#Map.Underlying)
* [type MethodSet](#MethodSet)
  * [func NewMethodSet(T Type) *MethodSet](#NewMethodSet)
  * [func (s *MethodSet) At(i int) *Selection](#MethodSet.At)
  * [func (s *MethodSet) Len() int](#MethodSet.Len)
  * [func (s *MethodSet) Lookup(pkg *Package, name string) *Selection](#MethodSet.Lookup)
  * [func (s *MethodSet) String() string](#MethodSet.String)
* [type Named](#Named)
  * [func NewNamed(obj *TypeName, underlying Type, methods []*Func) *Named](#NewNamed)
  * [func (t *Named) AddMethod(m *Func)](#Named.AddMethod)
  * [func (t *Named) Method(i int) *Func](#Named.Method)
  * [func (t *Named) NumMethods() int](#Named.NumMethods)
  * [func (t *Named) Obj() *TypeName](#Named.Obj)
  * [func (t *Named) SetUnderlying(underlying Type)](#Named.SetUnderlying)
  * [func (t *Named) String() string](#Named.String)
  * [func (t *Named) Underlying() Type](#Named.Underlying)
* [type Nil](#Nil)
  * [func (obj *Nil) Exported() bool](#Nil.Exported)
  * [func (obj *Nil) Id() string](#Nil.Id)
  * [func (obj *Nil) Name() string](#Nil.Name)
  * [func (obj *Nil) Parent() *Scope](#Nil.Parent)
  * [func (obj *Nil) Pkg() *Package](#Nil.Pkg)
  * [func (obj *Nil) Pos() token.Pos](#Nil.Pos)
  * [func (obj *Nil) String() string](#Nil.String)
  * [func (obj *Nil) Type() Type](#Nil.Type)
* [type Object](#Object)
  * [func LookupFieldOrMethod(T Type, addressable bool, pkg *Package, name string) (obj Object, index []int, indirect bool)](#LookupFieldOrMethod)
* [type Package](#Package)
  * [func NewPackage(path, name string) *Package](#NewPackage)
  * [func (pkg *Package) Complete() bool](#Package.Complete)
  * [func (pkg *Package) Imports() []*Package](#Package.Imports)
  * [func (pkg *Package) MarkComplete()](#Package.MarkComplete)
  * [func (pkg *Package) Name() string](#Package.Name)
  * [func (pkg *Package) Path() string](#Package.Path)
  * [func (pkg *Package) Scope() *Scope](#Package.Scope)
  * [func (pkg *Package) SetImports(list []*Package)](#Package.SetImports)
  * [func (pkg *Package) SetName(name string)](#Package.SetName)
  * [func (pkg *Package) String() string](#Package.String)
* [type PkgName](#PkgName)
  * [func NewPkgName(pos token.Pos, pkg *Package, name string, imported *Package) *PkgName](#NewPkgName)
  * [func (obj *PkgName) Exported() bool](#PkgName.Exported)
  * [func (obj *PkgName) Id() string](#PkgName.Id)
  * [func (obj *PkgName) Imported() *Package](#PkgName.Imported)
  * [func (obj *PkgName) Name() string](#PkgName.Name)
  * [func (obj *PkgName) Parent() *Scope](#PkgName.Parent)
  * [func (obj *PkgName) Pkg() *Package](#PkgName.Pkg)
  * [func (obj *PkgName) Pos() token.Pos](#PkgName.Pos)
  * [func (obj *PkgName) String() string](#PkgName.String)
  * [func (obj *PkgName) Type() Type](#PkgName.Type)
* [type Pointer](#Pointer)
  * [func NewPointer(elem Type) *Pointer](#NewPointer)
  * [func (p *Pointer) Elem() Type](#Pointer.Elem)
  * [func (p *Pointer) String() string](#Pointer.String)
  * [func (p *Pointer) Underlying() Type](#Pointer.Underlying)
* [type Qualifier](#Qualifier)
  * [func RelativeTo(pkg *Package) Qualifier](#RelativeTo)
* [type Scope](#Scope)
  * [func NewScope(parent *Scope, pos, end token.Pos, comment string) *Scope](#NewScope)
  * [func (s *Scope) Child(i int) *Scope](#Scope.Child)
  * [func (s *Scope) Contains(pos token.Pos) bool](#Scope.Contains)
  * [func (s *Scope) End() token.Pos](#Scope.End)
  * [func (s *Scope) Innermost(pos token.Pos) *Scope](#Scope.Innermost)
  * [func (s *Scope) Insert(obj Object) Object](#Scope.Insert)
  * [func (s *Scope) Len() int](#Scope.Len)
  * [func (s *Scope) Lookup(name string) Object](#Scope.Lookup)
  * [func (s *Scope) LookupParent(name string, pos token.Pos) (*Scope, Object)](#Scope.LookupParent)
  * [func (s *Scope) Names() []string](#Scope.Names)
  * [func (s *Scope) NumChildren() int](#Scope.NumChildren)
  * [func (s *Scope) Parent() *Scope](#Scope.Parent)
  * [func (s *Scope) Pos() token.Pos](#Scope.Pos)
  * [func (s *Scope) String() string](#Scope.String)
  * [func (s *Scope) WriteTo(w io.Writer, n int, recurse bool)](#Scope.WriteTo)
* [type Selection](#Selection)
  * [func (s *Selection) Index() []int](#Selection.Index)
  * [func (s *Selection) Indirect() bool](#Selection.Indirect)
  * [func (s *Selection) Kind() SelectionKind](#Selection.Kind)
  * [func (s *Selection) Obj() Object](#Selection.Obj)
  * [func (s *Selection) Recv() Type](#Selection.Recv)
  * [func (s *Selection) String() string](#Selection.String)
  * [func (s *Selection) Type() Type](#Selection.Type)
* [type SelectionKind](#SelectionKind)
* [type Signature](#Signature)
  * [func NewSignature(recv *Var, params, results *Tuple, variadic bool) *Signature](#NewSignature)
  * [func (s *Signature) Params() *Tuple](#Signature.Params)
  * [func (s *Signature) Recv() *Var](#Signature.Recv)
  * [func (s *Signature) Results() *Tuple](#Signature.Results)
  * [func (s *Signature) String() string](#Signature.String)
  * [func (s *Signature) Underlying() Type](#Signature.Underlying)
  * [func (s *Signature) Variadic() bool](#Signature.Variadic)
* [type Sizes](#Sizes)
  * [func SizesFor(compiler, arch string) Sizes](#SizesFor)
* [type Slice](#Slice)
  * [func NewSlice(elem Type) *Slice](#NewSlice)
  * [func (s *Slice) Elem() Type](#Slice.Elem)
  * [func (s *Slice) String() string](#Slice.String)
  * [func (s *Slice) Underlying() Type](#Slice.Underlying)
* [type StdSizes](#StdSizes)
  * [func (s *StdSizes) Alignof(T Type) int64](#StdSizes.Alignof)
  * [func (s *StdSizes) Offsetsof(fields []*Var) []int64](#StdSizes.Offsetsof)
  * [func (s *StdSizes) Sizeof(T Type) int64](#StdSizes.Sizeof)
* [type Struct](#Struct)
  * [func NewStruct(fields []*Var, tags []string) *Struct](#NewStruct)
  * [func (s *Struct) Field(i int) *Var](#Struct.Field)
  * [func (s *Struct) NumFields() int](#Struct.NumFields)
  * [func (s *Struct) String() string](#Struct.String)
  * [func (s *Struct) Tag(i int) string](#Struct.Tag)
  * [func (s *Struct) Underlying() Type](#Struct.Underlying)
* [type Tuple](#Tuple)
  * [func NewTuple(x ...*Var) *Tuple](#NewTuple)
  * [func (t *Tuple) At(i int) *Var](#Tuple.At)
  * [func (t *Tuple) Len() int](#Tuple.Len)
  * [func (t *Tuple) String() string](#Tuple.String)
  * [func (t *Tuple) Underlying() Type](#Tuple.Underlying)
* [type Type](#Type)
  * [func Default(typ Type) Type](#Default)
* [type TypeAndValue](#TypeAndValue)
  * [func Eval(fset *token.FileSet, pkg *Package, pos token.Pos, expr string) (_ TypeAndValue, err error)](#Eval)
  * [func (tv TypeAndValue) Addressable() bool](#TypeAndValue.Addressable)
  * [func (tv TypeAndValue) Assignable() bool](#TypeAndValue.Assignable)
  * [func (tv TypeAndValue) HasOk() bool](#TypeAndValue.HasOk)
  * [func (tv TypeAndValue) IsBuiltin() bool](#TypeAndValue.IsBuiltin)
  * [func (tv TypeAndValue) IsNil() bool](#TypeAndValue.IsNil)
  * [func (tv TypeAndValue) IsType() bool](#TypeAndValue.IsType)
  * [func (tv TypeAndValue) IsValue() bool](#TypeAndValue.IsValue)
  * [func (tv TypeAndValue) IsVoid() bool](#TypeAndValue.IsVoid)
* [type TypeName](#TypeName)
  * [func NewTypeName(pos token.Pos, pkg *Package, name string, typ Type) *TypeName](#NewTypeName)
  * [func (obj *TypeName) Exported() bool](#TypeName.Exported)
  * [func (obj *TypeName) Id() string](#TypeName.Id)
  * [func (obj *TypeName) IsAlias() bool](#TypeName.IsAlias)
  * [func (obj *TypeName) Name() string](#TypeName.Name)
  * [func (obj *TypeName) Parent() *Scope](#TypeName.Parent)
  * [func (obj *TypeName) Pkg() *Package](#TypeName.Pkg)
  * [func (obj *TypeName) Pos() token.Pos](#TypeName.Pos)
  * [func (obj *TypeName) String() string](#TypeName.String)
  * [func (obj *TypeName) Type() Type](#TypeName.Type)
* [type Var](#Var)
  * [func NewField(pos token.Pos, pkg *Package, name string, typ Type, embedded bool) *Var](#NewField)
  * [func NewParam(pos token.Pos, pkg *Package, name string, typ Type) *Var](#NewParam)
  * [func NewVar(pos token.Pos, pkg *Package, name string, typ Type) *Var](#NewVar)
  * [func (obj *Var) Anonymous() bool](#Var.Anonymous)
  * [func (obj *Var) Embedded() bool](#Var.Embedded)
  * [func (obj *Var) Exported() bool](#Var.Exported)
  * [func (obj *Var) Id() string](#Var.Id)
  * [func (obj *Var) IsField() bool](#Var.IsField)
  * [func (obj *Var) Name() string](#Var.Name)
  * [func (obj *Var) Parent() *Scope](#Var.Parent)
  * [func (obj *Var) Pkg() *Package](#Var.Pkg)
  * [func (obj *Var) Pos() token.Pos](#Var.Pos)
  * [func (obj *Var) String() string](#Var.String)
  * [func (obj *Var) Type() Type](#Var.Type)


#### <a id="pkg-examples">Examples</a>
* [Info](#example_Info)
* [MethodSet](#example_MethodSet)
* [Scope](#example_Scope)


#### <a id="pkg-files">Package files</a>
[api.go](https://golang.org/src/go/types/api.go) [assignments.go](https://golang.org/src/go/types/assignments.go) [builtins.go](https://golang.org/src/go/types/builtins.go) [call.go](https://golang.org/src/go/types/call.go) [check.go](https://golang.org/src/go/types/check.go) [conversions.go](https://golang.org/src/go/types/conversions.go) [decl.go](https://golang.org/src/go/types/decl.go) [errors.go](https://golang.org/src/go/types/errors.go) [eval.go](https://golang.org/src/go/types/eval.go) [expr.go](https://golang.org/src/go/types/expr.go) [exprstring.go](https://golang.org/src/go/types/exprstring.go) [gccgosizes.go](https://golang.org/src/go/types/gccgosizes.go) [initorder.go](https://golang.org/src/go/types/initorder.go) [interfaces.go](https://golang.org/src/go/types/interfaces.go) [labels.go](https://golang.org/src/go/types/labels.go) [lookup.go](https://golang.org/src/go/types/lookup.go) [methodset.go](https://golang.org/src/go/types/methodset.go) [object.go](https://golang.org/src/go/types/object.go) [objset.go](https://golang.org/src/go/types/objset.go) [operand.go](https://golang.org/src/go/types/operand.go) [package.go](https://golang.org/src/go/types/package.go) [predicates.go](https://golang.org/src/go/types/predicates.go) [resolver.go](https://golang.org/src/go/types/resolver.go) [return.go](https://golang.org/src/go/types/return.go) [scope.go](https://golang.org/src/go/types/scope.go) [selection.go](https://golang.org/src/go/types/selection.go) [sizes.go](https://golang.org/src/go/types/sizes.go) [stmt.go](https://golang.org/src/go/types/stmt.go) [type.go](https://golang.org/src/go/types/type.go) [typestring.go](https://golang.org/src/go/types/typestring.go) [typexpr.go](https://golang.org/src/go/types/typexpr.go) [universe.go](https://golang.org/src/go/types/universe.go) 




## <a id="pkg-variables">Variables</a>
Typ contains the predeclared *Basic types indexed by their
corresponding BasicKind.

The *Basic type for Typ[Byte] will have the name "uint8".
Use Universe.Lookup("byte").Type() to obtain the specific
alias basic type named "byte" (and analogous for "rune").


<pre>var <span id="Typ">Typ</span> = []*<a href="#Basic">Basic</a>{
    <a href="#Invalid">Invalid</a>: {<a href="#Invalid">Invalid</a>, 0, &#34;invalid type&#34;},

    <a href="#Bool">Bool</a>:          {<a href="#Bool">Bool</a>, <a href="#IsBoolean">IsBoolean</a>, &#34;bool&#34;},
    <a href="#Int">Int</a>:           {<a href="#Int">Int</a>, <a href="#IsInteger">IsInteger</a>, &#34;int&#34;},
    <a href="#Int8">Int8</a>:          {<a href="#Int8">Int8</a>, <a href="#IsInteger">IsInteger</a>, &#34;int8&#34;},
    <a href="#Int16">Int16</a>:         {<a href="#Int16">Int16</a>, <a href="#IsInteger">IsInteger</a>, &#34;int16&#34;},
    <a href="#Int32">Int32</a>:         {<a href="#Int32">Int32</a>, <a href="#IsInteger">IsInteger</a>, &#34;int32&#34;},
    <a href="#Int64">Int64</a>:         {<a href="#Int64">Int64</a>, <a href="#IsInteger">IsInteger</a>, &#34;int64&#34;},
    <a href="#Uint">Uint</a>:          {<a href="#Uint">Uint</a>, <a href="#IsInteger">IsInteger</a> | <a href="#IsUnsigned">IsUnsigned</a>, &#34;uint&#34;},
    <a href="#Uint8">Uint8</a>:         {<a href="#Uint8">Uint8</a>, <a href="#IsInteger">IsInteger</a> | <a href="#IsUnsigned">IsUnsigned</a>, &#34;uint8&#34;},
    <a href="#Uint16">Uint16</a>:        {<a href="#Uint16">Uint16</a>, <a href="#IsInteger">IsInteger</a> | <a href="#IsUnsigned">IsUnsigned</a>, &#34;uint16&#34;},
    <a href="#Uint32">Uint32</a>:        {<a href="#Uint32">Uint32</a>, <a href="#IsInteger">IsInteger</a> | <a href="#IsUnsigned">IsUnsigned</a>, &#34;uint32&#34;},
    <a href="#Uint64">Uint64</a>:        {<a href="#Uint64">Uint64</a>, <a href="#IsInteger">IsInteger</a> | <a href="#IsUnsigned">IsUnsigned</a>, &#34;uint64&#34;},
    <a href="#Uintptr">Uintptr</a>:       {<a href="#Uintptr">Uintptr</a>, <a href="#IsInteger">IsInteger</a> | <a href="#IsUnsigned">IsUnsigned</a>, &#34;uintptr&#34;},
    <a href="#Float32">Float32</a>:       {<a href="#Float32">Float32</a>, <a href="#IsFloat">IsFloat</a>, &#34;float32&#34;},
    <a href="#Float64">Float64</a>:       {<a href="#Float64">Float64</a>, <a href="#IsFloat">IsFloat</a>, &#34;float64&#34;},
    <a href="#Complex64">Complex64</a>:     {<a href="#Complex64">Complex64</a>, <a href="#IsComplex">IsComplex</a>, &#34;complex64&#34;},
    <a href="#Complex128">Complex128</a>:    {<a href="#Complex128">Complex128</a>, <a href="#IsComplex">IsComplex</a>, &#34;complex128&#34;},
    <a href="#String">String</a>:        {<a href="#String">String</a>, <a href="#IsString">IsString</a>, &#34;string&#34;},
    <a href="#UnsafePointer">UnsafePointer</a>: {<a href="#UnsafePointer">UnsafePointer</a>, 0, &#34;Pointer&#34;},

    <a href="#UntypedBool">UntypedBool</a>:    {<a href="#UntypedBool">UntypedBool</a>, <a href="#IsBoolean">IsBoolean</a> | <a href="#IsUntyped">IsUntyped</a>, &#34;untyped bool&#34;},
    <a href="#UntypedInt">UntypedInt</a>:     {<a href="#UntypedInt">UntypedInt</a>, <a href="#IsInteger">IsInteger</a> | <a href="#IsUntyped">IsUntyped</a>, &#34;untyped int&#34;},
    <a href="#UntypedRune">UntypedRune</a>:    {<a href="#UntypedRune">UntypedRune</a>, <a href="#IsInteger">IsInteger</a> | <a href="#IsUntyped">IsUntyped</a>, &#34;untyped rune&#34;},
    <a href="#UntypedFloat">UntypedFloat</a>:   {<a href="#UntypedFloat">UntypedFloat</a>, <a href="#IsFloat">IsFloat</a> | <a href="#IsUntyped">IsUntyped</a>, &#34;untyped float&#34;},
    <a href="#UntypedComplex">UntypedComplex</a>: {<a href="#UntypedComplex">UntypedComplex</a>, <a href="#IsComplex">IsComplex</a> | <a href="#IsUntyped">IsUntyped</a>, &#34;untyped complex&#34;},
    <a href="#UntypedString">UntypedString</a>:  {<a href="#UntypedString">UntypedString</a>, <a href="#IsString">IsString</a> | <a href="#IsUntyped">IsUntyped</a>, &#34;untyped string&#34;},
    <a href="#UntypedNil">UntypedNil</a>:     {<a href="#UntypedNil">UntypedNil</a>, <a href="#IsUntyped">IsUntyped</a>, &#34;untyped nil&#34;},
}</pre>

## <a id="AssertableTo">func</a> [AssertableTo](https://golang.org/src/go/types/api.go?s=12911:12955#L345)
<pre>func AssertableTo(V *<a href="#Interface">Interface</a>, T <a href="#Type">Type</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
AssertableTo reports whether a value of type V can be asserted to have type T.



## <a id="AssignableTo">func</a> [AssignableTo](https://golang.org/src/go/types/api.go?s=13111:13144#L351)
<pre>func AssignableTo(V, T <a href="#Type">Type</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
AssignableTo reports whether a value of type V is assignable to a variable of type T.



## <a id="CheckExpr">func</a> [CheckExpr](https://golang.org/src/go/types/eval.go?s=1880:1983#L46)
<pre>func CheckExpr(fset *<a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#FileSet">FileSet</a>, pkg *<a href="#Package">Package</a>, pos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>, expr <a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#Expr">Expr</a>, info *<a href="#Info">Info</a>) (err <a href="/pkg/builtin/#error">error</a>)</pre>
CheckExpr type checks the expression expr as if it had appeared at
position pos of package pkg. Type information about the expression
is recorded in info.

If pkg == nil, the Universe scope is used and the provided
position pos is ignored. If pkg != nil, and pos is invalid,
the package scope is used. Otherwise, pos must belong to the
package.

An error is returned if pos is not within the package or
if the node cannot be type-checked.

Note: Eval and CheckExpr should not be used instead of running Check
to compute types and values, but in addition to Check, as these
functions ignore the context in which an expression is used (e.g., an
assignment). Thus, top-level untyped constants will return an
untyped type rather then the respective context-specific type.



## <a id="Comparable">func</a> [Comparable](https://golang.org/src/go/types/predicates.go?s=1744:1772#L71)
<pre>func Comparable(T <a href="#Type">Type</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Comparable reports whether values of type T are comparable.



## <a id="ConvertibleTo">func</a> [ConvertibleTo](https://golang.org/src/go/types/api.go?s=13348:13382#L357)
<pre>func ConvertibleTo(V, T <a href="#Type">Type</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
ConvertibleTo reports whether a value of type V is convertible to a value of type T.



## <a id="DefPredeclaredTestFuncs">func</a> [DefPredeclaredTestFuncs](https://golang.org/src/go/types/universe.go?s=5251:5281#L172)
<pre>func DefPredeclaredTestFuncs()</pre>
DefPredeclaredTestFuncs defines the assert and trace built-ins.
These built-ins are intended for debugging and testing of this
package only.



## <a id="ExprString">func</a> [ExprString](https://golang.org/src/go/types/exprstring.go?s=439:473#L7)
<pre>func ExprString(x <a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#Expr">Expr</a>) <a href="/pkg/builtin/#string">string</a></pre>
ExprString returns the (possibly shortened) string representation for x.
Shortened representations are suitable for user interfaces but may not
necessarily follow Go syntax.



## <a id="Id">func</a> [Id](https://golang.org/src/go/types/object.go?s=2004:2045#L50)
<pre>func Id(pkg *<a href="#Package">Package</a>, name <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Id returns name if it is exported, otherwise it
returns the name qualified with the package path.



## <a id="Identical">func</a> [Identical](https://golang.org/src/go/types/predicates.go?s=2524:2554#L105)
<pre>func Identical(x, y <a href="#Type">Type</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Identical reports whether x and y are identical types.
Receivers of Signature types are ignored.



## <a id="IdenticalIgnoreTags">func</a> [IdenticalIgnoreTags](https://golang.org/src/go/types/predicates.go?s=2728:2768#L111)
<pre>func IdenticalIgnoreTags(x, y <a href="#Type">Type</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IdenticalIgnoreTags reports whether x and y are identical types if tags are ignored.
Receivers of Signature types are ignored.



## <a id="Implements">func</a> [Implements](https://golang.org/src/go/types/api.go?s=13555:13597#L363)
<pre>func Implements(V <a href="#Type">Type</a>, T *<a href="#Interface">Interface</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Implements reports whether type V implements interface T.



## <a id="IsInterface">func</a> [IsInterface](https://golang.org/src/go/types/predicates.go?s=1593:1624#L65)
<pre>func IsInterface(typ <a href="#Type">Type</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsInterface reports whether typ is an interface type.



## <a id="ObjectString">func</a> [ObjectString](https://golang.org/src/go/types/object.go?s=13740:13790#L440)
<pre>func ObjectString(obj <a href="#Object">Object</a>, qf <a href="#Qualifier">Qualifier</a>) <a href="/pkg/builtin/#string">string</a></pre>
ObjectString returns the string form of obj.
The Qualifier controls the printing of
package-level objects, and may be nil.



## <a id="SelectionString">func</a> [SelectionString](https://golang.org/src/go/types/selection.go?s=3594:3649#L109)
<pre>func SelectionString(s *<a href="#Selection">Selection</a>, qf <a href="#Qualifier">Qualifier</a>) <a href="/pkg/builtin/#string">string</a></pre>
SelectionString returns the string form of s.
The Qualifier controls the printing of
package-level objects, and may be nil.

Examples:


	"field (T) f int"
	"method (T) f(X) Y"
	"method expr (T) f(X) Y"



## <a id="TypeString">func</a> [TypeString](https://golang.org/src/go/types/typestring.go?s=2151:2197#L55)
<pre>func TypeString(typ <a href="#Type">Type</a>, qf <a href="#Qualifier">Qualifier</a>) <a href="/pkg/builtin/#string">string</a></pre>
TypeString returns the string representation of typ.
The Qualifier controls the printing of
package-level objects, and may be nil.



## <a id="WriteExpr">func</a> [WriteExpr](https://golang.org/src/go/types/exprstring.go?s=730:775#L16)
<pre>func WriteExpr(buf *<a href="/pkg/bytes/">bytes</a>.<a href="/pkg/bytes/#Buffer">Buffer</a>, x <a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#Expr">Expr</a>)</pre>
WriteExpr writes the (possibly shortened) string representation for x to buf.
Shortened representations are suitable for user interfaces but may not
necessarily follow Go syntax.



## <a id="WriteSignature">func</a> [WriteSignature](https://golang.org/src/go/types/typestring.go?s=7408:7476#L275)
<pre>func WriteSignature(buf *<a href="/pkg/bytes/">bytes</a>.<a href="/pkg/bytes/#Buffer">Buffer</a>, sig *<a href="#Signature">Signature</a>, qf <a href="#Qualifier">Qualifier</a>)</pre>
WriteSignature writes the representation of the signature sig to buf,
without a leading "func" keyword.
The Qualifier controls the printing of
package-level objects, and may be nil.



## <a id="WriteType">func</a> [WriteType](https://golang.org/src/go/types/typestring.go?s=2417:2474#L64)
<pre>func WriteType(buf *<a href="/pkg/bytes/">bytes</a>.<a href="/pkg/bytes/#Buffer">Buffer</a>, typ <a href="#Type">Type</a>, qf <a href="#Qualifier">Qualifier</a>)</pre>
WriteType writes the string representation of typ to buf.
The Qualifier controls the printing of
package-level objects, and may be nil.





## <a id="Array">type</a> [Array](https://golang.org/src/go/types/type.go?s=1709:1753#L84)
An Array represents an array type.


<pre>type Array struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewArray">func</a> [NewArray](https://golang.org/src/go/types/type.go?s=1881:1923#L91)
<pre>func NewArray(elem <a href="#Type">Type</a>, len <a href="/pkg/builtin/#int64">int64</a>) *<a href="#Array">Array</a></pre>
NewArray returns a new array type for the given element type and length.
A negative length indicates an unknown length.






### <a id="Array.Elem">func</a> (\*Array) [Elem](https://golang.org/src/go/types/type.go?s=2129:2156#L98)
<pre>func (a *<a href="#Array">Array</a>) Elem() <a href="#Type">Type</a></pre>
Elem returns element type of array a.




### <a id="Array.Len">func</a> (\*Array) [Len](https://golang.org/src/go/types/type.go?s=2042:2069#L95)
<pre>func (a *<a href="#Array">Array</a>) Len() <a href="/pkg/builtin/#int64">int64</a></pre>
Len returns the length of array a.
A negative result indicates an unknown length.




### <a id="Array.String">func</a> (\*Array) [String](https://golang.org/src/go/types/type.go?s=15597:15628#L471)
<pre>func (a *<a href="#Array">Array</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Array.Underlying">func</a> (\*Array) [Underlying](https://golang.org/src/go/types/type.go?s=15009:15042#L459)
<pre>func (a *<a href="#Array">Array</a>) Underlying() <a href="#Type">Type</a></pre>



## <a id="Basic">type</a> [Basic](https://golang.org/src/go/types/type.go?s=1304:1370#L68)
A Basic represents a basic type.


<pre>type Basic struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="Basic.Info">func</a> (\*Basic) [Info](https://golang.org/src/go/types/type.go?s=1528:1560#L78)
<pre>func (b *<a href="#Basic">Basic</a>) Info() <a href="#BasicInfo">BasicInfo</a></pre>
Info returns information about properties of basic type b.




### <a id="Basic.Kind">func</a> (\*Basic) [Kind](https://golang.org/src/go/types/type.go?s=1414:1446#L75)
<pre>func (b *<a href="#Basic">Basic</a>) Kind() <a href="#BasicKind">BasicKind</a></pre>
Kind returns the kind of basic type b.




### <a id="Basic.Name">func</a> (\*Basic) [Name](https://golang.org/src/go/types/type.go?s=1622:1651#L81)
<pre>func (b *<a href="#Basic">Basic</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the name of basic type b.




### <a id="Basic.String">func</a> (\*Basic) [String](https://golang.org/src/go/types/type.go?s=15531:15562#L470)
<pre>func (b *<a href="#Basic">Basic</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Basic.Underlying">func</a> (\*Basic) [Underlying](https://golang.org/src/go/types/type.go?s=14958:14991#L458)
<pre>func (b *<a href="#Basic">Basic</a>) Underlying() <a href="#Type">Type</a></pre>



## <a id="BasicInfo">type</a> [BasicInfo](https://golang.org/src/go/types/type.go?s=968:986#L50)
BasicInfo is a set of flags describing properties of a basic type.


<pre>type BasicInfo <a href="/pkg/builtin/#int">int</a></pre>


Properties of basic types.


<pre>const (
    <span id="IsBoolean">IsBoolean</span> <a href="#BasicInfo">BasicInfo</a> = 1 &lt;&lt; <a href="/pkg/builtin/#iota">iota</a>
    <span id="IsInteger">IsInteger</span>
    <span id="IsUnsigned">IsUnsigned</span>
    <span id="IsFloat">IsFloat</span>
    <span id="IsComplex">IsComplex</span>
    <span id="IsString">IsString</span>
    <span id="IsUntyped">IsUntyped</span>

    <span id="IsOrdered">IsOrdered</span>   = <a href="#IsInteger">IsInteger</a> | <a href="#IsFloat">IsFloat</a> | <a href="#IsString">IsString</a>
    <span id="IsNumeric">IsNumeric</span>   = <a href="#IsInteger">IsInteger</a> | <a href="#IsFloat">IsFloat</a> | <a href="#IsComplex">IsComplex</a>
    <span id="IsConstType">IsConstType</span> = <a href="#IsBoolean">IsBoolean</a> | <a href="#IsNumeric">IsNumeric</a> | <a href="#IsString">IsString</a>
)</pre>









## <a id="BasicKind">type</a> [BasicKind](https://golang.org/src/go/types/type.go?s=485:503#L10)
BasicKind describes the kind of basic type.


<pre>type BasicKind <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="Invalid">Invalid</span> <a href="#BasicKind">BasicKind</a> = <a href="/pkg/builtin/#iota">iota</a> <span class="comment">// type is invalid</span>

    <span class="comment">// predeclared types</span>
    <span id="Bool">Bool</span>
    <span id="Int">Int</span>
    <span id="Int8">Int8</span>
    <span id="Int16">Int16</span>
    <span id="Int32">Int32</span>
    <span id="Int64">Int64</span>
    <span id="Uint">Uint</span>
    <span id="Uint8">Uint8</span>
    <span id="Uint16">Uint16</span>
    <span id="Uint32">Uint32</span>
    <span id="Uint64">Uint64</span>
    <span id="Uintptr">Uintptr</span>
    <span id="Float32">Float32</span>
    <span id="Float64">Float64</span>
    <span id="Complex64">Complex64</span>
    <span id="Complex128">Complex128</span>
    <span id="String">String</span>
    <span id="UnsafePointer">UnsafePointer</span>

    <span class="comment">// types for untyped values</span>
    <span id="UntypedBool">UntypedBool</span>
    <span id="UntypedInt">UntypedInt</span>
    <span id="UntypedRune">UntypedRune</span>
    <span id="UntypedFloat">UntypedFloat</span>
    <span id="UntypedComplex">UntypedComplex</span>
    <span id="UntypedString">UntypedString</span>
    <span id="UntypedNil">UntypedNil</span>

    <span class="comment">// aliases</span>
    <span id="Byte">Byte</span> = <a href="#Uint8">Uint8</a>
    <span id="Rune">Rune</span> = <a href="#Int32">Int32</a>
)</pre>









## <a id="Builtin">type</a> [Builtin](https://golang.org/src/go/types/object.go?s=11624:11669#L327)
A Builtin represents a built-in function.
Builtins don't have a valid type.


<pre>type Builtin struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="Builtin.Exported">func</a> (\*Builtin) [Exported](https://golang.org/src/go/types/object.go?s=4259:4293#L131)
<pre>func (obj *<a href="#Builtin">Builtin</a>) Exported() <a href="/pkg/builtin/#bool">bool</a></pre>
Exported reports whether the object is exported (starts with a capital letter).
It doesn't take into account whether the object is in a local (function) scope
or not.




### <a id="Builtin.Id">func</a> (\*Builtin) [Id](https://golang.org/src/go/types/object.go?s=4383:4413#L134)
<pre>func (obj *<a href="#Builtin">Builtin</a>) Id() <a href="/pkg/builtin/#string">string</a></pre>
Id is a wrapper for Id(obj.Pkg(), obj.Name()).




### <a id="Builtin.Name">func</a> (\*Builtin) [Name](https://golang.org/src/go/types/object.go?s=3943:3975#L123)
<pre>func (obj *<a href="#Builtin">Builtin</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the object's (package-local, unqualified) name.




### <a id="Builtin.Parent">func</a> (\*Builtin) [Parent](https://golang.org/src/go/types/object.go?s=3521:3555#L113)
<pre>func (obj *<a href="#Builtin">Builtin</a>) Parent() *<a href="#Scope">Scope</a></pre>
Parent returns the scope in which the object is declared.
The result is nil for methods and struct fields.




### <a id="Builtin.Pkg">func</a> (\*Builtin) [Pkg](https://golang.org/src/go/types/object.go?s=3825:3858#L120)
<pre>func (obj *<a href="#Builtin">Builtin</a>) Pkg() *<a href="#Package">Package</a></pre>
Pkg returns the package to which the object belongs.
The result is nil for labels and objects in the Universe scope.




### <a id="Builtin.Pos">func</a> (\*Builtin) [Pos](https://golang.org/src/go/types/object.go?s=3647:3681#L116)
<pre>func (obj *<a href="#Builtin">Builtin</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>
Pos returns the declaration position of the object's identifier.




### <a id="Builtin.String">func</a> (\*Builtin) [String](https://golang.org/src/go/types/object.go?s=14293:14328#L452)
<pre>func (obj *<a href="#Builtin">Builtin</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Builtin.Type">func</a> (\*Builtin) [Type](https://golang.org/src/go/types/object.go?s=4032:4062#L126)
<pre>func (obj *<a href="#Builtin">Builtin</a>) Type() <a href="#Type">Type</a></pre>
Type returns the object's type.




## <a id="Chan">type</a> [Chan](https://golang.org/src/go/types/type.go?s=12513:12558#L382)
A Chan represents a channel type.


<pre>type Chan struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewChan">func</a> [NewChan](https://golang.org/src/go/types/type.go?s=12834:12876#L398)
<pre>func NewChan(dir <a href="#ChanDir">ChanDir</a>, elem <a href="#Type">Type</a>) *<a href="#Chan">Chan</a></pre>
NewChan returns a new channel type for the given direction and element type.






### <a id="Chan.Dir">func</a> (\*Chan) [Dir](https://golang.org/src/go/types/type.go?s=12950:12978#L403)
<pre>func (c *<a href="#Chan">Chan</a>) Dir() <a href="#ChanDir">ChanDir</a></pre>
Dir returns the direction of channel c.




### <a id="Chan.Elem">func</a> (\*Chan) [Elem](https://golang.org/src/go/types/type.go?s=13044:13070#L406)
<pre>func (c *<a href="#Chan">Chan</a>) Elem() <a href="#Type">Type</a></pre>
Elem returns the element type of channel c.




### <a id="Chan.String">func</a> (\*Chan) [String](https://golang.org/src/go/types/type.go?s=16125:16155#L479)
<pre>func (c *<a href="#Chan">Chan</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Chan.Underlying">func</a> (\*Chan) [Underlying](https://golang.org/src/go/types/type.go?s=15417:15449#L467)
<pre>func (c *<a href="#Chan">Chan</a>) Underlying() <a href="#Type">Type</a></pre>



## <a id="ChanDir">type</a> [ChanDir](https://golang.org/src/go/types/type.go?s=12610:12626#L388)
A ChanDir value indicates a channel direction.


<pre>type ChanDir <a href="/pkg/builtin/#int">int</a></pre>


The direction of a channel is indicated by one of these constants.


<pre>const (
    <span id="SendRecv">SendRecv</span> <a href="#ChanDir">ChanDir</a> = <a href="/pkg/builtin/#iota">iota</a>
    <span id="SendOnly">SendOnly</span>
    <span id="RecvOnly">RecvOnly</span>
)</pre>









## <a id="Checker">type</a> [Checker](https://golang.org/src/go/types/check.go?s=2959:4573#L62)
A Checker maintains the state of the type checker.
It must be created with NewChecker.


<pre>type Checker struct {
    *<a href="#Info">Info</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewChecker">func</a> [NewChecker](https://golang.org/src/go/types/check.go?s=6516:6601#L156)
<pre>func NewChecker(conf *<a href="#Config">Config</a>, fset *<a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#FileSet">FileSet</a>, pkg *<a href="#Package">Package</a>, info *<a href="#Info">Info</a>) *<a href="#Checker">Checker</a></pre>
NewChecker returns a new Checker instance for a given package.
Package files may be added incrementally via checker.Files.






### <a id="Checker.Files">func</a> (\*Checker) [Files](https://golang.org/src/go/types/check.go?s=8801:8853#L235)
<pre>func (check *<a href="#Checker">Checker</a>) Files(files []*<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#File">File</a>) <a href="/pkg/builtin/#error">error</a></pre>
Files checks the provided files as part of the checker's package.




## <a id="Config">type</a> [Config](https://golang.org/src/go/types/api.go?s=3684:5154#L84)
A Config specifies the configuration for type checking.
The zero value for Config is a ready-to-use default configuration.


<pre>type Config struct {
    <span class="comment">// If IgnoreFuncBodies is set, function bodies are not</span>
    <span class="comment">// type-checked.</span>
<span id="Config.IgnoreFuncBodies"></span>    IgnoreFuncBodies <a href="/pkg/builtin/#bool">bool</a>

    <span class="comment">// If FakeImportC is set, `import &#34;C&#34;` (for packages requiring Cgo)</span>
    <span class="comment">// declares an empty &#34;C&#34; package and errors are omitted for qualified</span>
    <span class="comment">// identifiers referring to package C (which won&#39;t find an object).</span>
    <span class="comment">// This feature is intended for the standard library cmd/api tool.</span>
    <span class="comment">//</span>
    <span class="comment">// Caution: Effects may be unpredictable due to follow-on errors.</span>
    <span class="comment">//          Do not use casually!</span>
<span id="Config.FakeImportC"></span>    FakeImportC <a href="/pkg/builtin/#bool">bool</a>

    <span class="comment">// If Error != nil, it is called with each error found</span>
    <span class="comment">// during type checking; err has dynamic type Error.</span>
    <span class="comment">// Secondary errors (for instance, to enumerate all types</span>
    <span class="comment">// involved in an invalid recursive type declaration) have</span>
    <span class="comment">// error strings that start with a &#39;\t&#39; character.</span>
    <span class="comment">// If Error == nil, type-checking stops with the first</span>
    <span class="comment">// error found.</span>
<span id="Config.Error"></span>    Error func(err <a href="/pkg/builtin/#error">error</a>)

    <span class="comment">// An importer is used to import packages referred to from</span>
    <span class="comment">// import declarations.</span>
    <span class="comment">// If the installed importer implements ImporterFrom, the type</span>
    <span class="comment">// checker calls ImportFrom instead of Import.</span>
    <span class="comment">// The type checker reports an error if an importer is needed</span>
    <span class="comment">// but none was installed.</span>
<span id="Config.Importer"></span>    Importer <a href="#Importer">Importer</a>

    <span class="comment">// If Sizes != nil, it provides the sizing functions for package unsafe.</span>
    <span class="comment">// Otherwise SizesFor(&#34;gc&#34;, &#34;amd64&#34;) is used instead.</span>
<span id="Config.Sizes"></span>    Sizes <a href="#Sizes">Sizes</a>

    <span class="comment">// If DisableUnusedImportCheck is set, packages are not checked</span>
    <span class="comment">// for unused imports.</span>
<span id="Config.DisableUnusedImportCheck"></span>    DisableUnusedImportCheck <a href="/pkg/builtin/#bool">bool</a>
}
</pre>











### <a id="Config.Check">func</a> (\*Config) [Check](https://golang.org/src/go/types/api.go?s=12626:12734#L339)
<pre>func (conf *<a href="#Config">Config</a>) Check(path <a href="/pkg/builtin/#string">string</a>, fset *<a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#FileSet">FileSet</a>, files []*<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#File">File</a>, info *<a href="#Info">Info</a>) (*<a href="#Package">Package</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Check type-checks a package and returns the resulting package object and
the first error if any. Additionally, if info != nil, Check populates each
of the non-nil maps in the Info struct.

The package is marked as complete if no errors occurred, otherwise it is
incomplete. See Config.Error for controlling behavior in the presence of
errors.

The package is specified by a list of *ast.Files and corresponding
file set, and the package path the package is identified with.
The clean path must not be empty or dot (".").




## <a id="Const">type</a> [Const](https://golang.org/src/go/types/object.go?s=6399:6448#L187)
A Const represents a declared constant.


<pre>type Const struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewConst">func</a> [NewConst](https://golang.org/src/go/types/object.go?s=6571:6663#L194)
<pre>func NewConst(pos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>, pkg *<a href="#Package">Package</a>, name <a href="/pkg/builtin/#string">string</a>, typ <a href="#Type">Type</a>, val <a href="/pkg/go/constant/">constant</a>.<a href="/pkg/go/constant/#Value">Value</a>) *<a href="#Const">Const</a></pre>
NewConst returns a new constant with value val.
The remaining arguments set the attributes found with all Objects.






### <a id="Const.Exported">func</a> (\*Const) [Exported](https://golang.org/src/go/types/object.go?s=4259:4293#L131)
<pre>func (obj *<a href="#Const">Const</a>) Exported() <a href="/pkg/builtin/#bool">bool</a></pre>
Exported reports whether the object is exported (starts with a capital letter).
It doesn't take into account whether the object is in a local (function) scope
or not.




### <a id="Const.Id">func</a> (\*Const) [Id](https://golang.org/src/go/types/object.go?s=4383:4413#L134)
<pre>func (obj *<a href="#Const">Const</a>) Id() <a href="/pkg/builtin/#string">string</a></pre>
Id is a wrapper for Id(obj.Pkg(), obj.Name()).




### <a id="Const.Name">func</a> (\*Const) [Name](https://golang.org/src/go/types/object.go?s=3943:3975#L123)
<pre>func (obj *<a href="#Const">Const</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the object's (package-local, unqualified) name.




### <a id="Const.Parent">func</a> (\*Const) [Parent](https://golang.org/src/go/types/object.go?s=3521:3555#L113)
<pre>func (obj *<a href="#Const">Const</a>) Parent() *<a href="#Scope">Scope</a></pre>
Parent returns the scope in which the object is declared.
The result is nil for methods and struct fields.




### <a id="Const.Pkg">func</a> (\*Const) [Pkg](https://golang.org/src/go/types/object.go?s=3825:3858#L120)
<pre>func (obj *<a href="#Const">Const</a>) Pkg() *<a href="#Package">Package</a></pre>
Pkg returns the package to which the object belongs.
The result is nil for labels and objects in the Universe scope.




### <a id="Const.Pos">func</a> (\*Const) [Pos](https://golang.org/src/go/types/object.go?s=3647:3681#L116)
<pre>func (obj *<a href="#Const">Const</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>
Pos returns the declaration position of the object's identifier.




### <a id="Const.String">func</a> (\*Const) [String](https://golang.org/src/go/types/object.go?s=13938:13971#L447)
<pre>func (obj *<a href="#Const">Const</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Const.Type">func</a> (\*Const) [Type](https://golang.org/src/go/types/object.go?s=4032:4062#L126)
<pre>func (obj *<a href="#Const">Const</a>) Type() <a href="#Type">Type</a></pre>
Type returns the object's type.




### <a id="Const.Val">func</a> (\*Const) [Val](https://golang.org/src/go/types/object.go?s=6791:6829#L199)
<pre>func (obj *<a href="#Const">Const</a>) Val() <a href="/pkg/go/constant/">constant</a>.<a href="/pkg/go/constant/#Value">Value</a></pre>
Val returns the constant's value.




## <a id="Error">type</a> [Error](https://golang.org/src/go/types/api.go?s=1516:1721#L31)
An Error describes a type-checking error; it implements the error interface.
A "soft" error is an error that still permits a valid interpretation of a
package (such as "unused variable"); "hard" errors may lead to unpredictable
behavior if ignored.


<pre>type Error struct {
<span id="Error.Fset"></span>    Fset *<a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#FileSet">FileSet</a> <span class="comment">// file set for interpretation of Pos</span>
<span id="Error.Pos"></span>    Pos  <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>      <span class="comment">// error position</span>
<span id="Error.Msg"></span>    Msg  <a href="/pkg/builtin/#string">string</a>         <span class="comment">// error message</span>
<span id="Error.Soft"></span>    Soft <a href="/pkg/builtin/#bool">bool</a>           <span class="comment">// if set, error is &#34;soft&#34;</span>
}
</pre>











### <a id="Error.Error">func</a> (Error) [Error](https://golang.org/src/go/types/api.go?s=1811:1842#L40)
<pre>func (err <a href="#Error">Error</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>
Error returns an error string formatted as follows:
filename:line:column: message




## <a id="Func">type</a> [Func](https://golang.org/src/go/types/object.go?s=10327:10425#L284)
A Func represents a declared function, concrete method, or abstract
(interface) method. Its Type() is always a *Signature.
An abstract method may belong to many interfaces due to embedding.


<pre>type Func struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="MissingMethod">func</a> [MissingMethod](https://golang.org/src/go/types/lookup.go?s=9112:9196#L245)
<pre>func MissingMethod(V <a href="#Type">Type</a>, T *<a href="#Interface">Interface</a>, static <a href="/pkg/builtin/#bool">bool</a>) (method *<a href="#Func">Func</a>, wrongType <a href="/pkg/builtin/#bool">bool</a>)</pre>
MissingMethod returns (nil, false) if V implements T, otherwise it
returns a missing method required by T and whether it is missing or
just has the wrong type.

For non-interface types V, or if static is set, V implements T if all
methods of T are present in V. Otherwise (V is an interface and static
is not set), MissingMethod only checks that methods of T which are also
present in V have matching types (e.g., for a type assertion x.(T) where
x is of interface type V).




### <a id="NewFunc">func</a> [NewFunc](https://golang.org/src/go/types/object.go?s=10524:10600#L291)
<pre>func NewFunc(pos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>, pkg *<a href="#Package">Package</a>, name <a href="/pkg/builtin/#string">string</a>, sig *<a href="#Signature">Signature</a>) *<a href="#Func">Func</a></pre>
NewFunc returns a new function with the given signature, representing
the function's type.






### <a id="Func.Exported">func</a> (\*Func) [Exported](https://golang.org/src/go/types/object.go?s=4259:4293#L131)
<pre>func (obj *<a href="#Func">Func</a>) Exported() <a href="/pkg/builtin/#bool">bool</a></pre>
Exported reports whether the object is exported (starts with a capital letter).
It doesn't take into account whether the object is in a local (function) scope
or not.




### <a id="Func.FullName">func</a> (\*Func) [FullName](https://golang.org/src/go/types/object.go?s=10865:10899#L302)
<pre>func (obj *<a href="#Func">Func</a>) FullName() <a href="/pkg/builtin/#string">string</a></pre>
FullName returns the package- or receiver-type-qualified name of
function or method obj.




### <a id="Func.Id">func</a> (\*Func) [Id](https://golang.org/src/go/types/object.go?s=4383:4413#L134)
<pre>func (obj *<a href="#Func">Func</a>) Id() <a href="/pkg/builtin/#string">string</a></pre>
Id is a wrapper for Id(obj.Pkg(), obj.Name()).




### <a id="Func.Name">func</a> (\*Func) [Name](https://golang.org/src/go/types/object.go?s=3943:3975#L123)
<pre>func (obj *<a href="#Func">Func</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the object's (package-local, unqualified) name.




### <a id="Func.Parent">func</a> (\*Func) [Parent](https://golang.org/src/go/types/object.go?s=3521:3555#L113)
<pre>func (obj *<a href="#Func">Func</a>) Parent() *<a href="#Scope">Scope</a></pre>
Parent returns the scope in which the object is declared.
The result is nil for methods and struct fields.




### <a id="Func.Pkg">func</a> (\*Func) [Pkg](https://golang.org/src/go/types/object.go?s=3825:3858#L120)
<pre>func (obj *<a href="#Func">Func</a>) Pkg() *<a href="#Package">Package</a></pre>
Pkg returns the package to which the object belongs.
The result is nil for labels and objects in the Universe scope.




### <a id="Func.Pos">func</a> (\*Func) [Pos](https://golang.org/src/go/types/object.go?s=3647:3681#L116)
<pre>func (obj *<a href="#Func">Func</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>
Pos returns the declaration position of the object's identifier.




### <a id="Func.Scope">func</a> (\*Func) [Scope](https://golang.org/src/go/types/object.go?s=11036:11067#L309)
<pre>func (obj *<a href="#Func">Func</a>) Scope() *<a href="#Scope">Scope</a></pre>
Scope returns the scope of the function's body block.




### <a id="Func.String">func</a> (\*Func) [String](https://golang.org/src/go/types/object.go?s=14151:14183#L450)
<pre>func (obj *<a href="#Func">Func</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Func.Type">func</a> (\*Func) [Type](https://golang.org/src/go/types/object.go?s=4032:4062#L126)
<pre>func (obj *<a href="#Func">Func</a>) Type() <a href="#Type">Type</a></pre>
Type returns the object's type.




## <a id="ImportMode">type</a> [ImportMode](https://golang.org/src/go/types/api.go?s=2467:2486#L57)
ImportMode is reserved for future use.


<pre>type ImportMode <a href="/pkg/builtin/#int">int</a></pre>











## <a id="Importer">type</a> [Importer](https://golang.org/src/go/types/api.go?s=2165:2423#L49)
An Importer resolves import paths to Packages.

CAUTION: This interface does not support the import of locally
vendored packages. See <a href="https://golang.org/s/go15vendor">https://golang.org/s/go15vendor</a>.
If possible, external implementations should implement ImporterFrom.


<pre>type Importer interface {
    <span class="comment">// Import returns the imported package for the given import path.</span>
    <span class="comment">// The semantics is like for ImporterFrom.ImportFrom except that</span>
    <span class="comment">// dir and mode are ignored (since they are not present).</span>
    Import(path <a href="/pkg/builtin/#string">string</a>) (*<a href="#Package">Package</a>, <a href="/pkg/builtin/#error">error</a>)
}</pre>











## <a id="ImporterFrom">type</a> [ImporterFrom](https://golang.org/src/go/types/api.go?s=2665:3553#L62)
An ImporterFrom resolves import paths to packages; it
supports vendoring per <a href="https://golang.org/s/go15vendor">https://golang.org/s/go15vendor</a>.
Use go/importer to obtain an ImporterFrom implementation.


<pre>type ImporterFrom interface {
    <span class="comment">// Importer is present for backward-compatibility. Calling</span>
    <span class="comment">// Import(path) is the same as calling ImportFrom(path, &#34;&#34;, 0);</span>
    <span class="comment">// i.e., locally vendored packages may not be found.</span>
    <span class="comment">// The types package does not call Import if an ImporterFrom</span>
    <span class="comment">// is present.</span>
    <a href="#Importer">Importer</a>

    <span class="comment">// ImportFrom returns the imported package for the given import</span>
    <span class="comment">// path when imported by a package file located in dir.</span>
    <span class="comment">// If the import failed, besides returning an error, ImportFrom</span>
    <span class="comment">// is encouraged to cache and return a package anyway, if one</span>
    <span class="comment">// was created. This will reduce package inconsistencies and</span>
    <span class="comment">// follow-on type checker errors due to the missing package.</span>
    <span class="comment">// The mode value must be 0; it is reserved for future use.</span>
    <span class="comment">// Two calls to ImportFrom with the same path and dir must</span>
    <span class="comment">// return the same package.</span>
    ImportFrom(path, dir <a href="/pkg/builtin/#string">string</a>, mode <a href="#ImportMode">ImportMode</a>) (*<a href="#Package">Package</a>, <a href="/pkg/builtin/#error">error</a>)
}</pre>











## <a id="Info">type</a> [Info](https://golang.org/src/go/types/api.go?s=5371:8910#L128)
Info holds result type information for a type-checked package.
Only the information for which a map is provided is collected.
If the package has type errors, the collected information may
be incomplete.


<pre>type Info struct {
<span id="Info.Types"></span>    <span class="comment">// Types maps expressions to their types, and for constant</span>
    <span class="comment">// expressions, also their values. Invalid expressions are</span>
    <span class="comment">// omitted.</span>
    <span class="comment">//</span>
    <span class="comment">// For (possibly parenthesized) identifiers denoting built-in</span>
    <span class="comment">// functions, the recorded signatures are call-site specific:</span>
    <span class="comment">// if the call result is not a constant, the recorded type is</span>
    <span class="comment">// an argument-specific signature. Otherwise, the recorded type</span>
    <span class="comment">// is invalid.</span>
    <span class="comment">//</span>
    <span class="comment">// The Types map does not record the type of every identifier,</span>
    <span class="comment">// only those that appear where an arbitrary expression is</span>
    <span class="comment">// permitted. For instance, the identifier f in a selector</span>
    <span class="comment">// expression x.f is found only in the Selections map, the</span>
    <span class="comment">// identifier z in a variable declaration &#39;var z int&#39; is found</span>
    <span class="comment">// only in the Defs map, and identifiers denoting packages in</span>
    <span class="comment">// qualified identifiers are collected in the Uses map.</span>
    Types map[<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#Expr">Expr</a>]<a href="#TypeAndValue">TypeAndValue</a>

<span id="Info.Defs"></span>    <span class="comment">// Defs maps identifiers to the objects they define (including</span>
    <span class="comment">// package names, dots &#34;.&#34; of dot-imports, and blank &#34;_&#34; identifiers).</span>
    <span class="comment">// For identifiers that do not denote objects (e.g., the package name</span>
    <span class="comment">// in package clauses, or symbolic variables t in t := x.(type) of</span>
    <span class="comment">// type switch headers), the corresponding objects are nil.</span>
    <span class="comment">//</span>
    <span class="comment">// For an embedded field, Defs returns the field *Var it defines.</span>
    <span class="comment">//</span>
    <span class="comment">// Invariant: Defs[id] == nil || Defs[id].Pos() == id.Pos()</span>
    Defs map[*<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#Ident">Ident</a>]<a href="#Object">Object</a>

<span id="Info.Uses"></span>    <span class="comment">// Uses maps identifiers to the objects they denote.</span>
    <span class="comment">//</span>
    <span class="comment">// For an embedded field, Uses returns the *TypeName it denotes.</span>
    <span class="comment">//</span>
    <span class="comment">// Invariant: Uses[id].Pos() != id.Pos()</span>
    Uses map[*<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#Ident">Ident</a>]<a href="#Object">Object</a>

<span id="Info.Implicits"></span>    <span class="comment">// Implicits maps nodes to their implicitly declared objects, if any.</span>
    <span class="comment">// The following node and object types may appear:</span>
    <span class="comment">//</span>
    <span class="comment">//     node               declared object</span>
    <span class="comment">//</span>
    <span class="comment">//     *ast.ImportSpec    *PkgName for imports without renames</span>
    <span class="comment">//     *ast.CaseClause    type-specific *Var for each type switch case clause (incl. default)</span>
    <span class="comment">//     *ast.Field         anonymous parameter *Var (incl. unnamed results)</span>
    <span class="comment">//</span>
    Implicits map[<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#Node">Node</a>]<a href="#Object">Object</a>

<span id="Info.Selections"></span>    <span class="comment">// Selections maps selector expressions (excluding qualified identifiers)</span>
    <span class="comment">// to their corresponding selections.</span>
    Selections map[*<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#SelectorExpr">SelectorExpr</a>]*<a href="#Selection">Selection</a>

<span id="Info.Scopes"></span>    <span class="comment">// Scopes maps ast.Nodes to the scopes they define. Package scopes are not</span>
    <span class="comment">// associated with a specific node but with all files belonging to a package.</span>
    <span class="comment">// Thus, the package scope can be found in the type-checked Package object.</span>
    <span class="comment">// Scopes nest, with the Universe scope being the outermost scope, enclosing</span>
    <span class="comment">// the package scope, which contains (one or more) files scopes, which enclose</span>
    <span class="comment">// function scopes which in turn enclose statement and function literal scopes.</span>
    <span class="comment">// Note that even though package-level functions are declared in the package</span>
    <span class="comment">// scope, the function scopes are embedded in the file scope of the file</span>
    <span class="comment">// containing the function declaration.</span>
    <span class="comment">//</span>
    <span class="comment">// The following node types may appear in Scopes:</span>
    <span class="comment">//</span>
    <span class="comment">//     *ast.File</span>
    <span class="comment">//     *ast.FuncType</span>
    <span class="comment">//     *ast.BlockStmt</span>
    <span class="comment">//     *ast.IfStmt</span>
    <span class="comment">//     *ast.SwitchStmt</span>
    <span class="comment">//     *ast.TypeSwitchStmt</span>
    <span class="comment">//     *ast.CaseClause</span>
    <span class="comment">//     *ast.CommClause</span>
    <span class="comment">//     *ast.ForStmt</span>
    <span class="comment">//     *ast.RangeStmt</span>
    <span class="comment">//</span>
    Scopes map[<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#Node">Node</a>]*<a href="#Scope">Scope</a>

<span id="Info.InitOrder"></span>    <span class="comment">// InitOrder is the list of package-level initializers in the order in which</span>
    <span class="comment">// they must be executed. Initializers referring to variables related by an</span>
    <span class="comment">// initialization dependency appear in topological order, the others appear</span>
    <span class="comment">// in source order. Variables without an initialization expression do not</span>
    <span class="comment">// appear in this list.</span>
    InitOrder []*<a href="#Initializer">Initializer</a>
}
</pre>





<a id="example_Info">Example</a>
<p>ExampleInfo prints various facts recorded by the type checker in a
types.Info struct: definitions of and references to each named object,
and the type, value, and mode of every expression in the package.
</p>





### <a id="Info.ObjectOf">func</a> (\*Info) [ObjectOf](https://golang.org/src/go/types/api.go?s=9535:9583#L237)
<pre>func (info *<a href="#Info">Info</a>) ObjectOf(id *<a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#Ident">Ident</a>) <a href="#Object">Object</a></pre>
ObjectOf returns the object denoted by the specified id,
or nil if not found.

If id is an embedded struct field, ObjectOf returns the field (*Var)
it defines, not the type (*TypeName) it uses.

Precondition: the Uses and Defs maps are populated.




### <a id="Info.TypeOf">func</a> (\*Info) [TypeOf](https://golang.org/src/go/types/api.go?s=9042:9083#L217)
<pre>func (info *<a href="#Info">Info</a>) TypeOf(e <a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#Expr">Expr</a>) <a href="#Type">Type</a></pre>
TypeOf returns the type of expression e, or nil if not found.
Precondition: the Types, Uses and Defs maps are populated.




## <a id="Initializer">type</a> [Initializer](https://golang.org/src/go/types/api.go?s=11755:11825#L310)
An Initializer describes a package-level variable, or a list of variables in case
of a multi-valued initialization expression, and the corresponding initialization
expression.


<pre>type Initializer struct {
<span id="Initializer.Lhs"></span>    Lhs []*<a href="#Var">Var</a> <span class="comment">// var Lhs = Rhs</span>
<span id="Initializer.Rhs"></span>    Rhs <a href="/pkg/go/ast/">ast</a>.<a href="/pkg/go/ast/#Expr">Expr</a>
}
</pre>











### <a id="Initializer.String">func</a> (\*Initializer) [String](https://golang.org/src/go/types/api.go?s=11827:11867#L315)
<pre>func (init *<a href="#Initializer">Initializer</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Interface">type</a> [Interface](https://golang.org/src/go/types/type.go?s=6899:7177#L234)
An Interface represents an interface type.


<pre>type Interface struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewInterface">func</a> [NewInterface](https://golang.org/src/go/types/type.go?s=8052:8117#L256)
<pre>func NewInterface(methods []*<a href="#Func">Func</a>, embeddeds []*<a href="#Named">Named</a>) *<a href="#Interface">Interface</a></pre>
NewInterface returns a new (incomplete) interface for the given methods and embedded types.
Each embedded type must have an underlying type of interface type.
NewInterface takes ownership of the provided methods and may modify their types by setting
missing receivers. To compute the method set of the interface, Complete must be called.

Deprecated: Use NewInterfaceType instead which allows any (even non-defined) interface types
to be embedded. This is necessary for interfaces that embed alias type names referring to
non-defined (literal) interface types.




### <a id="NewInterfaceType">func</a> [NewInterfaceType](https://golang.org/src/go/types/type.go?s=8763:8830#L270)
<pre>func NewInterfaceType(methods []*<a href="#Func">Func</a>, embeddeds []<a href="#Type">Type</a>) *<a href="#Interface">Interface</a></pre>
NewInterfaceType returns a new (incomplete) interface for the given methods and embedded types.
Each embedded type must have an underlying type of interface type (this property is not
verified for defined types, which may be in the process of being set up and which don't
have a valid underlying type yet).
NewInterfaceType takes ownership of the provided methods and may modify their types by setting
missing receivers. To compute the method set of the interface, Complete must be called.






### <a id="Interface.Complete">func</a> (\*Interface) [Complete](https://golang.org/src/go/types/type.go?s=11540:11581#L340)
<pre>func (t *<a href="#Interface">Interface</a>) Complete() *<a href="#Interface">Interface</a></pre>
Complete computes the interface's method set. It must be called by users of
NewInterfaceType and NewInterface after the interface's embedded types are
fully defined and before using the interface type in any way other than to
form other types. Complete returns the receiver.




### <a id="Interface.Embedded">func</a> (\*Interface) [Embedded](https://golang.org/src/go/types/type.go?s=10545:10587#L321)
<pre>func (t *<a href="#Interface">Interface</a>) Embedded(i <a href="/pkg/builtin/#int">int</a>) *<a href="#Named">Named</a></pre>
Embedded returns the i'th embedded defined (*Named) type of interface t for 0 <= i < t.NumEmbeddeds().
The result is nil if the i'th embedded type is not a defined type.

Deprecated: Use EmbeddedType which is not restricted to defined (*Named) types.




### <a id="Interface.EmbeddedType">func</a> (\*Interface) [EmbeddedType](https://golang.org/src/go/types/type.go?s=10736:10780#L324)
<pre>func (t *<a href="#Interface">Interface</a>) EmbeddedType(i <a href="/pkg/builtin/#int">int</a>) <a href="#Type">Type</a></pre>
EmbeddedType returns the i'th embedded type of interface t for 0 <= i < t.NumEmbeddeds().




### <a id="Interface.Empty">func</a> (\*Interface) [Empty](https://golang.org/src/go/types/type.go?s=11185:11217#L334)
<pre>func (t *<a href="#Interface">Interface</a>) Empty() <a href="/pkg/builtin/#bool">bool</a></pre>
Empty reports whether t is the empty interface.




### <a id="Interface.ExplicitMethod">func</a> (\*Interface) [ExplicitMethod](https://golang.org/src/go/types/type.go?s=10073:10120#L312)
<pre>func (t *<a href="#Interface">Interface</a>) ExplicitMethod(i <a href="/pkg/builtin/#int">int</a>) *<a href="#Func">Func</a></pre>
ExplicitMethod returns the i'th explicitly declared method of interface t for 0 <= i < t.NumExplicitMethods().
The methods are ordered by their unique Id.




### <a id="Interface.Method">func</a> (\*Interface) [Method](https://golang.org/src/go/types/type.go?s=11066:11105#L331)
<pre>func (t *<a href="#Interface">Interface</a>) Method(i <a href="/pkg/builtin/#int">int</a>) *<a href="#Func">Func</a></pre>
Method returns the i'th method of interface t for 0 <= i < t.NumMethods().
The methods are ordered by their unique Id.




### <a id="Interface.NumEmbeddeds">func</a> (\*Interface) [NumEmbeddeds](https://golang.org/src/go/types/type.go?s=10215:10253#L315)
<pre>func (t *<a href="#Interface">Interface</a>) NumEmbeddeds() <a href="/pkg/builtin/#int">int</a></pre>
NumEmbeddeds returns the number of embedded types in interface t.




### <a id="Interface.NumExplicitMethods">func</a> (\*Interface) [NumExplicitMethods](https://golang.org/src/go/types/type.go?s=9840:9884#L308)
<pre>func (t *<a href="#Interface">Interface</a>) NumExplicitMethods() <a href="/pkg/builtin/#int">int</a></pre>
NumExplicitMethods returns the number of explicitly declared methods of interface t.




### <a id="Interface.NumMethods">func</a> (\*Interface) [NumMethods](https://golang.org/src/go/types/type.go?s=10874:10910#L327)
<pre>func (t *<a href="#Interface">Interface</a>) NumMethods() <a href="/pkg/builtin/#int">int</a></pre>
NumMethods returns the total number of methods of interface t.




### <a id="Interface.String">func</a> (\*Interface) [String](https://golang.org/src/go/types/type.go?s=15993:16028#L477)
<pre>func (t *<a href="#Interface">Interface</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Interface.Underlying">func</a> (\*Interface) [Underlying](https://golang.org/src/go/types/type.go?s=15315:15352#L465)
<pre>func (t *<a href="#Interface">Interface</a>) Underlying() <a href="#Type">Type</a></pre>



## <a id="Label">type</a> [Label](https://golang.org/src/go/types/object.go?s=11274:11343#L315)
A Label represents a declared label.
Labels don't have a type.


<pre>type Label struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewLabel">func</a> [NewLabel](https://golang.org/src/go/types/object.go?s=11378:11440#L321)
<pre>func NewLabel(pos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>, pkg *<a href="#Package">Package</a>, name <a href="/pkg/builtin/#string">string</a>) *<a href="#Label">Label</a></pre>
NewLabel returns a new label.






### <a id="Label.Exported">func</a> (\*Label) [Exported](https://golang.org/src/go/types/object.go?s=4259:4293#L131)
<pre>func (obj *<a href="#Label">Label</a>) Exported() <a href="/pkg/builtin/#bool">bool</a></pre>
Exported reports whether the object is exported (starts with a capital letter).
It doesn't take into account whether the object is in a local (function) scope
or not.




### <a id="Label.Id">func</a> (\*Label) [Id](https://golang.org/src/go/types/object.go?s=4383:4413#L134)
<pre>func (obj *<a href="#Label">Label</a>) Id() <a href="/pkg/builtin/#string">string</a></pre>
Id is a wrapper for Id(obj.Pkg(), obj.Name()).




### <a id="Label.Name">func</a> (\*Label) [Name](https://golang.org/src/go/types/object.go?s=3943:3975#L123)
<pre>func (obj *<a href="#Label">Label</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the object's (package-local, unqualified) name.




### <a id="Label.Parent">func</a> (\*Label) [Parent](https://golang.org/src/go/types/object.go?s=3521:3555#L113)
<pre>func (obj *<a href="#Label">Label</a>) Parent() *<a href="#Scope">Scope</a></pre>
Parent returns the scope in which the object is declared.
The result is nil for methods and struct fields.




### <a id="Label.Pkg">func</a> (\*Label) [Pkg](https://golang.org/src/go/types/object.go?s=3825:3858#L120)
<pre>func (obj *<a href="#Label">Label</a>) Pkg() *<a href="#Package">Package</a></pre>
Pkg returns the package to which the object belongs.
The result is nil for labels and objects in the Universe scope.




### <a id="Label.Pos">func</a> (\*Label) [Pos](https://golang.org/src/go/types/object.go?s=3647:3681#L116)
<pre>func (obj *<a href="#Label">Label</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>
Pos returns the declaration position of the object's identifier.




### <a id="Label.String">func</a> (\*Label) [String](https://golang.org/src/go/types/object.go?s=14222:14255#L451)
<pre>func (obj *<a href="#Label">Label</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Label.Type">func</a> (\*Label) [Type](https://golang.org/src/go/types/object.go?s=4032:4062#L126)
<pre>func (obj *<a href="#Label">Label</a>) Type() <a href="#Type">Type</a></pre>
Type returns the object's type.




## <a id="Map">type</a> [Map](https://golang.org/src/go/types/type.go?s=12143:12178#L366)
A Map represents a map type.


<pre>type Map struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewMap">func</a> [NewMap](https://golang.org/src/go/types/type.go?s=12245:12277#L371)
<pre>func NewMap(key, elem <a href="#Type">Type</a>) *<a href="#Map">Map</a></pre>
NewMap returns a new map for the given key and element types.






### <a id="Map.Elem">func</a> (\*Map) [Elem](https://golang.org/src/go/types/type.go?s=12431:12456#L379)
<pre>func (m *<a href="#Map">Map</a>) Elem() <a href="#Type">Type</a></pre>
Elem returns the element type of map m.




### <a id="Map.Key">func</a> (\*Map) [Key](https://golang.org/src/go/types/type.go?s=12345:12369#L376)
<pre>func (m *<a href="#Map">Map</a>) Key() <a href="#Type">Type</a></pre>
Key returns the key type of map m.




### <a id="Map.String">func</a> (\*Map) [String](https://golang.org/src/go/types/type.go?s=16059:16088#L478)
<pre>func (m *<a href="#Map">Map</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Map.Underlying">func</a> (\*Map) [Underlying](https://golang.org/src/go/types/type.go?s=15366:15397#L466)
<pre>func (m *<a href="#Map">Map</a>) Underlying() <a href="#Type">Type</a></pre>



## <a id="MethodSet">type</a> [MethodSet](https://golang.org/src/go/types/methodset.go?s=485:529#L8)
A MethodSet is an ordered set of concrete or abstract (interface) methods;
a method is a MethodVal selection, and they are ordered by ascending m.Obj().Id().
The zero value for a MethodSet is a ready-to-use empty method set.


<pre>type MethodSet struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_MethodSet">Example</a>
<p>ExampleMethodSet prints the method sets of various types.
</p>



### <a id="NewMethodSet">func</a> [NewMethodSet](https://golang.org/src/go/types/methodset.go?s=1592:1628#L57)
<pre>func NewMethodSet(T <a href="#Type">Type</a>) *<a href="#MethodSet">MethodSet</a></pre>
NewMethodSet returns the method set for the given type T.
It always returns a non-nil method set, even if it is empty.






### <a id="MethodSet.At">func</a> (\*MethodSet) [At](https://golang.org/src/go/types/methodset.go?s=943:983#L30)
<pre>func (s *<a href="#MethodSet">MethodSet</a>) At(i <a href="/pkg/builtin/#int">int</a>) *<a href="#Selection">Selection</a></pre>
At returns the i'th method in s for 0 <= i < s.Len().




### <a id="MethodSet.Len">func</a> (\*MethodSet) [Len](https://golang.org/src/go/types/methodset.go?s=832:861#L27)
<pre>func (s *<a href="#MethodSet">MethodSet</a>) Len() <a href="/pkg/builtin/#int">int</a></pre>
Len returns the number of methods in s.




### <a id="MethodSet.Lookup">func</a> (\*MethodSet) [Lookup](https://golang.org/src/go/types/methodset.go?s=1088:1152#L33)
<pre>func (s *<a href="#MethodSet">MethodSet</a>) Lookup(pkg *<a href="#Package">Package</a>, name <a href="/pkg/builtin/#string">string</a>) *<a href="#Selection">Selection</a></pre>
Lookup returns the method with matching package and name, or nil if not found.




### <a id="MethodSet.String">func</a> (\*MethodSet) [String](https://golang.org/src/go/types/methodset.go?s=531:566#L12)
<pre>func (s *<a href="#MethodSet">MethodSet</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Named">type</a> [Named](https://golang.org/src/go/types/type.go?s=13126:13423#L409)
A Named represents a named type.


<pre>type Named struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewNamed">func</a> [NewNamed](https://golang.org/src/go/types/type.go?s=13672:13741#L418)
<pre>func NewNamed(obj *<a href="#TypeName">TypeName</a>, underlying <a href="#Type">Type</a>, methods []*<a href="#Func">Func</a>) *<a href="#Named">Named</a></pre>
NewNamed returns a new named type for the given type name, underlying type, and associated methods.
If the given type name obj doesn't have a type yet, its type is set to the returned named type.
The underlying type must not be a *Named.






### <a id="Named.AddMethod">func</a> (\*Named) [AddMethod](https://golang.org/src/go/types/type.go?s=14782:14816#L450)
<pre>func (t *<a href="#Named">Named</a>) AddMethod(m *<a href="#Func">Func</a>)</pre>
AddMethod adds method m unless it is already in the method list.




### <a id="Named.Method">func</a> (\*Named) [Method](https://golang.org/src/go/types/type.go?s=14295:14330#L436)
<pre>func (t *<a href="#Named">Named</a>) Method(i <a href="/pkg/builtin/#int">int</a>) *<a href="#Func">Func</a></pre>
Method returns the i'th method of named type t for 0 <= i < t.NumMethods().




### <a id="Named.NumMethods">func</a> (\*Named) [NumMethods](https://golang.org/src/go/types/type.go?s=14156:14188#L433)
<pre>func (t *<a href="#Named">Named</a>) NumMethods() <a href="/pkg/builtin/#int">int</a></pre>
NumMethods returns the number of explicit methods whose receiver is named type t.




### <a id="Named.Obj">func</a> (\*Named) [Obj](https://golang.org/src/go/types/type.go?s=14021:14052#L430)
<pre>func (t *<a href="#Named">Named</a>) Obj() *<a href="#TypeName">TypeName</a></pre>
Obj returns the type name for the named type t.




### <a id="Named.SetUnderlying">func</a> (\*Named) [SetUnderlying](https://golang.org/src/go/types/type.go?s=14423:14469#L439)
<pre>func (t *<a href="#Named">Named</a>) SetUnderlying(underlying <a href="#Type">Type</a>)</pre>
SetUnderlying sets the underlying type and marks t as complete.




### <a id="Named.String">func</a> (\*Named) [String](https://golang.org/src/go/types/type.go?s=16191:16222#L480)
<pre>func (t *<a href="#Named">Named</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Named.Underlying">func</a> (\*Named) [Underlying](https://golang.org/src/go/types/type.go?s=15468:15501#L468)
<pre>func (t *<a href="#Named">Named</a>) Underlying() <a href="#Type">Type</a></pre>



## <a id="Nil">type</a> [Nil](https://golang.org/src/go/types/object.go?s=11856:11883#L337)
Nil represents the predeclared value nil.


<pre>type Nil struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="Nil.Exported">func</a> (\*Nil) [Exported](https://golang.org/src/go/types/object.go?s=4259:4293#L131)
<pre>func (obj *<a href="#Nil">Nil</a>) Exported() <a href="/pkg/builtin/#bool">bool</a></pre>
Exported reports whether the object is exported (starts with a capital letter).
It doesn't take into account whether the object is in a local (function) scope
or not.




### <a id="Nil.Id">func</a> (\*Nil) [Id](https://golang.org/src/go/types/object.go?s=4383:4413#L134)
<pre>func (obj *<a href="#Nil">Nil</a>) Id() <a href="/pkg/builtin/#string">string</a></pre>
Id is a wrapper for Id(obj.Pkg(), obj.Name()).




### <a id="Nil.Name">func</a> (\*Nil) [Name](https://golang.org/src/go/types/object.go?s=3943:3975#L123)
<pre>func (obj *<a href="#Nil">Nil</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the object's (package-local, unqualified) name.




### <a id="Nil.Parent">func</a> (\*Nil) [Parent](https://golang.org/src/go/types/object.go?s=3521:3555#L113)
<pre>func (obj *<a href="#Nil">Nil</a>) Parent() *<a href="#Scope">Scope</a></pre>
Parent returns the scope in which the object is declared.
The result is nil for methods and struct fields.




### <a id="Nil.Pkg">func</a> (\*Nil) [Pkg](https://golang.org/src/go/types/object.go?s=3825:3858#L120)
<pre>func (obj *<a href="#Nil">Nil</a>) Pkg() *<a href="#Package">Package</a></pre>
Pkg returns the package to which the object belongs.
The result is nil for labels and objects in the Universe scope.




### <a id="Nil.Pos">func</a> (\*Nil) [Pos](https://golang.org/src/go/types/object.go?s=3647:3681#L116)
<pre>func (obj *<a href="#Nil">Nil</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>
Pos returns the declaration position of the object's identifier.




### <a id="Nil.String">func</a> (\*Nil) [String](https://golang.org/src/go/types/object.go?s=14364:14395#L453)
<pre>func (obj *<a href="#Nil">Nil</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Nil.Type">func</a> (\*Nil) [Type](https://golang.org/src/go/types/object.go?s=4032:4062#L126)
<pre>func (obj *<a href="#Nil">Nil</a>) Type() <a href="#Type">Type</a></pre>
Type returns the object's type.




## <a id="Object">type</a> [Object](https://golang.org/src/go/types/object.go?s=411:1898#L8)
An Object describes a named language entity such as a package,
constant, type, variable, function (incl. methods), or label.
All objects implement the Object interface.


<pre>type Object interface {
    Parent() *<a href="#Scope">Scope</a> <span class="comment">// scope in which this object is declared; nil for methods and struct fields</span>
    Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a> <span class="comment">// position of object identifier in declaration</span>
    Pkg() *<a href="#Package">Package</a>  <span class="comment">// package to which this object belongs; nil for labels and objects in the Universe scope</span>
    Name() <a href="/pkg/builtin/#string">string</a>   <span class="comment">// package local object name</span>
    Type() <a href="#Type">Type</a>     <span class="comment">// object type</span>
    Exported() <a href="/pkg/builtin/#bool">bool</a> <span class="comment">// reports whether the name starts with a capital letter</span>
    Id() <a href="/pkg/builtin/#string">string</a>     <span class="comment">// object name if exported, qualified name if not exported (see func Id)</span>

    <span class="comment">// String returns a human-readable string of the object.</span>
    String() <a href="/pkg/builtin/#string">string</a>
    <span class="comment">// contains filtered or unexported methods</span>
}</pre>









### <a id="LookupFieldOrMethod">func</a> [LookupFieldOrMethod](https://golang.org/src/go/types/lookup.go?s=1793:1911#L30)
<pre>func LookupFieldOrMethod(T <a href="#Type">Type</a>, addressable <a href="/pkg/builtin/#bool">bool</a>, pkg *<a href="#Package">Package</a>, name <a href="/pkg/builtin/#string">string</a>) (obj <a href="#Object">Object</a>, index []<a href="/pkg/builtin/#int">int</a>, indirect <a href="/pkg/builtin/#bool">bool</a>)</pre>
LookupFieldOrMethod looks up a field or method with given package and name
in T and returns the corresponding *Var or *Func, an index sequence, and a
bool indicating if there were any pointer indirections on the path to the
field or method. If addressable is set, T is the type of an addressable
variable (only matters for method lookups).

The last index entry is the field or method index in the (possibly embedded)
type where the entry was found, either:


	1) the list of declared methods of a named type; or
	2) the list of all methods (method set) of an interface type; or
	3) the list of fields of a struct type.

The earlier index entries are the indices of the embedded struct fields
traversed to get to the found entry, starting at depth 0.

If no entry is found, a nil object is returned. In this case, the returned
index and indirect values have the following meaning:


		- If index != nil, the index sequence points to an ambiguous entry
		(the same name appeared more than once at the same embedding level).
	
		- If indirect is set, a method with a pointer receiver type was found
	     but there was no pointer on the path from the actual receiver type to
		the method's formal receiver base type, nor was the receiver addressable.






## <a id="Package">type</a> [Package](https://golang.org/src/go/types/package.go?s=243:451#L3)
A Package describes a Go package.


<pre>type Package struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>




The Unsafe package is the package returned by an importer
for the import path "unsafe".


<pre>var <span id="Unsafe">Unsafe</span> *<a href="#Package">Package</a></pre>





### <a id="NewPackage">func</a> [NewPackage](https://golang.org/src/go/types/package.go?s=591:634#L14)
<pre>func NewPackage(path, name <a href="/pkg/builtin/#string">string</a>) *<a href="#Package">Package</a></pre>
NewPackage returns a new Package for the given package path and name.
The package is not complete and contains no explicit imports.






### <a id="Package.Complete">func</a> (\*Package) [Complete](https://golang.org/src/go/types/package.go?s=1375:1410#L35)
<pre>func (pkg *<a href="#Package">Package</a>) Complete() <a href="/pkg/builtin/#bool">bool</a></pre>
A package is complete if its scope contains (at least) all
exported objects; otherwise it is incomplete.




### <a id="Package.Imports">func</a> (\*Package) [Imports](https://golang.org/src/go/types/package.go?s=1857:1897#L46)
<pre>func (pkg *<a href="#Package">Package</a>) Imports() []*<a href="#Package">Package</a></pre>
Imports returns the list of packages directly imported by
pkg; the list is in source order.

If pkg was loaded from export data, Imports includes packages that
provide package-level objects referenced by pkg. This may be more or
less than the set of packages directly imported by pkg's source code.




### <a id="Package.MarkComplete">func</a> (\*Package) [MarkComplete](https://golang.org/src/go/types/package.go?s=1481:1515#L38)
<pre>func (pkg *<a href="#Package">Package</a>) MarkComplete()</pre>
MarkComplete marks a package as complete.




### <a id="Package.Name">func</a> (\*Package) [Name](https://golang.org/src/go/types/package.go?s=906:939#L23)
<pre>func (pkg *<a href="#Package">Package</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the package name.




### <a id="Package.Path">func</a> (\*Package) [Path](https://golang.org/src/go/types/package.go?s=817:850#L20)
<pre>func (pkg *<a href="#Package">Package</a>) Path() <a href="/pkg/builtin/#string">string</a></pre>
Path returns the package path.




### <a id="Package.Scope">func</a> (\*Package) [Scope](https://golang.org/src/go/types/package.go?s=1207:1241#L31)
<pre>func (pkg *<a href="#Package">Package</a>) Scope() *<a href="#Scope">Scope</a></pre>
Scope returns the (complete or incomplete) package scope
holding the objects declared at package level (TypeNames,
Consts, Vars, and Funcs).




### <a id="Package.SetImports">func</a> (\*Package) [SetImports](https://golang.org/src/go/types/package.go?s=2067:2114#L50)
<pre>func (pkg *<a href="#Package">Package</a>) SetImports(list []*<a href="#Package">Package</a>)</pre>
SetImports sets the list of explicitly imported packages to list.
It is the caller's responsibility to make sure list elements are unique.




### <a id="Package.SetName">func</a> (\*Package) [SetName](https://golang.org/src/go/types/package.go?s=995:1035#L26)
<pre>func (pkg *<a href="#Package">Package</a>) SetName(name <a href="/pkg/builtin/#string">string</a>)</pre>
SetName sets the package name.




### <a id="Package.String">func</a> (\*Package) [String](https://golang.org/src/go/types/package.go?s=2139:2174#L52)
<pre>func (pkg *<a href="#Package">Package</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="PkgName">type</a> [PkgName](https://golang.org/src/go/types/object.go?s=5720:5816#L170)
A PkgName represents an imported Go package.
PkgNames don't have a type.


<pre>type PkgName struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewPkgName">func</a> [NewPkgName](https://golang.org/src/go/types/object.go?s=5965:6050#L178)
<pre>func NewPkgName(pos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>, pkg *<a href="#Package">Package</a>, name <a href="/pkg/builtin/#string">string</a>, imported *<a href="#Package">Package</a>) *<a href="#PkgName">PkgName</a></pre>
NewPkgName returns a new PkgName object representing an imported package.
The remaining arguments set the attributes found with all Objects.






### <a id="PkgName.Exported">func</a> (\*PkgName) [Exported](https://golang.org/src/go/types/object.go?s=4259:4293#L131)
<pre>func (obj *<a href="#PkgName">PkgName</a>) Exported() <a href="/pkg/builtin/#bool">bool</a></pre>
Exported reports whether the object is exported (starts with a capital letter).
It doesn't take into account whether the object is in a local (function) scope
or not.




### <a id="PkgName.Id">func</a> (\*PkgName) [Id](https://golang.org/src/go/types/object.go?s=4383:4413#L134)
<pre>func (obj *<a href="#PkgName">PkgName</a>) Id() <a href="/pkg/builtin/#string">string</a></pre>
Id is a wrapper for Id(obj.Pkg(), obj.Name()).




### <a id="PkgName.Imported">func</a> (\*PkgName) [Imported](https://golang.org/src/go/types/object.go?s=6291:6330#L184)
<pre>func (obj *<a href="#PkgName">PkgName</a>) Imported() *<a href="#Package">Package</a></pre>
Imported returns the package that was imported.
It is distinct from Pkg(), which is the package containing the import statement.




### <a id="PkgName.Name">func</a> (\*PkgName) [Name](https://golang.org/src/go/types/object.go?s=3943:3975#L123)
<pre>func (obj *<a href="#PkgName">PkgName</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the object's (package-local, unqualified) name.




### <a id="PkgName.Parent">func</a> (\*PkgName) [Parent](https://golang.org/src/go/types/object.go?s=3521:3555#L113)
<pre>func (obj *<a href="#PkgName">PkgName</a>) Parent() *<a href="#Scope">Scope</a></pre>
Parent returns the scope in which the object is declared.
The result is nil for methods and struct fields.




### <a id="PkgName.Pkg">func</a> (\*PkgName) [Pkg](https://golang.org/src/go/types/object.go?s=3825:3858#L120)
<pre>func (obj *<a href="#PkgName">PkgName</a>) Pkg() *<a href="#Package">Package</a></pre>
Pkg returns the package to which the object belongs.
The result is nil for labels and objects in the Universe scope.




### <a id="PkgName.Pos">func</a> (\*PkgName) [Pos](https://golang.org/src/go/types/object.go?s=3647:3681#L116)
<pre>func (obj *<a href="#PkgName">PkgName</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>
Pos returns the declaration position of the object's identifier.




### <a id="PkgName.String">func</a> (\*PkgName) [String](https://golang.org/src/go/types/object.go?s=13867:13902#L446)
<pre>func (obj *<a href="#PkgName">PkgName</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="PkgName.Type">func</a> (\*PkgName) [Type](https://golang.org/src/go/types/object.go?s=4032:4062#L126)
<pre>func (obj *<a href="#PkgName">PkgName</a>) Type() <a href="#Type">Type</a></pre>
Type returns the object's type.




## <a id="Pointer">type</a> [Pointer](https://golang.org/src/go/types/type.go?s=3671:3721#L149)
A Pointer represents a pointer type.


<pre>type Pointer struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewPointer">func</a> [NewPointer](https://golang.org/src/go/types/type.go?s=3799:3834#L154)
<pre>func NewPointer(elem <a href="#Type">Type</a>) *<a href="#Pointer">Pointer</a></pre>
NewPointer returns a new pointer type for the given element (base) type.






### <a id="Pointer.Elem">func</a> (\*Pointer) [Elem](https://golang.org/src/go/types/type.go?s=3926:3955#L157)
<pre>func (p *<a href="#Pointer">Pointer</a>) Elem() <a href="#Type">Type</a></pre>
Elem returns the element type for the given pointer p.




### <a id="Pointer.String">func</a> (\*Pointer) [String](https://golang.org/src/go/types/type.go?s=15795:15828#L474)
<pre>func (p *<a href="#Pointer">Pointer</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Pointer.Underlying">func</a> (\*Pointer) [Underlying](https://golang.org/src/go/types/type.go?s=15162:15197#L462)
<pre>func (p *<a href="#Pointer">Pointer</a>) Underlying() <a href="#Type">Type</a></pre>



## <a id="Qualifier">type</a> [Qualifier](https://golang.org/src/go/types/typestring.go?s=781:817#L15)
A Qualifier controls how named package-level objects are printed in
calls to TypeString, ObjectString, and SelectionString.

These three formatting routines call the Qualifier for each
package-level object O, and if the Qualifier returns a non-empty
string p, the object is printed in the form p.O.
If it returns an empty string, only the object name O is printed.

Using a nil Qualifier is equivalent to using (*Package).Path: the
object is qualified by the import path, e.g., "encoding/json.Marshal".


<pre>type Qualifier func(*<a href="#Package">Package</a>) <a href="/pkg/builtin/#string">string</a></pre>









### <a id="RelativeTo">func</a> [RelativeTo](https://golang.org/src/go/types/typestring.go?s=917:956#L19)
<pre>func RelativeTo(pkg *<a href="#Package">Package</a>) <a href="#Qualifier">Qualifier</a></pre>
RelativeTo returns a Qualifier that fully qualifies members of
all packages other than pkg.






## <a id="Scope">type</a> [Scope](https://golang.org/src/go/types/scope.go?s=493:791#L12)
A Scope maintains a set of objects and links to its containing
(parent) and contained (children) scopes. Objects may be inserted
and looked up by name. The zero value for Scope is a ready-to-use
empty scope.


<pre>type Scope struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>




The Universe scope contains all predeclared objects of Go.
It is the outermost scope of any chain of nested scopes.


<pre>var <span id="Universe">Universe</span> *<a href="#Scope">Scope</a></pre>

<a id="example_Scope">Example</a>
<p>ExampleScope prints the tree of Scopes of a package created from a
set of parsed files.
</p>



### <a id="NewScope">func</a> [NewScope](https://golang.org/src/go/types/scope.go?s=915:986#L23)
<pre>func NewScope(parent *<a href="#Scope">Scope</a>, pos, end <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>, comment <a href="/pkg/builtin/#string">string</a>) *<a href="#Scope">Scope</a></pre>
NewScope returns a new, empty scope contained in the given parent
scope, if any. The comment is for debugging only.






### <a id="Scope.Child">func</a> (\*Scope) [Child](https://golang.org/src/go/types/scope.go?s=1820:1855#L54)
<pre>func (s *<a href="#Scope">Scope</a>) Child(i <a href="/pkg/builtin/#int">int</a>) *<a href="#Scope">Scope</a></pre>
Child returns the i'th child scope for 0 <= i < NumChildren().




### <a id="Scope.Contains">func</a> (\*Scope) [Contains](https://golang.org/src/go/types/scope.go?s=3985:4029#L111)
<pre>func (s *<a href="#Scope">Scope</a>) Contains(pos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Contains reports whether pos is within the scope's extent.
The result is guaranteed to be valid only if the type-checked
AST has complete position information.




### <a id="Scope.End">func</a> (\*Scope) [End](https://golang.org/src/go/types/scope.go?s=3766:3797#L106)
<pre>func (s *<a href="#Scope">Scope</a>) End() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>



### <a id="Scope.Innermost">func</a> (\*Scope) [Innermost](https://golang.org/src/go/types/scope.go?s=4347:4394#L120)
<pre>func (s *<a href="#Scope">Scope</a>) Innermost(pos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>) *<a href="#Scope">Scope</a></pre>
Innermost returns the innermost (child) scope containing
pos. If pos is not within any scope, the result is nil.
The result is also nil for the Universe scope.
The result is guaranteed to be valid only if the type-checked
AST has complete position information.




### <a id="Scope.Insert">func</a> (\*Scope) [Insert](https://golang.org/src/go/types/scope.go?s=3218:3259#L86)
<pre>func (s *<a href="#Scope">Scope</a>) Insert(obj <a href="#Object">Object</a>) <a href="#Object">Object</a></pre>
Insert attempts to insert an object obj into scope s.
If s already contains an alternative object alt with
the same name, Insert leaves s unchanged and returns alt.
Otherwise it inserts obj, sets the object's parent scope
if not already set, and returns nil.




### <a id="Scope.Len">func</a> (\*Scope) [Len](https://golang.org/src/go/types/scope.go?s=1348:1373#L36)
<pre>func (s *<a href="#Scope">Scope</a>) Len() <a href="/pkg/builtin/#int">int</a></pre>
Len returns the number of scope elements.




### <a id="Scope.Lookup">func</a> (\*Scope) [Lookup](https://golang.org/src/go/types/scope.go?s=2000:2042#L58)
<pre>func (s *<a href="#Scope">Scope</a>) Lookup(name <a href="/pkg/builtin/#string">string</a>) <a href="#Object">Object</a></pre>
Lookup returns the object in scope s with the given name if such an
object exists; otherwise the result is nil.




### <a id="Scope.LookupParent">func</a> (\*Scope) [LookupParent](https://golang.org/src/go/types/scope.go?s=2707:2780#L72)
<pre>func (s *<a href="#Scope">Scope</a>) LookupParent(name <a href="/pkg/builtin/#string">string</a>, pos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>) (*<a href="#Scope">Scope</a>, <a href="#Object">Object</a>)</pre>
LookupParent follows the parent chain of scopes starting with s until
it finds a scope where Lookup(name) returns a non-nil object, and then
returns that scope and object. If a valid position pos is provided,
only objects that were declared at or before pos are considered.
If no such scope and object exists, the result is (nil, nil).

Note that obj.Parent() may be different from the returned scope if the
object was inserted into the scope and already had a parent at that
time (see Insert, below). This can only happen for dot-imported objects
whose scope is the scope of the package that exported them.




### <a id="Scope.Names">func</a> (\*Scope) [Names](https://golang.org/src/go/types/scope.go?s=1459:1491#L39)
<pre>func (s *<a href="#Scope">Scope</a>) Names() []<a href="/pkg/builtin/#string">string</a></pre>
Names returns the scope's element names in sorted order.




### <a id="Scope.NumChildren">func</a> (\*Scope) [NumChildren](https://golang.org/src/go/types/scope.go?s=1692:1725#L51)
<pre>func (s *<a href="#Scope">Scope</a>) NumChildren() <a href="/pkg/builtin/#int">int</a></pre>
NumChildren returns the number of scopes nested in s.




### <a id="Scope.Parent">func</a> (\*Scope) [Parent](https://golang.org/src/go/types/scope.go?s=1250:1281#L33)
<pre>func (s *<a href="#Scope">Scope</a>) Parent() *<a href="#Scope">Scope</a></pre>
Parent returns the scope's containing (parent) scope.




### <a id="Scope.Pos">func</a> (\*Scope) [Pos](https://golang.org/src/go/types/scope.go?s=3717:3748#L105)
<pre>func (s *<a href="#Scope">Scope</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>
Pos and End describe the scope's source code extent [pos, end).
The results are guaranteed to be valid only if the type-checked
AST has complete position information. The extent is undefined
for Universe and package scopes.




### <a id="Scope.String">func</a> (\*Scope) [String](https://golang.org/src/go/types/scope.go?s=5516:5547#L168)
<pre>func (s *<a href="#Scope">Scope</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns a string representation of the scope, for debugging.




### <a id="Scope.WriteTo">func</a> (\*Scope) [WriteTo](https://golang.org/src/go/types/scope.go?s=5048:5105#L147)
<pre>func (s *<a href="#Scope">Scope</a>) WriteTo(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, n <a href="/pkg/builtin/#int">int</a>, recurse <a href="/pkg/builtin/#bool">bool</a>)</pre>
WriteTo writes a string representation of the scope to w,
with the scope elements sorted by name.
The level of indentation is controlled by n >= 0, with
n == 0 for no indentation.
If recurse is set, it also writes nested (children) scopes.




## <a id="Selection">type</a> [Selection](https://golang.org/src/go/types/selection.go?s=1093:1326#L30)
A Selection describes a selector expression x.f.
For the declarations:


	type T struct{ x int; E }
	type E struct{}
	func (e E) m() {}
	var p *T

the following relations exist:


	Selector    Kind          Recv    Obj    Type               Index     Indirect
	
	p.x         FieldVal      T       x      int                {0}       true
	p.m         MethodVal     *T      m      func (e *T) m()    {1, 0}    true
	T.m         MethodExpr    T       m      func m(_ T)        {1, 0}    false


<pre>type Selection struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="Selection.Index">func</a> (\*Selection) [Index](https://golang.org/src/go/types/selection.go?s=3090:3123#L92)
<pre>func (s *<a href="#Selection">Selection</a>) Index() []<a href="/pkg/builtin/#int">int</a></pre>
Index describes the path from x to f in x.f.
The last index entry is the field or method index of the type declaring f;
either:


	1) the list of declared methods of a named type; or
	2) the list of methods of an interface type; or
	3) the list of fields of a struct type.

The earlier index entries are the indices of the embedded fields implicitly
traversed to get from (the type of) x to f, starting at embedding depth 0.




### <a id="Selection.Indirect">func</a> (\*Selection) [Indirect](https://golang.org/src/go/types/selection.go?s=3239:3274#L96)
<pre>func (s *<a href="#Selection">Selection</a>) Indirect() <a href="/pkg/builtin/#bool">bool</a></pre>
Indirect reports whether any pointer indirection was required to get from
x to f in x.f.




### <a id="Selection.Kind">func</a> (\*Selection) [Kind](https://golang.org/src/go/types/selection.go?s=1364:1404#L39)
<pre>func (s *<a href="#Selection">Selection</a>) Kind() <a href="#SelectionKind">SelectionKind</a></pre>
Kind returns the selection kind.




### <a id="Selection.Obj">func</a> (\*Selection) [Obj](https://golang.org/src/go/types/selection.go?s=1620:1652#L46)
<pre>func (s *<a href="#Selection">Selection</a>) Obj() <a href="#Object">Object</a></pre>
Obj returns the object denoted by x.f; a *Var for
a field selection, and a *Func in all other cases.




### <a id="Selection.Recv">func</a> (\*Selection) [Recv](https://golang.org/src/go/types/selection.go?s=1462:1493#L42)
<pre>func (s *<a href="#Selection">Selection</a>) Recv() <a href="#Type">Type</a></pre>
Recv returns the type of x in x.f.




### <a id="Selection.String">func</a> (\*Selection) [String](https://golang.org/src/go/types/selection.go?s=3298:3333#L98)
<pre>func (s *<a href="#Selection">Selection</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Selection.Type">func</a> (\*Selection) [Type](https://golang.org/src/go/types/selection.go?s=1786:1817#L50)
<pre>func (s *<a href="#Selection">Selection</a>) Type() <a href="#Type">Type</a></pre>
Type returns the type of x.f, which may be different from the type of f.
See Selection for more information.




## <a id="SelectionKind">type</a> [SelectionKind](https://golang.org/src/go/types/selection.go?s=343:365#L6)
SelectionKind describes the kind of a selector expression x.f
(excluding qualified identifiers).


<pre>type SelectionKind <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="FieldVal">FieldVal</span>   <a href="#SelectionKind">SelectionKind</a> = <a href="/pkg/builtin/#iota">iota</a> <span class="comment">// x.f is a struct field selector</span>
    <span id="MethodVal">MethodVal</span>                       <span class="comment">// x.f is a method selector</span>
    <span id="MethodExpr">MethodExpr</span>                      <span class="comment">// x.f is a method expression</span>
)</pre>









## <a id="Signature">type</a> [Signature](https://golang.org/src/go/types/type.go?s=4745:5428#L187)
A Signature represents a (non-builtin) function or method type.
The receiver is ignored when comparing signatures for identity.


<pre>type Signature struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewSignature">func</a> [NewSignature](https://golang.org/src/go/types/type.go?s=5697:5775#L203)
<pre>func NewSignature(recv *<a href="#Var">Var</a>, params, results *<a href="#Tuple">Tuple</a>, variadic <a href="/pkg/builtin/#bool">bool</a>) *<a href="#Signature">Signature</a></pre>
NewSignature returns a new function type for the given receiver, parameters,
and results, either of which may be nil. If variadic is set, the function
is variadic, it must have at least one parameter, and the last parameter
must be of unnamed slice type.






### <a id="Signature.Params">func</a> (\*Signature) [Params](https://golang.org/src/go/types/type.go?s=6566:6601#L225)
<pre>func (s *<a href="#Signature">Signature</a>) Params() *<a href="#Tuple">Tuple</a></pre>
Params returns the parameters of signature s, or nil.




### <a id="Signature.Recv">func</a> (\*Signature) [Recv](https://golang.org/src/go/types/type.go?s=6458:6489#L222)
<pre>func (s *<a href="#Signature">Signature</a>) Recv() *<a href="#Var">Var</a></pre>
Recv returns the receiver of signature s (if a method), or nil if a
function. It is ignored when comparing signatures for identity.

For an abstract method, Recv returns the enclosing interface either
as a *Named or an *Interface. Due to embedding, an interface may
contain methods whose receiver type is a different interface.




### <a id="Signature.Results">func</a> (\*Signature) [Results](https://golang.org/src/go/types/type.go?s=6678:6714#L228)
<pre>func (s *<a href="#Signature">Signature</a>) Results() *<a href="#Tuple">Tuple</a></pre>
Results returns the results of signature s, or nil.




### <a id="Signature.String">func</a> (\*Signature) [String](https://golang.org/src/go/types/type.go?s=15927:15962#L476)
<pre>func (s *<a href="#Signature">Signature</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Signature.Underlying">func</a> (\*Signature) [Underlying](https://golang.org/src/go/types/type.go?s=15264:15301#L464)
<pre>func (s *<a href="#Signature">Signature</a>) Underlying() <a href="#Type">Type</a></pre>



### <a id="Signature.Variadic">func</a> (\*Signature) [Variadic](https://golang.org/src/go/types/type.go?s=6794:6829#L231)
<pre>func (s *<a href="#Signature">Signature</a>) Variadic() <a href="/pkg/builtin/#bool">bool</a></pre>
Variadic reports whether the signature s is variadic.




## <a id="Sizes">type</a> [Sizes](https://golang.org/src/go/types/sizes.go?s=265:769#L1)
Sizes defines the sizing functions for package unsafe.


<pre>type Sizes interface {
    <span class="comment">// Alignof returns the alignment of a variable of type T.</span>
    <span class="comment">// Alignof must implement the alignment guarantees required by the spec.</span>
    Alignof(T <a href="#Type">Type</a>) <a href="/pkg/builtin/#int64">int64</a>

    <span class="comment">// Offsetsof returns the offsets of the given struct fields, in bytes.</span>
    <span class="comment">// Offsetsof must implement the offset guarantees required by the spec.</span>
    Offsetsof(fields []*<a href="#Var">Var</a>) []<a href="/pkg/builtin/#int64">int64</a>

    <span class="comment">// Sizeof returns the size of a variable of type T.</span>
    <span class="comment">// Sizeof must implement the size guarantees required by the spec.</span>
    Sizeof(T <a href="#Type">Type</a>) <a href="/pkg/builtin/#int64">int64</a>
}</pre>









### <a id="SizesFor">func</a> [SizesFor](https://golang.org/src/go/types/sizes.go?s=4997:5039#L174)
<pre>func SizesFor(compiler, arch <a href="/pkg/builtin/#string">string</a>) <a href="#Sizes">Sizes</a></pre>
SizesFor returns the Sizes used by a compiler for an architecture.
The result is nil if a compiler/architecture pair is not known.

Supported architectures for compiler "gc":
"386", "arm", "arm64", "amd64", "amd64p32", "mips", "mipsle",
"mips64", "mips64le", "ppc64", "ppc64le", "riscv64", "s390x", "sparc64", "wasm".






## <a id="Slice">type</a> [Slice](https://golang.org/src/go/types/type.go?s=2212:2244#L101)
A Slice represents a slice type.


<pre>type Slice struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewSlice">func</a> [NewSlice](https://golang.org/src/go/types/type.go?s=2311:2342#L106)
<pre>func NewSlice(elem <a href="#Type">Type</a>) *<a href="#Slice">Slice</a></pre>
NewSlice returns a new slice type for the given element type.






### <a id="Slice.Elem">func</a> (\*Slice) [Elem](https://golang.org/src/go/types/type.go?s=2413:2440#L109)
<pre>func (s *<a href="#Slice">Slice</a>) Elem() <a href="#Type">Type</a></pre>
Elem returns the element type of slice s.




### <a id="Slice.String">func</a> (\*Slice) [String](https://golang.org/src/go/types/type.go?s=15663:15694#L472)
<pre>func (s *<a href="#Slice">Slice</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Slice.Underlying">func</a> (\*Slice) [Underlying](https://golang.org/src/go/types/type.go?s=15060:15093#L460)
<pre>func (s *<a href="#Slice">Slice</a>) Underlying() <a href="#Type">Type</a></pre>



## <a id="StdSizes">type</a> [StdSizes](https://golang.org/src/go/types/sizes.go?s=1657:1804#L33)
StdSizes is a convenience type for creating commonly used Sizes.
It makes the following simplifying assumptions:


		- The size of explicitly sized basic types (int16, etc.) is the
		  specified size.
		- The size of strings and interfaces is 2*WordSize.
		- The size of slices is 3*WordSize.
		- The size of an array of n elements corresponds to the size of
		  a struct of n consecutive fields of the array's element type.
	     - The size of a struct is the offset of the last field plus that
		  field's size. As with all element types, if the struct is used
		  in an array its size must first be aligned to a multiple of the
		  struct's alignment.
		- All other types have size WordSize.
		- Arrays and structs are aligned per spec definition; all other
		  types are naturally aligned with a maximum alignment MaxAlign.

*StdSizes implements Sizes.


<pre>type StdSizes struct {
<span id="StdSizes.WordSize"></span>    WordSize <a href="/pkg/builtin/#int64">int64</a> <span class="comment">// word size in bytes - must be &gt;= 4 (32bits)</span>
<span id="StdSizes.MaxAlign"></span>    MaxAlign <a href="/pkg/builtin/#int64">int64</a> <span class="comment">// maximum alignment in bytes - must be &gt;= 1</span>
}
</pre>











### <a id="StdSizes.Alignof">func</a> (\*StdSizes) [Alignof](https://golang.org/src/go/types/sizes.go?s=1806:1846#L38)
<pre>func (s *<a href="#StdSizes">StdSizes</a>) Alignof(T <a href="#Type">Type</a>) <a href="/pkg/builtin/#int64">int64</a></pre>



### <a id="StdSizes.Offsetsof">func</a> (\*StdSizes) [Offsetsof](https://golang.org/src/go/types/sizes.go?s=3001:3052#L82)
<pre>func (s *<a href="#StdSizes">StdSizes</a>) Offsetsof(fields []*<a href="#Var">Var</a>) []<a href="/pkg/builtin/#int64">int64</a></pre>



### <a id="StdSizes.Sizeof">func</a> (\*StdSizes) [Sizeof](https://golang.org/src/go/types/sizes.go?s=3479:3518#L110)
<pre>func (s *<a href="#StdSizes">StdSizes</a>) Sizeof(T <a href="#Type">Type</a>) <a href="/pkg/builtin/#int64">int64</a></pre>



## <a id="Struct">type</a> [Struct](https://golang.org/src/go/types/type.go?s=2498:2592#L112)
A Struct represents a struct type.


<pre>type Struct struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewStruct">func</a> [NewStruct](https://golang.org/src/go/types/type.go?s=2892:2944#L121)
<pre>func NewStruct(fields []*<a href="#Var">Var</a>, tags []<a href="/pkg/builtin/#string">string</a>) *<a href="#Struct">Struct</a></pre>
NewStruct returns a new struct with the given fields and corresponding field tags.
If a field with index i has a tag, tags[i] must be that tag, but len(tags) may be
only as long as required to hold the tag with the largest index i. Consequently,
if no field has a tag, tags may be nil.






### <a id="Struct.Field">func</a> (\*Struct) [Field](https://golang.org/src/go/types/type.go?s=3417:3451#L138)
<pre>func (s *<a href="#Struct">Struct</a>) Field(i <a href="/pkg/builtin/#int">int</a>) *<a href="#Var">Var</a></pre>
Field returns the i'th field for 0 <= i < NumFields().




### <a id="Struct.NumFields">func</a> (\*Struct) [NumFields](https://golang.org/src/go/types/type.go?s=3300:3332#L135)
<pre>func (s *<a href="#Struct">Struct</a>) NumFields() <a href="/pkg/builtin/#int">int</a></pre>
NumFields returns the number of fields in the struct (including blank and embedded fields).




### <a id="Struct.String">func</a> (\*Struct) [String](https://golang.org/src/go/types/type.go?s=15729:15761#L473)
<pre>func (s *<a href="#Struct">Struct</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Struct.Tag">func</a> (\*Struct) [Tag](https://golang.org/src/go/types/type.go?s=3536:3570#L141)
<pre>func (s *<a href="#Struct">Struct</a>) Tag(i <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#string">string</a></pre>
Tag returns the i'th field tag for 0 <= i < NumFields().




### <a id="Struct.Underlying">func</a> (\*Struct) [Underlying](https://golang.org/src/go/types/type.go?s=15111:15145#L461)
<pre>func (s *<a href="#Struct">Struct</a>) Underlying() <a href="#Type">Type</a></pre>



## <a id="Tuple">type</a> [Tuple](https://golang.org/src/go/types/type.go?s=4205:4239#L162)
A Tuple represents an ordered list of variables; a nil *Tuple is a valid (empty) tuple.
Tuples are used as components of signatures and to represent the type of multiple
assignments; they are not first class types of Go.


<pre>type Tuple struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewTuple">func</a> [NewTuple](https://golang.org/src/go/types/type.go?s=4298:4329#L167)
<pre>func NewTuple(x ...*<a href="#Var">Var</a>) *<a href="#Tuple">Tuple</a></pre>
NewTuple returns a new tuple for the given variables.






### <a id="Tuple.At">func</a> (\*Tuple) [At](https://golang.org/src/go/types/type.go?s=4558:4588#L183)
<pre>func (t *<a href="#Tuple">Tuple</a>) At(i <a href="/pkg/builtin/#int">int</a>) *<a href="#Var">Var</a></pre>
At returns the i'th variable of tuple t.




### <a id="Tuple.Len">func</a> (\*Tuple) [Len](https://golang.org/src/go/types/type.go?s=4434:4459#L175)
<pre>func (t *<a href="#Tuple">Tuple</a>) Len() <a href="/pkg/builtin/#int">int</a></pre>
Len returns the number variables of tuple t.




### <a id="Tuple.String">func</a> (\*Tuple) [String](https://golang.org/src/go/types/type.go?s=15861:15892#L475)
<pre>func (t *<a href="#Tuple">Tuple</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Tuple.Underlying">func</a> (\*Tuple) [Underlying](https://golang.org/src/go/types/type.go?s=15213:15246#L463)
<pre>func (t *<a href="#Tuple">Tuple</a>) Underlying() <a href="#Type">Type</a></pre>



## <a id="Type">type</a> [Type](https://golang.org/src/go/types/type.go?s=268:436#L1)
A Type represents a type of Go.
All types implement the Type interface.


<pre>type Type interface {
    <span class="comment">// Underlying returns the underlying type of a type.</span>
    Underlying() <a href="#Type">Type</a>

    <span class="comment">// String returns a string representation of a type.</span>
    String() <a href="/pkg/builtin/#string">string</a>
}</pre>









### <a id="Default">func</a> [Default](https://golang.org/src/go/types/predicates.go?s=8369:8396#L292)
<pre>func Default(typ <a href="#Type">Type</a>) <a href="#Type">Type</a></pre>
Default returns the default "typed" type for an "untyped" type;
it returns the incoming type for all other types. The default type
for untyped nil is untyped nil.






## <a id="TypeAndValue">type</a> [TypeAndValue](https://golang.org/src/go/types/api.go?s=9761:9842#L246)
TypeAndValue reports the type and value (for constants)
of the corresponding expression.


<pre>type TypeAndValue struct {
<span id="TypeAndValue.Type"></span>    Type  <a href="#Type">Type</a>
<span id="TypeAndValue.Value"></span>    Value <a href="/pkg/go/constant/">constant</a>.<a href="/pkg/go/constant/#Value">Value</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="Eval">func</a> [Eval](https://golang.org/src/go/types/eval.go?s=689:789#L14)
<pre>func Eval(fset *<a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#FileSet">FileSet</a>, pkg *<a href="#Package">Package</a>, pos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>, expr <a href="/pkg/builtin/#string">string</a>) (_ <a href="#TypeAndValue">TypeAndValue</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Eval returns the type and, if constant, the value for the
expression expr, evaluated at position pos of package pkg,
which must have been derived from type-checking an AST with
complete position information relative to the provided file
set.

The meaning of the parameters fset, pkg, and pos is the
same as in CheckExpr. An error is returned if expr cannot
be parsed successfully, or the resulting expr AST cannot be
type-checked.






### <a id="TypeAndValue.Addressable">func</a> (TypeAndValue) [Addressable](https://golang.org/src/go/types/api.go?s=11084:11125#L291)
<pre>func (tv <a href="#TypeAndValue">TypeAndValue</a>) Addressable() <a href="/pkg/builtin/#bool">bool</a></pre>
Addressable reports whether the corresponding expression
is addressable (<a href="https://golang.org/ref/spec#Address_operators">https://golang.org/ref/spec#Address_operators</a>).




### <a id="TypeAndValue.Assignable">func</a> (TypeAndValue) [Assignable](https://golang.org/src/go/types/api.go?s=11276:11316#L297)
<pre>func (tv <a href="#TypeAndValue">TypeAndValue</a>) Assignable() <a href="/pkg/builtin/#bool">bool</a></pre>
Assignable reports whether the corresponding expression
is assignable to (provided a value of the right type).




### <a id="TypeAndValue.HasOk">func</a> (TypeAndValue) [HasOk](https://golang.org/src/go/types/api.go?s=11479:11514#L303)
<pre>func (tv <a href="#TypeAndValue">TypeAndValue</a>) HasOk() <a href="/pkg/builtin/#bool">bool</a></pre>
HasOk reports whether the corresponding expression may be
used on the rhs of a comma-ok assignment.




### <a id="TypeAndValue.IsBuiltin">func</a> (TypeAndValue) [IsBuiltin](https://golang.org/src/go/types/api.go?s=10403:10442#L268)
<pre>func (tv <a href="#TypeAndValue">TypeAndValue</a>) IsBuiltin() <a href="/pkg/builtin/#bool">bool</a></pre>
IsBuiltin reports whether the corresponding expression denotes
a (possibly parenthesized) built-in function.




### <a id="TypeAndValue.IsNil">func</a> (TypeAndValue) [IsNil](https://golang.org/src/go/types/api.go?s=10861:10896#L285)
<pre>func (tv <a href="#TypeAndValue">TypeAndValue</a>) IsNil() <a href="/pkg/builtin/#bool">bool</a></pre>
IsNil reports whether the corresponding expression denotes the
predeclared value nil.




### <a id="TypeAndValue.IsType">func</a> (TypeAndValue) [IsType](https://golang.org/src/go/types/api.go?s=10219:10255#L262)
<pre>func (tv <a href="#TypeAndValue">TypeAndValue</a>) IsType() <a href="/pkg/builtin/#bool">bool</a></pre>
IsType reports whether the corresponding expression specifies a type.




### <a id="TypeAndValue.IsValue">func</a> (TypeAndValue) [IsValue](https://golang.org/src/go/types/api.go?s=10624:10661#L275)
<pre>func (tv <a href="#TypeAndValue">TypeAndValue</a>) IsValue() <a href="/pkg/builtin/#bool">bool</a></pre>
IsValue reports whether the corresponding expression is a value.
Builtins are not considered values. Constant values have a non-
nil Value.




### <a id="TypeAndValue.IsVoid">func</a> (TypeAndValue) [IsVoid](https://golang.org/src/go/types/api.go?s=10077:10113#L257)
<pre>func (tv <a href="#TypeAndValue">TypeAndValue</a>) IsVoid() <a href="/pkg/builtin/#bool">bool</a></pre>
IsVoid reports whether the corresponding expression
is a function call without results.




## <a id="TypeName">type</a> [TypeName](https://golang.org/src/go/types/object.go?s=7012:7044#L204)
A TypeName represents a name for a (defined or alias) type.


<pre>type TypeName struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewTypeName">func</a> [NewTypeName](https://golang.org/src/go/types/object.go?s=7402:7480#L215)
<pre>func NewTypeName(pos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>, pkg *<a href="#Package">Package</a>, name <a href="/pkg/builtin/#string">string</a>, typ <a href="#Type">Type</a>) *<a href="#TypeName">TypeName</a></pre>
NewTypeName returns a new type name denoting the given typ.
The remaining arguments set the attributes found with all Objects.

The typ argument may be a defined (Named) type or an alias type.
It may also be nil such that the returned TypeName can be used as
argument for NewNamed, which will set the TypeName's type as a side-
effect.






### <a id="TypeName.Exported">func</a> (\*TypeName) [Exported](https://golang.org/src/go/types/object.go?s=4259:4293#L131)
<pre>func (obj *<a href="#TypeName">TypeName</a>) Exported() <a href="/pkg/builtin/#bool">bool</a></pre>
Exported reports whether the object is exported (starts with a capital letter).
It doesn't take into account whether the object is in a local (function) scope
or not.




### <a id="TypeName.Id">func</a> (\*TypeName) [Id](https://golang.org/src/go/types/object.go?s=4383:4413#L134)
<pre>func (obj *<a href="#TypeName">TypeName</a>) Id() <a href="/pkg/builtin/#string">string</a></pre>
Id is a wrapper for Id(obj.Pkg(), obj.Name()).




### <a id="TypeName.IsAlias">func</a> (\*TypeName) [IsAlias](https://golang.org/src/go/types/object.go?s=7629:7664#L220)
<pre>func (obj *<a href="#TypeName">TypeName</a>) IsAlias() <a href="/pkg/builtin/#bool">bool</a></pre>
IsAlias reports whether obj is an alias name for a type.




### <a id="TypeName.Name">func</a> (\*TypeName) [Name](https://golang.org/src/go/types/object.go?s=3943:3975#L123)
<pre>func (obj *<a href="#TypeName">TypeName</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the object's (package-local, unqualified) name.




### <a id="TypeName.Parent">func</a> (\*TypeName) [Parent](https://golang.org/src/go/types/object.go?s=3521:3555#L113)
<pre>func (obj *<a href="#TypeName">TypeName</a>) Parent() *<a href="#Scope">Scope</a></pre>
Parent returns the scope in which the object is declared.
The result is nil for methods and struct fields.




### <a id="TypeName.Pkg">func</a> (\*TypeName) [Pkg](https://golang.org/src/go/types/object.go?s=3825:3858#L120)
<pre>func (obj *<a href="#TypeName">TypeName</a>) Pkg() *<a href="#Package">Package</a></pre>
Pkg returns the package to which the object belongs.
The result is nil for labels and objects in the Universe scope.




### <a id="TypeName.Pos">func</a> (\*TypeName) [Pos](https://golang.org/src/go/types/object.go?s=3647:3681#L116)
<pre>func (obj *<a href="#TypeName">TypeName</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>
Pos returns the declaration position of the object's identifier.




### <a id="TypeName.String">func</a> (\*TypeName) [String](https://golang.org/src/go/types/object.go?s=14009:14045#L448)
<pre>func (obj *<a href="#TypeName">TypeName</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="TypeName.Type">func</a> (\*TypeName) [Type](https://golang.org/src/go/types/object.go?s=4032:4062#L126)
<pre>func (obj *<a href="#TypeName">TypeName</a>) Type() <a href="#Type">Type</a></pre>
Type returns the object's type.




## <a id="Var">type</a> [Var](https://golang.org/src/go/types/object.go?s=8496:8702#L244)
A Variable represents a declared variable (including function parameters and results, and struct fields).


<pre>type Var struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewField">func</a> [NewField](https://golang.org/src/go/types/object.go?s=9401:9486#L265)
<pre>func NewField(pos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>, pkg *<a href="#Package">Package</a>, name <a href="/pkg/builtin/#string">string</a>, typ <a href="#Type">Type</a>, embedded <a href="/pkg/builtin/#bool">bool</a>) *<a href="#Var">Var</a></pre>
NewField returns a new variable representing a struct field.
For embedded fields, the name is the unqualified type name
/ under which the field is accessible.




### <a id="NewParam">func</a> [NewParam](https://golang.org/src/go/types/object.go?s=9028:9098#L258)
<pre>func NewParam(pos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>, pkg *<a href="#Package">Package</a>, name <a href="/pkg/builtin/#string">string</a>, typ <a href="#Type">Type</a>) *<a href="#Var">Var</a></pre>
NewParam returns a new variable representing a function parameter.




### <a id="NewVar">func</a> [NewVar](https://golang.org/src/go/types/object.go?s=8798:8866#L253)
<pre>func NewVar(pos <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a>, pkg *<a href="#Package">Package</a>, name <a href="/pkg/builtin/#string">string</a>, typ <a href="#Type">Type</a>) *<a href="#Var">Var</a></pre>
NewVar returns a new variable.
The arguments set the attributes found with all Objects.






### <a id="Var.Anonymous">func</a> (\*Var) [Anonymous](https://golang.org/src/go/types/object.go?s=9739:9771#L271)
<pre>func (obj *<a href="#Var">Var</a>) Anonymous() <a href="/pkg/builtin/#bool">bool</a></pre>
Anonymous reports whether the variable is an embedded field.
Same as Embedded; only present for backward-compatibility.




### <a id="Var.Embedded">func</a> (\*Var) [Embedded](https://golang.org/src/go/types/object.go?s=9860:9891#L274)
<pre>func (obj *<a href="#Var">Var</a>) Embedded() <a href="/pkg/builtin/#bool">bool</a></pre>
Embedded reports whether the variable is an embedded field.




### <a id="Var.Exported">func</a> (\*Var) [Exported](https://golang.org/src/go/types/object.go?s=4259:4293#L131)
<pre>func (obj *<a href="#Var">Var</a>) Exported() <a href="/pkg/builtin/#bool">bool</a></pre>
Exported reports whether the object is exported (starts with a capital letter).
It doesn't take into account whether the object is in a local (function) scope
or not.




### <a id="Var.Id">func</a> (\*Var) [Id](https://golang.org/src/go/types/object.go?s=4383:4413#L134)
<pre>func (obj *<a href="#Var">Var</a>) Id() <a href="/pkg/builtin/#string">string</a></pre>
Id is a wrapper for Id(obj.Pkg(), obj.Name()).




### <a id="Var.IsField">func</a> (\*Var) [IsField](https://golang.org/src/go/types/object.go?s=9976:10006#L277)
<pre>func (obj *<a href="#Var">Var</a>) IsField() <a href="/pkg/builtin/#bool">bool</a></pre>
IsField reports whether the variable is a struct field.




### <a id="Var.Name">func</a> (\*Var) [Name](https://golang.org/src/go/types/object.go?s=3943:3975#L123)
<pre>func (obj *<a href="#Var">Var</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the object's (package-local, unqualified) name.




### <a id="Var.Parent">func</a> (\*Var) [Parent](https://golang.org/src/go/types/object.go?s=3521:3555#L113)
<pre>func (obj *<a href="#Var">Var</a>) Parent() *<a href="#Scope">Scope</a></pre>
Parent returns the scope in which the object is declared.
The result is nil for methods and struct fields.




### <a id="Var.Pkg">func</a> (\*Var) [Pkg](https://golang.org/src/go/types/object.go?s=3825:3858#L120)
<pre>func (obj *<a href="#Var">Var</a>) Pkg() *<a href="#Package">Package</a></pre>
Pkg returns the package to which the object belongs.
The result is nil for labels and objects in the Universe scope.




### <a id="Var.Pos">func</a> (\*Var) [Pos](https://golang.org/src/go/types/object.go?s=3647:3681#L116)
<pre>func (obj *<a href="#Var">Var</a>) Pos() <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Pos">Pos</a></pre>
Pos returns the declaration position of the object's identifier.




### <a id="Var.String">func</a> (\*Var) [String](https://golang.org/src/go/types/object.go?s=14080:14111#L449)
<pre>func (obj *<a href="#Var">Var</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Var.Type">func</a> (\*Var) [Type](https://golang.org/src/go/types/object.go?s=4032:4062#L126)
<pre>func (obj *<a href="#Var">Var</a>) Type() <a href="#Type">Type</a></pre>
Type returns the object's type.








