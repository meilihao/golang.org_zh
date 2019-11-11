

# debug
`import "runtime/debug"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package debug contains facilities for programs to debug themselves while
they are running.




## <a id="pkg-index">Index</a>
* [func FreeOSMemory()](#FreeOSMemory)
* [func PrintStack()](#PrintStack)
* [func ReadGCStats(stats *GCStats)](#ReadGCStats)
* [func SetGCPercent(percent int) int](#SetGCPercent)
* [func SetMaxStack(bytes int) int](#SetMaxStack)
* [func SetMaxThreads(threads int) int](#SetMaxThreads)
* [func SetPanicOnFault(enabled bool) bool](#SetPanicOnFault)
* [func SetTraceback(level string)](#SetTraceback)
* [func Stack() []byte](#Stack)
* [func WriteHeapDump(fd uintptr)](#WriteHeapDump)
* [type BuildInfo](#BuildInfo)
  * [func ReadBuildInfo() (info *BuildInfo, ok bool)](#ReadBuildInfo)
* [type GCStats](#GCStats)
* [type Module](#Module)




#### <a id="pkg-files">Package files</a>
[garbage.go](https://golang.org/src/runtime/debug/garbage.go) [mod.go](https://golang.org/src/runtime/debug/mod.go) [stack.go](https://golang.org/src/runtime/debug/stack.go) [stubs.go](https://golang.org/src/runtime/debug/stubs.go) 






## <a id="FreeOSMemory">func</a> [FreeOSMemory](https://golang.org/src/runtime/debug/garbage.go?s=3859:3878#L89)
<pre>func FreeOSMemory()</pre>
FreeOSMemory forces a garbage collection followed by an
attempt to return as much memory to the operating system
as possible. (Even if this is not called, the runtime gradually
returns memory to the operating system in a background task.)



## <a id="PrintStack">func</a> [PrintStack](https://golang.org/src/runtime/debug/stack.go?s=383:400#L5)
<pre>func PrintStack()</pre>
PrintStack prints to standard error the stack trace returned by runtime.Stack.



## <a id="ReadGCStats">func</a> [ReadGCStats](https://golang.org/src/runtime/debug/garbage.go?s=1206:1238#L21)
<pre>func ReadGCStats(stats *<a href="#GCStats">GCStats</a>)</pre>
ReadGCStats reads statistics about garbage collection into stats.
The number of entries in the pause history is system-dependent;
stats.Pause slice will be reused if large enough, reallocated otherwise.
ReadGCStats may use the full capacity of the stats.Pause slice.
If stats.PauseQuantiles is non-empty, ReadGCStats fills it with quantiles
summarizing the distribution of pause time. For example, if
len(stats.PauseQuantiles) is 5, it will be filled with the minimum,
25%, 50%, 75%, and maximum pause times.



## <a id="SetGCPercent">func</a> [SetGCPercent](https://golang.org/src/runtime/debug/garbage.go?s=3526:3560#L81)
<pre>func SetGCPercent(percent <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
SetGCPercent sets the garbage collection target percentage:
a collection is triggered when the ratio of freshly allocated data
to live data remaining after the previous collection reaches this percentage.
SetGCPercent returns the previous setting.
The initial setting is the value of the GOGC environment variable
at startup, or 100 if the variable is not set.
A negative percentage disables garbage collection.



## <a id="SetMaxStack">func</a> [SetMaxStack](https://golang.org/src/runtime/debug/garbage.go?s=4361:4392#L103)
<pre>func SetMaxStack(bytes <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
SetMaxStack sets the maximum amount of memory that
can be used by a single goroutine stack.
If any goroutine exceeds this limit while growing its stack,
the program crashes.
SetMaxStack returns the previous setting.
The initial setting is 1 GB on 64-bit systems, 250 MB on 32-bit systems.

SetMaxStack is useful mainly for limiting the damage done by
goroutines that enter an infinite recursion. It only limits future
stack growth.



## <a id="SetMaxThreads">func</a> [SetMaxThreads](https://golang.org/src/runtime/debug/garbage.go?s=5205:5240#L121)
<pre>func SetMaxThreads(threads <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
SetMaxThreads sets the maximum number of operating system
threads that the Go program can use. If it attempts to use more than
this many, the program crashes.
SetMaxThreads returns the previous setting.
The initial setting is 10,000 threads.

The limit controls the number of operating system threads, not the number
of goroutines. A Go program creates a new thread only when a goroutine
is ready to run but all the existing threads are blocked in system calls, cgo calls,
or are locked to other goroutines due to use of runtime.LockOSThread.

SetMaxThreads is useful mainly for limiting the damage done by
programs that create an unbounded number of threads. The idea is
to take down the program before it takes down the operating system.



## <a id="SetPanicOnFault">func</a> [SetPanicOnFault](https://golang.org/src/runtime/debug/garbage.go?s=5865:5904#L134)
<pre>func SetPanicOnFault(enabled <a href="/pkg/builtin/#bool">bool</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
SetPanicOnFault controls the runtime's behavior when a program faults
at an unexpected (non-nil) address. Such faults are typically caused by
bugs such as runtime memory corruption, so the default response is to crash
the program. Programs working with memory-mapped files or unsafe
manipulation of memory may cause faults at non-nil addresses in less
dramatic situations; SetPanicOnFault allows such programs to request
that the runtime trigger only a panic, not a crash.
SetPanicOnFault applies only to the current goroutine.
It returns the previous setting.



## <a id="SetTraceback">func</a> [SetTraceback](https://golang.org/src/runtime/debug/garbage.go?s=6953:6984#L158)
<pre>func SetTraceback(level <a href="/pkg/builtin/#string">string</a>)</pre>
SetTraceback sets the amount of detail printed by the runtime in
the traceback it prints before exiting due to an unrecovered panic
or an internal runtime error.
The level argument takes the same values as the GOTRACEBACK
environment variable. For example, SetTraceback("all") ensure
that the program prints all goroutines when it crashes.
See the package runtime documentation for details.
If SetTraceback is called with a level lower than that of the
environment variable, the call is ignored.



## <a id="Stack">func</a> [Stack](https://golang.org/src/runtime/debug/stack.go?s=587:606#L11)
<pre>func Stack() []<a href="/pkg/builtin/#byte">byte</a></pre>
Stack returns a formatted stack trace of the goroutine that calls it.
It calls runtime.Stack with a large enough buffer to capture the entire trace.



## <a id="WriteHeapDump">func</a> [WriteHeapDump](https://golang.org/src/runtime/debug/garbage.go?s=6398:6428#L147)
<pre>func WriteHeapDump(fd <a href="/pkg/builtin/#uintptr">uintptr</a>)</pre>
WriteHeapDump writes a description of the heap and the objects in
it to the given file descriptor.

WriteHeapDump suspends the execution of all goroutines until the heap
dump is completely written.  Thus, the file descriptor must not be
connected to a pipe or socket whose other end is in the same Go
process; instead, use a temporary file or network socket.

The heap dump format is defined at <a href="https://golang.org/s/go15heapdump">https://golang.org/s/go15heapdump</a>.





## <a id="BuildInfo">type</a> [BuildInfo](https://golang.org/src/runtime/debug/mod.go?s=569:721#L13)
BuildInfo represents the build information read from
the running binary.


<pre>type BuildInfo struct {
<span id="BuildInfo.Path"></span>    Path <a href="/pkg/builtin/#string">string</a>    <span class="comment">// The main package path</span>
<span id="BuildInfo.Main"></span>    Main <a href="#Module">Module</a>    <span class="comment">// The main module information</span>
<span id="BuildInfo.Deps"></span>    Deps []*<a href="#Module">Module</a> <span class="comment">// Module dependencies</span>
}
</pre>









### <a id="ReadBuildInfo">func</a> [ReadBuildInfo](https://golang.org/src/runtime/debug/mod.go?s=404:451#L7)
<pre>func ReadBuildInfo() (info *<a href="#BuildInfo">BuildInfo</a>, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
ReadBuildInfo returns the build information embedded
in the running binary. The information is available only
in binaries built with module support.






## <a id="GCStats">type</a> [GCStats](https://golang.org/src/runtime/debug/garbage.go?s=279:671#L4)
GCStats collect information about recent garbage collections.


<pre>type GCStats struct {
<span id="GCStats.LastGC"></span>    LastGC         <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>       <span class="comment">// time of last collection</span>
<span id="GCStats.NumGC"></span>    NumGC          <a href="/pkg/builtin/#int64">int64</a>           <span class="comment">// number of garbage collections</span>
<span id="GCStats.PauseTotal"></span>    PauseTotal     <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>   <span class="comment">// total pause for all collections</span>
<span id="GCStats.Pause"></span>    Pause          []<a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a> <span class="comment">// pause history, most recent first</span>
<span id="GCStats.PauseEnd"></span>    PauseEnd       []<a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>     <span class="comment">// pause end times history, most recent first</span>
<span id="GCStats.PauseQuantiles"></span>    PauseQuantiles []<a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>
}
</pre>











## <a id="Module">type</a> [Module](https://golang.org/src/runtime/debug/mod.go?s=754:916#L20)
Module represents a module.


<pre>type Module struct {
<span id="Module.Path"></span>    Path    <a href="/pkg/builtin/#string">string</a>  <span class="comment">// module path</span>
<span id="Module.Version"></span>    Version <a href="/pkg/builtin/#string">string</a>  <span class="comment">// module version</span>
<span id="Module.Sum"></span>    Sum     <a href="/pkg/builtin/#string">string</a>  <span class="comment">// checksum</span>
<span id="Module.Replace"></span>    Replace *<a href="#Module">Module</a> <span class="comment">// replaced by this module</span>
}
</pre>














