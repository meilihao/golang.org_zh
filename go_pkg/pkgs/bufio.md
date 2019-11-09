# bufio
`import "bufio"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer
object, creating another object (Reader or Writer) that also implements
the interface but provides buffering and some help for textual I/O.

bufio 实现了带缓冲的I/O. 它封装了一个io.Reader或者io.Writer对象, 创建了另外一个对象(Reader或者Writer), 这个对象也实现了该接口, 并提供了缓冲机制, 同时支持一些文本I/O的帮助函数.

## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)](#ScanBytes)
* [func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)](#ScanLines)
* [func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)](#ScanRunes)
* [func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)](#ScanWords)
* [type ReadWriter](#ReadWriter)
  * [func NewReadWriter(r *Reader, w *Writer) *ReadWriter](#NewReadWriter)
* [type Reader](#Reader)
  * [func NewReader(rd io.Reader) *Reader](#NewReader)
  * [func NewReaderSize(rd io.Reader, size int) *Reader](#NewReaderSize)
  * [func (b *Reader) Buffered() int](#Reader.Buffered)
  * [func (b *Reader) Discard(n int) (discarded int, err error)](#Reader.Discard)
  * [func (b *Reader) Peek(n int) ([]byte, error)](#Reader.Peek)
  * [func (b *Reader) Read(p []byte) (n int, err error)](#Reader.Read)
  * [func (b *Reader) ReadByte() (byte, error)](#Reader.ReadByte)
  * [func (b *Reader) ReadBytes(delim byte) ([]byte, error)](#Reader.ReadBytes)
  * [func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)](#Reader.ReadLine)
  * [func (b *Reader) ReadRune() (r rune, size int, err error)](#Reader.ReadRune)
  * [func (b *Reader) ReadSlice(delim byte) (line []byte, err error)](#Reader.ReadSlice)
  * [func (b *Reader) ReadString(delim byte) (string, error)](#Reader.ReadString)
  * [func (b *Reader) Reset(r io.Reader)](#Reader.Reset)
  * [func (b *Reader) Size() int](#Reader.Size)
  * [func (b *Reader) UnreadByte() error](#Reader.UnreadByte)
  * [func (b *Reader) UnreadRune() error](#Reader.UnreadRune)
  * [func (b *Reader) WriteTo(w io.Writer) (n int64, err error)](#Reader.WriteTo)
* [type Scanner](#Scanner)
  * [func NewScanner(r io.Reader) *Scanner](#NewScanner)
  * [func (s *Scanner) Buffer(buf []byte, max int)](#Scanner.Buffer)
  * [func (s *Scanner) Bytes() []byte](#Scanner.Bytes)
  * [func (s *Scanner) Err() error](#Scanner.Err)
  * [func (s *Scanner) Scan() bool](#Scanner.Scan)
  * [func (s *Scanner) Split(split SplitFunc)](#Scanner.Split)
  * [func (s *Scanner) Text() string](#Scanner.Text)
* [type SplitFunc](#SplitFunc)
* [type Writer](#Writer)
  * [func NewWriter(w io.Writer) *Writer](#NewWriter)
  * [func NewWriterSize(w io.Writer, size int) *Writer](#NewWriterSize)
  * [func (b *Writer) Available() int](#Writer.Available)
  * [func (b *Writer) Buffered() int](#Writer.Buffered)
  * [func (b *Writer) Flush() error](#Writer.Flush)
  * [func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)](#Writer.ReadFrom)
  * [func (b *Writer) Reset(w io.Writer)](#Writer.Reset)
  * [func (b *Writer) Size() int](#Writer.Size)
  * [func (b *Writer) Write(p []byte) (nn int, err error)](#Writer.Write)
  * [func (b *Writer) WriteByte(c byte) error](#Writer.WriteByte)
  * [func (b *Writer) WriteRune(r rune) (size int, err error)](#Writer.WriteRune)
  * [func (b *Writer) WriteString(s string) (int, error)](#Writer.WriteString)


#### <a id="pkg-examples">Examples</a>
* [Scanner.Bytes](#example_Scanner_Bytes)
* [Scanner (Custom)](#example_Scanner_custom)
* [Scanner (EmptyFinalToken)](#example_Scanner_emptyFinalToken)
* [Scanner (Lines)](#example_Scanner_lines)
* [Scanner (Words)](#example_Scanner_words)
* [Writer](#example_Writer)


#### <a id="pkg-files">Package files</a>
[bufio.go](https://golang.org/src/bufio/bufio.go) [scan.go](https://golang.org/src/bufio/scan.go) 


## <a id="pkg-constants">Constants</a>

```go
const (
    // MaxScanTokenSize is the maximum size used to buffer a token
    // unless the user provides an explicit buffer with Scanner.Buffer.
    // The actual maximum token size may be smaller as the buffer
    // may need to include, for instance, a newline.
    //
    // 除非用户使用Scan.Buffer并指定缓冲容量, 那么MaxScanTokenSize就是用于缓冲一个token的最大容量.
    // 实际的最大token容量可能小于buffer容量,比如一个换行符.
    MaxScanTokenSize = 64 * 1024
)
```

## <a id="pkg-variables">Variables</a>

```go
var (
    ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte") // 非法调用UnreadByte
    ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune") // 非法调用UnreadRune
    ErrBufferFull        = errors.New("bufio: buffer full") // 缓冲区已满
    ErrNegativeCount     = errors.New("bufio: negative count") // 负计数
)
```

Errors returned by Scanner.

Scanner 返回的错误.


```go
var (
    ErrTooLong         = errors.New("bufio.Scanner: token too long") // token过长
    ErrNegativeAdvance = errors.New("bufio.Scanner: SplitFunc returns negative advance count") // SplitFunc返回了负的递进数
    ErrAdvanceTooFar   = errors.New("bufio.Scanner: SplitFunc returns advance count beyond input") // SplitFunc返回的递进数超出了输入
)
```


ErrFinalToken is a special sentinel error value. It is intended to be
returned by a Split function to indicate that the token being delivered
with the error is the last token and scanning should stop after this one.
After ErrFinalToken is received by Scan, scanning stops with no error.
The value is useful to stop processing early or when it is necessary to
deliver a final empty token. One could achieve the same behavior
with a custom error value but providing one here is tidier.
See the emptyFinalToken example for a use of this value.

ErrFinalToken是一个特殊的保护性的错误值. 分割函数通过返回它来表示这个带着该错误的token已经是最后一个token了,扫描操作应该停止了. Scan收到ErrFinalToken之后,扫描操作就会停止且不返回错误. 该值用于提前结束操作或有必要在最后传递一个空token的情况. 我们能用自定义的错误实现同样的逻辑,但这里提供该值显得更整洁. 可查看emptyFinalToken的例子来使用该值. 

```go
var ErrFinalToken = errors.New("final token") // 最后的token
```

## <a id="ScanBytes">func</a> [ScanBytes](https://golang.org/src/bufio/scan.go?s=9596:9674#L274)
<pre>func ScanBytes(data []<a href="/pkg/builtin/#byte">byte</a>, atEOF <a href="/pkg/builtin/#bool">bool</a>) (advance <a href="/pkg/builtin/#int">int</a>, token []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ScanBytes is a split function for a Scanner that returns each byte as a token.

ScanBytes 是用于Scanner的分割函数, 它会将每个字节作为一个token返回.

## <a id="ScanLines">func</a> [ScanLines](https://golang.org/src/bufio/scan.go?s=11802:11880#L335)
<pre>func ScanLines(data []<a href="/pkg/builtin/#byte">byte</a>, atEOF <a href="/pkg/builtin/#bool">bool</a>) (advance <a href="/pkg/builtin/#int">int</a>, token []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ScanLines is a split function for a Scanner that returns each line of
text, stripped of any trailing end-of-line marker. The returned line may
be empty. The end-of-line marker is one optional carriage return followed
by one mandatory newline. In regular expression notation, it is `\r?\n`.
The last non-empty line of input will be returned even if it has no
newline.

 ScanLines 是用于Scanner的分割函数, 它会将每一行文本(已去掉末尾的换行标记)作为一个token返回. 返回的行可以是空字符串. 换行标记为一个可选的回车符后跟一个必选的换行符,在正则表达式中用`\r?\n`表示. 最后一个非空行即使没有换行符也会作为一个token返回.


## <a id="ScanRunes">func</a> [ScanRunes](https://golang.org/src/bufio/scan.go?s=10243:10321#L289)
<pre>func ScanRunes(data []<a href="/pkg/builtin/#byte">byte</a>, atEOF <a href="/pkg/builtin/#bool">bool</a>) (advance <a href="/pkg/builtin/#int">int</a>, token []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ScanRunes is a split function for a Scanner that returns each
UTF-8-encoded rune as a token. The sequence of runes returned is
equivalent to that from a range loop over the input as a string, which
means that erroneous UTF-8 encodings translate to U+FFFD = "\xef\xbf\xbd".
Because of the Scan interface, this makes it impossible for the client to
distinguish correctly encoded replacement runes from encoding errors.

ScanRunes 是用于Scanner的分割函数, 它会将每个utf-8编码的unicode码值(rune)作为一个token返回. 返回的rune序列和range一个字符串的输出(rune序列)相同,这意味着错误的utf-8编码会被翻译为U+FFFD = "\xef\xbf\xbd". 因此,Scan接口不能从这些编码错误时返回的值中正确地区分出真正的编码替代字符(U+FFFD).

## <a id="ScanWords">func</a> [ScanWords](https://golang.org/src/bufio/scan.go?s=13096:13174#L380)
<pre>func ScanWords(data []<a href="/pkg/builtin/#byte">byte</a>, atEOF <a href="/pkg/builtin/#bool">bool</a>) (advance <a href="/pkg/builtin/#int">int</a>, token []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ScanWords is a split function for a Scanner that returns each
space-separated word of text, with surrounding spaces deleted. It will
never return an empty string. The definition of space is set by
unicode.IsSpace.


ScanWords 是用于Scanner的分割函数,它会将每个用空白符分隔的词(已去掉前后空白符)作为一个token返回. 它永远不会返回空字符串.空白字符是由unicode.IsSpace定义.


## <a id="ReadWriter">type</a> [ReadWriter](https://golang.org/src/bufio/bufio.go?s=18720:18764#L745)
ReadWriter stores pointers to a Reader and a Writer.
It implements io.ReadWriter.

ReadWriter 保存 Reader 和 Writer 的指针。它实现了 io.ReadWriter.

<pre>type ReadWriter struct {
    *<a href="#Reader">Reader</a>
    *<a href="#Writer">Writer</a>
}
</pre>









### <a id="NewReadWriter">func</a> [NewReadWriter](https://golang.org/src/bufio/bufio.go?s=18838:18890#L751)
<pre>func NewReadWriter(r *<a href="#Reader">Reader</a>, w *<a href="#Writer">Writer</a>) *<a href="#ReadWriter">ReadWriter</a></pre>
NewReadWriter allocates a new ReadWriter that dispatches to r and w.


NewReadWriter 会创建一个新的 ReadWriter，并将操作分派给r和w.



## <a id="Reader">type</a> [Reader](https://golang.org/src/bufio/bufio.go?s=829:1151#L21)
Reader implements buffering for an io.Reader object.

Reader 实现了带buffer机制的io.Reader.

<pre>type Reader struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewReader">func</a> [NewReader](https://golang.org/src/bufio/bufio.go?s=1757:1793#L51)
<pre>func NewReader(rd <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) *<a href="#Reader">Reader</a></pre>
NewReader returns a new Reader whose buffer has the default size.

NewReader返回一个新的Reader, 其内部buffer为默认大小.


### <a id="NewReaderSize">func</a> [NewReaderSize](https://golang.org/src/bufio/bufio.go?s=1414:1464#L36)
<pre>func NewReaderSize(rd <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, size <a href="/pkg/builtin/#int">int</a>) *<a href="#Reader">Reader</a></pre>
NewReaderSize returns a new Reader whose buffer has at least the specified
size. If the argument io.Reader is already a Reader with large enough
size, it returns the underlying Reader.

NewReaderSize 返回一个新的Reader,参数size指定了其内部buffer的最小值. 如果参数r已经是一个具有足够大缓冲区的Reader类型值,它会返回这个底层的Reader.




### <a id="Reader.Buffered">func</a> (\*Reader) [Buffered](https://golang.org/src/bufio/bufio.go?s=7796:7827#L308)
<pre>func (b *<a href="#Reader">Reader</a>) Buffered() <a href="/pkg/builtin/#int">int</a></pre>
Buffered returns the number of bytes that can be read from the current buffer.


Buffered 返回当前缓冲的可读字节数.

### <a id="Reader.Discard">func</a> (\*Reader) [Discard](https://golang.org/src/bufio/bufio.go?s=4163:4221#L153)
<pre>func (b *<a href="#Reader">Reader</a>) Discard(n <a href="/pkg/builtin/#int">int</a>) (discarded <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Discard skips the next n bytes, returning the number of bytes discarded.

If Discard skips fewer than n bytes, it also returns an error.
If 0 <= n <= b.Buffered(), Discard is guaranteed to succeed without
reading from the underlying io.Reader.

Discard 会跳过接下来的n个字节,并返回跳过的字节数.

如果Discard跳过的字节数比n小,它也会返回一个error.如果0 <= n <= b.Buffered(),Discard保证成功且不会从底层io.Reader中读取数据.


### <a id="Reader.Peek">func</a> (\*Reader) [Peek](https://golang.org/src/bufio/bufio.go?s=3366:3410#L119)
<pre>func (b *<a href="#Reader">Reader</a>) Peek(n <a href="/pkg/builtin/#int">int</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Peek returns the next n bytes without advancing the reader. The bytes stop
being valid at the next read call. If Peek returns fewer than n bytes, it
also returns an error explaining why the read is short. The error is
ErrBufferFull if n is larger than b's buffer size.

Calling Peek prevents a UnreadByte or UnreadRune call from succeeding
until the next read operation.

Peek 会返回后续的n个字节且不移动读取游标. 下一次调用read前, 返回的字节都是有效的. 如果Peek返回的字节数小于n,它也会返回一个error来说明原因. 如果n大于b的缓冲区大小,返回的错误是ErrBufferFull.

调用Peek可防止unreadyte或UnreadRune调用在下次读取操作之前成功.???

### <a id="Reader.Read">func</a> (\*Reader) [Read](https://golang.org/src/bufio/bufio.go?s=4864:4914#L187)
<pre>func (b *<a href="#Reader">Reader</a>) Read(p []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Read reads data into p.
It returns the number of bytes read into p.
The bytes are taken from at most one Read on the underlying Reader,
hence n may be less than len(p).
To read exactly len(p) bytes, use io.ReadFull(b, p).
At EOF, the count will be zero and err will be io.EOF.

Read 会读取数据到p中.它会返回读取到p中的字节数. 读取数据时,至多只会调用一次底层Reader的Read方法,因此n可能会小于len(p). 要精确读取到len(n)个字节, 请使用io.ReadFull(b, p).
读取到结尾时，返回值n是0且err是io.EOF.


### <a id="Reader.ReadByte">func</a> (\*Reader) [ReadByte](https://golang.org/src/bufio/bufio.go?s=5791:5832#L236)
<pre>func (b *<a href="#Reader">Reader</a>) ReadByte() (<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadByte reads and returns a single byte.
If no byte is available, returns an error.

ReadByte 会读取一个字节.如果没有字节可读，会返回一个error.


### <a id="Reader.ReadBytes">func</a> (\*Reader) [ReadBytes](https://golang.org/src/bufio/bufio.go?s=11350:11404#L419)
<pre>func (b *<a href="#Reader">Reader</a>) ReadBytes(delim <a href="/pkg/builtin/#byte">byte</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadBytes reads until the first occurrence of delim in the input,
returning a slice containing the data up to and including the delimiter.
If ReadBytes encounters an error before finding a delimiter,
it returns the data read before the error and the error itself (often io.EOF).
ReadBytes returns err != nil if and only if the returned data does not end in
delim.
For simple uses, a Scanner may be more convenient.

ReadBytes 会读取数据直到第一次遇到分隔符delim，返回一个包含已读取到数据的slice(包括分隔符). 如果ReadBytes在遇到终止符之前就捕获到一个错误,它会返回遇到错误之前已读取的数据和这个错误(经常是io.EOF). 当且仅当ReadBytes方法返回的数据不以delim结尾时, 返回的err != nil. 对于简单的使用, 或许使用Scanner会更方便.


### <a id="Reader.ReadLine">func</a> (\*Reader) [ReadLine](https://golang.org/src/bufio/bufio.go?s=10157:10224#L377)
<pre>func (b *<a href="#Reader">Reader</a>) ReadLine() (line []<a href="/pkg/builtin/#byte">byte</a>, isPrefix <a href="/pkg/builtin/#bool">bool</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadLine is a low-level line-reading primitive. Most callers should use
ReadBytes('\n') or ReadString('\n') instead or use a Scanner.

ReadLine tries to return a single line, not including the end-of-line bytes.
If the line was too long for the buffer then isPrefix is set and the
beginning of the line is returned. The rest of the line will be returned
from future calls. isPrefix will be false when returning the last fragment
of the line. The returned buffer is only valid until the next call to
ReadLine. ReadLine either returns a non-nil line or it returns an error,
never both.

The text returned from ReadLine does not include the line end ("\r\n" or "\n").
No indication or error is given if the input ends without a final line end.
Calling UnreadByte after ReadLine will always unread the last byte read
(possibly a character belonging to the line end) even if that byte is not
part of the line returned by ReadLine.


ReadLine 是一个底层的行读取原语.大多数情况下，应该使用ReadBytes('\n')或ReadString('\n')来代替,或者使用一个Scanner.

ReadLine会尝试返回一行数据, 不包括行尾的行结束符. 如果行太长超过了缓冲，返回值isPrefix会被设为true，并返回行的前面一部分. 该行剩下的部分将在之后的调用中返回. 返回值isPrefix会在返回该行最后一个片段时才设为false. 返回的数据仅在下一次调用ReadLine前有效. ReadLine会返回了一个非空行或者返回一个error，但是两者不会同时返回.

ReadLine返回的文本不会包含行结束符("\r\n"或者"\n"). 如果输入流结束时没有带最终的行结束符, 它不会给出提示或返回错误. 在ReadLine之后调用UnreadByte将总是撤销已读取的最后一个字节(可能是属于行结束符), 即便该字节并非ReadLine返回的行的一部分.

### <a id="Reader.ReadRune">func</a> (\*Reader) [ReadRune](https://golang.org/src/bufio/bufio.go?s=6786:6843#L275)
<pre>func (b *<a href="#Reader">Reader</a>) ReadRune() (r <a href="/pkg/builtin/#rune">rune</a>, size <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadRune reads a single UTF-8 encoded Unicode character and returns the
rune and its size in bytes. If the encoded rune is invalid, it consumes one byte
and returns unicode.ReplacementChar (U+FFFD) with a size of 1.

ReadRune 会读取单个UTF-8编码的Unicode字符, 并且返回该rune和它的字节大小. 如果编码是无效的，它(读取位置)会前移一个字节并且返回nicode.ReplacementChar(U+FFFD)且size为1.


### <a id="Reader.ReadSlice">func</a> (\*Reader) [ReadSlice](https://golang.org/src/bufio/bufio.go?s=8477:8540#L320)
<pre>func (b *<a href="#Reader">Reader</a>) ReadSlice(delim <a href="/pkg/builtin/#byte">byte</a>) (line []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadSlice reads until the first occurrence of delim in the input,
returning a slice pointing at the bytes in the buffer.
The bytes stop being valid at the next read.
If ReadSlice encounters an error before finding a delimiter,
it returns all the data in the buffer and the error itself (often io.EOF).
ReadSlice fails with error ErrBufferFull if the buffer fills without a delim.
Because the data returned from ReadSlice will be overwritten
by the next I/O operation, most clients should use
ReadBytes or ReadString instead.
ReadSlice returns err != nil if and only if line does not end in delim.

ReadSlice 会读取直到第一次遇到分隔符delim, 返回一个指向缓存中字节的slice. 这些字节只在下一次读取操作之前有效. 如果ReadSlice在遇到分隔符之前碰到了error, 它就会返回缓存中所有的数据和错误本身（经常是io.EOF）. 如果在读取到delim之前缓冲就被写满了，ReadSlice会失败并返回错误ErrBufferFull. 因为ReadSlice返回的数据会被下次的I/O操作覆盖，因此大多数用户会选择使用ReadBytes或者ReadString来代替. 当且仅当ReadSlice方法返回的数据不以delim结尾时，返回的err != nil.


### <a id="Reader.ReadString">func</a> (\*Reader) [ReadString](https://golang.org/src/bufio/bufio.go?s=12557:12612#L466)
<pre>func (b *<a href="#Reader">Reader</a>) ReadString(delim <a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadString reads until the first occurrence of delim in the input,
returning a string containing the data up to and including the delimiter.
If ReadString encounters an error before finding a delimiter,
it returns the data read before the error and the error itself (often io.EOF).
ReadString returns err != nil if and only if the returned data does not end in
delim.
For simple uses, a Scanner may be more convenient.

ReadString 会读取数据直到第一次遇到分隔符delim，返回一个包含已读取到数据的字符串(包括分隔符). 如果ReadString在遇到分隔符之前碰到了error, 它会返回遇到错误之前已读取的数据和这个错误（经常是io.EOF）. 当且仅当ReadString方法返回的数据不以delim结尾时，返回的err != nil. 对于简单的使用, 或许Scanner会更方便.


### <a id="Reader.Reset">func</a> (\*Reader) [Reset](https://golang.org/src/bufio/bufio.go?s=2059:2094#L60)
<pre>func (b *<a href="#Reader">Reader</a>) Reset(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>)</pre>
Reset discards any buffered data, resets all state, and switches
the buffered reader to read from r.

Reset会丢弃缓冲中的数据，重置所有的状态,将b切换到从r中读取数据.


### <a id="Reader.Size">func</a> (\*Reader) [Size](https://golang.org/src/bufio/bufio.go?s=1901:1928#L56)
<pre>func (b *<a href="#Reader">Reader</a>) Size() <a href="/pkg/builtin/#int">int</a></pre>
Size returns the size of the underlying buffer in bytes.

Size 返回 底层buffer的大小.


### <a id="Reader.UnreadByte">func</a> (\*Reader) [UnreadByte](https://golang.org/src/bufio/bufio.go?s=6267:6302#L255)
<pre>func (b *<a href="#Reader">Reader</a>) UnreadByte() <a href="/pkg/builtin/#error">error</a></pre>
UnreadByte unreads the last byte. Only the most recently read byte can be unread.

UnreadByte returns an error if the most recent method called on the
Reader was not a read operation. Notably, Peek is not considered a
read operation.

UnreadByte会撤消已读取的最后一个字节. 只有最近读出的字节才可以被撤销.

如Reader上最近的操作不是read操作, UnreadByte会返回一个error. 值得注意的是，Peek不被视为read操作.

### <a id="Reader.UnreadRune">func</a> (\*Reader) [UnreadRune](https://golang.org/src/bufio/bufio.go?s=7518:7553#L297)
<pre>func (b *<a href="#Reader">Reader</a>) UnreadRune() <a href="/pkg/builtin/#error">error</a></pre>
UnreadRune unreads the last rune. If the most recent method called on
the Reader was not a ReadRune, UnreadRune returns an error. (In this
regard it is stricter than UnreadByte, which will unread the last byte
from any read operation.)

UnreadRune 会撤消已读取的最后一个rune. 如果最近的在buffer上的操作不是ReadRune，则UnreadRune 会返回一个error. (在这个角度上看，这个函数比UnreadByte更严格，UnreadByte会将任意一个read操作的最后一个已读字节撤销.)


### <a id="Reader.WriteTo">func</a> (\*Reader) [WriteTo](https://golang.org/src/bufio/bufio.go?s=12904:12962#L475)
<pre>func (b *<a href="#Reader">Reader</a>) WriteTo(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) (n <a href="/pkg/builtin/#int64">int64</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
WriteTo implements io.WriterTo.
This may make multiple calls to the Read method of the underlying Reader.
If the underlying reader supports the WriteTo method,
this calls the underlying WriteTo without buffering.

WriteTo实现了io.WriterTo. 这可能会多次调用底层Reader的Read方法. 如果底层的reader支持WriteTo方法, 它会调用底层的WriteTo方法, 且不使用缓冲.


## <a id="Scanner">type</a> [Scanner](https://golang.org/src/bufio/scan.go?s=1184:1841#L20)
Scanner provides a convenient interface for reading data such as
a file of newline-delimited lines of text. Successive calls to
the Scan method will step through the 'tokens' of a file, skipping
the bytes between the tokens. The specification of a token is
defined by a split function of type SplitFunc; the default split
function breaks the input into lines with line termination stripped. Split
functions are defined in this package for scanning a file into
lines, bytes, UTF-8-encoded runes, and space-delimited words. The
client may instead provide a custom split function.

Scanning stops unrecoverably at EOF, the first I/O error, or a token too
large to fit in the buffer. When a scan stops, the reader may have
advanced arbitrarily far past the last token. Programs that need more
control over error handling or large tokens, or must run sequential scans
on a reader, should use bufio.Reader instead.

Scanner 提供了方便的读取数据的接口,比如从以换行符分隔的文本文件里读取. 连续调用Scan方法会步进式地从文件中读出tokens,并跳过tokens间的字节. token的格式是由SplitFunc类型的分割函数定义;默认的分割函数会将输入分割成多行,并去掉行尾的结束符. 本包预定义的分割函数可以将文件分割为行,字节,utf8编码的unicode码值和空白符分隔的words. 调用者可以定制自己的分割函数.

Scan操作在遇到EOF,第一个I/O错误,token超出缓冲区容量时会不可恢复地停止. 当扫描停止后，当前读取位置可能会远在最后一个获得的token后面. 如果程序需要更多对错误处理的控制或token很大，或必须在reader上连续扫描, 那么应使用bufio.Reader来代替.

<pre>type Scanner struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Scanner_custom">Example (Custom)</a>
<p>Use a Scanner with a custom split function (built by wrapping ScanWords) to validate 32-bit decimal input.
</p>

使用自定义的分割函数(通过封装ScanWords实现)来验证32位的十进制输入.

```go
package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
    // An artificial input source.
    // 一个手工输入源
	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
    // Create a custom split function by wrapping the existing ScanWords function.
    // 通过封装已有的ScanWords函数来实现自定义分割函数.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
    // Set the split function for the scanning operation.
    // 指定分割函数.
	scanner.Split(split)
    // Validate the input
    // 验证输入
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}
```

output:
```txt
1234
5678
Invalid input: strconv.ParseInt: parsing "1234567901234567890": value out of range
```

</p><a id="example_Scanner_emptyFinalToken">Example (EmptyFinalToken)</a>
<p>Use a Scanner with a custom split function to parse a comma-separated
list with an empty final value.
</p>

使用自定义分割函数解析使用逗号分割的且最后一项是空的序列.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    // Comma-separated list; last entry is empty.
    // 逗号分隔的序列; 最后一项是空的.
	const input = "1,2,3,4,"
	scanner := bufio.NewScanner(strings.NewReader(input))
    // Define a split function that separates on commas.
    // 定义一个使用逗号分割的分割函数.
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
        // but does not trigger an error to be returned from Scan itself.
        //
        // 最后一个传入的token可能是空字符串.
	    // 通过返回bufio.ErrFinalToken来告诉Scan: 这个token之后就没有数据了, 且(该错误)不会触发Scan返回错误.
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	// Scan.
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
```

