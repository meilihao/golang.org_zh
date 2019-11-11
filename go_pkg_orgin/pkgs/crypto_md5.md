

# md5
`import "crypto/md5"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package md5 implements the MD5 hash algorithm as defined in RFC 1321.

MD5 is cryptographically broken and should not be used for secure
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
[md5.go](https://golang.org/src/crypto/md5/md5.go) [md5block.go](https://golang.org/src/crypto/md5/md5block.go) [md5block_decl.go](https://golang.org/src/crypto/md5/md5block_decl.go) 


## <a id="pkg-constants">Constants</a>
The blocksize of MD5 in bytes.


<pre>const <span id="BlockSize">BlockSize</span> = 64</pre>The size of an MD5 checksum in bytes.


<pre>const <span id="Size">Size</span> = 16</pre>



## <a id="New">func</a> [New](https://golang.org/src/crypto/md5/md5.go?s=2566:2586#L103)
<pre>func New() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a></pre>
New returns a new hash.Hash computing the MD5 checksum. The Hash also
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

## <a id="Sum">func</a> [Sum](https://golang.org/src/crypto/md5/md5.go?s=4513:4545#L180)
<pre>func Sum(data []<a href="/pkg/builtin/#byte">byte</a>) [<a href="#Size">Size</a>]<a href="/pkg/builtin/#byte">byte</a></pre>
Sum returns the MD5 checksum of the data.



<a id="example_Sum">Example</a>


```go
```

output:
```txt
```






