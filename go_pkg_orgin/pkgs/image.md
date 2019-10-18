

# image
`import "image"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package image implements a basic 2-D image library.

The fundamental interface is called Image. An Image contains colors, which
are described in the image/color package.

Values of the Image interface are created either by calling functions such
as NewRGBA and NewPaletted, or by calling Decode on an io.Reader containing
image data in a format such as GIF, JPEG or PNG. Decoding any particular
image format requires the prior registration of a decoder function.
Registration is typically automatic as a side effect of initializing that
format's package so that, to decode a PNG image, it suffices to have


	import _ "image/png"

in a program's main package. The _ means to import a package purely for its
initialization side effects.

See "The Go image package" for more details:
<a href="https://golang.org/doc/articles/image_package.html">https://golang.org/doc/articles/image_package.html</a>


<a id="example_">Example</a>
<a id="example__decodeConfig">Example (DecodeConfig)</a>


## <a id="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func RegisterFormat(name, magic string, decode func(io.Reader) (Image, error), decodeConfig func(io.Reader) (Config, error))](#RegisterFormat)
* [type Alpha](#Alpha)
  * [func NewAlpha(r Rectangle) *Alpha](#NewAlpha)
  * [func (p *Alpha) AlphaAt(x, y int) color.Alpha](#Alpha.AlphaAt)
  * [func (p *Alpha) At(x, y int) color.Color](#Alpha.At)
  * [func (p *Alpha) Bounds() Rectangle](#Alpha.Bounds)
  * [func (p *Alpha) ColorModel() color.Model](#Alpha.ColorModel)
  * [func (p *Alpha) Opaque() bool](#Alpha.Opaque)
  * [func (p *Alpha) PixOffset(x, y int) int](#Alpha.PixOffset)
  * [func (p *Alpha) Set(x, y int, c color.Color)](#Alpha.Set)
  * [func (p *Alpha) SetAlpha(x, y int, c color.Alpha)](#Alpha.SetAlpha)
  * [func (p *Alpha) SubImage(r Rectangle) Image](#Alpha.SubImage)
* [type Alpha16](#Alpha16)
  * [func NewAlpha16(r Rectangle) *Alpha16](#NewAlpha16)
  * [func (p *Alpha16) Alpha16At(x, y int) color.Alpha16](#Alpha16.Alpha16At)
  * [func (p *Alpha16) At(x, y int) color.Color](#Alpha16.At)
  * [func (p *Alpha16) Bounds() Rectangle](#Alpha16.Bounds)
  * [func (p *Alpha16) ColorModel() color.Model](#Alpha16.ColorModel)
  * [func (p *Alpha16) Opaque() bool](#Alpha16.Opaque)
  * [func (p *Alpha16) PixOffset(x, y int) int](#Alpha16.PixOffset)
  * [func (p *Alpha16) Set(x, y int, c color.Color)](#Alpha16.Set)
  * [func (p *Alpha16) SetAlpha16(x, y int, c color.Alpha16)](#Alpha16.SetAlpha16)
  * [func (p *Alpha16) SubImage(r Rectangle) Image](#Alpha16.SubImage)
* [type CMYK](#CMYK)
  * [func NewCMYK(r Rectangle) *CMYK](#NewCMYK)
  * [func (p *CMYK) At(x, y int) color.Color](#CMYK.At)
  * [func (p *CMYK) Bounds() Rectangle](#CMYK.Bounds)
  * [func (p *CMYK) CMYKAt(x, y int) color.CMYK](#CMYK.CMYKAt)
  * [func (p *CMYK) ColorModel() color.Model](#CMYK.ColorModel)
  * [func (p *CMYK) Opaque() bool](#CMYK.Opaque)
  * [func (p *CMYK) PixOffset(x, y int) int](#CMYK.PixOffset)
  * [func (p *CMYK) Set(x, y int, c color.Color)](#CMYK.Set)
  * [func (p *CMYK) SetCMYK(x, y int, c color.CMYK)](#CMYK.SetCMYK)
  * [func (p *CMYK) SubImage(r Rectangle) Image](#CMYK.SubImage)
* [type Config](#Config)
  * [func DecodeConfig(r io.Reader) (Config, string, error)](#DecodeConfig)
* [type Gray](#Gray)
  * [func NewGray(r Rectangle) *Gray](#NewGray)
  * [func (p *Gray) At(x, y int) color.Color](#Gray.At)
  * [func (p *Gray) Bounds() Rectangle](#Gray.Bounds)
  * [func (p *Gray) ColorModel() color.Model](#Gray.ColorModel)
  * [func (p *Gray) GrayAt(x, y int) color.Gray](#Gray.GrayAt)
  * [func (p *Gray) Opaque() bool](#Gray.Opaque)
  * [func (p *Gray) PixOffset(x, y int) int](#Gray.PixOffset)
  * [func (p *Gray) Set(x, y int, c color.Color)](#Gray.Set)
  * [func (p *Gray) SetGray(x, y int, c color.Gray)](#Gray.SetGray)
  * [func (p *Gray) SubImage(r Rectangle) Image](#Gray.SubImage)
* [type Gray16](#Gray16)
  * [func NewGray16(r Rectangle) *Gray16](#NewGray16)
  * [func (p *Gray16) At(x, y int) color.Color](#Gray16.At)
  * [func (p *Gray16) Bounds() Rectangle](#Gray16.Bounds)
  * [func (p *Gray16) ColorModel() color.Model](#Gray16.ColorModel)
  * [func (p *Gray16) Gray16At(x, y int) color.Gray16](#Gray16.Gray16At)
  * [func (p *Gray16) Opaque() bool](#Gray16.Opaque)
  * [func (p *Gray16) PixOffset(x, y int) int](#Gray16.PixOffset)
  * [func (p *Gray16) Set(x, y int, c color.Color)](#Gray16.Set)
  * [func (p *Gray16) SetGray16(x, y int, c color.Gray16)](#Gray16.SetGray16)
  * [func (p *Gray16) SubImage(r Rectangle) Image](#Gray16.SubImage)
* [type Image](#Image)
  * [func Decode(r io.Reader) (Image, string, error)](#Decode)
* [type NRGBA](#NRGBA)
  * [func NewNRGBA(r Rectangle) *NRGBA](#NewNRGBA)
  * [func (p *NRGBA) At(x, y int) color.Color](#NRGBA.At)
  * [func (p *NRGBA) Bounds() Rectangle](#NRGBA.Bounds)
  * [func (p *NRGBA) ColorModel() color.Model](#NRGBA.ColorModel)
  * [func (p *NRGBA) NRGBAAt(x, y int) color.NRGBA](#NRGBA.NRGBAAt)
  * [func (p *NRGBA) Opaque() bool](#NRGBA.Opaque)
  * [func (p *NRGBA) PixOffset(x, y int) int](#NRGBA.PixOffset)
  * [func (p *NRGBA) Set(x, y int, c color.Color)](#NRGBA.Set)
  * [func (p *NRGBA) SetNRGBA(x, y int, c color.NRGBA)](#NRGBA.SetNRGBA)
  * [func (p *NRGBA) SubImage(r Rectangle) Image](#NRGBA.SubImage)
* [type NRGBA64](#NRGBA64)
  * [func NewNRGBA64(r Rectangle) *NRGBA64](#NewNRGBA64)
  * [func (p *NRGBA64) At(x, y int) color.Color](#NRGBA64.At)
  * [func (p *NRGBA64) Bounds() Rectangle](#NRGBA64.Bounds)
  * [func (p *NRGBA64) ColorModel() color.Model](#NRGBA64.ColorModel)
  * [func (p *NRGBA64) NRGBA64At(x, y int) color.NRGBA64](#NRGBA64.NRGBA64At)
  * [func (p *NRGBA64) Opaque() bool](#NRGBA64.Opaque)
  * [func (p *NRGBA64) PixOffset(x, y int) int](#NRGBA64.PixOffset)
  * [func (p *NRGBA64) Set(x, y int, c color.Color)](#NRGBA64.Set)
  * [func (p *NRGBA64) SetNRGBA64(x, y int, c color.NRGBA64)](#NRGBA64.SetNRGBA64)
  * [func (p *NRGBA64) SubImage(r Rectangle) Image](#NRGBA64.SubImage)
* [type NYCbCrA](#NYCbCrA)
  * [func NewNYCbCrA(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *NYCbCrA](#NewNYCbCrA)
  * [func (p *NYCbCrA) AOffset(x, y int) int](#NYCbCrA.AOffset)
  * [func (p *NYCbCrA) At(x, y int) color.Color](#NYCbCrA.At)
  * [func (p *NYCbCrA) ColorModel() color.Model](#NYCbCrA.ColorModel)
  * [func (p *NYCbCrA) NYCbCrAAt(x, y int) color.NYCbCrA](#NYCbCrA.NYCbCrAAt)
  * [func (p *NYCbCrA) Opaque() bool](#NYCbCrA.Opaque)
  * [func (p *NYCbCrA) SubImage(r Rectangle) Image](#NYCbCrA.SubImage)
* [type Paletted](#Paletted)
  * [func NewPaletted(r Rectangle, p color.Palette) *Paletted](#NewPaletted)
  * [func (p *Paletted) At(x, y int) color.Color](#Paletted.At)
  * [func (p *Paletted) Bounds() Rectangle](#Paletted.Bounds)
  * [func (p *Paletted) ColorIndexAt(x, y int) uint8](#Paletted.ColorIndexAt)
  * [func (p *Paletted) ColorModel() color.Model](#Paletted.ColorModel)
  * [func (p *Paletted) Opaque() bool](#Paletted.Opaque)
  * [func (p *Paletted) PixOffset(x, y int) int](#Paletted.PixOffset)
  * [func (p *Paletted) Set(x, y int, c color.Color)](#Paletted.Set)
  * [func (p *Paletted) SetColorIndex(x, y int, index uint8)](#Paletted.SetColorIndex)
  * [func (p *Paletted) SubImage(r Rectangle) Image](#Paletted.SubImage)
* [type PalettedImage](#PalettedImage)
* [type Point](#Point)
  * [func Pt(X, Y int) Point](#Pt)
  * [func (p Point) Add(q Point) Point](#Point.Add)
  * [func (p Point) Div(k int) Point](#Point.Div)
  * [func (p Point) Eq(q Point) bool](#Point.Eq)
  * [func (p Point) In(r Rectangle) bool](#Point.In)
  * [func (p Point) Mod(r Rectangle) Point](#Point.Mod)
  * [func (p Point) Mul(k int) Point](#Point.Mul)
  * [func (p Point) String() string](#Point.String)
  * [func (p Point) Sub(q Point) Point](#Point.Sub)
* [type RGBA](#RGBA)
  * [func NewRGBA(r Rectangle) *RGBA](#NewRGBA)
  * [func (p *RGBA) At(x, y int) color.Color](#RGBA.At)
  * [func (p *RGBA) Bounds() Rectangle](#RGBA.Bounds)
  * [func (p *RGBA) ColorModel() color.Model](#RGBA.ColorModel)
  * [func (p *RGBA) Opaque() bool](#RGBA.Opaque)
  * [func (p *RGBA) PixOffset(x, y int) int](#RGBA.PixOffset)
  * [func (p *RGBA) RGBAAt(x, y int) color.RGBA](#RGBA.RGBAAt)
  * [func (p *RGBA) Set(x, y int, c color.Color)](#RGBA.Set)
  * [func (p *RGBA) SetRGBA(x, y int, c color.RGBA)](#RGBA.SetRGBA)
  * [func (p *RGBA) SubImage(r Rectangle) Image](#RGBA.SubImage)
* [type RGBA64](#RGBA64)
  * [func NewRGBA64(r Rectangle) *RGBA64](#NewRGBA64)
  * [func (p *RGBA64) At(x, y int) color.Color](#RGBA64.At)
  * [func (p *RGBA64) Bounds() Rectangle](#RGBA64.Bounds)
  * [func (p *RGBA64) ColorModel() color.Model](#RGBA64.ColorModel)
  * [func (p *RGBA64) Opaque() bool](#RGBA64.Opaque)
  * [func (p *RGBA64) PixOffset(x, y int) int](#RGBA64.PixOffset)
  * [func (p *RGBA64) RGBA64At(x, y int) color.RGBA64](#RGBA64.RGBA64At)
  * [func (p *RGBA64) Set(x, y int, c color.Color)](#RGBA64.Set)
  * [func (p *RGBA64) SetRGBA64(x, y int, c color.RGBA64)](#RGBA64.SetRGBA64)
  * [func (p *RGBA64) SubImage(r Rectangle) Image](#RGBA64.SubImage)
* [type Rectangle](#Rectangle)
  * [func Rect(x0, y0, x1, y1 int) Rectangle](#Rect)
  * [func (r Rectangle) Add(p Point) Rectangle](#Rectangle.Add)
  * [func (r Rectangle) At(x, y int) color.Color](#Rectangle.At)
  * [func (r Rectangle) Bounds() Rectangle](#Rectangle.Bounds)
  * [func (r Rectangle) Canon() Rectangle](#Rectangle.Canon)
  * [func (r Rectangle) ColorModel() color.Model](#Rectangle.ColorModel)
  * [func (r Rectangle) Dx() int](#Rectangle.Dx)
  * [func (r Rectangle) Dy() int](#Rectangle.Dy)
  * [func (r Rectangle) Empty() bool](#Rectangle.Empty)
  * [func (r Rectangle) Eq(s Rectangle) bool](#Rectangle.Eq)
  * [func (r Rectangle) In(s Rectangle) bool](#Rectangle.In)
  * [func (r Rectangle) Inset(n int) Rectangle](#Rectangle.Inset)
  * [func (r Rectangle) Intersect(s Rectangle) Rectangle](#Rectangle.Intersect)
  * [func (r Rectangle) Overlaps(s Rectangle) bool](#Rectangle.Overlaps)
  * [func (r Rectangle) Size() Point](#Rectangle.Size)
  * [func (r Rectangle) String() string](#Rectangle.String)
  * [func (r Rectangle) Sub(p Point) Rectangle](#Rectangle.Sub)
  * [func (r Rectangle) Union(s Rectangle) Rectangle](#Rectangle.Union)
* [type Uniform](#Uniform)
  * [func NewUniform(c color.Color) *Uniform](#NewUniform)
  * [func (c *Uniform) At(x, y int) color.Color](#Uniform.At)
  * [func (c *Uniform) Bounds() Rectangle](#Uniform.Bounds)
  * [func (c *Uniform) ColorModel() color.Model](#Uniform.ColorModel)
  * [func (c *Uniform) Convert(color.Color) color.Color](#Uniform.Convert)
  * [func (c *Uniform) Opaque() bool](#Uniform.Opaque)
  * [func (c *Uniform) RGBA() (r, g, b, a uint32)](#Uniform.RGBA)
* [type YCbCr](#YCbCr)
  * [func NewYCbCr(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *YCbCr](#NewYCbCr)
  * [func (p *YCbCr) At(x, y int) color.Color](#YCbCr.At)
  * [func (p *YCbCr) Bounds() Rectangle](#YCbCr.Bounds)
  * [func (p *YCbCr) COffset(x, y int) int](#YCbCr.COffset)
  * [func (p *YCbCr) ColorModel() color.Model](#YCbCr.ColorModel)
  * [func (p *YCbCr) Opaque() bool](#YCbCr.Opaque)
  * [func (p *YCbCr) SubImage(r Rectangle) Image](#YCbCr.SubImage)
  * [func (p *YCbCr) YCbCrAt(x, y int) color.YCbCr](#YCbCr.YCbCrAt)
  * [func (p *YCbCr) YOffset(x, y int) int](#YCbCr.YOffset)
* [type YCbCrSubsampleRatio](#YCbCrSubsampleRatio)
  * [func (s YCbCrSubsampleRatio) String() string](#YCbCrSubsampleRatio.String)


#### <a id="pkg-examples">Examples</a>
* [Package](#example_)
* [Package (DecodeConfig)](#example__decodeConfig)


#### <a id="pkg-files">Package files</a>
[format.go](https://golang.org/src/image/format.go) [geom.go](https://golang.org/src/image/geom.go) [image.go](https://golang.org/src/image/image.go) [names.go](https://golang.org/src/image/names.go) [ycbcr.go](https://golang.org/src/image/ycbcr.go) 




## <a id="pkg-variables">Variables</a>

<pre>var (
    <span class="comment">// Black is an opaque black uniform image.</span>
    <span id="Black">Black</span> = <a href="#NewUniform">NewUniform</a>(<a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Black">Black</a>)
    <span class="comment">// White is an opaque white uniform image.</span>
    <span id="White">White</span> = <a href="#NewUniform">NewUniform</a>(<a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#White">White</a>)
    <span class="comment">// Transparent is a fully transparent uniform image.</span>
    <span id="Transparent">Transparent</span> = <a href="#NewUniform">NewUniform</a>(<a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Transparent">Transparent</a>)
    <span class="comment">// Opaque is a fully opaque uniform image.</span>
    <span id="Opaque">Opaque</span> = <a href="#NewUniform">NewUniform</a>(<a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Opaque">Opaque</a>)
)</pre>ErrFormat indicates that decoding encountered an unknown format.


<pre>var <span id="ErrFormat">ErrFormat</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;image: unknown format&#34;)</pre>

## <a id="RegisterFormat">func</a> [RegisterFormat](https://golang.org/src/image/format.go?s=1069:1193#L27)
<pre>func RegisterFormat(name, magic <a href="/pkg/builtin/#string">string</a>, decode func(<a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (<a href="#Image">Image</a>, <a href="/pkg/builtin/#error">error</a>), decodeConfig func(<a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (<a href="#Config">Config</a>, <a href="/pkg/builtin/#error">error</a>))</pre>
RegisterFormat registers an image format for use by Decode.
Name is the name of the format, like "jpeg" or "png".
Magic is the magic prefix that identifies the format's encoding. The magic
string can contain "?" wildcards that each match any one byte.
Decode is the function that decodes the encoded image.
DecodeConfig is the function that decodes just its configuration.





## <a id="Alpha">type</a> [Alpha](https://golang.org/src/image/image.go?s=14210:14512#L484)
Alpha is an in-memory image whose At method returns color.Alpha values.


<pre>type Alpha struct {
<span id="Alpha.Pix"></span>    <span class="comment">// Pix holds the image&#39;s pixels, as alpha values. The pixel at</span>
    <span class="comment">// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1].</span>
    Pix []<a href="/pkg/builtin/#uint8">uint8</a>
<span id="Alpha.Stride"></span>    <span class="comment">// Stride is the Pix stride (in bytes) between vertically adjacent pixels.</span>
    Stride <a href="/pkg/builtin/#int">int</a>
<span id="Alpha.Rect"></span>    <span class="comment">// Rect is the image&#39;s bounds.</span>
    Rect <a href="#Rectangle">Rectangle</a>
}
</pre>









### <a id="NewAlpha">func</a> [NewAlpha](https://golang.org/src/image/image.go?s=16380:16413#L569)
<pre>func NewAlpha(r <a href="#Rectangle">Rectangle</a>) *<a href="#Alpha">Alpha</a></pre>
NewAlpha returns a new Alpha image with the given bounds.






### <a id="Alpha.AlphaAt">func</a> (\*Alpha) [AlphaAt](https://golang.org/src/image/image.go?s=14708:14753#L502)
<pre>func (p *<a href="#Alpha">Alpha</a>) AlphaAt(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Alpha">Alpha</a></pre>



### <a id="Alpha.At">func</a> (\*Alpha) [At](https://golang.org/src/image/image.go?s=14638:14678#L498)
<pre>func (p *<a href="#Alpha">Alpha</a>) At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a></pre>



### <a id="Alpha.Bounds">func</a> (\*Alpha) [Bounds](https://golang.org/src/image/image.go?s=14584:14618#L496)
<pre>func (p *<a href="#Alpha">Alpha</a>) Bounds() <a href="#Rectangle">Rectangle</a></pre>



### <a id="Alpha.ColorModel">func</a> (\*Alpha) [ColorModel](https://golang.org/src/image/image.go?s=14514:14554#L494)
<pre>func (p *<a href="#Alpha">Alpha</a>) ColorModel() <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Model">Model</a></pre>



### <a id="Alpha.Opaque">func</a> (\*Alpha) [Opaque](https://golang.org/src/image/image.go?s=16042:16071#L551)
<pre>func (p *<a href="#Alpha">Alpha</a>) Opaque() <a href="/pkg/builtin/#bool">bool</a></pre>
Opaque scans the entire image and reports whether it is fully opaque.




### <a id="Alpha.PixOffset">func</a> (\*Alpha) [PixOffset](https://golang.org/src/image/image.go?s=14974:15013#L512)
<pre>func (p *<a href="#Alpha">Alpha</a>) PixOffset(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
PixOffset returns the index of the first element of Pix that corresponds to
the pixel at (x, y).




### <a id="Alpha.Set">func</a> (\*Alpha) [Set](https://golang.org/src/image/image.go?s=15074:15118#L516)
<pre>func (p *<a href="#Alpha">Alpha</a>) Set(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a>)</pre>



### <a id="Alpha.SetAlpha">func</a> (\*Alpha) [SetAlpha](https://golang.org/src/image/image.go?s=15248:15297#L524)
<pre>func (p *<a href="#Alpha">Alpha</a>) SetAlpha(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Alpha">Alpha</a>)</pre>



### <a id="Alpha.SubImage">func</a> (\*Alpha) [SubImage](https://golang.org/src/image/image.go?s=15536:15579#L534)
<pre>func (p *<a href="#Alpha">Alpha</a>) SubImage(r <a href="#Rectangle">Rectangle</a>) <a href="#Image">Image</a></pre>
SubImage returns an image representing the portion of the image p visible
through r. The returned value shares pixels with the original image.




## <a id="Alpha16">type</a> [Alpha16](https://golang.org/src/image/image.go?s=16581:16906#L576)
Alpha16 is an in-memory image whose At method returns color.Alpha16 values.


<pre>type Alpha16 struct {
<span id="Alpha16.Pix"></span>    <span class="comment">// Pix holds the image&#39;s pixels, as alpha values in big-endian format. The pixel at</span>
    <span class="comment">// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*2].</span>
    Pix []<a href="/pkg/builtin/#uint8">uint8</a>
<span id="Alpha16.Stride"></span>    <span class="comment">// Stride is the Pix stride (in bytes) between vertically adjacent pixels.</span>
    Stride <a href="/pkg/builtin/#int">int</a>
<span id="Alpha16.Rect"></span>    <span class="comment">// Rect is the image&#39;s bounds.</span>
    Rect <a href="#Rectangle">Rectangle</a>
}
</pre>









### <a id="NewAlpha16">func</a> [NewAlpha16](https://golang.org/src/image/image.go?s=18972:19009#L664)
<pre>func NewAlpha16(r <a href="#Rectangle">Rectangle</a>) *<a href="#Alpha16">Alpha16</a></pre>
NewAlpha16 returns a new Alpha16 image with the given bounds.






### <a id="Alpha16.Alpha16At">func</a> (\*Alpha16) [Alpha16At](https://golang.org/src/image/image.go?s=17112:17163#L594)
<pre>func (p *<a href="#Alpha16">Alpha16</a>) Alpha16At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Alpha16">Alpha16</a></pre>



### <a id="Alpha16.At">func</a> (\*Alpha16) [At](https://golang.org/src/image/image.go?s=17038:17080#L590)
<pre>func (p *<a href="#Alpha16">Alpha16</a>) At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a></pre>



### <a id="Alpha16.Bounds">func</a> (\*Alpha16) [Bounds](https://golang.org/src/image/image.go?s=16982:17018#L588)
<pre>func (p *<a href="#Alpha16">Alpha16</a>) Bounds() <a href="#Rectangle">Rectangle</a></pre>



### <a id="Alpha16.ColorModel">func</a> (\*Alpha16) [ColorModel](https://golang.org/src/image/image.go?s=16908:16950#L586)
<pre>func (p *<a href="#Alpha16">Alpha16</a>) ColorModel() <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Model">Model</a></pre>



### <a id="Alpha16.Opaque">func</a> (\*Alpha16) [Opaque](https://golang.org/src/image/image.go?s=18599:18630#L646)
<pre>func (p *<a href="#Alpha16">Alpha16</a>) Opaque() <a href="/pkg/builtin/#bool">bool</a></pre>
Opaque scans the entire image and reports whether it is fully opaque.




### <a id="Alpha16.PixOffset">func</a> (\*Alpha16) [PixOffset](https://golang.org/src/image/image.go?s=17422:17463#L604)
<pre>func (p *<a href="#Alpha16">Alpha16</a>) PixOffset(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
PixOffset returns the index of the first element of Pix that corresponds to
the pixel at (x, y).




### <a id="Alpha16.Set">func</a> (\*Alpha16) [Set](https://golang.org/src/image/image.go?s=17524:17570#L608)
<pre>func (p *<a href="#Alpha16">Alpha16</a>) Set(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a>)</pre>



### <a id="Alpha16.SetAlpha16">func</a> (\*Alpha16) [SetAlpha16](https://golang.org/src/image/image.go?s=17754:17809#L618)
<pre>func (p *<a href="#Alpha16">Alpha16</a>) SetAlpha16(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Alpha16">Alpha16</a>)</pre>



### <a id="Alpha16.SubImage">func</a> (\*Alpha16) [SubImage](https://golang.org/src/image/image.go?s=18087:18132#L629)
<pre>func (p *<a href="#Alpha16">Alpha16</a>) SubImage(r <a href="#Rectangle">Rectangle</a>) <a href="#Image">Image</a></pre>
SubImage returns an image representing the portion of the image p visible
through r. The returned value shares pixels with the original image.




## <a id="CMYK">type</a> [CMYK](https://golang.org/src/image/image.go?s=23591:23896#L832)
CMYK is an in-memory image whose At method returns color.CMYK values.


<pre>type CMYK struct {
<span id="CMYK.Pix"></span>    <span class="comment">// Pix holds the image&#39;s pixels, in C, M, Y, K order. The pixel at</span>
    <span class="comment">// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].</span>
    Pix []<a href="/pkg/builtin/#uint8">uint8</a>
<span id="CMYK.Stride"></span>    <span class="comment">// Stride is the Pix stride (in bytes) between vertically adjacent pixels.</span>
    Stride <a href="/pkg/builtin/#int">int</a>
<span id="CMYK.Rect"></span>    <span class="comment">// Rect is the image&#39;s bounds.</span>
    Rect <a href="#Rectangle">Rectangle</a>
}
</pre>









### <a id="NewCMYK">func</a> [NewCMYK](https://golang.org/src/image/image.go?s=25894:25925#L914)
<pre>func NewCMYK(r <a href="#Rectangle">Rectangle</a>) *<a href="#CMYK">CMYK</a></pre>
NewCMYK returns a new CMYK image with the given bounds.






### <a id="CMYK.At">func</a> (\*CMYK) [At](https://golang.org/src/image/image.go?s=24019:24058#L846)
<pre>func (p *<a href="#CMYK">CMYK</a>) At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a></pre>



### <a id="CMYK.Bounds">func</a> (\*CMYK) [Bounds](https://golang.org/src/image/image.go?s=23966:23999#L844)
<pre>func (p *<a href="#CMYK">CMYK</a>) Bounds() <a href="#Rectangle">Rectangle</a></pre>



### <a id="CMYK.CMYKAt">func</a> (\*CMYK) [CMYKAt](https://golang.org/src/image/image.go?s=24087:24129#L850)
<pre>func (p *<a href="#CMYK">CMYK</a>) CMYKAt(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#CMYK">CMYK</a></pre>



### <a id="CMYK.ColorModel">func</a> (\*CMYK) [ColorModel](https://golang.org/src/image/image.go?s=23898:23937#L842)
<pre>func (p *<a href="#CMYK">CMYK</a>) ColorModel() <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Model">Model</a></pre>



### <a id="CMYK.Opaque">func</a> (\*CMYK) [Opaque](https://golang.org/src/image/image.go?s=25788:25816#L909)
<pre>func (p *<a href="#CMYK">CMYK</a>) Opaque() <a href="/pkg/builtin/#bool">bool</a></pre>
Opaque scans the entire image and reports whether it is fully opaque.




### <a id="CMYK.PixOffset">func</a> (\*CMYK) [PixOffset](https://golang.org/src/image/image.go?s=24459:24497#L861)
<pre>func (p *<a href="#CMYK">CMYK</a>) PixOffset(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
PixOffset returns the index of the first element of Pix that corresponds to
the pixel at (x, y).




### <a id="CMYK.Set">func</a> (\*CMYK) [Set](https://golang.org/src/image/image.go?s=24558:24601#L865)
<pre>func (p *<a href="#CMYK">CMYK</a>) Set(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a>)</pre>



### <a id="CMYK.SetCMYK">func</a> (\*CMYK) [SetCMYK](https://golang.org/src/image/image.go?s=24871:24917#L878)
<pre>func (p *<a href="#CMYK">CMYK</a>) SetCMYK(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#CMYK">CMYK</a>)</pre>



### <a id="CMYK.SubImage">func</a> (\*CMYK) [SubImage](https://golang.org/src/image/image.go?s=25285:25327#L892)
<pre>func (p *<a href="#CMYK">CMYK</a>) SubImage(r <a href="#Rectangle">Rectangle</a>) <a href="#Image">Image</a></pre>
SubImage returns an image representing the portion of the image p visible
through r. The returned value shares pixels with the original image.




## <a id="Config">type</a> [Config](https://golang.org/src/image/image.go?s=1134:1202#L19)
Config holds an image's color model and dimensions.


<pre>type Config struct {
<span id="Config.ColorModel"></span>    ColorModel    <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Model">Model</a>
<span id="Config.Width"></span>    Width, Height <a href="/pkg/builtin/#int">int</a>
}
</pre>









### <a id="DecodeConfig">func</a> [DecodeConfig](https://golang.org/src/image/format.go?s=2883:2937#L91)
<pre>func DecodeConfig(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (<a href="#Config">Config</a>, <a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DecodeConfig decodes the color model and dimensions of an image that has
been encoded in a registered format. The string returned is the format name
used during format registration. Format registration is typically done by
an init function in the codec-specific package.






## <a id="Gray">type</a> [Gray](https://golang.org/src/image/image.go?s=19173:19473#L671)
Gray is an in-memory image whose At method returns color.Gray values.


<pre>type Gray struct {
<span id="Gray.Pix"></span>    <span class="comment">// Pix holds the image&#39;s pixels, as gray values. The pixel at</span>
    <span class="comment">// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1].</span>
    Pix []<a href="/pkg/builtin/#uint8">uint8</a>
<span id="Gray.Stride"></span>    <span class="comment">// Stride is the Pix stride (in bytes) between vertically adjacent pixels.</span>
    Stride <a href="/pkg/builtin/#int">int</a>
<span id="Gray.Rect"></span>    <span class="comment">// Rect is the image&#39;s bounds.</span>
    Rect <a href="#Rectangle">Rectangle</a>
}
</pre>









### <a id="NewGray">func</a> [NewGray](https://golang.org/src/image/image.go?s=21089:21120#L743)
<pre>func NewGray(r <a href="#Rectangle">Rectangle</a>) *<a href="#Gray">Gray</a></pre>
NewGray returns a new Gray image with the given bounds.






### <a id="Gray.At">func</a> (\*Gray) [At](https://golang.org/src/image/image.go?s=19596:19635#L685)
<pre>func (p *<a href="#Gray">Gray</a>) At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a></pre>



### <a id="Gray.Bounds">func</a> (\*Gray) [Bounds](https://golang.org/src/image/image.go?s=19543:19576#L683)
<pre>func (p *<a href="#Gray">Gray</a>) Bounds() <a href="#Rectangle">Rectangle</a></pre>



### <a id="Gray.ColorModel">func</a> (\*Gray) [ColorModel](https://golang.org/src/image/image.go?s=19475:19514#L681)
<pre>func (p *<a href="#Gray">Gray</a>) ColorModel() <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Model">Model</a></pre>



### <a id="Gray.GrayAt">func</a> (\*Gray) [GrayAt](https://golang.org/src/image/image.go?s=19664:19706#L689)
<pre>func (p *<a href="#Gray">Gray</a>) GrayAt(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Gray">Gray</a></pre>



### <a id="Gray.Opaque">func</a> (\*Gray) [Opaque](https://golang.org/src/image/image.go?s=20983:21011#L738)
<pre>func (p *<a href="#Gray">Gray</a>) Opaque() <a href="/pkg/builtin/#bool">bool</a></pre>
Opaque scans the entire image and reports whether it is fully opaque.




### <a id="Gray.PixOffset">func</a> (\*Gray) [PixOffset](https://golang.org/src/image/image.go?s=19925:19963#L699)
<pre>func (p *<a href="#Gray">Gray</a>) PixOffset(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
PixOffset returns the index of the first element of Pix that corresponds to
the pixel at (x, y).




### <a id="Gray.Set">func</a> (\*Gray) [Set](https://golang.org/src/image/image.go?s=20024:20067#L703)
<pre>func (p *<a href="#Gray">Gray</a>) Set(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a>)</pre>



### <a id="Gray.SetGray">func</a> (\*Gray) [SetGray](https://golang.org/src/image/image.go?s=20195:20241#L711)
<pre>func (p *<a href="#Gray">Gray</a>) SetGray(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Gray">Gray</a>)</pre>



### <a id="Gray.SubImage">func</a> (\*Gray) [SubImage](https://golang.org/src/image/image.go?s=20480:20522#L721)
<pre>func (p *<a href="#Gray">Gray</a>) SubImage(r <a href="#Rectangle">Rectangle</a>) <a href="#Image">Image</a></pre>
SubImage returns an image representing the portion of the image p visible
through r. The returned value shares pixels with the original image.




## <a id="Gray16">type</a> [Gray16](https://golang.org/src/image/image.go?s=21285:21608#L750)
Gray16 is an in-memory image whose At method returns color.Gray16 values.


<pre>type Gray16 struct {
<span id="Gray16.Pix"></span>    <span class="comment">// Pix holds the image&#39;s pixels, as gray values in big-endian format. The pixel at</span>
    <span class="comment">// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*2].</span>
    Pix []<a href="/pkg/builtin/#uint8">uint8</a>
<span id="Gray16.Stride"></span>    <span class="comment">// Stride is the Pix stride (in bytes) between vertically adjacent pixels.</span>
    Stride <a href="/pkg/builtin/#int">int</a>
<span id="Gray16.Rect"></span>    <span class="comment">// Rect is the image&#39;s bounds.</span>
    Rect <a href="#Rectangle">Rectangle</a>
}
</pre>









### <a id="NewGray16">func</a> [NewGray16](https://golang.org/src/image/image.go?s=23393:23428#L825)
<pre>func NewGray16(r <a href="#Rectangle">Rectangle</a>) *<a href="#Gray16">Gray16</a></pre>
NewGray16 returns a new Gray16 image with the given bounds.






### <a id="Gray16.At">func</a> (\*Gray16) [At](https://golang.org/src/image/image.go?s=21737:21778#L764)
<pre>func (p *<a href="#Gray16">Gray16</a>) At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a></pre>



### <a id="Gray16.Bounds">func</a> (\*Gray16) [Bounds](https://golang.org/src/image/image.go?s=21682:21717#L762)
<pre>func (p *<a href="#Gray16">Gray16</a>) Bounds() <a href="#Rectangle">Rectangle</a></pre>



### <a id="Gray16.ColorModel">func</a> (\*Gray16) [ColorModel](https://golang.org/src/image/image.go?s=21610:21651#L760)
<pre>func (p *<a href="#Gray16">Gray16</a>) ColorModel() <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Model">Model</a></pre>



### <a id="Gray16.Gray16At">func</a> (\*Gray16) [Gray16At](https://golang.org/src/image/image.go?s=21809:21857#L768)
<pre>func (p *<a href="#Gray16">Gray16</a>) Gray16At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Gray16">Gray16</a></pre>



### <a id="Gray16.Opaque">func</a> (\*Gray16) [Opaque](https://golang.org/src/image/image.go?s=23281:23311#L820)
<pre>func (p *<a href="#Gray16">Gray16</a>) Opaque() <a href="/pkg/builtin/#bool">bool</a></pre>
Opaque scans the entire image and reports whether it is fully opaque.




### <a id="Gray16.PixOffset">func</a> (\*Gray16) [PixOffset](https://golang.org/src/image/image.go?s=22114:22154#L778)
<pre>func (p *<a href="#Gray16">Gray16</a>) PixOffset(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
PixOffset returns the index of the first element of Pix that corresponds to
the pixel at (x, y).




### <a id="Gray16.Set">func</a> (\*Gray16) [Set](https://golang.org/src/image/image.go?s=22215:22260#L782)
<pre>func (p *<a href="#Gray16">Gray16</a>) Set(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a>)</pre>



### <a id="Gray16.SetGray16">func</a> (\*Gray16) [SetGray16](https://golang.org/src/image/image.go?s=22442:22494#L792)
<pre>func (p *<a href="#Gray16">Gray16</a>) SetGray16(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Gray16">Gray16</a>)</pre>



### <a id="Gray16.SubImage">func</a> (\*Gray16) [SubImage](https://golang.org/src/image/image.go?s=22772:22816#L803)
<pre>func (p *<a href="#Gray16">Gray16</a>) SubImage(r <a href="#Rectangle">Rectangle</a>) <a href="#Image">Image</a></pre>
SubImage returns an image representing the portion of the image p visible
through r. The returned value shares pixels with the original image.




## <a id="Image">type</a> [Image](https://golang.org/src/image/image.go?s=1293:1769#L26)
Image is a finite rectangular grid of color.Color values taken from a color
model.


<pre>type Image interface {
    <span class="comment">// ColorModel returns the Image&#39;s color model.</span>
    ColorModel() <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Model">Model</a>
    <span class="comment">// Bounds returns the domain for which At can return non-zero color.</span>
    <span class="comment">// The bounds do not necessarily contain the point (0, 0).</span>
    Bounds() <a href="#Rectangle">Rectangle</a>
    <span class="comment">// At returns the color of the pixel at (x, y).</span>
    <span class="comment">// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.</span>
    <span class="comment">// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.</span>
    At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a>
}</pre>









### <a id="Decode">func</a> [Decode](https://golang.org/src/image/format.go?s=2412:2459#L77)
<pre>func Decode(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (<a href="#Image">Image</a>, <a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Decode decodes an image that has been encoded in a registered format.
The string returned is the format name used during format registration.
Format registration is typically done by an init function in the codec-
specific package.






## <a id="NRGBA">type</a> [NRGBA](https://golang.org/src/image/image.go?s=8222:8528#L267)
NRGBA is an in-memory image whose At method returns color.NRGBA values.


<pre>type NRGBA struct {
<span id="NRGBA.Pix"></span>    <span class="comment">// Pix holds the image&#39;s pixels, in R, G, B, A order. The pixel at</span>
    <span class="comment">// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].</span>
    Pix []<a href="/pkg/builtin/#uint8">uint8</a>
<span id="NRGBA.Stride"></span>    <span class="comment">// Stride is the Pix stride (in bytes) between vertically adjacent pixels.</span>
    Stride <a href="/pkg/builtin/#int">int</a>
<span id="NRGBA.Rect"></span>    <span class="comment">// Rect is the image&#39;s bounds.</span>
    Rect <a href="#Rectangle">Rectangle</a>
}
</pre>









### <a id="NewNRGBA">func</a> [NewNRGBA](https://golang.org/src/image/image.go?s=10783:10816#L362)
<pre>func NewNRGBA(r <a href="#Rectangle">Rectangle</a>) *<a href="#NRGBA">NRGBA</a></pre>
NewNRGBA returns a new NRGBA image with the given bounds.






### <a id="NRGBA.At">func</a> (\*NRGBA) [At](https://golang.org/src/image/image.go?s=8654:8694#L281)
<pre>func (p *<a href="#NRGBA">NRGBA</a>) At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a></pre>



### <a id="NRGBA.Bounds">func</a> (\*NRGBA) [Bounds](https://golang.org/src/image/image.go?s=8600:8634#L279)
<pre>func (p *<a href="#NRGBA">NRGBA</a>) Bounds() <a href="#Rectangle">Rectangle</a></pre>



### <a id="NRGBA.ColorModel">func</a> (\*NRGBA) [ColorModel](https://golang.org/src/image/image.go?s=8530:8570#L277)
<pre>func (p *<a href="#NRGBA">NRGBA</a>) ColorModel() <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Model">Model</a></pre>



### <a id="NRGBA.NRGBAAt">func</a> (\*NRGBA) [NRGBAAt](https://golang.org/src/image/image.go?s=8724:8769#L285)
<pre>func (p *<a href="#NRGBA">NRGBA</a>) NRGBAAt(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#NRGBA">NRGBA</a></pre>



### <a id="NRGBA.Opaque">func</a> (\*NRGBA) [Opaque](https://golang.org/src/image/image.go?s=10440:10469#L344)
<pre>func (p *<a href="#NRGBA">NRGBA</a>) Opaque() <a href="/pkg/builtin/#bool">bool</a></pre>
Opaque scans the entire image and reports whether it is fully opaque.




### <a id="NRGBA.PixOffset">func</a> (\*NRGBA) [PixOffset](https://golang.org/src/image/image.go?s=9101:9140#L296)
<pre>func (p *<a href="#NRGBA">NRGBA</a>) PixOffset(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
PixOffset returns the index of the first element of Pix that corresponds to
the pixel at (x, y).




### <a id="NRGBA.Set">func</a> (\*NRGBA) [Set](https://golang.org/src/image/image.go?s=9201:9245#L300)
<pre>func (p *<a href="#NRGBA">NRGBA</a>) Set(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a>)</pre>



### <a id="NRGBA.SetNRGBA">func</a> (\*NRGBA) [SetNRGBA](https://golang.org/src/image/image.go?s=9517:9566#L313)
<pre>func (p *<a href="#NRGBA">NRGBA</a>) SetNRGBA(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#NRGBA">NRGBA</a>)</pre>



### <a id="NRGBA.SubImage">func</a> (\*NRGBA) [SubImage](https://golang.org/src/image/image.go?s=9934:9977#L327)
<pre>func (p *<a href="#NRGBA">NRGBA</a>) SubImage(r <a href="#Rectangle">Rectangle</a>) <a href="#Image">Image</a></pre>
SubImage returns an image representing the portion of the image p visible
through r. The returned value shares pixels with the original image.




## <a id="NRGBA64">type</a> [NRGBA64](https://golang.org/src/image/image.go?s=10984:11314#L369)
NRGBA64 is an in-memory image whose At method returns color.NRGBA64 values.


<pre>type NRGBA64 struct {
<span id="NRGBA64.Pix"></span>    <span class="comment">// Pix holds the image&#39;s pixels, in R, G, B, A order and big-endian format. The pixel at</span>
    <span class="comment">// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*8].</span>
    Pix []<a href="/pkg/builtin/#uint8">uint8</a>
<span id="NRGBA64.Stride"></span>    <span class="comment">// Stride is the Pix stride (in bytes) between vertically adjacent pixels.</span>
    Stride <a href="/pkg/builtin/#int">int</a>
<span id="NRGBA64.Rect"></span>    <span class="comment">// Rect is the image&#39;s bounds.</span>
    Rect <a href="#Rectangle">Rectangle</a>
}
</pre>









### <a id="NewNRGBA64">func</a> [NewNRGBA64](https://golang.org/src/image/image.go?s=14007:14044#L477)
<pre>func NewNRGBA64(r <a href="#Rectangle">Rectangle</a>) *<a href="#NRGBA64">NRGBA64</a></pre>
NewNRGBA64 returns a new NRGBA64 image with the given bounds.






### <a id="NRGBA64.At">func</a> (\*NRGBA64) [At](https://golang.org/src/image/image.go?s=11446:11488#L383)
<pre>func (p *<a href="#NRGBA64">NRGBA64</a>) At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a></pre>



### <a id="NRGBA64.Bounds">func</a> (\*NRGBA64) [Bounds](https://golang.org/src/image/image.go?s=11390:11426#L381)
<pre>func (p *<a href="#NRGBA64">NRGBA64</a>) Bounds() <a href="#Rectangle">Rectangle</a></pre>



### <a id="NRGBA64.ColorModel">func</a> (\*NRGBA64) [ColorModel](https://golang.org/src/image/image.go?s=11316:11358#L379)
<pre>func (p *<a href="#NRGBA64">NRGBA64</a>) ColorModel() <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Model">Model</a></pre>



### <a id="NRGBA64.NRGBA64At">func</a> (\*NRGBA64) [NRGBA64At](https://golang.org/src/image/image.go?s=11520:11571#L387)
<pre>func (p *<a href="#NRGBA64">NRGBA64</a>) NRGBA64At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#NRGBA64">NRGBA64</a></pre>



### <a id="NRGBA64.Opaque">func</a> (\*NRGBA64) [Opaque](https://golang.org/src/image/image.go?s=13634:13665#L459)
<pre>func (p *<a href="#NRGBA64">NRGBA64</a>) Opaque() <a href="/pkg/builtin/#bool">bool</a></pre>
Opaque scans the entire image and reports whether it is fully opaque.




### <a id="NRGBA64.PixOffset">func</a> (\*NRGBA64) [PixOffset](https://golang.org/src/image/image.go?s=12023:12064#L403)
<pre>func (p *<a href="#NRGBA64">NRGBA64</a>) PixOffset(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
PixOffset returns the index of the first element of Pix that corresponds to
the pixel at (x, y).




### <a id="NRGBA64.Set">func</a> (\*NRGBA64) [Set](https://golang.org/src/image/image.go?s=12125:12171#L407)
<pre>func (p *<a href="#NRGBA64">NRGBA64</a>) Set(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a>)</pre>



### <a id="NRGBA64.SetNRGBA64">func</a> (\*NRGBA64) [SetNRGBA64](https://golang.org/src/image/image.go?s=12575:12630#L424)
<pre>func (p *<a href="#NRGBA64">NRGBA64</a>) SetNRGBA64(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#NRGBA64">NRGBA64</a>)</pre>



### <a id="NRGBA64.SubImage">func</a> (\*NRGBA64) [SubImage](https://golang.org/src/image/image.go?s=13122:13167#L442)
<pre>func (p *<a href="#NRGBA64">NRGBA64</a>) SubImage(r <a href="#Rectangle">Rectangle</a>) <a href="#Image">Image</a></pre>
SubImage returns an image representing the portion of the image p visible
through r. The returned value shares pixels with the original image.




## <a id="NYCbCrA">type</a> [NYCbCrA](https://golang.org/src/image/ycbcr.go?s=5458:5518#L179)
NYCbCrA is an in-memory image of non-alpha-premultiplied Y'CbCr-with-alpha
colors. A and AStride are analogous to the Y and YStride fields of the
embedded YCbCr.


<pre>type NYCbCrA struct {
    <a href="#YCbCr">YCbCr</a>
<span id="NYCbCrA.A"></span>    A       []<a href="/pkg/builtin/#uint8">uint8</a>
<span id="NYCbCrA.AStride"></span>    AStride <a href="/pkg/builtin/#int">int</a>
}
</pre>









### <a id="NewNYCbCrA">func</a> [NewNYCbCrA](https://golang.org/src/image/ycbcr.go?s=7546:7619#L268)
<pre>func NewNYCbCrA(r <a href="#Rectangle">Rectangle</a>, subsampleRatio <a href="#YCbCrSubsampleRatio">YCbCrSubsampleRatio</a>) *<a href="#NYCbCrA">NYCbCrA</a></pre>
NewNYCbCrA returns a new NYCbCrA image with the given bounds and subsample
ratio.






### <a id="NYCbCrA.AOffset">func</a> (\*NYCbCrA) [AOffset](https://golang.org/src/image/ycbcr.go?s=6067:6106#L212)
<pre>func (p *<a href="#NYCbCrA">NYCbCrA</a>) AOffset(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
AOffset returns the index of the first element of A that corresponds to the
pixel at (x, y).




### <a id="NYCbCrA.At">func</a> (\*NYCbCrA) [At](https://golang.org/src/image/ycbcr.go?s=5595:5637#L189)
<pre>func (p *<a href="#NYCbCrA">NYCbCrA</a>) At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a></pre>



### <a id="NYCbCrA.ColorModel">func</a> (\*NYCbCrA) [ColorModel](https://golang.org/src/image/ycbcr.go?s=5520:5562#L185)
<pre>func (p *<a href="#NYCbCrA">NYCbCrA</a>) ColorModel() <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Model">Model</a></pre>



### <a id="NYCbCrA.NYCbCrAAt">func</a> (\*NYCbCrA) [NYCbCrAAt](https://golang.org/src/image/ycbcr.go?s=5669:5720#L193)
<pre>func (p *<a href="#NYCbCrA">NYCbCrA</a>) NYCbCrAAt(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#NYCbCrA">NYCbCrA</a></pre>



### <a id="NYCbCrA.Opaque">func</a> (\*NYCbCrA) [Opaque](https://golang.org/src/image/ycbcr.go?s=7180:7211#L249)
<pre>func (p *<a href="#NYCbCrA">NYCbCrA</a>) Opaque() <a href="/pkg/builtin/#bool">bool</a></pre>
Opaque scans the entire image and reports whether it is fully opaque.




### <a id="NYCbCrA.SubImage">func</a> (\*NYCbCrA) [SubImage](https://golang.org/src/image/ycbcr.go?s=6317:6362#L218)
<pre>func (p *<a href="#NYCbCrA">NYCbCrA</a>) SubImage(r <a href="#Rectangle">Rectangle</a>) <a href="#Image">Image</a></pre>
SubImage returns an image representing the portion of the image p visible
through r. The returned value shares pixels with the original image.




## <a id="Paletted">type</a> [Paletted](https://golang.org/src/image/image.go?s=26086:26453#L921)
Paletted is an in-memory image of uint8 indices into a given palette.


<pre>type Paletted struct {
<span id="Paletted.Pix"></span>    <span class="comment">// Pix holds the image&#39;s pixels, as palette indices. The pixel at</span>
    <span class="comment">// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1].</span>
    Pix []<a href="/pkg/builtin/#uint8">uint8</a>
<span id="Paletted.Stride"></span>    <span class="comment">// Stride is the Pix stride (in bytes) between vertically adjacent pixels.</span>
    Stride <a href="/pkg/builtin/#int">int</a>
<span id="Paletted.Rect"></span>    <span class="comment">// Rect is the image&#39;s bounds.</span>
    Rect <a href="#Rectangle">Rectangle</a>
<span id="Paletted.Palette"></span>    <span class="comment">// Palette is the image&#39;s palette.</span>
    Palette <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Palette">Palette</a>
}
</pre>









### <a id="NewPaletted">func</a> [NewPaletted](https://golang.org/src/image/image.go?s=28630:28686#L1024)
<pre>func NewPaletted(r <a href="#Rectangle">Rectangle</a>, p <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Palette">Palette</a>) *<a href="#Paletted">Paletted</a></pre>
NewPaletted returns a new Paletted image with the given width, height and
palette.






### <a id="Paletted.At">func</a> (\*Paletted) [At](https://golang.org/src/image/image.go?s=26578:26621#L937)
<pre>func (p *<a href="#Paletted">Paletted</a>) At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a></pre>



### <a id="Paletted.Bounds">func</a> (\*Paletted) [Bounds](https://golang.org/src/image/image.go?s=26521:26558#L935)
<pre>func (p *<a href="#Paletted">Paletted</a>) Bounds() <a href="#Rectangle">Rectangle</a></pre>



### <a id="Paletted.ColorIndexAt">func</a> (\*Paletted) [ColorIndexAt](https://golang.org/src/image/image.go?s=27143:27190#L962)
<pre>func (p *<a href="#Paletted">Paletted</a>) ColorIndexAt(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#uint8">uint8</a></pre>



### <a id="Paletted.ColorModel">func</a> (\*Paletted) [ColorModel](https://golang.org/src/image/image.go?s=26455:26498#L933)
<pre>func (p *<a href="#Paletted">Paletted</a>) ColorModel() <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Model">Model</a></pre>



### <a id="Paletted.Opaque">func</a> (\*Paletted) [Opaque](https://golang.org/src/image/image.go?s=28163:28195#L1000)
<pre>func (p *<a href="#Paletted">Paletted</a>) Opaque() <a href="/pkg/builtin/#bool">bool</a></pre>
Opaque scans the entire image and reports whether it is fully opaque.




### <a id="Paletted.PixOffset">func</a> (\*Paletted) [PixOffset](https://golang.org/src/image/image.go?s=26881:26923#L950)
<pre>func (p *<a href="#Paletted">Paletted</a>) PixOffset(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
PixOffset returns the index of the first element of Pix that corresponds to
the pixel at (x, y).




### <a id="Paletted.Set">func</a> (\*Paletted) [Set](https://golang.org/src/image/image.go?s=26984:27031#L954)
<pre>func (p *<a href="#Paletted">Paletted</a>) Set(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a>)</pre>



### <a id="Paletted.SetColorIndex">func</a> (\*Paletted) [SetColorIndex](https://golang.org/src/image/image.go?s=27283:27338#L970)
<pre>func (p *<a href="#Paletted">Paletted</a>) SetColorIndex(x, y <a href="/pkg/builtin/#int">int</a>, index <a href="/pkg/builtin/#uint8">uint8</a>)</pre>



### <a id="Paletted.SubImage">func</a> (\*Paletted) [SubImage](https://golang.org/src/image/image.go?s=27579:27625#L980)
<pre>func (p *<a href="#Paletted">Paletted</a>) SubImage(r <a href="#Rectangle">Rectangle</a>) <a href="#Image">Image</a></pre>
SubImage returns an image representing the portion of the image p visible
through r. The returned value shares pixels with the original image.




## <a id="PalettedImage">type</a> [PalettedImage](https://golang.org/src/image/image.go?s=2079:2215#L43)
PalettedImage is an image whose colors may come from a limited palette.
If m is a PalettedImage and m.ColorModel() returns a color.Palette p,
then m.At(x, y) should be equivalent to p[m.ColorIndexAt(x, y)]. If m's
color model is not a color.Palette, then ColorIndexAt's behavior is
undefined.


<pre>type PalettedImage interface {
    <span class="comment">// ColorIndexAt returns the palette index of the pixel at (x, y).</span>
    ColorIndexAt(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#uint8">uint8</a>
    <a href="#Image">Image</a>
}</pre>











## <a id="Point">type</a> [Point](https://golang.org/src/image/geom.go?s=286:317#L3)
A Point is an X, Y coordinate pair. The axes increase right and down.


<pre>type Point struct {
<span id="Point.X"></span>    X, Y <a href="/pkg/builtin/#int">int</a>
}
</pre>




ZP is the zero Point.

Deprecated: Use a literal image.Point{} instead.


<pre>var <span id="ZP">ZP</span> <a href="#Point">Point</a></pre>





### <a id="Pt">func</a> [Pt](https://golang.org/src/image/geom.go?s=1579:1602#L65)
<pre>func Pt(X, Y <a href="/pkg/builtin/#int">int</a>) <a href="#Point">Point</a></pre>
Pt is shorthand for Point{X, Y}.






### <a id="Point.Add">func</a> (Point) [Add](https://golang.org/src/image/geom.go?s=511:544#L13)
<pre>func (p <a href="#Point">Point</a>) Add(q <a href="#Point">Point</a>) <a href="#Point">Point</a></pre>
Add returns the vector p+q.




### <a id="Point.Div">func</a> (Point) [Div](https://golang.org/src/image/geom.go?s=823:854#L28)
<pre>func (p <a href="#Point">Point</a>) Div(k <a href="/pkg/builtin/#int">int</a>) <a href="#Point">Point</a></pre>
Div returns the vector p/k.




### <a id="Point.Eq">func</a> (Point) [Eq](https://golang.org/src/image/geom.go?s=1397:1428#L55)
<pre>func (p <a href="#Point">Point</a>) Eq(q <a href="#Point">Point</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Eq reports whether p and q are equal.




### <a id="Point.In">func</a> (Point) [In](https://golang.org/src/image/geom.go?s=925:960#L33)
<pre>func (p <a href="#Point">Point</a>) In(r <a href="#Rectangle">Rectangle</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
In reports whether p is in r.




### <a id="Point.Mod">func</a> (Point) [Mod](https://golang.org/src/image/geom.go?s=1164:1201#L40)
<pre>func (p <a href="#Point">Point</a>) Mod(r <a href="#Rectangle">Rectangle</a>) <a href="#Point">Point</a></pre>
Mod returns the point q in r such that p.X-q.X is a multiple of r's width
and p.Y-q.Y is a multiple of r's height.




### <a id="Point.Mul">func</a> (Point) [Mul](https://golang.org/src/image/geom.go?s=723:754#L23)
<pre>func (p <a href="#Point">Point</a>) Mul(k <a href="/pkg/builtin/#int">int</a>) <a href="#Point">Point</a></pre>
Mul returns the vector p*k.




### <a id="Point.String">func</a> (Point) [String](https://golang.org/src/image/geom.go?s=380:410#L8)
<pre>func (p <a href="#Point">Point</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns a string representation of p like "(3,4)".




### <a id="Point.Sub">func</a> (Point) [Sub](https://golang.org/src/image/geom.go?s=617:650#L18)
<pre>func (p <a href="#Point">Point</a>) Sub(q <a href="#Point">Point</a>) <a href="#Point">Point</a></pre>
Sub returns the vector p-q.




## <a id="RGBA">type</a> [RGBA](https://golang.org/src/image/image.go?s=2290:2595#L50)
RGBA is an in-memory image whose At method returns color.RGBA values.


<pre>type RGBA struct {
<span id="RGBA.Pix"></span>    <span class="comment">// Pix holds the image&#39;s pixels, in R, G, B, A order. The pixel at</span>
    <span class="comment">// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].</span>
    Pix []<a href="/pkg/builtin/#uint8">uint8</a>
<span id="RGBA.Stride"></span>    <span class="comment">// Stride is the Pix stride (in bytes) between vertically adjacent pixels.</span>
    Stride <a href="/pkg/builtin/#int">int</a>
<span id="RGBA.Rect"></span>    <span class="comment">// Rect is the image&#39;s bounds.</span>
    Rect <a href="#Rectangle">Rectangle</a>
}
</pre>









### <a id="NewRGBA">func</a> [NewRGBA](https://golang.org/src/image/image.go?s=4827:4858#L145)
<pre>func NewRGBA(r <a href="#Rectangle">Rectangle</a>) *<a href="#RGBA">RGBA</a></pre>
NewRGBA returns a new RGBA image with the given bounds.






### <a id="RGBA.At">func</a> (\*RGBA) [At](https://golang.org/src/image/image.go?s=2718:2757#L64)
<pre>func (p *<a href="#RGBA">RGBA</a>) At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a></pre>



### <a id="RGBA.Bounds">func</a> (\*RGBA) [Bounds](https://golang.org/src/image/image.go?s=2665:2698#L62)
<pre>func (p *<a href="#RGBA">RGBA</a>) Bounds() <a href="#Rectangle">Rectangle</a></pre>



### <a id="RGBA.ColorModel">func</a> (\*RGBA) [ColorModel](https://golang.org/src/image/image.go?s=2597:2636#L60)
<pre>func (p *<a href="#RGBA">RGBA</a>) ColorModel() <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Model">Model</a></pre>



### <a id="RGBA.Opaque">func</a> (\*RGBA) [Opaque](https://golang.org/src/image/image.go?s=4487:4515#L127)
<pre>func (p *<a href="#RGBA">RGBA</a>) Opaque() <a href="/pkg/builtin/#bool">bool</a></pre>
Opaque scans the entire image and reports whether it is fully opaque.




### <a id="RGBA.PixOffset">func</a> (\*RGBA) [PixOffset](https://golang.org/src/image/image.go?s=3158:3196#L79)
<pre>func (p *<a href="#RGBA">RGBA</a>) PixOffset(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
PixOffset returns the index of the first element of Pix that corresponds to
the pixel at (x, y).




### <a id="RGBA.RGBAAt">func</a> (\*RGBA) [RGBAAt](https://golang.org/src/image/image.go?s=2786:2828#L68)
<pre>func (p *<a href="#RGBA">RGBA</a>) RGBAAt(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#RGBA">RGBA</a></pre>



### <a id="RGBA.Set">func</a> (\*RGBA) [Set](https://golang.org/src/image/image.go?s=3257:3300#L83)
<pre>func (p *<a href="#RGBA">RGBA</a>) Set(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a>)</pre>



### <a id="RGBA.SetRGBA">func</a> (\*RGBA) [SetRGBA](https://golang.org/src/image/image.go?s=3570:3616#L96)
<pre>func (p *<a href="#RGBA">RGBA</a>) SetRGBA(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#RGBA">RGBA</a>)</pre>



### <a id="RGBA.SubImage">func</a> (\*RGBA) [SubImage](https://golang.org/src/image/image.go?s=3984:4026#L110)
<pre>func (p *<a href="#RGBA">RGBA</a>) SubImage(r <a href="#Rectangle">Rectangle</a>) <a href="#Image">Image</a></pre>
SubImage returns an image representing the portion of the image p visible
through r. The returned value shares pixels with the original image.




## <a id="RGBA64">type</a> [RGBA64](https://golang.org/src/image/image.go?s=5023:5352#L152)
RGBA64 is an in-memory image whose At method returns color.RGBA64 values.


<pre>type RGBA64 struct {
<span id="RGBA64.Pix"></span>    <span class="comment">// Pix holds the image&#39;s pixels, in R, G, B, A order and big-endian format. The pixel at</span>
    <span class="comment">// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*8].</span>
    Pix []<a href="/pkg/builtin/#uint8">uint8</a>
<span id="RGBA64.Stride"></span>    <span class="comment">// Stride is the Pix stride (in bytes) between vertically adjacent pixels.</span>
    Stride <a href="/pkg/builtin/#int">int</a>
<span id="RGBA64.Rect"></span>    <span class="comment">// Rect is the image&#39;s bounds.</span>
    Rect <a href="#Rectangle">Rectangle</a>
}
</pre>









### <a id="NewRGBA64">func</a> [NewRGBA64](https://golang.org/src/image/image.go?s=8022:8057#L260)
<pre>func NewRGBA64(r <a href="#Rectangle">Rectangle</a>) *<a href="#RGBA64">RGBA64</a></pre>
NewRGBA64 returns a new RGBA64 image with the given bounds.






### <a id="RGBA64.At">func</a> (\*RGBA64) [At](https://golang.org/src/image/image.go?s=5481:5522#L166)
<pre>func (p *<a href="#RGBA64">RGBA64</a>) At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a></pre>



### <a id="RGBA64.Bounds">func</a> (\*RGBA64) [Bounds](https://golang.org/src/image/image.go?s=5426:5461#L164)
<pre>func (p *<a href="#RGBA64">RGBA64</a>) Bounds() <a href="#Rectangle">Rectangle</a></pre>



### <a id="RGBA64.ColorModel">func</a> (\*RGBA64) [ColorModel](https://golang.org/src/image/image.go?s=5354:5395#L162)
<pre>func (p *<a href="#RGBA64">RGBA64</a>) ColorModel() <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Model">Model</a></pre>



### <a id="RGBA64.Opaque">func</a> (\*RGBA64) [Opaque](https://golang.org/src/image/image.go?s=7652:7682#L242)
<pre>func (p *<a href="#RGBA64">RGBA64</a>) Opaque() <a href="/pkg/builtin/#bool">bool</a></pre>
Opaque scans the entire image and reports whether it is fully opaque.




### <a id="RGBA64.PixOffset">func</a> (\*RGBA64) [PixOffset](https://golang.org/src/image/image.go?s=6051:6091#L186)
<pre>func (p *<a href="#RGBA64">RGBA64</a>) PixOffset(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
PixOffset returns the index of the first element of Pix that corresponds to
the pixel at (x, y).




### <a id="RGBA64.RGBA64At">func</a> (\*RGBA64) [RGBA64At](https://golang.org/src/image/image.go?s=5553:5601#L170)
<pre>func (p *<a href="#RGBA64">RGBA64</a>) RGBA64At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#RGBA64">RGBA64</a></pre>



### <a id="RGBA64.Set">func</a> (\*RGBA64) [Set](https://golang.org/src/image/image.go?s=6152:6197#L190)
<pre>func (p *<a href="#RGBA64">RGBA64</a>) Set(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a>)</pre>



### <a id="RGBA64.SetRGBA64">func</a> (\*RGBA64) [SetRGBA64](https://golang.org/src/image/image.go?s=6599:6651#L207)
<pre>func (p *<a href="#RGBA64">RGBA64</a>) SetRGBA64(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#RGBA64">RGBA64</a>)</pre>



### <a id="RGBA64.SubImage">func</a> (\*RGBA64) [SubImage](https://golang.org/src/image/image.go?s=7143:7187#L225)
<pre>func (p *<a href="#RGBA64">RGBA64</a>) SubImage(r <a href="#Rectangle">Rectangle</a>) <a href="#Image">Image</a></pre>
SubImage returns an image representing the portion of the image p visible
through r. The returned value shares pixels with the original image.




## <a id="Rectangle">type</a> [Rectangle](https://golang.org/src/image/geom.go?s=2049:2090#L77)
A Rectangle contains the points with Min.X <= X < Max.X, Min.Y <= Y < Max.Y.
It is well-formed if Min.X <= Max.X and likewise for Y. Points are always
well-formed. A rectangle's methods always return well-formed outputs for
well-formed inputs.

A Rectangle is also an Image whose bounds are the rectangle itself. At
returns color.Opaque for points in the rectangle and color.Transparent
otherwise.


<pre>type Rectangle struct {
<span id="Rectangle.Min"></span>    Min, Max <a href="#Point">Point</a>
}
</pre>




ZR is the zero Rectangle.

Deprecated: Use a literal image.Rectangle{} instead.


<pre>var <span id="ZR">ZR</span> <a href="#Rectangle">Rectangle</a></pre>





### <a id="Rect">func</a> [Rect](https://golang.org/src/image/geom.go?s=6341:6380#L256)
<pre>func Rect(x0, y0, x1, y1 <a href="/pkg/builtin/#int">int</a>) <a href="#Rectangle">Rectangle</a></pre>
Rect is shorthand for Rectangle{Pt(x0, y0), Pt(x1, y1)}. The returned
rectangle has minimum and maximum coordinates swapped if necessary so that
it is well-formed.






### <a id="Rectangle.Add">func</a> (Rectangle) [Add](https://golang.org/src/image/geom.go?s=2597:2638#L105)
<pre>func (r <a href="#Rectangle">Rectangle</a>) Add(p <a href="#Point">Point</a>) <a href="#Rectangle">Rectangle</a></pre>
Add returns the rectangle r translated by p.




### <a id="Rectangle.At">func</a> (Rectangle) [At](https://golang.org/src/image/geom.go?s=5719:5762#L231)
<pre>func (r <a href="#Rectangle">Rectangle</a>) At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a></pre>
At implements the Image interface.




### <a id="Rectangle.Bounds">func</a> (Rectangle) [Bounds](https://golang.org/src/image/geom.go?s=5887:5924#L239)
<pre>func (r <a href="#Rectangle">Rectangle</a>) Bounds() <a href="#Rectangle">Rectangle</a></pre>
Bounds implements the Image interface.




### <a id="Rectangle.Canon">func</a> (Rectangle) [Canon](https://golang.org/src/image/geom.go?s=5499:5535#L220)
<pre>func (r <a href="#Rectangle">Rectangle</a>) Canon() <a href="#Rectangle">Rectangle</a></pre>
Canon returns the canonical version of r. The returned rectangle has minimum
and maximum coordinates swapped if necessary so that it is well-formed.




### <a id="Rectangle.ColorModel">func</a> (Rectangle) [ColorModel](https://golang.org/src/image/geom.go?s=5986:6029#L244)
<pre>func (r <a href="#Rectangle">Rectangle</a>) ColorModel() <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Model">Model</a></pre>
ColorModel implements the Image interface.




### <a id="Rectangle.Dx">func</a> (Rectangle) [Dx](https://golang.org/src/image/geom.go?s=2270:2297#L87)
<pre>func (r <a href="#Rectangle">Rectangle</a>) Dx() <a href="/pkg/builtin/#int">int</a></pre>
Dx returns r's width.




### <a id="Rectangle.Dy">func</a> (Rectangle) [Dy](https://golang.org/src/image/geom.go?s=2355:2382#L92)
<pre>func (r <a href="#Rectangle">Rectangle</a>) Dy() <a href="/pkg/builtin/#int">int</a></pre>
Dy returns r's height.




### <a id="Rectangle.Empty">func</a> (Rectangle) [Empty](https://golang.org/src/image/geom.go?s=4500:4531#L190)
<pre>func (r <a href="#Rectangle">Rectangle</a>) Empty() <a href="/pkg/builtin/#bool">bool</a></pre>
Empty reports whether the rectangle contains no points.




### <a id="Rectangle.Eq">func</a> (Rectangle) [Eq](https://golang.org/src/image/geom.go?s=4694:4733#L196)
<pre>func (r <a href="#Rectangle">Rectangle</a>) Eq(s <a href="#Rectangle">Rectangle</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Eq reports whether r and s contain the same set of points. All empty
rectangles are considered equal.




### <a id="Rectangle.In">func</a> (Rectangle) [In](https://golang.org/src/image/geom.go?s=5067:5106#L208)
<pre>func (r <a href="#Rectangle">Rectangle</a>) In(s <a href="#Rectangle">Rectangle</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
In reports whether every point in r is in s.




### <a id="Rectangle.Inset">func</a> (Rectangle) [Inset](https://golang.org/src/image/geom.go?s=3122:3163#L123)
<pre>func (r <a href="#Rectangle">Rectangle</a>) Inset(n <a href="/pkg/builtin/#int">int</a>) <a href="#Rectangle">Rectangle</a></pre>
Inset returns the rectangle r inset by n, which may be negative. If either
of r's dimensions is less than 2*n then an empty rectangle near the center
of r will be returned.




### <a id="Rectangle.Intersect">func</a> (Rectangle) [Intersect](https://golang.org/src/image/geom.go?s=3567:3618#L143)
<pre>func (r <a href="#Rectangle">Rectangle</a>) Intersect(s <a href="#Rectangle">Rectangle</a>) <a href="#Rectangle">Rectangle</a></pre>
Intersect returns the largest rectangle contained by both r and s. If the
two rectangles do not overlap then the zero rectangle will be returned.




### <a id="Rectangle.Overlaps">func</a> (Rectangle) [Overlaps](https://golang.org/src/image/geom.go?s=4847:4892#L201)
<pre>func (r <a href="#Rectangle">Rectangle</a>) Overlaps(s <a href="#Rectangle">Rectangle</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Overlaps reports whether r and s have a non-empty intersection.




### <a id="Rectangle.Size">func</a> (Rectangle) [Size](https://golang.org/src/image/geom.go?s=2452:2483#L97)
<pre>func (r <a href="#Rectangle">Rectangle</a>) Size() <a href="#Point">Point</a></pre>
Size returns r's width and height.




### <a id="Rectangle.String">func</a> (Rectangle) [String](https://golang.org/src/image/geom.go?s=2159:2193#L82)
<pre>func (r <a href="#Rectangle">Rectangle</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns a string representation of r like "(3,4)-(6,5)".




### <a id="Rectangle.Sub">func</a> (Rectangle) [Sub](https://golang.org/src/image/geom.go?s=2793:2834#L113)
<pre>func (r <a href="#Rectangle">Rectangle</a>) Sub(p <a href="#Point">Point</a>) <a href="#Rectangle">Rectangle</a></pre>
Sub returns the rectangle r translated by -p.




### <a id="Rectangle.Union">func</a> (Rectangle) [Union](https://golang.org/src/image/geom.go?s=4130:4177#L167)
<pre>func (r <a href="#Rectangle">Rectangle</a>) Union(s <a href="#Rectangle">Rectangle</a>) <a href="#Rectangle">Rectangle</a></pre>
Union returns the smallest rectangle that contains both r and s.




## <a id="Uniform">type</a> [Uniform](https://golang.org/src/image/names.go?s=668:706#L14)
Uniform is an infinite-sized Image of uniform color.
It implements the color.Color, color.Model, and Image interfaces.


<pre>type Uniform struct {
<span id="Uniform.C"></span>    C <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a>
}
</pre>









### <a id="NewUniform">func</a> [NewUniform](https://golang.org/src/image/names.go?s=1213:1252#L40)
<pre>func NewUniform(c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a>) *<a href="#Uniform">Uniform</a></pre>





### <a id="Uniform.At">func</a> (\*Uniform) [At](https://golang.org/src/image/names.go?s=998:1040#L32)
<pre>func (c *<a href="#Uniform">Uniform</a>) At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a></pre>



### <a id="Uniform.Bounds">func</a> (\*Uniform) [Bounds](https://golang.org/src/image/names.go?s=903:939#L30)
<pre>func (c *<a href="#Uniform">Uniform</a>) Bounds() <a href="#Rectangle">Rectangle</a></pre>



### <a id="Uniform.ColorModel">func</a> (\*Uniform) [ColorModel](https://golang.org/src/image/names.go?s=777:819#L22)
<pre>func (c *<a href="#Uniform">Uniform</a>) ColorModel() <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Model">Model</a></pre>



### <a id="Uniform.Convert">func</a> (\*Uniform) [Convert](https://golang.org/src/image/names.go?s=835:885#L26)
<pre>func (c *<a href="#Uniform">Uniform</a>) Convert(<a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a></pre>



### <a id="Uniform.Opaque">func</a> (\*Uniform) [Opaque](https://golang.org/src/image/names.go?s=1130:1161#L35)
<pre>func (c *<a href="#Uniform">Uniform</a>) Opaque() <a href="/pkg/builtin/#bool">bool</a></pre>
Opaque scans the entire image and reports whether it is fully opaque.




### <a id="Uniform.RGBA">func</a> (\*Uniform) [RGBA](https://golang.org/src/image/names.go?s=708:752#L18)
<pre>func (c *<a href="#Uniform">Uniform</a>) RGBA() (r, g, b, a <a href="/pkg/builtin/#uint32">uint32</a>)</pre>



## <a id="YCbCr">type</a> [YCbCr](https://golang.org/src/image/ycbcr.go?s=1830:1977#L44)
YCbCr is an in-memory image of Y'CbCr colors. There is one Y sample per
pixel, but each Cb and Cr sample can span one or more pixels.
YStride is the Y slice index delta between vertically adjacent pixels.
CStride is the Cb and Cr slice index delta between vertically adjacent pixels
that map to separate chroma samples.
It is not an absolute requirement, but YStride and len(Y) are typically
multiples of 8, and:


	For 4:4:4, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/1.
	For 4:2:2, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/2.
	For 4:2:0, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/4.
	For 4:4:0, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/2.
	For 4:1:1, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/4.
	For 4:1:0, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/8.


<pre>type YCbCr struct {
<span id="YCbCr.Y"></span>    Y, Cb, Cr      []<a href="/pkg/builtin/#uint8">uint8</a>
<span id="YCbCr.YStride"></span>    YStride        <a href="/pkg/builtin/#int">int</a>
<span id="YCbCr.CStride"></span>    CStride        <a href="/pkg/builtin/#int">int</a>
<span id="YCbCr.SubsampleRatio"></span>    SubsampleRatio <a href="#YCbCrSubsampleRatio">YCbCrSubsampleRatio</a>
<span id="YCbCr.Rect"></span>    Rect           <a href="#Rectangle">Rectangle</a>
}
</pre>









### <a id="NewYCbCr">func</a> [NewYCbCr](https://golang.org/src/image/ycbcr.go?s=4872:4941#L159)
<pre>func NewYCbCr(r <a href="#Rectangle">Rectangle</a>, subsampleRatio <a href="#YCbCrSubsampleRatio">YCbCrSubsampleRatio</a>) *<a href="#YCbCr">YCbCr</a></pre>
NewYCbCr returns a new YCbCr image with the given bounds and subsample
ratio.






### <a id="YCbCr.At">func</a> (\*YCbCr) [At](https://golang.org/src/image/ycbcr.go?s=2105:2145#L60)
<pre>func (p *<a href="#YCbCr">YCbCr</a>) At(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a></pre>



### <a id="YCbCr.Bounds">func</a> (\*YCbCr) [Bounds](https://golang.org/src/image/ycbcr.go?s=2050:2084#L56)
<pre>func (p *<a href="#YCbCr">YCbCr</a>) Bounds() <a href="#Rectangle">Rectangle</a></pre>



### <a id="YCbCr.COffset">func</a> (\*YCbCr) [COffset](https://golang.org/src/image/ycbcr.go?s=2693:2730#L85)
<pre>func (p *<a href="#YCbCr">YCbCr</a>) COffset(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
COffset returns the index of the first element of Cb or Cr that corresponds
to the pixel at (x, y).




### <a id="YCbCr.ColorModel">func</a> (\*YCbCr) [ColorModel](https://golang.org/src/image/ycbcr.go?s=1979:2019#L52)
<pre>func (p *<a href="#YCbCr">YCbCr</a>) ColorModel() <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Model">Model</a></pre>



### <a id="YCbCr.Opaque">func</a> (\*YCbCr) [Opaque](https://golang.org/src/image/ycbcr.go?s=4122:4151#L127)
<pre>func (p *<a href="#YCbCr">YCbCr</a>) Opaque() <a href="/pkg/builtin/#bool">bool</a></pre>



### <a id="YCbCr.SubImage">func</a> (\*YCbCr) [SubImage](https://golang.org/src/image/ycbcr.go?s=3468:3511#L104)
<pre>func (p *<a href="#YCbCr">YCbCr</a>) SubImage(r <a href="#Rectangle">Rectangle</a>) <a href="#Image">Image</a></pre>
SubImage returns an image representing the portion of the image p visible
through r. The returned value shares pixels with the original image.




### <a id="YCbCr.YCbCrAt">func</a> (\*YCbCr) [YCbCrAt](https://golang.org/src/image/ycbcr.go?s=2175:2220#L64)
<pre>func (p *<a href="#YCbCr">YCbCr</a>) YCbCrAt(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#YCbCr">YCbCr</a></pre>



### <a id="YCbCr.YOffset">func</a> (\*YCbCr) [YOffset](https://golang.org/src/image/ycbcr.go?s=2488:2525#L79)
<pre>func (p *<a href="#YCbCr">YCbCr</a>) YOffset(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
YOffset returns the index of the first element of Y that corresponds to
the pixel at (x, y).




## <a id="YCbCrSubsampleRatio">type</a> [YCbCrSubsampleRatio](https://golang.org/src/image/ycbcr.go?s=278:306#L2)
YCbCrSubsampleRatio is the chroma subsample ratio used in a YCbCr image.


<pre>type YCbCrSubsampleRatio <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="YCbCrSubsampleRatio444">YCbCrSubsampleRatio444</span> <a href="#YCbCrSubsampleRatio">YCbCrSubsampleRatio</a> = <a href="/pkg/builtin/#iota">iota</a>
    <span id="YCbCrSubsampleRatio422">YCbCrSubsampleRatio422</span>
    <span id="YCbCrSubsampleRatio420">YCbCrSubsampleRatio420</span>
    <span id="YCbCrSubsampleRatio440">YCbCrSubsampleRatio440</span>
    <span id="YCbCrSubsampleRatio411">YCbCrSubsampleRatio411</span>
    <span id="YCbCrSubsampleRatio410">YCbCrSubsampleRatio410</span>
)</pre>









### <a id="YCbCrSubsampleRatio.String">func</a> (YCbCrSubsampleRatio) [String](https://golang.org/src/image/ycbcr.go?s=490:534#L13)
<pre>func (s <a href="#YCbCrSubsampleRatio">YCbCrSubsampleRatio</a>) String() <a href="/pkg/builtin/#string">string</a></pre>







