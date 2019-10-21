

# base64
`import "encoding/base64"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package base64 implements base64 encoding as specified by RFC 4648.


<a id="example_">Example</a>
```go
```

output:
```txt
```


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
  * [func (enc Encoding) Strict() *Encoding](#Encoding.Strict)
  * [func (enc Encoding) WithPadding(padding rune) *Encoding](#Encoding.WithPadding)


#### <a id="pkg-examples">Examples</a>
* [Package](#example_)
* [Encoding.DecodeString](#example_Encoding_DecodeString)
* [Encoding.EncodeToString](#example_Encoding_EncodeToString)
* [NewEncoder](#example_NewEncoder)


#### <a id="pkg-files">Package files</a>
[base64.go](https://golang.org/src/encoding/base64/base64.go) 


## <a id="pkg-constants">Constants</a>

<pre>const (
    <span id="StdPadding">StdPadding</span> <a href="/pkg/builtin/#rune">rune</a> = &#39;=&#39; <span class="comment">// Standard padding character</span>
    <span id="NoPadding">NoPadding</span>  <a href="/pkg/builtin/#rune">rune</a> = -1  <span class="comment">// No padding</span>
)</pre>

## <a id="pkg-variables">Variables</a>
RawStdEncoding is the standard raw, unpadded base64 encoding,
as defined in RFC 4648 section 3.2.
This is the same as StdEncoding but omits padding characters.


<pre>var <span id="RawStdEncoding">RawStdEncoding</span> = <a href="#StdEncoding">StdEncoding</a>.<a href="#WithPadding">WithPadding</a>(<a href="#NoPadding">NoPadding</a>)</pre>RawURLEncoding is the unpadded alternate base64 encoding defined in RFC 4648.
It is typically used in URLs and file names.
This is the same as URLEncoding but omits padding characters.


<pre>var <span id="RawURLEncoding">RawURLEncoding</span> = <a href="#URLEncoding">URLEncoding</a>.<a href="#WithPadding">WithPadding</a>(<a href="#NoPadding">NoPadding</a>)</pre>StdEncoding is the standard base64 encoding, as defined in
RFC 4648.


<pre>var <span id="StdEncoding">StdEncoding</span> = <a href="#NewEncoding">NewEncoding</a>(encodeStd)</pre>URLEncoding is the alternate base64 encoding defined in RFC 4648.
It is typically used in URLs and file names.


<pre>var <span id="URLEncoding">URLEncoding</span> = <a href="#NewEncoding">NewEncoding</a>(encodeURL)</pre>

## <a id="NewDecoder">func</a> [NewDecoder](https://golang.org/src/encoding/base64/base64.go?s=14692:14745#L586)
<pre>func NewDecoder(enc *<a href="#Encoding">Encoding</a>, r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a></pre>
NewDecoder constructs a new base64 stream decoder.



## <a id="NewEncoder">func</a> [NewEncoder](https://golang.org/src/encoding/base64/base64.go?s=6737:6795#L244)
<pre>func NewEncoder(enc *<a href="#Encoding">Encoding</a>, w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#WriteCloser">WriteCloser</a></pre>
NewEncoder returns a new base64 stream encoder. Data written to
the returned writer will be encoded using enc and then written to w.
Base64 encodings operate in 4-byte blocks; when finished
writing, the caller must Close the returned encoder to flush any
partially written blocks.


<a id="example_NewEncoder">Example</a>
```go
```

output:
```txt
```



## <a id="CorruptInputError">type</a> [CorruptInputError](https://golang.org/src/encoding/base64/base64.go?s=7161:7189#L261)

<pre>type CorruptInputError <a href="/pkg/builtin/#int64">int64</a></pre>











### <a id="CorruptInputError.Error">func</a> (CorruptInputError) [Error](https://golang.org/src/encoding/base64/base64.go?s=7191:7232#L263)
<pre>func (e <a href="#CorruptInputError">CorruptInputError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Encoding">type</a> [Encoding](https://golang.org/src/encoding/base64/base64.go?s=652:749#L13)
An Encoding is a radix 64 encoding/decoding scheme, defined by a
64-character alphabet. The most common encoding is the "base64"
encoding defined in RFC 4648 and used in MIME (RFC 2045) and PEM
(RFC 1421).  RFC 4648 also defines an alternate encoding, which is
the standard encoding with - and _ substituted for + and /.


<pre>type Encoding struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewEncoding">func</a> [NewEncoding](https://golang.org/src/encoding/base64/base64.go?s=1326:1368#L33)
<pre>func NewEncoding(encoder <a href="/pkg/builtin/#string">string</a>) *<a href="#Encoding">Encoding</a></pre>
NewEncoding returns a new padded Encoding defined by the given alphabet,
which must be a 64-byte string that does not contain the padding character
or CR / LF ('\r', '\n').
The resulting Encoding uses the default padding character ('='),
which may be changed or disabled via WithPadding.






### <a id="Encoding.Decode">func</a> (\*Encoding) [Decode](https://golang.org/src/encoding/base64/base64.go?s=11757:11820#L458)
<pre>func (enc *<a href="#Encoding">Encoding</a>) Decode(dst, src []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Decode decodes src using the encoding enc. It writes at most
DecodedLen(len(src)) bytes to dst and returns the number of bytes
written. If src contains invalid base64 data, it will return the
number of bytes successfully written and CorruptInputError.
New line characters (\r and \n) are ignored.




### <a id="Encoding.DecodeString">func</a> (\*Encoding) [DecodeString](https://golang.org/src/encoding/base64/base64.go?s=9654:9713#L370)
<pre>func (enc *<a href="#Encoding">Encoding</a>) DecodeString(s <a href="/pkg/builtin/#string">string</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DecodeString returns the bytes represented by the base64 string s.


<a id="example_Encoding_DecodeString">Example</a>
```go
```

output:
```txt
```


### <a id="Encoding.DecodedLen">func</a> (\*Encoding) [DecodedLen](https://golang.org/src/encoding/base64/base64.go?s=14931:14973#L592)
<pre>func (enc *<a href="#Encoding">Encoding</a>) DecodedLen(n <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
DecodedLen returns the maximum length in bytes of the decoded data
corresponding to n bytes of base64-encoded data.




### <a id="Encoding.Encode">func</a> (\*Encoding) [Encode](https://golang.org/src/encoding/base64/base64.go?s=3744:3788#L112)
<pre>func (enc *<a href="#Encoding">Encoding</a>) Encode(dst, src []<a href="/pkg/builtin/#byte">byte</a>)</pre>
Encode encodes src using the encoding enc, writing
EncodedLen(len(src)) bytes to dst.

The encoding pads the output to a multiple of 4 bytes,
so Encode is not appropriate for use on individual blocks
of a large data stream. Use NewEncoder() instead.




### <a id="Encoding.EncodeToString">func</a> (\*Encoding) [EncodeToString](https://golang.org/src/encoding/base64/base64.go?s=4932:4986#L164)
<pre>func (enc *<a href="#Encoding">Encoding</a>) EncodeToString(src []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#string">string</a></pre>
EncodeToString returns the base64 encoding of src.


<a id="example_Encoding_EncodeToString">Example</a>
```go
```

output:
```txt
```


### <a id="Encoding.EncodedLen">func</a> (\*Encoding) [EncodedLen](https://golang.org/src/encoding/base64/base64.go?s=6934:6976#L250)
<pre>func (enc *<a href="#Encoding">Encoding</a>) EncodedLen(n <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
EncodedLen returns the length in bytes of the base64 encoding
of an input buffer of length n.




### <a id="Encoding.Strict">func</a> (Encoding) [Strict](https://golang.org/src/encoding/base64/base64.go?s=2629:2667#L79)
<pre>func (enc <a href="#Encoding">Encoding</a>) Strict() *<a href="#Encoding">Encoding</a></pre>
Strict creates a new encoding identical to enc except with
strict decoding enabled. In this mode, the decoder requires that
trailing padding bits are zero, as described in RFC 4648 section 3.5.




### <a id="Encoding.WithPadding">func</a> (Encoding) [WithPadding](https://golang.org/src/encoding/base64/base64.go?s=2111:2166#L61)
<pre>func (enc <a href="#Encoding">Encoding</a>) WithPadding(padding <a href="/pkg/builtin/#rune">rune</a>) *<a href="#Encoding">Encoding</a></pre>
WithPadding creates a new encoding identical to enc except
with a specified padding character, or NoPadding to disable padding.
The padding character must not be '\r' or '\n', must not
be contained in the encoding's alphabet and must be a rune equal or
below '\xff'.







