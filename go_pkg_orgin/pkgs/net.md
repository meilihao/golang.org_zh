

# net
`import "net"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package net provides a portable interface for network I/O, including
TCP/IP, UDP, domain name resolution, and Unix domain sockets.

Although the package provides access to low-level networking
primitives, most clients will need only the basic interface provided
by the Dial, Listen, and Accept functions and the associated
Conn and Listener interfaces. The crypto/tls package uses
the same interfaces and similar Dial and Listen functions.

The Dial function connects to a server:


	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	// ...

The Listen function creates servers:


	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}

### Name Resolution
The method for resolving domain names, whether indirectly with functions like Dial
or directly with functions like LookupHost and LookupAddr, varies by operating system.

On Unix systems, the resolver has two options for resolving names.
It can use a pure Go resolver that sends DNS requests directly to the servers
listed in /etc/resolv.conf, or it can use a cgo-based resolver that calls C
library routines such as getaddrinfo and getnameinfo.

By default the pure Go resolver is used, because a blocked DNS request consumes
only a goroutine, while a blocked C call consumes an operating system thread.
When cgo is available, the cgo-based resolver is used instead under a variety of
conditions: on systems that do not let programs make direct DNS requests (OS X),
when the LOCALDOMAIN environment variable is present (even if empty),
when the RES_OPTIONS or HOSTALIASES environment variable is non-empty,
when the ASR_CONFIG environment variable is non-empty (OpenBSD only),
when /etc/resolv.conf or /etc/nsswitch.conf specify the use of features that the
Go resolver does not implement, and when the name being looked up ends in .local
or is an mDNS name.

The resolver decision can be overridden by setting the netdns value of the
GODEBUG environment variable (see package runtime) to go or cgo, as in:


	export GODEBUG=netdns=go    # force pure Go resolver
	export GODEBUG=netdns=cgo   # force cgo resolver

The decision can also be forced while building the Go source tree
by setting the netgo or netcgo build tag.

A numeric netdns setting, as in GODEBUG=netdns=1, causes the resolver
to print debugging information about its decisions.
To force a particular resolver while also printing debugging information,
join the two settings by a plus sign, as in GODEBUG=netdns=go+1.

On Plan 9, the resolver always accesses /net/cs and /net/dns.

On Windows, the resolver always uses C library functions, such as GetAddrInfo and DnsQuery.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func JoinHostPort(host, port string) string](#JoinHostPort)
* [func LookupAddr(addr string) (names []string, err error)](#LookupAddr)
* [func LookupCNAME(host string) (cname string, err error)](#LookupCNAME)
* [func LookupHost(host string) (addrs []string, err error)](#LookupHost)
* [func LookupPort(network, service string) (port int, err error)](#LookupPort)
* [func LookupTXT(name string) ([]string, error)](#LookupTXT)
* [func ParseCIDR(s string) (IP, *IPNet, error)](#ParseCIDR)
* [func Pipe() (Conn, Conn)](#Pipe)
* [func SplitHostPort(hostport string) (host, port string, err error)](#SplitHostPort)
* [type Addr](#Addr)
  * [func InterfaceAddrs() ([]Addr, error)](#InterfaceAddrs)
* [type AddrError](#AddrError)
  * [func (e *AddrError) Error() string](#AddrError.Error)
  * [func (e *AddrError) Temporary() bool](#AddrError.Temporary)
  * [func (e *AddrError) Timeout() bool](#AddrError.Timeout)
* [type Buffers](#Buffers)
  * [func (v *Buffers) Read(p []byte) (n int, err error)](#Buffers.Read)
  * [func (v *Buffers) WriteTo(w io.Writer) (n int64, err error)](#Buffers.WriteTo)
* [type Conn](#Conn)
  * [func Dial(network, address string) (Conn, error)](#Dial)
  * [func DialTimeout(network, address string, timeout time.Duration) (Conn, error)](#DialTimeout)
  * [func FileConn(f *os.File) (c Conn, err error)](#FileConn)
* [type DNSConfigError](#DNSConfigError)
  * [func (e *DNSConfigError) Error() string](#DNSConfigError.Error)
  * [func (e *DNSConfigError) Temporary() bool](#DNSConfigError.Temporary)
  * [func (e *DNSConfigError) Timeout() bool](#DNSConfigError.Timeout)
  * [func (e *DNSConfigError) Unwrap() error](#DNSConfigError.Unwrap)
* [type DNSError](#DNSError)
  * [func (e *DNSError) Error() string](#DNSError.Error)
  * [func (e *DNSError) Temporary() bool](#DNSError.Temporary)
  * [func (e *DNSError) Timeout() bool](#DNSError.Timeout)
* [type Dialer](#Dialer)
  * [func (d *Dialer) Dial(network, address string) (Conn, error)](#Dialer.Dial)
  * [func (d *Dialer) DialContext(ctx context.Context, network, address string) (Conn, error)](#Dialer.DialContext)
* [type Error](#Error)
* [type Flags](#Flags)
  * [func (f Flags) String() string](#Flags.String)
* [type HardwareAddr](#HardwareAddr)
  * [func ParseMAC(s string) (hw HardwareAddr, err error)](#ParseMAC)
  * [func (a HardwareAddr) String() string](#HardwareAddr.String)
* [type IP](#IP)
  * [func IPv4(a, b, c, d byte) IP](#IPv4)
  * [func LookupIP(host string) ([]IP, error)](#LookupIP)
  * [func ParseIP(s string) IP](#ParseIP)
  * [func (ip IP) DefaultMask() IPMask](#IP.DefaultMask)
  * [func (ip IP) Equal(x IP) bool](#IP.Equal)
  * [func (ip IP) IsGlobalUnicast() bool](#IP.IsGlobalUnicast)
  * [func (ip IP) IsInterfaceLocalMulticast() bool](#IP.IsInterfaceLocalMulticast)
  * [func (ip IP) IsLinkLocalMulticast() bool](#IP.IsLinkLocalMulticast)
  * [func (ip IP) IsLinkLocalUnicast() bool](#IP.IsLinkLocalUnicast)
  * [func (ip IP) IsLoopback() bool](#IP.IsLoopback)
  * [func (ip IP) IsMulticast() bool](#IP.IsMulticast)
  * [func (ip IP) IsUnspecified() bool](#IP.IsUnspecified)
  * [func (ip IP) MarshalText() ([]byte, error)](#IP.MarshalText)
  * [func (ip IP) Mask(mask IPMask) IP](#IP.Mask)
  * [func (ip IP) String() string](#IP.String)
  * [func (ip IP) To16() IP](#IP.To16)
  * [func (ip IP) To4() IP](#IP.To4)
  * [func (ip *IP) UnmarshalText(text []byte) error](#IP.UnmarshalText)
* [type IPAddr](#IPAddr)
  * [func ResolveIPAddr(network, address string) (*IPAddr, error)](#ResolveIPAddr)
  * [func (a *IPAddr) Network() string](#IPAddr.Network)
  * [func (a *IPAddr) String() string](#IPAddr.String)
* [type IPConn](#IPConn)
  * [func DialIP(network string, laddr, raddr *IPAddr) (*IPConn, error)](#DialIP)
  * [func ListenIP(network string, laddr *IPAddr) (*IPConn, error)](#ListenIP)
  * [func (c *IPConn) Close() error](#IPConn.Close)
  * [func (c *IPConn) File() (f *os.File, err error)](#IPConn.File)
  * [func (c *IPConn) LocalAddr() Addr](#IPConn.LocalAddr)
  * [func (c *IPConn) Read(b []byte) (int, error)](#IPConn.Read)
  * [func (c *IPConn) ReadFrom(b []byte) (int, Addr, error)](#IPConn.ReadFrom)
  * [func (c *IPConn) ReadFromIP(b []byte) (int, *IPAddr, error)](#IPConn.ReadFromIP)
  * [func (c *IPConn) ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error)](#IPConn.ReadMsgIP)
  * [func (c *IPConn) RemoteAddr() Addr](#IPConn.RemoteAddr)
  * [func (c *IPConn) SetDeadline(t time.Time) error](#IPConn.SetDeadline)
  * [func (c *IPConn) SetReadBuffer(bytes int) error](#IPConn.SetReadBuffer)
  * [func (c *IPConn) SetReadDeadline(t time.Time) error](#IPConn.SetReadDeadline)
  * [func (c *IPConn) SetWriteBuffer(bytes int) error](#IPConn.SetWriteBuffer)
  * [func (c *IPConn) SetWriteDeadline(t time.Time) error](#IPConn.SetWriteDeadline)
  * [func (c *IPConn) SyscallConn() (syscall.RawConn, error)](#IPConn.SyscallConn)
  * [func (c *IPConn) Write(b []byte) (int, error)](#IPConn.Write)
  * [func (c *IPConn) WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error)](#IPConn.WriteMsgIP)
  * [func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error)](#IPConn.WriteTo)
  * [func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error)](#IPConn.WriteToIP)
* [type IPMask](#IPMask)
  * [func CIDRMask(ones, bits int) IPMask](#CIDRMask)
  * [func IPv4Mask(a, b, c, d byte) IPMask](#IPv4Mask)
  * [func (m IPMask) Size() (ones, bits int)](#IPMask.Size)
  * [func (m IPMask) String() string](#IPMask.String)
* [type IPNet](#IPNet)
  * [func (n *IPNet) Contains(ip IP) bool](#IPNet.Contains)
  * [func (n *IPNet) Network() string](#IPNet.Network)
  * [func (n *IPNet) String() string](#IPNet.String)
* [type Interface](#Interface)
  * [func InterfaceByIndex(index int) (*Interface, error)](#InterfaceByIndex)
  * [func InterfaceByName(name string) (*Interface, error)](#InterfaceByName)
  * [func Interfaces() ([]Interface, error)](#Interfaces)
  * [func (ifi *Interface) Addrs() ([]Addr, error)](#Interface.Addrs)
  * [func (ifi *Interface) MulticastAddrs() ([]Addr, error)](#Interface.MulticastAddrs)
* [type InvalidAddrError](#InvalidAddrError)
  * [func (e InvalidAddrError) Error() string](#InvalidAddrError.Error)
  * [func (e InvalidAddrError) Temporary() bool](#InvalidAddrError.Temporary)
  * [func (e InvalidAddrError) Timeout() bool](#InvalidAddrError.Timeout)
* [type ListenConfig](#ListenConfig)
  * [func (lc *ListenConfig) Listen(ctx context.Context, network, address string) (Listener, error)](#ListenConfig.Listen)
  * [func (lc *ListenConfig) ListenPacket(ctx context.Context, network, address string) (PacketConn, error)](#ListenConfig.ListenPacket)
* [type Listener](#Listener)
  * [func FileListener(f *os.File) (ln Listener, err error)](#FileListener)
  * [func Listen(network, address string) (Listener, error)](#Listen)
* [type MX](#MX)
  * [func LookupMX(name string) ([]*MX, error)](#LookupMX)
* [type NS](#NS)
  * [func LookupNS(name string) ([]*NS, error)](#LookupNS)
* [type OpError](#OpError)
  * [func (e *OpError) Error() string](#OpError.Error)
  * [func (e *OpError) Temporary() bool](#OpError.Temporary)
  * [func (e *OpError) Timeout() bool](#OpError.Timeout)
  * [func (e *OpError) Unwrap() error](#OpError.Unwrap)
* [type PacketConn](#PacketConn)
  * [func FilePacketConn(f *os.File) (c PacketConn, err error)](#FilePacketConn)
  * [func ListenPacket(network, address string) (PacketConn, error)](#ListenPacket)
* [type ParseError](#ParseError)
  * [func (e *ParseError) Error() string](#ParseError.Error)
* [type Resolver](#Resolver)
  * [func (r *Resolver) LookupAddr(ctx context.Context, addr string) (names []string, err error)](#Resolver.LookupAddr)
  * [func (r *Resolver) LookupCNAME(ctx context.Context, host string) (cname string, err error)](#Resolver.LookupCNAME)
  * [func (r *Resolver) LookupHost(ctx context.Context, host string) (addrs []string, err error)](#Resolver.LookupHost)
  * [func (r *Resolver) LookupIPAddr(ctx context.Context, host string) ([]IPAddr, error)](#Resolver.LookupIPAddr)
  * [func (r *Resolver) LookupMX(ctx context.Context, name string) ([]*MX, error)](#Resolver.LookupMX)
  * [func (r *Resolver) LookupNS(ctx context.Context, name string) ([]*NS, error)](#Resolver.LookupNS)
  * [func (r *Resolver) LookupPort(ctx context.Context, network, service string) (port int, err error)](#Resolver.LookupPort)
  * [func (r *Resolver) LookupSRV(ctx context.Context, service, proto, name string) (cname string, addrs []*SRV, err error)](#Resolver.LookupSRV)
  * [func (r *Resolver) LookupTXT(ctx context.Context, name string) ([]string, error)](#Resolver.LookupTXT)
* [type SRV](#SRV)
  * [func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error)](#LookupSRV)
* [type TCPAddr](#TCPAddr)
  * [func ResolveTCPAddr(network, address string) (*TCPAddr, error)](#ResolveTCPAddr)
  * [func (a *TCPAddr) Network() string](#TCPAddr.Network)
  * [func (a *TCPAddr) String() string](#TCPAddr.String)
* [type TCPConn](#TCPConn)
  * [func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)](#DialTCP)
  * [func (c *TCPConn) Close() error](#TCPConn.Close)
  * [func (c *TCPConn) CloseRead() error](#TCPConn.CloseRead)
  * [func (c *TCPConn) CloseWrite() error](#TCPConn.CloseWrite)
  * [func (c *TCPConn) File() (f *os.File, err error)](#TCPConn.File)
  * [func (c *TCPConn) LocalAddr() Addr](#TCPConn.LocalAddr)
  * [func (c *TCPConn) Read(b []byte) (int, error)](#TCPConn.Read)
  * [func (c *TCPConn) ReadFrom(r io.Reader) (int64, error)](#TCPConn.ReadFrom)
  * [func (c *TCPConn) RemoteAddr() Addr](#TCPConn.RemoteAddr)
  * [func (c *TCPConn) SetDeadline(t time.Time) error](#TCPConn.SetDeadline)
  * [func (c *TCPConn) SetKeepAlive(keepalive bool) error](#TCPConn.SetKeepAlive)
  * [func (c *TCPConn) SetKeepAlivePeriod(d time.Duration) error](#TCPConn.SetKeepAlivePeriod)
  * [func (c *TCPConn) SetLinger(sec int) error](#TCPConn.SetLinger)
  * [func (c *TCPConn) SetNoDelay(noDelay bool) error](#TCPConn.SetNoDelay)
  * [func (c *TCPConn) SetReadBuffer(bytes int) error](#TCPConn.SetReadBuffer)
  * [func (c *TCPConn) SetReadDeadline(t time.Time) error](#TCPConn.SetReadDeadline)
  * [func (c *TCPConn) SetWriteBuffer(bytes int) error](#TCPConn.SetWriteBuffer)
  * [func (c *TCPConn) SetWriteDeadline(t time.Time) error](#TCPConn.SetWriteDeadline)
  * [func (c *TCPConn) SyscallConn() (syscall.RawConn, error)](#TCPConn.SyscallConn)
  * [func (c *TCPConn) Write(b []byte) (int, error)](#TCPConn.Write)
* [type TCPListener](#TCPListener)
  * [func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)](#ListenTCP)
  * [func (l *TCPListener) Accept() (Conn, error)](#TCPListener.Accept)
  * [func (l *TCPListener) AcceptTCP() (*TCPConn, error)](#TCPListener.AcceptTCP)
  * [func (l *TCPListener) Addr() Addr](#TCPListener.Addr)
  * [func (l *TCPListener) Close() error](#TCPListener.Close)
  * [func (l *TCPListener) File() (f *os.File, err error)](#TCPListener.File)
  * [func (l *TCPListener) SetDeadline(t time.Time) error](#TCPListener.SetDeadline)
  * [func (l *TCPListener) SyscallConn() (syscall.RawConn, error)](#TCPListener.SyscallConn)
* [type UDPAddr](#UDPAddr)
  * [func ResolveUDPAddr(network, address string) (*UDPAddr, error)](#ResolveUDPAddr)
  * [func (a *UDPAddr) Network() string](#UDPAddr.Network)
  * [func (a *UDPAddr) String() string](#UDPAddr.String)
* [type UDPConn](#UDPConn)
  * [func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)](#DialUDP)
  * [func ListenMulticastUDP(network string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)](#ListenMulticastUDP)
  * [func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)](#ListenUDP)
  * [func (c *UDPConn) Close() error](#UDPConn.Close)
  * [func (c *UDPConn) File() (f *os.File, err error)](#UDPConn.File)
  * [func (c *UDPConn) LocalAddr() Addr](#UDPConn.LocalAddr)
  * [func (c *UDPConn) Read(b []byte) (int, error)](#UDPConn.Read)
  * [func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error)](#UDPConn.ReadFrom)
  * [func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)](#UDPConn.ReadFromUDP)
  * [func (c *UDPConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error)](#UDPConn.ReadMsgUDP)
  * [func (c *UDPConn) RemoteAddr() Addr](#UDPConn.RemoteAddr)
  * [func (c *UDPConn) SetDeadline(t time.Time) error](#UDPConn.SetDeadline)
  * [func (c *UDPConn) SetReadBuffer(bytes int) error](#UDPConn.SetReadBuffer)
  * [func (c *UDPConn) SetReadDeadline(t time.Time) error](#UDPConn.SetReadDeadline)
  * [func (c *UDPConn) SetWriteBuffer(bytes int) error](#UDPConn.SetWriteBuffer)
  * [func (c *UDPConn) SetWriteDeadline(t time.Time) error](#UDPConn.SetWriteDeadline)
  * [func (c *UDPConn) SyscallConn() (syscall.RawConn, error)](#UDPConn.SyscallConn)
  * [func (c *UDPConn) Write(b []byte) (int, error)](#UDPConn.Write)
  * [func (c *UDPConn) WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error)](#UDPConn.WriteMsgUDP)
  * [func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error)](#UDPConn.WriteTo)
  * [func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)](#UDPConn.WriteToUDP)
* [type UnixAddr](#UnixAddr)
  * [func ResolveUnixAddr(network, address string) (*UnixAddr, error)](#ResolveUnixAddr)
  * [func (a *UnixAddr) Network() string](#UnixAddr.Network)
  * [func (a *UnixAddr) String() string](#UnixAddr.String)
* [type UnixConn](#UnixConn)
  * [func DialUnix(network string, laddr, raddr *UnixAddr) (*UnixConn, error)](#DialUnix)
  * [func ListenUnixgram(network string, laddr *UnixAddr) (*UnixConn, error)](#ListenUnixgram)
  * [func (c *UnixConn) Close() error](#UnixConn.Close)
  * [func (c *UnixConn) CloseRead() error](#UnixConn.CloseRead)
  * [func (c *UnixConn) CloseWrite() error](#UnixConn.CloseWrite)
  * [func (c *UnixConn) File() (f *os.File, err error)](#UnixConn.File)
  * [func (c *UnixConn) LocalAddr() Addr](#UnixConn.LocalAddr)
  * [func (c *UnixConn) Read(b []byte) (int, error)](#UnixConn.Read)
  * [func (c *UnixConn) ReadFrom(b []byte) (int, Addr, error)](#UnixConn.ReadFrom)
  * [func (c *UnixConn) ReadFromUnix(b []byte) (int, *UnixAddr, error)](#UnixConn.ReadFromUnix)
  * [func (c *UnixConn) ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error)](#UnixConn.ReadMsgUnix)
  * [func (c *UnixConn) RemoteAddr() Addr](#UnixConn.RemoteAddr)
  * [func (c *UnixConn) SetDeadline(t time.Time) error](#UnixConn.SetDeadline)
  * [func (c *UnixConn) SetReadBuffer(bytes int) error](#UnixConn.SetReadBuffer)
  * [func (c *UnixConn) SetReadDeadline(t time.Time) error](#UnixConn.SetReadDeadline)
  * [func (c *UnixConn) SetWriteBuffer(bytes int) error](#UnixConn.SetWriteBuffer)
  * [func (c *UnixConn) SetWriteDeadline(t time.Time) error](#UnixConn.SetWriteDeadline)
  * [func (c *UnixConn) SyscallConn() (syscall.RawConn, error)](#UnixConn.SyscallConn)
  * [func (c *UnixConn) Write(b []byte) (int, error)](#UnixConn.Write)
  * [func (c *UnixConn) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error)](#UnixConn.WriteMsgUnix)
  * [func (c *UnixConn) WriteTo(b []byte, addr Addr) (int, error)](#UnixConn.WriteTo)
  * [func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (int, error)](#UnixConn.WriteToUnix)
* [type UnixListener](#UnixListener)
  * [func ListenUnix(network string, laddr *UnixAddr) (*UnixListener, error)](#ListenUnix)
  * [func (l *UnixListener) Accept() (Conn, error)](#UnixListener.Accept)
  * [func (l *UnixListener) AcceptUnix() (*UnixConn, error)](#UnixListener.AcceptUnix)
  * [func (l *UnixListener) Addr() Addr](#UnixListener.Addr)
  * [func (l *UnixListener) Close() error](#UnixListener.Close)
  * [func (l *UnixListener) File() (f *os.File, err error)](#UnixListener.File)
  * [func (l *UnixListener) SetDeadline(t time.Time) error](#UnixListener.SetDeadline)
  * [func (l *UnixListener) SetUnlinkOnClose(unlink bool)](#UnixListener.SetUnlinkOnClose)
  * [func (l *UnixListener) SyscallConn() (syscall.RawConn, error)](#UnixListener.SyscallConn)
* [type UnknownNetworkError](#UnknownNetworkError)
  * [func (e UnknownNetworkError) Error() string](#UnknownNetworkError.Error)
  * [func (e UnknownNetworkError) Temporary() bool](#UnknownNetworkError.Temporary)
  * [func (e UnknownNetworkError) Timeout() bool](#UnknownNetworkError.Timeout)


#### <a id="pkg-examples">Examples</a>
* [CIDRMask](#example_CIDRMask)
* [IP.DefaultMask](#example_IP_DefaultMask)
* [IP.Mask](#example_IP_Mask)
* [IPv4](#example_IPv4)
* [IPv4Mask](#example_IPv4Mask)
* [Listener](#example_Listener)
* [ParseCIDR](#example_ParseCIDR)
* [ParseIP](#example_ParseIP)
* [UDPConn.WriteTo](#example_UDPConn_WriteTo)


#### <a id="pkg-files">Package files</a>
[addrselect.go](https://golang.org/src/net/addrselect.go) [cgo_linux.go](https://golang.org/src/net/cgo_linux.go) [cgo_resnew.go](https://golang.org/src/net/cgo_resnew.go) [cgo_socknew.go](https://golang.org/src/net/cgo_socknew.go) [cgo_unix.go](https://golang.org/src/net/cgo_unix.go) [conf.go](https://golang.org/src/net/conf.go) [dial.go](https://golang.org/src/net/dial.go) [dnsclient.go](https://golang.org/src/net/dnsclient.go) [dnsclient_unix.go](https://golang.org/src/net/dnsclient_unix.go) [dnsconfig_unix.go](https://golang.org/src/net/dnsconfig_unix.go) [error_posix.go](https://golang.org/src/net/error_posix.go) [error_unix.go](https://golang.org/src/net/error_unix.go) [fd_unix.go](https://golang.org/src/net/fd_unix.go) [file.go](https://golang.org/src/net/file.go) [file_unix.go](https://golang.org/src/net/file_unix.go) [hook.go](https://golang.org/src/net/hook.go) [hook_unix.go](https://golang.org/src/net/hook_unix.go) [hosts.go](https://golang.org/src/net/hosts.go) [interface.go](https://golang.org/src/net/interface.go) [interface_linux.go](https://golang.org/src/net/interface_linux.go) [ip.go](https://golang.org/src/net/ip.go) [iprawsock.go](https://golang.org/src/net/iprawsock.go) [iprawsock_posix.go](https://golang.org/src/net/iprawsock_posix.go) [ipsock.go](https://golang.org/src/net/ipsock.go) [ipsock_posix.go](https://golang.org/src/net/ipsock_posix.go) [lookup.go](https://golang.org/src/net/lookup.go) [lookup_unix.go](https://golang.org/src/net/lookup_unix.go) [mac.go](https://golang.org/src/net/mac.go) [net.go](https://golang.org/src/net/net.go) [nss.go](https://golang.org/src/net/nss.go) [parse.go](https://golang.org/src/net/parse.go) [pipe.go](https://golang.org/src/net/pipe.go) [port.go](https://golang.org/src/net/port.go) [port_unix.go](https://golang.org/src/net/port_unix.go) [rawconn.go](https://golang.org/src/net/rawconn.go) [sendfile_linux.go](https://golang.org/src/net/sendfile_linux.go) [sock_cloexec.go](https://golang.org/src/net/sock_cloexec.go) [sock_linux.go](https://golang.org/src/net/sock_linux.go) [sock_posix.go](https://golang.org/src/net/sock_posix.go) [sockaddr_posix.go](https://golang.org/src/net/sockaddr_posix.go) [sockopt_linux.go](https://golang.org/src/net/sockopt_linux.go) [sockopt_posix.go](https://golang.org/src/net/sockopt_posix.go) [sockoptip_linux.go](https://golang.org/src/net/sockoptip_linux.go) [sockoptip_posix.go](https://golang.org/src/net/sockoptip_posix.go) [splice_linux.go](https://golang.org/src/net/splice_linux.go) [tcpsock.go](https://golang.org/src/net/tcpsock.go) [tcpsock_posix.go](https://golang.org/src/net/tcpsock_posix.go) [tcpsockopt_posix.go](https://golang.org/src/net/tcpsockopt_posix.go) [tcpsockopt_unix.go](https://golang.org/src/net/tcpsockopt_unix.go) [udpsock.go](https://golang.org/src/net/udpsock.go) [udpsock_posix.go](https://golang.org/src/net/udpsock_posix.go) [unixsock.go](https://golang.org/src/net/unixsock.go) [unixsock_posix.go](https://golang.org/src/net/unixsock_posix.go) [writev_unix.go](https://golang.org/src/net/writev_unix.go) 


## <a id="pkg-constants">Constants</a>
IP address lengths (bytes).


<pre>const (
    <span id="IPv4len">IPv4len</span> = 4
    <span id="IPv6len">IPv6len</span> = 16
)</pre>

## <a id="pkg-variables">Variables</a>
Well-known IPv4 addresses


<pre>var (
    <span id="IPv4bcast">IPv4bcast</span>     = <a href="#IPv4">IPv4</a>(255, 255, 255, 255) <span class="comment">// limited broadcast</span>
    <span id="IPv4allsys">IPv4allsys</span>    = <a href="#IPv4">IPv4</a>(224, 0, 0, 1)       <span class="comment">// all systems</span>
    <span id="IPv4allrouter">IPv4allrouter</span> = <a href="#IPv4">IPv4</a>(224, 0, 0, 2)       <span class="comment">// all routers</span>
    <span id="IPv4zero">IPv4zero</span>      = <a href="#IPv4">IPv4</a>(0, 0, 0, 0)         <span class="comment">// all zeros</span>
)</pre>Well-known IPv6 addresses


<pre>var (
    <span id="IPv6zero">IPv6zero</span>                   = <a href="#IP">IP</a>{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    <span id="IPv6unspecified">IPv6unspecified</span>            = <a href="#IP">IP</a>{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    <span id="IPv6loopback">IPv6loopback</span>               = <a href="#IP">IP</a>{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
    <span id="IPv6interfacelocalallnodes">IPv6interfacelocalallnodes</span> = <a href="#IP">IP</a>{0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
    <span id="IPv6linklocalallnodes">IPv6linklocalallnodes</span>      = <a href="#IP">IP</a>{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
    <span id="IPv6linklocalallrouters">IPv6linklocalallrouters</span>    = <a href="#IP">IP</a>{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x02}
)</pre>DefaultResolver is the resolver used by the package-level Lookup
functions and by Dialers without a specified Resolver.


<pre>var <span id="DefaultResolver">DefaultResolver</span> = &amp;<a href="#Resolver">Resolver</a>{}</pre>Various errors contained in OpError.


<pre>var (
    <span id="ErrWriteToConnected">ErrWriteToConnected</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;use of WriteTo with pre-connected connection&#34;)
)</pre>

## <a id="JoinHostPort">func</a> [JoinHostPort](https://golang.org/src/net/ipsock.go?s=6810:6853#L217)
<pre>func JoinHostPort(host, port <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
JoinHostPort combines host and port into a network address of the
form "host:port". If host contains a colon, as found in literal
IPv6 addresses, then JoinHostPort returns "[host]:port".

See func Dial for a description of the host and port parameters.



## <a id="LookupAddr">func</a> [LookupAddr](https://golang.org/src/net/lookup.go?s=15484:15540#L437)
<pre>func LookupAddr(addr <a href="/pkg/builtin/#string">string</a>) (names []<a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
LookupAddr performs a reverse lookup for the given address, returning a list
of names mapping to that address.

When using the host C library resolver, at most one result will be
returned. To bypass the host resolver, use a custom Resolver.



## <a id="LookupCNAME">func</a> [LookupCNAME](https://golang.org/src/net/lookup.go?s=12081:12136#L358)
<pre>func LookupCNAME(host <a href="/pkg/builtin/#string">string</a>) (cname <a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
LookupCNAME returns the canonical name for the given host.
Callers that do not care about the canonical name can call
LookupHost or LookupIP directly; both take care of resolving
the canonical name as part of the lookup.

A canonical name is the final name after following zero
or more CNAME records.
LookupCNAME does not return an error if host does not
contain DNS "CNAME" records, as long as host resolves to
address records.



## <a id="LookupHost">func</a> [LookupHost](https://golang.org/src/net/lookup.go?s=5398:5454#L160)
<pre>func LookupHost(host <a href="/pkg/builtin/#string">string</a>) (addrs []<a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
LookupHost looks up the given host using the local resolver.
It returns a slice of that host's addresses.



## <a id="LookupPort">func</a> [LookupPort](https://golang.org/src/net/lookup.go?s=10817:10879#L322)
<pre>func LookupPort(network, service <a href="/pkg/builtin/#string">string</a>) (port <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
LookupPort looks up the port for the given network and service.



## <a id="LookupTXT">func</a> [LookupTXT](https://golang.org/src/net/lookup.go?s=14931:14976#L423)
<pre>func LookupTXT(name <a href="/pkg/builtin/#string">string</a>) ([]<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
LookupTXT returns the DNS TXT records for the given domain name.



## <a id="ParseCIDR">func</a> [ParseCIDR](https://golang.org/src/net/ip.go?s=16363:16407#L699)
<pre>func ParseCIDR(s <a href="/pkg/builtin/#string">string</a>) (<a href="#IP">IP</a>, *<a href="#IPNet">IPNet</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseCIDR parses s as a CIDR notation IP address and prefix length,
like "192.0.2.0/24" or "2001:db8::/32", as defined in
RFC 4632 and RFC 4291.

It returns the IP address and the network implied by the IP and
prefix length.
For example, ParseCIDR("192.0.2.1/24") returns the IP address
192.0.2.1 and the network 192.0.2.0/24.


<a id="example_ParseCIDR">Example</a>
```go
```

output:
```txt
```

## <a id="Pipe">func</a> [Pipe](https://golang.org/src/net/pipe.go?s=2871:2895#L108)
<pre>func Pipe() (<a href="#Conn">Conn</a>, <a href="#Conn">Conn</a>)</pre>
Pipe creates a synchronous, in-memory, full duplex
network connection; both ends implement the Conn interface.
Reads on one end are matched with writes on the other,
copying data directly between the two; there is no internal
buffering.



## <a id="SplitHostPort">func</a> [SplitHostPort](https://golang.org/src/net/ipsock.go?s=4766:4832#L146)
<pre>func SplitHostPort(hostport <a href="/pkg/builtin/#string">string</a>) (host, port <a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
SplitHostPort splits a network address of the form "host:port",
"host%zone:port", "[host]:port" or "[host%zone]:port" into host or
host%zone and port.

A literal IPv6 address in hostport must be enclosed in square
brackets, as in "[::1]:80", "[::1%lo0]:80".

See func Dial for a description of the hostport parameter, and host
and port results.





## <a id="Addr">type</a> [Addr](https://golang.org/src/net/net.go?s=3689:3875#L95)
Addr represents a network end point address.

The two methods Network and String conventionally return strings
that can be passed as the arguments to Dial, but the exact form
and meaning of the strings is up to the implementation.


<pre>type Addr interface {
    Network() <a href="/pkg/builtin/#string">string</a> <span class="comment">// name of the network (for example, &#34;tcp&#34;, &#34;udp&#34;)</span>
    String() <a href="/pkg/builtin/#string">string</a>  <span class="comment">// string form of address (for example, &#34;192.0.2.1:25&#34;, &#34;[2001:db8::1]:80&#34;)</span>
}</pre>









### <a id="InterfaceAddrs">func</a> [InterfaceAddrs](https://golang.org/src/net/interface.go?s=3418:3455#L105)
<pre>func InterfaceAddrs() ([]<a href="#Addr">Addr</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
InterfaceAddrs returns a list of the system's unicast interface
addresses.

The returned list does not identify the associated interface; use
Interfaces and Interface.Addrs for more detail.






## <a id="AddrError">type</a> [AddrError](https://golang.org/src/net/net.go?s=16776:16827#L528)

<pre>type AddrError struct {
<span id="AddrError.Err"></span>    Err  <a href="/pkg/builtin/#string">string</a>
<span id="AddrError.Addr"></span>    Addr <a href="/pkg/builtin/#string">string</a>
}
</pre>











### <a id="AddrError.Error">func</a> (\*AddrError) [Error](https://golang.org/src/net/net.go?s=16829:16863#L533)
<pre>func (e *<a href="#AddrError">AddrError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="AddrError.Temporary">func</a> (\*AddrError) [Temporary](https://golang.org/src/net/net.go?s=17039:17075#L545)
<pre>func (e *<a href="#AddrError">AddrError</a>) Temporary() <a href="/pkg/builtin/#bool">bool</a></pre>



### <a id="AddrError.Timeout">func</a> (\*AddrError) [Timeout](https://golang.org/src/net/net.go?s=16985:17019#L544)
<pre>func (e *<a href="#AddrError">AddrError</a>) Timeout() <a href="/pkg/builtin/#bool">bool</a></pre>



## <a id="Buffers">type</a> [Buffers](https://golang.org/src/net/net.go?s=20589:20610#L652)
Buffers contains zero or more runs of bytes to write.

On certain machines, for certain types of connections, this is
optimized into an OS-specific batch write operation (such as
"writev").


<pre>type Buffers [][]<a href="/pkg/builtin/#byte">byte</a></pre>











### <a id="Buffers.Read">func</a> (\*Buffers) [Read](https://golang.org/src/net/net.go?s=20973:21024#L675)
<pre>func (v *<a href="#Buffers">Buffers</a>) Read(p []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>



### <a id="Buffers.WriteTo">func</a> (\*Buffers) [WriteTo](https://golang.org/src/net/net.go?s=20687:20746#L659)
<pre>func (v *<a href="#Buffers">Buffers</a>) WriteTo(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) (n <a href="/pkg/builtin/#int64">int64</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>



## <a id="Conn">type</a> [Conn](https://golang.org/src/net/net.go?s=4005:6308#L103)
Conn is a generic stream-oriented network connection.

Multiple goroutines may invoke methods on a Conn simultaneously.


<pre>type Conn interface {
    <span class="comment">// Read reads data from the connection.</span>
    <span class="comment">// Read can be made to time out and return an Error with Timeout() == true</span>
    <span class="comment">// after a fixed time limit; see SetDeadline and SetReadDeadline.</span>
    Read(b []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)

    <span class="comment">// Write writes data to the connection.</span>
    <span class="comment">// Write can be made to time out and return an Error with Timeout() == true</span>
    <span class="comment">// after a fixed time limit; see SetDeadline and SetWriteDeadline.</span>
    Write(b []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)

    <span class="comment">// Close closes the connection.</span>
    <span class="comment">// Any blocked Read or Write operations will be unblocked and return errors.</span>
    Close() <a href="/pkg/builtin/#error">error</a>

    <span class="comment">// LocalAddr returns the local network address.</span>
    LocalAddr() <a href="#Addr">Addr</a>

    <span class="comment">// RemoteAddr returns the remote network address.</span>
    RemoteAddr() <a href="#Addr">Addr</a>

    <span class="comment">// SetDeadline sets the read and write deadlines associated</span>
    <span class="comment">// with the connection. It is equivalent to calling both</span>
    <span class="comment">// SetReadDeadline and SetWriteDeadline.</span>
    <span class="comment">//</span>
    <span class="comment">// A deadline is an absolute time after which I/O operations</span>
    <span class="comment">// fail with a timeout (see type Error) instead of</span>
    <span class="comment">// blocking. The deadline applies to all future and pending</span>
    <span class="comment">// I/O, not just the immediately following call to Read or</span>
    <span class="comment">// Write. After a deadline has been exceeded, the connection</span>
    <span class="comment">// can be refreshed by setting a deadline in the future.</span>
    <span class="comment">//</span>
    <span class="comment">// An idle timeout can be implemented by repeatedly extending</span>
    <span class="comment">// the deadline after successful Read or Write calls.</span>
    <span class="comment">//</span>
    <span class="comment">// A zero value for t means I/O operations will not time out.</span>
    <span class="comment">//</span>
    <span class="comment">// Note that if a TCP connection has keep-alive turned on,</span>
    <span class="comment">// which is the default unless overridden by Dialer.KeepAlive</span>
    <span class="comment">// or ListenConfig.KeepAlive, then a keep-alive failure may</span>
    <span class="comment">// also return a timeout error. On Unix systems a keep-alive</span>
    <span class="comment">// failure on I/O can be detected using</span>
    <span class="comment">// errors.Is(err, syscall.ETIMEDOUT).</span>
    SetDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a>

    <span class="comment">// SetReadDeadline sets the deadline for future Read calls</span>
    <span class="comment">// and any currently-blocked Read call.</span>
    <span class="comment">// A zero value for t means Read will not time out.</span>
    SetReadDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a>

    <span class="comment">// SetWriteDeadline sets the deadline for future Write calls</span>
    <span class="comment">// and any currently-blocked Write call.</span>
    <span class="comment">// Even if write times out, it may return n &gt; 0, indicating that</span>
    <span class="comment">// some of the data was successfully written.</span>
    <span class="comment">// A zero value for t means Write will not time out.</span>
    SetWriteDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a>
}</pre>









### <a id="Dial">func</a> [Dial](https://golang.org/src/net/dial.go?s=9799:9847#L306)
<pre>func Dial(network, address <a href="/pkg/builtin/#string">string</a>) (<a href="#Conn">Conn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Dial connects to the address on the named network.

Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only),
"udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4"
(IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and
"unixpacket".

For TCP and UDP networks, the address has the form "host:port".
The host must be a literal IP address, or a host name that can be
resolved to IP addresses.
The port must be a literal port number or a service name.
If the host is a literal IPv6 address it must be enclosed in square
brackets, as in "[2001:db8::1]:80" or "[fe80::1%zone]:80".
The zone specifies the scope of the literal IPv6 address as defined
in RFC 4007.
The functions JoinHostPort and SplitHostPort manipulate a pair of
host and port in this form.
When using TCP, and the host resolves to multiple IP addresses,
Dial will try each IP address in order until one succeeds.

Examples:


	Dial("tcp", "golang.org:http")
	Dial("tcp", "192.0.2.1:http")
	Dial("tcp", "198.51.100.1:80")
	Dial("udp", "[2001:db8::1]:domain")
	Dial("udp", "[fe80::1%lo0]:53")
	Dial("tcp", ":80")

For IP networks, the network must be "ip", "ip4" or "ip6" followed
by a colon and a literal protocol number or a protocol name, and
the address has the form "host". The host must be a literal IP
address or a literal IPv6 address with zone.
It depends on each operating system how the operating system
behaves with a non-well known protocol number such as "0" or "255".

Examples:


	Dial("ip4:1", "192.0.2.1")
	Dial("ip6:ipv6-icmp", "2001:db8::1")
	Dial("ip6:58", "fe80::1%lo0")

For TCP, UDP and IP networks, if the host is empty or a literal
unspecified IP address, as in ":80", "0.0.0.0:80" or "[::]:80" for
TCP and UDP, "", "0.0.0.0" or "::" for IP, the local system is
assumed.

For Unix networks, the address must be a file system path.




### <a id="DialTimeout">func</a> [DialTimeout](https://golang.org/src/net/dial.go?s=10311:10389#L321)
<pre>func DialTimeout(network, address <a href="/pkg/builtin/#string">string</a>, timeout <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>) (<a href="#Conn">Conn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DialTimeout acts like Dial but takes a timeout.

The timeout includes name resolution, if required.
When using TCP, and the host in the address parameter resolves to
multiple IP addresses, the timeout is spread over each consecutive
dial, such that each is given an appropriate fraction of the time
to connect.

See func Dial for a description of the network and address
parameters.




### <a id="FileConn">func</a> [FileConn](https://golang.org/src/net/file.go?s=659:704#L11)
<pre>func FileConn(f *<a href="/pkg/os/">os</a>.<a href="/pkg/os/#File">File</a>) (c <a href="#Conn">Conn</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
FileConn returns a copy of the network connection corresponding to
the open file f.
It is the caller's responsibility to close f when finished.
Closing c does not affect f, and closing f does not affect c.






## <a id="DNSConfigError">type</a> [DNSConfigError](https://golang.org/src/net/net.go?s=17681:17722#L561)
DNSConfigError represents an error reading the machine's DNS configuration.
(No longer used; kept for compatibility.)


<pre>type DNSConfigError struct {
<span id="DNSConfigError.Err"></span>    Err <a href="/pkg/builtin/#error">error</a>
}
</pre>











### <a id="DNSConfigError.Error">func</a> (\*DNSConfigError) [Error](https://golang.org/src/net/net.go?s=17783:17822#L566)
<pre>func (e *<a href="#DNSConfigError">DNSConfigError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="DNSConfigError.Temporary">func</a> (\*DNSConfigError) [Temporary](https://golang.org/src/net/net.go?s=17940:17981#L568)
<pre>func (e *<a href="#DNSConfigError">DNSConfigError</a>) Temporary() <a href="/pkg/builtin/#bool">bool</a></pre>



### <a id="DNSConfigError.Timeout">func</a> (\*DNSConfigError) [Timeout](https://golang.org/src/net/net.go?s=17881:17920#L567)
<pre>func (e *<a href="#DNSConfigError">DNSConfigError</a>) Timeout() <a href="/pkg/builtin/#bool">bool</a></pre>



### <a id="DNSConfigError.Unwrap">func</a> (\*DNSConfigError) [Unwrap](https://golang.org/src/net/net.go?s=17724:17763#L565)
<pre>func (e *<a href="#DNSConfigError">DNSConfigError</a>) Unwrap() <a href="/pkg/builtin/#error">error</a></pre>



## <a id="DNSError">type</a> [DNSError](https://golang.org/src/net/net.go?s=18137:18484#L576)
DNSError represents a DNS lookup error.


<pre>type DNSError struct {
<span id="DNSError.Err"></span>    Err         <a href="/pkg/builtin/#string">string</a> <span class="comment">// description of the error</span>
<span id="DNSError.Name"></span>    Name        <a href="/pkg/builtin/#string">string</a> <span class="comment">// name looked for</span>
<span id="DNSError.Server"></span>    Server      <a href="/pkg/builtin/#string">string</a> <span class="comment">// server used</span>
<span id="DNSError.IsTimeout"></span>    IsTimeout   <a href="/pkg/builtin/#bool">bool</a>   <span class="comment">// if true, timed out; not all timeouts set this</span>
<span id="DNSError.IsTemporary"></span>    IsTemporary <a href="/pkg/builtin/#bool">bool</a>   <span class="comment">// if true, error is temporary; not all errors set this</span>
<span id="DNSError.IsNotFound"></span>    IsNotFound  <a href="/pkg/builtin/#bool">bool</a>   <span class="comment">// if true, host could not be found</span>
}
</pre>











### <a id="DNSError.Error">func</a> (\*DNSError) [Error](https://golang.org/src/net/net.go?s=18486:18519#L585)
<pre>func (e *<a href="#DNSError">DNSError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="DNSError.Temporary">func</a> (\*DNSError) [Temporary](https://golang.org/src/net/net.go?s=19122:19157#L605)
<pre>func (e *<a href="#DNSError">DNSError</a>) Temporary() <a href="/pkg/builtin/#bool">bool</a></pre>
Temporary reports whether the DNS error is known to be temporary.
This is not always known; a DNS lookup may fail due to a temporary
error and return a DNSError for which Temporary returns false.




### <a id="DNSError.Timeout">func</a> (\*DNSError) [Timeout](https://golang.org/src/net/net.go?s=18859:18892#L600)
<pre>func (e *<a href="#DNSError">DNSError</a>) Timeout() <a href="/pkg/builtin/#bool">bool</a></pre>
Timeout reports whether the DNS lookup is known to have timed out.
This is not always known; a DNS lookup may fail due to a timeout
and return a DNSError for which Timeout returns false.




## <a id="Dialer">type</a> [Dialer](https://golang.org/src/net/dial.go?s=652:3413#L16)
A Dialer contains options for connecting to an address.

The zero value for each field is equivalent to dialing
without that option. Dialing with the zero value of Dialer
is therefore equivalent to just calling the Dial function.


<pre>type Dialer struct {
<span id="Dialer.Timeout"></span>    <span class="comment">// Timeout is the maximum amount of time a dial will wait for</span>
    <span class="comment">// a connect to complete. If Deadline is also set, it may fail</span>
    <span class="comment">// earlier.</span>
    <span class="comment">//</span>
    <span class="comment">// The default is no timeout.</span>
    <span class="comment">//</span>
    <span class="comment">// When using TCP and dialing a host name with multiple IP</span>
    <span class="comment">// addresses, the timeout may be divided between them.</span>
    <span class="comment">//</span>
    <span class="comment">// With or without a timeout, the operating system may impose</span>
    <span class="comment">// its own earlier timeout. For instance, TCP timeouts are</span>
    <span class="comment">// often around 3 minutes.</span>
    Timeout <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>

<span id="Dialer.Deadline"></span>    <span class="comment">// Deadline is the absolute point in time after which dials</span>
    <span class="comment">// will fail. If Timeout is set, it may fail earlier.</span>
    <span class="comment">// Zero means no deadline, or dependent on the operating system</span>
    <span class="comment">// as with the Timeout option.</span>
    Deadline <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>

<span id="Dialer.LocalAddr"></span>    <span class="comment">// LocalAddr is the local address to use when dialing an</span>
    <span class="comment">// address. The address must be of a compatible type for the</span>
    <span class="comment">// network being dialed.</span>
    <span class="comment">// If nil, a local address is automatically chosen.</span>
    LocalAddr <a href="#Addr">Addr</a>

<span id="Dialer.DualStack"></span>    <span class="comment">// DualStack previously enabled RFC 6555 Fast Fallback</span>
    <span class="comment">// support, also known as &#34;Happy Eyeballs&#34;, in which IPv4 is</span>
    <span class="comment">// tried soon if IPv6 appears to be misconfigured and</span>
    <span class="comment">// hanging.</span>
    <span class="comment">//</span>
    <span class="comment">// Deprecated: Fast Fallback is enabled by default. To</span>
    <span class="comment">// disable, set FallbackDelay to a negative value.</span>
    DualStack <a href="/pkg/builtin/#bool">bool</a>

<span id="Dialer.FallbackDelay"></span>    <span class="comment">// FallbackDelay specifies the length of time to wait before</span>
    <span class="comment">// spawning a RFC 6555 Fast Fallback connection. That is, this</span>
    <span class="comment">// is the amount of time to wait for IPv6 to succeed before</span>
    <span class="comment">// assuming that IPv6 is misconfigured and falling back to</span>
    <span class="comment">// IPv4.</span>
    <span class="comment">//</span>
    <span class="comment">// If zero, a default delay of 300ms is used.</span>
    <span class="comment">// A negative value disables Fast Fallback support.</span>
    FallbackDelay <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>

<span id="Dialer.KeepAlive"></span>    <span class="comment">// KeepAlive specifies the interval between keep-alive</span>
    <span class="comment">// probes for an active network connection.</span>
    <span class="comment">// If zero, keep-alive probes are sent with a default value</span>
    <span class="comment">// (currently 15 seconds), if supported by the protocol and operating</span>
    <span class="comment">// system. Network protocols or operating systems that do</span>
    <span class="comment">// not support keep-alives ignore this field.</span>
    <span class="comment">// If negative, keep-alive probes are disabled.</span>
    KeepAlive <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>

<span id="Dialer.Resolver"></span>    <span class="comment">// Resolver optionally specifies an alternate resolver to use.</span>
    Resolver *<a href="#Resolver">Resolver</a>

<span id="Dialer.Cancel"></span>    <span class="comment">// Cancel is an optional channel whose closure indicates that</span>
    <span class="comment">// the dial should be canceled. Not all types of dials support</span>
    <span class="comment">// cancellation.</span>
    <span class="comment">//</span>
    <span class="comment">// Deprecated: Use DialContext instead.</span>
    Cancel &lt;-chan struct{}

    <span class="comment">// If Control is not nil, it is called after creating the network</span>
    <span class="comment">// connection but before actually dialing.</span>
    <span class="comment">//</span>
    <span class="comment">// Network and address parameters passed to Control method are not</span>
    <span class="comment">// necessarily the ones passed to Dial. For example, passing &#34;tcp&#34; to Dial</span>
    <span class="comment">// will cause the Control function to be called with &#34;tcp4&#34; or &#34;tcp6&#34;.</span>
<span id="Dialer.Control"></span>    Control func(network, address <a href="/pkg/builtin/#string">string</a>, c <a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#RawConn">RawConn</a>) <a href="/pkg/builtin/#error">error</a>
}
</pre>











### <a id="Dialer.Dial">func</a> (\*Dialer) [Dial](https://golang.org/src/net/dial.go?s=10714:10774#L336)
<pre>func (d *<a href="#Dialer">Dialer</a>) Dial(network, address <a href="/pkg/builtin/#string">string</a>) (<a href="#Conn">Conn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Dial connects to the address on the named network.

See func Dial for a description of the network and address
parameters.




### <a id="Dialer.DialContext">func</a> (\*Dialer) [DialContext](https://golang.org/src/net/dial.go?s=11677:11765#L358)
<pre>func (d *<a href="#Dialer">Dialer</a>) DialContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, network, address <a href="/pkg/builtin/#string">string</a>) (<a href="#Conn">Conn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DialContext connects to the address on the named network using
the provided context.

The provided Context must be non-nil. If the context expires before
the connection is complete, an error is returned. Once successfully
connected, any expiration of the context will not affect the
connection.

When using TCP, and the host in the address parameter resolves to multiple
network addresses, any dial timeout (from d.Timeout or ctx) is spread
over each consecutive dial, such that each is given an appropriate
fraction of the time to connect.
For example, if a host has 4 IP addresses and the timeout is 1 minute,
the connect to each single address will be given 15 seconds to complete
before trying the next one.

See func Dial for a description of the network and address
parameters.




## <a id="Error">type</a> [Error](https://golang.org/src/net/net.go?s=13212:13333#L384)
An Error represents a network error.


<pre>type Error interface {
    <a href="/pkg/builtin/#error">error</a>
    Timeout() <a href="/pkg/builtin/#bool">bool</a>   <span class="comment">// Is the error a timeout?</span>
    Temporary() <a href="/pkg/builtin/#bool">bool</a> <span class="comment">// Is the error temporary?</span>
}</pre>











## <a id="Flags">type</a> [Flags](https://golang.org/src/net/interface.go?s=1342:1357#L28)

<pre>type Flags <a href="/pkg/builtin/#uint">uint</a></pre>



<pre>const (
    <span id="FlagUp">FlagUp</span>           <a href="#Flags">Flags</a> = 1 &lt;&lt; <a href="/pkg/builtin/#iota">iota</a> <span class="comment">// interface is up</span>
    <span id="FlagBroadcast">FlagBroadcast</span>                      <span class="comment">// interface supports broadcast access capability</span>
    <span id="FlagLoopback">FlagLoopback</span>                       <span class="comment">// interface is a loopback interface</span>
    <span id="FlagPointToPoint">FlagPointToPoint</span>                   <span class="comment">// interface belongs to a point-to-point link</span>
    <span id="FlagMulticast">FlagMulticast</span>                      <span class="comment">// interface supports multicast access capability</span>
)</pre>









### <a id="Flags.String">func</a> (Flags) [String](https://golang.org/src/net/interface.go?s=1846:1876#L46)
<pre>func (f <a href="#Flags">Flags</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="HardwareAddr">type</a> [HardwareAddr](https://golang.org/src/net/mac.go?s=268:292#L1)
A HardwareAddr represents a physical hardware address.


<pre>type HardwareAddr []<a href="/pkg/builtin/#byte">byte</a></pre>









### <a id="ParseMAC">func</a> [ParseMAC](https://golang.org/src/net/mac.go?s=1035:1087#L28)
<pre>func ParseMAC(s <a href="/pkg/builtin/#string">string</a>) (hw <a href="#HardwareAddr">HardwareAddr</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ParseMAC parses s as an IEEE 802 MAC-48, EUI-48, EUI-64, or a 20-octet
IP over InfiniBand link-layer address using one of the following formats:


	00:00:5e:00:53:01
	02:00:5e:10:00:00:00:01
	00:00:00:00:fe:80:00:00:00:00:00:00:02:00:5e:10:00:00:00:01
	00-00-5e-00-53-01
	02-00-5e-10-00-00-00-01
	00-00-00-00-fe-80-00-00-00-00-00-00-02-00-5e-10-00-00-00-01
	0000.5e00.5301
	0200.5e10.0000.0001
	0000.0000.fe80.0000.0000.0000.0200.5e10.0000.0001






### <a id="HardwareAddr.String">func</a> (HardwareAddr) [String](https://golang.org/src/net/mac.go?s=294:331#L2)
<pre>func (a <a href="#HardwareAddr">HardwareAddr</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="IP">type</a> [IP](https://golang.org/src/net/ip.go?s=946:960#L22)
An IP is a single IP address, a slice of bytes.
Functions in this package accept either 4-byte (IPv4)
or 16-byte (IPv6) slices as input.

Note that in this documentation, referring to an
IP address as an IPv4 address or an IPv6 address
is a semantic property of the address, not just the
length of the byte slice: a 16-byte slice can still
be an IPv4 address.


<pre>type IP []<a href="/pkg/builtin/#byte">byte</a></pre>









### <a id="IPv4">func</a> [IPv4](https://golang.org/src/net/ip.go?s=1216:1245#L35)
<pre>func IPv4(a, b, c, d <a href="/pkg/builtin/#byte">byte</a>) <a href="#IP">IP</a></pre>
IPv4 returns the IP address (in 16-byte form) of the
IPv4 address a.b.c.d.


<a id="example_IPv4">Example</a>
```go
```

output:
```txt
```


### <a id="LookupIP">func</a> [LookupIP](https://golang.org/src/net/lookup.go?s=6178:6218#L180)
<pre>func LookupIP(host <a href="/pkg/builtin/#string">string</a>) ([]<a href="#IP">IP</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
LookupIP looks up host using the local resolver.
It returns a slice of that host's IPv4 and IPv6 addresses.




### <a id="ParseIP">func</a> [ParseIP](https://golang.org/src/net/ip.go?s=15550:15575#L665)
<pre>func ParseIP(s <a href="/pkg/builtin/#string">string</a>) <a href="#IP">IP</a></pre>
ParseIP parses s as an IP address, returning the result.
The string s can be in dotted decimal ("192.0.2.1")
or IPv6 ("2001:db8::68") form.
If s is not a valid textual representation of an IP address,
ParseIP returns nil.


<a id="example_ParseIP">Example</a>
```go
```

output:
```txt
```




### <a id="IP.DefaultMask">func</a> (IP) [DefaultMask](https://golang.org/src/net/ip.go?s=5963:5996#L211)
<pre>func (ip <a href="#IP">IP</a>) DefaultMask() <a href="#IPMask">IPMask</a></pre>
DefaultMask returns the default IP mask for the IP address ip.
Only IPv4 addresses have default masks; DefaultMask returns
nil if ip is not a valid IPv4 address.


<a id="example_IP_DefaultMask">Example</a>
```go
```

output:
```txt
```


### <a id="IP.Equal">func</a> (IP) [Equal](https://golang.org/src/net/ip.go?s=10028:10057#L397)
<pre>func (ip <a href="#IP">IP</a>) Equal(x <a href="#IP">IP</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Equal reports whether ip and x are the same IP address.
An IPv4 address and that same address in IPv6 form are
considered to be equal.




### <a id="IP.IsGlobalUnicast">func</a> (IP) [IsGlobalUnicast](https://golang.org/src/net/ip.go?s=4686:4721#L155)
<pre>func (ip <a href="#IP">IP</a>) IsGlobalUnicast() <a href="/pkg/builtin/#bool">bool</a></pre>
IsGlobalUnicast reports whether ip is a global unicast
address.

The identification of global unicast addresses uses address type
identification as defined in RFC 1122, RFC 4632 and RFC 4291 with
the exception of IPv4 directed broadcast addresses.
It returns true even if ip is in IPv4 private address space or
local IPv6 unicast address space.




### <a id="IP.IsInterfaceLocalMulticast">func</a> (IP) [IsInterfaceLocalMulticast](https://golang.org/src/net/ip.go?s=3654:3699#L125)
<pre>func (ip <a href="#IP">IP</a>) IsInterfaceLocalMulticast() <a href="/pkg/builtin/#bool">bool</a></pre>
IsInterfaceLocalMulticast reports whether ip is
an interface-local multicast address.




### <a id="IP.IsLinkLocalMulticast">func</a> (IP) [IsLinkLocalMulticast](https://golang.org/src/net/ip.go?s=3852:3892#L131)
<pre>func (ip <a href="#IP">IP</a>) IsLinkLocalMulticast() <a href="/pkg/builtin/#bool">bool</a></pre>
IsLinkLocalMulticast reports whether ip is a link-local
multicast address.




### <a id="IP.IsLinkLocalUnicast">func</a> (IP) [IsLinkLocalUnicast](https://golang.org/src/net/ip.go?s=4131:4169#L140)
<pre>func (ip <a href="#IP">IP</a>) IsLinkLocalUnicast() <a href="/pkg/builtin/#bool">bool</a></pre>
IsLinkLocalUnicast reports whether ip is a link-local
unicast address.




### <a id="IP.IsLoopback">func</a> (IP) [IsLoopback](https://golang.org/src/net/ip.go?s=3230:3260#L108)
<pre>func (ip <a href="#IP">IP</a>) IsLoopback() <a href="/pkg/builtin/#bool">bool</a></pre>
IsLoopback reports whether ip is a loopback address.




### <a id="IP.IsMulticast">func</a> (IP) [IsMulticast](https://golang.org/src/net/ip.go?s=3415:3446#L116)
<pre>func (ip <a href="#IP">IP</a>) IsMulticast() <a href="/pkg/builtin/#bool">bool</a></pre>
IsMulticast reports whether ip is a multicast address.




### <a id="IP.IsUnspecified">func</a> (IP) [IsUnspecified](https://golang.org/src/net/ip.go?s=3079:3112#L103)
<pre>func (ip <a href="#IP">IP</a>) IsUnspecified() <a href="/pkg/builtin/#bool">bool</a></pre>
IsUnspecified reports whether ip is an unspecified address, either
the IPv4 address "0.0.0.0" or the IPv6 address "::".




### <a id="IP.MarshalText">func</a> (IP) [MarshalText](https://golang.org/src/net/ip.go?s=9278:9320#L368)
<pre>func (ip <a href="#IP">IP</a>) MarshalText() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
MarshalText implements the encoding.TextMarshaler interface.
The encoding is the same as returned by String, with one exception:
When len(ip) is zero, it returns an empty slice.




### <a id="IP.Mask">func</a> (IP) [Mask](https://golang.org/src/net/ip.go?s=6346:6379#L235)
<pre>func (ip <a href="#IP">IP</a>) Mask(mask <a href="#IPMask">IPMask</a>) <a href="#IP">IP</a></pre>
Mask returns the result of masking the IP address ip with mask.


<a id="example_IP_Mask">Example</a>
```go
```

output:
```txt
```


### <a id="IP.String">func</a> (IP) [String](https://golang.org/src/net/ip.go?s=7520:7548#L278)
<pre>func (ip <a href="#IP">IP</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns the string form of the IP address ip.
It returns one of 4 forms:


	- "<nil>", if ip has length 0
	- dotted decimal ("192.0.2.1"), if ip is an IPv4 or IP4-mapped IPv6 address
	- IPv6 ("2001:db8::1"), if ip is a valid IPv6 address
	- the hexadecimal form of ip, without punctuation, if no other cases apply




### <a id="IP.To16">func</a> (IP) [To16](https://golang.org/src/net/ip.go?s=5477:5499#L191)
<pre>func (ip <a href="#IP">IP</a>) To16() <a href="#IP">IP</a></pre>
To16 converts the IP address ip to a 16-byte representation.
If ip is not an IP address (it is the wrong length), To16 returns nil.




### <a id="IP.To4">func</a> (IP) [To4](https://golang.org/src/net/ip.go?s=5150:5171#L176)
<pre>func (ip <a href="#IP">IP</a>) To4() <a href="#IP">IP</a></pre>
To4 converts the IPv4 address ip to a 4-byte representation.
If ip is not an IPv4 address, To4 returns nil.




### <a id="IP.UnmarshalText">func</a> (\*IP) [UnmarshalText](https://golang.org/src/net/ip.go?s=9658:9704#L380)
<pre>func (ip *<a href="#IP">IP</a>) UnmarshalText(text []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
UnmarshalText implements the encoding.TextUnmarshaler interface.
The IP address is expected in a form accepted by ParseIP.




## <a id="IPAddr">type</a> [IPAddr](https://golang.org/src/net/iprawsock.go?s=1009:1084#L21)
IPAddr represents the address of an IP end point.


<pre>type IPAddr struct {
<span id="IPAddr.IP"></span>    IP   <a href="#IP">IP</a>
<span id="IPAddr.Zone"></span>    Zone <a href="/pkg/builtin/#string">string</a> <span class="comment">// IPv6 scoped addressing zone</span>
}
</pre>









### <a id="ResolveIPAddr">func</a> [ResolveIPAddr](https://golang.org/src/net/iprawsock.go?s=2073:2133#L67)
<pre>func ResolveIPAddr(network, address <a href="/pkg/builtin/#string">string</a>) (*<a href="#IPAddr">IPAddr</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ResolveIPAddr returns an address of IP end point.

The network must be an IP network name.

If the host in the address parameter is not a literal IP address,
ResolveIPAddr resolves the address to an address of IP end point.
Otherwise, it parses the address as a literal IP address.
The address parameter can use a host name, but this is not
recommended, because it will return at most one of the host name's
IP addresses.

See func Dial for a description of the network and address
parameters.






### <a id="IPAddr.Network">func</a> (\*IPAddr) [Network](https://golang.org/src/net/iprawsock.go?s=1139:1172#L27)
<pre>func (a *<a href="#IPAddr">IPAddr</a>) Network() <a href="/pkg/builtin/#string">string</a></pre>
Network returns the address's network name, "ip".




### <a id="IPAddr.String">func</a> (\*IPAddr) [String](https://golang.org/src/net/iprawsock.go?s=1190:1222#L29)
<pre>func (a *<a href="#IPAddr">IPAddr</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="IPConn">type</a> [IPConn](https://golang.org/src/net/iprawsock.go?s=2719:2747#L89)
IPConn is the implementation of the Conn and PacketConn interfaces
for IP network connections.


<pre>type IPConn struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="DialIP">func</a> [DialIP](https://golang.org/src/net/iprawsock.go?s=6142:6208#L201)
<pre>func DialIP(network <a href="/pkg/builtin/#string">string</a>, laddr, raddr *<a href="#IPAddr">IPAddr</a>) (*<a href="#IPConn">IPConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DialIP acts like Dial for IP networks.

The network must be an IP network name; see func Dial for details.

If laddr is nil, a local address is automatically chosen.
If the IP field of raddr is nil or an unspecified IP address, the
local system is assumed.




### <a id="ListenIP">func</a> [ListenIP](https://golang.org/src/net/iprawsock.go?s=6899:6960#L220)
<pre>func ListenIP(network <a href="/pkg/builtin/#string">string</a>, laddr *<a href="#IPAddr">IPAddr</a>) (*<a href="#IPConn">IPConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ListenIP acts like ListenPacket for IP networks.

The network must be an IP network name; see func Dial for details.

If the IP field of laddr is nil or an unspecified IP address,
ListenIP listens on all available IP addresses of the local system
except multicast IP addresses.






### <a id="IPConn.Close">func</a> (\*IPConn) [Close](https://golang.org/src/net/net.go?s=7068:7096#L194)
<pre>func (c *<a href="#IPConn">IPConn</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the connection.




### <a id="IPConn.File">func</a> (\*IPConn) [File](https://golang.org/src/net/net.go?s=9729:9774#L289)
<pre>func (c *<a href="#IPConn">IPConn</a>) File() (f *<a href="/pkg/os/">os</a>.<a href="/pkg/os/#File">File</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
File returns a copy of the underlying os.File.
It is the caller's responsibility to close f when finished.
Closing c does not affect f, and closing f does not affect c.

The returned os.File's file descriptor is different from the connection's.
Attempting to change properties of the original using this duplicate
may or may not have the desired effect.




### <a id="IPConn.LocalAddr">func</a> (\*IPConn) [LocalAddr](https://golang.org/src/net/net.go?s=7425:7456#L208)
<pre>func (c *<a href="#IPConn">IPConn</a>) LocalAddr() <a href="#Addr">Addr</a></pre>
LocalAddr returns the local network address.
The Addr returned is shared by all invocations of LocalAddr, so
do not modify it.




### <a id="IPConn.Read">func</a> (\*IPConn) [Read](https://golang.org/src/net/net.go?s=6487:6529#L170)
<pre>func (c *<a href="#IPConn">IPConn</a>) Read(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Read implements the Conn Read method.




### <a id="IPConn.ReadFrom">func</a> (\*IPConn) [ReadFrom](https://golang.org/src/net/iprawsock.go?s=3366:3420#L115)
<pre>func (c *<a href="#IPConn">IPConn</a>) ReadFrom(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="#Addr">Addr</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadFrom implements the PacketConn ReadFrom method.




### <a id="IPConn.ReadFromIP">func</a> (\*IPConn) [ReadFromIP](https://golang.org/src/net/iprawsock.go?s=3033:3092#L103)
<pre>func (c *<a href="#IPConn">IPConn</a>) ReadFromIP(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, *<a href="#IPAddr">IPAddr</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadFromIP acts like ReadFrom but returns an IPAddr.




### <a id="IPConn.ReadMsgIP">func</a> (\*IPConn) [ReadMsgIP](https://golang.org/src/net/iprawsock.go?s=4089:4176#L136)
<pre>func (c *<a href="#IPConn">IPConn</a>) ReadMsgIP(b, oob []<a href="/pkg/builtin/#byte">byte</a>) (n, oobn, flags <a href="/pkg/builtin/#int">int</a>, addr *<a href="#IPAddr">IPAddr</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadMsgIP reads a message from c, copying the payload into b and
the associated out-of-band data into oob. It returns the number of
bytes copied into b, the number of bytes copied into oob, the flags
that were set on the message and the source address of the message.

The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be
used to manipulate IP-level socket options in oob.




### <a id="IPConn.RemoteAddr">func</a> (\*IPConn) [RemoteAddr](https://golang.org/src/net/net.go?s=7650:7682#L218)
<pre>func (c *<a href="#IPConn">IPConn</a>) RemoteAddr() <a href="#Addr">Addr</a></pre>
RemoteAddr returns the remote network address.
The Addr returned is shared by all invocations of RemoteAddr, so
do not modify it.




### <a id="IPConn.SetDeadline">func</a> (\*IPConn) [SetDeadline](https://golang.org/src/net/net.go?s=7792:7837#L226)
<pre>func (c *<a href="#IPConn">IPConn</a>) SetDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetDeadline implements the Conn SetDeadline method.




### <a id="IPConn.SetReadBuffer">func</a> (\*IPConn) [SetReadBuffer](https://golang.org/src/net/net.go?s=8756:8801#L260)
<pre>func (c *<a href="#IPConn">IPConn</a>) SetReadBuffer(bytes <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetReadBuffer sets the size of the operating system's
receive buffer associated with the connection.




### <a id="IPConn.SetReadDeadline">func</a> (\*IPConn) [SetReadDeadline](https://golang.org/src/net/net.go?s=8092:8141#L237)
<pre>func (c *<a href="#IPConn">IPConn</a>) SetReadDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetReadDeadline implements the Conn SetReadDeadline method.




### <a id="IPConn.SetWriteBuffer">func</a> (\*IPConn) [SetWriteBuffer](https://golang.org/src/net/net.go?s=9109:9155#L272)
<pre>func (c *<a href="#IPConn">IPConn</a>) SetWriteBuffer(bytes <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetWriteBuffer sets the size of the operating system's
transmit buffer associated with the connection.




### <a id="IPConn.SetWriteDeadline">func</a> (\*IPConn) [SetWriteDeadline](https://golang.org/src/net/net.go?s=8402:8452#L248)
<pre>func (c *<a href="#IPConn">IPConn</a>) SetWriteDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetWriteDeadline implements the Conn SetWriteDeadline method.




### <a id="IPConn.SyscallConn">func</a> (\*IPConn) [SyscallConn](https://golang.org/src/net/iprawsock.go?s=2845:2900#L95)
<pre>func (c *<a href="#IPConn">IPConn</a>) SyscallConn() (<a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#RawConn">RawConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
SyscallConn returns a raw network connection.
This implements the syscall.Conn interface.




### <a id="IPConn.Write">func</a> (\*IPConn) [Write](https://golang.org/src/net/net.go?s=6790:6833#L182)
<pre>func (c *<a href="#IPConn">IPConn</a>) Write(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Write implements the Conn Write method.




### <a id="IPConn.WriteMsgIP">func</a> (\*IPConn) [WriteMsgIP](https://golang.org/src/net/iprawsock.go?s=5503:5584#L181)
<pre>func (c *<a href="#IPConn">IPConn</a>) WriteMsgIP(b, oob []<a href="/pkg/builtin/#byte">byte</a>, addr *<a href="#IPAddr">IPAddr</a>) (n, oobn <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
WriteMsgIP writes a message to addr via c, copying the payload from
b and the associated out-of-band data from oob. It returns the
number of payload and out-of-band bytes written.

The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be
used to manipulate IP-level socket options in oob.




### <a id="IPConn.WriteTo">func</a> (\*IPConn) [WriteTo](https://golang.org/src/net/iprawsock.go?s=4783:4841#L160)
<pre>func (c *<a href="#IPConn">IPConn</a>) WriteTo(b []<a href="/pkg/builtin/#byte">byte</a>, addr <a href="#Addr">Addr</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
WriteTo implements the PacketConn WriteTo method.




### <a id="IPConn.WriteToIP">func</a> (\*IPConn) [WriteToIP](https://golang.org/src/net/iprawsock.go?s=4456:4519#L148)
<pre>func (c *<a href="#IPConn">IPConn</a>) WriteToIP(b []<a href="/pkg/builtin/#byte">byte</a>, addr *<a href="#IPAddr">IPAddr</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
WriteToIP acts like WriteTo but takes an IPAddr.




## <a id="IPMask">type</a> [IPMask](https://golang.org/src/net/ip.go?s=994:1012#L25)
An IP mask is an IP address.


<pre>type IPMask []<a href="/pkg/builtin/#byte">byte</a></pre>









### <a id="CIDRMask">func</a> [CIDRMask](https://golang.org/src/net/ip.go?s=1801:1837#L61)
<pre>func CIDRMask(ones, bits <a href="/pkg/builtin/#int">int</a>) <a href="#IPMask">IPMask</a></pre>
CIDRMask returns an IPMask consisting of `ones' 1 bits
followed by 0s up to a total length of `bits' bits.
For a mask of this form, CIDRMask is the inverse of IPMask.Size.


<a id="example_CIDRMask">Example</a>
```go
```

output:
```txt
```


### <a id="IPv4Mask">func</a> [IPv4Mask](https://golang.org/src/net/ip.go?s=1499:1536#L49)
<pre>func IPv4Mask(a, b, c, d <a href="/pkg/builtin/#byte">byte</a>) <a href="#IPMask">IPMask</a></pre>
IPv4Mask returns the IP mask (in 4-byte form) of the
IPv4 mask a.b.c.d.


<a id="example_IPv4Mask">Example</a>
```go
```

output:
```txt
```




### <a id="IPMask.Size">func</a> (IPMask) [Size](https://golang.org/src/net/ip.go?s=11160:11199#L446)
<pre>func (m <a href="#IPMask">IPMask</a>) Size() (ones, bits <a href="/pkg/builtin/#int">int</a>)</pre>
Size returns the number of leading ones and total bits in the mask.
If the mask is not in the canonical form--ones followed by zeros--then
Size returns 0, 0.




### <a id="IPMask.String">func</a> (IPMask) [String](https://golang.org/src/net/ip.go?s=11357:11388#L455)
<pre>func (m <a href="#IPMask">IPMask</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns the hexadecimal form of m, with no punctuation.




## <a id="IPNet">type</a> [IPNet](https://golang.org/src/net/ip.go?s=1052:1133#L28)
An IPNet represents an IP network.


<pre>type IPNet struct {
<span id="IPNet.IP"></span>    IP   <a href="#IP">IP</a>     <span class="comment">// network number</span>
<span id="IPNet.Mask"></span>    Mask <a href="#IPMask">IPMask</a> <span class="comment">// network mask</span>
}
</pre>











### <a id="IPNet.Contains">func</a> (\*IPNet) [Contains](https://golang.org/src/net/ip.go?s=11853:11889#L486)
<pre>func (n *<a href="#IPNet">IPNet</a>) Contains(ip <a href="#IP">IP</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Contains reports whether the network includes ip.




### <a id="IPNet.Network">func</a> (\*IPNet) [Network](https://golang.org/src/net/ip.go?s=12173:12205#L504)
<pre>func (n *<a href="#IPNet">IPNet</a>) Network() <a href="/pkg/builtin/#string">string</a></pre>
Network returns the address's network name, "ip+net".




### <a id="IPNet.String">func</a> (\*IPNet) [String](https://golang.org/src/net/ip.go?s=12577:12608#L512)
<pre>func (n *<a href="#IPNet">IPNet</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns the CIDR notation of n like "192.0.2.0/24"
or "2001:db8::/48" as defined in RFC 4632 and RFC 4291.
If the mask is not in the canonical form, it returns the
string which consists of an IP address, followed by a slash
character and a mask expressed as hexadecimal form with no
punctuation like "198.51.100.0/c000ff00".




## <a id="Interface">type</a> [Interface](https://golang.org/src/net/interface.go?s=974:1340#L20)
Interface represents a mapping between network interface name
and index. It also represents network interface facility
information.


<pre>type Interface struct {
<span id="Interface.Index"></span>    Index        <a href="/pkg/builtin/#int">int</a>          <span class="comment">// positive integer that starts at one, zero is never used</span>
<span id="Interface.MTU"></span>    MTU          <a href="/pkg/builtin/#int">int</a>          <span class="comment">// maximum transmission unit</span>
<span id="Interface.Name"></span>    Name         <a href="/pkg/builtin/#string">string</a>       <span class="comment">// e.g., &#34;en0&#34;, &#34;lo0&#34;, &#34;eth0.100&#34;</span>
<span id="Interface.HardwareAddr"></span>    HardwareAddr <a href="#HardwareAddr">HardwareAddr</a> <span class="comment">// IEEE MAC-48, EUI-48 and EUI-64 form</span>
<span id="Interface.Flags"></span>    Flags        <a href="#Flags">Flags</a>        <span class="comment">// e.g., FlagUp, FlagLoopback, FlagMulticast</span>
}
</pre>









### <a id="InterfaceByIndex">func</a> [InterfaceByIndex](https://golang.org/src/net/interface.go?s=3822:3874#L118)
<pre>func InterfaceByIndex(index <a href="/pkg/builtin/#int">int</a>) (*<a href="#Interface">Interface</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
InterfaceByIndex returns the interface specified by index.

On Solaris, it returns one of the logical network interfaces
sharing the logical data link; for more precision use
InterfaceByName.




### <a id="InterfaceByName">func</a> [InterfaceByName](https://golang.org/src/net/interface.go?s=4551:4604#L143)
<pre>func InterfaceByName(name <a href="/pkg/builtin/#string">string</a>) (*<a href="#Interface">Interface</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
InterfaceByName returns the interface specified by name.




### <a id="Interfaces">func</a> [Interfaces](https://golang.org/src/net/interface.go?s=2963:3001#L89)
<pre>func Interfaces() ([]<a href="#Interface">Interface</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Interfaces returns a list of the system's network interfaces.






### <a id="Interface.Addrs">func</a> (\*Interface) [Addrs](https://golang.org/src/net/interface.go?s=2127:2172#L64)
<pre>func (ifi *<a href="#Interface">Interface</a>) Addrs() ([]<a href="#Addr">Addr</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Addrs returns a list of unicast interface addresses for a specific
interface.




### <a id="Interface.MulticastAddrs">func</a> (\*Interface) [MulticastAddrs](https://golang.org/src/net/interface.go?s=2553:2607#L77)
<pre>func (ifi *<a href="#Interface">Interface</a>) MulticastAddrs() ([]<a href="#Addr">Addr</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
MulticastAddrs returns a list of multicast, joined group addresses
for a specific interface.




## <a id="InvalidAddrError">type</a> [InvalidAddrError](https://golang.org/src/net/net.go?s=17342:17370#L553)

<pre>type InvalidAddrError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="InvalidAddrError.Error">func</a> (InvalidAddrError) [Error](https://golang.org/src/net/net.go?s=17372:17412#L555)
<pre>func (e <a href="#InvalidAddrError">InvalidAddrError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="InvalidAddrError.Temporary">func</a> (InvalidAddrError) [Temporary](https://golang.org/src/net/net.go?s=17496:17538#L557)
<pre>func (e <a href="#InvalidAddrError">InvalidAddrError</a>) Temporary() <a href="/pkg/builtin/#bool">bool</a></pre>



### <a id="InvalidAddrError.Timeout">func</a> (InvalidAddrError) [Timeout](https://golang.org/src/net/net.go?s=17436:17476#L556)
<pre>func (e <a href="#InvalidAddrError">InvalidAddrError</a>) Timeout() <a href="/pkg/builtin/#bool">bool</a></pre>



## <a id="ListenConfig">type</a> [ListenConfig](https://golang.org/src/net/dial.go?s=17749:18550#L588)
ListenConfig contains options for listening to an address.


<pre>type ListenConfig struct {
    <span class="comment">// If Control is not nil, it is called after creating the network</span>
    <span class="comment">// connection but before binding it to the operating system.</span>
    <span class="comment">//</span>
    <span class="comment">// Network and address parameters passed to Control method are not</span>
    <span class="comment">// necessarily the ones passed to Listen. For example, passing &#34;tcp&#34; to</span>
    <span class="comment">// Listen will cause the Control function to be called with &#34;tcp4&#34; or &#34;tcp6&#34;.</span>
<span id="ListenConfig.Control"></span>    Control func(network, address <a href="/pkg/builtin/#string">string</a>, c <a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#RawConn">RawConn</a>) <a href="/pkg/builtin/#error">error</a>

<span id="ListenConfig.KeepAlive"></span>    <span class="comment">// KeepAlive specifies the keep-alive period for network</span>
    <span class="comment">// connections accepted by this listener.</span>
    <span class="comment">// If zero, keep-alives are enabled if supported by the protocol</span>
    <span class="comment">// and operating system. Network protocols or operating systems</span>
    <span class="comment">// that do not support keep-alives ignore this field.</span>
    <span class="comment">// If negative, keep-alives are disabled.</span>
    KeepAlive <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>
}
</pre>











### <a id="ListenConfig.Listen">func</a> (\*ListenConfig) [Listen](https://golang.org/src/net/dial.go?s=18684:18778#L610)
<pre>func (lc *<a href="#ListenConfig">ListenConfig</a>) Listen(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, network, address <a href="/pkg/builtin/#string">string</a>) (<a href="#Listener">Listener</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Listen announces on the local network address.

See func Listen for a description of the network and address
parameters.




### <a id="ListenConfig.ListenPacket">func</a> (\*ListenConfig) [ListenPacket](https://golang.org/src/net/dial.go?s=19707:19809#L640)
<pre>func (lc *<a href="#ListenConfig">ListenConfig</a>) ListenPacket(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, network, address <a href="/pkg/builtin/#string">string</a>) (<a href="#PacketConn">PacketConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ListenPacket announces on the local network address.

See func ListenPacket for a description of the network and address
parameters.




## <a id="Listener">type</a> [Listener](https://golang.org/src/net/net.go?s=12868:13170#L371)
A Listener is a generic network listener for stream-oriented protocols.

Multiple goroutines may invoke methods on a Listener simultaneously.


<pre>type Listener interface {
    <span class="comment">// Accept waits for and returns the next connection to the listener.</span>
    Accept() (<a href="#Conn">Conn</a>, <a href="/pkg/builtin/#error">error</a>)

    <span class="comment">// Close closes the listener.</span>
    <span class="comment">// Any blocked Accept operations will be unblocked and return errors.</span>
    Close() <a href="/pkg/builtin/#error">error</a>

    <span class="comment">// Addr returns the listener&#39;s network address.</span>
    Addr() <a href="#Addr">Addr</a>
}</pre>





<a id="example_Listener">Example</a>
```go
```

output:
```txt
```




### <a id="FileListener">func</a> [FileListener](https://golang.org/src/net/file.go?s=1078:1132#L23)
<pre>func FileListener(f *<a href="/pkg/os/">os</a>.<a href="/pkg/os/#File">File</a>) (ln <a href="#Listener">Listener</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
FileListener returns a copy of the network listener corresponding
to the open file f.
It is the caller's responsibility to close ln when finished.
Closing ln does not affect f, and closing f does not affect ln.




### <a id="Listen">func</a> [Listen](https://golang.org/src/net/dial.go?s=21579:21633#L692)
<pre>func Listen(network, address <a href="/pkg/builtin/#string">string</a>) (<a href="#Listener">Listener</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Listen announces on the local network address.

The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".

For TCP networks, if the host in the address parameter is empty or
a literal unspecified IP address, Listen listens on all available
unicast and anycast IP addresses of the local system.
To only use IPv4, use network "tcp4".
The address can use a host name, but this is not recommended,
because it will create a listener for at most one of the host's IP
addresses.
If the port in the address parameter is empty or "0", as in
"127.0.0.1:" or "[::1]:0", a port number is automatically chosen.
The Addr method of Listener can be used to discover the chosen
port.

See func Dial for a description of the network and address
parameters.






## <a id="MX">type</a> [MX](https://golang.org/src/net/dnsclient.go?s=4904:4948#L184)
An MX represents a single DNS MX record.


<pre>type MX struct {
<span id="MX.Host"></span>    Host <a href="/pkg/builtin/#string">string</a>
<span id="MX.Pref"></span>    Pref <a href="/pkg/builtin/#uint16">uint16</a>
}
</pre>









### <a id="LookupMX">func</a> [LookupMX](https://golang.org/src/net/lookup.go?s=14204:14245#L403)
<pre>func LookupMX(name <a href="/pkg/builtin/#string">string</a>) ([]*<a href="#MX">MX</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
LookupMX returns the DNS MX records for the given domain name sorted by preference.






## <a id="NS">type</a> [NS](https://golang.org/src/net/dnsclient.go?s=5439:5470#L206)
An NS represents a single DNS NS record.


<pre>type NS struct {
<span id="NS.Host"></span>    Host <a href="/pkg/builtin/#string">string</a>
}
</pre>









### <a id="LookupNS">func</a> [LookupNS](https://golang.org/src/net/lookup.go?s=14577:14618#L413)
<pre>func LookupNS(name <a href="/pkg/builtin/#string">string</a>) ([]*<a href="#NS">NS</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
LookupNS returns the DNS NS records for the given domain name.






## <a id="OpError">type</a> [OpError](https://golang.org/src/net/net.go?s=14302:15076#L422)
OpError is the error type usually returned by functions in the net
package. It describes the operation, network type, and address of
an error.


<pre>type OpError struct {
<span id="OpError.Op"></span>    <span class="comment">// Op is the operation which caused the error, such as</span>
    <span class="comment">// &#34;read&#34; or &#34;write&#34;.</span>
    Op <a href="/pkg/builtin/#string">string</a>

<span id="OpError.Net"></span>    <span class="comment">// Net is the network type on which this error occurred,</span>
    <span class="comment">// such as &#34;tcp&#34; or &#34;udp6&#34;.</span>
    Net <a href="/pkg/builtin/#string">string</a>

    <span class="comment">// For operations involving a remote network connection, like</span>
    <span class="comment">// Dial, Read, or Write, Source is the corresponding local</span>
    <span class="comment">// network address.</span>
<span id="OpError.Source"></span>    Source <a href="#Addr">Addr</a>

<span id="OpError.Addr"></span>    <span class="comment">// Addr is the network address for which this error occurred.</span>
    <span class="comment">// For local operations, like Listen or SetDeadline, Addr is</span>
    <span class="comment">// the address of the local endpoint being manipulated.</span>
    <span class="comment">// For operations involving a remote network connection, like</span>
    <span class="comment">// Dial, Read, or Write, Addr is the remote address of that</span>
    <span class="comment">// connection.</span>
    Addr <a href="#Addr">Addr</a>

<span id="OpError.Err"></span>    <span class="comment">// Err is the error that occurred during the operation.</span>
    Err <a href="/pkg/builtin/#error">error</a>
}
</pre>











### <a id="OpError.Error">func</a> (\*OpError) [Error](https://golang.org/src/net/net.go?s=15129:15161#L450)
<pre>func (e *<a href="#OpError">OpError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="OpError.Temporary">func</a> (\*OpError) [Temporary](https://golang.org/src/net/net.go?s=16056:16090#L501)
<pre>func (e *<a href="#OpError">OpError</a>) Temporary() <a href="/pkg/builtin/#bool">bool</a></pre>



### <a id="OpError.Timeout">func</a> (\*OpError) [Timeout](https://golang.org/src/net/net.go?s=15815:15847#L488)
<pre>func (e *<a href="#OpError">OpError</a>) Timeout() <a href="/pkg/builtin/#bool">bool</a></pre>



### <a id="OpError.Unwrap">func</a> (\*OpError) [Unwrap](https://golang.org/src/net/net.go?s=15078:15110#L448)
<pre>func (e *<a href="#OpError">OpError</a>) Unwrap() <a href="/pkg/builtin/#error">error</a></pre>



## <a id="PacketConn">type</a> [PacketConn](https://golang.org/src/net/net.go?s=10061:12442#L300)
PacketConn is a generic packet-oriented network connection.

Multiple goroutines may invoke methods on a PacketConn simultaneously.


<pre>type PacketConn interface {
    <span class="comment">// ReadFrom reads a packet from the connection,</span>
    <span class="comment">// copying the payload into p. It returns the number of</span>
    <span class="comment">// bytes copied into p and the return address that</span>
    <span class="comment">// was on the packet.</span>
    <span class="comment">// It returns the number of bytes read (0 &lt;= n &lt;= len(p))</span>
    <span class="comment">// and any error encountered. Callers should always process</span>
    <span class="comment">// the n &gt; 0 bytes returned before considering the error err.</span>
    <span class="comment">// ReadFrom can be made to time out and return</span>
    <span class="comment">// an Error with Timeout() == true after a fixed time limit;</span>
    <span class="comment">// see SetDeadline and SetReadDeadline.</span>
    ReadFrom(p []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, addr <a href="#Addr">Addr</a>, err <a href="/pkg/builtin/#error">error</a>)

    <span class="comment">// WriteTo writes a packet with payload p to addr.</span>
    <span class="comment">// WriteTo can be made to time out and return</span>
    <span class="comment">// an Error with Timeout() == true after a fixed time limit;</span>
    <span class="comment">// see SetDeadline and SetWriteDeadline.</span>
    <span class="comment">// On packet-oriented connections, write timeouts are rare.</span>
    WriteTo(p []<a href="/pkg/builtin/#byte">byte</a>, addr <a href="#Addr">Addr</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)

    <span class="comment">// Close closes the connection.</span>
    <span class="comment">// Any blocked ReadFrom or WriteTo operations will be unblocked and return errors.</span>
    Close() <a href="/pkg/builtin/#error">error</a>

    <span class="comment">// LocalAddr returns the local network address.</span>
    LocalAddr() <a href="#Addr">Addr</a>

    <span class="comment">// SetDeadline sets the read and write deadlines associated</span>
    <span class="comment">// with the connection. It is equivalent to calling both</span>
    <span class="comment">// SetReadDeadline and SetWriteDeadline.</span>
    <span class="comment">//</span>
    <span class="comment">// A deadline is an absolute time after which I/O operations</span>
    <span class="comment">// fail with a timeout (see type Error) instead of</span>
    <span class="comment">// blocking. The deadline applies to all future and pending</span>
    <span class="comment">// I/O, not just the immediately following call to ReadFrom or</span>
    <span class="comment">// WriteTo. After a deadline has been exceeded, the connection</span>
    <span class="comment">// can be refreshed by setting a deadline in the future.</span>
    <span class="comment">//</span>
    <span class="comment">// An idle timeout can be implemented by repeatedly extending</span>
    <span class="comment">// the deadline after successful ReadFrom or WriteTo calls.</span>
    <span class="comment">//</span>
    <span class="comment">// A zero value for t means I/O operations will not time out.</span>
    SetDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a>

    <span class="comment">// SetReadDeadline sets the deadline for future ReadFrom calls</span>
    <span class="comment">// and any currently-blocked ReadFrom call.</span>
    <span class="comment">// A zero value for t means ReadFrom will not time out.</span>
    SetReadDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a>

    <span class="comment">// SetWriteDeadline sets the deadline for future WriteTo calls</span>
    <span class="comment">// and any currently-blocked WriteTo call.</span>
    <span class="comment">// Even if write times out, it may return n &gt; 0, indicating that</span>
    <span class="comment">// some of the data was successfully written.</span>
    <span class="comment">// A zero value for t means WriteTo will not time out.</span>
    SetWriteDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a>
}</pre>









### <a id="FilePacketConn">func</a> [FilePacketConn](https://golang.org/src/net/file.go?s=1519:1576#L35)
<pre>func FilePacketConn(f *<a href="/pkg/os/">os</a>.<a href="/pkg/os/#File">File</a>) (c <a href="#PacketConn">PacketConn</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
FilePacketConn returns a copy of the packet network connection
corresponding to the open file f.
It is the caller's responsibility to close f when finished.
Closing c does not affect f, and closing f does not affect c.




### <a id="ListenPacket">func</a> [ListenPacket](https://golang.org/src/net/dial.go?s=22729:22791#L719)
<pre>func ListenPacket(network, address <a href="/pkg/builtin/#string">string</a>) (<a href="#PacketConn">PacketConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ListenPacket announces on the local network address.

The network must be "udp", "udp4", "udp6", "unixgram", or an IP
transport. The IP transports are "ip", "ip4", or "ip6" followed by
a colon and a literal protocol number or a protocol name, as in
"ip:1" or "ip:icmp".

For UDP and IP networks, if the host in the address parameter is
empty or a literal unspecified IP address, ListenPacket listens on
all available IP addresses of the local system except multicast IP
addresses.
To only use IPv4, use network "udp4" or "ip4:proto".
The address can use a host name, but this is not recommended,
because it will create a listener for at most one of the host's IP
addresses.
If the port in the address parameter is empty or "0", as in
"127.0.0.1:" or "[::1]:0", a port number is automatically chosen.
The LocalAddr method of PacketConn can be used to discover the
chosen port.

See func Dial for a description of the network and address
parameters.






## <a id="ParseError">type</a> [ParseError](https://golang.org/src/net/net.go?s=16506:16690#L517)
A ParseError is the error type of literal network address parsers.


<pre>type ParseError struct {
<span id="ParseError.Type"></span>    <span class="comment">// Type is the type of string that was expected, such as</span>
    <span class="comment">// &#34;IP address&#34;, &#34;CIDR address&#34;.</span>
    Type <a href="/pkg/builtin/#string">string</a>

<span id="ParseError.Text"></span>    <span class="comment">// Text is the malformed text string.</span>
    Text <a href="/pkg/builtin/#string">string</a>
}
</pre>











### <a id="ParseError.Error">func</a> (\*ParseError) [Error](https://golang.org/src/net/net.go?s=16692:16727#L526)
<pre>func (e *<a href="#ParseError">ParseError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Resolver">type</a> [Resolver](https://golang.org/src/net/lookup.go?s=3212:4993#L110)
A Resolver looks up names and numbers.

A nil *Resolver is equivalent to a zero Resolver.


<pre>type Resolver struct {
<span id="Resolver.PreferGo"></span>    <span class="comment">// PreferGo controls whether Go&#39;s built-in DNS resolver is preferred</span>
    <span class="comment">// on platforms where it&#39;s available. It is equivalent to setting</span>
    <span class="comment">// GODEBUG=netdns=go, but scoped to just this resolver.</span>
    PreferGo <a href="/pkg/builtin/#bool">bool</a>

<span id="Resolver.StrictErrors"></span>    <span class="comment">// StrictErrors controls the behavior of temporary errors</span>
    <span class="comment">// (including timeout, socket errors, and SERVFAIL) when using</span>
    <span class="comment">// Go&#39;s built-in resolver. For a query composed of multiple</span>
    <span class="comment">// sub-queries (such as an A+AAAA address lookup, or walking the</span>
    <span class="comment">// DNS search list), this option causes such errors to abort the</span>
    <span class="comment">// whole query instead of returning a partial result. This is</span>
    <span class="comment">// not enabled by default because it may affect compatibility</span>
    <span class="comment">// with resolvers that process AAAA queries incorrectly.</span>
    StrictErrors <a href="/pkg/builtin/#bool">bool</a>

<span id="Resolver.Dial"></span>    <span class="comment">// Dial optionally specifies an alternate dialer for use by</span>
    <span class="comment">// Go&#39;s built-in DNS resolver to make TCP and UDP connections</span>
    <span class="comment">// to DNS services. The host in the address parameter will</span>
    <span class="comment">// always be a literal IP address and not a host name, and the</span>
    <span class="comment">// port in the address parameter will be a literal port number</span>
    <span class="comment">// and not a service name.</span>
    <span class="comment">// If the Conn returned is also a PacketConn, sent and received DNS</span>
    <span class="comment">// messages must adhere to RFC 1035 section 4.2.1, &#34;UDP usage&#34;.</span>
    <span class="comment">// Otherwise, DNS messages transmitted over Conn must adhere</span>
    <span class="comment">// to RFC 7766 section 5, &#34;Transport Protocol Selection&#34;.</span>
    <span class="comment">// If nil, the default dialer is used.</span>
    Dial func(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, network, address <a href="/pkg/builtin/#string">string</a>) (<a href="#Conn">Conn</a>, <a href="/pkg/builtin/#error">error</a>)
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="Resolver.LookupAddr">func</a> (\*Resolver) [LookupAddr](https://golang.org/src/net/lookup.go?s=15726:15817#L443)
<pre>func (r *<a href="#Resolver">Resolver</a>) LookupAddr(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, addr <a href="/pkg/builtin/#string">string</a>) (names []<a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
LookupAddr performs a reverse lookup for the given address, returning a list
of names mapping to that address.




### <a id="Resolver.LookupCNAME">func</a> (\*Resolver) [LookupCNAME](https://golang.org/src/net/lookup.go?s=12664:12754#L372)
<pre>func (r *<a href="#Resolver">Resolver</a>) LookupCNAME(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, host <a href="/pkg/builtin/#string">string</a>) (cname <a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
LookupCNAME returns the canonical name for the given host.
Callers that do not care about the canonical name can call
LookupHost or LookupIP directly; both take care of resolving
the canonical name as part of the lookup.

A canonical name is the final name after following zero
or more CNAME records.
LookupCNAME does not return an error if host does not
contain DNS "CNAME" records, as long as host resolves to
address records.




### <a id="Resolver.LookupHost">func</a> (\*Resolver) [LookupHost](https://golang.org/src/net/lookup.go?s=5635:5726#L166)
<pre>func (r *<a href="#Resolver">Resolver</a>) LookupHost(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, host <a href="/pkg/builtin/#string">string</a>) (addrs []<a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
LookupHost looks up the given host using the local resolver.
It returns a slice of that host's addresses.




### <a id="Resolver.LookupIPAddr">func</a> (\*Resolver) [LookupIPAddr](https://golang.org/src/net/lookup.go?s=6548:6631#L194)
<pre>func (r *<a href="#Resolver">Resolver</a>) LookupIPAddr(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, host <a href="/pkg/builtin/#string">string</a>) ([]<a href="#IPAddr">IPAddr</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
LookupIPAddr looks up host using the local resolver.
It returns a slice of that host's IPv4 and IPv6 addresses.




### <a id="Resolver.LookupMX">func</a> (\*Resolver) [LookupMX](https://golang.org/src/net/lookup.go?s=14399:14475#L408)
<pre>func (r *<a href="#Resolver">Resolver</a>) LookupMX(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, name <a href="/pkg/builtin/#string">string</a>) ([]*<a href="#MX">MX</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
LookupMX returns the DNS MX records for the given domain name sorted by preference.




### <a id="Resolver.LookupNS">func</a> (\*Resolver) [LookupNS](https://golang.org/src/net/lookup.go?s=14751:14827#L418)
<pre>func (r *<a href="#Resolver">Resolver</a>) LookupNS(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, name <a href="/pkg/builtin/#string">string</a>) ([]*<a href="#NS">NS</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
LookupNS returns the DNS NS records for the given domain name.




### <a id="Resolver.LookupPort">func</a> (\*Resolver) [LookupPort](https://golang.org/src/net/lookup.go?s=11027:11124#L327)
<pre>func (r *<a href="#Resolver">Resolver</a>) LookupPort(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, network, service <a href="/pkg/builtin/#string">string</a>) (port <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
LookupPort looks up the port for the given network and service.




### <a id="Resolver.LookupSRV">func</a> (\*Resolver) [LookupSRV](https://golang.org/src/net/lookup.go?s=13946:14064#L398)
<pre>func (r *<a href="#Resolver">Resolver</a>) LookupSRV(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, service, proto, name <a href="/pkg/builtin/#string">string</a>) (cname <a href="/pkg/builtin/#string">string</a>, addrs []*<a href="#SRV">SRV</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
LookupSRV tries to resolve an SRV query of the given service,
protocol, and domain name. The proto is "tcp" or "udp".
The returned records are sorted by priority and randomized
by weight within a priority.

LookupSRV constructs the DNS name to look up following RFC 2782.
That is, it looks up _service._proto.name. To accommodate services
publishing SRV records under non-standard names, if both service
and proto are empty strings, LookupSRV looks up name directly.




### <a id="Resolver.LookupTXT">func</a> (\*Resolver) [LookupTXT](https://golang.org/src/net/lookup.go?s=15112:15192#L428)
<pre>func (r *<a href="#Resolver">Resolver</a>) LookupTXT(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, name <a href="/pkg/builtin/#string">string</a>) ([]<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
LookupTXT returns the DNS TXT records for the given domain name.




## <a id="SRV">type</a> [SRV](https://golang.org/src/net/dnsclient.go?s=3623:3710#L130)
An SRV represents a single DNS SRV record.


<pre>type SRV struct {
<span id="SRV.Target"></span>    Target   <a href="/pkg/builtin/#string">string</a>
<span id="SRV.Port"></span>    Port     <a href="/pkg/builtin/#uint16">uint16</a>
<span id="SRV.Priority"></span>    Priority <a href="/pkg/builtin/#uint16">uint16</a>
<span id="SRV.Weight"></span>    Weight   <a href="/pkg/builtin/#uint16">uint16</a>
}
</pre>









### <a id="LookupSRV">func</a> [LookupSRV](https://golang.org/src/net/lookup.go?s=13286:13369#L385)
<pre>func LookupSRV(service, proto, name <a href="/pkg/builtin/#string">string</a>) (cname <a href="/pkg/builtin/#string">string</a>, addrs []*<a href="#SRV">SRV</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
LookupSRV tries to resolve an SRV query of the given service,
protocol, and domain name. The proto is "tcp" or "udp".
The returned records are sorted by priority and randomized
by weight within a priority.

LookupSRV constructs the DNS name to look up following RFC 2782.
That is, it looks up _service._proto.name. To accommodate services
publishing SRV records under non-standard names, if both service
and proto are empty strings, LookupSRV looks up name directly.






## <a id="TCPAddr">type</a> [TCPAddr](https://golang.org/src/net/tcpsock.go?s=388:474#L9)
TCPAddr represents the address of a TCP end point.


<pre>type TCPAddr struct {
<span id="TCPAddr.IP"></span>    IP   <a href="#IP">IP</a>
<span id="TCPAddr.Port"></span>    Port <a href="/pkg/builtin/#int">int</a>
<span id="TCPAddr.Zone"></span>    Zone <a href="/pkg/builtin/#string">string</a> <span class="comment">// IPv6 scoped addressing zone</span>
}
</pre>









### <a id="ResolveTCPAddr">func</a> [ResolveTCPAddr](https://golang.org/src/net/tcpsock.go?s=1596:1658#L58)
<pre>func ResolveTCPAddr(network, address <a href="/pkg/builtin/#string">string</a>) (*<a href="#TCPAddr">TCPAddr</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ResolveTCPAddr returns an address of TCP end point.

The network must be a TCP network name.

If the host in the address parameter is not a literal IP address or
the port is not a literal port number, ResolveTCPAddr resolves the
address to an address of TCP end point.
Otherwise, it parses the address as a pair of literal IP address
and port number.
The address parameter can use a host name, but this is not
recommended, because it will return at most one of the host name's
IP addresses.

See func Dial for a description of the network and address
parameters.






### <a id="TCPAddr.Network">func</a> (\*TCPAddr) [Network](https://golang.org/src/net/tcpsock.go?s=530:564#L16)
<pre>func (a *<a href="#TCPAddr">TCPAddr</a>) Network() <a href="/pkg/builtin/#string">string</a></pre>
Network returns the address's network name, "tcp".




### <a id="TCPAddr.String">func</a> (\*TCPAddr) [String](https://golang.org/src/net/tcpsock.go?s=583:616#L18)
<pre>func (a *<a href="#TCPAddr">TCPAddr</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="TCPConn">type</a> [TCPConn](https://golang.org/src/net/tcpsock.go?s=2118:2147#L75)
TCPConn is an implementation of the Conn interface for TCP network
connections.


<pre>type TCPConn struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="DialTCP">func</a> [DialTCP](https://golang.org/src/net/tcpsock.go?s=5677:5746#L196)
<pre>func DialTCP(network <a href="/pkg/builtin/#string">string</a>, laddr, raddr *<a href="#TCPAddr">TCPAddr</a>) (*<a href="#TCPConn">TCPConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DialTCP acts like Dial for TCP networks.

The network must be a TCP network name; see func Dial for details.

If laddr is nil, a local address is automatically chosen.
If the IP field of raddr is nil or an unspecified IP address, the
local system is assumed.






### <a id="TCPConn.Close">func</a> (\*TCPConn) [Close](https://golang.org/src/net/net.go?s=7068:7096#L194)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the connection.




### <a id="TCPConn.CloseRead">func</a> (\*TCPConn) [CloseRead](https://golang.org/src/net/tcpsock.go?s=2816:2851#L102)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) CloseRead() <a href="/pkg/builtin/#error">error</a></pre>
CloseRead shuts down the reading side of the TCP connection.
Most callers should just use Close.




### <a id="TCPConn.CloseWrite">func</a> (\*TCPConn) [CloseWrite](https://golang.org/src/net/tcpsock.go?s=3153:3189#L114)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) CloseWrite() <a href="/pkg/builtin/#error">error</a></pre>
CloseWrite shuts down the writing side of the TCP connection.
Most callers should just use Close.




### <a id="TCPConn.File">func</a> (\*TCPConn) [File](https://golang.org/src/net/net.go?s=9729:9774#L289)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) File() (f *<a href="/pkg/os/">os</a>.<a href="/pkg/os/#File">File</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
File returns a copy of the underlying os.File.
It is the caller's responsibility to close f when finished.
Closing c does not affect f, and closing f does not affect c.

The returned os.File's file descriptor is different from the connection's.
Attempting to change properties of the original using this duplicate
may or may not have the desired effect.




### <a id="TCPConn.LocalAddr">func</a> (\*TCPConn) [LocalAddr](https://golang.org/src/net/net.go?s=7425:7456#L208)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) LocalAddr() <a href="#Addr">Addr</a></pre>
LocalAddr returns the local network address.
The Addr returned is shared by all invocations of LocalAddr, so
do not modify it.




### <a id="TCPConn.Read">func</a> (\*TCPConn) [Read](https://golang.org/src/net/net.go?s=6487:6529#L170)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) Read(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Read implements the Conn Read method.




### <a id="TCPConn.ReadFrom">func</a> (\*TCPConn) [ReadFrom](https://golang.org/src/net/tcpsock.go?s=2436:2490#L89)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) ReadFrom(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (<a href="/pkg/builtin/#int64">int64</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadFrom implements the io.ReaderFrom ReadFrom method.




### <a id="TCPConn.RemoteAddr">func</a> (\*TCPConn) [RemoteAddr](https://golang.org/src/net/net.go?s=7650:7682#L218)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) RemoteAddr() <a href="#Addr">Addr</a></pre>
RemoteAddr returns the remote network address.
The Addr returned is shared by all invocations of RemoteAddr, so
do not modify it.




### <a id="TCPConn.SetDeadline">func</a> (\*TCPConn) [SetDeadline](https://golang.org/src/net/net.go?s=7792:7837#L226)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) SetDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetDeadline implements the Conn SetDeadline method.




### <a id="TCPConn.SetKeepAlive">func</a> (\*TCPConn) [SetKeepAlive](https://golang.org/src/net/tcpsock.go?s=4221:4273#L148)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) SetKeepAlive(keepalive <a href="/pkg/builtin/#bool">bool</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetKeepAlive sets whether the operating system should send
keep-alive messages on the connection.




### <a id="TCPConn.SetKeepAlivePeriod">func</a> (\*TCPConn) [SetKeepAlivePeriod](https://golang.org/src/net/tcpsock.go?s=4537:4596#L159)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) SetKeepAlivePeriod(d <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetKeepAlivePeriod sets period between keep-alives.




### <a id="TCPConn.SetLinger">func</a> (\*TCPConn) [SetLinger](https://golang.org/src/net/tcpsock.go?s=3875:3917#L136)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) SetLinger(sec <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetLinger sets the behavior of Close on a connection which still
has data waiting to be sent or to be acknowledged.

If sec < 0 (the default), the operating system finishes sending the
data in the background.

If sec == 0, the operating system discards any unsent or
unacknowledged data.

If sec > 0, the data is sent in the background as with sec < 0. On
some operating systems after sec seconds have elapsed any remaining
unsent data may be discarded.




### <a id="TCPConn.SetNoDelay">func</a> (\*TCPConn) [SetNoDelay](https://golang.org/src/net/tcpsock.go?s=5046:5094#L173)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) SetNoDelay(noDelay <a href="/pkg/builtin/#bool">bool</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetNoDelay controls whether the operating system should delay
packet transmission in hopes of sending fewer packets (Nagle's
algorithm).  The default is true (no delay), meaning that data is
sent as soon as possible after a Write.




### <a id="TCPConn.SetReadBuffer">func</a> (\*TCPConn) [SetReadBuffer](https://golang.org/src/net/net.go?s=8756:8801#L260)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) SetReadBuffer(bytes <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetReadBuffer sets the size of the operating system's
receive buffer associated with the connection.




### <a id="TCPConn.SetReadDeadline">func</a> (\*TCPConn) [SetReadDeadline](https://golang.org/src/net/net.go?s=8092:8141#L237)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) SetReadDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetReadDeadline implements the Conn SetReadDeadline method.




### <a id="TCPConn.SetWriteBuffer">func</a> (\*TCPConn) [SetWriteBuffer](https://golang.org/src/net/net.go?s=9109:9155#L272)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) SetWriteBuffer(bytes <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetWriteBuffer sets the size of the operating system's
transmit buffer associated with the connection.




### <a id="TCPConn.SetWriteDeadline">func</a> (\*TCPConn) [SetWriteDeadline](https://golang.org/src/net/net.go?s=8402:8452#L248)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) SetWriteDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetWriteDeadline implements the Conn SetWriteDeadline method.




### <a id="TCPConn.SyscallConn">func</a> (\*TCPConn) [SyscallConn](https://golang.org/src/net/tcpsock.go?s=2245:2301#L81)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) SyscallConn() (<a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#RawConn">RawConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
SyscallConn returns a raw network connection.
This implements the syscall.Conn interface.




### <a id="TCPConn.Write">func</a> (\*TCPConn) [Write](https://golang.org/src/net/net.go?s=6790:6833#L182)
<pre>func (c *<a href="#TCPConn">TCPConn</a>) Write(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Write implements the Conn Write method.




## <a id="TCPListener">type</a> [TCPListener](https://golang.org/src/net/tcpsock.go?s=6457:6512#L215)
TCPListener is a TCP network listener. Clients should typically
use variables of type Listener instead of assuming TCP.


<pre>type TCPListener struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="ListenTCP">func</a> [ListenTCP](https://golang.org/src/net/tcpsock.go?s=9391:9459#L314)
<pre>func ListenTCP(network <a href="/pkg/builtin/#string">string</a>, laddr *<a href="#TCPAddr">TCPAddr</a>) (*<a href="#TCPListener">TCPListener</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ListenTCP acts like Listen for TCP networks.

The network must be a TCP network name; see func Dial for details.

If the IP field of laddr is nil or an unspecified IP address,
ListenTCP listens on all available unicast and anycast IP addresses
of the local system.
If the Port field of laddr is 0, a port number is automatically
chosen.






### <a id="TCPListener.Accept">func</a> (\*TCPListener) [Accept](https://golang.org/src/net/tcpsock.go?s=7300:7344#L247)
<pre>func (l *<a href="#TCPListener">TCPListener</a>) Accept() (<a href="#Conn">Conn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Accept implements the Accept method in the Listener interface; it
waits for the next call and returns a generic Conn.




### <a id="TCPListener.AcceptTCP">func</a> (\*TCPListener) [AcceptTCP](https://golang.org/src/net/tcpsock.go?s=6923:6974#L234)
<pre>func (l *<a href="#TCPListener">TCPListener</a>) AcceptTCP() (*<a href="#TCPConn">TCPConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
AcceptTCP accepts the next incoming call and returns the new
connection.




### <a id="TCPListener.Addr">func</a> (\*TCPListener) [Addr](https://golang.org/src/net/tcpsock.go?s=8001:8034#L273)
<pre>func (l *<a href="#TCPListener">TCPListener</a>) Addr() <a href="#Addr">Addr</a></pre>
Addr returns the listener's network address, a *TCPAddr.
The Addr returned is shared by all invocations of Addr, so
do not modify it.




### <a id="TCPListener.Close">func</a> (\*TCPListener) [Close](https://golang.org/src/net/tcpsock.go?s=7639:7674#L260)
<pre>func (l *<a href="#TCPListener">TCPListener</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close stops listening on the TCP address.
Already Accepted connections are not closed.




### <a id="TCPListener.File">func</a> (\*TCPListener) [File](https://golang.org/src/net/tcpsock.go?s=8787:8839#L294)
<pre>func (l *<a href="#TCPListener">TCPListener</a>) File() (f *<a href="/pkg/os/">os</a>.<a href="/pkg/os/#File">File</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
File returns a copy of the underlying os.File.
It is the caller's responsibility to close f when finished.
Closing l does not affect f, and closing f does not affect l.

The returned os.File's file descriptor is different from the
connection's. Attempting to change properties of the original
using this duplicate may or may not have the desired effect.




### <a id="TCPListener.SetDeadline">func</a> (\*TCPListener) [SetDeadline](https://golang.org/src/net/tcpsock.go?s=8165:8217#L277)
<pre>func (l *<a href="#TCPListener">TCPListener</a>) SetDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetDeadline sets the deadline associated with the listener.
A zero time value disables the deadline.




### <a id="TCPListener.SyscallConn">func</a> (\*TCPListener) [SyscallConn](https://golang.org/src/net/tcpsock.go?s=6703:6763#L225)
<pre>func (l *<a href="#TCPListener">TCPListener</a>) SyscallConn() (<a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#RawConn">RawConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
SyscallConn returns a raw network connection.
This implements the syscall.Conn interface.

The returned RawConn only supports calling Control. Read and
Write return an error.




## <a id="UDPAddr">type</a> [UDPAddr](https://golang.org/src/net/udpsock.go?s=617:703#L15)
UDPAddr represents the address of a UDP end point.


<pre>type UDPAddr struct {
<span id="UDPAddr.IP"></span>    IP   <a href="#IP">IP</a>
<span id="UDPAddr.Port"></span>    Port <a href="/pkg/builtin/#int">int</a>
<span id="UDPAddr.Zone"></span>    Zone <a href="/pkg/builtin/#string">string</a> <span class="comment">// IPv6 scoped addressing zone</span>
}
</pre>









### <a id="ResolveUDPAddr">func</a> [ResolveUDPAddr](https://golang.org/src/net/udpsock.go?s=1825:1887#L64)
<pre>func ResolveUDPAddr(network, address <a href="/pkg/builtin/#string">string</a>) (*<a href="#UDPAddr">UDPAddr</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ResolveUDPAddr returns an address of UDP end point.

The network must be a UDP network name.

If the host in the address parameter is not a literal IP address or
the port is not a literal port number, ResolveUDPAddr resolves the
address to an address of UDP end point.
Otherwise, it parses the address as a pair of literal IP address
and port number.
The address parameter can use a host name, but this is not
recommended, because it will return at most one of the host name's
IP addresses.

See func Dial for a description of the network and address
parameters.






### <a id="UDPAddr.Network">func</a> (\*UDPAddr) [Network](https://golang.org/src/net/udpsock.go?s=759:793#L22)
<pre>func (a *<a href="#UDPAddr">UDPAddr</a>) Network() <a href="/pkg/builtin/#string">string</a></pre>
Network returns the address's network name, "udp".




### <a id="UDPAddr.String">func</a> (\*UDPAddr) [String](https://golang.org/src/net/udpsock.go?s=812:845#L24)
<pre>func (a *<a href="#UDPAddr">UDPAddr</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="UDPConn">type</a> [UDPConn](https://golang.org/src/net/udpsock.go?s=2364:2393#L81)
UDPConn is the implementation of the Conn and PacketConn interfaces
for UDP network connections.


<pre>type UDPConn struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="DialUDP">func</a> [DialUDP](https://golang.org/src/net/udpsock.go?s=5929:5998#L195)
<pre>func DialUDP(network <a href="/pkg/builtin/#string">string</a>, laddr, raddr *<a href="#UDPAddr">UDPAddr</a>) (*<a href="#UDPConn">UDPConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DialUDP acts like Dial for UDP networks.

The network must be a UDP network name; see func Dial for details.

If laddr is nil, a local address is automatically chosen.
If the IP field of raddr is nil or an unspecified IP address, the
local system is assumed.




### <a id="ListenMulticastUDP">func</a> [ListenMulticastUDP](https://golang.org/src/net/udpsock.go?s=8308:8397#L255)
<pre>func ListenMulticastUDP(network <a href="/pkg/builtin/#string">string</a>, ifi *<a href="#Interface">Interface</a>, gaddr *<a href="#UDPAddr">UDPAddr</a>) (*<a href="#UDPConn">UDPConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ListenMulticastUDP acts like ListenPacket for UDP networks but
takes a group address on a specific network interface.

The network must be a UDP network name; see func Dial for details.

ListenMulticastUDP listens on all available IP addresses of the
local system including the group, multicast IP address.
If ifi is nil, ListenMulticastUDP uses the system-assigned
multicast interface, although this is not recommended because the
assignment depends on platforms and sometimes it might require
routing configuration.
If the Port field of gaddr is 0, a port number is automatically
chosen.

ListenMulticastUDP is just for convenience of simple, small
applications. There are golang.org/x/net/ipv4 and
golang.org/x/net/ipv6 packages for general purpose uses.




### <a id="ListenUDP">func</a> [ListenUDP](https://golang.org/src/net/udpsock.go?s=6961:7025#L221)
<pre>func ListenUDP(network <a href="/pkg/builtin/#string">string</a>, laddr *<a href="#UDPAddr">UDPAddr</a>) (*<a href="#UDPConn">UDPConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ListenUDP acts like ListenPacket for UDP networks.

The network must be a UDP network name; see func Dial for details.

If the IP field of laddr is nil or an unspecified IP address,
ListenUDP listens on all available IP addresses of the local system
except multicast IP addresses.
If the Port field of laddr is 0, a port number is automatically
chosen.






### <a id="UDPConn.Close">func</a> (\*UDPConn) [Close](https://golang.org/src/net/net.go?s=7068:7096#L194)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the connection.




### <a id="UDPConn.File">func</a> (\*UDPConn) [File](https://golang.org/src/net/net.go?s=9729:9774#L289)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) File() (f *<a href="/pkg/os/">os</a>.<a href="/pkg/os/#File">File</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
File returns a copy of the underlying os.File.
It is the caller's responsibility to close f when finished.
Closing c does not affect f, and closing f does not affect c.

The returned os.File's file descriptor is different from the connection's.
Attempting to change properties of the original using this duplicate
may or may not have the desired effect.




### <a id="UDPConn.LocalAddr">func</a> (\*UDPConn) [LocalAddr](https://golang.org/src/net/net.go?s=7425:7456#L208)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) LocalAddr() <a href="#Addr">Addr</a></pre>
LocalAddr returns the local network address.
The Addr returned is shared by all invocations of LocalAddr, so
do not modify it.




### <a id="UDPConn.Read">func</a> (\*UDPConn) [Read](https://golang.org/src/net/net.go?s=6487:6529#L170)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) Read(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Read implements the Conn Read method.




### <a id="UDPConn.ReadFrom">func</a> (\*UDPConn) [ReadFrom](https://golang.org/src/net/udpsock.go?s=3017:3072#L107)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) ReadFrom(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="#Addr">Addr</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadFrom implements the PacketConn ReadFrom method.




### <a id="UDPConn.ReadFromUDP">func</a> (\*UDPConn) [ReadFromUDP](https://golang.org/src/net/udpsock.go?s=2681:2743#L95)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) ReadFromUDP(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, *<a href="#UDPAddr">UDPAddr</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadFromUDP acts like ReadFrom but returns a UDPAddr.




### <a id="UDPConn.ReadMsgUDP">func</a> (\*UDPConn) [ReadMsgUDP](https://golang.org/src/net/udpsock.go?s=3742:3832#L128)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) ReadMsgUDP(b, oob []<a href="/pkg/builtin/#byte">byte</a>) (n, oobn, flags <a href="/pkg/builtin/#int">int</a>, addr *<a href="#UDPAddr">UDPAddr</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadMsgUDP reads a message from c, copying the payload into b and
the associated out-of-band data into oob. It returns the number of
bytes copied into b, the number of bytes copied into oob, the flags
that were set on the message and the source address of the message.

The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be
used to manipulate IP-level socket options in oob.




### <a id="UDPConn.RemoteAddr">func</a> (\*UDPConn) [RemoteAddr](https://golang.org/src/net/net.go?s=7650:7682#L218)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) RemoteAddr() <a href="#Addr">Addr</a></pre>
RemoteAddr returns the remote network address.
The Addr returned is shared by all invocations of RemoteAddr, so
do not modify it.




### <a id="UDPConn.SetDeadline">func</a> (\*UDPConn) [SetDeadline](https://golang.org/src/net/net.go?s=7792:7837#L226)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) SetDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetDeadline implements the Conn SetDeadline method.




### <a id="UDPConn.SetReadBuffer">func</a> (\*UDPConn) [SetReadBuffer](https://golang.org/src/net/net.go?s=8756:8801#L260)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) SetReadBuffer(bytes <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetReadBuffer sets the size of the operating system's
receive buffer associated with the connection.




### <a id="UDPConn.SetReadDeadline">func</a> (\*UDPConn) [SetReadDeadline](https://golang.org/src/net/net.go?s=8092:8141#L237)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) SetReadDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetReadDeadline implements the Conn SetReadDeadline method.




### <a id="UDPConn.SetWriteBuffer">func</a> (\*UDPConn) [SetWriteBuffer](https://golang.org/src/net/net.go?s=9109:9155#L272)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) SetWriteBuffer(bytes <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetWriteBuffer sets the size of the operating system's
transmit buffer associated with the connection.




### <a id="UDPConn.SetWriteDeadline">func</a> (\*UDPConn) [SetWriteDeadline](https://golang.org/src/net/net.go?s=8402:8452#L248)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) SetWriteDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetWriteDeadline implements the Conn SetWriteDeadline method.




### <a id="UDPConn.SyscallConn">func</a> (\*UDPConn) [SyscallConn](https://golang.org/src/net/udpsock.go?s=2491:2547#L87)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) SyscallConn() (<a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#RawConn">RawConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
SyscallConn returns a raw network connection.
This implements the syscall.Conn interface.




### <a id="UDPConn.Write">func</a> (\*UDPConn) [Write](https://golang.org/src/net/net.go?s=6790:6833#L182)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) Write(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Write implements the Conn Write method.




### <a id="UDPConn.WriteMsgUDP">func</a> (\*UDPConn) [WriteMsgUDP](https://golang.org/src/net/udpsock.go?s=5282:5366#L175)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) WriteMsgUDP(b, oob []<a href="/pkg/builtin/#byte">byte</a>, addr *<a href="#UDPAddr">UDPAddr</a>) (n, oobn <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
WriteMsgUDP writes a message to addr via c if c isn't connected, or
to c's remote address if c is connected (in which case addr must be
nil). The payload is copied from b and the associated out-of-band
data is copied from oob. It returns the number of payload and
out-of-band bytes written.

The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be
used to manipulate IP-level socket options in oob.




### <a id="UDPConn.WriteTo">func</a> (\*UDPConn) [WriteTo](https://golang.org/src/net/udpsock.go?s=4443:4502#L152)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) WriteTo(b []<a href="/pkg/builtin/#byte">byte</a>, addr <a href="#Addr">Addr</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
WriteTo implements the PacketConn WriteTo method.


<a id="example_UDPConn_WriteTo">Example</a>
```go
```

output:
```txt
```


### <a id="UDPConn.WriteToUDP">func</a> (\*UDPConn) [WriteToUDP](https://golang.org/src/net/udpsock.go?s=4113:4179#L140)
<pre>func (c *<a href="#UDPConn">UDPConn</a>) WriteToUDP(b []<a href="/pkg/builtin/#byte">byte</a>, addr *<a href="#UDPAddr">UDPAddr</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
WriteToUDP acts like WriteTo but takes a UDPAddr.




## <a id="UnixAddr">type</a> [UnixAddr](https://golang.org/src/net/unixsock.go?s=556:606#L12)
UnixAddr represents the address of a Unix domain socket end point.


<pre>type UnixAddr struct {
<span id="UnixAddr.Name"></span>    Name <a href="/pkg/builtin/#string">string</a>
<span id="UnixAddr.Net"></span>    Net  <a href="/pkg/builtin/#string">string</a>
}
</pre>









### <a id="ResolveUnixAddr">func</a> [ResolveUnixAddr](https://golang.org/src/net/unixsock.go?s=1191:1255#L47)
<pre>func ResolveUnixAddr(network, address <a href="/pkg/builtin/#string">string</a>) (*<a href="#UnixAddr">UnixAddr</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ResolveUnixAddr returns an address of Unix domain socket end point.

The network must be a Unix network name.

See func Dial for a description of the network and address
parameters.






### <a id="UnixAddr.Network">func</a> (\*UnixAddr) [Network](https://golang.org/src/net/unixsock.go?s=694:729#L19)
<pre>func (a *<a href="#UnixAddr">UnixAddr</a>) Network() <a href="/pkg/builtin/#string">string</a></pre>
Network returns the address's network name, "unix", "unixgram" or
"unixpacket".




### <a id="UnixAddr.String">func</a> (\*UnixAddr) [String](https://golang.org/src/net/unixsock.go?s=749:783#L23)
<pre>func (a *<a href="#UnixAddr">UnixAddr</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="UnixConn">type</a> [UnixConn](https://golang.org/src/net/unixsock.go?s=1526:1556#L58)
UnixConn is an implementation of the Conn interface for connections
to Unix domain sockets.


<pre>type UnixConn struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="DialUnix">func</a> [DialUnix](https://golang.org/src/net/unixsock.go?s=5595:5667#L193)
<pre>func DialUnix(network <a href="/pkg/builtin/#string">string</a>, laddr, raddr *<a href="#UnixAddr">UnixAddr</a>) (*<a href="#UnixConn">UnixConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DialUnix acts like Dial for Unix networks.

The network must be a Unix network name; see func Dial for details.

If laddr is non-nil, it is used as the local address for the
connection.




### <a id="ListenUnixgram">func</a> [ListenUnixgram](https://golang.org/src/net/unixsock.go?s=9804:9875#L327)
<pre>func ListenUnixgram(network <a href="/pkg/builtin/#string">string</a>, laddr *<a href="#UnixAddr">UnixAddr</a>) (*<a href="#UnixConn">UnixConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ListenUnixgram acts like ListenPacket for Unix networks.

The network must be "unixgram".






### <a id="UnixConn.Close">func</a> (\*UnixConn) [Close](https://golang.org/src/net/net.go?s=7068:7096#L194)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the connection.




### <a id="UnixConn.CloseRead">func</a> (\*UnixConn) [CloseRead](https://golang.org/src/net/unixsock.go?s=1899:1935#L73)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) CloseRead() <a href="/pkg/builtin/#error">error</a></pre>
CloseRead shuts down the reading side of the Unix domain connection.
Most callers should just use Close.




### <a id="UnixConn.CloseWrite">func</a> (\*UnixConn) [CloseWrite](https://golang.org/src/net/unixsock.go?s=2245:2282#L85)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) CloseWrite() <a href="/pkg/builtin/#error">error</a></pre>
CloseWrite shuts down the writing side of the Unix domain connection.
Most callers should just use Close.




### <a id="UnixConn.File">func</a> (\*UnixConn) [File](https://golang.org/src/net/net.go?s=9729:9774#L289)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) File() (f *<a href="/pkg/os/">os</a>.<a href="/pkg/os/#File">File</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
File returns a copy of the underlying os.File.
It is the caller's responsibility to close f when finished.
Closing c does not affect f, and closing f does not affect c.

The returned os.File's file descriptor is different from the connection's.
Attempting to change properties of the original using this duplicate
may or may not have the desired effect.




### <a id="UnixConn.LocalAddr">func</a> (\*UnixConn) [LocalAddr](https://golang.org/src/net/net.go?s=7425:7456#L208)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) LocalAddr() <a href="#Addr">Addr</a></pre>
LocalAddr returns the local network address.
The Addr returned is shared by all invocations of LocalAddr, so
do not modify it.




### <a id="UnixConn.Read">func</a> (\*UnixConn) [Read](https://golang.org/src/net/net.go?s=6487:6529#L170)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) Read(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Read implements the Conn Read method.




### <a id="UnixConn.ReadFrom">func</a> (\*UnixConn) [ReadFrom](https://golang.org/src/net/unixsock.go?s=2879:2935#L108)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) ReadFrom(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="#Addr">Addr</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadFrom implements the PacketConn ReadFrom method.




### <a id="UnixConn.ReadFromUnix">func</a> (\*UnixConn) [ReadFromUnix](https://golang.org/src/net/unixsock.go?s=2540:2605#L96)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) ReadFromUnix(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, *<a href="#UnixAddr">UnixAddr</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadFromUnix acts like ReadFrom but returns a UnixAddr.




### <a id="UnixConn.ReadMsgUnix">func</a> (\*UnixConn) [ReadMsgUnix](https://golang.org/src/net/unixsock.go?s=3602:3695#L129)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) ReadMsgUnix(b, oob []<a href="/pkg/builtin/#byte">byte</a>) (n, oobn, flags <a href="/pkg/builtin/#int">int</a>, addr *<a href="#UnixAddr">UnixAddr</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadMsgUnix reads a message from c, copying the payload into b and
the associated out-of-band data into oob. It returns the number of
bytes copied into b, the number of bytes copied into oob, the flags
that were set on the message and the source address of the message.

Note that if len(b) == 0 and len(oob) > 0, this function will still
read (and discard) 1 byte from the connection.




### <a id="UnixConn.RemoteAddr">func</a> (\*UnixConn) [RemoteAddr](https://golang.org/src/net/net.go?s=7650:7682#L218)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) RemoteAddr() <a href="#Addr">Addr</a></pre>
RemoteAddr returns the remote network address.
The Addr returned is shared by all invocations of RemoteAddr, so
do not modify it.




### <a id="UnixConn.SetDeadline">func</a> (\*UnixConn) [SetDeadline](https://golang.org/src/net/net.go?s=7792:7837#L226)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) SetDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetDeadline implements the Conn SetDeadline method.




### <a id="UnixConn.SetReadBuffer">func</a> (\*UnixConn) [SetReadBuffer](https://golang.org/src/net/net.go?s=8756:8801#L260)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) SetReadBuffer(bytes <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetReadBuffer sets the size of the operating system's
receive buffer associated with the connection.




### <a id="UnixConn.SetReadDeadline">func</a> (\*UnixConn) [SetReadDeadline](https://golang.org/src/net/net.go?s=8092:8141#L237)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) SetReadDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetReadDeadline implements the Conn SetReadDeadline method.




### <a id="UnixConn.SetWriteBuffer">func</a> (\*UnixConn) [SetWriteBuffer](https://golang.org/src/net/net.go?s=9109:9155#L272)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) SetWriteBuffer(bytes <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetWriteBuffer sets the size of the operating system's
transmit buffer associated with the connection.




### <a id="UnixConn.SetWriteDeadline">func</a> (\*UnixConn) [SetWriteDeadline](https://golang.org/src/net/net.go?s=8402:8452#L248)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) SetWriteDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetWriteDeadline implements the Conn SetWriteDeadline method.




### <a id="UnixConn.SyscallConn">func</a> (\*UnixConn) [SyscallConn](https://golang.org/src/net/unixsock.go?s=1654:1711#L64)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) SyscallConn() (<a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#RawConn">RawConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
SyscallConn returns a raw network connection.
This implements the syscall.Conn interface.




### <a id="UnixConn.Write">func</a> (\*UnixConn) [Write](https://golang.org/src/net/net.go?s=6790:6833#L182)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) Write(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Write implements the Conn Write method.




### <a id="UnixConn.WriteMsgUnix">func</a> (\*UnixConn) [WriteMsgUnix](https://golang.org/src/net/unixsock.go?s=5018:5105#L174)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) WriteMsgUnix(b, oob []<a href="/pkg/builtin/#byte">byte</a>, addr *<a href="#UnixAddr">UnixAddr</a>) (n, oobn <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
WriteMsgUnix writes a message to addr via c, copying the payload
from b and the associated out-of-band data from oob. It returns the
number of payload and out-of-band bytes written.

Note that if len(b) == 0 and len(oob) > 0, this function will still
write 1 byte to the connection.




### <a id="UnixConn.WriteTo">func</a> (\*UnixConn) [WriteTo](https://golang.org/src/net/unixsock.go?s=4311:4371#L153)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) WriteTo(b []<a href="/pkg/builtin/#byte">byte</a>, addr <a href="#Addr">Addr</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
WriteTo implements the PacketConn WriteTo method.




### <a id="UnixConn.WriteToUnix">func</a> (\*UnixConn) [WriteToUnix](https://golang.org/src/net/unixsock.go?s=3978:4047#L141)
<pre>func (c *<a href="#UnixConn">UnixConn</a>) WriteToUnix(b []<a href="/pkg/builtin/#byte">byte</a>, addr *<a href="#UnixAddr">UnixAddr</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
WriteToUnix acts like WriteTo but takes a UnixAddr.




## <a id="UnixListener">type</a> [UnixListener](https://golang.org/src/net/unixsock.go?s=6287:6392#L210)
UnixListener is a Unix domain socket listener. Clients should
typically use variables of type Listener instead of assuming Unix
domain sockets.


<pre>type UnixListener struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="ListenUnix">func</a> [ListenUnix](https://golang.org/src/net/unixsock.go?s=9067:9138#L307)
<pre>func ListenUnix(network <a href="/pkg/builtin/#string">string</a>, laddr *<a href="#UnixAddr">UnixAddr</a>) (*<a href="#UnixListener">UnixListener</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ListenUnix acts like Listen for Unix networks.

The network must be "unix" or "unixpacket".






### <a id="UnixListener.Accept">func</a> (\*UnixListener) [Accept](https://golang.org/src/net/unixsock.go?s=7250:7295#L246)
<pre>func (l *<a href="#UnixListener">UnixListener</a>) Accept() (<a href="#Conn">Conn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Accept implements the Accept method in the Listener interface.
Returned connections will be of type *UnixConn.




### <a id="UnixListener.AcceptUnix">func</a> (\*UnixListener) [AcceptUnix](https://golang.org/src/net/unixsock.go?s=6877:6931#L233)
<pre>func (l *<a href="#UnixListener">UnixListener</a>) AcceptUnix() (*<a href="#UnixConn">UnixConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
AcceptUnix accepts the next incoming call and returns the new
connection.




### <a id="UnixListener.Addr">func</a> (\*UnixListener) [Addr](https://golang.org/src/net/unixsock.go?s=7942:7976#L272)
<pre>func (l *<a href="#UnixListener">UnixListener</a>) Addr() <a href="#Addr">Addr</a></pre>
Addr returns the listener's network address.
The Addr returned is shared by all invocations of Addr, so
do not modify it.




### <a id="UnixListener.Close">func</a> (\*UnixListener) [Close](https://golang.org/src/net/unixsock.go?s=7591:7627#L259)
<pre>func (l *<a href="#UnixListener">UnixListener</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close stops listening on the Unix address. Already accepted
connections are not closed.




### <a id="UnixListener.File">func</a> (\*UnixListener) [File](https://golang.org/src/net/unixsock.go?s=8730:8783#L293)
<pre>func (l *<a href="#UnixListener">UnixListener</a>) File() (f *<a href="/pkg/os/">os</a>.<a href="/pkg/os/#File">File</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
File returns a copy of the underlying os.File.
It is the caller's responsibility to close f when finished.
Closing l does not affect f, and closing f does not affect l.

The returned os.File's file descriptor is different from the
connection's. Attempting to change properties of the original
using this duplicate may or may not have the desired effect.




### <a id="UnixListener.SetDeadline">func</a> (\*UnixListener) [SetDeadline](https://golang.org/src/net/unixsock.go?s=8107:8160#L276)
<pre>func (l *<a href="#UnixListener">UnixListener</a>) SetDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetDeadline sets the deadline associated with the listener.
A zero time value disables the deadline.




### <a id="UnixListener.SetUnlinkOnClose">func</a> (\*UnixListener) [SetUnlinkOnClose](https://golang.org/src/net/unixsock_posix.go?s=5325:5377#L195)
<pre>func (l *<a href="#UnixListener">UnixListener</a>) SetUnlinkOnClose(unlink <a href="/pkg/builtin/#bool">bool</a>)</pre>
SetUnlinkOnClose sets whether the underlying socket file should be removed
from the file system when the listener is closed.

The default behavior is to unlink the socket file only when package net created it.
That is, when the listener and the underlying socket file were created by a call to
Listen or ListenUnix, then by default closing the listener will remove the socket file.
but if the listener was created by a call to FileListener to use an already existing
socket file, then by default closing the listener will not remove the socket file.




### <a id="UnixListener.SyscallConn">func</a> (\*UnixListener) [SyscallConn](https://golang.org/src/net/unixsock.go?s=6655:6716#L224)
<pre>func (l *<a href="#UnixListener">UnixListener</a>) SyscallConn() (<a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#RawConn">RawConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
SyscallConn returns a raw network connection.
This implements the syscall.Conn interface.

The returned RawConn only supports calling Control. Read and
Write return an error.




## <a id="UnknownNetworkError">type</a> [UnknownNetworkError](https://golang.org/src/net/net.go?s=17094:17125#L547)

<pre>type UnknownNetworkError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="UnknownNetworkError.Error">func</a> (UnknownNetworkError) [Error](https://golang.org/src/net/net.go?s=17127:17170#L549)
<pre>func (e <a href="#UnknownNetworkError">UnknownNetworkError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="UnknownNetworkError.Temporary">func</a> (UnknownNetworkError) [Temporary](https://golang.org/src/net/net.go?s=17278:17323#L551)
<pre>func (e <a href="#UnknownNetworkError">UnknownNetworkError</a>) Temporary() <a href="/pkg/builtin/#bool">bool</a></pre>



### <a id="UnknownNetworkError.Timeout">func</a> (UnknownNetworkError) [Timeout](https://golang.org/src/net/net.go?s=17215:17258#L550)
<pre>func (e <a href="#UnknownNetworkError">UnknownNetworkError</a>) Timeout() <a href="/pkg/builtin/#bool">bool</a></pre>






