

# unsafe
`import "unsafe"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package unsafe contains operations that step around the type safety of Go programs.

Packages that import unsafe may be non-portable and are not protected by the
Go 1 compatibility guidelines.

unsafe 包含了绕过Go类型安全的操作.

导入unsafe的包可能是不可移植的, 并且不受Go 1 兼容准则的保护.

## <a id="pkg-index">Index</a>
* [func Alignof(x ArbitraryType) uintptr](#Alignof)
* [func Offsetof(x ArbitraryType) uintptr](#Offsetof)
* [func Sizeof(x ArbitraryType) uintptr](#Sizeof)
* [type ArbitraryType](#ArbitraryType)
* [type Pointer](#Pointer)




#### <a id="pkg-files">Package files</a>
[unsafe.go](https://golang.org/src/unsafe/unsafe.go) 






## <a id="Alignof">func</a> [Alignof](https://golang.org/src/unsafe/unsafe.go?s=9177:9214#L195)
<pre>func Alignof(x <a href="#ArbitraryType">ArbitraryType</a>) <a href="/pkg/builtin/#uintptr">uintptr</a></pre>
Alignof takes an expression x of any type and returns the required alignment
of a hypothetical variable v as if v was declared via var v = x.
It is the largest value m such that the address of v is always zero mod m.
It is the same as the value returned by reflect.TypeOf(x).Align().
As a special case, if a variable s is of struct type and f is a field
within that struct, then Alignof(s.f) will return the required alignment
of a field of that type within a struct. This case is the same as the
value returned by reflect.TypeOf(s.f).FieldAlign().
The return value of Alignof is a Go constant.

Alignof 接收任何类型的表达式 x 并返回假想变量v(v是通过var v = x 声明的)所需的对齐值. 它返回m, 因此v的地址始(即当类型进行内存对齐后)终能整除 m. 它与reflect.TypeOf(x).Align()的返回值相同. 特殊情况是, 如果s是struct而f是其中的一个字段, 那么Alignof(s.f)将返回struct中该类型字段所需的对齐值. 这种情况与reflect.TypeOf(s.f).FieldAlign()的返回值相同. Alignof的返回值是Go的常量.

## <a id="Offsetof">func</a> [Offsetof](https://golang.org/src/unsafe/unsafe.go?s=8515:8553#L184)
<pre>func Offsetof(x <a href="#ArbitraryType">ArbitraryType</a>) <a href="/pkg/builtin/#uintptr">uintptr</a></pre>
Offsetof returns the offset within the struct of the field represented by x,
which must be of the form structValue.field. In other words, it returns the
number of bytes between the start of the struct and the start of the field.
The return value of Offsetof is a Go constant.

Offsetof 返回 x (x的格式必须为structValue.field) 表示的字段在struct中的偏移量. 换句话说, 它返回的是struct起始位置与字段起始位置之间的字节数. Offsetof的返回值是Go的常量.

## <a id="Sizeof">func</a> [Sizeof](https://golang.org/src/unsafe/unsafe.go?s=8189:8225#L178)
<pre>func Sizeof(x <a href="#ArbitraryType">ArbitraryType</a>) <a href="/pkg/builtin/#uintptr">uintptr</a></pre>
Sizeof takes an expression x of any type and returns the size in bytes
of a hypothetical variable v as if v was declared via var v = x.
The size does not include any memory possibly referenced by x.
For instance, if x is a slice, Sizeof returns the size of the slice
descriptor, not the size of the memory referenced by the slice.
The return value of Sizeof is a Go constant.

Sizeof 接收任何类型的表达式 x 并返回假想变量v(v是通过var v = x 声明的)所需的(内存大小的)字节数. 该大小不包含 x 可能引用的任何内存. 例如, x 是 slice, Sizeof仅返回slice描述符的大小，而不是该切片所引用的内存的大小. Sizeof的返回值是Go的常量.



## <a id="ArbitraryType">type</a> [ArbitraryType](https://golang.org/src/unsafe/unsafe.go?s=547:569#L5)
ArbitraryType is here for the purposes of documentation only and is not actually
part of the unsafe package. It represents the type of an arbitrary Go expression.

ArbitraryType 在此仅出于文档的说明目的, 实际上并不是unsafe的一部分. 它是任意Go类型的表达式.

<pre>type ArbitraryType <a href="/pkg/builtin/#int">int</a></pre>











## <a id="Pointer">type</a> [Pointer](https://golang.org/src/unsafe/unsafe.go?s=7766:7793#L170)
Pointer represents a pointer to an arbitrary type. There are four special operations
available for type Pointer that are not available for other types:


	- A pointer value of any type can be converted to a Pointer.
	- A Pointer can be converted to a pointer value of any type.
	- A uintptr can be converted to a Pointer.
	- A Pointer can be converted to a uintptr.

Pointer therefore allows a program to defeat the type system and read and write
arbitrary memory. It should be used with extreme care.

The following patterns involving Pointer are valid.
Code not using these patterns is likely to be invalid today
or to become invalid in the future.
Even the valid patterns below come with important caveats.

Running "go vet" can help find uses of Pointer that do not conform to these patterns,
but silence from "go vet" is not a guarantee that the code is valid.

(1) Conversion of a *T1 to Pointer to *T2.

Provided that T2 is no larger than T1 and that the two share an equivalent
memory layout, this conversion allows reinterpreting data of one type as
data of another type. An example is the implementation of
math.Float64bits:


	func Float64bits(f float64) uint64 {
		return *(*uint64)(unsafe.Pointer(&f))
	}

(2) Conversion of a Pointer to a uintptr (but not back to Pointer).

Converting a Pointer to a uintptr produces the memory address of the value
pointed at, as an integer. The usual use for such a uintptr is to print it.

Conversion of a uintptr back to Pointer is not valid in general.

A uintptr is an integer, not a reference.
Converting a Pointer to a uintptr creates an integer value
with no pointer semantics.
Even if a uintptr holds the address of some object,
the garbage collector will not update that uintptr's value
if the object moves, nor will that uintptr keep the object
from being reclaimed.

The remaining patterns enumerate the only valid conversions
from uintptr to Pointer.

(3) Conversion of a Pointer to a uintptr and back, with arithmetic.

If p points into an allocated object, it can be advanced through the object
by conversion to uintptr, addition of an offset, and conversion back to Pointer.


	p = unsafe.Pointer(uintptr(p) + offset)

The most common use of this pattern is to access fields in a struct
or elements of an array:


	// equivalent to f := unsafe.Pointer(&s.f)
	f := unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f))
	
	// equivalent to e := unsafe.Pointer(&x[i])
	e := unsafe.Pointer(uintptr(unsafe.Pointer(&x[0])) + i*unsafe.Sizeof(x[0]))

It is valid both to add and to subtract offsets from a pointer in this way.
It is also valid to use &^ to round pointers, usually for alignment.
In all cases, the result must continue to point into the original allocated object.

Unlike in C, it is not valid to advance a pointer just beyond the end of
its original allocation:


	// INVALID: end points outside allocated space.
	var s thing
	end = unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(s))
	
	// INVALID: end points outside allocated space.
	b := make([]byte, n)
	end = unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(n))

Note that both conversions must appear in the same expression, with only
the intervening arithmetic between them:


	// INVALID: uintptr cannot be stored in variable
	// before conversion back to Pointer.
	u := uintptr(p)
	p = unsafe.Pointer(u + offset)

Note that the pointer must point into an allocated object, so it may not be nil.


	// INVALID: conversion of nil pointer
	u := unsafe.Pointer(nil)
	p := unsafe.Pointer(uintptr(u) + offset)

(4) Conversion of a Pointer to a uintptr when calling syscall.Syscall.

The Syscall functions in package syscall pass their uintptr arguments directly
to the operating system, which then may, depending on the details of the call,
reinterpret some of them as pointers.
That is, the system call implementation is implicitly converting certain arguments
back from uintptr to pointer.

If a pointer argument must be converted to uintptr for use as an argument,
that conversion must appear in the call expression itself:


	syscall.Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(n))

The compiler handles a Pointer converted to a uintptr in the argument list of
a call to a function implemented in assembly by arranging that the referenced
allocated object, if any, is retained and not moved until the call completes,
even though from the types alone it would appear that the object is no longer
needed during the call.

For the compiler to recognize this pattern,
the conversion must appear in the argument list:


	// INVALID: uintptr cannot be stored in variable
	// before implicit conversion back to Pointer during system call.
	u := uintptr(unsafe.Pointer(p))
	syscall.Syscall(SYS_READ, uintptr(fd), u, uintptr(n))

(5) Conversion of the result of reflect.Value.Pointer or reflect.Value.UnsafeAddr
from uintptr to Pointer.

Package reflect's Value methods named Pointer and UnsafeAddr return type uintptr
instead of unsafe.Pointer to keep callers from changing the result to an arbitrary
type without first importing "unsafe". However, this means that the result is
fragile and must be converted to Pointer immediately after making the call,
in the same expression:


	p := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer()))

As in the cases above, it is invalid to store the result before the conversion:


	// INVALID: uintptr cannot be stored in variable
	// before conversion back to Pointer.
	u := reflect.ValueOf(new(int)).Pointer()
	p := (*int)(unsafe.Pointer(u))

(6) Conversion of a reflect.SliceHeader or reflect.StringHeader Data field to or from Pointer.

As in the previous case, the reflect data structures SliceHeader and StringHeader
declare the field Data as a uintptr to keep callers from changing the result to
an arbitrary type without first importing "unsafe". However, this means that
SliceHeader and StringHeader are only valid when interpreting the content
of an actual slice or string value.


	var s string
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // case 1
	hdr.Data = uintptr(unsafe.Pointer(p))              // case 6 (this case)
	hdr.Len = n

In this usage hdr.Data is really an alternate way to refer to the underlying
pointer in the string header, not a uintptr variable itself.

In general, reflect.SliceHeader and reflect.StringHeader should be used
only as *reflect.SliceHeader and *reflect.StringHeader pointing at actual
slices or strings, never as plain structs.
A program should not declare or allocate variables of these struct types.


	// INVALID: a directly-declared header will not hold Data as a reference.
	var hdr reflect.StringHeader
	hdr.Data = uintptr(unsafe.Pointer(p))
	hdr.Len = n
	s := *(*string)(unsafe.Pointer(&hdr)) // p possibly already lost


Pointer 表示 指向任意类型的指针. 指针类型有四种特殊操作，而其他类型则没有:

	- 任意类型的指针可以转换为一个Pointer类型值
	- 一个Pointer类型值可以转换为任意类型的指针
	- 一个uintptr类型值可以转换为一个Pointer类型值
	- 一个Pointer类型值可以转换为一个uintptr类型值

因此，指针允许程序绕过类型系统并读写任意内存. 使用时应格外小心.

以下涉及Pointer的模式是有效的. 不使用这些模式的代码在现今可能是无效的，或者将来可能是无效的. 甚至下面的有效模式也带有重要的警告.

运行"go vet"可以帮助查找不符合这些模式的Pointer用法，但是"go vet"未输出时不能保证该代码是有效的.

(1) 将 *T1 转成 Pointer 再 转成 *T2

前提是T2不大于T1, 且它们共享相同的内存布局, 这种转换允许将一种类型的数据重新解释为
其他类型的数据. math.Float64bits的实现即是例子:

	func Float64bits(f float64) uint64 {
		return *(*uint64)(unsafe.Pointer(&f))
	}

(2) 将Pointer转成uintptr(但不转回Pointer)

将Pointer转成uintptr会得到Pointer所指向的值的内存地址, 该地址是一个整数. 这种uintptr的通常用法是打印它.

通常, 将uintptr转换回Pointer是无效的.

uintptr是整数, 而不是引用. 将Pointer转成uintptr会创建一个整数值且没有指针语义. 即使该uintptr保存了某个对象的地址, `如果该对象被移动了, gc也不会更新该uintptr的值`; 同样`uintptr也不能保障该对象不被回收`. 

其余的模式列举了唯一有效的从uintptr到Pointer的转换.

(3) 将Pointer转成uintprt(算术操作后)再转回来

如果p指向已分配了空间的对象, 则可以通过先转成uintptr, 再添加偏移量, 再转回Pointer使得(指针的)指向前移.

	p = unsafe.Pointer(uintptr(p) + offset)

此模式最常见的用法是访问结构中的字段或数组的元素:


	// equivalent to f := unsafe.Pointer(&s.f)
	f := unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f))
	
	// equivalent to e := unsafe.Pointer(&x[i])
	e := unsafe.Pointer(uintptr(unsafe.Pointer(&x[0])) + i*unsafe.Sizeof(x[0]))

It is valid both to add and to subtract offsets from a pointer in this way.
It is also valid to use &^ to round pointers, usually for alignment.
In all cases, the result must continue to point into the original allocated object.

以这种方式从指针添加和减去偏移量都是有效的. 通常使用&^来四舍五入指针???（通常用于对齐）也是有效的. 在所有情况下，结果都必须继续指向原始分配的对象.

与C语言不同，将指针前移到原始分配空间末尾的后面是无效的:


	// INVALID: end points outside allocated space.
	var s thing
	end = unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(s))
	
	// INVALID: end points outside allocated space.
	b := make([]byte, n)
	end = unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(n))

注意, 这两个转换必须出现在相同的表达式中，它们之间仅有算术运算:


	// INVALID: uintptr cannot be stored in variable
	// before conversion back to Pointer.
	u := uintptr(p)
	p = unsafe.Pointer(u + offset) // 要先运行再转回去

注意, 指针必须指向已分配空间的对象，因此它不可能为nil.

	// INVALID: conversion of nil pointer
	u := unsafe.Pointer(nil)
	p := unsafe.Pointer(uintptr(u) + offset)

(4) 调用syscall.Syscall时将指针转换为uintptr

syscall包中的Syscall函数需要直接传uintptr给操作系统, 然后根据调用的详细信息，
重新将其中一些解释为指针. 也就是说，系统调用隐式地实现了将某些uintptr转换回指针的操作.

如果必须将指针参数转换为uintptr用作参数，那么该转换必须出现在调用表达式本身中：

	syscall.Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(n))

The compiler handles a Pointer converted to a uintptr in the argument list of
a call to a function implemented in assembly by arranging that the referenced
allocated object, if any, is retained and not moved until the call completes,
even though from the types alone it would appear that the object is no longer
needed during the call.

编译器处理call的参数列表并将Pointer转为 uintptr 是通过汇编实现的???, 它会设置这些已分配对象的引用, 所有引用会保留且不会移动直到调用完成，即使仅从类型来看, 调用期间这些对象不再被需要. 



为了使编译器能够识别这种模式，转换必须出现在参数列表中:


	// INVALID: uintptr cannot be stored in variable
	// before implicit conversion back to Pointer during system call.
	u := uintptr(unsafe.Pointer(p))
	syscall.Syscall(SYS_READ, uintptr(fd), u, uintptr(n))

(5) 将 reflect.Value.Pointer 或 reflect.Value.UnsafeAddr 的结果从 uintptr 转换为Pointer

reflect包的Value的Pointer和UnsafeAddr方法返回类型是uintptr而不是nsafe.Pointer, 以防止调用者在未事先导入"unsafe"包的情况下将结果更改为任意类型. 但是，这意味着结果很
脆弱，必须在调用后立即转换为Pointer, 并使用相同的表达式：


	p := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer()))

与上述例子一样，在转换之前存储结果是无效的:


	// INVALID: uintptr cannot be stored in variable
	// before conversion back to Pointer.
	u := reflect.ValueOf(new(int)).Pointer()
	p := (*int)(unsafe.Pointer(u))

(6) 将 Reflection.SliceHeader 或 reflect.StringHeader 数据字段转换为Pointer或从Pointer转换而来

与前面的例子一样，反射 SliceHeader 和 StringHeader 的数据结构将字段 Data 声明为 uintptr，以防止调用者在不事先导入 "unsafe" 包的情况下将结果更改为任意类型. 但是，这意味着 SliceHeader 和 StringHeader 仅在解释slice或字符串的实际内容时有效.


	var s string
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // case 1
	hdr.Data = uintptr(unsafe.Pointer(p))              // case 6 (this case)
	hdr.Len = n

在这种用法中，hdr.Data实际上是另一种引用string header中的底层指针的方式, 而不是uintptr变量本身???.

通常，reflect.SliceHeader和reflect.StringHeader 仅可用作* reflect.SliceHeader和* reflect.StringHeader指向实际的slicen或字符串, 切勿直接作为struct的类型. 程序不应声明或分配这些结构类型的变量.

	// INVALID: a directly-declared header will not hold Data as a reference.
	var hdr reflect.StringHeader
	hdr.Data = uintptr(unsafe.Pointer(p))
	hdr.Len = n
	s := *(*string)(unsafe.Pointer(&hdr)) // p possibly already lost



<pre>type Pointer *<a href="#ArbitraryType">ArbitraryType</a></pre>















