

# rpc
`import "net/rpc"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package rpc provides access to the exported methods of an object across a
network or other I/O connection.  A server registers an object, making it visible
as a service with the name of the type of the object.  After registration, exported
methods of the object will be accessible remotely.  A server may register multiple
objects (services) of different types but it is an error to register multiple
objects of the same type.

Only methods that satisfy these criteria will be made available for remote access;
other methods will be ignored:


	- the method's type is exported.
	- the method is exported.
	- the method has two arguments, both exported (or builtin) types.
	- the method's second argument is a pointer.
	- the method has return type error.

In effect, the method must look schematically like


	func (t *T) MethodName(argType T1, replyType *T2) error

where T1 and T2 can be marshaled by encoding/gob.
These requirements apply even if a different codec is used.
(In the future, these requirements may soften for custom codecs.)

The method's first argument represents the arguments provided by the caller; the
second argument represents the result parameters to be returned to the caller.
The method's return value, if non-nil, is passed back as a string that the client
sees as if created by errors.New.  If an error is returned, the reply parameter
will not be sent back to the client.

The server may handle requests on a single connection by calling ServeConn.  More
typically it will create a network listener and call Accept or, for an HTTP
listener, HandleHTTP and http.Serve.

A client wishing to use the service establishes a connection and then invokes
NewClient on the connection.  The convenience function Dial (DialHTTP) performs
both steps for a raw network connection (an HTTP connection).  The resulting
Client object has two methods, Call and Go, that specify the service and method to
call, a pointer containing the arguments, and a pointer to receive the result
parameters.

The Call method waits for the remote call to complete while the Go method
launches the call asynchronously and signals completion using the Call
structure's Done channel.

Unless an explicit codec is set up, package encoding/gob is used to
transport the data.

Here is a simple example.  A server wishes to export an object of type Arith:


	package server
	
	import "errors"
	
	type Args struct {
		A, B int
	}
	
	type Quotient struct {
		Quo, Rem int
	}
	
	type Arith int
	
	func (t *Arith) Multiply(args *Args, reply *int) error {
		*reply = args.A * args.B
		return nil
	}
	
	func (t *Arith) Divide(args *Args, quo *Quotient) error {
		if args.B == 0 {
			return errors.New("divide by zero")
		}
		quo.Quo = args.A / args.B
		quo.Rem = args.A % args.B
		return nil
	}

The server calls (for HTTP service):


	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

At this point, clients can see a service "Arith" with methods "Arith.Multiply" and
"Arith.Divide".  To invoke one, a client first dials the server:


	client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

Then it can make a remote call:


	// Synchronous call
	args := &server.Args{7,8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

or


	// Asynchronous call
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done	// will be equal to divCall
	// check errors, print, etc.

A server implementation will often provide a simple, type-safe wrapper for the
client.

The net/rpc package is frozen and is not accepting new features.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func Accept(lis net.Listener)](#Accept)
* [func HandleHTTP()](#HandleHTTP)
* [func Register(rcvr interface{}) error](#Register)
* [func RegisterName(name string, rcvr interface{}) error](#RegisterName)
* [func ServeCodec(codec ServerCodec)](#ServeCodec)
* [func ServeConn(conn io.ReadWriteCloser)](#ServeConn)
* [func ServeRequest(codec ServerCodec) error](#ServeRequest)
* [type Call](#Call)
* [type Client](#Client)
  * [func Dial(network, address string) (*Client, error)](#Dial)
  * [func DialHTTP(network, address string) (*Client, error)](#DialHTTP)
  * [func DialHTTPPath(network, address, path string) (*Client, error)](#DialHTTPPath)
  * [func NewClient(conn io.ReadWriteCloser) *Client](#NewClient)
  * [func NewClientWithCodec(codec ClientCodec) *Client](#NewClientWithCodec)
  * [func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error](#Client.Call)
  * [func (client *Client) Close() error](#Client.Close)
  * [func (client *Client) Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call](#Client.Go)
* [type ClientCodec](#ClientCodec)
* [type Request](#Request)
* [type Response](#Response)
* [type Server](#Server)
  * [func NewServer() *Server](#NewServer)
  * [func (server *Server) Accept(lis net.Listener)](#Server.Accept)
  * [func (server *Server) HandleHTTP(rpcPath, debugPath string)](#Server.HandleHTTP)
  * [func (server *Server) Register(rcvr interface{}) error](#Server.Register)
  * [func (server *Server) RegisterName(name string, rcvr interface{}) error](#Server.RegisterName)
  * [func (server *Server) ServeCodec(codec ServerCodec)](#Server.ServeCodec)
  * [func (server *Server) ServeConn(conn io.ReadWriteCloser)](#Server.ServeConn)
  * [func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request)](#Server.ServeHTTP)
  * [func (server *Server) ServeRequest(codec ServerCodec) error](#Server.ServeRequest)
* [type ServerCodec](#ServerCodec)
* [type ServerError](#ServerError)
  * [func (e ServerError) Error() string](#ServerError.Error)




#### <a id="pkg-files">Package files</a>
[client.go](https://golang.org/src/net/rpc/client.go) [debug.go](https://golang.org/src/net/rpc/debug.go) [server.go](https://golang.org/src/net/rpc/server.go) 


## <a id="pkg-constants">Constants</a>

<pre>const (
    <span class="comment">// Defaults used by HandleHTTP</span>
    <span id="DefaultRPCPath">DefaultRPCPath</span>   = &#34;/_goRPC_&#34;
    <span id="DefaultDebugPath">DefaultDebugPath</span> = &#34;/debug/rpc&#34;
)</pre>

## <a id="pkg-variables">Variables</a>
DefaultServer is the default instance of *Server.


<pre>var <span id="DefaultServer">DefaultServer</span> = <a href="#NewServer">NewServer</a>()</pre>
<pre>var <span id="ErrShutdown">ErrShutdown</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;connection is shut down&#34;)</pre>

## <a id="Accept">func</a> [Accept](https://golang.org/src/net/rpc/server.go?s=20693:20722#L675)
<pre>func Accept(lis <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Listener">Listener</a>)</pre>
Accept accepts connections on the listener and serves requests
to DefaultServer for each incoming connection.
Accept blocks; the caller typically invokes it in a go statement.



## <a id="HandleHTTP">func</a> [HandleHTTP](https://golang.org/src/net/rpc/server.go?s=21953:21970#L708)
<pre>func HandleHTTP()</pre>
HandleHTTP registers an HTTP handler for RPC messages to DefaultServer
on DefaultRPCPath and a debugging handler on DefaultDebugPath.
It is still necessary to invoke http.Serve(), typically in a go statement.



## <a id="Register">func</a> [Register](https://golang.org/src/net/rpc/server.go?s=18599:18636#L625)
<pre>func Register(rcvr interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Register publishes the receiver's methods in the DefaultServer.



## <a id="RegisterName">func</a> [RegisterName](https://golang.org/src/net/rpc/server.go?s=18795:18849#L629)
<pre>func RegisterName(name <a href="/pkg/builtin/#string">string</a>, rcvr interface{}) <a href="/pkg/builtin/#error">error</a></pre>
RegisterName is like Register but uses the provided name for the type
instead of the receiver's concrete type.



## <a id="ServeCodec">func</a> [ServeCodec](https://golang.org/src/net/rpc/server.go?s=20219:20253#L662)
<pre>func ServeCodec(codec <a href="#ServerCodec">ServerCodec</a>)</pre>
ServeCodec is like ServeConn but uses the specified codec to
decode requests and encode responses.



## <a id="ServeConn">func</a> [ServeConn](https://golang.org/src/net/rpc/server.go?s=20038:20077#L656)
<pre>func ServeConn(conn <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadWriteCloser">ReadWriteCloser</a>)</pre>
ServeConn runs the DefaultServer on a single connection.
ServeConn blocks, serving the connection until the client hangs up.
The caller typically invokes ServeConn in a go statement.
ServeConn uses the gob wire format (see package gob) on the
connection. To use an alternate codec, use ServeCodec.
See NewClient's comment for information about concurrent access.



## <a id="ServeRequest">func</a> [ServeRequest](https://golang.org/src/net/rpc/server.go?s=20418:20460#L668)
<pre>func ServeRequest(codec <a href="#ServerCodec">ServerCodec</a>) <a href="/pkg/builtin/#error">error</a></pre>
ServeRequest is like ServeCodec but synchronously serves a single request.
It does not close the codec upon completion.





## <a id="Call">type</a> [Call](https://golang.org/src/net/rpc/client.go?s=540:900#L19)
Call represents an active RPC.


<pre>type Call struct {
<span id="Call.ServiceMethod"></span>    ServiceMethod <a href="/pkg/builtin/#string">string</a>      <span class="comment">// The name of the service and method to call.</span>
<span id="Call.Args"></span>    Args          interface{} <span class="comment">// The argument to the function (*struct).</span>
<span id="Call.Reply"></span>    Reply         interface{} <span class="comment">// The reply from the function (*struct).</span>
<span id="Call.Error"></span>    Error         <a href="/pkg/builtin/#error">error</a>       <span class="comment">// After completion, the error status.</span>
<span id="Call.Done"></span>    Done          chan *<a href="#Call">Call</a>  <span class="comment">// Strobes when call is complete.</span>
}
</pre>











## <a id="Client">type</a> [Client](https://golang.org/src/net/rpc/client.go?s=1084:1360#L31)
Client represents an RPC Client.
There may be multiple outstanding Calls associated
with a single Client, and a Client may be used by
multiple goroutines simultaneously.


<pre>type Client struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="Dial">func</a> [Dial](https://golang.org/src/net/rpc/client.go?s=7475:7526#L264)
<pre>func Dial(network, address <a href="/pkg/builtin/#string">string</a>) (*<a href="#Client">Client</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Dial connects to an RPC server at the specified network address.




### <a id="DialHTTP">func</a> [DialHTTP](https://golang.org/src/net/rpc/client.go?s=6528:6583#L231)
<pre>func DialHTTP(network, address <a href="/pkg/builtin/#string">string</a>) (*<a href="#Client">Client</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DialHTTP connects to an HTTP RPC server at the specified network address
listening on the default HTTP RPC path.




### <a id="DialHTTPPath">func</a> [DialHTTPPath](https://golang.org/src/net/rpc/client.go?s=6737:6802#L237)
<pre>func DialHTTPPath(network, address, path <a href="/pkg/builtin/#string">string</a>) (*<a href="#Client">Client</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DialHTTPPath connects to an HTTP RPC server
at the specified network address and path.




### <a id="NewClient">func</a> [NewClient](https://golang.org/src/net/rpc/client.go?s=5326:5373#L183)
<pre>func NewClient(conn <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadWriteCloser">ReadWriteCloser</a>) *<a href="#Client">Client</a></pre>
NewClient returns a new Client to handle requests to the
set of services at the other end of the connection.
It adds a buffer to the write side of the connection so
the header and payload are sent as a unit.

The read and write halves of the connection are serialized independently,
so no interlocking is required. However each half may be accessed
concurrently so the implementation of conn should protect against
concurrent reads or concurrent writes.




### <a id="NewClientWithCodec">func</a> [NewClientWithCodec](https://golang.org/src/net/rpc/client.go?s=5647:5697#L191)
<pre>func NewClientWithCodec(codec <a href="#ClientCodec">ClientCodec</a>) *<a href="#Client">Client</a></pre>
NewClientWithCodec is like NewClient but uses the specified
codec to encode requests and decode responses.






### <a id="Client.Call">func</a> (\*Client) [Call](https://golang.org/src/net/rpc/client.go?s=9009:9100#L311)
<pre>func (client *<a href="#Client">Client</a>) Call(serviceMethod <a href="/pkg/builtin/#string">string</a>, args interface{}, reply interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Call invokes the named function, waits for it to complete, and returns its error status.




### <a id="Client.Close">func</a> (\*Client) [Close](https://golang.org/src/net/rpc/client.go?s=7764:7799#L274)
<pre>func (client *<a href="#Client">Client</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close calls the underlying codec's Close method. If the connection is already
shutting down, ErrShutdown is returned.




### <a id="Client.Go">func</a> (\*Client) [Go](https://golang.org/src/net/rpc/client.go?s=8284:8390#L289)
<pre>func (client *<a href="#Client">Client</a>) Go(serviceMethod <a href="/pkg/builtin/#string">string</a>, args interface{}, reply interface{}, done chan *<a href="#Call">Call</a>) *<a href="#Call">Call</a></pre>
Go invokes the function asynchronously. It returns the Call structure representing
the invocation. The done channel will signal when the call is complete by returning
the same Call object. If done is nil, Go will allocate a new channel.
If non-nil, done must be buffered or Go will deliberately crash.




## <a id="ClientCodec">type</a> [ClientCodec](https://golang.org/src/net/rpc/client.go?s=1890:2053#L53)
A ClientCodec implements writing of RPC requests and
reading of RPC responses for the client side of an RPC session.
The client calls WriteRequest to write a request to the connection
and calls ReadResponseHeader and ReadResponseBody in pairs
to read responses. The client calls Close when finished with the
connection. ReadResponseBody may be called with a nil
argument to force the body of the response to be read and then
discarded.
See NewClient's comment for information about concurrent access.


<pre>type ClientCodec interface {
    WriteRequest(*<a href="#Request">Request</a>, interface{}) <a href="/pkg/builtin/#error">error</a>
    ReadResponseHeader(*<a href="#Response">Response</a>) <a href="/pkg/builtin/#error">error</a>
    ReadResponseBody(interface{}) <a href="/pkg/builtin/#error">error</a>

    Close() <a href="/pkg/builtin/#error">error</a>
}</pre>











## <a id="Request">type</a> [Request](https://golang.org/src/net/rpc/server.go?s=5113:5299#L161)
Request is a header written before every RPC call. It is used internally
but documented here as an aid to debugging, such as when analyzing
network traffic.


<pre>type Request struct {
<span id="Request.ServiceMethod"></span>    ServiceMethod <a href="/pkg/builtin/#string">string</a> <span class="comment">// format: &#34;Service.Method&#34;</span>
<span id="Request.Seq"></span>    Seq           <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">// sequence number chosen by client</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











## <a id="Response">type</a> [Response](https://golang.org/src/net/rpc/server.go?s=5470:5699#L170)
Response is a header written before every RPC return. It is used internally
but documented here as an aid to debugging, such as when analyzing
network traffic.


<pre>type Response struct {
<span id="Response.ServiceMethod"></span>    ServiceMethod <a href="/pkg/builtin/#string">string</a> <span class="comment">// echoes that of the Request</span>
<span id="Response.Seq"></span>    Seq           <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">// echoes that of the request</span>
<span id="Response.Error"></span>    Error         <a href="/pkg/builtin/#string">string</a> <span class="comment">// error, if any.</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











## <a id="Server">type</a> [Server](https://golang.org/src/net/rpc/server.go?s=5737:5935#L178)
Server represents an RPC Server.


<pre>type Server struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewServer">func</a> [NewServer](https://golang.org/src/net/rpc/server.go?s=5972:5996#L187)
<pre>func NewServer() *<a href="#Server">Server</a></pre>
NewServer returns a new Server.






### <a id="Server.Accept">func</a> (\*Server) [Accept](https://golang.org/src/net/rpc/server.go?s=18334:18380#L613)
<pre>func (server *<a href="#Server">Server</a>) Accept(lis <a href="/pkg/net/">net</a>.<a href="/pkg/net/#Listener">Listener</a>)</pre>
Accept accepts connections on the listener and serves requests
for each incoming connection. Accept blocks until the listener
returns a non-nil error. The caller typically invokes Accept in a
go statement.




### <a id="Server.HandleHTTP">func</a> (\*Server) [HandleHTTP](https://golang.org/src/net/rpc/server.go?s=21597:21656#L700)
<pre>func (server *<a href="#Server">Server</a>) HandleHTTP(rpcPath, debugPath <a href="/pkg/builtin/#string">string</a>)</pre>
HandleHTTP registers an HTTP handler for RPC messages on rpcPath,
and a debugging handler on debugPath.
It is still necessary to invoke http.Serve(), typically in a go statement.




### <a id="Server.Register">func</a> (\*Server) [Register](https://golang.org/src/net/rpc/server.go?s=6943:6997#L214)
<pre>func (server *<a href="#Server">Server</a>) Register(rcvr interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Register publishes in the server the set of methods of the
receiver value that satisfy the following conditions:


	- exported method of exported type
	- two arguments, both of exported type
	- the second argument is a pointer
	- one return value, of type error

It returns an error if the receiver is not an exported type or has
no suitable methods. It also logs the error using package log.
The client accesses each method using a string of the form "Type.Method",
where Type is the receiver's concrete type.




### <a id="Server.RegisterName">func</a> (\*Server) [RegisterName](https://golang.org/src/net/rpc/server.go?s=7161:7232#L220)
<pre>func (server *<a href="#Server">Server</a>) RegisterName(name <a href="/pkg/builtin/#string">string</a>, rcvr interface{}) <a href="/pkg/builtin/#error">error</a></pre>
RegisterName is like Register but uses the provided name for the type
instead of the receiver's concrete type.




### <a id="Server.ServeCodec">func</a> (\*Server) [ServeCodec](https://golang.org/src/net/rpc/server.go?s=13814:13865#L444)
<pre>func (server *<a href="#Server">Server</a>) ServeCodec(codec <a href="#ServerCodec">ServerCodec</a>)</pre>
ServeCodec is like ServeConn but uses the specified codec to
decode requests and encode responses.




### <a id="Server.ServeConn">func</a> (\*Server) [ServeConn](https://golang.org/src/net/rpc/server.go?s=13471:13527#L431)
<pre>func (server *<a href="#Server">Server</a>) ServeConn(conn <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadWriteCloser">ReadWriteCloser</a>)</pre>
ServeConn runs the server on a single connection.
ServeConn blocks, serving the connection until the client hangs up.
The caller typically invokes ServeConn in a go statement.
ServeConn uses the gob wire format (see package gob) on the
connection. To use an alternate codec, use ServeCodec.
See NewClient's comment for information about concurrent access.




### <a id="Server.ServeHTTP">func</a> (\*Server) [ServeHTTP](https://golang.org/src/net/rpc/server.go?s=20925:20998#L681)
<pre>func (server *<a href="#Server">Server</a>) ServeHTTP(w <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#ResponseWriter">ResponseWriter</a>, req *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>)</pre>
ServeHTTP implements an http.Handler that answers RPC requests.




### <a id="Server.ServeRequest">func</a> (\*Server) [ServeRequest](https://golang.org/src/net/rpc/server.go?s=14686:14745#L474)
<pre>func (server *<a href="#Server">Server</a>) ServeRequest(codec <a href="#ServerCodec">ServerCodec</a>) <a href="/pkg/builtin/#error">error</a></pre>
ServeRequest is like ServeCodec but synchronously serves a single request.
It does not close the codec upon completion.




## <a id="ServerCodec">type</a> [ServerCodec](https://golang.org/src/net/rpc/server.go?s=19430:19655#L641)
A ServerCodec implements reading of RPC requests and writing of
RPC responses for the server side of an RPC session.
The server calls ReadRequestHeader and ReadRequestBody in pairs
to read requests from the connection, and it calls WriteResponse to
write a response back. The server calls Close when finished with the
connection. ReadRequestBody may be called with a nil
argument to force the body of the request to be read and discarded.
See NewClient's comment for information about concurrent access.


<pre>type ServerCodec interface {
    ReadRequestHeader(*<a href="#Request">Request</a>) <a href="/pkg/builtin/#error">error</a>
    ReadRequestBody(interface{}) <a href="/pkg/builtin/#error">error</a>
    WriteResponse(*<a href="#Response">Response</a>, interface{}) <a href="/pkg/builtin/#error">error</a>

    <span class="comment">// Close can be called multiple times and must be idempotent.</span>
    Close() <a href="/pkg/builtin/#error">error</a>
}</pre>











## <a id="ServerError">type</a> [ServerError](https://golang.org/src/net/rpc/client.go?s=365:388#L10)
ServerError represents an error that has been returned from
the remote side of the RPC connection.


<pre>type ServerError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="ServerError.Error">func</a> (ServerError) [Error](https://golang.org/src/net/rpc/client.go?s=390:425#L12)
<pre>func (e <a href="#ServerError">ServerError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>






