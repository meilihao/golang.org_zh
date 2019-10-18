

# reflect
`import "reflect"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package reflect implements run-time reflection, allowing a program to
manipulate objects with arbitrary types. The typical use is to take a value
with static type interface{} and extract its dynamic type information by
calling TypeOf, which returns a Type.

A call to ValueOf returns a Value representing the run-time data.
Zero takes a Type and returns a Value representing a zero value
for that type.

See "The Laws of Reflection" for an introduction to reflection in Go:
<a href="https://golang.org/doc/articles/laws_of_reflection.html">https://golang.org/doc/articles/laws_of_reflection.html</a>

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
Copy copies the contents of src into dst until either
dst has been filled or src has been exhausted.
It returns the number of elements copied.
Dst and src each must have kind Slice or Array, and
dst and src must have the same element type.

As a special case, src can have kind String if the element type of dst is kind Uint8.

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

DeepEqual 返回x和y是否"深度相等". 定义如下, 如果满足以下条件之一, 则两个相同类型的值是深度相等; 不同类型的值是永远不会深度相等.

如果数组对应的元素是深度相等, 那么它们就是深度相等.

如果struc对应的字段(无论是可导出的或不可导出的)是深度相等, 那么它们是深度相等.

如果func都是nil, 那么它们是深度相等; 否则不相等.

如果Interface中实际存储的值是深度相等, 那么它们是深度相等.

map深度相等需满足以下所有条件: 它们都是nil或都不是nil, 它们有相同的长度, 且它们是相同的map对象或对应key(使用Go相等匹配)映射的值是深度相等.

如果Pointer用Go的`==`比较返回相等或它们指向的值是深度相等, 那么它们是深度相等.

slice深度相等需满足以下所有条件: 它们都是nil或都不是nil, 它们有相同的长度, 且它们指向的底层数组相同(即&x[0] == &y[0])或它们对应的所有元素是深度相等. 注意: 一个非nil的slice和一个nil的slice(比如[]byte{}和[]byte(nil))是不相等的.

其他值: 数值, bool, string和channel, 如果使用Go的`==`比较返回相等, 那么它们是深度相等的.

通常, DeepEqual 是 Go 的`==`递归形式. 然而, 如果没有一些不一致, 这个想法是不可能实现的. 具体来说, 值可能与它本身不相等, 比如它们是func类型(通常无法比较), 或是浮点类型的NaN(在浮点数比较中它们不等于自身), 或者array, struct, interface包含了这样的值. 另一方面, 指针总是等于自身, 即使它们???指向或包含这样有问题的值???, 因为它们可用Go的`==`比较, 因此无论内容如何, 都满足深度相等的条件. DeepEqual已经定义了对slice和map的快捷比较方式: 如果x和y是相同的slice或map, 那么它们是深度相等的.

但DeepEqual遍历数据时可能会碰到循环. 在第二次及以后比较两个pointer时会沿用之前的相等(能比较到现在那么之前的比较结果必定是相等), 而不是重新检查它们指向的值, 这样可确保DeepEqual能终止.

## <a id="Swapper">func</a> [Swapper](https://golang.org/src/reflect/swapper.go?s=337:383#L3)
<pre>func Swapper(slice interface{}) func(i, j <a href="/pkg/builtin/#int">int</a>)</pre>
Swapper returns a function that swaps the elements in the provided
slice.

Swapper panics if the provided interface is not a slice.

Swapper 返回一个函数, 可交换slice参数中的元素.

如果提供的参数不是slice, Swapper 会 panic.


## <a id="ChanDir">type</a> [ChanDir](https://golang.org/src/reflect/type.go?s=11596:11612#L334)
ChanDir represents a channel type's direction.

ChanDir 返回channel类型的方向.


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

Kind 表示Type类型所表示的特定类型. Kind的零值不是一种有效的类型.


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
```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	for _, v := range []interface{}{"hi", 42, func() {}} {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind %s", v.Kind())
		}
	}

}
```

output:
```txt
hi
42
unhandled kind func
```




### <a id="Kind.String">func</a> (Kind) [String](https://golang.org/src/reflect/type.go?s=17788:17817#L586)
<pre>func (k <a href="#Kind">Kind</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns the name of k.

String 返回 k 的名称


## <a id="MapIter">type</a> [MapIter](https://golang.org/src/reflect/value.go?s=37394:37446#L1211)
A MapIter is an iterator for ranging over a map.
See Value.MapRange.

MapIter 是一个用于遍历map的迭代器. 可参考 Value.MapRange.

<pre>type MapIter struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="MapIter.Key">func</a> (\*MapIter) [Key](https://golang.org/src/reflect/value.go?s=37508:37538#L1217)
<pre>func (it *<a href="#MapIter">MapIter</a>) Key() <a href="#Value">Value</a></pre>
Key returns the key of the iterator's current map entry.

Key 返回迭代器中当前map项的key对应的Value类型.


### <a id="MapIter.Next">func</a> (\*MapIter) [Next](https://golang.org/src/reflect/value.go?s=38411:38441#L1247)
<pre>func (it *<a href="#MapIter">MapIter</a>) Next() <a href="/pkg/builtin/#bool">bool</a></pre>
Next advances the map iterator and reports whether there is another
entry. It returns false when the iterator is exhausted; subsequent
calls to Key, Value, or Next will panic.

Next 使得该迭代器迭代一次并报告是否存在下一项. 迭代完成时它会返回false; 此时调用Key, Value 或 Next 方法会panic.


### <a id="MapIter.Value">func</a> (\*MapIter) [Value](https://golang.org/src/reflect/value.go?s=37895:37927#L1231)
<pre>func (it *<a href="#MapIter">MapIter</a>) Value() <a href="#Value">Value</a></pre>
Value returns the value of the iterator's current map entry.

Value 返回迭代器中当前map项的value对应的Value类型.


## <a id="Method">type</a> [Method](https://golang.org/src/reflect/type.go?s=17129:17626#L564)
Method represents a single method.

Method 代表一个方法.

<pre>type Method struct {
<span id="Method.Name"></span>    <span class="comment">// Name is the method name.</span>
<span id="Method.PkgPath"></span>    <span class="comment">// PkgPath is the package path that qualifies a lower case (unexported)</span>
    <span class="comment">// method name. It is empty for upper case (exported) method names.</span>
    <span class="comment">// The combination of PkgPath and Name uniquely identifies a method</span>
    <span class="comment">// in a method set.</span>
    <span class="comment">// See https://golang.org/ref/spec#Uniqueness_of_identifiers</span>
    <span class="comment">//</span>
    <span class="comment">// Name 是方法的名称. </span>
    <span class="comment">// 当方法是不可导出时, PkgPath是其package路径; 可导出时该字段为空.</span>
    <span class="comment">// PkgPath和Name的组合可唯一标识方法集中的方法.</span>
    <span class="comment">// 可参考 https://golang.org/ref/spec#Uniqueness_of_identifiers</span>
    Name    <a href="/pkg/builtin/#string">string</a>
    PkgPath <a href="/pkg/builtin/#string">string</a>

<span id="Method.Type"></span>    Type  <a href="#Type">Type</a>  <span class="comment">// method type // 方法的类型</span>
<span id="Method.Func"></span>    Func  <a href="#Value">Value</a> <span class="comment">// func with receiver as first argument // func的地址. 该 func会将接收者作为第一个参数</span>
<span id="Method.Index"></span>    Index <a href="/pkg/builtin/#int">int</a>   <span class="comment">// index for Type.Method // Type.Method的索引</span>
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

SelectCase 描述了一次select操作中的一种方案. 方案的类型取决于Dir即通信的方向.

如果Dir是SelectDefault,该方式是默认的. 此时Chan和Send必须是零值.

如果Dir是SelectSend, 该方案表示一次send操作. 通常, Chan的底层值必须是一个channel, 且Send的底层值必须匹配该channel元素类型的值. 特例是, 如果Chan是一个零值, 该方案会被忽略, 且Send(无论是否零值)也会被忽略.

如果Dir是SelectRecv, 该方案表示一次receive操作. 通常, Chan的底层值必须是一个channel, 且Send必须是一个零值. 如果Chan是一个零值, 该方式会被忽略, 但Send必须还是零值. 当选择receive操作时, 收到的 Value 是由 Select 返回.

<pre>type SelectCase struct {
<span id="SelectCase.Dir"></span>    Dir  <a href="#SelectDir">SelectDir</a> <span class="comment">// direction of case</span>
<span id="SelectCase.Chan"></span>    Chan <a href="#Value">Value</a>     <span class="comment">// channel to use (for send or receive)</span>
<span id="SelectCase.Send"></span>    Send <a href="#Value">Value</a>     <span class="comment">// value to send (for send)</span>
}
</pre>











## <a id="SelectDir">type</a> [SelectDir](https://golang.org/src/reflect/value.go?s=62886:62904#L2108)
A SelectDir describes the communication direction of a select case.

SelectDir 表示一次select操作的通信方向.

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

SliceHeader是slice的runtime表示形式. 它不能被安全地和方便地使用, 且该形式可能会在以后的版本中发生变更.
而且, Data 字段不能保证它指向的数据不被gc回收, 因此程序中必须保留一个单独的, 类型正确的并指向底层数据的指针.

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

StringHeader 是string的runtime表示形式. 它不能被安全地和方便地使用, 且该形式可能会在以后的版本中发生变更.
而且, Data 字段不能保证它指向的数据不被gc回收, 因此程序中必须保留一个单独的, 类型正确的并指向底层数据的指针.

<pre>type StringHeader struct {
<span id="StringHeader.Data"></span>    Data <a href="/pkg/builtin/#uintptr">uintptr</a>
<span id="StringHeader.Len"></span>    Len  <a href="/pkg/builtin/#int">int</a>
}
</pre>











## <a id="StructField">type</a> [StructField](https://golang.org/src/reflect/type.go?s=29891:30415#L1085)
A StructField describes a single field in a struct.

StructField 描述了struct的单个字段.

<pre>type StructField struct {
<span id="StructField.Name"></span>    <span class="comment">// Name is the field name. // 字段名称</span>
    Name <a href="/pkg/builtin/#string">string</a>
<span id="StructField.PkgPath"></span>    <span class="comment">// PkgPath is the package path that qualifies a lower case (unexported)</span>
    <span class="comment">// field name. It is empty for upper case (exported) field names. // PkgPath是不可导出字段的package路径. 导出字段的该值为空.</span>
    <span class="comment">// See https://golang.org/ref/spec#Uniqueness_of_identifiers</span>
    PkgPath <a href="/pkg/builtin/#string">string</a>

<span id="StructField.Type"></span>    Type      <a href="#Type">Type</a>      <span class="comment">// field type // 字段类型</span>
<span id="StructField.Tag"></span>    Tag       <a href="#StructTag">StructTag</a> <span class="comment">// field tag string // 字段的tag</span>
<span id="StructField.Offset"></span>    Offset    <a href="/pkg/builtin/#uintptr">uintptr</a>   <span class="comment">// offset within struct, in bytes // 字段在struct的偏移量, 以byte为单位</span>
<span id="StructField.Index"></span>    Index     []<a href="/pkg/builtin/#int">int</a>     <span class="comment">// index sequence for Type.FieldByIndex // ???</span>
<span id="StructField.Anonymous"></span>    Anonymous <a href="/pkg/builtin/#bool">bool</a>      <span class="comment">// is an embedded field // 是否是内嵌字段</span>
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

StructTag 是 struct字段的tag.

按照惯例, 字符串tag是用空格分隔的`key:"value"`对. key是非空字符串, 由空格(U+0020 ' '), 双引号(U+0022 '"')和冒号(U+003A ':')除外的非控制字符组成. value使用"(U+0022)包裹并使用Go字符串字面量的语法.

<pre>type StructTag <a href="/pkg/builtin/#string">string</a></pre>





<a id="example_StructTag">Example</a>
```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	type S struct {
		F string `species:"gopher" color:"blue"`
	}

	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))

}
```

output:
```txt
blue gopher
```




### <a id="StructTag.Get">func</a> (StructTag) [Get](https://golang.org/src/reflect/type.go?s=31144:31187#L1115)
<pre>func (tag <a href="#StructTag">StructTag</a>) Get(key <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Get returns the value associated with key in the tag string.
If there is no such key in the tag, Get returns the empty string.
If the tag does not have the conventional format, the value
returned by Get is unspecified. To determine whether a tag is
explicitly set to the empty string, use Lookup.

Get 返回 tag 中与key关联的value. 如果在tag中没有该key, Get 会返回空字符串. 如果tag没采用惯例格式, 则 Get 返回的value是未知的. 如果要确定一个tag是否被设置成了空字符串, 请使用 Lookup.


### <a id="StructTag.Lookup">func</a> (StructTag) [Lookup](https://golang.org/src/reflect/type.go?s=31621:31684#L1126)
<pre>func (tag <a href="#StructTag">StructTag</a>) Lookup(key <a href="/pkg/builtin/#string">string</a>) (value <a href="/pkg/builtin/#string">string</a>, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
Lookup returns the value associated with key in the tag string.
If the key is present in the tag the value (which may be empty)
is returned. Otherwise the returned value will be the empty string.
The ok return value reports whether the value was explicitly set in
the tag string. If the tag does not have the conventional format,
the value returned by Lookup is unspecified.

Lookup 返回 tag 中与key关联的value. 如果key在tag中会返回value(可能是空字符串), 否则将返回空字符串. 返回值ok可明确表明该tag是否已设置. 如果tag没采用惯例格式, 则 Lookup 返回的value是未知的.


<a id="example_StructTag_Lookup">Example</a>
```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	type S struct {
		F0 string `alias:"field_0"`
		F1 string `alias:""`
		F2 string
	}

	s := S{}
	st := reflect.TypeOf(s)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if alias, ok := field.Tag.Lookup("alias"); ok {
			if alias == "" {
				fmt.Println("(blank)")
			} else {
				fmt.Println(alias)
			}
		} else {
			fmt.Println("(not specified)")
		}
	}

}
```

output:
```txt
field_0
(blank)
(not specified)
```

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

Type 表示 Go的类型.

并非所有的方法都适用于所有的类型. 在每种方法的文档中都注明了限制(如果存在). 在调用特定类型的方法前, 请使用 Kind 方法找出其类型. 调用不适合该类型的方法会导致运行时panic.

Type 可以用`==`比较,所以可作为map的key. 如果它们相等, 那么代表同一种类型.

```go
type Type interface {

    // Align returns the alignment in bytes of a value of
    // this type when allocated in memory.
    //
    // Align 返回 当从内存中分配该类型时会对齐的字节数.
    Align() int

    // FieldAlign returns the alignment in bytes of a value of
    // this type when used as a field in a struct.
    //
    // FieldAlign 返回 该类型作为struct字段时会对齐的字节数.
    FieldAlign() int

    // Method returns the i'th method in the type's method set.
    // It panics if i is not in the range [0, NumMethod()).
    //
    // For a non-interface type T or *T, the returned Method's Type and Func
    // fields describe a function whose first argument is the receiver.
    //
    // For an interface type, the returned Method's Type field gives the
    // method signature, without a receiver, and the Func field is nil.
    //
    // Only exported methods are accessible and they are sorted in
    // lexicographic order.
    //
    // Method 返回Type方法集中的第i个方法.
    // 如果i的范围不是[0, NumMethod()), 那么会panic.
    //
    // 对于非接口类型的T或*T, 它会返回Method类型. Method的Func字段是该func的地址, func会将receiver作为第一个参数.
    //
    // 对于接口类型, Method的Type字段会给出该方法的方法签名, 不包含receiver, 且Func字段为nil.
    //
    // 仅可访问导出的方法, 且它们是按照字典序排列的.
    Method(int) Method

    // MethodByName returns the method with that name in the type's
    // method set and a boolean indicating if the method was found.
    //
    // For a non-interface type T or *T, the returned Method's Type and Func
    // fields describe a function whose first argument is the receiver.
    //
    // For an interface type, the returned Method's Type field gives the
    // method signature, without a receiver, and the Func field is nil.
    MethodByName(string) (Method, bool)

    // NumMethod returns the number of exported methods in the type's method set.
    // NumMethod 返回该类型方法集中可导出的方法的数量.
    NumMethod() int

    // Name returns the type's name within its package for a defined type.
    // For other (non-defined) types it returns the empty string.
    Name() string

    // PkgPath returns a defined type's package path, that is, the import path
    // that uniquely identifies the package, such as "encoding/base64".
    // If the type was predeclared (string, error) or not defined (*T, struct{},
    // []int, or A where A is an alias for a non-defined type), the package path
    // will be the empty string.
    PkgPath() string

    // Size returns the number of bytes needed to store
    // a value of the given type; it is analogous to unsafe.Sizeof.
    Size() uintptr

    // String returns a string representation of the type.
    // The string representation may use shortened package names
    // (e.g., base64 instead of "encoding/base64") and is not
    // guaranteed to be unique among types. To test for type identity,
    // compare the Types directly.
    String() string

    // Kind returns the specific kind of this type.
    Kind() Kind

    // Implements reports whether the type implements the interface type u.
    Implements(u Type) bool

    // AssignableTo reports whether a value of the type is assignable to type u.
    AssignableTo(u Type) bool

    // ConvertibleTo reports whether a value of the type is convertible to type u.
    ConvertibleTo(u Type) bool

    // Comparable reports whether values of this type are comparable.
    Comparable() bool

    // Bits returns the size of the type in bits.
    // It panics if the type's Kind is not one of the
    // sized or unsized Int, Uint, Float, or Complex kinds.
    Bits() int

    // ChanDir returns a channel type's direction.
    // It panics if the type's Kind is not Chan.
    ChanDir() ChanDir

    // IsVariadic reports whether a function type's final input parameter
    // is a "..." parameter. If so, t.In(t.NumIn() - 1) returns the parameter's
    // implicit actual type []T.
    //
    // For concreteness, if t represents func(x int, y ... float64), then
    //
    //	t.NumIn() == 2
    //	t.In(0) is the reflect.Type for "int"
    //	t.In(1) is the reflect.Type for "[]float64"
    //	t.IsVariadic() == true
    //
    // IsVariadic panics if the type's Kind is not Func.
    IsVariadic() bool

    // Elem returns a type's element type.
    // It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
    Elem() Type

    // Field returns a struct type's i'th field.
    // It panics if the type's Kind is not Struct.
    // It panics if i is not in the range [0, NumField()).
    Field(i int) StructField

    // FieldByIndex returns the nested field corresponding
    // to the index sequence. It is equivalent to calling Field
    // successively for each index i.
    // It panics if the type's Kind is not Struct.
    FieldByIndex(index []int) StructField

    // FieldByName returns the struct field with the given name
    // and a boolean indicating if the field was found.
    FieldByName(name string) (StructField, bool)

    // FieldByNameFunc returns the struct field with a name
    // that satisfies the match function and a boolean indicating if
    // the field was found.
    //
    // FieldByNameFunc considers the fields in the struct itself
    // and then the fields in any embedded structs, in breadth first order,
    // stopping at the shallowest nesting depth containing one or more
    // fields satisfying the match function. If multiple fields at that depth
    // satisfy the match function, they cancel each other
    // and FieldByNameFunc returns no match.
    // This behavior mirrors Go's handling of name lookup in
    // structs containing embedded fields.
    FieldByNameFunc(match func(string) bool) (StructField, bool)

    // In returns the type of a function type's i'th input parameter.
    // It panics if the type's Kind is not Func.
    // It panics if i is not in the range [0, NumIn()).
    In(i int) Type

    // Key returns a map type's key type.
    // It panics if the type's Kind is not Map.
    Key() Type

    // Len returns an array type's length.
    // It panics if the type's Kind is not Array.
    Len() int

    // NumField returns a struct type's field count.
    // It panics if the type's Kind is not Struct.
    NumField() int

    // NumIn returns a function type's input parameter count.
    // It panics if the type's Kind is not Func.
    NumIn() int

    // NumOut returns a function type's output parameter count.
    // It panics if the type's Kind is not Func.
    NumOut() int

    // Out returns the type of a function type's i'th output parameter.
    // It panics if the type's Kind is not Func.
    // It panics if i is not in the range [0, NumOut()).
    Out(i int) Type
    // contains filtered or unexported methods
}
```



### <a id="ArrayOf">func</a> [ArrayOf](https://golang.org/src/reflect/type.go?s=78277:78316#L2801)
<pre>func ArrayOf(count <a href="/pkg/builtin/#int">int</a>, elem <a href="#Type">Type</a>) <a href="#Type">Type</a></pre>
ArrayOf returns the array type with the given count and element type.
For example, if t represents int, ArrayOf(5, t) represents [5]int.

If the resulting type would be larger than the available address space,
ArrayOf panics.

ArrayOf 返回指定个数和元素类型的数组类型. 例如, 如果t是int, 那么`ArrayOf(5, t)`表示`[5]int`.

如果生成的类型大于可用空间, 那么ArrayOf 会panic.


### <a id="ChanOf">func</a> [ChanOf](https://golang.org/src/reflect/type.go?s=49752:49789#L1766)
<pre>func ChanOf(dir <a href="#ChanDir">ChanDir</a>, t <a href="#Type">Type</a>) <a href="#Type">Type</a></pre>
ChanOf returns the channel type with the given direction and element type.
For example, if t represents int, ChanOf(RecvDir, t) represents <-chan int.

The gc runtime imposes a limit of 64 kB on channel element types.
If t's size is equal to or exceeds this limit, ChanOf panics.

ChanOf 返回指定通信方向和元素类型的channel类型. 例如, 如果t是int, `ChanOf(RecvDir, t)`表示`<-chan int`.

gc将通道元素类型限定在64k内. 如果它等于或大于限定, ChanOf 会panic.

### <a id="FuncOf">func</a> [FuncOf](https://golang.org/src/reflect/type.go?s=53751:53798#L1921)
<pre>func FuncOf(in, out []<a href="#Type">Type</a>, variadic <a href="/pkg/builtin/#bool">bool</a>) <a href="#Type">Type</a></pre>
FuncOf returns the function type with the given argument and result types.
For example if k represents int and e represents string,
FuncOf([]Type{k}, []Type{e}, false) represents func(int) string.

The variadic argument controls whether the function is variadic. FuncOf
panics if the in[len(in)-1] does not represent a slice and variadic is
true.

FuncOf 返回指定输入输出类型的function类型. 例如, 如果k是int, e是string,
`FuncOf([]Type{k}, []Type{e}, false)`表示`func(int) string`.

参数variadic表示该函数是否是可变的. 如果`in[len(in)-1]`不是slice且variadic是true, FuncOf 会panic.


### <a id="MapOf">func</a> [MapOf](https://golang.org/src/reflect/type.go?s=51294:51325#L1823)
<pre>func MapOf(key, elem <a href="#Type">Type</a>) <a href="#Type">Type</a></pre>
MapOf returns the map type with the given key and element types.
For example, if k represents int and e represents string,
MapOf(k, e) represents map[int]string.

If the key type is not a valid map key type (that is, if it does
not implement Go's == operator), MapOf panics.

MapOf 返回指定key和value类型的map类型. 例如, 如果k是int, e是string, `MapOf(k, e)`表示`map[int]string`.

如果key不是有效的map key(即它没有实现Go的`==`运算符). MapOf 会panic.


### <a id="PtrTo">func</a> [PtrTo](https://golang.org/src/reflect/type.go?s=39095:39118#L1373)
<pre>func PtrTo(t <a href="#Type">Type</a>) <a href="#Type">Type</a></pre>
PtrTo returns the pointer type with element t.
For example, if t represents type Foo, PtrTo(t) represents *Foo.

PtrTo 返回元素类型的指针. 例如, 如果t是Foo类型, `PtrTo(t)`表示`*Foo`.


### <a id="SliceOf">func</a> [SliceOf](https://golang.org/src/reflect/type.go?s=62434:62459#L2243)
<pre>func SliceOf(t <a href="#Type">Type</a>) <a href="#Type">Type</a></pre>
SliceOf returns the slice type with element type t.
For example, if t represents int, SliceOf(t) represents []int.

SliceOf 返回元素类型是t的slice类型.
例如, 如果t是int, `SliceOf(t)`表示`[]int`.


### <a id="StructOf">func</a> [StructOf](https://golang.org/src/reflect/type.go?s=64711:64751#L2324)
<pre>func StructOf(fields []<a href="#StructField">StructField</a>) <a href="#Type">Type</a></pre>
StructOf returns the struct type containing fields.
The Offset and Index fields are ignored and computed as they would be
by the compiler.

StructOf currently does not generate wrapper methods for embedded
fields and panics if passed unexported StructFields.
These limitations may be lifted in a future version.

StructOf 返回包含指定字段的struct类型. 字段中的Offset和Index会被忽略, 由编译器重新计算.

如果传入不可导出字段, StructOf 当前不会为嵌入字段生成封装方法并会panic. 这些限制可能会在将来的版本中取消.


<a id="example_StructOf">Example</a>
```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	typ := reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(float64(0)),
			Tag:  `json:"height"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(int(0)),
			Tag:  `json:"age"`,
		},
	})

	v := reflect.New(typ).Elem()
	v.Field(0).SetFloat(0.4)
	v.Field(1).SetInt(2)
	s := v.Addr().Interface()

	w := new(bytes.Buffer)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}

	fmt.Printf("value: %+v\n", s)
	fmt.Printf("json:  %s", w.Bytes())

	r := bytes.NewReader([]byte(`{"height":1.5,"age":10}`))
	if err := json.NewDecoder(r).Decode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)

}
```

