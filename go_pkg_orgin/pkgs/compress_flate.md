

# flate
`import "compress/flate"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package flate implements the DEFLATE compressed data format, described in
RFC 1951.  The gzip and zlib packages implement access to DEFLATE-based file
formats.


<a id="example__dictionary">Example (Dictionary)</a>
```go
```

output:
```txt
```
<p>A preset dictionary can be used to improve the compression ratio.
The downside to using a dictionary is that the compressor and decompressor
must agree in advance what dictionary to use.
</p><a id="example__reset">Example (Reset)</a>
```go
```

output:
```txt
```
<p>In performance critical applications, Reset can be used to discard the
current compressor or decompressor state and reinitialize them quickly
by taking advantage of previously allocated memory.
</p><a id="example__synchronization">Example (Synchronization)</a>
```go
```

output:
```txt
```
<p>DEFLATE is suitable for transmitting compressed data across the network.
</p>

## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func NewReader(r io.Reader) io.ReadCloser](#NewReader)
* [func NewReaderDict(r io.Reader, dict []byte) io.ReadCloser](#NewReaderDict)
* [type CorruptInputError](#CorruptInputError)
  * [func (e CorruptInputError) Error() string](#CorruptInputError.Error)
* [type InternalError](#InternalError)
  * [func (e InternalError) Error() string](#InternalError.Error)
* [type ReadError](#ReadError)
  * [func (e *ReadError) Error() string](#ReadError.Error)
* [type Reader](#Reader)
* [type Resetter](#Resetter)
* [type WriteError](#WriteError)
  * [func (e *WriteError) Error() string](#WriteError.Error)
* [type Writer](#Writer)
  * [func NewWriter(w io.Writer, level int) (*Writer, error)](#NewWriter)
  * [func NewWriterDict(w io.Writer, level int, dict []byte) (*Writer, error)](#NewWriterDict)
  * [func (w *Writer) Close() error](#Writer.Close)
  * [func (w *Writer) Flush() error](#Writer.Flush)
  * [func (w *Writer) Reset(dst io.Writer)](#Writer.Reset)
  * [func (w *Writer) Write(data []byte) (n int, err error)](#Writer.Write)


#### <a id="pkg-examples">Examples</a>
* [Package (Dictionary)](#example__dictionary)
* [Package (Reset)](#example__reset)
* [Package (Synchronization)](#example__synchronization)


#### <a id="pkg-files">Package files</a>
[deflate.go](https://golang.org/src/compress/flate/deflate.go) [deflatefast.go](https://golang.org/src/compress/flate/deflatefast.go) [dict_decoder.go](https://golang.org/src/compress/flate/dict_decoder.go) [huffman_bit_writer.go](https://golang.org/src/compress/flate/huffman_bit_writer.go) [huffman_code.go](https://golang.org/src/compress/flate/huffman_code.go) [inflate.go](https://golang.org/src/compress/flate/inflate.go) [token.go](https://golang.org/src/compress/flate/token.go) 


## <a id="pkg-constants">Constants</a>

<pre>const (
    <span id="NoCompression">NoCompression</span>      = 0
    <span id="BestSpeed">BestSpeed</span>          = 1
    <span id="BestCompression">BestCompression</span>    = 9
    <span id="DefaultCompression">DefaultCompression</span> = -1

    <span class="comment">// HuffmanOnly disables Lempel-Ziv match searching and only performs Huffman</span>
    <span class="comment">// entropy encoding. This mode is useful in compressing data that has</span>
    <span class="comment">// already been compressed with an LZ style algorithm (e.g. Snappy or LZ4)</span>
    <span class="comment">// that lacks an entropy encoder. Compression gains are achieved when</span>
    <span class="comment">// certain bytes in the input stream occur more frequently than others.</span>
    <span class="comment">//</span>
    <span class="comment">// Note that HuffmanOnly produces a compressed output that is</span>
    <span class="comment">// RFC 1951 compliant. That is, any valid DEFLATE decompressor will</span>
    <span class="comment">// continue to be able to decompress this output.</span>
    <span id="HuffmanOnly">HuffmanOnly</span> = -2
)</pre>



## <a id="NewReader">func</a> [NewReader](https://golang.org/src/compress/flate/inflate.go?s=19470:19511#L786)
<pre>func NewReader(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadCloser">ReadCloser</a></pre>
NewReader returns a new ReadCloser that can be used
to read the uncompressed version of r.
If r does not also implement io.ByteReader,
the decompressor may read more data than necessary from r.
It is the caller's responsibility to call Close on the ReadCloser
when finished reading.

The ReadCloser returned by NewReader also implements Resetter.



## <a id="NewReaderDict">func</a> [NewReaderDict](https://golang.org/src/compress/flate/inflate.go?s=20113:20171#L805)
<pre>func NewReaderDict(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, dict []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadCloser">ReadCloser</a></pre>
NewReaderDict is like NewReader but initializes the reader
with a preset dictionary. The returned Reader behaves as if
the uncompressed data stream started with the given dictionary,
which has already been read. NewReaderDict is typically used
to read data compressed by NewWriterDict.

The ReadCloser returned by NewReader also implements Resetter.





## <a id="CorruptInputError">type</a> [CorruptInputError](https://golang.org/src/compress/flate/inflate.go?s=957:985#L23)
A CorruptInputError reports the presence of corrupt input at a given offset.


<pre>type CorruptInputError <a href="/pkg/builtin/#int64">int64</a></pre>











### <a id="CorruptInputError.Error">func</a> (CorruptInputError) [Error](https://golang.org/src/compress/flate/inflate.go?s=987:1028#L25)
<pre>func (e <a href="#CorruptInputError">CorruptInputError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="InternalError">type</a> [InternalError](https://golang.org/src/compress/flate/inflate.go?s=1177:1202#L30)
An InternalError reports an error in the flate code itself.


<pre>type InternalError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="InternalError.Error">func</a> (InternalError) [Error](https://golang.org/src/compress/flate/inflate.go?s=1204:1241#L32)
<pre>func (e <a href="#InternalError">InternalError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="ReadError">type</a> [ReadError](https://golang.org/src/compress/flate/inflate.go?s=1395:1521#L37)
A ReadError reports an error encountered while reading input.

Deprecated: No longer returned.


<pre>type ReadError struct {
<span id="ReadError.Offset"></span>    Offset <a href="/pkg/builtin/#int64">int64</a> <span class="comment">// byte offset where error occurred</span>
<span id="ReadError.Err"></span>    Err    <a href="/pkg/builtin/#error">error</a> <span class="comment">// error returned by underlying Read</span>
}
</pre>











### <a id="ReadError.Error">func</a> (\*ReadError) [Error](https://golang.org/src/compress/flate/inflate.go?s=1523:1557#L42)
<pre>func (e *<a href="#ReadError">ReadError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Reader">type</a> [Reader](https://golang.org/src/compress/flate/inflate.go?s=8219:8270#L251)
The actual read interface needed by NewReader.
If the passed in io.Reader does not also have ReadByte,
the NewReader will introduce its own buffering.


<pre>type Reader interface {
    <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>
    <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ByteReader">ByteReader</a>
}</pre>











## <a id="Resetter">type</a> [Resetter](https://golang.org/src/compress/flate/inflate.go?s=2214:2399#L61)
Resetter resets a ReadCloser returned by NewReader or NewReaderDict
to switch to a new underlying Reader. This permits reusing a ReadCloser
instead of allocating a new one.


<pre>type Resetter interface {
    <span class="comment">// Reset discards any buffered data and resets the Resetter as if it was</span>
    <span class="comment">// newly initialized with the given reader.</span>
    Reset(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, dict []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a>
}</pre>











## <a id="WriteError">type</a> [WriteError](https://golang.org/src/compress/flate/inflate.go?s=1764:1892#L49)
A WriteError reports an error encountered while writing output.

Deprecated: No longer returned.


<pre>type WriteError struct {
<span id="WriteError.Offset"></span>    Offset <a href="/pkg/builtin/#int64">int64</a> <span class="comment">// byte offset where error occurred</span>
<span id="WriteError.Err"></span>    Err    <a href="/pkg/builtin/#error">error</a> <span class="comment">// error returned by underlying Write</span>
}
</pre>











### <a id="WriteError.Error">func</a> (\*WriteError) [Error](https://golang.org/src/compress/flate/inflate.go?s=1894:1929#L54)
<pre>func (e *<a href="#WriteError">WriteError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Writer">type</a> [Writer](https://golang.org/src/compress/flate/deflate.go?s=19492:19544#L691)
A Writer takes data written to it and writes the compressed
form of that data to an underlying writer (see NewWriter).


<pre>type Writer struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewWriter">func</a> [NewWriter](https://golang.org/src/compress/flate/deflate.go?s=18469:18524#L656)
<pre>func NewWriter(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, level <a href="/pkg/builtin/#int">int</a>) (*<a href="#Writer">Writer</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewWriter returns a new Writer compressing data at the given level.
Following zlib, levels range from 1 (BestSpeed) to 9 (BestCompression);
higher levels typically run slower but compress more. Level 0
(NoCompression) does not attempt any compression; it only adds the
necessary DEFLATE framing.
Level -1 (DefaultCompression) uses the default compression level.
Level -2 (HuffmanOnly) will use Huffman compression only, giving
a very fast compression for all types of input, but sacrificing considerable
compression efficiency.

If level is in the range [-2, 9] then the error returned will be nil.
Otherwise the error returned will be non-nil.




### <a id="NewWriterDict">func</a> [NewWriterDict](https://golang.org/src/compress/flate/deflate.go?s=18956:19028#L670)
<pre>func NewWriterDict(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, level <a href="/pkg/builtin/#int">int</a>, dict []<a href="/pkg/builtin/#byte">byte</a>) (*<a href="#Writer">Writer</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewWriterDict is like NewWriter but initializes the new
Writer with a preset dictionary. The returned Writer behaves
as if the dictionary had been written to it without producing
any compressed output. The compressed data written to w
can only be decompressed by a Reader initialized with the
same dictionary.






### <a id="Writer.Close">func</a> (\*Writer) [Close](https://golang.org/src/compress/flate/deflate.go?s=20445:20475#L718)
<pre>func (w *<a href="#Writer">Writer</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close flushes and closes the writer.




### <a id="Writer.Flush">func</a> (\*Writer) [Flush](https://golang.org/src/compress/flate/deflate.go?s=20263:20293#L711)
<pre>func (w *<a href="#Writer">Writer</a>) Flush() <a href="/pkg/builtin/#error">error</a></pre>
Flush flushes any pending data to the underlying writer.
It is useful mainly in compressed network protocols, to ensure that
a remote reader has enough data to reconstruct a packet.
Flush does not return until the data has been written.
Calling Flush when there is no pending data still causes the Writer
to emit a sync marker of at least 4 bytes.
If the underlying writer returns an error, Flush returns that error.

In the terminology of the zlib library, Flush is equivalent to Z_SYNC_FLUSH.




### <a id="Writer.Reset">func</a> (\*Writer) [Reset](https://golang.org/src/compress/flate/deflate.go?s=20658:20695#L725)
<pre>func (w *<a href="#Writer">Writer</a>) Reset(dst <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>)</pre>
Reset discards the writer's state and makes it equivalent to
the result of NewWriter or NewWriterDict called with dst
and w's level and dictionary.




### <a id="Writer.Write">func</a> (\*Writer) [Write](https://golang.org/src/compress/flate/deflate.go?s=19658:19712#L698)
<pre>func (w *<a href="#Writer">Writer</a>) Write(data []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Write writes data to w, which will eventually write the
compressed form of data to its underlying writer.







