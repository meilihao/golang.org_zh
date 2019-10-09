

# rc4
`import "crypto/rc4"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package rc4 implements RC4 encryption, as defined in Bruce Schneier's
Applied Cryptography.

RC4 is cryptographically broken and should not be used for secure
applications.




## <a id="pkg-index">Index</a>
* [type Cipher](#Cipher)
  * [func NewCipher(key []byte) (*Cipher, error)](#NewCipher)
  * [func (c *Cipher) Reset()](#Cipher.Reset)
  * [func (c *Cipher) XORKeyStream(dst, src []byte)](#Cipher.XORKeyStream)
* [type KeySizeError](#KeySizeError)
  * [func (k KeySizeError) Error() string](#KeySizeError.Error)




#### <a id="pkg-files">Package files</a>
[rc4.go](https://golang.org/src/crypto/rc4/rc4.go) 








## <a id="Cipher">type</a> [Cipher](https://golang.org/src/crypto/rc4/rc4.go?s=467:519#L8)
A Cipher is an instance of RC4 using a particular key.


<pre>type Cipher struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewCipher">func</a> [NewCipher](https://golang.org/src/crypto/rc4/rc4.go?s=778:821#L21)
<pre>func NewCipher(key []<a href="/pkg/builtin/#byte">byte</a>) (*<a href="#Cipher">Cipher</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewCipher creates and returns a new Cipher. The key argument should be the
RC4 key, at least 1 byte and at most 256 bytes.






### <a id="Cipher.Reset">func</a> (\*Cipher) [Reset](https://golang.org/src/crypto/rc4/rc4.go?s=1261:1285#L42)
<pre>func (c *<a href="#Cipher">Cipher</a>) Reset()</pre>
Reset zeros the key data and makes the Cipher unusable.

Deprecated: Reset can't guarantee that the key will be entirely removed from
the process's memory.




### <a id="Cipher.XORKeyStream">func</a> (\*Cipher) [XORKeyStream](https://golang.org/src/crypto/rc4/rc4.go?s=1472:1518#L51)
<pre>func (c *<a href="#Cipher">Cipher</a>) XORKeyStream(dst, src []<a href="/pkg/builtin/#byte">byte</a>)</pre>
XORKeyStream sets dst to the result of XORing src with the key stream.
Dst and src must overlap entirely or not at all.




## <a id="KeySizeError">type</a> [KeySizeError](https://golang.org/src/crypto/rc4/rc4.go?s=521:542#L13)

<pre>type KeySizeError <a href="/pkg/builtin/#int">int</a></pre>











### <a id="KeySizeError.Error">func</a> (KeySizeError) [Error](https://golang.org/src/crypto/rc4/rc4.go?s=544:580#L15)
<pre>func (k <a href="#KeySizeError">KeySizeError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>