output:
```txt
value: &{Height:0.4 Age:2}
json:  {"height":0.4,"age":2}
value: &{Height:1.5 Age:10}
```


### <a id="TypeOf">func</a> [TypeOf](https://golang.org/src/reflect/type.go?s=38787:38818#L1363)
<pre>func TypeOf(i interface{}) <a href="#Type">Type</a></pre>
TypeOf returns the reflection Type that represents the dynamic type of i.
If i is a nil interface value, TypeOf returns nil.

TypeOf 返回 动态类型i的反射类型. 如果i是interface{}(nil), TypeOf会返回nil.

<a id="example_TypeOf">Example</a>
```go
package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	// As interface types are only used for static typing, a
	// common idiom to find the reflection Type for an interface
	// type Foo is to use a *Foo value.
    //
    // 由于interface仅可用于静态类型, 因此为接口类型Foo查找其反射类型通常使用`*Foo`.
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()

	fileType := reflect.TypeOf((*os.File)(nil))
	fmt.Println(fileType.Implements(writerType))

}
```

output:
```txt
true
```


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

Value 是Go值的反射表示.

并非所有方法都适用于所有类型的值. 在每种方法的文档中都注明了限制(如果存在). 在调用特定类型的方法前, 请使用 Kind 方法找出其类型. 调用不适合该类型的方法会导致运行时panic.