output:
```txt
"1" "2" "3" "4" ""
```

<a id="example_Scanner_lines">Example (Lines)</a>
<p>The simplest use of a Scanner, to read standard input as a set of lines.
</p>

最简单的Scanner,用于从标准输入中读取多行数据.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
```


<a id="example_Scanner_words">Example (Words)</a>
<p>Use a Scanner to implement a simple word-count utility by scanning the
input as a sequence of space-delimited tokens.
</p>

使用Scanner读取到一系列使用空白符分割的tokens来实现一个简单的单词统计工具.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    // An artificial input source.
    // 手工输入源.
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
    // Set the split function for the scanning operation.
    // 设置分割函数.
	scanner.Split(bufio.ScanWords)
    // Count the words.
    // 统计单词数.
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
}
```

output:
```txt
15
```




### <a id="NewScanner">func</a> [NewScanner](https://golang.org/src/bufio/scan.go?s=3874:3911#L76)
<pre>func NewScanner(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) *<a href="#Scanner">Scanner</a></pre>
NewScanner returns a new Scanner to read from r.
The split function defaults to ScanLines.

NewScanner 创建一个从r中读取数据的Scanner. 默认的分割函数是ScanLines.




### <a id="Scanner.Buffer">func</a> (\*Scanner) [Buffer](https://golang.org/src/bufio/scan.go?s=9061:9106#L252)
<pre>func (s *<a href="#Scanner">Scanner</a>) Buffer(buf []<a href="/pkg/builtin/#byte">byte</a>, max <a href="/pkg/builtin/#int">int</a>)</pre>
Buffer sets the initial buffer to use when scanning and the maximum
size of buffer that may be allocated during scanning. The maximum
token size is the larger of max and cap(buf). If max <= cap(buf),
Scan will use this buffer only and do no allocation.

