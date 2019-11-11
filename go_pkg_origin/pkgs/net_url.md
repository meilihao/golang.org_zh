

# url
`import "net/url"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package url parses URLs and implements query escaping.




## <a id="pkg-index">Index</a>
* [func PathEscape(s string) string](#PathEscape)
* [func PathUnescape(s string) (string, error)](#PathUnescape)
* [func QueryEscape(s string) string](#QueryEscape)
* [func QueryUnescape(s string) (string, error)](#QueryUnescape)
* [type Error](#Error)
  * [func (e *Error) Error() string](#Error.Error)
  * [func (e *Error) Temporary() bool](#Error.Temporary)
  * [func (e *Error) Timeout() bool](#Error.Timeout)
  * [func (e *Error) Unwrap() error](#Error.Unwrap)
* [type EscapeError](#EscapeError)
  * [func (e EscapeError) Error() string](#EscapeError.Error)
* [type InvalidHostError](#InvalidHostError)
  * [func (e InvalidHostError) Error() string](#InvalidHostError.Error)
* [type URL](#URL)
  * [func Parse(rawurl string) (*URL, error)](#Parse)
  * [func ParseRequestURI(rawurl string) (*URL, error)](#ParseRequestURI)
  * [func (u *URL) EscapedPath() string](#URL.EscapedPath)
  * [func (u *URL) Hostname() string](#URL.Hostname)
  * [func (u *URL) IsAbs() bool](#URL.IsAbs)
  * [func (u *URL) MarshalBinary() (text []byte, err error)](#URL.MarshalBinary)
  * [func (u *URL) Parse(ref string) (*URL, error)](#URL.Parse)
  * [func (u *URL) Port() string](#URL.Port)
  * [func (u *URL) Query() Values](#URL.Query)
  * [func (u *URL) RequestURI() string](#URL.RequestURI)
  * [func (u *URL) ResolveReference(ref *URL) *URL](#URL.ResolveReference)
  * [func (u *URL) String() string](#URL.String)
  * [func (u *URL) UnmarshalBinary(text []byte) error](#URL.UnmarshalBinary)
* [type Userinfo](#Userinfo)
  * [func User(username string) *Userinfo](#User)
  * [func UserPassword(username, password string) *Userinfo](#UserPassword)
  * [func (u *Userinfo) Password() (string, bool)](#Userinfo.Password)
  * [func (u *Userinfo) String() string](#Userinfo.String)
  * [func (u *Userinfo) Username() string](#Userinfo.Username)
* [type Values](#Values)
  * [func ParseQuery(query string) (Values, error)](#ParseQuery)
  * [func (v Values) Add(key, value string)](#Values.Add)
  * [func (v Values) Del(key string)](#Values.Del)
  * [func (v Values) Encode() string](#Values.Encode)
  * [func (v Values) Get(key string) string](#Values.Get)
  * [func (v Values) Set(key, value string)](#Values.Set)


#### <a id="pkg-examples">Examples</a>
* [ParseQuery](#example_ParseQuery)
* [URL](#example_URL)
* [URL.EscapedPath](#example_URL_EscapedPath)
* [URL.Hostname](#example_URL_Hostname)
* [URL.IsAbs](#example_URL_IsAbs)
* [URL.MarshalBinary](#example_URL_MarshalBinary)
* [URL.Parse](#example_URL_Parse)
* [URL.Port](#example_URL_Port)
* [URL.Query](#example_URL_Query)
* [URL.RequestURI](#example_URL_RequestURI)
* [URL.ResolveReference](#example_URL_ResolveReference)
* [URL.String](#example_URL_String)
* [URL.UnmarshalBinary](#example_URL_UnmarshalBinary)
* [URL (Roundtrip)](#example_URL_roundtrip)
* [Values](#example_Values)


#### <a id="pkg-files">Package files</a>
[url.go](https://golang.org/src/net/url/url.go) 






## <a id="PathEscape">func</a> [PathEscape](https://golang.org/src/net/url/url.go?s=8053:8085#L268)
<pre>func PathEscape(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
PathEscape escapes the string so it can be safely placed inside a URL path segment,
replacing special characters (including /) with %XX sequences as needed.



## <a id="PathUnescape">func</a> [PathUnescape](https://golang.org/src/net/url/url.go?s=5543:5586#L181)
<pre>func PathUnescape(s <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
PathUnescape does the inverse transformation of PathEscape,
converting each 3-byte encoded substring of the form "%AB" into the
hex-decoded byte 0xAB. It returns an error if any % is not followed
by two hexadecimal digits.

PathUnescape is identical to QueryUnescape except that it does not
unescape '+' to ' ' (space).



## <a id="QueryEscape">func</a> [QueryEscape](https://golang.org/src/net/url/url.go?s=7811:7844#L262)
<pre>func QueryEscape(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
QueryEscape escapes the string so it can be safely placed
inside a URL query.



## <a id="QueryUnescape">func</a> [QueryUnescape](https://golang.org/src/net/url/url.go?s=5111:5155#L170)
<pre>func QueryUnescape(s <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
QueryUnescape does the inverse transformation of QueryEscape,
converting each 3-byte encoded substring of the form "%AB" into the
hex-decoded byte 0xAB.
It returns an error if any % is not followed by two hexadecimal
digits.





## <a id="Error">type</a> [Error](https://golang.org/src/net/url/url.go?s=623:679#L12)
Error reports an error and the operation and URL that caused it.


<pre>type Error struct {
<span id="Error.Op"></span>    Op  <a href="/pkg/builtin/#string">string</a>
<span id="Error.URL"></span>    URL <a href="/pkg/builtin/#string">string</a>
<span id="Error.Err"></span>    Err <a href="/pkg/builtin/#error">error</a>
}
</pre>











### <a id="Error.Error">func</a> (\*Error) [Error](https://golang.org/src/net/url/url.go?s=729:759#L19)
<pre>func (e *<a href="#Error">Error</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Error.Temporary">func</a> (\*Error) [Temporary](https://golang.org/src/net/url/url.go?s=926:958#L28)
<pre>func (e *<a href="#Error">Error</a>) Temporary() <a href="/pkg/builtin/#bool">bool</a></pre>



### <a id="Error.Timeout">func</a> (\*Error) [Timeout](https://golang.org/src/net/url/url.go?s=814:844#L21)
<pre>func (e *<a href="#Error">Error</a>) Timeout() <a href="/pkg/builtin/#bool">bool</a></pre>



### <a id="Error.Unwrap">func</a> (\*Error) [Unwrap](https://golang.org/src/net/url/url.go?s=681:711#L18)
<pre>func (e *<a href="#Error">Error</a>) Unwrap() <a href="/pkg/builtin/#error">error</a></pre>



## <a id="EscapeError">type</a> [EscapeError](https://golang.org/src/net/url/url.go?s=1586:1609#L71)

<pre>type EscapeError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="EscapeError.Error">func</a> (EscapeError) [Error](https://golang.org/src/net/url/url.go?s=1611:1646#L73)
<pre>func (e <a href="#EscapeError">EscapeError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="InvalidHostError">type</a> [InvalidHostError](https://golang.org/src/net/url/url.go?s=1709:1737#L77)

<pre>type InvalidHostError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="InvalidHostError.Error">func</a> (InvalidHostError) [Error](https://golang.org/src/net/url/url.go?s=1739:1779#L79)
<pre>func (e <a href="#InvalidHostError">InvalidHostError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="URL">type</a> [URL](https://golang.org/src/net/url/url.go?s=9835:10351#L346)
A URL represents a parsed URL (technically, a URI reference).

The general form represented is:


	[scheme:][//[userinfo@]host][/]path[?query][#fragment]

URLs that do not start with a slash after the scheme are interpreted as:


	scheme:opaque[?query][#fragment]

Note that the Path field is stored in decoded form: /%47%6f%2f becomes /Go/.
A consequence is that it is impossible to tell which slashes in the Path were
slashes in the raw URL and which were %2f. This distinction is rarely important,
but when it is, the code should use RawPath, an optional field which only gets
set if the default encoding is different from Path.

URL's String method uses the EscapedPath method to obtain the path. See the
EscapedPath method for more details.


<pre>type URL struct {
<span id="URL.Scheme"></span>    Scheme     <a href="/pkg/builtin/#string">string</a>
<span id="URL.Opaque"></span>    Opaque     <a href="/pkg/builtin/#string">string</a>    <span class="comment">// encoded opaque data</span>
<span id="URL.User"></span>    User       *<a href="#Userinfo">Userinfo</a> <span class="comment">// username and password information</span>
<span id="URL.Host"></span>    Host       <a href="/pkg/builtin/#string">string</a>    <span class="comment">// host or host:port</span>
<span id="URL.Path"></span>    Path       <a href="/pkg/builtin/#string">string</a>    <span class="comment">// path (relative paths may omit leading slash)</span>
<span id="URL.RawPath"></span>    RawPath    <a href="/pkg/builtin/#string">string</a>    <span class="comment">// encoded path hint (see EscapedPath method)</span>
<span id="URL.ForceQuery"></span>    ForceQuery <a href="/pkg/builtin/#bool">bool</a>      <span class="comment">// append a query (&#39;?&#39;) even if RawQuery is empty</span>
<span id="URL.RawQuery"></span>    RawQuery   <a href="/pkg/builtin/#string">string</a>    <span class="comment">// encoded query values, without &#39;?&#39;</span>
<span id="URL.Fragment"></span>    Fragment   <a href="/pkg/builtin/#string">string</a>    <span class="comment">// fragment for references, without &#39;#&#39;</span>
}
</pre>






<a id="example_URL">Example</a>


```go
```

output:
```txt
```

<a id="example_URL_roundtrip">Example (Roundtrip)</a>


```go
```

output:
```txt
```




### <a id="Parse">func</a> [Parse](https://golang.org/src/net/url/url.go?s=13280:13319#L462)
<pre>func Parse(rawurl <a href="/pkg/builtin/#string">string</a>) (*<a href="#URL">URL</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Parse parses rawurl into a URL structure.

The rawurl may be relative (a path, without a host) or absolute
(starting with a scheme). Trying to parse a hostname and path
without a scheme is invalid but may not necessarily return an
error, due to parsing ambiguities.




### <a id="ParseRequestURI">func</a> [ParseRequestURI](https://golang.org/src/net/url/url.go?s=13966:14015#L483)
<pre>func ParseRequestURI(rawurl <a href="/pkg/builtin/#string">string</a>) (*<a href="#URL">URL</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseRequestURI parses rawurl into a URL structure. It assumes that
rawurl was received in an HTTP request, so the rawurl is interpreted
only as an absolute URI or an absolute path.
The string rawurl is assumed not to have a #fragment suffix.
(Web browsers strip #fragment before sending the URL to a web server.)






### <a id="URL.EscapedPath">func</a> (\*URL) [EscapedPath](https://golang.org/src/net/url/url.go?s=20195:20229#L687)
<pre>func (u *<a href="#URL">URL</a>) EscapedPath() <a href="/pkg/builtin/#string">string</a></pre>
EscapedPath returns the escaped form of u.Path.
In general there are multiple possible escaped forms of any path.
EscapedPath returns u.RawPath when it is a valid escaping of u.Path.
Otherwise EscapedPath ignores u.RawPath and computes an escaped
form on its own.
The String and RequestURI methods use EscapedPath to construct
their results.
In general, code should call EscapedPath instead of
reading u.RawPath directly.



<a id="example_URL_EscapedPath">Example</a>


```go
```

output:
```txt
```


### <a id="URL.Hostname">func</a> (\*URL) [Hostname](https://golang.org/src/net/url/url.go?s=29753:29784#L1048)
<pre>func (u *<a href="#URL">URL</a>) Hostname() <a href="/pkg/builtin/#string">string</a></pre>
Hostname returns u.Host, stripping any valid port number if present.

If the result is enclosed in square brackets, as literal IPv6 addresses are,
the square brackets are removed from the result.



<a id="example_URL_Hostname">Example</a>


```go
```

output:
```txt
```


### <a id="URL.IsAbs">func</a> (\*URL) [IsAbs](https://golang.org/src/net/url/url.go?s=27318:27344#L964)
<pre>func (u *<a href="#URL">URL</a>) IsAbs() <a href="/pkg/builtin/#bool">bool</a></pre>
IsAbs reports whether the URL is absolute.
Absolute means that it has a non-empty scheme.



<a id="example_URL_IsAbs">Example</a>


```go
```

output:
```txt
```


### <a id="URL.MarshalBinary">func</a> (\*URL) [MarshalBinary](https://golang.org/src/net/url/url.go?s=30786:30840#L1082)
<pre>func (u *<a href="#URL">URL</a>) MarshalBinary() (text []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>


<a id="example_URL_MarshalBinary">Example</a>


```go
```

output:
```txt
```


### <a id="URL.Parse">func</a> (\*URL) [Parse](https://golang.org/src/net/url/url.go?s=27580:27625#L971)
<pre>func (u *<a href="#URL">URL</a>) Parse(ref <a href="/pkg/builtin/#string">string</a>) (*<a href="#URL">URL</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Parse parses a URL in the context of the receiver. The provided URL
may be relative or absolute. Parse returns nil, err on parse
failure, otherwise its return value is the same as ResolveReference.



<a id="example_URL_Parse">Example</a>


```go
```

output:
```txt
```


### <a id="URL.Port">func</a> (\*URL) [Port](https://golang.org/src/net/url/url.go?s=29989:30016#L1056)
<pre>func (u *<a href="#URL">URL</a>) Port() <a href="/pkg/builtin/#string">string</a></pre>
Port returns the port part of u.Host, without the leading colon.

If u.Host doesn't contain a valid numeric port, Port returns an empty string.



<a id="example_URL_Port">Example</a>


```go
```

output:
```txt
```


### <a id="URL.Query">func</a> (\*URL) [Query](https://golang.org/src/net/url/url.go?s=29031:29059#L1019)
<pre>func (u *<a href="#URL">URL</a>) Query() <a href="#Values">Values</a></pre>
Query parses RawQuery and returns the corresponding values.
It silently discards malformed value pairs.
To check errors use ParseQuery.



<a id="example_URL_Query">Example</a>


```go
```

output:
```txt
```


### <a id="URL.RequestURI">func</a> (\*URL) [RequestURI](https://golang.org/src/net/url/url.go?s=29223:29256#L1026)
<pre>func (u *<a href="#URL">URL</a>) RequestURI() <a href="/pkg/builtin/#string">string</a></pre>
RequestURI returns the encoded path?query or opaque?query
string that would be used in an HTTP request for u.



<a id="example_URL_RequestURI">Example</a>


```go
```

output:
```txt
```


### <a id="URL.ResolveReference">func</a> (\*URL) [ResolveReference](https://golang.org/src/net/url/url.go?s=28129:28174#L985)
<pre>func (u *<a href="#URL">URL</a>) ResolveReference(ref *<a href="#URL">URL</a>) *<a href="#URL">URL</a></pre>
ResolveReference resolves a URI reference to an absolute URI from
an absolute base URI u, per RFC 3986 Section 5.2. The URI reference
may be relative or absolute. ResolveReference always returns a new
URL instance, even if the returned URL is identical to either the
base or reference. If ref is an absolute URL, then ResolveReference
ignores base and returns a copy of ref.



<a id="example_URL_ResolveReference">Example</a>


```go
```

output:
```txt
```


### <a id="URL.String">func</a> (\*URL) [String](https://golang.org/src/net/url/url.go?s=22453:22482#L763)
<pre>func (u *<a href="#URL">URL</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String reassembles the URL into a valid URL string.
The general form of the result is one of:


	scheme:opaque?query#fragment
	scheme://userinfo@host/path?query#fragment

If u.Opaque is non-empty, String uses the first form;
otherwise it uses the second form.
Any non-ASCII characters in host are escaped.
To obtain the path, String uses u.EscapedPath().

In the second form, the following rules apply:


	- if u.Scheme is empty, scheme: is omitted.
	- if u.User is nil, userinfo@ is omitted.
	- if u.Host is empty, host/ is omitted.
	- if u.Scheme and u.Host are empty and u.User is nil,
	   the entire scheme://userinfo@host/ is omitted.
	- if u.Host is non-empty and u.Path begins with a /,
	   the form host/path does not add its own /.
	- if u.RawQuery is empty, ?query is omitted.
	- if u.Fragment is empty, #fragment is omitted.



<a id="example_URL_String">Example</a>


```go
```

output:
```txt
```


### <a id="URL.UnmarshalBinary">func</a> (\*URL) [UnmarshalBinary](https://golang.org/src/net/url/url.go?s=30878:30926#L1086)
<pre>func (u *<a href="#URL">URL</a>) UnmarshalBinary(text []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>


<a id="example_URL_UnmarshalBinary">Example</a>


```go
```

output:
```txt
```


## <a id="Userinfo">type</a> [Userinfo](https://golang.org/src/net/url/url.go?s=11258:11340#L380)
The Userinfo type is an immutable encapsulation of username and
password details for a URL. An existing Userinfo value is guaranteed
to have a username set (potentially empty, as allowed by RFC 2396),
and optionally a password.


<pre>type Userinfo struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="User">func</a> [User](https://golang.org/src/net/url/url.go?s=10437:10473#L360)
<pre>func User(username <a href="/pkg/builtin/#string">string</a>) *<a href="#Userinfo">Userinfo</a></pre>
User returns a Userinfo containing the provided username
and no password set.




### <a id="UserPassword">func</a> [UserPassword](https://golang.org/src/net/url/url.go?s=10914:10968#L372)
<pre>func UserPassword(username, password <a href="/pkg/builtin/#string">string</a>) *<a href="#Userinfo">Userinfo</a></pre>
UserPassword returns a Userinfo containing the provided username
and password.

This functionality should only be used with legacy web sites.
RFC 2396 warns that interpreting Userinfo this way
``is NOT RECOMMENDED, because the passing of authentication
information in clear text (such as URI) has proven to be a
security risk in almost every case where it has been used.''






### <a id="Userinfo.Password">func</a> (\*Userinfo) [Password](https://golang.org/src/net/url/url.go?s=11542:11586#L395)
<pre>func (u *<a href="#Userinfo">Userinfo</a>) Password() (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#bool">bool</a>)</pre>
Password returns the password in case it is set, and whether it is set.




### <a id="Userinfo.String">func</a> (\*Userinfo) [String](https://golang.org/src/net/url/url.go?s=11764:11798#L404)
<pre>func (u *<a href="#Userinfo">Userinfo</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns the encoded userinfo information in the standard form
of "username[:password]".




### <a id="Userinfo.Username">func</a> (\*Userinfo) [Username](https://golang.org/src/net/url/url.go?s=11376:11412#L387)
<pre>func (u *<a href="#Userinfo">Userinfo</a>) Username() <a href="/pkg/builtin/#string">string</a></pre>
Username returns the username.




## <a id="Values">type</a> [Values](https://golang.org/src/net/url/url.go?s=23981:24012#L816)
Values maps a string key to a list of values.
It is typically used for query parameters and form values.
Unlike in the http.Header map, the keys in a Values map
are case-sensitive.


<pre>type Values map[<a href="/pkg/builtin/#string">string</a>][]<a href="/pkg/builtin/#string">string</a></pre>






<a id="example_Values">Example</a>


```go
```

output:
```txt
```




### <a id="ParseQuery">func</a> [ParseQuery](https://golang.org/src/net/url/url.go?s=25201:25246#L859)
<pre>func ParseQuery(query <a href="/pkg/builtin/#string">string</a>) (<a href="#Values">Values</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseQuery parses the URL-encoded query string and returns
a map listing the values specified for each key.
ParseQuery always returns a non-nil map containing all the
valid query parameters found; err describes the first decoding error
encountered, if any.

Query is expected to be a list of key=value settings separated by
ampersands or semicolons. A setting without an equals sign is
interpreted as a key set to an empty value.



<a id="example_ParseQuery">Example</a>


```go
```

output:
```txt
```




### <a id="Values.Add">func</a> (Values) [Add](https://golang.org/src/net/url/url.go?s=24569:24607#L841)
<pre>func (v <a href="#Values">Values</a>) Add(key, value <a href="/pkg/builtin/#string">string</a>)</pre>
Add adds the value to key. It appends to any existing
values associated with key.




### <a id="Values.Del">func</a> (Values) [Del](https://golang.org/src/net/url/url.go?s=24692:24723#L846)
<pre>func (v <a href="#Values">Values</a>) Del(key <a href="/pkg/builtin/#string">string</a>)</pre>
Del deletes the values associated with key.




### <a id="Values.Encode">func</a> (Values) [Encode](https://golang.org/src/net/url/url.go?s=26009:26040#L901)
<pre>func (v <a href="#Values">Values</a>) Encode() <a href="/pkg/builtin/#string">string</a></pre>
Encode encodes the values into ``URL encoded'' form
("bar=baz&foo=quux") sorted by key.




### <a id="Values.Get">func</a> (Values) [Get](https://golang.org/src/net/url/url.go?s=24209:24247#L822)
<pre>func (v <a href="#Values">Values</a>) Get(key <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Get gets the first value associated with the given key.
If there are no values associated with the key, Get returns
the empty string. To access multiple values, use the map
directly.




### <a id="Values.Set">func</a> (Values) [Set](https://golang.org/src/net/url/url.go?s=24411:24449#L835)
<pre>func (v <a href="#Values">Values</a>) Set(key, value <a href="/pkg/builtin/#string">string</a>)</pre>
Set sets the key to value. It replaces any existing
values.







