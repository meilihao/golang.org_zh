

# cookiejar
`import "net/http/cookiejar"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package cookiejar implements an in-memory RFC 6265-compliant http.CookieJar.




## <a id="pkg-index">Index</a>
* [type Jar](#Jar)
  * [func New(o *Options) (*Jar, error)](#New)
  * [func (j *Jar) Cookies(u *url.URL) (cookies []*http.Cookie)](#Jar.Cookies)
  * [func (j *Jar) SetCookies(u *url.URL, cookies []*http.Cookie)](#Jar.SetCookies)
* [type Options](#Options)
* [type PublicSuffixList](#PublicSuffixList)


#### <a id="pkg-examples">Examples</a>
* [New](#example_New)


#### <a id="pkg-files">Package files</a>
[jar.go](https://golang.org/src/net/http/cookiejar/jar.go) [punycode.go](https://golang.org/src/net/http/cookiejar/punycode.go) 








## <a id="Jar">type</a> [Jar](https://golang.org/src/net/http/cookiejar/jar.go?s=1957:2301#L50)
Jar implements the http.CookieJar interface from the net/http package.


<pre>type Jar struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="New">func</a> [New](https://golang.org/src/net/http/cookiejar/jar.go?s=2387:2421#L67)
<pre>func New(o *<a href="#Options">Options</a>) (*<a href="#Jar">Jar</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
New returns a new cookie jar. A nil *Options is equivalent to a zero
Options.


<a id="example_New">Example</a>




### <a id="Jar.Cookies">func</a> (\*Jar) [Cookies](https://golang.org/src/net/http/cookiejar/jar.go?s=4642:4700#L144)
<pre>func (j *<a href="#Jar">Jar</a>) Cookies(u *<a href="/pkg/net/url/">url</a>.<a href="/pkg/net/url/#URL">URL</a>) (cookies []*<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Cookie">Cookie</a>)</pre>
Cookies implements the Cookies method of the http.CookieJar interface.

It returns an empty slice if the URL's scheme is not HTTP or HTTPS.




### <a id="Jar.SetCookies">func</a> (\*Jar) [SetCookies](https://golang.org/src/net/http/cookiejar/jar.go?s=6326:6386#L219)
<pre>func (j *<a href="#Jar">Jar</a>) SetCookies(u *<a href="/pkg/net/url/">url</a>.<a href="/pkg/net/url/#URL">URL</a>, cookies []*<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Cookie">Cookie</a>)</pre>
SetCookies implements the SetCookies method of the http.CookieJar interface.

It does nothing if the URL's scheme is not HTTP or HTTPS.




## <a id="Options">type</a> [Options](https://golang.org/src/net/http/cookiejar/jar.go?s=1537:1881#L39)
Options are the options for creating a new Jar.


<pre>type Options struct {
<span id="Options.PublicSuffixList"></span>    <span class="comment">// PublicSuffixList is the public suffix list that determines whether</span>
    <span class="comment">// an HTTP server can set a cookie for a domain.</span>
    <span class="comment">//</span>
    <span class="comment">// A nil value is valid and may be useful for testing but it is not</span>
    <span class="comment">// secure: it means that the HTTP server for foo.co.uk can set a cookie</span>
    <span class="comment">// for bar.co.uk.</span>
    PublicSuffixList <a href="#PublicSuffixList">PublicSuffixList</a>
}
</pre>











## <a id="PublicSuffixList">type</a> [PublicSuffixList](https://golang.org/src/net/http/cookiejar/jar.go?s=1003:1484#L24)
PublicSuffixList provides the public suffix of a domain. For example:


	- the public suffix of "example.com" is "com",
	- the public suffix of "foo1.foo2.foo3.co.uk" is "co.uk", and
	- the public suffix of "bar.pvt.k12.ma.us" is "pvt.k12.ma.us".

Implementations of PublicSuffixList must be safe for concurrent use by
multiple goroutines.

An implementation that always returns "" is valid and may be useful for
testing but it is not secure: it means that the HTTP server for foo.com can
set a cookie for bar.com.

A public suffix list implementation is in the package
golang.org/x/net/publicsuffix.


<pre>type PublicSuffixList interface {
    <span class="comment">// PublicSuffix returns the public suffix of domain.</span>
    <span class="comment">//</span>
    <span class="comment">// TODO: specify which of the caller and callee is responsible for IP</span>
    <span class="comment">// addresses, for leading and trailing dots, for case sensitivity, and</span>
    <span class="comment">// for IDN/Punycode.</span>
    PublicSuffix(domain <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a>

    <span class="comment">// String returns a description of the source of this public suffix</span>
    <span class="comment">// list. The description will typically contain something like a time</span>
    <span class="comment">// stamp or version number.</span>
    String() <a href="/pkg/builtin/#string">string</a>
}</pre>















