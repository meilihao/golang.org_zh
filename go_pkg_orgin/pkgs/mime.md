

# mime
`import "mime"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package mime implements parts of the MIME spec.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func AddExtensionType(ext, typ string) error](#AddExtensionType)
* [func ExtensionsByType(typ string) ([]string, error)](#ExtensionsByType)
* [func FormatMediaType(t string, param map[string]string) string](#FormatMediaType)
* [func ParseMediaType(v string) (mediatype string, params map[string]string, err error)](#ParseMediaType)
* [func TypeByExtension(ext string) string](#TypeByExtension)
* [type WordDecoder](#WordDecoder)
  * [func (d *WordDecoder) Decode(word string) (string, error)](#WordDecoder.Decode)
  * [func (d *WordDecoder) DecodeHeader(header string) (string, error)](#WordDecoder.DecodeHeader)
* [type WordEncoder](#WordEncoder)
  * [func (e WordEncoder) Encode(charset, s string) string](#WordEncoder.Encode)


#### <a id="pkg-examples">Examples</a>
* [WordDecoder.Decode](#example_WordDecoder_Decode)
* [WordDecoder.DecodeHeader](#example_WordDecoder_DecodeHeader)
* [WordEncoder.Encode](#example_WordEncoder_Encode)


#### <a id="pkg-files">Package files</a>
[encodedword.go](https://golang.org/src/mime/encodedword.go) [grammar.go](https://golang.org/src/mime/grammar.go) [mediatype.go](https://golang.org/src/mime/mediatype.go) [type.go](https://golang.org/src/mime/type.go) [type_unix.go](https://golang.org/src/mime/type_unix.go) 


## <a id="pkg-constants">Constants</a>

<pre>const (
    <span class="comment">// BEncoding represents Base64 encoding scheme as defined by RFC 2045.</span>
    <span id="BEncoding">BEncoding</span> = <a href="#WordEncoder">WordEncoder</a>(&#39;b&#39;)
    <span class="comment">// QEncoding represents the Q-encoding scheme as defined by RFC 2047.</span>
    <span id="QEncoding">QEncoding</span> = <a href="#WordEncoder">WordEncoder</a>(&#39;q&#39;)
)</pre>

## <a id="pkg-variables">Variables</a>
ErrInvalidMediaParameter is returned by ParseMediaType if
the media type value was found but there was an error parsing
the optional parameters


<pre>var <span id="ErrInvalidMediaParameter">ErrInvalidMediaParameter</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;mime: invalid media parameter&#34;)</pre>

## <a id="AddExtensionType">func</a> [AddExtensionType](https://golang.org/src/mime/type.go?s=4207:4251#L150)
<pre>func AddExtensionType(ext, typ <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
AddExtensionType sets the MIME type associated with
the extension ext to typ. The extension should begin with
a leading dot, as in ".html".



## <a id="ExtensionsByType">func</a> [ExtensionsByType](https://golang.org/src/mime/type.go?s=3786:3837#L133)
<pre>func ExtensionsByType(typ <a href="/pkg/builtin/#string">string</a>) ([]<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ExtensionsByType returns the extensions known to be associated with the MIME
type typ. The returned extensions will each begin with a leading dot, as in
".html". When typ has no associated extensions, ExtensionsByType returns an
nil slice.



## <a id="FormatMediaType">func</a> [FormatMediaType](https://golang.org/src/mime/mediatype.go?s=525:587#L10)
<pre>func FormatMediaType(t <a href="/pkg/builtin/#string">string</a>, param map[<a href="/pkg/builtin/#string">string</a>]<a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
FormatMediaType serializes mediatype t and the parameters
param as a media type conforming to RFC 2045 and RFC 2616.
The type and parameter names are written in lower-case.
When any of the arguments result in a standard violation then
FormatMediaType returns the empty string.



## <a id="ParseMediaType">func</a> [ParseMediaType](https://golang.org/src/mime/mediatype.go?s=3622:3707#L130)
<pre>func ParseMediaType(v <a href="/pkg/builtin/#string">string</a>) (mediatype <a href="/pkg/builtin/#string">string</a>, params map[<a href="/pkg/builtin/#string">string</a>]<a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ParseMediaType parses a media type value and any optional
parameters, per RFC 1521.  Media types are the values in
Content-Type and Content-Disposition headers (RFC 2183).
On success, ParseMediaType returns the media type converted
to lowercase and trimmed of white space and a non-nil map.
If there is an error parsing the optional parameter,
the media type will be returned along with the error
ErrInvalidMediaParameter.
The returned map, params, maps from the lowercase
attribute to the attribute value with its case preserved.



## <a id="TypeByExtension">func</a> [TypeByExtension](https://golang.org/src/mime/type.go?s=2782:2821#L96)
<pre>func TypeByExtension(ext <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
TypeByExtension returns the MIME type associated with the file extension ext.
The extension ext should begin with a leading dot, as in ".html".
When ext has no associated type, TypeByExtension returns "".

Extensions are looked up first case-sensitively, then case-insensitively.

The built-in table is small but on unix it is augmented by the local
system's mime.types file(s) if available under one or more of these
names:


	/etc/mime.types
	/etc/apache2/mime.types
	/etc/apache/mime.types

On Windows, MIME types are extracted from the registry.

Text types have the charset parameter set to "utf-8" by default.





## <a id="WordDecoder">type</a> [WordDecoder](https://golang.org/src/mime/encodedword.go?s=4854:5266#L177)
A WordDecoder decodes MIME headers containing RFC 2047 encoded-words.


<pre>type WordDecoder struct {
<span id="WordDecoder.CharsetReader"></span>    <span class="comment">// CharsetReader, if non-nil, defines a function to generate</span>
    <span class="comment">// charset-conversion readers, converting from the provided</span>
    <span class="comment">// charset into UTF-8.</span>
    <span class="comment">// Charsets are always lower-case. utf-8, iso-8859-1 and us-ascii charsets</span>
    <span class="comment">// are handled by default.</span>
    <span class="comment">// One of the CharsetReader&#39;s result values must be non-nil.</span>
    CharsetReader func(charset <a href="/pkg/builtin/#string">string</a>, input <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (<a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, <a href="/pkg/builtin/#error">error</a>)
}
</pre>











### <a id="WordDecoder.Decode">func</a> (\*WordDecoder) [Decode](https://golang.org/src/mime/encodedword.go?s=5312:5369#L188)
<pre>func (d *<a href="#WordDecoder">WordDecoder</a>) Decode(word <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Decode decodes an RFC 2047 encoded-word.



<a id="example_WordDecoder_Decode">Example</a>


```go
```

output:
```txt
```


### <a id="WordDecoder.DecodeHeader">func</a> (\*WordDecoder) [DecodeHeader](https://golang.org/src/mime/encodedword.go?s=6469:6534#L230)
<pre>func (d *<a href="#WordDecoder">WordDecoder</a>) DecodeHeader(header <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DecodeHeader decodes all encoded-words of the given string. It returns an
error if and only if CharsetReader of d returns an error.



<a id="example_WordDecoder_DecodeHeader">Example</a>


```go
```

output:
```txt
```


## <a id="WordEncoder">type</a> [WordEncoder](https://golang.org/src/mime/encodedword.go?s=329:350#L9)
A WordEncoder is an RFC 2047 encoded-word encoder.


<pre>type WordEncoder <a href="/pkg/builtin/#byte">byte</a></pre>











### <a id="WordEncoder.Encode">func</a> (WordEncoder) [Encode](https://golang.org/src/mime/encodedword.go?s=839:892#L25)
<pre>func (e <a href="#WordEncoder">WordEncoder</a>) Encode(charset, s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Encode returns the encoded-word form of s. If s is ASCII without special
characters, it is returned unchanged. The provided charset is the IANA
charset name of s. It is case insensitive.



<a id="example_WordEncoder_Encode">Example</a>


```go
```

output:
```txt
```





