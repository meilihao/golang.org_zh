

# adler32
`import "hash/adler32"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package adler32 implements the Adler-32 checksum.

It is defined in RFC 1950:


	Adler-32 is composed of two sums accumulated per byte: s1 is
	the sum of all bytes, s2 is the sum of all s1 values. Both sums
	are done modulo 65521. s1 is initialized to 1, s2 to zero.  The
	Adler-32 checksum is stored as s2*65536 + s1 in most-
	significant-byte first (network) order.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func Checksum(data []byte) uint32](#Checksum)
* [func New() hash.Hash32](#New)




#### <a id="pkg-files">Package files</a>
[adler32.go](https://golang.org/src/hash/adler32/adler32.go) 


## <a id="pkg-constants">Constants</a>
The size of an Adler-32 checksum in bytes.


<pre>const <span id="Size">Size</span> = 4</pre>



## <a id="Checksum">func</a> [Checksum](https://golang.org/src/hash/adler32/adler32.go?s=3187:3220#L124)
<pre>func Checksum(data []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#uint32">uint32</a></pre>
Checksum returns the Adler-32 checksum of data.



## <a id="New">func</a> [New](https://golang.org/src/hash/adler32/adler32.go?s=1345:1367#L33)
<pre>func New() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash32">Hash32</a></pre>
New returns a new hash.Hash32 computing the Adler-32 checksum. Its
Sum method will lay the value out in big-endian byte order. The
returned Hash32 also implements encoding.BinaryMarshaler and
encoding.BinaryUnmarshaler to marshal and unmarshal the internal
state of the hash.









