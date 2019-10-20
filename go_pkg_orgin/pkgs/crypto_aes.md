

# aes
`import "crypto/aes"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package aes implements AES encryption (formerly Rijndael), as defined in
U.S. Federal Information Processing Standards Publication 197.

The AES operations in this package are not implemented using constant-time algorithms.
An exception is when running on systems with enabled hardware support for AES
that makes these operations constant-time. Examples include amd64 systems using AES-NI
extensions and s390x systems using Message-Security-Assist extensions.
On such systems, when the result of NewCipher is passed to cipher.NewGCM,
the GHASH operation used by GCM is also constant-time.

AES 实现了 AES 加密算法(以前叫Rijndael), 由U.S. Federal Information Processing Standards Publication 197(美国联邦信息处理标准出版物197)定义.

本包中的AES操作未使用恒定时间的算法实现. 但在启用了AES硬件支持的系统上运行时, 操作是恒定时间的, 比如使用了AES-NI扩展的amd64和使用Message-Security-Assist扩展的s390x. 在这类系统上, 当将NewCipher的结果传递给cipher.NewGCM时, GCM使用的GHASH操作也是恒定时间的.


## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func NewCipher(key []byte) (cipher.Block, error)](#NewCipher)
* [type KeySizeError](#KeySizeError)
  * [func (k KeySizeError) Error() string](#KeySizeError.Error)




#### <a id="pkg-files">Package files</a>
[aes_gcm.go](https://golang.org/src/crypto/aes/aes_gcm.go) [block.go](https://golang.org/src/crypto/aes/block.go) [cipher.go](https://golang.org/src/crypto/aes/cipher.go) [cipher_asm.go](https://golang.org/src/crypto/aes/cipher_asm.go) [const.go](https://golang.org/src/crypto/aes/const.go) [modes.go](https://golang.org/src/crypto/aes/modes.go) 


## <a id="pkg-constants">Constants</a>
The AES block size in bytes.

AES block 的大小(字节).


<pre>const <span id="BlockSize">BlockSize</span> = 16</pre>



## <a id="NewCipher">func</a> [NewCipher](https://golang.org/src/crypto/aes/cipher.go?s=714:762#L22)
<pre>func NewCipher(key []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/crypto/cipher/">cipher</a>.<a href="/pkg/crypto/cipher/#Block">Block</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewCipher creates and returns a new cipher.Block.
The key argument should be the AES key,
either 16, 24, or 32 bytes to select
AES-128, AES-192, or AES-256.

NewCipher 创建一个cipher.Block接口. 参数key即AES的key. 其长度16,24,32分别对应AES-128, AES-192, AES-256.



## <a id="KeySizeError">type</a> [KeySizeError](https://golang.org/src/crypto/aes/cipher.go?s=417:438#L12)

<pre>type KeySizeError <a href="/pkg/builtin/#int">int</a></pre>











### <a id="KeySizeError.Error">func</a> (KeySizeError) [Error](https://golang.org/src/crypto/aes/cipher.go?s=440:476#L14)
<pre>func (k <a href="#KeySizeError">KeySizeError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>







