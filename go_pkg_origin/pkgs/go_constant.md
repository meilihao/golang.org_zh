

# constant
`import "go/constant"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package constant implements Values representing untyped
Go constants and their corresponding operations.

A special Unknown value may be used when a value
is unknown due to an error. Operations on unknown
values produce unknown values unless specified
otherwise.



<a id="example__complexNumbers">Example (ComplexNumbers)</a>


```go
```

output:
```txt
```


## <a id="pkg-index">Index</a>
* [func BitLen(x Value) int](#BitLen)
* [func BoolVal(x Value) bool](#BoolVal)
* [func Bytes(x Value) []byte](#Bytes)
* [func Compare(x_ Value, op token.Token, y_ Value) bool](#Compare)
* [func Float32Val(x Value) (float32, bool)](#Float32Val)
* [func Float64Val(x Value) (float64, bool)](#Float64Val)
* [func Int64Val(x Value) (int64, bool)](#Int64Val)
* [func Sign(x Value) int](#Sign)
* [func StringVal(x Value) string](#StringVal)
* [func Uint64Val(x Value) (uint64, bool)](#Uint64Val)
* [func Val(x Value) interface{}](#Val)
* [type Kind](#Kind)
* [type Value](#Value)
  * [func BinaryOp(x_ Value, op token.Token, y_ Value) Value](#BinaryOp)
  * [func Denom(x Value) Value](#Denom)
  * [func Imag(x Value) Value](#Imag)
  * [func Make(x interface{}) Value](#Make)
  * [func MakeBool(b bool) Value](#MakeBool)
  * [func MakeFloat64(x float64) Value](#MakeFloat64)
  * [func MakeFromBytes(bytes []byte) Value](#MakeFromBytes)
  * [func MakeFromLiteral(lit string, tok token.Token, zero uint) Value](#MakeFromLiteral)
  * [func MakeImag(x Value) Value](#MakeImag)
  * [func MakeInt64(x int64) Value](#MakeInt64)
  * [func MakeString(s string) Value](#MakeString)
  * [func MakeUint64(x uint64) Value](#MakeUint64)
  * [func MakeUnknown() Value](#MakeUnknown)
  * [func Num(x Value) Value](#Num)
  * [func Real(x Value) Value](#Real)
  * [func Shift(x Value, op token.Token, s uint) Value](#Shift)
  * [func ToComplex(x Value) Value](#ToComplex)
  * [func ToFloat(x Value) Value](#ToFloat)
  * [func ToInt(x Value) Value](#ToInt)
  * [func UnaryOp(op token.Token, y Value, prec uint) Value](#UnaryOp)


#### <a id="pkg-examples">Examples</a>
* [BinaryOp](#example_BinaryOp)
* [Compare](#example_Compare)
* [Sign](#example_Sign)
* [UnaryOp](#example_UnaryOp)
* [Val](#example_Val)
* [Package (ComplexNumbers)](#example__complexNumbers)


#### <a id="pkg-files">Package files</a>
[value.go](https://golang.org/src/go/constant/value.go) 






## <a id="BitLen">func</a> [BitLen](https://golang.org/src/go/constant/value.go?s=17113:17137#L620)
<pre>func BitLen(x <a href="#Value">Value</a>) <a href="/pkg/builtin/#int">int</a></pre>
BitLen returns the number of bits required to represent
the absolute value x in binary representation; x must be an Int or an Unknown.
If x is Unknown, the result is 0.



## <a id="BoolVal">func</a> [BoolVal](https://golang.org/src/go/constant/value.go?s=12837:12863#L453)
<pre>func BoolVal(x <a href="#Value">Value</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
BoolVal returns the Go boolean value of x, which must be a Bool or an Unknown.
If x is Unknown, the result is false.



## <a id="Bytes">func</a> [Bytes](https://golang.org/src/go/constant/value.go?s=18377:18403#L673)
<pre>func Bytes(x <a href="#Value">Value</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
Bytes returns the bytes for the absolute value of x in little-
endian binary representation; x must be an Int.



## <a id="Compare">func</a> [Compare](https://golang.org/src/go/constant/value.go?s=32329:32382#L1331)
<pre>func Compare(x_ <a href="#Value">Value</a>, op <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Token">Token</a>, y_ <a href="#Value">Value</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Compare returns the result of the comparison x op y.
The comparison must be defined for the operands.
If one of the operands is Unknown, the result is
false.



<a id="example_Compare">Example</a>


```go
```

output:
```txt
```

## <a id="Float32Val">func</a> [Float32Val](https://golang.org/src/go/constant/value.go?s=14345:14385#L510)
<pre>func Float32Val(x <a href="#Value">Value</a>) (<a href="/pkg/builtin/#float32">float32</a>, <a href="/pkg/builtin/#bool">bool</a>)</pre>
Float32Val is like Float64Val but for float32 instead of float64.



## <a id="Float64Val">func</a> [Float64Val](https://golang.org/src/go/constant/value.go?s=15123:15163#L535)
<pre>func Float64Val(x <a href="#Value">Value</a>) (<a href="/pkg/builtin/#float64">float64</a>, <a href="/pkg/builtin/#bool">bool</a>)</pre>
Float64Val returns the nearest Go float64 value of x and whether the result is exact;
x must be numeric or an Unknown, but not Complex. For values too small (too close to 0)
to represent as float64, Float64Val silently underflows to 0. The result sign always
matches the sign of x, even for 0.
If x is Unknown, the result is (0, false).



## <a id="Int64Val">func</a> [Int64Val](https://golang.org/src/go/constant/value.go?s=13529:13565#L480)
<pre>func Int64Val(x <a href="#Value">Value</a>) (<a href="/pkg/builtin/#int64">int64</a>, <a href="/pkg/builtin/#bool">bool</a>)</pre>
Int64Val returns the Go int64 value of x and whether the result is exact;
x must be an Int or an Unknown. If the result is not exact, its value is undefined.
If x is Unknown, the result is (0, false).



## <a id="Sign">func</a> [Sign](https://golang.org/src/go/constant/value.go?s=17549:17571#L636)
<pre>func Sign(x <a href="#Value">Value</a>) <a href="/pkg/builtin/#int">int</a></pre>
Sign returns -1, 0, or 1 depending on whether x < 0, x == 0, or x > 0;
x must be numeric or Unknown. For complex values x, the sign is 0 if x == 0,
otherwise it is != 0. If x is Unknown, the result is 1.



<a id="example_Sign">Example</a>


```go
```

output:
```txt
```

## <a id="StringVal">func</a> [StringVal](https://golang.org/src/go/constant/value.go?s=13135:13165#L466)
<pre>func StringVal(x <a href="#Value">Value</a>) <a href="/pkg/builtin/#string">string</a></pre>
StringVal returns the Go string value of x, which must be a String or an Unknown.
If x is Unknown, the result is "".



## <a id="Uint64Val">func</a> [Uint64Val](https://golang.org/src/go/constant/value.go?s=14019:14057#L496)
<pre>func Uint64Val(x <a href="#Value">Value</a>) (<a href="/pkg/builtin/#uint64">uint64</a>, <a href="/pkg/builtin/#bool">bool</a>)</pre>
Uint64Val returns the Go uint64 value of x and whether the result is exact;
x must be an Int or an Unknown. If the result is not exact, its value is undefined.
If x is Unknown, the result is (0, false).



## <a id="Val">func</a> [Val](https://golang.org/src/go/constant/value.go?s=16037:16066#L567)
<pre>func Val(x <a href="#Value">Value</a>) interface{}</pre>
Val returns the underlying value for a given constant. Since it returns an
interface, it is up to the caller to type assert the result to the expected
type. The possible dynamic return types are:


	x Kind             type of result
	-----------------------------------------
	Bool               bool
	String             string
	Int                int64 or *big.Int
	Float              *big.Float or *big.Rat
	everything else    nil



<a id="example_Val">Example</a>


```go
```

output:
```txt
```



## <a id="Kind">type</a> [Kind](https://golang.org/src/go/constant/value.go?s=621:634#L17)
Kind specifies the kind of value represented by a Value.


<pre>type Kind <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span class="comment">// unknown values</span>
    <span id="Unknown">Unknown</span> <a href="#Kind">Kind</a> = <a href="/pkg/builtin/#iota">iota</a>

    <span class="comment">// non-numeric values</span>
    <span id="Bool">Bool</span>
    <span id="String">String</span>

    <span class="comment">// numeric values</span>
    <span id="Int">Int</span>
    <span id="Float">Float</span>
    <span id="Complex">Complex</span>
)</pre>









## <a id="Value">type</a> [Value](https://golang.org/src/go/constant/value.go?s=816:1397#L34)
A Value represents the value of a Go constant.


<pre>type Value interface {
    <span class="comment">// Kind returns the value kind.</span>
    Kind() <a href="#Kind">Kind</a>

    <span class="comment">// String returns a short, quoted (human-readable) form of the value.</span>
    <span class="comment">// For numeric values, the result may be an approximation;</span>
    <span class="comment">// for String values the result may be a shortened string.</span>
    <span class="comment">// Use ExactString for a string representing a value exactly.</span>
    String() <a href="/pkg/builtin/#string">string</a>

    <span class="comment">// ExactString returns an exact, quoted (human-readable) form of the value.</span>
    <span class="comment">// If the Value is of Kind String, use StringVal to obtain the unquoted string.</span>
    ExactString() <a href="/pkg/builtin/#string">string</a>
    <span class="comment">// contains filtered or unexported methods</span>
}</pre>









### <a id="BinaryOp">func</a> [BinaryOp](https://golang.org/src/go/constant/value.go?s=27815:27870#L1095)
<pre>func BinaryOp(x_ <a href="#Value">Value</a>, op <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Token">Token</a>, y_ <a href="#Value">Value</a>) <a href="#Value">Value</a></pre>
BinaryOp returns the result of the binary expression x op y.
The operation must be defined for the operands. If one of the
operands is Unknown, the result is Unknown.
BinaryOp doesn't handle comparisons or shifts; use Compare
or Shift instead.

To force integer division of Int operands, use op == token.QUO_ASSIGN
instead of token.QUO; the result is guaranteed to be Int in this case.
Division by zero leads to a run-time panic.



<a id="example_BinaryOp">Example</a>


```go
```

output:
```txt
```


### <a id="Denom">func</a> [Denom](https://golang.org/src/go/constant/value.go?s=20200:20225#L759)
<pre>func Denom(x <a href="#Value">Value</a>) <a href="#Value">Value</a></pre>
Denom returns the denominator of x; x must be Int, Float, or Unknown.
If x is Unknown, or if it is too large or small to represent as a
fraction, the result is Unknown. Otherwise the result is an Int >= 1.




### <a id="Imag">func</a> [Imag](https://golang.org/src/go/constant/value.go?s=21348:21372#L807)
<pre>func Imag(x <a href="#Value">Value</a>) <a href="#Value">Value</a></pre>
Imag returns the imaginary part of x, which must be a numeric or unknown value.
If x is Unknown, the result is Unknown.




### <a id="Make">func</a> [Make](https://golang.org/src/go/constant/value.go?s=16625:16655#L598)
<pre>func Make(x interface{}) <a href="#Value">Value</a></pre>
Make returns the Value for x.


	type of x        result Kind
	----------------------------
	bool             Bool
	string           String
	int64            Int
	*big.Int         Int
	*big.Float       Float
	*big.Rat         Float
	anything else    Unknown




### <a id="MakeBool">func</a> [MakeBool](https://golang.org/src/go/constant/value.go?s=9714:9741#L335)
<pre>func MakeBool(b <a href="/pkg/builtin/#bool">bool</a>) <a href="#Value">Value</a></pre>
MakeBool returns the Bool value for b.




### <a id="MakeFloat64">func</a> [MakeFloat64](https://golang.org/src/go/constant/value.go?s=10228:10261#L353)
<pre>func MakeFloat64(x <a href="/pkg/builtin/#float64">float64</a>) <a href="#Value">Value</a></pre>
MakeFloat64 returns the Float value for x.
If x is not finite, the result is an Unknown.




### <a id="MakeFromBytes">func</a> [MakeFromBytes](https://golang.org/src/go/constant/value.go?s=18968:19006#L705)
<pre>func MakeFromBytes(bytes []<a href="/pkg/builtin/#byte">byte</a>) <a href="#Value">Value</a></pre>
MakeFromBytes returns the Int value given the bytes of its little-endian
binary representation. An empty byte slice argument represents 0.




### <a id="MakeFromLiteral">func</a> [MakeFromLiteral](https://golang.org/src/go/constant/value.go?s=10767:10833#L369)
<pre>func MakeFromLiteral(lit <a href="/pkg/builtin/#string">string</a>, tok <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Token">Token</a>, zero <a href="/pkg/builtin/#uint">uint</a>) <a href="#Value">Value</a></pre>
MakeFromLiteral returns the corresponding integer, floating-point,
imaginary, character, or string value for a Go literal string. The
tok value must be one of token.INT, token.FLOAT, token.IMAG,
token.CHAR, or token.STRING. The final argument must be zero.
If the literal string syntax is invalid, the result is an Unknown.




### <a id="MakeImag">func</a> [MakeImag](https://golang.org/src/go/constant/value.go?s=20674:20702#L781)
<pre>func MakeImag(x <a href="#Value">Value</a>) <a href="#Value">Value</a></pre>
MakeImag returns the Complex value x*i;
x must be Int, Float, or Unknown.
If x is Unknown, the result is Unknown.




### <a id="MakeInt64">func</a> [MakeInt64](https://golang.org/src/go/constant/value.go?s=9914:9943#L341)
<pre>func MakeInt64(x <a href="/pkg/builtin/#int64">int64</a>) <a href="#Value">Value</a></pre>
MakeInt64 returns the Int value for x.




### <a id="MakeString">func</a> [MakeString](https://golang.org/src/go/constant/value.go?s=9811:9842#L338)
<pre>func MakeString(s <a href="/pkg/builtin/#string">string</a>) <a href="#Value">Value</a></pre>
MakeString returns the String value for s.




### <a id="MakeUint64">func</a> [MakeUint64](https://golang.org/src/go/constant/value.go?s=10011:10042#L344)
<pre>func MakeUint64(x <a href="/pkg/builtin/#uint64">uint64</a>) <a href="#Value">Value</a></pre>
MakeUint64 returns the Int value for x.




### <a id="MakeUnknown">func</a> [MakeUnknown](https://golang.org/src/go/constant/value.go?s=9622:9646#L332)
<pre>func MakeUnknown() <a href="#Value">Value</a></pre>
MakeUnknown returns the Unknown value.




### <a id="Num">func</a> [Num](https://golang.org/src/go/constant/value.go?s=19650:19673#L737)
<pre>func Num(x <a href="#Value">Value</a>) <a href="#Value">Value</a></pre>
Num returns the numerator of x; x must be Int, Float, or Unknown.
If x is Unknown, or if it is too large or small to represent as a
fraction, the result is Unknown. Otherwise the result is an Int
with the same sign as x.




### <a id="Real">func</a> [Real](https://golang.org/src/go/constant/value.go?s=21016:21040#L794)
<pre>func Real(x <a href="#Value">Value</a>) <a href="#Value">Value</a></pre>
Real returns the real part of x, which must be a numeric or unknown value.
If x is Unknown, the result is Unknown.




### <a id="Shift">func</a> [Shift](https://golang.org/src/go/constant/value.go?s=31317:31366#L1275)
<pre>func Shift(x <a href="#Value">Value</a>, op <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Token">Token</a>, s <a href="/pkg/builtin/#uint">uint</a>) <a href="#Value">Value</a></pre>
Shift returns the result of the shift expression x op s
with op == token.SHL or token.SHR (<< or >>). x must be
an Int or an Unknown. If x is Unknown, the result is x.




### <a id="ToComplex">func</a> [ToComplex](https://golang.org/src/go/constant/value.go?s=23455:23484#L898)
<pre>func ToComplex(x <a href="#Value">Value</a>) <a href="#Value">Value</a></pre>
ToComplex converts x to a Complex value if x is representable as a Complex.
Otherwise it returns an Unknown.




### <a id="ToFloat">func</a> [ToFloat](https://golang.org/src/go/constant/value.go?s=23021:23048#L879)
<pre>func ToFloat(x <a href="#Value">Value</a>) <a href="#Value">Value</a></pre>
ToFloat converts x to a Float value if x is representable as a Float.
Otherwise it returns an Unknown.




### <a id="ToInt">func</a> [ToInt](https://golang.org/src/go/constant/value.go?s=21790:21815#L825)
<pre>func ToInt(x <a href="#Value">Value</a>) <a href="#Value">Value</a></pre>
ToInt converts x to an Int value if x is representable as an Int.
Otherwise it returns an Unknown.




### <a id="UnaryOp">func</a> [UnaryOp](https://golang.org/src/go/constant/value.go?s=24323:24377#L934)
<pre>func UnaryOp(op <a href="/pkg/go/token/">token</a>.<a href="/pkg/go/token/#Token">Token</a>, y <a href="#Value">Value</a>, prec <a href="/pkg/builtin/#uint">uint</a>) <a href="#Value">Value</a></pre>
UnaryOp returns the result of the unary expression op y.
The operation must be defined for the operand.
If prec > 0 it specifies the ^ (xor) result size in bits.
If y is Unknown, the result is Unknown.



<a id="example_UnaryOp">Example</a>


```go
```

output:
```txt
```







