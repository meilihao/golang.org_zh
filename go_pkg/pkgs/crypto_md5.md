

# md5
`import "crypto/md5"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package md5 implements the MD5 hash algorithm as defined in RFC 1321.

MD5 is cryptographically broken and should not be used for secure
applications.

md5 实现了定义在RFC 1321上的MD5哈希算法.

MD5已被破解, 不应用于安全类的应用.


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

MD5 的 blocksize(字节).

<pre>const <span id="BlockSize">BlockSize</span> = 64</pre>

The size of an MD5 checksum in bytes.

MD5 校验和的大小(字节).

<pre>const <span id="Size">Size</span> = 16</pre>



## <a id="New">func</a> [New](https://golang.org/src/crypto/md5/md5.go?s=2566:2586#L103)
<pre>func New() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a></pre>
New returns a new hash.Hash computing the MD5 checksum. The Hash also
implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
marshal and unmarshal the internal state of the hash.

New 返回hash.Hash, 用于计算MD5校验和. Hash也实现了 encoding.BinaryMarshaler 和 encoding.BinaryUnmarshaler, 用于marshal 和 unmarshal hash的内部状态.

<a id="example_New">Example</a>
```go
package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")
	fmt.Printf("%x", h.Sum(nil))
}
```

output:
```txt
e2c569be17396eca2a2e3c11578123ed
```
<a id="example_New_file">Example (File)</a>
```go
package main

import (
	"crypto/md5"
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

	h := md5.New()
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

## <a id="Sum">func</a> [Sum](https://golang.org/src/crypto/md5/md5.go?s=4513:4545#L180)
<pre>func Sum(data []<a href="/pkg/builtin/#byte">byte</a>) [<a href="#Size">Size</a>]<a href="/pkg/builtin/#byte">byte</a></pre>
Sum returns the MD5 checksum of the data.

Sum 返回参数data的MD5校验和.

<a id="example_Sum">Example</a>
```go
package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := []byte("These pretzels are making me thirsty.")
	fmt.Printf("%x", md5.Sum(data))
}
```

output:
```txt
b0804ec967f48520697662a204f5fe72
```






