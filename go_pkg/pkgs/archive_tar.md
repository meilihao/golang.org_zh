

# tar
`import "archive/tar"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package tar implements access to tar archives.

Tape archives (tar) are a file format for storing a sequence of files that
can be read and written in a streaming manner.
This package aims to cover most variations of the format,
including those produced by GNU and BSD tar tools.

tar 实现对tar格式文档的访问.

tar是一种文件格式, 用于存储一系列的文件, 这些文件可以以stream的方式读写.
本包旨在涵盖绝大多数tar的变种, 包括GNU和BSD生成的tar文档. 


<a id="example__minimal">Example (Minimal)</a>
```go
package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Create and add some files to the archive.
    // 创建一个buf用于写缓冲.
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	// Open and iterate through the files in the archive.
    // 打开tar文档并遍历其中的文件.
	tr := tar.NewReader(&buf)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}

}
```

output:
```txt
Contents of readme.txt:
This archive contains some text files.
Contents of gopher.txt:
Gopher names:
George
Geoffrey
Gonzo
Contents of todo.txt:
Get animal handling license.
```


## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [type Format](#Format)
  * [func (f Format) String() string](#Format.String)
* [type Header](#Header)
  * [func FileInfoHeader(fi os.FileInfo, link string) (*Header, error)](#FileInfoHeader)
  * [func (h *Header) FileInfo() os.FileInfo](#Header.FileInfo)
* [type Reader](#Reader)
  * [func NewReader(r io.Reader) *Reader](#NewReader)
  * [func (tr *Reader) Next() (*Header, error)](#Reader.Next)
  * [func (tr *Reader) Read(b []byte) (int, error)](#Reader.Read)
* [type Writer](#Writer)
  * [func NewWriter(w io.Writer) *Writer](#NewWriter)
  * [func (tw *Writer) Close() error](#Writer.Close)
  * [func (tw *Writer) Flush() error](#Writer.Flush)
  * [func (tw *Writer) Write(b []byte) (int, error)](#Writer.Write)
  * [func (tw *Writer) WriteHeader(hdr *Header) error](#Writer.WriteHeader)


#### <a id="pkg-examples">Examples</a>
* [Package (Minimal)](#example__minimal)


#### <a id="pkg-files">Package files</a>
[common.go](https://golang.org/src/archive/tar/common.go) [format.go](https://golang.org/src/archive/tar/format.go) [reader.go](https://golang.org/src/archive/tar/reader.go) [stat_actime1.go](https://golang.org/src/archive/tar/stat_actime1.go) [stat_unix.go](https://golang.org/src/archive/tar/stat_unix.go) [strconv.go](https://golang.org/src/archive/tar/strconv.go) [writer.go](https://golang.org/src/archive/tar/writer.go) 


## <a id="pkg-constants">Constants</a>
Type flags for Header.Typeflag.


```go
const (
    // Type '0' indicates a regular file.
    // '0'表示普通文件
    TypeReg  = '0'
    TypeRegA = '\x00' // Deprecated: Use TypeReg instead.

    // Type '1' to '6' are header-only flags and may not have a data body.
    // '1' ~ '6'是header-only flag, 可能没有具体数据.
    TypeLink    = '1' // Hard link // 硬链接
    TypeSymlink = '2' // Symbolic link // 软链接
    TypeChar    = '3' // Character device node // 字符设备节点
    TypeBlock   = '4' // Block device node // 块设备节点
    TypeDir     = '5' // Directory // 目录
    TypeFifo    = '6' // FIFO node // fifo节点

    // Type '7' is reserved. // 保留项
    TypeCont = '7'

    // Type 'x' is used by the PAX format to store key-value records that
    // are only relevant to the next file.
    // This package transparently handles these types.
    //
    // 'x'是PAX格式用于保存kv记录(该记录仅与下一个文件有关)
    // 本包透明地处理这些类型.
    TypeXHeader = 'x'

    // Type 'g' is used by the PAX format to store key-value records that
    // are relevant to all subsequent files.
    // This package only supports parsing and composing such headers,
    // but does not currently support persisting the global state across files.
    //
    // 'g'是PAX格式用于保存kv记录(该记录与所有后续文件有关).
    // 本包仅支持解析和组合这类header, 但目前还不支持在文件间保留全局状态???.
    TypeXGlobalHeader = 'g'

    // Type 'S' indicates a sparse file in the GNU format.
    // 表示GNU格式中的稀疏文件
    TypeGNUSparse = 'S'

    // Types 'L' and 'K' are used by the GNU format for a meta file
    // used to store the path or link name for the next file.
    // This package transparently handles these types.
    //
    // 'L' 和 'K' 是GNU格式的元数据文件用于存储下一个文件的路径或链接名称.
    // 本包透明地处理这些类型.
    TypeGNULongName = 'L'
    TypeGNULongLink = 'K'
)
```

## <a id="pkg-variables">Variables</a>

```go
var (
    ErrHeader          = errors.New("archive/tar: invalid tar header") // 无效的tar header
    ErrWriteTooLong    = errors.New("archive/tar: write too long") // 写入内容过长
    ErrFieldTooLong    = errors.New("archive/tar: header field too long") // header字段过长
    ErrWriteAfterClose = errors.New("archive/tar: write after close") // 在关闭后写入
)
```



## <a id="Format">type</a> [Format](https://golang.org/src/archive/tar/format.go?s=2010:2025#L36)
Format represents the tar archive format.

The original tar format was introduced in Unix V7.
Since then, there have been multiple competing formats attempting to
standardize or extend the V7 format to overcome its limitations.
The most common formats are the USTAR, PAX, and GNU formats,
each with their own advantages and limitations.

The following table captures the capabilities of each format:


	                  |  USTAR |       PAX |       GNU
	------------------+--------+-----------+----------
	Name              |   256B | unlimited | unlimited
	Linkname          |   100B | unlimited | unlimited
	Size              | uint33 | unlimited |    uint89
	Mode              | uint21 |    uint21 |    uint57
	Uid/Gid           | uint21 | unlimited |    uint57
	Uname/Gname       |    32B | unlimited |       32B
	ModTime           | uint33 | unlimited |     int89
	AccessTime        |    n/a | unlimited |     int89
	ChangeTime        |    n/a | unlimited |     int89
	Devmajor/Devminor | uint21 |    uint21 |    uint57
	------------------+--------+-----------+----------
	string encoding   |  ASCII |     UTF-8 |    binary
	sub-second times  |     no |       yes |        no
	sparse files      |     no |       yes |       yes

The table's upper portion shows the Header fields, where each format reports
the maximum number of bytes allowed for each string field and
the integer type used to store each numeric field
(where timestamps are stored as the number of seconds since the Unix epoch).

The table's lower portion shows specialized features of each format,
such as supported string encodings, support for sub-second timestamps,
or support for sparse files.

The Writer currently provides no support for sparse files.

Format 标识 tar文件的格式.

最初tar格式是在Unix V7引入的. 从那时起, 已有多种竞争的格式试图标准化或扩展V7格式以克服其局限性. 最常见的格式是STAR, PAX 和 GNU, 每种格式都有其自身的优点和局限性.

下表列举了每种格式的功能:


	                  |  USTAR |       PAX |       GNU
	------------------+--------+-----------+----------
	Name              |   256B | unlimited | unlimited
	Linkname          |   100B | unlimited | unlimited
	Size              | uint33 | unlimited |    uint89
	Mode              | uint21 |    uint21 |    uint57
	Uid/Gid           | uint21 | unlimited |    uint57
	Uname/Gname       |    32B | unlimited |       32B
	ModTime           | uint33 | unlimited |     int89
	AccessTime        |    n/a | unlimited |     int89
	ChangeTime        |    n/a | unlimited |     int89
	Devmajor/Devminor | uint21 |    uint21 |    uint57
	------------------+--------+-----------+----------
	string encoding   |  ASCII |     UTF-8 |    binary
	sub-second times  |     no |       yes |        no
	sparse files      |     no |       yes |       yes

表格的上部显示了每种格式每个字符串字段允许的最大字节数, 以及用于存储每个数值字段的整数类型(时间戳是存储在Unix纪元以来的秒数).

表格的下部显示了每种格式的特殊功能, 比如支持的字符串编码, 对亚秒及时间戳的支持或支持稀疏文件.

Writer 当前不支持稀疏文件.

<pre>type Format <a href="/pkg/builtin/#int">int</a></pre>


Constants to identify various tar formats.

标识各种tar格式的常量.

```go
const (

    // FormatUnknown indicates that the format is unknown.
    // 表示未知格式
    FormatUnknown Format

    // FormatUSTAR represents the USTAR header format defined in POSIX.1-1988.
    //
    // While this format is compatible with most tar readers,
    // the format has several limitations making it unsuitable for some usages.
    // Most notably, it cannot support sparse files, files larger than 8GiB,
    // filenames larger than 256 characters, and non-ASCII filenames.
    //
    // Reference:
    //	http://pubs.opengroup.org/onlinepubs/9699919799/utilities/pax.html#tag_20_92_13_06
    //
    // FormatUSTAR 表示 定义在POSIX.1-1988中的USTAR header.
    //
    // 虽然此格式与大多数tar reader兼容, 但该格式还是有一些限制, 使其不适合用于某些用途.
    // 最值得注意的是，它不支持稀疏文件, 大于8GiB的文件, 大于256个字符的文件名和非ASCII文件名.
    //
    // Reference:
    //	http://pubs.opengroup.org/onlinepubs/9699919799/utilities/pax.html#tag_20_92_13_06
    FormatUSTAR

    // FormatPAX represents the PAX header format defined in POSIX.1-2001.
    //
    // PAX extends USTAR by writing a special file with Typeflag TypeXHeader
    // preceding the original header. This file contains a set of key-value
    // records, which are used to overcome USTAR's shortcomings, in addition to
    // providing the ability to have sub-second resolution for timestamps.
    //
    // Some newer formats add their own extensions to PAX by defining their
    // own keys and assigning certain semantic meaning to the associated values.
    // For example, sparse file support in PAX is implemented using keys
    // defined by the GNU manual (e.g., "GNU.sparse.map").
    //
    // Reference:
    //	http://pubs.opengroup.org/onlinepubs/009695399/utilities/pax.html
    //
    // FormatPAX 表示定义在 POSIX.1-2001 中的PAX header.
    //
    // PAX通过在原始标头之前使用Typeflag TypeXHeader写入特殊文件来扩展USTAR.
    // 该文件包含一组键值记录，除了克服USTAR的缺点外，提供了亚秒级时间戳的功能.
    //
    // 一些较新的格式通过定义它们自己的扩展名到PAX, 并为相应的值分配某些语义.
    // 例如，PAX中的稀疏文件支持由GNU手册定义(例如,"GNU.sparse.map")的keys来实现.
    //
    // Reference:
    //	http://pubs.opengroup.org/onlinepubs/009695399/utilities/pax.html
    FormatPAX

    // FormatGNU represents the GNU header format.
    //
    // The GNU header format is older than the USTAR and PAX standards and
    // is not compatible with them. The GNU format supports
    // arbitrary file sizes, filenames of arbitrary encoding and length,
    // sparse files, and other features.
    //
    // It is recommended that PAX be chosen over GNU unless the target
    // application can only parse GNU formatted archives.
    //
    // Reference:
    //	https://www.gnu.org/software/tar/manual/html_node/Standard.html
    //
    // FormatGNU 表示 GNU header.
    //
    // GNU header格式早于USTAR 和 PAX 标准, 且与它们不兼容.  GNU格式支持任意文件大小，任意编码和长度的文件名，
    // 支持稀疏文件和其他功能.
    //
    // 除非目标应用只能解析GNU格式的tar文件, 否则推荐使用PAX.
    //
    // Reference:
    //	https://www.gnu.org/software/tar/manual/html_node/Standard.html
    FormatGNU
)
```


### <a id="Format.String">func</a> (Format) [String](https://golang.org/src/archive/tar/format.go?s=4648:4679#L107)
<pre>func (f <a href="#Format">Format</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Header">type</a> [Header](https://golang.org/src/archive/tar/common.go?s=4721:7248#L130)
A Header represents a single header in a tar archive.
Some fields may not be populated.

For forward compatibility, users that retrieve a Header from Reader.Next,
mutate it in some ways, and then pass it back to Writer.WriteHeader
should do so by creating a new Header and copying the fields
that they are interested in preserving.

Header 表示 tar文档中的当个header, 某些字段可能未填充.

为了向前兼容, 用户需要从Reader.Next获取一个Header, 以某种方式对其进行变换, 然后将其传递回Writer.WriteHeader, 无论创建新的Header还是拷贝他们感兴趣的字段都应这么做.

```go
type Header struct {
    // Typeflag is the type of header entry.
    // The zero value is automatically promoted to either TypeReg or TypeDir
    // depending on the presence of a trailing slash in Name.
    Typeflag byte

    Name     string // Name of file entry
    Linkname string // Target name of link (valid for TypeLink or TypeSymlink)

    Size  int64  // Logical file size in bytes
    Mode  int64  // Permission and mode bits
    Uid   int    // User ID of owner
    Gid   int    // Group ID of owner
    Uname string // User name of owner
    Gname string // Group name of owner

    // If the Format is unspecified, then Writer.WriteHeader rounds ModTime
    // to the nearest second and ignores the AccessTime and ChangeTime fields.
    //
    // To use AccessTime or ChangeTime, specify the Format as PAX or GNU.
    // To use sub-second resolution, specify the Format as PAX.
    ModTime    time.Time // Modification time
    AccessTime time.Time // Access time (requires either PAX or GNU support)
    ChangeTime time.Time // Change time (requires either PAX or GNU support)

    Devmajor int64 // Major device number (valid for TypeChar or TypeBlock)
    Devminor int64 // Minor device number (valid for TypeChar or TypeBlock)

    // Xattrs stores extended attributes as PAX records under the
    // "SCHILY.xattr." namespace.
    //
    // The following are semantically equivalent:
    //  h.Xattrs[key] = value
    //  h.PAXRecords["SCHILY.xattr."+key] = value
    //
    // When Writer.WriteHeader is called, the contents of Xattrs will take
    // precedence over those in PAXRecords.
    //
    // Deprecated: Use PAXRecords instead.
    Xattrs map[string]string // Go 1.3

    // PAXRecords is a map of PAX extended header records.
    //
    // User-defined records should have keys of the following form:
    //	VENDOR.keyword
    // Where VENDOR is some namespace in all uppercase, and keyword may
    // not contain the '=' character (e.g., "GOLANG.pkg.version").
    // The key and value should be non-empty UTF-8 strings.
    //
    // When Writer.WriteHeader is called, PAX records derived from the
    // other fields in Header take precedence over PAXRecords.
    PAXRecords map[string]string // Go 1.10

    // Format specifies the format of the tar header.
    //
    // This is set by Reader.Next as a best-effort guess at the format.
    // Since the Reader liberally reads some non-compliant files,
    // it is possible for this to be FormatUnknown.
    //
    // If the format is unspecified when Writer.WriteHeader is called,
    // then it uses the first format (in the order of USTAR, PAX, GNU)
    // capable of encoding this Header (see Format).
    Format Format // Go 1.10
}
```




### <a id="FileInfoHeader">func</a> [FileInfoHeader](https://golang.org/src/archive/tar/common.go?s=22088:22153#L619)
<pre>func FileInfoHeader(fi <a href="/pkg/os/">os</a>.<a href="/pkg/os/#FileInfo">FileInfo</a>, link <a href="/pkg/builtin/#string">string</a>) (*<a href="#Header">Header</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
FileInfoHeader creates a partially-populated Header from fi.
If fi describes a symlink, FileInfoHeader records link as the link target.
If fi describes a directory, a slash is appended to the name.

Since os.FileInfo's Name method only returns the base name of
the file it describes, it may be necessary to modify Header.Name
to provide the full path name of the file.






### <a id="Header.FileInfo">func</a> (\*Header) [FileInfo](https://golang.org/src/archive/tar/common.go?s=19247:19286#L519)
<pre>func (h *<a href="#Header">Header</a>) FileInfo() <a href="/pkg/os/">os</a>.<a href="/pkg/os/#FileInfo">FileInfo</a></pre>
FileInfo returns an os.FileInfo for the Header.




## <a id="Reader">type</a> [Reader](https://golang.org/src/archive/tar/reader.go?s=470:845#L9)
Reader provides sequential access to the contents of a tar archive.
Reader.Next advances to the next file in the archive (including the first),
and then Reader can be treated as an io.Reader to access the file's data.


<pre>type Reader struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewReader">func</a> [NewReader](https://golang.org/src/archive/tar/reader.go?s=986:1021#L29)
<pre>func NewReader(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) *<a href="#Reader">Reader</a></pre>
NewReader creates a new Reader reading from r.






### <a id="Reader.Next">func</a> (\*Reader) [Next](https://golang.org/src/archive/tar/reader.go?s=1328:1369#L38)
<pre>func (tr *<a href="#Reader">Reader</a>) Next() (*<a href="#Header">Header</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Next advances to the next entry in the tar archive.
The Header.Size determines how many bytes can be read for the next file.
Any remaining data in the current file is automatically discarded.

io.EOF is returned at the end of the input.




### <a id="Reader.Read">func</a> (\*Reader) [Read](https://golang.org/src/archive/tar/reader.go?s=19854:19899#L612)
<pre>func (tr *<a href="#Reader">Reader</a>) Read(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Read reads from the current file in the tar archive.
It returns (0, io.EOF) when it reaches the end of that file,
until Next is called to advance to the next file.

If the current file is sparse, then the regions marked as a hole
are read back as NUL-bytes.

Calling Read on special types like TypeLink, TypeSymlink, TypeChar,
TypeBlock, TypeDir, and TypeFifo returns (0, io.EOF) regardless of what
the Header.Size claims.




## <a id="Writer">type</a> [Writer](https://golang.org/src/archive/tar/writer.go?s=432:876#L9)
Writer provides sequential writing of a tar archive.
Write.WriteHeader begins a new file with the provided Header,
and then Writer can be treated as an io.Writer to supply that file's data.


<pre>type Writer struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewWriter">func</a> [NewWriter](https://golang.org/src/archive/tar/writer.go?s=926:961#L23)
<pre>func NewWriter(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) *<a href="#Writer">Writer</a></pre>
NewWriter creates a new Writer writing to w.






### <a id="Writer.Close">func</a> (\*Writer) [Close](https://golang.org/src/archive/tar/writer.go?s=14125:14156#L456)
<pre>func (tw *<a href="#Writer">Writer</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the tar archive by flushing the padding, and writing the footer.
If the current file (from a prior call to WriteHeader) is not fully written,
then this returns an error.




### <a id="Writer.Flush">func</a> (\*Writer) [Flush](https://golang.org/src/archive/tar/writer.go?s=1353:1384#L39)
<pre>func (tw *<a href="#Writer">Writer</a>) Flush() <a href="/pkg/builtin/#error">error</a></pre>
Flush finishes writing the current file's block padding.
The current file must be fully written before Flush can be called.

This is unnecessary as the next call to WriteHeader or Close
will implicitly flush out the file's padding.




### <a id="Writer.Write">func</a> (\*Writer) [Write](https://golang.org/src/archive/tar/writer.go?s=13015:13061#L421)
<pre>func (tw *<a href="#Writer">Writer</a>) Write(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Write writes to the current file in the tar archive.
Write returns the error ErrWriteTooLong if more than
Header.Size bytes are written after WriteHeader.

Calling Write on special types like TypeLink, TypeSymlink, TypeChar,
TypeBlock, TypeDir, and TypeFifo returns (0, ErrWriteTooLong) regardless
of what the Header.Size claims.




### <a id="Writer.WriteHeader">func</a> (\*Writer) [WriteHeader](https://golang.org/src/archive/tar/writer.go?s=1948:1996#L57)
<pre>func (tw *<a href="#Writer">Writer</a>) WriteHeader(hdr *<a href="#Header">Header</a>) <a href="/pkg/builtin/#error">error</a></pre>
WriteHeader writes hdr and prepares to accept the file's contents.
The Header.Size determines how many bytes can be written for the next file.
If the current file is not fully written, then this returns an error.
This implicitly flushes any padding necessary before writing the header.







