

# xml
`import "encoding/xml"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package xml implements a simple XML 1.0 parser that
understands XML name spaces.



<a id="example__customMarshalXML">Example (CustomMarshalXML)</a>


```go
```

output:
```txt
```

<a id="example__textMarshalXML">Example (TextMarshalXML)</a>


```go
```

output:
```txt
```


## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func Escape(w io.Writer, s []byte)](#Escape)
* [func EscapeText(w io.Writer, s []byte) error](#EscapeText)
* [func Marshal(v interface{}) ([]byte, error)](#Marshal)
* [func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)](#MarshalIndent)
* [func Unmarshal(data []byte, v interface{}) error](#Unmarshal)
* [type Attr](#Attr)
* [type CharData](#CharData)
  * [func (c CharData) Copy() CharData](#CharData.Copy)
* [type Comment](#Comment)
  * [func (c Comment) Copy() Comment](#Comment.Copy)
* [type Decoder](#Decoder)
  * [func NewDecoder(r io.Reader) *Decoder](#NewDecoder)
  * [func NewTokenDecoder(t TokenReader) *Decoder](#NewTokenDecoder)
  * [func (d *Decoder) Decode(v interface{}) error](#Decoder.Decode)
  * [func (d *Decoder) DecodeElement(v interface{}, start *StartElement) error](#Decoder.DecodeElement)
  * [func (d *Decoder) InputOffset() int64](#Decoder.InputOffset)
  * [func (d *Decoder) RawToken() (Token, error)](#Decoder.RawToken)
  * [func (d *Decoder) Skip() error](#Decoder.Skip)
  * [func (d *Decoder) Token() (Token, error)](#Decoder.Token)
* [type Directive](#Directive)
  * [func (d Directive) Copy() Directive](#Directive.Copy)
* [type Encoder](#Encoder)
  * [func NewEncoder(w io.Writer) *Encoder](#NewEncoder)
  * [func (enc *Encoder) Encode(v interface{}) error](#Encoder.Encode)
  * [func (enc *Encoder) EncodeElement(v interface{}, start StartElement) error](#Encoder.EncodeElement)
  * [func (enc *Encoder) EncodeToken(t Token) error](#Encoder.EncodeToken)
  * [func (enc *Encoder) Flush() error](#Encoder.Flush)
  * [func (enc *Encoder) Indent(prefix, indent string)](#Encoder.Indent)
* [type EndElement](#EndElement)
* [type Marshaler](#Marshaler)
* [type MarshalerAttr](#MarshalerAttr)
* [type Name](#Name)
* [type ProcInst](#ProcInst)
  * [func (p ProcInst) Copy() ProcInst](#ProcInst.Copy)
* [type StartElement](#StartElement)
  * [func (e StartElement) Copy() StartElement](#StartElement.Copy)
  * [func (e StartElement) End() EndElement](#StartElement.End)
* [type SyntaxError](#SyntaxError)
  * [func (e *SyntaxError) Error() string](#SyntaxError.Error)
* [type TagPathError](#TagPathError)
  * [func (e *TagPathError) Error() string](#TagPathError.Error)
* [type Token](#Token)
  * [func CopyToken(t Token) Token](#CopyToken)
* [type TokenReader](#TokenReader)
* [type UnmarshalError](#UnmarshalError)
  * [func (e UnmarshalError) Error() string](#UnmarshalError.Error)
* [type Unmarshaler](#Unmarshaler)
* [type UnmarshalerAttr](#UnmarshalerAttr)
* [type UnsupportedTypeError](#UnsupportedTypeError)
  * [func (e *UnsupportedTypeError) Error() string](#UnsupportedTypeError.Error)


#### <a id="pkg-examples">Examples</a>
* [Encoder](#example_Encoder)
* [MarshalIndent](#example_MarshalIndent)
* [Unmarshal](#example_Unmarshal)
* [Package (CustomMarshalXML)](#example__customMarshalXML)
* [Package (TextMarshalXML)](#example__textMarshalXML)


#### <a id="pkg-files">Package files</a>
[marshal.go](https://golang.org/src/encoding/xml/marshal.go) [read.go](https://golang.org/src/encoding/xml/read.go) [typeinfo.go](https://golang.org/src/encoding/xml/typeinfo.go) [xml.go](https://golang.org/src/encoding/xml/xml.go) 


## <a id="pkg-constants">Constants</a>

<pre>const (
    <span class="comment">// Header is a generic XML header suitable for use with the output of Marshal.</span>
    <span class="comment">// This is not automatically added to any output of this package,</span>
    <span class="comment">// it is provided as a convenience.</span>
    <span id="Header">Header</span> = `&lt;?xml version=&#34;1.0&#34; encoding=&#34;UTF-8&#34;?&gt;` + &#34;\n&#34;
)</pre>

