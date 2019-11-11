

# rand
`import "crypto/rand"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package rand implements a cryptographically secure
random number generator.

rand 实现了加密安全的随机数生成器.


## <a id="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func Int(rand io.Reader, max *big.Int) (n *big.Int, err error)](#Int)
* [func Prime(rand io.Reader, bits int) (p *big.Int, err error)](#Prime)
* [func Read(b []byte) (n int, err error)](#Read)


#### <a id="pkg-examples">Examples</a>
* [Read](#example_Read)


#### <a id="pkg-files">Package files</a>
[eagain.go](https://golang.org/src/crypto/rand/eagain.go) [rand.go](https://golang.org/src/crypto/rand/rand.go) [rand_batched.go](https://golang.org/src/crypto/rand/rand_batched.go) [rand_linux.go](https://golang.org/src/crypto/rand/rand_linux.go) [rand_unix.go](https://golang.org/src/crypto/rand/rand_unix.go) [util.go](https://golang.org/src/crypto/rand/util.go) 




## <a id="pkg-variables">Variables</a>
Reader is a global, shared instance of a cryptographically
secure random number generator.

On Linux and FreeBSD, Reader uses getrandom(2) if available, /dev/urandom otherwise.
On OpenBSD, Reader uses getentropy(2).
On other Unix-like systems, Reader reads from /dev/urandom.
On Windows systems, Reader uses the CryptGenRandom API.
On Wasm, Reader uses the Web Crypto API.

Reader 是一个全局共享的加密安全的随机数生成器.

Linux 和 FreeBSD 上, 优先使用`getrandom(2)`, 其次是`/dev/urandom`.
OpenBSD上, 使用`getentropy(2)`.
其他类Unix系统, 使用`/dev/urandom`.
Windows上, 使用 CryptGenRandom API.
Wasm上, 使用 Web Crypto API.


<pre>var <span id="Reader">Reader</span> <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a></pre>

## <a id="Int">func</a> [Int](https://golang.org/src/crypto/rand/util.go?s=3070:3132#L96)
<pre>func Int(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, max *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>) (n *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Int returns a uniform random value in [0, max). It panics if max <= 0.

Int 返回一个在[0, max)区间的随机值，如果max<=0则会panic.


## <a id="Prime">func</a> [Prime](https://golang.org/src/crypto/rand/util.go?s=1125:1185#L21)
<pre>func Prime(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, bits <a href="/pkg/builtin/#int">int</a>) (p *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Prime returns a number, p, of the given size, such that p is prime
with high probability.
Prime will return error for any error returned by rand.Read or if bits < 2.

Prime 返回一个给定大小的数值p, 该值很高可能是质数.

如果rand.Read出错，或者bits<2, 它会返回error.

## <a id="Read">func</a> [Read](https://golang.org/src/crypto/rand/rand.go?s=811:849#L13)
<pre>func Read(b []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Read is a helper function that calls Reader.Read using io.ReadFull.
On return, n == len(b) if and only if err == nil.

Read 是一个调用Reader.Read时会使用io.ReadFull的辅助函数. 当且仅当err == nil时，返回值n == len(b).

<a id="example_Read">Example</a>
<p>This example reads 10 cryptographically secure pseudorandom numbers from
rand.Reader and writes them to a byte slice.
</p>

本示例从rand.Reader中读取10个加密安全的伪随机数并将它们写入`[]byte`.

```go
package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
)

func main() {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// The slice should now contain random bytes instead of only zeroes.
	fmt.Println(bytes.Equal(b, make([]byte, c)))

}
```

output:
```txt
false
```





