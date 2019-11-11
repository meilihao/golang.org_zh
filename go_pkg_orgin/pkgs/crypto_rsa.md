

# rsa
`import "crypto/rsa"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package rsa implements RSA encryption as specified in PKCS#1.

RSA is a single, fundamental operation that is used in this package to
implement either public-key encryption or public-key signatures.

The original specification for encryption and signatures with RSA is PKCS#1
and the terms "RSA encryption" and "RSA signatures" by default refer to
PKCS#1 version 1.5. However, that specification has flaws and new designs
should use version two, usually called by just OAEP and PSS, where
possible.

Two sets of interfaces are included in this package. When a more abstract
interface isn't necessary, there are functions for encrypting/decrypting
with v1.5/OAEP and signing/verifying with v1.5/PSS. If one needs to abstract
over the public-key primitive, the PrivateKey struct implements the
Decrypter and Signer interfaces from the crypto package.

The RSA operations in this package are not implemented using constant-time algorithms.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func DecryptOAEP(hash hash.Hash, random io.Reader, priv *PrivateKey, ciphertext []byte, label []byte) ([]byte, error)](#DecryptOAEP)
* [func DecryptPKCS1v15(rand io.Reader, priv *PrivateKey, ciphertext []byte) ([]byte, error)](#DecryptPKCS1v15)
* [func DecryptPKCS1v15SessionKey(rand io.Reader, priv *PrivateKey, ciphertext []byte, key []byte) error](#DecryptPKCS1v15SessionKey)
* [func EncryptOAEP(hash hash.Hash, random io.Reader, pub *PublicKey, msg []byte, label []byte) ([]byte, error)](#EncryptOAEP)
* [func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) ([]byte, error)](#EncryptPKCS1v15)
* [func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error)](#SignPKCS1v15)
* [func SignPSS(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte, opts *PSSOptions) ([]byte, error)](#SignPSS)
* [func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error](#VerifyPKCS1v15)
* [func VerifyPSS(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte, opts *PSSOptions) error](#VerifyPSS)
* [type CRTValue](#CRTValue)
* [type OAEPOptions](#OAEPOptions)
* [type PKCS1v15DecryptOptions](#PKCS1v15DecryptOptions)
* [type PSSOptions](#PSSOptions)
  * [func (pssOpts *PSSOptions) HashFunc() crypto.Hash](#PSSOptions.HashFunc)
* [type PrecomputedValues](#PrecomputedValues)
* [type PrivateKey](#PrivateKey)
  * [func GenerateKey(random io.Reader, bits int) (*PrivateKey, error)](#GenerateKey)
  * [func GenerateMultiPrimeKey(random io.Reader, nprimes int, bits int) (*PrivateKey, error)](#GenerateMultiPrimeKey)
  * [func (priv *PrivateKey) Decrypt(rand io.Reader, ciphertext []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error)](#PrivateKey.Decrypt)
  * [func (priv *PrivateKey) Precompute()](#PrivateKey.Precompute)
  * [func (priv *PrivateKey) Public() crypto.PublicKey](#PrivateKey.Public)
  * [func (priv *PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)](#PrivateKey.Sign)
  * [func (priv *PrivateKey) Validate() error](#PrivateKey.Validate)
* [type PublicKey](#PublicKey)
  * [func (pub *PublicKey) Size() int](#PublicKey.Size)


#### <a id="pkg-examples">Examples</a>
* [DecryptOAEP](#example_DecryptOAEP)
* [DecryptPKCS1v15SessionKey](#example_DecryptPKCS1v15SessionKey)
* [EncryptOAEP](#example_EncryptOAEP)
* [SignPKCS1v15](#example_SignPKCS1v15)
* [VerifyPKCS1v15](#example_VerifyPKCS1v15)


#### <a id="pkg-files">Package files</a>
[pkcs1v15.go](https://golang.org/src/crypto/rsa/pkcs1v15.go) [pss.go](https://golang.org/src/crypto/rsa/pss.go) [rsa.go](https://golang.org/src/crypto/rsa/rsa.go) 


## <a id="pkg-constants">Constants</a>

<pre>const (
    <span class="comment">// PSSSaltLengthAuto causes the salt in a PSS signature to be as large</span>
    <span class="comment">// as possible when signing, and to be auto-detected when verifying.</span>
    <span id="PSSSaltLengthAuto">PSSSaltLengthAuto</span> = 0
    <span class="comment">// PSSSaltLengthEqualsHash causes the salt length to equal the length</span>
    <span class="comment">// of the hash used in the signature.</span>
    <span id="PSSSaltLengthEqualsHash">PSSSaltLengthEqualsHash</span> = -1
)</pre>

## <a id="pkg-variables">Variables</a>
ErrDecryption represents a failure to decrypt a message.
It is deliberately vague to avoid adaptive attacks.


<pre>var <span id="ErrDecryption">ErrDecryption</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;crypto/rsa: decryption error&#34;)</pre>ErrMessageTooLong is returned when attempting to encrypt a message which is
too large for the size of the public key.


<pre>var <span id="ErrMessageTooLong">ErrMessageTooLong</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;crypto/rsa: message too long for RSA public key size&#34;)</pre>ErrVerification represents a failure to verify a signature.
It is deliberately vague to avoid adaptive attacks.


<pre>var <span id="ErrVerification">ErrVerification</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;crypto/rsa: verification error&#34;)</pre>

## <a id="DecryptOAEP">func</a> [DecryptOAEP](https://golang.org/src/crypto/rsa/rsa.go?s=17122:17239#L559)
<pre>func DecryptOAEP(hash <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a>, random <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, priv *<a href="#PrivateKey">PrivateKey</a>, ciphertext []<a href="/pkg/builtin/#byte">byte</a>, label []<a href="/pkg/builtin/#byte">byte</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DecryptOAEP decrypts ciphertext using RSA-OAEP.

OAEP is parameterised by a hash function that is used as a random oracle.
Encryption and decryption of a given message must use the same hash function
and sha256.New() is a reasonable choice.

The random parameter, if not nil, is used to blind the private-key operation
and avoid timing side-channel attacks. Blinding is purely internal to this
function – the random data need not match that used when encrypting.

The label parameter must match the value given when encrypting. See
EncryptOAEP for details.



<a id="example_DecryptOAEP">Example</a>


```go
```

output:
```txt
```

## <a id="DecryptPKCS1v15">func</a> [DecryptPKCS1v15](https://golang.org/src/crypto/rsa/pkcs1v15.go?s=2396:2485#L66)
<pre>func DecryptPKCS1v15(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, priv *<a href="#PrivateKey">PrivateKey</a>, ciphertext []<a href="/pkg/builtin/#byte">byte</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DecryptPKCS1v15 decrypts a plaintext using RSA and the padding scheme from PKCS#1 v1.5.
If rand != nil, it uses RSA blinding to avoid timing side-channel attacks.

Note that whether this function returns an error or not discloses secret
information. If an attacker can cause this function to run repeatedly and
learn whether each instance returned an error then they can decrypt and
forge signatures as if they had the private key. See
DecryptPKCS1v15SessionKey for a way of solving this problem.



## <a id="DecryptPKCS1v15SessionKey">func</a> [DecryptPKCS1v15SessionKey](https://golang.org/src/crypto/rsa/pkcs1v15.go?s=4036:4137#L99)
<pre>func DecryptPKCS1v15SessionKey(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, priv *<a href="#PrivateKey">PrivateKey</a>, ciphertext []<a href="/pkg/builtin/#byte">byte</a>, key []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
DecryptPKCS1v15SessionKey decrypts a session key using RSA and the padding scheme from PKCS#1 v1.5.
If rand != nil, it uses RSA blinding to avoid timing side-channel attacks.
It returns an error if the ciphertext is the wrong length or if the
ciphertext is greater than the public modulus. Otherwise, no error is
returned. If the padding is valid, the resulting plaintext message is copied
into key. Otherwise, key is unchanged. These alternatives occur in constant
time. It is intended that the user of this function generate a random
session key beforehand and continue the protocol with the resulting value.
This will remove any possibility that an attacker can learn any information
about the plaintext.
See ``Chosen Ciphertext Attacks Against Protocols Based on the RSA
Encryption Standard PKCS #1'', Daniel Bleichenbacher, Advances in Cryptology
(Crypto '98).

Note that if the session key is too small then it may be possible for an
attacker to brute-force it. If they can do that then they can learn whether
a random value was used (because it'll be different for the same ciphertext)
and thus whether the padding was correct. This defeats the point of this
function. Using at least a 16-byte key will protect against this attack.



<a id="example_DecryptPKCS1v15SessionKey">Example</a>
<p>RSA is able to encrypt only a very limited amount of data. In order
to encrypt reasonable amounts of data a hybrid scheme is commonly
used: RSA is used to encrypt a key for a symmetric primitive like
AES-GCM.

Before encrypting, data is “padded” by embedding it in a known
structure. This is done for a number of reasons, but the most
obvious is to ensure that the value is large enough that the
exponentiation is larger than the modulus. (Otherwise it could be
decrypted with a square-root.)

In these designs, when using PKCS#1 v1.5, it&#39;s vitally important to
avoid disclosing whether the received RSA message was well-formed
(that is, whether the result of decrypting is a correctly padded
message) because this leaks secret information.
DecryptPKCS1v15SessionKey is designed for this situation and copies
the decrypted, symmetric key (if well-formed) in constant-time over
a buffer that contains a random key. Thus, if the RSA result isn&#39;t
well-formed, the implementation uses a random key in constant time.
</p>

```go
```

output:
```txt
```

## <a id="EncryptOAEP">func</a> [EncryptOAEP](https://golang.org/src/crypto/rsa/rsa.go?s=11973:12081#L366)
<pre>func EncryptOAEP(hash <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a>, random <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, pub *<a href="#PublicKey">PublicKey</a>, msg []<a href="/pkg/builtin/#byte">byte</a>, label []<a href="/pkg/builtin/#byte">byte</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
EncryptOAEP encrypts the given message with RSA-OAEP.

OAEP is parameterised by a hash function that is used as a random oracle.
Encryption and decryption of a given message must use the same hash function
and sha256.New() is a reasonable choice.

The random parameter is used as a source of entropy to ensure that
encrypting the same message twice doesn't result in the same ciphertext.

The label parameter may contain arbitrary data that will not be encrypted,
but which gives important context to the message. For example, if a given
public key is used to decrypt two types of messages then distinct label
values could be used to ensure that a ciphertext for one purpose cannot be
used for another by an attacker. If not required it can be empty.

The message must be no longer than the length of the public modulus minus
twice the hash length, minus a further 2.



<a id="example_EncryptOAEP">Example</a>


```go
```

output:
```txt
```

## <a id="EncryptPKCS1v15">func</a> [EncryptPKCS1v15](https://golang.org/src/crypto/rsa/pkcs1v15.go?s=1254:1334#L29)
<pre>func EncryptPKCS1v15(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, pub *<a href="#PublicKey">PublicKey</a>, msg []<a href="/pkg/builtin/#byte">byte</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
EncryptPKCS1v15 encrypts the given message with RSA and the padding
scheme from PKCS#1 v1.5.  The message must be no longer than the
length of the public modulus minus 11 bytes.

The rand parameter is used as a source of entropy to ensure that
encrypting the same message twice doesn't result in the same
ciphertext.

WARNING: use of this function to encrypt plaintexts other than
session keys is dangerous. Use RSA OAEP in new protocols.



## <a id="SignPKCS1v15">func</a> [SignPKCS1v15](https://golang.org/src/crypto/rsa/pkcs1v15.go?s=8780:8880#L222)
<pre>func SignPKCS1v15(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, priv *<a href="#PrivateKey">PrivateKey</a>, hash <a href="/pkg/crypto/">crypto</a>.<a href="/pkg/crypto/#Hash">Hash</a>, hashed []<a href="/pkg/builtin/#byte">byte</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
SignPKCS1v15 calculates the signature of hashed using
RSASSA-PKCS1-V1_5-SIGN from RSA PKCS#1 v1.5.  Note that hashed must
be the result of hashing the input message using the given hash
function. If hash is zero, hashed is signed directly. This isn't
advisable except for interoperability.

If rand is not nil then RSA blinding will be used to avoid timing
side-channel attacks.

This function is deterministic. Thus, if the set of possible
messages is small, an attacker may be able to build a map from
messages to signatures and identify the signed messages. As ever,
signatures provide authenticity, not confidentiality.



<a id="example_SignPKCS1v15">Example</a>


```go
```

output:
```txt
```

## <a id="SignPSS">func</a> [SignPSS](https://golang.org/src/crypto/rsa/pss.go?s=6982:7095#L239)
<pre>func SignPSS(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, priv *<a href="#PrivateKey">PrivateKey</a>, hash <a href="/pkg/crypto/">crypto</a>.<a href="/pkg/crypto/#Hash">Hash</a>, hashed []<a href="/pkg/builtin/#byte">byte</a>, opts *<a href="#PSSOptions">PSSOptions</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
SignPSS calculates the signature of hashed using RSASSA-PSS [1].
Note that hashed must be the result of hashing the input message using the
given hash function. The opts argument may be nil, in which case sensible
defaults are used.



## <a id="VerifyPKCS1v15">func</a> [VerifyPKCS1v15](https://golang.org/src/crypto/rsa/pkcs1v15.go?s=9772:9858#L258)
<pre>func VerifyPKCS1v15(pub *<a href="#PublicKey">PublicKey</a>, hash <a href="/pkg/crypto/">crypto</a>.<a href="/pkg/crypto/#Hash">Hash</a>, hashed []<a href="/pkg/builtin/#byte">byte</a>, sig []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
VerifyPKCS1v15 verifies an RSA PKCS#1 v1.5 signature.
hashed is the result of hashing the input message using the given hash
function and sig is the signature. A valid signature is indicated by
returning a nil error. If hash is zero then hashed is used directly. This
isn't advisable except for interoperability.



<a id="example_VerifyPKCS1v15">Example</a>


```go
```

output:
```txt
```

## <a id="VerifyPSS">func</a> [VerifyPSS](https://golang.org/src/crypto/rsa/pss.go?s=7805:7904#L264)
<pre>func VerifyPSS(pub *<a href="#PublicKey">PublicKey</a>, hash <a href="/pkg/crypto/">crypto</a>.<a href="/pkg/crypto/#Hash">Hash</a>, hashed []<a href="/pkg/builtin/#byte">byte</a>, sig []<a href="/pkg/builtin/#byte">byte</a>, opts *<a href="#PSSOptions">PSSOptions</a>) <a href="/pkg/builtin/#error">error</a></pre>
VerifyPSS verifies a PSS signature.
hashed is the result of hashing the input message using the given hash
function and sig is the signature. A valid signature is indicated by
returning a nil error. The opts argument may be nil, in which case sensible
defaults are used.





## <a id="CRTValue">type</a> [CRTValue](https://golang.org/src/crypto/rsa/rsa.go?s=5313:5484#L151)
CRTValue contains the precomputed Chinese remainder theorem values.


<pre>type CRTValue struct {
<span id="CRTValue.Exp"></span>    Exp   *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a> <span class="comment">// D mod (prime-1).</span>
<span id="CRTValue.Coeff"></span>    Coeff *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a> <span class="comment">// R·Coeff ≡ 1 mod Prime.</span>
<span id="CRTValue.R"></span>    R     *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a> <span class="comment">// product of primes prior to this (inc p and q).</span>
}
</pre>











## <a id="OAEPOptions">type</a> [OAEPOptions](https://golang.org/src/crypto/rsa/rsa.go?s=1791:2020#L45)
OAEPOptions is an interface for passing options to OAEP decryption using the
crypto.Decrypter interface.


<pre>type OAEPOptions struct {
<span id="OAEPOptions.Hash"></span>    <span class="comment">// Hash is the hash function that will be used when generating the mask.</span>
    Hash <a href="/pkg/crypto/">crypto</a>.<a href="/pkg/crypto/#Hash">Hash</a>
<span id="OAEPOptions.Label"></span>    <span class="comment">// Label is an arbitrary byte string that must be equal to the value</span>
    <span class="comment">// used when encrypting.</span>
    Label []<a href="/pkg/builtin/#byte">byte</a>
}
</pre>











## <a id="PKCS1v15DecryptOptions">type</a> [PKCS1v15DecryptOptions](https://golang.org/src/crypto/rsa/pkcs1v15.go?s=462:785#L11)
PKCS1v15DecrypterOpts is for passing options to PKCS#1 v1.5 decryption using
the crypto.Decrypter interface.


<pre>type PKCS1v15DecryptOptions struct {
<span id="PKCS1v15DecryptOptions.SessionKeyLen"></span>    <span class="comment">// SessionKeyLen is the length of the session key that is being</span>
    <span class="comment">// decrypted. If not zero, then a padding error during decryption will</span>
    <span class="comment">// cause a random plaintext of this length to be returned rather than</span>
    <span class="comment">// an error. These alternatives happen in constant time.</span>
    SessionKeyLen <a href="/pkg/builtin/#int">int</a>
}
</pre>











## <a id="PSSOptions">type</a> [PSSOptions](https://golang.org/src/crypto/rsa/pss.go?s=6062:6456#L210)
PSSOptions contains options for creating and verifying PSS signatures.


<pre>type PSSOptions struct {
<span id="PSSOptions.SaltLength"></span>    <span class="comment">// SaltLength controls the length of the salt used in the PSS</span>
    <span class="comment">// signature. It can either be a number of bytes, or one of the special</span>
    <span class="comment">// PSSSaltLength constants.</span>
    SaltLength <a href="/pkg/builtin/#int">int</a>

<span id="PSSOptions.Hash"></span>    <span class="comment">// Hash, if not zero, overrides the hash function passed to SignPSS.</span>
    <span class="comment">// This is the only way to specify the hash function when using the</span>
    <span class="comment">// crypto.Signer interface.</span>
    Hash <a href="/pkg/crypto/">crypto</a>.<a href="/pkg/crypto/#Hash">Hash</a>
}
</pre>











### <a id="PSSOptions.HashFunc">func</a> (\*PSSOptions) [HashFunc](https://golang.org/src/crypto/rsa/pss.go?s=6543:6592#L224)
<pre>func (pssOpts *<a href="#PSSOptions">PSSOptions</a>) HashFunc() <a href="/pkg/crypto/">crypto</a>.<a href="/pkg/crypto/#Hash">Hash</a></pre>
HashFunc returns pssOpts.Hash so that PSSOptions implements
crypto.SignerOpts.




## <a id="PrecomputedValues">type</a> [PrecomputedValues](https://golang.org/src/crypto/rsa/rsa.go?s=4875:5240#L139)

<pre>type PrecomputedValues struct {
<span id="PrecomputedValues.Dp"></span>    Dp, Dq *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a> <span class="comment">// D mod (P-1) (or mod Q-1)</span>
<span id="PrecomputedValues.Qinv"></span>    Qinv   *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a> <span class="comment">// Q^-1 mod P</span>

<span id="PrecomputedValues.CRTValues"></span>    <span class="comment">// CRTValues is used for the 3rd and subsequent primes. Due to a</span>
    <span class="comment">// historical accident, the CRT for the first two primes is handled</span>
    <span class="comment">// differently in PKCS#1 and interoperability is sufficiently</span>
    <span class="comment">// important that we mirror this.</span>
    CRTValues []<a href="#CRTValue">CRTValue</a>
}
</pre>











## <a id="PrivateKey">type</a> [PrivateKey](https://golang.org/src/crypto/rsa/rsa.go?s=2773:3071#L78)
A PrivateKey represents an RSA key


<pre>type PrivateKey struct {
    <a href="#PublicKey">PublicKey</a>            <span class="comment">// public part.</span>
<span id="PrivateKey.D"></span>    D         *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>   <span class="comment">// private exponent</span>
<span id="PrivateKey.Primes"></span>    Primes    []*<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a> <span class="comment">// prime factors of N, has &gt;= 2 elements.</span>

<span id="PrivateKey.Precomputed"></span>    <span class="comment">// Precomputed contains precomputed values that speed up private</span>
    <span class="comment">// operations, if available.</span>
    Precomputed <a href="#PrecomputedValues">PrecomputedValues</a>
}
</pre>









### <a id="GenerateKey">func</a> [GenerateKey](https://golang.org/src/crypto/rsa/rsa.go?s=6882:6947#L197)
<pre>func GenerateKey(random <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, bits <a href="/pkg/builtin/#int">int</a>) (*<a href="#PrivateKey">PrivateKey</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
GenerateKey generates an RSA keypair of the given bit size using the
random source random (for example, crypto/rand.Reader).




### <a id="GenerateMultiPrimeKey">func</a> [GenerateMultiPrimeKey](https://golang.org/src/crypto/rsa/rsa.go?s=7589:7677#L212)
<pre>func GenerateMultiPrimeKey(random <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, nprimes <a href="/pkg/builtin/#int">int</a>, bits <a href="/pkg/builtin/#int">int</a>) (*<a href="#PrivateKey">PrivateKey</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
GenerateMultiPrimeKey generates a multi-prime RSA keypair of the given bit
size and the given random source, as suggested in [1]. Although the public
keys are compatible (actually, indistinguishable) from the 2-prime case,
the private keys are not. Thus it may not be possible to export multi-prime
private keys in certain formats or to subsequently import them into other
code.

Table 1 in [2] suggests maximum numbers of primes for a given size.

[1] US patent 4405829 (1972, expired)
[2] <a href="http://www.cacr.math.uwaterloo.ca/techreports/2006/cacr2006-16.pdf">http://www.cacr.math.uwaterloo.ca/techreports/2006/cacr2006-16.pdf</a>






### <a id="PrivateKey.Decrypt">func</a> (\*PrivateKey) [Decrypt](https://golang.org/src/crypto/rsa/rsa.go?s=4081:4204#L111)
<pre>func (priv *<a href="#PrivateKey">PrivateKey</a>) Decrypt(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, ciphertext []<a href="/pkg/builtin/#byte">byte</a>, opts <a href="/pkg/crypto/">crypto</a>.<a href="/pkg/crypto/#DecrypterOpts">DecrypterOpts</a>) (plaintext []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Decrypt decrypts ciphertext with priv. If opts is nil or of type
*PKCS1v15DecryptOptions then PKCS#1 v1.5 decryption is performed. Otherwise
opts must have type *OAEPOptions and OAEP decryption is done.




### <a id="PrivateKey.Precompute">func</a> (\*PrivateKey) [Precompute](https://golang.org/src/crypto/rsa/rsa.go?s=13327:13363#L421)
<pre>func (priv *<a href="#PrivateKey">PrivateKey</a>) Precompute()</pre>
Precompute performs some calculations that speed up private key operations
in the future.




### <a id="PrivateKey.Public">func</a> (\*PrivateKey) [Public](https://golang.org/src/crypto/rsa/rsa.go?s=3129:3178#L89)
<pre>func (priv *<a href="#PrivateKey">PrivateKey</a>) Public() <a href="/pkg/crypto/">crypto</a>.<a href="/pkg/crypto/#PublicKey">PublicKey</a></pre>
Public returns the public key corresponding to priv.




### <a id="PrivateKey.Sign">func</a> (\*PrivateKey) [Sign](https://golang.org/src/crypto/rsa/rsa.go?s=3598:3697#L100)
<pre>func (priv *<a href="#PrivateKey">PrivateKey</a>) Sign(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, digest []<a href="/pkg/builtin/#byte">byte</a>, opts <a href="/pkg/crypto/">crypto</a>.<a href="/pkg/crypto/#SignerOpts">SignerOpts</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Sign signs digest with priv, reading randomness from rand. If opts is a
*PSSOptions then the PSS algorithm will be used, otherwise PKCS#1 v1.5 will
be used.

This method implements crypto.Signer, which is an interface to support keys
where the private part is kept in, for example, a hardware module. Common
uses should use the Sign* functions in this package directly.




### <a id="PrivateKey.Validate">func</a> (\*PrivateKey) [Validate](https://golang.org/src/crypto/rsa/rsa.go?s=5617:5657#L159)
<pre>func (priv *<a href="#PrivateKey">PrivateKey</a>) Validate() <a href="/pkg/builtin/#error">error</a></pre>
Validate performs basic sanity checks on the key.
It returns nil if the key is valid, or else an error describing a problem.




## <a id="PublicKey">type</a> [PublicKey](https://golang.org/src/crypto/rsa/rsa.go?s=1400:1479#L32)
A PublicKey represents the public part of an RSA key.


<pre>type PublicKey struct {
<span id="PublicKey.N"></span>    N *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a> <span class="comment">// modulus</span>
<span id="PublicKey.E"></span>    E <a href="/pkg/builtin/#int">int</a>      <span class="comment">// public exponent</span>
}
</pre>











### <a id="PublicKey.Size">func</a> (\*PublicKey) [Size](https://golang.org/src/crypto/rsa/rsa.go?s=1609:1641#L39)
<pre>func (pub *<a href="#PublicKey">PublicKey</a>) Size() <a href="/pkg/builtin/#int">int</a></pre>
Size returns the modulus size in bytes. Raw signatures and ciphertexts
for or by this public key will have the same size.