By default, Scan uses an internal buffer and sets the
maximum token size to MaxScanTokenSize.

Buffer panics if it is called after scanning has started.

Buffer 用于设置扫描时的初始缓冲区和扫描期间可能分配的缓冲区的最大大小. 最大的token大小是max和cap(buf)中的较大者. 如果max<=cap(buf), Scan使用该buf而不会重新分配(缓冲).

默认情况下, Scan会使用内置的缓存区,且将最大token的大小设置为MaxScanTokenSize.

如果在扫描开始后调用, Buffer方法会panic.



### <a id="Scanner.Bytes">func</a> (\*Scanner) [Bytes](https://golang.org/src/bufio/scan.go?s=4372:4404#L95)
<pre>func (s *<a href="#Scanner">Scanner</a>) Bytes() []<a href="/pkg/builtin/#byte">byte</a></pre>
Bytes returns the most recent token generated by a call to Scan.
The underlying array may point to data that will be overwritten
by a subsequent call to Scan. It does no allocation.

Bytes 返回最近一次Scan操作生成的token. 底层数组指向的数据可能会被下一次Scan操作覆盖.它不会分配内存.

<a id="example_Scanner_Bytes">Example</a>
```go
package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
  scanner := bufio.NewScanner(strings.NewReader("gopher"))
  for scanner.Scan() {
    fmt.Println(len(scanner.Bytes()) == 6)
  }
  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
  }
}
```

