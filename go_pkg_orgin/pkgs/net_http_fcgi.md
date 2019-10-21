

# fcgi
`import "net/http/fcgi"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package fcgi implements the FastCGI protocol.

See <a href="https://fast-cgi.github.io/">https://fast-cgi.github.io/</a> for an unofficial mirror of the
original documentation.

Currently only the responder role is supported.




## <a id="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func ProcessEnv(r *http.Request) map[string]string](#ProcessEnv)
* [func Serve(l net.Listener, handler http.Handler) error](#Serve)




#### <a id="pkg-files">Package files</a>
[child.go](https://golang.org/src/net/http/fcgi/child.go) [fcgi.go](https://golang.org/src/net/http/fcgi/fcgi.go) 




## <a id="pkg-variables">Variables</a>
ErrConnClosed is returned by Read when a handler attempts to read the body of
a request after the connection to the web server has been closed.


<pre>var <span id="ErrConnClosed">ErrConnClosed</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;fcgi: connection to web server closed&#34;)</pre>ErrRequestAborted is returned by Read when a handler attempts to read the
body of a request that has been aborted by the web server.


<pre>var <span id="ErrRequestAborted">ErrRequestAborted</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;fcgi: request aborted by web server&#34;)</pre>

## <a id="ProcessEnv">func</a> [ProcessEnv](https://golang.org/src/net/http/fcgi/child.go?s=9146:9196#L348)
<pre>func ProcessEnv(r *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>) map[<a href="/pkg/builtin/#string">string</a>]<a href="/pkg/builtin/#string">string</a></pre>
ProcessEnv returns FastCGI environment variables associated with the request r
for which no effort was made to be included in the request itself - the data
is hidden in the request's context. As an example, if REMOTE_USER is set for a
request, it will not be found anywhere in r, but it will be included in
ProcessEnv's response (via r's context).



## <a id="Serve">func</a> [Serve](https://golang.org/src/net/http/fcgi/child.go?s=8426:8480#L321)
<pre>func Serve(l <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Listener">Listener</a>, handler <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Handler">Handler</a>) <a href="/pkg/builtin/#error">error</a></pre>
Serve accepts incoming FastCGI connections on the listener l, creating a new
goroutine for each. The goroutine reads requests and then calls handler
to reply to them.
If l is nil, Serve accepts connections from os.Stdin.
If handler is nil, http.DefaultServeMux is used.