## <a id="pkg-variables">Variables</a>
HTMLAutoClose is the set of HTML elements that
should be considered to close automatically.

See the Decoder.Strict and Decoder.Entity fields' documentation.


<pre>var <span id="HTMLAutoClose">HTMLAutoClose</span> []<a href="/pkg/builtin/#string">string</a> = htmlAutoClose</pre>HTMLEntity is an entity map containing translations for the
standard HTML entity characters.

See the Decoder.Strict and Decoder.Entity fields' documentation.


<pre>var <span id="HTMLEntity">HTMLEntity</span> map[<a href="/pkg/builtin/#string">string</a>]<a href="/pkg/builtin/#string">string</a> = htmlEntity</pre>

## <a id="Escape">func</a> [Escape](https://golang.org/src/encoding/xml/xml.go?s=46103:46137#L1975)
<pre>func Escape(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, s []<a href="/pkg/builtin/#byte">byte</a>)</pre>
Escape is like EscapeText but omits the error return value.
It is provided for backwards compatibility with Go 1.0.
Code targeting Go 1.1 or later should use EscapeText.



## <a id="EscapeText">func</a> [EscapeText](https://golang.org/src/encoding/xml/xml.go?s=44221:44265#L1881)
<pre>func EscapeText(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, s []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
EscapeText writes to w the properly escaped XML equivalent
of the plain text data s.



## <a id="Marshal">func</a> [Marshal](https://golang.org/src/encoding/xml/marshal.go?s=3363:3406#L69)
<pre>func Marshal(v interface{}) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Marshal returns the XML encoding of v.

Marshal handles an array or slice by marshaling each of the elements.
Marshal handles a pointer by marshaling the value it points at or, if the
pointer is nil, by writing nothing. Marshal handles an interface value by
marshaling the value it contains or, if the interface value is nil, by
writing nothing. Marshal handles all other data by writing one or more XML
elements containing the data.

The name for the XML elements is taken from, in order of preference:


	- the tag on the XMLName field, if the data is a struct
	- the value of the XMLName field of type Name
	- the tag of the struct field used to obtain the data
	- the name of the struct field used to obtain the data
	- the name of the marshaled type

The XML element for a struct contains marshaled elements for each of the
exported fields of the struct, with these exceptions:


	- the XMLName field, described above, is omitted.
	- a field with tag "-" is omitted.
	- a field with tag "name,attr" becomes an attribute with
	  the given name in the XML element.
	- a field with tag ",attr" becomes an attribute with the
	  field name in the XML element.
	- a field with tag ",chardata" is written as character data,
	  not as an XML element.
	- a field with tag ",cdata" is written as character data
	  wrapped in one or more <![CDATA[ ... ]]> tags, not as an XML element.
	- a field with tag ",innerxml" is written verbatim, not subject
	  to the usual marshaling procedure.
	- a field with tag ",comment" is written as an XML comment, not
	  subject to the usual marshaling procedure. It must not contain
	  the "--" string within it.
	- a field with a tag including the "omitempty" option is omitted
	  if the field value is empty. The empty values are false, 0, any
	  nil pointer or interface value, and any array, slice, map, or
	  string of length zero.
	- an anonymous struct field is handled as if the fields of its
	  value were part of the outer struct.
	- a field implementing Marshaler is written by calling its MarshalXML
	  method.
	- a field implementing encoding.TextMarshaler is written by encoding the
	  result of its MarshalText method as text.

If a field uses a tag "a>b>c", then the element c will be nested inside
parent elements a and b. Fields that appear next to each other that name
the same parent will be enclosed in one XML element.

If the XML name for a struct field is defined by both the field tag and the
struct's XMLName field, the names must match.

See MarshalIndent for an example.

Marshal will return an error if asked to marshal a channel, function, or map.



## <a id="MarshalIndent">func</a> [MarshalIndent](https://golang.org/src/encoding/xml/marshal.go?s=5222:5294#L115)
<pre>func MarshalIndent(v interface{}, prefix, indent <a href="/pkg/builtin/#string">string</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
MarshalIndent works like Marshal, but each XML element begins on a new
indented line that starts with prefix and is followed by one or more
copies of indent according to the nesting depth.



<a id="example_MarshalIndent">Example</a>


```go
```

output:
```txt
```

## <a id="Unmarshal">func</a> [Unmarshal](https://golang.org/src/encoding/xml/read.go?s=5770:5818#L122)
<pre>func Unmarshal(data []<a href="/pkg/builtin/#byte">byte</a>, v interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Unmarshal parses the XML-encoded data and stores the result in
the value pointed to by v, which must be an arbitrary struct,
slice, or string. Well-formed data that does not fit into v is
discarded.

Because Unmarshal uses the reflect package, it can only assign
to exported (upper case) fields. Unmarshal uses a case-sensitive
comparison to match XML element names to tag values and struct
field names.

Unmarshal maps an XML element to a struct using the following rules.
In the rules, the tag of a field refers to the value associated with the
key 'xml' in the struct field's tag (see the example above).


	* If the struct has a field of type []byte or string with tag
	   ",innerxml", Unmarshal accumulates the raw XML nested inside the
	   element in that field. The rest of the rules still apply.
	
	* If the struct has a field named XMLName of type Name,
	   Unmarshal records the element name in that field.
	
	* If the XMLName field has an associated tag of the form
	   "name" or "namespace-URL name", the XML element must have
	   the given name (and, optionally, name space) or else Unmarshal
	   returns an error.
	
	* If the XML element has an attribute whose name matches a
	   struct field name with an associated tag containing ",attr" or
	   the explicit name in a struct field tag of the form "name,attr",
	   Unmarshal records the attribute value in that field.
	
	* If the XML element has an attribute not handled by the previous
	   rule and the struct has a field with an associated tag containing
	   ",any,attr", Unmarshal records the attribute value in the first
	   such field.
	
	* If the XML element contains character data, that data is
	   accumulated in the first struct field that has tag ",chardata".
	   The struct field may have type []byte or string.
	   If there is no such field, the character data is discarded.
	
	* If the XML element contains comments, they are accumulated in
	   the first struct field that has tag ",comment".  The struct
	   field may have type []byte or string. If there is no such
	   field, the comments are discarded.
	
	* If the XML element contains a sub-element whose name matches
	   the prefix of a tag formatted as "a" or "a>b>c", unmarshal
	   will descend into the XML structure looking for elements with the
	   given names, and will map the innermost elements to that struct
	   field. A tag starting with ">" is equivalent to one starting
	   with the field name followed by ">".
	
	* If the XML element contains a sub-element whose name matches
	   a struct field's XMLName tag and the struct field has no
	   explicit name tag as per the previous rule, unmarshal maps
	   the sub-element to that struct field.
	
	* If the XML element contains a sub-element whose name matches a
	   field without any mode flags (",attr", ",chardata", etc), Unmarshal
	   maps the sub-element to that struct field.
	
	* If the XML element contains a sub-element that hasn't matched any
	   of the above rules and the struct has a field with tag ",any",
	   unmarshal maps the sub-element to that struct field.
	
	* An anonymous struct field is handled as if the fields of its
	   value were part of the outer struct.
	
	* A struct field with tag "-" is never unmarshaled into.

If Unmarshal encounters a field type that implements the Unmarshaler
interface, Unmarshal calls its UnmarshalXML method to produce the value from
the XML element.  Otherwise, if the value implements
encoding.TextUnmarshaler, Unmarshal calls that value's UnmarshalText method.

Unmarshal maps an XML element to a string or []byte by saving the
concatenation of that element's character data in the string or
[]byte. The saved []byte is never nil.

Unmarshal maps an attribute value to a string or []byte by saving
the value in the string or slice.

Unmarshal maps an attribute value to an Attr by saving the attribute,
including its name, in the Attr.

Unmarshal maps an XML element or attribute value to a slice by
extending the length of the slice and mapping the element or attribute
to the newly created value.

Unmarshal maps an XML element or attribute value to a bool by
setting it to the boolean value represented by the string. Whitespace
is trimmed and ignored.

Unmarshal maps an XML element or attribute value to an integer or
floating-point field by setting the field to the result of
interpreting the string value in decimal. There is no check for
overflow. Whitespace is trimmed and ignored.

Unmarshal maps an XML element to a Name by recording the element
name.

Unmarshal maps an XML element to a pointer by setting the pointer
to a freshly allocated value and then mapping the element to that value.

A missing element or empty attribute value will be unmarshaled as a zero value.
If the field is a slice, a zero value will be appended to the field. Otherwise, the
field will be set to its zero value.



<a id="example_Unmarshal">Example</a>
<p>This example demonstrates unmarshaling an XML excerpt into a value with
some preset fields. Note that the Phone field isn&#39;t modified and that
the XML &lt;Company&gt; element is ignored. Also, the Groups field is assigned
considering the element path provided in its tag.
</p>

```go
```

output:
```txt
```



## <a id="Attr">type</a> [Attr](https://golang.org/src/encoding/xml/xml.go?s=1131:1177#L38)
An Attr represents an attribute in an XML element (Name=Value).


<pre>type Attr struct {
<span id="Attr.Name"></span>    Name  <a href="#Name">Name</a>
<span id="Attr.Value"></span>    Value <a href="/pkg/builtin/#string">string</a>
}
</pre>











## <a id="CharData">type</a> [CharData](https://golang.org/src/encoding/xml/xml.go?s=1966:1986#L74)
A CharData represents XML character data (raw text),
in which XML escape sequences have been replaced by
the characters they represent.


<pre>type CharData []<a href="/pkg/builtin/#byte">byte</a></pre>











### <a id="CharData.Copy">func</a> (CharData) [Copy](https://golang.org/src/encoding/xml/xml.go?s=2116:2149#L83)
<pre>func (c <a href="#CharData">CharData</a>) Copy() <a href="#CharData">CharData</a></pre>
Copy creates a new copy of CharData.




## <a id="Comment">type</a> [Comment](https://golang.org/src/encoding/xml/xml.go?s=2313:2332#L87)
A Comment represents an XML comment of the form <!--comment-->.
The bytes do not include the <!-- and --> comment markers.


<pre>type Comment []<a href="/pkg/builtin/#byte">byte</a></pre>











### <a id="Comment.Copy">func</a> (Comment) [Copy](https://golang.org/src/encoding/xml/xml.go?s=2373:2404#L90)
<pre>func (c <a href="#Comment">Comment</a>) Copy() <a href="#Comment">Comment</a></pre>
Copy creates a new copy of Comment.




## <a id="Decoder">type</a> [Decoder](https://golang.org/src/encoding/xml/xml.go?s=4169:6498#L147)
A Decoder represents an XML parser reading a particular input stream.
The parser assumes that its input is encoded in UTF-8.


<pre>type Decoder struct {
<span id="Decoder.Strict"></span>    <span class="comment">// Strict defaults to true, enforcing the requirements</span>
    <span class="comment">// of the XML specification.</span>
    <span class="comment">// If set to false, the parser allows input containing common</span>
    <span class="comment">// mistakes:</span>
    <span class="comment">//	* If an element is missing an end tag, the parser invents</span>
    <span class="comment">//	  end tags as necessary to keep the return values from Token</span>
    <span class="comment">//	  properly balanced.</span>
    <span class="comment">//	* In attribute values and character data, unknown or malformed</span>
    <span class="comment">//	  character entities (sequences beginning with &amp;) are left alone.</span>
    <span class="comment">//</span>
    <span class="comment">// Setting:</span>
    <span class="comment">//</span>
    <span class="comment">//	d.Strict = false</span>
    <span class="comment">//	d.AutoClose = xml.HTMLAutoClose</span>
    <span class="comment">//	d.Entity = xml.HTMLEntity</span>
    <span class="comment">//</span>
    <span class="comment">// creates a parser that can handle typical HTML.</span>
    <span class="comment">//</span>
    <span class="comment">// Strict mode does not enforce the requirements of the XML name spaces TR.</span>
    <span class="comment">// In particular it does not reject name space tags using undefined prefixes.</span>
    <span class="comment">// Such tags are recorded with the unknown prefix as the name space URL.</span>
    Strict <a href="/pkg/builtin/#bool">bool</a>

    <span class="comment">// When Strict == false, AutoClose indicates a set of elements to</span>
    <span class="comment">// consider closed immediately after they are opened, regardless</span>
    <span class="comment">// of whether an end element is present.</span>
<span id="Decoder.AutoClose"></span>    AutoClose []<a href="/pkg/builtin/#string">string</a>

<span id="Decoder.Entity"></span>    <span class="comment">// Entity can be used to map non-standard entity names to string replacements.</span>
    <span class="comment">// The parser behaves as if these standard mappings are present in the map,</span>
    <span class="comment">// regardless of the actual map content:</span>
    <span class="comment">//</span>
    <span class="comment">//	&#34;lt&#34;: &#34;&lt;&#34;,</span>
    <span class="comment">//	&#34;gt&#34;: &#34;&gt;&#34;,</span>
    <span class="comment">//	&#34;amp&#34;: &#34;&amp;&#34;,</span>
    <span class="comment">//	&#34;apos&#34;: &#34;&#39;&#34;,</span>
    <span class="comment">//	&#34;quot&#34;: `&#34;`,</span>
    Entity map[<a href="/pkg/builtin/#string">string</a>]<a href="/pkg/builtin/#string">string</a>

<span id="Decoder.CharsetReader"></span>    <span class="comment">// CharsetReader, if non-nil, defines a function to generate</span>
    <span class="comment">// charset-conversion readers, converting from the provided</span>
    <span class="comment">// non-UTF-8 charset into UTF-8. If CharsetReader is nil or</span>
    <span class="comment">// returns an error, parsing stops with an error. One of the</span>
    <span class="comment">// CharsetReader&#39;s result values must be non-nil.</span>
    CharsetReader func(charset <a href="/pkg/builtin/#string">string</a>, input <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (<a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, <a href="/pkg/builtin/#error">error</a>)

<span id="Decoder.DefaultSpace"></span>    <span class="comment">// DefaultSpace sets the default name space used for unadorned tags,</span>
    <span class="comment">// as if the entire XML stream were wrapped in an element containing</span>
    <span class="comment">// the attribute xmlns=&#34;DefaultSpace&#34;.</span>
    DefaultSpace <a href="/pkg/builtin/#string">string</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewDecoder">func</a> [NewDecoder](https://golang.org/src/encoding/xml/xml.go?s=6638:6675#L219)
<pre>func NewDecoder(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) *<a href="#Decoder">Decoder</a></pre>
NewDecoder creates a new XML parser reading from r.
If r does not implement io.ByteReader, NewDecoder will
do its own buffering.




### <a id="NewTokenDecoder">func</a> [NewTokenDecoder](https://golang.org/src/encoding/xml/xml.go?s=6895:6939#L231)
<pre>func NewTokenDecoder(t <a href="#TokenReader">TokenReader</a>) *<a href="#Decoder">Decoder</a></pre>
NewTokenDecoder creates a new XML parser using an underlying token stream.






### <a id="Decoder.Decode">func</a> (\*Decoder) [Decode](https://golang.org/src/encoding/xml/read.go?s=5973:6018#L128)
<pre>func (d *<a href="#Decoder">Decoder</a>) Decode(v interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Decode works like Unmarshal, except it reads the decoder
stream to find the start element.




### <a id="Decoder.DecodeElement">func</a> (\*Decoder) [DecodeElement](https://golang.org/src/encoding/xml/read.go?s=6293:6366#L136)
<pre>func (d *<a href="#Decoder">Decoder</a>) DecodeElement(v interface{}, start *<a href="#StartElement">StartElement</a>) <a href="/pkg/builtin/#error">error</a></pre>
DecodeElement works like Unmarshal except that it takes
a pointer to the start XML element to decode into v.
It is useful when a client reads some raw XML tokens itself
but also wants to defer to Unmarshal for some elements.




### <a id="Decoder.InputOffset">func</a> (\*Decoder) [InputOffset](https://golang.org/src/encoding/xml/xml.go?s=22360:22397#L914)
<pre>func (d *<a href="#Decoder">Decoder</a>) InputOffset() <a href="/pkg/builtin/#int64">int64</a></pre>
InputOffset returns the input stream byte offset of the current decoder position.
The offset gives the location of the end of the most recently returned token
and the beginning of the next token.




### <a id="Decoder.RawToken">func</a> (\*Decoder) [RawToken](https://golang.org/src/encoding/xml/xml.go?s=14520:14563#L531)
<pre>func (d *<a href="#Decoder">Decoder</a>) RawToken() (<a href="#Token">Token</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
RawToken is like Token but does not verify that
start and end elements match and does not translate
name space prefixes to their corresponding URLs.




### <a id="Decoder.Skip">func</a> (\*Decoder) [Skip](https://golang.org/src/encoding/xml/read.go?s=22007:22037#L730)
<pre>func (d *<a href="#Decoder">Decoder</a>) Skip() <a href="/pkg/builtin/#error">error</a></pre>
Skip reads tokens until it has consumed the end element
matching the most recent start element already consumed.
It recurs if it encounters a start element, so it can be used to
skip nested structures.
It returns nil if it finds an end element matching the start
element; otherwise it returns an error describing the problem.




### <a id="Decoder.Token">func</a> (\*Decoder) [Token](https://golang.org/src/encoding/xml/xml.go?s=8197:8237#L269)
<pre>func (d *<a href="#Decoder">Decoder</a>) Token() (<a href="#Token">Token</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Token returns the next XML token in the input stream.
At the end of the input stream, Token returns nil, io.EOF.

Slices of bytes in the returned token data refer to the
parser's internal buffer and remain valid only until the next
call to Token. To acquire a copy of the bytes, call CopyToken
or the token's Copy method.

Token expands self-closing elements such as <br/>
into separate start and end elements returned by successive calls.

Token guarantees that the StartElement and EndElement
tokens it returns are properly nested and matched:
if Token encounters an unexpected end element
or EOF before all expected end elements,
it will return an error.

Token implements XML name spaces as described by
<a href="https://www.w3.org/TR/REC-xml-names/">https://www.w3.org/TR/REC-xml-names/</a>.  Each of the
Name structures contained in the Token has the Space
set to the URL identifying its name space when known.
If Token encounters an unrecognized name space prefix,
it uses the prefix as the Space rather than report an error.




## <a id="Directive">type</a> [Directive](https://golang.org/src/encoding/xml/xml.go?s=2807:2828#L106)
A Directive represents an XML directive of the form <!text>.
The bytes do not include the <! and > markers.


<pre>type Directive []<a href="/pkg/builtin/#byte">byte</a></pre>











### <a id="Directive.Copy">func</a> (Directive) [Copy](https://golang.org/src/encoding/xml/xml.go?s=2871:2906#L109)
<pre>func (d <a href="#Directive">Directive</a>) Copy() <a href="#Directive">Directive</a></pre>
Copy creates a new copy of Directive.




## <a id="Encoder">type</a> [Encoder](https://golang.org/src/encoding/xml/marshal.go?s=5505:5539#L126)
An Encoder writes XML data to an output stream.


<pre>type Encoder struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>






<a id="example_Encoder">Example</a>


```go
```

output:
```txt
```




### <a id="NewEncoder">func</a> [NewEncoder](https://golang.org/src/encoding/xml/marshal.go?s=5595:5632#L131)
<pre>func NewEncoder(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) *<a href="#Encoder">Encoder</a></pre>
NewEncoder returns a new encoder that writes to w.






### <a id="Encoder.Encode">func</a> (\*Encoder) [Encode](https://golang.org/src/encoding/xml/marshal.go?s=6217:6264#L151)
<pre>func (enc *<a href="#Encoder">Encoder</a>) Encode(v interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Encode writes the XML encoding of v to the stream.

See the documentation for Marshal for details about the conversion
of Go values to XML.

Encode calls Flush before returning.




### <a id="Encoder.EncodeElement">func</a> (\*Encoder) [EncodeElement](https://golang.org/src/encoding/xml/marshal.go?s=6643:6717#L166)
<pre>func (enc *<a href="#Encoder">Encoder</a>) EncodeElement(v interface{}, start <a href="#StartElement">StartElement</a>) <a href="/pkg/builtin/#error">error</a></pre>
EncodeElement writes the XML encoding of v to the stream,
using start as the outermost tag in the encoding.

See the documentation for Marshal for details about the conversion
of Go values to XML.

EncodeElement calls Flush before returning.




### <a id="Encoder.EncodeToken">func</a> (\*Encoder) [EncodeToken](https://golang.org/src/encoding/xml/marshal.go?s=7625:7671#L192)
<pre>func (enc *<a href="#Encoder">Encoder</a>) EncodeToken(t <a href="#Token">Token</a>) <a href="/pkg/builtin/#error">error</a></pre>
EncodeToken writes the given XML token to the stream.
It returns an error if StartElement and EndElement tokens are not properly matched.

EncodeToken does not call Flush, because usually it is part of a larger operation
such as Encode or EncodeElement (or a custom Marshaler's MarshalXML invoked
during those), and those will call Flush when finished.
Callers that create an Encoder and then invoke EncodeToken directly, without
using Encode or EncodeElement, need to call Flush when finished to ensure
that the XML is written to the underlying writer.

EncodeToken allows writing a ProcInst with Target set to "xml" only as the first token
in the stream.




### <a id="Encoder.Flush">func</a> (\*Encoder) [Flush](https://golang.org/src/encoding/xml/marshal.go?s=10177:10210#L289)
<pre>func (enc *<a href="#Encoder">Encoder</a>) Flush() <a href="/pkg/builtin/#error">error</a></pre>
Flush flushes any buffered XML to the underlying writer.
See the EncodeToken documentation for details about when it is necessary.




### <a id="Encoder.Indent">func</a> (\*Encoder) [Indent](https://golang.org/src/encoding/xml/marshal.go?s=5922:5971#L140)
<pre>func (enc *<a href="#Encoder">Encoder</a>) Indent(prefix, indent <a href="/pkg/builtin/#string">string</a>)</pre>
Indent sets the encoder to generate XML in which each element
begins on a new indented line that starts with prefix and is followed by
one or more copies of indent according to the nesting depth.




## <a id="EndElement">type</a> [EndElement](https://golang.org/src/encoding/xml/xml.go?s=1782:1819#L67)
An EndElement represents an XML end element.


<pre>type EndElement struct {
<span id="EndElement.Name"></span>    Name <a href="#Name">Name</a>
}
</pre>











## <a id="Marshaler">type</a> [Marshaler](https://golang.org/src/encoding/xml/marshal.go?s=4324:4402#L93)
Marshaler is the interface implemented by objects that can marshal
themselves into valid XML elements.

MarshalXML encodes the receiver as zero or more XML elements.
By convention, arrays or slices are typically encoded as a sequence
of elements, one per entry.
Using start as the element tag is not required, but doing so
will enable Unmarshal to match the XML elements to the correct
struct field.
One common implementation strategy is to construct a separate
value with a layout corresponding to the desired XML and then
to encode it using e.EncodeElement.
Another common strategy is to use repeated calls to e.EncodeToken
to generate the XML output one token at a time.
The sequence of encoded tokens must make up zero or more valid
XML elements.


<pre>type Marshaler interface {
    MarshalXML(e *<a href="#Encoder">Encoder</a>, start <a href="#StartElement">StartElement</a>) <a href="/pkg/builtin/#error">error</a>
}</pre>











## <a id="MarshalerAttr">type</a> [MarshalerAttr](https://golang.org/src/encoding/xml/marshal.go?s=4949:5022#L108)
MarshalerAttr is the interface implemented by objects that can marshal
themselves into valid XML attributes.

MarshalXMLAttr returns an XML attribute with the encoded value of the receiver.
Using name as the attribute name is not required, but doing so
will enable Unmarshal to match the attribute to the correct
struct field.
If MarshalXMLAttr returns the zero attribute Attr{}, no attribute
will be generated in the output.
MarshalXMLAttr is used only for struct fields with the
"attr" option in the field tag.


<pre>type MarshalerAttr interface {
    MarshalXMLAttr(name <a href="#Name">Name</a>) (<a href="#Attr">Attr</a>, <a href="/pkg/builtin/#error">error</a>)
}</pre>











## <a id="Name">type</a> [Name](https://golang.org/src/encoding/xml/xml.go?s=1021:1062#L33)
A Name represents an XML name (Local) annotated
with a name space identifier (Space).
In tokens returned by Decoder.Token, the Space identifier
is given as a canonical URL, not the short prefix used
in the document being parsed.


<pre>type Name struct {
<span id="Name.Space"></span>    Space, Local <a href="/pkg/builtin/#string">string</a>
}
</pre>











## <a id="ProcInst">type</a> [ProcInst](https://golang.org/src/encoding/xml/xml.go?s=2521:2575#L93)
A ProcInst represents an XML processing instruction of the form <?target inst?>


<pre>type ProcInst struct {
<span id="ProcInst.Target"></span>    Target <a href="/pkg/builtin/#string">string</a>
<span id="ProcInst.Inst"></span>    Inst   []<a href="/pkg/builtin/#byte">byte</a>
}
</pre>











### <a id="ProcInst.Copy">func</a> (ProcInst) [Copy](https://golang.org/src/encoding/xml/xml.go?s=2617:2650#L99)
<pre>func (p <a href="#ProcInst">ProcInst</a>) Copy() <a href="#ProcInst">ProcInst</a></pre>
Copy creates a new copy of ProcInst.




## <a id="StartElement">type</a> [StartElement](https://golang.org/src/encoding/xml/xml.go?s=1385:1437#L48)
A StartElement represents an XML start element.


<pre>type StartElement struct {
<span id="StartElement.Name"></span>    Name <a href="#Name">Name</a>
<span id="StartElement.Attr"></span>    Attr []<a href="#Attr">Attr</a>
}
</pre>











### <a id="StartElement.Copy">func</a> (StartElement) [Copy](https://golang.org/src/encoding/xml/xml.go?s=1483:1524#L54)
<pre>func (e <a href="#StartElement">StartElement</a>) Copy() <a href="#StartElement">StartElement</a></pre>
Copy creates a new copy of StartElement.




### <a id="StartElement.End">func</a> (StartElement) [End](https://golang.org/src/encoding/xml/xml.go?s=1663:1701#L62)
<pre>func (e <a href="#StartElement">StartElement</a>) End() <a href="#EndElement">EndElement</a></pre>
End returns the corresponding XML end element.




## <a id="SyntaxError">type</a> [SyntaxError](https://golang.org/src/encoding/xml/xml.go?s=609:659#L19)
A SyntaxError represents a syntax error in the XML input stream.


<pre>type SyntaxError struct {
<span id="SyntaxError.Msg"></span>    Msg  <a href="/pkg/builtin/#string">string</a>
<span id="SyntaxError.Line"></span>    Line <a href="/pkg/builtin/#int">int</a>
}
</pre>











### <a id="SyntaxError.Error">func</a> (\*SyntaxError) [Error](https://golang.org/src/encoding/xml/xml.go?s=661:697#L24)
<pre>func (e *<a href="#SyntaxError">SyntaxError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="TagPathError">type</a> [TagPathError](https://golang.org/src/encoding/xml/typeinfo.go?s=8771:8868#L327)
A TagPathError represents an error in the unmarshaling process
caused by the use of field tags with conflicting paths.


<pre>type TagPathError struct {
<span id="TagPathError.Struct"></span>    Struct       <a href="/pkg/reflect/">reflect</a>.<a href="/pkg/reflect/#Type">Type</a>
<span id="TagPathError.Field1"></span>    Field1, Tag1 <a href="/pkg/builtin/#string">string</a>
<span id="TagPathError.Field2"></span>    Field2, Tag2 <a href="/pkg/builtin/#string">string</a>
}
</pre>











### <a id="TagPathError.Error">func</a> (\*TagPathError) [Error](https://golang.org/src/encoding/xml/typeinfo.go?s=8870:8907#L333)
<pre>func (e *<a href="#TagPathError">TagPathError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Token">type</a> [Token](https://golang.org/src/encoding/xml/xml.go?s=1310:1332#L45)
A Token is an interface holding one of the token types:
StartElement, EndElement, CharData, Comment, ProcInst, or Directive.


<pre>type Token interface{}</pre>









### <a id="CopyToken">func</a> [CopyToken](https://golang.org/src/encoding/xml/xml.go?s=2982:3011#L112)
<pre>func CopyToken(t <a href="#Token">Token</a>) <a href="#Token">Token</a></pre>
CopyToken returns a copy of a Token.






## <a id="TokenReader">type</a> [TokenReader](https://golang.org/src/encoding/xml/xml.go?s=3982:4036#L141)
A TokenReader is anything that can decode a stream of XML tokens, including a
Decoder.

When Token encounters an error or end-of-file condition after successfully
reading a token, it returns the token. It may return the (non-nil) error from
the same call or return the error (and a nil token) from a subsequent call.
An instance of this general case is that a TokenReader returning a non-nil
token at the end of the token stream may return either io.EOF or a nil error.
The next Read should return nil, io.EOF.

Implementations of Token are discouraged from returning a nil token with a
nil error. Callers should treat a return of nil, nil as indicating that
nothing happened; in particular it does not indicate EOF.


<pre>type TokenReader interface {
    Token() (<a href="#Token">Token</a>, <a href="/pkg/builtin/#error">error</a>)
}</pre>











## <a id="UnmarshalError">type</a> [UnmarshalError](https://golang.org/src/encoding/xml/read.go?s=6598:6624#L145)
An UnmarshalError represents an error in the unmarshaling process.


<pre>type UnmarshalError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="UnmarshalError.Error">func</a> (UnmarshalError) [Error](https://golang.org/src/encoding/xml/read.go?s=6626:6664#L147)
<pre>func (e <a href="#UnmarshalError">UnmarshalError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Unmarshaler">type</a> [Unmarshaler](https://golang.org/src/encoding/xml/read.go?s=7383:7465#L164)
Unmarshaler is the interface implemented by objects that can unmarshal
an XML element description of themselves.

UnmarshalXML decodes a single XML element
beginning with the given start element.
If it returns an error, the outer call to Unmarshal stops and
returns that error.
UnmarshalXML must consume exactly one XML element.
One common implementation strategy is to unmarshal into
a separate value with a layout matching the expected XML
using d.DecodeElement, and then to copy the data from
that value into the receiver.
Another common strategy is to use d.Token to process the
XML object one token at a time.
UnmarshalXML may not use d.RawToken.


<pre>type Unmarshaler interface {
    UnmarshalXML(d *<a href="#Decoder">Decoder</a>, start <a href="#StartElement">StartElement</a>) <a href="/pkg/builtin/#error">error</a>
}</pre>











## <a id="UnmarshalerAttr">type</a> [UnmarshalerAttr](https://golang.org/src/encoding/xml/read.go?s=7830:7899#L176)
UnmarshalerAttr is the interface implemented by objects that can unmarshal
an XML attribute description of themselves.

UnmarshalXMLAttr decodes a single XML attribute.
If it returns an error, the outer call to Unmarshal stops and
returns that error.
UnmarshalXMLAttr is used only for struct fields with the
"attr" option in the field tag.


<pre>type UnmarshalerAttr interface {
    UnmarshalXMLAttr(attr <a href="#Attr">Attr</a>) <a href="/pkg/builtin/#error">error</a>
}</pre>











## <a id="UnsupportedTypeError">type</a> [UnsupportedTypeError](https://golang.org/src/encoding/xml/marshal.go?s=28913:28968#L1013)
UnsupportedTypeError is returned when Marshal encounters a type
that cannot be converted into XML.


<pre>type UnsupportedTypeError struct {
<span id="UnsupportedTypeError.Type"></span>    Type <a href="/pkg/reflect/">reflect</a>.<a href="/pkg/reflect/#Type">Type</a>
}
</pre>











### <a id="UnsupportedTypeError.Error">func</a> (\*UnsupportedTypeError) [Error](https://golang.org/src/encoding/xml/marshal.go?s=28970:29015#L1017)
<pre>func (e *<a href="#UnsupportedTypeError">UnsupportedTypeError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>






