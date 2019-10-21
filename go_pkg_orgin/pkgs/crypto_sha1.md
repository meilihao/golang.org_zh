

# sha1
`import "crypto/sha1"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package sha1 implements the SHA-1 hash algorithm as defined in RFC 3174.

SHA-1 is cryptographically broken and should not be used for secure
applications.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func New() hash.Hash](#New)
* [func Sum(data []byte) [Size]byte](#Sum)


#### <a id="pkg-examples">Examples</a>
* [New](#example_New)
* [New (File)](#example_New_file)
* [Sum](#example_Sum)


#### <a id="pkg-files">Package files</a>
[sha1.go](https://golang.org/src/crypto/sha1/sha1.go) [sha1block.go](https://golang.org/src/crypto/sha1/sha1block.go) [sha1block_amd64.go](https://golang.org/src/crypto/sha1/sha1block_amd64.go) 


## <a id="pkg-constants">Constants</a>
The blocksize of SHA-1 in bytes.


<pre>const <span id="BlockSize">BlockSize</span> = 64</pre>The size of a SHA-1 checksum in bytes.


<pre>const <span id="Size">Size</span> = 20</pre>



## <a id="New">func</a> [New](https://golang.org/src/crypto/sha1/sha1.go?s=2818:2838#L111)
<pre>func New() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a></pre>
New returns a new hash.Hash computing the SHA1 checksum. The Hash also
implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
marshal and unmarshal the internal state of the hash.


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

## <a id="Sum">func</a> [Sum](https://golang.org/src/crypto/sha1/sha1.go?s=5970:6002#L251)
<pre>func Sum(data []<a href="/pkg/builtin/#byte">byte</a>) [<a href="#Size">Size</a>]<a href="/pkg/builtin/#byte">byte</a></pre>
Sum returns the SHA-1 checksum of the data.


<a id="example_Sum">Example</a>
```go
```

output:
```txt
```






