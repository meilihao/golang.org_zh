

# base32
`import "encoding/base32"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package base32 implements base32 encoding as specified by RFC 4648.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func NewDecoder(enc *Encoding, r io.Reader) io.Reader](#NewDecoder)
* [func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser](#NewEncoder)
* [type CorruptInputError](#CorruptInputError)
  * [func (e CorruptInputError) Error() string](#CorruptInputError.Error)
* [type Encoding](#Encoding)
  * [func NewEncoding(encoder string) *Encoding](#NewEncoding)
  * [func (enc *Encoding) Decode(dst, src []byte) (n int, err error)](#Encoding.Decode)
  * [func (enc *Encoding) DecodeString(s string) ([]byte, error)](#Encoding.DecodeString)
  * [func (enc *Encoding) DecodedLen(n int) int](#Encoding.DecodedLen)
  * [func (enc *Encoding) Encode(dst, src []byte)](#Encoding.Encode)
  * [func (enc *Encoding) EncodeToString(src []byte) string](#Encoding.EncodeToString)
  * [func (enc *Encoding) EncodedLen(n int) int](#Encoding.EncodedLen)
  * [func (enc Encoding) WithPadding(padding rune) *Encoding](#Encoding.WithPadding)


#### <a id="pkg-examples">Examples</a>
* [Encoding.DecodeString](#example_Encoding_DecodeString)
* [Encoding.EncodeToString](#example_Encoding_EncodeToString)
* [NewEncoder](#example_NewEncoder)


#### <a id="pkg-files">Package files</a>
[base32.go](https://golang.org/src/encoding/base32/base32.go) 


## <a id="pkg-constants">Constants</a>

<pre>const (
    <span id="StdPadding">StdPadding</span> <a href="/pkg/builtin/#rune">rune</a> = &#39;=&#39; <span class="comment">// Standard padding character</span>
    <span id="NoPadding">NoPadding</span>  <a href="/pkg/builtin/#rune">rune</a> = -1  <span class="comment">// No padding</span>
)</pre>

## <a id="pkg-variables">Variables</a>
HexEncoding is the ``Extended Hex Alphabet'' defined in RFC 4648.
It is typically used in DNS.


<pre>var <span id="HexEncoding">HexEncoding</span> = <a href="#NewEncoding">NewEncoding</a>(encodeHex)</pre>StdEncoding is the standard base32 encoding, as defined in
RFC 4648.


<pre>var <span id="StdEncoding">StdEncoding</span> = <a href="#NewEncoding">NewEncoding</a>(encodeStd)</pre>

## <a id="NewDecoder">func</a> [NewDecoder](https://golang.org/src/encoding/base32/base32.go?s=12588:12641#L512)
<pre>func NewDecoder(enc *<a href="#Encoding">Encoding</a>, r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a></pre>
NewDecoder constructs a new base32 stream decoder.



## <a id="NewEncoder">func</a> [NewEncoder](https://golang.org/src/encoding/base32/base32.go?s=5996:6054#L249)
<pre>func NewEncoder(enc *<a href="#Encoding">Encoding</a>, w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#WriteCloser">WriteCloser</a></pre>
NewEncoder returns a new base32 stream encoder. Data written to
the returned writer will be encoded using enc and then written to w.
Base32 encodings operate in 5-byte blocks; when finished
writing, the caller must Close the returned encoder to flush any
partially written blocks.


<a id="example_NewEncoder">Example</a>



## <a id="CorruptInputError">type</a> [CorruptInputError](https://golang.org/src/encoding/base32/base32.go?s=6341:6369#L266)

<pre>type CorruptInputError <a href="/pkg/builtin/#int64">int64</a></pre>











### <a id="CorruptInputError.Error">func</a> (CorruptInputError) [Error](https://golang.org/src/encoding/base32/base32.go?s=6371:6412#L268)
<pre>func (e <a href="#CorruptInputError">CorruptInputError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Encoding">type</a> [Encoding](https://golang.org/src/encoding/base32/base32.go?s=569:650#L13)
An Encoding is a radix 32 encoding/decoding scheme, defined by a
32-character alphabet. The most common is the "base32" encoding
introduced for SASL GSSAPI and standardized in RFC 4648.
The alternate "base32hex" encoding is used in DNSSEC.


<pre>type Encoding struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewEncoding">func</a> [NewEncoding](https://golang.org/src/encoding/base32/base32.go?s=964:1006#L29)
<pre>func NewEncoding(encoder <a href="/pkg/builtin/#string">string</a>) *<a href="#Encoding">Encoding</a></pre>
NewEncoding returns a new Encoding defined by the given alphabet,
which must be a 32-byte string.






### <a id="Encoding.Decode">func</a> (\*Encoding) [Decode](https://golang.org/src/encoding/base32/base32.go?s=9245:9308#L364)
<pre>func (enc *<a href="#Encoding">Encoding</a>) Decode(dst, src []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Decode decodes src using the encoding enc. It writes at most
DecodedLen(len(src)) bytes to dst and returns the number of bytes
written. If src contains invalid base32 data, it will return the
number of bytes successfully written and CorruptInputError.
New line characters (\r and \n) are ignored.




### <a id="Encoding.DecodeString">func</a> (\*Encoding) [DecodeString](https://golang.org/src/encoding/base32/base32.go?s=9470:9529#L371)
<pre>func (enc *<a href="#Encoding">Encoding</a>) DecodeString(s <a href="/pkg/builtin/#string">string</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DecodeString returns the bytes represented by the base32 string s.


<a id="example_Encoding_DecodeString">Example</a>


### <a id="Encoding.DecodedLen">func</a> (\*Encoding) [DecodedLen](https://golang.org/src/encoding/base32/base32.go?s=12827:12869#L518)
<pre>func (enc *<a href="#Encoding">Encoding</a>) DecodedLen(n <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
DecodedLen returns the maximum length in bytes of the decoded data
corresponding to n bytes of base32-encoded data.




### <a id="Encoding.Encode">func</a> (\*Encoding) [Encode](https://golang.org/src/encoding/base32/base32.go?s=2565:2609#L92)
<pre>func (enc *<a href="#Encoding">Encoding</a>) Encode(dst, src []<a href="/pkg/builtin/#byte">byte</a>)</pre>
Encode encodes src using the encoding enc, writing
EncodedLen(len(src)) bytes to dst.

The encoding pads the output to a multiple of 8 bytes,
so Encode is not appropriate for use on individual blocks
of a large data stream. Use NewEncoder() instead.




### <a id="Encoding.EncodeToString">func</a> (\*Encoding) [EncodeToString](https://golang.org/src/encoding/base32/base32.go?s=4156:4210#L168)
<pre>func (enc *<a href="#Encoding">Encoding</a>) EncodeToString(src []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#string">string</a></pre>
EncodeToString returns the base32 encoding of src.


<a id="example_Encoding_EncodeToString">Example</a>


### <a id="Encoding.EncodedLen">func</a> (\*Encoding) [EncodedLen](https://golang.org/src/encoding/base32/base32.go?s=6193:6235#L255)
<pre>func (enc *<a href="#Encoding">Encoding</a>) EncodedLen(n <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
EncodedLen returns the length in bytes of the base32 encoding
of an input buffer of length n.




### <a id="Encoding.WithPadding">func</a> (Encoding) [WithPadding](https://golang.org/src/encoding/base32/base32.go?s=1964:2019#L67)
<pre>func (enc <a href="#Encoding">Encoding</a>) WithPadding(padding <a href="/pkg/builtin/#rune">rune</a>) *<a href="#Encoding">Encoding</a></pre>
WithPadding creates a new encoding identical to enc except
with a specified padding character, or NoPadding to disable padding.
The padding character must not be '\r' or '\n', must not
be contained in the encoding's alphabet and must be a rune equal or
below '\xff'.








