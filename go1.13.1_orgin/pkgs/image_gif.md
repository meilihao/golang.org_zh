

# gif
`import "image/gif"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package gif implements a GIF image decoder and encoder.

The GIF specification is at <a href="https://www.w3.org/Graphics/GIF/spec-gif89a.txt">https://www.w3.org/Graphics/GIF/spec-gif89a.txt</a>.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func Decode(r io.Reader) (image.Image, error)](#Decode)
* [func DecodeConfig(r io.Reader) (image.Config, error)](#DecodeConfig)
* [func Encode(w io.Writer, m image.Image, o *Options) error](#Encode)
* [func EncodeAll(w io.Writer, g *GIF) error](#EncodeAll)
* [type GIF](#GIF)
  * [func DecodeAll(r io.Reader) (*GIF, error)](#DecodeAll)
* [type Options](#Options)




#### <a id="pkg-files">Package files</a>
[reader.go](https://golang.org/src/image/gif/reader.go) [writer.go](https://golang.org/src/image/gif/writer.go) 


## <a id="pkg-constants">Constants</a>
Disposal Methods.


<pre>const (
    <span id="DisposalNone">DisposalNone</span>       = 0x01
    <span id="DisposalBackground">DisposalBackground</span> = 0x02
    <span id="DisposalPrevious">DisposalPrevious</span>   = 0x03
)</pre>



## <a id="Decode">func</a> [Decode](https://golang.org/src/image/gif/reader.go?s=15206:15251#L551)
<pre>func Decode(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (<a href="/pkg/image/">image</a>.<a href="/pkg/image/#Image">Image</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Decode reads a GIF image from r and returns the first embedded
image as an image.Image.



## <a id="DecodeConfig">func</a> [DecodeConfig](https://golang.org/src/image/gif/reader.go?s=17445:17497#L613)
<pre>func DecodeConfig(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (<a href="/pkg/image/">image</a>.<a href="/pkg/image/#Config">Config</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DecodeConfig returns the global color model and dimensions of a GIF image
without decoding the entire image.



## <a id="Encode">func</a> [Encode](https://golang.org/src/image/gif/writer.go?s=10643:10700#L408)
<pre>func Encode(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, m <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Image">Image</a>, o *<a href="#Options">Options</a>) <a href="/pkg/builtin/#error">error</a></pre>
Encode writes the Image m to w in GIF format.



## <a id="EncodeAll">func</a> [EncodeAll](https://golang.org/src/image/gif/writer.go?s=9384:9425#L362)
<pre>func EncodeAll(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, g *<a href="#GIF">GIF</a>) <a href="/pkg/builtin/#error">error</a></pre>
EncodeAll writes the images in g to w in GIF format with the
given loop count and delay between frames.





## <a id="GIF">type</a> [GIF](https://golang.org/src/image/gif/reader.go?s=15437:16816#L560)
GIF represents the possibly multiple images stored in a GIF file.


<pre>type GIF struct {
<span id="GIF.Image"></span>    Image []*<a href="/pkg/image/">image</a>.<a href="/pkg/image/#Paletted">Paletted</a> <span class="comment">// The successive images.</span>
<span id="GIF.Delay"></span>    Delay []<a href="/pkg/builtin/#int">int</a>             <span class="comment">// The successive delay times, one per frame, in 100ths of a second.</span>
<span id="GIF.LoopCount"></span>    <span class="comment">// LoopCount controls the number of times an animation will be</span>
    <span class="comment">// restarted during display.</span>
    <span class="comment">// A LoopCount of 0 means to loop forever.</span>
    <span class="comment">// A LoopCount of -1 means to show each frame only once.</span>
    <span class="comment">// Otherwise, the animation is looped LoopCount+1 times.</span>
    LoopCount <a href="/pkg/builtin/#int">int</a>
<span id="GIF.Disposal"></span>    <span class="comment">// Disposal is the successive disposal methods, one per frame. For</span>
    <span class="comment">// backwards compatibility, a nil Disposal is valid to pass to EncodeAll,</span>
    <span class="comment">// and implies that each frame&#39;s disposal method is 0 (no disposal</span>
    <span class="comment">// specified).</span>
    Disposal []<a href="/pkg/builtin/#byte">byte</a>
<span id="GIF.Config"></span>    <span class="comment">// Config is the global color table (palette), width and height. A nil or</span>
    <span class="comment">// empty-color.Palette Config.ColorModel means that each frame has its own</span>
    <span class="comment">// color table and there is no global color table. Each frame&#39;s bounds must</span>
    <span class="comment">// be within the rectangle defined by the two points (0, 0) and</span>
    <span class="comment">// (Config.Width, Config.Height).</span>
    <span class="comment">//</span>
    <span class="comment">// For backwards compatibility, a zero-valued Config is valid to pass to</span>
    <span class="comment">// EncodeAll, and implies that the overall GIF&#39;s width and height equals</span>
    <span class="comment">// the first frame&#39;s bounds&#39; Rectangle.Max point.</span>
    Config <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Config">Config</a>
<span id="GIF.BackgroundIndex"></span>    <span class="comment">// BackgroundIndex is the background index in the global color table, for</span>
    <span class="comment">// use with the DisposalBackground disposal method.</span>
    BackgroundIndex <a href="/pkg/builtin/#byte">byte</a>
}
</pre>









### <a id="DecodeAll">func</a> [DecodeAll](https://golang.org/src/image/gif/reader.go?s=16917:16958#L591)
<pre>func DecodeAll(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (*<a href="#GIF">GIF</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DecodeAll reads a GIF image from r and returns the sequential frames
and timing information.






## <a id="Options">type</a> [Options](https://golang.org/src/image/gif/writer.go?s=8844:9272#L346)
Options are the encoding parameters.


<pre>type Options struct {
<span id="Options.NumColors"></span>    <span class="comment">// NumColors is the maximum number of colors used in the image.</span>
    <span class="comment">// It ranges from 1 to 256.</span>
    NumColors <a href="/pkg/builtin/#int">int</a>

<span id="Options.Quantizer"></span>    <span class="comment">// Quantizer is used to produce a palette with size NumColors.</span>
    <span class="comment">// palette.Plan9 is used in place of a nil Quantizer.</span>
    Quantizer <a href="/pkg/image/draw/">draw</a>.<a href="/pkg/image/draw/#Quantizer">Quantizer</a>

<span id="Options.Drawer"></span>    <span class="comment">// Drawer is used to convert the source image to the desired palette.</span>
    <span class="comment">// draw.FloydSteinberg is used in place of a nil Drawer.</span>
    Drawer <a href="/pkg/image/draw/">draw</a>.<a href="/pkg/image/draw/#Drawer">Drawer</a>
}
</pre>















