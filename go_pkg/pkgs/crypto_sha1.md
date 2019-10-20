

# sha1
`import "crypto/sha1"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package sha1 implements the SHA-1 hash algorithm as defined in RFC 3174.

SHA-1 is cryptographically broken and should not be used for secure
applications.

sha1 实现了定义在RFC 3174上的SHA-1哈希算法.

SHA-1已被破解, 不应用于安全类的应用.

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

SHA-1 的 blocksize(字节).


<pre>const <span id="BlockSize">BlockSize</span> = 64</pre>

The size of a SHA-1 checksum in bytes.

SHA-1 校验和的大小(字节).

<pre>const <span id="Size">Size</span> = 20</pre>



## <a id="New">func</a> [New](https://golang.org/src/crypto/sha1/sha1.go?s=2818:2838#L111)
<pre>func New() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a></pre>
New returns a new hash.Hash computing the SHA1 checksum. The Hash also
implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
marshal and unmarshal the internal state of the hash.

New 返回hash.Hash, 用于计算SHA1校验和. Hash也实现了 encoding.BinaryMarshaler 和 encoding.BinaryUnmarshaler, 用于marshal 和 unmarshal hash的内部状态.

<a id="example_New">Example</a>
```go
package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	h := sha1.New()
	io.WriteString(h, "His money is twice tainted:")
	io.WriteString(h, " 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))
}
```

output:
```txt
59 7f 6a 54 00 10 f9 4c 15 d7 18 06 a9 9a 2c 87 10 e7 47 bd
```
<a id="example_New_file">Example (File)</a>
```go
package main

import (
	"crypto/sha1"
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

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("% x", h.Sum(nil))
}
```

output:
```txt
2009/11/10 23:00:00 open file.txt: No such file or directory
```

## <a id="Sum">func</a> [Sum](https://golang.org/src/crypto/sha1/sha1.go?s=5970:6002#L251)
<pre>func Sum(data []<a href="/pkg/builtin/#byte">byte</a>) [<a href="#Size">Size</a>]<a href="/pkg/builtin/#byte">byte</a></pre>
Sum returns the SHA-1 checksum of the data.

Sum 返回参数data的SHA-1校验和.

<a id="example_Sum">Example</a>
```go
package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	data := []byte("This page intentionally left blank.")
	fmt.Printf("% x", sha1.Sum(data))
}
```

output:
```txt
af 06 49 23 bb f2 30 15 96 aa c4 c2 73 ba 32 17 8e bc 4a 96
```







