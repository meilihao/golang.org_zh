

# zip
`import "archive/zip"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package zip provides support for reading and writing ZIP archives.

See: <a href="https://www.pkware.com/appnote">https://www.pkware.com/appnote</a>

This package does not support disk spanning.

A note about ZIP64:

To be backwards compatible the FileHeader has both 32 and 64 bit Size
fields. The 64 bit fields will always contain the correct value and
for normal archives both fields will be the same. For files requiring
the ZIP64 format the 32 bit fields will be 0xffffffff and the 64 bit
fields must be used instead.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func RegisterCompressor(method uint16, comp Compressor)](#RegisterCompressor)
* [func RegisterDecompressor(method uint16, dcomp Decompressor)](#RegisterDecompressor)
* [type Compressor](#Compressor)
* [type Decompressor](#Decompressor)
* [type File](#File)
  * [func (f *File) DataOffset() (offset int64, err error)](#File.DataOffset)
  * [func (f *File) Open() (io.ReadCloser, error)](#File.Open)
* [type FileHeader](#FileHeader)
  * [func FileInfoHeader(fi os.FileInfo) (*FileHeader, error)](#FileInfoHeader)
  * [func (h *FileHeader) FileInfo() os.FileInfo](#FileHeader.FileInfo)
  * [func (h *FileHeader) ModTime() time.Time](#FileHeader.ModTime)
  * [func (h *FileHeader) Mode() (mode os.FileMode)](#FileHeader.Mode)
  * [func (h *FileHeader) SetModTime(t time.Time)](#FileHeader.SetModTime)
  * [func (h *FileHeader) SetMode(mode os.FileMode)](#FileHeader.SetMode)
* [type ReadCloser](#ReadCloser)
  * [func OpenReader(name string) (*ReadCloser, error)](#OpenReader)
  * [func (rc *ReadCloser) Close() error](#ReadCloser.Close)
* [type Reader](#Reader)
  * [func NewReader(r io.ReaderAt, size int64) (*Reader, error)](#NewReader)
  * [func (z *Reader) RegisterDecompressor(method uint16, dcomp Decompressor)](#Reader.RegisterDecompressor)
* [type Writer](#Writer)
  * [func NewWriter(w io.Writer) *Writer](#NewWriter)
  * [func (w *Writer) Close() error](#Writer.Close)
  * [func (w *Writer) Create(name string) (io.Writer, error)](#Writer.Create)
  * [func (w *Writer) CreateHeader(fh *FileHeader) (io.Writer, error)](#Writer.CreateHeader)
  * [func (w *Writer) Flush() error](#Writer.Flush)
  * [func (w *Writer) RegisterCompressor(method uint16, comp Compressor)](#Writer.RegisterCompressor)
  * [func (w *Writer) SetComment(comment string) error](#Writer.SetComment)
  * [func (w *Writer) SetOffset(n int64)](#Writer.SetOffset)


#### <a id="pkg-examples">Examples</a>
* [Reader](#example_Reader)
* [Writer](#example_Writer)
* [Writer.RegisterCompressor](#example_Writer_RegisterCompressor)


#### <a id="pkg-files">Package files</a>
[reader.go](https://golang.org/src/archive/zip/reader.go) [register.go](https://golang.org/src/archive/zip/register.go) [struct.go](https://golang.org/src/archive/zip/struct.go) [writer.go](https://golang.org/src/archive/zip/writer.go) 


## <a id="pkg-constants">Constants</a>
Compression methods.


<pre>const (
    <span id="Store">Store</span>   <a href="/pkg/builtin/#uint16">uint16</a> = 0 <span class="comment">// no compression</span>
    <span id="Deflate">Deflate</span> <a href="/pkg/builtin/#uint16">uint16</a> = 8 <span class="comment">// DEFLATE compressed</span>
)</pre>

## <a id="pkg-variables">Variables</a>

<pre>var (
    <span id="ErrFormat">ErrFormat</span>    = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;zip: not a valid zip file&#34;)
    <span id="ErrAlgorithm">ErrAlgorithm</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;zip: unsupported compression algorithm&#34;)
    <span id="ErrChecksum">ErrChecksum</span>  = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;zip: checksum error&#34;)
)</pre>

## <a id="RegisterCompressor">func</a> [RegisterCompressor](https://golang.org/src/archive/zip/register.go?s=3317:3372#L118)
<pre>func RegisterCompressor(method <a href="/pkg/builtin/#uint16">uint16</a>, comp <a href="#Compressor">Compressor</a>)</pre>
RegisterCompressor registers custom compressors for a specified method ID.
The common methods Store and Deflate are built in.



## <a id="RegisterDecompressor">func</a> [RegisterDecompressor](https://golang.org/src/archive/zip/register.go?s=3011:3071#L110)
<pre>func RegisterDecompressor(method <a href="/pkg/builtin/#uint16">uint16</a>, dcomp <a href="#Decompressor">Decompressor</a>)</pre>
RegisterDecompressor allows custom decompressors for a specified method ID.
The common methods Store and Deflate are built in.





## <a id="Compressor">type</a> [Compressor](https://golang.org/src/archive/zip/register.go?s=545:602#L10)
A Compressor returns a new compressing writer, writing to w.
The WriteCloser's Close method must be used to flush pending data to w.
The Compressor itself must be safe to invoke from multiple goroutines
simultaneously, but each returned writer will be used only by
one goroutine at a time.


<pre>type Compressor func(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) (<a href="/pkg/io/">io</a>.<a href="/pkg/io/#WriteCloser">WriteCloser</a>, <a href="/pkg/builtin/#error">error</a>)</pre>











## <a id="Decompressor">type</a> [Decompressor](https://golang.org/src/archive/zip/register.go?s=921:970#L17)
A Decompressor returns a new decompressing reader, reading from r.
The ReadCloser's Close method must be used to release associated resources.
The Decompressor itself must be safe to invoke from multiple goroutines
simultaneously, but each returned reader will be used only by
one goroutine at a time.


<pre>type Decompressor func(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadCloser">ReadCloser</a></pre>











## <a id="File">type</a> [File](https://golang.org/src/archive/zip/reader.go?s=639:759#L27)

<pre>type File struct {
    <a href="#FileHeader">FileHeader</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="File.DataOffset">func</a> (\*File) [DataOffset](https://golang.org/src/archive/zip/reader.go?s=3722:3775#L140)
<pre>func (f *<a href="#File">File</a>) DataOffset() (offset <a href="/pkg/builtin/#int64">int64</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
DataOffset returns the offset of the file's possibly-compressed
data, relative to the beginning of the zip file.

Most callers should instead use Open, which transparently
decompresses data and verifies checksums.




### <a id="File.Open">func</a> (\*File) [Open](https://golang.org/src/archive/zip/reader.go?s=4008:4052#L150)
<pre>func (f *<a href="#File">File</a>) Open() (<a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadCloser">ReadCloser</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Open returns a ReadCloser that provides access to the File's contents.
Multiple files may be read concurrently.




## <a id="FileHeader">type</a> [FileHeader](https://golang.org/src/archive/zip/struct.go?s=2650:5022#L72)
FileHeader describes a file within a zip file.
See the zip spec for details.


<pre>type FileHeader struct {
<span id="FileHeader.Name"></span>    <span class="comment">// Name is the name of the file.</span>
    <span class="comment">//</span>
    <span class="comment">// It must be a relative path, not start with a drive letter (such as &#34;C:&#34;),</span>
    <span class="comment">// and must use forward slashes instead of back slashes. A trailing slash</span>
    <span class="comment">// indicates that this file is a directory and should have no data.</span>
    <span class="comment">//</span>
    <span class="comment">// When reading zip files, the Name field is populated from</span>
    <span class="comment">// the zip file directly and is not validated for correctness.</span>
    <span class="comment">// It is the caller&#39;s responsibility to sanitize it as</span>
    <span class="comment">// appropriate, including canonicalizing slash directions,</span>
    <span class="comment">// validating that paths are relative, and preventing path</span>
    <span class="comment">// traversal through filenames (&#34;../../../&#34;).</span>
    Name <a href="/pkg/builtin/#string">string</a>

<span id="FileHeader.Comment"></span>    <span class="comment">// Comment is any arbitrary user-defined string shorter than 64KiB.</span>
    Comment <a href="/pkg/builtin/#string">string</a>

<span id="FileHeader.NonUTF8"></span>    <span class="comment">// NonUTF8 indicates that Name and Comment are not encoded in UTF-8.</span>
    <span class="comment">//</span>
    <span class="comment">// By specification, the only other encoding permitted should be CP-437,</span>
    <span class="comment">// but historically many ZIP readers interpret Name and Comment as whatever</span>
    <span class="comment">// the system&#39;s local character encoding happens to be.</span>
    <span class="comment">//</span>
    <span class="comment">// This flag should only be set if the user intends to encode a non-portable</span>
    <span class="comment">// ZIP file for a specific localized region. Otherwise, the Writer</span>
    <span class="comment">// automatically sets the ZIP format&#39;s UTF-8 flag for valid UTF-8 strings.</span>
    NonUTF8 <a href="/pkg/builtin/#bool">bool</a>

<span id="FileHeader.CreatorVersion"></span>    CreatorVersion <a href="/pkg/builtin/#uint16">uint16</a>
<span id="FileHeader.ReaderVersion"></span>    ReaderVersion  <a href="/pkg/builtin/#uint16">uint16</a>
<span id="FileHeader.Flags"></span>    Flags          <a href="/pkg/builtin/#uint16">uint16</a>

<span id="FileHeader.Method"></span>    <span class="comment">// Method is the compression method. If zero, Store is used.</span>
    Method <a href="/pkg/builtin/#uint16">uint16</a>

<span id="FileHeader.Modified"></span>    <span class="comment">// Modified is the modified time of the file.</span>
    <span class="comment">//</span>
    <span class="comment">// When reading, an extended timestamp is preferred over the legacy MS-DOS</span>
    <span class="comment">// date field, and the offset between the times is used as the timezone.</span>
    <span class="comment">// If only the MS-DOS date is present, the timezone is assumed to be UTC.</span>
    <span class="comment">//</span>
    <span class="comment">// When writing, an extended timestamp (which is timezone-agnostic) is</span>
    <span class="comment">// always emitted. The legacy MS-DOS date field is encoded according to the</span>
    <span class="comment">// location of the Modified time.</span>
    Modified     <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>
<span id="FileHeader.ModifiedTime"></span>    ModifiedTime <a href="/pkg/builtin/#uint16">uint16</a> <span class="comment">// Deprecated: Legacy MS-DOS date; use Modified instead.</span>
<span id="FileHeader.ModifiedDate"></span>    ModifiedDate <a href="/pkg/builtin/#uint16">uint16</a> <span class="comment">// Deprecated: Legacy MS-DOS time; use Modified instead.</span>

<span id="FileHeader.CRC32"></span>    CRC32              <a href="/pkg/builtin/#uint32">uint32</a>
<span id="FileHeader.CompressedSize"></span>    CompressedSize     <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">// Deprecated: Use CompressedSize64 instead.</span>
<span id="FileHeader.UncompressedSize"></span>    UncompressedSize   <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">// Deprecated: Use UncompressedSize64 instead.</span>
<span id="FileHeader.CompressedSize64"></span>    CompressedSize64   <a href="/pkg/builtin/#uint64">uint64</a>
<span id="FileHeader.UncompressedSize64"></span>    UncompressedSize64 <a href="/pkg/builtin/#uint64">uint64</a>
<span id="FileHeader.Extra"></span>    Extra              []<a href="/pkg/builtin/#byte">byte</a>
<span id="FileHeader.ExternalAttrs"></span>    ExternalAttrs      <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">// Meaning depends on CreatorVersion</span>
}
</pre>









### <a id="FileInfoHeader">func</a> [FileInfoHeader](https://golang.org/src/archive/zip/struct.go?s=6203:6259#L164)
<pre>func FileInfoHeader(fi <a href="/pkg/os/">os</a>.<a href="/pkg/os/#FileInfo">FileInfo</a>) (*<a href="#FileHeader">FileHeader</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
FileInfoHeader creates a partially-populated FileHeader from an
os.FileInfo.
Because os.FileInfo's Name method returns only the base name of
the file it describes, it may be necessary to modify the Name field
of the returned header to provide the full path name of the file.
If compression is desired, callers should set the FileHeader.Method
field; it is unset by default.






### <a id="FileHeader.FileInfo">func</a> (\*FileHeader) [FileInfo](https://golang.org/src/archive/zip/struct.go?s=5079:5122#L131)
<pre>func (h *<a href="#FileHeader">FileHeader</a>) FileInfo() <a href="/pkg/os/">os</a>.<a href="/pkg/os/#FileInfo">FileInfo</a></pre>
FileInfo returns an os.FileInfo for the FileHeader.




### <a id="FileHeader.ModTime">func</a> (\*FileHeader) [ModTime](https://golang.org/src/archive/zip/struct.go?s=8486:8526#L239)
<pre>func (h *<a href="#FileHeader">FileHeader</a>) ModTime() <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a></pre>
ModTime returns the modification time in UTC using the legacy
ModifiedDate and ModifiedTime fields.

Deprecated: Use Modified instead.




### <a id="FileHeader.Mode">func</a> (\*FileHeader) [Mode](https://golang.org/src/archive/zip/struct.go?s=9339:9385#L273)
<pre>func (h *<a href="#FileHeader">FileHeader</a>) Mode() (mode <a href="/pkg/os/">os</a>.<a href="/pkg/os/#FileMode">FileMode</a>)</pre>
Mode returns the permission and mode bits for the FileHeader.




### <a id="FileHeader.SetModTime">func</a> (\*FileHeader) [SetModTime](https://golang.org/src/archive/zip/struct.go?s=8728:8772#L247)
<pre>func (h *<a href="#FileHeader">FileHeader</a>) SetModTime(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>)</pre>
SetModTime sets the Modified, ModifiedTime, and ModifiedDate fields
to the given time in UTC.

Deprecated: Use Modified instead.




### <a id="FileHeader.SetMode">func</a> (\*FileHeader) [SetMode](https://golang.org/src/archive/zip/struct.go?s=9760:9806#L287)
<pre>func (h *<a href="#FileHeader">FileHeader</a>) SetMode(mode <a href="/pkg/os/">os</a>.<a href="/pkg/os/#FileMode">FileMode</a>)</pre>
SetMode changes the permission and mode bits for the FileHeader.




## <a id="ReadCloser">type</a> [ReadCloser](https://golang.org/src/archive/zip/reader.go?s=591:637#L22)

<pre>type ReadCloser struct {
    <a href="#Reader">Reader</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="OpenReader">func</a> [OpenReader](https://golang.org/src/archive/zip/reader.go?s=911:960#L40)
<pre>func OpenReader(name <a href="/pkg/builtin/#string">string</a>) (*<a href="#ReadCloser">ReadCloser</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
OpenReader will open the Zip file specified by name and return a ReadCloser.






### <a id="ReadCloser.Close">func</a> (\*ReadCloser) [Close](https://golang.org/src/archive/zip/reader.go?s=3432:3467#L131)
<pre>func (rc *<a href="#ReadCloser">ReadCloser</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the Zip file, rendering it unusable for I/O.




## <a id="Reader">type</a> [Reader](https://golang.org/src/archive/zip/reader.go?s=456:589#L15)

<pre>type Reader struct {
<span id="Reader.File"></span>    File    []*<a href="#File">File</a>
<span id="Reader.Comment"></span>    Comment <a href="/pkg/builtin/#string">string</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Reader">Example</a>
```go
```

output:
```txt
```




### <a id="NewReader">func</a> [NewReader](https://golang.org/src/archive/zip/reader.go?s=1328:1386#L61)
<pre>func NewReader(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReaderAt">ReaderAt</a>, size <a href="/pkg/builtin/#int64">int64</a>) (*<a href="#Reader">Reader</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewReader returns a new Reader reading from r, which is assumed to
have the given size in bytes.






### <a id="Reader.RegisterDecompressor">func</a> (\*Reader) [RegisterDecompressor](https://golang.org/src/archive/zip/reader.go?s=3014:3086#L115)
<pre>func (z *<a href="#Reader">Reader</a>) RegisterDecompressor(method <a href="/pkg/builtin/#uint16">uint16</a>, dcomp <a href="#Decompressor">Decompressor</a>)</pre>
RegisterDecompressor registers or overrides a custom decompressor for a
specific method ID. If a decompressor for a given method is not found,
Reader will default to looking up the decompressor at the package level.




## <a id="Writer">type</a> [Writer](https://golang.org/src/archive/zip/writer.go?s=448:781#L14)
Writer implements a zip file writer.


<pre>type Writer struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Writer">Example</a>
```go
```

output:
```txt
```




### <a id="NewWriter">func</a> [NewWriter](https://golang.org/src/archive/zip/writer.go?s=894:929#L33)
<pre>func NewWriter(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) *<a href="#Writer">Writer</a></pre>
NewWriter returns a new Writer writing a zip file to w.






### <a id="Writer.Close">func</a> (\*Writer) [Close](https://golang.org/src/archive/zip/writer.go?s=1977:2007#L66)
<pre>func (w *<a href="#Writer">Writer</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close finishes writing the zip file by writing the central directory.
It does not close the underlying writer.




### <a id="Writer.Create">func</a> (\*Writer) [Create](https://golang.org/src/archive/zip/writer.go?s=6761:6816#L207)
<pre>func (w *<a href="#Writer">Writer</a>) Create(name <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Create adds a file to the zip file using the provided name.
It returns a Writer to which the file contents should be written.
The file contents will be compressed using the Deflate method.
The name must be a relative path: it must not start with a drive
letter (e.g. C:) or leading slash, and only forward slashes are
allowed. To create a directory instead of a file, add a trailing
slash to the name.
The file's contents must be written to the io.Writer before the next
call to Create, CreateHeader, or Close.




### <a id="Writer.CreateHeader">func</a> (\*Writer) [CreateHeader](https://golang.org/src/archive/zip/writer.go?s=8164:8228#L245)
<pre>func (w *<a href="#Writer">Writer</a>) CreateHeader(fh *<a href="#FileHeader">FileHeader</a>) (<a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
CreateHeader adds a file to the zip archive using the provided FileHeader
for the file metadata. Writer takes ownership of fh and may mutate
its fields. The caller must not modify fh after calling CreateHeader.

This returns a Writer to which the file contents should be written.
The file's contents must be written to the io.Writer before the next
call to Create, CreateHeader, or Close.




### <a id="Writer.Flush">func</a> (\*Writer) [Flush](https://golang.org/src/archive/zip/writer.go?s=1509:1539#L50)
<pre>func (w *<a href="#Writer">Writer</a>) Flush() <a href="/pkg/builtin/#error">error</a></pre>
Flush flushes any buffered data to the underlying writer.
Calling Flush is not normally necessary; calling Close is sufficient.




### <a id="Writer.RegisterCompressor">func</a> (\*Writer) [RegisterCompressor](https://golang.org/src/archive/zip/writer.go?s=13353:13420#L401)
<pre>func (w *<a href="#Writer">Writer</a>) RegisterCompressor(method <a href="/pkg/builtin/#uint16">uint16</a>, comp <a href="#Compressor">Compressor</a>)</pre>
RegisterCompressor registers or overrides a custom compressor for a specific
method ID. If a compressor for a given method is not found, Writer will
default to looking up the compressor at the package level.


<a id="example_Writer_RegisterCompressor">Example</a>
```go
```

output:
```txt
```


### <a id="Writer.SetComment">func</a> (\*Writer) [SetComment](https://golang.org/src/archive/zip/writer.go?s=1686:1735#L56)
<pre>func (w *<a href="#Writer">Writer</a>) SetComment(comment <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetComment sets the end-of-central-directory comment field.
It can only be called before Close.




### <a id="Writer.SetOffset">func</a> (\*Writer) [SetOffset](https://golang.org/src/archive/zip/writer.go?s=1237:1272#L41)
<pre>func (w *<a href="#Writer">Writer</a>) SetOffset(n <a href="/pkg/builtin/#int64">int64</a>)</pre>
SetOffset sets the offset of the beginning of the zip data within the
underlying writer. It should be used when the zip data is appended to an
existing file, such as a binary executable.
It must be called before any data is written.







