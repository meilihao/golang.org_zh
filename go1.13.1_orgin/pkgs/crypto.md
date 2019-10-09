

# crypto
`import "crypto"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package crypto collects common cryptographic constants.




## <a id="pkg-index">Index</a>
* [func RegisterHash(h Hash, f func() hash.Hash)](#RegisterHash)
* [type Decrypter](#Decrypter)
* [type DecrypterOpts](#DecrypterOpts)
* [type Hash](#Hash)
  * [func (h Hash) Available() bool](#Hash.Available)
  * [func (h Hash) HashFunc() Hash](#Hash.HashFunc)
  * [func (h Hash) New() hash.Hash](#Hash.New)
  * [func (h Hash) Size() int](#Hash.Size)
* [type PrivateKey](#PrivateKey)
* [type PublicKey](#PublicKey)
* [type Signer](#Signer)
* [type SignerOpts](#SignerOpts)




#### <a id="pkg-files">Package files</a>
[crypto.go](https://golang.org/src/crypto/crypto.go) 






## <a id="RegisterHash">func</a> [RegisterHash](https://golang.org/src/crypto/crypto.go?s=3098:3143#L90)
<pre>func RegisterHash(h <a href="#Hash">Hash</a>, f func() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a>)</pre>
RegisterHash registers a function that returns a new instance of the given
hash function. This is intended to be called from the init function in
packages that implement hash functions.





## <a id="Decrypter">type</a> [Decrypter](https://golang.org/src/crypto/crypto.go?s=5031:5403#L137)
Decrypter is an interface for an opaque private key that can be used for
asymmetric decryption operations. An example would be an RSA key
kept in a hardware module.


<pre>type Decrypter interface {
    <span class="comment">// Public returns the public key corresponding to the opaque,</span>
    <span class="comment">// private key.</span>
    Public() <a href="#PublicKey">PublicKey</a>

    <span class="comment">// Decrypt decrypts msg. The opts argument should be appropriate for</span>
    <span class="comment">// the primitive used. See the documentation in each implementation for</span>
    <span class="comment">// details.</span>
    Decrypt(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, msg []<a href="/pkg/builtin/#byte">byte</a>, opts <a href="#DecrypterOpts">DecrypterOpts</a>) (plaintext []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)
}</pre>











## <a id="DecrypterOpts">type</a> [DecrypterOpts](https://golang.org/src/crypto/crypto.go?s=5405:5435#L148)

<pre>type DecrypterOpts interface{}</pre>











## <a id="Hash">type</a> [Hash](https://golang.org/src/crypto/crypto.go?s=364:378#L6)
Hash identifies a cryptographic hash function that is implemented in another
package.


<pre>type Hash <a href="/pkg/builtin/#uint">uint</a></pre>



<pre>const (
    <span id="MD4">MD4</span>         <a href="#Hash">Hash</a> = 1 + <a href="/pkg/builtin/#iota">iota</a> <span class="comment">// import golang.org/x/crypto/md4</span>
    <span id="MD5">MD5</span>                         <span class="comment">// import crypto/md5</span>
    <span id="SHA1">SHA1</span>                        <span class="comment">// import crypto/sha1</span>
    <span id="SHA224">SHA224</span>                      <span class="comment">// import crypto/sha256</span>
    <span id="SHA256">SHA256</span>                      <span class="comment">// import crypto/sha256</span>
    <span id="SHA384">SHA384</span>                      <span class="comment">// import crypto/sha512</span>
    <span id="SHA512">SHA512</span>                      <span class="comment">// import crypto/sha512</span>
    <span id="MD5SHA1">MD5SHA1</span>                     <span class="comment">// no implementation; MD5+SHA1 used for TLS RSA</span>
    <span id="RIPEMD160">RIPEMD160</span>                   <span class="comment">// import golang.org/x/crypto/ripemd160</span>
    <span id="SHA3_224">SHA3_224</span>                    <span class="comment">// import golang.org/x/crypto/sha3</span>
    <span id="SHA3_256">SHA3_256</span>                    <span class="comment">// import golang.org/x/crypto/sha3</span>
    <span id="SHA3_384">SHA3_384</span>                    <span class="comment">// import golang.org/x/crypto/sha3</span>
    <span id="SHA3_512">SHA3_512</span>                    <span class="comment">// import golang.org/x/crypto/sha3</span>
    <span id="SHA512_224">SHA512_224</span>                  <span class="comment">// import crypto/sha512</span>
    <span id="SHA512_256">SHA512_256</span>                  <span class="comment">// import crypto/sha512</span>
    <span id="BLAKE2s_256">BLAKE2s_256</span>                 <span class="comment">// import golang.org/x/crypto/blake2s</span>
    <span id="BLAKE2b_256">BLAKE2b_256</span>                 <span class="comment">// import golang.org/x/crypto/blake2b</span>
    <span id="BLAKE2b_384">BLAKE2b_384</span>                 <span class="comment">// import golang.org/x/crypto/blake2b</span>
    <span id="BLAKE2b_512">BLAKE2b_512</span>                 <span class="comment">// import golang.org/x/crypto/blake2b</span>

)</pre>









### <a id="Hash.Available">func</a> (Hash) [Available](https://golang.org/src/crypto/crypto.go?s=2827:2857#L83)
<pre>func (h <a href="#Hash">Hash</a>) Available() <a href="/pkg/builtin/#bool">bool</a></pre>
Available reports whether the given hash function is linked into the binary.




### <a id="Hash.HashFunc">func</a> (Hash) [HashFunc](https://golang.org/src/crypto/crypto.go?s=458:487#L9)
<pre>func (h <a href="#Hash">Hash</a>) HashFunc() <a href="#Hash">Hash</a></pre>
HashFunc simply returns the value of h so that Hash implements SignerOpts.




### <a id="Hash.New">func</a> (Hash) [New](https://golang.org/src/crypto/crypto.go?s=2544:2573#L72)
<pre>func (h <a href="#Hash">Hash</a>) New() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a></pre>
New returns a new hash.Hash calculating the given hash function. New panics
if the hash function is not linked into the binary.




### <a id="Hash.Size">func</a> (Hash) [Size](https://golang.org/src/crypto/crypto.go?s=2225:2249#L61)
<pre>func (h <a href="#Hash">Hash</a>) Size() <a href="/pkg/builtin/#int">int</a></pre>
Size returns the length, in bytes, of a digest resulting from the given hash
function. It doesn't require that the hash function in question be linked
into the program.




## <a id="PrivateKey">type</a> [PrivateKey](https://golang.org/src/crypto/crypto.go?s=3411:3438#L101)
PrivateKey represents a private key using an unspecified algorithm.


<pre>type PrivateKey interface{}</pre>











## <a id="PublicKey">type</a> [PublicKey](https://golang.org/src/crypto/crypto.go?s=3312:3338#L98)
PublicKey represents a public key using an unspecified algorithm.


<pre>type PublicKey interface{}</pre>











## <a id="Signer">type</a> [Signer](https://golang.org/src/crypto/crypto.go?s=3587:4582#L105)
Signer is an interface for an opaque private key that can be used for
signing operations. For example, an RSA key kept in a hardware module.


<pre>type Signer interface {
    <span class="comment">// Public returns the public key corresponding to the opaque,</span>
    <span class="comment">// private key.</span>
    Public() <a href="#PublicKey">PublicKey</a>

    <span class="comment">// Sign signs digest with the private key, possibly using entropy from</span>
    <span class="comment">// rand. For an RSA key, the resulting signature should be either a</span>
    <span class="comment">// PKCS#1 v1.5 or PSS signature (as indicated by opts). For an (EC)DSA</span>
    <span class="comment">// key, it should be a DER-serialised, ASN.1 signature structure.</span>
    <span class="comment">//</span>
    <span class="comment">// Hash implements the SignerOpts interface and, in most cases, one can</span>
    <span class="comment">// simply pass in the hash function used as opts. Sign may also attempt</span>
    <span class="comment">// to type assert opts to other types in order to obtain algorithm</span>
    <span class="comment">// specific values. See the documentation in each package for details.</span>
    <span class="comment">//</span>
    <span class="comment">// Note that when a signature of a hash of a larger message is needed,</span>
    <span class="comment">// the caller is responsible for hashing the larger message and passing</span>
    <span class="comment">// the hash (as digest) and the hash function (as opts) to Sign.</span>
    Sign(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, digest []<a href="/pkg/builtin/#byte">byte</a>, opts <a href="#SignerOpts">SignerOpts</a>) (signature []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)
}</pre>











## <a id="SignerOpts">type</a> [SignerOpts](https://golang.org/src/crypto/crypto.go?s=4642:4855#L127)
SignerOpts contains options for signing with a Signer.


<pre>type SignerOpts interface {
    <span class="comment">// HashFunc returns an identifier for the hash function used to produce</span>
    <span class="comment">// the message passed to Signer.Sign, or else zero to indicate that no</span>
    <span class="comment">// hashing was done.</span>
    HashFunc() <a href="#Hash">Hash</a>
}</pre>