output:
```txt
true
```

<p>Return the most recent call to Scan as a []byte.
</p>

### <a id="Scanner.Err">func</a> (\*Scanner) [Err](https://golang.org/src/bufio/scan.go?s=4094:4123#L85)
<pre>func (s *<a href="#Scanner">Scanner</a>) Err() <a href="/pkg/builtin/#error">error</a></pre>
Err returns the first non-EOF error that was encountered by the Scanner.

Err 返回Scanner遇到的第一个非EOF的错误.


### <a id="Scanner.Scan">func</a> (\*Scanner) [Scan](https://golang.org/src/bufio/scan.go?s=5721:5750#L124)
<pre>func (s *<a href="#Scanner">Scanner</a>) Scan() <a href="/pkg/builtin/#bool">bool</a></pre>
Scan advances the Scanner to the next token, which will then be
available through the Bytes or Text method. It returns false when the
scan stops, either by reaching the end of the input or an error.
After Scan returns false, the Err method will return any error that
occurred during scanning, except that if it was io.EOF, Err
will return nil.
Scan panics if the split function returns too many empty
tokens without advancing the input. This is a common error mode for
scanners.

Scan 会步进Scanner以获得下一个token, 该token可以通过 Bytes 或 Text 方法获得. 当扫描到达输入的结尾或者遇到错误而停止时,它会返回false. 在Scan方法返回false后,Err方法将返回扫描时遇到的任何错误,除了io.EOF，此时Err会返回nil. 若分割函数返回太多个空token而没有步进输入时，那么它就会panic. 这是scanner的一个常见错误模式.



