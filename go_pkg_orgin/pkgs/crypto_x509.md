

# x509
`import "crypto/x509"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package x509 parses X.509-encoded keys and certificates.

On UNIX systems the environment variables SSL_CERT_FILE and SSL_CERT_DIR
can be used to override the system default locations for the SSL certificate
file and SSL certificate files directory, respectively.




## <a id="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func CreateCertificate(rand io.Reader, template, parent *Certificate, pub, priv interface{}) (cert []byte, err error)](#CreateCertificate)
* [func CreateCertificateRequest(rand io.Reader, template *CertificateRequest, priv interface{}) (csr []byte, err error)](#CreateCertificateRequest)
* [func DecryptPEMBlock(b *pem.Block, password []byte) ([]byte, error)](#DecryptPEMBlock)
* [func EncryptPEMBlock(rand io.Reader, blockType string, data, password []byte, alg PEMCipher) (*pem.Block, error)](#EncryptPEMBlock)
* [func IsEncryptedPEMBlock(b *pem.Block) bool](#IsEncryptedPEMBlock)
* [func MarshalECPrivateKey(key *ecdsa.PrivateKey) ([]byte, error)](#MarshalECPrivateKey)
* [func MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte](#MarshalPKCS1PrivateKey)
* [func MarshalPKCS1PublicKey(key *rsa.PublicKey) []byte](#MarshalPKCS1PublicKey)
* [func MarshalPKCS8PrivateKey(key interface{}) ([]byte, error)](#MarshalPKCS8PrivateKey)
* [func MarshalPKIXPublicKey(pub interface{}) ([]byte, error)](#MarshalPKIXPublicKey)
* [func ParseCRL(crlBytes []byte) (*pkix.CertificateList, error)](#ParseCRL)
* [func ParseDERCRL(derBytes []byte) (*pkix.CertificateList, error)](#ParseDERCRL)
* [func ParseECPrivateKey(der []byte) (*ecdsa.PrivateKey, error)](#ParseECPrivateKey)
* [func ParsePKCS1PrivateKey(der []byte) (*rsa.PrivateKey, error)](#ParsePKCS1PrivateKey)
* [func ParsePKCS1PublicKey(der []byte) (*rsa.PublicKey, error)](#ParsePKCS1PublicKey)
* [func ParsePKCS8PrivateKey(der []byte) (key interface{}, err error)](#ParsePKCS8PrivateKey)
* [func ParsePKIXPublicKey(derBytes []byte) (pub interface{}, err error)](#ParsePKIXPublicKey)
* [type CertPool](#CertPool)
  * [func NewCertPool() *CertPool](#NewCertPool)
  * [func SystemCertPool() (*CertPool, error)](#SystemCertPool)
  * [func (s *CertPool) AddCert(cert *Certificate)](#CertPool.AddCert)
  * [func (s *CertPool) AppendCertsFromPEM(pemCerts []byte) (ok bool)](#CertPool.AppendCertsFromPEM)
  * [func (s *CertPool) Subjects() [][]byte](#CertPool.Subjects)
* [type Certificate](#Certificate)
  * [func ParseCertificate(asn1Data []byte) (*Certificate, error)](#ParseCertificate)
  * [func ParseCertificates(asn1Data []byte) ([]*Certificate, error)](#ParseCertificates)
  * [func (c *Certificate) CheckCRLSignature(crl *pkix.CertificateList) error](#Certificate.CheckCRLSignature)
  * [func (c *Certificate) CheckSignature(algo SignatureAlgorithm, signed, signature []byte) error](#Certificate.CheckSignature)
  * [func (c *Certificate) CheckSignatureFrom(parent *Certificate) error](#Certificate.CheckSignatureFrom)
  * [func (c *Certificate) CreateCRL(rand io.Reader, priv interface{}, revokedCerts []pkix.RevokedCertificate, now, expiry time.Time) (crlBytes []byte, err error)](#Certificate.CreateCRL)
  * [func (c *Certificate) Equal(other *Certificate) bool](#Certificate.Equal)
  * [func (c *Certificate) Verify(opts VerifyOptions) (chains [][]*Certificate, err error)](#Certificate.Verify)
  * [func (c *Certificate) VerifyHostname(h string) error](#Certificate.VerifyHostname)
* [type CertificateInvalidError](#CertificateInvalidError)
  * [func (e CertificateInvalidError) Error() string](#CertificateInvalidError.Error)
* [type CertificateRequest](#CertificateRequest)
  * [func ParseCertificateRequest(asn1Data []byte) (*CertificateRequest, error)](#ParseCertificateRequest)
  * [func (c *CertificateRequest) CheckSignature() error](#CertificateRequest.CheckSignature)
* [type ConstraintViolationError](#ConstraintViolationError)
  * [func (ConstraintViolationError) Error() string](#ConstraintViolationError.Error)
* [type ExtKeyUsage](#ExtKeyUsage)
* [type HostnameError](#HostnameError)
  * [func (h HostnameError) Error() string](#HostnameError.Error)
* [type InsecureAlgorithmError](#InsecureAlgorithmError)
  * [func (e InsecureAlgorithmError) Error() string](#InsecureAlgorithmError.Error)
* [type InvalidReason](#InvalidReason)
* [type KeyUsage](#KeyUsage)
* [type PEMCipher](#PEMCipher)
* [type PublicKeyAlgorithm](#PublicKeyAlgorithm)
  * [func (algo PublicKeyAlgorithm) String() string](#PublicKeyAlgorithm.String)
* [type SignatureAlgorithm](#SignatureAlgorithm)
  * [func (algo SignatureAlgorithm) String() string](#SignatureAlgorithm.String)
* [type SystemRootsError](#SystemRootsError)
  * [func (se SystemRootsError) Error() string](#SystemRootsError.Error)
* [type UnhandledCriticalExtension](#UnhandledCriticalExtension)
  * [func (h UnhandledCriticalExtension) Error() string](#UnhandledCriticalExtension.Error)
* [type UnknownAuthorityError](#UnknownAuthorityError)
  * [func (e UnknownAuthorityError) Error() string](#UnknownAuthorityError.Error)
* [type VerifyOptions](#VerifyOptions)


#### <a id="pkg-examples">Examples</a>
* [Certificate.Verify](#example_Certificate_Verify)
* [ParsePKIXPublicKey](#example_ParsePKIXPublicKey)


#### <a id="pkg-files">Package files</a>
[cert_pool.go](https://golang.org/src/crypto/x509/cert_pool.go) [pem_decrypt.go](https://golang.org/src/crypto/x509/pem_decrypt.go) [pkcs1.go](https://golang.org/src/crypto/x509/pkcs1.go) [pkcs8.go](https://golang.org/src/crypto/x509/pkcs8.go) [root.go](https://golang.org/src/crypto/x509/root.go) [root_linux.go](https://golang.org/src/crypto/x509/root_linux.go) [root_unix.go](https://golang.org/src/crypto/x509/root_unix.go) [sec1.go](https://golang.org/src/crypto/x509/sec1.go) [verify.go](https://golang.org/src/crypto/x509/verify.go) [x509.go](https://golang.org/src/crypto/x509/x509.go) 




## <a id="pkg-variables">Variables</a>
ErrUnsupportedAlgorithm results from attempting to perform an operation that
involves algorithms that are not currently implemented.


<pre>var <span id="ErrUnsupportedAlgorithm">ErrUnsupportedAlgorithm</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;x509: cannot verify signature: algorithm unimplemented&#34;)</pre>IncorrectPasswordError is returned when an incorrect password is detected.


<pre>var <span id="IncorrectPasswordError">IncorrectPasswordError</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;x509: decryption password incorrect&#34;)</pre>

## <a id="CreateCertificate">func</a> [CreateCertificate](https://golang.org/src/crypto/x509/x509.go?s=68468:68585#L2126)
<pre>func CreateCertificate(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, template, parent *<a href="#Certificate">Certificate</a>, pub, priv interface{}) (cert []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
CreateCertificate creates a new X.509v3 certificate based on a template.
The following members of template are used:


	- AuthorityKeyId
	- BasicConstraintsValid
	- CRLDistributionPoints
	- DNSNames
	- EmailAddresses
	- ExcludedDNSDomains
	- ExcludedEmailAddresses
	- ExcludedIPRanges
	- ExcludedURIDomains
	- ExtKeyUsage
	- ExtraExtensions
	- IsCA
	- IssuingCertificateURL
	- KeyUsage
	- MaxPathLen
	- MaxPathLenZero
	- NotAfter
	- NotBefore
	- OCSPServer
	- PermittedDNSDomains
	- PermittedDNSDomainsCritical
	- PermittedEmailAddresses
	- PermittedIPRanges
	- PermittedURIDomains
	- PolicyIdentifiers
	- SerialNumber
	- SignatureAlgorithm
	- Subject
	- SubjectKeyId
	- URIs
	- UnknownExtKeyUsage

The certificate is signed by parent. If parent is equal to template then the
certificate is self-signed. The parameter pub is the public key of the
signee and priv is the private key of the signer.

The returned slice is the certificate in DER encoding.

The currently supported key types are *rsa.PublicKey, *ecdsa.PublicKey and
ed25519.PublicKey. pub must be a supported key type, and priv must be a
crypto.Signer with a supported public key.

The AuthorityKeyId will be taken from the SubjectKeyId of parent, if any,
unless the resulting certificate is self-signed. Otherwise the value from
template will be used.



## <a id="CreateCertificateRequest">func</a> [CreateCertificateRequest](https://golang.org/src/crypto/x509/x509.go?s=78680:78797#L2456)
<pre>func CreateCertificateRequest(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, template *<a href="#CertificateRequest">CertificateRequest</a>, priv interface{}) (csr []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
CreateCertificateRequest creates a new certificate request based on a
template. The following members of template are used:


	- SignatureAlgorithm
	- Subject
	- DNSNames
	- EmailAddresses
	- IPAddresses
	- URIs
	- ExtraExtensions
	- Attributes (deprecated)

priv is the private key to sign the CSR with, and the corresponding public
key will be included in the CSR. It must implement crypto.Signer and its
Public() method must return a *rsa.PublicKey or a *ecdsa.PublicKey or a
ed25519.PublicKey. (A *rsa.PrivateKey, *ecdsa.PrivateKey or
ed25519.PrivateKey satisfies this.)

The returned slice is the certificate request in DER encoding.



## <a id="DecryptPEMBlock">func</a> [DecryptPEMBlock](https://golang.org/src/crypto/x509/pem_decrypt.go?s=3143:3210#L105)
<pre>func DecryptPEMBlock(b *<a href="/pkg/encoding/pem/">pem</a>.<a href="/pkg/encoding/pem/#Block">Block</a>, password []<a href="/pkg/builtin/#byte">byte</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DecryptPEMBlock takes a password encrypted PEM block and the password used to
encrypt it and returns a slice of decrypted DER encoded bytes. It inspects
the DEK-Info header to determine the algorithm used for decryption. If no
DEK-Info header is present, an error is returned. If an incorrect password
is detected an IncorrectPasswordError is returned. Because of deficiencies
in the encrypted-PEM format, it's not always possible to detect an incorrect
password. In these cases no error will be returned but the decrypted DER
bytes will be random noise.



## <a id="EncryptPEMBlock">func</a> [EncryptPEMBlock](https://golang.org/src/crypto/x509/pem_decrypt.go?s=5072:5184#L173)
<pre>func EncryptPEMBlock(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, blockType <a href="/pkg/builtin/#string">string</a>, data, password []<a href="/pkg/builtin/#byte">byte</a>, alg <a href="#PEMCipher">PEMCipher</a>) (*<a href="/pkg/encoding/pem/">pem</a>.<a href="/pkg/encoding/pem/#Block">Block</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
EncryptPEMBlock returns a PEM block of the specified type holding the
given DER-encoded data encrypted with the specified algorithm and
password.



## <a id="IsEncryptedPEMBlock">func</a> [IsEncryptedPEMBlock](https://golang.org/src/crypto/x509/pem_decrypt.go?s=2314:2357#L89)
<pre>func IsEncryptedPEMBlock(b *<a href="/pkg/encoding/pem/">pem</a>.<a href="/pkg/encoding/pem/#Block">Block</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsEncryptedPEMBlock returns if the PEM block is password encrypted.



## <a id="MarshalECPrivateKey">func</a> [MarshalECPrivateKey](https://golang.org/src/crypto/x509/sec1.go?s=1256:1319#L33)
<pre>func MarshalECPrivateKey(key *<a href="/pkg/crypto/ecdsa/">ecdsa</a>.<a href="/pkg/crypto/ecdsa/#PrivateKey">PrivateKey</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
MarshalECPrivateKey converts an EC private key to SEC 1, ASN.1 DER form.

This kind of key is commonly encoded in PEM blocks of type "EC PRIVATE KEY".
For a more flexible key format which is not EC specific, use
MarshalPKCS8PrivateKey.



## <a id="MarshalPKCS1PrivateKey">func</a> [MarshalPKCS1PrivateKey](https://golang.org/src/crypto/x509/pkcs1.go?s=2856:2911#L94)
<pre>func MarshalPKCS1PrivateKey(key *<a href="/pkg/crypto/rsa/">rsa</a>.<a href="/pkg/crypto/rsa/#PrivateKey">PrivateKey</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
MarshalPKCS1PrivateKey converts an RSA private key to PKCS#1, ASN.1 DER form.

This kind of key is commonly encoded in PEM blocks of type "RSA PRIVATE KEY".
For a more flexible key format which is not RSA specific, use
MarshalPKCS8PrivateKey.



## <a id="MarshalPKCS1PublicKey">func</a> [MarshalPKCS1PublicKey](https://golang.org/src/crypto/x509/pkcs1.go?s=4608:4661#L157)
<pre>func MarshalPKCS1PublicKey(key *<a href="/pkg/crypto/rsa/">rsa</a>.<a href="/pkg/crypto/rsa/#PublicKey">PublicKey</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
MarshalPKCS1PublicKey converts an RSA public key to PKCS#1, ASN.1 DER form.

This kind of key is commonly encoded in PEM blocks of type "RSA PUBLIC KEY".



## <a id="MarshalPKCS8PrivateKey">func</a> [MarshalPKCS8PrivateKey](https://golang.org/src/crypto/x509/pkcs8.go?s=3121:3181#L78)
<pre>func MarshalPKCS8PrivateKey(key interface{}) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
MarshalPKCS8PrivateKey converts an RSA private key to PKCS#8, ASN.1 DER form.

The following key types are currently supported: *rsa.PrivateKey, *ecdsa.PrivateKey
and ed25519.PrivateKey. Unsupported key types result in an error.

This kind of key is commonly encoded in PEM blocks of type "PRIVATE KEY".



## <a id="MarshalPKIXPublicKey">func</a> [MarshalPKIXPublicKey](https://golang.org/src/crypto/x509/x509.go?s=3534:3592#L103)
<pre>func MarshalPKIXPublicKey(pub interface{}) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
MarshalPKIXPublicKey converts a public key to PKIX, ASN.1 DER form.

The following key types are currently supported: *rsa.PublicKey, *ecdsa.PublicKey
and ed25519.PublicKey. Unsupported key types result in an error.

This kind of key is commonly encoded in PEM blocks of type "PUBLIC KEY".



## <a id="ParseCRL">func</a> [ParseCRL](https://golang.org/src/crypto/x509/x509.go?s=71208:71269#L2224)
<pre>func ParseCRL(crlBytes []<a href="/pkg/builtin/#byte">byte</a>) (*<a href="/pkg/crypto/x509/pkix/">pkix</a>.<a href="/pkg/crypto/x509/pkix/#CertificateList">CertificateList</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseCRL parses a CRL from the given bytes. It's often the case that PEM
encoded CRLs will appear where they should be DER encoded, so this function
will transparently handle PEM encoding as long as there isn't any leading
garbage.



## <a id="ParseDERCRL">func</a> [ParseDERCRL](https://golang.org/src/crypto/x509/x509.go?s=71526:71590#L2235)
<pre>func ParseDERCRL(derBytes []<a href="/pkg/builtin/#byte">byte</a>) (*<a href="/pkg/crypto/x509/pkix/">pkix</a>.<a href="/pkg/crypto/x509/pkix/#CertificateList">CertificateList</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseDERCRL parses a DER encoded CRL from the given bytes.



## <a id="ParseECPrivateKey">func</a> [ParseECPrivateKey](https://golang.org/src/crypto/x509/sec1.go?s=903:964#L24)
<pre>func ParseECPrivateKey(der []<a href="/pkg/builtin/#byte">byte</a>) (*<a href="/pkg/crypto/ecdsa/">ecdsa</a>.<a href="/pkg/crypto/ecdsa/#PrivateKey">PrivateKey</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseECPrivateKey parses an EC public key in SEC 1, ASN.1 DER form.

This kind of key is commonly encoded in PEM blocks of type "EC PUBLIC KEY".



## <a id="ParsePKCS1PrivateKey">func</a> [ParsePKCS1PrivateKey](https://golang.org/src/crypto/x509/pkcs1.go?s=1137:1199#L37)
<pre>func ParsePKCS1PrivateKey(der []<a href="/pkg/builtin/#byte">byte</a>) (*<a href="/pkg/crypto/rsa/">rsa</a>.<a href="/pkg/crypto/rsa/#PrivateKey">PrivateKey</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParsePKCS1PrivateKey parses an RSA private key in PKCS#1, ASN.1 DER form.

This kind of key is commonly encoded in PEM blocks of type "RSA PRIVATE KEY".



## <a id="ParsePKCS1PublicKey">func</a> [ParsePKCS1PublicKey](https://golang.org/src/crypto/x509/pkcs1.go?s=3739:3799#L128)
<pre>func ParsePKCS1PublicKey(der []<a href="/pkg/builtin/#byte">byte</a>) (*<a href="/pkg/crypto/rsa/">rsa</a>.<a href="/pkg/crypto/rsa/#PublicKey">PublicKey</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParsePKCS1PublicKey parses an RSA public key in PKCS#1, ASN.1 DER form.

This kind of key is commonly encoded in PEM blocks of type "RSA PUBLIC KEY".



## <a id="ParsePKCS8PrivateKey">func</a> [ParsePKCS8PrivateKey](https://golang.org/src/crypto/x509/pkcs8.go?s=839:905#L23)
<pre>func ParsePKCS8PrivateKey(der []<a href="/pkg/builtin/#byte">byte</a>) (key interface{}, err <a href="/pkg/builtin/#error">error</a>)</pre>
ParsePKCS8PrivateKey parses an unencrypted private key in PKCS#8, ASN.1 DER form.

It returns a *rsa.PrivateKey, a *ecdsa.PrivateKey, or a ed25519.PrivateKey.
More types might be supported in the future.

This kind of key is commonly encoded in PEM blocks of type "PRIVATE KEY".



## <a id="ParsePKIXPublicKey">func</a> [ParsePKIXPublicKey](https://golang.org/src/crypto/x509/x509.go?s=1334:1403#L43)
<pre>func ParsePKIXPublicKey(derBytes []<a href="/pkg/builtin/#byte">byte</a>) (pub interface{}, err <a href="/pkg/builtin/#error">error</a>)</pre>
ParsePKIXPublicKey parses a public key in PKIX, ASN.1 DER form.

It returns a *rsa.PublicKey, *dsa.PublicKey, *ecdsa.PublicKey, or
ed25519.PublicKey. More types might be supported in the future.

This kind of key is commonly encoded in PEM blocks of type "PUBLIC KEY".


<a id="example_ParsePKIXPublicKey">Example</a>
```go
```

output:
```txt
```



## <a id="CertPool">type</a> [CertPool](https://golang.org/src/crypto/x509/cert_pool.go?s=261:382#L4)
CertPool is a set of certificates.


<pre>type CertPool struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewCertPool">func</a> [NewCertPool](https://golang.org/src/crypto/x509/cert_pool.go?s=430:458#L11)
<pre>func NewCertPool() *<a href="#CertPool">CertPool</a></pre>
NewCertPool returns a new, empty CertPool.




### <a id="SystemCertPool">func</a> [SystemCertPool](https://golang.org/src/crypto/x509/cert_pool.go?s=1351:1391#L45)
<pre>func SystemCertPool() (*<a href="#CertPool">CertPool</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
SystemCertPool returns a copy of the system cert pool.

Any mutations to the returned pool are not written to disk and do
not affect any other pool returned by SystemCertPool.

New changes in the system cert pool might not be reflected
in subsequent calls.






### <a id="CertPool.AddCert">func</a> (\*CertPool) [AddCert](https://golang.org/src/crypto/x509/cert_pool.go?s=2402:2447#L91)
<pre>func (s *<a href="#CertPool">CertPool</a>) AddCert(cert *<a href="#Certificate">Certificate</a>)</pre>
AddCert adds a certificate to a pool.




### <a id="CertPool.AppendCertsFromPEM">func</a> (\*CertPool) [AppendCertsFromPEM](https://golang.org/src/crypto/x509/cert_pool.go?s=3200:3264#L118)
<pre>func (s *<a href="#CertPool">CertPool</a>) AppendCertsFromPEM(pemCerts []<a href="/pkg/builtin/#byte">byte</a>) (ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
AppendCertsFromPEM attempts to parse a series of PEM encoded certificates.
It appends any certificates found to s and reports whether any certificates
were successfully parsed.

On many Linux systems, /etc/ssl/cert.pem will contain the system wide set
of root CAs in a format suitable for this function.




### <a id="CertPool.Subjects">func</a> (\*CertPool) [Subjects](https://golang.org/src/crypto/x509/cert_pool.go?s=3691:3729#L143)
<pre>func (s *<a href="#CertPool">CertPool</a>) Subjects() [][]<a href="/pkg/builtin/#byte">byte</a></pre>
Subjects returns a list of the DER-encoded subjects of
all of the certificates in the pool.




## <a id="Certificate">type</a> [Certificate](https://golang.org/src/crypto/x509/x509.go?s=21660:25707#L651)
A Certificate represents an X.509 certificate.


<pre>type Certificate struct {
<span id="Certificate.Raw"></span>    Raw                     []<a href="/pkg/builtin/#byte">byte</a> <span class="comment">// Complete ASN.1 DER content (certificate, signature algorithm and signature).</span>
<span id="Certificate.RawTBSCertificate"></span>    RawTBSCertificate       []<a href="/pkg/builtin/#byte">byte</a> <span class="comment">// Certificate part of raw ASN.1 DER content.</span>
<span id="Certificate.RawSubjectPublicKeyInfo"></span>    RawSubjectPublicKeyInfo []<a href="/pkg/builtin/#byte">byte</a> <span class="comment">// DER encoded SubjectPublicKeyInfo.</span>
<span id="Certificate.RawSubject"></span>    RawSubject              []<a href="/pkg/builtin/#byte">byte</a> <span class="comment">// DER encoded Subject</span>
<span id="Certificate.RawIssuer"></span>    RawIssuer               []<a href="/pkg/builtin/#byte">byte</a> <span class="comment">// DER encoded Issuer</span>

<span id="Certificate.Signature"></span>    Signature          []<a href="/pkg/builtin/#byte">byte</a>
<span id="Certificate.SignatureAlgorithm"></span>    SignatureAlgorithm <a href="#SignatureAlgorithm">SignatureAlgorithm</a>

<span id="Certificate.PublicKeyAlgorithm"></span>    PublicKeyAlgorithm <a href="#PublicKeyAlgorithm">PublicKeyAlgorithm</a>
<span id="Certificate.PublicKey"></span>    PublicKey          interface{}

<span id="Certificate.Version"></span>    Version             <a href="/pkg/builtin/#int">int</a>
<span id="Certificate.SerialNumber"></span>    SerialNumber        *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>
<span id="Certificate.Issuer"></span>    Issuer              <a href="/pkg/crypto/x509/pkix/">pkix</a>.<a href="/pkg/crypto/x509/pkix/#Name">Name</a>
<span id="Certificate.Subject"></span>    Subject             <a href="/pkg/crypto/x509/pkix/">pkix</a>.<a href="/pkg/crypto/x509/pkix/#Name">Name</a>
<span id="Certificate.NotBefore"></span>    NotBefore, NotAfter <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a> <span class="comment">// Validity bounds.</span>
<span id="Certificate.KeyUsage"></span>    KeyUsage            <a href="#KeyUsage">KeyUsage</a>

<span id="Certificate.Extensions"></span>    <span class="comment">// Extensions contains raw X.509 extensions. When parsing certificates,</span>
    <span class="comment">// this can be used to extract non-critical extensions that are not</span>
    <span class="comment">// parsed by this package. When marshaling certificates, the Extensions</span>
    <span class="comment">// field is ignored, see ExtraExtensions.</span>
    Extensions []<a href="/pkg/crypto/x509/pkix/">pkix</a>.<a href="/pkg/crypto/x509/pkix/#Extension">Extension</a>

<span id="Certificate.ExtraExtensions"></span>    <span class="comment">// ExtraExtensions contains extensions to be copied, raw, into any</span>
    <span class="comment">// marshaled certificates. Values override any extensions that would</span>
    <span class="comment">// otherwise be produced based on the other fields. The ExtraExtensions</span>
    <span class="comment">// field is not populated when parsing certificates, see Extensions.</span>
    ExtraExtensions []<a href="/pkg/crypto/x509/pkix/">pkix</a>.<a href="/pkg/crypto/x509/pkix/#Extension">Extension</a>

<span id="Certificate.UnhandledCriticalExtensions"></span>    <span class="comment">// UnhandledCriticalExtensions contains a list of extension IDs that</span>
    <span class="comment">// were not (fully) processed when parsing. Verify will fail if this</span>
    <span class="comment">// slice is non-empty, unless verification is delegated to an OS</span>
    <span class="comment">// library which understands all the critical extensions.</span>
    <span class="comment">//</span>
    <span class="comment">// Users can access these extensions using Extensions and can remove</span>
    <span class="comment">// elements from this slice if they believe that they have been</span>
    <span class="comment">// handled.</span>
    UnhandledCriticalExtensions []<a href="/pkg/encoding/asn1/">asn1</a>.<a href="/pkg/encoding/asn1/#ObjectIdentifier">ObjectIdentifier</a>

<span id="Certificate.ExtKeyUsage"></span>    ExtKeyUsage        []<a href="#ExtKeyUsage">ExtKeyUsage</a>           <span class="comment">// Sequence of extended key usages.</span>
<span id="Certificate.UnknownExtKeyUsage"></span>    UnknownExtKeyUsage []<a href="/pkg/encoding/asn1/">asn1</a>.<a href="/pkg/encoding/asn1/#ObjectIdentifier">ObjectIdentifier</a> <span class="comment">// Encountered extended key usages unknown to this package.</span>

<span id="Certificate.BasicConstraintsValid"></span>    <span class="comment">// BasicConstraintsValid indicates whether IsCA, MaxPathLen,</span>
    <span class="comment">// and MaxPathLenZero are valid.</span>
    BasicConstraintsValid <a href="/pkg/builtin/#bool">bool</a>
<span id="Certificate.IsCA"></span>    IsCA                  <a href="/pkg/builtin/#bool">bool</a>

<span id="Certificate.MaxPathLen"></span>    <span class="comment">// MaxPathLen and MaxPathLenZero indicate the presence and</span>
    <span class="comment">// value of the BasicConstraints&#39; &#34;pathLenConstraint&#34;.</span>
    <span class="comment">//</span>
    <span class="comment">// When parsing a certificate, a positive non-zero MaxPathLen</span>
    <span class="comment">// means that the field was specified, -1 means it was unset,</span>
    <span class="comment">// and MaxPathLenZero being true mean that the field was</span>
    <span class="comment">// explicitly set to zero. The case of MaxPathLen==0 with MaxPathLenZero==false</span>
    <span class="comment">// should be treated equivalent to -1 (unset).</span>
    <span class="comment">//</span>
    <span class="comment">// When generating a certificate, an unset pathLenConstraint</span>
    <span class="comment">// can be requested with either MaxPathLen == -1 or using the</span>
    <span class="comment">// zero value for both MaxPathLen and MaxPathLenZero.</span>
    MaxPathLen <a href="/pkg/builtin/#int">int</a>
<span id="Certificate.MaxPathLenZero"></span>    <span class="comment">// MaxPathLenZero indicates that BasicConstraintsValid==true</span>
    <span class="comment">// and MaxPathLen==0 should be interpreted as an actual</span>
    <span class="comment">// maximum path length of zero. Otherwise, that combination is</span>
    <span class="comment">// interpreted as MaxPathLen not being set.</span>
    MaxPathLenZero <a href="/pkg/builtin/#bool">bool</a>

<span id="Certificate.SubjectKeyId"></span>    SubjectKeyId   []<a href="/pkg/builtin/#byte">byte</a>
<span id="Certificate.AuthorityKeyId"></span>    AuthorityKeyId []<a href="/pkg/builtin/#byte">byte</a>

    <span class="comment">// RFC 5280, 4.2.2.1 (Authority Information Access)</span>
<span id="Certificate.OCSPServer"></span>    OCSPServer            []<a href="/pkg/builtin/#string">string</a>
<span id="Certificate.IssuingCertificateURL"></span>    IssuingCertificateURL []<a href="/pkg/builtin/#string">string</a>

    <span class="comment">// Subject Alternate Name values. (Note that these values may not be valid</span>
    <span class="comment">// if invalid values were contained within a parsed certificate. For</span>
    <span class="comment">// example, an element of DNSNames may not be a valid DNS domain name.)</span>
<span id="Certificate.DNSNames"></span>    DNSNames       []<a href="/pkg/builtin/#string">string</a>
<span id="Certificate.EmailAddresses"></span>    EmailAddresses []<a href="/pkg/builtin/#string">string</a>
<span id="Certificate.IPAddresses"></span>    IPAddresses    []<a href="/pkg/net/">net</a>.<a href="/pkg/net/#IP">IP</a>
<span id="Certificate.URIs"></span>    URIs           []*<a href="/pkg/net/url/">url</a>.<a href="/pkg/net/url/#URL">URL</a>

    <span class="comment">// Name constraints</span>
<span id="Certificate.PermittedDNSDomainsCritical"></span>    PermittedDNSDomainsCritical <a href="/pkg/builtin/#bool">bool</a> <span class="comment">// if true then the name constraints are marked critical.</span>
<span id="Certificate.PermittedDNSDomains"></span>    PermittedDNSDomains         []<a href="/pkg/builtin/#string">string</a>
<span id="Certificate.ExcludedDNSDomains"></span>    ExcludedDNSDomains          []<a href="/pkg/builtin/#string">string</a>
<span id="Certificate.PermittedIPRanges"></span>    PermittedIPRanges           []*<a href="/pkg/net/">net</a>.<a href="/pkg/net/#IPNet">IPNet</a>
<span id="Certificate.ExcludedIPRanges"></span>    ExcludedIPRanges            []*<a href="/pkg/net/">net</a>.<a href="/pkg/net/#IPNet">IPNet</a>
<span id="Certificate.PermittedEmailAddresses"></span>    PermittedEmailAddresses     []<a href="/pkg/builtin/#string">string</a>
<span id="Certificate.ExcludedEmailAddresses"></span>    ExcludedEmailAddresses      []<a href="/pkg/builtin/#string">string</a>
<span id="Certificate.PermittedURIDomains"></span>    PermittedURIDomains         []<a href="/pkg/builtin/#string">string</a>
<span id="Certificate.ExcludedURIDomains"></span>    ExcludedURIDomains          []<a href="/pkg/builtin/#string">string</a>

    <span class="comment">// CRL Distribution Points</span>
<span id="Certificate.CRLDistributionPoints"></span>    CRLDistributionPoints []<a href="/pkg/builtin/#string">string</a>

<span id="Certificate.PolicyIdentifiers"></span>    PolicyIdentifiers []<a href="/pkg/encoding/asn1/">asn1</a>.<a href="/pkg/encoding/asn1/#ObjectIdentifier">ObjectIdentifier</a>
}
</pre>









### <a id="ParseCertificate">func</a> [ParseCertificate](https://golang.org/src/crypto/x509/x509.go?s=53077:53137#L1601)
<pre>func ParseCertificate(asn1Data []<a href="/pkg/builtin/#byte">byte</a>) (*<a href="#Certificate">Certificate</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseCertificate parses a single certificate from the given ASN.1 DER data.




### <a id="ParseCertificates">func</a> [ParseCertificates](https://golang.org/src/crypto/x509/x509.go?s=53513:53576#L1616)
<pre>func ParseCertificates(asn1Data []<a href="/pkg/builtin/#byte">byte</a>) ([]*<a href="#Certificate">Certificate</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseCertificates parses one or more certificates from the given ASN.1 DER
data. The certificates must be concatenated with no intermediate padding.






### <a id="Certificate.CheckCRLSignature">func</a> (\*Certificate) [CheckCRLSignature](https://golang.org/src/crypto/x509/x509.go?s=33520:33592#L966)
<pre>func (c *<a href="#Certificate">Certificate</a>) CheckCRLSignature(crl *<a href="/pkg/crypto/x509/pkix/">pkix</a>.<a href="/pkg/crypto/x509/pkix/#CertificateList">CertificateList</a>) <a href="/pkg/builtin/#error">error</a></pre>
CheckCRLSignature checks that the signature in crl is from c.




### <a id="Certificate.CheckSignature">func</a> (\*Certificate) [CheckSignature](https://golang.org/src/crypto/x509/x509.go?s=30196:30289#L859)
<pre>func (c *<a href="#Certificate">Certificate</a>) CheckSignature(algo <a href="#SignatureAlgorithm">SignatureAlgorithm</a>, signed, signature []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
CheckSignature verifies that signature is a valid signature over signed from
c's public key.




### <a id="Certificate.CheckSignatureFrom">func</a> (\*Certificate) [CheckSignatureFrom](https://golang.org/src/crypto/x509/x509.go?s=29134:29201#L831)
<pre>func (c *<a href="#Certificate">Certificate</a>) CheckSignatureFrom(parent *<a href="#Certificate">Certificate</a>) <a href="/pkg/builtin/#error">error</a></pre>
CheckSignatureFrom verifies that the signature on c is a valid signature
from parent.




### <a id="Certificate.CreateCRL">func</a> (\*Certificate) [CreateCRL](https://golang.org/src/crypto/x509/x509.go?s=71955:72112#L2247)
<pre>func (c *<a href="#Certificate">Certificate</a>) CreateCRL(rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, priv interface{}, revokedCerts []<a href="/pkg/crypto/x509/pkix/">pkix</a>.<a href="/pkg/crypto/x509/pkix/#RevokedCertificate">RevokedCertificate</a>, now, expiry <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) (crlBytes []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
CreateCRL returns a DER encoded CRL, signed by this Certificate, that
contains the given list of revoked certificates.




### <a id="Certificate.Equal">func</a> (\*Certificate) [Equal](https://golang.org/src/crypto/x509/x509.go?s=26546:26598#L772)
<pre>func (c *<a href="#Certificate">Certificate</a>) Equal(other *<a href="#Certificate">Certificate</a>) <a href="/pkg/builtin/#bool">bool</a></pre>



### <a id="Certificate.Verify">func</a> (\*Certificate) [Verify](https://golang.org/src/crypto/x509/verify.go?s=23966:24051#L714)
<pre>func (c *<a href="#Certificate">Certificate</a>) Verify(opts <a href="#VerifyOptions">VerifyOptions</a>) (chains [][]*<a href="#Certificate">Certificate</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Verify attempts to verify c by building one or more chains from c to a
certificate in opts.Roots, using certificates in opts.Intermediates if
needed. If successful, it returns one or more chains where the first
element of the chain is c and the last element is from opts.Roots.

If opts.Roots is nil and system roots are unavailable the returned error
will be of type SystemRootsError.

Name constraints in the intermediates will be applied to all names claimed
in the chain, not just opts.DNSName. Thus it is invalid for a leaf to claim
example.com if an intermediate doesn't permit it, even if example.com is not
the name being validated. Note that DirectoryName constraints are not
supported.

Extended Key Usage values are enforced down a chain, so an intermediate or
root that enumerates EKUs prevents a leaf from asserting an EKU not in that
list.

WARNING: this function doesn't do any revocation checking.


<a id="example_Certificate_Verify">Example</a>
```go
```

output:
```txt
```


### <a id="Certificate.VerifyHostname">func</a> (\*Certificate) [VerifyHostname](https://golang.org/src/crypto/x509/verify.go?s=31014:31066#L986)
<pre>func (c *<a href="#Certificate">Certificate</a>) VerifyHostname(h <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
VerifyHostname returns nil if c is a valid certificate for the named host.
Otherwise it returns an error describing the mismatch.




## <a id="CertificateInvalidError">type</a> [CertificateInvalidError](https://golang.org/src/crypto/x509/verify.go?s=2642:2739#L62)
CertificateInvalidError results when an odd error occurs. Users of this
library probably want to handle all these errors uniformly.


<pre>type CertificateInvalidError struct {
<span id="CertificateInvalidError.Cert"></span>    Cert   *<a href="#Certificate">Certificate</a>
<span id="CertificateInvalidError.Reason"></span>    Reason <a href="#InvalidReason">InvalidReason</a>
<span id="CertificateInvalidError.Detail"></span>    Detail <a href="/pkg/builtin/#string">string</a>
}
</pre>











### <a id="CertificateInvalidError.Error">func</a> (CertificateInvalidError) [Error](https://golang.org/src/crypto/x509/verify.go?s=2741:2788#L68)
<pre>func (e <a href="#CertificateInvalidError">CertificateInvalidError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="CertificateRequest">type</a> [CertificateRequest](https://golang.org/src/crypto/x509/x509.go?s=73706:75268#L2308)
CertificateRequest represents a PKCS #10, certificate signature request.


<pre>type CertificateRequest struct {
<span id="CertificateRequest.Raw"></span>    Raw                      []<a href="/pkg/builtin/#byte">byte</a> <span class="comment">// Complete ASN.1 DER content (CSR, signature algorithm and signature).</span>
<span id="CertificateRequest.RawTBSCertificateRequest"></span>    RawTBSCertificateRequest []<a href="/pkg/builtin/#byte">byte</a> <span class="comment">// Certificate request info part of raw ASN.1 DER content.</span>
<span id="CertificateRequest.RawSubjectPublicKeyInfo"></span>    RawSubjectPublicKeyInfo  []<a href="/pkg/builtin/#byte">byte</a> <span class="comment">// DER encoded SubjectPublicKeyInfo.</span>
<span id="CertificateRequest.RawSubject"></span>    RawSubject               []<a href="/pkg/builtin/#byte">byte</a> <span class="comment">// DER encoded Subject.</span>

<span id="CertificateRequest.Version"></span>    Version            <a href="/pkg/builtin/#int">int</a>
<span id="CertificateRequest.Signature"></span>    Signature          []<a href="/pkg/builtin/#byte">byte</a>
<span id="CertificateRequest.SignatureAlgorithm"></span>    SignatureAlgorithm <a href="#SignatureAlgorithm">SignatureAlgorithm</a>

<span id="CertificateRequest.PublicKeyAlgorithm"></span>    PublicKeyAlgorithm <a href="#PublicKeyAlgorithm">PublicKeyAlgorithm</a>
<span id="CertificateRequest.PublicKey"></span>    PublicKey          interface{}

<span id="CertificateRequest.Subject"></span>    Subject <a href="/pkg/crypto/x509/pkix/">pkix</a>.<a href="/pkg/crypto/x509/pkix/#Name">Name</a>

<span id="CertificateRequest.Attributes"></span>    <span class="comment">// Attributes contains the CSR attributes that can parse as</span>
    <span class="comment">// pkix.AttributeTypeAndValueSET.</span>
    <span class="comment">//</span>
    <span class="comment">// Deprecated: Use Extensions and ExtraExtensions instead for parsing and</span>
    <span class="comment">// generating the requestedExtensions attribute.</span>
    Attributes []<a href="/pkg/crypto/x509/pkix/">pkix</a>.<a href="/pkg/crypto/x509/pkix/#AttributeTypeAndValueSET">AttributeTypeAndValueSET</a>

<span id="CertificateRequest.Extensions"></span>    <span class="comment">// Extensions contains all requested extensions, in raw form. When parsing</span>
    <span class="comment">// CSRs, this can be used to extract extensions that are not parsed by this</span>
    <span class="comment">// package.</span>
    Extensions []<a href="/pkg/crypto/x509/pkix/">pkix</a>.<a href="/pkg/crypto/x509/pkix/#Extension">Extension</a>

<span id="CertificateRequest.ExtraExtensions"></span>    <span class="comment">// ExtraExtensions contains extensions to be copied, raw, into any CSR</span>
    <span class="comment">// marshaled by CreateCertificateRequest. Values override any extensions</span>
    <span class="comment">// that would otherwise be produced based on the other fields but are</span>
    <span class="comment">// overridden by any extensions specified in Attributes.</span>
    <span class="comment">//</span>
    <span class="comment">// The ExtraExtensions field is not populated by ParseCertificateRequest,</span>
    <span class="comment">// see Extensions instead.</span>
    ExtraExtensions []<a href="/pkg/crypto/x509/pkix/">pkix</a>.<a href="/pkg/crypto/x509/pkix/#Extension">Extension</a>

    <span class="comment">// Subject Alternate Name values.</span>
<span id="CertificateRequest.DNSNames"></span>    DNSNames       []<a href="/pkg/builtin/#string">string</a>
<span id="CertificateRequest.EmailAddresses"></span>    EmailAddresses []<a href="/pkg/builtin/#string">string</a>
<span id="CertificateRequest.IPAddresses"></span>    IPAddresses    []<a href="/pkg/net/">net</a>.<a href="/pkg/net/#IP">IP</a>
<span id="CertificateRequest.URIs"></span>    URIs           []*<a href="/pkg/net/url/">url</a>.<a href="/pkg/net/url/#URL">URL</a>
}
</pre>









### <a id="ParseCertificateRequest">func</a> [ParseCertificateRequest](https://golang.org/src/crypto/x509/x509.go?s=83278:83352#L2627)
<pre>func ParseCertificateRequest(asn1Data []<a href="/pkg/builtin/#byte">byte</a>) (*<a href="#CertificateRequest">CertificateRequest</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseCertificateRequest parses a single certificate request from the
given ASN.1 DER data.






### <a id="CertificateRequest.CheckSignature">func</a> (\*CertificateRequest) [CheckSignature](https://golang.org/src/crypto/x509/x509.go?s=85070:85121#L2688)
<pre>func (c *<a href="#CertificateRequest">CertificateRequest</a>) CheckSignature() <a href="/pkg/builtin/#error">error</a></pre>
CheckSignature reports whether the signature on c is valid.




## <a id="ConstraintViolationError">type</a> [ConstraintViolationError](https://golang.org/src/crypto/x509/x509.go?s=26363:26401#L766)
ConstraintViolationError results when a requested usage is not permitted by
a certificate. For example: checking a signature when the public key isn't a
certificate signing key.


<pre>type ConstraintViolationError struct{}
</pre>











### <a id="ConstraintViolationError.Error">func</a> (ConstraintViolationError) [Error](https://golang.org/src/crypto/x509/x509.go?s=26403:26449#L768)
<pre>func (<a href="#ConstraintViolationError">ConstraintViolationError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="ExtKeyUsage">type</a> [ExtKeyUsage](https://golang.org/src/crypto/x509/x509.go?s=19722:19742#L592)
ExtKeyUsage represents an extended set of actions that are valid for a given key.
Each of the ExtKeyUsage* constants define a unique action.


<pre>type ExtKeyUsage <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="ExtKeyUsageAny">ExtKeyUsageAny</span> <a href="#ExtKeyUsage">ExtKeyUsage</a> = <a href="/pkg/builtin/#iota">iota</a>
    <span id="ExtKeyUsageServerAuth">ExtKeyUsageServerAuth</span>
    <span id="ExtKeyUsageClientAuth">ExtKeyUsageClientAuth</span>
    <span id="ExtKeyUsageCodeSigning">ExtKeyUsageCodeSigning</span>
    <span id="ExtKeyUsageEmailProtection">ExtKeyUsageEmailProtection</span>
    <span id="ExtKeyUsageIPSECEndSystem">ExtKeyUsageIPSECEndSystem</span>
    <span id="ExtKeyUsageIPSECTunnel">ExtKeyUsageIPSECTunnel</span>
    <span id="ExtKeyUsageIPSECUser">ExtKeyUsageIPSECUser</span>
    <span id="ExtKeyUsageTimeStamping">ExtKeyUsageTimeStamping</span>
    <span id="ExtKeyUsageOCSPSigning">ExtKeyUsageOCSPSigning</span>
    <span id="ExtKeyUsageMicrosoftServerGatedCrypto">ExtKeyUsageMicrosoftServerGatedCrypto</span>
    <span id="ExtKeyUsageNetscapeServerGatedCrypto">ExtKeyUsageNetscapeServerGatedCrypto</span>
    <span id="ExtKeyUsageMicrosoftCommercialCodeSigning">ExtKeyUsageMicrosoftCommercialCodeSigning</span>
    <span id="ExtKeyUsageMicrosoftKernelCodeSigning">ExtKeyUsageMicrosoftKernelCodeSigning</span>
)</pre>









## <a id="HostnameError">type</a> [HostnameError](https://golang.org/src/crypto/x509/verify.go?s=3933:4008#L94)
HostnameError results when the set of authorized names doesn't match the
requested name.


<pre>type HostnameError struct {
<span id="HostnameError.Certificate"></span>    Certificate *<a href="#Certificate">Certificate</a>
<span id="HostnameError.Host"></span>    Host        <a href="/pkg/builtin/#string">string</a>
}
</pre>











### <a id="HostnameError.Error">func</a> (HostnameError) [Error](https://golang.org/src/crypto/x509/verify.go?s=4010:4047#L99)
<pre>func (h <a href="#HostnameError">HostnameError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="InsecureAlgorithmError">type</a> [InsecureAlgorithmError](https://golang.org/src/crypto/x509/x509.go?s=25977:26023#L757)
An InsecureAlgorithmError


<pre>type InsecureAlgorithmError <a href="#SignatureAlgorithm">SignatureAlgorithm</a></pre>











### <a id="InsecureAlgorithmError.Error">func</a> (InsecureAlgorithmError) [Error](https://golang.org/src/crypto/x509/x509.go?s=26025:26071#L759)
<pre>func (e <a href="#InsecureAlgorithmError">InsecureAlgorithmError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="InvalidReason">type</a> [InvalidReason](https://golang.org/src/crypto/x509/verify.go?s=444:466#L14)

<pre>type InvalidReason <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span class="comment">// NotAuthorizedToSign results when a certificate is signed by another</span>
    <span class="comment">// which isn&#39;t marked as a CA certificate.</span>
    <span id="NotAuthorizedToSign">NotAuthorizedToSign</span> <a href="#InvalidReason">InvalidReason</a> = <a href="/pkg/builtin/#iota">iota</a>
    <span class="comment">// Expired results when a certificate has expired, based on the time</span>
    <span class="comment">// given in the VerifyOptions.</span>
    <span id="Expired">Expired</span>
    <span class="comment">// CANotAuthorizedForThisName results when an intermediate or root</span>
    <span class="comment">// certificate has a name constraint which doesn&#39;t permit a DNS or</span>
    <span class="comment">// other name (including IP address) in the leaf certificate.</span>
    <span id="CANotAuthorizedForThisName">CANotAuthorizedForThisName</span>
    <span class="comment">// TooManyIntermediates results when a path length constraint is</span>
    <span class="comment">// violated.</span>
    <span id="TooManyIntermediates">TooManyIntermediates</span>
    <span class="comment">// IncompatibleUsage results when the certificate&#39;s key usage indicates</span>
    <span class="comment">// that it may only be used for a different purpose.</span>
    <span id="IncompatibleUsage">IncompatibleUsage</span>
    <span class="comment">// NameMismatch results when the subject name of a parent certificate</span>
    <span class="comment">// does not match the issuer name in the child.</span>
    <span id="NameMismatch">NameMismatch</span>
    <span class="comment">// NameConstraintsWithoutSANs results when a leaf certificate doesn&#39;t</span>
    <span class="comment">// contain a Subject Alternative Name extension, but a CA certificate</span>
    <span class="comment">// contains name constraints, and the Common Name can be interpreted as</span>
    <span class="comment">// a hostname.</span>
    <span class="comment">//</span>
    <span class="comment">// You can avoid this error by setting the experimental GODEBUG environment</span>
    <span class="comment">// variable to &#34;x509ignoreCN=1&#34;, disabling Common Name matching entirely.</span>
    <span class="comment">// This behavior might become the default in the future.</span>
    <span id="NameConstraintsWithoutSANs">NameConstraintsWithoutSANs</span>
    <span class="comment">// UnconstrainedName results when a CA certificate contains permitted</span>
    <span class="comment">// name constraints, but leaf certificate contains a name of an</span>
    <span class="comment">// unsupported or unconstrained type.</span>
    <span id="UnconstrainedName">UnconstrainedName</span>
    <span class="comment">// TooManyConstraints results when the number of comparison operations</span>
    <span class="comment">// needed to check a certificate exceeds the limit set by</span>
    <span class="comment">// VerifyOptions.MaxConstraintComparisions. This limit exists to</span>
    <span class="comment">// prevent pathological certificates can consuming excessive amounts of</span>
    <span class="comment">// CPU time to verify.</span>
    <span id="TooManyConstraints">TooManyConstraints</span>
    <span class="comment">// CANotAuthorizedForExtKeyUsage results when an intermediate or root</span>
    <span class="comment">// certificate does not permit a requested extended key usage.</span>
    <span id="CANotAuthorizedForExtKeyUsage">CANotAuthorizedForExtKeyUsage</span>
)</pre>









## <a id="KeyUsage">type</a> [KeyUsage](https://golang.org/src/crypto/x509/x509.go?s=17381:17398#L547)
KeyUsage represents the set of actions that are valid for a given key. It's
a bitmap of the KeyUsage* constants.


<pre>type KeyUsage <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="KeyUsageDigitalSignature">KeyUsageDigitalSignature</span> <a href="#KeyUsage">KeyUsage</a> = 1 &lt;&lt; <a href="/pkg/builtin/#iota">iota</a>
    <span id="KeyUsageContentCommitment">KeyUsageContentCommitment</span>
    <span id="KeyUsageKeyEncipherment">KeyUsageKeyEncipherment</span>
    <span id="KeyUsageDataEncipherment">KeyUsageDataEncipherment</span>
    <span id="KeyUsageKeyAgreement">KeyUsageKeyAgreement</span>
    <span id="KeyUsageCertSign">KeyUsageCertSign</span>
    <span id="KeyUsageCRLSign">KeyUsageCRLSign</span>
    <span id="KeyUsageEncipherOnly">KeyUsageEncipherOnly</span>
    <span id="KeyUsageDecipherOnly">KeyUsageDecipherOnly</span>
)</pre>









## <a id="PEMCipher">type</a> [PEMCipher](https://golang.org/src/crypto/x509/pem_decrypt.go?s=472:490#L13)

<pre>type PEMCipher <a href="/pkg/builtin/#int">int</a></pre>


Possible values for the EncryptPEMBlock encryption algorithm.


<pre>const (
    <span id="PEMCipherDES">PEMCipherDES</span> <a href="#PEMCipher">PEMCipher</a>
    <span id="PEMCipher3DES">PEMCipher3DES</span>
    <span id="PEMCipherAES128">PEMCipherAES128</span>
    <span id="PEMCipherAES192">PEMCipherAES192</span>
    <span id="PEMCipherAES256">PEMCipherAES256</span>
)</pre>









## <a id="PublicKeyAlgorithm">type</a> [PublicKeyAlgorithm](https://golang.org/src/crypto/x509/x509.go?s=5883:5910#L212)

<pre>type PublicKeyAlgorithm <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="UnknownPublicKeyAlgorithm">UnknownPublicKeyAlgorithm</span> <a href="#PublicKeyAlgorithm">PublicKeyAlgorithm</a> = <a href="/pkg/builtin/#iota">iota</a>
    <span id="RSA">RSA</span>
    <span id="DSA">DSA</span>
    <span id="ECDSA">ECDSA</span>
    <span id="Ed25519">Ed25519</span>
)</pre>









### <a id="PublicKeyAlgorithm.String">func</a> (PublicKeyAlgorithm) [String](https://golang.org/src/crypto/x509/x509.go?s=6116:6162#L229)
<pre>func (algo <a href="#PublicKeyAlgorithm">PublicKeyAlgorithm</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="SignatureAlgorithm">type</a> [SignatureAlgorithm](https://golang.org/src/crypto/x509/x509.go?s=5183:5210#L172)

<pre>type SignatureAlgorithm <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="UnknownSignatureAlgorithm">UnknownSignatureAlgorithm</span> <a href="#SignatureAlgorithm">SignatureAlgorithm</a> = <a href="/pkg/builtin/#iota">iota</a>
    <span id="MD2WithRSA">MD2WithRSA</span>
    <span id="MD5WithRSA">MD5WithRSA</span>
    <span id="SHA1WithRSA">SHA1WithRSA</span>
    <span id="SHA256WithRSA">SHA256WithRSA</span>
    <span id="SHA384WithRSA">SHA384WithRSA</span>
    <span id="SHA512WithRSA">SHA512WithRSA</span>
    <span id="DSAWithSHA1">DSAWithSHA1</span>
    <span id="DSAWithSHA256">DSAWithSHA256</span>
    <span id="ECDSAWithSHA1">ECDSAWithSHA1</span>
    <span id="ECDSAWithSHA256">ECDSAWithSHA256</span>
    <span id="ECDSAWithSHA384">ECDSAWithSHA384</span>
    <span id="ECDSAWithSHA512">ECDSAWithSHA512</span>
    <span id="SHA256WithRSAPSS">SHA256WithRSAPSS</span>
    <span id="SHA384WithRSAPSS">SHA384WithRSAPSS</span>
    <span id="SHA512WithRSAPSS">SHA512WithRSAPSS</span>
    <span id="PureEd25519">PureEd25519</span>
)</pre>









### <a id="SignatureAlgorithm.String">func</a> (SignatureAlgorithm) [String](https://golang.org/src/crypto/x509/x509.go?s=5688:5734#L203)
<pre>func (algo <a href="#SignatureAlgorithm">SignatureAlgorithm</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="SystemRootsError">type</a> [SystemRootsError](https://golang.org/src/crypto/x509/verify.go?s=6034:6077#L162)
SystemRootsError results when we fail to load the system root certificates.


<pre>type SystemRootsError struct {
<span id="SystemRootsError.Err"></span>    Err <a href="/pkg/builtin/#error">error</a>
}
</pre>











### <a id="SystemRootsError.Error">func</a> (SystemRootsError) [Error](https://golang.org/src/crypto/x509/verify.go?s=6079:6120#L166)
<pre>func (se <a href="#SystemRootsError">SystemRootsError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="UnhandledCriticalExtension">type</a> [UnhandledCriticalExtension](https://golang.org/src/crypto/x509/x509.go?s=33744:33784#L971)

<pre>type UnhandledCriticalExtension struct{}
</pre>











### <a id="UnhandledCriticalExtension.Error">func</a> (UnhandledCriticalExtension) [Error](https://golang.org/src/crypto/x509/x509.go?s=33786:33836#L973)
<pre>func (h <a href="#UnhandledCriticalExtension">UnhandledCriticalExtension</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="UnknownAuthorityError">type</a> [UnknownAuthorityError](https://golang.org/src/crypto/x509/verify.go?s=5131:5435#L135)
UnknownAuthorityError results when the certificate issuer is unknown


<pre>type UnknownAuthorityError struct {
<span id="UnknownAuthorityError.Cert"></span>    Cert *<a href="#Certificate">Certificate</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="UnknownAuthorityError.Error">func</a> (UnknownAuthorityError) [Error](https://golang.org/src/crypto/x509/verify.go?s=5437:5482#L145)
<pre>func (e <a href="#UnknownAuthorityError">UnknownAuthorityError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="VerifyOptions">type</a> [VerifyOptions](https://golang.org/src/crypto/x509/verify.go?s=6646:7592#L180)
VerifyOptions contains parameters for Certificate.Verify. It's a structure
because other PKIX verification APIs have ended up needing many options.


<pre>type VerifyOptions struct {
<span id="VerifyOptions.DNSName"></span>    DNSName       <a href="/pkg/builtin/#string">string</a>
<span id="VerifyOptions.Intermediates"></span>    Intermediates *<a href="#CertPool">CertPool</a>
<span id="VerifyOptions.Roots"></span>    Roots         *<a href="#CertPool">CertPool</a> <span class="comment">// if nil, the system roots are used</span>
<span id="VerifyOptions.CurrentTime"></span>    CurrentTime   <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a> <span class="comment">// if zero, the current time is used</span>
    <span class="comment">// KeyUsage specifies which Extended Key Usage values are acceptable. A leaf</span>
    <span class="comment">// certificate is accepted if it contains any of the listed values. An empty</span>
    <span class="comment">// list means ExtKeyUsageServerAuth. To accept any key usage, include</span>
    <span class="comment">// ExtKeyUsageAny.</span>
    <span class="comment">//</span>
    <span class="comment">// Certificate chains are required to nest these extended key usage values.</span>
    <span class="comment">// (This matches the Windows CryptoAPI behavior, but not the spec.)</span>
<span id="VerifyOptions.KeyUsages"></span>    KeyUsages []<a href="#ExtKeyUsage">ExtKeyUsage</a>
<span id="VerifyOptions.MaxConstraintComparisions"></span>    <span class="comment">// MaxConstraintComparisions is the maximum number of comparisons to</span>
    <span class="comment">// perform when checking a given certificate&#39;s name constraints. If</span>
    <span class="comment">// zero, a sensible default is used. This limit prevents pathological</span>
    <span class="comment">// certificates from consuming excessive amounts of CPU time when</span>
    <span class="comment">// validating.</span>
    MaxConstraintComparisions <a href="/pkg/builtin/#int">int</a>
}
</pre>














