

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

zip 提供了对zip格式文件的读写支持.

参见: <a href="https://www.pkware.com/appnote">https://www.pkware.com/appnote</a>

本包不支持disk spanning.

关于ZIP64的注意事项:

为了向下兼容, FileHeader同时拥有32位和64位的Size字段. 64位字段将始终包含正确的值,且对于普通文档,这两个字段的值都是相同的. 但对于使用ZIP64格式的文件,其32位字段的值为0xffffffff,且必须使用64位字段.


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

压缩方式.


<pre>const (
    <span id="Store">Store</span>   <a href="/pkg/builtin/#uint16">uint16</a> = 0 <span class="comment">// no compression // 不压缩</span>
    <span id="Deflate">Deflate</span> <a href="/pkg/builtin/#uint16">uint16</a> = 8 <span class="comment">// DEFLATE compressed // 压缩</span>
)</pre>

## <a id="pkg-variables">Variables</a>

```go
var (
    ErrFormat    = errors.New("zip: not a valid zip file") // 无效的zip文件
    ErrAlgorithm = errors.New("zip: unsupported compression algorithm") // 不支持的压缩算法
    ErrChecksum  = errors.New("zip: checksum error") // 校验和错误
)
```

## <a id="RegisterCompressor">func</a> [RegisterCompressor](https://golang.org/src/archive/zip/register.go?s=3317:3372#L118)
<pre>func RegisterCompressor(method <a href="/pkg/builtin/#uint16">uint16</a>, comp <a href="#Compressor">Compressor</a>)</pre>
RegisterCompressor registers custom compressors for a specified method ID.
The common methods Store and Deflate are built in.

RegisterCompressor 使用指定的方法ID注册自定义压缩器. 已内置常用的方法Store和Deflate.


## <a id="RegisterDecompressor">func</a> [RegisterDecompressor](https://golang.org/src/archive/zip/register.go?s=3011:3071#L110)
<pre>func RegisterDecompressor(method <a href="/pkg/builtin/#uint16">uint16</a>, dcomp <a href="#Decompressor">Decompressor</a>)</pre>
RegisterDecompressor allows custom decompressors for a specified method ID.
The common methods Store and Deflate are built in.


RegisterDecompressor 使用指定的方法ID注册自定义解压器. 已内置常用的方法Store和Deflate.


## <a id="Compressor">type</a> [Compressor](https://golang.org/src/archive/zip/register.go?s=545:602#L10)
A Compressor returns a new compressing writer, writing to w.
The WriteCloser's Close method must be used to flush pending data to w.
The Compressor itself must be safe to invoke from multiple goroutines
simultaneously, but each returned writer will be used only by
one goroutine at a time.

Compressor 返回一个会将数据写入w的压缩处理器. 它必须使用WriteCloser的Close方法将剩余数据写入w. Compressor本身必须能同时地被多个goroutine调用, 但每次返回的writer一次只能由一个goroutine调用.

<pre>type Compressor func(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) (<a href="/pkg/io/">io</a>.<a href="/pkg/io/#WriteCloser">WriteCloser</a>, <a href="/pkg/builtin/#error">error</a>)</pre>











## <a id="Decompressor">type</a> [Decompressor](https://golang.org/src/archive/zip/register.go?s=921:970#L17)
A Decompressor returns a new decompressing reader, reading from r.
The ReadCloser's Close method must be used to release associated resources.
The Decompressor itself must be safe to invoke from multiple goroutines
simultaneously, but each returned reader will be used only by
one goroutine at a time.

Decompressor 返回一个从r中读取数据的解压处理器. 它必须使用ReadCloser的Close方法来释放相关的资源. Decompressor本身必须能同时地被多个goroutine调用, 但每次返回的reader一次只能由一个goroutine调用.

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


DataOffset 返回相对于zip文档起始位置,该文件(数据可能被压缩了)的偏移量.

大多数调用者应使用Open来代替(DataOffset),它会透明地解压缩数据并验证校验和.



