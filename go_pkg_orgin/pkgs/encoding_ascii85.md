

# ascii85
`import "encoding/ascii85"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package ascii85 implements the ascii85 data encoding
as used in the btoa tool and Adobe's PostScript and PDF document formats.




## <a id="pkg-index">Index</a>
* [func Decode(dst, src []byte, flush bool) (ndst, nsrc int, err error)](#Decode)
* [func Encode(dst, src []byte) int](#Encode)
* [func MaxEncodedLen(n int) int](#MaxEncodedLen)
* [func NewDecoder(r io.Reader) io.Reader](#NewDecoder)
* [func NewEncoder(w io.Writer) io.WriteCloser](#NewEncoder)
* [type CorruptInputError](#CorruptInputError)
  * [func (e CorruptInputError) Error() string](#CorruptInputError.Error)




#### <a id="pkg-files">Package files</a>
[ascii85.go](https://golang.org/src/encoding/ascii85/ascii85.go) 






## <a id="Decode">func</a> [Decode](https://golang.org/src/encoding/ascii85/ascii85.go?s=4411:4479#L179)
<pre>func Decode(dst, src []<a href="/pkg/builtin/#byte">byte</a>, flush <a href="/pkg/builtin/#bool">bool</a>) (ndst, nsrc <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Decode decodes src into dst, returning both the number
of bytes written to dst and the number consumed from src.
If src contains invalid ascii85 data, Decode will return the
number of bytes successfully written and a CorruptInputError.
Decode ignores space and control characters in src.
Often, ascii85-encoded data is wrapped in <~ and ~> symbols.
Decode expects these to have been stripped by the caller.

If flush is true, Decode assumes that src represents the
end of the input stream and processes it completely rather
than wait for the completion of another 32-bit block.

NewDecoder wraps an io.Reader interface around Decode.



## <a id="Encode">func</a> [Encode](https://golang.org/src/encoding/ascii85/ascii85.go?s=781:813#L17)
<pre>func Encode(dst, src []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#int">int</a></pre>
Encode encodes src into at most MaxEncodedLen(len(src))
bytes of dst, returning the actual number of bytes written.

The encoding handles 4-byte chunks, using a special encoding
for the last fragment, so Encode is not appropriate for use on
individual blocks of a large data stream. Use NewEncoder() instead.

Often, ascii85-encoded data is wrapped in <~ and ~> symbols.
Encode does not add these.



## <a id="MaxEncodedLen">func</a> [MaxEncodedLen](https://golang.org/src/encoding/ascii85/ascii85.go?s=1788:1817#L76)
<pre>func MaxEncodedLen(n <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
MaxEncodedLen returns the maximum length of an encoding of n source bytes.



## <a id="NewDecoder">func</a> [NewDecoder](https://golang.org/src/encoding/ascii85/ascii85.go?s=5662:5700#L236)
<pre>func NewDecoder(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a></pre>
NewDecoder constructs a new ascii85 stream decoder.



## <a id="NewEncoder">func</a> [NewEncoder](https://golang.org/src/encoding/ascii85/ascii85.go?s=2132:2175#L83)
<pre>func NewEncoder(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#WriteCloser">WriteCloser</a></pre>
NewEncoder returns a new ascii85 stream encoder. Data written to
the returned writer will be encoded and then written to w.
Ascii85 encodings operate in 32-bit blocks; when finished
writing, the caller must Close the returned encoder to flush any
trailing partial block.





## <a id="CorruptInputError">type</a> [CorruptInputError](https://golang.org/src/encoding/ascii85/ascii85.go?s=3580:3608#L159)

<pre>type CorruptInputError <a href="/pkg/builtin/#int64">int64</a></pre>











### <a id="CorruptInputError.Error">func</a> (CorruptInputError) [Error](https://golang.org/src/encoding/ascii85/ascii85.go?s=3610:3651#L161)
<pre>func (e <a href="#CorruptInputError">CorruptInputError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>






