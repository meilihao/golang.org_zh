

# mail
`import "net/mail"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package mail implements parsing of mail messages.

For the most part, this package follows the syntax as specified by RFC 5322 and
extended by RFC 6532.
Notable divergences:


	* Obsolete address formats are not parsed, including addresses with
	  embedded route information.
	* The full range of spacing (the CFWS syntax element) is not supported,
	  such as breaking addresses across lines.
	* No unicode normalization is performed.
	* The special characters ()[]:;@\, are allowed to appear unquoted in names.




## <a id="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func ParseDate(date string) (time.Time, error)](#ParseDate)
* [type Address](#Address)
  * [func ParseAddress(address string) (*Address, error)](#ParseAddress)
  * [func ParseAddressList(list string) ([]*Address, error)](#ParseAddressList)
  * [func (a *Address) String() string](#Address.String)
* [type AddressParser](#AddressParser)
  * [func (p *AddressParser) Parse(address string) (*Address, error)](#AddressParser.Parse)
  * [func (p *AddressParser) ParseList(list string) ([]*Address, error)](#AddressParser.ParseList)
* [type Header](#Header)
  * [func (h Header) AddressList(key string) ([]*Address, error)](#Header.AddressList)
  * [func (h Header) Date() (time.Time, error)](#Header.Date)
  * [func (h Header) Get(key string) string](#Header.Get)
* [type Message](#Message)
  * [func ReadMessage(r io.Reader) (msg *Message, err error)](#ReadMessage)


#### <a id="pkg-examples">Examples</a>
* [ParseAddress](#example_ParseAddress)
* [ParseAddressList](#example_ParseAddressList)
* [ReadMessage](#example_ReadMessage)


#### <a id="pkg-files">Package files</a>
[message.go](https://golang.org/src/net/mail/message.go) 




## <a id="pkg-variables">Variables</a>

<pre>var <span id="ErrHeaderNotPresent">ErrHeaderNotPresent</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;mail: header not in message&#34;)</pre>

## <a id="ParseDate">func</a> [ParseDate](https://golang.org/src/net/mail/message.go?s=2394:2440#L89)
<pre>func ParseDate(date <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseDate parses an RFC 5322 date string.





## <a id="Address">type</a> [Address](https://golang.org/src/net/mail/message.go?s=3851:3951#L136)
Address represents a single mail address.
An address such as "Barry Gibbs <bg@example.com>" is represented
as Address{Name: "Barry Gibbs", Address: "bg@example.com"}.


<pre>type Address struct {
<span id="Address.Name"></span>    Name    <a href="/pkg/builtin/#string">string</a> <span class="comment">// Proper name; may be empty.</span>
<span id="Address.Address"></span>    Address <a href="/pkg/builtin/#string">string</a> <span class="comment">// user@domain</span>
}
</pre>









### <a id="ParseAddress">func</a> [ParseAddress](https://golang.org/src/net/mail/message.go?s=4039:4090#L142)
<pre>func ParseAddress(address <a href="/pkg/builtin/#string">string</a>) (*<a href="#Address">Address</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseAddress parses a single RFC 5322 address, e.g. "Barry Gibbs <bg@example.com>"


<a id="example_ParseAddress">Example</a>


### <a id="ParseAddressList">func</a> [ParseAddressList](https://golang.org/src/net/mail/message.go?s=4219:4273#L147)
<pre>func ParseAddressList(list <a href="/pkg/builtin/#string">string</a>) ([]*<a href="#Address">Address</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseAddressList parses the given string as a list of addresses.


<a id="example_ParseAddressList">Example</a>




### <a id="Address.String">func</a> (\*Address) [String](https://golang.org/src/net/mail/message.go?s=5217:5250#L172)
<pre>func (a *<a href="#Address">Address</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String formats the address as a valid RFC 5322 address.
If the address's name contains non-ASCII characters
the name will be rendered according to RFC 2047.




## <a id="AddressParser">type</a> [AddressParser](https://golang.org/src/net/mail/message.go?s=4380:4515#L152)
An AddressParser is an RFC 5322 address parser.


<pre>type AddressParser struct {
<span id="AddressParser.WordDecoder"></span>    <span class="comment">// WordDecoder optionally specifies a decoder for RFC 2047 encoded-words.</span>
    WordDecoder *<a href="/pkg/mime/">mime</a>.<a href="/pkg/mime/#WordDecoder">WordDecoder</a>
}
</pre>











### <a id="AddressParser.Parse">func</a> (\*AddressParser) [Parse](https://golang.org/src/net/mail/message.go?s=4624:4687#L159)
<pre>func (p *<a href="#AddressParser">AddressParser</a>) Parse(address <a href="/pkg/builtin/#string">string</a>) (*<a href="#Address">Address</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Parse parses a single RFC 5322 address of the
form "Gogh Fir <gf@example.com>" or "foo@example.com".




### <a id="AddressParser.ParseList">func</a> (\*AddressParser) [ParseList](https://golang.org/src/net/mail/message.go?s=4909:4975#L165)
<pre>func (p *<a href="#AddressParser">AddressParser</a>) ParseList(list <a href="/pkg/builtin/#string">string</a>) ([]*<a href="#Address">Address</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseList parses the given string as a list of comma-separated addresses
of the form "Gogh Fir <gf@example.com>" or "foo@example.com".




## <a id="Header">type</a> [Header](https://golang.org/src/net/mail/message.go?s=2743:2774#L101)
A Header represents the key-value pairs in a mail message header.


<pre>type Header map[<a href="/pkg/builtin/#string">string</a>][]<a href="/pkg/builtin/#string">string</a></pre>











### <a id="Header.AddressList">func</a> (Header) [AddressList](https://golang.org/src/net/mail/message.go?s=3508:3567#L125)
<pre>func (h <a href="#Header">Header</a>) AddressList(key <a href="/pkg/builtin/#string">string</a>) ([]*<a href="#Address">Address</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
AddressList parses the named header field as a list of addresses.




### <a id="Header.Date">func</a> (Header) [Date](https://golang.org/src/net/mail/message.go?s=3286:3327#L116)
<pre>func (h <a href="#Header">Header</a>) Date() (<a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Date parses the Date header field.




### <a id="Header.Get">func</a> (Header) [Get](https://golang.org/src/net/mail/message.go?s=3094:3132#L109)
<pre>func (h <a href="#Header">Header</a>) Get(key <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Get gets the first value associated with the given key.
It is case insensitive; CanonicalMIMEHeaderKey is used
to canonicalize the provided key.
If there are no values associated with the key, Get returns "".
To access multiple values of a key, or to use non-canonical keys,
access the map directly.




## <a id="Message">type</a> [Message](https://golang.org/src/net/mail/message.go?s=1006:1062#L35)
A Message represents a parsed mail message.


<pre>type Message struct {
<span id="Message.Header"></span>    Header <a href="#Header">Header</a>
<span id="Message.Body"></span>    Body   <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>
}
</pre>









### <a id="ReadMessage">func</a> [ReadMessage](https://golang.org/src/net/mail/message.go?s=1206:1261#L43)
<pre>func ReadMessage(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (msg *<a href="#Message">Message</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadMessage reads a message from r.
The headers are parsed, and the body of the message will be available
for reading from msg.Body.


<a id="example_ReadMessage">Example</a>








