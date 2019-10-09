

# runtime
`import "runtime"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package runtime contains operations that interact with Go's runtime system,
such as functions to control goroutines. It also includes the low-level type information
used by the reflect package; see reflect's documentation for the programmable
interface to the run-time type system.

### Environment Variables
The following environment variables ($name or %name%, depending on the host
operating system) control the run-time behavior of Go programs. The meanings
and use may change from release to release.

The GOGC variable sets the initial garbage collection target percentage.
A collection is triggered when the ratio of freshly allocated data to live data
remaining after the previous collection reaches this percentage. The default
is GOGC=100. Setting GOGC=off disables the garbage collector entirely.
The runtime/debug package's SetGCPercent function allows changing this
percentage at run time. See <a href="https://golang.org/pkg/runtime/debug/#SetGCPercent">https://golang.org/pkg/runtime/debug/#SetGCPercent</a>.

The GODEBUG variable controls debugging variables within the runtime.
It is a comma-separated list of name=val pairs setting these named variables:


	allocfreetrace: setting allocfreetrace=1 causes every allocation to be
	profiled and a stack trace printed on each object's allocation and free.
	
	clobberfree: setting clobberfree=1 causes the garbage collector to
	clobber the memory content of an object with bad content when it frees
	the object.
	
	cgocheck: setting cgocheck=0 disables all checks for packages
	using cgo to incorrectly pass Go pointers to non-Go code.
	Setting cgocheck=1 (the default) enables relatively cheap
	checks that may miss some errors.  Setting cgocheck=2 enables
	expensive checks that should not miss any errors, but will
	cause your program to run slower.
	
	efence: setting efence=1 causes the allocator to run in a mode
	where each object is allocated on a unique page and addresses are
	never recycled.
	
	gccheckmark: setting gccheckmark=1 enables verification of the
	garbage collector's concurrent mark phase by performing a
	second mark pass while the world is stopped.  If the second
	pass finds a reachable object that was not found by concurrent
	mark, the garbage collector will panic.
	
	gcpacertrace: setting gcpacertrace=1 causes the garbage collector to
	print information about the internal state of the concurrent pacer.
	
	gcshrinkstackoff: setting gcshrinkstackoff=1 disables moving goroutines
	onto smaller stacks. In this mode, a goroutine's stack can only grow.
	
	gcstoptheworld: setting gcstoptheworld=1 disables concurrent garbage collection,
	making every garbage collection a stop-the-world event. Setting gcstoptheworld=2
	also disables concurrent sweeping after the garbage collection finishes.
	
	gctrace: setting gctrace=1 causes the garbage collector to emit a single line to standard
	error at each collection, summarizing the amount of memory collected and the
	length of the pause. The format of this line is subject to change.
	Currently, it is:
		gc # @#s #%: #+#+# ms clock, #+#/#/#+# ms cpu, #->#-># MB, # MB goal, # P
	where the fields are as follows:
		gc #        the GC number, incremented at each GC
		@#s         time in seconds since program start
		#%          percentage of time spent in GC since program start
		#+...+#     wall-clock/CPU times for the phases of the GC
		#->#-># MB  heap size at GC start, at GC end, and live heap
		# MB goal   goal heap size
		# P         number of processors used
	The phases are stop-the-world (STW) sweep termination, concurrent
	mark and scan, and STW mark termination. The CPU times
	for mark/scan are broken down in to assist time (GC performed in
	line with allocation), background GC time, and idle GC time.
	If the line ends with "(forced)", this GC was forced by a
	runtime.GC() call.
	
	Setting gctrace to any value > 0 also causes the garbage collector
	to emit a summary when memory is released back to the system.
	This process of returning memory to the system is called scavenging.
	The format of this summary is subject to change.
	Currently it is:
		scvg#: # MB released  printed only if non-zero
		scvg#: inuse: # idle: # sys: # released: # consumed: # (MB)
	where the fields are as follows:
		scvg#        the scavenge cycle number, incremented at each scavenge
		inuse: #     MB used or partially used spans
		idle: #      MB spans pending scavenging
		sys: #       MB mapped from the system
		released: #  MB released to the system
		consumed: #  MB allocated from the system
	
	madvdontneed: setting madvdontneed=1 will use MADV_DONTNEED
	instead of MADV_FREE on Linux when returning memory to the
	kernel. This is less efficient, but causes RSS numbers to drop
	more quickly.
	
	memprofilerate: setting memprofilerate=X will update the value of runtime.MemProfileRate.
	When set to 0 memory profiling is disabled.  Refer to the description of
	MemProfileRate for the default value.
	
	invalidptr: defaults to invalidptr=1, causing the garbage collector and stack
	copier to crash the program if an invalid pointer value (for example, 1)
	is found in a pointer-typed location. Setting invalidptr=0 disables this check.
	This should only be used as a temporary workaround to diagnose buggy code.
	The real fix is to not store integers in pointer-typed locations.
	
	sbrk: setting sbrk=1 replaces the memory allocator and garbage collector
	with a trivial allocator that obtains memory from the operating system and
	never reclaims any memory.
	
	scavenge: scavenge=1 enables debugging mode of heap scavenger.
	
	scheddetail: setting schedtrace=X and scheddetail=1 causes the scheduler to emit
	detailed multiline info every X milliseconds, describing state of the scheduler,
	processors, threads and goroutines.
	
	schedtrace: setting schedtrace=X causes the scheduler to emit a single line to standard
	error every X milliseconds, summarizing the scheduler state.
	
	tracebackancestors: setting tracebackancestors=N extends tracebacks with the stacks at
	which goroutines were created, where N limits the number of ancestor goroutines to
	report. This also extends the information returned by runtime.Stack. Ancestor's goroutine
	IDs will refer to the ID of the goroutine at the time of creation; it's possible for this
	ID to be reused for another goroutine. Setting N to 0 will report no ancestry information.

The net, net/http, and crypto/tls packages also refer to debugging variables in GODEBUG.
See the documentation for those packages for details.

The GOMAXPROCS variable limits the number of operating system threads that
can execute user-level Go code simultaneously. There is no limit to the number of threads
that can be blocked in system calls on behalf of Go code; those do not count against
the GOMAXPROCS limit. This package's GOMAXPROCS function queries and changes
the limit.

The GORACE variable configures the race detector, for programs built using -race.
See <a href="https://golang.org/doc/articles/race_detector.html">https://golang.org/doc/articles/race_detector.html</a> for details.

The GOTRACEBACK variable controls the amount of output generated when a Go
program fails due to an unrecovered panic or an unexpected runtime condition.
By default, a failure prints a stack trace for the current goroutine,
eliding functions internal to the run-time system, and then exits with exit code 2.
The failure prints stack traces for all goroutines if there is no current goroutine
or the failure is internal to the run-time.
GOTRACEBACK=none omits the goroutine stack traces entirely.
GOTRACEBACK=single (the default) behaves as described above.
GOTRACEBACK=all adds stack traces for all user-created goroutines.
GOTRACEBACK=system is like ``all'' but adds stack frames for run-time functions
and shows goroutines created internally by the run-time.
GOTRACEBACK=crash is like ``system'' but crashes in an operating system-specific
manner instead of exiting. For example, on Unix systems, the crash raises
SIGABRT to trigger a core dump.
For historical reasons, the GOTRACEBACK settings 0, 1, and 2 are synonyms for
none, all, and system, respectively.
The runtime/debug package's SetTraceback function allows increasing the
amount of output at run time, but it cannot reduce the amount below that
specified by the environment variable.
See <a href="https://golang.org/pkg/runtime/debug/#SetTraceback">https://golang.org/pkg/runtime/debug/#SetTraceback</a>.

The GOARCH, GOOS, GOPATH, and GOROOT environment variables complete
the set of Go environment variables. They influence the building of Go programs
(see <a href="https://golang.org/cmd/go">https://golang.org/cmd/go</a> and <a href="https://golang.org/pkg/go/build">https://golang.org/pkg/go/build</a>).
GOARCH, GOOS, and GOROOT are recorded at compile time and made available by
constants or functions in this package, but they do not influence the execution
of the run-time system.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func BlockProfile(p []BlockProfileRecord) (n int, ok bool)](#BlockProfile)
* [func Breakpoint()](#Breakpoint)
* [func CPUProfile() []byte](#CPUProfile)
* [func Caller(skip int) (pc uintptr, file string, line int, ok bool)](#Caller)
* [func Callers(skip int, pc []uintptr) int](#Callers)
* [func GC()](#GC)
* [func GOMAXPROCS(n int) int](#GOMAXPROCS)
* [func GOROOT() string](#GOROOT)
* [func Goexit()](#Goexit)
* [func GoroutineProfile(p []StackRecord) (n int, ok bool)](#GoroutineProfile)
* [func Gosched()](#Gosched)
* [func KeepAlive(x interface{})](#KeepAlive)
* [func LockOSThread()](#LockOSThread)
* [func MemProfile(p []MemProfileRecord, inuseZero bool) (n int, ok bool)](#MemProfile)
* [func MutexProfile(p []BlockProfileRecord) (n int, ok bool)](#MutexProfile)
* [func NumCPU() int](#NumCPU)
* [func NumCgoCall() int64](#NumCgoCall)
* [func NumGoroutine() int](#NumGoroutine)
* [func ReadMemStats(m *MemStats)](#ReadMemStats)
* [func ReadTrace() []byte](#ReadTrace)
* [func SetBlockProfileRate(rate int)](#SetBlockProfileRate)
* [func SetCPUProfileRate(hz int)](#SetCPUProfileRate)
* [func SetCgoTraceback(version int, traceback, context, symbolizer unsafe.Pointer)](#SetCgoTraceback)
* [func SetFinalizer(obj interface{}, finalizer interface{})](#SetFinalizer)
* [func SetMutexProfileFraction(rate int) int](#SetMutexProfileFraction)
* [func Stack(buf []byte, all bool) int](#Stack)
* [func StartTrace() error](#StartTrace)
* [func StopTrace()](#StopTrace)
* [func ThreadCreateProfile(p []StackRecord) (n int, ok bool)](#ThreadCreateProfile)
* [func UnlockOSThread()](#UnlockOSThread)
* [func Version() string](#Version)
* [type BlockProfileRecord](#BlockProfileRecord)
* [type Error](#Error)
* [type Frame](#Frame)
* [type Frames](#Frames)
  * [func CallersFrames(callers []uintptr) *Frames](#CallersFrames)
  * [func (ci *Frames) Next() (frame Frame, more bool)](#Frames.Next)
* [type Func](#Func)
  * [func FuncForPC(pc uintptr) *Func](#FuncForPC)
  * [func (f *Func) Entry() uintptr](#Func.Entry)
  * [func (f *Func) FileLine(pc uintptr) (file string, line int)](#Func.FileLine)
  * [func (f *Func) Name() string](#Func.Name)
* [type MemProfileRecord](#MemProfileRecord)
  * [func (r *MemProfileRecord) InUseBytes() int64](#MemProfileRecord.InUseBytes)
  * [func (r *MemProfileRecord) InUseObjects() int64](#MemProfileRecord.InUseObjects)
  * [func (r *MemProfileRecord) Stack() []uintptr](#MemProfileRecord.Stack)
* [type MemStats](#MemStats)
* [type StackRecord](#StackRecord)
  * [func (r *StackRecord) Stack() []uintptr](#StackRecord.Stack)
* [type TypeAssertionError](#TypeAssertionError)
  * [func (e *TypeAssertionError) Error() string](#TypeAssertionError.Error)
  * [func (*TypeAssertionError) RuntimeError()](#TypeAssertionError.RuntimeError)


#### <a id="pkg-examples">Examples</a>
* [Frames](#example_Frames)


#### <a id="pkg-files">Package files</a>
[alg.go](https://golang.org/src/runtime/alg.go) [atomic_pointer.go](https://golang.org/src/runtime/atomic_pointer.go) [cgo.go](https://golang.org/src/runtime/cgo.go) [cgo_mmap.go](https://golang.org/src/runtime/cgo_mmap.go) [cgo_sigaction.go](https://golang.org/src/runtime/cgo_sigaction.go) [cgocall.go](https://golang.org/src/runtime/cgocall.go) [cgocallback.go](https://golang.org/src/runtime/cgocallback.go) [cgocheck.go](https://golang.org/src/runtime/cgocheck.go) [chan.go](https://golang.org/src/runtime/chan.go) [compiler.go](https://golang.org/src/runtime/compiler.go) [complex.go](https://golang.org/src/runtime/complex.go) [cpuflags.go](https://golang.org/src/runtime/cpuflags.go) [cpuflags_amd64.go](https://golang.org/src/runtime/cpuflags_amd64.go) [cpuprof.go](https://golang.org/src/runtime/cpuprof.go) [cputicks.go](https://golang.org/src/runtime/cputicks.go) [debug.go](https://golang.org/src/runtime/debug.go) [debugcall.go](https://golang.org/src/runtime/debugcall.go) [debuglog.go](https://golang.org/src/runtime/debuglog.go) [debuglog_off.go](https://golang.org/src/runtime/debuglog_off.go) [defs_linux_amd64.go](https://golang.org/src/runtime/defs_linux_amd64.go) [env_posix.go](https://golang.org/src/runtime/env_posix.go) [error.go](https://golang.org/src/runtime/error.go) [extern.go](https://golang.org/src/runtime/extern.go) [fastlog2.go](https://golang.org/src/runtime/fastlog2.go) [fastlog2table.go](https://golang.org/src/runtime/fastlog2table.go) [float.go](https://golang.org/src/runtime/float.go) [hash64.go](https://golang.org/src/runtime/hash64.go) [heapdump.go](https://golang.org/src/runtime/heapdump.go) [iface.go](https://golang.org/src/runtime/iface.go) [lfstack.go](https://golang.org/src/runtime/lfstack.go) [lfstack_64bit.go](https://golang.org/src/runtime/lfstack_64bit.go) [lock_futex.go](https://golang.org/src/runtime/lock_futex.go) [malloc.go](https://golang.org/src/runtime/malloc.go) [map.go](https://golang.org/src/runtime/map.go) [map_fast32.go](https://golang.org/src/runtime/map_fast32.go) [map_fast64.go](https://golang.org/src/runtime/map_fast64.go) [map_faststr.go](https://golang.org/src/runtime/map_faststr.go) [mbarrier.go](https://golang.org/src/runtime/mbarrier.go) [mbitmap.go](https://golang.org/src/runtime/mbitmap.go) [mcache.go](https://golang.org/src/runtime/mcache.go) [mcentral.go](https://golang.org/src/runtime/mcentral.go) [mem_linux.go](https://golang.org/src/runtime/mem_linux.go) [mfinal.go](https://golang.org/src/runtime/mfinal.go) [mfixalloc.go](https://golang.org/src/runtime/mfixalloc.go) [mgc.go](https://golang.org/src/runtime/mgc.go) [mgclarge.go](https://golang.org/src/runtime/mgclarge.go) [mgcmark.go](https://golang.org/src/runtime/mgcmark.go) [mgcscavenge.go](https://golang.org/src/runtime/mgcscavenge.go) [mgcstack.go](https://golang.org/src/runtime/mgcstack.go) [mgcsweep.go](https://golang.org/src/runtime/mgcsweep.go) [mgcsweepbuf.go](https://golang.org/src/runtime/mgcsweepbuf.go) [mgcwork.go](https://golang.org/src/runtime/mgcwork.go) [mheap.go](https://golang.org/src/runtime/mheap.go) [mprof.go](https://golang.org/src/runtime/mprof.go) [msan0.go](https://golang.org/src/runtime/msan0.go) [msize.go](https://golang.org/src/runtime/msize.go) [mstats.go](https://golang.org/src/runtime/mstats.go) [mwbbuf.go](https://golang.org/src/runtime/mwbbuf.go) [netpoll.go](https://golang.org/src/runtime/netpoll.go) [netpoll_epoll.go](https://golang.org/src/runtime/netpoll_epoll.go) [os_linux.go](https://golang.org/src/runtime/os_linux.go) [os_linux_generic.go](https://golang.org/src/runtime/os_linux_generic.go) [os_linux_noauxv.go](https://golang.org/src/runtime/os_linux_noauxv.go) [os_nonopenbsd.go](https://golang.org/src/runtime/os_nonopenbsd.go) [panic.go](https://golang.org/src/runtime/panic.go) [plugin.go](https://golang.org/src/runtime/plugin.go) [print.go](https://golang.org/src/runtime/print.go) [proc.go](https://golang.org/src/runtime/proc.go) [profbuf.go](https://golang.org/src/runtime/profbuf.go) [proflabel.go](https://golang.org/src/runtime/proflabel.go) [race0.go](https://golang.org/src/runtime/race0.go) [rdebug.go](https://golang.org/src/runtime/rdebug.go) [relax_stub.go](https://golang.org/src/runtime/relax_stub.go) [runtime.go](https://golang.org/src/runtime/runtime.go) [runtime1.go](https://golang.org/src/runtime/runtime1.go) [runtime2.go](https://golang.org/src/runtime/runtime2.go) [rwmutex.go](https://golang.org/src/runtime/rwmutex.go) [select.go](https://golang.org/src/runtime/select.go) [sema.go](https://golang.org/src/runtime/sema.go) [signal_amd64x.go](https://golang.org/src/runtime/signal_amd64x.go) [signal_linux_amd64.go](https://golang.org/src/runtime/signal_linux_amd64.go) [signal_sighandler.go](https://golang.org/src/runtime/signal_sighandler.go) [signal_unix.go](https://golang.org/src/runtime/signal_unix.go) [sigqueue.go](https://golang.org/src/runtime/sigqueue.go) [sigqueue_note.go](https://golang.org/src/runtime/sigqueue_note.go) [sigtab_linux_generic.go](https://golang.org/src/runtime/sigtab_linux_generic.go) [sizeclasses.go](https://golang.org/src/runtime/sizeclasses.go) [slice.go](https://golang.org/src/runtime/slice.go) [softfloat64.go](https://golang.org/src/runtime/softfloat64.go) [stack.go](https://golang.org/src/runtime/stack.go) [string.go](https://golang.org/src/runtime/string.go) [stubs.go](https://golang.org/src/runtime/stubs.go) [stubs2.go](https://golang.org/src/runtime/stubs2.go) [stubs3.go](https://golang.org/src/runtime/stubs3.go) [stubs_amd64x.go](https://golang.org/src/runtime/stubs_amd64x.go) [stubs_linux.go](https://golang.org/src/runtime/stubs_linux.go) [symtab.go](https://golang.org/src/runtime/symtab.go) [sys_nonppc64x.go](https://golang.org/src/runtime/sys_nonppc64x.go) [sys_x86.go](https://golang.org/src/runtime/sys_x86.go) [time.go](https://golang.org/src/runtime/time.go) [timestub.go](https://golang.org/src/runtime/timestub.go) [timestub2.go](https://golang.org/src/runtime/timestub2.go) [trace.go](https://golang.org/src/runtime/trace.go) [traceback.go](https://golang.org/src/runtime/traceback.go) [type.go](https://golang.org/src/runtime/type.go) [typekind.go](https://golang.org/src/runtime/typekind.go) [utf8.go](https://golang.org/src/runtime/utf8.go) [vdso_elf64.go](https://golang.org/src/runtime/vdso_elf64.go) [vdso_linux.go](https://golang.org/src/runtime/vdso_linux.go) [vdso_linux_amd64.go](https://golang.org/src/runtime/vdso_linux_amd64.go) [write_err.go](https://golang.org/src/runtime/write_err.go) 


## <a id="pkg-constants">Constants</a>
Compiler is the name of the compiler toolchain that built the
running binary. Known toolchains are:


	gc      Also known as cmd/compile.
	gccgo   The gccgo front end, part of the GCC compiler suite.


<pre>const <span id="Compiler">Compiler</span> = &#34;gc&#34;</pre>GOARCH is the running program's architecture target:
one of 386, amd64, arm, s390x, and so on.


<pre>const <span id="GOARCH">GOARCH</span> <a href="/pkg/builtin/#string">string</a> = <a href="/pkg/runtime/internal/sys/">sys</a>.<a href="/pkg/runtime/internal/sys/#GOARCH">GOARCH</a></pre>GOOS is the running program's operating system target:
one of darwin, freebsd, linux, and so on.
To view possible combinations of GOOS and GOARCH, run "go tool dist list".


<pre>const <span id="GOOS">GOOS</span> <a href="/pkg/builtin/#string">string</a> = <a href="/pkg/runtime/internal/sys/">sys</a>.<a href="/pkg/runtime/internal/sys/#GOOS">GOOS</a></pre>

## <a id="pkg-variables">Variables</a>
MemProfileRate controls the fraction of memory allocations
that are recorded and reported in the memory profile.
The profiler aims to sample an average of
one allocation per MemProfileRate bytes allocated.

To include every allocated block in the profile, set MemProfileRate to 1.
To turn off profiling entirely, set MemProfileRate to 0.

The tools that process the memory profiles assume that the
profile rate is constant across the lifetime of the program
and equal to the current value. Programs that change the
memory profiling rate should do so just once, as early as
possible in the execution of the program (for example,
at the beginning of main).


<pre>var <span id="MemProfileRate">MemProfileRate</span> <a href="/pkg/builtin/#int">int</a> = 512 * 1024</pre>

## <a id="BlockProfile">func</a> [BlockProfile](https://golang.org/src/runtime/mprof.go?s=19026:19084#L624)
<pre>func BlockProfile(p []<a href="#BlockProfileRecord">BlockProfileRecord</a>) (n <a href="/pkg/builtin/#int">int</a>, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
BlockProfile returns n, the number of records in the current blocking profile.
If len(p) >= n, BlockProfile copies the profile into p and returns n, true.
If len(p) < n, BlockProfile does not change p and returns n, false.

Most clients should use the runtime/pprof package or
the testing package's -test.blockprofile flag instead
of calling BlockProfile directly.



## <a id="Breakpoint">func</a> [Breakpoint](https://golang.org/src/runtime/proc.go?s=99618:99635#L3489)
<pre>func Breakpoint()</pre>
Breakpoint executes a breakpoint trap.



## <a id="CPUProfile">func</a> [CPUProfile](https://golang.org/src/runtime/cpuprof.go?s=5826:5850#L178)
<pre>func CPUProfile() []<a href="/pkg/builtin/#byte">byte</a></pre>
CPUProfile panics.
It formerly provided raw access to chunks of
a pprof-format profile generated by the runtime.
The details of generating that format have changed,
so this functionality has been removed.

Deprecated: Use the runtime/pprof package,
or the handlers in the net/http/pprof package,
or the testing package's -test.cpuprofile flag instead.



## <a id="Caller">func</a> [Caller](https://golang.org/src/runtime/extern.go?s=9344:9410#L170)
<pre>func Caller(skip <a href="/pkg/builtin/#int">int</a>) (pc <a href="/pkg/builtin/#uintptr">uintptr</a>, file <a href="/pkg/builtin/#string">string</a>, line <a href="/pkg/builtin/#int">int</a>, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
Caller reports file and line number information about function invocations on
the calling goroutine's stack. The argument skip is the number of stack frames
to ascend, with 0 identifying the caller of Caller.  (For historical reasons the
meaning of skip differs between Caller and Callers.) The return values report the
program counter, file name, and line number within the file of the corresponding
call. The boolean ok is false if it was not possible to recover the information.



## <a id="Callers">func</a> [Callers](https://golang.org/src/runtime/extern.go?s=10396:10436#L194)
<pre>func Callers(skip <a href="/pkg/builtin/#int">int</a>, pc []<a href="/pkg/builtin/#uintptr">uintptr</a>) <a href="/pkg/builtin/#int">int</a></pre>
Callers fills the slice pc with the return program counters of function invocations
on the calling goroutine's stack. The argument skip is the number of stack frames
to skip before recording in pc, with 0 identifying the frame for Callers itself and
1 identifying the caller of Callers.
It returns the number of entries written to pc.

To translate these PCs into symbolic information such as function
names and line numbers, use CallersFrames. CallersFrames accounts
for inlined functions and adjusts the return program counters into
call program counters. Iterating over the returned slice of PCs
directly is discouraged, as is using FuncForPC on any of the
returned PCs, since these cannot account for inlining or return
program counter adjustment.
go:noinline



## <a id="GC">func</a> [GC](https://golang.org/src/runtime/mgc.go?s=37788:37797#L1024)
<pre>func GC()</pre>
GC runs a garbage collection and blocks the caller until the
garbage collection is complete. It may also block the entire
program.



## <a id="GOMAXPROCS">func</a> [GOMAXPROCS](https://golang.org/src/runtime/debug.go?s=533:559#L7)
<pre>func GOMAXPROCS(n <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
GOMAXPROCS sets the maximum number of CPUs that can be executing
simultaneously and returns the previous setting. If n < 1, it does not
change the current setting.
The number of logical CPUs on the local machine can be queried with NumCPU.
This call will go away when the scheduler improves.



## <a id="GOROOT">func</a> [GOROOT](https://golang.org/src/runtime/extern.go?s=10815:10835#L207)
<pre>func GOROOT() <a href="/pkg/builtin/#string">string</a></pre>
GOROOT returns the root of the Go tree. It uses the
GOROOT environment variable, if set at process start,
or else the root used during the Go build.



## <a id="Goexit">func</a> [Goexit](https://golang.org/src/runtime/panic.go?s=17291:17304#L532)
<pre>func Goexit()</pre>
Goexit terminates the goroutine that calls it. No other goroutine is affected.
Goexit runs all deferred calls before terminating the goroutine. Because Goexit
is not a panic, any recover calls in those deferred functions will return nil.

Calling Goexit from the main goroutine terminates that goroutine
without func main returning. Since func main has not returned,
the program continues execution of other goroutines.
If all other goroutines exit, the program crashes.



## <a id="GoroutineProfile">func</a> [GoroutineProfile](https://golang.org/src/runtime/mprof.go?s=21480:21535#L710)
<pre>func GoroutineProfile(p []<a href="#StackRecord">StackRecord</a>) (n <a href="/pkg/builtin/#int">int</a>, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
GoroutineProfile returns n, the number of records in the active goroutine stack profile.
If len(p) >= n, GoroutineProfile copies the profile into p and returns n, true.
If len(p) < n, GoroutineProfile does not change p and returns n, false.

Most clients should use the runtime/pprof package instead
of calling GoroutineProfile directly.



## <a id="Gosched">func</a> [Gosched](https://golang.org/src/runtime/proc.go?s=8875:8889#L257)
<pre>func Gosched()</pre>
Gosched yields the processor, allowing other goroutines to run. It does not
suspend the current goroutine, so execution resumes automatically.



## <a id="KeepAlive">func</a> [KeepAlive](https://golang.org/src/runtime/mfinal.go?s=14896:14925#L436)
<pre>func KeepAlive(x interface{})</pre>
KeepAlive marks its argument as currently reachable.
This ensures that the object is not freed, and its finalizer is not run,
before the point in the program where KeepAlive is called.

A very simplified example showing where KeepAlive is required:


	type File struct { d int }
	d, err := syscall.Open("/file/path", syscall.O_RDONLY, 0)
	// ... do something if err != nil ...
	p := &File{d}
	runtime.SetFinalizer(p, func(p *File) { syscall.Close(p.d) })
	var buf [10]byte
	n, err := syscall.Read(p.d, buf[:])
	// Ensure p is not finalized until Read returns.
	runtime.KeepAlive(p)
	// No more uses of p after this point.

Without the KeepAlive call, the finalizer could run at the start of
syscall.Read, closing the file descriptor before syscall.Read makes
the actual system call.



## <a id="LockOSThread">func</a> [LockOSThread](https://golang.org/src/runtime/proc.go?s=100728:100747#L3522)
<pre>func LockOSThread()</pre>
LockOSThread wires the calling goroutine to its current operating system thread.
The calling goroutine will always execute in that thread,
and no other goroutine will execute in it,
until the calling goroutine has made as many calls to
UnlockOSThread as to LockOSThread.
If the calling goroutine exits without unlocking the thread,
the thread will be terminated.

All init functions are run on the startup thread. Calling LockOSThread
from an init function will cause the main function to be invoked on
that thread.

A goroutine should call LockOSThread before calling OS services or
non-Go library functions that depend on per-thread state.



## <a id="MemProfile">func</a> [MemProfile](https://golang.org/src/runtime/mprof.go?s=16356:16426#L533)
<pre>func MemProfile(p []<a href="#MemProfileRecord">MemProfileRecord</a>, inuseZero <a href="/pkg/builtin/#bool">bool</a>) (n <a href="/pkg/builtin/#int">int</a>, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
MemProfile returns a profile of memory allocated and freed per allocation
site.

MemProfile returns n, the number of records in the current memory profile.
If len(p) >= n, MemProfile copies the profile into p and returns n, true.
If len(p) < n, MemProfile does not change p and returns n, false.

If inuseZero is true, the profile includes allocation records
where r.AllocBytes > 0 but r.AllocBytes == r.FreeBytes.
These are sites where memory was allocated, but it has all
been released back to the runtime.

The returned profile may be up to two garbage collection cycles old.
This is to avoid skewing the profile toward allocations; because
allocations happen in real time but frees are delayed until the garbage
collector performs sweeping, the profile only accounts for allocations
that have had a chance to be freed by the garbage collector.

Most clients should use the runtime/pprof package or
the testing package's -test.memprofile flag instead
of calling MemProfile directly.



## <a id="MutexProfile">func</a> [MutexProfile](https://golang.org/src/runtime/mprof.go?s=20015:20073#L659)
<pre>func MutexProfile(p []<a href="#BlockProfileRecord">BlockProfileRecord</a>) (n <a href="/pkg/builtin/#int">int</a>, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
MutexProfile returns n, the number of records in the current mutex profile.
If len(p) >= n, MutexProfile copies the profile into p and returns n, true.
Otherwise, MutexProfile does not change p, and returns n, false.

Most clients should use the runtime/pprof package
instead of calling MutexProfile directly.



## <a id="NumCPU">func</a> [NumCPU](https://golang.org/src/runtime/debug.go?s=1169:1186#L33)
<pre>func NumCPU() <a href="/pkg/builtin/#int">int</a></pre>
NumCPU returns the number of logical CPUs usable by the current process.

The set of available CPUs is checked by querying the operating system
at process startup. Changes to operating system CPU allocation after
process startup are not reflected.



## <a id="NumCgoCall">func</a> [NumCgoCall](https://golang.org/src/runtime/debug.go?s=1285:1308#L38)
<pre>func NumCgoCall() <a href="/pkg/builtin/#int64">int64</a></pre>
NumCgoCall returns the number of cgo calls made by the current process.



## <a id="NumGoroutine">func</a> [NumGoroutine](https://golang.org/src/runtime/debug.go?s=1520:1543#L47)
<pre>func NumGoroutine() <a href="/pkg/builtin/#int">int</a></pre>
NumGoroutine returns the number of goroutines that currently exist.



## <a id="ReadMemStats">func</a> [ReadMemStats](https://golang.org/src/runtime/mstats.go?s=15992:16022#L433)
<pre>func ReadMemStats(m *<a href="#MemStats">MemStats</a>)</pre>
ReadMemStats populates m with memory allocator statistics.

The returned memory allocator statistics are up to date as of the
call to ReadMemStats. This is in contrast with a heap profile,
which is a snapshot as of the most recently completed garbage
collection cycle.



## <a id="ReadTrace">func</a> [ReadTrace](https://golang.org/src/runtime/trace.go?s=15491:15514#L350)
<pre>func ReadTrace() []<a href="/pkg/builtin/#byte">byte</a></pre>
ReadTrace returns the next chunk of binary tracing data, blocking until data
is available. If tracing is turned off and all the data accumulated while it
was on has been returned, ReadTrace returns nil. The caller must copy the
returned data before calling ReadTrace again.
ReadTrace must be called from one goroutine at a time.



## <a id="SetBlockProfileRate">func</a> [SetBlockProfileRate](https://golang.org/src/runtime/mprof.go?s=11147:11181#L370)
<pre>func SetBlockProfileRate(rate <a href="/pkg/builtin/#int">int</a>)</pre>
SetBlockProfileRate controls the fraction of goroutine blocking events
that are reported in the blocking profile. The profiler aims to sample
an average of one blocking event per rate nanoseconds spent blocked.

To include every blocking event in the profile, pass rate = 1.
To turn off profiling entirely, pass rate <= 0.



## <a id="SetCPUProfileRate">func</a> [SetCPUProfileRate](https://golang.org/src/runtime/cpuprof.go?s=1880:1910#L44)
<pre>func SetCPUProfileRate(hz <a href="/pkg/builtin/#int">int</a>)</pre>
SetCPUProfileRate sets the CPU profiling rate to hz samples per second.
If hz <= 0, SetCPUProfileRate turns off profiling.
If the profiler is on, the rate cannot be changed without first turning it off.

Most clients should use the runtime/pprof package or
the testing package's -test.cpuprofile flag instead of calling
SetCPUProfileRate directly.



## <a id="SetCgoTraceback">func</a> [SetCgoTraceback](https://golang.org/src/runtime/traceback.go?s=40862:40942#L1179)
<pre>func SetCgoTraceback(version <a href="/pkg/builtin/#int">int</a>, traceback, context, symbolizer <a href="/pkg/unsafe/">unsafe</a>.<a href="/pkg/unsafe/#Pointer">Pointer</a>)</pre>
SetCgoTraceback records three C functions to use to gather
traceback information from C code and to convert that traceback
information into symbolic information. These are used when printing
stack traces for a program that uses cgo.

The traceback and context functions may be called from a signal
handler, and must therefore use only async-signal safe functions.
The symbolizer function may be called while the program is
crashing, and so must be cautious about using memory.  None of the
functions may call back into Go.

The context function will be called with a single argument, a
pointer to a struct:


	struct {
		Context uintptr
	}

In C syntax, this struct will be


	struct {
		uintptr_t Context;
	};

If the Context field is 0, the context function is being called to
record the current traceback context. It should record in the
Context field whatever information is needed about the current
point of execution to later produce a stack trace, probably the
stack pointer and PC. In this case the context function will be
called from C code.

If the Context field is not 0, then it is a value returned by a
previous call to the context function. This case is called when the
context is no longer needed; that is, when the Go code is returning
to its C code caller. This permits the context function to release
any associated resources.

While it would be correct for the context function to record a
complete a stack trace whenever it is called, and simply copy that
out in the traceback function, in a typical program the context
function will be called many times without ever recording a
traceback for that context. Recording a complete stack trace in a
call to the context function is likely to be inefficient.

The traceback function will be called with a single argument, a
pointer to a struct:


	struct {
		Context    uintptr
		SigContext uintptr
		Buf        *uintptr
		Max        uintptr
	}

In C syntax, this struct will be


	struct {
		uintptr_t  Context;
		uintptr_t  SigContext;
		uintptr_t* Buf;
		uintptr_t  Max;
	};

The Context field will be zero to gather a traceback from the
current program execution point. In this case, the traceback
function will be called from C code.

Otherwise Context will be a value previously returned by a call to
the context function. The traceback function should gather a stack
trace from that saved point in the program execution. The traceback
function may be called from an execution thread other than the one
that recorded the context, but only when the context is known to be
valid and unchanging. The traceback function may also be called
deeper in the call stack on the same thread that recorded the
context. The traceback function may be called multiple times with
the same Context value; it will usually be appropriate to cache the
result, if possible, the first time this is called for a specific
context value.

If the traceback function is called from a signal handler on a Unix
system, SigContext will be the signal context argument passed to
the signal handler (a C ucontext_t* cast to uintptr_t). This may be
used to start tracing at the point where the signal occurred. If
the traceback function is not called from a signal handler,
SigContext will be zero.

Buf is where the traceback information should be stored. It should
be PC values, such that Buf[0] is the PC of the caller, Buf[1] is
the PC of that function's caller, and so on.  Max is the maximum
number of entries to store.  The function should store a zero to
indicate the top of the stack, or that the caller is on a different
stack, presumably a Go stack.

Unlike runtime.Callers, the PC values returned should, when passed
to the symbolizer function, return the file/line of the call
instruction.  No additional subtraction is required or appropriate.

On all platforms, the traceback function is invoked when a call from
Go to C to Go requests a stack trace. On linux/amd64, linux/ppc64le,
and freebsd/amd64, the traceback function is also invoked when a
signal is received by a thread that is executing a cgo call. The
traceback function should not make assumptions about when it is
called, as future versions of Go may make additional calls.

The symbolizer function will be called with a single argument, a
pointer to a struct:


	struct {
		PC      uintptr // program counter to fetch information for
		File    *byte   // file name (NUL terminated)
		Lineno  uintptr // line number
		Func    *byte   // function name (NUL terminated)
		Entry   uintptr // function entry point
		More    uintptr // set non-zero if more info for this PC
		Data    uintptr // unused by runtime, available for function
	}

In C syntax, this struct will be


	struct {
		uintptr_t PC;
		char*     File;
		uintptr_t Lineno;
		char*     Func;
		uintptr_t Entry;
		uintptr_t More;
		uintptr_t Data;
	};

The PC field will be a value returned by a call to the traceback
function.

The first time the function is called for a particular traceback,
all the fields except PC will be 0. The function should fill in the
other fields if possible, setting them to 0/nil if the information
is not available. The Data field may be used to store any useful
information across calls. The More field should be set to non-zero
if there is more information for this PC, zero otherwise. If More
is set non-zero, the function will be called again with the same
PC, and may return different information (this is intended for use
with inlined functions). If More is zero, the function will be
called with the next PC value in the traceback. When the traceback
is complete, the function will be called once more with PC set to
zero; this may be used to free any information. Each call will
leave the fields of the struct set to the same values they had upon
return, except for the PC field when the More field is zero. The
function must not keep a copy of the struct pointer between calls.

When calling SetCgoTraceback, the version argument is the version
number of the structs that the functions expect to receive.
Currently this must be zero.

The symbolizer function may be nil, in which case the results of
the traceback function will be displayed as numbers. If the
traceback function is nil, the symbolizer function will never be
called. The context function may be nil, in which case the
traceback function will only be called with the context field set
to zero.  If the context function is nil, then calls from Go to C
to Go will not show a traceback for the C portion of the call stack.

SetCgoTraceback should be called only once, ideally from an init function.



## <a id="SetFinalizer">func</a> [SetFinalizer](https://golang.org/src/runtime/mfinal.go?s=10402:10459#L299)
<pre>func SetFinalizer(obj interface{}, finalizer interface{})</pre>
SetFinalizer sets the finalizer associated with obj to the provided
finalizer function. When the garbage collector finds an unreachable block
with an associated finalizer, it clears the association and runs
finalizer(obj) in a separate goroutine. This makes obj reachable again,
but now without an associated finalizer. Assuming that SetFinalizer
is not called again, the next time the garbage collector sees
that obj is unreachable, it will free obj.

SetFinalizer(obj, nil) clears any finalizer associated with obj.

The argument obj must be a pointer to an object allocated by calling
new, by taking the address of a composite literal, or by taking the
address of a local variable.
The argument finalizer must be a function that takes a single argument
to which obj's type can be assigned, and can have arbitrary ignored return
values. If either of these is not true, SetFinalizer may abort the
program.

Finalizers are run in dependency order: if A points at B, both have
finalizers, and they are otherwise unreachable, only the finalizer
for A runs; once A is freed, the finalizer for B can run.
If a cyclic structure includes a block with a finalizer, that
cycle is not guaranteed to be garbage collected and the finalizer
is not guaranteed to run, because there is no ordering that
respects the dependencies.

The finalizer is scheduled to run at some arbitrary time after the
program can no longer reach the object to which obj points.
There is no guarantee that finalizers will run before a program exits,
so typically they are useful only for releasing non-memory resources
associated with an object during a long-running program.
For example, an os.File object could use a finalizer to close the
associated operating system file descriptor when a program discards
an os.File without calling Close, but it would be a mistake
to depend on a finalizer to flush an in-memory I/O buffer such as a
bufio.Writer, because the buffer would not be flushed at program exit.

It is not guaranteed that a finalizer will run if the size of *obj is
zero bytes.

It is not guaranteed that a finalizer will run for objects allocated
in initializers for package-level variables. Such objects may be
linker-allocated, not heap-allocated.

A finalizer may run as soon as an object becomes unreachable.
In order to use finalizers correctly, the program must ensure that
the object is reachable until it is no longer required.
Objects stored in global variables, or that can be found by tracing
pointers from a global variable, are reachable. For other objects,
pass the object to a call of the KeepAlive function to mark the
last point in the function where the object must be reachable.

For example, if p points to a struct that contains a file descriptor d,
and p has a finalizer that closes that file descriptor, and if the last
use of p in a function is a call to syscall.Write(p.d, buf, size), then
p may be unreachable as soon as the program enters syscall.Write. The
finalizer may run at that moment, closing p.d, causing syscall.Write
to fail because it is writing to a closed file descriptor (or, worse,
to an entirely different file descriptor opened by a different goroutine).
To avoid this problem, call runtime.KeepAlive(p) after the call to
syscall.Write.

A single goroutine runs all finalizers for a program, sequentially.
If a finalizer must run for a long time, it should do so by starting
a new goroutine.



## <a id="SetMutexProfileFraction">func</a> [SetMutexProfileFraction](https://golang.org/src/runtime/mprof.go?s=12654:12696#L429)
<pre>func SetMutexProfileFraction(rate <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
SetMutexProfileFraction controls the fraction of mutex contention events
that are reported in the mutex profile. On average 1/rate events are
reported. The previous rate is returned.

To turn off profiling entirely, pass rate 0.
To just read the current rate, pass rate < 0.
(For n>1 the details of sampling may change.)



## <a id="Stack">func</a> [Stack](https://golang.org/src/runtime/mprof.go?s=22812:22848#L770)
<pre>func Stack(buf []<a href="/pkg/builtin/#byte">byte</a>, all <a href="/pkg/builtin/#bool">bool</a>) <a href="/pkg/builtin/#int">int</a></pre>
Stack formats a stack trace of the calling goroutine into buf
and returns the number of bytes written to buf.
If all is true, Stack formats stack traces of all other goroutines
into buf after the trace for the current goroutine.



## <a id="StartTrace">func</a> [StartTrace](https://golang.org/src/runtime/trace.go?s=9984:10007#L172)
<pre>func StartTrace() <a href="/pkg/builtin/#error">error</a></pre>
StartTrace enables tracing for the current process.
While tracing, the data will be buffered and available via ReadTrace.
StartTrace returns an error if tracing is already enabled.
Most clients should use the runtime/trace package or the testing package's
-test.trace flag instead of calling StartTrace directly.



## <a id="StopTrace">func</a> [StopTrace](https://golang.org/src/runtime/trace.go?s=13236:13252#L263)
<pre>func StopTrace()</pre>
StopTrace stops tracing, if it was previously enabled.
StopTrace only returns after all the reads for the trace have completed.



## <a id="ThreadCreateProfile">func</a> [ThreadCreateProfile](https://golang.org/src/runtime/mprof.go?s=20812:20870#L688)
<pre>func ThreadCreateProfile(p []<a href="#StackRecord">StackRecord</a>) (n <a href="/pkg/builtin/#int">int</a>, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
ThreadCreateProfile returns n, the number of records in the thread creation profile.
If len(p) >= n, ThreadCreateProfile copies the profile into p and returns n, true.
If len(p) < n, ThreadCreateProfile does not change p and returns n, false.

Most clients should use the runtime/pprof package instead
of calling ThreadCreateProfile directly.



## <a id="UnlockOSThread">func</a> [UnlockOSThread](https://golang.org/src/runtime/proc.go?s=102346:102367#L3574)
<pre>func UnlockOSThread()</pre>
UnlockOSThread undoes an earlier call to LockOSThread.
If this drops the number of active LockOSThread calls on the
calling goroutine to zero, it unwires the calling goroutine from
its fixed operating system thread.
If there are no active LockOSThread calls, this is a no-op.

Before calling UnlockOSThread, the caller must ensure that the OS
thread is suitable for running other goroutines. If the caller made
any permanent changes to the state of the thread that would affect
other goroutines, it should not call this function and thus leave
the goroutine locked to the OS thread until the goroutine (and
hence the thread) exits.



## <a id="Version">func</a> [Version](https://golang.org/src/runtime/extern.go?s=11085:11106#L218)
<pre>func Version() <a href="/pkg/builtin/#string">string</a></pre>
Version returns the Go tree's version string.
It is either the commit hash and date at the time of the build or,
when possible, a release tag like "go1.3".





## <a id="BlockProfileRecord">type</a> [BlockProfileRecord](https://golang.org/src/runtime/mprof.go?s=18564:18639#L611)
BlockProfileRecord describes blocking events originated
at a particular call sequence (stack trace).


<pre>type BlockProfileRecord struct {
<span id="BlockProfileRecord.Count"></span>    Count  <a href="/pkg/builtin/#int64">int64</a>
<span id="BlockProfileRecord.Cycles"></span>    Cycles <a href="/pkg/builtin/#int64">int64</a>
    <a href="#StackRecord">StackRecord</a>
}
</pre>











## <a id="Error">type</a> [Error](https://golang.org/src/runtime/error.go?s=256:492#L1)
The Error interface identifies a run time error.


<pre>type Error interface {
    <a href="/pkg/builtin/#error">error</a>

    <span class="comment">// RuntimeError is a no-op function but</span>
    <span class="comment">// serves to distinguish types that are run time</span>
    <span class="comment">// errors from ordinary errors: a type is a</span>
    <span class="comment">// run time error if it has a RuntimeError method.</span>
    RuntimeError()
}</pre>











## <a id="Frame">type</a> [Frame](https://golang.org/src/runtime/symtab.go?s=647:1923#L15)
Frame is the information returned by Frames for each call frame.


<pre>type Frame struct {
<span id="Frame.PC"></span>    <span class="comment">// PC is the program counter for the location in this frame.</span>
    <span class="comment">// For a frame that calls another frame, this will be the</span>
    <span class="comment">// program counter of a call instruction. Because of inlining,</span>
    <span class="comment">// multiple frames may have the same PC value, but different</span>
    <span class="comment">// symbolic information.</span>
    PC <a href="/pkg/builtin/#uintptr">uintptr</a>

<span id="Frame.Func"></span>    <span class="comment">// Func is the Func value of this call frame. This may be nil</span>
    <span class="comment">// for non-Go code or fully inlined functions.</span>
    Func *<a href="#Func">Func</a>

<span id="Frame.Function"></span>    <span class="comment">// Function is the package path-qualified function name of</span>
    <span class="comment">// this call frame. If non-empty, this string uniquely</span>
    <span class="comment">// identifies a single function in the program.</span>
    <span class="comment">// This may be the empty string if not known.</span>
    <span class="comment">// If Func is not nil then Function == Func.Name().</span>
    Function <a href="/pkg/builtin/#string">string</a>

<span id="Frame.File"></span>    <span class="comment">// File and Line are the file name and line number of the</span>
    <span class="comment">// location in this frame. For non-leaf frames, this will be</span>
    <span class="comment">// the location of a call. These may be the empty string and</span>
    <span class="comment">// zero, respectively, if not known.</span>
    File <a href="/pkg/builtin/#string">string</a>
<span id="Frame.Line"></span>    Line <a href="/pkg/builtin/#int">int</a>

<span id="Frame.Entry"></span>    <span class="comment">// Entry point program counter for the function; may be zero</span>
    <span class="comment">// if not known. If Func is not nil then Entry ==</span>
    <span class="comment">// Func.Entry().</span>
    Entry <a href="/pkg/builtin/#uintptr">uintptr</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











## <a id="Frames">type</a> [Frames](https://golang.org/src/runtime/symtab.go?s=359:577#L5)
Frames may be used to get function/file/line information for a
slice of PC values returned by Callers.


<pre>type Frames struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Frames">Example</a>




### <a id="CallersFrames">func</a> [CallersFrames](https://golang.org/src/runtime/symtab.go?s=2110:2155#L55)
<pre>func CallersFrames(callers []<a href="/pkg/builtin/#uintptr">uintptr</a>) *<a href="#Frames">Frames</a></pre>
CallersFrames takes a slice of PC values returned by Callers and
prepares to return function/file/line information.
Do not change the slice until you are done with the Frames.






### <a id="Frames.Next">func</a> (\*Frames) [Next](https://golang.org/src/runtime/symtab.go?s=2362:2411#L63)
<pre>func (ci *<a href="#Frames">Frames</a>) Next() (frame <a href="#Frame">Frame</a>, more <a href="/pkg/builtin/#bool">bool</a>)</pre>
Next returns frame information for the next caller.
If more is false, there are no more callers (the Frame value is valid).




## <a id="Func">type</a> [Func](https://golang.org/src/runtime/symtab.go?s=6200:6281#L188)
A Func represents a Go function in the running binary.


<pre>type Func struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="FuncForPC">func</a> [FuncForPC](https://golang.org/src/runtime/symtab.go?s=15788:15820#L480)
<pre>func FuncForPC(pc <a href="/pkg/builtin/#uintptr">uintptr</a>) *<a href="#Func">Func</a></pre>
FuncForPC returns a *Func describing the function that contains the
given program counter address, or else nil.

If pc represents multiple functions because of inlining, it returns
the a *Func describing the innermost function, but with an entry
of the outermost function.






### <a id="Func.Entry">func</a> (\*Func) [Entry](https://golang.org/src/runtime/symtab.go?s=16962:16992#L520)
<pre>func (f *<a href="#Func">Func</a>) Entry() <a href="/pkg/builtin/#uintptr">uintptr</a></pre>
Entry returns the entry address of the function.




### <a id="Func.FileLine">func</a> (\*Func) [FileLine](https://golang.org/src/runtime/symtab.go?s=17321:17380#L533)
<pre>func (f *<a href="#Func">Func</a>) FileLine(pc <a href="/pkg/builtin/#uintptr">uintptr</a>) (file <a href="/pkg/builtin/#string">string</a>, line <a href="/pkg/builtin/#int">int</a>)</pre>
FileLine returns the file name and line number of the
source code corresponding to the program counter pc.
The result will not be accurate if pc is not a program
counter within f.




### <a id="Func.Name">func</a> (\*Func) [Name](https://golang.org/src/runtime/symtab.go?s=16702:16730#L507)
<pre>func (f *<a href="#Func">Func</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the name of the function.




## <a id="MemProfileRecord">type</a> [MemProfileRecord](https://golang.org/src/runtime/mprof.go?s=14465:14742#L487)
A MemProfileRecord describes the live objects allocated
by a particular call sequence (stack trace).


<pre>type MemProfileRecord struct {
<span id="MemProfileRecord.AllocBytes"></span>    AllocBytes, FreeBytes     <a href="/pkg/builtin/#int64">int64</a>       <span class="comment">// number of bytes allocated, freed</span>
<span id="MemProfileRecord.AllocObjects"></span>    AllocObjects, FreeObjects <a href="/pkg/builtin/#int64">int64</a>       <span class="comment">// number of objects allocated, freed</span>
<span id="MemProfileRecord.Stack0"></span>    Stack0                    [32]<a href="/pkg/builtin/#uintptr">uintptr</a> <span class="comment">// stack trace for this record; ends at first 0 entry</span>
}
</pre>











### <a id="MemProfileRecord.InUseBytes">func</a> (\*MemProfileRecord) [InUseBytes](https://golang.org/src/runtime/mprof.go?s=14819:14864#L494)
<pre>func (r *<a href="#MemProfileRecord">MemProfileRecord</a>) InUseBytes() <a href="/pkg/builtin/#int64">int64</a></pre>
InUseBytes returns the number of bytes in use (AllocBytes - FreeBytes).




### <a id="MemProfileRecord.InUseObjects">func</a> (\*MemProfileRecord) [InUseObjects](https://golang.org/src/runtime/mprof.go?s=14987:15034#L497)
<pre>func (r *<a href="#MemProfileRecord">MemProfileRecord</a>) InUseObjects() <a href="/pkg/builtin/#int64">int64</a></pre>
InUseObjects returns the number of objects in use (AllocObjects - FreeObjects).




### <a id="MemProfileRecord.Stack">func</a> (\*MemProfileRecord) [Stack](https://golang.org/src/runtime/mprof.go?s=15165:15209#L503)
<pre>func (r *<a href="#MemProfileRecord">MemProfileRecord</a>) Stack() []<a href="/pkg/builtin/#uintptr">uintptr</a></pre>
Stack returns the stack trace associated with the record,
a prefix of r.Stack0.




## <a id="MemStats">type</a> [MemStats](https://golang.org/src/runtime/mstats.go?s=5486:14902#L135)
A MemStats records statistics about the memory allocator.


<pre>type MemStats struct {

<span id="MemStats.Alloc"></span>    <span class="comment">// Alloc is bytes of allocated heap objects.</span>
    <span class="comment">//</span>
    <span class="comment">// This is the same as HeapAlloc (see below).</span>
    Alloc <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.TotalAlloc"></span>    <span class="comment">// TotalAlloc is cumulative bytes allocated for heap objects.</span>
    <span class="comment">//</span>
    <span class="comment">// TotalAlloc increases as heap objects are allocated, but</span>
    <span class="comment">// unlike Alloc and HeapAlloc, it does not decrease when</span>
    <span class="comment">// objects are freed.</span>
    TotalAlloc <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.Sys"></span>    <span class="comment">// Sys is the total bytes of memory obtained from the OS.</span>
    <span class="comment">//</span>
    <span class="comment">// Sys is the sum of the XSys fields below. Sys measures the</span>
    <span class="comment">// virtual address space reserved by the Go runtime for the</span>
    <span class="comment">// heap, stacks, and other internal data structures. It&#39;s</span>
    <span class="comment">// likely that not all of the virtual address space is backed</span>
    <span class="comment">// by physical memory at any given moment, though in general</span>
    <span class="comment">// it all was at some point.</span>
    Sys <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.Lookups"></span>    <span class="comment">// Lookups is the number of pointer lookups performed by the</span>
    <span class="comment">// runtime.</span>
    <span class="comment">//</span>
    <span class="comment">// This is primarily useful for debugging runtime internals.</span>
    Lookups <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.Mallocs"></span>    <span class="comment">// Mallocs is the cumulative count of heap objects allocated.</span>
    <span class="comment">// The number of live objects is Mallocs - Frees.</span>
    Mallocs <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.Frees"></span>    <span class="comment">// Frees is the cumulative count of heap objects freed.</span>
    Frees <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.HeapAlloc"></span>    <span class="comment">// HeapAlloc is bytes of allocated heap objects.</span>
    <span class="comment">//</span>
    <span class="comment">// &#34;Allocated&#34; heap objects include all reachable objects, as</span>
    <span class="comment">// well as unreachable objects that the garbage collector has</span>
    <span class="comment">// not yet freed. Specifically, HeapAlloc increases as heap</span>
    <span class="comment">// objects are allocated and decreases as the heap is swept</span>
    <span class="comment">// and unreachable objects are freed. Sweeping occurs</span>
    <span class="comment">// incrementally between GC cycles, so these two processes</span>
    <span class="comment">// occur simultaneously, and as a result HeapAlloc tends to</span>
    <span class="comment">// change smoothly (in contrast with the sawtooth that is</span>
    <span class="comment">// typical of stop-the-world garbage collectors).</span>
    HeapAlloc <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.HeapSys"></span>    <span class="comment">// HeapSys is bytes of heap memory obtained from the OS.</span>
    <span class="comment">//</span>
    <span class="comment">// HeapSys measures the amount of virtual address space</span>
    <span class="comment">// reserved for the heap. This includes virtual address space</span>
    <span class="comment">// that has been reserved but not yet used, which consumes no</span>
    <span class="comment">// physical memory, but tends to be small, as well as virtual</span>
    <span class="comment">// address space for which the physical memory has been</span>
    <span class="comment">// returned to the OS after it became unused (see HeapReleased</span>
    <span class="comment">// for a measure of the latter).</span>
    <span class="comment">//</span>
    <span class="comment">// HeapSys estimates the largest size the heap has had.</span>
    HeapSys <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.HeapIdle"></span>    <span class="comment">// HeapIdle is bytes in idle (unused) spans.</span>
    <span class="comment">//</span>
    <span class="comment">// Idle spans have no objects in them. These spans could be</span>
    <span class="comment">// (and may already have been) returned to the OS, or they can</span>
    <span class="comment">// be reused for heap allocations, or they can be reused as</span>
    <span class="comment">// stack memory.</span>
    <span class="comment">//</span>
    <span class="comment">// HeapIdle minus HeapReleased estimates the amount of memory</span>
    <span class="comment">// that could be returned to the OS, but is being retained by</span>
    <span class="comment">// the runtime so it can grow the heap without requesting more</span>
    <span class="comment">// memory from the OS. If this difference is significantly</span>
    <span class="comment">// larger than the heap size, it indicates there was a recent</span>
    <span class="comment">// transient spike in live heap size.</span>
    HeapIdle <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.HeapInuse"></span>    <span class="comment">// HeapInuse is bytes in in-use spans.</span>
    <span class="comment">//</span>
    <span class="comment">// In-use spans have at least one object in them. These spans</span>
    <span class="comment">// can only be used for other objects of roughly the same</span>
    <span class="comment">// size.</span>
    <span class="comment">//</span>
    <span class="comment">// HeapInuse minus HeapAlloc estimates the amount of memory</span>
    <span class="comment">// that has been dedicated to particular size classes, but is</span>
    <span class="comment">// not currently being used. This is an upper bound on</span>
    <span class="comment">// fragmentation, but in general this memory can be reused</span>
    <span class="comment">// efficiently.</span>
    HeapInuse <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.HeapReleased"></span>    <span class="comment">// HeapReleased is bytes of physical memory returned to the OS.</span>
    <span class="comment">//</span>
    <span class="comment">// This counts heap memory from idle spans that was returned</span>
    <span class="comment">// to the OS and has not yet been reacquired for the heap.</span>
    HeapReleased <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.HeapObjects"></span>    <span class="comment">// HeapObjects is the number of allocated heap objects.</span>
    <span class="comment">//</span>
    <span class="comment">// Like HeapAlloc, this increases as objects are allocated and</span>
    <span class="comment">// decreases as the heap is swept and unreachable objects are</span>
    <span class="comment">// freed.</span>
    HeapObjects <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.StackInuse"></span>    <span class="comment">// StackInuse is bytes in stack spans.</span>
    <span class="comment">//</span>
    <span class="comment">// In-use stack spans have at least one stack in them. These</span>
    <span class="comment">// spans can only be used for other stacks of the same size.</span>
    <span class="comment">//</span>
    <span class="comment">// There is no StackIdle because unused stack spans are</span>
    <span class="comment">// returned to the heap (and hence counted toward HeapIdle).</span>
    StackInuse <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.StackSys"></span>    <span class="comment">// StackSys is bytes of stack memory obtained from the OS.</span>
    <span class="comment">//</span>
    <span class="comment">// StackSys is StackInuse, plus any memory obtained directly</span>
    <span class="comment">// from the OS for OS thread stacks (which should be minimal).</span>
    StackSys <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.MSpanInuse"></span>    <span class="comment">// MSpanInuse is bytes of allocated mspan structures.</span>
    MSpanInuse <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.MSpanSys"></span>    <span class="comment">// MSpanSys is bytes of memory obtained from the OS for mspan</span>
    <span class="comment">// structures.</span>
    MSpanSys <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.MCacheInuse"></span>    <span class="comment">// MCacheInuse is bytes of allocated mcache structures.</span>
    MCacheInuse <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.MCacheSys"></span>    <span class="comment">// MCacheSys is bytes of memory obtained from the OS for</span>
    <span class="comment">// mcache structures.</span>
    MCacheSys <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.BuckHashSys"></span>    <span class="comment">// BuckHashSys is bytes of memory in profiling bucket hash tables.</span>
    BuckHashSys <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.GCSys"></span>    <span class="comment">// GCSys is bytes of memory in garbage collection metadata.</span>
    GCSys <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.OtherSys"></span>    <span class="comment">// OtherSys is bytes of memory in miscellaneous off-heap</span>
    <span class="comment">// runtime allocations.</span>
    OtherSys <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.NextGC"></span>    <span class="comment">// NextGC is the target heap size of the next GC cycle.</span>
    <span class="comment">//</span>
    <span class="comment">// The garbage collector&#39;s goal is to keep HeapAlloc ≤ NextGC.</span>
    <span class="comment">// At the end of each GC cycle, the target for the next cycle</span>
    <span class="comment">// is computed based on the amount of reachable data and the</span>
    <span class="comment">// value of GOGC.</span>
    NextGC <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.LastGC"></span>    <span class="comment">// LastGC is the time the last garbage collection finished, as</span>
    <span class="comment">// nanoseconds since 1970 (the UNIX epoch).</span>
    LastGC <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.PauseTotalNs"></span>    <span class="comment">// PauseTotalNs is the cumulative nanoseconds in GC</span>
    <span class="comment">// stop-the-world pauses since the program started.</span>
    <span class="comment">//</span>
    <span class="comment">// During a stop-the-world pause, all goroutines are paused</span>
    <span class="comment">// and only the garbage collector can run.</span>
    PauseTotalNs <a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.PauseNs"></span>    <span class="comment">// PauseNs is a circular buffer of recent GC stop-the-world</span>
    <span class="comment">// pause times in nanoseconds.</span>
    <span class="comment">//</span>
    <span class="comment">// The most recent pause is at PauseNs[(NumGC+255)%256]. In</span>
    <span class="comment">// general, PauseNs[N%256] records the time paused in the most</span>
    <span class="comment">// recent N%256th GC cycle. There may be multiple pauses per</span>
    <span class="comment">// GC cycle; this is the sum of all pauses during a cycle.</span>
    PauseNs [256]<a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.PauseEnd"></span>    <span class="comment">// PauseEnd is a circular buffer of recent GC pause end times,</span>
    <span class="comment">// as nanoseconds since 1970 (the UNIX epoch).</span>
    <span class="comment">//</span>
    <span class="comment">// This buffer is filled the same way as PauseNs. There may be</span>
    <span class="comment">// multiple pauses per GC cycle; this records the end of the</span>
    <span class="comment">// last pause in a cycle.</span>
    PauseEnd [256]<a href="/pkg/builtin/#uint64">uint64</a>

<span id="MemStats.NumGC"></span>    <span class="comment">// NumGC is the number of completed GC cycles.</span>
    NumGC <a href="/pkg/builtin/#uint32">uint32</a>

<span id="MemStats.NumForcedGC"></span>    <span class="comment">// NumForcedGC is the number of GC cycles that were forced by</span>
    <span class="comment">// the application calling the GC function.</span>
    NumForcedGC <a href="/pkg/builtin/#uint32">uint32</a>

<span id="MemStats.GCCPUFraction"></span>    <span class="comment">// GCCPUFraction is the fraction of this program&#39;s available</span>
    <span class="comment">// CPU time used by the GC since the program started.</span>
    <span class="comment">//</span>
    <span class="comment">// GCCPUFraction is expressed as a number between 0 and 1,</span>
    <span class="comment">// where 0 means GC has consumed none of this program&#39;s CPU. A</span>
    <span class="comment">// program&#39;s available CPU time is defined as the integral of</span>
    <span class="comment">// GOMAXPROCS since the program started. That is, if</span>
    <span class="comment">// GOMAXPROCS is 2 and a program has been running for 10</span>
    <span class="comment">// seconds, its &#34;available CPU&#34; is 20 seconds. GCCPUFraction</span>
    <span class="comment">// does not include CPU time used for write barrier activity.</span>
    <span class="comment">//</span>
    <span class="comment">// This is the same as the fraction of CPU reported by</span>
    <span class="comment">// GODEBUG=gctrace=1.</span>
    GCCPUFraction <a href="/pkg/builtin/#float64">float64</a>

<span id="MemStats.EnableGC"></span>    <span class="comment">// EnableGC indicates that GC is enabled. It is always true,</span>
    <span class="comment">// even if GOGC=off.</span>
    EnableGC <a href="/pkg/builtin/#bool">bool</a>

<span id="MemStats.DebugGC"></span>    <span class="comment">// DebugGC is currently unused.</span>
    DebugGC <a href="/pkg/builtin/#bool">bool</a>

<span id="MemStats.BySize"></span>    <span class="comment">// BySize reports per-size class allocation statistics.</span>
    <span class="comment">//</span>
    <span class="comment">// BySize[N] gives statistics for allocations of size S where</span>
    <span class="comment">// BySize[N-1].Size &lt; S ≤ BySize[N].Size.</span>
    <span class="comment">//</span>
    <span class="comment">// This does not report allocations larger than BySize[60].Size.</span>
    BySize [61]struct {
        <span class="comment">// Size is the maximum byte size of an object in this</span>
        <span class="comment">// size class.</span>
        Size <a href="/pkg/builtin/#uint32">uint32</a>

        <span class="comment">// Mallocs is the cumulative count of heap objects</span>
        <span class="comment">// allocated in this size class. The cumulative bytes</span>
        <span class="comment">// of allocation is Size*Mallocs. The number of live</span>
        <span class="comment">// objects in this size class is Mallocs - Frees.</span>
        Mallocs <a href="/pkg/builtin/#uint64">uint64</a>

        <span class="comment">// Frees is the cumulative count of heap objects freed</span>
        <span class="comment">// in this size class.</span>
        Frees <a href="/pkg/builtin/#uint64">uint64</a>
    }
}
</pre>











## <a id="StackRecord">type</a> [StackRecord](https://golang.org/src/runtime/mprof.go?s=13296:13397#L454)
A StackRecord describes a single execution stack.


<pre>type StackRecord struct {
<span id="StackRecord.Stack0"></span>    Stack0 [32]<a href="/pkg/builtin/#uintptr">uintptr</a> <span class="comment">// stack trace for this record; ends at first 0 entry</span>
}
</pre>











### <a id="StackRecord.Stack">func</a> (\*StackRecord) [Stack](https://golang.org/src/runtime/mprof.go?s=13485:13524#L460)
<pre>func (r *<a href="#StackRecord">StackRecord</a>) Stack() []<a href="/pkg/builtin/#uintptr">uintptr</a></pre>
Stack returns the stack trace associated with the record,
a prefix of r.Stack0.




## <a id="TypeAssertionError">type</a> [TypeAssertionError](https://golang.org/src/runtime/error.go?s=552:731#L11)
A TypeAssertionError explains a failed type assertion.


<pre>type TypeAssertionError struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="TypeAssertionError.Error">func</a> (\*TypeAssertionError) [Error](https://golang.org/src/runtime/error.go?s=779:822#L20)
<pre>func (e *<a href="#TypeAssertionError">TypeAssertionError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="TypeAssertionError.RuntimeError">func</a> (\*TypeAssertionError) [RuntimeError](https://golang.org/src/runtime/error.go?s=733:774#L18)
<pre>func (*<a href="#TypeAssertionError">TypeAssertionError</a>) RuntimeError()</pre>







