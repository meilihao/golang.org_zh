

# html
`import "html"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package html provides functions for escaping and unescaping HTML text.




## <a id="pkg-index">Index</a>
* [func EscapeString(s string) string](#EscapeString)
* [func UnescapeString(s string) string](#UnescapeString)


#### <a id="pkg-examples">Examples</a>
* [EscapeString](#example_EscapeString)
* [UnescapeString](#example_UnescapeString)


#### <a id="pkg-files">Package files</a>
[entity.go](https://golang.org/src/html/entity.go) [escape.go](https://golang.org/src/html/escape.go) 






## <a id="EscapeString">func</a> [EscapeString](https://golang.org/src/html/escape.go?s=4324:4358#L168)
<pre>func EscapeString(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
EscapeString escapes special characters like "<" to become "&lt;". It
escapes only five such characters: <, >, &, ' and ".
UnescapeString(EscapeString(s)) == s always holds, but the converse isn't
always true.



<a id="example_EscapeString">Example</a>


```go
```

output:
```txt
```

## <a id="UnescapeString">func</a> [UnescapeString](https://golang.org/src/html/escape.go?s=4699:4735#L177)
<pre>func UnescapeString(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
UnescapeString unescapes entities like "&lt;" to become "<". It unescapes a
larger range of entities than EscapeString escapes. For example, "&aacute;"
unescapes to "á", as does "&#225;" and "&#xE1;".
UnescapeString(EscapeString(s)) == s always holds, but the converse isn't
always true.



<a id="example_UnescapeString">Example</a>


```go
```

output:
```txt
```






