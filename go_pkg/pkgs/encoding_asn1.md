

# asn1
`import "encoding/asn1"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package asn1 implements parsing of DER-encoded ASN.1 data structures,
as defined in ITU-T Rec X.690.

See also ``A Layman's Guide to a Subset of ASN.1, BER, and DER,''
<a href="http://luca.ntop.org/Teaching/Appunti/asn1.html">http://luca.ntop.org/Teaching/Appunti/asn1.html</a>.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func Marshal(val interface{}) ([]byte, error)](#Marshal)
* [func MarshalWithParams(val interface{}, params string) ([]byte, error)](#MarshalWithParams)
* [func Unmarshal(b []byte, val interface{}) (rest []byte, err error)](#Unmarshal)
* [func UnmarshalWithParams(b []byte, val interface{}, params string) (rest []byte, err error)](#UnmarshalWithParams)
* [type BitString](#BitString)
  * [func (b BitString) At(i int) int](#BitString.At)
  * [func (b BitString) RightAlign() []byte](#BitString.RightAlign)
* [type Enumerated](#Enumerated)
* [type Flag](#Flag)
* [type ObjectIdentifier](#ObjectIdentifier)
  * [func (oi ObjectIdentifier) Equal(other ObjectIdentifier) bool](#ObjectIdentifier.Equal)
  * [func (oi ObjectIdentifier) String() string](#ObjectIdentifier.String)
* [type RawContent](#RawContent)
* [type RawValue](#RawValue)
* [type StructuralError](#StructuralError)
  * [func (e StructuralError) Error() string](#StructuralError.Error)
* [type SyntaxError](#SyntaxError)
  * [func (e SyntaxError) Error() string](#SyntaxError.Error)




#### <a id="pkg-files">Package files</a>
[asn1.go](https://golang.org/src/encoding/asn1/asn1.go) [common.go](https://golang.org/src/encoding/asn1/common.go) [marshal.go](https://golang.org/src/encoding/asn1/marshal.go) 


## <a id="pkg-constants">Constants</a>
ASN.1 tags represent the type of the following object.


<pre>const (
    <span id="TagBoolean">TagBoolean</span>         = 1
    <span id="TagInteger">TagInteger</span>         = 2
    <span id="TagBitString">TagBitString</span>       = 3
    <span id="TagOctetString">TagOctetString</span>     = 4
    <span id="TagNull">TagNull</span>            = 5
    <span id="TagOID">TagOID</span>             = 6
    <span id="TagEnum">TagEnum</span>            = 10
    <span id="TagUTF8String">TagUTF8String</span>      = 12
    <span id="TagSequence">TagSequence</span>        = 16
    <span id="TagSet">TagSet</span>             = 17
    <span id="TagNumericString">TagNumericString</span>   = 18
    <span id="TagPrintableString">TagPrintableString</span> = 19
    <span id="TagT61String">TagT61String</span>       = 20
    <span id="TagIA5String">TagIA5String</span>       = 22
    <span id="TagUTCTime">TagUTCTime</span>         = 23
    <span id="TagGeneralizedTime">TagGeneralizedTime</span> = 24
    <span id="TagGeneralString">TagGeneralString</span>   = 27
)</pre>ASN.1 class types represent the namespace of the tag.


<pre>const (
    <span id="ClassUniversal">ClassUniversal</span>       = 0
    <span id="ClassApplication">ClassApplication</span>     = 1
    <span id="ClassContextSpecific">ClassContextSpecific</span> = 2
    <span id="ClassPrivate">ClassPrivate</span>         = 3
)</pre>

## <a id="pkg-variables">Variables</a>
NullBytes contains bytes representing the DER-encoded ASN.1 NULL type.


<pre>var <span id="NullBytes">NullBytes</span> = []<a href="/pkg/builtin/#byte">byte</a>{<a href="#TagNull">TagNull</a>, 0}</pre>NullRawValue is a RawValue with its Tag set to the ASN.1 NULL type tag (5).


<pre>var <span id="NullRawValue">NullRawValue</span> = <a href="#RawValue">RawValue</a>{<a href="#RawValue.Tag">Tag</a>: <a href="#TagNull">TagNull</a>}</pre>

## <a id="Marshal">func</a> [Marshal](https://golang.org/src/encoding/asn1/marshal.go?s=15427:15472#L667)
<pre>func Marshal(val interface{}) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Marshal returns the ASN.1 encoding of val.

In addition to the struct tags recognised by Unmarshal, the following can be
used:


	ia5:         causes strings to be marshaled as ASN.1, IA5String values
	omitempty:   causes empty slices to be skipped
	printable:   causes strings to be marshaled as ASN.1, PrintableString values
	utf8:        causes strings to be marshaled as ASN.1, UTF8String values
	utc:         causes time.Time to be marshaled as ASN.1, UTCTime values
	generalized: causes time.Time to be marshaled as ASN.1, GeneralizedTime values



## <a id="MarshalWithParams">func</a> [MarshalWithParams](https://golang.org/src/encoding/asn1/marshal.go?s=15658:15728#L673)
<pre>func MarshalWithParams(val interface{}, params <a href="/pkg/builtin/#string">string</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
MarshalWithParams allows field parameters to be specified for the
top-level element. The form of the params is the same as the field tags.



## <a id="Unmarshal">func</a> [Unmarshal](https://golang.org/src/encoding/asn1/asn1.go?s=29745:29811#L1043)
<pre>func Unmarshal(b []<a href="/pkg/builtin/#byte">byte</a>, val interface{}) (rest []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Unmarshal parses the DER-encoded ASN.1 data structure b
and uses the reflect package to fill in an arbitrary value pointed at by val.
Because Unmarshal uses the reflect package, the structs
being written to must use upper case field names.

An ASN.1 INTEGER can be written to an int, int32, int64,
or *big.Int (from the math/big package).
If the encoded value does not fit in the Go type,
Unmarshal returns a parse error.

An ASN.1 BIT STRING can be written to a BitString.

An ASN.1 OCTET STRING can be written to a []byte.

An ASN.1 OBJECT IDENTIFIER can be written to an
ObjectIdentifier.

An ASN.1 ENUMERATED can be written to an Enumerated.

An ASN.1 UTCTIME or GENERALIZEDTIME can be written to a time.Time.

An ASN.1 PrintableString, IA5String, or NumericString can be written to a string.

Any of the above ASN.1 values can be written to an interface{}.
The value stored in the interface has the corresponding Go type.
For integers, that type is int64.

An ASN.1 SEQUENCE OF x or SET OF x can be written
to a slice if an x can be written to the slice's element type.

An ASN.1 SEQUENCE or SET can be written to a struct
if each of the elements in the sequence can be
written to the corresponding element in the struct.

The following tags on struct fields have special meaning to Unmarshal:


	application specifies that an APPLICATION tag is used
	private     specifies that a PRIVATE tag is used
	default:x   sets the default value for optional integer fields (only used if optional is also present)
	explicit    specifies that an additional, explicit tag wraps the implicit one
	optional    marks the field as ASN.1 OPTIONAL
	set         causes a SET, rather than a SEQUENCE type to be expected
	tag:x       specifies the ASN.1 tag number; implies ASN.1 CONTEXT SPECIFIC

If the type of the first field of a structure is RawContent then the raw
ASN1 contents of the struct will be stored in it.

If the type name of a slice element ends with "SET" then it's treated as if
the "set" tag was set on it. This can be used with nested slices where a
struct tag cannot be given.

Other ASN.1 types are not supported; if it encounters them,
Unmarshal returns a parse error.



## <a id="UnmarshalWithParams">func</a> [UnmarshalWithParams](https://golang.org/src/encoding/asn1/asn1.go?s=30004:30095#L1049)
<pre>func UnmarshalWithParams(b []<a href="/pkg/builtin/#byte">byte</a>, val interface{}, params <a href="/pkg/builtin/#string">string</a>) (rest []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
UnmarshalWithParams allows field parameters to be specified for the
top-level element. The form of the params is the same as the field tags.





## <a id="BitString">type</a> [BitString](https://golang.org/src/encoding/asn1/asn1.go?s=4194:4301#L148)
BitString is the structure to use when you want an ASN.1 BIT STRING type. A
bit string is padded up to the nearest byte in memory and the number of
valid bits is recorded. Padding bits will be zero.


<pre>type BitString struct {
<span id="BitString.Bytes"></span>    Bytes     []<a href="/pkg/builtin/#byte">byte</a> <span class="comment">// bits packed into bytes.</span>
<span id="BitString.BitLength"></span>    BitLength <a href="/pkg/builtin/#int">int</a>    <span class="comment">// length in bits.</span>
}
</pre>











### <a id="BitString.At">func</a> (BitString) [At](https://golang.org/src/encoding/asn1/asn1.go?s=4395:4427#L155)
<pre>func (b <a href="#BitString">BitString</a>) At(i <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
At returns the bit at the given index. If the index is out of range it
returns false.




### <a id="BitString.RightAlign">func</a> (BitString) [RightAlign](https://golang.org/src/encoding/asn1/asn1.go?s=4667:4705#L166)
<pre>func (b <a href="#BitString">BitString</a>) RightAlign() []<a href="/pkg/builtin/#byte">byte</a></pre>
RightAlign returns a slice where the padding bits are at the beginning. The
slice may share memory with the BitString.




## <a id="Enumerated">type</a> [Enumerated](https://golang.org/src/encoding/asn1/asn1.go?s=7446:7465#L284)
An Enumerated is represented as a plain int.


<pre>type Enumerated <a href="/pkg/builtin/#int">int</a></pre>











## <a id="Flag">type</a> [Flag](https://golang.org/src/encoding/asn1/asn1.go?s=7534:7548#L289)
A Flag accepts any data and is set to true if present.


<pre>type Flag <a href="/pkg/builtin/#bool">bool</a></pre>











## <a id="ObjectIdentifier">type</a> [ObjectIdentifier](https://golang.org/src/encoding/asn1/asn1.go?s=5834:5861#L211)
An ObjectIdentifier represents an ASN.1 OBJECT IDENTIFIER.


<pre>type ObjectIdentifier []<a href="/pkg/builtin/#int">int</a></pre>











### <a id="ObjectIdentifier.Equal">func</a> (ObjectIdentifier) [Equal](https://golang.org/src/encoding/asn1/asn1.go?s=5932:5993#L214)
<pre>func (oi <a href="#ObjectIdentifier">ObjectIdentifier</a>) Equal(other <a href="#ObjectIdentifier">ObjectIdentifier</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Equal reports whether oi and other represent the same identifier.




### <a id="ObjectIdentifier.String">func</a> (ObjectIdentifier) [String](https://golang.org/src/encoding/asn1/asn1.go?s=6139:6181#L227)
<pre>func (oi <a href="#ObjectIdentifier">ObjectIdentifier</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="RawContent">type</a> [RawContent](https://golang.org/src/encoding/asn1/asn1.go?s=12926:12948#L479)
RawContent is used to signal that the undecoded, DER data needs to be
preserved for a struct. To use it, the first field of the struct must have
this type. It's an error for any of the other fields to have this type.


<pre>type RawContent []<a href="/pkg/builtin/#byte">byte</a></pre>











## <a id="RawValue">type</a> [RawValue](https://golang.org/src/encoding/asn1/asn1.go?s=12572:12698#L469)
A RawValue represents an undecoded ASN.1 object.


<pre>type RawValue struct {
<span id="RawValue.Class"></span>    Class, Tag <a href="/pkg/builtin/#int">int</a>
<span id="RawValue.IsCompound"></span>    IsCompound <a href="/pkg/builtin/#bool">bool</a>
<span id="RawValue.Bytes"></span>    Bytes      []<a href="/pkg/builtin/#byte">byte</a>
<span id="RawValue.FullBytes"></span>    FullBytes  []<a href="/pkg/builtin/#byte">byte</a> <span class="comment">// includes the tag and length</span>
}
</pre>











## <a id="StructuralError">type</a> [StructuralError](https://golang.org/src/encoding/asn1/asn1.go?s=1177:1220#L25)
A StructuralError suggests that the ASN.1 data is valid, but the Go type
which is receiving it doesn't match.


<pre>type StructuralError struct {
<span id="StructuralError.Msg"></span>    Msg <a href="/pkg/builtin/#string">string</a>
}
</pre>











### <a id="StructuralError.Error">func</a> (StructuralError) [Error](https://golang.org/src/encoding/asn1/asn1.go?s=1222:1261#L29)
<pre>func (e <a href="#StructuralError">StructuralError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="SyntaxError">type</a> [SyntaxError](https://golang.org/src/encoding/asn1/asn1.go?s=1366:1405#L32)
A SyntaxError suggests that the ASN.1 data is invalid.


<pre>type SyntaxError struct {
<span id="SyntaxError.Msg"></span>    Msg <a href="/pkg/builtin/#string">string</a>
}
</pre>











### <a id="SyntaxError.Error">func</a> (SyntaxError) [Error](https://golang.org/src/encoding/asn1/asn1.go?s=1407:1442#L36)
<pre>func (e <a href="#SyntaxError">SyntaxError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>







