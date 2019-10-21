

# math
`import "math"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package math provides basic constants and mathematical functions.

This package does not guarantee bit-identical results across architectures.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func Abs(x float64) float64](#Abs)
* [func Acos(x float64) float64](#Acos)
* [func Acosh(x float64) float64](#Acosh)
* [func Asin(x float64) float64](#Asin)
* [func Asinh(x float64) float64](#Asinh)
* [func Atan(x float64) float64](#Atan)
* [func Atan2(y, x float64) float64](#Atan2)
* [func Atanh(x float64) float64](#Atanh)
* [func Cbrt(x float64) float64](#Cbrt)
* [func Ceil(x float64) float64](#Ceil)
* [func Copysign(x, y float64) float64](#Copysign)
* [func Cos(x float64) float64](#Cos)
* [func Cosh(x float64) float64](#Cosh)
* [func Dim(x, y float64) float64](#Dim)
* [func Erf(x float64) float64](#Erf)
* [func Erfc(x float64) float64](#Erfc)
* [func Erfcinv(x float64) float64](#Erfcinv)
* [func Erfinv(x float64) float64](#Erfinv)
* [func Exp(x float64) float64](#Exp)
* [func Exp2(x float64) float64](#Exp2)
* [func Expm1(x float64) float64](#Expm1)
* [func Float32bits(f float32) uint32](#Float32bits)
* [func Float32frombits(b uint32) float32](#Float32frombits)
* [func Float64bits(f float64) uint64](#Float64bits)
* [func Float64frombits(b uint64) float64](#Float64frombits)
* [func Floor(x float64) float64](#Floor)
* [func Frexp(f float64) (frac float64, exp int)](#Frexp)
* [func Gamma(x float64) float64](#Gamma)
* [func Hypot(p, q float64) float64](#Hypot)
* [func Ilogb(x float64) int](#Ilogb)
* [func Inf(sign int) float64](#Inf)
* [func IsInf(f float64, sign int) bool](#IsInf)
* [func IsNaN(f float64) (is bool)](#IsNaN)
* [func J0(x float64) float64](#J0)
* [func J1(x float64) float64](#J1)
* [func Jn(n int, x float64) float64](#Jn)
* [func Ldexp(frac float64, exp int) float64](#Ldexp)
* [func Lgamma(x float64) (lgamma float64, sign int)](#Lgamma)
* [func Log(x float64) float64](#Log)
* [func Log10(x float64) float64](#Log10)
* [func Log1p(x float64) float64](#Log1p)
* [func Log2(x float64) float64](#Log2)
* [func Logb(x float64) float64](#Logb)
* [func Max(x, y float64) float64](#Max)
* [func Min(x, y float64) float64](#Min)
* [func Mod(x, y float64) float64](#Mod)
* [func Modf(f float64) (int float64, frac float64)](#Modf)
* [func NaN() float64](#NaN)
* [func Nextafter(x, y float64) (r float64)](#Nextafter)
* [func Nextafter32(x, y float32) (r float32)](#Nextafter32)
* [func Pow(x, y float64) float64](#Pow)
* [func Pow10(n int) float64](#Pow10)
* [func Remainder(x, y float64) float64](#Remainder)
* [func Round(x float64) float64](#Round)
* [func RoundToEven(x float64) float64](#RoundToEven)
* [func Signbit(x float64) bool](#Signbit)
* [func Sin(x float64) float64](#Sin)
* [func Sincos(x float64) (sin, cos float64)](#Sincos)
* [func Sinh(x float64) float64](#Sinh)
* [func Sqrt(x float64) float64](#Sqrt)
* [func Tan(x float64) float64](#Tan)
* [func Tanh(x float64) float64](#Tanh)
* [func Trunc(x float64) float64](#Trunc)
* [func Y0(x float64) float64](#Y0)
* [func Y1(x float64) float64](#Y1)
* [func Yn(n int, x float64) float64](#Yn)


#### <a id="pkg-examples">Examples</a>
* [Abs](#example_Abs)
* [Acos](#example_Acos)
* [Acosh](#example_Acosh)
* [Asin](#example_Asin)
* [Asinh](#example_Asinh)
* [Atan](#example_Atan)
* [Atan2](#example_Atan2)
* [Atanh](#example_Atanh)
* [Ceil](#example_Ceil)
* [Cos](#example_Cos)
* [Cosh](#example_Cosh)
* [Floor](#example_Floor)
* [Log](#example_Log)
* [Log10](#example_Log10)
* [Log2](#example_Log2)
* [Mod](#example_Mod)
* [Pow](#example_Pow)
* [Pow10](#example_Pow10)
* [Round](#example_Round)
* [RoundToEven](#example_RoundToEven)
* [Sin](#example_Sin)
* [Sincos](#example_Sincos)
* [Sinh](#example_Sinh)
* [Sqrt](#example_Sqrt)
* [Tan](#example_Tan)
* [Tanh](#example_Tanh)


#### <a id="pkg-files">Package files</a>
[abs.go](https://golang.org/src/math/abs.go) [acosh.go](https://golang.org/src/math/acosh.go) [asin.go](https://golang.org/src/math/asin.go) [asinh.go](https://golang.org/src/math/asinh.go) [atan.go](https://golang.org/src/math/atan.go) [atan2.go](https://golang.org/src/math/atan2.go) [atanh.go](https://golang.org/src/math/atanh.go) [bits.go](https://golang.org/src/math/bits.go) [cbrt.go](https://golang.org/src/math/cbrt.go) [const.go](https://golang.org/src/math/const.go) [copysign.go](https://golang.org/src/math/copysign.go) [dim.go](https://golang.org/src/math/dim.go) [erf.go](https://golang.org/src/math/erf.go) [erfinv.go](https://golang.org/src/math/erfinv.go) [exp.go](https://golang.org/src/math/exp.go) [exp_asm.go](https://golang.org/src/math/exp_asm.go) [expm1.go](https://golang.org/src/math/expm1.go) [floor.go](https://golang.org/src/math/floor.go) [frexp.go](https://golang.org/src/math/frexp.go) [gamma.go](https://golang.org/src/math/gamma.go) [hypot.go](https://golang.org/src/math/hypot.go) [j0.go](https://golang.org/src/math/j0.go) [j1.go](https://golang.org/src/math/j1.go) [jn.go](https://golang.org/src/math/jn.go) [ldexp.go](https://golang.org/src/math/ldexp.go) [lgamma.go](https://golang.org/src/math/lgamma.go) [log.go](https://golang.org/src/math/log.go) [log10.go](https://golang.org/src/math/log10.go) [log1p.go](https://golang.org/src/math/log1p.go) [logb.go](https://golang.org/src/math/logb.go) [mod.go](https://golang.org/src/math/mod.go) [modf.go](https://golang.org/src/math/modf.go) [nextafter.go](https://golang.org/src/math/nextafter.go) [pow.go](https://golang.org/src/math/pow.go) [pow10.go](https://golang.org/src/math/pow10.go) [remainder.go](https://golang.org/src/math/remainder.go) [signbit.go](https://golang.org/src/math/signbit.go) [sin.go](https://golang.org/src/math/sin.go) [sincos.go](https://golang.org/src/math/sincos.go) [sinh.go](https://golang.org/src/math/sinh.go) [sqrt.go](https://golang.org/src/math/sqrt.go) [tan.go](https://golang.org/src/math/tan.go) [tanh.go](https://golang.org/src/math/tanh.go) [trig_reduce.go](https://golang.org/src/math/trig_reduce.go) [unsafe.go](https://golang.org/src/math/unsafe.go) 


## <a id="pkg-constants">Constants</a>
Mathematical constants.


<pre>const (
    <span id="E">E</span>   = 2.71828182845904523536028747135266249775724709369995957496696763 <span class="comment">// https://oeis.org/A001113</span>
    <span id="Pi">Pi</span>  = 3.14159265358979323846264338327950288419716939937510582097494459 <span class="comment">// https://oeis.org/A000796</span>
    <span id="Phi">Phi</span> = 1.61803398874989484820458683436563811772030917980576286213544862 <span class="comment">// https://oeis.org/A001622</span>

    <span id="Sqrt2">Sqrt2</span>   = 1.41421356237309504880168872420969807856967187537694807317667974 <span class="comment">// https://oeis.org/A002193</span>
    <span id="SqrtE">SqrtE</span>   = 1.64872127070012814684865078781416357165377610071014801157507931 <span class="comment">// https://oeis.org/A019774</span>
    <span id="SqrtPi">SqrtPi</span>  = 1.77245385090551602729816748334114518279754945612238712821380779 <span class="comment">// https://oeis.org/A002161</span>
    <span id="SqrtPhi">SqrtPhi</span> = 1.27201964951406896425242246173749149171560804184009624861664038 <span class="comment">// https://oeis.org/A139339</span>

    <span id="Ln2">Ln2</span>    = 0.693147180559945309417232121458176568075500134360255254120680009 <span class="comment">// https://oeis.org/A002162</span>
    <span id="Log2E">Log2E</span>  = 1 / <a href="#Ln2">Ln2</a>
    <span id="Ln10">Ln10</span>   = 2.30258509299404568401799145468436420760110148862877297603332790 <span class="comment">// https://oeis.org/A002392</span>
    <span id="Log10E">Log10E</span> = 1 / <a href="#Ln10">Ln10</a>
)</pre>Floating-point limit values.
Max is the largest finite value representable by the type.
SmallestNonzero is the smallest positive, non-zero value representable by the type.


<pre>const (
    <span id="MaxFloat32">MaxFloat32</span>             = 3.40282346638528859811704183484516925440e+38  <span class="comment">// 2**127 * (2**24 - 1) / 2**23</span>
    <span id="SmallestNonzeroFloat32">SmallestNonzeroFloat32</span> = 1.401298464324817070923729583289916131280e-45 <span class="comment">// 1 / 2**(127 - 1 + 23)</span>

    <span id="MaxFloat64">MaxFloat64</span>             = 1.797693134862315708145274237317043567981e+308 <span class="comment">// 2**1023 * (2**53 - 1) / 2**52</span>
    <span id="SmallestNonzeroFloat64">SmallestNonzeroFloat64</span> = 4.940656458412465441765687928682213723651e-324 <span class="comment">// 1 / 2**(1023 - 1 + 52)</span>
)</pre>Integer limit values.


<pre>const (
    <span id="MaxInt8">MaxInt8</span>   = 1&lt;&lt;7 - 1
    <span id="MinInt8">MinInt8</span>   = -1 &lt;&lt; 7
    <span id="MaxInt16">MaxInt16</span>  = 1&lt;&lt;15 - 1
    <span id="MinInt16">MinInt16</span>  = -1 &lt;&lt; 15
    <span id="MaxInt32">MaxInt32</span>  = 1&lt;&lt;31 - 1
    <span id="MinInt32">MinInt32</span>  = -1 &lt;&lt; 31
    <span id="MaxInt64">MaxInt64</span>  = 1&lt;&lt;63 - 1
    <span id="MinInt64">MinInt64</span>  = -1 &lt;&lt; 63
    <span id="MaxUint8">MaxUint8</span>  = 1&lt;&lt;8 - 1
    <span id="MaxUint16">MaxUint16</span> = 1&lt;&lt;16 - 1
    <span id="MaxUint32">MaxUint32</span> = 1&lt;&lt;32 - 1
    <span id="MaxUint64">MaxUint64</span> = 1&lt;&lt;64 - 1
)</pre>



## <a id="Abs">func</a> [Abs](https://golang.org/src/math/abs.go?s=278:305#L2)
<pre>func Abs(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Abs returns the absolute value of x.

Special cases are:


	Abs(±Inf) = +Inf
	Abs(NaN) = NaN


<a id="example_Abs">Example</a>
```go
```

output:
```txt
```

## <a id="Acos">func</a> [Acos](https://golang.org/src/math/asin.go?s=897:925#L41)
<pre>func Acos(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Acos returns the arccosine, in radians, of x.

Special case is:


	Acos(x) = NaN if x < -1 or x > 1


<a id="example_Acos">Example</a>
```go
```

output:
```txt
```

## <a id="Acosh">func</a> [Acosh](https://golang.org/src/math/acosh.go?s=1294:1323#L32)
<pre>func Acosh(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Acosh returns the inverse hyperbolic cosine of x.

Special cases are:


	Acosh(+Inf) = +Inf
	Acosh(x) = NaN if x < 1
	Acosh(NaN) = NaN


<a id="example_Acosh">Example</a>
```go
```

output:
```txt
```

## <a id="Asin">func</a> [Asin](https://golang.org/src/math/asin.go?s=434:462#L9)
<pre>func Asin(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Asin returns the arcsine, in radians, of x.

Special cases are:


	Asin(±0) = ±0
	Asin(x) = NaN if x < -1 or x > 1


<a id="example_Asin">Example</a>
```go
```

output:
```txt
```

## <a id="Asinh">func</a> [Asinh](https://golang.org/src/math/asinh.go?s=1228:1257#L29)
<pre>func Asinh(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Asinh returns the inverse hyperbolic sine of x.

Special cases are:


	Asinh(±0) = ±0
	Asinh(±Inf) = ±Inf
	Asinh(NaN) = NaN


<a id="example_Asinh">Example</a>
```go
```

output:
```txt
```

## <a id="Atan">func</a> [Atan](https://golang.org/src/math/atan.go?s=2901:2929#L85)
<pre>func Atan(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Atan returns the arctangent, in radians, of x.

Special cases are:


	Atan(±0) = ±0
	Atan(±Inf) = ±Pi/2


<a id="example_Atan">Example</a>
```go
```

output:
```txt
```

## <a id="Atan2">func</a> [Atan2](https://golang.org/src/math/atan2.go?s=770:802#L19)
<pre>func Atan2(y, x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Atan2 returns the arc tangent of y/x, using
the signs of the two to determine the quadrant
of the return value.

Special cases are (in order):


	Atan2(y, NaN) = NaN
	Atan2(NaN, x) = NaN
	Atan2(+0, x>=0) = +0
	Atan2(-0, x>=0) = -0
	Atan2(+0, x<=-0) = +Pi
	Atan2(-0, x<=-0) = -Pi
	Atan2(y>0, 0) = +Pi/2
	Atan2(y<0, 0) = -Pi/2
	Atan2(+Inf, +Inf) = +Pi/4
	Atan2(-Inf, +Inf) = -Pi/4
	Atan2(+Inf, -Inf) = 3Pi/4
	Atan2(-Inf, -Inf) = -3Pi/4
	Atan2(y, +Inf) = 0
	Atan2(y>0, -Inf) = +Pi
	Atan2(y<0, -Inf) = -Pi
	Atan2(+Inf, x) = +Pi/2
	Atan2(-Inf, x) = -Pi/2


<a id="example_Atan2">Example</a>
```go
```

output:
```txt
```

## <a id="Atanh">func</a> [Atanh](https://golang.org/src/math/atanh.go?s=1450:1479#L37)
<pre>func Atanh(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Atanh returns the inverse hyperbolic tangent of x.

Special cases are:


	Atanh(1) = +Inf
	Atanh(±0) = ±0
	Atanh(-1) = -Inf
	Atanh(x) = NaN if x < -1 or x > 1
	Atanh(NaN) = NaN


<a id="example_Atanh">Example</a>
```go
```

output:
```txt
```

## <a id="Cbrt">func</a> [Cbrt](https://golang.org/src/math/cbrt.go?s=807:835#L15)
<pre>func Cbrt(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Cbrt returns the cube root of x.

Special cases are:


	Cbrt(±0) = ±0
	Cbrt(±Inf) = ±Inf
	Cbrt(NaN) = NaN



## <a id="Ceil">func</a> [Ceil](https://golang.org/src/math/floor.go?s=720:748#L26)
<pre>func Ceil(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Ceil returns the least integer value greater than or equal to x.

Special cases are:


	Ceil(±0) = ±0
	Ceil(±Inf) = ±Inf
	Ceil(NaN) = NaN


<a id="example_Ceil">Example</a>
```go
```

output:
```txt
```

## <a id="Copysign">func</a> [Copysign](https://golang.org/src/math/copysign.go?s=248:283#L1)
<pre>func Copysign(x, y <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Copysign returns a value with the magnitude
of x and the sign of y.



## <a id="Cos">func</a> [Cos](https://golang.org/src/math/sin.go?s=3697:3724#L107)
<pre>func Cos(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Cos returns the cosine of the radian argument x.

Special cases are:


	Cos(±Inf) = NaN
	Cos(NaN) = NaN


<a id="example_Cos">Example</a>
```go
```

output:
```txt
```

## <a id="Cosh">func</a> [Cosh](https://golang.org/src/math/sinh.go?s=1448:1476#L62)
<pre>func Cosh(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Cosh returns the hyperbolic cosine of x.

Special cases are:


	Cosh(±0) = 1
	Cosh(±Inf) = +Inf
	Cosh(NaN) = NaN


<a id="example_Cosh">Example</a>
```go
```

output:
```txt
```

## <a id="Dim">func</a> [Dim](https://golang.org/src/math/dim.go?s=324:354#L3)
<pre>func Dim(x, y <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Dim returns the maximum of x-y or 0.

Special cases are:


	Dim(+Inf, +Inf) = NaN
	Dim(-Inf, -Inf) = NaN
	Dim(x, NaN) = Dim(NaN, x) = NaN



## <a id="Erf">func</a> [Erf](https://golang.org/src/math/erf.go?s=8274:8301#L178)
<pre>func Erf(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Erf returns the error function of x.

Special cases are:


	Erf(+Inf) = 1
	Erf(-Inf) = -1
	Erf(NaN) = NaN



## <a id="Erfc">func</a> [Erfc](https://golang.org/src/math/erf.go?s=10048:10076#L257)
<pre>func Erfc(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Erfc returns the complementary error function of x.

Special cases are:


	Erfc(+Inf) = 0
	Erfc(-Inf) = 2
	Erfc(NaN) = NaN



## <a id="Erfcinv">func</a> [Erfcinv](https://golang.org/src/math/erfinv.go?s=3384:3415#L115)
<pre>func Erfcinv(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Erfcinv returns the inverse of Erfc(x).

Special cases are:


	Erfcinv(0) = +Inf
	Erfcinv(2) = -Inf
	Erfcinv(x) = NaN if x < 0 or x > 2
	Erfcinv(NaN) = NaN



## <a id="Erfinv">func</a> [Erfinv](https://golang.org/src/math/erfinv.go?s=2367:2397#L66)
<pre>func Erfinv(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Erfinv returns the inverse error function of x.

Special cases are:


	Erfinv(1) = +Inf
	Erfinv(-1) = -Inf
	Erfinv(x) = NaN if x < -1 or x > 1
	Erfinv(NaN) = NaN



## <a id="Exp">func</a> [Exp](https://golang.org/src/math/exp.go?s=368:395#L4)
<pre>func Exp(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Exp returns e**x, the base-e exponential of x.

Special cases are:


	Exp(+Inf) = +Inf
	Exp(NaN) = NaN

Very large values overflow to 0 or +Inf.
Very small values underflow to 1.



## <a id="Exp2">func</a> [Exp2](https://golang.org/src/math/exp.go?s=4061:4089#L125)
<pre>func Exp2(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Exp2 returns 2**x, the base-2 exponential of x.

Special cases are the same as Exp.



## <a id="Expm1">func</a> [Expm1](https://golang.org/src/math/expm1.go?s=5195:5224#L114)
<pre>func Expm1(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Expm1 returns e**x - 1, the base-e exponential of x minus 1.
It is more accurate than Exp(x) - 1 when x is near zero.

Special cases are:


	Expm1(+Inf) = +Inf
	Expm1(-Inf) = -1
	Expm1(NaN) = NaN

Very large values overflow to -1 or +Inf.



## <a id="Float32bits">func</a> [Float32bits](https://golang.org/src/math/unsafe.go?s=363:397#L2)
<pre>func Float32bits(f <a href="/pkg/builtin/#float32">float32</a>) <a href="/pkg/builtin/#uint32">uint32</a></pre>
Float32bits returns the IEEE 754 binary representation of f,
with the sign bit of f and the result in the same bit position.
Float32bits(Float32frombits(x)) == x.



## <a id="Float32frombits">func</a> [Float32frombits](https://golang.org/src/math/unsafe.go?s=660:698#L8)
<pre>func Float32frombits(b <a href="/pkg/builtin/#uint32">uint32</a>) <a href="/pkg/builtin/#float32">float32</a></pre>
Float32frombits returns the floating-point number corresponding
to the IEEE 754 binary representation b, with the sign bit of b
and the result in the same bit position.
Float32frombits(Float32bits(x)) == x.



## <a id="Float64bits">func</a> [Float64bits](https://golang.org/src/math/unsafe.go?s=919:953#L13)
<pre>func Float64bits(f <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#uint64">uint64</a></pre>
Float64bits returns the IEEE 754 binary representation of f,
with the sign bit of f and the result in the same bit position,
and Float64bits(Float64frombits(x)) == x.



## <a id="Float64frombits">func</a> [Float64frombits](https://golang.org/src/math/unsafe.go?s=1216:1254#L19)
<pre>func Float64frombits(b <a href="/pkg/builtin/#uint64">uint64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Float64frombits returns the floating-point number corresponding
to the IEEE 754 binary representation b, with the sign bit of b
and the result in the same bit position.
Float64frombits(Float64bits(x)) == x.



## <a id="Floor">func</a> [Floor](https://golang.org/src/math/floor.go?s=332:361#L3)
<pre>func Floor(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Floor returns the greatest integer value less than or equal to x.

Special cases are:


	Floor(±0) = ±0
	Floor(±Inf) = ±Inf
	Floor(NaN) = NaN


<a id="example_Floor">Example</a>
```go
```

output:
```txt
```

## <a id="Frexp">func</a> [Frexp](https://golang.org/src/math/frexp.go?s=469:514#L6)
<pre>func Frexp(f <a href="/pkg/builtin/#float64">float64</a>) (frac <a href="/pkg/builtin/#float64">float64</a>, exp <a href="/pkg/builtin/#int">int</a>)</pre>
Frexp breaks f into a normalized fraction
and an integral power of two.
It returns frac and exp satisfying f == frac × 2**exp,
with the absolute value of frac in the interval [½, 1).

Special cases are:


	Frexp(±0) = ±0, 0
	Frexp(±Inf) = ±Inf, 0
	Frexp(NaN) = NaN, 0



## <a id="Gamma">func</a> [Gamma](https://golang.org/src/math/gamma.go?s=4029:4058#L120)
<pre>func Gamma(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Gamma returns the Gamma function of x.

Special cases are:


	Gamma(+Inf) = +Inf
	Gamma(+0) = +Inf
	Gamma(-0) = -Inf
	Gamma(x) = NaN for integer x < 0
	Gamma(-Inf) = NaN
	Gamma(NaN) = NaN



## <a id="Hypot">func</a> [Hypot](https://golang.org/src/math/hypot.go?s=464:496#L9)
<pre>func Hypot(p, q <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Hypot returns Sqrt(p*p + q*q), taking care to avoid
unnecessary overflow and underflow.

Special cases are:


	Hypot(±Inf, q) = +Inf
	Hypot(p, ±Inf) = +Inf
	Hypot(NaN, q) = NaN
	Hypot(p, NaN) = NaN



## <a id="Ilogb">func</a> [Ilogb](https://golang.org/src/math/logb.go?s=641:666#L22)
<pre>func Ilogb(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#int">int</a></pre>
Ilogb returns the binary exponent of x as an integer.

Special cases are:


	Ilogb(±Inf) = MaxInt32
	Ilogb(0) = MinInt32
	Ilogb(NaN) = MaxInt32



## <a id="Inf">func</a> [Inf](https://golang.org/src/math/bits.go?s=491:517#L10)
<pre>func Inf(sign <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.



## <a id="IsInf">func</a> [IsInf](https://golang.org/src/math/bits.go?s=1289:1325#L36)
<pre>func IsInf(f <a href="/pkg/builtin/#float64">float64</a>, sign <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsInf reports whether f is an infinity, according to sign.
If sign > 0, IsInf reports whether f is positive infinity.
If sign < 0, IsInf reports whether f is negative infinity.
If sign == 0, IsInf reports whether f is either infinity.



## <a id="IsNaN">func</a> [IsNaN](https://golang.org/src/math/bits.go?s=791:822#L24)
<pre>func IsNaN(f <a href="/pkg/builtin/#float64">float64</a>) (is <a href="/pkg/builtin/#bool">bool</a>)</pre>
IsNaN reports whether f is an IEEE 754 ``not-a-number'' value.



## <a id="J0">func</a> [J0](https://golang.org/src/math/j0.go?s=2892:2918#L66)
<pre>func J0(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
J0 returns the order-zero Bessel function of the first kind.

Special cases are:


	J0(±Inf) = 0
	J0(0) = 1
	J0(NaN) = NaN



## <a id="J1">func</a> [J1](https://golang.org/src/math/j1.go?s=2945:2971#L64)
<pre>func J1(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
J1 returns the order-one Bessel function of the first kind.

Special cases are:


	J1(±Inf) = 0
	J1(NaN) = NaN



## <a id="Jn">func</a> [Jn](https://golang.org/src/math/jn.go?s=1826:1859#L43)
<pre>func Jn(n <a href="/pkg/builtin/#int">int</a>, x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Jn returns the order-n Bessel function of the first kind.

Special cases are:


	Jn(n, ±Inf) = 0
	Jn(n, NaN) = NaN



## <a id="Ldexp">func</a> [Ldexp](https://golang.org/src/math/ldexp.go?s=342:383#L4)
<pre>func Ldexp(frac <a href="/pkg/builtin/#float64">float64</a>, exp <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Ldexp is the inverse of Frexp.
It returns frac × 2**exp.

Special cases are:


	Ldexp(±0, exp) = ±0
	Ldexp(±Inf, exp) = ±Inf
	Ldexp(NaN, exp) = NaN



## <a id="Lgamma">func</a> [Lgamma](https://golang.org/src/math/lgamma.go?s=7007:7056#L164)
<pre>func Lgamma(x <a href="/pkg/builtin/#float64">float64</a>) (lgamma <a href="/pkg/builtin/#float64">float64</a>, sign <a href="/pkg/builtin/#int">int</a>)</pre>
Lgamma returns the natural logarithm and sign (-1 or +1) of Gamma(x).

Special cases are:


	Lgamma(+Inf) = +Inf
	Lgamma(0) = +Inf
	Lgamma(-integer) = +Inf
	Lgamma(-Inf) = -Inf
	Lgamma(NaN) = NaN



## <a id="Log">func</a> [Log](https://golang.org/src/math/log.go?s=2809:2836#L70)
<pre>func Log(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Log returns the natural logarithm of x.

Special cases are:


	Log(+Inf) = +Inf
	Log(0) = -Inf
	Log(x < 0) = NaN
	Log(NaN) = NaN


<a id="example_Log">Example</a>
```go
```

output:
```txt
```

## <a id="Log10">func</a> [Log10](https://golang.org/src/math/log10.go?s=265:294#L1)
<pre>func Log10(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Log10 returns the decimal logarithm of x.
The special cases are the same as for Log.


<a id="example_Log10">Example</a>
```go
```

output:
```txt
```

## <a id="Log1p">func</a> [Log1p](https://golang.org/src/math/log1p.go?s=3722:3751#L85)
<pre>func Log1p(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Log1p returns the natural logarithm of 1 plus its argument x.
It is more accurate than Log(1 + x) when x is near zero.

Special cases are:


	Log1p(+Inf) = +Inf
	Log1p(±0) = ±0
	Log1p(-1) = -Inf
	Log1p(x < -1) = NaN
	Log1p(NaN) = NaN



## <a id="Log2">func</a> [Log2](https://golang.org/src/math/log10.go?s=448:476#L7)
<pre>func Log2(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Log2 returns the binary logarithm of x.
The special cases are the same as for Log.


<a id="example_Log2">Example</a>
```go
```

output:
```txt
```

## <a id="Logb">func</a> [Logb](https://golang.org/src/math/logb.go?s=300:328#L3)
<pre>func Logb(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Logb returns the binary exponent of x.

Special cases are:


	Logb(±Inf) = +Inf
	Logb(0) = -Inf
	Logb(NaN) = NaN



## <a id="Max">func</a> [Max](https://golang.org/src/math/dim.go?s=816:846#L25)
<pre>func Max(x, y <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Max returns the larger of x or y.

Special cases are:


	Max(x, +Inf) = Max(+Inf, x) = +Inf
	Max(x, NaN) = Max(NaN, x) = NaN
	Max(+0, ±0) = Max(±0, +0) = +0
	Max(-0, -0) = -0



## <a id="Min">func</a> [Min](https://golang.org/src/math/dim.go?s=1285:1315#L52)
<pre>func Min(x, y <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Min returns the smaller of x or y.

Special cases are:


	Min(x, -Inf) = Min(-Inf, x) = -Inf
	Min(x, NaN) = Min(NaN, x) = NaN
	Min(-0, ±0) = Min(±0, -0) = -0



## <a id="Mod">func</a> [Mod](https://golang.org/src/math/mod.go?s=483:513#L11)
<pre>func Mod(x, y <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Mod returns the floating-point remainder of x/y.
The magnitude of the result is less than y and its
sign agrees with that of x.

Special cases are:


	Mod(±Inf, y) = NaN
	Mod(NaN, y) = NaN
	Mod(x, 0) = NaN
	Mod(x, ±Inf) = x
	Mod(x, NaN) = NaN


<a id="example_Mod">Example</a>
```go
```

output:
```txt
```

## <a id="Modf">func</a> [Modf](https://golang.org/src/math/modf.go?s=368:416#L3)
<pre>func Modf(f <a href="/pkg/builtin/#float64">float64</a>) (int <a href="/pkg/builtin/#float64">float64</a>, frac <a href="/pkg/builtin/#float64">float64</a>)</pre>
Modf returns integer and fractional floating-point numbers
that sum to f. Both values have the same sign as f.

Special cases are:


	Modf(±Inf) = ±Inf, NaN
	Modf(NaN) = NaN, NaN



## <a id="NaN">func</a> [NaN](https://golang.org/src/math/bits.go?s=671:689#L21)
<pre>func NaN() <a href="/pkg/builtin/#float64">float64</a></pre>
NaN returns an IEEE 754 ``not-a-number'' value.



## <a id="Nextafter">func</a> [Nextafter](https://golang.org/src/math/nextafter.go?s=917:957#L25)
<pre>func Nextafter(x, y <a href="/pkg/builtin/#float64">float64</a>) (r <a href="/pkg/builtin/#float64">float64</a>)</pre>
Nextafter returns the next representable float64 value after x towards y.

Special cases are:


	Nextafter(x, x)   = x
	Nextafter(NaN, y) = NaN
	Nextafter(x, NaN) = NaN



## <a id="Nextafter32">func</a> [Nextafter32](https://golang.org/src/math/nextafter.go?s=363:405#L3)
<pre>func Nextafter32(x, y <a href="/pkg/builtin/#float32">float32</a>) (r <a href="/pkg/builtin/#float32">float32</a>)</pre>
Nextafter32 returns the next representable float32 value after x towards y.

Special cases are:


	Nextafter32(x, x)   = x
	Nextafter32(NaN, y) = NaN
	Nextafter32(x, NaN) = NaN



## <a id="Pow">func</a> [Pow](https://golang.org/src/math/pow.go?s=1186:1216#L28)
<pre>func Pow(x, y <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Pow returns x**y, the base-x exponential of y.

Special cases are (in order):


	Pow(x, ±0) = 1 for any x
	Pow(1, y) = 1 for any y
	Pow(x, 1) = x for any x
	Pow(NaN, y) = NaN
	Pow(x, NaN) = NaN
	Pow(±0, y) = ±Inf for y an odd integer < 0
	Pow(±0, -Inf) = +Inf
	Pow(±0, +Inf) = +0
	Pow(±0, y) = +Inf for finite y < 0 and not an odd integer
	Pow(±0, y) = ±0 for y an odd integer > 0
	Pow(±0, y) = +0 for finite y > 0 and not an odd integer
	Pow(-1, ±Inf) = 1
	Pow(x, +Inf) = +Inf for |x| > 1
	Pow(x, -Inf) = +0 for |x| > 1
	Pow(x, +Inf) = +0 for |x| < 1
	Pow(x, -Inf) = +Inf for |x| < 1
	Pow(+Inf, y) = +Inf for y > 0
	Pow(+Inf, y) = +0 for y < 0
	Pow(-Inf, y) = Pow(-0, -y)
	Pow(x, y) = NaN for finite x < 0 and finite non-integer y


<a id="example_Pow">Example</a>
```go
```

output:
```txt
```

## <a id="Pow10">func</a> [Pow10](https://golang.org/src/math/pow10.go?s=980:1005#L20)
<pre>func Pow10(n <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Pow10 returns 10**n, the base-10 exponential of n.

Special cases are:


	Pow10(n) =    0 for n < -323
	Pow10(n) = +Inf for n > 308


<a id="example_Pow10">Example</a>
```go
```

output:
```txt
```

## <a id="Remainder">func</a> [Remainder](https://golang.org/src/math/remainder.go?s=1283:1319#L27)
<pre>func Remainder(x, y <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Remainder returns the IEEE 754 floating-point remainder of x/y.

Special cases are:


	Remainder(±Inf, y) = NaN
	Remainder(NaN, y) = NaN
	Remainder(x, 0) = NaN
	Remainder(x, ±Inf) = x
	Remainder(x, NaN) = NaN



## <a id="Round">func</a> [Round](https://golang.org/src/math/floor.go?s=1237:1266#L54)
<pre>func Round(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Round returns the nearest integer, rounding half away from zero.

Special cases are:


	Round(±0) = ±0
	Round(±Inf) = ±Inf
	Round(NaN) = NaN


<a id="example_Round">Example</a>
```go
```

output:
```txt
```

## <a id="RoundToEven">func</a> [RoundToEven](https://golang.org/src/math/floor.go?s=2165:2200#L91)
<pre>func RoundToEven(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
RoundToEven returns the nearest integer, rounding ties to even.

Special cases are:


	RoundToEven(±0) = ±0
	RoundToEven(±Inf) = ±Inf
	RoundToEven(NaN) = NaN


<a id="example_RoundToEven">Example</a>
```go
```

output:
```txt
```

## <a id="Signbit">func</a> [Signbit](https://golang.org/src/math/signbit.go?s=233:261#L1)
<pre>func Signbit(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Signbit reports whether x is negative or negative zero.



## <a id="Sin">func</a> [Sin](https://golang.org/src/math/sin.go?s=5063:5090#L168)
<pre>func Sin(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Sin returns the sine of the radian argument x.

Special cases are:


	Sin(±0) = ±0
	Sin(±Inf) = NaN
	Sin(NaN) = NaN


<a id="example_Sin">Example</a>
```go
```

output:
```txt
```

## <a id="Sincos">func</a> [Sincos](https://golang.org/src/math/sincos.go?s=376:417#L5)
<pre>func Sincos(x <a href="/pkg/builtin/#float64">float64</a>) (sin, cos <a href="/pkg/builtin/#float64">float64</a>)</pre>
Sincos returns Sin(x), Cos(x).

Special cases are:


	Sincos(±0) = ±0, 1
	Sincos(±Inf) = NaN, NaN
	Sincos(NaN) = NaN, NaN


<a id="example_Sincos">Example</a>
```go
```

output:
```txt
```

## <a id="Sinh">func</a> [Sinh](https://golang.org/src/math/sinh.go?s=564:592#L15)
<pre>func Sinh(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Sinh returns the hyperbolic sine of x.

Special cases are:


	Sinh(±0) = ±0
	Sinh(±Inf) = ±Inf
	Sinh(NaN) = NaN


<a id="example_Sinh">Example</a>
```go
```

output:
```txt
```

## <a id="Sqrt">func</a> [Sqrt](https://golang.org/src/math/sqrt.go?s=3702:3730#L82)
<pre>func Sqrt(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Sqrt returns the square root of x.

Special cases are:


	Sqrt(+Inf) = +Inf
	Sqrt(±0) = ±0
	Sqrt(x < 0) = NaN
	Sqrt(NaN) = NaN


<a id="example_Sqrt">Example</a>
```go
```

output:
```txt
```

## <a id="Tan">func</a> [Tan](https://golang.org/src/math/tan.go?s=2594:2621#L72)
<pre>func Tan(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Tan returns the tangent of the radian argument x.

Special cases are:


	Tan(±0) = ±0
	Tan(±Inf) = NaN
	Tan(NaN) = NaN


<a id="example_Tan">Example</a>
```go
```

output:
```txt
```

## <a id="Tanh">func</a> [Tanh](https://golang.org/src/math/tanh.go?s=2211:2239#L64)
<pre>func Tanh(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Tanh returns the hyperbolic tangent of x.

Special cases are:


	Tanh(±0) = ±0
	Tanh(±Inf) = ±1
	Tanh(NaN) = NaN


<a id="example_Tanh">Example</a>
```go
```

output:
```txt
```

## <a id="Trunc">func</a> [Trunc](https://golang.org/src/math/floor.go?s=933:962#L38)
<pre>func Trunc(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Trunc returns the integer value of x.

Special cases are:


	Trunc(±0) = ±0
	Trunc(±Inf) = ±Inf
	Trunc(NaN) = NaN



## <a id="Y0">func</a> [Y0](https://golang.org/src/math/j0.go?s=4799:4825#L144)
<pre>func Y0(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Y0 returns the order-zero Bessel function of the second kind.

Special cases are:


	Y0(+Inf) = 0
	Y0(0) = -Inf
	Y0(x < 0) = NaN
	Y0(NaN) = NaN



## <a id="Y1">func</a> [Y1](https://golang.org/src/math/j1.go?s=4735:4761#L144)
<pre>func Y1(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Y1 returns the order-one Bessel function of the second kind.

Special cases are:


	Y1(+Inf) = 0
	Y1(0) = -Inf
	Y1(x < 0) = NaN
	Y1(NaN) = NaN



## <a id="Yn">func</a> [Yn](https://golang.org/src/math/jn.go?s=6078:6111#L223)
<pre>func Yn(n <a href="/pkg/builtin/#int">int</a>, x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Yn returns the order-n Bessel function of the second kind.

Special cases are:


	Yn(n, +Inf) = 0
	Yn(n ≥ 0, 0) = -Inf
	Yn(n < 0, 0) = +Inf if n is odd, -Inf if n is even
	Yn(n, x < 0) = NaN
	Yn(n, NaN) = NaN