### <a id="Scanner.Split">func</a> (\*Scanner) [Split](https://golang.org/src/bufio/scan.go?s=9374:9414#L264)
<pre>func (s *<a href="#Scanner">Scanner</a>) Split(split <a href="#SplitFunc">SplitFunc</a>)</pre>
Split sets the split function for the Scanner.
The default split function is ScanLines.

Split panics if it is called after scanning has started.

Split 用于设置Scanner的分割函数,默认是ScanLines.

如果在扫描开始后调用,Split方法会panic.


### <a id="Scanner.Text">func</a> (\*Scanner) [Text](https://golang.org/src/bufio/scan.go?s=4542:4573#L101)
<pre>func (s *<a href="#Scanner">Scanner</a>) Text() <a href="/pkg/builtin/#string">string</a></pre>
Text returns the most recent token generated by a call to Scan
as a newly allocated string holding its bytes.

Text 会获取最近一次调用Scan生成的token,并申请创建一个字符串保存该token.


## <a id="SplitFunc">type</a> [SplitFunc](https://golang.org/src/bufio/scan.go?s=3047:3130#L55)
SplitFunc is the signature of the split function used to tokenize the
input. The arguments are an initial substring of the remaining unprocessed
data and a flag, atEOF, that reports whether the Reader has no more data
to give. The return values are the number of bytes to advance the input
and the next token to return to the user, if any, plus an error, if any.

