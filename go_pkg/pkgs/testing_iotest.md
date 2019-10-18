

# iotest
`import "testing/iotest"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package iotest implements Readers and Writers useful mainly for testing.




## <a id="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func DataErrReader(r io.Reader) io.Reader](#DataErrReader)
* [func HalfReader(r io.Reader) io.Reader](#HalfReader)
* [func NewReadLogger(prefix string, r io.Reader) io.Reader](#NewReadLogger)
* [func NewWriteLogger(prefix string, w io.Writer) io.Writer](#NewWriteLogger)
* [func OneByteReader(r io.Reader) io.Reader](#OneByteReader)
* [func TimeoutReader(r io.Reader) io.Reader](#TimeoutReader)
* [func TruncateWriter(w io.Writer, n int64) io.Writer](#TruncateWriter)




#### <a id="pkg-files">Package files</a>
[logger.go](https://golang.org/src/testing/iotest/logger.go) [reader.go](https://golang.org/src/testing/iotest/reader.go) [writer.go](https://golang.org/src/testing/iotest/writer.go) 




## <a id="pkg-variables">Variables</a>

<pre>var <span id="ErrTimeout">ErrTimeout</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;timeout&#34;)</pre>

## <a id="DataErrReader">func</a> [DataErrReader](https://golang.org/src/testing/iotest/reader.go?s=1273:1314#L35)
<pre>func DataErrReader(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a></pre>
DataErrReader changes the way errors are handled by a Reader. Normally, a
Reader returns an error (typically EOF) from the first Read call after the
last piece of data is read. DataErrReader wraps a Reader and changes its
behavior so the final error is returned along with the final data, instead
of in the first call after the final data.



## <a id="HalfReader">func</a> [HalfReader](https://golang.org/src/testing/iotest/reader.go?s=719:757#L20)
<pre>func HalfReader(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a></pre>
HalfReader returns a Reader that implements Read
by reading half as many requested bytes from r.



## <a id="NewReadLogger">func</a> [NewReadLogger](https://golang.org/src/testing/iotest/logger.go?s=1203:1259#L42)
<pre>func NewReadLogger(prefix <a href="/pkg/builtin/#string">string</a>, r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a></pre>
NewReadLogger returns a reader that behaves like r except
that it logs (using log.Printf) each read to standard error,
printing the prefix and the hexadecimal data read.



## <a id="NewWriteLogger">func</a> [NewWriteLogger](https://golang.org/src/testing/iotest/logger.go?s=659:716#L20)
<pre>func NewWriteLogger(prefix <a href="/pkg/builtin/#string">string</a>, w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a></pre>
NewWriteLogger returns a writer that behaves like w except
that it logs (using log.Printf) each write to standard error,
printing the prefix and the hexadecimal data written.



## <a id="OneByteReader">func</a> [OneByteReader](https://golang.org/src/testing/iotest/reader.go?s=381:422#L5)
<pre>func OneByteReader(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a></pre>
OneByteReader returns a Reader that implements
each non-empty Read by reading one byte from r.



## <a id="TimeoutReader">func</a> [TimeoutReader](https://golang.org/src/testing/iotest/reader.go?s=1969:2010#L65)
<pre>func TimeoutReader(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a></pre>
TimeoutReader returns ErrTimeout on the second read
with no data. Subsequent calls to read succeed.



## <a id="TruncateWriter">func</a> [TruncateWriter](https://golang.org/src/testing/iotest/writer.go?s=278:329#L1)
<pre>func TruncateWriter(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, n <a href="/pkg/builtin/#int64">int64</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a></pre>
TruncateWriter returns a Writer that writes to w
but stops silently after n bytes.









