

# sha256
`import "crypto/sha256"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package sha256 implements the SHA224 and SHA256 hash algorithms as defined
in FIPS 180-4.

sha1 实现了定义在FIPS 180-4的SHA224 和 SHA256哈希算法.


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

SHA256 和 SHA224 blocksize(字节).

<pre>const <span id="BlockSize">BlockSize</span> = 64</pre>

The size of a SHA256 checksum in bytes.

SHA256 校验和的大小(字节).

<pre>const <span id="Size">Size</span> = 32</pre>

The size of a SHA224 checksum in bytes.

SHA224 校验和的大小(字节).


<pre>const <span id="Size224">Size224</span> = 28</pre>



## <a id="New">func</a> [New](https://golang.org/src/crypto/sha256/sha256.go?s=3821:3841#L151)
<pre>func New() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a></pre>
New returns a new hash.Hash computing the SHA256 checksum. The Hash
also implements encoding.BinaryMarshaler and
encoding.BinaryUnmarshaler to marshal and unmarshal the internal
state of the hash.

New 返回hash.Hash, 用于计算SHA256校验和. Hash也实现了 encoding.BinaryMarshaler 和 encoding.BinaryUnmarshaler, 用于marshal 和 unmarshal hash的内部状态.


<a id="example_New">Example</a>
```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("hello world\n"))
	fmt.Printf("%x", h.Sum(nil))
}
```

output:
```txt
a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447
```
<a id="example_New_file">Example (File)</a>
```go
package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))
}
```

output:
```txt
2009/11/10 23:00:00 open file.txt: No such file or directory
```

## <a id="New224">func</a> [New224](https://golang.org/src/crypto/sha256/sha256.go?s=3951:3974#L158)
<pre>func New224() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a></pre>
New224 returns a new hash.Hash computing the SHA224 checksum.

New224 返回hash.Hash, 用于计算SHA224校验和.

## <a id="Sum224">func</a> [Sum224](https://golang.org/src/crypto/sha256/sha256.go?s=5787:5834#L252)
<pre>func Sum224(data []<a href="/pkg/builtin/#byte">byte</a>) (sum224 [<a href="#Size224">Size224</a>]<a href="/pkg/builtin/#byte">byte</a>)</pre>
Sum224 returns the SHA224 checksum of the data.

Sum224 返回参数data的SHA224校验和.

## <a id="Sum256">func</a> [Sum256](https://golang.org/src/crypto/sha256/sha256.go?s=5634:5669#L244)
<pre>func Sum256(data []<a href="/pkg/builtin/#byte">byte</a>) [<a href="#Size">Size</a>]<a href="/pkg/builtin/#byte">byte</a></pre>
Sum256 returns the SHA256 checksum of the data.

Sum256 返回参数data的SHA256校验和.


<a id="example_Sum256">Example</a>
```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	sum := sha256.Sum256([]byte("hello world\n"))
	fmt.Printf("%x", sum)
}
```

output:
```txt
a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447
```






