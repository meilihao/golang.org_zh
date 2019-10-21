

# jsonrpc
`import "net/rpc/jsonrpc"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package jsonrpc implements a JSON-RPC 1.0 ClientCodec and ServerCodec
for the rpc package.
For JSON-RPC 2.0 support, see <a href="https://godoc.org/?q=json-rpc+2.0">https://godoc.org/?q=json-rpc+2.0</a>




## <a id="pkg-index">Index</a>
* [func Dial(network, address string) (*rpc.Client, error)](#Dial)
* [func NewClient(conn io.ReadWriteCloser) *rpc.Client](#NewClient)
* [func NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec](#NewClientCodec)
* [func NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec](#NewServerCodec)
* [func ServeConn(conn io.ReadWriteCloser)](#ServeConn)




#### <a id="pkg-files">Package files</a>
[client.go](https://golang.org/src/net/rpc/jsonrpc/client.go) [server.go](https://golang.org/src/net/rpc/jsonrpc/server.go) 






## <a id="Dial">func</a> [Dial](https://golang.org/src/net/rpc/jsonrpc/client.go?s=2889:2944#L108)
<pre>func Dial(network, address <a href="/pkg/builtin/#string">string</a>) (*<a href="/pkg/net/rpc/">rpc</a>.<a href="/pkg/net/rpc/#Client">Client</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Dial connects to a JSON-RPC server at the specified network address.



## <a id="NewClient">func</a> [NewClient](https://golang.org/src/net/rpc/jsonrpc/client.go?s=2707:2758#L103)
<pre>func NewClient(conn <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadWriteCloser">ReadWriteCloser</a>) *<a href="/pkg/net/rpc/">rpc</a>.<a href="/pkg/net/rpc/#Client">Client</a></pre>
NewClient returns a new rpc.Client to handle requests to the
set of services at the other end of the connection.



## <a id="NewClientCodec">func</a> [NewClientCodec](https://golang.org/src/net/rpc/jsonrpc/client.go?s=1034:1094#L27)
<pre>func NewClientCodec(conn <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadWriteCloser">ReadWriteCloser</a>) <a href="/pkg/net/rpc/">rpc</a>.<a href="/pkg/net/rpc/#ClientCodec">ClientCodec</a></pre>
NewClientCodec returns a new rpc.ClientCodec using JSON-RPC on conn.



## <a id="NewServerCodec">func</a> [NewServerCodec](https://golang.org/src/net/rpc/jsonrpc/server.go?s=998:1058#L27)
<pre>func NewServerCodec(conn <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadWriteCloser">ReadWriteCloser</a>) <a href="/pkg/net/rpc/">rpc</a>.<a href="/pkg/net/rpc/#ServerCodec">ServerCodec</a></pre>
NewServerCodec returns a new rpc.ServerCodec using JSON-RPC on conn.



## <a id="ServeConn">func</a> [ServeConn](https://golang.org/src/net/rpc/jsonrpc/server.go?s=3191:3230#L122)
<pre>func ServeConn(conn <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadWriteCloser">ReadWriteCloser</a>)</pre>
ServeConn runs the JSON-RPC server on a single connection.
ServeConn blocks, serving the connection until the client hangs up.
The caller typically invokes ServeConn in a go statement.








