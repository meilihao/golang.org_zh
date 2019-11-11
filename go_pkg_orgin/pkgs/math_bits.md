

# bits
`import "math/bits"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package bits implements bit counting and manipulation
functions for the predeclared unsigned integer types.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func Add(x, y, carry uint) (sum, carryOut uint)](#Add)
* [func Add32(x, y, carry uint32) (sum, carryOut uint32)](#Add32)
* [func Add64(x, y, carry uint64) (sum, carryOut uint64)](#Add64)
* [func Div(hi, lo, y uint) (quo, rem uint)](#Div)
* [func Div32(hi, lo, y uint32) (quo, rem uint32)](#Div32)
* [func Div64(hi, lo, y uint64) (quo, rem uint64)](#Div64)
* [func LeadingZeros(x uint) int](#LeadingZeros)
* [func LeadingZeros16(x uint16) int](#LeadingZeros16)
* [func LeadingZeros32(x uint32) int](#LeadingZeros32)
* [func LeadingZeros64(x uint64) int](#LeadingZeros64)
* [func LeadingZeros8(x uint8) int](#LeadingZeros8)
* [func Len(x uint) int](#Len)
* [func Len16(x uint16) (n int)](#Len16)
* [func Len32(x uint32) (n int)](#Len32)
* [func Len64(x uint64) (n int)](#Len64)
* [func Len8(x uint8) int](#Len8)
* [func Mul(x, y uint) (hi, lo uint)](#Mul)
* [func Mul32(x, y uint32) (hi, lo uint32)](#Mul32)
* [func Mul64(x, y uint64) (hi, lo uint64)](#Mul64)
* [func OnesCount(x uint) int](#OnesCount)
* [func OnesCount16(x uint16) int](#OnesCount16)
* [func OnesCount32(x uint32) int](#OnesCount32)
* [func OnesCount64(x uint64) int](#OnesCount64)
* [func OnesCount8(x uint8) int](#OnesCount8)
* [func Reverse(x uint) uint](#Reverse)
* [func Reverse16(x uint16) uint16](#Reverse16)
* [func Reverse32(x uint32) uint32](#Reverse32)
* [func Reverse64(x uint64) uint64](#Reverse64)
* [func Reverse8(x uint8) uint8](#Reverse8)
* [func ReverseBytes(x uint) uint](#ReverseBytes)
* [func ReverseBytes16(x uint16) uint16](#ReverseBytes16)
* [func ReverseBytes32(x uint32) uint32](#ReverseBytes32)
* [func ReverseBytes64(x uint64) uint64](#ReverseBytes64)
* [func RotateLeft(x uint, k int) uint](#RotateLeft)
* [func RotateLeft16(x uint16, k int) uint16](#RotateLeft16)
* [func RotateLeft32(x uint32, k int) uint32](#RotateLeft32)
* [func RotateLeft64(x uint64, k int) uint64](#RotateLeft64)
* [func RotateLeft8(x uint8, k int) uint8](#RotateLeft8)
* [func Sub(x, y, borrow uint) (diff, borrowOut uint)](#Sub)
* [func Sub32(x, y, borrow uint32) (diff, borrowOut uint32)](#Sub32)
* [func Sub64(x, y, borrow uint64) (diff, borrowOut uint64)](#Sub64)
* [func TrailingZeros(x uint) int](#TrailingZeros)
* [func TrailingZeros16(x uint16) int](#TrailingZeros16)
* [func TrailingZeros32(x uint32) int](#TrailingZeros32)
* [func TrailingZeros64(x uint64) int](#TrailingZeros64)
* [func TrailingZeros8(x uint8) int](#TrailingZeros8)


#### <a id="pkg-examples">Examples</a>
* [LeadingZeros16](#example_LeadingZeros16)
* [LeadingZeros32](#example_LeadingZeros32)
* [LeadingZeros64](#example_LeadingZeros64)
* [LeadingZeros8](#example_LeadingZeros8)
* [Len16](#example_Len16)
* [Len32](#example_Len32)
* [Len64](#example_Len64)
* [Len8](#example_Len8)
* [OnesCount](#example_OnesCount)
* [OnesCount16](#example_OnesCount16)
* [OnesCount32](#example_OnesCount32)
* [OnesCount64](#example_OnesCount64)
* [OnesCount8](#example_OnesCount8)
* [Reverse16](#example_Reverse16)
* [Reverse32](#example_Reverse32)
* [Reverse64](#example_Reverse64)
* [Reverse8](#example_Reverse8)
* [ReverseBytes16](#example_ReverseBytes16)
* [ReverseBytes32](#example_ReverseBytes32)
* [ReverseBytes64](#example_ReverseBytes64)
* [RotateLeft16](#example_RotateLeft16)
* [RotateLeft32](#example_RotateLeft32)
* [RotateLeft64](#example_RotateLeft64)
* [RotateLeft8](#example_RotateLeft8)
* [TrailingZeros16](#example_TrailingZeros16)
* [TrailingZeros32](#example_TrailingZeros32)
* [TrailingZeros64](#example_TrailingZeros64)
* [TrailingZeros8](#example_TrailingZeros8)


#### <a id="pkg-files">Package files</a>
[bits.go](https://golang.org/src/math/bits/bits.go) [bits_errors.go](https://golang.org/src/math/bits/bits_errors.go) [bits_tables.go](https://golang.org/src/math/bits/bits_tables.go) 


## <a id="pkg-constants">Constants</a>
UintSize is the size of a uint in bits.


<pre>const <span id="UintSize">UintSize</span> = uintSize</pre>



## <a id="Add">func</a> [Add](https://golang.org/src/math/bits/bits.go?s=10577:10624#L344)
<pre>func Add(x, y, carry <a href="/pkg/builtin/#uint">uint</a>) (sum, carryOut <a href="/pkg/builtin/#uint">uint</a>)</pre>
Add returns the sum with carry of x, y and carry: sum = x + y + carry.
The carry input must be 0 or 1; otherwise the behavior is undefined.
The carryOut output is guaranteed to be 0 or 1.

This function's execution time does not depend on the inputs.



## <a id="Add32">func</a> [Add32](https://golang.org/src/math/bits/bits.go?s=11093:11146#L358)
<pre>func Add32(x, y, carry <a href="/pkg/builtin/#uint32">uint32</a>) (sum, carryOut <a href="/pkg/builtin/#uint32">uint32</a>)</pre>
Add32 returns the sum with carry of x, y and carry: sum = x + y + carry.
The carry input must be 0 or 1; otherwise the behavior is undefined.
The carryOut output is guaranteed to be 0 or 1.

This function's execution time does not depend on the inputs.



## <a id="Add64">func</a> [Add64](https://golang.org/src/math/bits/bits.go?s=11528:11581#L370)
<pre>func Add64(x, y, carry <a href="/pkg/builtin/#uint64">uint64</a>) (sum, carryOut <a href="/pkg/builtin/#uint64">uint64</a>)</pre>
Add64 returns the sum with carry of x, y and carry: sum = x + y + carry.
The carry input must be 0 or 1; otherwise the behavior is undefined.
The carryOut output is guaranteed to be 0 or 1.

This function's execution time does not depend on the inputs.



## <a id="Div">func</a> [Div](https://golang.org/src/math/bits/bits.go?s=15135:15175#L476)
<pre>func Div(hi, lo, y <a href="/pkg/builtin/#uint">uint</a>) (quo, rem <a href="/pkg/builtin/#uint">uint</a>)</pre>
Div returns the quotient and remainder of (hi, lo) divided by y:
quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits' upper
half in parameter hi and the lower half in parameter lo.
Div panics for y == 0 (division by zero) or y <= hi (quotient overflow).



## <a id="Div32">func</a> [Div32](https://golang.org/src/math/bits/bits.go?s=15633:15679#L489)
<pre>func Div32(hi, lo, y <a href="/pkg/builtin/#uint32">uint32</a>) (quo, rem <a href="/pkg/builtin/#uint32">uint32</a>)</pre>
Div32 returns the quotient and remainder of (hi, lo) divided by y:
quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits' upper
half in parameter hi and the lower half in parameter lo.
Div32 panics for y == 0 (division by zero) or y <= hi (quotient overflow).



## <a id="Div64">func</a> [Div64](https://golang.org/src/math/bits/bits.go?s=16106:16152#L502)
<pre>func Div64(hi, lo, y <a href="/pkg/builtin/#uint64">uint64</a>) (quo, rem <a href="/pkg/builtin/#uint64">uint64</a>)</pre>
Div64 returns the quotient and remainder of (hi, lo) divided by y:
quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits' upper
half in parameter hi and the lower half in parameter lo.
Div64 panics for y == 0 (division by zero) or y <= hi (quotient overflow).



## <a id="LeadingZeros">func</a> [LeadingZeros](https://golang.org/src/math/bits/bits.go?s=574:603#L9)
<pre>func LeadingZeros(x <a href="/pkg/builtin/#uint">uint</a>) <a href="/pkg/builtin/#int">int</a></pre>
LeadingZeros returns the number of leading zero bits in x; the result is UintSize for x == 0.



## <a id="LeadingZeros16">func</a> [LeadingZeros16](https://golang.org/src/math/bits/bits.go?s=874:907#L15)
<pre>func LeadingZeros16(x <a href="/pkg/builtin/#uint16">uint16</a>) <a href="/pkg/builtin/#int">int</a></pre>
LeadingZeros16 returns the number of leading zero bits in x; the result is 16 for x == 0.



<a id="example_LeadingZeros16">Example</a>


```go
```

output:
```txt
```

## <a id="LeadingZeros32">func</a> [LeadingZeros32](https://golang.org/src/math/bits/bits.go?s=1027:1060#L18)
<pre>func LeadingZeros32(x <a href="/pkg/builtin/#uint32">uint32</a>) <a href="/pkg/builtin/#int">int</a></pre>
LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.



<a id="example_LeadingZeros32">Example</a>


```go
```

output:
```txt
```

## <a id="LeadingZeros64">func</a> [LeadingZeros64](https://golang.org/src/math/bits/bits.go?s=1180:1213#L21)
<pre>func LeadingZeros64(x <a href="/pkg/builtin/#uint64">uint64</a>) <a href="/pkg/builtin/#int">int</a></pre>
LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.



<a id="example_LeadingZeros64">Example</a>


```go
```

output:
```txt
```

## <a id="LeadingZeros8">func</a> [LeadingZeros8](https://golang.org/src/math/bits/bits.go?s=725:756#L12)
<pre>func LeadingZeros8(x <a href="/pkg/builtin/#uint8">uint8</a>) <a href="/pkg/builtin/#int">int</a></pre>
LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.



<a id="example_LeadingZeros8">Example</a>


```go
```

output:
```txt
```

## <a id="Len">func</a> [Len](https://golang.org/src/math/bits/bits.go?s=9325:9345#L286)
<pre>func Len(x <a href="/pkg/builtin/#uint">uint</a>) <a href="/pkg/builtin/#int">int</a></pre>
Len returns the minimum number of bits required to represent x; the result is 0 for x == 0.



## <a id="Len16">func</a> [Len16](https://golang.org/src/math/bits/bits.go?s=9671:9699#L299)
<pre>func Len16(x <a href="/pkg/builtin/#uint16">uint16</a>) (n <a href="/pkg/builtin/#int">int</a>)</pre>
Len16 returns the minimum number of bits required to represent x; the result is 0 for x == 0.



<a id="example_Len16">Example</a>


```go
```

output:
```txt
```

## <a id="Len32">func</a> [Len32](https://golang.org/src/math/bits/bits.go?s=9867:9895#L308)
<pre>func Len32(x <a href="/pkg/builtin/#uint32">uint32</a>) (n <a href="/pkg/builtin/#int">int</a>)</pre>
Len32 returns the minimum number of bits required to represent x; the result is 0 for x == 0.



<a id="example_Len32">Example</a>


```go
```

output:
```txt
```

## <a id="Len64">func</a> [Len64](https://golang.org/src/math/bits/bits.go?s=10104:10132#L321)
<pre>func Len64(x <a href="/pkg/builtin/#uint64">uint64</a>) (n <a href="/pkg/builtin/#int">int</a>)</pre>
Len64 returns the minimum number of bits required to represent x; the result is 0 for x == 0.



<a id="example_Len64">Example</a>


```go
```

output:
```txt
```

## <a id="Len8">func</a> [Len8](https://golang.org/src/math/bits/bits.go?s=9522:9544#L294)
<pre>func Len8(x <a href="/pkg/builtin/#uint8">uint8</a>) <a href="/pkg/builtin/#int">int</a></pre>
Len8 returns the minimum number of bits required to represent x; the result is 0 for x == 0.



<a id="example_Len8">Example</a>


```go
```

output:
```txt
```

## <a id="Mul">func</a> [Mul](https://golang.org/src/math/bits/bits.go?s=13808:13841#L429)
<pre>func Mul(x, y <a href="/pkg/builtin/#uint">uint</a>) (hi, lo <a href="/pkg/builtin/#uint">uint</a>)</pre>
Mul returns the full-width product of x and y: (hi, lo) = x * y
with the product bits' upper half returned in hi and the lower
half returned in lo.

This function's execution time does not depend on the inputs.



## <a id="Mul32">func</a> [Mul32](https://golang.org/src/math/bits/bits.go?s=14220:14259#L443)
<pre>func Mul32(x, y <a href="/pkg/builtin/#uint32">uint32</a>) (hi, lo <a href="/pkg/builtin/#uint32">uint32</a>)</pre>
Mul32 returns the 64-bit product of x and y: (hi, lo) = x * y
with the product bits' upper half returned in hi and the lower
half returned in lo.

This function's execution time does not depend on the inputs.



## <a id="Mul64">func</a> [Mul64](https://golang.org/src/math/bits/bits.go?s=14566:14605#L454)
<pre>func Mul64(x, y <a href="/pkg/builtin/#uint64">uint64</a>) (hi, lo <a href="/pkg/builtin/#uint64">uint64</a>)</pre>
Mul64 returns the 128-bit product of x and y: (hi, lo) = x * y
with the product bits' upper half returned in hi and the lower
half returned in lo.

This function's execution time does not depend on the inputs.



## <a id="OnesCount">func</a> [OnesCount](https://golang.org/src/math/bits/bits.go?s=3908:3934#L101)
<pre>func OnesCount(x <a href="/pkg/builtin/#uint">uint</a>) <a href="/pkg/builtin/#int">int</a></pre>
OnesCount returns the number of one bits ("population count") in x.



<a id="example_OnesCount">Example</a>


```go
```

output:
```txt
```

## <a id="OnesCount16">func</a> [OnesCount16](https://golang.org/src/math/bits/bits.go?s=4230:4260#L114)
<pre>func OnesCount16(x <a href="/pkg/builtin/#uint16">uint16</a>) <a href="/pkg/builtin/#int">int</a></pre>
OnesCount16 returns the number of one bits ("population count") in x.



<a id="example_OnesCount16">Example</a>


```go
```

output:
```txt
```

## <a id="OnesCount32">func</a> [OnesCount32](https://golang.org/src/math/bits/bits.go?s=4384:4414#L119)
<pre>func OnesCount32(x <a href="/pkg/builtin/#uint32">uint32</a>) <a href="/pkg/builtin/#int">int</a></pre>
OnesCount32 returns the number of one bits ("population count") in x.



<a id="example_OnesCount32">Example</a>


```go
```

output:
```txt
```

## <a id="OnesCount64">func</a> [OnesCount64](https://golang.org/src/math/bits/bits.go?s=4582:4612#L124)
<pre>func OnesCount64(x <a href="/pkg/builtin/#uint64">uint64</a>) <a href="/pkg/builtin/#int">int</a></pre>
OnesCount64 returns the number of one bits ("population count") in x.



<a id="example_OnesCount64">Example</a>


```go
```

output:
```txt
```

## <a id="OnesCount8">func</a> [OnesCount8](https://golang.org/src/math/bits/bits.go?s=4099:4127#L109)
<pre>func OnesCount8(x <a href="/pkg/builtin/#uint8">uint8</a>) <a href="/pkg/builtin/#int">int</a></pre>
OnesCount8 returns the number of one bits ("population count") in x.



<a id="example_OnesCount8">Example</a>


```go
```

output:
```txt
```

## <a id="Reverse">func</a> [Reverse](https://golang.org/src/math/bits/bits.go?s=7247:7272#L210)
<pre>func Reverse(x <a href="/pkg/builtin/#uint">uint</a>) <a href="/pkg/builtin/#uint">uint</a></pre>
Reverse returns the value of x with its bits in reversed order.



## <a id="Reverse16">func</a> [Reverse16](https://golang.org/src/math/bits/bits.go?s=7563:7594#L223)
<pre>func Reverse16(x <a href="/pkg/builtin/#uint16">uint16</a>) <a href="/pkg/builtin/#uint16">uint16</a></pre>
Reverse16 returns the value of x with its bits in reversed order.



<a id="example_Reverse16">Example</a>


```go
```

output:
```txt
```

## <a id="Reverse32">func</a> [Reverse32](https://golang.org/src/math/bits/bits.go?s=7728:7759#L228)
<pre>func Reverse32(x <a href="/pkg/builtin/#uint32">uint32</a>) <a href="/pkg/builtin/#uint32">uint32</a></pre>
Reverse32 returns the value of x with its bits in reversed order.



<a id="example_Reverse32">Example</a>


```go
```

output:
```txt
```

## <a id="Reverse64">func</a> [Reverse64](https://golang.org/src/math/bits/bits.go?s=7974:8005#L237)
<pre>func Reverse64(x <a href="/pkg/builtin/#uint64">uint64</a>) <a href="/pkg/builtin/#uint64">uint64</a></pre>
Reverse64 returns the value of x with its bits in reversed order.



<a id="example_Reverse64">Example</a>


```go
```

output:
```txt
```

## <a id="Reverse8">func</a> [Reverse8](https://golang.org/src/math/bits/bits.go?s=7441:7469#L218)
<pre>func Reverse8(x <a href="/pkg/builtin/#uint8">uint8</a>) <a href="/pkg/builtin/#uint8">uint8</a></pre>
Reverse8 returns the value of x with its bits in reversed order.



<a id="example_Reverse8">Example</a>


```go
```

output:
```txt
```

## <a id="ReverseBytes">func</a> [ReverseBytes](https://golang.org/src/math/bits/bits.go?s=8317:8347#L250)
<pre>func ReverseBytes(x <a href="/pkg/builtin/#uint">uint</a>) <a href="/pkg/builtin/#uint">uint</a></pre>
ReverseBytes returns the value of x with its bytes in reversed order.

This function's execution time does not depend on the inputs.



## <a id="ReverseBytes16">func</a> [ReverseBytes16](https://golang.org/src/math/bits/bits.go?s=8601:8637#L260)
<pre>func ReverseBytes16(x <a href="/pkg/builtin/#uint16">uint16</a>) <a href="/pkg/builtin/#uint16">uint16</a></pre>
ReverseBytes16 returns the value of x with its bytes in reversed order.

This function's execution time does not depend on the inputs.



<a id="example_ReverseBytes16">Example</a>


```go
```

output:
```txt
```

## <a id="ReverseBytes32">func</a> [ReverseBytes32](https://golang.org/src/math/bits/bits.go?s=8806:8842#L267)
<pre>func ReverseBytes32(x <a href="/pkg/builtin/#uint32">uint32</a>) <a href="/pkg/builtin/#uint32">uint32</a></pre>
ReverseBytes32 returns the value of x with its bytes in reversed order.

This function's execution time does not depend on the inputs.



<a id="example_ReverseBytes32">Example</a>


```go
```

output:
```txt
```

## <a id="ReverseBytes64">func</a> [ReverseBytes64](https://golang.org/src/math/bits/bits.go?s=9065:9101#L276)
<pre>func ReverseBytes64(x <a href="/pkg/builtin/#uint64">uint64</a>) <a href="/pkg/builtin/#uint64">uint64</a></pre>
ReverseBytes64 returns the value of x with its bytes in reversed order.

This function's execution time does not depend on the inputs.



<a id="example_ReverseBytes64">Example</a>


```go
```

output:
```txt
```

## <a id="RotateLeft">func</a> [RotateLeft](https://golang.org/src/math/bits/bits.go?s=5791:5826#L160)
<pre>func RotateLeft(x <a href="/pkg/builtin/#uint">uint</a>, k <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#uint">uint</a></pre>
RotateLeft returns the value of x rotated left by (k mod UintSize) bits.
To rotate x right by k bits, call RotateLeft(x, -k).

This function's execution time does not depend on the inputs.



## <a id="RotateLeft16">func</a> [RotateLeft16](https://golang.org/src/math/bits/bits.go?s=6437:6478#L181)
<pre>func RotateLeft16(x <a href="/pkg/builtin/#uint16">uint16</a>, k <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#uint16">uint16</a></pre>
RotateLeft16 returns the value of x rotated left by (k mod 16) bits.
To rotate x right by k bits, call RotateLeft16(x, -k).

This function's execution time does not depend on the inputs.



<a id="example_RotateLeft16">Example</a>


```go
```

output:
```txt
```

## <a id="RotateLeft32">func</a> [RotateLeft32](https://golang.org/src/math/bits/bits.go?s=6744:6785#L191)
<pre>func RotateLeft32(x <a href="/pkg/builtin/#uint32">uint32</a>, k <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#uint32">uint32</a></pre>
RotateLeft32 returns the value of x rotated left by (k mod 32) bits.
To rotate x right by k bits, call RotateLeft32(x, -k).

This function's execution time does not depend on the inputs.



<a id="example_RotateLeft32">Example</a>


```go
```

output:
```txt
```

## <a id="RotateLeft64">func</a> [RotateLeft64](https://golang.org/src/math/bits/bits.go?s=7051:7092#L201)
<pre>func RotateLeft64(x <a href="/pkg/builtin/#uint64">uint64</a>, k <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#uint64">uint64</a></pre>
RotateLeft64 returns the value of x rotated left by (k mod 64) bits.
To rotate x right by k bits, call RotateLeft64(x, -k).

This function's execution time does not depend on the inputs.



<a id="example_RotateLeft64">Example</a>


```go
```

output:
```txt
```

## <a id="RotateLeft8">func</a> [RotateLeft8](https://golang.org/src/math/bits/bits.go?s=6134:6172#L171)
<pre>func RotateLeft8(x <a href="/pkg/builtin/#uint8">uint8</a>, k <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#uint8">uint8</a></pre>
RotateLeft8 returns the value of x rotated left by (k mod 8) bits.
To rotate x right by k bits, call RotateLeft8(x, -k).

This function's execution time does not depend on the inputs.



<a id="example_RotateLeft8">Example</a>


```go
```

output:
```txt
```

## <a id="Sub">func</a> [Sub](https://golang.org/src/math/bits/bits.go?s=12171:12221#L386)
<pre>func Sub(x, y, borrow <a href="/pkg/builtin/#uint">uint</a>) (diff, borrowOut <a href="/pkg/builtin/#uint">uint</a>)</pre>
Sub returns the difference of x, y and borrow: diff = x - y - borrow.
The borrow input must be 0 or 1; otherwise the behavior is undefined.
The borrowOut output is guaranteed to be 0 or 1.

This function's execution time does not depend on the inputs.



## <a id="Sub32">func</a> [Sub32](https://golang.org/src/math/bits/bits.go?s=12693:12749#L400)
<pre>func Sub32(x, y, borrow <a href="/pkg/builtin/#uint32">uint32</a>) (diff, borrowOut <a href="/pkg/builtin/#uint32">uint32</a>)</pre>
Sub32 returns the difference of x, y and borrow, diff = x - y - borrow.
The borrow input must be 0 or 1; otherwise the behavior is undefined.
The borrowOut output is guaranteed to be 0 or 1.

This function's execution time does not depend on the inputs.



## <a id="Sub64">func</a> [Sub64](https://golang.org/src/math/bits/bits.go?s=13375:13431#L415)
<pre>func Sub64(x, y, borrow <a href="/pkg/builtin/#uint64">uint64</a>) (diff, borrowOut <a href="/pkg/builtin/#uint64">uint64</a>)</pre>
Sub64 returns the difference of x, y and borrow: diff = x - y - borrow.
The borrow input must be 0 or 1; otherwise the behavior is undefined.
The borrowOut output is guaranteed to be 0 or 1.

This function's execution time does not depend on the inputs.



## <a id="TrailingZeros">func</a> [TrailingZeros](https://golang.org/src/math/bits/bits.go?s=1929:1959#L43)
<pre>func TrailingZeros(x <a href="/pkg/builtin/#uint">uint</a>) <a href="/pkg/builtin/#int">int</a></pre>
TrailingZeros returns the number of trailing zero bits in x; the result is UintSize for x == 0.



## <a id="TrailingZeros16">func</a> [TrailingZeros16](https://golang.org/src/math/bits/bits.go?s=2310:2344#L56)
<pre>func TrailingZeros16(x <a href="/pkg/builtin/#uint16">uint16</a>) <a href="/pkg/builtin/#int">int</a></pre>
TrailingZeros16 returns the number of trailing zero bits in x; the result is 16 for x == 0.



<a id="example_TrailingZeros16">Example</a>


```go
```

output:
```txt
```

## <a id="TrailingZeros32">func</a> [TrailingZeros32](https://golang.org/src/math/bits/bits.go?s=2568:2602#L65)
<pre>func TrailingZeros32(x <a href="/pkg/builtin/#uint32">uint32</a>) <a href="/pkg/builtin/#int">int</a></pre>
TrailingZeros32 returns the number of trailing zero bits in x; the result is 32 for x == 0.



<a id="example_TrailingZeros32">Example</a>


```go
```

output:
```txt
```

## <a id="TrailingZeros64">func</a> [TrailingZeros64](https://golang.org/src/math/bits/bits.go?s=2820:2854#L74)
<pre>func TrailingZeros64(x <a href="/pkg/builtin/#uint64">uint64</a>) <a href="/pkg/builtin/#int">int</a></pre>
TrailingZeros64 returns the number of trailing zero bits in x; the result is 64 for x == 0.



<a id="example_TrailingZeros64">Example</a>


```go
```

output:
```txt
```

## <a id="TrailingZeros8">func</a> [TrailingZeros8](https://golang.org/src/math/bits/bits.go?s=2153:2185#L51)
<pre>func TrailingZeros8(x <a href="/pkg/builtin/#uint8">uint8</a>) <a href="/pkg/builtin/#int">int</a></pre>
TrailingZeros8 returns the number of trailing zero bits in x; the result is 8 for x == 0.



<a id="example_TrailingZeros8">Example</a>


```go
```

output:
```txt
```






