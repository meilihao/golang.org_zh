

# png
`import "image/png"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package png implements a PNG image decoder and encoder.

The PNG specification is at <a href="https://www.w3.org/TR/PNG/">https://www.w3.org/TR/PNG/</a>.




## <a id="pkg-index">Index</a>
* [func Decode(r io.Reader) (image.Image, error)](#Decode)
* [func DecodeConfig(r io.Reader) (image.Config, error)](#DecodeConfig)
* [func Encode(w io.Writer, m image.Image) error](#Encode)
* [type CompressionLevel](#CompressionLevel)
* [type Encoder](#Encoder)
  * [func (enc *Encoder) Encode(w io.Writer, m image.Image) error](#Encoder.Encode)
* [type EncoderBuffer](#EncoderBuffer)
* [type EncoderBufferPool](#EncoderBufferPool)
* [type FormatError](#FormatError)
  * [func (e FormatError) Error() string](#FormatError.Error)
* [type UnsupportedError](#UnsupportedError)
  * [func (e UnsupportedError) Error() string](#UnsupportedError.Error)


#### <a id="pkg-examples">Examples</a>
* [Decode](#example_Decode)
* [Encode](#example_Encode)


#### <a id="pkg-files">Package files</a>
[paeth.go](https://golang.org/src/image/png/paeth.go) [reader.go](https://golang.org/src/image/png/reader.go) [writer.go](https://golang.org/src/image/png/writer.go) 






## <a id="Decode">func</a> [Decode](https://golang.org/src/image/png/reader.go?s=24412:24457#L942)
<pre>func Decode(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (<a href="/pkg/image/">image</a>.<a href="/pkg/image/#Image">Image</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Decode reads a PNG image from r and returns it as an image.Image.
The type of Image returned depends on the PNG contents.


<a id="example_Decode">Example</a>

## <a id="DecodeConfig">func</a> [DecodeConfig](https://golang.org/src/image/png/reader.go?s=24912:24964#L966)
<pre>func DecodeConfig(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (<a href="/pkg/image/">image</a>.<a href="/pkg/image/#Config">Config</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DecodeConfig returns the color model and dimensions of a PNG image without
decoding the entire image.



## <a id="Encode">func</a> [Encode](https://golang.org/src/image/png/writer.go?s=12983:13028#L548)
<pre>func Encode(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, m <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Image">Image</a>) <a href="/pkg/builtin/#error">error</a></pre>
Encode writes the Image m to w in PNG format. Any Image may be
encoded, but images that are not image.NRGBA might be encoded lossily.


<a id="example_Encode">Example</a>



## <a id="CompressionLevel">type</a> [CompressionLevel](https://golang.org/src/image/png/writer.go?s=1155:1180#L44)

<pre>type CompressionLevel <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="DefaultCompression">DefaultCompression</span> <a href="#CompressionLevel">CompressionLevel</a> = 0
    <span id="NoCompression">NoCompression</span>      <a href="#CompressionLevel">CompressionLevel</a> = -1
    <span id="BestSpeed">BestSpeed</span>          <a href="#CompressionLevel">CompressionLevel</a> = -2
    <span id="BestCompression">BestCompression</span>    <a href="#CompressionLevel">CompressionLevel</a> = -3
)</pre>









## <a id="Encoder">type</a> [Encoder](https://golang.org/src/image/png/writer.go?s=328:527#L9)
Encoder configures encoding PNG images.


<pre>type Encoder struct {
<span id="Encoder.CompressionLevel"></span>    CompressionLevel <a href="#CompressionLevel">CompressionLevel</a>

<span id="Encoder.BufferPool"></span>    <span class="comment">// BufferPool optionally specifies a buffer pool to get temporary</span>
    <span class="comment">// EncoderBuffers when encoding an image.</span>
    BufferPool <a href="#EncoderBufferPool">EncoderBufferPool</a>
}
</pre>











### <a id="Encoder.Encode">func</a> (\*Encoder) [Encode](https://golang.org/src/image/png/writer.go?s=13121:13181#L554)
<pre>func (enc *<a href="#Encoder">Encoder</a>) Encode(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, m <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Image">Image</a>) <a href="/pkg/builtin/#error">error</a></pre>
Encode writes the Image m to w in PNG format.




## <a id="EncoderBuffer">type</a> [EncoderBuffer](https://golang.org/src/image/png/writer.go?s=858:884#L26)
EncoderBuffer holds the buffers used for encoding PNG images.


<pre>type EncoderBuffer encoder</pre>











## <a id="EncoderBufferPool">type</a> [EncoderBufferPool](https://golang.org/src/image/png/writer.go?s=712:791#L20)
EncoderBufferPool is an interface for getting and returning temporary
instances of the EncoderBuffer struct. This can be used to reuse buffers
when encoding multiple images.


<pre>type EncoderBufferPool interface {
    Get() *<a href="#EncoderBuffer">EncoderBuffer</a>
    Put(*<a href="#EncoderBuffer">EncoderBuffer</a>)
}</pre>











## <a id="FormatError">type</a> [FormatError](https://golang.org/src/image/png/reader.go?s=2386:2409#L114)
A FormatError reports that the input is not a valid PNG.


<pre>type FormatError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="FormatError.Error">func</a> (FormatError) [Error](https://golang.org/src/image/png/reader.go?s=2411:2446#L116)
<pre>func (e <a href="#FormatError">FormatError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="UnsupportedError">type</a> [UnsupportedError](https://golang.org/src/image/png/reader.go?s=2642:2670#L121)
An UnsupportedError reports that the input uses a valid but unimplemented PNG feature.


<pre>type UnsupportedError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="UnsupportedError.Error">func</a> (UnsupportedError) [Error](https://golang.org/src/image/png/reader.go?s=2672:2712#L123)
<pre>func (e <a href="#UnsupportedError">UnsupportedError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>







