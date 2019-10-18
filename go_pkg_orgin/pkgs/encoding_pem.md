

# pem
`import "encoding/pem"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package pem implements the PEM data encoding, which originated in Privacy
Enhanced Mail. The most common use of PEM encoding today is in TLS keys and
certificates. See RFC 1421.




## <a id="pkg-index">Index</a>
* [func Encode(out io.Writer, b *Block) error](#Encode)
* [func EncodeToMemory(b *Block) []byte](#EncodeToMemory)
* [type Block](#Block)
  * [func Decode(data []byte) (p *Block, rest []byte)](#Decode)


#### <a id="pkg-examples">Examples</a>
* [Decode](#example_Decode)
* [Encode](#example_Encode)


#### <a id="pkg-files">Package files</a>
[pem.go](https://golang.org/src/encoding/pem/pem.go) 






## <a id="Encode">func</a> [Encode](https://golang.org/src/encoding/pem/pem.go?s=6968:7010#L254)
<pre>func Encode(out <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, b *<a href="#Block">Block</a>) <a href="/pkg/builtin/#error">error</a></pre>
Encode writes the PEM encoding of b to out.


<a id="example_Encode">Example</a>

## <a id="EncodeToMemory">func</a> [EncodeToMemory](https://golang.org/src/encoding/pem/pem.go?s=8702:8738#L324)
<pre>func EncodeToMemory(b *<a href="#Block">Block</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
EncodeToMemory returns the PEM encoding of b.

If b has invalid headers and cannot be encoded,
EncodeToMemory returns nil. If it is important to
report details about this error case, use Encode instead.





## <a id="Block">type</a> [Block](https://golang.org/src/encoding/pem/pem.go?s=669:934#L17)
A Block represents a PEM encoded structure.

The encoded form is:


	-----BEGIN Type-----
	Headers
	base64-encoded Bytes
	-----END Type-----

where Headers is a possibly empty sequence of Key: Value lines.


<pre>type Block struct {
<span id="Block.Type"></span>    Type    <a href="/pkg/builtin/#string">string</a>            <span class="comment">// The type, taken from the preamble (i.e. &#34;RSA PRIVATE KEY&#34;).</span>
<span id="Block.Headers"></span>    Headers map[<a href="/pkg/builtin/#string">string</a>]<a href="/pkg/builtin/#string">string</a> <span class="comment">// Optional headers.</span>
<span id="Block.Bytes"></span>    Bytes   []<a href="/pkg/builtin/#byte">byte</a>            <span class="comment">// The decoded bytes of the contents. Typically a DER encoded ASN.1 structure.</span>
}
</pre>









### <a id="Decode">func</a> [Decode](https://golang.org/src/encoding/pem/pem.go?s=2505:2553#L76)
<pre>func Decode(data []<a href="/pkg/builtin/#byte">byte</a>) (p *<a href="#Block">Block</a>, rest []<a href="/pkg/builtin/#byte">byte</a>)</pre>
Decode will find the next PEM formatted block (certificate, private key
etc) in the input. It returns that block and the remainder of the input. If
no PEM data is found, p is nil and the whole of the input is returned in
rest.


<a id="example_Decode">Example</a>








