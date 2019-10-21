

# des
`import "crypto/des"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package des implements the Data Encryption Standard (DES) and the
Triple Data Encryption Algorithm (TDEA) as defined
in U.S. Federal Information Processing Standards Publication 46-3.

DES is cryptographically broken and should not be used for secure
applications.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func NewCipher(key []byte) (cipher.Block, error)](#NewCipher)
* [func NewTripleDESCipher(key []byte) (cipher.Block, error)](#NewTripleDESCipher)
* [type KeySizeError](#KeySizeError)
  * [func (k KeySizeError) Error() string](#KeySizeError.Error)


#### <a id="pkg-examples">Examples</a>
* [NewTripleDESCipher](#example_NewTripleDESCipher)


#### <a id="pkg-files">Package files</a>
[block.go](https://golang.org/src/crypto/des/block.go) [cipher.go](https://golang.org/src/crypto/des/cipher.go) [const.go](https://golang.org/src/crypto/des/const.go) 


## <a id="pkg-constants">Constants</a>
The DES block size in bytes.


<pre>const <span id="BlockSize">BlockSize</span> = 8</pre>



## <a id="NewCipher">func</a> [NewCipher](https://golang.org/src/crypto/des/cipher.go?s=586:634#L19)
<pre>func NewCipher(key []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/crypto/cipher/">cipher</a>.<a href="/pkg/crypto/cipher/#Block">Block</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewCipher creates and returns a new cipher.Block.



## <a id="NewTripleDESCipher">func</a> [NewTripleDESCipher](https://golang.org/src/crypto/des/cipher.go?s=1708:1765#L63)
<pre>func NewTripleDESCipher(key []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/crypto/cipher/">cipher</a>.<a href="/pkg/crypto/cipher/#Block">Block</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewTripleDESCipher creates and returns a new cipher.Block.


<a id="example_NewTripleDESCipher">Example</a>
```go
```

output:
```txt
```



## <a id="KeySizeError">type</a> [KeySizeError](https://golang.org/src/crypto/des/cipher.go?s=311:332#L7)

<pre>type KeySizeError <a href="/pkg/builtin/#int">int</a></pre>











### <a id="KeySizeError.Error">func</a> (KeySizeError) [Error](https://golang.org/src/crypto/des/cipher.go?s=334:370#L9)
<pre>func (k <a href="#KeySizeError">KeySizeError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>






