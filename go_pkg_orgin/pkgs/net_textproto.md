

# textproto
`import "net/textproto"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package textproto implements generic support for text-based request/response
protocols in the style of HTTP, NNTP, and SMTP.

The package provides:

Error, which represents a numeric error response from
a server.

Pipeline, to manage pipelined requests and responses
in a client.

Reader, to read numeric response code lines,
key: value headers, lines wrapped with leading spaces
on continuation lines, and whole text blocks ending
with a dot on a line by itself.

Writer, to write dot-encoded text blocks.

Conn, a convenient packaging of Reader, Writer, and Pipeline for use
with a single network connection.




## <a id="pkg-index">Index</a>
* [func CanonicalMIMEHeaderKey(s string) string](#CanonicalMIMEHeaderKey)
* [func TrimBytes(b []byte) []byte](#TrimBytes)
* [func TrimString(s string) string](#TrimString)
* [type Conn](#Conn)
  * [func Dial(network, addr string) (*Conn, error)](#Dial)
  * [func NewConn(conn io.ReadWriteCloser) *Conn](#NewConn)
  * [func (c *Conn) Close() error](#Conn.Close)
  * [func (c *Conn) Cmd(format string, args ...interface{}) (id uint, err error)](#Conn.Cmd)
* [type Error](#Error)
  * [func (e *Error) Error() string](#Error.Error)
* [type MIMEHeader](#MIMEHeader)
  * [func (h MIMEHeader) Add(key, value string)](#MIMEHeader.Add)
  * [func (h MIMEHeader) Del(key string)](#MIMEHeader.Del)
  * [func (h MIMEHeader) Get(key string) string](#MIMEHeader.Get)
  * [func (h MIMEHeader) Set(key, value string)](#MIMEHeader.Set)
* [type Pipeline](#Pipeline)
  * [func (p *Pipeline) EndRequest(id uint)](#Pipeline.EndRequest)
  * [func (p *Pipeline) EndResponse(id uint)](#Pipeline.EndResponse)
  * [func (p *Pipeline) Next() uint](#Pipeline.Next)
  * [func (p *Pipeline) StartRequest(id uint)](#Pipeline.StartRequest)
  * [func (p *Pipeline) StartResponse(id uint)](#Pipeline.StartResponse)
* [type ProtocolError](#ProtocolError)
  * [func (p ProtocolError) Error() string](#ProtocolError.Error)
* [type Reader](#Reader)
  * [func NewReader(r *bufio.Reader) *Reader](#NewReader)
  * [func (r *Reader) DotReader() io.Reader](#Reader.DotReader)
  * [func (r *Reader) ReadCodeLine(expectCode int) (code int, message string, err error)](#Reader.ReadCodeLine)
  * [func (r *Reader) ReadContinuedLine() (string, error)](#Reader.ReadContinuedLine)
  * [func (r *Reader) ReadContinuedLineBytes() ([]byte, error)](#Reader.ReadContinuedLineBytes)
  * [func (r *Reader) ReadDotBytes() ([]byte, error)](#Reader.ReadDotBytes)
  * [func (r *Reader) ReadDotLines() ([]string, error)](#Reader.ReadDotLines)
  * [func (r *Reader) ReadLine() (string, error)](#Reader.ReadLine)
  * [func (r *Reader) ReadLineBytes() ([]byte, error)](#Reader.ReadLineBytes)
  * [func (r *Reader) ReadMIMEHeader() (MIMEHeader, error)](#Reader.ReadMIMEHeader)
  * [func (r *Reader) ReadResponse(expectCode int) (code int, message string, err error)](#Reader.ReadResponse)
* [type Writer](#Writer)
  * [func NewWriter(w *bufio.Writer) *Writer](#NewWriter)
  * [func (w *Writer) DotWriter() io.WriteCloser](#Writer.DotWriter)
  * [func (w *Writer) PrintfLine(format string, args ...interface{}) error](#Writer.PrintfLine)




#### <a id="pkg-files">Package files</a>
[header.go](https://golang.org/src/net/textproto/header.go) [pipeline.go](https://golang.org/src/net/textproto/pipeline.go) [reader.go](https://golang.org/src/net/textproto/reader.go) [textproto.go](https://golang.org/src/net/textproto/textproto.go) [writer.go](https://golang.org/src/net/textproto/writer.go) 






## <a id="CanonicalMIMEHeaderKey">func</a> [CanonicalMIMEHeaderKey](https://golang.org/src/net/textproto/reader.go?s=15120:15164#L559)
<pre>func CanonicalMIMEHeaderKey(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
CanonicalMIMEHeaderKey returns the canonical format of the
MIME header key s. The canonicalization converts the first
letter and any letter following a hyphen to upper case;
the rest are converted to lowercase. For example, the
canonical key for "accept-encoding" is "Accept-Encoding".
MIME header keys are assumed to be ASCII only.
If s contains a space or invalid header field bytes, it is
returned without modifications.



## <a id="TrimBytes">func</a> [TrimBytes](https://golang.org/src/net/textproto/textproto.go?s=3409:3440#L127)
<pre>func TrimBytes(b []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
TrimBytes returns b without leading and trailing ASCII space.



## <a id="TrimString">func</a> [TrimString](https://golang.org/src/net/textproto/textproto.go?s=3172:3204#L116)
<pre>func TrimString(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
TrimString returns s without leading and trailing ASCII space.





## <a id="Conn">type</a> [Conn](https://golang.org/src/net/textproto/textproto.go?s=1556:1627#L48)
A Conn represents a textual network protocol connection.
It consists of a Reader and Writer to manage I/O
and a Pipeline to sequence concurrent requests on the connection.
These embedded types carry methods with them;
see the documentation of those types for details.


<pre>type Conn struct {
    <a href="#Reader">Reader</a>
    <a href="#Writer">Writer</a>
    <a href="#Pipeline">Pipeline</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="Dial">func</a> [Dial](https://golang.org/src/net/textproto/textproto.go?s=2064:2110#L71)
<pre>func Dial(network, addr <a href="/pkg/builtin/#string">string</a>) (*<a href="#Conn">Conn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Dial connects to the given address on the given network using net.Dial
and then returns a new Conn for the connection.




### <a id="NewConn">func</a> [NewConn](https://golang.org/src/net/textproto/textproto.go?s=1679:1722#L56)
<pre>func NewConn(conn <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadWriteCloser">ReadWriteCloser</a>) *<a href="#Conn">Conn</a></pre>
NewConn returns a new Conn using conn for I/O.






### <a id="Conn.Close">func</a> (\*Conn) [Close](https://golang.org/src/net/textproto/textproto.go?s=1882:1910#L65)
<pre>func (c *<a href="#Conn">Conn</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the connection.




### <a id="Conn.Cmd">func</a> (\*Conn) [Cmd](https://golang.org/src/net/textproto/textproto.go?s=2883:2958#L104)
<pre>func (c *<a href="#Conn">Conn</a>) Cmd(format <a href="/pkg/builtin/#string">string</a>, args ...interface{}) (id <a href="/pkg/builtin/#uint">uint</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Cmd is a convenience method that sends a command after
waiting its turn in the pipeline. The command text is the
result of formatting format with args and appending \r\n.
Cmd returns the id of the command, for use with StartResponse and EndResponse.

For example, a client might run a HELP command that returns a dot-body
by using:


	id, err := c.Cmd("HELP")
	if err != nil {
		return nil, err
	}
	
	c.StartResponse(id)
	defer c.EndResponse(id)
	
	if _, _, err = c.ReadCodeLine(110); err != nil {
		return nil, err
	}
	text, err := c.ReadDotBytes()
	if err != nil {
		return nil, err
	}
	return c.ReadCodeLine(250)




## <a id="Error">type</a> [Error](https://golang.org/src/net/textproto/textproto.go?s=951:995#L26)
An Error represents a numeric error response from a server.


<pre>type Error struct {
<span id="Error.Code"></span>    Code <a href="/pkg/builtin/#int">int</a>
<span id="Error.Msg"></span>    Msg  <a href="/pkg/builtin/#string">string</a>
}
</pre>











### <a id="Error.Error">func</a> (\*Error) [Error](https://golang.org/src/net/textproto/textproto.go?s=997:1027#L31)
<pre>func (e *<a href="#Error">Error</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="MIMEHeader">type</a> [MIMEHeader](https://golang.org/src/net/textproto/header.go?s=261:296#L1)
A MIMEHeader represents a MIME-style header mapping
keys to sets of values.


<pre>type MIMEHeader map[<a href="/pkg/builtin/#string">string</a>][]<a href="/pkg/builtin/#string">string</a></pre>











### <a id="MIMEHeader.Add">func</a> (MIMEHeader) [Add](https://golang.org/src/net/textproto/header.go?s=403:445#L3)
<pre>func (h <a href="#MIMEHeader">MIMEHeader</a>) Add(key, value <a href="/pkg/builtin/#string">string</a>)</pre>
Add adds the key, value pair to the header.
It appends to any existing values associated with key.




### <a id="MIMEHeader.Del">func</a> (MIMEHeader) [Del](https://golang.org/src/net/textproto/header.go?s=1281:1316#L33)
<pre>func (h <a href="#MIMEHeader">MIMEHeader</a>) Del(key <a href="/pkg/builtin/#string">string</a>)</pre>
Del deletes the values associated with key.




### <a id="MIMEHeader.Get">func</a> (MIMEHeader) [Get](https://golang.org/src/net/textproto/header.go?s=1073:1115#L21)
<pre>func (h <a href="#MIMEHeader">MIMEHeader</a>) Get(key <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Get gets the first value associated with the given key.
It is case insensitive; CanonicalMIMEHeaderKey is used
to canonicalize the provided key.
If there are no values associated with the key, Get returns "".
To access multiple values of a key, or to use non-canonical keys,
access the map directly.




### <a id="MIMEHeader.Set">func</a> (MIMEHeader) [Set](https://golang.org/src/net/textproto/header.go?s=657:699#L11)
<pre>func (h <a href="#MIMEHeader">MIMEHeader</a>) Set(key, value <a href="/pkg/builtin/#string">string</a>)</pre>
Set sets the header entries associated with key to
the single element value. It replaces any existing
values associated with key.




## <a id="Pipeline">type</a> [Pipeline](https://golang.org/src/net/textproto/pipeline.go?s=816:916#L18)
A Pipeline manages a pipelined in-order request/response sequence.

To use a Pipeline p to manage multiple clients on a connection,
each client should run:


	id := p.Next()	// take a number
	
	p.StartRequest(id)	// wait for turn to send request
	«send request»
	p.EndRequest(id)	// notify Pipeline that request is sent
	
	p.StartResponse(id)	// wait for turn to read response
	«read response»
	p.EndResponse(id)	// notify Pipeline that response is read

A pipelined server can use the same calls to ensure that
responses computed in parallel are written in the correct order.


<pre>type Pipeline struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="Pipeline.EndRequest">func</a> (\*Pipeline) [EndRequest](https://golang.org/src/net/textproto/pipeline.go?s=1368:1406#L42)
<pre>func (p *<a href="#Pipeline">Pipeline</a>) EndRequest(id <a href="/pkg/builtin/#uint">uint</a>)</pre>
EndRequest notifies p that the request with the given id has been sent
(or, if this is a server, received).




### <a id="Pipeline.EndResponse">func</a> (\*Pipeline) [EndResponse](https://golang.org/src/net/textproto/pipeline.go?s=1734:1773#L54)
<pre>func (p *<a href="#Pipeline">Pipeline</a>) EndResponse(id <a href="/pkg/builtin/#uint">uint</a>)</pre>
EndResponse notifies p that the response with the given id has been received
(or, if this is a server, sent).




### <a id="Pipeline.Next">func</a> (\*Pipeline) [Next](https://golang.org/src/net/textproto/pipeline.go?s=975:1005#L26)
<pre>func (p *<a href="#Pipeline">Pipeline</a>) Next() <a href="/pkg/builtin/#uint">uint</a></pre>
Next returns the next id for a request/response pair.




### <a id="Pipeline.StartRequest">func</a> (\*Pipeline) [StartRequest](https://golang.org/src/net/textproto/pipeline.go?s=1187:1227#L36)
<pre>func (p *<a href="#Pipeline">Pipeline</a>) StartRequest(id <a href="/pkg/builtin/#uint">uint</a>)</pre>
StartRequest blocks until it is time to send (or, if this is a server, receive)
the request with the given id.




### <a id="Pipeline.StartResponse">func</a> (\*Pipeline) [StartResponse](https://golang.org/src/net/textproto/pipeline.go?s=1549:1590#L48)
<pre>func (p *<a href="#Pipeline">Pipeline</a>) StartResponse(id <a href="/pkg/builtin/#uint">uint</a>)</pre>
StartResponse blocks until it is time to receive (or, if this is a server, send)
the request with the given id.




## <a id="ProtocolError">type</a> [ProtocolError](https://golang.org/src/net/textproto/textproto.go?s=1185:1210#L37)
A ProtocolError describes a protocol violation such
as an invalid response or a hung-up connection.


<pre>type ProtocolError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="ProtocolError.Error">func</a> (ProtocolError) [Error](https://golang.org/src/net/textproto/textproto.go?s=1212:1249#L39)
<pre>func (p <a href="#ProtocolError">ProtocolError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Reader">type</a> [Reader](https://golang.org/src/net/textproto/reader.go?s=379:497#L9)
A Reader implements convenience methods for reading requests
or responses from a text protocol network connection.


<pre>type Reader struct {
<span id="Reader.R"></span>    R *<a href="/pkg/bufio/">bufio</a>.<a href="/pkg/bufio/#Reader">Reader</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewReader">func</a> [NewReader](https://golang.org/src/net/textproto/reader.go?s=714:753#L20)
<pre>func NewReader(r *<a href="/pkg/bufio/">bufio</a>.<a href="/pkg/bufio/#Reader">Reader</a>) *<a href="#Reader">Reader</a></pre>
NewReader returns a new Reader reading from r.

To avoid denial of service attacks, the provided bufio.Reader
should be reading from an io.LimitReader or similar Reader to bound
the size of responses.






### <a id="Reader.DotReader">func</a> (\*Reader) [DotReader](https://golang.org/src/net/textproto/reader.go?s=8463:8501#L289)
<pre>func (r *<a href="#Reader">Reader</a>) DotReader() <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a></pre>
DotReader returns a new Reader that satisfies Reads using the
decoded text of a dot-encoded block read from r.
The returned Reader is only valid until the next call
to a method on r.

Dot encoding is a common framing used for data blocks
in text protocols such as SMTP.  The data consists of a sequence
of lines, each of which ends in "\r\n".  The sequence itself
ends at a line containing just a dot: ".\r\n".  Lines beginning
with a dot are escaped with an additional dot to avoid
looking like the end of the sequence.

The decoded form returned by the Reader's Read method
rewrites the "\r\n" line endings into the simpler "\n",
removes leading dot escapes if present, and stops with error io.EOF
after consuming (and discarding) the end-of-sequence line.




### <a id="Reader.ReadCodeLine">func</a> (\*Reader) [ReadCodeLine](https://golang.org/src/net/textproto/reader.go?s=5830:5913#L212)
<pre>func (r *<a href="#Reader">Reader</a>) ReadCodeLine(expectCode <a href="/pkg/builtin/#int">int</a>) (code <a href="/pkg/builtin/#int">int</a>, message <a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadCodeLine reads a response code line of the form


	code message

where code is a three-digit status code and the message
extends to the rest of the line. An example of such a line is:


	220 plan9.bell-labs.com ESMTP

If the prefix of the status does not match the digits in expectCode,
ReadCodeLine returns with err set to &Error{code, message}.
For example, if expectCode is 31, an error will be returned if
the status is not in the range [310,319].

If the response is multi-line, ReadCodeLine returns an error.

An expectCode <= 0 disables the check of the status code.




### <a id="Reader.ReadContinuedLine">func</a> (\*Reader) [ReadContinuedLine](https://golang.org/src/net/textproto/reader.go?s=2276:2328#L82)
<pre>func (r *<a href="#Reader">Reader</a>) ReadContinuedLine() (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadContinuedLine reads a possibly continued line from r,
eliding the final trailing ASCII white space.
Lines after the first are considered continuations if they
begin with a space or tab character. In the returned data,
continuation lines are separated from the previous line
only by a single space: the newline and leading white space
are removed.

For example, consider this input:


	Line 1
	  continued...
	Line 2

The first call to ReadContinuedLine will return "Line 1 continued..."
and the second will return "Line 2".

A line consisting of only white space is never continued.




### <a id="Reader.ReadContinuedLineBytes">func</a> (\*Reader) [ReadContinuedLineBytes](https://golang.org/src/net/textproto/reader.go?s=2794:2851#L103)
<pre>func (r *<a href="#Reader">Reader</a>) ReadContinuedLineBytes() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadContinuedLineBytes is like ReadContinuedLine but
returns a []byte instead of a string.




### <a id="Reader.ReadDotBytes">func</a> (\*Reader) [ReadDotBytes](https://golang.org/src/net/textproto/reader.go?s=10877:10924#L405)
<pre>func (r *<a href="#Reader">Reader</a>) ReadDotBytes() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadDotBytes reads a dot-encoding and returns the decoded data.

See the documentation for the DotReader method for details about dot-encoding.




### <a id="Reader.ReadDotLines">func</a> (\*Reader) [ReadDotLines](https://golang.org/src/net/textproto/reader.go?s=11187:11236#L413)
<pre>func (r *<a href="#Reader">Reader</a>) ReadDotLines() ([]<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadDotLines reads a dot-encoding and returns a slice
containing the decoded lines, with the final \r\n or \n elided from each.

See the documentation for the DotReader method for details about dot-encoding.




### <a id="Reader.ReadLine">func</a> (\*Reader) [ReadLine](https://golang.org/src/net/textproto/reader.go?s=918:961#L27)
<pre>func (r *<a href="#Reader">Reader</a>) ReadLine() (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadLine reads a single line from r,
eliding the final \n or \r\n from the returned string.




### <a id="Reader.ReadLineBytes">func</a> (\*Reader) [ReadLineBytes](https://golang.org/src/net/textproto/reader.go?s=1101:1149#L33)
<pre>func (r *<a href="#Reader">Reader</a>) ReadLineBytes() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadLineBytes is like ReadLine but returns a []byte instead of a string.




### <a id="Reader.ReadMIMEHeader">func</a> (\*Reader) [ReadMIMEHeader](https://golang.org/src/net/textproto/reader.go?s=12321:12374#L461)
<pre>func (r *<a href="#Reader">Reader</a>) ReadMIMEHeader() (<a href="#MIMEHeader">MIMEHeader</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadMIMEHeader reads a MIME-style header from r.
The header is a sequence of possibly continued Key: Value lines
ending in a blank line.
The returned map m maps CanonicalMIMEHeaderKey(key) to a
sequence of values in the same order encountered in the input.

For example, consider this input:


	My-Key: Value 1
	Long-Key: Even
	       Longer Value
	My-Key: Value 2

Given that input, ReadMIMEHeader returns the map:


	map[string][]string{
		"My-Key": {"Value 1", "Value 2"},
		"Long-Key": {"Even Longer Value"},
	}




### <a id="Reader.ReadResponse">func</a> (\*Reader) [ReadResponse](https://golang.org/src/net/textproto/reader.go?s=6987:7070#L247)
<pre>func (r *<a href="#Reader">Reader</a>) ReadResponse(expectCode <a href="/pkg/builtin/#int">int</a>) (code <a href="/pkg/builtin/#int">int</a>, message <a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadResponse reads a multi-line response of the form:


	code-message line 1
	code-message line 2
	...
	code message line n

where code is a three-digit status code. The first line starts with the
code and a hyphen. The response is terminated by a line that starts
with the same code followed by a space. Each line in message is
separated by a newline (\n).

See page 36 of RFC 959 (<a href="https://www.ietf.org/rfc/rfc959.txt">https://www.ietf.org/rfc/rfc959.txt</a>) for
details of another form of response accepted:


	code-message line 1
	message line 2
	...
	code message line n

If the prefix of the status does not match the digits in expectCode,
ReadResponse returns with err set to &Error{code, message}.
For example, if expectCode is 31, an error will be returned if
the status is not in the range [310,319].

An expectCode <= 0 disables the check of the status code.




## <a id="Writer">type</a> [Writer](https://golang.org/src/net/textproto/writer.go?s=332:389#L5)
A Writer implements convenience methods for writing
requests or responses to a text protocol network connection.


<pre>type Writer struct {
<span id="Writer.W"></span>    W *<a href="/pkg/bufio/">bufio</a>.<a href="/pkg/bufio/#Writer">Writer</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewWriter">func</a> [NewWriter](https://golang.org/src/net/textproto/writer.go?s=439:478#L11)
<pre>func NewWriter(w *<a href="/pkg/bufio/">bufio</a>.<a href="/pkg/bufio/#Writer">Writer</a>) *<a href="#Writer">Writer</a></pre>
NewWriter returns a new Writer writing to w.






### <a id="Writer.DotWriter">func</a> (\*Writer) [DotWriter](https://golang.org/src/net/textproto/writer.go?s=1208:1251#L33)
<pre>func (w *<a href="#Writer">Writer</a>) DotWriter() <a href="/pkg/io/">io</a>.<a href="/pkg/io/#WriteCloser">WriteCloser</a></pre>
DotWriter returns a writer that can be used to write a dot-encoding to w.
It takes care of inserting leading dots when necessary,
translating line-ending \n into \r\n, and adding the final .\r\n line
when the DotWriter is closed. The caller should close the
DotWriter before the next call to a method on w.

See the documentation for Reader's DotReader method for details about dot-encoding.




### <a id="Writer.PrintfLine">func</a> (\*Writer) [PrintfLine](https://golang.org/src/net/textproto/writer.go?s=635:704#L19)
<pre>func (w *<a href="#Writer">Writer</a>) PrintfLine(format <a href="/pkg/builtin/#string">string</a>, args ...interface{}) <a href="/pkg/builtin/#error">error</a></pre>
PrintfLine writes the formatted output followed by \r\n.








