

# utf16
`import "unicode/utf16"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package utf16 implements encoding and decoding of UTF-16 sequences.




## <a id="pkg-index">Index</a>
* [func Decode(s []uint16) []rune](#Decode)
* [func DecodeRune(r1, r2 rune) rune](#DecodeRune)
* [func Encode(s []rune) []uint16](#Encode)
* [func EncodeRune(r rune) (r1, r2 rune)](#EncodeRune)
* [func IsSurrogate(r rune) bool](#IsSurrogate)




#### <a id="pkg-files">Package files</a>
[utf16.go](https://golang.org/src/unicode/utf16/utf16.go) 






## <a id="Decode">func</a> [Decode](https://golang.org/src/unicode/utf16/utf16.go?s=2378:2408#L78)
<pre>func Decode(s []<a href="/pkg/builtin/#uint16">uint16</a>) []<a href="/pkg/builtin/#rune">rune</a></pre>
Decode returns the Unicode code point sequence represented
by the UTF-16 encoding s.



## <a id="DecodeRune">func</a> [DecodeRune](https://golang.org/src/unicode/utf16/utf16.go?s=1164:1197#L27)
<pre>func DecodeRune(r1, r2 <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#rune">rune</a></pre>
DecodeRune returns the UTF-16 decoding of a surrogate pair.
If the pair is not a valid UTF-16 surrogate pair, DecodeRune returns
the Unicode replacement code point U+FFFD.



## <a id="Encode">func</a> [Encode](https://golang.org/src/unicode/utf16/utf16.go?s=1790:1820#L46)
<pre>func Encode(s []<a href="/pkg/builtin/#rune">rune</a>) []<a href="/pkg/builtin/#uint16">uint16</a></pre>
Encode returns the UTF-16 encoding of the Unicode code point sequence s.



## <a id="EncodeRune">func</a> [EncodeRune](https://golang.org/src/unicode/utf16/utf16.go?s=1530:1567#L37)
<pre>func EncodeRune(r <a href="/pkg/builtin/#rune">rune</a>) (r1, r2 <a href="/pkg/builtin/#rune">rune</a>)</pre>
EncodeRune returns the UTF-16 surrogate pair r1, r2 for the given rune.
If the rune is not a valid Unicode code point or does not need encoding,
EncodeRune returns U+FFFD, U+FFFD.



## <a id="IsSurrogate">func</a> [IsSurrogate](https://golang.org/src/unicode/utf16/utf16.go?s=916:945#L20)
<pre>func IsSurrogate(r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsSurrogate reports whether the specified Unicode code point
can appear in a surrogate pair.








