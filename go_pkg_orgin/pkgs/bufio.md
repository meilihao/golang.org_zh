

# bufio
`import "bufio"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer
object, creating another object (Reader or Writer) that also implements
the interface but provides buffering and some help for textual I/O.




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

<pre>const (
    <span class="comment">// MaxScanTokenSize is the maximum size used to buffer a token</span>
    <span class="comment">// unless the user provides an explicit buffer with Scanner.Buffer.</span>
    <span class="comment">// The actual maximum token size may be smaller as the buffer</span>
    <span class="comment">// may need to include, for instance, a newline.</span>
    <span id="MaxScanTokenSize">MaxScanTokenSize</span> = 64 * 1024
)</pre>

## <a id="pkg-variables">Variables</a>

<pre>var (
    <span id="ErrInvalidUnreadByte">ErrInvalidUnreadByte</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;bufio: invalid use of UnreadByte&#34;)
    <span id="ErrInvalidUnreadRune">ErrInvalidUnreadRune</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;bufio: invalid use of UnreadRune&#34;)
    <span id="ErrBufferFull">ErrBufferFull</span>        = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;bufio: buffer full&#34;)
    <span id="ErrNegativeCount">ErrNegativeCount</span>     = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;bufio: negative count&#34;)
)</pre>Errors returned by Scanner.


<pre>var (
    <span id="ErrTooLong">ErrTooLong</span>         = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;bufio.Scanner: token too long&#34;)
    <span id="ErrNegativeAdvance">ErrNegativeAdvance</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;bufio.Scanner: SplitFunc returns negative advance count&#34;)
    <span id="ErrAdvanceTooFar">ErrAdvanceTooFar</span>   = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;bufio.Scanner: SplitFunc returns advance count beyond input&#34;)
)</pre>ErrFinalToken is a special sentinel error value. It is intended to be
returned by a Split function to indicate that the token being delivered
with the error is the last token and scanning should stop after this one.
After ErrFinalToken is received by Scan, scanning stops with no error.
The value is useful to stop processing early or when it is necessary to
deliver a final empty token. One could achieve the same behavior
with a custom error value but providing one here is tidier.
See the emptyFinalToken example for a use of this value.


<pre>var <span id="ErrFinalToken">ErrFinalToken</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;final token&#34;)</pre>

## <a id="ScanBytes">func</a> [ScanBytes](https://golang.org/src/bufio/scan.go?s=9596:9674#L274)
<pre>func ScanBytes(data []<a href="/pkg/builtin/#byte">byte</a>, atEOF <a href="/pkg/builtin/#bool">bool</a>) (advance <a href="/pkg/builtin/#int">int</a>, token []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ScanBytes is a split function for a Scanner that returns each byte as a token.



## <a id="ScanLines">func</a> [ScanLines](https://golang.org/src/bufio/scan.go?s=11802:11880#L335)
<pre>func ScanLines(data []<a href="/pkg/builtin/#byte">byte</a>, atEOF <a href="/pkg/builtin/#bool">bool</a>) (advance <a href="/pkg/builtin/#int">int</a>, token []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ScanLines is a split function for a Scanner that returns each line of
text, stripped of any trailing end-of-line marker. The returned line may
be empty. The end-of-line marker is one optional carriage return followed
by one mandatory newline. In regular expression notation, it is `\r?\n`.
The last non-empty line of input will be returned even if it has no
newline.



## <a id="ScanRunes">func</a> [ScanRunes](https://golang.org/src/bufio/scan.go?s=10243:10321#L289)
<pre>func ScanRunes(data []<a href="/pkg/builtin/#byte">byte</a>, atEOF <a href="/pkg/builtin/#bool">bool</a>) (advance <a href="/pkg/builtin/#int">int</a>, token []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ScanRunes is a split function for a Scanner that returns each
UTF-8-encoded rune as a token. The sequence of runes returned is
equivalent to that from a range loop over the input as a string, which
means that erroneous UTF-8 encodings translate to U+FFFD = "\xef\xbf\xbd".
Because of the Scan interface, this makes it impossible for the client to
distinguish correctly encoded replacement runes from encoding errors.



## <a id="ScanWords">func</a> [ScanWords](https://golang.org/src/bufio/scan.go?s=13096:13174#L380)
<pre>func ScanWords(data []<a href="/pkg/builtin/#byte">byte</a>, atEOF <a href="/pkg/builtin/#bool">bool</a>) (advance <a href="/pkg/builtin/#int">int</a>, token []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ScanWords is a split function for a Scanner that returns each
space-separated word of text, with surrounding spaces deleted. It will
never return an empty string. The definition of space is set by
unicode.IsSpace.





## <a id="ReadWriter">type</a> [ReadWriter](https://golang.org/src/bufio/bufio.go?s=18720:18764#L745)
ReadWriter stores pointers to a Reader and a Writer.
It implements io.ReadWriter.


<pre>type ReadWriter struct {
    *<a href="#Reader">Reader</a>
    *<a href="#Writer">Writer</a>
}
</pre>









### <a id="NewReadWriter">func</a> [NewReadWriter](https://golang.org/src/bufio/bufio.go?s=18838:18890#L751)
<pre>func NewReadWriter(r *<a href="#Reader">Reader</a>, w *<a href="#Writer">Writer</a>) *<a href="#ReadWriter">ReadWriter</a></pre>
NewReadWriter allocates a new ReadWriter that dispatches to r and w.






## <a id="Reader">type</a> [Reader](https://golang.org/src/bufio/bufio.go?s=829:1151#L21)
Reader implements buffering for an io.Reader object.


<pre>type Reader struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewReader">func</a> [NewReader](https://golang.org/src/bufio/bufio.go?s=1757:1793#L51)
<pre>func NewReader(rd <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) *<a href="#Reader">Reader</a></pre>
NewReader returns a new Reader whose buffer has the default size.




### <a id="NewReaderSize">func</a> [NewReaderSize](https://golang.org/src/bufio/bufio.go?s=1414:1464#L36)
<pre>func NewReaderSize(rd <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, size <a href="/pkg/builtin/#int">int</a>) *<a href="#Reader">Reader</a></pre>
NewReaderSize returns a new Reader whose buffer has at least the specified
size. If the argument io.Reader is already a Reader with large enough
size, it returns the underlying Reader.






### <a id="Reader.Buffered">func</a> (\*Reader) [Buffered](https://golang.org/src/bufio/bufio.go?s=7796:7827#L308)
<pre>func (b *<a href="#Reader">Reader</a>) Buffered() <a href="/pkg/builtin/#int">int</a></pre>
Buffered returns the number of bytes that can be read from the current buffer.




### <a id="Reader.Discard">func</a> (\*Reader) [Discard](https://golang.org/src/bufio/bufio.go?s=4163:4221#L153)
<pre>func (b *<a href="#Reader">Reader</a>) Discard(n <a href="/pkg/builtin/#int">int</a>) (discarded <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Discard skips the next n bytes, returning the number of bytes discarded.

If Discard skips fewer than n bytes, it also returns an error.
If 0 <= n <= b.Buffered(), Discard is guaranteed to succeed without
reading from the underlying io.Reader.




### <a id="Reader.Peek">func</a> (\*Reader) [Peek](https://golang.org/src/bufio/bufio.go?s=3366:3410#L119)
<pre>func (b *<a href="#Reader">Reader</a>) Peek(n <a href="/pkg/builtin/#int">int</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Peek returns the next n bytes without advancing the reader. The bytes stop
being valid at the next read call. If Peek returns fewer than n bytes, it
also returns an error explaining why the read is short. The error is
ErrBufferFull if n is larger than b's buffer size.

Calling Peek prevents a UnreadByte or UnreadRune call from succeeding
until the next read operation.




### <a id="Reader.Read">func</a> (\*Reader) [Read](https://golang.org/src/bufio/bufio.go?s=4864:4914#L187)
<pre>func (b *<a href="#Reader">Reader</a>) Read(p []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Read reads data into p.
It returns the number of bytes read into p.
The bytes are taken from at most one Read on the underlying Reader,
hence n may be less than len(p).
To read exactly len(p) bytes, use io.ReadFull(b, p).
At EOF, the count will be zero and err will be io.EOF.




### <a id="Reader.ReadByte">func</a> (\*Reader) [ReadByte](https://golang.org/src/bufio/bufio.go?s=5791:5832#L236)
<pre>func (b *<a href="#Reader">Reader</a>) ReadByte() (<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadByte reads and returns a single byte.
If no byte is available, returns an error.




### <a id="Reader.ReadBytes">func</a> (\*Reader) [ReadBytes](https://golang.org/src/bufio/bufio.go?s=11350:11404#L419)
<pre>func (b *<a href="#Reader">Reader</a>) ReadBytes(delim <a href="/pkg/builtin/#byte">byte</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadBytes reads until the first occurrence of delim in the input,
returning a slice containing the data up to and including the delimiter.
If ReadBytes encounters an error before finding a delimiter,
it returns the data read before the error and the error itself (often io.EOF).
ReadBytes returns err != nil if and only if the returned data does not end in
delim.
For simple uses, a Scanner may be more convenient.




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




### <a id="Reader.ReadRune">func</a> (\*Reader) [ReadRune](https://golang.org/src/bufio/bufio.go?s=6786:6843#L275)
<pre>func (b *<a href="#Reader">Reader</a>) ReadRune() (r <a href="/pkg/builtin/#rune">rune</a>, size <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadRune reads a single UTF-8 encoded Unicode character and returns the
rune and its size in bytes. If the encoded rune is invalid, it consumes one byte
and returns unicode.ReplacementChar (U+FFFD) with a size of 1.




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




### <a id="Reader.ReadString">func</a> (\*Reader) [ReadString](https://golang.org/src/bufio/bufio.go?s=12557:12612#L466)
<pre>func (b *<a href="#Reader">Reader</a>) ReadString(delim <a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadString reads until the first occurrence of delim in the input,
returning a string containing the data up to and including the delimiter.
If ReadString encounters an error before finding a delimiter,
it returns the data read before the error and the error itself (often io.EOF).
ReadString returns err != nil if and only if the returned data does not end in
delim.
For simple uses, a Scanner may be more convenient.




### <a id="Reader.Reset">func</a> (\*Reader) [Reset](https://golang.org/src/bufio/bufio.go?s=2059:2094#L60)
<pre>func (b *<a href="#Reader">Reader</a>) Reset(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>)</pre>
Reset discards any buffered data, resets all state, and switches
the buffered reader to read from r.




### <a id="Reader.Size">func</a> (\*Reader) [Size](https://golang.org/src/bufio/bufio.go?s=1901:1928#L56)
<pre>func (b *<a href="#Reader">Reader</a>) Size() <a href="/pkg/builtin/#int">int</a></pre>
Size returns the size of the underlying buffer in bytes.




### <a id="Reader.UnreadByte">func</a> (\*Reader) [UnreadByte](https://golang.org/src/bufio/bufio.go?s=6267:6302#L255)
<pre>func (b *<a href="#Reader">Reader</a>) UnreadByte() <a href="/pkg/builtin/#error">error</a></pre>
UnreadByte unreads the last byte. Only the most recently read byte can be unread.

UnreadByte returns an error if the most recent method called on the
Reader was not a read operation. Notably, Peek is not considered a
read operation.




### <a id="Reader.UnreadRune">func</a> (\*Reader) [UnreadRune](https://golang.org/src/bufio/bufio.go?s=7518:7553#L297)
<pre>func (b *<a href="#Reader">Reader</a>) UnreadRune() <a href="/pkg/builtin/#error">error</a></pre>
UnreadRune unreads the last rune. If the most recent method called on
the Reader was not a ReadRune, UnreadRune returns an error. (In this
regard it is stricter than UnreadByte, which will unread the last byte
from any read operation.)




### <a id="Reader.WriteTo">func</a> (\*Reader) [WriteTo](https://golang.org/src/bufio/bufio.go?s=12904:12962#L475)
<pre>func (b *<a href="#Reader">Reader</a>) WriteTo(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) (n <a href="/pkg/builtin/#int64">int64</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
WriteTo implements io.WriterTo.
This may make multiple calls to the Read method of the underlying Reader.
If the underlying reader supports the WriteTo method,
this calls the underlying WriteTo without buffering.




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


<pre>type Scanner struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Scanner_custom">Example (Custom)</a>
```go
```

output:
```txt
```
<p>Use a Scanner with a custom split function (built by wrapping ScanWords) to validate
32-bit decimal input.
</p><a id="example_Scanner_emptyFinalToken">Example (EmptyFinalToken)</a>
```go
```

output:
```txt
```
<p>Use a Scanner with a custom split function to parse a comma-separated
list with an empty final value.
</p><a id="example_Scanner_lines">Example (Lines)</a>
```go
```

output:
```txt
```
<p>The simplest use of a Scanner, to read standard input as a set of lines.
</p><a id="example_Scanner_words">Example (Words)</a>
```go
```

output:
```txt
```
<p>Use a Scanner to implement a simple word-count utility by scanning the
input as a sequence of space-delimited tokens.
</p>



### <a id="NewScanner">func</a> [NewScanner](https://golang.org/src/bufio/scan.go?s=3874:3911#L76)
<pre>func NewScanner(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) *<a href="#Scanner">Scanner</a></pre>
NewScanner returns a new Scanner to read from r.
The split function defaults to ScanLines.






### <a id="Scanner.Buffer">func</a> (\*Scanner) [Buffer](https://golang.org/src/bufio/scan.go?s=9061:9106#L252)
<pre>func (s *<a href="#Scanner">Scanner</a>) Buffer(buf []<a href="/pkg/builtin/#byte">byte</a>, max <a href="/pkg/builtin/#int">int</a>)</pre>
Buffer sets the initial buffer to use when scanning and the maximum
size of buffer that may be allocated during scanning. The maximum
token size is the larger of max and cap(buf). If max <= cap(buf),
Scan will use this buffer only and do no allocation.

By default, Scan uses an internal buffer and sets the
maximum token size to MaxScanTokenSize.

Buffer panics if it is called after scanning has started.




### <a id="Scanner.Bytes">func</a> (\*Scanner) [Bytes](https://golang.org/src/bufio/scan.go?s=4372:4404#L95)
<pre>func (s *<a href="#Scanner">Scanner</a>) Bytes() []<a href="/pkg/builtin/#byte">byte</a></pre>
Bytes returns the most recent token generated by a call to Scan.
The underlying array may point to data that will be overwritten
by a subsequent call to Scan. It does no allocation.


<a id="example_Scanner_Bytes">Example</a>
```go
```

output:
```txt
```
<p>Return the most recent call to Scan as a []byte.
</p>

### <a id="Scanner.Err">func</a> (\*Scanner) [Err](https://golang.org/src/bufio/scan.go?s=4094:4123#L85)
<pre>func (s *<a href="#Scanner">Scanner</a>) Err() <a href="/pkg/builtin/#error">error</a></pre>
Err returns the first non-EOF error that was encountered by the Scanner.




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




### <a id="Scanner.Split">func</a> (\*Scanner) [Split](https://golang.org/src/bufio/scan.go?s=9374:9414#L264)
<pre>func (s *<a href="#Scanner">Scanner</a>) Split(split <a href="#SplitFunc">SplitFunc</a>)</pre>
Split sets the split function for the Scanner.
The default split function is ScanLines.

Split panics if it is called after scanning has started.




### <a id="Scanner.Text">func</a> (\*Scanner) [Text](https://golang.org/src/bufio/scan.go?s=4542:4573#L101)
<pre>func (s *<a href="#Scanner">Scanner</a>) Text() <a href="/pkg/builtin/#string">string</a></pre>
Text returns the most recent token generated by a call to Scan
as a newly allocated string holding its bytes.




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


<pre>type SplitFunc func(data []<a href="/pkg/builtin/#byte">byte</a>, atEOF <a href="/pkg/builtin/#bool">bool</a>) (advance <a href="/pkg/builtin/#int">int</a>, token []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>











## <a id="Writer">type</a> [Writer](https://golang.org/src/bufio/bufio.go?s=14169:14238#L534)
Writer implements buffering for an io.Writer object.
If an error occurs writing to a Writer, no more data will be
accepted and all subsequent writes, and Flush, will return the error.
After all data has been written, the client should call the
Flush method to guarantee all data has been forwarded to
the underlying io.Writer.


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




### <a id="NewWriter">func</a> [NewWriter](https://golang.org/src/bufio/bufio.go?s=14753:14788#L560)
<pre>func NewWriter(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) *<a href="#Writer">Writer</a></pre>
NewWriter returns a new Writer whose buffer has the default size.




### <a id="NewWriterSize">func</a> [NewWriterSize](https://golang.org/src/bufio/bufio.go?s=14434:14483#L544)
<pre>func NewWriterSize(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, size <a href="/pkg/builtin/#int">int</a>) *<a href="#Writer">Writer</a></pre>
NewWriterSize returns a new Writer whose buffer has at least the specified
size. If the argument io.Writer is already a Writer with large enough
size, it returns the underlying Writer.






### <a id="Writer.Available">func</a> (\*Writer) [Available](https://golang.org/src/bufio/bufio.go?s=15592:15624#L600)
<pre>func (b *<a href="#Writer">Writer</a>) Available() <a href="/pkg/builtin/#int">int</a></pre>
Available returns how many bytes are unused in the buffer.




### <a id="Writer.Buffered">func</a> (\*Writer) [Buffered](https://golang.org/src/bufio/bufio.go?s=15742:15773#L603)
<pre>func (b *<a href="#Writer">Writer</a>) Buffered() <a href="/pkg/builtin/#int">int</a></pre>
Buffered returns the number of bytes that have been written into the current buffer.




### <a id="Writer.Flush">func</a> (\*Writer) [Flush](https://golang.org/src/bufio/bufio.go?s=15189:15219#L576)
<pre>func (b *<a href="#Writer">Writer</a>) Flush() <a href="/pkg/builtin/#error">error</a></pre>
Flush writes any buffered data to the underlying io.Writer.




### <a id="Writer.ReadFrom">func</a> (\*Writer) [ReadFrom](https://golang.org/src/bufio/bufio.go?s=17892:17951#L700)
<pre>func (b *<a href="#Writer">Writer</a>) ReadFrom(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (n <a href="/pkg/builtin/#int64">int64</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadFrom implements io.ReaderFrom. If the underlying writer
supports the ReadFrom method, and b has no buffered data yet,
this calls the underlying ReadFrom without buffering.




### <a id="Writer.Reset">func</a> (\*Writer) [Reset](https://golang.org/src/bufio/bufio.go?s=15053:15088#L569)
<pre>func (b *<a href="#Writer">Writer</a>) Reset(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>)</pre>
Reset discards any unflushed buffered data, clears any error, and
resets b to write its output to w.




### <a id="Writer.Size">func</a> (\*Writer) [Size](https://golang.org/src/bufio/bufio.go?s=14895:14922#L565)
<pre>func (b *<a href="#Writer">Writer</a>) Size() <a href="/pkg/builtin/#int">int</a></pre>
Size returns the size of the underlying buffer in bytes.




### <a id="Writer.Write">func</a> (\*Writer) [Write](https://golang.org/src/bufio/bufio.go?s=15966:16018#L609)
<pre>func (b *<a href="#Writer">Writer</a>) Write(p []<a href="/pkg/builtin/#byte">byte</a>) (nn <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Write writes the contents of p into the buffer.
It returns the number of bytes written.
If nn < len(p), it also returns an error explaining
why the write is short.




### <a id="Writer.WriteByte">func</a> (\*Writer) [WriteByte](https://golang.org/src/bufio/bufio.go?s=16442:16482#L634)
<pre>func (b *<a href="#Writer">Writer</a>) WriteByte(c <a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
WriteByte writes a single byte.




### <a id="Writer.WriteRune">func</a> (\*Writer) [WriteRune](https://golang.org/src/bufio/bufio.go?s=16728:16784#L648)
<pre>func (b *<a href="#Writer">Writer</a>) WriteRune(r <a href="/pkg/builtin/#rune">rune</a>) (size <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
WriteRune writes a single Unicode code point, returning
the number of bytes written and any error.




### <a id="Writer.WriteString">func</a> (\*Writer) [WriteString](https://golang.org/src/bufio/bufio.go?s=17416:17467#L679)
<pre>func (b *<a href="#Writer">Writer</a>) WriteString(s <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
WriteString writes a string.
It returns the number of bytes written.
If the count is less than len(s), it also returns an error explaining
why the write is short.