Scanning stops if the function returns an error, in which case some of
the input may be discarded.

Otherwise, the Scanner advances the input. If the token is not nil,
the Scanner returns it to the user. If the token is nil, the
Scanner reads more data and continues scanning; if there is no more
data--if atEOF was true--the Scanner returns. If the data does not
yet hold a complete token, for instance if it has no newline while
scanning lines, a SplitFunc can return (0, nil, nil) to signal the
Scanner to read more data into the slice and try again with a
longer slice starting at the same point in the input.

The function is never called with an empty data slice unless atEOF
is true. If atEOF is true, however, data may be non-empty and,
as always, holds unprocessed text.


SplitFunc 是用于对输入进行分词的分割函数的签名. 总共有两个参数:data是原始slice中剩余未处理数据的部分,atEOF标记表示是否Reader接口已经不能提供更多的数据. 返回值是解析位置前进的字节数，将要返回给调用者的token切片，以及可能遇到的错误.

如果函数返回error, 扫描将停止，在这种情况下，输入可能会被丢弃.

否则Scanner会步进输入. 如果token不为nil, Scanner会将它返回给用户; 如果是nil, Scanner会继续扫描以读取更多数据, 如果没有数据了, Scanner会返回值为true的atEOF. 如果数据不是一个完整的token，例如需要一整行数据但data里没有换行符, SplitFunc可以返回(0, nil, nil)来表示 Scanner 需要读取更多的数据到切片中并使用更长的切片再在相同的位置重新读取.

