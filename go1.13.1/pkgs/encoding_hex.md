

# hex
`import "encoding/hex"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package hex implements hexadecimal encoding and decoding.




## <a id="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func Decode(dst, src []byte) (int, error)](#Decode)
* [func DecodeString(s string) ([]byte, error)](#DecodeString)
* [func DecodedLen(x int) int](#DecodedLen)
* [func Dump(data []byte) string](#Dump)
* [func Dumper(w io.Writer) io.WriteCloser](#Dumper)
* [func Encode(dst, src []byte) int](#Encode)
* [func EncodeToString(src []byte) string](#EncodeToString)
* [func EncodedLen(n int) int](#EncodedLen)
* [func NewDecoder(r io.Reader) io.Reader](#NewDecoder)
* [func NewEncoder(w io.Writer) io.Writer](#NewEncoder)
* [type InvalidByteError](#InvalidByteError)
  * [func (e InvalidByteError) Error() string](#InvalidByteError.Error)


#### <a id="pkg-examples">Examples</a>
* [Decode](#example_Decode)
* [DecodeString](#example_DecodeString)
* [Dump](#example_Dump)
* [Dumper](#example_Dumper)
* [Encode](#example_Encode)
* [EncodeToString](#example_EncodeToString)


#### <a id="pkg-files">Package files</a>
[hex.go](https://golang.org/src/encoding/hex/hex.go) 




## <a id="pkg-variables">Variables</a>
ErrLength reports an attempt to decode an odd-length input
using Decode or DecodeString.
The stream-based Decoder returns io.ErrUnexpectedEOF instead of ErrLength.


<pre>var <span id="ErrLength">ErrLength</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;encoding/hex: odd length hex string&#34;)</pre>

## <a id="Decode">func</a> [Decode](https://golang.org/src/encoding/hex/hex.go?s=1767:1808#L48)
<pre>func Decode(dst, src []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Decode decodes src into DecodedLen(len(src)) bytes,
returning the actual number of bytes written to dst.

Decode expects that src contains only hexadecimal
characters and that src has even length.
If the input is malformed, Decode returns the number
of bytes decoded before the error.


<a id="example_Decode">Example</a>

## <a id="DecodeString">func</a> [DecodeString](https://golang.org/src/encoding/hex/hex.go?s=3100:3143#L100)
<pre>func DecodeString(s <a href="/pkg/builtin/#string">string</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DecodeString returns the bytes represented by the hexadecimal string s.

DecodeString expects that src contains only hexadecimal
characters and that src has even length.
If the input is malformed, DecodeString returns
the bytes decoded before the error.


<a id="example_DecodeString">Example</a>

## <a id="DecodedLen">func</a> [DecodedLen](https://golang.org/src/encoding/hex/hex.go?s=1417:1443#L39)
<pre>func DecodedLen(x <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
DecodedLen returns the length of a decoding of x source bytes.
Specifically, it returns x / 2.



## <a id="Dump">func</a> [Dump](https://golang.org/src/encoding/hex/hex.go?s=3521:3550#L110)
<pre>func Dump(data []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#string">string</a></pre>
Dump returns a string that contains a hex dump of the given data. The format
of the hex dump matches the output of `hexdump -C` on the command line.


<a id="example_Dump">Example</a>

## <a id="Dumper">func</a> [Dumper](https://golang.org/src/encoding/hex/hex.go?s=6082:6121#L205)
<pre>func Dumper(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#WriteCloser">WriteCloser</a></pre>
Dumper returns a WriteCloser that writes a hex dump of all written data to
w. The format of the dump matches the output of `hexdump -C` on the command
line.


<a id="example_Dumper">Example</a>

## <a id="Encode">func</a> [Encode](https://golang.org/src/encoding/hex/hex.go?s=687:719#L15)
<pre>func Encode(dst, src []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#int">int</a></pre>
Encode encodes src into EncodedLen(len(src))
bytes of dst. As a convenience, it returns the number
of bytes written to dst, but this value is always EncodedLen(len(src)).
Encode implements hexadecimal encoding.


<a id="example_Encode">Example</a>

## <a id="EncodeToString">func</a> [EncodeToString](https://golang.org/src/encoding/hex/hex.go?s=2704:2742#L88)
<pre>func EncodeToString(src []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#string">string</a></pre>
EncodeToString returns the hexadecimal encoding of src.


<a id="example_EncodeToString">Example</a>

## <a id="EncodedLen">func</a> [EncodedLen](https://golang.org/src/encoding/hex/hex.go?s=419:445#L9)
<pre>func EncodedLen(n <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
EncodedLen returns the length of an encoding of n source bytes.
Specifically, it returns n * 2.



## <a id="NewDecoder">func</a> [NewDecoder](https://golang.org/src/encoding/hex/hex.go?s=4942:4980#L166)
<pre>func NewDecoder(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a></pre>
NewDecoder returns an io.Reader that decodes hexadecimal characters from r.
NewDecoder expects that r contain only an even number of hexadecimal characters.



## <a id="NewEncoder">func</a> [NewEncoder](https://golang.org/src/encoding/hex/hex.go?s=4219:4257#L137)
<pre>func NewEncoder(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a></pre>
NewEncoder returns an io.Writer that writes lowercase hexadecimal characters to w.





## <a id="InvalidByteError">type</a> [InvalidByteError](https://golang.org/src/encoding/hex/hex.go?s=1178:1204#L31)
InvalidByteError values describe errors resulting from an invalid byte in a hex string.


<pre>type InvalidByteError <a href="/pkg/builtin/#byte">byte</a></pre>











### <a id="InvalidByteError.Error">func</a> (InvalidByteError) [Error](https://golang.org/src/encoding/hex/hex.go?s=1206:1246#L33)
<pre>func (e <a href="#InvalidByteError">InvalidByteError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>







