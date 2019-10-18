

# crc64
`import "hash/crc64"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package crc64 implements the 64-bit cyclic redundancy check, or CRC-64,
checksum. See <a href="https://en.wikipedia.org/wiki/Cyclic_redundancy_check">https://en.wikipedia.org/wiki/Cyclic_redundancy_check</a> for
information.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func Checksum(data []byte, tab *Table) uint64](#Checksum)
* [func New(tab *Table) hash.Hash64](#New)
* [func Update(crc uint64, tab *Table, p []byte) uint64](#Update)
* [type Table](#Table)
  * [func MakeTable(poly uint64) *Table](#MakeTable)




#### <a id="pkg-files">Package files</a>
[crc64.go](https://golang.org/src/hash/crc64/crc64.go) 


## <a id="pkg-constants">Constants</a>
Predefined polynomials.


<pre>const (
    <span class="comment">// The ISO polynomial, defined in ISO 3309 and used in HDLC.</span>
    <span id="ISO">ISO</span> = 0xD800000000000000

    <span class="comment">// The ECMA polynomial, defined in ECMA 182.</span>
    <span id="ECMA">ECMA</span> = 0xC96C5795D7870F42
)</pre>The size of a CRC-64 checksum in bytes.


<pre>const <span id="Size">Size</span> = 8</pre>



## <a id="Checksum">func</a> [Checksum](https://golang.org/src/hash/crc64/crc64.go?s=5400:5445#L202)
<pre>func Checksum(data []<a href="/pkg/builtin/#byte">byte</a>, tab *<a href="#Table">Table</a>) <a href="/pkg/builtin/#uint64">uint64</a></pre>
Checksum returns the CRC-64 checksum of data
using the polynomial represented by the Table.



## <a id="New">func</a> [New](https://golang.org/src/hash/crc64/crc64.go?s=2354:2386#L90)
<pre>func New(tab *<a href="#Table">Table</a>) <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash64">Hash64</a></pre>
New creates a new hash.Hash64 computing the CRC-64 checksum using the
polynomial represented by the Table. Its Sum method will lay the
value out in big-endian byte order. The returned Hash64 also
implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
marshal and unmarshal the internal state of the hash.



## <a id="Update">func</a> [Update](https://golang.org/src/hash/crc64/crc64.go?s=4878:4930#L184)
<pre>func Update(crc <a href="/pkg/builtin/#uint64">uint64</a>, tab *<a href="#Table">Table</a>, p []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#uint64">uint64</a></pre>
Update returns the result of adding the bytes in p to the crc.





## <a id="Table">type</a> [Table](https://golang.org/src/hash/crc64/crc64.go?s=721:743#L19)
Table is a 256-word table representing the polynomial for efficient processing.


<pre>type Table [256]<a href="/pkg/builtin/#uint64">uint64</a></pre>









### <a id="MakeTable">func</a> [MakeTable](https://golang.org/src/hash/crc64/crc64.go?s=1214:1248#L38)
<pre>func MakeTable(poly <a href="/pkg/builtin/#uint64">uint64</a>) *<a href="#Table">Table</a></pre>
MakeTable returns a Table constructed from the specified polynomial.
The contents of this Table must not be modified.










