

# color
`import "image/color"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package color implements a basic color library.




## <a id="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func CMYKToRGB(c, m, y, k uint8) (uint8, uint8, uint8)](#CMYKToRGB)
* [func RGBToCMYK(r, g, b uint8) (uint8, uint8, uint8, uint8)](#RGBToCMYK)
* [func RGBToYCbCr(r, g, b uint8) (uint8, uint8, uint8)](#RGBToYCbCr)
* [func YCbCrToRGB(y, cb, cr uint8) (uint8, uint8, uint8)](#YCbCrToRGB)
* [type Alpha](#Alpha)
  * [func (c Alpha) RGBA() (r, g, b, a uint32)](#Alpha.RGBA)
* [type Alpha16](#Alpha16)
  * [func (c Alpha16) RGBA() (r, g, b, a uint32)](#Alpha16.RGBA)
* [type CMYK](#CMYK)
  * [func (c CMYK) RGBA() (uint32, uint32, uint32, uint32)](#CMYK.RGBA)
* [type Color](#Color)
* [type Gray](#Gray)
  * [func (c Gray) RGBA() (r, g, b, a uint32)](#Gray.RGBA)
* [type Gray16](#Gray16)
  * [func (c Gray16) RGBA() (r, g, b, a uint32)](#Gray16.RGBA)
* [type Model](#Model)
  * [func ModelFunc(f func(Color) Color) Model](#ModelFunc)
* [type NRGBA](#NRGBA)
  * [func (c NRGBA) RGBA() (r, g, b, a uint32)](#NRGBA.RGBA)
* [type NRGBA64](#NRGBA64)
  * [func (c NRGBA64) RGBA() (r, g, b, a uint32)](#NRGBA64.RGBA)
* [type NYCbCrA](#NYCbCrA)
  * [func (c NYCbCrA) RGBA() (uint32, uint32, uint32, uint32)](#NYCbCrA.RGBA)
* [type Palette](#Palette)
  * [func (p Palette) Convert(c Color) Color](#Palette.Convert)
  * [func (p Palette) Index(c Color) int](#Palette.Index)
* [type RGBA](#RGBA)
  * [func (c RGBA) RGBA() (r, g, b, a uint32)](#RGBA.RGBA)
* [type RGBA64](#RGBA64)
  * [func (c RGBA64) RGBA() (r, g, b, a uint32)](#RGBA64.RGBA)
* [type YCbCr](#YCbCr)
  * [func (c YCbCr) RGBA() (uint32, uint32, uint32, uint32)](#YCbCr.RGBA)




#### <a id="pkg-files">Package files</a>
[color.go](https://golang.org/src/image/color/color.go) [ycbcr.go](https://golang.org/src/image/color/ycbcr.go) 




## <a id="pkg-variables">Variables</a>
Standard colors.


<pre>var (
    <span id="Black">Black</span>       = <a href="#Gray16">Gray16</a>{0}
    <span id="White">White</span>       = <a href="#Gray16">Gray16</a>{0xffff}
    <span id="Transparent">Transparent</span> = <a href="#Alpha16">Alpha16</a>{0}
    <span id="Opaque">Opaque</span>      = <a href="#Alpha16">Alpha16</a>{0xffff}
)</pre>

## <a id="CMYKToRGB">func</a> [CMYKToRGB](https://golang.org/src/image/color/ycbcr.go?s=9901:9955#L326)
<pre>func CMYKToRGB(c, m, y, k <a href="/pkg/builtin/#uint8">uint8</a>) (<a href="/pkg/builtin/#uint8">uint8</a>, <a href="/pkg/builtin/#uint8">uint8</a>, <a href="/pkg/builtin/#uint8">uint8</a>)</pre>
CMYKToRGB converts a CMYK quadruple to an RGB triple.



## <a id="RGBToCMYK">func</a> [RGBToCMYK](https://golang.org/src/image/color/ycbcr.go?s=9499:9557#L305)
<pre>func RGBToCMYK(r, g, b <a href="/pkg/builtin/#uint8">uint8</a>) (<a href="/pkg/builtin/#uint8">uint8</a>, <a href="/pkg/builtin/#uint8">uint8</a>, <a href="/pkg/builtin/#uint8">uint8</a>, <a href="/pkg/builtin/#uint8">uint8</a>)</pre>
RGBToCMYK converts an RGB triple to a CMYK quadruple.



## <a id="RGBToYCbCr">func</a> [RGBToYCbCr](https://golang.org/src/image/color/ycbcr.go?s=232:284#L1)
<pre>func RGBToYCbCr(r, g, b <a href="/pkg/builtin/#uint8">uint8</a>) (<a href="/pkg/builtin/#uint8">uint8</a>, <a href="/pkg/builtin/#uint8">uint8</a>, <a href="/pkg/builtin/#uint8">uint8</a>)</pre>
RGBToYCbCr converts an RGB triple to a Y'CbCr triple.



## <a id="YCbCrToRGB">func</a> [YCbCrToRGB](https://golang.org/src/image/color/ycbcr.go?s=1556:1610#L48)
<pre>func YCbCrToRGB(y, cb, cr <a href="/pkg/builtin/#uint8">uint8</a>) (<a href="/pkg/builtin/#uint8">uint8</a>, <a href="/pkg/builtin/#uint8">uint8</a>, <a href="/pkg/builtin/#uint8">uint8</a>)</pre>
YCbCrToRGB converts a Y'CbCr triple to an RGB triple.





## <a id="Alpha">type</a> [Alpha](https://golang.org/src/image/color/color.go?s=2362:2392#L89)
Alpha represents an 8-bit alpha color.


<pre>type Alpha struct {
<span id="Alpha.A"></span>    A <a href="/pkg/builtin/#uint8">uint8</a>
}
</pre>











### <a id="Alpha.RGBA">func</a> (Alpha) [RGBA](https://golang.org/src/image/color/color.go?s=2394:2435#L93)
<pre>func (c <a href="#Alpha">Alpha</a>) RGBA() (r, g, b, a <a href="/pkg/builtin/#uint32">uint32</a>)</pre>



## <a id="Alpha16">type</a> [Alpha16](https://golang.org/src/image/color/color.go?s=2534:2567#L100)
Alpha16 represents a 16-bit alpha color.


<pre>type Alpha16 struct {
<span id="Alpha16.A"></span>    A <a href="/pkg/builtin/#uint16">uint16</a>
}
</pre>











### <a id="Alpha16.RGBA">func</a> (Alpha16) [RGBA](https://golang.org/src/image/color/color.go?s=2569:2612#L104)
<pre>func (c <a href="#Alpha16">Alpha16</a>) RGBA() (r, g, b, a <a href="/pkg/builtin/#uint32">uint32</a>)</pre>



## <a id="CMYK">type</a> [CMYK](https://golang.org/src/image/color/ycbcr.go?s=10352:10390#L338)
CMYK represents a fully opaque CMYK color, having 8 bits for each of cyan,
magenta, yellow and black.

It is not associated with any particular color profile.


<pre>type CMYK struct {
<span id="CMYK.C"></span>    C, M, Y, K <a href="/pkg/builtin/#uint8">uint8</a>
}
</pre>











### <a id="CMYK.RGBA">func</a> (CMYK) [RGBA](https://golang.org/src/image/color/ycbcr.go?s=10392:10445#L342)
<pre>func (c <a href="#CMYK">CMYK</a>) RGBA() (<a href="/pkg/builtin/#uint32">uint32</a>, <a href="/pkg/builtin/#uint32">uint32</a>, <a href="/pkg/builtin/#uint32">uint32</a>, <a href="/pkg/builtin/#uint32">uint32</a>)</pre>



## <a id="Color">type</a> [Color](https://golang.org/src/image/color/color.go?s=335:744#L1)
Color can convert itself to alpha-premultiplied 16-bits per channel RGBA.
The conversion may be lossy.


<pre>type Color interface {
    <span class="comment">// RGBA returns the alpha-premultiplied red, green, blue and alpha values</span>
    <span class="comment">// for the color. Each value ranges within [0, 0xffff], but is represented</span>
    <span class="comment">// by a uint32 so that multiplying by a blend factor up to 0xffff will not</span>
    <span class="comment">// overflow.</span>
    <span class="comment">//</span>
    <span class="comment">// An alpha-premultiplied color component c has been scaled by alpha (a),</span>
    <span class="comment">// so has valid values 0 &lt;= c &lt;= a.</span>
    RGBA() (r, g, b, a <a href="/pkg/builtin/#uint32">uint32</a>)
}</pre>











## <a id="Gray">type</a> [Gray](https://golang.org/src/image/color/color.go?s=2699:2728#L110)
Gray represents an 8-bit grayscale color.


<pre>type Gray struct {
<span id="Gray.Y"></span>    Y <a href="/pkg/builtin/#uint8">uint8</a>
}
</pre>











### <a id="Gray.RGBA">func</a> (Gray) [RGBA](https://golang.org/src/image/color/color.go?s=2730:2770#L114)
<pre>func (c <a href="#Gray">Gray</a>) RGBA() (r, g, b, a <a href="/pkg/builtin/#uint32">uint32</a>)</pre>



## <a id="Gray16">type</a> [Gray16](https://golang.org/src/image/color/color.go?s=2878:2910#L121)
Gray16 represents a 16-bit grayscale color.


<pre>type Gray16 struct {
<span id="Gray16.Y"></span>    Y <a href="/pkg/builtin/#uint16">uint16</a>
}
</pre>











### <a id="Gray16.RGBA">func</a> (Gray16) [RGBA](https://golang.org/src/image/color/color.go?s=2912:2954#L125)
<pre>func (c <a href="#Gray16">Gray16</a>) RGBA() (r, g, b, a <a href="/pkg/builtin/#uint32">uint32</a>)</pre>



## <a id="Model">type</a> [Model](https://golang.org/src/image/color/color.go?s=3098:3146#L132)
Model can convert any Color to one from its own color model. The conversion
may be lossy.


<pre>type Model interface {
    Convert(c <a href="#Color">Color</a>) <a href="#Color">Color</a>
}</pre>




Models for the standard color types.


<pre>var (
    <span id="RGBAModel">RGBAModel</span>    <a href="#Model">Model</a> = <a href="#ModelFunc">ModelFunc</a>(rgbaModel)
    <span id="RGBA64Model">RGBA64Model</span>  <a href="#Model">Model</a> = <a href="#ModelFunc">ModelFunc</a>(rgba64Model)
    <span id="NRGBAModel">NRGBAModel</span>   <a href="#Model">Model</a> = <a href="#ModelFunc">ModelFunc</a>(nrgbaModel)
    <span id="NRGBA64Model">NRGBA64Model</span> <a href="#Model">Model</a> = <a href="#ModelFunc">ModelFunc</a>(nrgba64Model)
    <span id="AlphaModel">AlphaModel</span>   <a href="#Model">Model</a> = <a href="#ModelFunc">ModelFunc</a>(alphaModel)
    <span id="Alpha16Model">Alpha16Model</span> <a href="#Model">Model</a> = <a href="#ModelFunc">ModelFunc</a>(alpha16Model)
    <span id="GrayModel">GrayModel</span>    <a href="#Model">Model</a> = <a href="#ModelFunc">ModelFunc</a>(grayModel)
    <span id="Gray16Model">Gray16Model</span>  <a href="#Model">Model</a> = <a href="#ModelFunc">ModelFunc</a>(gray16Model)
)</pre>
CMYKModel is the Model for CMYK colors.


<pre>var <span id="CMYKModel">CMYKModel</span> <a href="#Model">Model</a> = <a href="#ModelFunc">ModelFunc</a>(cmykModel)</pre>
NYCbCrAModel is the Model for non-alpha-premultiplied Y'CbCr-with-alpha
colors.


<pre>var <span id="NYCbCrAModel">NYCbCrAModel</span> <a href="#Model">Model</a> = <a href="#ModelFunc">ModelFunc</a>(nYCbCrAModel)</pre>
YCbCrModel is the Model for Y'CbCr colors.


<pre>var <span id="YCbCrModel">YCbCrModel</span> <a href="#Model">Model</a> = <a href="#ModelFunc">ModelFunc</a>(yCbCrModel)</pre>





### <a id="ModelFunc">func</a> [ModelFunc](https://golang.org/src/image/color/color.go?s=3221:3262#L137)
<pre>func ModelFunc(f func(<a href="#Color">Color</a>) <a href="#Color">Color</a>) <a href="#Model">Model</a></pre>
ModelFunc returns a Model that invokes f to implement the conversion.






## <a id="NRGBA">type</a> [NRGBA](https://golang.org/src/image/color/color.go?s=1635:1674#L46)
NRGBA represents a non-alpha-premultiplied 32-bit color.


<pre>type NRGBA struct {
<span id="NRGBA.R"></span>    R, G, B, A <a href="/pkg/builtin/#uint8">uint8</a>
}
</pre>











### <a id="NRGBA.RGBA">func</a> (NRGBA) [RGBA](https://golang.org/src/image/color/color.go?s=1676:1717#L50)
<pre>func (c <a href="#NRGBA">NRGBA</a>) RGBA() (r, g, b, a <a href="/pkg/builtin/#uint32">uint32</a>)</pre>



## <a id="NRGBA64">type</a> [NRGBA64](https://golang.org/src/image/color/color.go?s=2058:2100#L70)
NRGBA64 represents a non-alpha-premultiplied 64-bit color,
having 16 bits for each of red, green, blue and alpha.


<pre>type NRGBA64 struct {
<span id="NRGBA64.R"></span>    R, G, B, A <a href="/pkg/builtin/#uint16">uint16</a>
}
</pre>











### <a id="NRGBA64.RGBA">func</a> (NRGBA64) [RGBA](https://golang.org/src/image/color/color.go?s=2102:2145#L74)
<pre>func (c <a href="#NRGBA64">NRGBA64</a>) RGBA() (r, g, b, a <a href="/pkg/builtin/#uint32">uint32</a>)</pre>



## <a id="NYCbCrA">type</a> [NYCbCrA](https://golang.org/src/image/color/ycbcr.go?s=7878:7917#L232)
NYCbCrA represents a non-alpha-premultiplied Y'CbCr-with-alpha color, having
8 bits each for one luma, two chroma and one alpha component.


<pre>type NYCbCrA struct {
    <a href="#YCbCr">YCbCr</a>
<span id="NYCbCrA.A"></span>    A <a href="/pkg/builtin/#uint8">uint8</a>
}
</pre>











### <a id="NYCbCrA.RGBA">func</a> (NYCbCrA) [RGBA](https://golang.org/src/image/color/ycbcr.go?s=7919:7975#L237)
<pre>func (c <a href="#NYCbCrA">NYCbCrA</a>) RGBA() (<a href="/pkg/builtin/#uint32">uint32</a>, <a href="/pkg/builtin/#uint32">uint32</a>, <a href="/pkg/builtin/#uint32">uint32</a>, <a href="/pkg/builtin/#uint32">uint32</a>)</pre>



## <a id="Palette">type</a> [Palette](https://golang.org/src/image/color/color.go?s=6539:6559#L270)
Palette is a palette of colors.


<pre>type Palette []<a href="#Color">Color</a></pre>











### <a id="Palette.Convert">func</a> (Palette) [Convert](https://golang.org/src/image/color/color.go?s=6637:6676#L273)
<pre>func (p <a href="#Palette">Palette</a>) Convert(c <a href="#Color">Color</a>) <a href="#Color">Color</a></pre>
Convert returns the palette color closest to c in Euclidean R,G,B space.




### <a id="Palette.Index">func</a> (Palette) [Index](https://golang.org/src/image/color/color.go?s=6830:6865#L282)
<pre>func (p <a href="#Palette">Palette</a>) Index(c <a href="#Color">Color</a>) <a href="/pkg/builtin/#int">int</a></pre>
Index returns the index of the palette color closest to c in Euclidean
R,G,B,A space.




## <a id="RGBA">type</a> [RGBA](https://golang.org/src/image/color/color.go?s=983:1021#L16)
RGBA represents a traditional 32-bit alpha-premultiplied color, having 8
bits for each of red, green, blue and alpha.

An alpha-premultiplied color component C has been scaled by alpha (A), so
has valid values 0 <= C <= A.


<pre>type RGBA struct {
<span id="RGBA.R"></span>    R, G, B, A <a href="/pkg/builtin/#uint8">uint8</a>
}
</pre>











### <a id="RGBA.RGBA">func</a> (RGBA) [RGBA](https://golang.org/src/image/color/color.go?s=1023:1063#L20)
<pre>func (c <a href="#RGBA">RGBA</a>) RGBA() (r, g, b, a <a href="/pkg/builtin/#uint32">uint32</a>)</pre>



## <a id="RGBA64">type</a> [RGBA64](https://golang.org/src/image/color/color.go?s=1425:1466#L37)
RGBA64 represents a 64-bit alpha-premultiplied color, having 16 bits for
each of red, green, blue and alpha.

An alpha-premultiplied color component C has been scaled by alpha (A), so
has valid values 0 <= C <= A.


<pre>type RGBA64 struct {
<span id="RGBA64.R"></span>    R, G, B, A <a href="/pkg/builtin/#uint16">uint16</a>
}
</pre>











### <a id="RGBA64.RGBA">func</a> (RGBA64) [RGBA](https://golang.org/src/image/color/color.go?s=1468:1510#L41)
<pre>func (c <a href="#RGBA64">RGBA64</a>) RGBA() (r, g, b, a <a href="/pkg/builtin/#uint32">uint32</a>)</pre>



## <a id="YCbCr">type</a> [YCbCr](https://golang.org/src/image/color/ycbcr.go?s=5890:5928#L157)
YCbCr represents a fully opaque 24-bit Y'CbCr color, having 8 bits each for
one luma and two chroma components.

JPEG, VP8, the MPEG family and other codecs use this color model. Such
codecs often use the terms YUV and Y'CbCr interchangeably, but strictly
speaking, the term YUV applies only to analog video signals, and Y' (luma)
is Y (luminance) after applying gamma correction.

Conversion between RGB and Y'CbCr is lossy and there are multiple, slightly
different formulae for converting between the two. This package follows
the JFIF specification at <a href="https://www.w3.org/Graphics/JPEG/jfif3.pdf">https://www.w3.org/Graphics/JPEG/jfif3.pdf</a>.


<pre>type YCbCr struct {
<span id="YCbCr.Y"></span>    Y, Cb, Cr <a href="/pkg/builtin/#uint8">uint8</a>
}
</pre>











### <a id="YCbCr.RGBA">func</a> (YCbCr) [RGBA](https://golang.org/src/image/color/ycbcr.go?s=5930:5984#L161)
<pre>func (c <a href="#YCbCr">YCbCr</a>) RGBA() (<a href="/pkg/builtin/#uint32">uint32</a>, <a href="/pkg/builtin/#uint32">uint32</a>, <a href="/pkg/builtin/#uint32">uint32</a>, <a href="/pkg/builtin/#uint32">uint32</a>)</pre>






