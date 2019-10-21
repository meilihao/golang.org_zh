

# gzip
`import "compress/gzip"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package gzip implements reading and writing of gzip format compressed files,
as specified in RFC 1952.


<a id="example__writerReader">Example (WriterReader)</a>
```go
```

output:
```txt
```


## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [type Header](#Header)
* [type Reader](#Reader)
  * [func NewReader(r io.Reader) (*Reader, error)](#NewReader)
  * [func (z *Reader) Close() error](#Reader.Close)
  * [func (z *Reader) Multistream(ok bool)](#Reader.Multistream)
  * [func (z *Reader) Read(p []byte) (n int, err error)](#Reader.Read)
  * [func (z *Reader) Reset(r io.Reader) error](#Reader.Reset)
* [type Writer](#Writer)
  * [func NewWriter(w io.Writer) *Writer](#NewWriter)
  * [func NewWriterLevel(w io.Writer, level int) (*Writer, error)](#NewWriterLevel)
  * [func (z *Writer) Close() error](#Writer.Close)
  * [func (z *Writer) Flush() error](#Writer.Flush)
  * [func (z *Writer) Reset(w io.Writer)](#Writer.Reset)
  * [func (z *Writer) Write(p []byte) (int, error)](#Writer.Write)


#### <a id="pkg-examples">Examples</a>
* [Reader.Multistream](#example_Reader_Multistream)
* [Package (WriterReader)](#example__writerReader)


#### <a id="pkg-files">Package files</a>
[gunzip.go](https://golang.org/src/compress/gzip/gunzip.go) [gzip.go](https://golang.org/src/compress/gzip/gzip.go) 


## <a id="pkg-constants">Constants</a>
These constants are copied from the flate package, so that code that imports
"compress/gzip" does not also have to import "compress/flate".


<pre>const (
    <span id="NoCompression">NoCompression</span>      = <a href="/pkg/compress/flate/">flate</a>.<a href="/pkg/compress/flate/#NoCompression">NoCompression</a>
    <span id="BestSpeed">BestSpeed</span>          = <a href="/pkg/compress/flate/">flate</a>.<a href="/pkg/compress/flate/#BestSpeed">BestSpeed</a>
    <span id="BestCompression">BestCompression</span>    = <a href="/pkg/compress/flate/">flate</a>.<a href="/pkg/compress/flate/#BestCompression">BestCompression</a>
    <span id="DefaultCompression">DefaultCompression</span> = <a href="/pkg/compress/flate/">flate</a>.<a href="/pkg/compress/flate/#DefaultCompression">DefaultCompression</a>
    <span id="HuffmanOnly">HuffmanOnly</span>        = <a href="/pkg/compress/flate/">flate</a>.<a href="/pkg/compress/flate/#HuffmanOnly">HuffmanOnly</a>
)</pre>

## <a id="pkg-variables">Variables</a>

<pre>var (
    <span class="comment">// ErrChecksum is returned when reading GZIP data that has an invalid checksum.</span>
    <span id="ErrChecksum">ErrChecksum</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;gzip: invalid checksum&#34;)
    <span class="comment">// ErrHeader is returned when reading GZIP data that has an invalid header.</span>
    <span id="ErrHeader">ErrHeader</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;gzip: invalid header&#34;)
)</pre>



## <a id="Header">type</a> [Header](https://golang.org/src/compress/gzip/gunzip.go?s=1297:1500#L42)
The gzip file stores a header giving metadata about the compressed file.
That header is exposed as the fields of the Writer and Reader structs.

Strings must be UTF-8 encoded and may only contain Unicode code points
U+0001 through U+00FF, due to limitations of the GZIP file format.


<pre>type Header struct {
<span id="Header.Comment"></span>    Comment <a href="/pkg/builtin/#string">string</a>    <span class="comment">// comment</span>
<span id="Header.Extra"></span>    Extra   []<a href="/pkg/builtin/#byte">byte</a>    <span class="comment">// &#34;extra data&#34;</span>
<span id="Header.ModTime"></span>    ModTime <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a> <span class="comment">// modification time</span>
<span id="Header.Name"></span>    Name    <a href="/pkg/builtin/#string">string</a>    <span class="comment">// file name</span>
<span id="Header.OS"></span>    OS      <a href="/pkg/builtin/#byte">byte</a>      <span class="comment">// operating system type</span>
}
</pre>











## <a id="Reader">type</a> [Reader](https://golang.org/src/compress/gzip/gunzip.go?s=2199:2512#L64)
A Reader is an io.Reader that can be read to retrieve
uncompressed data from a gzip-format compressed file.

In general, a gzip file can be a concatenation of gzip files,
each with its own header. Reads from the Reader
return the concatenation of the uncompressed data of each.
Only the first header is recorded in the Reader fields.

Gzip files store a length and checksum of the uncompressed data.
The Reader will return an ErrChecksum when Read
reaches the end of the uncompressed data if it does not
have the expected length or checksum. Clients should treat data
returned by Read as tentative until they receive the io.EOF
marking the end of the data.


<pre>type Reader struct {
    <a href="#Header">Header</a> <span class="comment">// valid after NewReader or Reader.Reset</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewReader">func</a> [NewReader](https://golang.org/src/compress/gzip/gunzip.go?s=2831:2875#L82)
<pre>func NewReader(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (*<a href="#Reader">Reader</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewReader creates a new Reader reading the given reader.
If r does not also implement io.ByteReader,
the decompressor may read more data than necessary from r.

It is the caller's responsibility to call Close on the Reader when done.

The Reader.Header fields will be valid in the Reader returned.






### <a id="Reader.Close">func</a> (\*Reader) [Close](https://golang.org/src/compress/gzip/gunzip.go?s=8554:8584#L282)
<pre>func (z *<a href="#Reader">Reader</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the Reader. It does not close the underlying io.Reader.
In order for the GZIP checksum to be verified, the reader must be
fully consumed until the io.EOF.




### <a id="Reader.Multistream">func</a> (\*Reader) [Multistream](https://golang.org/src/compress/gzip/gunzip.go?s=4446:4483#L123)
<pre>func (z *<a href="#Reader">Reader</a>) Multistream(ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
Multistream controls whether the reader supports multistream files.

If enabled (the default), the Reader expects the input to be a sequence
of individually gzipped data streams, each with its own header and
trailer, ending at EOF. The effect is that the concatenation of a sequence
of gzipped files is treated as equivalent to the gzip of the concatenation
of the sequence. This is standard behavior for gzip readers.

Calling Multistream(false) disables this behavior; disabling the behavior
can be useful when reading file formats that distinguish individual gzip
data streams or mix gzip data streams with other data streams.
In this mode, when the Reader reaches the end of the data stream,
Read returns io.EOF. The underlying reader must implement io.ByteReader
in order to be left positioned just after the gzip stream.
To start the next stream, call z.Reset(r) followed by z.Multistream(false).
If there is no next stream, z.Reset(r) will return io.EOF.


<a id="example_Reader_Multistream">Example</a>
```go
```

output:
```txt
```


### <a id="Reader.Read">func</a> (\*Reader) [Read](https://golang.org/src/compress/gzip/gunzip.go?s=7491:7541#L236)
<pre>func (z *<a href="#Reader">Reader</a>) Read(p []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Read implements io.Reader, reading uncompressed bytes from its underlying Reader.




### <a id="Reader.Reset">func</a> (\*Reader) [Reset](https://golang.org/src/compress/gzip/gunzip.go?s=3184:3225#L93)
<pre>func (z *<a href="#Reader">Reader</a>) Reset(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) <a href="/pkg/builtin/#error">error</a></pre>
Reset discards the Reader z's state and makes it equivalent to the
result of its original state from NewReader, but reading from r instead.
This permits reusing a Reader rather than allocating a new one.




## <a id="Writer">type</a> [Writer](https://golang.org/src/compress/gzip/gzip.go?s=706:1052#L18)
A Writer is an io.WriteCloser.
Writes to a Writer are compressed and written to w.


<pre>type Writer struct {
    <a href="#Header">Header</a> <span class="comment">// written at first call to Write, Flush, or Close</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewWriter">func</a> [NewWriter](https://golang.org/src/compress/gzip/gzip.go?s=1411:1446#L39)
<pre>func NewWriter(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) *<a href="#Writer">Writer</a></pre>
NewWriter returns a new Writer.
Writes to the returned writer are compressed and written to w.

It is the caller's responsibility to call Close on the Writer when done.
Writes may be buffered and not flushed until Close.

Callers that wish to set the fields in Writer.Header must do so before
the first call to Write, Flush, or Close.




### <a id="NewWriterLevel">func</a> [NewWriterLevel](https://golang.org/src/compress/gzip/gzip.go?s=1836:1896#L50)
<pre>func NewWriterLevel(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, level <a href="/pkg/builtin/#int">int</a>) (*<a href="#Writer">Writer</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewWriterLevel is like NewWriter but specifies the compression level instead
of assuming DefaultCompression.

The compression level can be DefaultCompression, NoCompression, HuffmanOnly
or any integer value between BestSpeed and BestCompression inclusive.
The error returned will be nil if the level is valid.






### <a id="Writer.Close">func</a> (\*Writer) [Close](https://golang.org/src/compress/gzip/gzip.go?s=5974:6004#L218)
<pre>func (z *<a href="#Writer">Writer</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the Writer by flushing any unwritten data to the underlying
io.Writer and writing the GZIP footer.
It does not close the underlying io.Writer.




### <a id="Writer.Flush">func</a> (\*Writer) [Flush](https://golang.org/src/compress/gzip/gzip.go?s=5582:5612#L198)
<pre>func (z *<a href="#Writer">Writer</a>) Flush() <a href="/pkg/builtin/#error">error</a></pre>
Flush flushes any pending compressed data to the underlying writer.

It is useful mainly in compressed network protocols, to ensure that
a remote reader has enough data to reconstruct a packet. Flush does
not return until the data has been written. If the underlying
writer returns an error, Flush returns that error.

In the terminology of the zlib library, Flush is equivalent to Z_SYNC_FLUSH.




### <a id="Writer.Reset">func</a> (\*Writer) [Reset](https://golang.org/src/compress/gzip/gzip.go?s=2567:2602#L78)
<pre>func (z *<a href="#Writer">Writer</a>) Reset(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>)</pre>
Reset discards the Writer z's state and makes it equivalent to the
result of its original state from NewWriter or NewWriterLevel, but
writing to w instead. This permits reusing a Writer rather than
allocating a new one.




### <a id="Writer.Write">func</a> (\*Writer) [Write](https://golang.org/src/compress/gzip/gzip.go?s=3862:3907#L129)
<pre>func (z *<a href="#Writer">Writer</a>) Write(p []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Write writes a compressed form of p to the underlying io.Writer. The
compressed bytes are not necessarily flushed until the Writer is closed.







