

# reflect
`import "reflect"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
reflect 包了实现运行时反射, 允许程序操纵任意类型的对象. 典型的用法是用静态类型interface{}保存值并通过调用TypeOf返回Type类型来获取其动态类型信息.

ValueOf调用会返回一个代表运行时数据的Value. Zero 传入 Type 会返回表示该类型零值的 Value.

请参阅"Reflection法则"以了解 Go 中的反射：<a href="https://golang.org/doc/articles/laws_of_reflection.html">https://golang.org/doc/articles/laws_of_reflection.html</a>

## <a id="pkg-index">Index</a>
* [func Copy(dst, src Value) int](#Copy)
* [func DeepEqual(x, y interface{}) bool](#DeepEqual)
* [func Swapper(slice interface{}) func(i, j int)](#Swapper)
* [type ChanDir](#ChanDir)
  * [func (d ChanDir) String() string](#ChanDir.String)
* [type Kind](#Kind)
  * [func (k Kind) String() string](#Kind.String)
* [type MapIter](#MapIter)
  * [func (it *MapIter) Key() Value](#MapIter.Key)
  * [func (it *MapIter) Next() bool](#MapIter.Next)
  * [func (it *MapIter) Value() Value](#MapIter.Value)
* [type Method](#Method)
* [type SelectCase](#SelectCase)
* [type SelectDir](#SelectDir)
* [type SliceHeader](#SliceHeader)
* [type StringHeader](#StringHeader)
* [type StructField](#StructField)
* [type StructTag](#StructTag)
  * [func (tag StructTag) Get(key string) string](#StructTag.Get)
  * [func (tag StructTag) Lookup(key string) (value string, ok bool)](#StructTag.Lookup)
* [type Type](#Type)
  * [func ArrayOf(count int, elem Type) Type](#ArrayOf)
  * [func ChanOf(dir ChanDir, t Type) Type](#ChanOf)
  * [func FuncOf(in, out []Type, variadic bool) Type](#FuncOf)
  * [func MapOf(key, elem Type) Type](#MapOf)
  * [func PtrTo(t Type) Type](#PtrTo)
  * [func SliceOf(t Type) Type](#SliceOf)
  * [func StructOf(fields []StructField) Type](#StructOf)
  * [func TypeOf(i interface{}) Type](#TypeOf)
* [type Value](#Value)
  * [func Append(s Value, x ...Value) Value](#Append)
  * [func AppendSlice(s, t Value) Value](#AppendSlice)
  * [func Indirect(v Value) Value](#Indirect)
  * [func MakeChan(typ Type, buffer int) Value](#MakeChan)
  * [func MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value](#MakeFunc)
  * [func MakeMap(typ Type) Value](#MakeMap)
  * [func MakeMapWithSize(typ Type, n int) Value](#MakeMapWithSize)
  * [func MakeSlice(typ Type, len, cap int) Value](#MakeSlice)
  * [func New(typ Type) Value](#New)
  * [func NewAt(typ Type, p unsafe.Pointer) Value](#NewAt)
  * [func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)](#Select)
  * [func ValueOf(i interface{}) Value](#ValueOf)
  * [func Zero(typ Type) Value](#Zero)
  * [func (v Value) Addr() Value](#Value.Addr)
  * [func (v Value) Bool() bool](#Value.Bool)
  * [func (v Value) Bytes() []byte](#Value.Bytes)
  * [func (v Value) Call(in []Value) []Value](#Value.Call)
  * [func (v Value) CallSlice(in []Value) []Value](#Value.CallSlice)
  * [func (v Value) CanAddr() bool](#Value.CanAddr)
  * [func (v Value) CanInterface() bool](#Value.CanInterface)
  * [func (v Value) CanSet() bool](#Value.CanSet)
  * [func (v Value) Cap() int](#Value.Cap)
  * [func (v Value) Close()](#Value.Close)
  * [func (v Value) Complex() complex128](#Value.Complex)
  * [func (v Value) Convert(t Type) Value](#Value.Convert)
  * [func (v Value) Elem() Value](#Value.Elem)
  * [func (v Value) Field(i int) Value](#Value.Field)
  * [func (v Value) FieldByIndex(index []int) Value](#Value.FieldByIndex)
  * [func (v Value) FieldByName(name string) Value](#Value.FieldByName)
  * [func (v Value) FieldByNameFunc(match func(string) bool) Value](#Value.FieldByNameFunc)
  * [func (v Value) Float() float64](#Value.Float)
  * [func (v Value) Index(i int) Value](#Value.Index)
  * [func (v Value) Int() int64](#Value.Int)
  * [func (v Value) Interface() (i interface{})](#Value.Interface)
  * [func (v Value) InterfaceData() [2]uintptr](#Value.InterfaceData)
  * [func (v Value) IsNil() bool](#Value.IsNil)
  * [func (v Value) IsValid() bool](#Value.IsValid)
  * [func (v Value) IsZero() bool](#Value.IsZero)
  * [func (v Value) Kind() Kind](#Value.Kind)
  * [func (v Value) Len() int](#Value.Len)
  * [func (v Value) MapIndex(key Value) Value](#Value.MapIndex)
  * [func (v Value) MapKeys() []Value](#Value.MapKeys)
  * [func (v Value) MapRange() *MapIter](#Value.MapRange)
  * [func (v Value) Method(i int) Value](#Value.Method)
  * [func (v Value) MethodByName(name string) Value](#Value.MethodByName)
  * [func (v Value) NumField() int](#Value.NumField)
  * [func (v Value) NumMethod() int](#Value.NumMethod)
  * [func (v Value) OverflowComplex(x complex128) bool](#Value.OverflowComplex)
  * [func (v Value) OverflowFloat(x float64) bool](#Value.OverflowFloat)
  * [func (v Value) OverflowInt(x int64) bool](#Value.OverflowInt)
  * [func (v Value) OverflowUint(x uint64) bool](#Value.OverflowUint)
  * [func (v Value) Pointer() uintptr](#Value.Pointer)
  * [func (v Value) Recv() (x Value, ok bool)](#Value.Recv)
  * [func (v Value) Send(x Value)](#Value.Send)
  * [func (v Value) Set(x Value)](#Value.Set)
  * [func (v Value) SetBool(x bool)](#Value.SetBool)
  * [func (v Value) SetBytes(x []byte)](#Value.SetBytes)
  * [func (v Value) SetCap(n int)](#Value.SetCap)
  * [func (v Value) SetComplex(x complex128)](#Value.SetComplex)
  * [func (v Value) SetFloat(x float64)](#Value.SetFloat)
  * [func (v Value) SetInt(x int64)](#Value.SetInt)
  * [func (v Value) SetLen(n int)](#Value.SetLen)
  * [func (v Value) SetMapIndex(key, elem Value)](#Value.SetMapIndex)
  * [func (v Value) SetPointer(x unsafe.Pointer)](#Value.SetPointer)
  * [func (v Value) SetString(x string)](#Value.SetString)
  * [func (v Value) SetUint(x uint64)](#Value.SetUint)
  * [func (v Value) Slice(i, j int) Value](#Value.Slice)
  * [func (v Value) Slice3(i, j, k int) Value](#Value.Slice3)
  * [func (v Value) String() string](#Value.String)
  * [func (v Value) TryRecv() (x Value, ok bool)](#Value.TryRecv)
  * [func (v Value) TrySend(x Value) bool](#Value.TrySend)
  * [func (v Value) Type() Type](#Value.Type)
  * [func (v Value) Uint() uint64](#Value.Uint)
  * [func (v Value) UnsafeAddr() uintptr](#Value.UnsafeAddr)
* [type ValueError](#ValueError)
  * [func (e *ValueError) Error() string](#ValueError.Error)


#### <a id="pkg-examples">Examples</a>
* [Kind](#example_Kind)
* [MakeFunc](#example_MakeFunc)
* [StructOf](#example_StructOf)
* [StructTag](#example_StructTag)
* [StructTag.Lookup](#example_StructTag_Lookup)
* [TypeOf](#example_TypeOf)


#### <a id="pkg-files">Package files</a>
[deepequal.go](https://golang.org/src/reflect/deepequal.go) [makefunc.go](https://golang.org/src/reflect/makefunc.go) [swapper.go](https://golang.org/src/reflect/swapper.go) [type.go](https://golang.org/src/reflect/type.go) [value.go](https://golang.org/src/reflect/value.go) 






## <a id="Copy">func</a> [Copy](https://golang.org/src/reflect/value.go?s=61204:61233#L2041)
<pre>func Copy(dst, src <a href="#Value">Value</a>) <a href="/pkg/builtin/#int">int</a></pre>
Copy 会将src的内容复制到dst中直至dst被填满或src耗尽. 它会返回被复制元素的数量. dst和src必须是slice或map,且必须有相同的元素类型.

特殊情况: 如果dst的元素类型是unit8, 那么src可以是string.

## <a id="DeepEqual">func</a> [DeepEqual](https://golang.org/src/reflect/deepequal.go?s=6112:6149#L177)
<pre>func DeepEqual(x, y interface{}) <a href="/pkg/builtin/#bool">bool</a></pre>
DeepEqual reports whether x and y are ``deeply equal,'' defined as follows.
Two values of identical type are deeply equal if one of the following cases applies.
Values of distinct types are never deeply equal.

Array values are deeply equal when their corresponding elements are deeply equal.

Struct values are deeply equal if their corresponding fields,
both exported and unexported, are deeply equal.

Func values are deeply equal if both are nil; otherwise they are not deeply equal.

Interface values are deeply equal if they hold deeply equal concrete values.

Map values are deeply equal when all of the following are true:
they are both nil or both non-nil, they have the same length,
and either they are the same map object or their corresponding keys
(matched using Go equality) map to deeply equal values.

Pointer values are deeply equal if they are equal using Go's == operator
or if they point to deeply equal values.

Slice values are deeply equal when all of the following are true:
they are both nil or both non-nil, they have the same length,
and either they point to the same initial entry of the same underlying array
(that is, &x[0] == &y[0]) or their corresponding elements (up to length) are deeply equal.
Note that a non-nil empty slice and a nil slice (for example, []byte{} and []byte(nil))
are not deeply equal.

Other values - numbers, bools, strings, and channels - are deeply equal
if they are equal using Go's == operator.

In general DeepEqual is a recursive relaxation of Go's == operator.
However, this idea is impossible to implement without some inconsistency.
Specifically, it is possible for a value to be unequal to itself,
either because it is of func type (uncomparable in general)
or because it is a floating-point NaN value (not equal to itself in floating-point comparison),
or because it is an array, struct, or interface containing
such a value.
On the other hand, pointer values are always equal to themselves,
even if they point at or contain such problematic values,
because they compare equal using Go's == operator, and that
is a sufficient condition to be deeply equal, regardless of content.
DeepEqual has been defined so that the same short-cut applies
to slices and maps: if x and y are the same slice or the same map,
they are deeply equal regardless of content.

As DeepEqual traverses the data values it may find a cycle. The
second and subsequent times that DeepEqual compares two pointer
values that have been compared before, it treats the values as
equal rather than examining the values to which they point.
This ensures that DeepEqual terminates.



## <a id="Swapper">func</a> [Swapper](https://golang.org/src/reflect/swapper.go?s=337:383#L3)
<pre>func Swapper(slice interface{}) func(i, j <a href="/pkg/builtin/#int">int</a>)</pre>
Swapper returns a function that swaps the elements in the provided
slice.

Swapper panics if the provided interface is not a slice.





## <a id="ChanDir">type</a> [ChanDir](https://golang.org/src/reflect/type.go?s=11596:11612#L334)
ChanDir represents a channel type's direction.


<pre>type ChanDir <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="RecvDir">RecvDir</span> <a href="#ChanDir">ChanDir</a>             = 1 &lt;&lt; <a href="/pkg/builtin/#iota">iota</a> <span class="comment">// &lt;-chan</span>
    <span id="SendDir">SendDir</span>                                 <span class="comment">// chan&lt;-</span>
    <span id="BothDir">BothDir</span> = <a href="#RecvDir">RecvDir</a> | <a href="#SendDir">SendDir</a>             <span class="comment">// chan</span>
)</pre>









### <a id="ChanDir.String">func</a> (ChanDir) [String](https://golang.org/src/reflect/type.go?s=28775:28807#L1035)
<pre>func (d <a href="#ChanDir">ChanDir</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Kind">type</a> [Kind](https://golang.org/src/reflect/type.go?s=8395:8409#L220)
A Kind represents the specific kind of type that a Type represents.
The zero Kind is not a valid kind.


<pre>type Kind <a href="/pkg/builtin/#uint">uint</a></pre>



<pre>const (
    <span id="Invalid">Invalid</span> <a href="#Kind">Kind</a> = <a href="/pkg/builtin/#iota">iota</a>
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
    <span id="Array">Array</span>
    <span id="Chan">Chan</span>
    <span id="Func">Func</span>
    <span id="Interface">Interface</span>
    <span id="Map">Map</span>
    <span id="Ptr">Ptr</span>
    <span id="Slice">Slice</span>
    <span id="String">String</span>
    <span id="Struct">Struct</span>
    <span id="UnsafePointer">UnsafePointer</span>
)</pre>



<a id="example_Kind">Example</a>






### <a id="Kind.String">func</a> (Kind) [String](https://golang.org/src/reflect/type.go?s=17788:17817#L586)
<pre>func (k <a href="#Kind">Kind</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns the name of k.




## <a id="MapIter">type</a> [MapIter](https://golang.org/src/reflect/value.go?s=37394:37446#L1211)
A MapIter is an iterator for ranging over a map.
See Value.MapRange.


<pre>type MapIter struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="MapIter.Key">func</a> (\*MapIter) [Key](https://golang.org/src/reflect/value.go?s=37508:37538#L1217)
<pre>func (it *<a href="#MapIter">MapIter</a>) Key() <a href="#Value">Value</a></pre>
Key returns the key of the iterator's current map entry.




### <a id="MapIter.Next">func</a> (\*MapIter) [Next](https://golang.org/src/reflect/value.go?s=38411:38441#L1247)
<pre>func (it *<a href="#MapIter">MapIter</a>) Next() <a href="/pkg/builtin/#bool">bool</a></pre>
Next advances the map iterator and reports whether there is another
entry. It returns false when the iterator is exhausted; subsequent
calls to Key, Value, or Next will panic.




### <a id="MapIter.Value">func</a> (\*MapIter) [Value](https://golang.org/src/reflect/value.go?s=37895:37927#L1231)
<pre>func (it *<a href="#MapIter">MapIter</a>) Value() <a href="#Value">Value</a></pre>
Value returns the value of the iterator's current map entry.




## <a id="Method">type</a> [Method](https://golang.org/src/reflect/type.go?s=17129:17626#L564)
Method represents a single method.


<pre>type Method struct {
<span id="Method.Name"></span>    <span class="comment">// Name is the method name.</span>
<span id="Method.PkgPath"></span>    <span class="comment">// PkgPath is the package path that qualifies a lower case (unexported)</span>
    <span class="comment">// method name. It is empty for upper case (exported) method names.</span>
    <span class="comment">// The combination of PkgPath and Name uniquely identifies a method</span>
    <span class="comment">// in a method set.</span>
    <span class="comment">// See https://golang.org/ref/spec#Uniqueness_of_identifiers</span>
    Name    <a href="/pkg/builtin/#string">string</a>
    PkgPath <a href="/pkg/builtin/#string">string</a>

<span id="Method.Type"></span>    Type  <a href="#Type">Type</a>  <span class="comment">// method type</span>
<span id="Method.Func"></span>    Func  <a href="#Value">Value</a> <span class="comment">// func with receiver as first argument</span>
<span id="Method.Index"></span>    Index <a href="/pkg/builtin/#int">int</a>   <span class="comment">// index for Type.Method</span>
}
</pre>











## <a id="SelectCase">type</a> [SelectCase](https://golang.org/src/reflect/value.go?s=64057:64220#L2136)
A SelectCase describes a single case in a select operation.
The kind of case depends on Dir, the communication direction.

If Dir is SelectDefault, the case represents a default case.
Chan and Send must be zero Values.

If Dir is SelectSend, the case represents a send operation.
Normally Chan's underlying value must be a channel, and Send's underlying value must be
assignable to the channel's element type. As a special case, if Chan is a zero Value,
then the case is ignored, and the field Send will also be ignored and may be either zero
or non-zero.

If Dir is SelectRecv, the case represents a receive operation.
Normally Chan's underlying value must be a channel and Send must be a zero Value.
If Chan is a zero Value, then the case is ignored, but Send must still be a zero Value.
When a receive operation is selected, the received Value is returned by Select.


<pre>type SelectCase struct {
<span id="SelectCase.Dir"></span>    Dir  <a href="#SelectDir">SelectDir</a> <span class="comment">// direction of case</span>
<span id="SelectCase.Chan"></span>    Chan <a href="#Value">Value</a>     <span class="comment">// channel to use (for send or receive)</span>
<span id="SelectCase.Send"></span>    Send <a href="#Value">Value</a>     <span class="comment">// value to send (for send)</span>
}
</pre>











## <a id="SelectDir">type</a> [SelectDir](https://golang.org/src/reflect/value.go?s=62886:62904#L2108)
A SelectDir describes the communication direction of a select case.


<pre>type SelectDir <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="SelectSend">SelectSend</span>    <a href="#SelectDir">SelectDir</a> <span class="comment">// case Chan &lt;- Send</span>
    <span id="SelectRecv">SelectRecv</span>              <span class="comment">// case &lt;-Chan:</span>
    <span id="SelectDefault">SelectDefault</span>           <span class="comment">// default</span>
)</pre>









## <a id="SliceHeader">type</a> [SliceHeader](https://golang.org/src/reflect/value.go?s=58765:58826#L1954)
SliceHeader is the runtime representation of a slice.
It cannot be used safely or portably and its representation may
change in a later release.
Moreover, the Data field is not sufficient to guarantee the data
it references will not be garbage collected, so programs must keep
a separate, correctly typed pointer to the underlying data.


<pre>type SliceHeader struct {
<span id="SliceHeader.Data"></span>    Data <a href="/pkg/builtin/#uintptr">uintptr</a>
<span id="SliceHeader.Len"></span>    Len  <a href="/pkg/builtin/#int">int</a>
<span id="SliceHeader.Cap"></span>    Cap  <a href="/pkg/builtin/#int">int</a>
}
</pre>











## <a id="StringHeader">type</a> [StringHeader](https://golang.org/src/reflect/value.go?s=58219:58271#L1937)
StringHeader is the runtime representation of a string.
It cannot be used safely or portably and its representation may
change in a later release.
Moreover, the Data field is not sufficient to guarantee the data
it references will not be garbage collected, so programs must keep
a separate, correctly typed pointer to the underlying data.


<pre>type StringHeader struct {
<span id="StringHeader.Data"></span>    Data <a href="/pkg/builtin/#uintptr">uintptr</a>
<span id="StringHeader.Len"></span>    Len  <a href="/pkg/builtin/#int">int</a>
}
</pre>











## <a id="StructField">type</a> [StructField](https://golang.org/src/reflect/type.go?s=29891:30415#L1085)
A StructField describes a single field in a struct.


<pre>type StructField struct {
<span id="StructField.Name"></span>    <span class="comment">// Name is the field name.</span>
    Name <a href="/pkg/builtin/#string">string</a>
<span id="StructField.PkgPath"></span>    <span class="comment">// PkgPath is the package path that qualifies a lower case (unexported)</span>
    <span class="comment">// field name. It is empty for upper case (exported) field names.</span>
    <span class="comment">// See https://golang.org/ref/spec#Uniqueness_of_identifiers</span>
    PkgPath <a href="/pkg/builtin/#string">string</a>

<span id="StructField.Type"></span>    Type      <a href="#Type">Type</a>      <span class="comment">// field type</span>
<span id="StructField.Tag"></span>    Tag       <a href="#StructTag">StructTag</a> <span class="comment">// field tag string</span>
<span id="StructField.Offset"></span>    Offset    <a href="/pkg/builtin/#uintptr">uintptr</a>   <span class="comment">// offset within struct, in bytes</span>
<span id="StructField.Index"></span>    Index     []<a href="/pkg/builtin/#int">int</a>     <span class="comment">// index sequence for Type.FieldByIndex</span>
<span id="StructField.Anonymous"></span>    Anonymous <a href="/pkg/builtin/#bool">bool</a>      <span class="comment">// is an embedded field</span>
}
</pre>











## <a id="StructTag">type</a> [StructTag](https://golang.org/src/reflect/type.go?s=30809:30830#L1108)
A StructTag is the tag string in a struct field.

By convention, tag strings are a concatenation of
optionally space-separated key:"value" pairs.
Each key is a non-empty string consisting of non-control
characters other than space (U+0020 ' '), quote (U+0022 '"'),
and colon (U+003A ':').  Each value is quoted using U+0022 '"'
characters and Go string literal syntax.


<pre>type StructTag <a href="/pkg/builtin/#string">string</a></pre>





<a id="example_StructTag">Example</a>






### <a id="StructTag.Get">func</a> (StructTag) [Get](https://golang.org/src/reflect/type.go?s=31144:31187#L1115)
<pre>func (tag <a href="#StructTag">StructTag</a>) Get(key <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Get returns the value associated with key in the tag string.
If there is no such key in the tag, Get returns the empty string.
If the tag does not have the conventional format, the value
returned by Get is unspecified. To determine whether a tag is
explicitly set to the empty string, use Lookup.




### <a id="StructTag.Lookup">func</a> (StructTag) [Lookup](https://golang.org/src/reflect/type.go?s=31621:31684#L1126)
<pre>func (tag <a href="#StructTag">StructTag</a>) Lookup(key <a href="/pkg/builtin/#string">string</a>) (value <a href="/pkg/builtin/#string">string</a>, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
Lookup returns the value associated with key in the tag string.
If the key is present in the tag the value (which may be empty)
is returned. Otherwise the returned value will be the empty string.
The ok return value reports whether the value was explicitly set in
the tag string. If the tag does not have the conventional format,
the value returned by Lookup is unspecified.


<a id="example_StructTag_Lookup">Example</a>


## <a id="Type">type</a> [Type](https://golang.org/src/reflect/type.go?s=1321:7563#L28)
Type is the representation of a Go type.

Not all methods apply to all kinds of types. Restrictions,
if any, are noted in the documentation for each method.
Use the Kind method to find out the kind of type before
calling kind-specific methods. Calling a method
inappropriate to the kind of type causes a run-time panic.

Type values are comparable, such as with the == operator,
so they can be used as map keys.
Two Type values are equal if they represent identical types.


<pre>type Type interface {

    <span class="comment">// Align returns the alignment in bytes of a value of</span>
    <span class="comment">// this type when allocated in memory.</span>
    Align() <a href="/pkg/builtin/#int">int</a>

    <span class="comment">// FieldAlign returns the alignment in bytes of a value of</span>
    <span class="comment">// this type when used as a field in a struct.</span>
    FieldAlign() <a href="/pkg/builtin/#int">int</a>

    <span class="comment">// Method returns the i&#39;th method in the type&#39;s method set.</span>
    <span class="comment">// It panics if i is not in the range [0, NumMethod()).</span>
    <span class="comment">//</span>
    <span class="comment">// For a non-interface type T or *T, the returned Method&#39;s Type and Func</span>
    <span class="comment">// fields describe a function whose first argument is the receiver.</span>
    <span class="comment">//</span>
    <span class="comment">// For an interface type, the returned Method&#39;s Type field gives the</span>
    <span class="comment">// method signature, without a receiver, and the Func field is nil.</span>
    <span class="comment">//</span>
    <span class="comment">// Only exported methods are accessible and they are sorted in</span>
    <span class="comment">// lexicographic order.</span>
    Method(<a href="/pkg/builtin/#int">int</a>) <a href="#Method">Method</a>

    <span class="comment">// MethodByName returns the method with that name in the type&#39;s</span>
    <span class="comment">// method set and a boolean indicating if the method was found.</span>
    <span class="comment">//</span>
    <span class="comment">// For a non-interface type T or *T, the returned Method&#39;s Type and Func</span>
    <span class="comment">// fields describe a function whose first argument is the receiver.</span>
    <span class="comment">//</span>
    <span class="comment">// For an interface type, the returned Method&#39;s Type field gives the</span>
    <span class="comment">// method signature, without a receiver, and the Func field is nil.</span>
    MethodByName(<a href="/pkg/builtin/#string">string</a>) (<a href="#Method">Method</a>, <a href="/pkg/builtin/#bool">bool</a>)

    <span class="comment">// NumMethod returns the number of exported methods in the type&#39;s method set.</span>
    NumMethod() <a href="/pkg/builtin/#int">int</a>

    <span class="comment">// Name returns the type&#39;s name within its package for a defined type.</span>
    <span class="comment">// For other (non-defined) types it returns the empty string.</span>
    Name() <a href="/pkg/builtin/#string">string</a>

    <span class="comment">// PkgPath returns a defined type&#39;s package path, that is, the import path</span>
    <span class="comment">// that uniquely identifies the package, such as &#34;encoding/base64&#34;.</span>
    <span class="comment">// If the type was predeclared (string, error) or not defined (*T, struct{},</span>
    <span class="comment">// []int, or A where A is an alias for a non-defined type), the package path</span>
    <span class="comment">// will be the empty string.</span>
    PkgPath() <a href="/pkg/builtin/#string">string</a>

    <span class="comment">// Size returns the number of bytes needed to store</span>
    <span class="comment">// a value of the given type; it is analogous to unsafe.Sizeof.</span>
    Size() <a href="/pkg/builtin/#uintptr">uintptr</a>

    <span class="comment">// String returns a string representation of the type.</span>
    <span class="comment">// The string representation may use shortened package names</span>
    <span class="comment">// (e.g., base64 instead of &#34;encoding/base64&#34;) and is not</span>
    <span class="comment">// guaranteed to be unique among types. To test for type identity,</span>
    <span class="comment">// compare the Types directly.</span>
    String() <a href="/pkg/builtin/#string">string</a>

    <span class="comment">// Kind returns the specific kind of this type.</span>
    Kind() <a href="#Kind">Kind</a>

    <span class="comment">// Implements reports whether the type implements the interface type u.</span>
    Implements(u <a href="#Type">Type</a>) <a href="/pkg/builtin/#bool">bool</a>

    <span class="comment">// AssignableTo reports whether a value of the type is assignable to type u.</span>
    AssignableTo(u <a href="#Type">Type</a>) <a href="/pkg/builtin/#bool">bool</a>

    <span class="comment">// ConvertibleTo reports whether a value of the type is convertible to type u.</span>
    ConvertibleTo(u <a href="#Type">Type</a>) <a href="/pkg/builtin/#bool">bool</a>

    <span class="comment">// Comparable reports whether values of this type are comparable.</span>
    Comparable() <a href="/pkg/builtin/#bool">bool</a>

    <span class="comment">// Bits returns the size of the type in bits.</span>
    <span class="comment">// It panics if the type&#39;s Kind is not one of the</span>
    <span class="comment">// sized or unsized Int, Uint, Float, or Complex kinds.</span>
    Bits() <a href="/pkg/builtin/#int">int</a>

    <span class="comment">// ChanDir returns a channel type&#39;s direction.</span>
    <span class="comment">// It panics if the type&#39;s Kind is not Chan.</span>
    ChanDir() <a href="#ChanDir">ChanDir</a>

    <span class="comment">// IsVariadic reports whether a function type&#39;s final input parameter</span>
    <span class="comment">// is a &#34;...&#34; parameter. If so, t.In(t.NumIn() - 1) returns the parameter&#39;s</span>
    <span class="comment">// implicit actual type []T.</span>
    <span class="comment">//</span>
    <span class="comment">// For concreteness, if t represents func(x int, y ... float64), then</span>
    <span class="comment">//</span>
    <span class="comment">//	t.NumIn() == 2</span>
    <span class="comment">//	t.In(0) is the reflect.Type for &#34;int&#34;</span>
    <span class="comment">//	t.In(1) is the reflect.Type for &#34;[]float64&#34;</span>
    <span class="comment">//	t.IsVariadic() == true</span>
    <span class="comment">//</span>
    <span class="comment">// IsVariadic panics if the type&#39;s Kind is not Func.</span>
    IsVariadic() <a href="/pkg/builtin/#bool">bool</a>

    <span class="comment">// Elem returns a type&#39;s element type.</span>
    <span class="comment">// It panics if the type&#39;s Kind is not Array, Chan, Map, Ptr, or Slice.</span>
    Elem() <a href="#Type">Type</a>

    <span class="comment">// Field returns a struct type&#39;s i&#39;th field.</span>
    <span class="comment">// It panics if the type&#39;s Kind is not Struct.</span>
    <span class="comment">// It panics if i is not in the range [0, NumField()).</span>
    Field(i <a href="/pkg/builtin/#int">int</a>) <a href="#StructField">StructField</a>

    <span class="comment">// FieldByIndex returns the nested field corresponding</span>
    <span class="comment">// to the index sequence. It is equivalent to calling Field</span>
    <span class="comment">// successively for each index i.</span>
    <span class="comment">// It panics if the type&#39;s Kind is not Struct.</span>
    FieldByIndex(index []<a href="/pkg/builtin/#int">int</a>) <a href="#StructField">StructField</a>

    <span class="comment">// FieldByName returns the struct field with the given name</span>
    <span class="comment">// and a boolean indicating if the field was found.</span>
    FieldByName(name <a href="/pkg/builtin/#string">string</a>) (<a href="#StructField">StructField</a>, <a href="/pkg/builtin/#bool">bool</a>)

    <span class="comment">// FieldByNameFunc returns the struct field with a name</span>
    <span class="comment">// that satisfies the match function and a boolean indicating if</span>
    <span class="comment">// the field was found.</span>
    <span class="comment">//</span>
    <span class="comment">// FieldByNameFunc considers the fields in the struct itself</span>
    <span class="comment">// and then the fields in any embedded structs, in breadth first order,</span>
    <span class="comment">// stopping at the shallowest nesting depth containing one or more</span>
    <span class="comment">// fields satisfying the match function. If multiple fields at that depth</span>
    <span class="comment">// satisfy the match function, they cancel each other</span>
    <span class="comment">// and FieldByNameFunc returns no match.</span>
    <span class="comment">// This behavior mirrors Go&#39;s handling of name lookup in</span>
    <span class="comment">// structs containing embedded fields.</span>
    FieldByNameFunc(match func(<a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a>) (<a href="#StructField">StructField</a>, <a href="/pkg/builtin/#bool">bool</a>)

    <span class="comment">// In returns the type of a function type&#39;s i&#39;th input parameter.</span>
    <span class="comment">// It panics if the type&#39;s Kind is not Func.</span>
    <span class="comment">// It panics if i is not in the range [0, NumIn()).</span>
    In(i <a href="/pkg/builtin/#int">int</a>) <a href="#Type">Type</a>

    <span class="comment">// Key returns a map type&#39;s key type.</span>
    <span class="comment">// It panics if the type&#39;s Kind is not Map.</span>
    Key() <a href="#Type">Type</a>

    <span class="comment">// Len returns an array type&#39;s length.</span>
    <span class="comment">// It panics if the type&#39;s Kind is not Array.</span>
    Len() <a href="/pkg/builtin/#int">int</a>

    <span class="comment">// NumField returns a struct type&#39;s field count.</span>
    <span class="comment">// It panics if the type&#39;s Kind is not Struct.</span>
    NumField() <a href="/pkg/builtin/#int">int</a>

    <span class="comment">// NumIn returns a function type&#39;s input parameter count.</span>
    <span class="comment">// It panics if the type&#39;s Kind is not Func.</span>
    NumIn() <a href="/pkg/builtin/#int">int</a>

    <span class="comment">// NumOut returns a function type&#39;s output parameter count.</span>
    <span class="comment">// It panics if the type&#39;s Kind is not Func.</span>
    NumOut() <a href="/pkg/builtin/#int">int</a>

    <span class="comment">// Out returns the type of a function type&#39;s i&#39;th output parameter.</span>
    <span class="comment">// It panics if the type&#39;s Kind is not Func.</span>
    <span class="comment">// It panics if i is not in the range [0, NumOut()).</span>
    Out(i <a href="/pkg/builtin/#int">int</a>) <a href="#Type">Type</a>
    <span class="comment">// contains filtered or unexported methods</span>
}</pre>









### <a id="ArrayOf">func</a> [ArrayOf](https://golang.org/src/reflect/type.go?s=78277:78316#L2801)
<pre>func ArrayOf(count <a href="/pkg/builtin/#int">int</a>, elem <a href="#Type">Type</a>) <a href="#Type">Type</a></pre>
ArrayOf returns the array type with the given count and element type.
For example, if t represents int, ArrayOf(5, t) represents [5]int.

If the resulting type would be larger than the available address space,
ArrayOf panics.




### <a id="ChanOf">func</a> [ChanOf](https://golang.org/src/reflect/type.go?s=49752:49789#L1766)
<pre>func ChanOf(dir <a href="#ChanDir">ChanDir</a>, t <a href="#Type">Type</a>) <a href="#Type">Type</a></pre>
ChanOf returns the channel type with the given direction and element type.
For example, if t represents int, ChanOf(RecvDir, t) represents <-chan int.

The gc runtime imposes a limit of 64 kB on channel element types.
If t's size is equal to or exceeds this limit, ChanOf panics.




### <a id="FuncOf">func</a> [FuncOf](https://golang.org/src/reflect/type.go?s=53751:53798#L1921)
<pre>func FuncOf(in, out []<a href="#Type">Type</a>, variadic <a href="/pkg/builtin/#bool">bool</a>) <a href="#Type">Type</a></pre>
FuncOf returns the function type with the given argument and result types.
For example if k represents int and e represents string,
FuncOf([]Type{k}, []Type{e}, false) represents func(int) string.

The variadic argument controls whether the function is variadic. FuncOf
panics if the in[len(in)-1] does not represent a slice and variadic is
true.




### <a id="MapOf">func</a> [MapOf](https://golang.org/src/reflect/type.go?s=51294:51325#L1823)
<pre>func MapOf(key, elem <a href="#Type">Type</a>) <a href="#Type">Type</a></pre>
MapOf returns the map type with the given key and element types.
For example, if k represents int and e represents string,
MapOf(k, e) represents map[int]string.

If the key type is not a valid map key type (that is, if it does
not implement Go's == operator), MapOf panics.




### <a id="PtrTo">func</a> [PtrTo](https://golang.org/src/reflect/type.go?s=39095:39118#L1373)
<pre>func PtrTo(t <a href="#Type">Type</a>) <a href="#Type">Type</a></pre>
PtrTo returns the pointer type with element t.
For example, if t represents type Foo, PtrTo(t) represents *Foo.




### <a id="SliceOf">func</a> [SliceOf](https://golang.org/src/reflect/type.go?s=62434:62459#L2243)
<pre>func SliceOf(t <a href="#Type">Type</a>) <a href="#Type">Type</a></pre>
SliceOf returns the slice type with element type t.
For example, if t represents int, SliceOf(t) represents []int.




### <a id="StructOf">func</a> [StructOf](https://golang.org/src/reflect/type.go?s=64711:64751#L2324)
<pre>func StructOf(fields []<a href="#StructField">StructField</a>) <a href="#Type">Type</a></pre>
StructOf returns the struct type containing fields.
The Offset and Index fields are ignored and computed as they would be
by the compiler.

StructOf currently does not generate wrapper methods for embedded
fields and panics if passed unexported StructFields.
These limitations may be lifted in a future version.


<a id="example_StructOf">Example</a>


### <a id="TypeOf">func</a> [TypeOf](https://golang.org/src/reflect/type.go?s=38787:38818#L1363)
<pre>func TypeOf(i interface{}) <a href="#Type">Type</a></pre>
TypeOf returns the reflection Type that represents the dynamic type of i.
If i is a nil interface value, TypeOf returns nil.


<a id="example_TypeOf">Example</a>




## <a id="Value">type</a> [Value](https://golang.org/src/reflect/value.go?s=1303:2522#L26)
Value is the reflection interface to a Go value.

Not all methods apply to all kinds of values. Restrictions,
if any, are noted in the documentation for each method.
Use the Kind method to find out the kind of value before
calling kind-specific methods. Calling a method
inappropriate to the kind of type causes a run time panic.

The zero Value represents no value.
Its IsValid method returns false, its Kind method returns Invalid,
its String method returns "<invalid Value>", and all other methods panic.
Most functions and methods never return an invalid value.
If one does, its documentation states the conditions explicitly.

A Value can be used concurrently by multiple goroutines provided that
the underlying Go value can be used concurrently for the equivalent
direct operations.

To compare two Values, compare the results of the Interface method.
Using == on two Values does not compare the underlying values
they represent.


<pre>type Value struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="Append">func</a> [Append](https://golang.org/src/reflect/value.go?s=60336:60374#L2014)
<pre>func Append(s <a href="#Value">Value</a>, x ...<a href="#Value">Value</a>) <a href="#Value">Value</a></pre>
Append appends the values x to a slice s and returns the resulting slice.
As in Go, each x's value must be assignable to the slice's element type.




### <a id="AppendSlice">func</a> [AppendSlice](https://golang.org/src/reflect/value.go?s=60643:60677#L2025)
<pre>func AppendSlice(s, t <a href="#Value">Value</a>) <a href="#Value">Value</a></pre>
AppendSlice appends a slice t to a slice s and returns the resulting slice.
The slices s and t must have the same element type.




### <a id="Indirect">func</a> [Indirect](https://golang.org/src/reflect/value.go?s=68611:68639#L2297)
<pre>func Indirect(v <a href="#Value">Value</a>) <a href="#Value">Value</a></pre>
Indirect returns the value that v points to.
If v is a nil pointer, Indirect returns a zero Value.
If v is not a pointer, Indirect returns v.




### <a id="MakeChan">func</a> [MakeChan](https://golang.org/src/reflect/value.go?s=67671:67712#L2263)
<pre>func MakeChan(typ <a href="#Type">Type</a>, buffer <a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
MakeChan creates a new channel with the specified type and buffer size.




### <a id="MakeFunc">func</a> [MakeFunc](https://golang.org/src/reflect/makefunc.go?s=1662:1732#L38)
<pre>func MakeFunc(typ <a href="#Type">Type</a>, fn func(args []<a href="#Value">Value</a>) (results []<a href="#Value">Value</a>)) <a href="#Value">Value</a></pre>
MakeFunc returns a new function of the given Type
that wraps the function fn. When called, that new function
does the following:


	- converts its arguments to a slice of Values.
	- runs results := fn(args).
	- returns the results as a slice of Values, one per formal result.

The implementation fn can assume that the argument Value slice
has the number and type of arguments given by typ.
If typ describes a variadic function, the final Value is itself
a slice representing the variadic arguments, as in the
body of a variadic function. The result Value slice returned by fn
must have the number and type of results given by typ.

The Value.Call method allows the caller to invoke a typed function
in terms of Values; in contrast, MakeFunc allows the caller to implement
a typed function in terms of Values.

The Examples section of the documentation includes an illustration
of how to use MakeFunc to build a swap function for different types.


<a id="example_MakeFunc">Example</a>


### <a id="MakeMap">func</a> [MakeMap](https://golang.org/src/reflect/value.go?s=68085:68113#L2279)
<pre>func MakeMap(typ <a href="#Type">Type</a>) <a href="#Value">Value</a></pre>
MakeMap creates a new map with the specified type.




### <a id="MakeMapWithSize">func</a> [MakeMapWithSize](https://golang.org/src/reflect/value.go?s=68263:68306#L2285)
<pre>func MakeMapWithSize(typ <a href="#Type">Type</a>, n <a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
MakeMapWithSize creates a new map with the specified type
and initial space for approximately n elements.




### <a id="MakeSlice">func</a> [MakeSlice](https://golang.org/src/reflect/value.go?s=67146:67190#L2244)
<pre>func MakeSlice(typ <a href="#Type">Type</a>, len, cap <a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
MakeSlice creates a new zero-initialized slice value
for the specified slice type, length, and capacity.




### <a id="New">func</a> [New](https://golang.org/src/reflect/value.go?s=69829:69853#L2339)
<pre>func New(typ <a href="#Type">Type</a>) <a href="#Value">Value</a></pre>
New returns a Value representing a pointer to a new zero value
for the specified type. That is, the returned Value's Type is PtrTo(typ).




### <a id="NewAt">func</a> [NewAt](https://golang.org/src/reflect/value.go?s=70110:70154#L2351)
<pre>func NewAt(typ <a href="#Type">Type</a>, p <a href="/pkg/unsafe/">unsafe</a>.<a href="/pkg/unsafe/#Pointer">Pointer</a>) <a href="#Value">Value</a></pre>
NewAt returns a Value representing a pointer to a value of the
specified type, using p as that pointer.




### <a id="Select">func</a> [Select](https://golang.org/src/reflect/value.go?s=64713:64782#L2149)
<pre>func Select(cases []<a href="#SelectCase">SelectCase</a>) (chosen <a href="/pkg/builtin/#int">int</a>, recv <a href="#Value">Value</a>, recvOK <a href="/pkg/builtin/#bool">bool</a>)</pre>
Select executes a select operation described by the list of cases.
Like the Go select statement, it blocks until at least one of the cases
can proceed, makes a uniform pseudo-random choice,
and then executes that case. It returns the index of the chosen case
and, if that case was a receive operation, the value received and a
boolean indicating whether the value corresponds to a send on the channel
(as opposed to a zero value received because the channel is closed).




### <a id="ValueOf">func</a> [ValueOf](https://golang.org/src/reflect/value.go?s=68830:68863#L2306)
<pre>func ValueOf(i interface{}) <a href="#Value">Value</a></pre>
ValueOf returns a new Value initialized to the concrete value
stored in the interface i. ValueOf(nil) returns the zero Value.




### <a id="Zero">func</a> [Zero](https://golang.org/src/reflect/value.go?s=69466:69491#L2325)
<pre>func Zero(typ <a href="#Type">Type</a>) <a href="#Value">Value</a></pre>
Zero returns a Value representing the zero value for the specified type.
The result is different from the zero value of the Value struct,
which represents no value at all.
For example, Zero(TypeOf(42)) returns a Value with Kind Int and value 0.
The returned value is neither addressable nor settable.






### <a id="Value.Addr">func</a> (Value) [Addr](https://golang.org/src/reflect/value.go?s=7759:7786#L246)
<pre>func (v <a href="#Value">Value</a>) Addr() <a href="#Value">Value</a></pre>
Addr returns a pointer value representing the address of v.
It panics if CanAddr() returns false.
Addr is typically used to obtain a pointer to a struct field
or slice element in order to call a method that requires a
pointer receiver.




### <a id="Value.Bool">func</a> (Value) [Bool](https://golang.org/src/reflect/value.go?s=8012:8038#L255)
<pre>func (v <a href="#Value">Value</a>) Bool() <a href="/pkg/builtin/#bool">bool</a></pre>
Bool returns v's underlying value.
It panics if v's kind is not Bool.




### <a id="Value.Bytes">func</a> (Value) [Bytes](https://golang.org/src/reflect/value.go?s=8185:8214#L262)
<pre>func (v <a href="#Value">Value</a>) Bytes() []<a href="/pkg/builtin/#byte">byte</a></pre>
Bytes returns v's underlying value.
It panics if v's underlying value is not a slice of bytes.




### <a id="Value.Call">func</a> (Value) [Call](https://golang.org/src/reflect/value.go?s=9975:10014#L308)
<pre>func (v <a href="#Value">Value</a>) Call(in []<a href="#Value">Value</a>) []<a href="#Value">Value</a></pre>
Call calls the function v with the input arguments in.
For example, if len(in) == 3, v.Call(in) represents the Go call v(in[0], in[1], in[2]).
Call panics if v's Kind is not Func.
It returns the output results as Values.
As in Go, each input argument must be assignable to the
type of the function's corresponding input parameter.
If v is a variadic function, Call creates the variadic slice parameter
itself, copying in the corresponding values.




### <a id="Value.CallSlice">func</a> (Value) [CallSlice](https://golang.org/src/reflect/value.go?s=10552:10596#L321)
<pre>func (v <a href="#Value">Value</a>) CallSlice(in []<a href="#Value">Value</a>) []<a href="#Value">Value</a></pre>
CallSlice calls the variadic function v with the input arguments in,
assigning the slice in[len(in)-1] to v's final variadic argument.
For example, if len(in) == 3, v.CallSlice(in) represents the Go call v(in[0], in[1], in[2]...).
CallSlice panics if v's Kind is not Func or if v is not variadic.
It returns the output results as Values.
As in Go, each input argument must be assignable to the
type of the function's corresponding input parameter.




### <a id="Value.CanAddr">func</a> (Value) [CanAddr](https://golang.org/src/reflect/value.go?s=9081:9110#L287)
<pre>func (v <a href="#Value">Value</a>) CanAddr() <a href="/pkg/builtin/#bool">bool</a></pre>
CanAddr reports whether the value's address can be obtained with Addr.
Such values are called addressable. A value is addressable if it is
an element of a slice, an element of an addressable array,
a field of an addressable struct, or the result of dereferencing a pointer.
If CanAddr returns false, calling Addr will panic.




### <a id="Value.CanInterface">func</a> (Value) [CanInterface](https://golang.org/src/reflect/value.go?s=30628:30662#L980)
<pre>func (v <a href="#Value">Value</a>) CanInterface() <a href="/pkg/builtin/#bool">bool</a></pre>
CanInterface reports whether Interface can be used without panicking.




### <a id="Value.CanSet">func</a> (Value) [CanSet](https://golang.org/src/reflect/value.go?s=9425:9453#L296)
<pre>func (v <a href="#Value">Value</a>) CanSet() <a href="/pkg/builtin/#bool">bool</a></pre>
CanSet reports whether the value of v can be changed.
A Value can be changed only if it is addressable and was not
obtained by the use of unexported struct fields.
If CanSet returns false, calling Set or any type-specific
setter (e.g., SetBool, SetInt) will panic.




### <a id="Value.Cap">func</a> (Value) [Cap](https://golang.org/src/reflect/value.go?s=24283:24307#L749)
<pre>func (v <a href="#Value">Value</a>) Cap() <a href="/pkg/builtin/#int">int</a></pre>
Cap returns v's capacity.
It panics if v's Kind is not Array, Chan, or Slice.




### <a id="Value.Close">func</a> (Value) [Close](https://golang.org/src/reflect/value.go?s=24646:24668#L765)
<pre>func (v <a href="#Value">Value</a>) Close()</pre>
Close closes the channel v.
It panics if v's Kind is not Chan.




### <a id="Value.Complex">func</a> (Value) [Complex](https://golang.org/src/reflect/value.go?s=24848:24883#L773)
<pre>func (v <a href="#Value">Value</a>) Complex() <a href="/pkg/builtin/#complex128">complex128</a></pre>
Complex returns v's underlying value, as a complex128.
It panics if v's Kind is not Complex64 or Complex128




### <a id="Value.Convert">func</a> (Value) [Convert](https://golang.org/src/reflect/value.go?s=71619:71655#L2399)
<pre>func (v <a href="#Value">Value</a>) Convert(t <a href="#Type">Type</a>) <a href="#Value">Value</a></pre>
Convert returns the value v converted to type t.
If the usual Go conversion rules do not allow conversion
of the value v to type t, Convert panics.




### <a id="Value.Elem">func</a> (Value) [Elem](https://golang.org/src/reflect/value.go?s=25266:25293#L788)
<pre>func (v <a href="#Value">Value</a>) Elem() <a href="#Value">Value</a></pre>
Elem returns the value that the interface v contains
or that the pointer v points to.
It panics if v's Kind is not Interface or Ptr.
It returns the zero Value if v is nil.




### <a id="Value.Field">func</a> (Value) [Field](https://golang.org/src/reflect/value.go?s=26096:26129#L825)
<pre>func (v <a href="#Value">Value</a>) Field(i <a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
Field returns the i'th field of the struct v.
It panics if v's Kind is not Struct or i is out of range.




### <a id="Value.FieldByIndex">func</a> (Value) [FieldByIndex](https://golang.org/src/reflect/value.go?s=27159:27205#L857)
<pre>func (v <a href="#Value">Value</a>) FieldByIndex(index []<a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
FieldByIndex returns the nested field corresponding to index.
It panics if v's Kind is not struct.




### <a id="Value.FieldByName">func</a> (Value) [FieldByName](https://golang.org/src/reflect/value.go?s=27686:27731#L879)
<pre>func (v <a href="#Value">Value</a>) FieldByName(name <a href="/pkg/builtin/#string">string</a>) <a href="#Value">Value</a></pre>
FieldByName returns the struct field with the given name.
It returns the zero Value if no field was found.
It panics if v's Kind is not struct.




### <a id="Value.FieldByNameFunc">func</a> (Value) [FieldByNameFunc](https://golang.org/src/reflect/value.go?s=28036:28097#L891)
<pre>func (v <a href="#Value">Value</a>) FieldByNameFunc(match func(<a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a>) <a href="#Value">Value</a></pre>
FieldByNameFunc returns the struct field with a name
that satisfies the match function.
It panics if v's Kind is not struct.
It returns the zero Value if no field was found.




### <a id="Value.Float">func</a> (Value) [Float](https://golang.org/src/reflect/value.go?s=28307:28337#L900)
<pre>func (v <a href="#Value">Value</a>) Float() <a href="/pkg/builtin/#float64">float64</a></pre>
Float returns v's underlying value, as a float64.
It panics if v's Kind is not Float32 or Float64




### <a id="Value.Index">func</a> (Value) [Index](https://golang.org/src/reflect/value.go?s=28677:28710#L915)
<pre>func (v <a href="#Value">Value</a>) Index(i <a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
Index returns v's i'th element.
It panics if v's Kind is not Array, Slice, or String or i is out of range.




### <a id="Value.Int">func</a> (Value) [Int](https://golang.org/src/reflect/value.go?s=30233:30259#L961)
<pre>func (v <a href="#Value">Value</a>) Int() <a href="/pkg/builtin/#int64">int64</a></pre>
Int returns v's underlying value, as an int64.
It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64.




### <a id="Value.Interface">func</a> (Value) [Interface](https://golang.org/src/reflect/value.go?s=30985:31027#L992)
<pre>func (v <a href="#Value">Value</a>) Interface() (i interface{})</pre>
Interface returns v's current value as an interface{}.
It is equivalent to:


	var i interface{} = (v's underlying value)

It panics if the Value was obtained by accessing
unexported struct fields.




### <a id="Value.InterfaceData">func</a> (Value) [InterfaceData](https://golang.org/src/reflect/value.go?s=32081:32122#L1028)
<pre>func (v <a href="#Value">Value</a>) InterfaceData() [2]<a href="/pkg/builtin/#uintptr">uintptr</a></pre>
InterfaceData returns the interface v's value as a uintptr pair.
It panics if v's Kind is not Interface.




### <a id="Value.IsNil">func</a> (Value) [IsNil](https://golang.org/src/reflect/value.go?s=32868:32895#L1046)
<pre>func (v <a href="#Value">Value</a>) IsNil() <a href="/pkg/builtin/#bool">bool</a></pre>
IsNil reports whether its argument v is nil. The argument must be
a chan, func, interface, map, pointer, or slice value; if it is
not, IsNil panics. Note that IsNil is not always equivalent to a
regular comparison with nil in Go. For example, if v was created
by calling ValueOf with an uninitialized interface variable i,
i==nil will be true but v.IsNil will panic as v will be the zero
Value.




### <a id="Value.IsValid">func</a> (Value) [IsValid](https://golang.org/src/reflect/value.go?s=33651:33680#L1071)
<pre>func (v <a href="#Value">Value</a>) IsValid() <a href="/pkg/builtin/#bool">bool</a></pre>
IsValid reports whether v represents a value.
It returns false if v is the zero Value.
If IsValid returns false, all other methods except String panic.
Most functions and methods never return an invalid value.
If one does, its documentation states the conditions explicitly.




### <a id="Value.IsZero">func</a> (Value) [IsZero](https://golang.org/src/reflect/value.go?s=33807:33835#L1077)
<pre>func (v <a href="#Value">Value</a>) IsZero() <a href="/pkg/builtin/#bool">bool</a></pre>
IsZero reports whether v is the zero value for its type.
It panics if the argument is invalid.




### <a id="Value.Kind">func</a> (Value) [Kind](https://golang.org/src/reflect/value.go?s=34854:34880#L1117)
<pre>func (v <a href="#Value">Value</a>) Kind() <a href="#Kind">Kind</a></pre>
Kind returns v's Kind.
If v is the zero Value (IsValid returns false), Kind returns Invalid.




### <a id="Value.Len">func</a> (Value) [Len](https://golang.org/src/reflect/value.go?s=34998:35022#L1123)
<pre>func (v <a href="#Value">Value</a>) Len() <a href="/pkg/builtin/#int">int</a></pre>
Len returns v's length.
It panics if v's Kind is not Array, Chan, Map, Slice, or String.




### <a id="Value.MapIndex">func</a> (Value) [MapIndex](https://golang.org/src/reflect/value.go?s=35734:35774#L1147)
<pre>func (v <a href="#Value">Value</a>) MapIndex(key <a href="#Value">Value</a>) <a href="#Value">Value</a></pre>
MapIndex returns the value associated with key in the map v.
It panics if v's Kind is not Map.
It returns the zero Value if key is not found in the map or if v represents a nil map.
As in Go, the key's value must be assignable to the map's key type.




### <a id="Value.MapKeys">func</a> (Value) [MapKeys](https://golang.org/src/reflect/value.go?s=36734:36766#L1180)
<pre>func (v <a href="#Value">Value</a>) MapKeys() []<a href="#Value">Value</a></pre>
MapKeys returns a slice containing all the keys present in the map,
in unspecified order.
It panics if v's Kind is not Map.
It returns an empty slice if v represents a nil map.




### <a id="Value.MapRange">func</a> (Value) [MapRange](https://golang.org/src/reflect/value.go?s=39097:39131#L1275)
<pre>func (v <a href="#Value">Value</a>) MapRange() *<a href="#MapIter">MapIter</a></pre>
MapRange returns a range iterator for a map.
It panics if v's Kind is not Map.

Call Next to advance the iterator, and Key/Value to access each entry.
Next returns false when the iterator is exhausted.
MapRange follows the same iteration semantics as a range statement.

Example:


	iter := reflect.ValueOf(m).MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		...
	}




### <a id="Value.Method">func</a> (Value) [Method](https://golang.org/src/reflect/value.go?s=39880:39914#L1297)
<pre>func (v <a href="#Value">Value</a>) Method(i <a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
Method returns a function value corresponding to v's i'th method.
The arguments to a Call on the returned function should not include
a receiver; the returned function will always use v as the receiver.
Method panics if i is out of range or if v is a nil interface value.




### <a id="Value.MethodByName">func</a> (Value) [MethodByName](https://golang.org/src/reflect/value.go?s=40926:40972#L1329)
<pre>func (v <a href="#Value">Value</a>) MethodByName(name <a href="/pkg/builtin/#string">string</a>) <a href="#Value">Value</a></pre>
MethodByName returns a function value corresponding to the method
of v with the given name.
The arguments to a Call on the returned function should not include
a receiver; the returned function will always use v as the receiver.
It returns the zero Value if no method was found.




### <a id="Value.NumField">func</a> (Value) [NumField](https://golang.org/src/reflect/value.go?s=41298:41327#L1345)
<pre>func (v <a href="#Value">Value</a>) NumField() <a href="/pkg/builtin/#int">int</a></pre>
NumField returns the number of fields in the struct v.
It panics if v's Kind is not Struct.




### <a id="Value.NumMethod">func</a> (Value) [NumMethod](https://golang.org/src/reflect/value.go?s=40448:40478#L1314)
<pre>func (v <a href="#Value">Value</a>) NumMethod() <a href="/pkg/builtin/#int">int</a></pre>
NumMethod returns the number of exported methods in the value's method set.




### <a id="Value.OverflowComplex">func</a> (Value) [OverflowComplex](https://golang.org/src/reflect/value.go?s=41562:41611#L1353)
<pre>func (v <a href="#Value">Value</a>) OverflowComplex(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
OverflowComplex reports whether the complex128 x cannot be represented by v's type.
It panics if v's Kind is not Complex64 or Complex128.




### <a id="Value.OverflowFloat">func</a> (Value) [OverflowFloat](https://golang.org/src/reflect/value.go?s=41956:42000#L1366)
<pre>func (v <a href="#Value">Value</a>) OverflowFloat(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
OverflowFloat reports whether the float64 x cannot be represented by v's type.
It panics if v's Kind is not Float32 or Float64.




### <a id="Value.OverflowInt">func</a> (Value) [OverflowInt](https://golang.org/src/reflect/value.go?s=42433:42473#L1386)
<pre>func (v <a href="#Value">Value</a>) OverflowInt(x <a href="/pkg/builtin/#int64">int64</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
OverflowInt reports whether the int64 x cannot be represented by v's type.
It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64.




### <a id="Value.OverflowUint">func</a> (Value) [OverflowUint](https://golang.org/src/reflect/value.go?s=42866:42908#L1399)
<pre>func (v <a href="#Value">Value</a>) OverflowUint(x <a href="/pkg/builtin/#uint64">uint64</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
OverflowUint reports whether the uint64 x cannot be represented by v's type.
It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64.




### <a id="Value.Pointer">func</a> (Value) [Pointer](https://golang.org/src/reflect/value.go?s=43879:43911#L1424)
<pre>func (v <a href="#Value">Value</a>) Pointer() <a href="/pkg/builtin/#uintptr">uintptr</a></pre>
Pointer returns v's value as a uintptr.
It returns uintptr instead of unsafe.Pointer so that
code using reflect cannot obtain unsafe.Pointers
without importing the unsafe package explicitly.
It panics if v's Kind is not Chan, Func, Map, Ptr, Slice, or UnsafePointer.

If v's Kind is Func, the returned pointer is an underlying
code pointer, but not necessarily enough to identify a
single function uniquely. The only guarantee is that the
result is zero if and only if v is a nil func Value.

If v's Kind is Slice, the returned pointer is to the first
element of the slice. If the slice is nil the returned value
is 0.  If the slice is empty but non-nil the return value is non-zero.




### <a id="Value.Recv">func</a> (Value) [Recv](https://golang.org/src/reflect/value.go?s=45087:45127#L1460)
<pre>func (v <a href="#Value">Value</a>) Recv() (x <a href="#Value">Value</a>, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
Recv receives and returns a value from the channel v.
It panics if v's Kind is not Chan.
The receive blocks until a value is ready.
The boolean value ok is true if the value x corresponds to a send
on the channel, false if it is a zero value received because the channel is closed.




### <a id="Value.Send">func</a> (Value) [Send](https://golang.org/src/reflect/value.go?s=45934:45962#L1493)
<pre>func (v <a href="#Value">Value</a>) Send(x <a href="#Value">Value</a>)</pre>
Send sends x on the channel v.
It panics if v's kind is not Chan or if x's type is not the same type as v's element type.
As in Go, x's value must be assignable to the channel's element type.




### <a id="Value.Set">func</a> (Value) [Set](https://golang.org/src/reflect/value.go?s=46618:46645#L1520)
<pre>func (v <a href="#Value">Value</a>) Set(x <a href="#Value">Value</a>)</pre>
Set assigns x to the value v.
It panics if CanSet returns false.
As in Go, x's value must be assignable to v's type.




### <a id="Value.SetBool">func</a> (Value) [SetBool](https://golang.org/src/reflect/value.go?s=47059:47089#L1537)
<pre>func (v <a href="#Value">Value</a>) SetBool(x <a href="/pkg/builtin/#bool">bool</a>)</pre>
SetBool sets v's underlying value.
It panics if v's Kind is not Bool or if CanSet() is false.




### <a id="Value.SetBytes">func</a> (Value) [SetBytes](https://golang.org/src/reflect/value.go?s=47255:47288#L1545)
<pre>func (v <a href="#Value">Value</a>) SetBytes(x []<a href="/pkg/builtin/#byte">byte</a>)</pre>
SetBytes sets v's underlying value.
It panics if v's underlying value is not a slice of bytes.




### <a id="Value.SetCap">func</a> (Value) [SetCap](https://golang.org/src/reflect/value.go?s=49485:49513#L1629)
<pre>func (v <a href="#Value">Value</a>) SetCap(n <a href="/pkg/builtin/#int">int</a>)</pre>
SetCap sets v's capacity to n.
It panics if v's Kind is not Slice or if n is smaller than the length or
greater than the capacity of the slice.




### <a id="Value.SetComplex">func</a> (Value) [SetComplex](https://golang.org/src/reflect/value.go?s=47875:47914#L1567)
<pre>func (v <a href="#Value">Value</a>) SetComplex(x <a href="/pkg/builtin/#complex128">complex128</a>)</pre>
SetComplex sets v's underlying value to x.
It panics if v's Kind is not Complex64 or Complex128, or if CanSet() is false.




### <a id="Value.SetFloat">func</a> (Value) [SetFloat](https://golang.org/src/reflect/value.go?s=48263:48297#L1581)
<pre>func (v <a href="#Value">Value</a>) SetFloat(x <a href="/pkg/builtin/#float64">float64</a>)</pre>
SetFloat sets v's underlying value to x.
It panics if v's Kind is not Float32 or Float64, or if CanSet() is false.




### <a id="Value.SetInt">func</a> (Value) [SetInt](https://golang.org/src/reflect/value.go?s=48645:48675#L1595)
<pre>func (v <a href="#Value">Value</a>) SetInt(x <a href="/pkg/builtin/#int64">int64</a>)</pre>
SetInt sets v's underlying value to x.
It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64, or if CanSet() is false.




### <a id="Value.SetLen">func</a> (Value) [SetLen](https://golang.org/src/reflect/value.go?s=49133:49161#L1616)
<pre>func (v <a href="#Value">Value</a>) SetLen(n <a href="/pkg/builtin/#int">int</a>)</pre>
SetLen sets v's length to n.
It panics if v's Kind is not Slice or if n is negative or
greater than the capacity of the slice.




### <a id="Value.SetMapIndex">func</a> (Value) [SetMapIndex](https://golang.org/src/reflect/value.go?s=50058:50101#L1645)
<pre>func (v <a href="#Value">Value</a>) SetMapIndex(key, elem <a href="#Value">Value</a>)</pre>
SetMapIndex sets the element associated with key in the map v to elem.
It panics if v's Kind is not Map.
If elem is the zero Value, SetMapIndex deletes the key from the map.
Otherwise if v holds a nil map, SetMapIndex will panic.
As in Go, key's elem must be assignable to the map's key type,
and elem's value must be assignable to the map's elem type.




### <a id="Value.SetPointer">func</a> (Value) [SetPointer](https://golang.org/src/reflect/value.go?s=51350:51393#L1696)
<pre>func (v <a href="#Value">Value</a>) SetPointer(x <a href="/pkg/unsafe/">unsafe</a>.<a href="/pkg/unsafe/#Pointer">Pointer</a>)</pre>
SetPointer sets the unsafe.Pointer value v to x.
It panics if v's Kind is not UnsafePointer.




### <a id="Value.SetString">func</a> (Value) [SetString](https://golang.org/src/reflect/value.go?s=51586:51620#L1704)
<pre>func (v <a href="#Value">Value</a>) SetString(x <a href="/pkg/builtin/#string">string</a>)</pre>
SetString sets v's underlying value to x.
It panics if v's Kind is not String or if CanSet() is false.




### <a id="Value.SetUint">func</a> (Value) [SetUint](https://golang.org/src/reflect/value.go?s=50833:50865#L1674)
<pre>func (v <a href="#Value">Value</a>) SetUint(x <a href="/pkg/builtin/#uint64">uint64</a>)</pre>
SetUint sets v's underlying value to x.
It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64, or if CanSet() is false.




### <a id="Value.Slice">func</a> (Value) [Slice](https://golang.org/src/reflect/value.go?s=51845:51881#L1713)
<pre>func (v <a href="#Value">Value</a>) Slice(i, j <a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
Slice returns v[i:j].
It panics if v's Kind is not Array, Slice or String, or if v is an unaddressable array,
or if the indexes are out of bounds.




### <a id="Value.Slice3">func</a> (Value) [Slice3](https://golang.org/src/reflect/value.go?s=53463:53503#L1775)
<pre>func (v <a href="#Value">Value</a>) Slice3(i, j, k <a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
Slice3 is the 3-index form of the slice operation: it returns v[i:j:k].
It panics if v's Kind is not Array or Slice, or if v is an unaddressable array,
or if the indexes are out of bounds.




### <a id="Value.String">func</a> (Value) [String](https://golang.org/src/reflect/value.go?s=55048:55078#L1830)
<pre>func (v <a href="#Value">Value</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns the string v's underlying value, as a string.
String is a special case because of Go's String method convention.
Unlike the other getters, it does not panic if v's Kind is not String.
Instead, it returns a string of the form "<T value>" where T is v's type.
The fmt package treats Values specially. It does not call their String
method implicitly but instead prints the concrete values they hold.




### <a id="Value.TryRecv">func</a> (Value) [TryRecv](https://golang.org/src/reflect/value.go?s=55748:55791#L1847)
<pre>func (v <a href="#Value">Value</a>) TryRecv() (x <a href="#Value">Value</a>, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
TryRecv attempts to receive a value from the channel v but will not block.
It panics if v's Kind is not Chan.
If the receive delivers a value, x is the transferred value and ok is true.
If the receive cannot finish without blocking, x is the zero Value and ok is false.
If the channel is closed, x is the zero value for the channel's element type and ok is false.




### <a id="Value.TrySend">func</a> (Value) [TrySend](https://golang.org/src/reflect/value.go?s=56074:56110#L1857)
<pre>func (v <a href="#Value">Value</a>) TrySend(x <a href="#Value">Value</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
TrySend attempts to send x on the channel v but will not block.
It panics if v's Kind is not Chan.
It reports whether the value was sent.
As in Go, x's value must be assignable to the channel's element type.




### <a id="Value.Type">func</a> (Value) [Type](https://golang.org/src/reflect/value.go?s=56202:56228#L1864)
<pre>func (v <a href="#Value">Value</a>) Type() <a href="#Type">Type</a></pre>
Type returns v's type.




### <a id="Value.Uint">func</a> (Value) [Uint](https://golang.org/src/reflect/value.go?s=57071:57099#L1897)
<pre>func (v <a href="#Value">Value</a>) Uint() <a href="/pkg/builtin/#uint64">uint64</a></pre>
Uint returns v's underlying value, as a uint64.
It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64.




### <a id="Value.UnsafeAddr">func</a> (Value) [UnsafeAddr](https://golang.org/src/reflect/value.go?s=57609:57644#L1920)
<pre>func (v <a href="#Value">Value</a>) UnsafeAddr() <a href="/pkg/builtin/#uintptr">uintptr</a></pre>
UnsafeAddr returns a pointer to v's data.
It is for advanced clients that also import the "unsafe" package.
It panics if v is not addressable.




## <a id="ValueError">type</a> [ValueError](https://golang.org/src/reflect/value.go?s=4981:5035#L147)
A ValueError occurs when a Value method is invoked on
a Value that does not support it. Such cases are documented
in the description of each method.


<pre>type ValueError struct {
<span id="ValueError.Method"></span>    Method <a href="/pkg/builtin/#string">string</a>
<span id="ValueError.Kind"></span>    Kind   <a href="#Kind">Kind</a>
}
</pre>











### <a id="ValueError.Error">func</a> (\*ValueError) [Error](https://golang.org/src/reflect/value.go?s=5037:5072#L152)
<pre>func (e *<a href="#ValueError">ValueError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>