### <a id="File.Open">func</a> (\*File) [Open](https://golang.org/src/archive/zip/reader.go?s=4008:4052#L150)
<pre>func (f *<a href="#File">File</a>) Open() (<a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadCloser">ReadCloser</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Open returns a ReadCloser that provides access to the File's contents.
Multiple files may be read concurrently.

Open方法返回一个io.ReadCloser接口, 提供了读取文件内容的方法. 多个文件可以并发读取.


## <a id="FileHeader">type</a> [FileHeader](https://golang.org/src/archive/zip/struct.go?s=2650:5022#L72)
FileHeader describes a file within a zip file.
See the zip spec for details.

FileHeader描述zip文档中的一个文件. 细节请见zip的定义.


```go
type FileHeader struct {
    // Name is the name of the file.
    //
    // It must be a relative path, not start with a drive letter (such as "C:"),
    // and must use forward slashes instead of back slashes. A trailing slash
    // indicates that this file is a directory and should have no data.
    //
    // When reading zip files, the Name field is populated from
    // the zip file directly and is not validated for correctness.
    // It is the caller's responsibility to sanitize it as
    // appropriate, including canonicalizing slash directions,
    // validating that paths are relative, and preventing path
    // traversal through filenames ("../../../").
    //
    // Name 是文件的名称
    //
    // 它必须是相对路径, 不能以盘符(比如"C:")开头, 并且必须使用正斜杠而不是反斜杠. 名称末尾的斜杠表示此文件是目录, 不应包含任何数据.
    //
    // 读取zip文件时, 能从zip文件中直接获取到Name, 但还未验证是否正确. 调用者有责任进行适当地处理, 包括规范化的斜线方向, 验证路径是相对的, 并防止通过路径
    // 遍历文件名("../../../").
    Name string

    // Comment is any arbitrary user-defined string shorter than 64KiB.
    // Comment 是 用户自定义的短于64KB的任意字符串.
    Comment string

    // NonUTF8 indicates that Name and Comment are not encoded in UTF-8.
    //
    // By specification, the only other encoding permitted should be CP-437,
    // but historically many ZIP readers interpret Name and Comment as whatever
    // the system's local character encoding happens to be.
    //
    // This flag should only be set if the user intends to encode a non-portable
    // ZIP file for a specific localized region. Otherwise, the Writer
    // automatically sets the ZIP format's UTF-8 flag for valid UTF-8 strings.
    //
    // NonUTF8 表示 Name和Comment未使用UTF-8编码.
    //
    // 根据规范，唯一允许的其他编码应为CP-437，但历史上许多ZIP reader将 Name 和 Comment 解释为系统的本地字符编码.
    //
    // 仅当用户打算为特定的本地区域编码不可移植的zip文件时，才应设置此标志. 否则，将自动为有效的UTF-8字符串设置ZIP格式的UTF-8 flag.
    NonUTF8 bool // Go 1.10

    CreatorVersion uint16
    ReaderVersion  uint16
    Flags          uint16

    // Method is the compression method. If zero, Store is used.
    // Method 是 压缩方式. 如果是零值就使用 Store.
    Method uint16

    // Modified is the modified time of the file.
    //
    // When reading, an extended timestamp is preferred over the legacy MS-DOS
    // date field, and the offset between the times is used as the timezone.
    // If only the MS-DOS date is present, the timezone is assumed to be UTC.
    //
    // When writing, an extended timestamp (which is timezone-agnostic) is
    // always emitted. The legacy MS-DOS date field is encoded according to the
    // location of the Modified time.
    //
    // Modified 是文件的修改时间.
    //
    // 读取时，扩展时间戳优先于过时的MS-DOS日期字段，并且时间之间的偏移量作为时区. 如果仅存在MS-DOS日期，则将时区假定为UTC.
    //
    // 写入时，将使用扩展的时间戳（与时区无关）. 过时的MS-DOS日期字段将根据Modified进行设置.
    Modified     time.Time // Go 1.10
    ModifiedTime uint16 // Deprecated: Legacy MS-DOS date; use Modified instead. // 已弃用: 过时的MS-DOS日期, 用 Modified 代替.
    ModifiedDate uint16 // Deprecated: Legacy MS-DOS time; use Modified instead. // 已弃用: 过时的MS-DOS时间, 用 Modified 代替.

    CRC32              uint32
    CompressedSize     uint32 // Deprecated: Use CompressedSize64 instead. // 已弃用: 请使用CompressedSize64
    UncompressedSize   uint32 // Deprecated: Use UncompressedSize64 instead. // 已弃用: 请使用UncompressedSize64
    CompressedSize64   uint64 // Go 1.1
    UncompressedSize64 uint64 // Go 1.1
    Extra              []byte
    ExternalAttrs      uint32 // Meaning depends on CreatorVersion // 其含义依赖于CreatorVersion
}
```








### <a id="FileInfoHeader">func</a> [FileInfoHeader](https://golang.org/src/archive/zip/struct.go?s=6203:6259#L164)
<pre>func FileInfoHeader(fi <a href="/pkg/os/">os</a>.<a href="/pkg/os/#FileInfo">FileInfo</a>) (*<a href="#FileHeader">FileHeader</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
FileInfoHeader creates a partially-populated FileHeader from an
os.FileInfo.
Because os.FileInfo's Name method returns only the base name of
the file it describes, it may be necessary to modify the Name field
of the returned header to provide the full path name of the file.
If compression is desired, callers should set the FileHeader.Method
field; it is unset by default.

FileInfoHeader 返回一个基于fi填充了部分字段的Header. 因为os.FileInfo的Name方法返回的文件名不带路径,有可能需要将Header的Name字段改为文件的完整路径名. 如果需要压缩，则调用者应设置FileHeader.Method的字段, 默认情况下未设置.




### <a id="FileHeader.FileInfo">func</a> (\*FileHeader) [FileInfo](https://golang.org/src/archive/zip/struct.go?s=5079:5122#L131)
<pre>func (h *<a href="#FileHeader">FileHeader</a>) FileInfo() <a href="/pkg/os/">os</a>.<a href="/pkg/os/#FileInfo">FileInfo</a></pre>
FileInfo returns an os.FileInfo for the FileHeader.

FileInfo 为 FileHeader 返回一个os.FileInfo.


### <a id="FileHeader.ModTime">func</a> (\*FileHeader) [ModTime](https://golang.org/src/archive/zip/struct.go?s=8486:8526#L239)
<pre>func (h *<a href="#FileHeader">FileHeader</a>) ModTime() <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a></pre>
ModTime returns the modification time in UTC using the legacy
ModifiedDate and ModifiedTime fields.

Deprecated: Use Modified instead.

ModTime 使用过时的ModifiedDate和ModifiedTime字段以返回UTC格式的修改时间.

不推荐使用：改用Modified.


### <a id="FileHeader.Mode">func</a> (\*FileHeader) [Mode](https://golang.org/src/archive/zip/struct.go?s=9339:9385#L273)
<pre>func (h *<a href="#FileHeader">FileHeader</a>) Mode() (mode <a href="/pkg/os/">os</a>.<a href="/pkg/os/#FileMode">FileMode</a>)</pre>
Mode returns the permission and mode bits for the FileHeader.

Mode 返回h的权限和模式位.


### <a id="FileHeader.SetModTime">func</a> (\*FileHeader) [SetModTime](https://golang.org/src/archive/zip/struct.go?s=8728:8772#L247)
<pre>func (h *<a href="#FileHeader">FileHeader</a>) SetModTime(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>)</pre>
SetModTime sets the Modified, ModifiedTime, and ModifiedDate fields
to the given time in UTC.

Deprecated: Use Modified instead.

SetModTime 使用给定的UTC时间来设置Modified, ModifiedTime 和 ModifiedDate字段.

不推荐使用：改用Modified.


### <a id="FileHeader.SetMode">func</a> (\*FileHeader) [SetMode](https://golang.org/src/archive/zip/struct.go?s=9760:9806#L287)
<pre>func (h *<a href="#FileHeader">FileHeader</a>) SetMode(mode <a href="/pkg/os/">os</a>.<a href="/pkg/os/#FileMode">FileMode</a>)</pre>
SetMode changes the permission and mode bits for the FileHeader.

SetMode 修改FileHeader的权限和模式位.


## <a id="ReadCloser">type</a> [ReadCloser](https://golang.org/src/archive/zip/reader.go?s=591:637#L22)

<pre>type ReadCloser struct {
    <a href="#Reader">Reader</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="OpenReader">func</a> [OpenReader](https://golang.org/src/archive/zip/reader.go?s=911:960#L40)
<pre>func OpenReader(name <a href="/pkg/builtin/#string">string</a>) (*<a href="#ReadCloser">ReadCloser</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
OpenReader will open the Zip file specified by name and return a ReadCloser.


OpenReader 会打开指定name的zip文档并返回一个ReadCloser.



### <a id="ReadCloser.Close">func</a> (\*ReadCloser) [Close](https://golang.org/src/archive/zip/reader.go?s=3432:3467#L131)
<pre>func (rc *<a href="#ReadCloser">ReadCloser</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the Zip file, rendering it unusable for I/O.

Close会关闭zip文档. 使其无法使用I/O.


## <a id="Reader">type</a> [Reader](https://golang.org/src/archive/zip/reader.go?s=456:589#L15)

<pre>type Reader struct {
<span id="Reader.File"></span>    File    []*<a href="#File">File</a>
<span id="Reader.Comment"></span>    Comment <a href="/pkg/builtin/#string">string</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Reader">Example</a>
```go
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Open a zip archive for reading.
	r, err := zip.OpenReader("testdata/readme.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
    //
    // 遍历文档, 并输出包含的文件的内容.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}
```




### <a id="NewReader">func</a> [NewReader](https://golang.org/src/archive/zip/reader.go?s=1328:1386#L61)
<pre>func NewReader(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReaderAt">ReaderAt</a>, size <a href="/pkg/builtin/#int64">int64</a>) (*<a href="#Reader">Reader</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewReader returns a new Reader reading from r, which is assumed to
have the given size in bytes.

NewReader 返回一个从r读取数据的Reader, 它被假定为具有给定的字节大小.




### <a id="Reader.RegisterDecompressor">func</a> (\*Reader) [RegisterDecompressor](https://golang.org/src/archive/zip/reader.go?s=3014:3086#L115)
<pre>func (z *<a href="#Reader">Reader</a>) RegisterDecompressor(method <a href="/pkg/builtin/#uint16">uint16</a>, dcomp <a href="#Decompressor">Decompressor</a>)</pre>
RegisterDecompressor registers or overrides a custom decompressor for a
specific method ID. If a decompressor for a given method is not found,
Reader will default to looking up the decompressor at the package level.

RegisterDecompressor 会注册或覆盖特定方法ID的自定义解压器. 如果没有找到给定方法的解压器, Reader将默认在包级别查找解压器.


## <a id="Writer">type</a> [Writer](https://golang.org/src/archive/zip/writer.go?s=448:781#L14)
Writer implements a zip file writer.

Writer是zip文档的写入器.

<pre>type Writer struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Writer">Example</a>
```go
package main

import (
	"archive/zip"
	"bytes"
	"log"
)

func main() {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
    // 确保检查Close方法返回的错误.
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
}
```



### <a id="NewWriter">func</a> [NewWriter](https://golang.org/src/archive/zip/writer.go?s=894:929#L33)
<pre>func NewWriter(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) *<a href="#Writer">Writer</a></pre>
NewWriter returns a new Writer writing a zip file to w.

NewWriter 创建一个向zip文档写入的Writer.




### <a id="Writer.Close">func</a> (\*Writer) [Close](https://golang.org/src/archive/zip/writer.go?s=1977:2007#L66)
<pre>func (w *<a href="#Writer">Writer</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close finishes writing the zip file by writing the central directory.
It does not close the underlying writer.

Close方法通过写入central directory信息来关闭该Writer. 它不会关闭底层的io.Writer接口.


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

Create 会使用指定文件名将该文件加入到zip文档中. 它会返回一个用于写入该文件内容的io.Writer. 文件内容将使用Deflate方法压缩. 文件名必须是相对路径: 不能以盘符(比如C:)或者斜杠(/)开头,且只接受'/'作为路径分隔符. 要创建目录时，请在名称后追加斜杠. 文件的内容必须在下一次调用Create, CreateHeader或Close方法之前写入.


### <a id="Writer.CreateHeader">func</a> (\*Writer) [CreateHeader](https://golang.org/src/archive/zip/writer.go?s=8164:8228#L245)
<pre>func (w *<a href="#Writer">Writer</a>) CreateHeader(fh *<a href="#FileHeader">FileHeader</a>) (<a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
CreateHeader adds a file to the zip archive using the provided FileHeader
for the file metadata. Writer takes ownership of fh and may mutate
its fields. The caller must not modify fh after calling CreateHeader.

This returns a Writer to which the file contents should be written.
The file's contents must be written to the io.Writer before the next
call to Create, CreateHeader, or Close.

Create 会使用指定的FileHeader表示文件元信息并将该文件加入到zip文档中. Writer获得fh的所有权, 可能会修改fh的字段. 调用CreateHeader后不能修改fh.

该方法会返回一个Writer用于写入文件内容. 文件内容必须在下一次调用Create, CreateHeader或Close方法之前写入.

### <a id="Writer.Flush">func</a> (\*Writer) [Flush](https://golang.org/src/archive/zip/writer.go?s=1509:1539#L50)
<pre>func (w *<a href="#Writer">Writer</a>) Flush() <a href="/pkg/builtin/#error">error</a></pre>
Flush flushes any buffered data to the underlying writer.
Calling Flush is not normally necessary; calling Close is sufficient.

Flush 会将所有的缓冲数据写入底层的io.Writer. Flush方法不是必须的, 调用Close方法就足够了.


### <a id="Writer.RegisterCompressor">func</a> (\*Writer) [RegisterCompressor](https://golang.org/src/archive/zip/writer.go?s=13353:13420#L401)
<pre>func (w *<a href="#Writer">Writer</a>) RegisterCompressor(method <a href="/pkg/builtin/#uint16">uint16</a>, comp <a href="#Compressor">Compressor</a>)</pre>
RegisterCompressor registers or overrides a custom compressor for a specific
method ID. If a compressor for a given method is not found, Writer will
default to looking up the compressor at the package level.

RegisterCompressor 会注册或覆盖特定方法ID的自定义压缩器. 如果没有找到给定方法的压缩器, Writer将默认在包级别查找压缩器.

<a id="example_Writer_RegisterCompressor">Example</a>
```go
package main

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"io"
)

func main() {
	// Override the default Deflate compressor with a higher compression level.
    // 用具有更高压缩比的压缩器覆盖默认的Deflate压缩器.

	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Register a custom Deflate compressor.
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	// Proceed to add files to w.
}
```


### <a id="Writer.SetComment">func</a> (\*Writer) [SetComment](https://golang.org/src/archive/zip/writer.go?s=1686:1735#L56)
<pre>func (w *<a href="#Writer">Writer</a>) SetComment(comment <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetComment sets the end-of-central-directory comment field.
It can only be called before Close.

SetComment 设置end-of-central-directory的注释字段. 它只能在 Close 之前调用.


### <a id="Writer.SetOffset">func</a> (\*Writer) [SetOffset](https://golang.org/src/archive/zip/writer.go?s=1237:1272#L41)
<pre>func (w *<a href="#Writer">Writer</a>) SetOffset(n <a href="/pkg/builtin/#int64">int64</a>)</pre>
SetOffset sets the offset of the beginning of the zip data within the
underlying writer. It should be used when the zip data is appended to an
existing file, such as a binary executable.
It must be called before any data is written.


SetOffset 用于设置将zip数据写入底层io.Writer中的开始偏移量. 当将zip数据附加到现有文件（比如一个二进制可执行文件）中时就应使用它. 必须在写入任何数据前先调用它.





