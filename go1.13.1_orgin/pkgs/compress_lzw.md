

# lzw
`import "compress/lzw"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package lzw implements the Lempel-Ziv-Welch compressed data format,
described in T. A. Welch, ``A Technique for High-Performance Data
Compression'', Computer, 17(6) (June 1984), pp 8-19.

In particular, it implements LZW as used by the GIF and PDF file
formats, which means variable-width codes up to 12 bits and the first
two non-literal codes are a clear code and an EOF code.

The TIFF file format uses a similar but incompatible version of the LZW
algorithm. See the golang.org/x/image/tiff/lzw package for an
implementation.




## <a id="pkg-index">Index</a>
* [func NewReader(r io.Reader, order Order, litWidth int) io.ReadCloser](#NewReader)
* [func NewWriter(w io.Writer, order Order, litWidth int) io.WriteCloser](#NewWriter)
* [type Order](#Order)




#### <a id="pkg-files">Package files</a>
[reader.go](https://golang.org/src/compress/lzw/reader.go) [writer.go](https://golang.org/src/compress/lzw/writer.go) 






## <a id="NewReader">func</a> [NewReader](https://golang.org/src/compress/lzw/reader.go?s=6772:6840#L229)
<pre>func NewReader(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, order <a href="#Order">Order</a>, litWidth <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadCloser">ReadCloser</a></pre>
NewReader creates a new io.ReadCloser.
Reads from the returned io.ReadCloser read and decompress data from r.
If r does not also implement io.ByteReader,
the decompressor may read more data than necessary from r.
It is the caller's responsibility to call Close on the ReadCloser when
finished reading.
The number of bits to use for literal codes, litWidth, must be in the
range [2,8] and is typically 8. It must equal the litWidth
used during compression.



## <a id="NewWriter">func</a> [NewWriter](https://golang.org/src/compress/lzw/writer.go?s=6435:6504#L231)
<pre>func NewWriter(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, order <a href="#Order">Order</a>, litWidth <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#WriteCloser">WriteCloser</a></pre>
NewWriter creates a new io.WriteCloser.
Writes to the returned io.WriteCloser are compressed and written to w.
It is the caller's responsibility to call Close on the WriteCloser when
finished writing.
The number of bits to use for literal codes, litWidth, must be in the
range [2,8] and is typically 8. Input bytes must be less than 1<<litWidth.





## <a id="Order">type</a> [Order](https://golang.org/src/compress/lzw/reader.go?s=938:952#L19)
Order specifies the bit ordering in an LZW data stream.


<pre>type Order <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span class="comment">// LSB means Least Significant Bits first, as used in the GIF file format.</span>
    <span id="LSB">LSB</span> <a href="#Order">Order</a> = <a href="/pkg/builtin/#iota">iota</a>
    <span class="comment">// MSB means Most Significant Bits first, as used in the TIFF and PDF</span>
    <span class="comment">// file formats.</span>
    <span id="MSB">MSB</span>
)</pre>













