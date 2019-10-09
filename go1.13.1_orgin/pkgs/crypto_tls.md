

# tls
`import "crypto/tls"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package tls partially implements TLS 1.2, as specified in RFC 5246,
and TLS 1.3, as specified in RFC 8446.

TLS 1.3 is available on an opt-out basis in Go 1.13. To disable
it, set the GODEBUG environment variable (comma-separated key=value
options) such that it includes "tls13=0".




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func Listen(network, laddr string, config *Config) (net.Listener, error)](#Listen)
* [func NewListener(inner net.Listener, config *Config) net.Listener](#NewListener)
* [type Certificate](#Certificate)
  * [func LoadX509KeyPair(certFile, keyFile string) (Certificate, error)](#LoadX509KeyPair)
  * [func X509KeyPair(certPEMBlock, keyPEMBlock []byte) (Certificate, error)](#X509KeyPair)
* [type CertificateRequestInfo](#CertificateRequestInfo)
* [type ClientAuthType](#ClientAuthType)
* [type ClientHelloInfo](#ClientHelloInfo)
* [type ClientSessionCache](#ClientSessionCache)
  * [func NewLRUClientSessionCache(capacity int) ClientSessionCache](#NewLRUClientSessionCache)
* [type ClientSessionState](#ClientSessionState)
* [type Config](#Config)
  * [func (c *Config) BuildNameToCertificate()](#Config.BuildNameToCertificate)
  * [func (c *Config) Clone() *Config](#Config.Clone)
  * [func (c *Config) SetSessionTicketKeys(keys [][32]byte)](#Config.SetSessionTicketKeys)
* [type Conn](#Conn)
  * [func Client(conn net.Conn, config *Config) *Conn](#Client)
  * [func Dial(network, addr string, config *Config) (*Conn, error)](#Dial)
  * [func DialWithDialer(dialer *net.Dialer, network, addr string, config *Config) (*Conn, error)](#DialWithDialer)
  * [func Server(conn net.Conn, config *Config) *Conn](#Server)
  * [func (c *Conn) Close() error](#Conn.Close)
  * [func (c *Conn) CloseWrite() error](#Conn.CloseWrite)
  * [func (c *Conn) ConnectionState() ConnectionState](#Conn.ConnectionState)
  * [func (c *Conn) Handshake() error](#Conn.Handshake)
  * [func (c *Conn) LocalAddr() net.Addr](#Conn.LocalAddr)
  * [func (c *Conn) OCSPResponse() []byte](#Conn.OCSPResponse)
  * [func (c *Conn) Read(b []byte) (int, error)](#Conn.Read)
  * [func (c *Conn) RemoteAddr() net.Addr](#Conn.RemoteAddr)
  * [func (c *Conn) SetDeadline(t time.Time) error](#Conn.SetDeadline)
  * [func (c *Conn) SetReadDeadline(t time.Time) error](#Conn.SetReadDeadline)
  * [func (c *Conn) SetWriteDeadline(t time.Time) error](#Conn.SetWriteDeadline)
  * [func (c *Conn) VerifyHostname(host string) error](#Conn.VerifyHostname)
  * [func (c *Conn) Write(b []byte) (int, error)](#Conn.Write)
* [type ConnectionState](#ConnectionState)
  * [func (cs *ConnectionState) ExportKeyingMaterial(label string, context []byte, length int) ([]byte, error)](#ConnectionState.ExportKeyingMaterial)
* [type CurveID](#CurveID)
* [type RecordHeaderError](#RecordHeaderError)
  * [func (e RecordHeaderError) Error() string](#RecordHeaderError.Error)
* [type RenegotiationSupport](#RenegotiationSupport)
* [type SignatureScheme](#SignatureScheme)


#### <a id="pkg-examples">Examples</a>
* [Config (KeyLogWriter)](#example_Config_keyLogWriter)
* [Dial](#example_Dial)
* [LoadX509KeyPair](#example_LoadX509KeyPair)
* [X509KeyPair](#example_X509KeyPair)
* [X509KeyPair (HttpServer)](#example_X509KeyPair_httpServer)


#### <a id="pkg-files">Package files</a>
[alert.go](https://golang.org/src/crypto/tls/alert.go) [auth.go](https://golang.org/src/crypto/tls/auth.go) [cipher_suites.go](https://golang.org/src/crypto/tls/cipher_suites.go) [common.go](https://golang.org/src/crypto/tls/common.go) [conn.go](https://golang.org/src/crypto/tls/conn.go) [handshake_client.go](https://golang.org/src/crypto/tls/handshake_client.go) [handshake_client_tls13.go](https://golang.org/src/crypto/tls/handshake_client_tls13.go) [handshake_messages.go](https://golang.org/src/crypto/tls/handshake_messages.go) [handshake_server.go](https://golang.org/src/crypto/tls/handshake_server.go) [handshake_server_tls13.go](https://golang.org/src/crypto/tls/handshake_server_tls13.go) [key_agreement.go](https://golang.org/src/crypto/tls/key_agreement.go) [key_schedule.go](https://golang.org/src/crypto/tls/key_schedule.go) [prf.go](https://golang.org/src/crypto/tls/prf.go) [ticket.go](https://golang.org/src/crypto/tls/ticket.go) [tls.go](https://golang.org/src/crypto/tls/tls.go) 


## <a id="pkg-constants">Constants</a>
A list of cipher suite IDs that are, or have been, implemented by this
package.

Taken from <a href="https://www.iana.org/assignments/tls-parameters/tls-parameters.xml">https://www.iana.org/assignments/tls-parameters/tls-parameters.xml</a>


<pre>const (
    <span class="comment">// TLS 1.0 - 1.2 cipher suites.</span>
    <span id="TLS_RSA_WITH_RC4_128_SHA">TLS_RSA_WITH_RC4_128_SHA</span>                <a href="/pkg/builtin/#uint16">uint16</a> = 0x0005
    <span id="TLS_RSA_WITH_3DES_EDE_CBC_SHA">TLS_RSA_WITH_3DES_EDE_CBC_SHA</span>           <a href="/pkg/builtin/#uint16">uint16</a> = 0x000a
    <span id="TLS_RSA_WITH_AES_128_CBC_SHA">TLS_RSA_WITH_AES_128_CBC_SHA</span>            <a href="/pkg/builtin/#uint16">uint16</a> = 0x002f
    <span id="TLS_RSA_WITH_AES_256_CBC_SHA">TLS_RSA_WITH_AES_256_CBC_SHA</span>            <a href="/pkg/builtin/#uint16">uint16</a> = 0x0035
    <span id="TLS_RSA_WITH_AES_128_CBC_SHA256">TLS_RSA_WITH_AES_128_CBC_SHA256</span>         <a href="/pkg/builtin/#uint16">uint16</a> = 0x003c
    <span id="TLS_RSA_WITH_AES_128_GCM_SHA256">TLS_RSA_WITH_AES_128_GCM_SHA256</span>         <a href="/pkg/builtin/#uint16">uint16</a> = 0x009c
    <span id="TLS_RSA_WITH_AES_256_GCM_SHA384">TLS_RSA_WITH_AES_256_GCM_SHA384</span>         <a href="/pkg/builtin/#uint16">uint16</a> = 0x009d
    <span id="TLS_ECDHE_ECDSA_WITH_RC4_128_SHA">TLS_ECDHE_ECDSA_WITH_RC4_128_SHA</span>        <a href="/pkg/builtin/#uint16">uint16</a> = 0xc007
    <span id="TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA">TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA</span>    <a href="/pkg/builtin/#uint16">uint16</a> = 0xc009
    <span id="TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA">TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA</span>    <a href="/pkg/builtin/#uint16">uint16</a> = 0xc00a
    <span id="TLS_ECDHE_RSA_WITH_RC4_128_SHA">TLS_ECDHE_RSA_WITH_RC4_128_SHA</span>          <a href="/pkg/builtin/#uint16">uint16</a> = 0xc011
    <span id="TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA">TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA</span>     <a href="/pkg/builtin/#uint16">uint16</a> = 0xc012
    <span id="TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA">TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA</span>      <a href="/pkg/builtin/#uint16">uint16</a> = 0xc013
    <span id="TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA">TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA</span>      <a href="/pkg/builtin/#uint16">uint16</a> = 0xc014
    <span id="TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256">TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256</span> <a href="/pkg/builtin/#uint16">uint16</a> = 0xc023
    <span id="TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256">TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256</span>   <a href="/pkg/builtin/#uint16">uint16</a> = 0xc027
    <span id="TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256">TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256</span>   <a href="/pkg/builtin/#uint16">uint16</a> = 0xc02f
    <span id="TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256">TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256</span> <a href="/pkg/builtin/#uint16">uint16</a> = 0xc02b
    <span id="TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384">TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384</span>   <a href="/pkg/builtin/#uint16">uint16</a> = 0xc030
    <span id="TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384">TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384</span> <a href="/pkg/builtin/#uint16">uint16</a> = 0xc02c
    <span id="TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305">TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305</span>    <a href="/pkg/builtin/#uint16">uint16</a> = 0xcca8
    <span id="TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305">TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305</span>  <a href="/pkg/builtin/#uint16">uint16</a> = 0xcca9

    <span class="comment">// TLS 1.3 cipher suites.</span>
    <span id="TLS_AES_128_GCM_SHA256">TLS_AES_128_GCM_SHA256</span>       <a href="/pkg/builtin/#uint16">uint16</a> = 0x1301
    <span id="TLS_AES_256_GCM_SHA384">TLS_AES_256_GCM_SHA384</span>       <a href="/pkg/builtin/#uint16">uint16</a> = 0x1302
    <span id="TLS_CHACHA20_POLY1305_SHA256">TLS_CHACHA20_POLY1305_SHA256</span> <a href="/pkg/builtin/#uint16">uint16</a> = 0x1303

    <span class="comment">// TLS_FALLBACK_SCSV isn&#39;t a standard cipher suite but an indicator</span>
    <span class="comment">// that the client is doing version fallback. See RFC 7507.</span>
    <span id="TLS_FALLBACK_SCSV">TLS_FALLBACK_SCSV</span> <a href="/pkg/builtin/#uint16">uint16</a> = 0x5600
)</pre>
<pre>const (
    <span id="VersionTLS10">VersionTLS10</span> = 0x0301
    <span id="VersionTLS11">VersionTLS11</span> = 0x0302
    <span id="VersionTLS12">VersionTLS12</span> = 0x0303
    <span id="VersionTLS13">VersionTLS13</span> = 0x0304

    <span class="comment">// Deprecated: SSLv3 is cryptographically broken, and will be</span>
    <span class="comment">// removed in Go 1.14. See golang.org/issue/32716.</span>
    <span id="VersionSSL30">VersionSSL30</span> = 0x0300
)</pre>



## <a id="Listen">func</a> [Listen](https://golang.org/src/crypto/tls/tls.go?s=2529:2601#L71)
<pre>func Listen(network, laddr <a href="/pkg/builtin/#string">string</a>, config *<a href="#Config">Config</a>) (<a href="/pkg/net/">net</a>.<a href="/pkg/net/#Listener">Listener</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Listen creates a TLS listener accepting connections on the
given network address using net.Listen.
The configuration config must be non-nil and must include
at least one certificate or else set GetCertificate.



## <a id="NewListener">func</a> [NewListener](https://golang.org/src/crypto/tls/tls.go?s=2167:2232#L60)
<pre>func NewListener(inner <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Listener">Listener</a>, config *<a href="#Config">Config</a>) <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Listener">Listener</a></pre>
NewListener creates a Listener which accepts connections from an inner
Listener and wraps each connection with Server.
The configuration config must be non-nil and must include
at least one certificate or else set GetCertificate.





## <a id="Certificate">type</a> [Certificate](https://golang.org/src/crypto/tls/common.go?s=36047:36965#L988)
A Certificate is a chain of one or more certificates, leaf first.


<pre>type Certificate struct {
<span id="Certificate.Certificate"></span>    Certificate [][]<a href="/pkg/builtin/#byte">byte</a>
<span id="Certificate.PrivateKey"></span>    <span class="comment">// PrivateKey contains the private key corresponding to the public key in</span>
<span id="Certificate.Leaf"></span>    <span class="comment">// Leaf. This must implement crypto.Signer with an RSA, ECDSA or Ed25519 PublicKey.</span>
    <span class="comment">// For a server up to TLS 1.2, it can also implement crypto.Decrypter with</span>
    <span class="comment">// an RSA PublicKey.</span>
    PrivateKey <a href="/pkg/crypto/">crypto</a>.<a href="/pkg/crypto/#PrivateKey">PrivateKey</a>
<span id="Certificate.OCSPStaple"></span>    <span class="comment">// OCSPStaple contains an optional OCSP response which will be served</span>
    <span class="comment">// to clients that request it.</span>
    OCSPStaple []<a href="/pkg/builtin/#byte">byte</a>
<span id="Certificate.SignedCertificateTimestamps"></span>    <span class="comment">// SignedCertificateTimestamps contains an optional list of Signed</span>
    <span class="comment">// Certificate Timestamps which will be served to clients that request it.</span>
    SignedCertificateTimestamps [][]<a href="/pkg/builtin/#byte">byte</a>
    <span class="comment">// Leaf is the parsed form of the leaf certificate, which may be</span>
    <span class="comment">// initialized using x509.ParseCertificate to reduce per-handshake</span>
    <span class="comment">// processing for TLS clients doing client authentication. If nil, the</span>
    <span class="comment">// leaf certificate will be parsed as needed.</span>
    Leaf *<a href="/pkg/crypto/x509/">x509</a>.<a href="/pkg/crypto/x509/#Certificate">Certificate</a>
}
</pre>









### <a id="LoadX509KeyPair">func</a> [LoadX509KeyPair](https://golang.org/src/crypto/tls/tls.go?s=5662:5729#L175)
<pre>func LoadX509KeyPair(certFile, keyFile <a href="/pkg/builtin/#string">string</a>) (<a href="#Certificate">Certificate</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
LoadX509KeyPair reads and parses a public/private key pair from a pair
of files. The files must contain PEM encoded data. The certificate file
may contain intermediate certificates following the leaf certificate to
form a certificate chain. On successful return, Certificate.Leaf will
be nil because the parsed form of the certificate is not retained.


<a id="example_LoadX509KeyPair">Example</a>


### <a id="X509KeyPair">func</a> [X509KeyPair](https://golang.org/src/crypto/tls/tls.go?s=6170:6241#L190)
<pre>func X509KeyPair(certPEMBlock, keyPEMBlock []<a href="/pkg/builtin/#byte">byte</a>) (<a href="#Certificate">Certificate</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
X509KeyPair parses a public/private key pair from a pair of
PEM encoded data. On successful return, Certificate.Leaf will be nil because
the parsed form of the certificate is not retained.


<a id="example_X509KeyPair">Example</a>
<a id="example_X509KeyPair_httpServer">Example (HttpServer)</a>




## <a id="CertificateRequestInfo">type</a> [CertificateRequestInfo](https://golang.org/src/crypto/tls/common.go?s=14463:14919#L383)
CertificateRequestInfo contains information from a server's
CertificateRequest message, which is used to demand a certificate and proof
of control from a client.


<pre>type CertificateRequestInfo struct {
<span id="CertificateRequestInfo.AcceptableCAs"></span>    <span class="comment">// AcceptableCAs contains zero or more, DER-encoded, X.501</span>
    <span class="comment">// Distinguished Names. These are the names of root or intermediate CAs</span>
    <span class="comment">// that the server wishes the returned certificate to be signed by. An</span>
    <span class="comment">// empty slice indicates that the server has no preference.</span>
    AcceptableCAs [][]<a href="/pkg/builtin/#byte">byte</a>

<span id="CertificateRequestInfo.SignatureSchemes"></span>    <span class="comment">// SignatureSchemes lists the signature schemes that the server is</span>
    <span class="comment">// willing to verify.</span>
    SignatureSchemes []<a href="#SignatureScheme">SignatureScheme</a>
}
</pre>











## <a id="ClientAuthType">type</a> [ClientAuthType](https://golang.org/src/crypto/tls/common.go?s=8937:8960#L249)
ClientAuthType declares the policy the server will follow for
TLS Client Authentication.


<pre>type ClientAuthType <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="NoClientCert">NoClientCert</span> <a href="#ClientAuthType">ClientAuthType</a> = <a href="/pkg/builtin/#iota">iota</a>
    <span id="RequestClientCert">RequestClientCert</span>
    <span id="RequireAnyClientCert">RequireAnyClientCert</span>
    <span id="VerifyClientCertIfGiven">VerifyClientCertIfGiven</span>
    <span id="RequireAndVerifyClientCert">RequireAndVerifyClientCert</span>
)</pre>









## <a id="ClientHelloInfo">type</a> [ClientHelloInfo](https://golang.org/src/crypto/tls/common.go?s=12391:14290#L335)
ClientHelloInfo contains information from a ClientHello message in order to
guide certificate selection in the GetCertificate callback.


<pre>type ClientHelloInfo struct {
<span id="ClientHelloInfo.CipherSuites"></span>    <span class="comment">// CipherSuites lists the CipherSuites supported by the client (e.g.</span>
    <span class="comment">// TLS_AES_128_GCM_SHA256, TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256).</span>
    CipherSuites []<a href="/pkg/builtin/#uint16">uint16</a>

<span id="ClientHelloInfo.ServerName"></span>    <span class="comment">// ServerName indicates the name of the server requested by the client</span>
    <span class="comment">// in order to support virtual hosting. ServerName is only set if the</span>
    <span class="comment">// client is using SNI (see RFC 4366, Section 3.1).</span>
    ServerName <a href="/pkg/builtin/#string">string</a>

<span id="ClientHelloInfo.SupportedCurves"></span>    <span class="comment">// SupportedCurves lists the elliptic curves supported by the client.</span>
    <span class="comment">// SupportedCurves is set only if the Supported Elliptic Curves</span>
    <span class="comment">// Extension is being used (see RFC 4492, Section 5.1.1).</span>
    SupportedCurves []<a href="#CurveID">CurveID</a>

<span id="ClientHelloInfo.SupportedPoints"></span>    <span class="comment">// SupportedPoints lists the point formats supported by the client.</span>
    <span class="comment">// SupportedPoints is set only if the Supported Point Formats Extension</span>
    <span class="comment">// is being used (see RFC 4492, Section 5.1.2).</span>
    SupportedPoints []<a href="/pkg/builtin/#uint8">uint8</a>

<span id="ClientHelloInfo.SignatureSchemes"></span>    <span class="comment">// SignatureSchemes lists the signature and hash schemes that the client</span>
    <span class="comment">// is willing to verify. SignatureSchemes is set only if the Signature</span>
    <span class="comment">// Algorithms Extension is being used (see RFC 5246, Section 7.4.1.4.1).</span>
    SignatureSchemes []<a href="#SignatureScheme">SignatureScheme</a>

<span id="ClientHelloInfo.SupportedProtos"></span>    <span class="comment">// SupportedProtos lists the application protocols supported by the client.</span>
    <span class="comment">// SupportedProtos is set only if the Application-Layer Protocol</span>
    <span class="comment">// Negotiation Extension is being used (see RFC 7301, Section 3.1).</span>
    <span class="comment">//</span>
    <span class="comment">// Servers can select a protocol by setting Config.NextProtos in a</span>
    <span class="comment">// GetConfigForClient return value.</span>
    SupportedProtos []<a href="/pkg/builtin/#string">string</a>

<span id="ClientHelloInfo.SupportedVersions"></span>    <span class="comment">// SupportedVersions lists the TLS versions supported by the client.</span>
    <span class="comment">// For TLS versions less than 1.3, this is extrapolated from the max</span>
    <span class="comment">// version advertised by the client, so values other than the greatest</span>
    <span class="comment">// might be rejected if used.</span>
    SupportedVersions []<a href="/pkg/builtin/#uint16">uint16</a>

<span id="ClientHelloInfo.Conn"></span>    <span class="comment">// Conn is the underlying net.Conn for the connection. Do not read</span>
    <span class="comment">// from, or write to, this connection; that will cause the TLS</span>
    <span class="comment">// connection to fail.</span>
    Conn <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Conn">Conn</a>
}
</pre>











## <a id="ClientSessionCache">type</a> [ClientSessionCache](https://golang.org/src/crypto/tls/common.go?s=10819:11352#L293)
ClientSessionCache is a cache of ClientSessionState objects that can be used
by a client to resume a TLS session with a given server. ClientSessionCache
implementations should expect to be called concurrently from different
goroutines. Up to TLS 1.2, only ticket-based resumption is supported, not
SessionID-based resumption. In TLS 1.3 they were merged into PSK modes, which
are supported via this interface.


<pre>type ClientSessionCache interface {
    <span class="comment">// Get searches for a ClientSessionState associated with the given key.</span>
    <span class="comment">// On return, ok is true if one was found.</span>
    Get(sessionKey <a href="/pkg/builtin/#string">string</a>) (session *<a href="#ClientSessionState">ClientSessionState</a>, ok <a href="/pkg/builtin/#bool">bool</a>)

    <span class="comment">// Put adds the ClientSessionState to the cache with the given key. It might</span>
    <span class="comment">// get called multiple times in a connection if a TLS 1.3 server provides</span>
    <span class="comment">// more than one session ticket. If called with a nil *ClientSessionState,</span>
    <span class="comment">// it should remove the cache entry.</span>
    Put(sessionKey <a href="/pkg/builtin/#string">string</a>, cs *<a href="#ClientSessionState">ClientSessionState</a>)
}</pre>









### <a id="NewLRUClientSessionCache">func</a> [NewLRUClientSessionCache](https://golang.org/src/crypto/tls/common.go?s=37517:37579#L1031)
<pre>func NewLRUClientSessionCache(capacity <a href="/pkg/builtin/#int">int</a>) <a href="#ClientSessionCache">ClientSessionCache</a></pre>
NewLRUClientSessionCache returns a ClientSessionCache with the given
capacity that uses an LRU strategy. If capacity is < 1, a default capacity
is used instead.






## <a id="ClientSessionState">type</a> [ClientSessionState](https://golang.org/src/crypto/tls/common.go?s=9457:10389#L272)
ClientSessionState contains the state needed by clients to resume TLS
sessions.


<pre>type ClientSessionState struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











## <a id="Config">type</a> [Config](https://golang.org/src/crypto/tls/common.go?s=16307:25155#L428)
A Config structure is used to configure a TLS client or server.
After one has been passed to a TLS function it must not be
modified. A Config may be reused; the tls package will also not
modify it.


<pre>type Config struct {
<span id="Config.Rand"></span>    <span class="comment">// Rand provides the source of entropy for nonces and RSA blinding.</span>
    <span class="comment">// If Rand is nil, TLS uses the cryptographic random reader in package</span>
    <span class="comment">// crypto/rand.</span>
    <span class="comment">// The Reader must be safe for use by multiple goroutines.</span>
    Rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>

<span id="Config.Time"></span>    <span class="comment">// Time returns the current time as the number of seconds since the epoch.</span>
    <span class="comment">// If Time is nil, TLS uses time.Now.</span>
    Time func() <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>

<span id="Config.Certificates"></span>    <span class="comment">// Certificates contains one or more certificate chains to present to</span>
    <span class="comment">// the other side of the connection. Server configurations must include</span>
    <span class="comment">// at least one certificate or else set GetCertificate. Clients doing</span>
    <span class="comment">// client-authentication may set either Certificates or</span>
<span id="Config.GetClientCertificate"></span>    <span class="comment">// GetClientCertificate.</span>
    Certificates []<a href="#Certificate">Certificate</a>

<span id="Config.NameToCertificate"></span>    <span class="comment">// NameToCertificate maps from a certificate name to an element of</span>
    <span class="comment">// Certificates. Note that a certificate name can be of the form</span>
    <span class="comment">// &#39;*.example.com&#39; and so doesn&#39;t have to be a domain name as such.</span>
    <span class="comment">// See Config.BuildNameToCertificate</span>
    <span class="comment">// The nil value causes the first element of Certificates to be used</span>
    <span class="comment">// for all connections.</span>
    NameToCertificate map[<a href="/pkg/builtin/#string">string</a>]*<a href="#Certificate">Certificate</a>

<span id="Config.GetCertificate"></span>    <span class="comment">// GetCertificate returns a Certificate based on the given</span>
    <span class="comment">// ClientHelloInfo. It will only be called if the client supplies SNI</span>
    <span class="comment">// information or if Certificates is empty.</span>
    <span class="comment">//</span>
    <span class="comment">// If GetCertificate is nil or returns nil, then the certificate is</span>
    <span class="comment">// retrieved from NameToCertificate. If NameToCertificate is nil, the</span>
    <span class="comment">// first element of Certificates will be used.</span>
    GetCertificate func(*<a href="#ClientHelloInfo">ClientHelloInfo</a>) (*<a href="#Certificate">Certificate</a>, <a href="/pkg/builtin/#error">error</a>)

    <span class="comment">// GetClientCertificate, if not nil, is called when a server requests a</span>
    <span class="comment">// certificate from a client. If set, the contents of Certificates will</span>
    <span class="comment">// be ignored.</span>
    <span class="comment">//</span>
    <span class="comment">// If GetClientCertificate returns an error, the handshake will be</span>
    <span class="comment">// aborted and that error will be returned. Otherwise</span>
    <span class="comment">// GetClientCertificate must return a non-nil Certificate. If</span>
    <span class="comment">// Certificate.Certificate is empty then no certificate will be sent to</span>
    <span class="comment">// the server. If this is unacceptable to the server then it may abort</span>
    <span class="comment">// the handshake.</span>
    <span class="comment">//</span>
    <span class="comment">// GetClientCertificate may be called multiple times for the same</span>
    <span class="comment">// connection if renegotiation occurs or if TLS 1.3 is in use.</span>
    GetClientCertificate func(*<a href="#CertificateRequestInfo">CertificateRequestInfo</a>) (*<a href="#Certificate">Certificate</a>, <a href="/pkg/builtin/#error">error</a>)

<span id="Config.GetConfigForClient"></span>    <span class="comment">// GetConfigForClient, if not nil, is called after a ClientHello is</span>
    <span class="comment">// received from a client. It may return a non-nil Config in order to</span>
    <span class="comment">// change the Config that will be used to handle this connection. If</span>
    <span class="comment">// the returned Config is nil, the original Config will be used. The</span>
    <span class="comment">// Config returned by this callback may not be subsequently modified.</span>
    <span class="comment">//</span>
    <span class="comment">// If GetConfigForClient is nil, the Config passed to Server() will be</span>
    <span class="comment">// used for all connections.</span>
    <span class="comment">//</span>
    <span class="comment">// Uniquely for the fields in the returned Config, session ticket keys</span>
    <span class="comment">// will be duplicated from the original Config if not set.</span>
    <span class="comment">// Specifically, if SetSessionTicketKeys was called on the original</span>
    <span class="comment">// config but not on the returned config then the ticket keys from the</span>
    <span class="comment">// original config will be copied into the new config before use.</span>
    <span class="comment">// Otherwise, if SessionTicketKey was set in the original config but</span>
    <span class="comment">// not in the returned config then it will be copied into the returned</span>
    <span class="comment">// config before use. If neither of those cases applies then the key</span>
    <span class="comment">// material from the returned config will be used for session tickets.</span>
    GetConfigForClient func(*<a href="#ClientHelloInfo">ClientHelloInfo</a>) (*<a href="#Config">Config</a>, <a href="/pkg/builtin/#error">error</a>)

<span id="Config.VerifyPeerCertificate"></span>    <span class="comment">// VerifyPeerCertificate, if not nil, is called after normal</span>
    <span class="comment">// certificate verification by either a TLS client or server. It</span>
    <span class="comment">// receives the raw ASN.1 certificates provided by the peer and also</span>
    <span class="comment">// any verified chains that normal processing found. If it returns a</span>
    <span class="comment">// non-nil error, the handshake is aborted and that error results.</span>
    <span class="comment">//</span>
    <span class="comment">// If normal verification fails then the handshake will abort before</span>
    <span class="comment">// considering this callback. If normal verification is disabled by</span>
    <span class="comment">// setting InsecureSkipVerify, or (for a server) when ClientAuth is</span>
    <span class="comment">// RequestClientCert or RequireAnyClientCert, then this callback will</span>
    <span class="comment">// be considered but the verifiedChains argument will always be nil.</span>
    VerifyPeerCertificate func(rawCerts [][]<a href="/pkg/builtin/#byte">byte</a>, verifiedChains [][]*<a href="/pkg/crypto/x509/">x509</a>.<a href="/pkg/crypto/x509/#Certificate">Certificate</a>) <a href="/pkg/builtin/#error">error</a>

<span id="Config.RootCAs"></span>    <span class="comment">// RootCAs defines the set of root certificate authorities</span>
    <span class="comment">// that clients use when verifying server certificates.</span>
    <span class="comment">// If RootCAs is nil, TLS uses the host&#39;s root CA set.</span>
    RootCAs *<a href="/pkg/crypto/x509/">x509</a>.<a href="/pkg/crypto/x509/#CertPool">CertPool</a>

<span id="Config.NextProtos"></span>    <span class="comment">// NextProtos is a list of supported application level protocols, in</span>
    <span class="comment">// order of preference.</span>
    NextProtos []<a href="/pkg/builtin/#string">string</a>

<span id="Config.ServerName"></span>    <span class="comment">// ServerName is used to verify the hostname on the returned</span>
    <span class="comment">// certificates unless InsecureSkipVerify is given. It is also included</span>
    <span class="comment">// in the client&#39;s handshake to support virtual hosting unless it is</span>
    <span class="comment">// an IP address.</span>
    ServerName <a href="/pkg/builtin/#string">string</a>

<span id="Config.ClientAuth"></span>    <span class="comment">// ClientAuth determines the server&#39;s policy for</span>
    <span class="comment">// TLS Client Authentication. The default is NoClientCert.</span>
    ClientAuth <a href="#ClientAuthType">ClientAuthType</a>

<span id="Config.ClientCAs"></span>    <span class="comment">// ClientCAs defines the set of root certificate authorities</span>
    <span class="comment">// that servers use if required to verify a client certificate</span>
    <span class="comment">// by the policy in ClientAuth.</span>
    ClientCAs *<a href="/pkg/crypto/x509/">x509</a>.<a href="/pkg/crypto/x509/#CertPool">CertPool</a>

<span id="Config.InsecureSkipVerify"></span>    <span class="comment">// InsecureSkipVerify controls whether a client verifies the</span>
    <span class="comment">// server&#39;s certificate chain and host name.</span>
    <span class="comment">// If InsecureSkipVerify is true, TLS accepts any certificate</span>
    <span class="comment">// presented by the server and any host name in that certificate.</span>
    <span class="comment">// In this mode, TLS is susceptible to man-in-the-middle attacks.</span>
    <span class="comment">// This should be used only for testing.</span>
    InsecureSkipVerify <a href="/pkg/builtin/#bool">bool</a>

<span id="Config.CipherSuites"></span>    <span class="comment">// CipherSuites is a list of supported cipher suites for TLS versions up to</span>
    <span class="comment">// TLS 1.2. If CipherSuites is nil, a default list of secure cipher suites</span>
    <span class="comment">// is used, with a preference order based on hardware performance. The</span>
    <span class="comment">// default cipher suites might change over Go versions. Note that TLS 1.3</span>
    <span class="comment">// ciphersuites are not configurable.</span>
    CipherSuites []<a href="/pkg/builtin/#uint16">uint16</a>

<span id="Config.PreferServerCipherSuites"></span>    <span class="comment">// PreferServerCipherSuites controls whether the server selects the</span>
    <span class="comment">// client&#39;s most preferred ciphersuite, or the server&#39;s most preferred</span>
    <span class="comment">// ciphersuite. If true then the server&#39;s preference, as expressed in</span>
    <span class="comment">// the order of elements in CipherSuites, is used.</span>
    PreferServerCipherSuites <a href="/pkg/builtin/#bool">bool</a>

<span id="Config.SessionTicketsDisabled"></span>    <span class="comment">// SessionTicketsDisabled may be set to true to disable session ticket and</span>
    <span class="comment">// PSK (resumption) support. Note that on clients, session ticket support is</span>
    <span class="comment">// also disabled if ClientSessionCache is nil.</span>
    SessionTicketsDisabled <a href="/pkg/builtin/#bool">bool</a>

<span id="Config.SessionTicketKey"></span>    <span class="comment">// SessionTicketKey is used by TLS servers to provide session resumption.</span>
    <span class="comment">// See RFC 5077 and the PSK mode of RFC 8446. If zero, it will be filled</span>
    <span class="comment">// with random data before the first server handshake.</span>
    <span class="comment">//</span>
    <span class="comment">// If multiple servers are terminating connections for the same host</span>
    <span class="comment">// they should all have the same SessionTicketKey. If the</span>
    <span class="comment">// SessionTicketKey leaks, previously recorded and future TLS</span>
    <span class="comment">// connections using that key might be compromised.</span>
    SessionTicketKey [32]<a href="/pkg/builtin/#byte">byte</a>

<span id="Config.ClientSessionCache"></span>    <span class="comment">// ClientSessionCache is a cache of ClientSessionState entries for TLS</span>
    <span class="comment">// session resumption. It is only used by clients.</span>
    ClientSessionCache <a href="#ClientSessionCache">ClientSessionCache</a>

<span id="Config.MinVersion"></span>    <span class="comment">// MinVersion contains the minimum SSL/TLS version that is acceptable.</span>
    <span class="comment">// If zero, then TLS 1.0 is taken as the minimum.</span>
    MinVersion <a href="/pkg/builtin/#uint16">uint16</a>

<span id="Config.MaxVersion"></span>    <span class="comment">// MaxVersion contains the maximum SSL/TLS version that is acceptable.</span>
    <span class="comment">// If zero, then the maximum version supported by this package is used,</span>
    <span class="comment">// which is currently TLS 1.3.</span>
    MaxVersion <a href="/pkg/builtin/#uint16">uint16</a>

<span id="Config.CurvePreferences"></span>    <span class="comment">// CurvePreferences contains the elliptic curves that will be used in</span>
    <span class="comment">// an ECDHE handshake, in preference order. If empty, the default will</span>
    <span class="comment">// be used. The client will use the first preference as the type for</span>
    <span class="comment">// its key share in TLS 1.3. This may change in the future.</span>
    CurvePreferences []<a href="#CurveID">CurveID</a>

<span id="Config.DynamicRecordSizingDisabled"></span>    <span class="comment">// DynamicRecordSizingDisabled disables adaptive sizing of TLS records.</span>
    <span class="comment">// When true, the largest possible TLS record size is always used. When</span>
    <span class="comment">// false, the size of TLS records may be adjusted in an attempt to</span>
    <span class="comment">// improve latency.</span>
    DynamicRecordSizingDisabled <a href="/pkg/builtin/#bool">bool</a>

<span id="Config.Renegotiation"></span>    <span class="comment">// Renegotiation controls what types of renegotiation are supported.</span>
    <span class="comment">// The default, none, is correct for the vast majority of applications.</span>
    Renegotiation <a href="#RenegotiationSupport">RenegotiationSupport</a>

<span id="Config.KeyLogWriter"></span>    <span class="comment">// KeyLogWriter optionally specifies a destination for TLS master secrets</span>
    <span class="comment">// in NSS key log format that can be used to allow external programs</span>
    <span class="comment">// such as Wireshark to decrypt TLS connections.</span>
    <span class="comment">// See https://developer.mozilla.org/en-US/docs/Mozilla/Projects/NSS/Key_Log_Format.</span>
    <span class="comment">// Use of KeyLogWriter compromises security and should only be</span>
    <span class="comment">// used for debugging.</span>
    KeyLogWriter <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Config_keyLogWriter">Example (KeyLogWriter)</a>






### <a id="Config.BuildNameToCertificate">func</a> (\*Config) [BuildNameToCertificate](https://golang.org/src/crypto/tls/common.go?s=34698:34739#L940)
<pre>func (c *<a href="#Config">Config</a>) BuildNameToCertificate()</pre>
BuildNameToCertificate parses c.Certificates and builds c.NameToCertificate
from the CommonName and SubjectAlternateName fields of each of the leaf
certificates.




### <a id="Config.Clone">func</a> (\*Config) [Clone](https://golang.org/src/crypto/tls/common.go?s=26487:26519#L649)
<pre>func (c *<a href="#Config">Config</a>) Clone() *<a href="#Config">Config</a></pre>
Clone returns a shallow clone of c. It is safe to clone a Config that is
being used concurrently by a TLS client or server.




### <a id="Config.SetSessionTicketKeys">func</a> (\*Config) [SetSessionTicketKeys](https://golang.org/src/crypto/tls/common.go?s=29607:29661#L737)
<pre>func (c *<a href="#Config">Config</a>) SetSessionTicketKeys(keys [][32]<a href="/pkg/builtin/#byte">byte</a>)</pre>
SetSessionTicketKeys updates the session ticket keys for a server. The first
key will be used when creating new tickets, while all keys can be used for
decrypting tickets. It is safe to call this function while the server is
running in order to rotate the session ticket keys. The function will panic
if keys is empty.




## <a id="Conn">type</a> [Conn](https://golang.org/src/crypto/tls/conn.go?s=434:3993#L15)
A Conn represents a secured connection.
It implements the net.Conn interface.


<pre>type Conn struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="Client">func</a> [Client](https://golang.org/src/crypto/tls/tls.go?s=1413:1461#L36)
<pre>func Client(conn <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Conn">Conn</a>, config *<a href="#Config">Config</a>) *<a href="#Conn">Conn</a></pre>
Client returns a new TLS client side connection
using conn as the underlying transport.
The config cannot be nil: users must set either ServerName or
InsecureSkipVerify in the config.




### <a id="Dial">func</a> [Dial](https://golang.org/src/crypto/tls/tls.go?s=5164:5226#L166)
<pre>func Dial(network, addr <a href="/pkg/builtin/#string">string</a>, config *<a href="#Config">Config</a>) (*<a href="#Conn">Conn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Dial connects to the given network address using net.Dial
and then initiates a TLS handshake, returning the resulting
TLS connection.
Dial interprets a nil configuration as equivalent to
the zero configuration; see the documentation of Config
for the defaults.


<a id="example_Dial">Example</a>


### <a id="DialWithDialer">func</a> [DialWithDialer](https://golang.org/src/crypto/tls/tls.go?s=3510:3602#L95)
<pre>func DialWithDialer(dialer *<a href="/pkg/net/">net</a>.<a href="/pkg/net/#Dialer">Dialer</a>, network, addr <a href="/pkg/builtin/#string">string</a>, config *<a href="#Config">Config</a>) (*<a href="#Conn">Conn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DialWithDialer connects to the given network address using dialer.Dial and
then initiates a TLS handshake, returning the resulting TLS connection. Any
timeout or deadline given in the dialer apply to connection and TLS
handshake as a whole.

DialWithDialer interprets a nil configuration as equivalent to the zero
configuration; see the documentation of Config for the defaults.




### <a id="Server">func</a> [Server](https://golang.org/src/crypto/tls/tls.go?s=1121:1169#L28)
<pre>func Server(conn <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Conn">Conn</a>, config *<a href="#Config">Config</a>) *<a href="#Conn">Conn</a></pre>
Server returns a new TLS server side connection
using conn as the underlying transport.
The configuration config must be non-nil and must include
at least one certificate or else set GetCertificate.






### <a id="Conn.Close">func</a> (\*Conn) [Close](https://golang.org/src/crypto/tls/conn.go?s=39450:39478#L1275)
<pre>func (c *<a href="#Conn">Conn</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the connection.




### <a id="Conn.CloseWrite">func</a> (\*Conn) [CloseWrite](https://golang.org/src/crypto/tls/conn.go?s=40555:40588#L1314)
<pre>func (c *<a href="#Conn">Conn</a>) CloseWrite() <a href="/pkg/builtin/#error">error</a></pre>
CloseWrite shuts down the writing side of the connection. It should only be
called once the handshake has completed and does not call CloseWrite on the
underlying connection. Most callers should just use Close.




### <a id="Conn.ConnectionState">func</a> (\*Conn) [ConnectionState](https://golang.org/src/crypto/tls/conn.go?s=41859:41907#L1372)
<pre>func (c *<a href="#Conn">Conn</a>) ConnectionState() <a href="#ConnectionState">ConnectionState</a></pre>
ConnectionState returns basic TLS details about the connection.




### <a id="Conn.Handshake">func</a> (\*Conn) [Handshake](https://golang.org/src/crypto/tls/conn.go?s=41102:41134#L1337)
<pre>func (c *<a href="#Conn">Conn</a>) Handshake() <a href="/pkg/builtin/#error">error</a></pre>
Handshake runs the client or server handshake
protocol if it has not yet been run.
Most uses of this package need not call Handshake
explicitly: the first Read or Write will call it automatically.




### <a id="Conn.LocalAddr">func</a> (\*Conn) [LocalAddr](https://golang.org/src/crypto/tls/conn.go?s=4156:4191#L108)
<pre>func (c *<a href="#Conn">Conn</a>) LocalAddr() <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Addr">Addr</a></pre>
LocalAddr returns the local network address.




### <a id="Conn.OCSPResponse">func</a> (\*Conn) [OCSPResponse](https://golang.org/src/crypto/tls/conn.go?s=42931:42967#L1409)
<pre>func (c *<a href="#Conn">Conn</a>) OCSPResponse() []<a href="/pkg/builtin/#byte">byte</a></pre>
OCSPResponse returns the stapled OCSP response from the TLS server, if
any. (Only valid for client connections.)




### <a id="Conn.Read">func</a> (\*Conn) [Read](https://golang.org/src/crypto/tls/conn.go?s=38163:38205#L1231)
<pre>func (c *<a href="#Conn">Conn</a>) Read(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Read can be made to time out and return a net.Error with Timeout() == true
after a fixed time limit; see SetDeadline and SetReadDeadline.




### <a id="Conn.RemoteAddr">func</a> (\*Conn) [RemoteAddr](https://golang.org/src/crypto/tls/conn.go?s=4274:4310#L113)
<pre>func (c *<a href="#Conn">Conn</a>) RemoteAddr() <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Addr">Addr</a></pre>
RemoteAddr returns the remote network address.




### <a id="Conn.SetDeadline">func</a> (\*Conn) [SetDeadline](https://golang.org/src/crypto/tls/conn.go?s=4594:4639#L120)
<pre>func (c *<a href="#Conn">Conn</a>) SetDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetDeadline sets the read and write deadlines associated with the connection.
A zero value for t means Read and Write will not time out.
After a Write has timed out, the TLS state is corrupt and all future writes will return the same error.




### <a id="Conn.SetReadDeadline">func</a> (\*Conn) [SetReadDeadline](https://golang.org/src/crypto/tls/conn.go?s=4799:4848#L126)
<pre>func (c *<a href="#Conn">Conn</a>) SetReadDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetReadDeadline sets the read deadline on the underlying connection.
A zero value for t means Read will not time out.




### <a id="Conn.SetWriteDeadline">func</a> (\*Conn) [SetWriteDeadline](https://golang.org/src/crypto/tls/conn.go?s=5122:5172#L133)
<pre>func (c *<a href="#Conn">Conn</a>) SetWriteDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetWriteDeadline sets the write deadline on the underlying connection.
A zero value for t means Write will not time out.
After a Write has timed out, the TLS state is corrupt and all future writes will return the same error.




### <a id="Conn.VerifyHostname">func</a> (\*Conn) [VerifyHostname](https://golang.org/src/crypto/tls/conn.go?s=43226:43274#L1419)
<pre>func (c *<a href="#Conn">Conn</a>) VerifyHostname(host <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
VerifyHostname checks that the peer certificate chain is valid for
connecting to host. If so, it returns nil; if not, it returns an error
describing the problem.




### <a id="Conn.Write">func</a> (\*Conn) [Write](https://golang.org/src/crypto/tls/conn.go?s=34095:34138#L1071)
<pre>func (c *<a href="#Conn">Conn</a>) Write(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Write writes data to the connection.




## <a id="ConnectionState">type</a> [ConnectionState](https://golang.org/src/crypto/tls/common.go?s=6689:8406#L214)
ConnectionState records basic TLS details about the connection.


<pre>type ConnectionState struct {
<span id="ConnectionState.Version"></span>    Version                     <a href="/pkg/builtin/#uint16">uint16</a>                <span class="comment">// TLS version used by the connection (e.g. VersionTLS12)</span>
<span id="ConnectionState.HandshakeComplete"></span>    HandshakeComplete           <a href="/pkg/builtin/#bool">bool</a>                  <span class="comment">// TLS handshake is complete</span>
<span id="ConnectionState.DidResume"></span>    DidResume                   <a href="/pkg/builtin/#bool">bool</a>                  <span class="comment">// connection resumes a previous TLS connection</span>
<span id="ConnectionState.CipherSuite"></span>    CipherSuite                 <a href="/pkg/builtin/#uint16">uint16</a>                <span class="comment">// cipher suite in use (TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, ...)</span>
<span id="ConnectionState.NegotiatedProtocol"></span>    NegotiatedProtocol          <a href="/pkg/builtin/#string">string</a>                <span class="comment">// negotiated next protocol (not guaranteed to be from Config.NextProtos)</span>
<span id="ConnectionState.NegotiatedProtocolIsMutual"></span>    NegotiatedProtocolIsMutual  <a href="/pkg/builtin/#bool">bool</a>                  <span class="comment">// negotiated protocol was advertised by server (client side only)</span>
<span id="ConnectionState.ServerName"></span>    ServerName                  <a href="/pkg/builtin/#string">string</a>                <span class="comment">// server name requested by client, if any (server side only)</span>
<span id="ConnectionState.PeerCertificates"></span>    PeerCertificates            []*<a href="/pkg/crypto/x509/">x509</a>.<a href="/pkg/crypto/x509/#Certificate">Certificate</a>   <span class="comment">// certificate chain presented by remote peer</span>
<span id="ConnectionState.VerifiedChains"></span>    VerifiedChains              [][]*<a href="/pkg/crypto/x509/">x509</a>.<a href="/pkg/crypto/x509/#Certificate">Certificate</a> <span class="comment">// verified chains built from PeerCertificates</span>
<span id="ConnectionState.SignedCertificateTimestamps"></span>    SignedCertificateTimestamps [][]<a href="/pkg/builtin/#byte">byte</a>              <span class="comment">// SCTs from the peer, if any</span>
<span id="ConnectionState.OCSPResponse"></span>    OCSPResponse                []<a href="/pkg/builtin/#byte">byte</a>                <span class="comment">// stapled OCSP response from peer, if any</span>

<span id="ConnectionState.TLSUnique"></span>    <span class="comment">// TLSUnique contains the &#34;tls-unique&#34; channel binding value (see RFC</span>
    <span class="comment">// 5929, section 3). For resumed sessions this value will be nil</span>
    <span class="comment">// because resumption does not include enough context (see</span>
    <span class="comment">// https://mitls.org/pages/attacks/3SHAKE#channelbindings). This will</span>
    <span class="comment">// change in future versions of Go once the TLS master-secret fix has</span>
    <span class="comment">// been standardized and implemented. It is not defined in TLS 1.3.</span>
    TLSUnique []<a href="/pkg/builtin/#byte">byte</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="ConnectionState.ExportKeyingMaterial">func</a> (\*ConnectionState) [ExportKeyingMaterial](https://golang.org/src/crypto/tls/common.go?s=8692:8797#L243)
<pre>func (cs *<a href="#ConnectionState">ConnectionState</a>) ExportKeyingMaterial(label <a href="/pkg/builtin/#string">string</a>, context []<a href="/pkg/builtin/#byte">byte</a>, length <a href="/pkg/builtin/#int">int</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ExportKeyingMaterial returns length bytes of exported key material in a new
slice as defined in RFC 5705. If context is nil, it is not used as part of
the seed. If the connection was set to allow renegotiation via
Config.Renegotiation, this function will return an error.




## <a id="CurveID">type</a> [CurveID](https://golang.org/src/crypto/tls/common.go?s=3396:3415#L103)
CurveID is the type of a TLS identifier for an elliptic curve. See
<a href="https://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-8">https://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-8</a>.

In TLS 1.3, this type is called NamedGroup, but at this time this library
only supports Elliptic Curve based groups. See RFC 8446, Section 4.2.7.


<pre>type CurveID <a href="/pkg/builtin/#uint16">uint16</a></pre>



<pre>const (
    <span id="CurveP256">CurveP256</span> <a href="#CurveID">CurveID</a> = 23
    <span id="CurveP384">CurveP384</span> <a href="#CurveID">CurveID</a> = 24
    <span id="CurveP521">CurveP521</span> <a href="#CurveID">CurveID</a> = 25
    <span id="X25519">X25519</span>    <a href="#CurveID">CurveID</a> = 29
)</pre>









## <a id="RecordHeaderError">type</a> [RecordHeaderError](https://golang.org/src/crypto/tls/conn.go?s=17421:17892#L544)
RecordHeaderError is returned when a TLS record header is invalid.


<pre>type RecordHeaderError struct {
<span id="RecordHeaderError.Msg"></span>    <span class="comment">// Msg contains a human readable string that describes the error.</span>
    Msg <a href="/pkg/builtin/#string">string</a>
<span id="RecordHeaderError.RecordHeader"></span>    <span class="comment">// RecordHeader contains the five bytes of TLS record header that</span>
    <span class="comment">// triggered the error.</span>
    RecordHeader [5]<a href="/pkg/builtin/#byte">byte</a>
<span id="RecordHeaderError.Conn"></span>    <span class="comment">// Conn provides the underlying net.Conn in the case that a client</span>
    <span class="comment">// sent an initial handshake that didn&#39;t look like TLS.</span>
    <span class="comment">// It is nil if there&#39;s already been a handshake or a TLS alert has</span>
    <span class="comment">// been written to the connection.</span>
    Conn <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Conn">Conn</a>
}
</pre>











### <a id="RecordHeaderError.Error">func</a> (RecordHeaderError) [Error](https://golang.org/src/crypto/tls/conn.go?s=17894:17935#L557)
<pre>func (e <a href="#RecordHeaderError">RecordHeaderError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="RenegotiationSupport">type</a> [RenegotiationSupport](https://golang.org/src/crypto/tls/common.go?s=15715:15744#L409)
RenegotiationSupport enumerates the different levels of support for TLS
renegotiation. TLS renegotiation is the act of performing subsequent
handshakes on a connection after the first. This significantly complicates
the state machine and has been the source of numerous, subtle security
issues. Initiating a renegotiation is not supported, but support for
accepting renegotiation requests may be enabled.

Even when enabled, the server may not change its identity between handshakes
(i.e. the leaf certificate must be the same). Additionally, concurrent
handshake and application data flow is not permitted so renegotiation can
only be used with protocols that synchronise with the renegotiation, such as
HTTPS.

Renegotiation is not defined in TLS 1.3.


<pre>type RenegotiationSupport <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span class="comment">// RenegotiateNever disables renegotiation.</span>
    <span id="RenegotiateNever">RenegotiateNever</span> <a href="#RenegotiationSupport">RenegotiationSupport</a> = <a href="/pkg/builtin/#iota">iota</a>

    <span class="comment">// RenegotiateOnceAsClient allows a remote server to request</span>
    <span class="comment">// renegotiation once per connection.</span>
    <span id="RenegotiateOnceAsClient">RenegotiateOnceAsClient</span>

    <span class="comment">// RenegotiateFreelyAsClient allows a remote server to repeatedly</span>
    <span class="comment">// request renegotiation.</span>
    <span id="RenegotiateFreelyAsClient">RenegotiateFreelyAsClient</span>
)</pre>









## <a id="SignatureScheme">type</a> [SignatureScheme](https://golang.org/src/crypto/tls/common.go?s=11456:11483#L307)
SignatureScheme identifies a signature algorithm supported by TLS. See
RFC 8446, Section 4.2.3.


<pre>type SignatureScheme <a href="/pkg/builtin/#uint16">uint16</a></pre>



<pre>const (
    <span class="comment">// RSASSA-PKCS1-v1_5 algorithms.</span>
    <span id="PKCS1WithSHA256">PKCS1WithSHA256</span> <a href="#SignatureScheme">SignatureScheme</a> = 0x0401
    <span id="PKCS1WithSHA384">PKCS1WithSHA384</span> <a href="#SignatureScheme">SignatureScheme</a> = 0x0501
    <span id="PKCS1WithSHA512">PKCS1WithSHA512</span> <a href="#SignatureScheme">SignatureScheme</a> = 0x0601

    <span class="comment">// RSASSA-PSS algorithms with public key OID rsaEncryption.</span>
    <span id="PSSWithSHA256">PSSWithSHA256</span> <a href="#SignatureScheme">SignatureScheme</a> = 0x0804
    <span id="PSSWithSHA384">PSSWithSHA384</span> <a href="#SignatureScheme">SignatureScheme</a> = 0x0805
    <span id="PSSWithSHA512">PSSWithSHA512</span> <a href="#SignatureScheme">SignatureScheme</a> = 0x0806

    <span class="comment">// ECDSA algorithms. Only constrained to a specific curve in TLS 1.3.</span>
    <span id="ECDSAWithP256AndSHA256">ECDSAWithP256AndSHA256</span> <a href="#SignatureScheme">SignatureScheme</a> = 0x0403
    <span id="ECDSAWithP384AndSHA384">ECDSAWithP384AndSHA384</span> <a href="#SignatureScheme">SignatureScheme</a> = 0x0503
    <span id="ECDSAWithP521AndSHA512">ECDSAWithP521AndSHA512</span> <a href="#SignatureScheme">SignatureScheme</a> = 0x0603

    <span class="comment">// EdDSA algorithms.</span>
    <span id="Ed25519">Ed25519</span> <a href="#SignatureScheme">SignatureScheme</a> = 0x0807

    <span class="comment">// Legacy signature and hash algorithms for TLS 1.2.</span>
    <span id="PKCS1WithSHA1">PKCS1WithSHA1</span> <a href="#SignatureScheme">SignatureScheme</a> = 0x0201
    <span id="ECDSAWithSHA1">ECDSAWithSHA1</span> <a href="#SignatureScheme">SignatureScheme</a> = 0x0203
)</pre>