Value 的零值表示没有值. 它的 IsValid 方法会返回false, Kind 的方法会返回 Invalid, String 方法会返回"<invalid Value>", 除此以外的方法都会panic. 大多数函数和方法从不返回无效的值. 如果真有, 其文档会明确地说明条件.

Value 可用由多个goroutine并发使用, 前提是其底层的Go值同样可以被并发使用.

要比较两个 Value, 请用它们 Interface 方法的返回值来比较, 直接使用`==`比较并不会比较它们所表示的底层值.

<pre>type Value struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="Append">func</a> [Append](https://golang.org/src/reflect/value.go?s=60336:60374#L2014)
<pre>func Append(s <a href="#Value">Value</a>, x ...<a href="#Value">Value</a>) <a href="#Value">Value</a></pre>
Append appends the values x to a slice s and returns the resulting slice.
As in Go, each x's value must be assignable to the slice's element type.

Append 会将x append到s(slice)中并返回slice. 与Go中(正常操作)一样, x的value必须要匹配slice的元素类型.


### <a id="AppendSlice">func</a> [AppendSlice](https://golang.org/src/reflect/value.go?s=60643:60677#L2025)
<pre>func AppendSlice(s, t <a href="#Value">Value</a>) <a href="#Value">Value</a></pre>
AppendSlice appends a slice t to a slice s and returns the resulting slice.
The slices s and t must have the same element type.

AppendSlice 会将t append到s并返回slice. s和t的元素类型必须一致.


### <a id="Indirect">func</a> [Indirect](https://golang.org/src/reflect/value.go?s=68611:68639#L2297)
<pre>func Indirect(v <a href="#Value">Value</a>) <a href="#Value">Value</a></pre>
Indirect returns the value that v points to.
If v is a nil pointer, Indirect returns a zero Value.
If v is not a pointer, Indirect returns v.

Indirect 返回 v 所指向的值. 如果v是nil指针, Indirect 会返回一个 Value 的零值. 如果v不是指针, Indirect 会panic.


### <a id="MakeChan">func</a> [MakeChan](https://golang.org/src/reflect/value.go?s=67671:67712#L2263)
<pre>func MakeChan(typ <a href="#Type">Type</a>, buffer <a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
MakeChan creates a new channel with the specified type and buffer size.

MakeChan 使用指定类型和缓冲大小创建一个channel.


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

MakeFunc 根据给定类型封装函数fn并返回一个新函数, 新函数被调用时会执行如下操作:


        - 将它的参数转换成`[]Value`形式.
        - 执行 results := fn(args).
        - 以`[]Value`形式返回执行结果, 再转成给定类型的返回值对应的具体类型.

Value.Call 方法允许调用者根据Values调用一个类型确定的函数. 相反, MakeFunc允许调用者根据Values实现一个类型确定的函数.

下面的例子演示了如何使用MakeFunc为不同类型构建swap函数.

<a id="example_MakeFunc">Example</a>
```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// swap is the implementation passed to MakeFunc.
	// It must work in terms of reflect.Values so that it is possible
	// to write code without knowing beforehand what the types
	// will be.
    //
    // swap 是会传递给MakeFunc的实现.
    // 它必须基于reflect.Value来实现, 以便于在有可能不知道类型的情况下编写代码.
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}

	// makeSwap expects fptr to be a pointer to a nil function.
	// It sets that pointer to a new function created with MakeFunc.
	// When the function is invoked, reflect turns the arguments
	// into Values, calls swap, and then turns swap's result slice
	// into the values returned by the new function.
    //
    // makeSwap 期望fptr是指向nil函数的指针.
    // 它会将指针指向由MakeFunc创建的新函数.
    // 调用时, reflect将参数转换为Values, 然后调用swap, 接着将swap的slice返回值转化为新函数的返回值.
	makeSwap := func(fptr interface{}) {
		// fptr is a pointer to a function.
		// Obtain the function value itself (likely nil) as a reflect.Value
		// so that we can query its type and then set the value.
        //
        // fptr是指向函数的指针. 通过reflect.Value获取函数本身(可能为nil), 以便于我们查询其类型, 然后赋值.
		fn := reflect.ValueOf(fptr).Elem()

		// Make a function of the right type.
        // 使用正确的类型来生成函数.
		v := reflect.MakeFunc(fn.Type(), swap)

		// Assign it to the value fn represents.
        // 将其赋值给fn.
		fn.Set(v)
	}

	// Make and call a swap function for ints.
    // 为ints生成并调用swap
	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(0, 1))

	// Make and call a swap function for float64s.
    // 为float64s生成并调用swap
	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	fmt.Println(floatSwap(2.72, 3.14))

}
```

output:
```txt
1 0
3.14 2.72
```

### <a id="MakeMap">func</a> [MakeMap](https://golang.org/src/reflect/value.go?s=68085:68113#L2279)
<pre>func MakeMap(typ <a href="#Type">Type</a>) <a href="#Value">Value</a></pre>
MakeMap creates a new map with the specified type.

MakeMap 根据指定类型创建一个map.


### <a id="MakeMapWithSize">func</a> [MakeMapWithSize](https://golang.org/src/reflect/value.go?s=68263:68306#L2285)
<pre>func MakeMapWithSize(typ <a href="#Type">Type</a>, n <a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
MakeMapWithSize creates a new map with the specified type
and initial space for approximately n elements.

MakeMapWithSize 根据指定类型和初始空间大小创建一个map.


### <a id="MakeSlice">func</a> [MakeSlice](https://golang.org/src/reflect/value.go?s=67146:67190#L2244)
<pre>func MakeSlice(typ <a href="#Type">Type</a>, len, cap <a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
MakeSlice creates a new zero-initialized slice value
for the specified slice type, length, and capacity.

MakeSlice 根据指定slice类型, length和capacity创建已初始化好的slice.


### <a id="New">func</a> [New](https://golang.org/src/reflect/value.go?s=69829:69853#L2339)
<pre>func New(typ <a href="#Type">Type</a>) <a href="#Value">Value</a></pre>
New returns a Value representing a pointer to a new zero value
for the specified type. That is, the returned Value's Type is PtrTo(typ).

New 返回一个Value, 其表示指向给定类型零值的指针, 即返回Value的类型是`PtrTo(typ)`.


### <a id="NewAt">func</a> [NewAt](https://golang.org/src/reflect/value.go?s=70110:70154#L2351)
<pre>func NewAt(typ <a href="#Type">Type</a>, p <a href="/pkg/unsafe/">unsafe</a>.<a href="/pkg/unsafe/#Pointer">Pointer</a>) <a href="#Value">Value</a></pre>
NewAt returns a Value representing a pointer to a value of the
specified type, using p as that pointer.

NewAt 返回一个Value, 其表示指向给定类型零值的指针, 并使用p作为那个指针.


### <a id="Select">func</a> [Select](https://golang.org/src/reflect/value.go?s=64713:64782#L2149)
<pre>func Select(cases []<a href="#SelectCase">SelectCase</a>) (chosen <a href="/pkg/builtin/#int">int</a>, recv <a href="#Value">Value</a>, recvOK <a href="/pkg/builtin/#bool">bool</a>)</pre>
Select executes a select operation described by the list of cases.
Like the Go select statement, it blocks until at least one of the cases
can proceed, makes a uniform pseudo-random choice,
and then executes that case. It returns the index of the chosen case
and, if that case was a receive operation, the value received and a
boolean indicating whether the value corresponds to a send on the channel
(as opposed to a zero value received because the channel is closed).

Select 执行给定cases参数中的一个Select操作. 像Go select语法, 它会阻塞直到至少一种情况可用, 再做一个伪随机的选择, 然后执行它. 它返回所选方案的索引, 如果是接收操作, 返回收到的Value和一个bool, 该bool表示Value是否是send到channel上的值(而不是因为channel已关闭而返回的零值)


### <a id="ValueOf">func</a> [ValueOf](https://golang.org/src/reflect/value.go?s=68830:68863#L2306)
<pre>func ValueOf(i interface{}) <a href="#Value">Value</a></pre>
ValueOf returns a new Value initialized to the concrete value
stored in the interface i. ValueOf(nil) returns the zero Value.

ValueOf 返回一个已初始化为i中保存的具体值的Value.  ValueOf(nil)返回一个Value的零值.


### <a id="Zero">func</a> [Zero](https://golang.org/src/reflect/value.go?s=69466:69491#L2325)
<pre>func Zero(typ <a href="#Type">Type</a>) <a href="#Value">Value</a></pre>
Zero returns a Value representing the zero value for the specified type.
The result is different from the zero value of the Value struct,
which represents no value at all.
For example, Zero(TypeOf(42)) returns a Value with Kind Int and value 0.
The returned value is neither addressable nor settable.

Zero 返回一个表示指定类型零值的Value. 该值与Value的零值不同, 因为Value的零值不表示任何值. 例如, `Zero(TypeOf(42))`返回一个Value, 其类型是 Int, 具体内容是0. 返回的值即不可用寻址, 也不可以设置.




### <a id="Value.Addr">func</a> (Value) [Addr](https://golang.org/src/reflect/value.go?s=7759:7786#L246)
<pre>func (v <a href="#Value">Value</a>) Addr() <a href="#Value">Value</a></pre>
Addr returns a pointer value representing the address of v.
It panics if CanAddr() returns false.
Addr is typically used to obtain a pointer to a struct field
or slice element in order to call a method that requires a
pointer receiver.

Addr 返回表示v地址的指针的Value. 如果 CanAddr()返回false, 它会panic. Addr 通常用于获取指向struct字段或slice元素的指针, 以便于调用需要指针接收器的方法.


### <a id="Value.Bool">func</a> (Value) [Bool](https://golang.org/src/reflect/value.go?s=8012:8038#L255)
<pre>func (v <a href="#Value">Value</a>) Bool() <a href="/pkg/builtin/#bool">bool</a></pre>
Bool returns v's underlying value.
It panics if v's kind is not Bool.

Bool 返回v的底层值. 如果v的类型不是 Bool 则会panic.


### <a id="Value.Bytes">func</a> (Value) [Bytes](https://golang.org/src/reflect/value.go?s=8185:8214#L262)
<pre>func (v <a href="#Value">Value</a>) Bytes() []<a href="/pkg/builtin/#byte">byte</a></pre>
Bytes returns v's underlying value.
It panics if v's underlying value is not a slice of bytes.

Bytes 返回v的底层值. 如果v的类型不是slice或bytes 则会panic.


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

Call 使用入参in调用函数v. 例如如果len(in)==3, 那么v.Call(in)等价于Go的调用`v(in[0], in[1], in[2])`. v的类型不是 Func, 那么它会panic. 与在Go中一样, 每个入参必须匹配函数对应的入参类型. 如果v是一个可变参数的函数, Call 会自己创建保存可变参数的slice,并复制相应的值.


### <a id="Value.CallSlice">func</a> (Value) [CallSlice](https://golang.org/src/reflect/value.go?s=10552:10596#L321)
<pre>func (v <a href="#Value">Value</a>) CallSlice(in []<a href="#Value">Value</a>) []<a href="#Value">Value</a></pre>
CallSlice calls the variadic function v with the input arguments in,
assigning the slice in[len(in)-1] to v's final variadic argument.
For example, if len(in) == 3, v.CallSlice(in) represents the Go call v(in[0], in[1], in[2]...).
CallSlice panics if v's Kind is not Func or if v is not variadic.
It returns the output results as Values.
As in Go, each input argument must be assignable to the
type of the function's corresponding input parameter.


CallSlice 使用入参in调用带可变参数的函数v, 并将`in[len(in)-1]`分配给v的可变参数.
例如, 如果len(in) == 3, 那么`v.CallSlice(in)等价于`Go的调用`v(in[0], in[1], in[2]...)`. 如果v的类型不是 Func 或 v 不是可变参数的函数, 那么它会panic. 它返回的类型是`[]Value`. 与在Go中一样, 每个入参必须匹配函数对应的入参类型.


### <a id="Value.CanAddr">func</a> (Value) [CanAddr](https://golang.org/src/reflect/value.go?s=9081:9110#L287)
<pre>func (v <a href="#Value">Value</a>) CanAddr() <a href="/pkg/builtin/#bool">bool</a></pre>
CanAddr reports whether the value's address can be obtained with Addr.
Such values are called addressable. A value is addressable if it is
an element of a slice, an element of an addressable array,
a field of an addressable struct, or the result of dereferencing a pointer.
If CanAddr returns false, calling Addr will panic.

CanAddr 判断能否通过 Addr 获取v的地址. 这样的值称为可寻址的. slice的元素, 可寻址数组的元素, 可寻址struct的字段或解指针的结果都是可寻址的. 如果 CanAddr 返回false, 那么调用Addr会panic.


### <a id="Value.CanInterface">func</a> (Value) [CanInterface](https://golang.org/src/reflect/value.go?s=30628:30662#L980)
<pre>func (v <a href="#Value">Value</a>) CanInterface() <a href="/pkg/builtin/#bool">bool</a></pre>
CanInterface reports whether Interface can be used without panicking.

CanInterface 判断 Interface 是否可用(不会引起panic).


### <a id="Value.CanSet">func</a> (Value) [CanSet](https://golang.org/src/reflect/value.go?s=9425:9453#L296)
<pre>func (v <a href="#Value">Value</a>) CanSet() <a href="/pkg/builtin/#bool">bool</a></pre>
CanSet reports whether the value of v can be changed.
A Value can be changed only if it is addressable and was not
obtained by the use of unexported struct fields.
If CanSet returns false, calling Set or any type-specific
setter (e.g., SetBool, SetInt) will panic.

CanSet 判断v的值是否可以改变. 仅有该Value是可寻址的且不是通过struct的未导出字段获取的才可以改变. 如果 CanSet 返回false, 调用 Set 或其他任意Setter类型(比如SetBool, SetInt等)的函数都会panic.




### <a id="Value.Cap">func</a> (Value) [Cap](https://golang.org/src/reflect/value.go?s=24283:24307#L749)
<pre>func (v <a href="#Value">Value</a>) Cap() <a href="/pkg/builtin/#int">int</a></pre>
Cap returns v's capacity.
It panics if v's Kind is not Array, Chan, or Slice.

Cap 返回v的容量. 如果v的类型不是 Array, Chan, Slice, 那么它会panic


### <a id="Value.Close">func</a> (Value) [Close](https://golang.org/src/reflect/value.go?s=24646:24668#L765)
<pre>func (v <a href="#Value">Value</a>) Close()</pre>
Close closes the channel v.
It panics if v's Kind is not Chan.

Close 关闭channel v. 如果v的类型不是 Chan, 它会panic.


### <a id="Value.Complex">func</a> (Value) [Complex](https://golang.org/src/reflect/value.go?s=24848:24883#L773)
<pre>func (v <a href="#Value">Value</a>) Complex() <a href="/pkg/builtin/#complex128">complex128</a></pre>
Complex returns v's underlying value, as a complex128.
It panics if v's Kind is not Complex64 or Complex128

Complex 返回 x的底层值, 比如complex128. 如果v的类型不是 Complex64 或 Complex128, 那么它会panic.


### <a id="Value.Convert">func</a> (Value) [Convert](https://golang.org/src/reflect/value.go?s=71619:71655#L2399)
<pre>func (v <a href="#Value">Value</a>) Convert(t <a href="#Type">Type</a>) <a href="#Value">Value</a></pre>
Convert returns the value v converted to type t.
If the usual Go conversion rules do not allow conversion
of the value v to type t, Convert panics.

Convert 将v转成t类型. 如果通常的 Go 转换规则不允许将 v 转换为 t 类型, 那么它会panic.


### <a id="Value.Elem">func</a> (Value) [Elem](https://golang.org/src/reflect/value.go?s=25266:25293#L788)
<pre>func (v <a href="#Value">Value</a>) Elem() <a href="#Value">Value</a></pre>
Elem returns the value that the interface v contains
or that the pointer v points to.
It panics if v's Kind is not Interface or Ptr.
It returns the zero Value if v is nil.

Elem 返回接口v中存储的值或指针v指向的值. 如果v的类型不是 Interface 或 Ptr, 那么它会panic. 如果v为nil, 它会返回Value的零值.


### <a id="Value.Field">func</a> (Value) [Field](https://golang.org/src/reflect/value.go?s=26096:26129#L825)
<pre>func (v <a href="#Value">Value</a>) Field(i <a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
Field returns the i'th field of the struct v.
It panics if v's Kind is not Struct or i is out of range.

Field 表示struct v中第i个字段. 如果v的类型不是 Struct 或 i 超出范围, 那么它会panic.


### <a id="Value.FieldByIndex">func</a> (Value) [FieldByIndex](https://golang.org/src/reflect/value.go?s=27159:27205#L857)
<pre>func (v <a href="#Value">Value</a>) FieldByIndex(index []<a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
FieldByIndex returns the nested field corresponding to index.
It panics if v's Kind is not struct.

FieldByIndex 通过索引返回相应的字段. 如果v的类型不是 Struct, 那么它会panic.


### <a id="Value.FieldByName">func</a> (Value) [FieldByName](https://golang.org/src/reflect/value.go?s=27686:27731#L879)
<pre>func (v <a href="#Value">Value</a>) FieldByName(name <a href="/pkg/builtin/#string">string</a>) <a href="#Value">Value</a></pre>
FieldByName returns the struct field with the given name.
It returns the zero Value if no field was found.
It panics if v's Kind is not struct.

FieldByName 通过给定的name查找struct的字段. 如果该字段未找到则返回零值Value. 如果v的类型不是 Struct, 那么它会panic.


### <a id="Value.FieldByNameFunc">func</a> (Value) [FieldByNameFunc](https://golang.org/src/reflect/value.go?s=28036:28097#L891)
<pre>func (v <a href="#Value">Value</a>) FieldByNameFunc(match func(<a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a>) <a href="#Value">Value</a></pre>
FieldByNameFunc returns the struct field with a name
that satisfies the match function.
It panics if v's Kind is not struct.
It returns the zero Value if no field was found.

FieldByNameFunc 返回满足函数match的struct字段. 如果v的类型不是 Struct, 那么它会panic. 如果未找到则返回零值Value.


### <a id="Value.Float">func</a> (Value) [Float](https://golang.org/src/reflect/value.go?s=28307:28337#L900)
<pre>func (v <a href="#Value">Value</a>) Float() <a href="/pkg/builtin/#float64">float64</a></pre>
Float returns v's underlying value, as a float64.
It panics if v's Kind is not Float32 or Float64

Float 返回v的底层值, 比如float64. 如果v的类型不是Float32 或 Float64, 那么它会panic.


### <a id="Value.Index">func</a> (Value) [Index](https://golang.org/src/reflect/value.go?s=28677:28710#L915)
<pre>func (v <a href="#Value">Value</a>) Index(i <a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
Index returns v's i'th element.
It panics if v's Kind is not Array, Slice, or String or i is out of range.

Index 返回v中对应的第i个元素. 如果v的类型不是Array, Slice, 或 String 或 i 超出返回, 那么它会panic.


### <a id="Value.Int">func</a> (Value) [Int](https://golang.org/src/reflect/value.go?s=30233:30259#L961)
<pre>func (v <a href="#Value">Value</a>) Int() <a href="/pkg/builtin/#int64">int64</a></pre>
Int returns v's underlying value, as an int64.
It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64.

Init 返回v的底层值, 比如int64. 如果k的类型不是 Int, Int8, Int16, Int32, 或 Int64, 那么它会panic.



### <a id="Value.Interface">func</a> (Value) [Interface](https://golang.org/src/reflect/value.go?s=30985:31027#L992)
<pre>func (v <a href="#Value">Value</a>) Interface() (i interface{})</pre>
Interface returns v's current value as an interface{}.
It is equivalent to:


	var i interface{} = (v's underlying value)

It panics if the Value was obtained by accessing
unexported struct fields.

Interface 以interface{}的形式返回v的当前值, 等价于:
```go
var i interface{} = (v's underlying value)
```

如果Value是通过struct的未导出字段获取的, 那么它会panic.


### <a id="Value.InterfaceData">func</a> (Value) [InterfaceData](https://golang.org/src/reflect/value.go?s=32081:32122#L1028)
<pre>func (v <a href="#Value">Value</a>) InterfaceData() [2]<a href="/pkg/builtin/#uintptr">uintptr</a></pre>
InterfaceData returns the interface v's value as a uintptr pair.
It panics if v's Kind is not Interface.

InterfaceData 返回接口v中的值(以uintptr对). 如果v的类型不是 Interface, 那么它会panic.


### <a id="Value.IsNil">func</a> (Value) [IsNil](https://golang.org/src/reflect/value.go?s=32868:32895#L1046)
<pre>func (v <a href="#Value">Value</a>) IsNil() <a href="/pkg/builtin/#bool">bool</a></pre>
IsNil reports whether its argument v is nil. The argument must be
a chan, func, interface, map, pointer, or slice value; if it is
not, IsNil panics. Note that IsNil is not always equivalent to a
regular comparison with nil in Go. For example, if v was created
by calling ValueOf with an uninitialized interface variable i,
i==nil will be true but v.IsNil will panic as v will be the zero
Value.

IsNil 判断v是否为nil. v必须是chan, func, interface, map, pointer或slice, 除此之外, 它会panic. 注意: IsNil不总是等同于Go中的与nil的常规比较. 例如, v是通过未初始化的interface变量i调用 ValueOf 创建的, 那么i==nil为true. 但 v.IsNil会panic, 因为v是Value的零值.




### <a id="Value.IsValid">func</a> (Value) [IsValid](https://golang.org/src/reflect/value.go?s=33651:33680#L1071)
<pre>func (v <a href="#Value">Value</a>) IsValid() <a href="/pkg/builtin/#bool">bool</a></pre>
IsValid reports whether v represents a value.
It returns false if v is the zero Value.
If IsValid returns false, all other methods except String panic.
Most functions and methods never return an invalid value.
If one does, its documentation states the conditions explicitly.

IsValid 判断v是否表示一个值. 如果v是零值Value, 那么它会返回false. 如果 IsValid 返回false, 那么除 String 外所有的其他方法都会panic. 大部分函数和方法绝不返回无效值. 如果返回了的话, 其文档会明确说明条件.


### <a id="Value.IsZero">func</a> (Value) [IsZero](https://golang.org/src/reflect/value.go?s=33807:33835#L1077)
<pre>func (v <a href="#Value">Value</a>) IsZero() <a href="/pkg/builtin/#bool">bool</a></pre>
IsZero reports whether v is the zero value for its type.
It panics if the argument is invalid.

IsZero 判断v是否是其类型的零值. 如果参数无效, 那么它会panic.


### <a id="Value.Kind">func</a> (Value) [Kind](https://golang.org/src/reflect/value.go?s=34854:34880#L1117)
<pre>func (v <a href="#Value">Value</a>) Kind() <a href="#Kind">Kind</a></pre>
Kind returns v's Kind.
If v is the zero Value (IsValid returns false), Kind returns Invalid.

Kind 返回v的Kind类型. 如果v是零值Value(即IsValid返回false), 那么 Kind 返回 Invalid.


### <a id="Value.Len">func</a> (Value) [Len](https://golang.org/src/reflect/value.go?s=34998:35022#L1123)
<pre>func (v <a href="#Value">Value</a>) Len() <a href="/pkg/builtin/#int">int</a></pre>
Len returns v's length.
It panics if v's Kind is not Array, Chan, Map, Slice, or String.

Len 返回v的长度. 如果v的类型不是Array, Chan, Map, Slice, 或 String, 那么它会panic.


### <a id="Value.MapIndex">func</a> (Value) [MapIndex](https://golang.org/src/reflect/value.go?s=35734:35774#L1147)
<pre>func (v <a href="#Value">Value</a>) MapIndex(key <a href="#Value">Value</a>) <a href="#Value">Value</a></pre>
MapIndex returns the value associated with key in the map v.
It panics if v's Kind is not Map.
It returns the zero Value if key is not found in the map or if v represents a nil map.
As in Go, the key's value must be assignable to the map's key type.


MapIndex 返回map v中key对应的value. 如果v的Kind不是Map, 那么它会panic. 如果key在map中不存在或v是nil的map, 它会返回零值Value. 和在Go中一样, 参数key的值必须与map key的类型一致.

### <a id="Value.MapKeys">func</a> (Value) [MapKeys](https://golang.org/src/reflect/value.go?s=36734:36766#L1180)
<pre>func (v <a href="#Value">Value</a>) MapKeys() []<a href="#Value">Value</a></pre>
MapKeys returns a slice containing all the keys present in the map,
in unspecified order.
It panics if v's Kind is not Map.
It returns an empty slice if v represents a nil map.

MapKeys 返回未排序的包含map中所有key的slice. 如果v的 Kind 不是 Map, 那么它会panic. 如果v是nil的map, 它会返回空的slice.


### <a id="Value.MapRange">func</a> (Value) [MapRange](https://golang.org/src/reflect/value.go?s=39097:39131#L1275)
<pre>func (v <a href="#Value">Value</a>) MapRange() *<a href="#MapIter">MapIter</a></pre>
MapRange returns a range iterator for a map.
It panics if v's Kind is not Map.

Call Next to advance the iterator, and Key/Value to access each entry.
Next returns false when the iterator is exhausted.
MapRange follows the same iteration semantics as a range statement.

MapRange 返回一个map的迭代器.如果v的 Kind 不是 Map, 那么它会panic.

调用 Next 会迭代一次, 通过 Key/Value 可访问每一项. 如果 Next 返回false, 表示迭代已用尽. MapRange 遵循与range相同的迭代语法.

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

Method 返回v的第i个方法. 调用返回的函数不用receiver, 其始终使用v作为receiver. 如果i超出范围或v的接口值是nil, 那么它会panic.


### <a id="Value.MethodByName">func</a> (Value) [MethodByName](https://golang.org/src/reflect/value.go?s=40926:40972#L1329)
<pre>func (v <a href="#Value">Value</a>) MethodByName(name <a href="/pkg/builtin/#string">string</a>) <a href="#Value">Value</a></pre>
MethodByName returns a function value corresponding to the method
of v with the given name.
The arguments to a Call on the returned function should not include
a receiver; the returned function will always use v as the receiver.
It returns the zero Value if no method was found.

MethodByName 根据给定的name返回对应的函数. 调用返回的函数不用receiver, 其始终使用v作为receiver. 如果没有找到返回零值Value.


### <a id="Value.NumField">func</a> (Value) [NumField](https://golang.org/src/reflect/value.go?s=41298:41327#L1345)
<pre>func (v <a href="#Value">Value</a>) NumField() <a href="/pkg/builtin/#int">int</a></pre>
NumField returns the number of fields in the struct v.
It panics if v's Kind is not Struct.

NumField 返回struct v的字段数量. 如果v的 Kind 不是 Struct, 那么它会panic.


### <a id="Value.NumMethod">func</a> (Value) [NumMethod](https://golang.org/src/reflect/value.go?s=40448:40478#L1314)
<pre>func (v <a href="#Value">Value</a>) NumMethod() <a href="/pkg/builtin/#int">int</a></pre>
NumMethod returns the number of exported methods in the value's method set.

NumMethod 返回 v 方法集中可导出方法的数量.

### <a id="Value.OverflowComplex">func</a> (Value) [OverflowComplex](https://golang.org/src/reflect/value.go?s=41562:41611#L1353)
<pre>func (v <a href="#Value">Value</a>) OverflowComplex(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
OverflowComplex reports whether the complex128 x cannot be represented by v's type.
It panics if v's Kind is not Complex64 or Complex128.

OverflowComplex 判断 x 是否不能表示v. 如果v的 Kind 不是 Complex64 或 Complex128, 那么它会panic.


### <a id="Value.OverflowFloat">func</a> (Value) [OverflowFloat](https://golang.org/src/reflect/value.go?s=41956:42000#L1366)
<pre>func (v <a href="#Value">Value</a>) OverflowFloat(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
OverflowFloat reports whether the float64 x cannot be represented by v's type.
It panics if v's Kind is not Float32 or Float64.

OverflowFloat 判断 x 是否不能表示v, 如果v的类型不是Float32 或 Float64, 那么它会panic.


### <a id="Value.OverflowInt">func</a> (Value) [OverflowInt](https://golang.org/src/reflect/value.go?s=42433:42473#L1386)
<pre>func (v <a href="#Value">Value</a>) OverflowInt(x <a href="/pkg/builtin/#int64">int64</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
OverflowInt reports whether the int64 x cannot be represented by v's type.
It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64.

OverflowInt 判断 x 是否不能表示v. 如果v的 Kind 不是Int, Int8, Int16, Int32 或 Int64, 那么它会panic.


### <a id="Value.OverflowUint">func</a> (Value) [OverflowUint](https://golang.org/src/reflect/value.go?s=42866:42908#L1399)
<pre>func (v <a href="#Value">Value</a>) OverflowUint(x <a href="/pkg/builtin/#uint64">uint64</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
OverflowUint reports whether the uint64 x cannot be represented by v's type.
It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64.


OverflowUint 判断 x 是否不能表示v. 如果v的 Kind 不是Uint, Uintptr, Uint8, Uint16, Uint32 或 Uint64, 那么它会panic.

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

Pointer 以uintptr的形式返回v的值. 它用uintptr代替unsafe.Pointer作为返回类型, 因此如果不导入`unsafe`包, 那么反射的代码无法获取到unsafe.Pointer. 如果v的 Kind 不是Chan, Func, Map, Ptr, Slice 或 UnsafePointer, 那么它会panic.

如果v的 Kind 是 Func, 则返回的指针是底层代码的指针, 但不一定足以唯一地标识一个函数. 唯一能保证的是, 如果v是nil的func Value, 那么它会返回0.

如果v的 Kind 是 Slice, 返回的指针指向该slice的第一个元素. 如果slice是nil, 那么返回0. 如果slice是空的但不为nil, 则返回值不为零.

### <a id="Value.Recv">func</a> (Value) [Recv](https://golang.org/src/reflect/value.go?s=45087:45127#L1460)
<pre>func (v <a href="#Value">Value</a>) Recv() (x <a href="#Value">Value</a>, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
Recv receives and returns a value from the channel v.
It panics if v's Kind is not Chan.
The receive blocks until a value is ready.
The boolean value ok is true if the value x corresponds to a send
on the channel, false if it is a zero value received because the channel is closed.

Recv 接收并返回一个channel v中的值. 如果v的 Kind 不是 Chan, 那么它是panic. 该操纵会阻塞直至有一个值是ready. bool类型的ok如果为true, 表示x是被发送到channel上的值; 如果为false, 表示由于channel已关闭, x是零值.


### <a id="Value.Send">func</a> (Value) [Send](https://golang.org/src/reflect/value.go?s=45934:45962#L1493)
<pre>func (v <a href="#Value">Value</a>) Send(x <a href="#Value">Value</a>)</pre>
Send sends x on the channel v.
It panics if v's kind is not Chan or if x's type is not the same type as v's element type.
As in Go, x's value must be assignable to the channel's element type.

Send 将 x 发送到 channel 上. 如果v 的 Kind 不是 Chan 或者 x的类型与v的元素类型一致, 那么它会panic. 与在Go中一样, x的值必须是能赋值给channel的类型.


### <a id="Value.Set">func</a> (Value) [Set](https://golang.org/src/reflect/value.go?s=46618:46645#L1520)
<pre>func (v <a href="#Value">Value</a>) Set(x <a href="#Value">Value</a>)</pre>
Set assigns x to the value v.
It panics if CanSet returns false.
As in Go, x's value must be assignable to v's type.


Set 将 x 赋值给v, 如果 CanSet 返回 false, 那么它会panic. 和在Go中一样, x的值必须是能赋值给v的类型.

### <a id="Value.SetBool">func</a> (Value) [SetBool](https://golang.org/src/reflect/value.go?s=47059:47089#L1537)
<pre>func (v <a href="#Value">Value</a>) SetBool(x <a href="/pkg/builtin/#bool">bool</a>)</pre>
SetBool sets v's underlying value.
It panics if v's Kind is not Bool or if CanSet() is false.

SetBool 设置v的底层值. 如果v的 Kind 不是 Bool 或 CanSet() 返回 false, 那么它会panic.


### <a id="Value.SetBytes">func</a> (Value) [SetBytes](https://golang.org/src/reflect/value.go?s=47255:47288#L1545)
<pre>func (v <a href="#Value">Value</a>) SetBytes(x []<a href="/pkg/builtin/#byte">byte</a>)</pre>
SetBytes sets v's underlying value.
It panics if v's underlying value is not a slice of bytes.

SetBytes 设置v的底层值. 如果v的 底层值不是`[]byte`, 那么它会panic.


### <a id="Value.SetCap">func</a> (Value) [SetCap](https://golang.org/src/reflect/value.go?s=49485:49513#L1629)
<pre>func (v <a href="#Value">Value</a>) SetCap(n <a href="/pkg/builtin/#int">int</a>)</pre>
SetCap sets v's capacity to n.
It panics if v's Kind is not Slice or if n is smaller than the length or
greater than the capacity of the slice.

SetCap 将v的容量设置为n. 如果v的 Kind 不是 Slice 或 n 小于 length 或 大于 slice的现有容量, 那么它会panic.


### <a id="Value.SetComplex">func</a> (Value) [SetComplex](https://golang.org/src/reflect/value.go?s=47875:47914#L1567)
<pre>func (v <a href="#Value">Value</a>) SetComplex(x <a href="/pkg/builtin/#complex128">complex128</a>)</pre>
SetComplex sets v's underlying value to x.
It panics if v's Kind is not Complex64 or Complex128, or if CanSet() is false.

SetComplex 将v的底层值设为x. 如果v 的 Kind 不是 Complex64 或 Complex128 或 CanSet() 返回false, 那么它会panic.


### <a id="Value.SetFloat">func</a> (Value) [SetFloat](https://golang.org/src/reflect/value.go?s=48263:48297#L1581)
<pre>func (v <a href="#Value">Value</a>) SetFloat(x <a href="/pkg/builtin/#float64">float64</a>)</pre>
SetFloat sets v's underlying value to x.
It panics if v's Kind is not Float32 or Float64, or if CanSet() is false.

SetFloat 将v的底层值设为x. 如果v 的 Kind 不是 Float32, Float64 或 CanSet() 返回false, 那么它会panic.


### <a id="Value.SetInt">func</a> (Value) [SetInt](https://golang.org/src/reflect/value.go?s=48645:48675#L1595)
<pre>func (v <a href="#Value">Value</a>) SetInt(x <a href="/pkg/builtin/#int64">int64</a>)</pre>
SetInt sets v's underlying value to x.
It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64, or if CanSet() is false.

SetInt 将v的底层值设为x. 如果v 的 Kind 不是 Int, Int8, Int16, Int32, Int64 或 CanSet() 返回false, 那么它会panic.


### <a id="Value.SetLen">func</a> (Value) [SetLen](https://golang.org/src/reflect/value.go?s=49133:49161#L1616)
<pre>func (v <a href="#Value">Value</a>) SetLen(n <a href="/pkg/builtin/#int">int</a>)</pre>
SetLen sets v's length to n.
It panics if v's Kind is not Slice or if n is negative or
greater than the capacity of the slice.

SetLen 将v的length设置为n. 如果v的 Kind 不是 Slice 或 n 是 负数 或 大于 slice的现有容量, 那么它会panic.


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

SetPointer 将unsafe.Pointer类型的值赋值给x. 如果v的 Kind 不是 UnsafePointer, 那么它会panic.


### <a id="Value.SetString">func</a> (Value) [SetString](https://golang.org/src/reflect/value.go?s=51586:51620#L1704)
<pre>func (v <a href="#Value">Value</a>) SetString(x <a href="/pkg/builtin/#string">string</a>)</pre>
SetString sets v's underlying value to x.
It panics if v's Kind is not String or if CanSet() is false.

SetString 将v的底层值设为x. 如果v的 Kind 不是 String 或 CanSet() 返回 false, 那么它会panic.


### <a id="Value.SetUint">func</a> (Value) [SetUint](https://golang.org/src/reflect/value.go?s=50833:50865#L1674)
<pre>func (v <a href="#Value">Value</a>) SetUint(x <a href="/pkg/builtin/#uint64">uint64</a>)</pre>
SetUint sets v's underlying value to x.
It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64, or if CanSet() is false.

SetUint 将v的底层值设为x. 如果v的 Kind 不是 Uint, Uintptr, Uint8, Uint16, Uint32, Uint64 或 CanSet() 返回false, 那么它会panic.


### <a id="Value.Slice">func</a> (Value) [Slice](https://golang.org/src/reflect/value.go?s=51845:51881#L1713)
<pre>func (v <a href="#Value">Value</a>) Slice(i, j <a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
Slice returns v[i:j].
It panics if v's Kind is not Array, Slice or String, or if v is an unaddressable array,
or if the indexes are out of bounds.

Slice 返回 v[i:j]. 如果v的 Kind 不是 Array, Slice, String 或 v是不可寻址的数组, 或 索引超出范围, 那么它会panic.


### <a id="Value.Slice3">func</a> (Value) [Slice3](https://golang.org/src/reflect/value.go?s=53463:53503#L1775)
<pre>func (v <a href="#Value">Value</a>) Slice3(i, j, k <a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
Slice3 is the 3-index form of the slice operation: it returns v[i:j:k].
It panics if v's Kind is not Array or Slice, or if v is an unaddressable array,
or if the indexes are out of bounds.

Slice3 是 slice的3索引操作形式, 会返回 v[i:j:k]. 如果v的 Kind 不是 Array, Slice 或 v是不可寻址的数组, 或 索引超出范围, 那么它会panic.


### <a id="Value.String">func</a> (Value) [String](https://golang.org/src/reflect/value.go?s=55048:55078#L1830)
<pre>func (v <a href="#Value">Value</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns the string v's underlying value, as a string.
String is a special case because of Go's String method convention.
Unlike the other getters, it does not panic if v's Kind is not String.
Instead, it returns a string of the form "<T value>" where T is v's type.
The fmt package treats Values specially. It does not call their String
method implicitly but instead prints the concrete values they hold.

String 以字符串形式返回字符串v的底层值. 由于Go的 String 方法约定, 因此String 是一种特殊情况. 不像其他的getter, 即使v的 Kind 不是 String, 它也不会panic, 而是以"<T value>"是格式返回字符串, 其中T是v的类型. fmt包会特别对待Value, 它不会隐式地调用它们的 String 方法, 而是直接打印它们持有的实际值.


### <a id="Value.TryRecv">func</a> (Value) [TryRecv](https://golang.org/src/reflect/value.go?s=55748:55791#L1847)
<pre>func (v <a href="#Value">Value</a>) TryRecv() (x <a href="#Value">Value</a>, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
TryRecv attempts to receive a value from the channel v but will not block.
It panics if v's Kind is not Chan.
If the receive delivers a value, x is the transferred value and ok is true.
If the receive cannot finish without blocking, x is the zero Value and ok is false.
If the channel is closed, x is the zero value for the channel's element type and ok is false.

TryRecv 尝试不阻塞地从channel v中接收一个值. 如果v的 Kind 不是 Chan, 那么它会panic. 如果实际传递了一个值, x是传输的值, ok为true. 如果不阻塞就不能收到值，那么x是零值Value, ok为false. 如果channel已关闭, 那么ok为false, x是channel元素类型的零值.


### <a id="Value.TrySend">func</a> (Value) [TrySend](https://golang.org/src/reflect/value.go?s=56074:56110#L1857)
<pre>func (v <a href="#Value">Value</a>) TrySend(x <a href="#Value">Value</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
TrySend attempts to send x on the channel v but will not block.
It panics if v's Kind is not Chan.
It reports whether the value was sent.
As in Go, x's value must be assignable to the channel's element type.

TryRecv 尝试不阻塞地将x发送到channel v上. 如果v的 Kind 不是 Chan, 那么它会panic. 它会返回是否发送成功. 和在Go中一样, x的值必须要匹配channel元素的类型.



### <a id="Value.Type">func</a> (Value) [Type](https://golang.org/src/reflect/value.go?s=56202:56228#L1864)
<pre>func (v <a href="#Value">Value</a>) Type() <a href="#Type">Type</a></pre>
Type returns v's type.

Type 返回v的类型.


### <a id="Value.Uint">func</a> (Value) [Uint](https://golang.org/src/reflect/value.go?s=57071:57099#L1897)
<pre>func (v <a href="#Value">Value</a>) Uint() <a href="/pkg/builtin/#uint64">uint64</a></pre>
Uint returns v's underlying value, as a uint64.
It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64.

Uint 以uint64的形式返回v的底层值. 如果v的 Kind 不是Uint, Uintptr, Uint8, Uint16, Uint32 或 Uint64, 那么它会panic.


### <a id="Value.UnsafeAddr">func</a> (Value) [UnsafeAddr](https://golang.org/src/reflect/value.go?s=57609:57644#L1920)
<pre>func (v <a href="#Value">Value</a>) UnsafeAddr() <a href="/pkg/builtin/#uintptr">uintptr</a></pre>
UnsafeAddr returns a pointer to v's data.
It is for advanced clients that also import the "unsafe" package.
It panics if v is not addressable.

UnsafeAddr 返回一个指向v的数据的指针. 它适用于还导入了`unsafe`包的高级用户. 如果v是不可寻址的, 那么它会panic.


## <a id="ValueError">type</a> [ValueError](https://golang.org/src/reflect/value.go?s=4981:5035#L147)
A ValueError occurs when a Value method is invoked on
a Value that does not support it. Such cases are documented
in the description of each method.

ValueError 在不支持 Value 方法的 Value 上该方法, 则会发生 ValueError. 此类情况已记录在每个方法的说明中.

<pre>type ValueError struct {
<span id="ValueError.Method"></span>    Method <a href="/pkg/builtin/#string">string</a>
<span id="ValueError.Kind"></span>    Kind   <a href="#Kind">Kind</a>
}
</pre>









### <a id="ValueError.Error">func</a> (\*ValueError) [Error](https://golang.org/src/reflect/value.go?s=5037:5072#L152)
<pre>func (e *<a href="#ValueError">ValueError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>


# Bugs
FieldByName and related functions consider struct field names to be equal if the names are equal, even if they are unexported names originating in different packages. The practical effect of this is that the result of t.FieldByName("x") is not well defined if the struct type t contains multiple fields named x (embedded from different packages). FieldByName may return one of the fields named x or may report that there are none. See https://golang.org/issue/4876 for more details.

如果名称相同, 那么FieldByName和相关函数会认为struct字段是相同的, 即使它们是来自不同包的未导出字段. 实际的影响是: 如果struct t包含了多个名为x的字段(嵌在不同的package中), t.FieldByName("x") 的返回结果是不确定的. FieldByName可能会返回它们其中一个名为x的字段或没有找到字段. 详细信息请参考https://golang.org/issue/4876.