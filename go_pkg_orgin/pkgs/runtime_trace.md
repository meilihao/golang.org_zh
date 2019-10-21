

# trace
`import "runtime/trace"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package trace contains facilities for programs to generate traces
for the Go execution tracer.

### Tracing runtime activities
The execution trace captures a wide range of execution events such as
goroutine creation/blocking/unblocking, syscall enter/exit/block,
GC-related events, changes of heap size, processor start/stop, etc.
A precise nanosecond-precision timestamp and a stack trace is
captured for most events. The generated trace can be interpreted
using `go tool trace`.

Support for tracing tests and benchmarks built with the standard
testing package is built into `go test`. For example, the following
command runs the test in the current directory and writes the trace
file (trace.out).


	go test -trace=test.out

This runtime/trace package provides APIs to add equivalent tracing
support to a standalone program. See the Example that demonstrates
how to use this API to enable tracing.

There is also a standard HTTP interface to trace data. Adding the
following line will install a handler under the /debug/pprof/trace URL
to download a live trace:


	import _ "net/http/pprof"

See the net/http/pprof package for more details about all of the
debug endpoints installed by this import.

### User annotation
Package trace provides user annotation APIs that can be used to
log interesting events during execution.

There are three types of user annotations: log messages, regions,
and tasks.

Log emits a timestamped message to the execution trace along with
additional information such as the category of the message and
which goroutine called Log. The execution tracer provides UIs to filter
and group goroutines using the log category and the message supplied
in Log.

A region is for logging a time interval during a goroutine's execution.
By definition, a region starts and ends in the same goroutine.
Regions can be nested to represent subintervals.
For example, the following code records four regions in the execution
trace to trace the durations of sequential steps in a cappuccino making
operation.


	trace.WithRegion(ctx, "makeCappuccino", func() {
	
	   // orderID allows to identify a specific order
	   // among many cappuccino order region records.
	   trace.Log(ctx, "orderID", orderID)
	
	   trace.WithRegion(ctx, "steamMilk", steamMilk)
	   trace.WithRegion(ctx, "extractCoffee", extractCoffee)
	   trace.WithRegion(ctx, "mixMilkCoffee", mixMilkCoffee)
	})

A task is a higher-level component that aids tracing of logical
operations such as an RPC request, an HTTP request, or an
interesting local operation which may require multiple goroutines
working together. Since tasks can involve multiple goroutines,
they are tracked via a context.Context object. NewTask creates
a new task and embeds it in the returned context.Context object.
Log messages and regions are attached to the task, if any, in the
Context passed to Log and WithRegion.

For example, assume that we decided to froth milk, extract coffee,
and mix milk and coffee in separate goroutines. With a task,
the trace tool can identify the goroutines involved in a specific
cappuccino order.


	ctx, task := trace.NewTask(ctx, "makeCappuccino")
	trace.Log(ctx, "orderID", orderID)
	
	milk := make(chan bool)
	espresso := make(chan bool)
	
	go func() {
	        trace.WithRegion(ctx, "steamMilk", steamMilk)
	        milk <- true
	}()
	go func() {
	        trace.WithRegion(ctx, "extractCoffee", extractCoffee)
	        espresso <- true
	}()
	go func() {
	        defer task.End() // When assemble is done, the order is complete.
	        <-espresso
	        <-milk
	        trace.WithRegion(ctx, "mixMilkCoffee", mixMilkCoffee)
	}()

The trace tool computes the latency of a task by measuring the
time between the task creation and the task end and provides
latency distributions for each task type found in the trace.


<a id="example_">Example</a>
```go
```

output:
```txt
```
<p>Example demonstrates the use of the trace package to trace
the execution of a Go program. The trace output will be
written to the file trace.out
</p>

## <a id="pkg-index">Index</a>
* [func IsEnabled() bool](#IsEnabled)
* [func Log(ctx context.Context, category, message string)](#Log)
* [func Logf(ctx context.Context, category, format string, args ...interface{})](#Logf)
* [func Start(w io.Writer) error](#Start)
* [func Stop()](#Stop)
* [func WithRegion(ctx context.Context, regionType string, fn func())](#WithRegion)
* [type Region](#Region)
  * [func StartRegion(ctx context.Context, regionType string) *Region](#StartRegion)
  * [func (r *Region) End()](#Region.End)
* [type Task](#Task)
  * [func NewTask(pctx context.Context, taskType string) (ctx context.Context, task *Task)](#NewTask)
  * [func (t *Task) End()](#Task.End)


#### <a id="pkg-examples">Examples</a>
* [Package](#example_)


#### <a id="pkg-files">Package files</a>
[annotation.go](https://golang.org/src/runtime/trace/annotation.go) [trace.go](https://golang.org/src/runtime/trace/trace.go) 






## <a id="IsEnabled">func</a> [IsEnabled](https://golang.org/src/runtime/trace/annotation.go?s=5752:5773#L167)
<pre>func IsEnabled() <a href="/pkg/builtin/#bool">bool</a></pre>
IsEnabled reports whether tracing is enabled.
The information is advisory only. The tracing status
may have changed by the time this function returns.



## <a id="Log">func</a> [Log](https://golang.org/src/runtime/trace/annotation.go?s=2863:2918#L81)
<pre>func Log(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, category, message <a href="/pkg/builtin/#string">string</a>)</pre>
Log emits a one-off event with the given category and message.
Category can be empty and the API assumes there are only a handful of
unique categories in the system.



## <a id="Logf">func</a> [Logf](https://golang.org/src/runtime/trace/annotation.go?s=3064:3140#L87)
<pre>func Logf(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, category, format <a href="/pkg/builtin/#string">string</a>, args ...interface{})</pre>
Logf is like Log, but the value is formatted using the specified format spec.



## <a id="Start">func</a> [Start](https://golang.org/src/runtime/trace/trace.go?s=4570:4599#L110)
<pre>func Start(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) <a href="/pkg/builtin/#error">error</a></pre>
Start enables tracing for the current program.
While tracing, the trace will be buffered and written to w.
Start returns an error if tracing is already enabled.



## <a id="Stop">func</a> [Stop](https://golang.org/src/runtime/trace/trace.go?s=4988:4999#L132)
<pre>func Stop()</pre>
Stop stops the current tracing, if any.
Stop only returns after all the writes for the trace have completed.



## <a id="WithRegion">func</a> [WithRegion](https://golang.org/src/runtime/trace/annotation.go?s=3767:3833#L108)
<pre>func WithRegion(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, regionType <a href="/pkg/builtin/#string">string</a>, fn func())</pre>
WithRegion starts a region associated with its calling goroutine, runs fn,
and then ends the region. If the context carries a task, the region is
associated with the task. Otherwise, the region is attached to the background
task.

The regionType is used to classify regions, so there should be only a
handful of unique region types.





## <a id="Region">type</a> [Region](https://golang.org/src/runtime/trace/annotation.go?s=5345:5405#L149)
Region is a region of code whose execution time interval is traced.


<pre>type Region struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="StartRegion">func</a> [StartRegion](https://golang.org/src/runtime/trace/annotation.go?s=5058:5122#L139)
<pre>func StartRegion(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, regionType <a href="/pkg/builtin/#string">string</a>) *<a href="#Region">Region</a></pre>
StartRegion starts a region and returns a function for marking the
end of the region. The returned Region's End function must be called
from the same goroutine where the region was started.
Within each goroutine, regions must nest. That is, regions started
after this region must be ended before this region can be ended.
Recommended usage is


	defer trace.StartRegion(ctx, "myTracedRegion").End()






### <a id="Region.End">func</a> (\*Region) [End](https://golang.org/src/runtime/trace/annotation.go?s=5483:5505#L157)
<pre>func (r *<a href="#Region">Region</a>) End()</pre>
End marks the end of the traced code region.




## <a id="Task">type</a> [Task](https://golang.org/src/runtime/trace/annotation.go?s=2320:2388#L59)
Task is a data type for tracing a user-defined, logical operation.


<pre>type Task struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewTask">func</a> [NewTask](https://golang.org/src/runtime/trace/annotation.go?s=1131:1216#L24)
<pre>func NewTask(pctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, taskType <a href="/pkg/builtin/#string">string</a>) (ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, task *<a href="#Task">Task</a>)</pre>
NewTask creates a task instance with the type taskType and returns
it along with a Context that carries the task.
If the input context contains a task, the new task is its subtask.

The taskType is used to classify task instances. Analysis tools
like the Go execution tracer may assume there are only a bounded
number of unique task types in the system.

The returned end function is used to mark the task's end.
The trace tool measures task latency as the time between task creation
and when the end function is called, and provides the latency
distribution per task type.
If the end function is called multiple times, only the first
call is used in the latency measurement.


	ctx, task := trace.NewTask(ctx, "awesomeTask")
	trace.WithRegion(ctx, "preparation", prepWork)
	// preparation of the task
	go func() {  // continue processing the task in a separate goroutine.
	    defer task.End()
	    trace.WithRegion(ctx, "remainingWork", remainingWork)
	}()






### <a id="Task.End">func</a> (\*Task) [End](https://golang.org/src/runtime/trace/annotation.go?s=2453:2473#L65)
<pre>func (t *<a href="#Task">Task</a>) End()</pre>
End marks the end of the operation represented by the Task.