除非atEOF为真，否则永远不会在空切片data上调用本函数.然而，如果atEOF为真，data却可能是非空的, 与往常一样会包含未处理的文本.

<pre>type SplitFunc func(data []<a href="/pkg/builtin/#byte">byte</a>, atEOF <a href="/pkg/builtin/#bool">bool</a>) (advance <a href="/pkg/builtin/#int">int</a>, token []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>











## <a id="Writer">type</a> [Writer](https://golang.org/src/bufio/bufio.go?s=14169:14238#L534)
Writer implements buffering for an io.Writer object.
If an error occurs writing to a Writer, no more data will be
accepted and all subsequent writes, and Flush, will return the error.
After all data has been written, the client should call the
Flush method to guarantee all data has been forwarded to
the underlying io.Writer.

Writer 实现了带buffer机制的io.Writer对象. 如果在写数据到Writer时出现错误，那么将不会再有数据被写入, 而且所有随后的写操作和Flush方法都会返回error. 当所有数据被写入后, 使用者应调用Flush方法以确保所有数据都已写入底层的io.Writer.

<pre>type Writer struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Writer">Example</a>
```go
package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  w := bufio.NewWriter(os.Stdout)
  fmt.Fprint(w, "Hello, ")
  fmt.Fprint(w, "world!")
  w.Flush() // Don't forget to flush! // 不要忘记调用Flush方法!
}
```

output:
```txt
Hello, world!
```




### <a id="NewWriter">func</a> [NewWriter](https://golang.org/src/bufio/bufio.go?s=14753:14788#L560)
<pre>func NewWriter(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) *<a href="#Writer">Writer</a></pre>
NewWriter returns a new Writer whose buffer has the default size.

NewWriter 创建一个带默认大小的缓冲区的Writer.


### <a id="NewWriterSize">func</a> [NewWriterSize](https://golang.org/src/bufio/bufio.go?s=14434:14483#L544)
<pre>func NewWriterSize(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, size <a href="/pkg/builtin/#int">int</a>) *<a href="#Writer">Writer</a></pre>
NewWriterSize returns a new Writer whose buffer has at least the specified
size. If the argument io.Writer is already a Writer with large enough
size, it returns the underlying Writer.

NewWriterSize 会创建一个Writer,其拥有至少是size大小的缓冲区. 如果参数w已经是带有足够大缓冲的Writer了，那么它会返回底层的Writer.




### <a id="Writer.Available">func</a> (\*Writer) [Available](https://golang.org/src/bufio/bufio.go?s=15592:15624#L600)
<pre>func (b *<a href="#Writer">Writer</a>) Available() <a href="/pkg/builtin/#int">int</a></pre>
Available returns how many bytes are unused in the buffer.

Available 返回buffer中未使用空间的字节数.


### <a id="Writer.Buffered">func</a> (\*Writer) [Buffered](https://golang.org/src/bufio/bufio.go?s=15742:15773#L603)
<pre>func (b *<a href="#Writer">Writer</a>) Buffered() <a href="/pkg/builtin/#int">int</a></pre>
Buffered returns the number of bytes that have been written into the current buffer.

Buffered 返回已写到当前buffer中的字节数.


### <a id="Writer.Flush">func</a> (\*Writer) [Flush](https://golang.org/src/bufio/bufio.go?s=15189:15219#L576)
<pre>func (b *<a href="#Writer">Writer</a>) Flush() <a href="/pkg/builtin/#error">error</a></pre>
Flush writes any buffered data to the underlying io.Writer.

Flush 会将所有缓冲的数据写入底层的io.Writer.


### <a id="Writer.ReadFrom">func</a> (\*Writer) [ReadFrom](https://golang.org/src/bufio/bufio.go?s=17892:17951#L700)
<pre>func (b *<a href="#Writer">Writer</a>) ReadFrom(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (n <a href="/pkg/builtin/#int64">int64</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadFrom implements io.ReaderFrom. If the underlying writer
supports the ReadFrom method, and b has no buffered data yet,
this calls the underlying ReadFrom without buffering.

ReadFrom 实现了 io.ReaderFrom 接口. 如果底层的writer支持ReadFrom方法, 且b还未缓存过数据, 那么本方法会使用底层的ReadFrom且不缓冲数据.


### <a id="Writer.Reset">func</a> (\*Writer) [Reset](https://golang.org/src/bufio/bufio.go?s=15053:15088#L569)
<pre>func (b *<a href="#Writer">Writer</a>) Reset(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>)</pre>
Reset discards any unflushed buffered data, clears any error, and
resets b to write its output to w.

Reset 会丢弃缓冲中未刷新的数据，清除所有错误，将b的写入定向到w.


### <a id="Writer.Size">func</a> (\*Writer) [Size](https://golang.org/src/bufio/bufio.go?s=14895:14922#L565)
<pre>func (b *<a href="#Writer">Writer</a>) Size() <a href="/pkg/builtin/#int">int</a></pre>
Size returns the size of the underlying buffer in bytes.

Size 返回底层缓冲的大小(字节).


### <a id="Writer.Write">func</a> (\*Writer) [Write](https://golang.org/src/bufio/bufio.go?s=15966:16018#L609)
<pre>func (b *<a href="#Writer">Writer</a>) Write(p []<a href="/pkg/builtin/#byte">byte</a>) (nn <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Write writes the contents of p into the buffer.
It returns the number of bytes written.
If nn < len(p), it also returns an error explaining
why the write is short.

Write会将p的内容写入缓冲,并返回写入的字节数.如果nn < len(p), 它会返回error来解释为什么写入的数据会短缺.


### <a id="Writer.WriteByte">func</a> (\*Writer) [WriteByte](https://golang.org/src/bufio/bufio.go?s=16442:16482#L634)
<pre>func (b *<a href="#Writer">Writer</a>) WriteByte(c <a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
WriteByte writes a single byte.

WriteByte 会写入单个字节.


### <a id="Writer.WriteRune">func</a> (\*Writer) [WriteRune](https://golang.org/src/bufio/bufio.go?s=16728:16784#L648)
<pre>func (b *<a href="#Writer">Writer</a>) WriteRune(r <a href="/pkg/builtin/#rune">rune</a>) (size <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
WriteRune writes a single Unicode code point, returning
the number of bytes written and any error.

WriteRune 会写入单个的Unicode码点,并返回写的字节数和遇到的错误.


### <a id="Writer.WriteString">func</a> (\*Writer) [WriteString](https://golang.org/src/bufio/bufio.go?s=17416:17467#L679)
<pre>func (b *<a href="#Writer">Writer</a>) WriteString(s <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
WriteString writes a string.
It returns the number of bytes written.
If the count is less than len(s), it also returns an error explaining
why the write is short.

WriteString 会写入一个字符串,并返回写入的字节数. 如果写入的字节数比len(s)少，它会返回error来解释为什么写入的数据会短缺.





