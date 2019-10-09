

# smtp
`import "net/smtp"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package smtp implements the Simple Mail Transfer Protocol as defined in RFC 5321.
It also implements the following extensions:


	8BITMIME  RFC 1652
	AUTH      RFC 2554
	STARTTLS  RFC 3207

Additional extensions may be handled by clients.

The smtp package is frozen and is not accepting new features.
Some external packages provide more functionality. See:


	<a href="https://godoc.org/?q=smtp">https://godoc.org/?q=smtp</a>


<a id="example_">Example</a>


## <a id="pkg-index">Index</a>
* [func SendMail(addr string, a Auth, from string, to []string, msg []byte) error](#SendMail)
* [type Auth](#Auth)
  * [func CRAMMD5Auth(username, secret string) Auth](#CRAMMD5Auth)
  * [func PlainAuth(identity, username, password, host string) Auth](#PlainAuth)
* [type Client](#Client)
  * [func Dial(addr string) (*Client, error)](#Dial)
  * [func NewClient(conn net.Conn, host string) (*Client, error)](#NewClient)
  * [func (c *Client) Auth(a Auth) error](#Client.Auth)
  * [func (c *Client) Close() error](#Client.Close)
  * [func (c *Client) Data() (io.WriteCloser, error)](#Client.Data)
  * [func (c *Client) Extension(ext string) (bool, string)](#Client.Extension)
  * [func (c *Client) Hello(localName string) error](#Client.Hello)
  * [func (c *Client) Mail(from string) error](#Client.Mail)
  * [func (c *Client) Noop() error](#Client.Noop)
  * [func (c *Client) Quit() error](#Client.Quit)
  * [func (c *Client) Rcpt(to string) error](#Client.Rcpt)
  * [func (c *Client) Reset() error](#Client.Reset)
  * [func (c *Client) StartTLS(config *tls.Config) error](#Client.StartTLS)
  * [func (c *Client) TLSConnectionState() (state tls.ConnectionState, ok bool)](#Client.TLSConnectionState)
  * [func (c *Client) Verify(addr string) error](#Client.Verify)
* [type ServerInfo](#ServerInfo)


#### <a id="pkg-examples">Examples</a>
* [Package](#example_)
* [PlainAuth](#example_PlainAuth)
* [SendMail](#example_SendMail)


#### <a id="pkg-files">Package files</a>
[auth.go](https://golang.org/src/net/smtp/auth.go) [smtp.go](https://golang.org/src/net/smtp/smtp.go) 






## <a id="SendMail">func</a> [SendMail](https://golang.org/src/net/smtp/smtp.go?s=9230:9308#L309)
<pre>func SendMail(addr <a href="/pkg/builtin/#string">string</a>, a <a href="#Auth">Auth</a>, from <a href="/pkg/builtin/#string">string</a>, to []<a href="/pkg/builtin/#string">string</a>, msg []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
SendMail connects to the server at addr, switches to TLS if
possible, authenticates with the optional mechanism a if possible,
and then sends an email from address from, to addresses to, with
message msg.
The addr must include a port, as in "mail.example.com:smtp".

The addresses in the to parameter are the SMTP RCPT addresses.

The msg parameter should be an RFC 822-style email with headers
first, a blank line, and then the message body. The lines of msg
should be CRLF terminated. The msg headers should usually include
fields such as "From", "To", "Subject", and "Cc".  Sending "Bcc"
messages is accomplished by including an email address in the to
parameter but not including it in the msg headers.

The SendMail function and the net/smtp package are low-level
mechanisms and provide no support for DKIM signing, MIME
attachments (see the mime/multipart package), or other mail
functionality. Higher-level packages exist outside of the standard
library.


<a id="example_SendMail">Example</a>



## <a id="Auth">type</a> [Auth](https://golang.org/src/net/smtp/auth.go?s=292:1191#L5)
Auth is implemented by an SMTP authentication mechanism.


<pre>type Auth interface {
    <span class="comment">// Start begins an authentication with a server.</span>
    <span class="comment">// It returns the name of the authentication protocol</span>
    <span class="comment">// and optionally data to include in the initial AUTH message</span>
    <span class="comment">// sent to the server. It can return proto == &#34;&#34; to indicate</span>
    <span class="comment">// that the authentication should be skipped.</span>
    <span class="comment">// If it returns a non-nil error, the SMTP client aborts</span>
    <span class="comment">// the authentication attempt and closes the connection.</span>
    Start(server *<a href="#ServerInfo">ServerInfo</a>) (proto <a href="/pkg/builtin/#string">string</a>, toServer []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)

    <span class="comment">// Next continues the authentication. The server has just sent</span>
    <span class="comment">// the fromServer data. If more is true, the server expects a</span>
    <span class="comment">// response, which Next should return as toServer; otherwise</span>
    <span class="comment">// Next should return toServer == nil.</span>
    <span class="comment">// If Next returns a non-nil error, the SMTP client aborts</span>
    <span class="comment">// the authentication attempt and closes the connection.</span>
    Next(fromServer []<a href="/pkg/builtin/#byte">byte</a>, more <a href="/pkg/builtin/#bool">bool</a>) (toServer []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)
}</pre>









### <a id="CRAMMD5Auth">func</a> [CRAMMD5Auth](https://golang.org/src/net/smtp/auth.go?s=3375:3421#L84)
<pre>func CRAMMD5Auth(username, secret <a href="/pkg/builtin/#string">string</a>) <a href="#Auth">Auth</a></pre>
CRAMMD5Auth returns an Auth that implements the CRAM-MD5 authentication
mechanism as defined in RFC 2195.
The returned Auth uses the given username and secret to authenticate
to the server using the challenge-response mechanism.




### <a id="PlainAuth">func</a> [PlainAuth](https://golang.org/src/net/smtp/auth.go?s=2004:2066#L44)
<pre>func PlainAuth(identity, username, password, host <a href="/pkg/builtin/#string">string</a>) <a href="#Auth">Auth</a></pre>
PlainAuth returns an Auth that implements the PLAIN authentication
mechanism as defined in RFC 4616. The returned Auth uses the given
username and password to authenticate to host and act as identity.
Usually identity should be the empty string, to act as username.

PlainAuth will only send the credentials if the connection is using TLS
or is connected to localhost. Otherwise authentication will fail with an
error, without sending the credentials.


<a id="example_PlainAuth">Example</a>




## <a id="Client">type</a> [Client](https://golang.org/src/net/smtp/smtp.go?s=751:1341#L20)
A Client represents a client connection to an SMTP server.


<pre>type Client struct {
<span id="Client.Text"></span>    <span class="comment">// Text is the textproto.Conn used by the Client. It is exported to allow for</span>
    <span class="comment">// clients to add extensions.</span>
    Text *<a href="/pkg/net/textproto/">textproto</a>.<a href="/pkg/net/textproto/#Conn">Conn</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="Dial">func</a> [Dial](https://golang.org/src/net/smtp/smtp.go?s=1473:1512#L41)
<pre>func Dial(addr <a href="/pkg/builtin/#string">string</a>) (*<a href="#Client">Client</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Dial returns a new Client connected to an SMTP server at addr.
The addr must include a port, as in "mail.example.com:smtp".




### <a id="NewClient">func</a> [NewClient](https://golang.org/src/net/smtp/smtp.go?s=1785:1844#L52)
<pre>func NewClient(conn <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Conn">Conn</a>, host <a href="/pkg/builtin/#string">string</a>) (*<a href="#Client">Client</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewClient returns a new Client using an existing connection and host as a
server name to be used when authenticating.






### <a id="Client.Auth">func</a> (\*Client) [Auth](https://golang.org/src/net/smtp/smtp.go?s=5600:5635#L189)
<pre>func (c *<a href="#Client">Client</a>) Auth(a <a href="#Auth">Auth</a>) <a href="/pkg/builtin/#error">error</a></pre>
Auth authenticates a client using the provided authentication mechanism.
A failed authentication closes the connection.
Only servers that advertise the AUTH extension support this function.




### <a id="Client.Close">func</a> (\*Client) [Close](https://golang.org/src/net/smtp/smtp.go?s=2129:2159#L65)
<pre>func (c *<a href="#Client">Client</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the connection.




### <a id="Client.Data">func</a> (\*Client) [Data](https://golang.org/src/net/smtp/smtp.go?s=7974:8021#L279)
<pre>func (c *<a href="#Client">Client</a>) Data() (<a href="/pkg/io/">io</a>.<a href="/pkg/io/#WriteCloser">WriteCloser</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Data issues a DATA command to the server and returns a writer that
can be used to write the mail headers and body. The caller should
close the writer before calling any more methods on c. A call to
Data must be preceded by one or more calls to Rcpt.




### <a id="Client.Extension">func</a> (\*Client) [Extension](https://golang.org/src/net/smtp/smtp.go?s=10558:10611#L370)
<pre>func (c *<a href="#Client">Client</a>) Extension(ext <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#bool">bool</a>, <a href="/pkg/builtin/#string">string</a>)</pre>
Extension reports whether an extension is support by the server.
The extension name is case-insensitive. If the extension is supported,
Extension also returns a string that contains any parameters the
server specifies for the extension.




### <a id="Client.Hello">func</a> (\*Client) [Hello](https://golang.org/src/net/smtp/smtp.go?s=2710:2756#L86)
<pre>func (c *<a href="#Client">Client</a>) Hello(localName <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Hello sends a HELO or EHLO to the server as the given host name.
Calling this method is only necessary if the client needs control
over the host name used. The client will introduce itself as "localhost"
automatically otherwise. If Hello is called, it must be called before
any of the other methods.




### <a id="Client.Mail">func</a> (\*Client) [Mail](https://golang.org/src/net/smtp/smtp.go?s=6862:6902#L236)
<pre>func (c *<a href="#Client">Client</a>) Mail(from <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Mail issues a MAIL command to the server using the provided email address.
If the server supports the 8BITMIME extension, Mail adds the BODY=8BITMIME
parameter.
This initiates a mail transaction and is followed by one or more Rcpt calls.




### <a id="Client.Noop">func</a> (\*Client) [Noop](https://golang.org/src/net/smtp/smtp.go?s=11125:11154#L394)
<pre>func (c *<a href="#Client">Client</a>) Noop() <a href="/pkg/builtin/#error">error</a></pre>
Noop sends the NOOP command to the server. It does nothing but check
that the connection to the server is okay.




### <a id="Client.Quit">func</a> (\*Client) [Quit](https://golang.org/src/net/smtp/smtp.go?s=11328:11357#L403)
<pre>func (c *<a href="#Client">Client</a>) Quit() <a href="/pkg/builtin/#error">error</a></pre>
Quit sends the QUIT command and closes the connection to the server.




### <a id="Client.Rcpt">func</a> (\*Client) [Rcpt](https://golang.org/src/net/smtp/smtp.go?s=7383:7421#L256)
<pre>func (c *<a href="#Client">Client</a>) Rcpt(to <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Rcpt issues a RCPT command to the server using the provided email address.
A call to Rcpt must be preceded by a call to Mail and may be followed by
a Data call or another Rcpt call.




### <a id="Client.Reset">func</a> (\*Client) [Reset](https://golang.org/src/net/smtp/smtp.go?s=10875:10905#L384)
<pre>func (c *<a href="#Client">Client</a>) Reset() <a href="/pkg/builtin/#error">error</a></pre>
Reset sends the RSET command to the server, aborting the current mail
transaction.




### <a id="Client.StartTLS">func</a> (\*Client) [StartTLS](https://golang.org/src/net/smtp/smtp.go?s=4352:4403#L146)
<pre>func (c *<a href="#Client">Client</a>) StartTLS(config *<a href="/pkg/crypto/tls/">tls</a>.<a href="/pkg/crypto/tls/#Config">Config</a>) <a href="/pkg/builtin/#error">error</a></pre>
StartTLS sends the STARTTLS command and encrypts all further communication.
Only servers that advertise the STARTTLS extension support this function.




### <a id="Client.TLSConnectionState">func</a> (\*Client) [TLSConnectionState](https://golang.org/src/net/smtp/smtp.go?s=4774:4848#L163)
<pre>func (c *<a href="#Client">Client</a>) TLSConnectionState() (state <a href="/pkg/crypto/tls/">tls</a>.<a href="/pkg/crypto/tls/#ConnectionState">ConnectionState</a>, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
TLSConnectionState returns the client's TLS connection state.
The return values are their zero values if StartTLS did
not succeed.




### <a id="Client.Verify">func</a> (\*Client) [Verify](https://golang.org/src/net/smtp/smtp.go?s=5188:5230#L175)
<pre>func (c *<a href="#Client">Client</a>) Verify(addr <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Verify checks the validity of an email address on the server.
If Verify returns nil, the address is valid. A non-nil return
does not necessarily indicate an invalid address. Many servers
will not verify addresses for security reasons.




## <a id="ServerInfo">type</a> [ServerInfo](https://golang.org/src/net/smtp/auth.go?s=1249:1426#L25)
ServerInfo records information about an SMTP server.


<pre>type ServerInfo struct {
<span id="ServerInfo.Name"></span>    Name <a href="/pkg/builtin/#string">string</a>   <span class="comment">// SMTP server name</span>
<span id="ServerInfo.TLS"></span>    TLS  <a href="/pkg/builtin/#bool">bool</a>     <span class="comment">// using TLS, with valid certificate for Name</span>
<span id="ServerInfo.Auth"></span>    Auth []<a href="/pkg/builtin/#string">string</a> <span class="comment">// advertised authentication mechanisms</span>
}
</pre>















