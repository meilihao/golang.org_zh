

# cgi
`import "net/http/cgi"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package cgi implements CGI (Common Gateway Interface) as specified
in RFC 3875.

Note that using CGI means starting a new process to handle each
request, which is typically less efficient than using a
long-running server. This package is intended primarily for
compatibility with existing systems.




## <a id="pkg-index">Index</a>
* [func Request() (*http.Request, error)](#Request)
* [func RequestFromMap(params map[string]string) (*http.Request, error)](#RequestFromMap)
* [func Serve(handler http.Handler) error](#Serve)
* [type Handler](#Handler)
  * [func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request)](#Handler.ServeHTTP)




#### <a id="pkg-files">Package files</a>
[child.go](https://golang.org/src/net/http/cgi/child.go) [host.go](https://golang.org/src/net/http/cgi/host.go) 






## <a id="Request">func</a> [Request](https://golang.org/src/net/http/cgi/child.go?s=604:641#L19)
<pre>func Request() (*<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Request returns the HTTP request as represented in the current
environment. This assumes the current program is being run
by a web server in a CGI environment.
The returned Request's Body is populated, if applicable.



## <a id="RequestFromMap">func</a> [RequestFromMap](https://golang.org/src/net/http/cgi/child.go?s=1163:1231#L42)
<pre>func RequestFromMap(params map[<a href="/pkg/builtin/#string">string</a>]<a href="/pkg/builtin/#string">string</a>) (*<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
RequestFromMap creates an http.Request from CGI variables.
The returned Request's Body field is not populated.



## <a id="Serve">func</a> [Serve](https://golang.org/src/net/http/cgi/child.go?s=3917:3955#L136)
<pre>func Serve(handler <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Handler">Handler</a>) <a href="/pkg/builtin/#error">error</a></pre>
Serve executes the provided Handler on the currently active CGI
request, if any. If there's no current CGI environment
an error is returned. The provided handler may be nil to use
http.DefaultServeMux.





## <a id="Handler">type</a> [Handler](https://golang.org/src/net/http/cgi/host.go?s=1282:2432#L37)
Handler runs an executable in a subprocess with a CGI environment.


<pre>type Handler struct {
<span id="Handler.Path"></span>    Path <a href="/pkg/builtin/#string">string</a> <span class="comment">// path to the CGI executable</span>
<span id="Handler.Root"></span>    Root <a href="/pkg/builtin/#string">string</a> <span class="comment">// root URI prefix of handler or empty for &#34;/&#34;</span>

<span id="Handler.Dir"></span>    <span class="comment">// Dir specifies the CGI executable&#39;s working directory.</span>
    <span class="comment">// If Dir is empty, the base directory of Path is used.</span>
    <span class="comment">// If Path has no base directory, the current working</span>
    <span class="comment">// directory is used.</span>
    Dir <a href="/pkg/builtin/#string">string</a>

<span id="Handler.Env"></span>    Env        []<a href="/pkg/builtin/#string">string</a>    <span class="comment">// extra environment variables to set, if any, as &#34;key=value&#34;</span>
<span id="Handler.InheritEnv"></span>    InheritEnv []<a href="/pkg/builtin/#string">string</a>    <span class="comment">// environment variables to inherit from host, as &#34;key&#34;</span>
<span id="Handler.Logger"></span>    Logger     *<a href="/pkg/log/">log</a>.<a href="/pkg/log/#Logger">Logger</a> <span class="comment">// optional log for errors or nil to use log.Print</span>
<span id="Handler.Args"></span>    Args       []<a href="/pkg/builtin/#string">string</a>    <span class="comment">// optional arguments to pass to child process</span>
<span id="Handler.Stderr"></span>    Stderr     <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>   <span class="comment">// optional stderr for the child process; nil means os.Stderr</span>

<span id="Handler.PathLocationHandler"></span>    <span class="comment">// PathLocationHandler specifies the root http Handler that</span>
    <span class="comment">// should handle internal redirects when the CGI process</span>
    <span class="comment">// returns a Location header value starting with a &#34;/&#34;, as</span>
    <span class="comment">// specified in RFC 3875 § 6.3.2. This will likely be</span>
    <span class="comment">// http.DefaultServeMux.</span>
    <span class="comment">//</span>
    <span class="comment">// If nil, a CGI response with a local URI path is instead sent</span>
    <span class="comment">// back to the client and not redirected internally.</span>
    PathLocationHandler <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Handler">Handler</a>
}
</pre>











### <a id="Handler.ServeHTTP">func</a> (\*Handler) [ServeHTTP](https://golang.org/src/net/http/cgi/host.go?s=3102:3172#L96)
<pre>func (h *<a href="#Handler">Handler</a>) ServeHTTP(rw <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#ResponseWriter">ResponseWriter</a>, req *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>)</pre>






