

# sha512
`import "crypto/sha512"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package sha512 implements the SHA-384, SHA-512, SHA-512/224, and SHA-512/256
hash algorithms as defined in FIPS 180-4.

All the hash.Hash implementations returned by this package also
implement encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
marshal and unmarshal the internal state of the hash.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func New() hash.Hash](#New)
* [func New384() hash.Hash](#New384)
* [func New512_224() hash.Hash](#New512_224)
* [func New512_256() hash.Hash](#New512_256)
* [func Sum384(data []byte) (sum384 [Size384]byte)](#Sum384)
* [func Sum512(data []byte) [Size]byte](#Sum512)
* [func Sum512_224(data []byte) (sum224 [Size224]byte)](#Sum512_224)
* [func Sum512_256(data []byte) (sum256 [Size256]byte)](#Sum512_256)




#### <a id="pkg-files">Package files</a>
[sha512.go](https://golang.org/src/crypto/sha512/sha512.go) [sha512block.go](https://golang.org/src/crypto/sha512/sha512block.go) [sha512block_amd64.go](https://golang.org/src/crypto/sha512/sha512block_amd64.go) 


## <a id="pkg-constants">Constants</a>

<pre>const (
    <span class="comment">// Size is the size, in bytes, of a SHA-512 checksum.</span>
    <span id="Size">Size</span> = 64

    <span class="comment">// Size224 is the size, in bytes, of a SHA-512/224 checksum.</span>
    <span id="Size224">Size224</span> = 28

    <span class="comment">// Size256 is the size, in bytes, of a SHA-512/256 checksum.</span>
    <span id="Size256">Size256</span> = 32

    <span class="comment">// Size384 is the size, in bytes, of a SHA-384 checksum.</span>
    <span id="Size384">Size384</span> = 48

    <span class="comment">// BlockSize is the block size, in bytes, of the SHA-512/224,</span>
    <span class="comment">// SHA-512/256, SHA-384 and SHA-512 hash functions.</span>
    <span id="BlockSize">BlockSize</span> = 128
)</pre>



## <a id="New">func</a> [New](https://golang.org/src/crypto/sha512/sha512.go?s=5606:5626#L203)
<pre>func New() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a></pre>
New returns a new hash.Hash computing the SHA-512 checksum.



## <a id="New384">func</a> [New384](https://golang.org/src/crypto/sha512/sha512.go?s=6100:6123#L224)
<pre>func New384() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a></pre>
New384 returns a new hash.Hash computing the SHA-384 checksum.



## <a id="New512_224">func</a> [New512_224](https://golang.org/src/crypto/sha512/sha512.go?s=5766:5793#L210)
<pre>func New512_224() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a></pre>
New512_224 returns a new hash.Hash computing the SHA-512/224 checksum.



## <a id="New512_256">func</a> [New512_256](https://golang.org/src/crypto/sha512/sha512.go?s=5937:5964#L217)
<pre>func New512_256() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a></pre>
New512_256 returns a new hash.Hash computing the SHA-512/256 checksum.



## <a id="Sum384">func</a> [Sum384](https://golang.org/src/crypto/sha512/sha512.go?s=8416:8463#L330)
<pre>func Sum384(data []<a href="/pkg/builtin/#byte">byte</a>) (sum384 [<a href="#Size384">Size384</a>]<a href="/pkg/builtin/#byte">byte</a>)</pre>
Sum384 returns the SHA384 checksum of the data.



## <a id="Sum512">func</a> [Sum512](https://golang.org/src/crypto/sha512/sha512.go?s=8239:8274#L322)
<pre>func Sum512(data []<a href="/pkg/builtin/#byte">byte</a>) [<a href="#Size">Size</a>]<a href="/pkg/builtin/#byte">byte</a></pre>
Sum512 returns the SHA512 checksum of the data.



## <a id="Sum512_224">func</a> [Sum512_224](https://golang.org/src/crypto/sha512/sha512.go?s=8653:8704#L340)
<pre>func Sum512_224(data []<a href="/pkg/builtin/#byte">byte</a>) (sum224 [<a href="#Size224">Size224</a>]<a href="/pkg/builtin/#byte">byte</a>)</pre>
Sum512_224 returns the Sum512/224 checksum of the data.



## <a id="Sum512_256">func</a> [Sum512_256](https://golang.org/src/crypto/sha512/sha512.go?s=8898:8949#L350)
<pre>func Sum512_256(data []<a href="/pkg/builtin/#byte">byte</a>) (sum256 [<a href="#Size256">Size256</a>]<a href="/pkg/builtin/#byte">byte</a>)</pre>
Sum512_256 returns the Sum512/256 checksum of the data.









