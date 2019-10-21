

# pprof
`import "net/http/pprof"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package pprof serves via its HTTP server runtime profiling data
in the format expected by the pprof visualization tool.

The package is typically only imported for the side effect of
registering its HTTP handlers.
The handled paths all begin with /debug/pprof/.

To use pprof, link this package into your program:


	import _ "net/http/pprof"

If your application is not already running an http server, you
need to start one. Add "net/http" and "log" to your imports and
the following code to your main function:


	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

Then use the pprof tool to look at the heap profile:


	go tool pprof <a href="http://localhost:6060/debug/pprof/heap">http://localhost:6060/debug/pprof/heap</a>

Or to look at a 30-second CPU profile:


	go tool pprof <a href="http://localhost:6060/debug/pprof/profile?seconds=30">http://localhost:6060/debug/pprof/profile?seconds=30</a>

Or to look at the goroutine blocking profile, after calling
runtime.SetBlockProfileRate in your program:


	go tool pprof <a href="http://localhost:6060/debug/pprof/block">http://localhost:6060/debug/pprof/block</a>

Or to collect a 5-second execution trace:


	wget <a href="http://localhost:6060/debug/pprof/trace?seconds=5">http://localhost:6060/debug/pprof/trace?seconds=5</a>

Or to look at the holders of contended mutexes, after calling
runtime.SetMutexProfileFraction in your program:


	go tool pprof <a href="http://localhost:6060/debug/pprof/mutex">http://localhost:6060/debug/pprof/mutex</a>

To view all available profiles, open <a href="http://localhost:6060/debug/pprof/">http://localhost:6060/debug/pprof/</a>
in your browser.

For a study of the facility in action, visit


	<a href="https://blog.golang.org/2011/06/profiling-go-programs.html">https://blog.golang.org/2011/06/profiling-go-programs.html</a>




## <a id="pkg-index">Index</a>
* [func Cmdline(w http.ResponseWriter, r *http.Request)](#Cmdline)
* [func Handler(name string) http.Handler](#Handler)
* [func Index(w http.ResponseWriter, r *http.Request)](#Index)
* [func Profile(w http.ResponseWriter, r *http.Request)](#Profile)
* [func Symbol(w http.ResponseWriter, r *http.Request)](#Symbol)
* [func Trace(w http.ResponseWriter, r *http.Request)](#Trace)




#### <a id="pkg-files">Package files</a>
[pprof.go](https://golang.org/src/net/http/pprof/pprof.go) 






## <a id="Cmdline">func</a> [Cmdline](https://golang.org/src/net/http/pprof/pprof.go?s=2316:2368#L73)
<pre>func Cmdline(w <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#ResponseWriter">ResponseWriter</a>, r *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>)</pre>
Cmdline responds with the running program's
command line, with arguments separated by NUL bytes.
The package initialization registers it as /debug/pprof/cmdline.



## <a id="Handler">func</a> [Handler](https://golang.org/src/net/http/pprof/pprof.go?s=6833:6871#L211)
<pre>func Handler(name <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Handler">Handler</a></pre>
Handler returns an HTTP handler that serves the named profile.



## <a id="Index">func</a> [Index](https://golang.org/src/net/http/pprof/pprof.go?s=8810:8860#L254)
<pre>func Index(w <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#ResponseWriter">ResponseWriter</a>, r *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>)</pre>
Index responds with the pprof-formatted profile named by the request.
For example, "/debug/pprof/heap" serves the "heap" profile.
Index responds to a request for "/debug/pprof/" with an HTML page
listing the available profiles.



## <a id="Profile">func</a> [Profile](https://golang.org/src/net/http/pprof/pprof.go?s=3453:3505#L106)
<pre>func Profile(w <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#ResponseWriter">ResponseWriter</a>, r *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>)</pre>
Profile responds with the pprof-formatted cpu profile.
Profiling lasts for duration specified in seconds GET parameter, or for 30 seconds if not specified.
The package initialization registers it as /debug/pprof/profile.



## <a id="Symbol">func</a> [Symbol](https://golang.org/src/net/http/pprof/pprof.go?s=5606:5657#L164)
<pre>func Symbol(w <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#ResponseWriter">ResponseWriter</a>, r *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>)</pre>
Symbol looks up the program counters listed in the request,
responding with a table mapping program counters to function names.
The package initialization registers it as /debug/pprof/symbol.



## <a id="Trace">func</a> [Trace](https://golang.org/src/net/http/pprof/pprof.go?s=4562:4612#L135)
<pre>func Trace(w <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#ResponseWriter">ResponseWriter</a>, r *<a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Request">Request</a>)</pre>
Trace responds with the execution trace in binary form.
Tracing lasts for duration specified in seconds GET parameter, or for 1 second if not specified.
The package initialization registers it as /debug/pprof/trace.








