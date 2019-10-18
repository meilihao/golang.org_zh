

# ecdsa
`import "crypto/ecdsa"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package ecdsa implements the Elliptic Curve Digital Signature Algorithm, as
defined in FIPS 186-3.

This implementation  derives the nonce from an AES-CTR CSPRNG keyed by
ChopMD(256, SHA2-512(priv.D || entropy || hash)). The CSPRNG key is IRO by
a result of Coron; the AES-CTR stream is IRO under standard assumptions.


<a id="example_">Example</a>


## <a id="pkg-index">Index</a>
* [func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)](#Sign)
* [func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool](#Verify)
* [type PrivateKey](#PrivateKey)
  * [func GenerateKey(c elliptic.Curve, rand io.Reader) (*PrivateKey, error)](#GenerateKey)
  * [func (priv *PrivateKey) Public() crypto.PublicKey](#PrivateKey.Public)
  * [func (priv *PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)](#PrivateKey.Sign)
* [type PublicKey](#PublicKey)


#### <a id="pkg-examples">Examples</a>
* [Package](#example_)


#### <a id="pkg-files">Package files</a>
[ecdsa.go](https://golang.org/src/crypto/ecdsa/ecdsa.go) 






## <a id="Sign">func</a> [Sign](https://golang.org/src/crypto/ecdsa/ecdsa.go?s=4798:4881#L145)
<pre>func Sign(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, priv *<a href="#PrivateKey">PrivateKey</a>, hash []<a href="/pkg/builtin/#byte">byte</a>) (r, s *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Sign signs a hash (which should be the result of hashing a larger message)
using the private key, priv. If the hash is longer than the bit-length of the
private key's curve order, the hash will be truncated to that length.  It
returns the signature as a pair of integers. The security of the private key
depends on the entropy of rand.



## <a id="Verify">func</a> [Verify](https://golang.org/src/crypto/ecdsa/ecdsa.go?s=6579:6639#L223)
<pre>func Verify(pub *<a href="#PublicKey">PublicKey</a>, hash []<a href="/pkg/builtin/#byte">byte</a>, r, s *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Verify verifies the signature in r, s of hash using the public key, pub. Its
return value records whether the signature is valid.





## <a id="PrivateKey">type</a> [PrivateKey](https://golang.org/src/crypto/ecdsa/ecdsa.go?s=1527:1576#L44)
PrivateKey represents an ECDSA private key.


<pre>type PrivateKey struct {
    <a href="#PublicKey">PublicKey</a>
<span id="PrivateKey.D"></span>    D *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>
}
</pre>









### <a id="GenerateKey">func</a> [GenerateKey](https://golang.org/src/crypto/ecdsa/ecdsa.go?s=2955:3026#L94)
<pre>func GenerateKey(c <a href="/pkg/crypto/elliptic/">elliptic</a>.<a href="/pkg/crypto/elliptic/#Curve">Curve</a>, rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (*<a href="#PrivateKey">PrivateKey</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
GenerateKey generates a public and private key pair.






### <a id="PrivateKey.Public">func</a> (\*PrivateKey) [Public](https://golang.org/src/crypto/ecdsa/ecdsa.go?s=1681:1730#L54)
<pre>func (priv *<a href="#PrivateKey">PrivateKey</a>) Public() <a href="/pkg/crypto/">crypto</a>.<a href="/pkg/crypto/#PublicKey">PublicKey</a></pre>
Public returns the public key corresponding to priv.




### <a id="PrivateKey.Sign">func</a> (\*PrivateKey) [Sign](https://golang.org/src/crypto/ecdsa/ecdsa.go?s=2196:2295#L65)
<pre>func (priv *<a href="#PrivateKey">PrivateKey</a>) Sign(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, digest []<a href="/pkg/builtin/#byte">byte</a>, opts <a href="/pkg/crypto/">crypto</a>.<a href="/pkg/crypto/#SignerOpts">SignerOpts</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Sign signs digest with priv, reading randomness from rand. The opts argument
is not currently used but, in keeping with the crypto.Signer interface,
should be the hash function used to digest the message.

This method implements crypto.Signer, which is an interface to support keys
where the private part is kept in, for example, a hardware module. Common
uses should use the Sign function in this package directly.




## <a id="PublicKey">type</a> [PublicKey](https://golang.org/src/crypto/ecdsa/ecdsa.go?s=1422:1478#L38)
PublicKey represents an ECDSA public key.


<pre>type PublicKey struct {
    <a href="/pkg/crypto/elliptic/">elliptic</a>.<a href="/pkg/crypto/elliptic/#Curve">Curve</a>
<span id="PublicKey.X"></span>    X, Y *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>
}
</pre>















