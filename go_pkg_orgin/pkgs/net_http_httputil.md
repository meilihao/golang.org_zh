

# httputil
`import "net/http/httputil"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package httputil provides HTTP utility functions, complementing the
more common ones in the net/http package.




## <a id="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func DumpRequest(req *http.Request, body bool) ([]byte, error)](#DumpRequest)
* [func DumpRequestOut(req *http.Request, body bool) ([]byte, error)](#DumpRequestOut)
* [func DumpResponse(resp *http.Response, body bool) ([]byte, error)](#DumpResponse)
* [func NewChunkedReader(r io.Reader) io.Reader](#NewChunkedReader)
* [func NewChunkedWriter(w io.Writer) io.WriteCloser](#NewChunkedWriter)
* [type BufferPool](#BufferPool)
* [type ClientConn](#ClientConn)
  * [func NewClientConn(c net.Conn, r *bufio.Reader) *ClientConn](#NewClientConn)
  * [func NewProxyClientConn(c net.Conn, r *bufio.Reader) *ClientConn](#NewProxyClientConn)
  * [func (cc *ClientConn) Close() error](#ClientConn.Close)
  * [func (cc *ClientConn) Do(req *http.Request) (*http.Response, error)](#ClientConn.Do)
  * [func (cc *ClientConn) Hijack() (c net.Conn, r *bufio.Reader)](#ClientConn.Hijack)
  * [func (cc *ClientConn) Pending() int](#ClientConn.Pending)
  * [func (cc *ClientConn) Read(req *http.Request) (resp *http.Response, err error)](#ClientConn.Read)
  * [func (cc *ClientConn) Write(req *http.Request) error](#ClientConn.Write)
* [type ReverseProxy](#ReverseProxy)
  * [func NewSingleHostReverseProxy(target *url.URL) *ReverseProxy](#NewSingleHostReverseProxy)
  * [func (p *ReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request)](#ReverseProxy.ServeHTTP)
* [type ServerConn](#ServerConn)
  * [func NewServerConn(c net.Conn, r *bufio.Reader) *ServerConn](#NewServerConn)
  * [func (sc *ServerConn) Close() error](#ServerConn.Close)
  * [func (sc *ServerConn) Hijack() (net.Conn, *bufio.Reader)](#ServerConn.Hijack)
  * [func (sc *ServerConn) Pending() int](#ServerConn.Pending)
  * [func (sc *ServerConn) Read() (*http.Request, error)](#ServerConn.Read)
  * [func (sc *ServerConn) Write(req *http.Request, resp *http.Response) error](#ServerConn.Write)


#### <a id="pkg-examples">Examples</a>
* [DumpRequest](#example_DumpRequest)
* [DumpRequestOut](#example_DumpRequestOut)
* [DumpResponse](#example_DumpResponse)
* [ReverseProxy](#example_ReverseProxy)


#### <a id="pkg-files">Package files</a>
[dump.go](https://golang.org/src/net/http/httputil/dump.go) [httputil.go](https://golang.org/src/net/http/httputil/httputil.go) [persist.go](https://golang.org/src/net/http/httputil/persist.go) [reverseproxy.go](https://golang.org/src/net/http/httputil/reverseproxy.go) 




## <a id="pkg-variables">Variables</a>

<pre>var (
    <span class="comment">// Deprecated: No longer used.</span>
    <span id="ErrPersistEOF">ErrPersistEOF</span> = &amp;<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#ProtocolError">ProtocolError</a>{<a href="/pkg/net/http/#ProtocolError.ErrorString">ErrorString</a>: &#34;persistent connection closed&#34;}

    <span class="comment">// Deprecated: No longer used.</span>
    <span id="ErrClosed">ErrClosed</span> = &amp;<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#ProtocolError">ProtocolError</a>{<a href="/pkg/net/http/#ProtocolError.ErrorString">ErrorString</a>: &#34;connection closed by user&#34;}

    <span class="comment">// Deprecated: No longer used.</span>
    <span id="ErrPipeline">ErrPipeline</span> = &amp;<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#ProtocolError">ProtocolError</a>{<a href="/pkg/net/http/#ProtocolError.ErrorString">ErrorString</a>: &#34;pipeline error&#34;}
)</pre>ErrLineTooLong is returned when reading malformed chunked data
with lines that are too long.


<pre>var <span id="ErrLineTooLong">ErrLineTooLong</span> = <a href="/pkg/net/http/internal/">internal</a>.<a href="/pkg/net/http/internal/#ErrLineTooLong">ErrLineTooLong</a></pre>

## <a id="DumpRequest">func</a> [DumpRequest](https://golang.org/src/net/http/httputil/dump.go?s=5638:5700#L181)
<pre>func DumpRequest(req *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>, body <a href="/pkg/builtin/#bool">bool</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DumpRequest returns the given request in its HTTP/1.x wire
representation. It should only be used by servers to debug client
requests. The returned representation is an approximation only;
some details of the initial request are lost while parsing it into
an http.Request. In particular, the order and case of header field
names are lost. The order of values in multi-valued headers is kept
intact. HTTP/2 requests are dumped in HTTP/1.x form, not in their
original binary representations.

If body is true, DumpRequest also returns the body. To do so, it
consumes req.Body and then replaces it with a new io.ReadCloser
that yields the same bytes. If DumpRequest returns an error,
the state of req is undefined.

The documentation for http.Request.Write details which fields
of req are included in the dump.


<a id="example_DumpRequest">Example</a>
```go
```

output:
```txt
```

## <a id="DumpRequestOut">func</a> [DumpRequestOut](https://golang.org/src/net/http/httputil/dump.go?s=1848:1913#L56)
<pre>func DumpRequestOut(req *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>, body <a href="/pkg/builtin/#bool">bool</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DumpRequestOut is like DumpRequest but for outgoing client requests. It
includes any headers that the standard http.Transport adds, such as
User-Agent.


<a id="example_DumpRequestOut">Example</a>
```go
```

output:
```txt
```

## <a id="DumpResponse">func</a> [DumpResponse](https://golang.org/src/net/http/httputil/dump.go?s=8166:8231#L271)
<pre>func DumpResponse(resp *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Response">Response</a>, body <a href="/pkg/builtin/#bool">bool</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DumpResponse is like DumpRequest but dumps a response.


<a id="example_DumpResponse">Example</a>
```go
```

output:
```txt
```

## <a id="NewChunkedReader">func</a> [NewChunkedReader](https://golang.org/src/net/http/httputil/httputil.go?s=688:732#L10)
<pre>func NewChunkedReader(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a></pre>
NewChunkedReader returns a new chunkedReader that translates the data read from r
out of HTTP "chunked" format before returning it.
The chunkedReader returns io.EOF when the final 0-length chunk is read.

NewChunkedReader is not needed by normal applications. The http package
automatically decodes chunking when reading response bodies.



## <a id="NewChunkedWriter">func</a> [NewChunkedWriter](https://golang.org/src/net/http/httputil/httputil.go?s=1431:1480#L25)
<pre>func NewChunkedWriter(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#WriteCloser">WriteCloser</a></pre>
NewChunkedWriter returns a new chunkedWriter that translates writes into HTTP
"chunked" format before writing them to w. Closing the returned chunkedWriter
sends the final 0-length chunk that marks the end of the stream but does
not send the final CRLF that appears after trailers; trailers and the last
CRLF must be written separately.

NewChunkedWriter is not needed by normal applications. The http
package adds chunking automatically if handlers don't set a
Content-Length header. Using NewChunkedWriter inside a handler
would result in double chunking or chunking with a Content-Length
length, both of which are wrong.





## <a id="BufferPool">type</a> [BufferPool](https://golang.org/src/net/http/httputil/reverseproxy.go?s=2617:2673#L73)
A BufferPool is an interface for getting and returning temporary
byte slices for use by io.CopyBuffer.


<pre>type BufferPool interface {
    Get() []<a href="/pkg/builtin/#byte">byte</a>
    Put([]<a href="/pkg/builtin/#byte">byte</a>)
}</pre>











## <a id="ClientConn">type</a> [ClientConn](https://golang.org/src/net/http/httputil/persist.go?s=5973:6341#L220)
ClientConn is an artifact of Go's early HTTP implementation.
It is low-level, old, and unused by Go's current HTTP stack.
We should have deleted it before Go 1.

Deprecated: Use Client or Transport in package net/http instead.


<pre>type ClientConn struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewClientConn">func</a> [NewClientConn](https://golang.org/src/net/http/httputil/persist.go?s=6591:6650#L238)
<pre>func NewClientConn(c <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Conn">Conn</a>, r *<a href="/pkg/bufio/">bufio</a>.<a href="/pkg/bufio/#Reader">Reader</a>) *<a href="#ClientConn">ClientConn</a></pre>
NewClientConn is an artifact of Go's early HTTP implementation.
It is low-level, old, and unused by Go's current HTTP stack.
We should have deleted it before Go 1.

Deprecated: Use the Client or Transport in package net/http instead.




### <a id="NewProxyClientConn">func</a> [NewProxyClientConn](https://golang.org/src/net/http/httputil/persist.go?s=7083:7147#L255)
<pre>func NewProxyClientConn(c <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Conn">Conn</a>, r *<a href="/pkg/bufio/">bufio</a>.<a href="/pkg/bufio/#Reader">Reader</a>) *<a href="#ClientConn">ClientConn</a></pre>
NewProxyClientConn is an artifact of Go's early HTTP implementation.
It is low-level, old, and unused by Go's current HTTP stack.
We should have deleted it before Go 1.

Deprecated: Use the Client or Transport in package net/http instead.






### <a id="ClientConn.Close">func</a> (\*ClientConn) [Close](https://golang.org/src/net/http/httputil/persist.go?s=7765:7800#L276)
<pre>func (cc *<a href="#ClientConn">ClientConn</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close calls Hijack and then also closes the underlying connection.




### <a id="ClientConn.Do">func</a> (\*ClientConn) [Do](https://golang.org/src/net/http/httputil/persist.go?s=11094:11161#L415)
<pre>func (cc *<a href="#ClientConn">ClientConn</a>) Do(req *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>) (*<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Response">Response</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Do is convenience method that writes a request and reads a response.




### <a id="ClientConn.Hijack">func</a> (\*ClientConn) [Hijack](https://golang.org/src/net/http/httputil/persist.go?s=7541:7601#L265)
<pre>func (cc *<a href="#ClientConn">ClientConn</a>) Hijack() (c <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Conn">Conn</a>, r *<a href="/pkg/bufio/">bufio</a>.<a href="/pkg/bufio/#Reader">Reader</a>)</pre>
Hijack detaches the ClientConn and returns the underlying connection as well
as the read-side bufio which may have some left over data. Hijack may be
called before the user or Read have signaled the end of the keep-alive
logic. The user should not call Hijack while Read or Write is in progress.




### <a id="ClientConn.Pending">func</a> (\*ClientConn) [Pending](https://golang.org/src/net/http/httputil/persist.go?s=9367:9402#L343)
<pre>func (cc *<a href="#ClientConn">ClientConn</a>) Pending() <a href="/pkg/builtin/#int">int</a></pre>
Pending returns the number of unanswered requests
that have been sent on the connection.




### <a id="ClientConn.Read">func</a> (\*ClientConn) [Read](https://golang.org/src/net/http/httputil/persist.go?s=9747:9825#L353)
<pre>func (cc *<a href="#ClientConn">ClientConn</a>) Read(req *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>) (resp *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Response">Response</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Read reads the next response from the wire. A valid response might be
returned together with an ErrPersistEOF, which means that the remote
requested that this be the last request serviced. Read can be called
concurrently with Write, but not with another Read.




### <a id="ClientConn.Write">func</a> (\*ClientConn) [Write](https://golang.org/src/net/http/httputil/persist.go?s=8267:8319#L289)
<pre>func (cc *<a href="#ClientConn">ClientConn</a>) Write(req *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>) <a href="/pkg/builtin/#error">error</a></pre>
Write writes a request. An ErrPersistEOF error is returned if the connection
has been closed in an HTTP keep-alive sense. If req.Close equals true, the
keep-alive connection is logically closed after this request and the opposing
server is informed. An ErrUnexpectedEOF indicates the remote closed the
underlying TCP connection, which is usually considered as graceful close.




## <a id="ReverseProxy">type</a> [ReverseProxy](https://golang.org/src/net/http/httputil/reverseproxy.go?s=490:2506#L17)
ReverseProxy is an HTTP Handler that takes an incoming request and
sends it to another server, proxying the response back to the
client.


<pre>type ReverseProxy struct {
<span id="ReverseProxy.Director"></span>    <span class="comment">// Director must be a function which modifies</span>
    <span class="comment">// the request into a new request to be sent</span>
    <span class="comment">// using Transport. Its response is then copied</span>
    <span class="comment">// back to the original client unmodified.</span>
    <span class="comment">// Director must not access the provided Request</span>
    <span class="comment">// after returning.</span>
    Director func(*<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>)

    <span class="comment">// The transport used to perform proxy requests.</span>
    <span class="comment">// If nil, http.DefaultTransport is used.</span>
<span id="ReverseProxy.Transport"></span>    Transport <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#RoundTripper">RoundTripper</a>

<span id="ReverseProxy.FlushInterval"></span>    <span class="comment">// FlushInterval specifies the flush interval</span>
    <span class="comment">// to flush to the client while copying the</span>
    <span class="comment">// response body.</span>
    <span class="comment">// If zero, no periodic flushing is done.</span>
    <span class="comment">// A negative value means to flush immediately</span>
    <span class="comment">// after each write to the client.</span>
    <span class="comment">// The FlushInterval is ignored when ReverseProxy</span>
    <span class="comment">// recognizes a response as a streaming response;</span>
    <span class="comment">// for such responses, writes are flushed to the client</span>
    <span class="comment">// immediately.</span>
    FlushInterval <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>

<span id="ReverseProxy.ErrorLog"></span>    <span class="comment">// ErrorLog specifies an optional logger for errors</span>
    <span class="comment">// that occur when attempting to proxy the request.</span>
    <span class="comment">// If nil, logging is done via the log package&#39;s standard logger.</span>
    ErrorLog *<a href="/pkg/log/">log</a>.<a href="/pkg/log/#Logger">Logger</a>

<span id="ReverseProxy.BufferPool"></span>    <span class="comment">// BufferPool optionally specifies a buffer pool to</span>
    <span class="comment">// get byte slices for use by io.CopyBuffer when</span>
    <span class="comment">// copying HTTP response bodies.</span>
    BufferPool <a href="#BufferPool">BufferPool</a>

<span id="ReverseProxy.ModifyResponse"></span>    <span class="comment">// ModifyResponse is an optional function that modifies the</span>
    <span class="comment">// Response from the backend. It is called if the backend</span>
    <span class="comment">// returns a response at all, with any HTTP status code.</span>
    <span class="comment">// If the backend is unreachable, the optional ErrorHandler is</span>
    <span class="comment">// called without any call to ModifyResponse.</span>
    <span class="comment">//</span>
    <span class="comment">// If ModifyResponse returns an error, ErrorHandler is called</span>
    <span class="comment">// with its error value. If ErrorHandler is nil, its default</span>
    <span class="comment">// implementation is used.</span>
    ModifyResponse func(*<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Response">Response</a>) <a href="/pkg/builtin/#error">error</a>

<span id="ReverseProxy.ErrorHandler"></span>    <span class="comment">// ErrorHandler is an optional function that handles errors</span>
    <span class="comment">// reaching the backend or errors from ModifyResponse.</span>
    <span class="comment">//</span>
    <span class="comment">// If nil, the default is to log the provided error and return</span>
    <span class="comment">// a 502 Status Bad Gateway response.</span>
    ErrorHandler func(<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#ResponseWriter">ResponseWriter</a>, *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>, <a href="/pkg/builtin/#error">error</a>)
}
</pre>





<a id="example_ReverseProxy">Example</a>
```go
```

output:
```txt
```




### <a id="NewSingleHostReverseProxy">func</a> [NewSingleHostReverseProxy](https://golang.org/src/net/http/httputil/reverseproxy.go?s=3318:3379#L97)
<pre>func NewSingleHostReverseProxy(target *<a href="/pkg/net/url/">url</a>.<a href="/pkg/net/url/#URL">URL</a>) *<a href="#ReverseProxy">ReverseProxy</a></pre>
NewSingleHostReverseProxy returns a new ReverseProxy that routes
URLs to the scheme, host, and base path provided in target. If the
target's path is "/base" and the incoming request was for "/dir",
the target request will be for /base/dir.
NewSingleHostReverseProxy does not rewrite the Host header.
To rewrite Host headers, use ReverseProxy directly with a custom
Director policy.






### <a id="ReverseProxy.ServeHTTP">func</a> (\*ReverseProxy) [ServeHTTP](https://golang.org/src/net/http/httputil/reverseproxy.go?s=5504:5579#L167)
<pre>func (p *<a href="#ReverseProxy">ReverseProxy</a>) ServeHTTP(rw <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#ResponseWriter">ResponseWriter</a>, req *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>)</pre>



## <a id="ServerConn">type</a> [ServerConn](https://golang.org/src/net/http/httputil/persist.go?s=1010:1327#L27)
ServerConn is an artifact of Go's early HTTP implementation.
It is low-level, old, and unused by Go's current HTTP stack.
We should have deleted it before Go 1.

Deprecated: Use the Server in package net/http instead.


<pre>type ServerConn struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewServerConn">func</a> [NewServerConn](https://golang.org/src/net/http/httputil/persist.go?s=1564:1623#L44)
<pre>func NewServerConn(c <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Conn">Conn</a>, r *<a href="/pkg/bufio/">bufio</a>.<a href="/pkg/bufio/#Reader">Reader</a>) *<a href="#ServerConn">ServerConn</a></pre>
NewServerConn is an artifact of Go's early HTTP implementation.
It is low-level, old, and unused by Go's current HTTP stack.
We should have deleted it before Go 1.

Deprecated: Use the Server in package net/http instead.






### <a id="ServerConn.Close">func</a> (\*ServerConn) [Close](https://golang.org/src/net/http/httputil/persist.go?s=2265:2300#L66)
<pre>func (sc *<a href="#ServerConn">ServerConn</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close calls Hijack and then also closes the underlying connection.




### <a id="ServerConn.Hijack">func</a> (\*ServerConn) [Hijack](https://golang.org/src/net/http/httputil/persist.go?s=2038:2094#L55)
<pre>func (sc *<a href="#ServerConn">ServerConn</a>) Hijack() (<a href="/pkg/net/">net</a>.<a href="/pkg/net/#Conn">Conn</a>, *<a href="/pkg/bufio/">bufio</a>.<a href="/pkg/bufio/#Reader">Reader</a>)</pre>
Hijack detaches the ServerConn and returns the underlying connection as well
as the read-side bufio which may have some left over data. Hijack may be
called before Read has signaled the end of the keep-alive logic. The user
should not call Hijack while Read or Write is in progress.




### <a id="ServerConn.Pending">func</a> (\*ServerConn) [Pending](https://golang.org/src/net/http/httputil/persist.go?s=4387:4422#L156)
<pre>func (sc *<a href="#ServerConn">ServerConn</a>) Pending() <a href="/pkg/builtin/#int">int</a></pre>
Pending returns the number of unanswered requests
that have been received on the connection.




### <a id="ServerConn.Read">func</a> (\*ServerConn) [Read](https://golang.org/src/net/http/httputil/persist.go?s=2635:2686#L78)
<pre>func (sc *<a href="#ServerConn">ServerConn</a>) Read() (*<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Read returns the next request on the wire. An ErrPersistEOF is returned if
it is gracefully determined that there are no more requests (e.g. after the
first request on an HTTP/1.0 connection, or after a Connection:close on a
HTTP/1.1 connection).




### <a id="ServerConn.Write">func</a> (\*ServerConn) [Write](https://golang.org/src/net/http/httputil/persist.go?s=4734:4807#L165)
<pre>func (sc *<a href="#ServerConn">ServerConn</a>) Write(req *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>, resp *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Response">Response</a>) <a href="/pkg/builtin/#error">error</a></pre>
Write writes resp in response to req. To close the connection gracefully, set the
Response.Close field to true. Write should be considered operational until
it returns an error, regardless of any errors returned on the Read side.







