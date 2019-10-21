

# io
`import "io"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package io provides basic interfaces to I/O primitives.
Its primary job is to wrap existing implementations of such primitives,
such as those in package os, into shared public interfaces that
abstract the functionality, plus some other related primitives.

Because these interfaces and primitives wrap lower-level operations with
various implementations, unless otherwise informed clients should not
assume they are safe for parallel execution.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func Copy(dst Writer, src Reader) (written int64, err error)](#Copy)
* [func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)](#CopyBuffer)
* [func CopyN(dst Writer, src Reader, n int64) (written int64, err error)](#CopyN)
* [func Pipe() (*PipeReader, *PipeWriter)](#Pipe)
* [func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)](#ReadAtLeast)
* [func ReadFull(r Reader, buf []byte) (n int, err error)](#ReadFull)
* [func WriteString(w Writer, s string) (n int, err error)](#WriteString)
* [type ByteReader](#ByteReader)
* [type ByteScanner](#ByteScanner)
* [type ByteWriter](#ByteWriter)
* [type Closer](#Closer)
* [type LimitedReader](#LimitedReader)
  * [func (l *LimitedReader) Read(p []byte) (n int, err error)](#LimitedReader.Read)
* [type PipeReader](#PipeReader)
  * [func (r *PipeReader) Close() error](#PipeReader.Close)
  * [func (r *PipeReader) CloseWithError(err error) error](#PipeReader.CloseWithError)
  * [func (r *PipeReader) Read(data []byte) (n int, err error)](#PipeReader.Read)
* [type PipeWriter](#PipeWriter)
  * [func (w *PipeWriter) Close() error](#PipeWriter.Close)
  * [func (w *PipeWriter) CloseWithError(err error) error](#PipeWriter.CloseWithError)
  * [func (w *PipeWriter) Write(data []byte) (n int, err error)](#PipeWriter.Write)
* [type ReadCloser](#ReadCloser)
* [type ReadSeeker](#ReadSeeker)
* [type ReadWriteCloser](#ReadWriteCloser)
* [type ReadWriteSeeker](#ReadWriteSeeker)
* [type ReadWriter](#ReadWriter)
* [type Reader](#Reader)
  * [func LimitReader(r Reader, n int64) Reader](#LimitReader)
  * [func MultiReader(readers ...Reader) Reader](#MultiReader)
  * [func TeeReader(r Reader, w Writer) Reader](#TeeReader)
* [type ReaderAt](#ReaderAt)
* [type ReaderFrom](#ReaderFrom)
* [type RuneReader](#RuneReader)
* [type RuneScanner](#RuneScanner)
* [type SectionReader](#SectionReader)
  * [func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader](#NewSectionReader)
  * [func (s *SectionReader) Read(p []byte) (n int, err error)](#SectionReader.Read)
  * [func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)](#SectionReader.ReadAt)
  * [func (s *SectionReader) Seek(offset int64, whence int) (int64, error)](#SectionReader.Seek)
  * [func (s *SectionReader) Size() int64](#SectionReader.Size)
* [type Seeker](#Seeker)
* [type StringWriter](#StringWriter)
* [type WriteCloser](#WriteCloser)
* [type WriteSeeker](#WriteSeeker)
* [type Writer](#Writer)
  * [func MultiWriter(writers ...Writer) Writer](#MultiWriter)
* [type WriterAt](#WriterAt)
* [type WriterTo](#WriterTo)


#### <a id="pkg-examples">Examples</a>
* [Copy](#example_Copy)
* [CopyBuffer](#example_CopyBuffer)
* [CopyN](#example_CopyN)
* [LimitReader](#example_LimitReader)
* [MultiReader](#example_MultiReader)
* [MultiWriter](#example_MultiWriter)
* [Pipe](#example_Pipe)
* [ReadAtLeast](#example_ReadAtLeast)
* [ReadFull](#example_ReadFull)
* [SectionReader](#example_SectionReader)
* [SectionReader.ReadAt](#example_SectionReader_ReadAt)
* [SectionReader.Seek](#example_SectionReader_Seek)
* [TeeReader](#example_TeeReader)
* [WriteString](#example_WriteString)


#### <a id="pkg-files">Package files</a>
[io.go](https://golang.org/src/io/io.go) [multi.go](https://golang.org/src/io/multi.go) [pipe.go](https://golang.org/src/io/pipe.go) 


## <a id="pkg-constants">Constants</a>
Seek whence values.


<pre>const (
    <span id="SeekStart">SeekStart</span>   = 0 <span class="comment">// seek relative to the origin of the file</span>
    <span id="SeekCurrent">SeekCurrent</span> = 1 <span class="comment">// seek relative to the current offset</span>
    <span id="SeekEnd">SeekEnd</span>     = 2 <span class="comment">// seek relative to the end</span>
)</pre>

## <a id="pkg-variables">Variables</a>
EOF is the error returned by Read when no more input is available.
Functions should return EOF only to signal a graceful end of input.
If the EOF occurs unexpectedly in a structured data stream,
the appropriate error is either ErrUnexpectedEOF or some other error
giving more detail.


<pre>var <span id="EOF">EOF</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;EOF&#34;)</pre>ErrClosedPipe is the error used for read or write operations on a closed pipe.


<pre>var <span id="ErrClosedPipe">ErrClosedPipe</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;io: read/write on closed pipe&#34;)</pre>ErrNoProgress is returned by some clients of an io.Reader when
many calls to Read have failed to return any data or error,
usually the sign of a broken io.Reader implementation.


<pre>var <span id="ErrNoProgress">ErrNoProgress</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;multiple Read calls return no data or error&#34;)</pre>ErrShortBuffer means that a read required a longer buffer than was provided.


<pre>var <span id="ErrShortBuffer">ErrShortBuffer</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;short buffer&#34;)</pre>ErrShortWrite means that a write accepted fewer bytes than requested
but failed to return an explicit error.


<pre>var <span id="ErrShortWrite">ErrShortWrite</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;short write&#34;)</pre>ErrUnexpectedEOF means that EOF was encountered in the
middle of reading a fixed-size block or data structure.


<pre>var <span id="ErrUnexpectedEOF">ErrUnexpectedEOF</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;unexpected EOF&#34;)</pre>

## <a id="Copy">func</a> [Copy](https://golang.org/src/io/io.go?s=12796:12856#L353)
<pre>func Copy(dst <a href="#Writer">Writer</a>, src <a href="#Reader">Reader</a>) (written <a href="/pkg/builtin/#int64">int64</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Copy copies from src to dst until either EOF is reached
on src or an error occurs. It returns the number of bytes
copied and the first error encountered while copying, if any.

A successful Copy returns err == nil, not err == EOF.
Because Copy is defined to read from src until EOF, it does
not treat an EOF from Read as an error to be reported.

If src implements the WriterTo interface,
the copy is implemented by calling src.WriteTo(dst).
Otherwise, if dst implements the ReaderFrom interface,
the copy is implemented by calling dst.ReadFrom(src).

Copy 会将src的数据拷贝到dst，直到在src上碰到EOF或发生错误. 它会返回已拷贝的字节数和遇到的第一个错误(如果有的话).

Copy方法成功时返回值err应为nil而非EOF. 因为Copy被定义为从src读取直到EOF，它不会将读取到的EOF视为错误.

如果src实现了WriterTo接口，本函数会调用src.WriteTo(dst)进行拷贝. 否则如果dst实现了ReaderFrom接口，本函数会调用dst.ReadFrom(src)进行拷贝.

<a id="example_Copy">Example</a>
```go
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}
```

output:
```txt
some io.Reader stream to be read
```

## <a id="CopyBuffer">func</a> [CopyBuffer](https://golang.org/src/io/io.go?s=13136:13214#L361)
<pre>func CopyBuffer(dst <a href="#Writer">Writer</a>, src <a href="#Reader">Reader</a>, buf []<a href="/pkg/builtin/#byte">byte</a>) (written <a href="/pkg/builtin/#int64">int64</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
CopyBuffer is identical to Copy except that it stages through the
provided buffer (if one is required) rather than allocating a
temporary one. If buf is nil, one is allocated; otherwise if it has
zero length, CopyBuffer panics.


<a id="example_CopyBuffer">Example</a>
```go
```

output:
```txt
```

## <a id="CopyN">func</a> [CopyN](https://golang.org/src/io/io.go?s=11951:12021#L329)
<pre>func CopyN(dst <a href="#Writer">Writer</a>, src <a href="#Reader">Reader</a>, n <a href="/pkg/builtin/#int64">int64</a>) (written <a href="/pkg/builtin/#int64">int64</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
CopyN copies n bytes (or until an error) from src to dst.
It returns the number of bytes copied and the earliest
error encountered while copying.
On return, written == n if and only if err == nil.

If dst implements the ReaderFrom interface,
the copy is implemented using it.


<a id="example_CopyN">Example</a>
```go
```

output:
```txt
```

## <a id="Pipe">func</a> [Pipe](https://golang.org/src/io/pipe.go?s=4782:4820#L176)
<pre>func Pipe() (*<a href="#PipeReader">PipeReader</a>, *<a href="#PipeWriter">PipeWriter</a>)</pre>
Pipe creates a synchronous in-memory pipe.
It can be used to connect code expecting an io.Reader
with code expecting an io.Writer.

Reads and Writes on the pipe are matched one to one
except when multiple Reads are needed to consume a single Write.
That is, each Write to the PipeWriter blocks until it has satisfied
one or more Reads from the PipeReader that fully consume
the written data.
The data is copied directly from the Write to the corresponding
Read (or Reads); there is no internal buffering.

It is safe to call Read and Write in parallel with each other or with Close.
Parallel calls to Read and parallel calls to Write are also safe:
the individual calls will be gated sequentially.


<a id="example_Pipe">Example</a>
```go
```

output:
```txt
```

## <a id="ReadAtLeast">func</a> [ReadAtLeast](https://golang.org/src/io/io.go?s=10829:10895#L294)
<pre>func ReadAtLeast(r <a href="#Reader">Reader</a>, buf []<a href="/pkg/builtin/#byte">byte</a>, min <a href="/pkg/builtin/#int">int</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadAtLeast reads from r into buf until it has read at least min bytes.
It returns the number of bytes copied and an error if fewer bytes were read.
The error is EOF only if no bytes were read.
If an EOF happens after reading fewer than min bytes,
ReadAtLeast returns ErrUnexpectedEOF.
If min is greater than the length of buf, ReadAtLeast returns ErrShortBuffer.
On return, n >= min if and only if err == nil.
If r returns an error having read at least min bytes, the error is dropped.


<a id="example_ReadAtLeast">Example</a>
```go
```

output:
```txt
```

## <a id="ReadFull">func</a> [ReadFull](https://golang.org/src/io/io.go?s=11557:11611#L318)
<pre>func ReadFull(r <a href="#Reader">Reader</a>, buf []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadFull reads exactly len(buf) bytes from r into buf.
It returns the number of bytes copied and an error if fewer bytes were read.
The error is EOF only if no bytes were read.
If an EOF happens after reading some but not all the bytes,
ReadFull returns ErrUnexpectedEOF.
On return, n == len(buf) if and only if err == nil.
If r returns an error having read at least len(buf) bytes, the error is dropped.


<a id="example_ReadFull">Example</a>
```go
```

output:
```txt
```

## <a id="WriteString">func</a> [WriteString](https://golang.org/src/io/io.go?s=10163:10218#L279)
<pre>func WriteString(w <a href="#Writer">Writer</a>, s <a href="/pkg/builtin/#string">string</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
WriteString writes the contents of the string s to w, which accepts a slice of bytes.
If w implements StringWriter, its WriteString method is invoked directly.
Otherwise, w.Write is called exactly once.


<a id="example_WriteString">Example</a>
```go
```

output:
```txt
```



## <a id="ByteReader">type</a> [ByteReader](https://golang.org/src/io/io.go?s=8610:8665#L229)
ByteReader is the interface that wraps the ReadByte method.

ReadByte reads and returns the next byte from the input or
any error encountered. If ReadByte returns an error, no input
byte was consumed, and the returned byte value is undefined.


<pre>type ByteReader interface {
    ReadByte() (<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)
}</pre>











## <a id="ByteScanner">type</a> [ByteScanner](https://golang.org/src/io/io.go?s=8966:9028#L240)
ByteScanner is the interface that adds the UnreadByte method to the
basic ReadByte method.

UnreadByte causes the next call to ReadByte to return the same byte
as the previous call to ReadByte.
It may be an error to call UnreadByte twice without an intervening
call to ReadByte.


<pre>type ByteScanner interface {
    <a href="#ByteReader">ByteReader</a>
    UnreadByte() <a href="/pkg/builtin/#error">error</a>
}</pre>











## <a id="ByteWriter">type</a> [ByteWriter](https://golang.org/src/io/io.go?s=9094:9148#L246)
ByteWriter is the interface that wraps the WriteByte method.


<pre>type ByteWriter interface {
    WriteByte(c <a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a>
}</pre>











## <a id="Closer">type</a> [Closer](https://golang.org/src/io/io.go?s=4043:4083#L88)
Closer is the interface that wraps the basic Close method.

The behavior of Close after the first call is undefined.
Specific implementations may document their own behavior.


<pre>type Closer interface {
    Close() <a href="/pkg/builtin/#error">error</a>
}</pre>











## <a id="LimitedReader">type</a> [LimitedReader](https://golang.org/src/io/io.go?s=14790:14883#L426)
A LimitedReader reads from R but limits the amount of
data returned to just N bytes. Each call to Read
updates N to reflect the new amount remaining.
Read returns EOF when N <= 0 or when the underlying R returns EOF.


<pre>type LimitedReader struct {
<span id="LimitedReader.R"></span>    R <a href="#Reader">Reader</a> <span class="comment">// underlying reader</span>
<span id="LimitedReader.N"></span>    N <a href="/pkg/builtin/#int64">int64</a>  <span class="comment">// max bytes remaining</span>
}
</pre>











### <a id="LimitedReader.Read">func</a> (\*LimitedReader) [Read](https://golang.org/src/io/io.go?s=14885:14942#L431)
<pre>func (l *<a href="#LimitedReader">LimitedReader</a>) Read(p []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>



## <a id="PipeReader">type</a> [PipeReader](https://golang.org/src/io/pipe.go?s=2378:2413#L107)
A PipeReader is the read half of a pipe.


<pre>type PipeReader struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="PipeReader.Close">func</a> (\*PipeReader) [Close](https://golang.org/src/io/pipe.go?s=2861:2895#L122)
<pre>func (r *<a href="#PipeReader">PipeReader</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the reader; subsequent writes to the
write half of the pipe will return the error ErrClosedPipe.




### <a id="PipeReader.CloseWithError">func</a> (\*PipeReader) [CloseWithError](https://golang.org/src/io/pipe.go?s=3046:3098#L128)
<pre>func (r *<a href="#PipeReader">PipeReader</a>) CloseWithError(err <a href="/pkg/builtin/#error">error</a>) <a href="/pkg/builtin/#error">error</a></pre>
CloseWithError closes the reader; subsequent writes
to the write half of the pipe will return the error err.




### <a id="PipeReader.Read">func</a> (\*PipeReader) [Read](https://golang.org/src/io/pipe.go?s=2659:2716#L116)
<pre>func (r *<a href="#PipeReader">PipeReader</a>) Read(data []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Read implements the standard Read interface:
it reads data from the pipe, blocking until a writer
arrives or the write end is closed.
If the write end is closed with an error, that error is
returned as err; otherwise err is EOF.




## <a id="PipeWriter">type</a> [PipeWriter](https://golang.org/src/io/pipe.go?s=3176:3211#L133)
A PipeWriter is the write half of a pipe.


<pre>type PipeWriter struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="PipeWriter.Close">func</a> (\*PipeWriter) [Close](https://golang.org/src/io/pipe.go?s=3691:3725#L148)
<pre>func (w *<a href="#PipeWriter">PipeWriter</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the writer; subsequent reads from the
read half of the pipe will return no bytes and EOF.




### <a id="PipeWriter.CloseWithError">func</a> (\*PipeWriter) [CloseWithError](https://golang.org/src/io/pipe.go?s=3955:4007#L157)
<pre>func (w *<a href="#PipeWriter">PipeWriter</a>) CloseWithError(err <a href="/pkg/builtin/#error">error</a>) <a href="/pkg/builtin/#error">error</a></pre>
CloseWithError closes the writer; subsequent reads from the
read half of the pipe will return no bytes and the error err,
or EOF if err is nil.

CloseWithError always returns nil.




### <a id="PipeWriter.Write">func</a> (\*PipeWriter) [Write](https://golang.org/src/io/pipe.go?s=3494:3552#L142)
<pre>func (w *<a href="#PipeWriter">PipeWriter</a>) Write(data []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Write implements the standard Write interface:
it writes data to the pipe, blocking until one or more readers
have consumed all the data or the read end is closed.
If the read end is closed with an error, that err is
returned as err; otherwise err is ErrClosedPipe.




## <a id="ReadCloser">type</a> [ReadCloser](https://golang.org/src/io/io.go?s=4977:5022#L116)
ReadCloser is the interface that groups the basic Read and Close methods.


<pre>type ReadCloser interface {
    <a href="#Reader">Reader</a>
    <a href="#Closer">Closer</a>
}</pre>











## <a id="ReadSeeker">type</a> [ReadSeeker](https://golang.org/src/io/io.go?s=5376:5421#L135)
ReadSeeker is the interface that groups the basic Read and Seek methods.


<pre>type ReadSeeker interface {
    <a href="#Reader">Reader</a>
    <a href="#Seeker">Seeker</a>
}</pre>











## <a id="ReadWriteCloser">type</a> [ReadWriteCloser](https://golang.org/src/io/io.go?s=5240:5298#L128)
ReadWriteCloser is the interface that groups the basic Read, Write and Close methods.


<pre>type ReadWriteCloser interface {
    <a href="#Reader">Reader</a>
    <a href="#Writer">Writer</a>
    <a href="#Closer">Closer</a>
}</pre>











## <a id="ReadWriteSeeker">type</a> [ReadWriteSeeker](https://golang.org/src/io/io.go?s=5637:5695#L147)
ReadWriteSeeker is the interface that groups the basic Read, Write and Seek methods.


<pre>type ReadWriteSeeker interface {
    <a href="#Reader">Reader</a>
    <a href="#Writer">Writer</a>
    <a href="#Seeker">Seeker</a>
}</pre>











## <a id="ReadWriter">type</a> [ReadWriter](https://golang.org/src/io/io.go?s=4853:4898#L110)
ReadWriter is the interface that groups the basic Read and Write methods.


<pre>type ReadWriter interface {
    <a href="#Reader">Reader</a>
    <a href="#Writer">Writer</a>
}</pre>











## <a id="Reader">type</a> [Reader](https://golang.org/src/io/io.go?s=3303:3363#L67)
Reader is the interface that wraps the basic Read method.

Read reads up to len(p) bytes into p. It returns the number of bytes
read (0 <= n <= len(p)) and any error encountered. Even if Read
returns n < len(p), it may use all of p as scratch space during the call.
If some data is available but not len(p) bytes, Read conventionally
returns what is available instead of waiting for more.

When Read encounters an error or end-of-file condition after
successfully reading n > 0 bytes, it returns the number of
bytes read. It may return the (non-nil) error from the same call
or return the error (and n == 0) from a subsequent call.
An instance of this general case is that a Reader returning
a non-zero number of bytes at the end of the input stream may
return either err == EOF or err == nil. The next Read should
return 0, EOF.

Callers should always process the n > 0 bytes returned before
considering the error err. Doing so correctly handles I/O errors
that happen after reading some bytes and also both of the
allowed EOF behaviors.

Implementations of Read are discouraged from returning a
zero byte count with a nil error, except when len(p) == 0.
Callers should treat a return of 0 and nil as indicating that
nothing happened; in particular it does not indicate EOF.

Implementations must not retain p.


<pre>type Reader interface {
    Read(p []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)
}</pre>









### <a id="LimitReader">func</a> [LimitReader](https://golang.org/src/io/io.go?s=14485:14527#L420)
<pre>func LimitReader(r <a href="#Reader">Reader</a>, n <a href="/pkg/builtin/#int64">int64</a>) <a href="#Reader">Reader</a></pre>
LimitReader returns a Reader that reads from r
but stops with EOF after n bytes.
The underlying implementation is a *LimitedReader.


<a id="example_LimitReader">Example</a>
```go
```

output:
```txt
```


### <a id="MultiReader">func</a> [MultiReader](https://golang.org/src/io/multi.go?s=1269:1311#L38)
<pre>func MultiReader(readers ...<a href="#Reader">Reader</a>) <a href="#Reader">Reader</a></pre>
MultiReader returns a Reader that's the logical concatenation of
the provided input readers. They're read sequentially. Once all
inputs have returned EOF, Read will return EOF.  If any of the readers
return a non-nil, non-EOF error, Read will return that error.


<a id="example_MultiReader">Example</a>
```go
```

output:
```txt
```


### <a id="TeeReader">func</a> [TeeReader](https://golang.org/src/io/io.go?s=16906:16947#L515)
<pre>func TeeReader(r <a href="#Reader">Reader</a>, w <a href="#Writer">Writer</a>) <a href="#Reader">Reader</a></pre>
TeeReader returns a Reader that writes to w what it reads from r.
All reads from r performed through it are matched with
corresponding writes to w. There is no internal buffering -
the write must complete before the read completes.
Any error encountered while writing is reported as a read error.


<a id="example_TeeReader">Example</a>
```go
```

output:
```txt
```




## <a id="ReaderAt">type</a> [ReaderAt](https://golang.org/src/io/io.go?s=7545:7620#L201)
ReaderAt is the interface that wraps the basic ReadAt method.

ReadAt reads len(p) bytes into p starting at offset off in the
underlying input source. It returns the number of bytes
read (0 <= n <= len(p)) and any error encountered.

When ReadAt returns n < len(p), it returns a non-nil error
explaining why more bytes were not returned. In this respect,
ReadAt is stricter than Read.

Even if ReadAt returns n < len(p), it may use all of p as scratch
space during the call. If some data is available but not len(p) bytes,
ReadAt blocks until either all the data is available or an error occurs.
In this respect ReadAt is different from Read.

If the n = len(p) bytes returned by ReadAt are at the end of the
input source, ReadAt may return either err == EOF or err == nil.

If ReadAt is reading from an input source with a seek offset,
ReadAt should not affect nor be affected by the underlying
seek offset.

Clients of ReadAt can execute parallel ReadAt calls on the
same input source.

Implementations must not retain p.


<pre>type ReaderAt interface {
    ReadAt(p []<a href="/pkg/builtin/#byte">byte</a>, off <a href="/pkg/builtin/#int64">int64</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)
}</pre>











## <a id="ReaderFrom">type</a> [ReaderFrom](https://golang.org/src/io/io.go?s=5991:6061#L160)
ReaderFrom is the interface that wraps the ReadFrom method.

ReadFrom reads data from r until EOF or error.
The return value n is the number of bytes read.
Any error except io.EOF encountered during the read is also returned.

The Copy function uses ReaderFrom if available.


<pre>type ReaderFrom interface {
    ReadFrom(r <a href="#Reader">Reader</a>) (n <a href="/pkg/builtin/#int64">int64</a>, err <a href="/pkg/builtin/#error">error</a>)
}</pre>











## <a id="RuneReader">type</a> [RuneReader](https://golang.org/src/io/io.go?s=9372:9443#L255)
RuneReader is the interface that wraps the ReadRune method.

ReadRune reads a single UTF-8 encoded Unicode character
and returns the rune and its size in bytes. If no character is
available, err will be set.


<pre>type RuneReader interface {
    ReadRune() (r <a href="/pkg/builtin/#rune">rune</a>, size <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)
}</pre>











## <a id="RuneScanner">type</a> [RuneScanner](https://golang.org/src/io/io.go?s=9744:9806#L266)
RuneScanner is the interface that adds the UnreadRune method to the
basic ReadRune method.

UnreadRune causes the next call to ReadRune to return the same rune
as the previous call to ReadRune.
It may be an error to call UnreadRune twice without an intervening
call to ReadRune.


<pre>type RuneScanner interface {
    <a href="#RuneReader">RuneReader</a>
    UnreadRune() <a href="/pkg/builtin/#error">error</a>
}</pre>











## <a id="SectionReader">type</a> [SectionReader](https://golang.org/src/io/io.go?s=15408:15492#L451)
SectionReader implements Read, Seek, and ReadAt on a section
of an underlying ReaderAt.


<pre>type SectionReader struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_SectionReader">Example</a>
```go
```

output:
```txt
```




### <a id="NewSectionReader">func</a> [NewSectionReader](https://golang.org/src/io/io.go?s=15195:15263#L445)
<pre>func NewSectionReader(r <a href="#ReaderAt">ReaderAt</a>, off <a href="/pkg/builtin/#int64">int64</a>, n <a href="/pkg/builtin/#int64">int64</a>) *<a href="#SectionReader">SectionReader</a></pre>
NewSectionReader returns a SectionReader that reads from r
starting at offset off and stops with EOF after n bytes.






### <a id="SectionReader.Read">func</a> (\*SectionReader) [Read](https://golang.org/src/io/io.go?s=15494:15551#L458)
<pre>func (s *<a href="#SectionReader">SectionReader</a>) Read(p []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>



### <a id="SectionReader.ReadAt">func</a> (\*SectionReader) [ReadAt](https://golang.org/src/io/io.go?s=16155:16225#L491)
<pre>func (s *<a href="#SectionReader">SectionReader</a>) ReadAt(p []<a href="/pkg/builtin/#byte">byte</a>, off <a href="/pkg/builtin/#int64">int64</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>

<a id="example_SectionReader_ReadAt">Example</a>
```go
```

output:
```txt
```


### <a id="SectionReader.Seek">func</a> (\*SectionReader) [Seek](https://golang.org/src/io/io.go?s=15828:15897#L473)
<pre>func (s *<a href="#SectionReader">SectionReader</a>) Seek(offset <a href="/pkg/builtin/#int64">int64</a>, whence <a href="/pkg/builtin/#int">int</a>) (<a href="/pkg/builtin/#int64">int64</a>, <a href="/pkg/builtin/#error">error</a>)</pre>

<a id="example_SectionReader_Seek">Example</a>
```go
```

output:
```txt
```


### <a id="SectionReader.Size">func</a> (\*SectionReader) [Size](https://golang.org/src/io/io.go?s=16528:16564#L508)
<pre>func (s *<a href="#SectionReader">SectionReader</a>) Size() <a href="/pkg/builtin/#int64">int64</a></pre>
Size returns the size of the section in bytes.




## <a id="Seeker">type</a> [Seeker](https://golang.org/src/io/io.go?s=4702:4774#L105)
Seeker is the interface that wraps the basic Seek method.

Seek sets the offset for the next Read or Write to offset,
interpreted according to whence:
SeekStart means relative to the start of the file,
SeekCurrent means relative to the current offset, and
SeekEnd means relative to the end.
Seek returns the new offset relative to the start of the
file and an error, if any.

Seeking to an offset before the start of the file is an error.
Seeking to any positive offset is legal, but the behavior of subsequent
I/O operations on the underlying object is implementation-dependent.


<pre>type Seeker interface {
    Seek(offset <a href="/pkg/builtin/#int64">int64</a>, whence <a href="/pkg/builtin/#int">int</a>) (<a href="/pkg/builtin/#int64">int64</a>, <a href="/pkg/builtin/#error">error</a>)
}</pre>











## <a id="StringWriter">type</a> [StringWriter](https://golang.org/src/io/io.go?s=9876:9949#L272)
StringWriter is the interface that wraps the WriteString method.


<pre>type StringWriter interface {
    WriteString(s <a href="/pkg/builtin/#string">string</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)
}</pre>











## <a id="WriteCloser">type</a> [WriteCloser](https://golang.org/src/io/io.go?s=5103:5149#L122)
WriteCloser is the interface that groups the basic Write and Close methods.


<pre>type WriteCloser interface {
    <a href="#Writer">Writer</a>
    <a href="#Closer">Closer</a>
}</pre>











## <a id="WriteSeeker">type</a> [WriteSeeker](https://golang.org/src/io/io.go?s=5501:5547#L141)
WriteSeeker is the interface that groups the basic Write and Seek methods.


<pre>type WriteSeeker interface {
    <a href="#Writer">Writer</a>
    <a href="#Seeker">Seeker</a>
}</pre>











## <a id="Writer">type</a> [Writer](https://golang.org/src/io/io.go?s=3794:3855#L80)
Writer is the interface that wraps the basic Write method.

Write writes len(p) bytes from p to the underlying data stream.
It returns the number of bytes written from p (0 <= n <= len(p))
and any error encountered that caused the write to stop early.
Write must return a non-nil error if it returns n < len(p).
Write must not modify the slice data, even temporarily.

Implementations must not retain p.


<pre>type Writer interface {
    Write(p []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)
}</pre>









### <a id="MultiWriter">func</a> [MultiWriter](https://golang.org/src/io/multi.go?s=2446:2488#L92)
<pre>func MultiWriter(writers ...<a href="#Writer">Writer</a>) <a href="#Writer">Writer</a></pre>
MultiWriter creates a writer that duplicates its writes to all the
provided writers, similar to the Unix tee(1) command.

Each write is written to each listed writer, one at a time.
If a listed writer returns an error, that overall write operation
stops and returns the error; it does not continue down the list.


<a id="example_MultiWriter">Example</a>
```go
```

output:
```txt
```




## <a id="WriterAt">type</a> [WriterAt](https://golang.org/src/io/io.go?s=8275:8351#L220)
WriterAt is the interface that wraps the basic WriteAt method.

WriteAt writes len(p) bytes from p to the underlying data stream
at offset off. It returns the number of bytes written from p (0 <= n <= len(p))
and any error encountered that caused the write to stop early.
WriteAt must return a non-nil error if it returns n < len(p).

If WriteAt is writing to a destination with a seek offset,
WriteAt should not affect nor be affected by the underlying
seek offset.

Clients of WriteAt can execute parallel WriteAt calls on the same
destination if the ranges do not overlap.

Implementations must not retain p.


<pre>type WriterAt interface {
    WriteAt(p []<a href="/pkg/builtin/#byte">byte</a>, off <a href="/pkg/builtin/#int64">int64</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)
}</pre>











## <a id="WriterTo">type</a> [WriterTo](https://golang.org/src/io/io.go?s=6381:6448#L171)
WriterTo is the interface that wraps the WriteTo method.

WriteTo writes data to w until there's no more data to write or
when an error occurs. The return value n is the number of bytes
written. Any error encountered during the write is also returned.

The Copy function uses WriterTo if available.


<pre>type WriterTo interface {
    WriteTo(w <a href="#Writer">Writer</a>) (n <a href="/pkg/builtin/#int64">int64</a>, err <a href="/pkg/builtin/#error">error</a>)
}</pre>














