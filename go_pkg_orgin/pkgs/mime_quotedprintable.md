

# quotedprintable
`import "mime/quotedprintable"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package quotedprintable implements quoted-printable encoding as specified by
RFC 2045.




## <a id="pkg-index">Index</a>
* [type Reader](#Reader)
  * [func NewReader(r io.Reader) *Reader](#NewReader)
  * [func (r *Reader) Read(p []byte) (n int, err error)](#Reader.Read)
* [type Writer](#Writer)
  * [func NewWriter(w io.Writer) *Writer](#NewWriter)
  * [func (w *Writer) Close() error](#Writer.Close)
  * [func (w *Writer) Write(p []byte) (n int, err error)](#Writer.Write)


#### <a id="pkg-examples">Examples</a>
* [NewReader](#example_NewReader)
* [NewWriter](#example_NewWriter)


#### <a id="pkg-files">Package files</a>
[reader.go](https://golang.org/src/mime/quotedprintable/reader.go) [writer.go](https://golang.org/src/mime/quotedprintable/writer.go) 








## <a id="Reader">type</a> [Reader](https://golang.org/src/mime/quotedprintable/reader.go?s=362:485#L7)
Reader is a quoted-printable decoder.


<pre>type Reader struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewReader">func</a> [NewReader](https://golang.org/src/mime/quotedprintable/reader.go?s=552:587#L14)
<pre>func NewReader(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) *<a href="#Reader">Reader</a></pre>
NewReader returns a quoted-printable reader, decoding from r.



<a id="example_NewReader">Example</a>


```go
```

output:
```txt
```




### <a id="Reader.Read">func</a> (\*Reader) [Read](https://golang.org/src/mime/quotedprintable/reader.go?s=1504:1554#L62)
<pre>func (r *<a href="#Reader">Reader</a>) Read(p []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Read reads and decodes quoted-printable data from the underlying reader.




## <a id="Writer">type</a> [Writer](https://golang.org/src/mime/quotedprintable/writer.go?s=294:491#L2)
A Writer is a quoted-printable writer that implements io.WriteCloser.


<pre>type Writer struct {
<span id="Writer.Binary"></span>    <span class="comment">// Binary mode treats the writer&#39;s input as pure binary and processes end of</span>
    <span class="comment">// line bytes as binary data.</span>
    Binary <a href="/pkg/builtin/#bool">bool</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewWriter">func</a> [NewWriter](https://golang.org/src/mime/quotedprintable/writer.go?s=545:580#L14)
<pre>func NewWriter(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) *<a href="#Writer">Writer</a></pre>
NewWriter returns a new Writer that writes to w.



<a id="example_NewWriter">Example</a>


```go
```

output:
```txt
```




### <a id="Writer.Close">func</a> (\*Writer) [Close](https://golang.org/src/mime/quotedprintable/writer.go?s=1499:1529#L57)
<pre>func (w *<a href="#Writer">Writer</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the Writer, flushing any unwritten data to the underlying
io.Writer, but does not close the underlying io.Writer.




### <a id="Writer.Write">func</a> (\*Writer) [Write](https://golang.org/src/mime/quotedprintable/writer.go?s=822:873#L21)
<pre>func (w *<a href="#Writer">Writer</a>) Write(p []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Write encodes p using quoted-printable encoding and writes it to the
underlying io.Writer. It limits line length to 76 characters. The encoded
bytes are not necessarily flushed until the Writer is closed.







