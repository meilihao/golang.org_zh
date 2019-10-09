

# ed25519
`import "crypto/ed25519"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package ed25519 implements the Ed25519 signature algorithm. See
<a href="https://ed25519.cr.yp.to/">https://ed25519.cr.yp.to/</a>.

These functions are also compatible with the “Ed25519” function defined in
RFC 8032. However, unlike RFC 8032's formulation, this package's private key
representation includes a public key suffix to make multiple signing
operations with the same key more efficient. This package refers to the RFC
8032 private key as the “seed”.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func GenerateKey(rand io.Reader) (PublicKey, PrivateKey, error)](#GenerateKey)
* [func Sign(privateKey PrivateKey, message []byte) []byte](#Sign)
* [func Verify(publicKey PublicKey, message, sig []byte) bool](#Verify)
* [type PrivateKey](#PrivateKey)
  * [func NewKeyFromSeed(seed []byte) PrivateKey](#NewKeyFromSeed)
  * [func (priv PrivateKey) Public() crypto.PublicKey](#PrivateKey.Public)
  * [func (priv PrivateKey) Seed() []byte](#PrivateKey.Seed)
  * [func (priv PrivateKey) Sign(rand io.Reader, message []byte, opts crypto.SignerOpts) (signature []byte, err error)](#PrivateKey.Sign)
* [type PublicKey](#PublicKey)




#### <a id="pkg-files">Package files</a>
[ed25519.go](https://golang.org/src/crypto/ed25519/ed25519.go) 


## <a id="pkg-constants">Constants</a>

<pre>const (
    <span class="comment">// PublicKeySize is the size, in bytes, of public keys as used in this package.</span>
    <span id="PublicKeySize">PublicKeySize</span> = 32
    <span class="comment">// PrivateKeySize is the size, in bytes, of private keys as used in this package.</span>
    <span id="PrivateKeySize">PrivateKeySize</span> = 64
    <span class="comment">// SignatureSize is the size, in bytes, of signatures generated and verified by this package.</span>
    <span id="SignatureSize">SignatureSize</span> = 64
    <span class="comment">// SeedSize is the size, in bytes, of private key seeds. These are the private key representations used by RFC 8032.</span>
    <span id="SeedSize">SeedSize</span> = 32
)</pre>



## <a id="GenerateKey">func</a> [GenerateKey](https://golang.org/src/crypto/ed25519/ed25519.go?s=2707:2770#L67)
<pre>func GenerateKey(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (<a href="#PublicKey">PublicKey</a>, <a href="#PrivateKey">PrivateKey</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
GenerateKey generates a public/private key pair using entropy from rand.
If rand is nil, crypto/rand.Reader will be used.



## <a id="Sign">func</a> [Sign](https://golang.org/src/crypto/ed25519/ed25519.go?s=3996:4051#L114)
<pre>func Sign(privateKey <a href="#PrivateKey">PrivateKey</a>, message []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
Sign signs the message with privateKey and returns a signature. It will
panic if len(privateKey) is not PrivateKeySize.



## <a id="Verify">func</a> [Verify](https://golang.org/src/crypto/ed25519/ed25519.go?s=5320:5378#L163)
<pre>func Verify(publicKey <a href="#PublicKey">PublicKey</a>, message, sig []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Verify reports whether sig is a valid signature of message by publicKey. It
will panic if len(publicKey) is not PublicKeySize.





## <a id="PrivateKey">type</a> [PrivateKey](https://golang.org/src/crypto/ed25519/ed25519.go?s=1488:1510#L34)
PrivateKey is the type of Ed25519 private keys. It implements crypto.Signer.


<pre>type PrivateKey []<a href="/pkg/builtin/#byte">byte</a></pre>









### <a id="NewKeyFromSeed">func</a> [NewKeyFromSeed](https://golang.org/src/crypto/ed25519/ed25519.go?s=3316:3359#L88)
<pre>func NewKeyFromSeed(seed []<a href="/pkg/builtin/#byte">byte</a>) <a href="#PrivateKey">PrivateKey</a></pre>
NewKeyFromSeed calculates a private key from a seed. It will panic if
len(seed) is not SeedSize. This function is provided for interoperability
with RFC 8032. RFC 8032's private keys correspond to seeds in this
package.






### <a id="PrivateKey.Public">func</a> (PrivateKey) [Public](https://golang.org/src/crypto/ed25519/ed25519.go?s=1567:1615#L37)
<pre>func (priv <a href="#PrivateKey">PrivateKey</a>) Public() <a href="/pkg/crypto/">crypto</a>.<a href="/pkg/crypto/#PublicKey">PublicKey</a></pre>
Public returns the PublicKey corresponding to priv.




### <a id="PrivateKey.Seed">func</a> (PrivateKey) [Seed](https://golang.org/src/crypto/ed25519/ed25519.go?s=1898:1934#L46)
<pre>func (priv <a href="#PrivateKey">PrivateKey</a>) Seed() []<a href="/pkg/builtin/#byte">byte</a></pre>
Seed returns the private key seed corresponding to priv. It is provided for
interoperability with RFC 8032. RFC 8032's private keys correspond to seeds
in this package.




### <a id="PrivateKey.Sign">func</a> (PrivateKey) [Sign](https://golang.org/src/crypto/ed25519/ed25519.go?s=2319:2432#L57)
<pre>func (priv <a href="#PrivateKey">PrivateKey</a>) Sign(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, message []<a href="/pkg/builtin/#byte">byte</a>, opts <a href="/pkg/crypto/">crypto</a>.<a href="/pkg/crypto/#SignerOpts">SignerOpts</a>) (signature []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Sign signs the given message with priv.
Ed25519 performs two passes over messages to be signed and therefore cannot
handle pre-hashed messages. Thus opts.HashFunc() must return zero to
indicate the message hasn't been hashed. This can be achieved by passing
crypto.Hash(0) as the value for opts.




## <a id="PublicKey">type</a> [PublicKey](https://golang.org/src/crypto/ed25519/ed25519.go?s=1385:1406#L31)
PublicKey is the type of Ed25519 public keys.


<pre>type PublicKey []<a href="/pkg/builtin/#byte">byte</a></pre>















