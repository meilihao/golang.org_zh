

# hmac
`import "crypto/hmac"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package hmac implements the Keyed-Hash Message Authentication Code (HMAC) as
defined in U.S. Federal Information Processing Standards Publication 198.
An HMAC is a cryptographic hash that uses a key to sign a message.
The receiver verifies the hash by recomputing it using the same key.

Receivers should be careful to use Equal to compare MACs in order to avoid
timing side-channels:


	// ValidMAC reports whether messageMAC is a valid HMAC tag for message.
	func ValidMAC(message, messageMAC, key []byte) bool {
		mac := hmac.New(sha256.New, key)
		mac.Write(message)
		expectedMAC := mac.Sum(nil)
		return hmac.Equal(messageMAC, expectedMAC)
	}

hmac 实现了定义在U.S. Federal Information Processing Standards Publication 198中的Keyed-Hash Message Authentication Code (HMAC, 密钥相关的哈希运算消息认证码). HMAC是使用一个key对消息进行签名的加密函数算法. 接收者可通过同样的key重新计算hash来验证它.

接收者应使用谨慎的 Equal 来比较MAC，以避免timing side-channels(时序侧道攻击).


## <a id="pkg-index">Index</a>
* [func Equal(mac1, mac2 []byte) bool](#Equal)
* [func New(h func() hash.Hash, key []byte) hash.Hash](#New)




#### <a id="pkg-files">Package files</a>
[hmac.go](https://golang.org/src/crypto/hmac/hmac.go) 






## <a id="Equal">func</a> [Equal](https://golang.org/src/crypto/hmac/hmac.go?s=2543:2577#L86)
<pre>func Equal(mac1, mac2 []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Equal compares two MACs for equality without leaking timing information.

Equal 比较两个 MAC 是否相等, 且不会泄露时序信息.

## <a id="New">func</a> [New](https://golang.org/src/crypto/hmac/hmac.go?s=1932:1982#L60)
<pre>func New(h func() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a>, key []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a></pre>
New returns a new HMAC hash using the given hash.Hash type and key.
Note that unlike other hash implementations in the standard library,
the returned Hash does not implement encoding.BinaryMarshaler
or encoding.BinaryUnmarshaler.


New 根据给定的 hash.Hash 和 key 返回一个用于 HMAC 的 hash.Hash. 注意与标准库中其他hash实现不同, 返回的 Hash 没有实现 encoding.BinaryMarshaler 和 encoding.BinaryUnmarshaler.







