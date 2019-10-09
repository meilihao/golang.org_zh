

# atomic
`import "sync/atomic"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package atomic provides low-level atomic memory primitives
useful for implementing synchronization algorithms.

These functions require great care to be used correctly.
Except for special, low-level applications, synchronization is better
done with channels or the facilities of the sync package.
Share memory by communicating;
don't communicate by sharing memory.

The swap operation, implemented by the SwapT functions, is the atomic
equivalent of:


	old = *addr
	*addr = new
	return old

The compare-and-swap operation, implemented by the CompareAndSwapT
functions, is the atomic equivalent of:


	if *addr == old {
		*addr = new
		return true
	}
	return false

The add operation, implemented by the AddT functions, is the atomic
equivalent of:


	*addr += delta
	return *addr

The load and store operations, implemented by the LoadT and StoreT
functions, are the atomic equivalents of "return *addr" and
"*addr = val".




## <a id="pkg-index">Index</a>
* [func AddInt32(addr *int32, delta int32) (new int32)](#AddInt32)
* [func AddInt64(addr *int64, delta int64) (new int64)](#AddInt64)
* [func AddUint32(addr *uint32, delta uint32) (new uint32)](#AddUint32)
* [func AddUint64(addr *uint64, delta uint64) (new uint64)](#AddUint64)
* [func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)](#AddUintptr)
* [func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)](#CompareAndSwapInt32)
* [func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)](#CompareAndSwapInt64)
* [func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)](#CompareAndSwapPointer)
* [func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)](#CompareAndSwapUint32)
* [func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)](#CompareAndSwapUint64)
* [func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)](#CompareAndSwapUintptr)
* [func LoadInt32(addr *int32) (val int32)](#LoadInt32)
* [func LoadInt64(addr *int64) (val int64)](#LoadInt64)
* [func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)](#LoadPointer)
* [func LoadUint32(addr *uint32) (val uint32)](#LoadUint32)
* [func LoadUint64(addr *uint64) (val uint64)](#LoadUint64)
* [func LoadUintptr(addr *uintptr) (val uintptr)](#LoadUintptr)
* [func StoreInt32(addr *int32, val int32)](#StoreInt32)
* [func StoreInt64(addr *int64, val int64)](#StoreInt64)
* [func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)](#StorePointer)
* [func StoreUint32(addr *uint32, val uint32)](#StoreUint32)
* [func StoreUint64(addr *uint64, val uint64)](#StoreUint64)
* [func StoreUintptr(addr *uintptr, val uintptr)](#StoreUintptr)
* [func SwapInt32(addr *int32, new int32) (old int32)](#SwapInt32)
* [func SwapInt64(addr *int64, new int64) (old int64)](#SwapInt64)
* [func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)](#SwapPointer)
* [func SwapUint32(addr *uint32, new uint32) (old uint32)](#SwapUint32)
* [func SwapUint64(addr *uint64, new uint64) (old uint64)](#SwapUint64)
* [func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)](#SwapUintptr)
* [type Value](#Value)
  * [func (v *Value) Load() (x interface{})](#Value.Load)
  * [func (v *Value) Store(x interface{})](#Value.Store)


#### <a id="pkg-examples">Examples</a>
* [Value (Config)](#example_Value_config)
* [Value (ReadMostly)](#example_Value_readMostly)


#### <a id="pkg-files">Package files</a>
[doc.go](https://golang.org/src/sync/atomic/doc.go) [value.go](https://golang.org/src/sync/atomic/value.go) 






## <a id="AddInt32">func</a> [AddInt32](https://golang.org/src/sync/atomic/doc.go?s=3572:3623#L83)
<pre>func AddInt32(addr *<a href="/pkg/builtin/#int32">int32</a>, delta <a href="/pkg/builtin/#int32">int32</a>) (new <a href="/pkg/builtin/#int32">int32</a>)</pre>
AddInt32 atomically adds delta to *addr and returns the new value.



## <a id="AddInt64">func</a> [AddInt64](https://golang.org/src/sync/atomic/doc.go?s=3977:4028#L91)
<pre>func AddInt64(addr *<a href="/pkg/builtin/#int64">int64</a>, delta <a href="/pkg/builtin/#int64">int64</a>) (new <a href="/pkg/builtin/#int64">int64</a>)</pre>
AddInt64 atomically adds delta to *addr and returns the new value.



## <a id="AddUint32">func</a> [AddUint32](https://golang.org/src/sync/atomic/doc.go?s=3850:3905#L88)
<pre>func AddUint32(addr *<a href="/pkg/builtin/#uint32">uint32</a>, delta <a href="/pkg/builtin/#uint32">uint32</a>) (new <a href="/pkg/builtin/#uint32">uint32</a>)</pre>
AddUint32 atomically adds delta to *addr and returns the new value.
To subtract a signed positive constant value c from x, do AddUint32(&x, ^uint32(c-1)).
In particular, to decrement x, do AddUint32(&x, ^uint32(0)).



## <a id="AddUint64">func</a> [AddUint64](https://golang.org/src/sync/atomic/doc.go?s=4255:4310#L96)
<pre>func AddUint64(addr *<a href="/pkg/builtin/#uint64">uint64</a>, delta <a href="/pkg/builtin/#uint64">uint64</a>) (new <a href="/pkg/builtin/#uint64">uint64</a>)</pre>
AddUint64 atomically adds delta to *addr and returns the new value.
To subtract a signed positive constant value c from x, do AddUint64(&x, ^uint64(c-1)).
In particular, to decrement x, do AddUint64(&x, ^uint64(0)).



## <a id="AddUintptr">func</a> [AddUintptr](https://golang.org/src/sync/atomic/doc.go?s=4384:4443#L99)
<pre>func AddUintptr(addr *<a href="/pkg/builtin/#uintptr">uintptr</a>, delta <a href="/pkg/builtin/#uintptr">uintptr</a>) (new <a href="/pkg/builtin/#uintptr">uintptr</a>)</pre>
AddUintptr atomically adds delta to *addr and returns the new value.



## <a id="CompareAndSwapInt32">func</a> [CompareAndSwapInt32](https://golang.org/src/sync/atomic/doc.go?s=2620:2688#L65)
<pre>func CompareAndSwapInt32(addr *<a href="/pkg/builtin/#int32">int32</a>, old, new <a href="/pkg/builtin/#int32">int32</a>) (swapped <a href="/pkg/builtin/#bool">bool</a>)</pre>
CompareAndSwapInt32 executes the compare-and-swap operation for an int32 value.



## <a id="CompareAndSwapInt64">func</a> [CompareAndSwapInt64](https://golang.org/src/sync/atomic/doc.go?s=2773:2841#L68)
<pre>func CompareAndSwapInt64(addr *<a href="/pkg/builtin/#int64">int64</a>, old, new <a href="/pkg/builtin/#int64">int64</a>) (swapped <a href="/pkg/builtin/#bool">bool</a>)</pre>
CompareAndSwapInt64 executes the compare-and-swap operation for an int64 value.



## <a id="CompareAndSwapPointer">func</a> [CompareAndSwapPointer](https://golang.org/src/sync/atomic/doc.go?s=3412:3500#L80)
<pre>func CompareAndSwapPointer(addr *<a href="/pkg/unsafe/">unsafe</a>.<a href="/pkg/unsafe/#Pointer">Pointer</a>, old, new <a href="/pkg/unsafe/">unsafe</a>.<a href="/pkg/unsafe/#Pointer">Pointer</a>) (swapped <a href="/pkg/builtin/#bool">bool</a>)</pre>
CompareAndSwapPointer executes the compare-and-swap operation for a unsafe.Pointer value.



## <a id="CompareAndSwapUint32">func</a> [CompareAndSwapUint32](https://golang.org/src/sync/atomic/doc.go?s=2927:2998#L71)
<pre>func CompareAndSwapUint32(addr *<a href="/pkg/builtin/#uint32">uint32</a>, old, new <a href="/pkg/builtin/#uint32">uint32</a>) (swapped <a href="/pkg/builtin/#bool">bool</a>)</pre>
CompareAndSwapUint32 executes the compare-and-swap operation for a uint32 value.



## <a id="CompareAndSwapUint64">func</a> [CompareAndSwapUint64](https://golang.org/src/sync/atomic/doc.go?s=3084:3155#L74)
<pre>func CompareAndSwapUint64(addr *<a href="/pkg/builtin/#uint64">uint64</a>, old, new <a href="/pkg/builtin/#uint64">uint64</a>) (swapped <a href="/pkg/builtin/#bool">bool</a>)</pre>
CompareAndSwapUint64 executes the compare-and-swap operation for a uint64 value.



## <a id="CompareAndSwapUintptr">func</a> [CompareAndSwapUintptr](https://golang.org/src/sync/atomic/doc.go?s=3243:3317#L77)
<pre>func CompareAndSwapUintptr(addr *<a href="/pkg/builtin/#uintptr">uintptr</a>, old, new <a href="/pkg/builtin/#uintptr">uintptr</a>) (swapped <a href="/pkg/builtin/#bool">bool</a>)</pre>
CompareAndSwapUintptr executes the compare-and-swap operation for a uintptr value.



## <a id="LoadInt32">func</a> [LoadInt32](https://golang.org/src/sync/atomic/doc.go?s=4482:4521#L102)
<pre>func LoadInt32(addr *<a href="/pkg/builtin/#int32">int32</a>) (val <a href="/pkg/builtin/#int32">int32</a>)</pre>
LoadInt32 atomically loads *addr.



## <a id="LoadInt64">func</a> [LoadInt64](https://golang.org/src/sync/atomic/doc.go?s=4560:4599#L105)
<pre>func LoadInt64(addr *<a href="/pkg/builtin/#int64">int64</a>) (val <a href="/pkg/builtin/#int64">int64</a>)</pre>
LoadInt64 atomically loads *addr.



## <a id="LoadPointer">func</a> [LoadPointer](https://golang.org/src/sync/atomic/doc.go?s=4890:4949#L117)
<pre>func LoadPointer(addr *<a href="/pkg/unsafe/">unsafe</a>.<a href="/pkg/unsafe/#Pointer">Pointer</a>) (val <a href="/pkg/unsafe/">unsafe</a>.<a href="/pkg/unsafe/#Pointer">Pointer</a>)</pre>
LoadPointer atomically loads *addr.



## <a id="LoadUint32">func</a> [LoadUint32](https://golang.org/src/sync/atomic/doc.go?s=4639:4681#L108)
<pre>func LoadUint32(addr *<a href="/pkg/builtin/#uint32">uint32</a>) (val <a href="/pkg/builtin/#uint32">uint32</a>)</pre>
LoadUint32 atomically loads *addr.



## <a id="LoadUint64">func</a> [LoadUint64](https://golang.org/src/sync/atomic/doc.go?s=4721:4763#L111)
<pre>func LoadUint64(addr *<a href="/pkg/builtin/#uint64">uint64</a>) (val <a href="/pkg/builtin/#uint64">uint64</a>)</pre>
LoadUint64 atomically loads *addr.



## <a id="LoadUintptr">func</a> [LoadUintptr](https://golang.org/src/sync/atomic/doc.go?s=4804:4849#L114)
<pre>func LoadUintptr(addr *<a href="/pkg/builtin/#uintptr">uintptr</a>) (val <a href="/pkg/builtin/#uintptr">uintptr</a>)</pre>
LoadUintptr atomically loads *addr.



## <a id="StoreInt32">func</a> [StoreInt32](https://golang.org/src/sync/atomic/doc.go?s=4999:5038#L120)
<pre>func StoreInt32(addr *<a href="/pkg/builtin/#int32">int32</a>, val <a href="/pkg/builtin/#int32">int32</a>)</pre>
StoreInt32 atomically stores val into *addr.



## <a id="StoreInt64">func</a> [StoreInt64](https://golang.org/src/sync/atomic/doc.go?s=5088:5127#L123)
<pre>func StoreInt64(addr *<a href="/pkg/builtin/#int64">int64</a>, val <a href="/pkg/builtin/#int64">int64</a>)</pre>
StoreInt64 atomically stores val into *addr.



## <a id="StorePointer">func</a> [StorePointer](https://golang.org/src/sync/atomic/doc.go?s=5462:5521#L135)
<pre>func StorePointer(addr *<a href="/pkg/unsafe/">unsafe</a>.<a href="/pkg/unsafe/#Pointer">Pointer</a>, val <a href="/pkg/unsafe/">unsafe</a>.<a href="/pkg/unsafe/#Pointer">Pointer</a>)</pre>
StorePointer atomically stores val into *addr.



## <a id="StoreUint32">func</a> [StoreUint32](https://golang.org/src/sync/atomic/doc.go?s=5178:5220#L126)
<pre>func StoreUint32(addr *<a href="/pkg/builtin/#uint32">uint32</a>, val <a href="/pkg/builtin/#uint32">uint32</a>)</pre>
StoreUint32 atomically stores val into *addr.



## <a id="StoreUint64">func</a> [StoreUint64](https://golang.org/src/sync/atomic/doc.go?s=5271:5313#L129)
<pre>func StoreUint64(addr *<a href="/pkg/builtin/#uint64">uint64</a>, val <a href="/pkg/builtin/#uint64">uint64</a>)</pre>
StoreUint64 atomically stores val into *addr.



## <a id="StoreUintptr">func</a> [StoreUintptr](https://golang.org/src/sync/atomic/doc.go?s=5365:5410#L132)
<pre>func StoreUintptr(addr *<a href="/pkg/builtin/#uintptr">uintptr</a>, val <a href="/pkg/builtin/#uintptr">uintptr</a>)</pre>
StoreUintptr atomically stores val into *addr.



## <a id="SwapInt32">func</a> [SwapInt32](https://golang.org/src/sync/atomic/doc.go?s=1754:1804#L47)
<pre>func SwapInt32(addr *<a href="/pkg/builtin/#int32">int32</a>, new <a href="/pkg/builtin/#int32">int32</a>) (old <a href="/pkg/builtin/#int32">int32</a>)</pre>
SwapInt32 atomically stores new into *addr and returns the previous *addr value.



## <a id="SwapInt64">func</a> [SwapInt64](https://golang.org/src/sync/atomic/doc.go?s=1890:1940#L50)
<pre>func SwapInt64(addr *<a href="/pkg/builtin/#int64">int64</a>, new <a href="/pkg/builtin/#int64">int64</a>) (old <a href="/pkg/builtin/#int64">int64</a>)</pre>
SwapInt64 atomically stores new into *addr and returns the previous *addr value.



## <a id="SwapPointer">func</a> [SwapPointer](https://golang.org/src/sync/atomic/doc.go?s=2456:2535#L62)
<pre>func SwapPointer(addr *<a href="/pkg/unsafe/">unsafe</a>.<a href="/pkg/unsafe/#Pointer">Pointer</a>, new <a href="/pkg/unsafe/">unsafe</a>.<a href="/pkg/unsafe/#Pointer">Pointer</a>) (old <a href="/pkg/unsafe/">unsafe</a>.<a href="/pkg/unsafe/#Pointer">Pointer</a>)</pre>
SwapPointer atomically stores new into *addr and returns the previous *addr value.



## <a id="SwapUint32">func</a> [SwapUint32](https://golang.org/src/sync/atomic/doc.go?s=2027:2081#L53)
<pre>func SwapUint32(addr *<a href="/pkg/builtin/#uint32">uint32</a>, new <a href="/pkg/builtin/#uint32">uint32</a>) (old <a href="/pkg/builtin/#uint32">uint32</a>)</pre>
SwapUint32 atomically stores new into *addr and returns the previous *addr value.



## <a id="SwapUint64">func</a> [SwapUint64](https://golang.org/src/sync/atomic/doc.go?s=2168:2222#L56)
<pre>func SwapUint64(addr *<a href="/pkg/builtin/#uint64">uint64</a>, new <a href="/pkg/builtin/#uint64">uint64</a>) (old <a href="/pkg/builtin/#uint64">uint64</a>)</pre>
SwapUint64 atomically stores new into *addr and returns the previous *addr value.



## <a id="SwapUintptr">func</a> [SwapUintptr](https://golang.org/src/sync/atomic/doc.go?s=2310:2368#L59)
<pre>func SwapUintptr(addr *<a href="/pkg/builtin/#uintptr">uintptr</a>, new <a href="/pkg/builtin/#uintptr">uintptr</a>) (old <a href="/pkg/builtin/#uintptr">uintptr</a>)</pre>
SwapUintptr atomically stores new into *addr and returns the previous *addr value.





## <a id="Value">type</a> [Value](https://golang.org/src/sync/atomic/value.go?s=436:472#L6)
A Value provides an atomic load and store of a consistently typed value.
The zero value for a Value returns nil from Load.
Once Store has been called, a Value must not be copied.

A Value must not be copied after first use.


<pre>type Value struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Value_config">Example (Config)</a>
<p>The following example shows how to use Value for periodic program config updates
and propagation of the changes to worker goroutines.
</p><a id="example_Value_readMostly">Example (ReadMostly)</a>
<p>The following example shows how to maintain a scalable frequently read,
but infrequently updated data structure using copy-on-write idiom.
</p>





### <a id="Value.Load">func</a> (\*Value) [Load](https://golang.org/src/sync/atomic/value.go?s=723:761#L18)
<pre>func (v *<a href="#Value">Value</a>) Load() (x interface{})</pre>
Load returns the value set by the most recent Store.
It returns nil if there has been no call to Store for this Value.




### <a id="Value.Store">func</a> (\*Value) [Store](https://golang.org/src/sync/atomic/value.go?s=1233:1269#L35)
<pre>func (v *<a href="#Value">Value</a>) Store(x interface{})</pre>
Store sets the value of the Value to x.
All calls to Store for a given Value must use values of the same concrete type.
Store of an inconsistent type panics, as does Store(nil).








