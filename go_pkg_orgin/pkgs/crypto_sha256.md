

# sha256
`import "crypto/sha256"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package sha256 implements the SHA224 and SHA256 hash algorithms as defined
in FIPS 180-4.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func New() hash.Hash](#New)
* [func New224() hash.Hash](#New224)
* [func Sum224(data []byte) (sum224 [Size224]byte)](#Sum224)
* [func Sum256(data []byte) [Size]byte](#Sum256)


#### <a id="pkg-examples">Examples</a>
* [New](#example_New)
* [New (File)](#example_New_file)
* [Sum256](#example_Sum256)


#### <a id="pkg-files">Package files</a>
[sha256.go](https://golang.org/src/crypto/sha256/sha256.go) [sha256block.go](https://golang.org/src/crypto/sha256/sha256block.go) [sha256block_amd64.go](https://golang.org/src/crypto/sha256/sha256block_amd64.go) [sha256block_decl.go](https://golang.org/src/crypto/sha256/sha256block_decl.go) 


## <a id="pkg-constants">Constants</a>
The blocksize of SHA256 and SHA224 in bytes.


<pre>const <span id="BlockSize">BlockSize</span> = 64</pre>The size of a SHA256 checksum in bytes.


<pre>const <span id="Size">Size</span> = 32</pre>The size of a SHA224 checksum in bytes.


<pre>const <span id="Size224">Size224</span> = 28</pre>



## <a id="New">func</a> [New](https://golang.org/src/crypto/sha256/sha256.go?s=3821:3841#L151)
<pre>func New() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a></pre>
New returns a new hash.Hash computing the SHA256 checksum. The Hash
also implements encoding.BinaryMarshaler and
encoding.BinaryUnmarshaler to marshal and unmarshal the internal
state of the hash.


<a id="example_New">Example</a>
```go
```

output:
```txt
```
<a id="example_New_file">Example (File)</a>
```go
```

output:
```txt
```

## <a id="New224">func</a> [New224](https://golang.org/src/crypto/sha256/sha256.go?s=3951:3974#L158)
<pre>func New224() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a></pre>
New224 returns a new hash.Hash computing the SHA224 checksum.



## <a id="Sum224">func</a> [Sum224](https://golang.org/src/crypto/sha256/sha256.go?s=5787:5834#L252)
<pre>func Sum224(data []<a href="/pkg/builtin/#byte">byte</a>) (sum224 [<a href="#Size224">Size224</a>]<a href="/pkg/builtin/#byte">byte</a>)</pre>
Sum224 returns the SHA224 checksum of the data.



## <a id="Sum256">func</a> [Sum256](https://golang.org/src/crypto/sha256/sha256.go?s=5634:5669#L244)
<pre>func Sum256(data []<a href="/pkg/builtin/#byte">byte</a>) [<a href="#Size">Size</a>]<a href="/pkg/builtin/#byte">byte</a></pre>
Sum256 returns the SHA256 checksum of the data.


<a id="example_Sum256">Example</a>
```go
```

output:
```txt
```






