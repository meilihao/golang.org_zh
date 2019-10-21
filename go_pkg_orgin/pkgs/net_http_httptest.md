

# httptest
`import "net/http/httptest"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package httptest provides utilities for HTTP testing.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func NewRequest(method, target string, body io.Reader) *http.Request](#NewRequest)
* [type ResponseRecorder](#ResponseRecorder)
  * [func NewRecorder() *ResponseRecorder](#NewRecorder)
  * [func (rw *ResponseRecorder) Flush()](#ResponseRecorder.Flush)
  * [func (rw *ResponseRecorder) Header() http.Header](#ResponseRecorder.Header)
  * [func (rw *ResponseRecorder) Result() *http.Response](#ResponseRecorder.Result)
  * [func (rw *ResponseRecorder) Write(buf []byte) (int, error)](#ResponseRecorder.Write)
  * [func (rw *ResponseRecorder) WriteHeader(code int)](#ResponseRecorder.WriteHeader)
  * [func (rw *ResponseRecorder) WriteString(str string) (int, error)](#ResponseRecorder.WriteString)
* [type Server](#Server)
  * [func NewServer(handler http.Handler) *Server](#NewServer)
  * [func NewTLSServer(handler http.Handler) *Server](#NewTLSServer)
  * [func NewUnstartedServer(handler http.Handler) *Server](#NewUnstartedServer)
  * [func (s *Server) Certificate() *x509.Certificate](#Server.Certificate)
  * [func (s *Server) Client() *http.Client](#Server.Client)
  * [func (s *Server) Close()](#Server.Close)
  * [func (s *Server) CloseClientConnections()](#Server.CloseClientConnections)
  * [func (s *Server) Start()](#Server.Start)
  * [func (s *Server) StartTLS()](#Server.StartTLS)


#### <a id="pkg-examples">Examples</a>
* [NewTLSServer](#example_NewTLSServer)
* [ResponseRecorder](#example_ResponseRecorder)
* [Server](#example_Server)


#### <a id="pkg-files">Package files</a>
[httptest.go](https://golang.org/src/net/http/httptest/httptest.go) [recorder.go](https://golang.org/src/net/http/httptest/recorder.go) [server.go](https://golang.org/src/net/http/httptest/server.go) 


## <a id="pkg-constants">Constants</a>
DefaultRemoteAddr is the default remote address to return in RemoteAddr if
an explicit DefaultRemoteAddr isn't set on ResponseRecorder.


<pre>const <span id="DefaultRemoteAddr">DefaultRemoteAddr</span> = &#34;1.2.3.4&#34;</pre>



## <a id="NewRequest">func</a> [NewRequest](https://golang.org/src/net/http/httptest/httptest.go?s=1162:1230#L31)
<pre>func NewRequest(method, target <a href="/pkg/builtin/#string">string</a>, body <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a></pre>
NewRequest returns a new incoming server Request, suitable
for passing to an http.Handler for testing.

The target is the RFC 7230 "request-target": it may be either a
path or an absolute URL. If target is an absolute URL, the host name
from the URL is used. Otherwise, "example.com" is used.

The TLS field is set to a non-nil dummy value if target has scheme
"https".

The Request.Proto is always HTTP/1.1.

An empty method means "GET".

The provided body may be nil. If the body is of type *bytes.Reader,
*strings.Reader, or *bytes.Buffer, the Request.ContentLength is
set.

NewRequest panics on error for ease of use in testing, where a
panic is acceptable.

To generate a client HTTP request instead of a server request, see
the NewRequest function in the net/http package.





## <a id="ResponseRecorder">type</a> [ResponseRecorder](https://golang.org/src/net/http/httptest/recorder.go?s=413:1378#L10)
ResponseRecorder is an implementation of http.ResponseWriter that
records its mutations for later inspection in tests.


<pre>type ResponseRecorder struct {
<span id="ResponseRecorder.Code"></span>    <span class="comment">// Code is the HTTP response code set by WriteHeader.</span>
    <span class="comment">//</span>
    <span class="comment">// Note that if a Handler never calls WriteHeader or Write,</span>
    <span class="comment">// this might end up being 0, rather than the implicit</span>
    <span class="comment">// http.StatusOK. To get the implicit value, use the Result</span>
    <span class="comment">// method.</span>
    Code <a href="/pkg/builtin/#int">int</a>

<span id="ResponseRecorder.HeaderMap"></span>    <span class="comment">// HeaderMap contains the headers explicitly set by the Handler.</span>
    <span class="comment">// It is an internal detail.</span>
    <span class="comment">//</span>
    <span class="comment">// Deprecated: HeaderMap exists for historical compatibility</span>
    <span class="comment">// and should not be used. To access the headers returned by a handler,</span>
    <span class="comment">// use the Response.Header map as returned by the Result method.</span>
    HeaderMap <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Header">Header</a>

<span id="ResponseRecorder.Body"></span>    <span class="comment">// Body is the buffer to which the Handler&#39;s Write calls are sent.</span>
    <span class="comment">// If nil, the Writes are silently discarded.</span>
    Body *<a href="/pkg/bytes/">bytes</a>.<a href="/pkg/bytes/#Buffer">Buffer</a>

<span id="ResponseRecorder.Flushed"></span>    <span class="comment">// Flushed is whether the Handler called Flush.</span>
    Flushed <a href="/pkg/builtin/#bool">bool</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_ResponseRecorder">Example</a>
```go
```

output:
```txt
```




### <a id="NewRecorder">func</a> [NewRecorder](https://golang.org/src/net/http/httptest/recorder.go?s=1436:1472#L40)
<pre>func NewRecorder() *<a href="#ResponseRecorder">ResponseRecorder</a></pre>
NewRecorder returns an initialized ResponseRecorder.






### <a id="ResponseRecorder.Flush">func</a> (\*ResponseRecorder) [Flush](https://golang.org/src/net/http/httptest/recorder.go?s=3712:3747#L129)
<pre>func (rw *<a href="#ResponseRecorder">ResponseRecorder</a>) Flush()</pre>
Flush implements http.Flusher. To test whether Flush was
called, see rw.Flushed.




### <a id="ResponseRecorder.Header">func</a> (\*ResponseRecorder) [Header](https://golang.org/src/net/http/httptest/recorder.go?s=2013:2061#L56)
<pre>func (rw *<a href="#ResponseRecorder">ResponseRecorder</a>) Header() <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Header">Header</a></pre>
Header implements http.ResponseWriter. It returns the response
headers to mutate within a handler. To test the headers that were
written after a handler completes, use the Result method and see
the returned Response value's Header.




### <a id="ResponseRecorder.Result">func</a> (\*ResponseRecorder) [Result](https://golang.org/src/net/http/httptest/recorder.go?s=4458:4509#L151)
<pre>func (rw *<a href="#ResponseRecorder">ResponseRecorder</a>) Result() *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Response">Response</a></pre>
Result returns the response generated by the handler.

The returned Response will have at least its StatusCode,
Header, Body, and optionally Trailer populated.
More fields may be populated in the future, so callers should
not DeepEqual the result in tests.

The Response.Header is a snapshot of the headers at the time of the
first write call, or at the time of this call, if the handler never
did a write.

The Response.Body is guaranteed to be non-nil and Body.Read call is
guaranteed to not return any error other than io.EOF.

Result must only be called after the handler has finished running.




### <a id="ResponseRecorder.Write">func</a> (\*ResponseRecorder) [Write](https://golang.org/src/net/http/httptest/recorder.go?s=2927:2985#L96)
<pre>func (rw *<a href="#ResponseRecorder">ResponseRecorder</a>) Write(buf []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Write implements http.ResponseWriter. The data in buf is written to
rw.Body, if not nil.




### <a id="ResponseRecorder.WriteHeader">func</a> (\*ResponseRecorder) [WriteHeader](https://golang.org/src/net/http/httptest/recorder.go?s=3396:3445#L115)
<pre>func (rw *<a href="#ResponseRecorder">ResponseRecorder</a>) WriteHeader(code <a href="/pkg/builtin/#int">int</a>)</pre>
WriteHeader implements http.ResponseWriter.




### <a id="ResponseRecorder.WriteString">func</a> (\*ResponseRecorder) [WriteString](https://golang.org/src/net/http/httptest/recorder.go?s=3180:3244#L106)
<pre>func (rw *<a href="#ResponseRecorder">ResponseRecorder</a>) WriteString(str <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
WriteString implements io.StringWriter. The data in str is written
to rw.Body, if not nil.




## <a id="Server">type</a> [Server](https://golang.org/src/net/http/httptest/server.go?s=477:1449#L16)
A Server is an HTTP server listening on a system-chosen port on the
local loopback interface, for use in end-to-end HTTP tests.


<pre>type Server struct {
<span id="Server.URL"></span>    URL      <a href="/pkg/builtin/#string">string</a> <span class="comment">// base URL of form http://ipaddr:port with no trailing slash</span>
<span id="Server.Listener"></span>    Listener <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Listener">Listener</a>

<span id="Server.TLS"></span>    <span class="comment">// TLS is the optional TLS configuration, populated with a new config</span>
    <span class="comment">// after TLS is started. If set on an unstarted server before StartTLS</span>
    <span class="comment">// is called, existing fields are copied into the new config.</span>
    TLS *<a href="/pkg/crypto/tls/">tls</a>.<a href="/pkg/crypto/tls/#Config">Config</a>

<span id="Server.Config"></span>    <span class="comment">// Config may be changed after calling NewUnstartedServer and</span>
    <span class="comment">// before Start or StartTLS.</span>
    Config *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Server">Server</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Server">Example</a>
```go
```

output:
```txt
```




### <a id="NewServer">func</a> [NewServer](https://golang.org/src/net/http/httptest/server.go?s=2824:2868#L88)
<pre>func NewServer(handler <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Handler">Handler</a>) *<a href="#Server">Server</a></pre>
NewServer starts and returns a new Server.
The caller should call Close when finished, to shut it down.




### <a id="NewTLSServer">func</a> [NewTLSServer](https://golang.org/src/net/http/httptest/server.go?s=4921:4968#L168)
<pre>func NewTLSServer(handler <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Handler">Handler</a>) *<a href="#Server">Server</a></pre>
NewTLSServer starts and returns a new Server using TLS.
The caller should call Close when finished, to shut it down.


<a id="example_NewTLSServer">Example</a>
```go
```

output:
```txt
```


### <a id="NewUnstartedServer">func</a> [NewUnstartedServer](https://golang.org/src/net/http/httptest/server.go?s=3149:3202#L100)
<pre>func NewUnstartedServer(handler <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Handler">Handler</a>) *<a href="#Server">Server</a></pre>
NewUnstartedServer returns a new Server but doesn't start it.

After changing its configuration, the caller should call Start or
StartTLS.

The caller should call Close when finished, to shut it down.






### <a id="Server.Certificate">func</a> (\*Server) [Certificate](https://golang.org/src/net/http/httptest/server.go?s=8164:8212#L273)
<pre>func (s *<a href="#Server">Server</a>) Certificate() *<a href="/pkg/crypto/x509/">x509</a>.<a href="/pkg/crypto/x509/#Certificate">Certificate</a></pre>
Certificate returns the certificate used by the server, or nil if
the server doesn't use TLS.




### <a id="Server.Client">func</a> (\*Server) [Client](https://golang.org/src/net/http/httptest/server.go?s=8438:8476#L280)
<pre>func (s *<a href="#Server">Server</a>) Client() *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Client">Client</a></pre>
Client returns an HTTP client configured for making requests to the server.
It is configured to trust the server's TLS test certificate and will
close its idle connections on Server.Close.




### <a id="Server.Close">func</a> (\*Server) [Close](https://golang.org/src/net/http/httptest/server.go?s=5205:5229#L180)
<pre>func (s *<a href="#Server">Server</a>) Close()</pre>
Close shuts down the server and blocks until all outstanding
requests on this server have completed.




### <a id="Server.CloseClientConnections">func</a> (\*Server) [CloseClientConnections](https://golang.org/src/net/http/httptest/server.go?s=7450:7491#L244)
<pre>func (s *<a href="#Server">Server</a>) CloseClientConnections()</pre>
CloseClientConnections closes any open HTTP connections to the test Server.




### <a id="Server.Start">func</a> (\*Server) [Start](https://golang.org/src/net/http/httptest/server.go?s=3354:3378#L108)
<pre>func (s *<a href="#Server">Server</a>) Start()</pre>
Start starts a server from NewUnstartedServer.




### <a id="Server.StartTLS">func</a> (\*Server) [StartTLS](https://golang.org/src/net/http/httptest/server.go?s=3745:3772#L125)
<pre>func (s *<a href="#Server">Server</a>) StartTLS()</pre>
StartTLS starts TLS on a server from NewUnstartedServer.







