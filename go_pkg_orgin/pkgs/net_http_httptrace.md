

# httptrace
`import "net/http/httptrace"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package httptrace provides mechanisms to trace the events within
HTTP client requests.


<a id="example_">Example</a>
```go
```

output:
```txt
```


## <a id="pkg-index">Index</a>
* [func WithClientTrace(ctx context.Context, trace *ClientTrace) context.Context](#WithClientTrace)
* [type ClientTrace](#ClientTrace)
  * [func ContextClientTrace(ctx context.Context) *ClientTrace](#ContextClientTrace)
* [type DNSDoneInfo](#DNSDoneInfo)
* [type DNSStartInfo](#DNSStartInfo)
* [type GotConnInfo](#GotConnInfo)
* [type WroteRequestInfo](#WroteRequestInfo)


#### <a id="pkg-examples">Examples</a>
* [Package](#example_)


#### <a id="pkg-files">Package files</a>
[trace.go](https://golang.org/src/net/http/httptrace/trace.go) 






## <a id="WithClientTrace">func</a> [WithClientTrace](https://golang.org/src/net/http/httptrace/trace.go?s=991:1068#L24)
<pre>func WithClientTrace(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, trace *<a href="#ClientTrace">ClientTrace</a>) <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a></pre>
WithClientTrace returns a new context based on the provided parent
ctx. HTTP client requests made with the returned context will use
the provided trace hooks, in addition to any previous hooks
registered with ctx. Any hooks defined in the provided trace will
be called first.





## <a id="ClientTrace">type</a> [ClientTrace](https://golang.org/src/net/http/httptrace/trace.go?s=2359:5819#L70)
ClientTrace is a set of hooks to run at various stages of an outgoing
HTTP request. Any particular hook may be nil. Functions may be
called concurrently from different goroutines and some may be called
after the request has completed or failed.

ClientTrace currently traces a single HTTP request & response
during a single round trip and has no hooks that span a series
of redirected requests.

See <a href="https://blog.golang.org/http-tracing">https://blog.golang.org/http-tracing</a> for more.


<pre>type ClientTrace struct {
<span id="ClientTrace.GetConn"></span>    <span class="comment">// GetConn is called before a connection is created or</span>
    <span class="comment">// retrieved from an idle pool. The hostPort is the</span>
    <span class="comment">// &#34;host:port&#34; of the target or proxy. GetConn is called even</span>
    <span class="comment">// if there&#39;s already an idle cached connection available.</span>
    GetConn func(hostPort <a href="/pkg/builtin/#string">string</a>)

<span id="ClientTrace.GotConn"></span>    <span class="comment">// GotConn is called after a successful connection is</span>
    <span class="comment">// obtained. There is no hook for failure to obtain a</span>
    <span class="comment">// connection; instead, use the error from</span>
    <span class="comment">// Transport.RoundTrip.</span>
    GotConn func(<a href="#GotConnInfo">GotConnInfo</a>)

<span id="ClientTrace.PutIdleConn"></span>    <span class="comment">// PutIdleConn is called when the connection is returned to</span>
    <span class="comment">// the idle pool. If err is nil, the connection was</span>
    <span class="comment">// successfully returned to the idle pool. If err is non-nil,</span>
    <span class="comment">// it describes why not. PutIdleConn is not called if</span>
    <span class="comment">// connection reuse is disabled via Transport.DisableKeepAlives.</span>
    <span class="comment">// PutIdleConn is called before the caller&#39;s Response.Body.Close</span>
    <span class="comment">// call returns.</span>
    <span class="comment">// For HTTP/2, this hook is not currently used.</span>
    PutIdleConn func(err <a href="/pkg/builtin/#error">error</a>)

<span id="ClientTrace.GotFirstResponseByte"></span>    <span class="comment">// GotFirstResponseByte is called when the first byte of the response</span>
    <span class="comment">// headers is available.</span>
    GotFirstResponseByte func()

<span id="ClientTrace.Got100Continue"></span>    <span class="comment">// Got100Continue is called if the server replies with a &#34;100</span>
    <span class="comment">// Continue&#34; response.</span>
    Got100Continue func()

<span id="ClientTrace.Got1xxResponse"></span>    <span class="comment">// Got1xxResponse is called for each 1xx informational response header</span>
    <span class="comment">// returned before the final non-1xx response. Got1xxResponse is called</span>
    <span class="comment">// for &#34;100 Continue&#34; responses, even if Got100Continue is also defined.</span>
    <span class="comment">// If it returns an error, the client request is aborted with that error value.</span>
    Got1xxResponse func(code <a href="/pkg/builtin/#int">int</a>, header <a href="/pkg/net/textproto/">textproto</a>.<a href="/pkg/net/textproto/#MIMEHeader">MIMEHeader</a>) <a href="/pkg/builtin/#error">error</a>

<span id="ClientTrace.DNSStart"></span>    <span class="comment">// DNSStart is called when a DNS lookup begins.</span>
    DNSStart func(<a href="#DNSStartInfo">DNSStartInfo</a>)

<span id="ClientTrace.DNSDone"></span>    <span class="comment">// DNSDone is called when a DNS lookup ends.</span>
    DNSDone func(<a href="#DNSDoneInfo">DNSDoneInfo</a>)

<span id="ClientTrace.ConnectStart"></span>    <span class="comment">// ConnectStart is called when a new connection&#39;s Dial begins.</span>
    <span class="comment">// If net.Dialer.DualStack (IPv6 &#34;Happy Eyeballs&#34;) support is</span>
    <span class="comment">// enabled, this may be called multiple times.</span>
    ConnectStart func(network, addr <a href="/pkg/builtin/#string">string</a>)

<span id="ClientTrace.ConnectDone"></span>    <span class="comment">// ConnectDone is called when a new connection&#39;s Dial</span>
    <span class="comment">// completes. The provided err indicates whether the</span>
    <span class="comment">// connection completedly successfully.</span>
    <span class="comment">// If net.Dialer.DualStack (&#34;Happy Eyeballs&#34;) support is</span>
    <span class="comment">// enabled, this may be called multiple times.</span>
    ConnectDone func(network, addr <a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)

<span id="ClientTrace.TLSHandshakeStart"></span>    <span class="comment">// TLSHandshakeStart is called when the TLS handshake is started. When</span>
    <span class="comment">// connecting to a HTTPS site via a HTTP proxy, the handshake happens after</span>
    <span class="comment">// the CONNECT request is processed by the proxy.</span>
    TLSHandshakeStart func()

<span id="ClientTrace.TLSHandshakeDone"></span>    <span class="comment">// TLSHandshakeDone is called after the TLS handshake with either the</span>
    <span class="comment">// successful handshake&#39;s connection state, or a non-nil error on handshake</span>
    <span class="comment">// failure.</span>
    TLSHandshakeDone func(<a href="/pkg/crypto/tls/">tls</a>.<a href="/pkg/crypto/tls/#ConnectionState">ConnectionState</a>, <a href="/pkg/builtin/#error">error</a>)

<span id="ClientTrace.WroteHeaderField"></span>    <span class="comment">// WroteHeaderField is called after the Transport has written</span>
    <span class="comment">// each request header. At the time of this call the values</span>
    <span class="comment">// might be buffered and not yet written to the network.</span>
    WroteHeaderField func(key <a href="/pkg/builtin/#string">string</a>, value []<a href="/pkg/builtin/#string">string</a>)

<span id="ClientTrace.WroteHeaders"></span>    <span class="comment">// WroteHeaders is called after the Transport has written</span>
    <span class="comment">// all request headers.</span>
    WroteHeaders func()

<span id="ClientTrace.Wait100Continue"></span>    <span class="comment">// Wait100Continue is called if the Request specified</span>
    <span class="comment">// &#34;Expect: 100-continue&#34; and the Transport has written the</span>
    <span class="comment">// request headers but is waiting for &#34;100 Continue&#34; from the</span>
    <span class="comment">// server before writing the request body.</span>
    Wait100Continue func()

<span id="ClientTrace.WroteRequest"></span>    <span class="comment">// WroteRequest is called with the result of writing the</span>
    <span class="comment">// request and any body. It may be called multiple times</span>
    <span class="comment">// in the case of retried requests.</span>
    WroteRequest func(<a href="#WroteRequestInfo">WroteRequestInfo</a>)
}
</pre>









### <a id="ContextClientTrace">func</a> [ContextClientTrace](https://golang.org/src/net/http/httptrace/trace.go?s=560:617#L14)
<pre>func ContextClientTrace(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>) *<a href="#ClientTrace">ClientTrace</a></pre>
ContextClientTrace returns the ClientTrace associated with the
provided context. If none, it returns nil.






## <a id="DNSDoneInfo">type</a> [DNSDoneInfo](https://golang.org/src/net/http/httptrace/trace.go?s=7116:7492#L206)
DNSDoneInfo contains information about the results of a DNS lookup.


<pre>type DNSDoneInfo struct {
<span id="DNSDoneInfo.Addrs"></span>    <span class="comment">// Addrs are the IPv4 and/or IPv6 addresses found in the DNS</span>
    <span class="comment">// lookup. The contents of the slice should not be mutated.</span>
    Addrs []<a href="/pkg/net/">net</a>.<a href="/pkg/net/#IPAddr">IPAddr</a>

<span id="DNSDoneInfo.Err"></span>    <span class="comment">// Err is any error that occurred during the DNS lookup.</span>
    Err <a href="/pkg/builtin/#error">error</a>

<span id="DNSDoneInfo.Coalesced"></span>    <span class="comment">// Coalesced is whether the Addrs were shared with another</span>
    <span class="comment">// caller who was doing the same DNS lookup concurrently.</span>
    Coalesced <a href="/pkg/builtin/#bool">bool</a>
}
</pre>











## <a id="DNSStartInfo">type</a> [DNSStartInfo](https://golang.org/src/net/http/httptrace/trace.go?s=7002:7043#L201)
DNSStartInfo contains information about a DNS request.


<pre>type DNSStartInfo struct {
<span id="DNSStartInfo.Host"></span>    Host <a href="/pkg/builtin/#string">string</a>
}
</pre>











## <a id="GotConnInfo">type</a> [GotConnInfo](https://golang.org/src/net/http/httptrace/trace.go?s=7794:8303#L228)
GotConnInfo is the argument to the ClientTrace.GotConn function and
contains information about the obtained connection.


<pre>type GotConnInfo struct {
<span id="GotConnInfo.Conn"></span>    <span class="comment">// Conn is the connection that was obtained. It is owned by</span>
    <span class="comment">// the http.Transport and should not be read, written or</span>
    <span class="comment">// closed by users of ClientTrace.</span>
    Conn <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Conn">Conn</a>

<span id="GotConnInfo.Reused"></span>    <span class="comment">// Reused is whether this connection has been previously</span>
    <span class="comment">// used for another HTTP request.</span>
    Reused <a href="/pkg/builtin/#bool">bool</a>

<span id="GotConnInfo.WasIdle"></span>    <span class="comment">// WasIdle is whether this connection was obtained from an</span>
    <span class="comment">// idle pool.</span>
    WasIdle <a href="/pkg/builtin/#bool">bool</a>

<span id="GotConnInfo.IdleTime"></span>    <span class="comment">// IdleTime reports how long the connection was previously</span>
    <span class="comment">// idle, if WasIdle is true.</span>
    IdleTime <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>
}
</pre>











## <a id="WroteRequestInfo">type</a> [WroteRequestInfo](https://golang.org/src/net/http/httptrace/trace.go?s=5900:6003#L158)
WroteRequestInfo contains information provided to the WroteRequest
hook.


<pre>type WroteRequestInfo struct {
<span id="WroteRequestInfo.Err"></span>    <span class="comment">// Err is any error encountered while writing the Request.</span>
    Err <a href="/pkg/builtin/#error">error</a>
}
</pre>














