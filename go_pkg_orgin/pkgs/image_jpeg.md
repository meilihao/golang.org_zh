

# jpeg
`import "image/jpeg"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package jpeg implements a JPEG image decoder and encoder.

JPEG is defined in ITU-T T.81: <a href="https://www.w3.org/Graphics/JPEG/itu-t81.pdf">https://www.w3.org/Graphics/JPEG/itu-t81.pdf</a>.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func Decode(r io.Reader) (image.Image, error)](#Decode)
* [func DecodeConfig(r io.Reader) (image.Config, error)](#DecodeConfig)
* [func Encode(w io.Writer, m image.Image, o *Options) error](#Encode)
* [type FormatError](#FormatError)
  * [func (e FormatError) Error() string](#FormatError.Error)
* [type Options](#Options)
* [type Reader](#Reader)
* [type UnsupportedError](#UnsupportedError)
  * [func (e UnsupportedError) Error() string](#UnsupportedError.Error)




#### <a id="pkg-files">Package files</a>
[fdct.go](https://golang.org/src/image/jpeg/fdct.go) [huffman.go](https://golang.org/src/image/jpeg/huffman.go) [idct.go](https://golang.org/src/image/jpeg/idct.go) [reader.go](https://golang.org/src/image/jpeg/reader.go) [scan.go](https://golang.org/src/image/jpeg/scan.go) [writer.go](https://golang.org/src/image/jpeg/writer.go) 


## <a id="pkg-constants">Constants</a>
DefaultQuality is the default quality encoding parameter.


<pre>const <span id="DefaultQuality">DefaultQuality</span> = 75</pre>



## <a id="Decode">func</a> [Decode](https://golang.org/src/image/jpeg/reader.go?s=22312:22357#L767)
<pre>func Decode(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (<a href="/pkg/image/">image</a>.<a href="/pkg/image/#Image">Image</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Decode reads a JPEG image from r and returns it as an image.Image.



## <a id="DecodeConfig">func</a> [DecodeConfig](https://golang.org/src/image/jpeg/reader.go?s=22514:22566#L774)
<pre>func DecodeConfig(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (<a href="/pkg/image/">image</a>.<a href="/pkg/image/#Config">Config</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DecodeConfig returns the color model and dimensions of a JPEG image without
decoding the entire image.



## <a id="Encode">func</a> [Encode](https://golang.org/src/image/jpeg/writer.go?s=16086:16143#L565)
<pre>func Encode(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, m <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Image">Image</a>, o *<a href="#Options">Options</a>) <a href="/pkg/builtin/#error">error</a></pre>
Encode writes the Image m to w in JPEG 4:2:0 baseline format with the given
options. Default parameters are used if a nil *Options is passed.





## <a id="FormatError">type</a> [FormatError](https://golang.org/src/image/jpeg/reader.go?s=583:606#L11)
A FormatError reports that the input is not a valid JPEG.


<pre>type FormatError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="FormatError.Error">func</a> (FormatError) [Error](https://golang.org/src/image/jpeg/reader.go?s=608:643#L13)
<pre>func (e <a href="#FormatError">FormatError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Options">type</a> [Options](https://golang.org/src/image/jpeg/writer.go?s=15900:15936#L559)
Options are the encoding parameters.
Quality ranges from 1 to 100 inclusive, higher is better.


<pre>type Options struct {
<span id="Options.Quality"></span>    Quality <a href="/pkg/builtin/#int">int</a>
}
</pre>











## <a id="Reader">type</a> [Reader](https://golang.org/src/image/jpeg/reader.go?s=2944:2995#L84)
Deprecated: Reader is not used by the image/jpeg package and should
not be used by others. It is kept for compatibility.


<pre>type Reader interface {
    <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ByteReader">ByteReader</a>
    <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>
}</pre>











## <a id="UnsupportedError">type</a> [UnsupportedError](https://golang.org/src/image/jpeg/reader.go?s=783:811#L16)
An UnsupportedError reports that the input uses a valid but unimplemented JPEG feature.


<pre>type UnsupportedError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="UnsupportedError.Error">func</a> (UnsupportedError) [Error](https://golang.org/src/image/jpeg/reader.go?s=813:853#L18)
<pre>func (e <a href="#UnsupportedError">UnsupportedError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>






