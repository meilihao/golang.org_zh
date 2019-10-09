

# crc32
`import "hash/crc32"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package crc32 implements the 32-bit cyclic redundancy check, or CRC-32,
checksum. See <a href="https://en.wikipedia.org/wiki/Cyclic_redundancy_check">https://en.wikipedia.org/wiki/Cyclic_redundancy_check</a> for
information.

Polynomials are represented in LSB-first form also known as reversed representation.

See <a href="https://en.wikipedia.org/wiki/Mathematics_of_cyclic_redundancy_checks#Reversed_representations_and_reciprocal_polynomials">https://en.wikipedia.org/wiki/Mathematics_of_cyclic_redundancy_checks#Reversed_representations_and_reciprocal_polynomials</a>
for information.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func Checksum(data []byte, tab *Table) uint32](#Checksum)
* [func ChecksumIEEE(data []byte) uint32](#ChecksumIEEE)
* [func New(tab *Table) hash.Hash32](#New)
* [func NewIEEE() hash.Hash32](#NewIEEE)
* [func Update(crc uint32, tab *Table, p []byte) uint32](#Update)
* [type Table](#Table)
  * [func MakeTable(poly uint32) *Table](#MakeTable)


#### <a id="pkg-examples">Examples</a>
* [MakeTable](#example_MakeTable)


#### <a id="pkg-files">Package files</a>
[crc32.go](https://golang.org/src/hash/crc32/crc32.go) [crc32_amd64.go](https://golang.org/src/hash/crc32/crc32_amd64.go) [crc32_generic.go](https://golang.org/src/hash/crc32/crc32_generic.go) 


## <a id="pkg-constants">Constants</a>
Predefined polynomials.


<pre>const (
    <span class="comment">// IEEE is by far and away the most common CRC-32 polynomial.</span>
    <span class="comment">// Used by ethernet (IEEE 802.3), v.42, fddi, gzip, zip, png, ...</span>
    <span id="IEEE">IEEE</span> = 0xedb88320

    <span class="comment">// Castagnoli&#39;s polynomial, used in iSCSI.</span>
    <span class="comment">// Has better error detection characteristics than IEEE.</span>
    <span class="comment">// https://dx.doi.org/10.1109/26.231911</span>
    <span id="Castagnoli">Castagnoli</span> = 0x82f63b78

    <span class="comment">// Koopman&#39;s polynomial.</span>
    <span class="comment">// Also has better error detection characteristics than IEEE.</span>
    <span class="comment">// https://dx.doi.org/10.1109/DSN.2002.1028931</span>
    <span id="Koopman">Koopman</span> = 0xeb31d82e
)</pre>The size of a CRC-32 checksum in bytes.


<pre>const <span id="Size">Size</span> = 4</pre>

## <a id="pkg-variables">Variables</a>
IEEETable is the table for the IEEE polynomial.


<pre>var <span id="IEEETable">IEEETable</span> = simpleMakeTable(<a href="#IEEE">IEEE</a>)</pre>

## <a id="Checksum">func</a> [Checksum](https://golang.org/src/hash/crc32/crc32.go?s=7383:7428#L237)
<pre>func Checksum(data []<a href="/pkg/builtin/#byte">byte</a>, tab *<a href="#Table">Table</a>) <a href="/pkg/builtin/#uint32">uint32</a></pre>
Checksum returns the CRC-32 checksum of data
using the polynomial represented by the Table.



## <a id="ChecksumIEEE">func</a> [ChecksumIEEE](https://golang.org/src/hash/crc32/crc32.go?s=7544:7581#L241)
<pre>func ChecksumIEEE(data []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#uint32">uint32</a></pre>
ChecksumIEEE returns the CRC-32 checksum of data
using the IEEE polynomial.



## <a id="New">func</a> [New](https://golang.org/src/hash/crc32/crc32.go?s=4699:4731#L137)
<pre>func New(tab *<a href="#Table">Table</a>) <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash32">Hash32</a></pre>
New creates a new hash.Hash32 computing the CRC-32 checksum using the
polynomial represented by the Table. Its Sum method will lay the
value out in big-endian byte order. The returned Hash32 also
implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
marshal and unmarshal the internal state of the hash.



## <a id="NewIEEE">func</a> [NewIEEE](https://golang.org/src/hash/crc32/crc32.go?s=5130:5156#L149)
<pre>func NewIEEE() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash32">Hash32</a></pre>
NewIEEE creates a new hash.Hash32 computing the CRC-32 checksum using
the IEEE polynomial. Its Sum method will lay the value out in
big-endian byte order. The returned Hash32 also implements
encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal
and unmarshal the internal state of the hash.



## <a id="Update">func</a> [Update](https://golang.org/src/hash/crc32/crc32.go?s=6349:6401#L200)
<pre>func Update(crc <a href="/pkg/builtin/#uint32">uint32</a>, tab *<a href="#Table">Table</a>, p []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#uint32">uint32</a></pre>
Update returns the result of adding the bytes in p to the crc.





## <a id="Table">type</a> [Table](https://golang.org/src/hash/crc32/crc32.go?s=1280:1302#L32)
Table is a 256-word table representing the polynomial for efficient processing.


<pre>type Table [256]<a href="/pkg/builtin/#uint32">uint32</a></pre>









### <a id="MakeTable">func</a> [MakeTable](https://golang.org/src/hash/crc32/crc32.go?s=4035:4069#L114)
<pre>func MakeTable(poly <a href="/pkg/builtin/#uint32">uint32</a>) *<a href="#Table">Table</a></pre>
MakeTable returns a Table constructed from the specified polynomial.
The contents of this Table must not be modified.


<a id="example_MakeTable">Example</a>








