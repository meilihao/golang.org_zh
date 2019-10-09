

# zlib
`import "compress/zlib"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package zlib implements reading and writing of zlib format compressed data,
as specified in RFC 1950.

The implementation provides filters that uncompress during reading
and compress during writing.  For example, to write compressed data
to a buffer:


	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write([]byte("hello, world\n"))
	w.Close()

and to read that data back:


	r, err := zlib.NewReader(&b)
	io.Copy(os.Stdout, r)
	r.Close()




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func NewReader(r io.Reader) (io.ReadCloser, error)](#NewReader)
* [func NewReaderDict(r io.Reader, dict []byte) (io.ReadCloser, error)](#NewReaderDict)
* [type Resetter](#Resetter)
* [type Writer](#Writer)
  * [func NewWriter(w io.Writer) *Writer](#NewWriter)
  * [func NewWriterLevel(w io.Writer, level int) (*Writer, error)](#NewWriterLevel)
  * [func NewWriterLevelDict(w io.Writer, level int, dict []byte) (*Writer, error)](#NewWriterLevelDict)
  * [func (z *Writer) Close() error](#Writer.Close)
  * [func (z *Writer) Flush() error](#Writer.Flush)
  * [func (z *Writer) Reset(w io.Writer)](#Writer.Reset)
  * [func (z *Writer) Write(p []byte) (n int, err error)](#Writer.Write)


#### <a id="pkg-examples">Examples</a>
* [NewReader](#example_NewReader)
* [NewWriter](#example_NewWriter)


#### <a id="pkg-files">Package files</a>
[reader.go](https://golang.org/src/compress/zlib/reader.go) [writer.go](https://golang.org/src/compress/zlib/writer.go) 


## <a id="pkg-constants">Constants</a>
These constants are copied from the flate package, so that code that imports
"compress/zlib" does not also have to import "compress/flate".


<pre>const (
    <span id="NoCompression">NoCompression</span>      = <a href="/pkg/compress/flate/">flate</a>.<a href="/pkg/compress/flate/#NoCompression">NoCompression</a>
    <span id="BestSpeed">BestSpeed</span>          = <a href="/pkg/compress/flate/">flate</a>.<a href="/pkg/compress/flate/#BestSpeed">BestSpeed</a>
    <span id="BestCompression">BestCompression</span>    = <a href="/pkg/compress/flate/">flate</a>.<a href="/pkg/compress/flate/#BestCompression">BestCompression</a>
    <span id="DefaultCompression">DefaultCompression</span> = <a href="/pkg/compress/flate/">flate</a>.<a href="/pkg/compress/flate/#DefaultCompression">DefaultCompression</a>
    <span id="HuffmanOnly">HuffmanOnly</span>        = <a href="/pkg/compress/flate/">flate</a>.<a href="/pkg/compress/flate/#HuffmanOnly">HuffmanOnly</a>
)</pre>

## <a id="pkg-variables">Variables</a>

<pre>var (
    <span class="comment">// ErrChecksum is returned when reading ZLIB data that has an invalid checksum.</span>
    <span id="ErrChecksum">ErrChecksum</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;zlib: invalid checksum&#34;)
    <span class="comment">// ErrDictionary is returned when reading ZLIB data that has an invalid dictionary.</span>
    <span id="ErrDictionary">ErrDictionary</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;zlib: invalid dictionary&#34;)
    <span class="comment">// ErrHeader is returned when reading ZLIB data that has an invalid header.</span>
    <span id="ErrHeader">ErrHeader</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;zlib: invalid header&#34;)
)</pre>

## <a id="NewReader">func</a> [NewReader](https://golang.org/src/compress/zlib/reader.go?s=2006:2056#L60)
<pre>func NewReader(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (<a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadCloser">ReadCloser</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewReader creates a new ReadCloser.
Reads from the returned ReadCloser read and decompress data from r.
If r does not implement io.ByteReader, the decompressor may read more
data than necessary from r.
It is the caller's responsibility to call Close on the ReadCloser when done.

The ReadCloser returned by NewReader also implements Resetter.


<a id="example_NewReader">Example</a>

## <a id="NewReaderDict">func</a> [NewReaderDict](https://golang.org/src/compress/zlib/reader.go?s=2412:2479#L69)
<pre>func NewReaderDict(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, dict []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadCloser">ReadCloser</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewReaderDict is like NewReader but uses a preset dictionary.
NewReaderDict ignores the dictionary if the compressed data does not refer to it.
If the compressed data refers to a different dictionary, NewReaderDict returns ErrDictionary.

The ReadCloser returned by NewReaderDict also implements Resetter.





## <a id="Resetter">type</a> [Resetter](https://golang.org/src/compress/zlib/reader.go?s=1456:1641#L47)
Resetter resets a ReadCloser returned by NewReader or NewReaderDict
to switch to a new underlying Reader. This permits reusing a ReadCloser
instead of allocating a new one.


<pre>type Resetter interface {
    <span class="comment">// Reset discards any buffered data and resets the Resetter as if it was</span>
    <span class="comment">// newly initialized with the given reader.</span>
    Reset(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, dict []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a>
}</pre>











## <a id="Writer">type</a> [Writer](https://golang.org/src/compress/zlib/writer.go?s=753:945#L18)
A Writer takes data written to it and writes the compressed
form of that data to an underlying writer (see NewWriter).


<pre>type Writer struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewWriter">func</a> [NewWriter](https://golang.org/src/compress/zlib/writer.go?s=1182:1217#L34)
<pre>func NewWriter(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) *<a href="#Writer">Writer</a></pre>
NewWriter creates a new Writer.
Writes to the returned Writer are compressed and written to w.

It is the caller's responsibility to call Close on the Writer when done.
Writes may be buffered and not flushed until Close.


<a id="example_NewWriter">Example</a>


### <a id="NewWriterLevel">func</a> [NewWriterLevel](https://golang.org/src/compress/zlib/writer.go?s=1616:1676#L45)
<pre>func NewWriterLevel(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, level <a href="/pkg/builtin/#int">int</a>) (*<a href="#Writer">Writer</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewWriterLevel is like NewWriter but specifies the compression level instead
of assuming DefaultCompression.

The compression level can be DefaultCompression, NoCompression, HuffmanOnly
or any integer value between BestSpeed and BestCompression inclusive.
The error returned will be nil if the level is valid.




### <a id="NewWriterLevelDict">func</a> [NewWriterLevelDict](https://golang.org/src/compress/zlib/writer.go?s=1925:2002#L54)
<pre>func NewWriterLevelDict(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, level <a href="/pkg/builtin/#int">int</a>, dict []<a href="/pkg/builtin/#byte">byte</a>) (*<a href="#Writer">Writer</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewWriterLevelDict is like NewWriterLevel but specifies a dictionary to
compress with.

The dictionary may be nil. If not, its contents should not be modified until
the Writer is closed.






### <a id="Writer.Close">func</a> (\*Writer) [Close](https://golang.org/src/compress/zlib/writer.go?s=4930:4960#L167)
<pre>func (z *<a href="#Writer">Writer</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the Writer, flushing any unwritten data to the underlying
io.Writer, but does not close the underlying io.Writer.




### <a id="Writer.Flush">func</a> (\*Writer) [Flush](https://golang.org/src/compress/zlib/writer.go?s=4630:4660#L154)
<pre>func (z *<a href="#Writer">Writer</a>) Flush() <a href="/pkg/builtin/#error">error</a></pre>
Flush flushes the Writer to its underlying io.Writer.




### <a id="Writer.Reset">func</a> (\*Writer) [Reset](https://golang.org/src/compress/zlib/writer.go?s=2368:2403#L68)
<pre>func (z *<a href="#Writer">Writer</a>) Reset(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>)</pre>
Reset clears the state of the Writer z such that it is equivalent to its
initial state from NewWriterLevel or NewWriterLevelDict, but instead writing
to w.




### <a id="Writer.Write">func</a> (\*Writer) [Write](https://golang.org/src/compress/zlib/writer.go?s=4287:4338#L134)
<pre>func (z *<a href="#Writer">Writer</a>) Write(p []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Write writes a compressed form of p to the underlying io.Writer. The
compressed bytes are not necessarily flushed until the Writer is closed or
explicitly flushed.








