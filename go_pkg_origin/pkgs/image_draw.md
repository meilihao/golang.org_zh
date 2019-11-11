

# draw
`import "image/draw"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package draw provides image composition functions.

See "The Go image/draw package" for an introduction to this package:
<a href="https://golang.org/doc/articles/image_draw.html">https://golang.org/doc/articles/image_draw.html</a>




## <a id="pkg-index">Index</a>
* [func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)](#Draw)
* [func DrawMask(dst Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, op Op)](#DrawMask)
* [type Drawer](#Drawer)
* [type Image](#Image)
* [type Op](#Op)
  * [func (op Op) Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point)](#Op.Draw)
* [type Quantizer](#Quantizer)


#### <a id="pkg-examples">Examples</a>
* [Drawer (FloydSteinberg)](#example_Drawer_floydSteinberg)


#### <a id="pkg-files">Package files</a>
[draw.go](https://golang.org/src/image/draw/draw.go) 






## <a id="Draw">func</a> [Draw](https://golang.org/src/image/draw/draw.go?s=2817:2896#L90)
<pre>func Draw(dst <a href="#Image">Image</a>, r <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Rectangle">Rectangle</a>, src <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Image">Image</a>, sp <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Point">Point</a>, op <a href="#Op">Op</a>)</pre>
Draw calls DrawMask with a nil mask.



## <a id="DrawMask">func</a> [DrawMask](https://golang.org/src/image/draw/draw.go?s=3138:3255#L96)
<pre>func DrawMask(dst <a href="#Image">Image</a>, r <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Rectangle">Rectangle</a>, src <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Image">Image</a>, sp <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Point">Point</a>, mask <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Image">Image</a>, mp <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Point">Point</a>, op <a href="#Op">Op</a>)</pre>
DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
in dst with the result of a Porter-Duff composition. A nil mask is treated as opaque.





## <a id="Drawer">type</a> [Drawer](https://golang.org/src/image/draw/draw.go?s=1342:1564#L40)
Drawer contains the Draw method.


<pre>type Drawer interface {
    <span class="comment">// Draw aligns r.Min in dst with sp in src and then replaces the</span>
    <span class="comment">// rectangle r in dst with the result of drawing src on dst.</span>
    Draw(dst <a href="#Image">Image</a>, r <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Rectangle">Rectangle</a>, src <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Image">Image</a>, sp <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Point">Point</a>)
}</pre>




FloydSteinberg is a Drawer that is the Src Op with Floyd-Steinberg error
diffusion.


<pre>var <span id="FloydSteinberg">FloydSteinberg</span> <a href="#Drawer">Drawer</a> = floydSteinberg{}</pre>


<a id="example_Drawer_floydSteinberg">Example (FloydSteinberg)</a>


```go
```

output:
```txt
```






## <a id="Image">type</a> [Image](https://golang.org/src/image/draw/draw.go?s=572:639#L11)
Image is an image.Image with a Set method to change a single pixel.


<pre>type Image interface {
    <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Image">Image</a>
    Set(x, y <a href="/pkg/builtin/#int">int</a>, c <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Color">Color</a>)
}</pre>











## <a id="Op">type</a> [Op](https://golang.org/src/image/draw/draw.go?s=956:967#L24)
Op is a Porter-Duff compositing operator.


<pre>type Op <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span class="comment">// Over specifies ``(src in mask) over dst&#39;&#39;.</span>
    <span id="Over">Over</span> <a href="#Op">Op</a> = <a href="/pkg/builtin/#iota">iota</a>
    <span class="comment">// Src specifies ``src in mask&#39;&#39;.</span>
    <span id="Src">Src</span>
)</pre>









### <a id="Op.Draw">func</a> (Op) [Draw](https://golang.org/src/image/draw/draw.go?s=1169:1249#L35)
<pre>func (op <a href="#Op">Op</a>) Draw(dst <a href="#Image">Image</a>, r <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Rectangle">Rectangle</a>, src <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Image">Image</a>, sp <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Point">Point</a>)</pre>
Draw implements the Drawer interface by calling the Draw function with this
Op.




## <a id="Quantizer">type</a> [Quantizer](https://golang.org/src/image/draw/draw.go?s=687:909#L17)
Quantizer produces a palette for an image.


<pre>type Quantizer interface {
    <span class="comment">// Quantize appends up to cap(p) - len(p) colors to p and returns the</span>
    <span class="comment">// updated palette suitable for converting m to a paletted image.</span>
    Quantize(p <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Palette">Palette</a>, m <a href="/pkg/image/">image</a>.<a href="/pkg/image/#Image">Image</a>) <a href="/pkg/image/color/">color</a>.<a href="/pkg/image/color/#Palette">Palette</a>
}</pre>














