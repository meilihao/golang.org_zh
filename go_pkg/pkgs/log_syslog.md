

# syslog
`import "log/syslog"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package syslog provides a simple interface to the system log
service. It can send messages to the syslog daemon using UNIX
domain sockets, UDP or TCP.

Only one call to Dial is necessary. On write failures,
the syslog client will attempt to reconnect to the server
and write again.

The syslog package is frozen and is not accepting new features.
Some external packages provide more functionality. See:


	<a href="https://godoc.org/?q=syslog">https://godoc.org/?q=syslog</a>

syslog 为系统日志服务提供了一个简单的接口. 它能通过UNIX domain socket, UDP 或 TCP发送消息给syslog守护进程.

Dial 函数只需要调用一次. 如果写入失败， syslog 客户端会尝试重新连接服务器并再次写入.

syslog 包处于冻结状态，不会接受新的功能. 一些其它包提供更多的功能. 参考：


	<a href="https://godoc.org/?q=syslog">https://godoc.org/?q=syslog</a>



## <a id="pkg-index">Index</a>
* [func NewLogger(p Priority, logFlag int) (*log.Logger, error)](#NewLogger)
* [type Priority](#Priority)
* [type Writer](#Writer)
  * [func Dial(network, raddr string, priority Priority, tag string) (*Writer, error)](#Dial)
  * [func New(priority Priority, tag string) (*Writer, error)](#New)
  * [func (w *Writer) Alert(m string) error](#Writer.Alert)
  * [func (w *Writer) Close() error](#Writer.Close)
  * [func (w *Writer) Crit(m string) error](#Writer.Crit)
  * [func (w *Writer) Debug(m string) error](#Writer.Debug)
  * [func (w *Writer) Emerg(m string) error](#Writer.Emerg)
  * [func (w *Writer) Err(m string) error](#Writer.Err)
  * [func (w *Writer) Info(m string) error](#Writer.Info)
  * [func (w *Writer) Notice(m string) error](#Writer.Notice)
  * [func (w *Writer) Warning(m string) error](#Writer.Warning)
  * [func (w *Writer) Write(b []byte) (int, error)](#Writer.Write)


#### <a id="pkg-examples">Examples</a>
* [Dial](#example_Dial)


#### <a id="pkg-files">Package files</a>
[doc.go](https://golang.org/src/log/syslog/doc.go) [syslog.go](https://golang.org/src/log/syslog/syslog.go) [syslog_unix.go](https://golang.org/src/log/syslog/syslog_unix.go) 






## <a id="NewLogger">func</a> [NewLogger](https://golang.org/src/log/syslog/syslog.go?s=7405:7465#L299)
<pre>func NewLogger(p <a href="#Priority">Priority</a>, logFlag <a href="/pkg/builtin/#int">int</a>) (*<a href="/pkg/log/">log</a>.<a href="/pkg/log/#Logger">Logger</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewLogger creates a log.Logger whose output is written to the
system log service with the specified priority, a combination of
the syslog facility and severity. The logFlag argument is the flag
set passed through to log.New to create the Logger.

NewLogger 创建一个 log.Logger 对象，它会把携带特定优先级(syslog 类型和严重程度的组合)的消息写到系统日志服务. 参数 logFlag 和 log.New 的 flag 参数作用相同.



## <a id="Priority">type</a> [Priority](https://golang.org/src/log/syslog/syslog.go?s=521:538#L14)
The Priority is a combination of the syslog facility and
severity. For example, LOG_ALERT | LOG_FTP sends an alert severity
message from the FTP facility. The default severity is LOG_EMERG;
the default facility is LOG_KERN.

Priority 是 syslog 类型和严重程度的结合. 比如， LOG_ALERT | LOG_FTP 是从 FTP 设备发送过来的一个警告消息. 默认的严重程度是 LOG_EMERG，默认的类型是 LOG_KERN.


<pre>type Priority <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (

    <span class="comment">// From /usr/include/sys/syslog.h.</span>
    <span class="comment">// These are the same on Linux, BSD, and OS X.</span>
    <span id="LOG_EMERG">LOG_EMERG</span> <a href="#Priority">Priority</a> = <a href="/pkg/builtin/#iota">iota</a>
    <span id="LOG_ALERT">LOG_ALERT</span>
    <span id="LOG_CRIT">LOG_CRIT</span>
    <span id="LOG_ERR">LOG_ERR</span>
    <span id="LOG_WARNING">LOG_WARNING</span>
    <span id="LOG_NOTICE">LOG_NOTICE</span>
    <span id="LOG_INFO">LOG_INFO</span>
    <span id="LOG_DEBUG">LOG_DEBUG</span>
)</pre>

<pre>const (

    <span class="comment">// From /usr/include/sys/syslog.h.</span>
    <span class="comment">// These are the same up to LOG_FTP on Linux, BSD, and OS X.</span>
    <span id="LOG_KERN">LOG_KERN</span> <a href="#Priority">Priority</a> = <a href="/pkg/builtin/#iota">iota</a> &lt;&lt; 3
    <span id="LOG_USER">LOG_USER</span>
    <span id="LOG_MAIL">LOG_MAIL</span>
    <span id="LOG_DAEMON">LOG_DAEMON</span>
    <span id="LOG_AUTH">LOG_AUTH</span>
    <span id="LOG_SYSLOG">LOG_SYSLOG</span>
    <span id="LOG_LPR">LOG_LPR</span>
    <span id="LOG_NEWS">LOG_NEWS</span>
    <span id="LOG_UUCP">LOG_UUCP</span>
    <span id="LOG_CRON">LOG_CRON</span>
    <span id="LOG_AUTHPRIV">LOG_AUTHPRIV</span>
    <span id="LOG_FTP">LOG_FTP</span>

    <span id="LOG_LOCAL0">LOG_LOCAL0</span>
    <span id="LOG_LOCAL1">LOG_LOCAL1</span>
    <span id="LOG_LOCAL2">LOG_LOCAL2</span>
    <span id="LOG_LOCAL3">LOG_LOCAL3</span>
    <span id="LOG_LOCAL4">LOG_LOCAL4</span>
    <span id="LOG_LOCAL5">LOG_LOCAL5</span>
    <span id="LOG_LOCAL6">LOG_LOCAL6</span>
    <span id="LOG_LOCAL7">LOG_LOCAL7</span>
)</pre>









## <a id="Writer">type</a> [Writer](https://golang.org/src/log/syslog/syslog.go?s=1273:1432#L66)
A Writer is a connection to a syslog server.

Writer 是一个 syslog 服务器的链接.

<pre>type Writer struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="Dial">func</a> [Dial](https://golang.org/src/log/syslog/syslog.go?s=2740:2820#L108)
<pre>func Dial(network, raddr <a href="/pkg/builtin/#string">string</a>, priority <a href="#Priority">Priority</a>, tag <a href="/pkg/builtin/#string">string</a>) (*<a href="#Writer">Writer</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Dial establishes a connection to a log daemon by connecting to
address raddr on the specified network. Each write to the returned
writer sends a log message with the facility and severity
(from priority) and tag. If tag is empty, the os.Args[0] is used.
If network is empty, Dial will connect to the local syslog server.
Otherwise, see the documentation for net.Dial for valid values
of network and raddr.

Dial 通过network和raddr 建立与日志守护程序的连接. 每次向返回的writer写入都会发送一条日志消息，其中包含`类型和严重性(来自优先级)`和tag. 如果tag参数为空，则使用 os.Args[0]. 如果network参数为空，Dial 将连接到本地系统日志服务. 否则，请参阅 net.Dial 的文档以传递有效的network和 raddr.


<a id="example_Dial">Example</a>
```go
package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.Dial("tcp", "localhost:1234",
		syslog.LOG_WARNING|syslog.LOG_DAEMON, "demotag")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(sysLog, "This is a daemon warning with demotag.")
	sysLog.Emerg("And this is a daemon emergency with demotag.")
}
```

output:
```txt
./prog.go:10:17: undefined: syslog.Dial
./prog.go:11:3: undefined: syslog.LOG_WARNING
./prog.go:11:22: undefined: syslog.LOG_DAEMON
```


### <a id="New">func</a> [New](https://golang.org/src/log/syslog/syslog.go?s=2215:2271#L97)
<pre>func New(priority <a href="#Priority">Priority</a>, tag <a href="/pkg/builtin/#string">string</a>) (*<a href="#Writer">Writer</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
New establishes a new connection to the system log daemon. Each
write to the returned writer sends a log message with the given
priority (a combination of the syslog facility and severity) and
prefix tag. If tag is empty, the os.Args[0] is used.

New 与系统日志守护进程建立一个新的连接. 每次向返回的writer写入都会发送一条携带指定优先级(syslog 类型和严重程度的组合)和前缀tag的日志消息. 如果参数 tag 为空，会使用os.Args[0].




### <a id="Writer.Alert">func</a> (\*Writer) [Alert](https://golang.org/src/log/syslog/syslog.go?s=4429:4467#L190)
<pre>func (w *<a href="#Writer">Writer</a>) Alert(m <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Alert logs a message with severity LOG_ALERT, ignoring the severity
passed to New.

Alert 记录一个严重程度是 LOG_ALERT 的消息，忽略传递给 New 里的 severity 参数.



### <a id="Writer.Close">func</a> (\*Writer) [Close](https://golang.org/src/log/syslog/syslog.go?s=3996:4026#L169)
<pre>func (w *<a href="#Writer">Writer</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes a connection to the syslog daemon.

Close 关闭与syslog daemon的连接.


### <a id="Writer.Crit">func</a> (\*Writer) [Crit](https://golang.org/src/log/syslog/syslog.go?s=4613:4650#L197)
<pre>func (w *<a href="#Writer">Writer</a>) Crit(m <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Crit logs a message with severity LOG_CRIT, ignoring the severity
passed to New.

Crit 记录一个严重程度是 LOG_CRIT 的消息，忽略传递给 New 里的 severity 参数.


### <a id="Writer.Debug">func</a> (\*Writer) [Debug](https://golang.org/src/log/syslog/syslog.go?s=5541:5579#L232)
<pre>func (w *<a href="#Writer">Writer</a>) Debug(m <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Debug logs a message with severity LOG_DEBUG, ignoring the severity
passed to New.

Debug 记录一个严重程度是 LOG_DEBUG 的消息，忽略传递给 New 里的 severity 参数.


### <a id="Writer.Emerg">func</a> (\*Writer) [Emerg](https://golang.org/src/log/syslog/syslog.go?s=4243:4281#L183)
<pre>func (w *<a href="#Writer">Writer</a>) Emerg(m <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Emerg logs a message with severity LOG_EMERG, ignoring the severity
passed to New.

Emerg 记录一个严重程度是 LOG_EMERG 的消息，忽略传递给 New 里的 severity 参数.


### <a id="Writer.Err">func</a> (\*Writer) [Err](https://golang.org/src/log/syslog/syslog.go?s=4793:4829#L204)
<pre>func (w *<a href="#Writer">Writer</a>) Err(m <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Err logs a message with severity LOG_ERR, ignoring the severity
passed to New.

Err 记录一个严重程度是 LOG_ERR 的消息，忽略传递给 New 里的 severity 参数.


### <a id="Writer.Info">func</a> (\*Writer) [Info](https://golang.org/src/log/syslog/syslog.go?s=5357:5394#L225)
<pre>func (w *<a href="#Writer">Writer</a>) Info(m <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Info logs a message with severity LOG_INFO, ignoring the severity
passed to New.

Info 记录一个严重程度是 Info 的消息，忽略传递给 New 里的 severity 参数.


### <a id="Writer.Notice">func</a> (\*Writer) [Notice](https://golang.org/src/log/syslog/syslog.go?s=5171:5210#L218)
<pre>func (w *<a href="#Writer">Writer</a>) Notice(m <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Notice logs a message with severity LOG_NOTICE, ignoring the
severity passed to New.

Notice 记录一个严重程度是 Notice 的消息，忽略传递给 New 里的 severity 参数.


### <a id="Writer.Warning">func</a> (\*Writer) [Warning](https://golang.org/src/log/syslog/syslog.go?s=4979:5019#L211)
<pre>func (w *<a href="#Writer">Writer</a>) Warning(m <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Warning logs a message with severity LOG_WARNING, ignoring the
severity passed to New.

Warning 记录一个严重程度是 Warning 的消息，忽略传递给 New 里的 severity 参数.


### <a id="Writer.Write">func</a> (\*Writer) [Write](https://golang.org/src/log/syslog/syslog.go?s=3847:3892#L164)
<pre>func (w *<a href="#Writer">Writer</a>) Write(b []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Write sends a log message to the syslog daemon.

Write 向syslog daemon发送一条log消息.





