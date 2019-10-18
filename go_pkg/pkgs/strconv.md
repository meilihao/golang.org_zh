

# strconv
`import "strconv"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package strconv implements conversions to and from string representations
of basic data types.

### Numeric Conversions
The most common numeric conversions are Atoi (string to int) and Itoa (int to string).


	i, err := strconv.Atoi("-42")
	s := strconv.Itoa(-42)

These assume decimal and the Go int type.

ParseBool, ParseFloat, ParseInt, and ParseUint convert strings to values:


	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	i, err := strconv.ParseInt("-42", 10, 64)
	u, err := strconv.ParseUint("42", 10, 64)

The parse functions return the widest type (float64, int64, and uint64),
but if the size argument specifies a narrower width the result can be
converted to that narrower type without data loss:


	s := "2147483647" // biggest int32
	i64, err := strconv.ParseInt(s, 10, 32)
	...
	i := int32(i64)

FormatBool, FormatFloat, FormatInt, and FormatUint convert values to strings:


	s := strconv.FormatBool(true)
	s := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s := strconv.FormatInt(-42, 16)
	s := strconv.FormatUint(42, 16)

AppendBool, AppendFloat, AppendInt, and AppendUint are similar but
append the formatted value to a destination slice.

### String Conversions
Quote and QuoteToASCII convert strings to quoted Go string literals.
The latter guarantees that the result is an ASCII string, by escaping
any non-ASCII Unicode with \u:


	q := strconv.Quote("Hello, 世界")
	q := strconv.QuoteToASCII("Hello, 世界")

QuoteRune and QuoteRuneToASCII are similar but accept runes and
return quoted Go rune literals.

Unquote and UnquoteChar unquote Go string and rune literals.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func AppendBool(dst []byte, b bool) []byte](#AppendBool)
* [func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte](#AppendFloat)
* [func AppendInt(dst []byte, i int64, base int) []byte](#AppendInt)
* [func AppendQuote(dst []byte, s string) []byte](#AppendQuote)
* [func AppendQuoteRune(dst []byte, r rune) []byte](#AppendQuoteRune)
* [func AppendQuoteRuneToASCII(dst []byte, r rune) []byte](#AppendQuoteRuneToASCII)
* [func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte](#AppendQuoteRuneToGraphic)
* [func AppendQuoteToASCII(dst []byte, s string) []byte](#AppendQuoteToASCII)
* [func AppendQuoteToGraphic(dst []byte, s string) []byte](#AppendQuoteToGraphic)
* [func AppendUint(dst []byte, i uint64, base int) []byte](#AppendUint)
* [func Atoi(s string) (int, error)](#Atoi)
* [func CanBackquote(s string) bool](#CanBackquote)
* [func FormatBool(b bool) string](#FormatBool)
* [func FormatFloat(f float64, fmt byte, prec, bitSize int) string](#FormatFloat)
* [func FormatInt(i int64, base int) string](#FormatInt)
* [func FormatUint(i uint64, base int) string](#FormatUint)
* [func IsGraphic(r rune) bool](#IsGraphic)
* [func IsPrint(r rune) bool](#IsPrint)
* [func Itoa(i int) string](#Itoa)
* [func ParseBool(str string) (bool, error)](#ParseBool)
* [func ParseFloat(s string, bitSize int) (float64, error)](#ParseFloat)
* [func ParseInt(s string, base int, bitSize int) (i int64, err error)](#ParseInt)
* [func ParseUint(s string, base int, bitSize int) (uint64, error)](#ParseUint)
* [func Quote(s string) string](#Quote)
* [func QuoteRune(r rune) string](#QuoteRune)
* [func QuoteRuneToASCII(r rune) string](#QuoteRuneToASCII)
* [func QuoteRuneToGraphic(r rune) string](#QuoteRuneToGraphic)
* [func QuoteToASCII(s string) string](#QuoteToASCII)
* [func QuoteToGraphic(s string) string](#QuoteToGraphic)
* [func Unquote(s string) (string, error)](#Unquote)
* [func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)](#UnquoteChar)
* [type NumError](#NumError)
  * [func (e *NumError) Error() string](#NumError.Error)


#### <a id="pkg-examples">Examples</a>
* [AppendBool](#example_AppendBool)
* [AppendFloat](#example_AppendFloat)
* [AppendInt](#example_AppendInt)
* [AppendQuote](#example_AppendQuote)
* [AppendQuoteRune](#example_AppendQuoteRune)
* [AppendQuoteRuneToASCII](#example_AppendQuoteRuneToASCII)
* [AppendQuoteToASCII](#example_AppendQuoteToASCII)
* [AppendUint](#example_AppendUint)
* [Atoi](#example_Atoi)
* [CanBackquote](#example_CanBackquote)
* [FormatBool](#example_FormatBool)
* [FormatFloat](#example_FormatFloat)
* [FormatInt](#example_FormatInt)
* [FormatUint](#example_FormatUint)
* [IsGraphic](#example_IsGraphic)
* [IsPrint](#example_IsPrint)
* [Itoa](#example_Itoa)
* [NumError](#example_NumError)
* [ParseBool](#example_ParseBool)
* [ParseFloat](#example_ParseFloat)
* [ParseInt](#example_ParseInt)
* [ParseUint](#example_ParseUint)
* [Quote](#example_Quote)
* [QuoteRune](#example_QuoteRune)
* [QuoteRuneToASCII](#example_QuoteRuneToASCII)
* [QuoteRuneToGraphic](#example_QuoteRuneToGraphic)
* [QuoteToASCII](#example_QuoteToASCII)
* [QuoteToGraphic](#example_QuoteToGraphic)
* [Unquote](#example_Unquote)
* [UnquoteChar](#example_UnquoteChar)


#### <a id="pkg-files">Package files</a>
[atob.go](https://golang.org/src/strconv/atob.go) [atof.go](https://golang.org/src/strconv/atof.go) [atoi.go](https://golang.org/src/strconv/atoi.go) [decimal.go](https://golang.org/src/strconv/decimal.go) [doc.go](https://golang.org/src/strconv/doc.go) [extfloat.go](https://golang.org/src/strconv/extfloat.go) [ftoa.go](https://golang.org/src/strconv/ftoa.go) [isprint.go](https://golang.org/src/strconv/isprint.go) [itoa.go](https://golang.org/src/strconv/itoa.go) [quote.go](https://golang.org/src/strconv/quote.go) 


## <a id="pkg-constants">Constants</a>
IntSize is the size in bits of an int or uint value.


<pre>const <span id="IntSize">IntSize</span> = intSize</pre>

## <a id="pkg-variables">Variables</a>
ErrRange indicates that a value is out of range for the target type.


<pre>var <span id="ErrRange">ErrRange</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;value out of range&#34;)</pre>ErrSyntax indicates that a value does not have the right syntax for the target type.


<pre>var <span id="ErrSyntax">ErrSyntax</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;invalid syntax&#34;)</pre>

## <a id="AppendBool">func</a> [AppendBool](https://golang.org/src/strconv/atob.go?s=852:894#L20)
<pre>func AppendBool(dst []<a href="/pkg/builtin/#byte">byte</a>, b <a href="/pkg/builtin/#bool">bool</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
AppendBool appends "true" or "false", according to the value of b,
to dst and returns the extended buffer.


<a id="example_AppendBool">Example</a>

## <a id="AppendFloat">func</a> [AppendFloat](https://golang.org/src/strconv/ftoa.go?s=1994:2069#L43)
<pre>func AppendFloat(dst []<a href="/pkg/builtin/#byte">byte</a>, f <a href="/pkg/builtin/#float64">float64</a>, fmt <a href="/pkg/builtin/#byte">byte</a>, prec, bitSize <a href="/pkg/builtin/#int">int</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
AppendFloat appends the string form of the floating-point number f,
as generated by FormatFloat, to dst and returns the extended buffer.


<a id="example_AppendFloat">Example</a>

## <a id="AppendInt">func</a> [AppendInt](https://golang.org/src/strconv/itoa.go?s=1214:1266#L30)
<pre>func AppendInt(dst []<a href="/pkg/builtin/#byte">byte</a>, i <a href="/pkg/builtin/#int64">int64</a>, base <a href="/pkg/builtin/#int">int</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
AppendInt appends the string form of the integer i,
as generated by FormatInt, to dst and returns the extended buffer.


<a id="example_AppendInt">Example</a>

## <a id="AppendQuote">func</a> [AppendQuote](https://golang.org/src/strconv/quote.go?s=3473:3518#L120)
<pre>func AppendQuote(dst []<a href="/pkg/builtin/#byte">byte</a>, s <a href="/pkg/builtin/#string">string</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
AppendQuote appends a double-quoted Go string literal representing s,
as generated by Quote, to dst and returns the extended buffer.


<a id="example_AppendQuote">Example</a>

## <a id="AppendQuoteRune">func</a> [AppendQuoteRune](https://golang.org/src/strconv/quote.go?s=5187:5234#L159)
<pre>func AppendQuoteRune(dst []<a href="/pkg/builtin/#byte">byte</a>, r <a href="/pkg/builtin/#rune">rune</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
AppendQuoteRune appends a single-quoted Go character literal representing the rune,
as generated by QuoteRune, to dst and returns the extended buffer.


<a id="example_AppendQuoteRune">Example</a>

## <a id="AppendQuoteRuneToASCII">func</a> [AppendQuoteRuneToASCII](https://golang.org/src/strconv/quote.go?s=5796:5850#L173)
<pre>func AppendQuoteRuneToASCII(dst []<a href="/pkg/builtin/#byte">byte</a>, r <a href="/pkg/builtin/#rune">rune</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
AppendQuoteRuneToASCII appends a single-quoted Go character literal representing the rune,
as generated by QuoteRuneToASCII, to dst and returns the extended buffer.


<a id="example_AppendQuoteRuneToASCII">Example</a>

## <a id="AppendQuoteRuneToGraphic">func</a> [AppendQuoteRuneToGraphic](https://golang.org/src/strconv/quote.go?s=6421:6477#L187)
<pre>func AppendQuoteRuneToGraphic(dst []<a href="/pkg/builtin/#byte">byte</a>, r <a href="/pkg/builtin/#rune">rune</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
AppendQuoteRuneToGraphic appends a single-quoted Go character literal representing the rune,
as generated by QuoteRuneToGraphic, to dst and returns the extended buffer.



## <a id="AppendQuoteToASCII">func</a> [AppendQuoteToASCII](https://golang.org/src/strconv/quote.go?s=4033:4085#L133)
<pre>func AppendQuoteToASCII(dst []<a href="/pkg/builtin/#byte">byte</a>, s <a href="/pkg/builtin/#string">string</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
AppendQuoteToASCII appends a double-quoted Go string literal representing s,
as generated by QuoteToASCII, to dst and returns the extended buffer.


<a id="example_AppendQuoteToASCII">Example</a>

## <a id="AppendQuoteToGraphic">func</a> [AppendQuoteToGraphic](https://golang.org/src/strconv/quote.go?s=4609:4663#L146)
<pre>func AppendQuoteToGraphic(dst []<a href="/pkg/builtin/#byte">byte</a>, s <a href="/pkg/builtin/#string">string</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
AppendQuoteToGraphic appends a double-quoted Go string literal representing s,
as generated by QuoteToGraphic, to dst and returns the extended buffer.



## <a id="AppendUint">func</a> [AppendUint](https://golang.org/src/strconv/itoa.go?s=1574:1628#L40)
<pre>func AppendUint(dst []<a href="/pkg/builtin/#byte">byte</a>, i <a href="/pkg/builtin/#uint64">uint64</a>, base <a href="/pkg/builtin/#int">int</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
AppendUint appends the string form of the unsigned integer i,
as generated by FormatUint, to dst and returns the extended buffer.


<a id="example_AppendUint">Example</a>

## <a id="Atoi">func</a> [Atoi](https://golang.org/src/strconv/atoi.go?s=5576:5608#L208)
<pre>func Atoi(s <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.


<a id="example_Atoi">Example</a>

## <a id="CanBackquote">func</a> [CanBackquote](https://golang.org/src/strconv/quote.go?s=6697:6729#L194)
<pre>func CanBackquote(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
CanBackquote reports whether the string s can be represented
unchanged as a single-line backquoted string without control
characters other than tab.


<a id="example_CanBackquote">Example</a>

## <a id="FormatBool">func</a> [FormatBool](https://golang.org/src/strconv/atob.go?s=660:690#L11)
<pre>func FormatBool(b <a href="/pkg/builtin/#bool">bool</a>) <a href="/pkg/builtin/#string">string</a></pre>
FormatBool returns "true" or "false" according to the value of b.


<a id="example_FormatBool">Example</a>

## <a id="FormatFloat">func</a> [FormatFloat](https://golang.org/src/strconv/ftoa.go?s=1697:1760#L37)
<pre>func FormatFloat(f <a href="/pkg/builtin/#float64">float64</a>, fmt <a href="/pkg/builtin/#byte">byte</a>, prec, bitSize <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#string">string</a></pre>
FormatFloat converts the floating-point number f to a string,
according to the format fmt and precision prec. It rounds the
result assuming that the original was obtained from a floating-point
value of bitSize bits (32 for float32, 64 for float64).

The format fmt is one of
'b' (-ddddp±ddd, a binary exponent),
'e' (-d.dddde±dd, a decimal exponent),
'E' (-d.ddddE±dd, a decimal exponent),
'f' (-ddd.dddd, no exponent),
'g' ('e' for large exponents, 'f' otherwise),
'G' ('E' for large exponents, 'f' otherwise),
'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).

The precision prec controls the number of digits (excluding the exponent)
printed by the 'e', 'E', 'f', 'g', 'G', 'x', and 'X' formats.
For 'e', 'E', 'f', 'x', and 'X', it is the number of digits after the decimal point.
For 'g' and 'G' it is the maximum number of significant digits (trailing
zeros are removed).
The special precision -1 uses the smallest number of digits
necessary such that ParseFloat will return f exactly.


<a id="example_FormatFloat">Example</a>

## <a id="FormatInt">func</a> [FormatInt](https://golang.org/src/strconv/itoa.go?s=784:824#L15)
<pre>func FormatInt(i <a href="/pkg/builtin/#int64">int64</a>, base <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#string">string</a></pre>
FormatInt returns the string representation of i in the given base,
for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
for digit values >= 10.


<a id="example_FormatInt">Example</a>

## <a id="FormatUint">func</a> [FormatUint](https://golang.org/src/strconv/itoa.go?s=434:476#L4)
<pre>func FormatUint(i <a href="/pkg/builtin/#uint64">uint64</a>, base <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#string">string</a></pre>
FormatUint returns the string representation of i in the given base,
for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
for digit values >= 10.


<a id="example_FormatUint">Example</a>

## <a id="IsGraphic">func</a> [IsGraphic](https://golang.org/src/strconv/quote.go?s=14365:14392#L518)
<pre>func IsGraphic(r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsGraphic reports whether the rune is defined as a Graphic by Unicode. Such
characters include letters, marks, numbers, punctuation, symbols, and
spaces, from categories L, M, N, P, S, and Zs.


<a id="example_IsGraphic">Example</a>

## <a id="IsPrint">func</a> [IsPrint](https://golang.org/src/strconv/quote.go?s=12881:12906#L472)
<pre>func IsPrint(r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsPrint reports whether the rune is defined as printable by Go, with
the same definition as unicode.IsPrint: letters, numbers, punctuation,
symbols and ASCII space.


<a id="example_IsPrint">Example</a>

## <a id="Itoa">func</a> [Itoa](https://golang.org/src/strconv/itoa.go?s=1028:1051#L24)
<pre>func Itoa(i <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#string">string</a></pre>
Itoa is equivalent to FormatInt(int64(i), 10).


<a id="example_Itoa">Example</a>

## <a id="ParseBool">func</a> [ParseBool](https://golang.org/src/strconv/atob.go?s=351:391#L1)
<pre>func ParseBool(str <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#bool">bool</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseBool returns the boolean value represented by the string.
It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
Any other value returns an error.


<a id="example_ParseBool">Example</a>

## <a id="ParseFloat">func</a> [ParseFloat](https://golang.org/src/strconv/atof.go?s=14691:14746#L650)
<pre>func ParseFloat(s <a href="/pkg/builtin/#string">string</a>, bitSize <a href="/pkg/builtin/#int">int</a>) (<a href="/pkg/builtin/#float64">float64</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseFloat converts the string s to a floating-point number
with the precision specified by bitSize: 32 for float32, or 64 for float64.
When bitSize=32, the result still has type float64, but it will be
convertible to float32 without changing its value.

ParseFloat accepts decimal and hexadecimal floating-point number syntax.
If s is well-formed and near a valid floating-point number,
ParseFloat returns the nearest floating-point number rounded
using IEEE754 unbiased rounding.
(Parsing a hexadecimal floating-point value only rounds when
there are more bits in the hexadecimal representation than
will fit in the mantissa.)

The errors that ParseFloat returns have concrete type *NumError
and include err.Num = s.

If s is not syntactically well-formed, ParseFloat returns err.Err = ErrSyntax.

If s is syntactically well-formed but is more than 1/2 ULP
away from the largest floating point number of the given size,
ParseFloat returns f = ±Inf, err.Err = ErrRange.

ParseFloat recognizes the strings "NaN", "+Inf", and "-Inf" as their
respective special floating point values. It ignores case when matching.


<a id="example_ParseFloat">Example</a>

## <a id="ParseInt">func</a> [ParseInt](https://golang.org/src/strconv/atoi.go?s=4679:4746#L163)
<pre>func ParseInt(s <a href="/pkg/builtin/#string">string</a>, base <a href="/pkg/builtin/#int">int</a>, bitSize <a href="/pkg/builtin/#int">int</a>) (i <a href="/pkg/builtin/#int64">int64</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ParseInt interprets a string s in the given base (0, 2 to 36) and
bit size (0 to 64) and returns the corresponding value i.

If base == 0, the base is implied by the string's prefix:
base 2 for "0b", base 8 for "0" or "0o", base 16 for "0x",
and base 10 otherwise. Also, for base == 0 only, underscore
characters are permitted per the Go integer literal syntax.
If base is below 0, is 1, or is above 36, an error is returned.

The bitSize argument specifies the integer type
that the result must fit into. Bit sizes 0, 8, 16, 32, and 64
correspond to int, int8, int16, int32, and int64.
If bitSize is below 0 or above 64, an error is returned.

The errors that ParseInt returns have concrete type *NumError
and include err.Num = s. If s is empty or contains invalid
digits, err.Err = ErrSyntax and the returned value is 0;
if the value corresponding to s cannot be represented by a
signed integer of the given size, err.Err = ErrRange and the
returned value is the maximum magnitude integer of the
appropriate bitSize and sign.


<a id="example_ParseInt">Example</a>

## <a id="ParseUint">func</a> [ParseUint](https://golang.org/src/strconv/atoi.go?s=1789:1852#L48)
<pre>func ParseUint(s <a href="/pkg/builtin/#string">string</a>, base <a href="/pkg/builtin/#int">int</a>, bitSize <a href="/pkg/builtin/#int">int</a>) (<a href="/pkg/builtin/#uint64">uint64</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseUint is like ParseInt but for unsigned numbers.


<a id="example_ParseUint">Example</a>

## <a id="Quote">func</a> [Quote](https://golang.org/src/strconv/quote.go?s=3261:3288#L114)
<pre>func Quote(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Quote returns a double-quoted Go string literal representing s. The
returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for
control characters and non-printable characters as defined by
IsPrint.


<a id="example_Quote">Example</a>

## <a id="QuoteRune">func</a> [QuoteRune](https://golang.org/src/strconv/quote.go?s=4950:4979#L153)
<pre>func QuoteRune(r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#string">string</a></pre>
QuoteRune returns a single-quoted Go character literal representing the
rune. The returned string uses Go escape sequences (\t, \n, \xFF, \u0100)
for control characters and non-printable characters as defined by IsPrint.


<a id="example_QuoteRune">Example</a>

## <a id="QuoteRuneToASCII">func</a> [QuoteRuneToASCII](https://golang.org/src/strconv/quote.go?s=5539:5575#L167)
<pre>func QuoteRuneToASCII(r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#string">string</a></pre>
QuoteRuneToASCII returns a single-quoted Go character literal representing
the rune. The returned string uses Go escape sequences (\t, \n, \xFF,
\u0100) for non-ASCII characters and non-printable characters as defined
by IsPrint.


<a id="example_QuoteRuneToASCII">Example</a>

## <a id="QuoteRuneToGraphic">func</a> [QuoteRuneToGraphic](https://golang.org/src/strconv/quote.go?s=6158:6196#L181)
<pre>func QuoteRuneToGraphic(r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#string">string</a></pre>
QuoteRuneToGraphic returns a single-quoted Go character literal representing
the rune. The returned string uses Go escape sequences (\t, \n, \xFF,
\u0100) for non-ASCII characters and non-printable characters as defined
by IsGraphic.


<a id="example_QuoteRuneToGraphic">Example</a>

## <a id="QuoteToASCII">func</a> [QuoteToASCII](https://golang.org/src/strconv/quote.go?s=3801:3835#L127)
<pre>func QuoteToASCII(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
QuoteToASCII returns a double-quoted Go string literal representing s.
The returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for
non-ASCII characters and non-printable characters as defined by IsPrint.


<a id="example_QuoteToASCII">Example</a>

## <a id="QuoteToGraphic">func</a> [QuoteToGraphic](https://golang.org/src/strconv/quote.go?s=4371:4407#L140)
<pre>func QuoteToGraphic(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
QuoteToGraphic returns a double-quoted Go string literal representing s.
The returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for
non-ASCII characters and non-printable characters as defined by IsGraphic.


<a id="example_QuoteToGraphic">Example</a>

## <a id="Unquote">func</a> [Unquote](https://golang.org/src/strconv/quote.go?s=10231:10269#L357)
<pre>func Unquote(s <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Unquote interprets s as a single-quoted, double-quoted,
or backquoted Go string literal, returning the string value
that s quotes.  (If s is single-quoted, it would be a Go
character literal; Unquote returns the corresponding
one-character string.)


<a id="example_Unquote">Example</a>

## <a id="UnquoteChar">func</a> [UnquoteChar](https://golang.org/src/strconv/quote.go?s=8246:8337#L241)
<pre>func UnquoteChar(s <a href="/pkg/builtin/#string">string</a>, quote <a href="/pkg/builtin/#byte">byte</a>) (value <a href="/pkg/builtin/#rune">rune</a>, multibyte <a href="/pkg/builtin/#bool">bool</a>, tail <a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
UnquoteChar decodes the first character or byte in the escaped string
or character literal represented by the string s.
It returns four values:


	1) value, the decoded Unicode code point or byte value;
	2) multibyte, a boolean indicating whether the decoded character requires a multibyte UTF-8 representation;
	3) tail, the remainder of the string after the character; and
	4) an error that will be nil if the character is syntactically valid.

The second argument, quote, specifies the type of literal being parsed
and therefore which escaped quote character is permitted.
If set to a single quote, it permits the sequence \' and disallows unescaped '.
If set to a double quote, it permits \" and disallows unescaped ".
If set to zero, it does not permit either escape and allows both quote characters to appear unescaped.


<a id="example_UnquoteChar">Example</a>



## <a id="NumError">type</a> [NumError](https://golang.org/src/strconv/atoi.go?s=809:1023#L14)
A NumError records a failed conversion.


<pre>type NumError struct {
<span id="NumError.Func"></span>    Func <a href="/pkg/builtin/#string">string</a> <span class="comment">// the failing function (ParseBool, ParseInt, ParseUint, ParseFloat)</span>
<span id="NumError.Num"></span>    Num  <a href="/pkg/builtin/#string">string</a> <span class="comment">// the input</span>
<span id="NumError.Err"></span>    Err  <a href="/pkg/builtin/#error">error</a>  <span class="comment">// the reason the conversion failed (e.g. ErrRange, ErrSyntax, etc.)</span>
}
</pre>





<a id="example_NumError">Example</a>






### <a id="NumError.Error">func</a> (\*NumError) [Error](https://golang.org/src/strconv/atoi.go?s=1025:1058#L20)
<pre>func (e *<a href="#NumError">NumError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>







