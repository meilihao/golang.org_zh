

# builtin
`import "builtin"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package builtin provides documentation for Go's predeclared identifiers.
The items documented here are not actually in package builtin
but their descriptions here allow godoc to present documentation
for the language's special identifiers.

builtin包提供了Go的预声明标识符的文档. 这里列出的条目实际上并不是在本包中,对它们的描述是为了让godoc给golang的特殊标识符提供文档.



## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func append(slice []Type, elems ...Type) []Type](#append)
* [func cap(v Type) int](#cap)
* [func close(c chan<- Type)](#close)
* [func complex(r, i FloatType) ComplexType](#complex)
* [func copy(dst, src []Type) int](#copy)
* [func delete(m map[Type]Type1, key Type)](#delete)
* [func imag(c ComplexType) FloatType](#imag)
* [func len(v Type) int](#len)
* [func make(t Type, size ...IntegerType) Type](#make)
* [func new(Type) *Type](#new)
* [func panic(v interface{})](#panic)
* [func print(args ...Type)](#print)
* [func println(args ...Type)](#println)
* [func real(c ComplexType) FloatType](#real)
* [func recover() interface{}](#recover)
* [type ComplexType](#ComplexType)
* [type FloatType](#FloatType)
* [type IntegerType](#IntegerType)
* [type Type](#Type)
* [type Type1](#Type1)
* [type bool](#bool)
* [type byte](#byte)
* [type complex128](#complex128)
* [type complex64](#complex64)
* [type error](#error)
* [type float32](#float32)
* [type float64](#float64)
* [type int](#int)
* [type int16](#int16)
* [type int32](#int32)
* [type int64](#int64)
* [type int8](#int8)
* [type rune](#rune)
* [type string](#string)
* [type uint](#uint)
* [type uint16](#uint16)
* [type uint32](#uint32)
* [type uint64](#uint64)
* [type uint8](#uint8)
* [type uintptr](#uintptr)




#### <a id="pkg-files">Package files</a>
[builtin.go](https://golang.org/src/builtin/builtin.go) 


## <a id="pkg-constants">Constants</a>

true and false are the two untyped boolean values.

true和false是两个无类型的布尔值.

```go
const (
    true  = 0 == 0 // Untyped bool.
    false = 0 != 0 // Untyped bool.
)
```

iota is a predeclared identifier representing the untyped integer ordinal number of the current const specification in a (usually parenthesized) const declaration. It is zero-indexed.

iota为预声明的标识符，它表示const声明中（通常在括号中）, 无类型的整数的递增序数. 它从0开始索引.

```go
const iota = 0 // Untyped int.
```

## <a id="pkg-variables">Variables</a>

nil is a predeclared identifier representing the zero value for a pointer, channel, func, interface, map, or slice type.

nil为预声明的标识符，它表示pointer, channel, func, interface、map或slice类型的零值.

```go
var nil Type // Type must be a pointer, channel, func, interface, map, or slice type
```

## <a id="append">func append</a>
<pre>func append(slice []Type, elems ...Type) []Type</pre>

The append built-in function appends elements to the end of a slice. If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated. Append returns the updated slice. It is therefore necessary to store the result of append, often in the variable holding the slice itself:

内建函数append会将元素追加到slice的末尾. 若它有足够的容量， 原来的切片最终会容纳新的元素;否则就会重新分配一个新的底层数组. append会返回更新后的slice. 因此必须存储追加后的结果, 通常是赋值给该slice本身的变量:

<pre>
slice = append(slice, elem1, elem2)
slice = append(slice, anotherSlice...)
</pre>

As a special case, it is legal to append a string to a byte slice, like this:

作为一种特殊的情况, 将字符追加到字节数组之后是合法的, 就像这样:

<pre>slice = append([]byte("hello "), "world"...)</pre>

## <a id="cap">func cap</a>
<pre>
func cap(v Type) int
</pre>
The cap built-in function returns the capacity of v, according to its type:

内建函数cap会返回 v 的容量, 这取决于具体类型:
<pre>
Array: the number of elements in v (same as len(v)).
Pointer to array: the number of elements in *v (same as len(v)).
Slice: the maximum length the slice can reach when resliced;
if v is nil, cap(v) is zero.
Channel: the channel buffer capacity, in units of elements;
if v is nil, cap(v) is zero.

Array: v 中元素的数量(和len(v)相等).
*Array: *v 中元素的数量(和len(v)相等).
Slice: (不重新分配底层数组时)切片能接受的最大长度;如果v是nil,那么cap(v)为0.
Channel: channel的缓冲容量;如果v是nil,那么cap(v)为0.
</pre>
For some arguments, such as a simple array expression, the result can be a constant. See the Go language specification's "Length and capacity" section for details.

对于某些参数，例如简单的数组表达式，结果是常量. 有关详细信息，请参见Go语言规范的"Length and capacity"部分.

## <a id="cap">func close</a>
<pre>
func close(c chan<- Type)
</pre>
The close built-in function closes a channel, which must be either bidirectional or send-only. It should be executed only by the sender, never the receiver, and has the effect of shutting down the channel after the last sent value is received. After the last value has been received from a closed channel c, any receive from c will succeed without blocking, returning the zero value for the channel element. The form

内建函数close会关闭一个channel, 该信道必须为双向的或只发送的. 它应当只由发送者执行, 而不是由接收者执行, 在最后发送的值被接收后再关闭会更好. 当最后一个值从已关闭的channel c中取出后, 任何从 c 的接收操作都不会被阻塞, 而是返回该channel元素类型的零值. 形式是:
<pre>
x, ok := <-c
</pre>
will also set ok to false for a closed channel.

在已close的channel上接收将会将ok置为false.

## <a id="complex">func complex</a>
<pre>
func complex(r, i FloatType) ComplexType
</pre>
The complex built-in function constructs a complex value from two floating-point values. The real and imaginary parts must be of the same size, either float32 or float64 (or assignable to them), and the return value will be the corresponding complex type (complex64 for float32, complex128 for float64).

内建函数complex会将两个浮点数值构造成一个复数值. 其实部和虚部的类型必须相同，即 float32 或 float64, 其返回值即为对应的复数类型(complex64对应float32, complex128对应float64).


## <a id="copy">func copy</a>
<pre>
func copy(dst, src []Type) int
</pre>
The copy built-in function copies elements from a source slice into a destination slice. (As a special case, it also will copy bytes from a string to a slice of bytes.) The source and destination may overlap. Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).

内建函数copy将元素从源slice复制到目标slice. (特殊情况是，它也能将从字符串的字节序列复制到byte slice中). 源和目标可能会重叠. copy返回被复制的元素数量, 它会是len(src)和len(dst)中的较小者.

## <a id="delete">func delete</a>
<pre>
func delete(m map[Type]Type1, key Type)
</pre>
The delete built-in function deletes the element with the specified key (m[key]) from the map. If m is nil or there is no such element, delete is a no-op.

内建函数delete会按照指定的key将元素从map中删除. 若m为nil或无此元素,delete操作即为空操作.

## <a id="imag">func imag</a>
<pre>
func imag(c ComplexType) FloatType
</pre>
The imag built-in function returns the imaginary part of the complex number c. The return value will be floating point type corresponding to the type of c.

内建函数imag会返回复数 c 的虚部. 其返回值为对应于 c 类型的浮点类型.

## <a id="len">func len</a>
<pre>
func len(v Type) int
</pre>
The len built-in function returns the length of v, according to its type:

内建函数len会返回 v 的长度, 这取决于具体类型:
<pre>
Array: the number of elements in v.
Pointer to array: the number of elements in *v (even if v is nil).
Slice, or map: the number of elements in v; if v is nil, len(v) is zero.
String: the number of bytes in v.
Channel: the number of elements queued (unread) in the channel buffer;
         if v is nil, len(v) is zero.

Array: v 中元素的数量.
*Array: *v 中元素的数量(即使v为nil).
Slice或者map: v 中元素的数量;如果v是nil,那么len(v)为0.
String: v 中的字节数.
Channel: channel中未取出元素的数量;如果v是nil,那么len(v)为0.
</pre>
For some arguments, such as a string literal or a simple array expression, the result can be a constant. See the Go language specification's "Length and capacity" section for details.

对于某些参数，例如简单的数组表达式或字符串，结果是常量. 有关详细信息，请参见Go语言规范的"Length and capacity"部分.

## <a id="make">func make</a>
<pre>
func make(t Type, size ...IntegerType) Type
</pre>
The make built-in function allocates and initializes an object of type slice, map, or chan (only). Like new, the first argument is a type, not a value. Unlike new, make's return type is the same as the type of its argument, not a pointer to it. The specification of the result depends on the type:

内建函数make仅可以分配并初始化一个类型为slice, map或chan的对象. 与 new 相同的是，其第一个参数为类型而非值. 不同的是， make的返回类型与其参数相同, 而非指向它的指针. 其具体结果取决于具体的类型:

<pre>
Slice: The size specifies the length. The capacity of the slice is
equal to its length. A second integer argument may be provided to
specify a different capacity; it must be no smaller than the
length. For example, make([]int, 0, 10) allocates an underlying array
of size 10 and returns a slice of length 0 and capacity 10 that is
backed by this underlying array.
Map: An empty map is allocated with enough space to hold the
specified number of elements. The size may be omitted, in which case
a small starting size is allocated.
Channel: The channel's buffer is initialized with the specified
buffer capacity. If zero, or the size is omitted, the channel is
unbuffered.

Slice: size 指定了其长度, 该切片的容量等于其长度. 第二个整数参数用于指定不同的容量;它必须不小于其长度,因此 make([]int, 0, 10) 会分配一个长度为0，容量为10的slice.
Map: 会创建一个空map,有足够的空间容纳指定数量的元素.size可以省略, 这种情况下
就会使用一个小的起始大小.
Channel: 初始化一个带指定缓冲容量的channel.若size为零或被省略, 该channel即为无缓冲的.
</pre>

## <a id="new">func new</a>
<pre>
func new(Type) *Type
</pre>

The new built-in function allocates memory. The first argument is a type, not a value, and the value returned is a pointer to a newly allocated zero value of that type.

内建函数new会分配内存. 其第一个参数为类型而非值, 其返回值为指向新分配的该类型的零值的指针.

## <a id="panic">func panic</a>
<pre>
func panic(v interface{})
</pre>
The panic built-in function stops normal execution of the current goroutine. When a function F calls panic, normal execution of F stops immediately. Any functions whose execution was deferred by F are run in the usual way, and then F returns to its caller. To the caller G, the invocation of F then behaves like a call to panic, terminating G's execution and running any deferred functions. This continues until all functions in the executing goroutine have stopped, in reverse order. At that point, the program is terminated with a non-zero exit code. This termination sequence is called panicking and can be controlled by the built-in function recover.

内建函数panic会停止当前goroutine的正常执行. 当函数 F 调用 panic 时, F 的正常执行就会立刻停止. 任何由 F 延迟的函数执行都会照旧执行，接着 F 就返回给其调用者. 对于其调用者 G，调用了F就相当于调用panic，即终止 G 的执行并运行任何被延迟的函数. 这会持续到该goroutine中所有函数都按相反的顺序停止执行为止. 这时, 程序会终止并报告错误情况, 包括引发panic的参数. 此终止顺序被称为恐慌过程, 并可通过内建函数recover 控制.
## <a id="print">func print</a>
<pre>
func print(args ...Type)
</pre>
The print built-in function formats its arguments in an implementation-specific way and writes the result to standard error. Print is useful for bootstrapping and debugging; it is not guaranteed to stay in the language.

内建函数print以特定的方法格式化参数并将结果写入stderr，常用于自举和调试;但不保证未来还会包含在golang里.

## <a id="println">func println</a>
<pre>
func println(args ...Type)
</pre>
The println built-in function formats its arguments in an implementation-specific way and writes the result to standard error. Spaces are always added between arguments and a newline is appended. Println is useful for bootstrapping and debugging; it is not guaranteed to stay in the language.

内建函数println以特定的方法格式化参数并将结果写入stderr, 但会在输出的参数间添加空格,且输出结束后会换行. 常用于自举和调试;但不保证未来还会包含在golang里.

## <a id="real">func real</a>
<pre>
func real(c ComplexType) FloatType
</pre>
The real built-in function returns the real part of the complex number c. The return value will be floating point type corresponding to the type of c.

内建函数real会返回复数 c 的实部. 其返回值为对应于 c 类型的浮点类型.

## <a id="recover">func recover</a>
<pre>
func recover() interface{}
</pre>
The recover built-in function allows a program to manage behavior of a panicking goroutine. Executing a call to recover inside a deferred function (but not any function called by it) stops the panicking sequence by restoring normal execution and retrieves the error value passed to the call of panic. If recover is called outside the deferred function it will not stop a panicking sequence. In this case, or when the goroutine is not panicking, or if the argument supplied to panic was nil, recover returns nil. Thus the return value from recover reports whether the goroutine is panicking.

内建函数recover允许程序管理panic过程中的goroutine. 在defer函数（而不是任何被它调用的函数）中调用 recover 会通过恢复正常的执行顺序并取出传给 panic 调用时的错误值来停止panic过程. 若在defer函数外调用 recover, 它将不会停止panic过程. 在此情况下，若该goroutine没有panic, 或提供给 panic 的参数为 nil，那么 recover 就会返回 nil. 因此 recover 的返回值就反应了该goroutine是否处于panic中.


## <a id="ComplexType">type</a> [ComplexType](https://golang.org/src/builtin/builtin.go?s=4111:4137#L113)
ComplexType is here for the purposes of documentation only. It is a
stand-in for either complex type: complex64 or complex128.


ComplexType 在此只用作文档说明. 它代表复数类型: complex64 或 complex128.


<pre>type ComplexType complex64</pre>











## <a id="FloatType">type</a> [FloatType](https://golang.org/src/builtin/builtin.go?s=3954:3976#L109)
FloatType is here for the purposes of documentation only. It is a stand-in
for either float type: float32 or float64.

FloatType 在此只用作文档说明. 它代表浮点数类型: float32 或 float64.

<pre>type FloatType float32</pre>











## <a id="IntegerType">type</a> [IntegerType](https://golang.org/src/builtin/builtin.go?s=3808:3828#L105)
IntegerType is here for the purposes of documentation only. It is a stand-in
for any integer type: int, uint, int8 etc.

IntegerType 在此只用作文档说明. 它代表整数类型: int, uint, int8等.

<pre>type IntegerType int</pre>











## <a id="Type">type</a> [Type](https://golang.org/src/builtin/builtin.go?s=3490:3503#L96)
Type is here for the purposes of documentation only. It is a stand-in
for any Go type, but represents the same type for any given function
invocation.

Type 在此只用作文档说明. 它代表所有的Go类型， 但对任何给定的函数来说, 它都表示一种相同的类型.

<pre>type Type int</pre>











## <a id="Type1">type</a> [Type1](https://golang.org/src/builtin/builtin.go?s=3666:3680#L101)
Type1 is here for the purposes of documentation only. It is a stand-in
for any Go type, but represents the same type for any given function
invocation.

Type1 在此只用作文档说明. 它代表所有的Go类型， 但对任何给定的函数来说, 它都表示一种相同的类型.


<pre>type Type1 int</pre>



## <a id="bool">type bool</a>

bool is the set of boolean values, true and false.

bool 表示布尔值，即 true 和 false.
<pre>
type bool bool
</pre>
## <a id="byte">type byte</a>

byte is an alias for uint8 and is equivalent to uint8 in all ways. It is used, by convention, to distinguish byte values from 8-bit unsigned integer values.

byte 为 uint8 的别名, 它完全等价于 uint8. 习惯上用它来区别byte值和8位的无符号整数值.
<pre>
type byte = uint8
</pre>
## <a id="complex128">type complex128</a>

complex128 is the set of all complex numbers with float64 real and imaginary parts.

complex128 是实部和虚部都为 float64 的复数.

<pre>
type complex128 complex128
</pre>
## <a id="complex64">type complex64</a>

complex64 is the set of all complex numbers with float32 real and imaginary parts.

complex64 是实部和虚部都为 float32 的复数.
<pre>
type complex64 complex64
</pre>

## <a id="error">type error</a>

The error built-in interface type is the conventional interface for representing an error condition, with the nil value representing no error.

内建接口error表示错误情况的惯用接口, nil 值即表示没有错误.
<pre>
type error interface {
    Error() string
}
</pre>
## <a id="float32">type float32</a>

float32 is the set of all IEEE-754 32-bit floating-point numbers.

float32 表示IEEE-754 32位浮点数.
<pre>
type float32 float32
</pre>
## <a id="float64">type float64</a>

float64 is the set of all IEEE-754 64-bit floating-point numbers.

float64 表示IEEE-754 64位浮点数.
<pre>
type float64 float64
</pre>
## <a id="int">type int</a>

int is a signed integer type that is at least 32 bits in size. It is a distinct type, however, and not an alias for, say, int32.

int 是带符号整数类型，其大小至少为32位. 它是一种独特的类型, 而不是某种别名, 比如int32的.
<pre>
type int int
</pre>
## <a id="int16">type int16</a>

int16 is the set of all signed 16-bit integers. Range: -32768 through 32767.

int16 表示带符号的16位整数, 范围: -32768 ~ 32767.
<pre>
type int16 int16
</pre>
## <a id="int32">type int32</a>

int32 is the set of all signed 32-bit integers. Range: -2147483648 through 2147483647.

int32 表示带符号的32位整数. 范围: -2147483648 ~ 2147483647.
<pre>
type int32 int32
</pre>

## <a id="int64">type int64</a>

int64 is the set of all signed 64-bit integers. Range: -9223372036854775808 through 9223372036854775807.

int64 表示带符号的64位整数. 范围: -9223372036854775808 ~ 9223372036854775807.
<pre>
type int64 int64
</pre>
## <a id="int8">type int8</a>

int8 is the set of all signed 8-bit integers. Range: -128 through 127.

int8 表示带符号的8位整数. 范围: -128 ~ 127.
<pre>
type int8 int8
</pre>
## <a id="rune">type rune</a>

rune is an alias for int32 and is equivalent to int32 in all ways. It is used, by convention, to distinguish character values from integer values.

rune 是 int32 的别名，它完全等价于 int32. 习惯上用它来区别字符值和整数值.
<pre>
type rune = int32
</pre>
## <a id="string">type string</a>

string is the set of all strings of 8-bit bytes, conventionally but not necessarily representing UTF-8-encoded text. A string may be empty, but not nil. Values of string type are immutable.

string 是8位字节字符的集合，习惯上用于表示以UTF-8编码的文本, 但并不必须如此. string 可为空，但不能为 nil. string 类型的值是不可变的.
<pre>
type string string
</pre>
## <a id="uint">type uint</a>

uint is an unsigned integer type that is at least 32 bits in size. It is a distinct type, however, and not an alias for, say, uint32.

uint 是无符号整数类型，其大小至少为32位. 它是一种独特的类型，而不是某种别名, 比如uint32的.
<pre>
type uint uint
</pre>
## <a id="uint16">type uint16</a>

uint16 is the set of all unsigned 16-bit integers. Range: 0 through 65535.

uint16 表示无符号的16位整数. 范围: 0 ~ 65535.
<pre>
type uint16 uint16
</pre>
## <a id="uint32">type uint32</a>

uint32 is the set of all unsigned 32-bit integers. Range: 0 through 4294967295.

uint32 表示无符号的32位整数. 范围: 0 ~ 4294967295.
<pre>
type uint32 uint32
</pre>
## <a id="uint64">type uint64</a>

uint64 is the set of all unsigned 64-bit integers. Range: 0 through 18446744073709551615.

uint64 表示无符号的64位整数. 范围: 0 ~ 18446744073709551615.
<pre>
type uint64 uint64
</pre>
## <a id="uint8">type uint8</a>

uint8 is the set of all unsigned 8-bit integers. Range: 0 through 255.

uint8 表示无符号的8位整数. 范围: 0 ~ 255.
<pre>
type uint8 uint8
</pre>
## <a id="uintptr">type uintptr</a>

uintptr is an integer type that is large enough to hold the bit pattern of any pointer.

uintptr 为整数类型，其大小足以容纳任何指针的位模式.
<pre>
type uintptr uintptr
</pre>