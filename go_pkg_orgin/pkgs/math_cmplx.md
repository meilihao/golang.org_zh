

# cmplx
`import "math/cmplx"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package cmplx provides basic constants and mathematical functions for
complex numbers.




## <a id="pkg-index">Index</a>
* [func Abs(x complex128) float64](#Abs)
* [func Acos(x complex128) complex128](#Acos)
* [func Acosh(x complex128) complex128](#Acosh)
* [func Asin(x complex128) complex128](#Asin)
* [func Asinh(x complex128) complex128](#Asinh)
* [func Atan(x complex128) complex128](#Atan)
* [func Atanh(x complex128) complex128](#Atanh)
* [func Conj(x complex128) complex128](#Conj)
* [func Cos(x complex128) complex128](#Cos)
* [func Cosh(x complex128) complex128](#Cosh)
* [func Cot(x complex128) complex128](#Cot)
* [func Exp(x complex128) complex128](#Exp)
* [func Inf() complex128](#Inf)
* [func IsInf(x complex128) bool](#IsInf)
* [func IsNaN(x complex128) bool](#IsNaN)
* [func Log(x complex128) complex128](#Log)
* [func Log10(x complex128) complex128](#Log10)
* [func NaN() complex128](#NaN)
* [func Phase(x complex128) float64](#Phase)
* [func Polar(x complex128) (r, θ float64)](#Polar)
* [func Pow(x, y complex128) complex128](#Pow)
* [func Rect(r, θ float64) complex128](#Rect)
* [func Sin(x complex128) complex128](#Sin)
* [func Sinh(x complex128) complex128](#Sinh)
* [func Sqrt(x complex128) complex128](#Sqrt)
* [func Tan(x complex128) complex128](#Tan)
* [func Tanh(x complex128) complex128](#Tanh)


#### <a id="pkg-examples">Examples</a>
* [Abs](#example_Abs)
* [Exp](#example_Exp)
* [Polar](#example_Polar)


#### <a id="pkg-files">Package files</a>
[abs.go](https://golang.org/src/math/cmplx/abs.go) [asin.go](https://golang.org/src/math/cmplx/asin.go) [conj.go](https://golang.org/src/math/cmplx/conj.go) [exp.go](https://golang.org/src/math/cmplx/exp.go) [isinf.go](https://golang.org/src/math/cmplx/isinf.go) [isnan.go](https://golang.org/src/math/cmplx/isnan.go) [log.go](https://golang.org/src/math/cmplx/log.go) [phase.go](https://golang.org/src/math/cmplx/phase.go) [polar.go](https://golang.org/src/math/cmplx/polar.go) [pow.go](https://golang.org/src/math/cmplx/pow.go) [rect.go](https://golang.org/src/math/cmplx/rect.go) [sin.go](https://golang.org/src/math/cmplx/sin.go) [sqrt.go](https://golang.org/src/math/cmplx/sqrt.go) [tan.go](https://golang.org/src/math/cmplx/tan.go) 






## <a id="Abs">func</a> [Abs](https://golang.org/src/math/cmplx/abs.go?s=349:379#L2)
<pre>func Abs(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Abs returns the absolute value (also called the modulus) of x.



<a id="example_Abs">Example</a>


```go
```

output:
```txt
```

## <a id="Acos">func</a> [Acos](https://golang.org/src/math/cmplx/asin.go?s=2730:2764#L77)
<pre>func Acos(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Acos returns the inverse cosine of x.



## <a id="Acosh">func</a> [Acosh](https://golang.org/src/math/cmplx/asin.go?s=2882:2917#L83)
<pre>func Acosh(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Acosh returns the inverse hyperbolic cosine of x.



## <a id="Asin">func</a> [Asin](https://golang.org/src/math/cmplx/asin.go?s=1687:1721#L41)
<pre>func Asin(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Asin returns the inverse sine of x.



## <a id="Asinh">func</a> [Asinh](https://golang.org/src/math/cmplx/asin.go?s=2101:2136#L54)
<pre>func Asinh(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Asinh returns the inverse hyperbolic sine of x.



## <a id="Atan">func</a> [Atan](https://golang.org/src/math/cmplx/asin.go?s=3945:3979#L125)
<pre>func Atan(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Atan returns the inverse tangent of x.



## <a id="Atanh">func</a> [Atanh](https://golang.org/src/math/cmplx/asin.go?s=4320:4355#L145)
<pre>func Atanh(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Atanh returns the inverse hyperbolic tangent of x.



## <a id="Conj">func</a> [Conj](https://golang.org/src/math/cmplx/conj.go?s=219:253#L1)
<pre>func Conj(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Conj returns the complex conjugate of x.



## <a id="Cos">func</a> [Cos](https://golang.org/src/math/cmplx/sin.go?s=2612:2645#L88)
<pre>func Cos(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Cos returns the cosine of x.



## <a id="Cosh">func</a> [Cosh](https://golang.org/src/math/cmplx/sin.go?s=3059:3093#L107)
<pre>func Cosh(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Cosh returns the hyperbolic cosine of x.



## <a id="Cot">func</a> [Cot](https://golang.org/src/math/cmplx/tan.go?s=4310:4343#L167)
<pre>func Cot(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Cot returns the cotangent of x.



## <a id="Exp">func</a> [Exp](https://golang.org/src/math/cmplx/exp.go?s=1624:1657#L41)
<pre>func Exp(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Exp returns e**x, the base-e exponential of x.



<a id="example_Exp">Example</a>
<p>ExampleExp computes Euler&#39;s identity.
</p>

```go
```

output:
```txt
```

## <a id="Inf">func</a> [Inf](https://golang.org/src/math/cmplx/isinf.go?s=434:455#L8)
<pre>func Inf() <a href="/pkg/builtin/#complex128">complex128</a></pre>
Inf returns a complex infinity, complex(+Inf, +Inf).



## <a id="IsInf">func</a> [IsInf](https://golang.org/src/math/cmplx/isinf.go?s=257:286#L1)
<pre>func IsInf(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsInf reports whether either real(x) or imag(x) is an infinity.



## <a id="IsNaN">func</a> [IsNaN](https://golang.org/src/math/cmplx/isnan.go?s=279:308#L1)
<pre>func IsNaN(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsNaN reports whether either real(x) or imag(x) is NaN
and neither is an infinity.



## <a id="Log">func</a> [Log](https://golang.org/src/math/cmplx/log.go?s=1845:1878#L47)
<pre>func Log(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Log returns the natural logarithm of x.



## <a id="Log10">func</a> [Log10](https://golang.org/src/math/cmplx/log.go?s=1973:2008#L52)
<pre>func Log10(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Log10 returns the decimal logarithm of x.



## <a id="NaN">func</a> [NaN](https://golang.org/src/math/cmplx/isnan.go?s=525:546#L12)
<pre>func NaN() <a href="/pkg/builtin/#complex128">complex128</a></pre>
NaN returns a complex ``not-a-number'' value.



## <a id="Phase">func</a> [Phase](https://golang.org/src/math/cmplx/phase.go?s=299:331#L1)
<pre>func Phase(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#float64">float64</a></pre>
Phase returns the phase (also called the argument) of x.
The returned value is in the range [-Pi, Pi].



## <a id="Polar">func</a> [Polar](https://golang.org/src/math/cmplx/polar.go?s=301:341#L1)
<pre>func Polar(x <a href="/pkg/builtin/#complex128">complex128</a>) (r, θ <a href="/pkg/builtin/#float64">float64</a>)</pre>
Polar returns the absolute value r and phase θ of x,
such that x = r * e**θi.
The phase is in the range [-Pi, Pi].



<a id="example_Polar">Example</a>


```go
```

output:
```txt
```

## <a id="Pow">func</a> [Pow](https://golang.org/src/math/cmplx/pow.go?s=1709:1745#L39)
<pre>func Pow(x, y <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Pow returns x**y, the base-x exponential of y.
For generalized compatibility with math.Pow:


	Pow(0, ±0) returns 1+0i
	Pow(0, c) for real(c)<0 returns Inf+0i if imag(c) is zero, otherwise Inf+Inf i.



## <a id="Rect">func</a> [Rect](https://golang.org/src/math/cmplx/rect.go?s=257:292#L1)
<pre>func Rect(r, θ <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Rect returns the complex number x with polar coordinates r, θ.



## <a id="Sin">func</a> [Sin](https://golang.org/src/math/cmplx/sin.go?s=1600:1633#L43)
<pre>func Sin(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Sin returns the sine of x.



## <a id="Sinh">func</a> [Sinh](https://golang.org/src/math/cmplx/sin.go?s=2082:2116#L63)
<pre>func Sinh(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Sinh returns the hyperbolic sine of x.



## <a id="Sqrt">func</a> [Sqrt](https://golang.org/src/math/cmplx/sqrt.go?s=1928:1962#L48)
<pre>func Sqrt(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Sqrt returns the square root of x.
The result r is chosen so that real(r) ≥ 0 and imag(r) has the same sign as imag(x).



## <a id="Tan">func</a> [Tan](https://golang.org/src/math/cmplx/tan.go?s=1834:1867#L49)
<pre>func Tan(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Tan returns the tangent of x.



## <a id="Tanh">func</a> [Tanh](https://golang.org/src/math/cmplx/tan.go?s=2396:2430#L73)
<pre>func Tanh(x <a href="/pkg/builtin/#complex128">complex128</a>) <a href="/pkg/builtin/#complex128">complex128</a></pre>
Tanh returns the hyperbolic tangent of x.








