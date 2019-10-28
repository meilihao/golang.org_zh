

# hash
`import "hash"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package hash provides interfaces for hash functions.

hash 提供了hash函数所需的接口.


<a id="example__binaryMarshaler">Example (BinaryMarshaler)</a>
```go
package main

import (
	"bytes"
	"crypto/sha256"
	"encoding"
	"fmt"
	"log"
)

func main() {
	const (
		input1 = "The tunneling gopher digs downwards, "
		input2 = "unaware of what he will find."
	)

	first := sha256.New()
	first.Write([]byte(input1))

	marshaler, ok := first.(encoding.BinaryMarshaler)
	if !ok {
		log.Fatal("first does not implement encoding.BinaryMarshaler")
	}
	state, err := marshaler.MarshalBinary()
	if err != nil {
		log.Fatal("unable to marshal hash:", err)
	}

	second := sha256.New()

	unmarshaler, ok := second.(encoding.BinaryUnmarshaler)
	if !ok {
		log.Fatal("second does not implement encoding.BinaryUnmarshaler")
	}
	if err := unmarshaler.UnmarshalBinary(state); err != nil {
		log.Fatal("unable to unmarshal hash:", err)
	}

	first.Write([]byte(input2))
	second.Write([]byte(input2))

	fmt.Printf("%x\n", first.Sum(nil))
	fmt.Println(bytes.Equal(first.Sum(nil), second.Sum(nil)))
}
```

optput:
```txt
57d51a066f3a39942649cd9a76c77e97ceab246756ff3888659e6aa5a07f4a52
true
```

## <a id="pkg-index">Index</a>
* [type Hash](#Hash)
* [type Hash32](#Hash32)
* [type Hash64](#Hash64)


#### <a id="pkg-examples">Examples</a>
* [Package (BinaryMarshaler)](#example__binaryMarshaler)


#### <a id="pkg-files">Package files</a>
[hash.go](https://golang.org/src/hash/hash.go) 








## <a id="Hash">type</a> [Hash](https://golang.org/src/hash/hash.go?s=1238:1887#L16)
Hash is the common interface implemented by all hash functions.

Hash implementations in the standard library (e.g. hash/crc32 and
crypto/sha256) implement the encoding.BinaryMarshaler and
encoding.BinaryUnmarshaler interfaces. Marshaling a hash implementation
allows its internal state to be saved and used for additional processing
later, without having to re-write the data previously written to the hash.
The hash state may contain portions of the input in its original form,
which users are expected to handle for any possible security implications.

Compatibility: Any future changes to hash or crypto packages will endeavor
to maintain compatibility with state encoded using previous versions.
That is, any released versions of the packages should be able to
decode data written with any previously released version,
subject to issues such as security fixes.
See the Go compatibility document for background: <a href="https://golang.org/doc/go1compat">https://golang.org/doc/go1compat</a>

 Hash是所有hash函数都需要实现的公共接口.

 标准库中的Hash实现(比如hash/crc32 和 crypto/sha256)也实现了encoding.BinaryMarshaler 和 encoding.BinaryUnmarshaler 接口. 它们可以将内部的state保存起来, 以便在以后处理, 而不必重新写入先前写入hash的数据. hash state 可能包含原始的输入的部分???, 应由用户处理以应对任何可能的安全隐患.

兼容性: 将来对hash或crypto package的任何变更都将努力保持与以前版本的state编码兼容. 也就是说, 任何正式版本都能解码任何先前发行版本编码的数据, 便于安全修复. 有关背景信息, 请参阅Go兼容性文档: <a href="https://golang.org/doc/go1compat">https://golang.org/doc/go1compat</a>

```go
type Hash interface {
    // Write (via the embedded io.Writer interface) adds more data to the running hash.
    // It never returns an error.
    //
    // Write (通过嵌入io.Writer接口) 将更多的数据传入正在运行的hash函数.
    // 该方法不会返回错误.
    io.Writer

    // Sum appends the current hash to b and returns the resulting slice.
    // It does not change the underlying hash state.
    //
    // Sum 将当前(计算所得)的hash值追加到b后面,并返回slice类型的结果.
    // 该方法不会改变底层hash函数的状态.
    Sum(b []byte) []byte

    // Reset resets the Hash to its initial state.
    // Reset 将hash函数重置到初始状态.
    Reset()

    // Size returns the number of bytes Sum will return.
    // Size 返回Sum方法结果的长度.
    Size() int

    // BlockSize returns the hash's underlying block size.
    // The Write method must be able to accept any amount
    // of data, but it may operate more efficiently if all writes
    // are a multiple of the block size.
    //
    // BlockSize 返回hash函数的底层块大小.
    // Write方法必须能接受任意数量的数据,但如果写入的数据量都是块大小的倍数的话,那么该方法的处理会更高效.
    BlockSize() int
}
```

## <a id="Hash32">type</a> [Hash32](https://golang.org/src/hash/hash.go?s=1965:2012#L39)
Hash32 is the common interface implemented by all 32-bit hash functions.

Hash32 是所有32位hash函数都需要实现的公共接口.

<pre>type Hash32 interface {
    <a href="#Hash">Hash</a>
    Sum32() <a href="/pkg/builtin/#uint32">uint32</a>
}</pre>











## <a id="Hash64">type</a> [Hash64](https://golang.org/src/hash/hash.go?s=2090:2137#L45)
Hash64 is the common interface implemented by all 64-bit hash functions.

Hash64 是所有64位hash函数都需要实现的公共接口.

<pre>type Hash64 interface {
    <a href="#Hash">Hash</a>
    Sum64() <a href="/pkg/builtin/#uint64">uint64</a>
}</pre>















