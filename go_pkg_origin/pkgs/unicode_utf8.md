

# utf8
`import "unicode/utf8"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package utf8 implements functions and constants to support text encoded in
UTF-8. It includes functions to translate between runes and UTF-8 byte sequences.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func DecodeLastRune(p []byte) (r rune, size int)](#DecodeLastRune)
* [func DecodeLastRuneInString(s string) (r rune, size int)](#DecodeLastRuneInString)
* [func DecodeRune(p []byte) (r rune, size int)](#DecodeRune)
* [func DecodeRuneInString(s string) (r rune, size int)](#DecodeRuneInString)
* [func EncodeRune(p []byte, r rune) int](#EncodeRune)
* [func FullRune(p []byte) bool](#FullRune)
* [func FullRuneInString(s string) bool](#FullRuneInString)
* [func RuneCount(p []byte) int](#RuneCount)
* [func RuneCountInString(s string) (n int)](#RuneCountInString)
* [func RuneLen(r rune) int](#RuneLen)
* [func RuneStart(b byte) bool](#RuneStart)
* [func Valid(p []byte) bool](#Valid)
* [func ValidRune(r rune) bool](#ValidRune)
* [func ValidString(s string) bool](#ValidString)


#### <a id="pkg-examples">Examples</a>
* [DecodeLastRune](#example_DecodeLastRune)
* [DecodeLastRuneInString](#example_DecodeLastRuneInString)
* [DecodeRune](#example_DecodeRune)
* [DecodeRuneInString](#example_DecodeRuneInString)
* [EncodeRune](#example_EncodeRune)
* [FullRune](#example_FullRune)
* [FullRuneInString](#example_FullRuneInString)
* [RuneCount](#example_RuneCount)
* [RuneCountInString](#example_RuneCountInString)
* [RuneLen](#example_RuneLen)
* [RuneStart](#example_RuneStart)
* [Valid](#example_Valid)
* [ValidRune](#example_ValidRune)
* [ValidString](#example_ValidString)


#### <a id="pkg-files">Package files</a>
[utf8.go](https://golang.org/src/unicode/utf8/utf8.go) 


## <a id="pkg-constants">Constants</a>
Numbers fundamental to the encoding.


<pre>const (
    <span id="RuneError">RuneError</span> = &#39;\uFFFD&#39;     <span class="comment">// the &#34;error&#34; Rune or &#34;Unicode replacement character&#34;</span>
    <span id="RuneSelf">RuneSelf</span>  = 0x80         <span class="comment">// characters below Runeself are represented as themselves in a single byte.</span>
    <span id="MaxRune">MaxRune</span>   = &#39;\U0010FFFF&#39; <span class="comment">// Maximum valid Unicode code point.</span>
    <span id="UTFMax">UTFMax</span>    = 4            <span class="comment">// maximum number of bytes of a UTF-8 encoded Unicode character.</span>
)</pre>



## <a id="DecodeLastRune">func</a> [DecodeLastRune](https://golang.org/src/unicode/utf8/utf8.go?s=8257:8305#L236)
<pre>func DecodeLastRune(p []<a href="/pkg/builtin/#byte">byte</a>) (r <a href="/pkg/builtin/#rune">rune</a>, size <a href="/pkg/builtin/#int">int</a>)</pre>
DecodeLastRune unpacks the last UTF-8 encoding in p and returns the rune and
its width in bytes. If p is empty it returns (RuneError, 0). Otherwise, if
the encoding is invalid, it returns (RuneError, 1). Both are impossible
results for correct, non-empty UTF-8.

An encoding is invalid if it is incorrect UTF-8, encodes a rune that is
out of range, or is not the shortest possible UTF-8 encoding for the
value. No other validation is performed.



<a id="example_DecodeLastRune">Example</a>


```go
```

output:
```txt
```

## <a id="DecodeLastRuneInString">func</a> [DecodeLastRuneInString](https://golang.org/src/unicode/utf8/utf8.go?s=9267:9323#L276)
<pre>func DecodeLastRuneInString(s <a href="/pkg/builtin/#string">string</a>) (r <a href="/pkg/builtin/#rune">rune</a>, size <a href="/pkg/builtin/#int">int</a>)</pre>
DecodeLastRuneInString is like DecodeLastRune but its input is a string. If
s is empty it returns (RuneError, 0). Otherwise, if the encoding is invalid,
it returns (RuneError, 1). Both are impossible results for correct,
non-empty UTF-8.

An encoding is invalid if it is incorrect UTF-8, encodes a rune that is
out of range, or is not the shortest possible UTF-8 encoding for the
value. No other validation is performed.



<a id="example_DecodeLastRuneInString">Example</a>


```go
```

output:
```txt
```

## <a id="DecodeRune">func</a> [DecodeRune](https://golang.org/src/unicode/utf8/utf8.go?s=5271:5315#L140)
<pre>func DecodeRune(p []<a href="/pkg/builtin/#byte">byte</a>) (r <a href="/pkg/builtin/#rune">rune</a>, size <a href="/pkg/builtin/#int">int</a>)</pre>
DecodeRune unpacks the first UTF-8 encoding in p and returns the rune and
its width in bytes. If p is empty it returns (RuneError, 0). Otherwise, if
the encoding is invalid, it returns (RuneError, 1). Both are impossible
results for correct, non-empty UTF-8.

An encoding is invalid if it is incorrect UTF-8, encodes a rune that is
out of range, or is not the shortest possible UTF-8 encoding for the
value. No other validation is performed.



<a id="example_DecodeRune">Example</a>


```go
```

output:
```txt
```

## <a id="DecodeRuneInString">func</a> [DecodeRuneInString](https://golang.org/src/unicode/utf8/utf8.go?s=6744:6796#L188)
<pre>func DecodeRuneInString(s <a href="/pkg/builtin/#string">string</a>) (r <a href="/pkg/builtin/#rune">rune</a>, size <a href="/pkg/builtin/#int">int</a>)</pre>
DecodeRuneInString is like DecodeRune but its input is a string. If s is
empty it returns (RuneError, 0). Otherwise, if the encoding is invalid, it
returns (RuneError, 1). Both are impossible results for correct, non-empty
UTF-8.

An encoding is invalid if it is incorrect UTF-8, encodes a rune that is
out of range, or is not the shortest possible UTF-8 encoding for the
value. No other validation is performed.



<a id="example_DecodeRuneInString">Example</a>


```go
```

output:
```txt
```

## <a id="EncodeRune">func</a> [EncodeRune](https://golang.org/src/unicode/utf8/utf8.go?s=10383:10420#L330)
<pre>func EncodeRune(p []<a href="/pkg/builtin/#byte">byte</a>, r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#int">int</a></pre>
EncodeRune writes into p (which must be large enough) the UTF-8 encoding of the rune.
It returns the number of bytes written.



<a id="example_EncodeRune">Example</a>


```go
```

output:
```txt
```

## <a id="FullRune">func</a> [FullRune](https://golang.org/src/unicode/utf8/utf8.go?s=3991:4019#L93)
<pre>func FullRune(p []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
FullRune reports whether the bytes in p begin with a full UTF-8 encoding of a rune.
An invalid encoding is considered a full Rune since it will convert as a width-1 error rune.



<a id="example_FullRune">Example</a>


```go
```

output:
```txt
```

## <a id="FullRuneInString">func</a> [FullRuneInString](https://golang.org/src/unicode/utf8/utf8.go?s=4426:4462#L113)
<pre>func FullRuneInString(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
FullRuneInString is like FullRune but its input is a string.



<a id="example_FullRuneInString">Example</a>


```go
```

output:
```txt
```

## <a id="RuneCount">func</a> [RuneCount](https://golang.org/src/unicode/utf8/utf8.go?s=11249:11277#L362)
<pre>func RuneCount(p []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#int">int</a></pre>
RuneCount returns the number of runes in p. Erroneous and short
encodings are treated as single runes of width 1 byte.



<a id="example_RuneCount">Example</a>


```go
```

output:
```txt
```

## <a id="RuneCountInString">func</a> [RuneCountInString](https://golang.org/src/unicode/utf8/utf8.go?s=11921:11961#L399)
<pre>func RuneCountInString(s <a href="/pkg/builtin/#string">string</a>) (n <a href="/pkg/builtin/#int">int</a>)</pre>
RuneCountInString is like RuneCount but its input is a string.



<a id="example_RuneCountInString">Example</a>


```go
```

output:
```txt
```

## <a id="RuneLen">func</a> [RuneLen](https://golang.org/src/unicode/utf8/utf8.go?s=9987:10011#L310)
<pre>func RuneLen(r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#int">int</a></pre>
RuneLen returns the number of bytes required to encode the rune.
It returns -1 if the rune is not a valid value to encode in UTF-8.



<a id="example_RuneLen">Example</a>


```go
```

output:
```txt
```

## <a id="RuneStart">func</a> [RuneStart](https://golang.org/src/unicode/utf8/utf8.go?s=12700:12727#L436)
<pre>func RuneStart(b <a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
RuneStart reports whether the byte could be the first byte of an encoded,
possibly invalid rune. Second and subsequent bytes always have the top two
bits set to 10.



<a id="example_RuneStart">Example</a>


```go
```

output:
```txt
```

## <a id="Valid">func</a> [Valid](https://golang.org/src/unicode/utf8/utf8.go?s=12830:12855#L439)
<pre>func Valid(p []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Valid reports whether p consists entirely of valid UTF-8-encoded runes.



<a id="example_Valid">Example</a>


```go
```

output:
```txt
```

## <a id="ValidRune">func</a> [ValidRune](https://golang.org/src/unicode/utf8/utf8.go?s=14223:14250#L504)
<pre>func ValidRune(r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
ValidRune reports whether r can be legally encoded as UTF-8.
Code points that are out of range or a surrogate half are illegal.



<a id="example_ValidRune">Example</a>


```go
```

output:
```txt
```

## <a id="ValidString">func</a> [ValidString](https://golang.org/src/unicode/utf8/utf8.go?s=13497:13528#L471)
<pre>func ValidString(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
ValidString reports whether s consists entirely of valid UTF-8-encoded runes.



<a id="example_ValidString">Example</a>


```go
```

output:
```txt
```






