

# cipher
`import "crypto/cipher"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package cipher implements standard block cipher modes that can be wrapped
around low-level block cipher implementations.
See <a href="https://csrc.nist.gov/groups/ST/toolkit/BCM/current_modes.html">https://csrc.nist.gov/groups/ST/toolkit/BCM/current_modes.html</a>
and NIST Special Publication 800-38A.




## <a id="pkg-index">Index</a>
* [type AEAD](#AEAD)
  * [func NewGCM(cipher Block) (AEAD, error)](#NewGCM)
  * [func NewGCMWithNonceSize(cipher Block, size int) (AEAD, error)](#NewGCMWithNonceSize)
  * [func NewGCMWithTagSize(cipher Block, tagSize int) (AEAD, error)](#NewGCMWithTagSize)
* [type Block](#Block)
* [type BlockMode](#BlockMode)
  * [func NewCBCDecrypter(b Block, iv []byte) BlockMode](#NewCBCDecrypter)
  * [func NewCBCEncrypter(b Block, iv []byte) BlockMode](#NewCBCEncrypter)
* [type Stream](#Stream)
  * [func NewCFBDecrypter(block Block, iv []byte) Stream](#NewCFBDecrypter)
  * [func NewCFBEncrypter(block Block, iv []byte) Stream](#NewCFBEncrypter)
  * [func NewCTR(block Block, iv []byte) Stream](#NewCTR)
  * [func NewOFB(b Block, iv []byte) Stream](#NewOFB)
* [type StreamReader](#StreamReader)
  * [func (r StreamReader) Read(dst []byte) (n int, err error)](#StreamReader.Read)
* [type StreamWriter](#StreamWriter)
  * [func (w StreamWriter) Close() error](#StreamWriter.Close)
  * [func (w StreamWriter) Write(src []byte) (n int, err error)](#StreamWriter.Write)


#### <a id="pkg-examples">Examples</a>
* [NewCBCDecrypter](#example_NewCBCDecrypter)
* [NewCBCEncrypter](#example_NewCBCEncrypter)
* [NewCFBDecrypter](#example_NewCFBDecrypter)
* [NewCFBEncrypter](#example_NewCFBEncrypter)
* [NewCTR](#example_NewCTR)
* [NewGCM (Decrypt)](#example_NewGCM_decrypt)
* [NewGCM (Encrypt)](#example_NewGCM_encrypt)
* [NewOFB](#example_NewOFB)
* [StreamReader](#example_StreamReader)
* [StreamWriter](#example_StreamWriter)


#### <a id="pkg-files">Package files</a>
[cbc.go](https://golang.org/src/crypto/cipher/cbc.go) [cfb.go](https://golang.org/src/crypto/cipher/cfb.go) [cipher.go](https://golang.org/src/crypto/cipher/cipher.go) [ctr.go](https://golang.org/src/crypto/cipher/ctr.go) [gcm.go](https://golang.org/src/crypto/cipher/gcm.go) [io.go](https://golang.org/src/crypto/cipher/io.go) [ofb.go](https://golang.org/src/crypto/cipher/ofb.go) [xor_amd64.go](https://golang.org/src/crypto/cipher/xor_amd64.go) 








## <a id="AEAD">type</a> [AEAD](https://golang.org/src/crypto/cipher/gcm.go?s=459:1799#L7)
AEAD is a cipher mode providing authenticated encryption with associated
data. For a description of the methodology, see


	<a href="https://en.wikipedia.org/wiki/Authenticated_encryption">https://en.wikipedia.org/wiki/Authenticated_encryption</a>


<pre>type AEAD interface {
    <span class="comment">// NonceSize returns the size of the nonce that must be passed to Seal</span>
    <span class="comment">// and Open.</span>
    NonceSize() <a href="/pkg/builtin/#int">int</a>

    <span class="comment">// Overhead returns the maximum difference between the lengths of a</span>
    <span class="comment">// plaintext and its ciphertext.</span>
    Overhead() <a href="/pkg/builtin/#int">int</a>

    <span class="comment">// Seal encrypts and authenticates plaintext, authenticates the</span>
    <span class="comment">// additional data and appends the result to dst, returning the updated</span>
    <span class="comment">// slice. The nonce must be NonceSize() bytes long and unique for all</span>
    <span class="comment">// time, for a given key.</span>
    <span class="comment">//</span>
    <span class="comment">// To reuse plaintext&#39;s storage for the encrypted output, use plaintext[:0]</span>
    <span class="comment">// as dst. Otherwise, the remaining capacity of dst must not overlap plaintext.</span>
    Seal(dst, nonce, plaintext, additionalData []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a>

    <span class="comment">// Open decrypts and authenticates ciphertext, authenticates the</span>
    <span class="comment">// additional data and, if successful, appends the resulting plaintext</span>
    <span class="comment">// to dst, returning the updated slice. The nonce must be NonceSize()</span>
    <span class="comment">// bytes long and both it and the additional data must match the</span>
    <span class="comment">// value passed to Seal.</span>
    <span class="comment">//</span>
    <span class="comment">// To reuse ciphertext&#39;s storage for the decrypted output, use ciphertext[:0]</span>
    <span class="comment">// as dst. Otherwise, the remaining capacity of dst must not overlap plaintext.</span>
    <span class="comment">//</span>
    <span class="comment">// Even if the function fails, the contents of dst, up to its capacity,</span>
    <span class="comment">// may be overwritten.</span>
    Open(dst, nonce, ciphertext, additionalData []<a href="/pkg/builtin/#byte">byte</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)
}</pre>









### <a id="NewGCM">func</a> [NewGCM](https://golang.org/src/crypto/cipher/gcm.go?s=3384:3423#L74)
<pre>func NewGCM(cipher <a href="#Block">Block</a>) (<a href="#AEAD">AEAD</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewGCM returns the given 128-bit, block cipher wrapped in Galois Counter Mode
with the standard nonce length.

In general, the GHASH operation performed by this implementation of GCM is not constant-time.
An exception is when the underlying Block was created by aes.NewCipher
on systems with hardware support for AES. See the crypto/aes package documentation for details.


<a id="example_NewGCM_decrypt">Example (Decrypt)</a>
<a id="example_NewGCM_encrypt">Example (Encrypt)</a>


### <a id="NewGCMWithNonceSize">func</a> [NewGCMWithNonceSize](https://golang.org/src/crypto/cipher/gcm.go?s=3858:3920#L84)
<pre>func NewGCMWithNonceSize(cipher <a href="#Block">Block</a>, size <a href="/pkg/builtin/#int">int</a>) (<a href="#AEAD">AEAD</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewGCMWithNonceSize returns the given 128-bit, block cipher wrapped in Galois
Counter Mode, which accepts nonces of the given length.

Only use this function if you require compatibility with an existing
cryptosystem that uses non-standard nonce lengths. All other users should use
NewGCM, which is faster and more resistant to misuse.




### <a id="NewGCMWithTagSize">func</a> [NewGCMWithTagSize](https://golang.org/src/crypto/cipher/gcm.go?s=4379:4442#L96)
<pre>func NewGCMWithTagSize(cipher <a href="#Block">Block</a>, tagSize <a href="/pkg/builtin/#int">int</a>) (<a href="#AEAD">AEAD</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewGCMWithTagSize returns the given 128-bit, block cipher wrapped in Galois
Counter Mode, which generates tags with the given length.

Tag sizes between 12 and 16 bytes are allowed.

Only use this function if you require compatibility with an existing
cryptosystem that uses non-standard tag lengths. All other users should use
NewGCM, which is more resistant to misuse.






## <a id="Block">type</a> [Block](https://golang.org/src/crypto/cipher/cipher.go?s=636:992#L5)
A Block represents an implementation of block cipher
using a given key. It provides the capability to encrypt
or decrypt individual blocks. The mode implementations
extend that capability to streams of blocks.


<pre>type Block interface {
    <span class="comment">// BlockSize returns the cipher&#39;s block size.</span>
    BlockSize() <a href="/pkg/builtin/#int">int</a>

    <span class="comment">// Encrypt encrypts the first block in src into dst.</span>
    <span class="comment">// Dst and src must overlap entirely or not at all.</span>
    Encrypt(dst, src []<a href="/pkg/builtin/#byte">byte</a>)

    <span class="comment">// Decrypt decrypts the first block in src into dst.</span>
    <span class="comment">// Dst and src must overlap entirely or not at all.</span>
    Decrypt(dst, src []<a href="/pkg/builtin/#byte">byte</a>)
}</pre>











## <a id="BlockMode">type</a> [BlockMode](https://golang.org/src/crypto/cipher/cipher.go?s=1745:2452#L35)
A BlockMode represents a block cipher running in a block-based mode (CBC,
ECB etc).


<pre>type BlockMode interface {
    <span class="comment">// BlockSize returns the mode&#39;s block size.</span>
    BlockSize() <a href="/pkg/builtin/#int">int</a>

    <span class="comment">// CryptBlocks encrypts or decrypts a number of blocks. The length of</span>
    <span class="comment">// src must be a multiple of the block size. Dst and src must overlap</span>
    <span class="comment">// entirely or not at all.</span>
    <span class="comment">//</span>
    <span class="comment">// If len(dst) &lt; len(src), CryptBlocks should panic. It is acceptable</span>
    <span class="comment">// to pass a dst bigger than src, and in that case, CryptBlocks will</span>
    <span class="comment">// only update dst[:len(src)] and will not touch the rest of dst.</span>
    <span class="comment">//</span>
    <span class="comment">// Multiple calls to CryptBlocks behave as if the concatenation of</span>
    <span class="comment">// the src buffers was passed in a single run. That is, BlockMode</span>
    <span class="comment">// maintains state and does not reset at each CryptBlocks call.</span>
    CryptBlocks(dst, src []<a href="/pkg/builtin/#byte">byte</a>)
}</pre>









### <a id="NewCBCDecrypter">func</a> [NewCBCDecrypter](https://golang.org/src/crypto/cipher/cbc.go?s=2910:2960#L95)
<pre>func NewCBCDecrypter(b <a href="#Block">Block</a>, iv []<a href="/pkg/builtin/#byte">byte</a>) <a href="#BlockMode">BlockMode</a></pre>
NewCBCDecrypter returns a BlockMode which decrypts in cipher block chaining
mode, using the given Block. The length of iv must be the same as the
Block's block size and must match the iv used to encrypt the data.


<a id="example_NewCBCDecrypter">Example</a>


### <a id="NewCBCEncrypter">func</a> [NewCBCEncrypter](https://golang.org/src/crypto/cipher/cbc.go?s=1185:1235#L35)
<pre>func NewCBCEncrypter(b <a href="#Block">Block</a>, iv []<a href="/pkg/builtin/#byte">byte</a>) <a href="#BlockMode">BlockMode</a></pre>
NewCBCEncrypter returns a BlockMode which encrypts in cipher block chaining
mode, using the given Block. The length of iv must be the same as the
Block's block size.


<a id="example_NewCBCEncrypter">Example</a>




## <a id="Stream">type</a> [Stream](https://golang.org/src/crypto/cipher/cipher.go?s=1034:1653#L19)
A Stream represents a stream cipher.


<pre>type Stream interface {
    <span class="comment">// XORKeyStream XORs each byte in the given slice with a byte from the</span>
    <span class="comment">// cipher&#39;s key stream. Dst and src must overlap entirely or not at all.</span>
    <span class="comment">//</span>
    <span class="comment">// If len(dst) &lt; len(src), XORKeyStream should panic. It is acceptable</span>
    <span class="comment">// to pass a dst bigger than src, and in that case, XORKeyStream will</span>
    <span class="comment">// only update dst[:len(src)] and will not touch the rest of dst.</span>
    <span class="comment">//</span>
    <span class="comment">// Multiple calls to XORKeyStream behave as if the concatenation of</span>
    <span class="comment">// the src buffers was passed in a single run. That is, Stream</span>
    <span class="comment">// maintains state and does not reset at each XORKeyStream call.</span>
    XORKeyStream(dst, src []<a href="/pkg/builtin/#byte">byte</a>)
}</pre>









### <a id="NewCFBDecrypter">func</a> [NewCFBDecrypter](https://golang.org/src/crypto/cipher/cfb.go?s=1480:1531#L50)
<pre>func NewCFBDecrypter(block <a href="#Block">Block</a>, iv []<a href="/pkg/builtin/#byte">byte</a>) <a href="#Stream">Stream</a></pre>
NewCFBDecrypter returns a Stream which decrypts with cipher feedback mode,
using the given Block. The iv must be the same length as the Block's block
size.


<a id="example_NewCFBDecrypter">Example</a>


### <a id="NewCFBEncrypter">func</a> [NewCFBEncrypter](https://golang.org/src/crypto/cipher/cfb.go?s=1225:1276#L43)
<pre>func NewCFBEncrypter(block <a href="#Block">Block</a>, iv []<a href="/pkg/builtin/#byte">byte</a>) <a href="#Stream">Stream</a></pre>
NewCFBEncrypter returns a Stream which encrypts with cipher feedback mode,
using the given Block. The iv must be the same length as the Block's block
size.


<a id="example_NewCFBEncrypter">Example</a>


### <a id="NewCTR">func</a> [NewCTR](https://golang.org/src/crypto/cipher/ctr.go?s=955:997#L25)
<pre>func NewCTR(block <a href="#Block">Block</a>, iv []<a href="/pkg/builtin/#byte">byte</a>) <a href="#Stream">Stream</a></pre>
NewCTR returns a Stream which encrypts/decrypts using the given Block in
counter mode. The length of iv must be the same as the Block's block size.


<a id="example_NewCTR">Example</a>


### <a id="NewOFB">func</a> [NewOFB](https://golang.org/src/crypto/cipher/ofb.go?s=502:540#L11)
<pre>func NewOFB(b <a href="#Block">Block</a>, iv []<a href="/pkg/builtin/#byte">byte</a>) <a href="#Stream">Stream</a></pre>
NewOFB returns a Stream that encrypts or decrypts using the block cipher b
in output feedback mode. The initialization vector iv's length must be equal
to b's block size.


<a id="example_NewOFB">Example</a>




## <a id="StreamReader">type</a> [StreamReader](https://golang.org/src/crypto/cipher/io.go?s=426:477#L4)
StreamReader wraps a Stream into an io.Reader. It calls XORKeyStream
to process each slice of data which passes through.


<pre>type StreamReader struct {
<span id="StreamReader.S"></span>    S <a href="#Stream">Stream</a>
<span id="StreamReader.R"></span>    R <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>
}
</pre>





<a id="example_StreamReader">Example</a>






### <a id="StreamReader.Read">func</a> (StreamReader) [Read](https://golang.org/src/crypto/cipher/io.go?s=479:536#L9)
<pre>func (r <a href="#StreamReader">StreamReader</a>) Read(dst []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>



## <a id="StreamWriter">type</a> [StreamWriter](https://golang.org/src/crypto/cipher/io.go?s=934:1010#L20)
StreamWriter wraps a Stream into an io.Writer. It calls XORKeyStream
to process each slice of data which passes through. If any Write call
returns short then the StreamWriter is out of sync and must be discarded.
A StreamWriter has no internal buffering; Close does not need
to be called to flush write data.


<pre>type StreamWriter struct {
<span id="StreamWriter.S"></span>    S   <a href="#Stream">Stream</a>
<span id="StreamWriter.W"></span>    W   <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>
<span id="StreamWriter.Err"></span>    Err <a href="/pkg/builtin/#error">error</a> <span class="comment">// unused</span>
}
</pre>





<a id="example_StreamWriter">Example</a>






### <a id="StreamWriter.Close">func</a> (StreamWriter) [Close](https://golang.org/src/crypto/cipher/io.go?s=1386:1421#L38)
<pre>func (w <a href="#StreamWriter">StreamWriter</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the underlying Writer and returns its Close return value, if the Writer
is also an io.Closer. Otherwise it returns nil.




### <a id="StreamWriter.Write">func</a> (StreamWriter) [Write](https://golang.org/src/crypto/cipher/io.go?s=1012:1070#L26)
<pre>func (w <a href="#StreamWriter">StreamWriter</a>) Write(src []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>







