

# json
`import "encoding/json"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package json implements encoding and decoding of JSON as defined in
RFC 7159. The mapping between JSON and Go values is described
in the documentation for the Marshal and Unmarshal functions.

See "JSON and Go" for an introduction to this package:
<a href="https://golang.org/doc/articles/json_and_go.html">https://golang.org/doc/articles/json_and_go.html</a>



<a id="example__customMarshalJSON">Example (CustomMarshalJSON)</a>


```go
```

output:
```txt
```

<a id="example__textMarshalJSON">Example (TextMarshalJSON)</a>


```go
```

output:
```txt
```


## <a id="pkg-index">Index</a>
* [func Compact(dst *bytes.Buffer, src []byte) error](#Compact)
* [func HTMLEscape(dst *bytes.Buffer, src []byte)](#HTMLEscape)
* [func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error](#Indent)
* [func Marshal(v interface{}) ([]byte, error)](#Marshal)
* [func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)](#MarshalIndent)
* [func Unmarshal(data []byte, v interface{}) error](#Unmarshal)
* [func Valid(data []byte) bool](#Valid)
* [type Decoder](#Decoder)
  * [func NewDecoder(r io.Reader) *Decoder](#NewDecoder)
  * [func (dec *Decoder) Buffered() io.Reader](#Decoder.Buffered)
  * [func (dec *Decoder) Decode(v interface{}) error](#Decoder.Decode)
  * [func (dec *Decoder) DisallowUnknownFields()](#Decoder.DisallowUnknownFields)
  * [func (dec *Decoder) More() bool](#Decoder.More)
  * [func (dec *Decoder) Token() (Token, error)](#Decoder.Token)
  * [func (dec *Decoder) UseNumber()](#Decoder.UseNumber)
* [type Delim](#Delim)
  * [func (d Delim) String() string](#Delim.String)
* [type Encoder](#Encoder)
  * [func NewEncoder(w io.Writer) *Encoder](#NewEncoder)
  * [func (enc *Encoder) Encode(v interface{}) error](#Encoder.Encode)
  * [func (enc *Encoder) SetEscapeHTML(on bool)](#Encoder.SetEscapeHTML)
  * [func (enc *Encoder) SetIndent(prefix, indent string)](#Encoder.SetIndent)
* [type InvalidUTF8Error](#InvalidUTF8Error)
  * [func (e *InvalidUTF8Error) Error() string](#InvalidUTF8Error.Error)
* [type InvalidUnmarshalError](#InvalidUnmarshalError)
  * [func (e *InvalidUnmarshalError) Error() string](#InvalidUnmarshalError.Error)
* [type Marshaler](#Marshaler)
* [type MarshalerError](#MarshalerError)
  * [func (e *MarshalerError) Error() string](#MarshalerError.Error)
  * [func (e *MarshalerError) Unwrap() error](#MarshalerError.Unwrap)
* [type Number](#Number)
  * [func (n Number) Float64() (float64, error)](#Number.Float64)
  * [func (n Number) Int64() (int64, error)](#Number.Int64)
  * [func (n Number) String() string](#Number.String)
* [type RawMessage](#RawMessage)
  * [func (m RawMessage) MarshalJSON() ([]byte, error)](#RawMessage.MarshalJSON)
  * [func (m *RawMessage) UnmarshalJSON(data []byte) error](#RawMessage.UnmarshalJSON)
* [type SyntaxError](#SyntaxError)
  * [func (e *SyntaxError) Error() string](#SyntaxError.Error)
* [type Token](#Token)
* [type UnmarshalFieldError](#UnmarshalFieldError)
  * [func (e *UnmarshalFieldError) Error() string](#UnmarshalFieldError.Error)
* [type UnmarshalTypeError](#UnmarshalTypeError)
  * [func (e *UnmarshalTypeError) Error() string](#UnmarshalTypeError.Error)
* [type Unmarshaler](#Unmarshaler)
* [type UnsupportedTypeError](#UnsupportedTypeError)
  * [func (e *UnsupportedTypeError) Error() string](#UnsupportedTypeError.Error)
* [type UnsupportedValueError](#UnsupportedValueError)
  * [func (e *UnsupportedValueError) Error() string](#UnsupportedValueError.Error)


#### <a id="pkg-examples">Examples</a>
* [Decoder](#example_Decoder)
* [Decoder.Decode (Stream)](#example_Decoder_Decode_stream)
* [Decoder.Token](#example_Decoder_Token)
* [HTMLEscape](#example_HTMLEscape)
* [Indent](#example_Indent)
* [Marshal](#example_Marshal)
* [MarshalIndent](#example_MarshalIndent)
* [RawMessage (Marshal)](#example_RawMessage_marshal)
* [RawMessage (Unmarshal)](#example_RawMessage_unmarshal)
* [Unmarshal](#example_Unmarshal)
* [Valid](#example_Valid)
* [Package (CustomMarshalJSON)](#example__customMarshalJSON)
* [Package (TextMarshalJSON)](#example__textMarshalJSON)


#### <a id="pkg-files">Package files</a>
[decode.go](https://golang.org/src/encoding/json/decode.go) [encode.go](https://golang.org/src/encoding/json/encode.go) [fold.go](https://golang.org/src/encoding/json/fold.go) [indent.go](https://golang.org/src/encoding/json/indent.go) [scanner.go](https://golang.org/src/encoding/json/scanner.go) [stream.go](https://golang.org/src/encoding/json/stream.go) [tables.go](https://golang.org/src/encoding/json/tables.go) [tags.go](https://golang.org/src/encoding/json/tags.go) 






## <a id="Compact">func</a> [Compact](https://golang.org/src/encoding/json/indent.go?s=284:333#L1)
<pre>func Compact(dst *<a href="/pkg/bytes/">bytes</a>.<a href="/pkg/bytes/#Buffer">Buffer</a>, src []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
Compact appends to dst the JSON-encoded src with
insignificant space characters elided.



## <a id="HTMLEscape">func</a> [HTMLEscape](https://golang.org/src/encoding/json/encode.go?s=7613:7659#L185)
<pre>func HTMLEscape(dst *<a href="/pkg/bytes/">bytes</a>.<a href="/pkg/bytes/#Buffer">Buffer</a>, src []<a href="/pkg/builtin/#byte">byte</a>)</pre>
HTMLEscape appends to dst the JSON-encoded src with <, >, &, U+2028 and U+2029
characters inside string literals changed to \u003c, \u003e, \u0026, \u2028, \u2029
so that the JSON will be safe to embed inside HTML <script> tags.
For historical reasons, web browsers don't honor standard HTML
escaping within <script> tags, so an alternative JSON encoding must
be used.



<a id="example_HTMLEscape">Example</a>


```go
```

output:
```txt
```

## <a id="Indent">func</a> [Indent](https://golang.org/src/encoding/json/indent.go?s=2192:2263#L69)
<pre>func Indent(dst *<a href="/pkg/bytes/">bytes</a>.<a href="/pkg/bytes/#Buffer">Buffer</a>, src []<a href="/pkg/builtin/#byte">byte</a>, prefix, indent <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Indent appends to dst an indented form of the JSON-encoded src.
Each element in a JSON object or array begins on a new,
indented line beginning with prefix followed by one or more
copies of indent according to the indentation nesting.
The data appended to dst does not begin with the prefix nor
any indentation, to make it easier to embed inside other formatted JSON data.
Although leading space characters (space, tab, carriage return, newline)
at the beginning of src are dropped, trailing space characters
at the end of src are preserved and copied to dst.
For example, if src has no trailing spaces, neither will dst;
if src ends in a trailing newline, so will dst.



<a id="example_Indent">Example</a>


```go
```

output:
```txt
```

## <a id="Marshal">func</a> [Marshal](https://golang.org/src/encoding/json/encode.go?s=6471:6514#L148)
<pre>func Marshal(v interface{}) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Marshal returns the JSON encoding of v.

Marshal traverses the value v recursively.
If an encountered value implements the Marshaler interface
and is not a nil pointer, Marshal calls its MarshalJSON method
to produce JSON. If no MarshalJSON method is present but the
value implements encoding.TextMarshaler instead, Marshal calls
its MarshalText method and encodes the result as a JSON string.
The nil pointer exception is not strictly necessary
but mimics a similar, necessary exception in the behavior of
UnmarshalJSON.

Otherwise, Marshal uses the following type-dependent default encodings:

Boolean values encode as JSON booleans.

Floating point, integer, and Number values encode as JSON numbers.

String values encode as JSON strings coerced to valid UTF-8,
replacing invalid bytes with the Unicode replacement rune.
So that the JSON will be safe to embed inside HTML <script> tags,
the string is encoded using HTMLEscape,
which replaces "<", ">", "&", U+2028, and U+2029 are escaped
to "\u003c","\u003e", "\u0026", "\u2028", and "\u2029".
This replacement can be disabled when using an Encoder,
by calling SetEscapeHTML(false).

Array and slice values encode as JSON arrays, except that
[]byte encodes as a base64-encoded string, and a nil slice
encodes as the null JSON value.

Struct values encode as JSON objects.
Each exported struct field becomes a member of the object, using the
field name as the object key, unless the field is omitted for one of the
reasons given below.

The encoding of each struct field can be customized by the format string
stored under the "json" key in the struct field's tag.
The format string gives the name of the field, possibly followed by a
comma-separated list of options. The name may be empty in order to
specify options without overriding the default field name.

The "omitempty" option specifies that the field should be omitted
from the encoding if the field has an empty value, defined as
false, 0, a nil pointer, a nil interface value, and any empty array,
slice, map, or string.

As a special case, if the field tag is "-", the field is always omitted.
Note that a field with name "-" can still be generated using the tag "-,".

Examples of struct field tags and their meanings:


	// Field appears in JSON as key "myName".
	Field int `json:"myName"`
	
	// Field appears in JSON as key "myName" and
	// the field is omitted from the object if its value is empty,
	// as defined above.
	Field int `json:"myName,omitempty"`
	
	// Field appears in JSON as key "Field" (the default), but
	// the field is skipped if empty.
	// Note the leading comma.
	Field int `json:",omitempty"`
	
	// Field is ignored by this package.
	Field int `json:"-"`
	
	// Field appears in JSON as key "-".
	Field int `json:"-,"`

The "string" option signals that a field is stored as JSON inside a
JSON-encoded string. It applies only to fields of string, floating point,
integer, or boolean types. This extra level of encoding is sometimes used
when communicating with JavaScript programs:


	Int64String int64 `json:",string"`

The key name will be used if it's a non-empty string consisting of
only Unicode letters, digits, and ASCII punctuation except quotation
marks, backslash, and comma.

Anonymous struct fields are usually marshaled as if their inner exported fields
were fields in the outer struct, subject to the usual Go visibility rules amended
as described in the next paragraph.
An anonymous struct field with a name given in its JSON tag is treated as
having that name, rather than being anonymous.
An anonymous struct field of interface type is treated the same as having
that type as its name, rather than being anonymous.

The Go visibility rules for struct fields are amended for JSON when
deciding which field to marshal or unmarshal. If there are
multiple fields at the same level, and that level is the least
nested (and would therefore be the nesting level selected by the
usual Go rules), the following extra rules apply:

1) Of those fields, if any are JSON-tagged, only tagged fields are considered,
even if there are multiple untagged fields that would otherwise conflict.

2) If there is exactly one field (tagged or not according to the first rule), that is selected.

3) Otherwise there are multiple fields, and all are ignored; no error occurs.

Handling of anonymous struct fields is new in Go 1.1.
Prior to Go 1.1, anonymous struct fields were ignored. To force ignoring of
an anonymous struct field in both current and earlier versions, give the field
a JSON tag of "-".

Map values encode as JSON objects. The map's key type must either be a
string, an integer type, or implement encoding.TextMarshaler. The map keys
are sorted and used as JSON object keys by applying the following rules,
subject to the UTF-8 coercion described for string values above:


	- keys of any string type are used directly
	- encoding.TextMarshalers are marshaled
	- integer keys are converted to strings

Pointer values encode as the value pointed to.
A nil pointer encodes as the null JSON value.

Interface values encode as the value contained in the interface.
A nil interface value encodes as the null JSON value.

Channel, complex, and function values cannot be encoded in JSON.
Attempting to encode such a value causes Marshal to return
an UnsupportedTypeError.

JSON cannot represent cyclic data structures and Marshal does not
handle them. Passing cyclic structures to Marshal will result in
an infinite recursion.



<a id="example_Marshal">Example</a>


```go
```

output:
```txt
```

## <a id="MarshalIndent">func</a> [MarshalIndent](https://golang.org/src/encoding/json/encode.go?s=6964:7036#L166)
<pre>func MarshalIndent(v interface{}, prefix, indent <a href="/pkg/builtin/#string">string</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
MarshalIndent is like Marshal but applies Indent to format the output.
Each JSON element in the output will begin on a new line beginning with prefix
followed by one or more copies of indent according to the indentation nesting.



<a id="example_MarshalIndent">Example</a>


```go
```

output:
```txt
```

## <a id="Unmarshal">func</a> [Unmarshal](https://golang.org/src/encoding/json/decode.go?s=4043:4091#L85)
<pre>func Unmarshal(data []<a href="/pkg/builtin/#byte">byte</a>, v interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Unmarshal parses the JSON-encoded data and stores the result
in the value pointed to by v. If v is nil or not a pointer,
Unmarshal returns an InvalidUnmarshalError.

Unmarshal uses the inverse of the encodings that
Marshal uses, allocating maps, slices, and pointers as necessary,
with the following additional rules:

To unmarshal JSON into a pointer, Unmarshal first handles the case of
the JSON being the JSON literal null. In that case, Unmarshal sets
the pointer to nil. Otherwise, Unmarshal unmarshals the JSON into
the value pointed at by the pointer. If the pointer is nil, Unmarshal
allocates a new value for it to point to.

To unmarshal JSON into a value implementing the Unmarshaler interface,
Unmarshal calls that value's UnmarshalJSON method, including
when the input is a JSON null.
Otherwise, if the value implements encoding.TextUnmarshaler
and the input is a JSON quoted string, Unmarshal calls that value's
UnmarshalText method with the unquoted form of the string.

To unmarshal JSON into a struct, Unmarshal matches incoming object
keys to the keys used by Marshal (either the struct field name or its tag),
preferring an exact match but also accepting a case-insensitive match. By
default, object keys which don't have a corresponding struct field are
ignored (see Decoder.DisallowUnknownFields for an alternative).

To unmarshal JSON into an interface value,
Unmarshal stores one of these in the interface value:


	bool, for JSON booleans
	float64, for JSON numbers
	string, for JSON strings
	[]interface{}, for JSON arrays
	map[string]interface{}, for JSON objects
	nil for JSON null

To unmarshal a JSON array into a slice, Unmarshal resets the slice length
to zero and then appends each element to the slice.
As a special case, to unmarshal an empty JSON array into a slice,
Unmarshal replaces the slice with a new empty slice.

To unmarshal a JSON array into a Go array, Unmarshal decodes
JSON array elements into corresponding Go array elements.
If the Go array is smaller than the JSON array,
the additional JSON array elements are discarded.
If the JSON array is smaller than the Go array,
the additional Go array elements are set to zero values.

To unmarshal a JSON object into a map, Unmarshal first establishes a map to
use. If the map is nil, Unmarshal allocates a new map. Otherwise Unmarshal
reuses the existing map, keeping existing entries. Unmarshal then stores
key-value pairs from the JSON object into the map. The map's key type must
either be a string, an integer, or implement encoding.TextUnmarshaler.

If a JSON value is not appropriate for a given target type,
or if a JSON number overflows the target type, Unmarshal
skips that field and completes the unmarshaling as best it can.
If no more serious errors are encountered, Unmarshal returns
an UnmarshalTypeError describing the earliest such error. In any
case, it's not guaranteed that all the remaining fields following
the problematic one will be unmarshaled into the target object.

The JSON null value unmarshals into an interface, map, pointer, or slice
by setting that Go value to nil. Because null is often used in JSON to mean
``not present,'' unmarshaling a JSON null into any other Go type has no effect
on the value and produces no error.

When unmarshaling quoted strings, invalid UTF-8 or
invalid UTF-16 surrogate pairs are not treated as an error.
Instead, they are replaced by the Unicode replacement
character U+FFFD.



<a id="example_Unmarshal">Example</a>


```go
```

output:
```txt
```

## <a id="Valid">func</a> [Valid](https://golang.org/src/encoding/json/scanner.go?s=648:676#L9)
<pre>func Valid(data []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Valid reports whether data is a valid JSON encoding.



<a id="example_Valid">Example</a>


```go
```

output:
```txt
```



## <a id="Decoder">type</a> [Decoder](https://golang.org/src/encoding/json/stream.go?s=276:517#L4)
A Decoder reads and decodes JSON values from an input stream.


<pre>type Decoder struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>






<a id="example_Decoder">Example</a>
<p>This example uses a Decoder to decode a stream of distinct JSON values.
</p>

```go
```

output:
```txt
```




### <a id="NewDecoder">func</a> [NewDecoder](https://golang.org/src/encoding/json/stream.go?s=683:720#L21)
<pre>func NewDecoder(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) *<a href="#Decoder">Decoder</a></pre>
NewDecoder returns a new decoder that reads from r.

The decoder introduces its own buffering and may
read data from r beyond the JSON values requested.






### <a id="Decoder.Buffered">func</a> (\*Decoder) [Buffered](https://golang.org/src/encoding/json/stream.go?s=2242:2282#L73)
<pre>func (dec *<a href="#Decoder">Decoder</a>) Buffered() <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a></pre>
Buffered returns a reader of the data remaining in the Decoder's
buffer. The reader is valid until the next call to Decode.




### <a id="Decoder.Decode">func</a> (\*Decoder) [Decode](https://golang.org/src/encoding/json/stream.go?s=1425:1472#L39)
<pre>func (dec *<a href="#Decoder">Decoder</a>) Decode(v interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Decode reads the next JSON-encoded value from its
input and stores it in the value pointed to by v.

See the documentation for Unmarshal for details about
the conversion of JSON into a Go value.



<a id="example_Decoder_Decode_stream">Example (Stream)</a>
<p>This example uses a Decoder to decode a streaming array of JSON objects.
</p>

```go
```

output:
```txt
```


### <a id="Decoder.DisallowUnknownFields">func</a> (\*Decoder) [DisallowUnknownFields](https://golang.org/src/encoding/json/stream.go?s=1132:1175#L32)
<pre>func (dec *<a href="#Decoder">Decoder</a>) DisallowUnknownFields()</pre>
DisallowUnknownFields causes the Decoder to return an error when the destination
is a struct and the input contains object keys which do not match any
non-ignored, exported fields in the destination.




### <a id="Decoder.More">func</a> (\*Decoder) [More](https://golang.org/src/encoding/json/stream.go?s=12320:12351#L471)
<pre>func (dec *<a href="#Decoder">Decoder</a>) More() <a href="/pkg/builtin/#bool">bool</a></pre>
More reports whether there is another element in the
current array or object being parsed.




### <a id="Decoder.Token">func</a> (\*Decoder) [Token](https://golang.org/src/encoding/json/stream.go?s=9461:9503#L356)
<pre>func (dec *<a href="#Decoder">Decoder</a>) Token() (<a href="#Token">Token</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Token returns the next JSON token in the input stream.
At the end of the input stream, Token returns nil, io.EOF.

Token guarantees that the delimiters [ ] { } it returns are
properly nested and matched: if Token encounters an unexpected
delimiter in the input, it will return an error.

The input stream consists of basic JSON values—bool, string,
number, and null—along with delimiters [ ] { } of type Delim
to mark the start and end of arrays and objects.
Commas and colons are elided.



<a id="example_Decoder_Token">Example</a>
<p>This example uses a Decoder to decode a stream of distinct JSON values.
</p>

```go
```

output:
```txt
```


### <a id="Decoder.UseNumber">func</a> (\*Decoder) [UseNumber](https://golang.org/src/encoding/json/stream.go?s=863:894#L27)
<pre>func (dec *<a href="#Decoder">Decoder</a>) UseNumber()</pre>
UseNumber causes the Decoder to unmarshal a number into an interface{} as a
Number instead of as a float64.




## <a id="Delim">type</a> [Delim](https://golang.org/src/encoding/json/stream.go?s=8866:8881#L339)
A Delim is a JSON array or object delimiter, one of [ ] { or }.


<pre>type Delim <a href="/pkg/builtin/#rune">rune</a></pre>











### <a id="Delim.String">func</a> (Delim) [String](https://golang.org/src/encoding/json/stream.go?s=8883:8913#L341)
<pre>func (d <a href="#Delim">Delim</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Encoder">type</a> [Encoder](https://golang.org/src/encoding/json/stream.go?s=4405:4556#L167)
An Encoder writes JSON values to an output stream.


<pre>type Encoder struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewEncoder">func</a> [NewEncoder](https://golang.org/src/encoding/json/stream.go?s=4612:4649#L178)
<pre>func NewEncoder(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) *<a href="#Encoder">Encoder</a></pre>
NewEncoder returns a new encoder that writes to w.






### <a id="Encoder.Encode">func</a> (\*Encoder) [Encode](https://golang.org/src/encoding/json/stream.go?s=4885:4932#L187)
<pre>func (enc *<a href="#Encoder">Encoder</a>) Encode(v interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Encode writes the JSON encoding of v to the stream,
followed by a newline character.

See the documentation for Marshal for details about the
conversion of Go values to JSON.




### <a id="Encoder.SetEscapeHTML">func</a> (\*Encoder) [SetEscapeHTML](https://golang.org/src/encoding/json/stream.go?s=6487:6529#L239)
<pre>func (enc *<a href="#Encoder">Encoder</a>) SetEscapeHTML(on <a href="/pkg/builtin/#bool">bool</a>)</pre>
SetEscapeHTML specifies whether problematic HTML characters
should be escaped inside JSON quoted strings.
The default behavior is to escape &, <, and > to \u0026, \u003c, and \u003e
to avoid certain safety problems that can arise when embedding JSON in HTML.

In non-HTML settings where the escaping interferes with the readability
of the output, SetEscapeHTML(false) disables this behavior.




### <a id="Encoder.SetIndent">func</a> (\*Encoder) [SetIndent](https://golang.org/src/encoding/json/stream.go?s=5964:6016#L227)
<pre>func (enc *<a href="#Encoder">Encoder</a>) SetIndent(prefix, indent <a href="/pkg/builtin/#string">string</a>)</pre>
SetIndent instructs the encoder to format each subsequent encoded
value as if indented by the package-level function Indent(dst, src, prefix, indent).
Calling SetIndent("", "") disables indentation.




## <a id="InvalidUTF8Error">type</a> [InvalidUTF8Error](https://golang.org/src/encoding/json/encode.go?s=9255:9345#L245)
Before Go 1.2, an InvalidUTF8Error was returned by Marshal when
attempting to encode a string value with invalid UTF-8 sequences.
As of Go 1.2, Marshal instead coerces the string to valid UTF-8 by
replacing invalid bytes with the Unicode replacement rune U+FFFD.

Deprecated: No longer used; kept for compatibility.


<pre>type InvalidUTF8Error struct {
<span id="InvalidUTF8Error.S"></span>    S <a href="/pkg/builtin/#string">string</a> <span class="comment">// the whole string value that caused the error</span>
}
</pre>











### <a id="InvalidUTF8Error.Error">func</a> (\*InvalidUTF8Error) [Error](https://golang.org/src/encoding/json/encode.go?s=9347:9388#L249)
<pre>func (e *<a href="#InvalidUTF8Error">InvalidUTF8Error</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="InvalidUnmarshalError">type</a> [InvalidUnmarshalError](https://golang.org/src/encoding/json/decode.go?s=6229:6285#L144)
An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
(The argument to Unmarshal must be a non-nil pointer.)


<pre>type InvalidUnmarshalError struct {
<span id="InvalidUnmarshalError.Type"></span>    Type <a href="/pkg/reflect/">reflect</a>.<a href="/pkg/reflect/#Type">Type</a>
}
</pre>











### <a id="InvalidUnmarshalError.Error">func</a> (\*InvalidUnmarshalError) [Error](https://golang.org/src/encoding/json/decode.go?s=6287:6333#L148)
<pre>func (e *<a href="#InvalidUnmarshalError">InvalidUnmarshalError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Marshaler">type</a> [Marshaler](https://golang.org/src/encoding/json/encode.go?s=8424:8483#L216)
Marshaler is the interface implemented by types that
can marshal themselves into valid JSON.


<pre>type Marshaler interface {
    MarshalJSON() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)
}</pre>











## <a id="MarshalerError">type</a> [MarshalerError](https://golang.org/src/encoding/json/encode.go?s=9547:9608#L254)
A MarshalerError represents an error from calling a MarshalJSON or MarshalText method.


<pre>type MarshalerError struct {
<span id="MarshalerError.Type"></span>    Type <a href="/pkg/reflect/">reflect</a>.<a href="/pkg/reflect/#Type">Type</a>
<span id="MarshalerError.Err"></span>    Err  <a href="/pkg/builtin/#error">error</a>
}
</pre>











### <a id="MarshalerError.Error">func</a> (\*MarshalerError) [Error](https://golang.org/src/encoding/json/encode.go?s=9610:9649#L259)
<pre>func (e *<a href="#MarshalerError">MarshalerError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="MarshalerError.Unwrap">func</a> (\*MarshalerError) [Unwrap](https://golang.org/src/encoding/json/encode.go?s=9748:9787#L263)
<pre>func (e *<a href="#MarshalerError">MarshalerError</a>) Unwrap() <a href="/pkg/builtin/#error">error</a></pre>



## <a id="Number">type</a> [Number](https://golang.org/src/encoding/json/decode.go?s=7038:7056#L177)
A Number represents a JSON number literal.


<pre>type Number <a href="/pkg/builtin/#string">string</a></pre>











### <a id="Number.Float64">func</a> (Number) [Float64](https://golang.org/src/encoding/json/decode.go?s=7206:7248#L183)
<pre>func (n <a href="#Number">Number</a>) Float64() (<a href="/pkg/builtin/#float64">float64</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Float64 returns the number as a float64.




### <a id="Number.Int64">func</a> (Number) [Int64](https://golang.org/src/encoding/json/decode.go?s=7337:7375#L188)
<pre>func (n <a href="#Number">Number</a>) Int64() (<a href="/pkg/builtin/#int64">int64</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Int64 returns the number as an int64.




### <a id="Number.String">func</a> (Number) [String](https://golang.org/src/encoding/json/decode.go?s=7108:7139#L180)
<pre>func (n <a href="#Number">Number</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns the literal text of the number.




## <a id="RawMessage">type</a> [RawMessage](https://golang.org/src/encoding/json/stream.go?s=6715:6737#L246)
RawMessage is a raw encoded JSON value.
It implements Marshaler and Unmarshaler and can
be used to delay JSON decoding or precompute a JSON encoding.


<pre>type RawMessage []<a href="/pkg/builtin/#byte">byte</a></pre>






<a id="example_RawMessage_marshal">Example (Marshal)</a>
<p>This example uses RawMessage to use a precomputed JSON during marshal.
</p>

```go
```

output:
```txt
```

<a id="example_RawMessage_unmarshal">Example (Unmarshal)</a>
<p>This example uses RawMessage to delay parsing part of a JSON message.
</p>

```go
```

output:
```txt
```






### <a id="RawMessage.MarshalJSON">func</a> (RawMessage) [MarshalJSON](https://golang.org/src/encoding/json/stream.go?s=6791:6840#L249)
<pre>func (m <a href="#RawMessage">RawMessage</a>) MarshalJSON() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
MarshalJSON returns m as the JSON encoding of m.




### <a id="RawMessage.UnmarshalJSON">func</a> (\*RawMessage) [UnmarshalJSON](https://golang.org/src/encoding/json/stream.go?s=6952:7005#L257)
<pre>func (m *<a href="#RawMessage">RawMessage</a>) UnmarshalJSON(data []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
UnmarshalJSON sets *m to a copy of data.




## <a id="SyntaxError">type</a> [SyntaxError](https://golang.org/src/encoding/json/scanner.go?s=1150:1276#L30)
A SyntaxError is a description of a JSON syntax error.


<pre>type SyntaxError struct {
<span id="SyntaxError.Offset"></span>    Offset <a href="/pkg/builtin/#int64">int64</a> <span class="comment">// error occurred after reading Offset bytes</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="SyntaxError.Error">func</a> (\*SyntaxError) [Error](https://golang.org/src/encoding/json/scanner.go?s=1278:1314#L35)
<pre>func (e *<a href="#SyntaxError">SyntaxError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Token">type</a> [Token](https://golang.org/src/encoding/json/stream.go?s=7463:7485#L277)
A Token holds a value of one of these types:


	Delim, for the four JSON delimiters [ ] { }
	bool, for JSON booleans
	float64, for JSON numbers
	Number, for JSON numbers
	string, for JSON string literals
	nil, for JSON null


<pre>type Token interface{}</pre>











## <a id="UnmarshalFieldError">type</a> [UnmarshalFieldError](https://golang.org/src/encoding/json/decode.go?s=5801:5897#L132)
An UnmarshalFieldError describes a JSON object key that
led to an unexported (and therefore unwritable) struct field.

Deprecated: No longer used; kept for compatibility.


<pre>type UnmarshalFieldError struct {
<span id="UnmarshalFieldError.Key"></span>    Key   <a href="/pkg/builtin/#string">string</a>
<span id="UnmarshalFieldError.Type"></span>    Type  <a href="/pkg/reflect/">reflect</a>.<a href="/pkg/reflect/#Type">Type</a>
<span id="UnmarshalFieldError.Field"></span>    Field <a href="/pkg/reflect/">reflect</a>.<a href="/pkg/reflect/#StructField">StructField</a>
}
</pre>











### <a id="UnmarshalFieldError.Error">func</a> (\*UnmarshalFieldError) [Error](https://golang.org/src/encoding/json/decode.go?s=5899:5943#L138)
<pre>func (e *<a href="#UnmarshalFieldError">UnmarshalFieldError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="UnmarshalTypeError">type</a> [UnmarshalTypeError](https://golang.org/src/encoding/json/decode.go?s=4921:5306#L113)
An UnmarshalTypeError describes a JSON value that was
not appropriate for a value of a specific Go type.


<pre>type UnmarshalTypeError struct {
<span id="UnmarshalTypeError.Value"></span>    Value  <a href="/pkg/builtin/#string">string</a>       <span class="comment">// description of JSON value - &#34;bool&#34;, &#34;array&#34;, &#34;number -5&#34;</span>
<span id="UnmarshalTypeError.Type"></span>    Type   <a href="/pkg/reflect/">reflect</a>.<a href="/pkg/reflect/#Type">Type</a> <span class="comment">// type of Go value it could not be assigned to</span>
<span id="UnmarshalTypeError.Offset"></span>    Offset <a href="/pkg/builtin/#int64">int64</a>        <span class="comment">// error occurred after reading Offset bytes</span>
<span id="UnmarshalTypeError.Struct"></span>    Struct <a href="/pkg/builtin/#string">string</a>       <span class="comment">// name of the struct type containing the field</span>
<span id="UnmarshalTypeError.Field"></span>    Field  <a href="/pkg/builtin/#string">string</a>       <span class="comment">// the full path from root node to the field</span>
}
</pre>











### <a id="UnmarshalTypeError.Error">func</a> (\*UnmarshalTypeError) [Error](https://golang.org/src/encoding/json/decode.go?s=5308:5351#L121)
<pre>func (e *<a href="#UnmarshalTypeError">UnmarshalTypeError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Unmarshaler">type</a> [Unmarshaler](https://golang.org/src/encoding/json/decode.go?s=4749:4808#L107)
Unmarshaler is the interface implemented by types
that can unmarshal a JSON description of themselves.
The input can be assumed to be a valid encoding of
a JSON value. UnmarshalJSON must copy the JSON data
if it wishes to retain the data after returning.

By convention, to approximate the behavior of Unmarshal itself,
Unmarshalers implement UnmarshalJSON([]byte("null")) as a no-op.


<pre>type Unmarshaler interface {
    UnmarshalJSON([]<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a>
}</pre>











## <a id="UnsupportedTypeError">type</a> [UnsupportedTypeError](https://golang.org/src/encoding/json/encode.go?s=8591:8646#L222)
An UnsupportedTypeError is returned by Marshal when attempting
to encode an unsupported value type.


<pre>type UnsupportedTypeError struct {
<span id="UnsupportedTypeError.Type"></span>    Type <a href="/pkg/reflect/">reflect</a>.<a href="/pkg/reflect/#Type">Type</a>
}
</pre>











### <a id="UnsupportedTypeError.Error">func</a> (\*UnsupportedTypeError) [Error](https://golang.org/src/encoding/json/encode.go?s=8648:8693#L226)
<pre>func (e *<a href="#UnsupportedTypeError">UnsupportedTypeError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="UnsupportedValueError">type</a> [UnsupportedValueError](https://golang.org/src/encoding/json/encode.go?s=8752:8824#L230)

<pre>type UnsupportedValueError struct {
<span id="UnsupportedValueError.Value"></span>    Value <a href="/pkg/reflect/">reflect</a>.<a href="/pkg/reflect/#Value">Value</a>
<span id="UnsupportedValueError.Str"></span>    Str   <a href="/pkg/builtin/#string">string</a>
}
</pre>











### <a id="UnsupportedValueError.Error">func</a> (\*UnsupportedValueError) [Error](https://golang.org/src/encoding/json/encode.go?s=8826:8872#L235)
<pre>func (e *<a href="#UnsupportedValueError">UnsupportedValueError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>






