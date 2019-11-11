

# rand
`import "crypto/rand"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package rand implements a cryptographically secure
random number generator.




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


<pre>var <span id="Reader">Reader</span> <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a></pre>

## <a id="Int">func</a> [Int](https://golang.org/src/crypto/rand/util.go?s=3070:3132#L96)
<pre>func Int(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, max *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>) (n *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Int returns a uniform random value in [0, max). It panics if max <= 0.



## <a id="Prime">func</a> [Prime](https://golang.org/src/crypto/rand/util.go?s=1125:1185#L21)
<pre>func Prime(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, bits <a href="/pkg/builtin/#int">int</a>) (p *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Prime returns a number, p, of the given size, such that p is prime
with high probability.
Prime will return error for any error returned by rand.Read or if bits < 2.



## <a id="Read">func</a> [Read](https://golang.org/src/crypto/rand/rand.go?s=811:849#L13)
<pre>func Read(b []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Read is a helper function that calls Reader.Read using io.ReadFull.
On return, n == len(b) if and only if err == nil.



<a id="example_Read">Example</a>
<p>This example reads 10 cryptographically secure pseudorandom numbers from
rand.Reader and writes them to a byte slice.
</p>

```go
```

output:
```txt
```






