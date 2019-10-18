

# big
`import "math/big"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package big implements arbitrary-precision arithmetic (big numbers).
The following numeric types are supported:


	Int    signed integers
	Rat    rational numbers
	Float  floating-point numbers

The zero value for an Int, Rat, or Float correspond to 0. Thus, new
values can be declared in the usual ways and denote 0 without further
initialization:


	var x Int        // &x is an *Int of value 0
	var r = &Rat{}   // r is a *Rat of value 0
	y := new(Float)  // y is a *Float of value 0

Alternatively, new values can be allocated and initialized with factory
functions of the form:


	func NewT(v V) *T

For instance, NewInt(x) returns an *Int set to the value of the int64
argument x, NewRat(a, b) returns a *Rat set to the fraction a/b where
a and b are int64 values, and NewFloat(f) returns a *Float initialized
to the float64 argument f. More flexibility is provided with explicit
setters, for instance:


	var z1 Int
	z1.SetUint64(123)                 // z1 := 123
	z2 := new(Rat).SetFloat64(1.25)   // z2 := 5/4
	z3 := new(Float).SetInt(z1)       // z3 := 123.0

Setters, numeric operations and predicates are represented as methods of
the form:


	func (z *T) SetV(v V) *T          // z = v
	func (z *T) Unary(x *T) *T        // z = unary x
	func (z *T) Binary(x, y *T) *T    // z = x binary y
	func (x *T) Pred() P              // p = pred(x)

with T one of Int, Rat, or Float. For unary and binary operations, the
result is the receiver (usually named z in that case; see below); if it
is one of the operands x or y it may be safely overwritten (and its memory
reused).

Arithmetic expressions are typically written as a sequence of individual
method calls, with each call corresponding to an operation. The receiver
denotes the result and the method arguments are the operation's operands.
For instance, given three *Int values a, b and c, the invocation


	c.Add(a, b)

computes the sum a + b and stores the result in c, overwriting whatever
value was held in c before. Unless specified otherwise, operations permit
aliasing of parameters, so it is perfectly ok to write


	sum.Add(sum, x)

to accumulate values x in a sum.

(By always passing in a result value via the receiver, memory use can be
much better controlled. Instead of having to allocate new memory for each
result, an operation can reuse the space allocated for the result value,
and overwrite that value with the new result in the process.)

Notational convention: Incoming method parameters (including the receiver)
are named consistently in the API to clarify their use. Incoming operands
are usually named x, y, a, b, and so on, but never z. A parameter specifying
the result is named z (typically the receiver).

For instance, the arguments for (*Int).Add are named x and y, and because
the receiver specifies the result destination, it is called z:


	func (z *Int) Add(x, y *Int) *Int

Methods of this form typically return the incoming receiver as well, to
enable simple call chaining.

Methods which don't require a result value to be passed in (for instance,
Int.Sign), simply return the result. In this case, the receiver is typically
the first operand, named x:


	func (x *Int) Sign() int

Various methods support conversions between strings and corresponding
numeric values, and vice versa: *Int, *Rat, and *Float values implement
the Stringer interface for a (default) string representation of the value,
but also provide SetString methods to initialize a value from a string in
a variety of supported formats (see the respective SetString documentation).

Finally, *Int, *Rat, and *Float satisfy the fmt package's Scanner interface
for scanning and (except for *Rat) the Formatter interface for formatted
printing.


<a id="example__eConvergents">Example (EConvergents)</a>
<p>This example demonstrates how to use big.Rat to compute the
first 15 terms in the sequence of rational convergents for
the constant e (base of natural logarithm).
</p><a id="example__fibonacci">Example (Fibonacci)</a>
<p>This example demonstrates how to use big.Int to compute the smallest
Fibonacci number with 100 decimal digits and to test whether it is prime.
</p><a id="example__sqrt2">Example (Sqrt2)</a>
<p>This example shows how to use big.Float to compute the square root of 2 with
a precision of 200 bits, and how to print the result as a decimal number.
</p>

## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func Jacobi(x, y *Int) int](#Jacobi)
* [type Accuracy](#Accuracy)
  * [func (i Accuracy) String() string](#Accuracy.String)
* [type ErrNaN](#ErrNaN)
  * [func (err ErrNaN) Error() string](#ErrNaN.Error)
* [type Float](#Float)
  * [func NewFloat(x float64) *Float](#NewFloat)
  * [func ParseFloat(s string, base int, prec uint, mode RoundingMode) (f *Float, b int, err error)](#ParseFloat)
  * [func (z *Float) Abs(x *Float) *Float](#Float.Abs)
  * [func (x *Float) Acc() Accuracy](#Float.Acc)
  * [func (z *Float) Add(x, y *Float) *Float](#Float.Add)
  * [func (x *Float) Append(buf []byte, fmt byte, prec int) []byte](#Float.Append)
  * [func (x *Float) Cmp(y *Float) int](#Float.Cmp)
  * [func (z *Float) Copy(x *Float) *Float](#Float.Copy)
  * [func (x *Float) Float32() (float32, Accuracy)](#Float.Float32)
  * [func (x *Float) Float64() (float64, Accuracy)](#Float.Float64)
  * [func (x *Float) Format(s fmt.State, format rune)](#Float.Format)
  * [func (z *Float) GobDecode(buf []byte) error](#Float.GobDecode)
  * [func (x *Float) GobEncode() ([]byte, error)](#Float.GobEncode)
  * [func (x *Float) Int(z *Int) (*Int, Accuracy)](#Float.Int)
  * [func (x *Float) Int64() (int64, Accuracy)](#Float.Int64)
  * [func (x *Float) IsInf() bool](#Float.IsInf)
  * [func (x *Float) IsInt() bool](#Float.IsInt)
  * [func (x *Float) MantExp(mant *Float) (exp int)](#Float.MantExp)
  * [func (x *Float) MarshalText() (text []byte, err error)](#Float.MarshalText)
  * [func (x *Float) MinPrec() uint](#Float.MinPrec)
  * [func (x *Float) Mode() RoundingMode](#Float.Mode)
  * [func (z *Float) Mul(x, y *Float) *Float](#Float.Mul)
  * [func (z *Float) Neg(x *Float) *Float](#Float.Neg)
  * [func (z *Float) Parse(s string, base int) (f *Float, b int, err error)](#Float.Parse)
  * [func (x *Float) Prec() uint](#Float.Prec)
  * [func (z *Float) Quo(x, y *Float) *Float](#Float.Quo)
  * [func (x *Float) Rat(z *Rat) (*Rat, Accuracy)](#Float.Rat)
  * [func (z *Float) Scan(s fmt.ScanState, ch rune) error](#Float.Scan)
  * [func (z *Float) Set(x *Float) *Float](#Float.Set)
  * [func (z *Float) SetFloat64(x float64) *Float](#Float.SetFloat64)
  * [func (z *Float) SetInf(signbit bool) *Float](#Float.SetInf)
  * [func (z *Float) SetInt(x *Int) *Float](#Float.SetInt)
  * [func (z *Float) SetInt64(x int64) *Float](#Float.SetInt64)
  * [func (z *Float) SetMantExp(mant *Float, exp int) *Float](#Float.SetMantExp)
  * [func (z *Float) SetMode(mode RoundingMode) *Float](#Float.SetMode)
  * [func (z *Float) SetPrec(prec uint) *Float](#Float.SetPrec)
  * [func (z *Float) SetRat(x *Rat) *Float](#Float.SetRat)
  * [func (z *Float) SetString(s string) (*Float, bool)](#Float.SetString)
  * [func (z *Float) SetUint64(x uint64) *Float](#Float.SetUint64)
  * [func (x *Float) Sign() int](#Float.Sign)
  * [func (x *Float) Signbit() bool](#Float.Signbit)
  * [func (z *Float) Sqrt(x *Float) *Float](#Float.Sqrt)
  * [func (x *Float) String() string](#Float.String)
  * [func (z *Float) Sub(x, y *Float) *Float](#Float.Sub)
  * [func (x *Float) Text(format byte, prec int) string](#Float.Text)
  * [func (x *Float) Uint64() (uint64, Accuracy)](#Float.Uint64)
  * [func (z *Float) UnmarshalText(text []byte) error](#Float.UnmarshalText)
* [type Int](#Int)
  * [func NewInt(x int64) *Int](#NewInt)
  * [func (z *Int) Abs(x *Int) *Int](#Int.Abs)
  * [func (z *Int) Add(x, y *Int) *Int](#Int.Add)
  * [func (z *Int) And(x, y *Int) *Int](#Int.And)
  * [func (z *Int) AndNot(x, y *Int) *Int](#Int.AndNot)
  * [func (x *Int) Append(buf []byte, base int) []byte](#Int.Append)
  * [func (z *Int) Binomial(n, k int64) *Int](#Int.Binomial)
  * [func (x *Int) Bit(i int) uint](#Int.Bit)
  * [func (x *Int) BitLen() int](#Int.BitLen)
  * [func (x *Int) Bits() []Word](#Int.Bits)
  * [func (x *Int) Bytes() []byte](#Int.Bytes)
  * [func (x *Int) Cmp(y *Int) (r int)](#Int.Cmp)
  * [func (x *Int) CmpAbs(y *Int) int](#Int.CmpAbs)
  * [func (z *Int) Div(x, y *Int) *Int](#Int.Div)
  * [func (z *Int) DivMod(x, y, m *Int) (*Int, *Int)](#Int.DivMod)
  * [func (z *Int) Exp(x, y, m *Int) *Int](#Int.Exp)
  * [func (x *Int) Format(s fmt.State, ch rune)](#Int.Format)
  * [func (z *Int) GCD(x, y, a, b *Int) *Int](#Int.GCD)
  * [func (z *Int) GobDecode(buf []byte) error](#Int.GobDecode)
  * [func (x *Int) GobEncode() ([]byte, error)](#Int.GobEncode)
  * [func (x *Int) Int64() int64](#Int.Int64)
  * [func (x *Int) IsInt64() bool](#Int.IsInt64)
  * [func (x *Int) IsUint64() bool](#Int.IsUint64)
  * [func (z *Int) Lsh(x *Int, n uint) *Int](#Int.Lsh)
  * [func (x *Int) MarshalJSON() ([]byte, error)](#Int.MarshalJSON)
  * [func (x *Int) MarshalText() (text []byte, err error)](#Int.MarshalText)
  * [func (z *Int) Mod(x, y *Int) *Int](#Int.Mod)
  * [func (z *Int) ModInverse(g, n *Int) *Int](#Int.ModInverse)
  * [func (z *Int) ModSqrt(x, p *Int) *Int](#Int.ModSqrt)
  * [func (z *Int) Mul(x, y *Int) *Int](#Int.Mul)
  * [func (z *Int) MulRange(a, b int64) *Int](#Int.MulRange)
  * [func (z *Int) Neg(x *Int) *Int](#Int.Neg)
  * [func (z *Int) Not(x *Int) *Int](#Int.Not)
  * [func (z *Int) Or(x, y *Int) *Int](#Int.Or)
  * [func (x *Int) ProbablyPrime(n int) bool](#Int.ProbablyPrime)
  * [func (z *Int) Quo(x, y *Int) *Int](#Int.Quo)
  * [func (z *Int) QuoRem(x, y, r *Int) (*Int, *Int)](#Int.QuoRem)
  * [func (z *Int) Rand(rnd *rand.Rand, n *Int) *Int](#Int.Rand)
  * [func (z *Int) Rem(x, y *Int) *Int](#Int.Rem)
  * [func (z *Int) Rsh(x *Int, n uint) *Int](#Int.Rsh)
  * [func (z *Int) Scan(s fmt.ScanState, ch rune) error](#Int.Scan)
  * [func (z *Int) Set(x *Int) *Int](#Int.Set)
  * [func (z *Int) SetBit(x *Int, i int, b uint) *Int](#Int.SetBit)
  * [func (z *Int) SetBits(abs []Word) *Int](#Int.SetBits)
  * [func (z *Int) SetBytes(buf []byte) *Int](#Int.SetBytes)
  * [func (z *Int) SetInt64(x int64) *Int](#Int.SetInt64)
  * [func (z *Int) SetString(s string, base int) (*Int, bool)](#Int.SetString)
  * [func (z *Int) SetUint64(x uint64) *Int](#Int.SetUint64)
  * [func (x *Int) Sign() int](#Int.Sign)
  * [func (z *Int) Sqrt(x *Int) *Int](#Int.Sqrt)
  * [func (x *Int) String() string](#Int.String)
  * [func (z *Int) Sub(x, y *Int) *Int](#Int.Sub)
  * [func (x *Int) Text(base int) string](#Int.Text)
  * [func (x *Int) TrailingZeroBits() uint](#Int.TrailingZeroBits)
  * [func (x *Int) Uint64() uint64](#Int.Uint64)
  * [func (z *Int) UnmarshalJSON(text []byte) error](#Int.UnmarshalJSON)
  * [func (z *Int) UnmarshalText(text []byte) error](#Int.UnmarshalText)
  * [func (z *Int) Xor(x, y *Int) *Int](#Int.Xor)
* [type Rat](#Rat)
  * [func NewRat(a, b int64) *Rat](#NewRat)
  * [func (z *Rat) Abs(x *Rat) *Rat](#Rat.Abs)
  * [func (z *Rat) Add(x, y *Rat) *Rat](#Rat.Add)
  * [func (x *Rat) Cmp(y *Rat) int](#Rat.Cmp)
  * [func (x *Rat) Denom() *Int](#Rat.Denom)
  * [func (x *Rat) Float32() (f float32, exact bool)](#Rat.Float32)
  * [func (x *Rat) Float64() (f float64, exact bool)](#Rat.Float64)
  * [func (x *Rat) FloatString(prec int) string](#Rat.FloatString)
  * [func (z *Rat) GobDecode(buf []byte) error](#Rat.GobDecode)
  * [func (x *Rat) GobEncode() ([]byte, error)](#Rat.GobEncode)
  * [func (z *Rat) Inv(x *Rat) *Rat](#Rat.Inv)
  * [func (x *Rat) IsInt() bool](#Rat.IsInt)
  * [func (x *Rat) MarshalText() (text []byte, err error)](#Rat.MarshalText)
  * [func (z *Rat) Mul(x, y *Rat) *Rat](#Rat.Mul)
  * [func (z *Rat) Neg(x *Rat) *Rat](#Rat.Neg)
  * [func (x *Rat) Num() *Int](#Rat.Num)
  * [func (z *Rat) Quo(x, y *Rat) *Rat](#Rat.Quo)
  * [func (x *Rat) RatString() string](#Rat.RatString)
  * [func (z *Rat) Scan(s fmt.ScanState, ch rune) error](#Rat.Scan)
  * [func (z *Rat) Set(x *Rat) *Rat](#Rat.Set)
  * [func (z *Rat) SetFloat64(f float64) *Rat](#Rat.SetFloat64)
  * [func (z *Rat) SetFrac(a, b *Int) *Rat](#Rat.SetFrac)
  * [func (z *Rat) SetFrac64(a, b int64) *Rat](#Rat.SetFrac64)
  * [func (z *Rat) SetInt(x *Int) *Rat](#Rat.SetInt)
  * [func (z *Rat) SetInt64(x int64) *Rat](#Rat.SetInt64)
  * [func (z *Rat) SetString(s string) (*Rat, bool)](#Rat.SetString)
  * [func (z *Rat) SetUint64(x uint64) *Rat](#Rat.SetUint64)
  * [func (x *Rat) Sign() int](#Rat.Sign)
  * [func (x *Rat) String() string](#Rat.String)
  * [func (z *Rat) Sub(x, y *Rat) *Rat](#Rat.Sub)
  * [func (z *Rat) UnmarshalText(text []byte) error](#Rat.UnmarshalText)
* [type RoundingMode](#RoundingMode)
  * [func (i RoundingMode) String() string](#RoundingMode.String)
* [type Word](#Word)


#### <a id="pkg-examples">Examples</a>
* [Float.Add](#example_Float_Add)
* [Float.Cmp](#example_Float_Cmp)
* [Float.Scan](#example_Float_Scan)
* [Float (Shift)](#example_Float_shift)
* [Int.Scan](#example_Int_Scan)
* [Int.SetString](#example_Int_SetString)
* [Rat.Scan](#example_Rat_Scan)
* [Rat.SetString](#example_Rat_SetString)
* [RoundingMode](#example_RoundingMode)
* [Package (EConvergents)](#example__eConvergents)
* [Package (Fibonacci)](#example__fibonacci)
* [Package (Sqrt2)](#example__sqrt2)


#### <a id="pkg-files">Package files</a>
[accuracy_string.go](https://golang.org/src/math/big/accuracy_string.go) [arith.go](https://golang.org/src/math/big/arith.go) [arith_amd64.go](https://golang.org/src/math/big/arith_amd64.go) [arith_decl.go](https://golang.org/src/math/big/arith_decl.go) [decimal.go](https://golang.org/src/math/big/decimal.go) [doc.go](https://golang.org/src/math/big/doc.go) [float.go](https://golang.org/src/math/big/float.go) [floatconv.go](https://golang.org/src/math/big/floatconv.go) [floatmarsh.go](https://golang.org/src/math/big/floatmarsh.go) [ftoa.go](https://golang.org/src/math/big/ftoa.go) [int.go](https://golang.org/src/math/big/int.go) [intconv.go](https://golang.org/src/math/big/intconv.go) [intmarsh.go](https://golang.org/src/math/big/intmarsh.go) [nat.go](https://golang.org/src/math/big/nat.go) [natconv.go](https://golang.org/src/math/big/natconv.go) [prime.go](https://golang.org/src/math/big/prime.go) [rat.go](https://golang.org/src/math/big/rat.go) [ratconv.go](https://golang.org/src/math/big/ratconv.go) [ratmarsh.go](https://golang.org/src/math/big/ratmarsh.go) [roundingmode_string.go](https://golang.org/src/math/big/roundingmode_string.go) [sqrt.go](https://golang.org/src/math/big/sqrt.go) 


## <a id="pkg-constants">Constants</a>
Exponent and precision limits.


<pre>const (
    <span id="MaxExp">MaxExp</span>  = <a href="/pkg/math/">math</a>.<a href="/pkg/math/#MaxInt32">MaxInt32</a>  <span class="comment">// largest supported exponent</span>
    <span id="MinExp">MinExp</span>  = <a href="/pkg/math/">math</a>.<a href="/pkg/math/#MinInt32">MinInt32</a>  <span class="comment">// smallest supported exponent</span>
    <span id="MaxPrec">MaxPrec</span> = <a href="/pkg/math/">math</a>.<a href="/pkg/math/#MaxUint32">MaxUint32</a> <span class="comment">// largest (theoretically) supported precision; likely memory-limited</span>
)</pre>MaxBase is the largest number base accepted for string conversions.


<pre>const <span id="MaxBase">MaxBase</span> = 10 + (&#39;z&#39; - &#39;a&#39; + 1) + (&#39;Z&#39; - &#39;A&#39; + 1)</pre>



## <a id="Jacobi">func</a> [Jacobi](https://golang.org/src/math/big/int.go?s=20459:20485#L782)
<pre>func Jacobi(x, y *<a href="#Int">Int</a>) <a href="/pkg/builtin/#int">int</a></pre>
Jacobi returns the Jacobi symbol (x/y), either +1, -1, or 0.
The y argument must be an odd integer.





## <a id="Accuracy">type</a> [Accuracy](https://golang.org/src/math/big/float.go?s=5878:5896#L138)
Accuracy describes the rounding error produced by the most recent
operation that generated a Float value, relative to the exact value.


<pre>type Accuracy <a href="/pkg/builtin/#int8">int8</a></pre>


Constants describing the Accuracy of a Float.


<pre>const (
    <span id="Below">Below</span> <a href="#Accuracy">Accuracy</a> = -1
    <span id="Exact">Exact</span> <a href="#Accuracy">Accuracy</a> = 0
    <span id="Above">Above</span> <a href="#Accuracy">Accuracy</a> = +1
)</pre>









### <a id="Accuracy.String">func</a> (Accuracy) [String](https://golang.org/src/math/big/accuracy_string.go?s=183:216#L1)
<pre>func (i <a href="#Accuracy">Accuracy</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="ErrNaN">type</a> [ErrNaN](https://golang.org/src/math/big/float.go?s=3256:3290#L67)
An ErrNaN panic is raised by a Float operation that would lead to
a NaN under IEEE-754 rules. An ErrNaN implements the error interface.


<pre>type ErrNaN struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="ErrNaN.Error">func</a> (ErrNaN) [Error](https://golang.org/src/math/big/float.go?s=3292:3324#L71)
<pre>func (err <a href="#ErrNaN">ErrNaN</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Float">type</a> [Float](https://golang.org/src/math/big/float.go?s=3000:3112#L55)
A nonzero finite Float represents a multi-precision floating point number


	sign × mantissa × 2**exponent

with 0.5 <= mantissa < 1.0, and MinExp <= exponent <= MaxExp.
A Float may also be zero (+0, -0) or infinite (+Inf, -Inf).
All Floats are ordered, and the ordering of two Floats x and y
is defined by x.Cmp(y).

Each Float value also has a precision, rounding mode, and accuracy.
The precision is the maximum number of mantissa bits available to
represent the value. The rounding mode specifies how a result should
be rounded to fit into the mantissa bits, and accuracy describes the
rounding error with respect to the exact result.

Unless specified otherwise, all operations (including setters) that
specify a *Float variable for the result (usually via the receiver
with the exception of MantExp), round the numeric result according
to the precision and rounding mode of the result variable.

If the provided result precision is 0 (see below), it is set to the
precision of the argument with the largest precision value before any
rounding takes place, and the rounding mode remains unchanged. Thus,
uninitialized Floats provided as result arguments will have their
precision set to a reasonable value determined by the operands, and
their mode is the zero value for RoundingMode (ToNearestEven).

By setting the desired precision to 24 or 53 and using matching rounding
mode (typically ToNearestEven), Float operations produce the same results
as the corresponding float32 or float64 IEEE-754 arithmetic for operands
that correspond to normal (i.e., not denormal) float32 or float64 numbers.
Exponent underflow and overflow lead to a 0 or an Infinity for different
values than IEEE-754 because Float exponents have a much larger range.

The zero (uninitialized) value for a Float is ready to use and represents
the number +0.0 exactly, with precision 0 and rounding mode ToNearestEven.

Operations always take pointer arguments (*Float) rather
than Float values, and each unique Float value requires
its own unique *Float pointer. To "copy" a Float value,
an existing (or newly allocated) Float must be set to
a new value using the Float.Set method; shallow copies
of Floats are not supported and may lead to errors.


<pre>type Float struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Float_shift">Example (Shift)</a>




### <a id="NewFloat">func</a> [NewFloat](https://golang.org/src/math/big/float.go?s=3502:3533#L78)
<pre>func NewFloat(x <a href="/pkg/builtin/#float64">float64</a>) *<a href="#Float">Float</a></pre>
NewFloat allocates and returns a new Float set to x,
with precision 53 and rounding mode ToNearestEven.
NewFloat panics with ErrNaN if x is a NaN.




### <a id="ParseFloat">func</a> [ParseFloat](https://golang.org/src/math/big/floatconv.go?s=7935:8029#L279)
<pre>func ParseFloat(s <a href="/pkg/builtin/#string">string</a>, base <a href="/pkg/builtin/#int">int</a>, prec <a href="/pkg/builtin/#uint">uint</a>, mode <a href="#RoundingMode">RoundingMode</a>) (f *<a href="#Float">Float</a>, b <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ParseFloat is like f.Parse(s, base) with f set to the given precision
and rounding mode.






### <a id="Float.Abs">func</a> (\*Float) [Abs](https://golang.org/src/math/big/float.go?s=31799:31835#L1161)
<pre>func (z *<a href="#Float">Float</a>) Abs(x *<a href="#Float">Float</a>) *<a href="#Float">Float</a></pre>
Abs sets z to the (possibly rounded) value |x| (the absolute value of x)
and returns z.




### <a id="Float.Acc">func</a> (\*Float) [Acc](https://golang.org/src/math/big/float.go?s=7803:7833#L218)
<pre>func (x *<a href="#Float">Float</a>) Acc() <a href="#Accuracy">Accuracy</a></pre>
Acc returns the accuracy of x produced by the most recent operation.




### <a id="Float.Add">func</a> (\*Float) [Add](https://golang.org/src/math/big/float.go?s=39582:39621#L1428)
<pre>func (z *<a href="#Float">Float</a>) Add(x, y *<a href="#Float">Float</a>) *<a href="#Float">Float</a></pre>
Add sets z to the rounded sum x+y and returns z. If z's precision is 0,
it is changed to the larger of x's or y's precision before the operation.
Rounding is performed according to z's precision and rounding mode; and
z's accuracy reports the result error relative to the exact (not rounded)
result. Add panics with ErrNaN if x and y are infinities with opposite
signs. The value of z is undefined in that case.


<a id="example_Float_Add">Example</a>


### <a id="Float.Append">func</a> (\*Float) [Append](https://golang.org/src/math/big/ftoa.go?s=2522:2583#L53)
<pre>func (x *<a href="#Float">Float</a>) Append(buf []<a href="/pkg/builtin/#byte">byte</a>, fmt <a href="/pkg/builtin/#byte">byte</a>, prec <a href="/pkg/builtin/#int">int</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
Append appends to buf the string form of the floating-point number x,
as generated by x.Text, and returns the extended buffer.




### <a id="Float.Cmp">func</a> (\*Float) [Cmp](https://golang.org/src/math/big/float.go?s=44468:44501#L1661)
<pre>func (x *<a href="#Float">Float</a>) Cmp(y *<a href="#Float">Float</a>) <a href="/pkg/builtin/#int">int</a></pre>
Cmp compares x and y and returns:


	-1 if x <  y
	 0 if x == y (incl. -0 == 0, -Inf == -Inf, and +Inf == +Inf)
	+1 if x >  y


<a id="example_Float_Cmp">Example</a>


### <a id="Float.Copy">func</a> (\*Float) [Copy](https://golang.org/src/math/big/float.go?s=18781:18818#L657)
<pre>func (z *<a href="#Float">Float</a>) Copy(x *<a href="#Float">Float</a>) *<a href="#Float">Float</a></pre>
Copy sets z to x, with the same precision, rounding mode, and
accuracy as x, and returns z. x is not changed even if z and
x are the same.




### <a id="Float.Float32">func</a> (\*Float) [Float32](https://golang.org/src/math/big/float.go?s=22187:22232#L820)
<pre>func (x *<a href="#Float">Float</a>) Float32() (<a href="/pkg/builtin/#float32">float32</a>, <a href="#Accuracy">Accuracy</a>)</pre>
Float32 returns the float32 value nearest to x. If x is too small to be
represented by a float32 (|x| < math.SmallestNonzeroFloat32), the result
is (0, Below) or (-0, Above), respectively, depending on the sign of x.
If x is too large to be represented by a float32 (|x| > math.MaxFloat32),
the result is (+Inf, Above) or (-Inf, Below), depending on the sign of x.




### <a id="Float.Float64">func</a> (\*Float) [Float64](https://golang.org/src/math/big/float.go?s=26040:26085#L940)
<pre>func (x *<a href="#Float">Float</a>) Float64() (<a href="/pkg/builtin/#float64">float64</a>, <a href="#Accuracy">Accuracy</a>)</pre>
Float64 returns the float64 value nearest to x. If x is too small to be
represented by a float64 (|x| < math.SmallestNonzeroFloat64), the result
is (0, Below) or (-0, Above), respectively, depending on the sign of x.
If x is too large to be represented by a float64 (|x| > math.MaxFloat64),
the result is (+Inf, Above) or (-Inf, Below), depending on the sign of x.




### <a id="Float.Format">func</a> (\*Float) [Format](https://golang.org/src/math/big/ftoa.go?s=12434:12482#L456)
<pre>func (x *<a href="#Float">Float</a>) Format(s <a href="/pkg/fmt/">fmt</a>.<a href="/pkg/fmt/#State">State</a>, format <a href="/pkg/builtin/#rune">rune</a>)</pre>
Format implements fmt.Formatter. It accepts all the regular
formats for floating-point numbers ('b', 'e', 'E', 'f', 'F',
'g', 'G', 'x') as well as 'p' and 'v'. See (*Float).Text for the
interpretation of 'p'. The 'v' format is handled like 'g'.
Format also supports specification of the minimum precision
in digits, the output field width, as well as the format flags
'+' and ' ' for sign control, '0' for space or zero padding,
and '-' for left or right justification. See the fmt package
for details.




### <a id="Float.GobDecode">func</a> (\*Float) [GobDecode](https://golang.org/src/math/big/floatmarsh.go?s=1925:1968#L54)
<pre>func (z *<a href="#Float">Float</a>) GobDecode(buf []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
GobDecode implements the gob.GobDecoder interface.
The result is rounded per the precision and rounding mode of
z unless z's precision is 0, in which case z is set exactly
to the decoded value.




### <a id="Float.GobEncode">func</a> (\*Float) [GobEncode](https://golang.org/src/math/big/floatmarsh.go?s=523:566#L10)
<pre>func (x *<a href="#Float">Float</a>) GobEncode() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
GobEncode implements the gob.GobEncoder interface.
The Float value and all its attributes (precision,
rounding mode, accuracy) are marshaled.




### <a id="Float.Int">func</a> (\*Float) [Int](https://golang.org/src/math/big/float.go?s=29766:29810#L1061)
<pre>func (x *<a href="#Float">Float</a>) Int(z *<a href="#Int">Int</a>) (*<a href="#Int">Int</a>, <a href="#Accuracy">Accuracy</a>)</pre>
Int returns the result of truncating x towards zero;
or nil if x is an infinity.
The result is Exact if x.IsInt(); otherwise it is Below
for x > 0, and Above for x < 0.
If a non-nil *Int argument z is provided, Int stores
the result in z instead of allocating a new Int.




### <a id="Float.Int64">func</a> (\*Float) [Int64](https://golang.org/src/math/big/float.go?s=20918:20959#L765)
<pre>func (x *<a href="#Float">Float</a>) Int64() (<a href="/pkg/builtin/#int64">int64</a>, <a href="#Accuracy">Accuracy</a>)</pre>
Int64 returns the integer resulting from truncating x towards zero.
If math.MinInt64 <= x <= math.MaxInt64, the result is Exact if x is
an integer, and Above (x < 0) or Below (x > 0) otherwise.
The result is (math.MinInt64, Above) for x < math.MinInt64,
and (math.MaxInt64, Below) for x > math.MaxInt64.




### <a id="Float.IsInf">func</a> (\*Float) [IsInf](https://golang.org/src/math/big/float.go?s=10041:10069#L326)
<pre>func (x *<a href="#Float">Float</a>) IsInf() <a href="/pkg/builtin/#bool">bool</a></pre>
IsInf reports whether x is +Inf or -Inf.




### <a id="Float.IsInt">func</a> (\*Float) [IsInt](https://golang.org/src/math/big/float.go?s=10173:10201#L332)
<pre>func (x *<a href="#Float">Float</a>) IsInt() <a href="/pkg/builtin/#bool">bool</a></pre>
IsInt reports whether x is an integer.
±Inf values are not integers.




### <a id="Float.MantExp">func</a> (\*Float) [MantExp](https://golang.org/src/math/big/float.go?s=8678:8724#L256)
<pre>func (x *<a href="#Float">Float</a>) MantExp(mant *<a href="#Float">Float</a>) (exp <a href="/pkg/builtin/#int">int</a>)</pre>
MantExp breaks x into its mantissa and exponent components
and returns the exponent. If a non-nil mant argument is
provided its value is set to the mantissa of x, with the
same precision and rounding mode as x. The components
satisfy x == mant × 2**exp, with 0.5 <= |mant| < 1.0.
Calling MantExp with a nil argument is an efficient way to
get the exponent of the receiver.

Special cases are:


	(  ±0).MantExp(mant) = 0, with mant set to   ±0
	(±Inf).MantExp(mant) = 0, with mant set to ±Inf

x and mant may be the same in which case x is set to its
mantissa value.




### <a id="Float.MarshalText">func</a> (\*Float) [MarshalText](https://golang.org/src/math/big/floatmarsh.go?s=2781:2835#L91)
<pre>func (x *<a href="#Float">Float</a>) MarshalText() (text []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
MarshalText implements the encoding.TextMarshaler interface.
Only the Float value is marshaled (in full precision), other
attributes such as precision or accuracy are ignored.




### <a id="Float.MinPrec">func</a> (\*Float) [MinPrec](https://golang.org/src/math/big/float.go?s=7505:7535#L205)
<pre>func (x *<a href="#Float">Float</a>) MinPrec() <a href="/pkg/builtin/#uint">uint</a></pre>
MinPrec returns the minimum precision required to represent x exactly
(i.e., the smallest prec before x.SetPrec(prec) would start rounding x).
The result is 0 for |x| == 0 and |x| == Inf.




### <a id="Float.Mode">func</a> (\*Float) [Mode](https://golang.org/src/math/big/float.go?s=7675:7710#L213)
<pre>func (x *<a href="#Float">Float</a>) Mode() <a href="#RoundingMode">RoundingMode</a></pre>
Mode returns the rounding mode of x.




### <a id="Float.Mul">func</a> (\*Float) [Mul](https://golang.org/src/math/big/float.go?s=42700:42739#L1569)
<pre>func (z *<a href="#Float">Float</a>) Mul(x, y *<a href="#Float">Float</a>) *<a href="#Float">Float</a></pre>
Mul sets z to the rounded product x*y and returns z.
Precision, rounding, and accuracy reporting are as for Add.
Mul panics with ErrNaN if one operand is zero and the other
operand an infinity. The value of z is undefined in that case.




### <a id="Float.Neg">func</a> (\*Float) [Neg](https://golang.org/src/math/big/float.go?s=31968:32004#L1169)
<pre>func (z *<a href="#Float">Float</a>) Neg(x *<a href="#Float">Float</a>) *<a href="#Float">Float</a></pre>
Neg sets z to the (possibly rounded) value of x with its sign negated,
and returns z.




### <a id="Float.Parse">func</a> (\*Float) [Parse](https://golang.org/src/math/big/floatconv.go?s=7239:7309#L251)
<pre>func (z *<a href="#Float">Float</a>) Parse(s <a href="/pkg/builtin/#string">string</a>, base <a href="/pkg/builtin/#int">int</a>) (f *<a href="#Float">Float</a>, b <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Parse parses s which must contain a text representation of a floating-
point number with a mantissa in the given conversion base (the exponent
is always a decimal number), or a string representing an infinite value.

For base 0, an underscore character ``_'' may appear between a base
prefix and an adjacent digit, and between successive digits; such
underscores do not change the value of the number, or the returned
digit count. Incorrect placement of underscores is reported as an
error if there are no other errors. If base != 0, underscores are
not recognized and thus terminate scanning like any other character
that is not a valid radix point or digit.

It sets z to the (possibly rounded) value of the corresponding floating-
point value, and returns z, the actual base b, and an error err, if any.
The entire string (not just a prefix) must be consumed for success.
If z's precision is 0, it is changed to 64 before rounding takes effect.
The number must be of the form:


	number    = [ sign ] ( float | "inf" | "Inf" ) .
	sign      = "+" | "-" .
	float     = ( mantissa | prefix pmantissa ) [ exponent ] .
	prefix    = "0" [ "b" | "B" | "o" | "O" | "x" | "X" ] .
	mantissa  = digits "." [ digits ] | digits | "." digits .
	pmantissa = [ "_" ] digits "." [ digits ] | [ "_" ] digits | "." digits .
	exponent  = ( "e" | "E" | "p" | "P" ) [ sign ] digits .
	digits    = digit { [ "_" ] digit } .
	digit     = "0" ... "9" | "a" ... "z" | "A" ... "Z" .

The base argument must be 0, 2, 8, 10, or 16. Providing an invalid base
argument will lead to a run-time panic.

For base 0, the number prefix determines the actual base: A prefix of
``0b'' or ``0B'' selects base 2, ``0o'' or ``0O'' selects base 8, and
``0x'' or ``0X'' selects base 16. Otherwise, the actual base is 10 and
no prefix is accepted. The octal prefix "0" is not supported (a leading
"0" is simply considered a "0").

A "p" or "P" exponent indicates a base 2 (rather then base 10) exponent;
for instance, "0x1.fffffffffffffp1023" (using base 0) represents the
maximum float64 value. For hexadecimal mantissae, the exponent character
must be one of 'p' or 'P', if present (an "e" or "E" exponent indicator
cannot be distinguished from a mantissa digit).

The returned *Float f is nil and the value of z is valid but not
defined if an error is reported.




### <a id="Float.Prec">func</a> (\*Float) [Prec](https://golang.org/src/math/big/float.go?s=7254:7281#L198)
<pre>func (x *<a href="#Float">Float</a>) Prec() <a href="/pkg/builtin/#uint">uint</a></pre>
Prec returns the mantissa precision of x in bits.
The result may be 0 for |x| == 0 and |x| == Inf.




### <a id="Float.Quo">func</a> (\*Float) [Quo](https://golang.org/src/math/big/float.go?s=43621:43660#L1614)
<pre>func (z *<a href="#Float">Float</a>) Quo(x, y *<a href="#Float">Float</a>) *<a href="#Float">Float</a></pre>
Quo sets z to the rounded quotient x/y and returns z.
Precision, rounding, and accuracy reporting are as for Add.
Quo panics with ErrNaN if both operands are zero or infinities.
The value of z is undefined in that case.




### <a id="Float.Rat">func</a> (\*Float) [Rat](https://golang.org/src/math/big/float.go?s=30865:30909#L1117)
<pre>func (x *<a href="#Float">Float</a>) Rat(z *<a href="#Rat">Rat</a>) (*<a href="#Rat">Rat</a>, <a href="#Accuracy">Accuracy</a>)</pre>
Rat returns the rational number corresponding to x;
or nil if x is an infinity.
The result is Exact if x is not an Inf.
If a non-nil *Rat argument z is provided, Rat stores
the result in z instead of allocating a new Rat.




### <a id="Float.Scan">func</a> (\*Float) [Scan](https://golang.org/src/math/big/floatconv.go?s=8439:8491#L290)
<pre>func (z *<a href="#Float">Float</a>) Scan(s <a href="/pkg/fmt/">fmt</a>.<a href="/pkg/fmt/#ScanState">ScanState</a>, ch <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#error">error</a></pre>
Scan is a support routine for fmt.Scanner; it sets z to the value of
the scanned number. It accepts formats whose verbs are supported by
fmt.Scan for floating point values, which are:
'b' (binary), 'e', 'E', 'f', 'F', 'g' and 'G'.
Scan doesn't handle ±Inf.


<a id="example_Float_Scan">Example</a>


### <a id="Float.Set">func</a> (\*Float) [Set](https://golang.org/src/math/big/float.go?s=18319:18355#L633)
<pre>func (z *<a href="#Float">Float</a>) Set(x *<a href="#Float">Float</a>) *<a href="#Float">Float</a></pre>
Set sets z to the (possibly rounded) value of x and returns z.
If z's precision is 0, it is changed to the precision of x
before setting z (and rounding will have no effect).
Rounding is performed according to z's precision and rounding
mode; and z's accuracy reports the result error relative to the
exact (not rounded) result.




### <a id="Float.SetFloat64">func</a> (\*Float) [SetFloat64](https://golang.org/src/math/big/float.go?s=15682:15726#L531)
<pre>func (z *<a href="#Float">Float</a>) SetFloat64(x <a href="/pkg/builtin/#float64">float64</a>) *<a href="#Float">Float</a></pre>
SetFloat64 sets z to the (possibly rounded) value of x and returns z.
If z's precision is 0, it is changed to 53 (and rounding will have
no effect). SetFloat64 panics with ErrNaN if x is a NaN.




### <a id="Float.SetInf">func</a> (\*Float) [SetInf](https://golang.org/src/math/big/float.go?s=17867:17910#L620)
<pre>func (z *<a href="#Float">Float</a>) SetInf(signbit <a href="/pkg/builtin/#bool">bool</a>) *<a href="#Float">Float</a></pre>
SetInf sets z to the infinite Float -Inf if signbit is
set, or +Inf if signbit is not set, and returns z. The
precision of z is unchanged and the result is always
Exact.




### <a id="Float.SetInt">func</a> (\*Float) [SetInt](https://golang.org/src/math/big/float.go?s=16858:16895#L579)
<pre>func (z *<a href="#Float">Float</a>) SetInt(x *<a href="#Int">Int</a>) *<a href="#Float">Float</a></pre>
SetInt sets z to the (possibly rounded) value of x and returns z.
If z's precision is 0, it is changed to the larger of x.BitLen()
or 64 (and rounding will have no effect).




### <a id="Float.SetInt64">func</a> (\*Float) [SetInt64](https://golang.org/src/math/big/float.go?s=15244:15284#L518)
<pre>func (z *<a href="#Float">Float</a>) SetInt64(x <a href="/pkg/builtin/#int64">int64</a>) *<a href="#Float">Float</a></pre>
SetInt64 sets z to the (possibly rounded) value of x and returns z.
If z's precision is 0, it is changed to 64 (and rounding will have
no effect).




### <a id="Float.SetMantExp">func</a> (\*Float) [SetMantExp](https://golang.org/src/math/big/float.go?s=9667:9722#L307)
<pre>func (z *<a href="#Float">Float</a>) SetMantExp(mant *<a href="#Float">Float</a>, exp <a href="/pkg/builtin/#int">int</a>) *<a href="#Float">Float</a></pre>
SetMantExp sets z to mant × 2**exp and returns z.
The result z has the same precision and rounding mode
as mant. SetMantExp is an inverse of MantExp but does
not require 0.5 <= |mant| < 1.0. Specifically:


	mant := new(Float)
	new(Float).SetMantExp(mant, x.MantExp(mant)).Cmp(x) == 0

Special cases are:


	z.SetMantExp(  ±0, exp) =   ±0
	z.SetMantExp(±Inf, exp) = ±Inf

z and mant may be the same in which case z's exponent
is set to exp.




### <a id="Float.SetMode">func</a> (\*Float) [SetMode](https://golang.org/src/math/big/float.go?s=7054:7103#L190)
<pre>func (z *<a href="#Float">Float</a>) SetMode(mode <a href="#RoundingMode">RoundingMode</a>) *<a href="#Float">Float</a></pre>
SetMode sets z's rounding mode to mode and returns an exact z.
z remains unchanged otherwise.
z.SetMode(z.Mode()) is a cheap way to set z's accuracy to Exact.




### <a id="Float.SetPrec">func</a> (\*Float) [SetPrec](https://golang.org/src/math/big/float.go?s=6398:6439#L154)
<pre>func (z *<a href="#Float">Float</a>) SetPrec(prec <a href="/pkg/builtin/#uint">uint</a>) *<a href="#Float">Float</a></pre>
SetPrec sets z's precision to prec and returns the (possibly) rounded
value of z. Rounding occurs according to z's rounding mode if the mantissa
cannot be represented in prec bits without loss of precision.
SetPrec(0) maps all finite values to ±0; infinite values remain unchanged.
If prec > MaxPrec, it is set to MaxPrec.




### <a id="Float.SetRat">func</a> (\*Float) [SetRat](https://golang.org/src/math/big/float.go?s=17463:17500#L603)
<pre>func (z *<a href="#Float">Float</a>) SetRat(x *<a href="#Rat">Rat</a>) *<a href="#Float">Float</a></pre>
SetRat sets z to the (possibly rounded) value of x and returns z.
If z's precision is 0, it is changed to the largest of a.BitLen(),
b.BitLen(), or 64; with x = a/b.




### <a id="Float.SetString">func</a> (\*Float) [SetString](https://golang.org/src/math/big/floatconv.go?s=637:687#L12)
<pre>func (z *<a href="#Float">Float</a>) SetString(s <a href="/pkg/builtin/#string">string</a>) (*<a href="#Float">Float</a>, <a href="/pkg/builtin/#bool">bool</a>)</pre>
SetString sets z to the value of s and returns z and a boolean indicating
success. s must be a floating-point number of the same format as accepted
by Parse, with base argument 0. The entire string (not just a prefix) must
be valid for success. If the operation failed, the value of z is undefined
but the returned value is nil.




### <a id="Float.SetUint64">func</a> (\*Float) [SetUint64](https://golang.org/src/math/big/float.go?s=15010:15052#L511)
<pre>func (z *<a href="#Float">Float</a>) SetUint64(x <a href="/pkg/builtin/#uint64">uint64</a>) *<a href="#Float">Float</a></pre>
SetUint64 sets z to the (possibly rounded) value of x and returns z.
If z's precision is 0, it is changed to 64 (and rounding will have
no effect).




### <a id="Float.Sign">func</a> (\*Float) [Sign](https://golang.org/src/math/big/float.go?s=7928:7954#L228)
<pre>func (x *<a href="#Float">Float</a>) Sign() <a href="/pkg/builtin/#int">int</a></pre>
Sign returns:


	-1 if x <   0
	 0 if x is ±0
	+1 if x >   0




### <a id="Float.Signbit">func</a> (\*Float) [Signbit](https://golang.org/src/math/big/float.go?s=9947:9977#L321)
<pre>func (x *<a href="#Float">Float</a>) Signbit() <a href="/pkg/builtin/#bool">bool</a></pre>
Signbit reports whether x is negative or negative zero.




### <a id="Float.Sqrt">func</a> (\*Float) [Sqrt](https://golang.org/src/math/big/sqrt.go?s=521:558#L11)
<pre>func (z *<a href="#Float">Float</a>) Sqrt(x *<a href="#Float">Float</a>) *<a href="#Float">Float</a></pre>
Sqrt sets z to the rounded square root of x, and returns it.

If z's precision is 0, it is changed to x's precision before the
operation. Rounding is performed according to z's precision and
rounding mode.

The function panics if z < 0. The value of z is undefined in that
case.




### <a id="Float.String">func</a> (\*Float) [String](https://golang.org/src/math/big/ftoa.go?s=2328:2359#L47)
<pre>func (x *<a href="#Float">Float</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String formats x like x.Text('g', 10).
(String must be called explicitly, Float.Format does not support %s verb.)




### <a id="Float.Sub">func</a> (\*Float) [Sub](https://golang.org/src/math/big/float.go?s=41299:41338#L1502)
<pre>func (z *<a href="#Float">Float</a>) Sub(x, y *<a href="#Float">Float</a>) *<a href="#Float">Float</a></pre>
Sub sets z to the rounded difference x-y and returns z.
Precision, rounding, and accuracy reporting are as for Add.
Sub panics with ErrNaN if x and y are infinities with equal
signs. The value of z is undefined in that case.




### <a id="Float.Text">func</a> (\*Float) [Text](https://golang.org/src/math/big/ftoa.go?s=2000:2050#L37)
<pre>func (x *<a href="#Float">Float</a>) Text(format <a href="/pkg/builtin/#byte">byte</a>, prec <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#string">string</a></pre>
Text converts the floating-point number x to a string according
to the given format and precision prec. The format is one of:


	'e'	-d.dddde±dd, decimal exponent, at least two (possibly 0) exponent digits
	'E'	-d.ddddE±dd, decimal exponent, at least two (possibly 0) exponent digits
	'f'	-ddddd.dddd, no exponent
	'g'	like 'e' for large exponents, like 'f' otherwise
	'G'	like 'E' for large exponents, like 'f' otherwise
	'x'	-0xd.dddddp±dd, hexadecimal mantissa, decimal power of two exponent
	'p'	-0x.dddp±dd, hexadecimal mantissa, decimal power of two exponent (non-standard)
	'b'	-ddddddp±dd, decimal mantissa, decimal power of two exponent (non-standard)

For the power-of-two exponent formats, the mantissa is printed in normalized form:


	'x'	hexadecimal mantissa in [1, 2), or 0
	'p'	hexadecimal mantissa in [½, 1), or 0
	'b'	decimal integer mantissa using x.Prec() bits, or 0

Note that the 'x' form is the one used by most other languages and libraries.

If format is a different character, Text returns a "%" followed by the
unrecognized format character.

The precision prec controls the number of digits (excluding the exponent)
printed by the 'e', 'E', 'f', 'g', 'G', and 'x' formats.
For 'e', 'E', 'f', and 'x', it is the number of digits after the decimal point.
For 'g' and 'G' it is the total number of digits. A negative precision selects
the smallest number of decimal digits necessary to identify the value x uniquely
using x.Prec() mantissa bits.
The prec value is ignored for the 'b' and 'p' formats.




### <a id="Float.Uint64">func</a> (\*Float) [Uint64](https://golang.org/src/math/big/float.go?s=19978:20021#L720)
<pre>func (x *<a href="#Float">Float</a>) Uint64() (<a href="/pkg/builtin/#uint64">uint64</a>, <a href="#Accuracy">Accuracy</a>)</pre>
Uint64 returns the unsigned integer resulting from truncating x
towards zero. If 0 <= x <= math.MaxUint64, the result is Exact
if x is an integer and Below otherwise.
The result is (0, Above) for x < 0, and (math.MaxUint64, Below)
for x > math.MaxUint64.




### <a id="Float.UnmarshalText">func</a> (\*Float) [UnmarshalText](https://golang.org/src/math/big/floatmarsh.go?s=3155:3203#L103)
<pre>func (z *<a href="#Float">Float</a>) UnmarshalText(text []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
UnmarshalText implements the encoding.TextUnmarshaler interface.
The result is rounded per the precision and rounding mode of z.
If z's precision is 0, it is changed to 64 before rounding takes
effect.




## <a id="Int">type</a> [Int](https://golang.org/src/math/big/int.go?s=724:804#L15)
An Int represents a signed multi-precision integer.
The zero value for an Int represents the value 0.

Operations always take pointer arguments (*Int) rather
than Int values, and each unique Int value requires
its own unique *Int pointer. To "copy" an Int value,
an existing (or newly allocated) Int must be set to
a new value using the Int.Set method; shallow copies
of Ints are not supported and may lead to errors.


<pre>type Int struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewInt">func</a> [NewInt](https://golang.org/src/math/big/int.go?s=1394:1419#L58)
<pre>func NewInt(x <a href="/pkg/builtin/#int64">int64</a>) *<a href="#Int">Int</a></pre>
NewInt allocates and returns a new Int set to x.






### <a id="Int.Abs">func</a> (\*Int) [Abs](https://golang.org/src/math/big/int.go?s=2460:2490#L92)
<pre>func (z *<a href="#Int">Int</a>) Abs(x *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
Abs sets z to |x| (the absolute value of x) and returns z.




### <a id="Int.Add">func</a> (\*Int) [Add](https://golang.org/src/math/big/int.go?s=2717:2750#L106)
<pre>func (z *<a href="#Int">Int</a>) Add(x, y *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
Add sets z to the sum x+y and returns z.




### <a id="Int.And">func</a> (\*Int) [And](https://golang.org/src/math/big/int.go?s=26168:26201#L1016)
<pre>func (z *<a href="#Int">Int</a>) And(x, y *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
And sets z = x & y and returns z.




### <a id="Int.AndNot">func</a> (\*Int) [AndNot](https://golang.org/src/math/big/int.go?s=26849:26885#L1046)
<pre>func (z *<a href="#Int">Int</a>) AndNot(x, y *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
AndNot sets z = x &^ y and returns z.




### <a id="Int.Append">func</a> (\*Int) [Append](https://golang.org/src/math/big/intconv.go?s=864:913#L20)
<pre>func (x *<a href="#Int">Int</a>) Append(buf []<a href="/pkg/builtin/#byte">byte</a>, base <a href="/pkg/builtin/#int">int</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
Append appends the string representation of x, as generated by
x.Text(base), to buf and returns the extended buffer.




### <a id="Int.Binomial">func</a> (\*Int) [Binomial](https://golang.org/src/math/big/int.go?s=4533:4572#L187)
<pre>func (z *<a href="#Int">Int</a>) Binomial(n, k <a href="/pkg/builtin/#int64">int64</a>) *<a href="#Int">Int</a></pre>
Binomial sets z to the binomial coefficient of (n, k) and returns z.




### <a id="Int.Bit">func</a> (\*Int) [Bit](https://golang.org/src/math/big/int.go?s=25280:25309#L976)
<pre>func (x *<a href="#Int">Int</a>) Bit(i <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#uint">uint</a></pre>
Bit returns the value of the i'th bit of x. That is, it
returns (x>>i)&1. The bit index i must be >= 0.




### <a id="Int.BitLen">func</a> (\*Int) [BitLen](https://golang.org/src/math/big/int.go?s=11672:11698#L445)
<pre>func (x *<a href="#Int">Int</a>) BitLen() <a href="/pkg/builtin/#int">int</a></pre>
BitLen returns the length of the absolute value of x in bits.
The bit length of 0 is 0.




### <a id="Int.Bits">func</a> (\*Int) [Bits](https://golang.org/src/math/big/int.go?s=1908:1935#L76)
<pre>func (x *<a href="#Int">Int</a>) Bits() []<a href="#Word">Word</a></pre>
Bits provides raw (unchecked but fast) access to x by returning its
absolute value as a little-endian Word slice. The result and x share
the same underlying array.
Bits is intended to support implementation of missing low-level Int
functionality outside this package; it should be avoided otherwise.




### <a id="Int.Bytes">func</a> (\*Int) [Bytes](https://golang.org/src/math/big/int.go?s=11477:11505#L438)
<pre>func (x *<a href="#Int">Int</a>) Bytes() []<a href="/pkg/builtin/#byte">byte</a></pre>
Bytes returns the absolute value of x as a big-endian byte slice.




### <a id="Int.Cmp">func</a> (\*Int) [Cmp](https://golang.org/src/math/big/int.go?s=7878:7911#L310)
<pre>func (x *<a href="#Int">Int</a>) Cmp(y *<a href="#Int">Int</a>) (r <a href="/pkg/builtin/#int">int</a>)</pre>
Cmp compares x and y and returns:


	-1 if x <  y
	 0 if x == y
	+1 if x >  y




### <a id="Int.CmpAbs">func</a> (\*Int) [CmpAbs](https://golang.org/src/math/big/int.go?s=8280:8312#L335)
<pre>func (x *<a href="#Int">Int</a>) CmpAbs(y *<a href="#Int">Int</a>) <a href="/pkg/builtin/#int">int</a></pre>
CmpAbs compares the absolute values of x and y and returns:


	-1 if |x| <  |y|
	 0 if |x| == |y|
	+1 if |x| >  |y|




### <a id="Int.Div">func</a> (\*Int) [Div](https://golang.org/src/math/big/int.go?s=6294:6327#L237)
<pre>func (z *<a href="#Int">Int</a>) Div(x, y *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
Div sets z to the quotient x/y for y != 0 and returns z.
If y == 0, a division-by-zero run-time panic occurs.
Div implements Euclidean division (unlike Go); see DivMod for more details.




### <a id="Int.DivMod">func</a> (\*Int) [DivMod](https://golang.org/src/math/big/int.go?s=7499:7546#L286)
<pre>func (z *<a href="#Int">Int</a>) DivMod(x, y, m *<a href="#Int">Int</a>) (*<a href="#Int">Int</a>, *<a href="#Int">Int</a>)</pre>
DivMod sets z to the quotient x div y and m to the modulus x mod y
and returns the pair (z, m) for y != 0.
If y == 0, a division-by-zero run-time panic occurs.

DivMod implements Euclidean division and modulus (unlike Go):


	q = x div y  such that
	m = x - y*q  with 0 <= m < |y|

(See Raymond T. Boute, ``The Euclidean definition of the functions
div and mod''. ACM Transactions on Programming Languages and
Systems (TOPLAS), 14(2):127-144, New York, NY, USA, 4/1992.
ACM press.)
See QuoRem for T-division and modulus (like Go).




### <a id="Int.Exp">func</a> (\*Int) [Exp](https://golang.org/src/math/big/int.go?s=12242:12278#L461)
<pre>func (z *<a href="#Int">Int</a>) Exp(x, y, m *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
Exp sets z = x**y mod |m| (i.e. the sign of m is ignored), and returns z.
If m == nil or m == 0, z = x**y unless y <= 0 then z = 1. If m > 0, y < 0,
and x and n are not relatively prime, z is unchanged and nil is returned.

Modular exponentation of inputs of a particular size is not a
cryptographically constant-time operation.




### <a id="Int.Format">func</a> (\*Int) [Format](https://golang.org/src/math/big/intconv.go?s=2003:2045#L57)
<pre>func (x *<a href="#Int">Int</a>) Format(s <a href="/pkg/fmt/">fmt</a>.<a href="/pkg/fmt/#State">State</a>, ch <a href="/pkg/builtin/#rune">rune</a>)</pre>
Format implements fmt.Formatter. It accepts the formats
'b' (binary), 'o' (octal with 0 prefix), 'O' (octal with 0o prefix),
'd' (decimal), 'x' (lowercase hexadecimal), and
'X' (uppercase hexadecimal).
Also supported are the full suite of package fmt's format
flags for integral types, including '+' and ' ' for sign
control, '#' for leading zero in octal and for hexadecimal,
a leading "0x" or "0X" for "%#x" and "%#X" respectively,
specification of minimum digits precision, output field
width, space or zero padding, and '-' for left or right
justification.




### <a id="Int.GCD">func</a> (\*Int) [GCD](https://golang.org/src/math/big/int.go?s=13198:13237#L497)
<pre>func (z *<a href="#Int">Int</a>) GCD(x, y, a, b *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
GCD sets z to the greatest common divisor of a and b, which both must
be > 0, and returns z.
If x or y are not nil, GCD sets their value such that z = a*x + b*y.
If either a or b is <= 0, GCD sets z = x = y = 0.




### <a id="Int.GobDecode">func</a> (\*Int) [GobDecode](https://golang.org/src/math/big/intmarsh.go?s=796:837#L23)
<pre>func (z *<a href="#Int">Int</a>) GobDecode(buf []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
GobDecode implements the gob.GobDecoder interface.




### <a id="Int.GobEncode">func</a> (\*Int) [GobEncode](https://golang.org/src/math/big/intmarsh.go?s=412:453#L8)
<pre>func (x *<a href="#Int">Int</a>) GobEncode() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
GobEncode implements the gob.GobEncoder interface.




### <a id="Int.Int64">func</a> (\*Int) [Int64](https://golang.org/src/math/big/int.go?s=8801:8828#L361)
<pre>func (x *<a href="#Int">Int</a>) Int64() <a href="/pkg/builtin/#int64">int64</a></pre>
Int64 returns the int64 representation of x.
If x cannot be represented in an int64, the result is undefined.




### <a id="Int.IsInt64">func</a> (\*Int) [IsInt64](https://golang.org/src/math/big/int.go?s=9129:9157#L376)
<pre>func (x *<a href="#Int">Int</a>) IsInt64() <a href="/pkg/builtin/#bool">bool</a></pre>
IsInt64 reports whether x can be represented as an int64.




### <a id="Int.IsUint64">func</a> (\*Int) [IsUint64](https://golang.org/src/math/big/int.go?s=9331:9360#L385)
<pre>func (x *<a href="#Int">Int</a>) IsUint64() <a href="/pkg/builtin/#bool">bool</a></pre>
IsUint64 reports whether x can be represented as a uint64.




### <a id="Int.Lsh">func</a> (\*Int) [Lsh](https://golang.org/src/math/big/int.go?s=24678:24716#L952)
<pre>func (z *<a href="#Int">Int</a>) Lsh(x *<a href="#Int">Int</a>, n <a href="/pkg/builtin/#uint">uint</a>) *<a href="#Int">Int</a></pre>
Lsh sets z = x << n and returns z.




### <a id="Int.MarshalJSON">func</a> (\*Int) [MarshalJSON](https://golang.org/src/math/big/intmarsh.go?s=1831:1874#L59)
<pre>func (x *<a href="#Int">Int</a>) MarshalJSON() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
MarshalJSON implements the json.Marshaler interface.




### <a id="Int.MarshalText">func</a> (\*Int) [MarshalText](https://golang.org/src/math/big/intmarsh.go?s=1186:1238#L39)
<pre>func (x *<a href="#Int">Int</a>) MarshalText() (text []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
MarshalText implements the encoding.TextMarshaler interface.




### <a id="Int.Mod">func</a> (\*Int) [Mod](https://golang.org/src/math/big/int.go?s=6693:6726#L254)
<pre>func (z *<a href="#Int">Int</a>) Mod(x, y *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
Mod sets z to the modulus x%y for y != 0 and returns z.
If y == 0, a division-by-zero run-time panic occurs.
Mod implements Euclidean modulus (unlike Go); see DivMod for more details.




### <a id="Int.ModInverse">func</a> (\*Int) [ModInverse](https://golang.org/src/math/big/int.go?s=19833:19873#L752)
<pre>func (z *<a href="#Int">Int</a>) ModInverse(g, n *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
ModInverse sets z to the multiplicative inverse of g in the ring ℤ/nℤ
and returns z. If g and n are not relatively prime, g has no multiplicative
inverse in the ring ℤ/nℤ.  In this case, z is unchanged and the return value
is nil.




### <a id="Int.ModSqrt">func</a> (\*Int) [ModSqrt](https://golang.org/src/math/big/int.go?s=24033:24070#L925)
<pre>func (z *<a href="#Int">Int</a>) ModSqrt(x, p *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
ModSqrt sets z to a square root of x mod p if such a square root exists, and
returns z. The modulus p must be an odd prime. If x is not a square mod p,
ModSqrt leaves z unchanged and returns nil. This function panics if p is
not an odd integer.




### <a id="Int.Mul">func</a> (\*Int) [Mul](https://golang.org/src/math/big/int.go?s=3668:3701#L148)
<pre>func (z *<a href="#Int">Int</a>) Mul(x, y *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
Mul sets z to the product x*y and returns z.




### <a id="Int.MulRange">func</a> (\*Int) [MulRange](https://golang.org/src/math/big/int.go?s=4117:4156#L166)
<pre>func (z *<a href="#Int">Int</a>) MulRange(a, b <a href="/pkg/builtin/#int64">int64</a>) *<a href="#Int">Int</a></pre>
MulRange sets z to the product of all integers
in the range [a, b] inclusively and returns z.
If a > b (empty range), the result is 1.




### <a id="Int.Neg">func</a> (\*Int) [Neg](https://golang.org/src/math/big/int.go?s=2566:2596#L99)
<pre>func (z *<a href="#Int">Int</a>) Neg(x *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
Neg sets z to -x and returns z.




### <a id="Int.Not">func</a> (\*Int) [Not](https://golang.org/src/math/big/int.go?s=29110:29140#L1139)
<pre>func (z *<a href="#Int">Int</a>) Not(x *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
Not sets z = ^x and returns z.




### <a id="Int.Or">func</a> (\*Int) [Or](https://golang.org/src/math/big/int.go?s=27663:27695#L1079)
<pre>func (z *<a href="#Int">Int</a>) Or(x, y *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
Or sets z = x | y and returns z.




### <a id="Int.ProbablyPrime">func</a> (\*Int) [ProbablyPrime](https://golang.org/src/math/big/prime.go?s=1085:1124#L16)
<pre>func (x *<a href="#Int">Int</a>) ProbablyPrime(n <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
ProbablyPrime reports whether x is probably prime,
applying the Miller-Rabin test with n pseudorandomly chosen bases
as well as a Baillie-PSW test.

If x is prime, ProbablyPrime returns true.
If x is chosen randomly and not prime, ProbablyPrime probably returns false.
The probability of returning true for a randomly chosen non-prime is at most ¼ⁿ.

ProbablyPrime is 100% accurate for inputs less than 2⁶⁴.
See Menezes et al., Handbook of Applied Cryptography, 1997, pp. 145-149,
and FIPS 186-4 Appendix F for further discussion of the error probabilities.

ProbablyPrime is not suitable for judging primes that an adversary may
have crafted to fool the test.

As of Go 1.8, ProbablyPrime(0) is allowed and applies only a Baillie-PSW test.
Before Go 1.8, ProbablyPrime applied only the Miller-Rabin tests, and ProbablyPrime(0) panicked.




### <a id="Int.Quo">func</a> (\*Int) [Quo](https://golang.org/src/math/big/int.go?s=4979:5012#L201)
<pre>func (z *<a href="#Int">Int</a>) Quo(x, y *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
Quo sets z to the quotient x/y for y != 0 and returns z.
If y == 0, a division-by-zero run-time panic occurs.
Quo implements truncated division (like Go); see QuoRem for more details.




### <a id="Int.QuoRem">func</a> (\*Int) [QuoRem](https://golang.org/src/math/big/int.go?s=5895:5942#L228)
<pre>func (z *<a href="#Int">Int</a>) QuoRem(x, y, r *<a href="#Int">Int</a>) (*<a href="#Int">Int</a>, *<a href="#Int">Int</a>)</pre>
QuoRem sets z to the quotient x/y and r to the remainder x%y
and returns the pair (z, r) for y != 0.
If y == 0, a division-by-zero run-time panic occurs.

QuoRem implements T-division and modulus (like Go):


	q = x/y      with the result truncated to zero
	r = x - y*q

(See Daan Leijen, ``Division and Modulus for Computer Scientists''.)
See DivMod for Euclidean division and modulus (unlike Go).




### <a id="Int.Rand">func</a> (\*Int) [Rand](https://golang.org/src/math/big/int.go?s=19395:19442#L738)
<pre>func (z *<a href="#Int">Int</a>) Rand(rnd *<a href="/pkg/math/rand/">rand</a>.<a href="/pkg/math/rand/#Rand">Rand</a>, n *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
Rand sets z to a pseudo-random number in [0, n) and returns z.

As this uses the math/rand package, it must not be used for
security-sensitive work. Use crypto/rand.Int instead.




### <a id="Int.Rem">func</a> (\*Int) [Rem](https://golang.org/src/math/big/int.go?s=5321:5354#L210)
<pre>func (z *<a href="#Int">Int</a>) Rem(x, y *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
Rem sets z to the remainder x%y for y != 0 and returns z.
If y == 0, a division-by-zero run-time panic occurs.
Rem implements truncated modulus (like Go); see QuoRem for more details.




### <a id="Int.Rsh">func</a> (\*Int) [Rsh](https://golang.org/src/math/big/int.go?s=24814:24852#L959)
<pre>func (z *<a href="#Int">Int</a>) Rsh(x *<a href="#Int">Int</a>, n <a href="/pkg/builtin/#uint">uint</a>) *<a href="#Int">Int</a></pre>
Rsh sets z = x >> n and returns z.




### <a id="Int.Scan">func</a> (\*Int) [Scan](https://golang.org/src/math/big/intconv.go?s=6459:6509#L228)
<pre>func (z *<a href="#Int">Int</a>) Scan(s <a href="/pkg/fmt/">fmt</a>.<a href="/pkg/fmt/#ScanState">ScanState</a>, ch <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#error">error</a></pre>
Scan is a support routine for fmt.Scanner; it sets z to the value of
the scanned number. It accepts the formats 'b' (binary), 'o' (octal),
'd' (decimal), 'x' (lowercase hexadecimal), and 'X' (uppercase hexadecimal).


<a id="example_Int_Scan">Example</a>


### <a id="Int.Set">func</a> (\*Int) [Set](https://golang.org/src/math/big/int.go?s=1488:1518#L63)
<pre>func (z *<a href="#Int">Int</a>) Set(x *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
Set sets z to x and returns z.




### <a id="Int.SetBit">func</a> (\*Int) [SetBit](https://golang.org/src/math/big/int.go?s=25824:25872#L999)
<pre>func (z *<a href="#Int">Int</a>) SetBit(x *<a href="#Int">Int</a>, i <a href="/pkg/builtin/#int">int</a>, b <a href="/pkg/builtin/#uint">uint</a>) *<a href="#Int">Int</a></pre>
SetBit sets z to x, with x's i'th bit set to b (0 or 1).
That is, if b is 1 SetBit sets z = x | (1 << i);
if b is 0 SetBit sets z = x &^ (1 << i). If b is not 0 or 1,
SetBit will panic.




### <a id="Int.SetBits">func</a> (\*Int) [SetBits](https://golang.org/src/math/big/int.go?s=2304:2342#L85)
<pre>func (z *<a href="#Int">Int</a>) SetBits(abs []<a href="#Word">Word</a>) *<a href="#Int">Int</a></pre>
SetBits provides raw (unchecked but fast) access to z by setting its
value to abs, interpreted as a little-endian Word slice, and returning
z. The result and abs share the same underlying array.
SetBits is intended to support implementation of missing low-level Int
functionality outside this package; it should be avoided otherwise.




### <a id="Int.SetBytes">func</a> (\*Int) [SetBytes](https://golang.org/src/math/big/int.go?s=11309:11348#L431)
<pre>func (z *<a href="#Int">Int</a>) SetBytes(buf []<a href="/pkg/builtin/#byte">byte</a>) *<a href="#Int">Int</a></pre>
SetBytes interprets buf as the bytes of a big-endian unsigned
integer, sets z to that value, and returns z.




### <a id="Int.SetInt64">func</a> (\*Int) [SetInt64](https://golang.org/src/math/big/int.go?s=1053:1089#L39)
<pre>func (z *<a href="#Int">Int</a>) SetInt64(x <a href="/pkg/builtin/#int64">int64</a>) *<a href="#Int">Int</a></pre>
SetInt64 sets z to x and returns z.




### <a id="Int.SetString">func</a> (\*Int) [SetString](https://golang.org/src/math/big/int.go?s=10645:10701#L412)
<pre>func (z *<a href="#Int">Int</a>) SetString(s <a href="/pkg/builtin/#string">string</a>, base <a href="/pkg/builtin/#int">int</a>) (*<a href="#Int">Int</a>, <a href="/pkg/builtin/#bool">bool</a>)</pre>
SetString sets z to the value of s, interpreted in the given base,
and returns z and a boolean indicating success. The entire string
(not just a prefix) must be valid for success. If SetString fails,
the value of z is undefined but the returned value is nil.

The base argument must be 0 or a value between 2 and MaxBase.
For base 0, the number prefix determines the actual base: A prefix of
``0b'' or ``0B'' selects base 2, ``0'', ``0o'' or ``0O'' selects base 8,
and ``0x'' or ``0X'' selects base 16. Otherwise, the selected base is 10
and no prefix is accepted.

For bases <= 36, lower and upper case letters are considered the same:
The letters 'a' to 'z' and 'A' to 'Z' represent digit values 10 to 35.
For bases > 36, the upper case letters 'A' to 'Z' represent the digit
values 36 to 61.

For base 0, an underscore character ``_'' may appear between a base
prefix and an adjacent digit, and between successive digits; such
underscores do not change the value of the number.
Incorrect placement of underscores is reported as an error if there
are no other errors. If base != 0, underscores are not recognized
and act like any other character that is not a valid digit.


<a id="example_Int_SetString">Example</a>


### <a id="Int.SetUint64">func</a> (\*Int) [SetUint64](https://golang.org/src/math/big/int.go?s=1245:1283#L51)
<pre>func (z *<a href="#Int">Int</a>) SetUint64(x <a href="/pkg/builtin/#uint64">uint64</a>) *<a href="#Int">Int</a></pre>
SetUint64 sets z to x and returns z.




### <a id="Int.Sign">func</a> (\*Int) [Sign](https://golang.org/src/math/big/int.go?s=911:935#L28)
<pre>func (x *<a href="#Int">Int</a>) Sign() <a href="/pkg/builtin/#int">int</a></pre>
Sign returns:


	-1 if x <  0
	 0 if x == 0
	+1 if x >  0




### <a id="Int.Sqrt">func</a> (\*Int) [Sqrt](https://golang.org/src/math/big/int.go?s=29492:29523#L1155)
<pre>func (z *<a href="#Int">Int</a>) Sqrt(x *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
Sqrt sets z to ⌊√x⌋, the largest integer such that z² ≤ x, and returns z.
It panics if x is negative.




### <a id="Int.String">func</a> (\*Int) [String](https://golang.org/src/math/big/intconv.go?s=1099:1128#L29)
<pre>func (x *<a href="#Int">Int</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns the decimal representation of x as generated by
x.Text(10).




### <a id="Int.Sub">func</a> (\*Int) [Sub](https://golang.org/src/math/big/int.go?s=3194:3227#L127)
<pre>func (z *<a href="#Int">Int</a>) Sub(x, y *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
Sub sets z to the difference x-y and returns z.




### <a id="Int.Text">func</a> (\*Int) [Text](https://golang.org/src/math/big/intconv.go?s=625:660#L11)
<pre>func (x *<a href="#Int">Int</a>) Text(base <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#string">string</a></pre>
Text returns the string representation of x in the given base.
Base must be between 2 and 62, inclusive. The result uses the
lower-case letters 'a' to 'z' for digit values 10 to 35, and
the upper-case letters 'A' to 'Z' for digit values 36 to 61.
No prefix (such as "0x") is added to the string. If x is a nil
pointer it returns "<nil>".




### <a id="Int.TrailingZeroBits">func</a> (\*Int) [TrailingZeroBits](https://golang.org/src/math/big/int.go?s=11820:11857#L451)
<pre>func (x *<a href="#Int">Int</a>) TrailingZeroBits() <a href="/pkg/builtin/#uint">uint</a></pre>
TrailingZeroBits returns the number of consecutive least significant zero
bits of |x|.




### <a id="Int.Uint64">func</a> (\*Int) [Uint64](https://golang.org/src/math/big/int.go?s=9012:9041#L371)
<pre>func (x *<a href="#Int">Int</a>) Uint64() <a href="/pkg/builtin/#uint64">uint64</a></pre>
Uint64 returns the uint64 representation of x.
If x cannot be represented in a uint64, the result is undefined.




### <a id="Int.UnmarshalJSON">func</a> (\*Int) [UnmarshalJSON](https://golang.org/src/math/big/intmarsh.go?s=1964:2010#L64)
<pre>func (z *<a href="#Int">Int</a>) UnmarshalJSON(text []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
UnmarshalJSON implements the json.Unmarshaler interface.




### <a id="Int.UnmarshalText">func</a> (\*Int) [UnmarshalText](https://golang.org/src/math/big/intmarsh.go?s=1395:1441#L47)
<pre>func (z *<a href="#Int">Int</a>) UnmarshalText(text []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
UnmarshalText implements the encoding.TextUnmarshaler interface.




### <a id="Int.Xor">func</a> (\*Int) [Xor](https://golang.org/src/math/big/int.go?s=28433:28466#L1109)
<pre>func (z *<a href="#Int">Int</a>) Xor(x, y *<a href="#Int">Int</a>) *<a href="#Int">Int</a></pre>
Xor sets z = x ^ y and returns z.




## <a id="Rat">type</a> [Rat](https://golang.org/src/math/big/rat.go?s=705:905#L13)
A Rat represents a quotient a/b of arbitrary precision.
The zero value for a Rat represents the value 0.

Operations always take pointer arguments (*Rat) rather
than Rat values, and each unique Rat value requires
its own unique *Rat pointer. To "copy" a Rat value,
an existing (or newly allocated) Rat must be set to
a new value using the Rat.Set method; shallow copies
of Rats are not supported and may lead to errors.


<pre>type Rat struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewRat">func</a> [NewRat](https://golang.org/src/math/big/rat.go?s=971:999#L21)
<pre>func NewRat(a, b <a href="/pkg/builtin/#int64">int64</a>) *<a href="#Rat">Rat</a></pre>
NewRat creates a new Rat with numerator a and denominator b.






### <a id="Rat.Abs">func</a> (\*Rat) [Abs](https://golang.org/src/math/big/rat.go?s=9365:9395#L349)
<pre>func (z *<a href="#Rat">Rat</a>) Abs(x *<a href="#Rat">Rat</a>) *<a href="#Rat">Rat</a></pre>
Abs sets z to |x| (the absolute value of x) and returns z.




### <a id="Rat.Add">func</a> (\*Rat) [Add](https://golang.org/src/math/big/rat.go?s=12293:12326#L480)
<pre>func (z *<a href="#Rat">Rat</a>) Add(x, y *<a href="#Rat">Rat</a>) *<a href="#Rat">Rat</a></pre>
Add sets z to the sum x+y and returns z.




### <a id="Rat.Cmp">func</a> (\*Rat) [Cmp](https://golang.org/src/math/big/rat.go?s=12124:12153#L472)
<pre>func (x *<a href="#Rat">Rat</a>) Cmp(y *<a href="#Rat">Rat</a>) <a href="/pkg/builtin/#int">int</a></pre>
Cmp compares x and y and returns:


	-1 if x <  y
	 0 if x == y
	+1 if x >  y




### <a id="Rat.Denom">func</a> (\*Rat) [Denom](https://golang.org/src/math/big/rat.go?s=10636:10662#L406)
<pre>func (x *<a href="#Rat">Rat</a>) Denom() *<a href="#Int">Int</a></pre>
Denom returns the denominator of x; it is always > 0.
The result is a reference to x's denominator; it
may change if a new value is assigned to x, and vice versa.




### <a id="Rat.Float32">func</a> (\*Rat) [Float32](https://golang.org/src/math/big/rat.go?s=7459:7506#L261)
<pre>func (x *<a href="#Rat">Rat</a>) Float32() (f <a href="/pkg/builtin/#float32">float32</a>, exact <a href="/pkg/builtin/#bool">bool</a>)</pre>
Float32 returns the nearest float32 value for x and a bool indicating
whether f represents x exactly. If the magnitude of x is too large to
be represented by a float32, f is an infinity and exact is false.
The sign of f always matches the sign of x, even if f == 0.




### <a id="Rat.Float64">func</a> (\*Rat) [Float64](https://golang.org/src/math/big/rat.go?s=7944:7991#L277)
<pre>func (x *<a href="#Rat">Rat</a>) Float64() (f <a href="/pkg/builtin/#float64">float64</a>, exact <a href="/pkg/builtin/#bool">bool</a>)</pre>
Float64 returns the nearest float64 value for x and a bool indicating
whether f represents x exactly. If the magnitude of x is too large to
be represented by a float64, f is an infinity and exact is false.
The sign of f always matches the sign of x, even if f == 0.




### <a id="Rat.FloatString">func</a> (\*Rat) [FloatString](https://golang.org/src/math/big/ratconv.go?s=8783:8825#L314)
<pre>func (x *<a href="#Rat">Rat</a>) FloatString(prec <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#string">string</a></pre>
FloatString returns a string representation of x in decimal form with prec
digits of precision after the radix point. The last digit is rounded to
nearest, with halves rounded away from zero.




### <a id="Rat.GobDecode">func</a> (\*Rat) [GobDecode](https://golang.org/src/math/big/ratmarsh.go?s=1058:1099#L32)
<pre>func (z *<a href="#Rat">Rat</a>) GobDecode(buf []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
GobDecode implements the gob.GobDecoder interface.




### <a id="Rat.GobEncode">func</a> (\*Rat) [GobEncode](https://golang.org/src/math/big/ratmarsh.go?s=432:473#L9)
<pre>func (x *<a href="#Rat">Rat</a>) GobEncode() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
GobEncode implements the gob.GobEncoder interface.




### <a id="Rat.Inv">func</a> (\*Rat) [Inv](https://golang.org/src/math/big/rat.go?s=9622:9652#L363)
<pre>func (z *<a href="#Rat">Rat</a>) Inv(x *<a href="#Rat">Rat</a>) *<a href="#Rat">Rat</a></pre>
Inv sets z to 1/x and returns z.




### <a id="Rat.IsInt">func</a> (\*Rat) [IsInt](https://golang.org/src/math/big/rat.go?s=10112:10138#L391)
<pre>func (x *<a href="#Rat">Rat</a>) IsInt() <a href="/pkg/builtin/#bool">bool</a></pre>
IsInt reports whether the denominator of x is 1.




### <a id="Rat.MarshalText">func</a> (\*Rat) [MarshalText](https://golang.org/src/math/big/ratmarsh.go?s=1555:1607#L51)
<pre>func (x *<a href="#Rat">Rat</a>) MarshalText() (text []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
MarshalText implements the encoding.TextMarshaler interface.




### <a id="Rat.Mul">func</a> (\*Rat) [Mul](https://golang.org/src/math/big/rat.go?s=12788:12821#L500)
<pre>func (z *<a href="#Rat">Rat</a>) Mul(x, y *<a href="#Rat">Rat</a>) *<a href="#Rat">Rat</a></pre>
Mul sets z to the product x*y and returns z.




### <a id="Rat.Neg">func</a> (\*Rat) [Neg](https://golang.org/src/math/big/rat.go?s=9473:9503#L356)
<pre>func (z *<a href="#Rat">Rat</a>) Neg(x *<a href="#Rat">Rat</a>) *<a href="#Rat">Rat</a></pre>
Neg sets z to -x and returns z.




### <a id="Rat.Num">func</a> (\*Rat) [Num](https://golang.org/src/math/big/rat.go?s=10421:10445#L399)
<pre>func (x *<a href="#Rat">Rat</a>) Num() *<a href="#Int">Int</a></pre>
Num returns the numerator of x; it may be <= 0.
The result is a reference to x's numerator; it
may change if a new value is assigned to x, and vice versa.
The sign of the numerator corresponds to the sign of x.




### <a id="Rat.Quo">func</a> (\*Rat) [Quo](https://golang.org/src/math/big/rat.go?s=13180:13213#L515)
<pre>func (z *<a href="#Rat">Rat</a>) Quo(x, y *<a href="#Rat">Rat</a>) *<a href="#Rat">Rat</a></pre>
Quo sets z to the quotient x/y and returns z.
If y == 0, a division-by-zero run-time panic occurs.




### <a id="Rat.RatString">func</a> (\*Rat) [RatString](https://golang.org/src/math/big/ratconv.go?s=8484:8516#L304)
<pre>func (x *<a href="#Rat">Rat</a>) RatString() <a href="/pkg/builtin/#string">string</a></pre>
RatString returns a string representation of x in the form "a/b" if b != 1,
and in the form "a" if b == 1.




### <a id="Rat.Scan">func</a> (\*Rat) [Scan](https://golang.org/src/math/big/ratconv.go?s=595:645#L16)
<pre>func (z *<a href="#Rat">Rat</a>) Scan(s <a href="/pkg/fmt/">fmt</a>.<a href="/pkg/fmt/#ScanState">ScanState</a>, ch <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#error">error</a></pre>
Scan is a support routine for fmt.Scanner. It accepts the formats
'e', 'E', 'f', 'F', 'g', 'G', and 'v'. All formats are equivalent.


<a id="example_Rat_Scan">Example</a>


### <a id="Rat.Set">func</a> (\*Rat) [Set](https://golang.org/src/math/big/rat.go?s=9209:9239#L340)
<pre>func (z *<a href="#Rat">Rat</a>) Set(x *<a href="#Rat">Rat</a>) *<a href="#Rat">Rat</a></pre>
Set sets z to x (by making a copy of x) and returns z.




### <a id="Rat.SetFloat64">func</a> (\*Rat) [SetFloat64](https://golang.org/src/math/big/rat.go?s=1132:1172#L27)
<pre>func (z *<a href="#Rat">Rat</a>) SetFloat64(f <a href="/pkg/builtin/#float64">float64</a>) *<a href="#Rat">Rat</a></pre>
SetFloat64 sets z to exactly f and returns z.
If f is not finite, SetFloat returns nil.




### <a id="Rat.SetFrac">func</a> (\*Rat) [SetFrac](https://golang.org/src/math/big/rat.go?s=8191:8228#L290)
<pre>func (z *<a href="#Rat">Rat</a>) SetFrac(a, b *<a href="#Int">Int</a>) *<a href="#Rat">Rat</a></pre>
SetFrac sets z to a/b and returns z.




### <a id="Rat.SetFrac64">func</a> (\*Rat) [SetFrac64](https://golang.org/src/math/big/rat.go?s=8531:8571#L305)
<pre>func (z *<a href="#Rat">Rat</a>) SetFrac64(a, b <a href="/pkg/builtin/#int64">int64</a>) *<a href="#Rat">Rat</a></pre>
SetFrac64 sets z to a/b and returns z.




### <a id="Rat.SetInt">func</a> (\*Rat) [SetInt](https://golang.org/src/math/big/rat.go?s=8801:8834#L319)
<pre>func (z *<a href="#Rat">Rat</a>) SetInt(x *<a href="#Int">Int</a>) *<a href="#Rat">Rat</a></pre>
SetInt sets z to x (by making a copy of x) and returns z.




### <a id="Rat.SetInt64">func</a> (\*Rat) [SetInt64](https://golang.org/src/math/big/rat.go?s=8924:8960#L326)
<pre>func (z *<a href="#Rat">Rat</a>) SetInt64(x <a href="/pkg/builtin/#int64">int64</a>) *<a href="#Rat">Rat</a></pre>
SetInt64 sets z to x and returns z.




### <a id="Rat.SetString">func</a> (\*Rat) [SetString](https://golang.org/src/math/big/ratconv.go?s=2124:2170#L47)
<pre>func (z *<a href="#Rat">Rat</a>) SetString(s <a href="/pkg/builtin/#string">string</a>) (*<a href="#Rat">Rat</a>, <a href="/pkg/builtin/#bool">bool</a>)</pre>
SetString sets z to the value of s and returns z and a boolean indicating
success. s can be given as a (possibly signed) fraction "a/b", or as a
floating-point number optionally followed by an exponent.
If a fraction is provided, both the dividend and the divisor may be a
decimal integer or independently use a prefix of ``0b'', ``0'' or ``0o'',
or ``0x'' (or their upper-case variants) to denote a binary, octal, or
hexadecimal integer, respectively. The divisor may not be signed.
If a floating-point number is provided, it may be in decimal form or
use any of the same prefixes as above but for ``0'' to denote a non-decimal
mantissa. A leading ``0'' is considered a decimal leading 0; it does not
indicate octal representation in this case.
An optional base-10 ``e'' or base-2 ``p'' (or their upper-case variants)
exponent may be provided as well, except for hexadecimal floats which
only accept an (optional) ``p'' exponent (because an ``e'' or ``E'' cannot
be distinguished from a mantissa digit).
The entire string, not just a prefix, must be valid for success. If the
operation failed, the value of z is undefined but the returned value is nil.


<a id="example_Rat_SetString">Example</a>


### <a id="Rat.SetUint64">func</a> (\*Rat) [SetUint64](https://golang.org/src/math/big/rat.go?s=9056:9094#L333)
<pre>func (z *<a href="#Rat">Rat</a>) SetUint64(x <a href="/pkg/builtin/#uint64">uint64</a>) *<a href="#Rat">Rat</a></pre>
SetUint64 sets z to x and returns z.




### <a id="Rat.Sign">func</a> (\*Rat) [Sign](https://golang.org/src/math/big/rat.go?s=10011:10035#L386)
<pre>func (x *<a href="#Rat">Rat</a>) Sign() <a href="/pkg/builtin/#int">int</a></pre>
Sign returns:


	-1 if x <  0
	 0 if x == 0
	+1 if x >  0




### <a id="Rat.String">func</a> (\*Rat) [String](https://golang.org/src/math/big/ratconv.go?s=8047:8076#L285)
<pre>func (x *<a href="#Rat">Rat</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns a string representation of x in the form "a/b" (even if b == 1).




### <a id="Rat.Sub">func</a> (\*Rat) [Sub](https://golang.org/src/math/big/rat.go?s=12542:12575#L490)
<pre>func (z *<a href="#Rat">Rat</a>) Sub(x, y *<a href="#Rat">Rat</a>) *<a href="#Rat">Rat</a></pre>
Sub sets z to the difference x-y and returns z.




### <a id="Rat.UnmarshalText">func</a> (\*Rat) [UnmarshalText](https://golang.org/src/math/big/ratmarsh.go?s=1752:1798#L59)
<pre>func (z *<a href="#Rat">Rat</a>) UnmarshalText(text []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
UnmarshalText implements the encoding.TextUnmarshaler interface.




## <a id="RoundingMode">type</a> [RoundingMode](https://golang.org/src/math/big/float.go?s=5178:5200#L122)
RoundingMode determines how a Float value is rounded to the
desired precision. Rounding may change the Float value; the
rounding error is described by the Float's Accuracy.


<pre>type RoundingMode <a href="/pkg/builtin/#byte">byte</a></pre>


These constants define supported rounding modes.


<pre>const (
    <span id="ToNearestEven">ToNearestEven</span> <a href="#RoundingMode">RoundingMode</a> = <a href="/pkg/builtin/#iota">iota</a> <span class="comment">// == IEEE 754-2008 roundTiesToEven</span>
    <span id="ToNearestAway">ToNearestAway</span>                     <span class="comment">// == IEEE 754-2008 roundTiesToAway</span>
    <span id="ToZero">ToZero</span>                            <span class="comment">// == IEEE 754-2008 roundTowardZero</span>
    <span id="AwayFromZero">AwayFromZero</span>                      <span class="comment">// no IEEE 754-2008 equivalent</span>
    <span id="ToNegativeInf">ToNegativeInf</span>                     <span class="comment">// == IEEE 754-2008 roundTowardNegative</span>
    <span id="ToPositiveInf">ToPositiveInf</span>                     <span class="comment">// == IEEE 754-2008 roundTowardPositive</span>
)</pre>



<a id="example_RoundingMode">Example</a>






### <a id="RoundingMode.String">func</a> (RoundingMode) [String](https://golang.org/src/math/big/roundingmode_string.go?s=263:300#L1)
<pre>func (i <a href="#RoundingMode">RoundingMode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Word">type</a> [Word](https://golang.org/src/math/big/arith.go?s=602:616#L6)
A Word represents a single digit of a multi-precision unsigned integer.


<pre>type Word <a href="/pkg/builtin/#uint">uint</a></pre>















