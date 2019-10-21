

# testing
`import "testing"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package testing provides support for automated testing of Go packages.
It is intended to be used in concert with the ``go test'' command, which automates
execution of any function of the form


	func TestXxx(*testing.T)

where Xxx does not start with a lowercase letter. The function name
serves to identify the test routine.

Within these functions, use the Error, Fail or related methods to signal failure.

To write a new test suite, create a file whose name ends _test.go that
contains the TestXxx functions as described here. Put the file in the same
package as the one being tested. The file will be excluded from regular
package builds but will be included when the ``go test'' command is run.
For more detail, run ``go help test'' and ``go help testflag''.

A simple test function looks like this:


	func TestAbs(t *testing.T) {
	    got := Abs(-1)
	    if got != 1 {
	        t.Errorf("Abs(-1) = %d; want 1", got)
	    }
	}

### Benchmarks
Functions of the form


	func BenchmarkXxx(*testing.B)

are considered benchmarks, and are executed by the "go test" command when
its -bench flag is provided. Benchmarks are run sequentially.

For a description of the testing flags, see
<a href="https://golang.org/cmd/go/#hdr-Testing_flags">https://golang.org/cmd/go/#hdr-Testing_flags</a>

A sample benchmark function looks like this:


	func BenchmarkHello(b *testing.B) {
	    for i := 0; i < b.N; i++ {
	        fmt.Sprintf("hello")
	    }
	}

The benchmark function must run the target code b.N times.
During benchmark execution, b.N is adjusted until the benchmark function lasts
long enough to be timed reliably. The output


	BenchmarkHello    10000000    282 ns/op

means that the loop ran 10000000 times at a speed of 282 ns per loop.

If a benchmark needs some expensive setup before running, the timer
may be reset:


	func BenchmarkBigLen(b *testing.B) {
	    big := NewBig()
	    b.ResetTimer()
	    for i := 0; i < b.N; i++ {
	        big.Len()
	    }
	}

If a benchmark needs to test performance in a parallel setting, it may use
the RunParallel helper function; such benchmarks are intended to be used with
the go test -cpu flag:


	func BenchmarkTemplateParallel(b *testing.B) {
	    templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	    b.RunParallel(func(pb *testing.PB) {
	        var buf bytes.Buffer
	        for pb.Next() {
	            buf.Reset()
	            templ.Execute(&buf, "World")
	        }
	    })
	}

### Examples
The package also runs and verifies example code. Example functions may
include a concluding line comment that begins with "Output:" and is compared with
the standard output of the function when the tests are run. (The comparison
ignores leading and trailing space.) These are examples of an example:


	func ExampleHello() {
	    fmt.Println("hello")
	    // Output: hello
	}
	
	func ExampleSalutations() {
	    fmt.Println("hello, and")
	    fmt.Println("goodbye")
	    // Output:
	    // hello, and
	    // goodbye
	}

The comment prefix "Unordered output:" is like "Output:", but matches any
line order:


	func ExamplePerm() {
	    for _, value := range Perm(4) {
	        fmt.Println(value)
	    }
	    // Unordered output: 4
	    // 2
	    // 1
	    // 3
	    // 0
	}

Example functions without output comments are compiled but not executed.

The naming convention to declare examples for the package, a function F, a type T and
method M on type T are:


	func Example() { ... }
	func ExampleF() { ... }
	func ExampleT() { ... }
	func ExampleT_M() { ... }

Multiple example functions for a package/type/function/method may be provided by
appending a distinct suffix to the name. The suffix must start with a
lower-case letter.


	func Example_suffix() { ... }
	func ExampleF_suffix() { ... }
	func ExampleT_suffix() { ... }
	func ExampleT_M_suffix() { ... }

The entire test file is presented as the example when it contains a single
example function, at least one other function, type, variable, or constant
declaration, and no test or benchmark functions.

### Skipping
Tests or benchmarks may be skipped at run time with a call to
the Skip method of *T or *B:


	func TestTimeConsuming(t *testing.T) {
	    if testing.Short() {
	        t.Skip("skipping test in short mode.")
	    }
	    ...
	}

### Subtests and Sub-benchmarks
The Run methods of T and B allow defining subtests and sub-benchmarks,
without having to define separate functions for each. This enables uses
like table-driven benchmarks and creating hierarchical tests.
It also provides a way to share common setup and tear-down code:


	func TestFoo(t *testing.T) {
	    // <setup code>
	    t.Run("A=1", func(t *testing.T) { ... })
	    t.Run("A=2", func(t *testing.T) { ... })
	    t.Run("B=1", func(t *testing.T) { ... })
	    // <tear-down code>
	}

Each subtest and sub-benchmark has a unique name: the combination of the name
of the top-level test and the sequence of names passed to Run, separated by
slashes, with an optional trailing sequence number for disambiguation.

The argument to the -run and -bench command-line flags is an unanchored regular
expression that matches the test's name. For tests with multiple slash-separated
elements, such as subtests, the argument is itself slash-separated, with
expressions matching each name element in turn. Because it is unanchored, an
empty expression matches any string.
For example, using "matching" to mean "whose name contains":


	go test -run ''      # Run all tests.
	go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".
	go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".
	go test -run /A=1    # For all top-level tests, run subtests matching "A=1".

Subtests can also be used to control parallelism. A parent test will only
complete once all of its subtests complete. In this example, all tests are
run in parallel with each other, and only with each other, regardless of
other top-level tests that may be defined:


	func TestGroupedParallel(t *testing.T) {
	    for _, tc := range tests {
	        tc := tc // capture range variable
	        t.Run(tc.Name, func(t *testing.T) {
	            t.Parallel()
	            ...
	        })
	    }
	}

The race detector kills the program if it exceeds 8192 concurrent goroutines,
so use care when running parallel tests with the -race flag set.

Run does not return until parallel subtests have completed, providing a way
to clean up after a group of parallel tests:


	func TestTeardownParallel(t *testing.T) {
	    // This Run will not return until the parallel tests finish.
	    t.Run("group", func(t *testing.T) {
	        t.Run("Test1", parallelTest1)
	        t.Run("Test2", parallelTest2)
	        t.Run("Test3", parallelTest3)
	    })
	    // <tear-down code>
	}

### Main
It is sometimes necessary for a test program to do extra setup or teardown
before or after testing. It is also sometimes necessary for a test to control
which code runs on the main thread. To support these and other cases,
if a test file contains a function:


	func TestMain(m *testing.M)

then the generated test will call TestMain(m) instead of running the tests
directly. TestMain runs in the main goroutine and can do whatever setup
and teardown is necessary around a call to m.Run. It should then call
os.Exit with the result of m.Run. When TestMain is called, flag.Parse has
not been run. If TestMain depends on command-line flags, including those
of the testing package, it should call flag.Parse explicitly.

A simple implementation of TestMain is:


	func TestMain(m *testing.M) {
		// call flag.Parse() here if TestMain uses flags
		os.Exit(m.Run())
	}




## <a id="pkg-index">Index</a>
* [func AllocsPerRun(runs int, f func()) (avg float64)](#AllocsPerRun)
* [func CoverMode() string](#CoverMode)
* [func Coverage() float64](#Coverage)
* [func Init()](#Init)
* [func Main(matchString func(pat, str string) (bool, error), tests []InternalTest, benchmarks []InternalBenchmark, examples []InternalExample)](#Main)
* [func RegisterCover(c Cover)](#RegisterCover)
* [func RunBenchmarks(matchString func(pat, str string) (bool, error), benchmarks []InternalBenchmark)](#RunBenchmarks)
* [func RunExamples(matchString func(pat, str string) (bool, error), examples []InternalExample) (ok bool)](#RunExamples)
* [func RunTests(matchString func(pat, str string) (bool, error), tests []InternalTest) (ok bool)](#RunTests)
* [func Short() bool](#Short)
* [func Verbose() bool](#Verbose)
* [type B](#B)
  * [func (c *B) Error(args ...interface{})](#B.Error)
  * [func (c *B) Errorf(format string, args ...interface{})](#B.Errorf)
  * [func (c *B) Fail()](#B.Fail)
  * [func (c *B) FailNow()](#B.FailNow)
  * [func (c *B) Failed() bool](#B.Failed)
  * [func (c *B) Fatal(args ...interface{})](#B.Fatal)
  * [func (c *B) Fatalf(format string, args ...interface{})](#B.Fatalf)
  * [func (c *B) Helper()](#B.Helper)
  * [func (c *B) Log(args ...interface{})](#B.Log)
  * [func (c *B) Logf(format string, args ...interface{})](#B.Logf)
  * [func (c *B) Name() string](#B.Name)
  * [func (b *B) ReportAllocs()](#B.ReportAllocs)
  * [func (b *B) ReportMetric(n float64, unit string)](#B.ReportMetric)
  * [func (b *B) ResetTimer()](#B.ResetTimer)
  * [func (b *B) Run(name string, f func(b *B)) bool](#B.Run)
  * [func (b *B) RunParallel(body func(*PB))](#B.RunParallel)
  * [func (b *B) SetBytes(n int64)](#B.SetBytes)
  * [func (b *B) SetParallelism(p int)](#B.SetParallelism)
  * [func (c *B) Skip(args ...interface{})](#B.Skip)
  * [func (c *B) SkipNow()](#B.SkipNow)
  * [func (c *B) Skipf(format string, args ...interface{})](#B.Skipf)
  * [func (c *B) Skipped() bool](#B.Skipped)
  * [func (b *B) StartTimer()](#B.StartTimer)
  * [func (b *B) StopTimer()](#B.StopTimer)
* [type BenchmarkResult](#BenchmarkResult)
  * [func Benchmark(f func(b *B)) BenchmarkResult](#Benchmark)
  * [func (r BenchmarkResult) AllocedBytesPerOp() int64](#BenchmarkResult.AllocedBytesPerOp)
  * [func (r BenchmarkResult) AllocsPerOp() int64](#BenchmarkResult.AllocsPerOp)
  * [func (r BenchmarkResult) MemString() string](#BenchmarkResult.MemString)
  * [func (r BenchmarkResult) NsPerOp() int64](#BenchmarkResult.NsPerOp)
  * [func (r BenchmarkResult) String() string](#BenchmarkResult.String)
* [type Cover](#Cover)
* [type CoverBlock](#CoverBlock)
* [type InternalBenchmark](#InternalBenchmark)
* [type InternalExample](#InternalExample)
* [type InternalTest](#InternalTest)
* [type M](#M)
  * [func MainStart(deps testDeps, tests []InternalTest, benchmarks []InternalBenchmark, examples []InternalExample) *M](#MainStart)
  * [func (m *M) Run() int](#M.Run)
* [type PB](#PB)
  * [func (pb *PB) Next() bool](#PB.Next)
* [type T](#T)
  * [func (c *T) Error(args ...interface{})](#T.Error)
  * [func (c *T) Errorf(format string, args ...interface{})](#T.Errorf)
  * [func (c *T) Fail()](#T.Fail)
  * [func (c *T) FailNow()](#T.FailNow)
  * [func (c *T) Failed() bool](#T.Failed)
  * [func (c *T) Fatal(args ...interface{})](#T.Fatal)
  * [func (c *T) Fatalf(format string, args ...interface{})](#T.Fatalf)
  * [func (c *T) Helper()](#T.Helper)
  * [func (c *T) Log(args ...interface{})](#T.Log)
  * [func (c *T) Logf(format string, args ...interface{})](#T.Logf)
  * [func (c *T) Name() string](#T.Name)
  * [func (t *T) Parallel()](#T.Parallel)
  * [func (t *T) Run(name string, f func(t *T)) bool](#T.Run)
  * [func (c *T) Skip(args ...interface{})](#T.Skip)
  * [func (c *T) SkipNow()](#T.SkipNow)
  * [func (c *T) Skipf(format string, args ...interface{})](#T.Skipf)
  * [func (c *T) Skipped() bool](#T.Skipped)
* [type TB](#TB)


#### <a id="pkg-examples">Examples</a>
* [B.ReportMetric](#example_B_ReportMetric)
* [B.RunParallel](#example_B_RunParallel)


#### <a id="pkg-files">Package files</a>
[allocs.go](https://golang.org/src/testing/allocs.go) [benchmark.go](https://golang.org/src/testing/benchmark.go) [cover.go](https://golang.org/src/testing/cover.go) [example.go](https://golang.org/src/testing/example.go) [match.go](https://golang.org/src/testing/match.go) [run_example.go](https://golang.org/src/testing/run_example.go) [testing.go](https://golang.org/src/testing/testing.go) 






## <a id="AllocsPerRun">func</a> [AllocsPerRun](https://golang.org/src/testing/allocs.go?s=669:720#L10)
<pre>func AllocsPerRun(runs <a href="/pkg/builtin/#int">int</a>, f func()) (avg <a href="/pkg/builtin/#float64">float64</a>)</pre>
AllocsPerRun returns the average number of allocations during calls to f.
Although the return value has type float64, it will always be an integral value.

To compute the number of allocations, the function will first be run once as
a warm-up. The average number of allocations over the specified number of
runs will then be measured and returned.

AllocsPerRun sets GOMAXPROCS to 1 during its measurement and will restore
it before returning.



## <a id="CoverMode">func</a> [CoverMode](https://golang.org/src/testing/testing.go?s=14759:14782#L371)
<pre>func CoverMode() <a href="/pkg/builtin/#string">string</a></pre>
CoverMode reports what the test coverage mode is set to. The
values are "set", "count", or "atomic". The return value will be
empty if test coverage is not enabled.



## <a id="Coverage">func</a> [Coverage](https://golang.org/src/testing/cover.go?s=1633:1656#L38)
<pre>func Coverage() <a href="/pkg/builtin/#float64">float64</a></pre>
Coverage reports the current code coverage as a fraction in the range [0, 1].
If coverage is not enabled, Coverage returns 0.

When running a large set of sequential test cases, checking Coverage after each one
can be useful for identifying which test cases exercise new code paths.
It is not a replacement for the reports generated by 'go test -cover' and
'go tool cover'.



## <a id="Init">func</a> [Init](https://golang.org/src/testing/testing.go?s=9123:9134#L249)
<pre>func Init()</pre>
Init registers testing flags. These flags are automatically registered by
the "go test" command before running test functions, so Init is only needed
when calling functions such as Benchmark without using "go test".

Init has no effect if it was already called.



## <a id="Main">func</a> [Main](https://golang.org/src/testing/testing.go?s=34952:35092#L1033)
<pre>func Main(matchString func(pat, str <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#bool">bool</a>, <a href="/pkg/builtin/#error">error</a>), tests []<a href="#InternalTest">InternalTest</a>, benchmarks []<a href="#InternalBenchmark">InternalBenchmark</a>, examples []<a href="#InternalExample">InternalExample</a>)</pre>
Main is an internal function, part of the implementation of the "go test" command.
It was exported because it is cross-package and predates "internal" packages.
It is no longer used by "go test" but preserved, as much as possible, for other
systems that simulate "go test" using Main, but Main sometimes cannot be updated as
new functionality is added to the testing package.
Systems simulating "go test" should be updated to use MainStart.



## <a id="RegisterCover">func</a> [RegisterCover](https://golang.org/src/testing/cover.go?s=2096:2123#L57)
<pre>func RegisterCover(c <a href="#Cover">Cover</a>)</pre>
RegisterCover records the coverage data accumulators for the tests.
NOTE: This function is internal to the testing infrastructure and may change.
It is not covered (yet) by the Go 1 compatibility guidelines.



## <a id="RunBenchmarks">func</a> [RunBenchmarks](https://golang.org/src/testing/benchmark.go?s=13648:13747#L483)
<pre>func RunBenchmarks(matchString func(pat, str <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#bool">bool</a>, <a href="/pkg/builtin/#error">error</a>), benchmarks []<a href="#InternalBenchmark">InternalBenchmark</a>)</pre>
An internal function but exported because it is cross-package; part of the implementation
of the "go test" command.



## <a id="RunExamples">func</a> [RunExamples](https://golang.org/src/testing/example.go?s=454:557#L14)
<pre>func RunExamples(matchString func(pat, str <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#bool">bool</a>, <a href="/pkg/builtin/#error">error</a>), examples []<a href="#InternalExample">InternalExample</a>) (ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
An internal function but exported because it is cross-package; part of the implementation
of the "go test" command.



## <a id="RunTests">func</a> [RunTests](https://golang.org/src/testing/testing.go?s=38708:38802#L1164)
<pre>func RunTests(matchString func(pat, str <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#bool">bool</a>, <a href="/pkg/builtin/#error">error</a>), tests []<a href="#InternalTest">InternalTest</a>) (ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
An internal function but exported because it is cross-package; part of the implementation
of the "go test" command.



## <a id="Short">func</a> [Short](https://golang.org/src/testing/testing.go?s=14330:14347#L356)
<pre>func Short() <a href="/pkg/builtin/#bool">bool</a></pre>
Short reports whether the -test.short flag is set.



## <a id="Verbose">func</a> [Verbose](https://golang.org/src/testing/testing.go?s=14859:14878#L376)
<pre>func Verbose() <a href="/pkg/builtin/#bool">bool</a></pre>
Verbose reports whether the -test.v flag is set.





## <a id="B">type</a> [B](https://golang.org/src/testing/benchmark.go?s=2425:3313#L82)
B is a type passed to Benchmark functions to manage benchmark
timing and to specify the number of iterations to run.

A benchmark ends when its Benchmark function returns or calls any of the methods
FailNow, Fatal, Fatalf, SkipNow, Skip, or Skipf. Those methods must be called
only from the goroutine running the Benchmark function.
The other reporting methods, such as the variations of Log and Error,
may be called simultaneously from multiple goroutines.

Like in tests, benchmark logs are accumulated during execution
and dumped to standard error when done. Unlike in tests, benchmark logs
are always printed, so as not to hide output whose existence may be
affecting benchmark results.


<pre>type B struct {
<span id="B.N"></span>    N <a href="/pkg/builtin/#int">int</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="B.Error">func</a> (\*B) [Error](https://golang.org/src/testing/testing.go?s=24490:24533#L686)
<pre>func (c *<a href="#B">B</a>) Error(args ...interface{})</pre>
Error is equivalent to Log followed by Fail.




### <a id="B.Errorf">func</a> (\*B) [Errorf](https://golang.org/src/testing/testing.go?s=24629:24688#L692)
<pre>func (c *<a href="#B">B</a>) Errorf(format <a href="/pkg/builtin/#string">string</a>, args ...interface{})</pre>
Errorf is equivalent to Logf followed by Fail.




### <a id="B.Fail">func</a> (\*B) [Fail](https://golang.org/src/testing/testing.go?s=20928:20951#L591)
<pre>func (c *<a href="#B">B</a>) Fail()</pre>
Fail marks the function as having failed but continues execution.




### <a id="B.FailNow">func</a> (\*B) [FailNow](https://golang.org/src/testing/testing.go?s=21819:21845#L620)
<pre>func (c *<a href="#B">B</a>) FailNow()</pre>
FailNow marks the function as having failed and stops its execution
by calling runtime.Goexit (which then runs all deferred calls in the
current goroutine).
Execution will continue at the next test or benchmark.
FailNow must be called from the goroutine running the
test or benchmark function, not from other goroutines
created during the test. Calling FailNow does not stop
those other goroutines.




### <a id="B.Failed">func</a> (\*B) [Failed](https://golang.org/src/testing/testing.go?s=21261:21291#L605)
<pre>func (c *<a href="#B">B</a>) Failed() <a href="/pkg/builtin/#bool">bool</a></pre>
Failed reports whether the function has failed.




### <a id="B.Fatal">func</a> (\*B) [Fatal](https://golang.org/src/testing/testing.go?s=24792:24835#L698)
<pre>func (c *<a href="#B">B</a>) Fatal(args ...interface{})</pre>
Fatal is equivalent to Log followed by FailNow.




### <a id="B.Fatalf">func</a> (\*B) [Fatalf](https://golang.org/src/testing/testing.go?s=24937:24996#L704)
<pre>func (c *<a href="#B">B</a>) Fatalf(format <a href="/pkg/builtin/#string">string</a>, args ...interface{})</pre>
Fatalf is equivalent to Logf followed by FailNow.




### <a id="B.Helper">func</a> (\*B) [Helper](https://golang.org/src/testing/testing.go?s=26324:26349#L751)
<pre>func (c *<a href="#B">B</a>) Helper()</pre>
Helper marks the calling function as a test helper function.
When printing file and line information, that function will be skipped.
Helper may be called simultaneously from multiple goroutines.




### <a id="B.Log">func</a> (\*B) [Log](https://golang.org/src/testing/testing.go?s=23899:23940#L676)
<pre>func (c *<a href="#B">B</a>) Log(args ...interface{})</pre>
Log formats its arguments using default formatting, analogous to Println,
and records the text in the error log. For tests, the text will be printed only if
the test fails or the -test.v flag is set. For benchmarks, the text is always
printed to avoid having performance depend on the value of the -test.v flag.




### <a id="B.Logf">func</a> (\*B) [Logf](https://golang.org/src/testing/testing.go?s=24343:24400#L683)
<pre>func (c *<a href="#B">B</a>) Logf(format <a href="/pkg/builtin/#string">string</a>, args ...interface{})</pre>
Logf formats its arguments according to the format, analogous to Printf, and
records the text in the error log. A final newline is added if not provided. For
tests, the text will be printed only if the test fails or the -test.v flag is
set. For benchmarks, the text is always printed to avoid having performance
depend on the value of the -test.v flag.




### <a id="B.Name">func</a> (\*B) [Name](https://golang.org/src/testing/testing.go?s=20684:20714#L577)
<pre>func (c *<a href="#B">B</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the name of the running test or benchmark.




### <a id="B.ReportAllocs">func</a> (\*B) [ReportAllocs](https://golang.org/src/testing/benchmark.go?s=5071:5097#L164)
<pre>func (b *<a href="#B">B</a>) ReportAllocs()</pre>
ReportAllocs enables malloc statistics for this benchmark.
It is equivalent to setting -test.benchmem, but it only affects the
benchmark function that calls ReportAllocs.




### <a id="B.ReportMetric">func</a> (\*B) [ReportMetric](https://golang.org/src/testing/benchmark.go?s=9573:9621#L325)
<pre>func (b *<a href="#B">B</a>) ReportMetric(n <a href="/pkg/builtin/#float64">float64</a>, unit <a href="/pkg/builtin/#string">string</a>)</pre>
ReportMetric adds "n unit" to the reported benchmark results.
If the metric is per-iteration, the caller should divide by b.N,
and by convention units should end in "/op".
ReportMetric overrides any previously reported value for the same unit.
ReportMetric panics if unit is the empty string or if unit contains
any whitespace.
If unit is a unit normally reported by the benchmark framework itself
(such as "allocs/op"), ReportMetric will override that metric.
Setting "ns/op" to 0 will suppress that built-in metric.


<a id="example_B_ReportMetric">Example</a>
```go
```

output:
```txt
```


### <a id="B.ResetTimer">func</a> (\*B) [ResetTimer](https://golang.org/src/testing/benchmark.go?s=4244:4268#L136)
<pre>func (b *<a href="#B">B</a>) ResetTimer()</pre>
ResetTimer zeroes the elapsed benchmark time and memory allocation counters
and deletes user-reported metrics.
It does not affect whether the timer is running.




### <a id="B.Run">func</a> (\*B) [Run](https://golang.org/src/testing/benchmark.go?s=16694:16741#L584)
<pre>func (b *<a href="#B">B</a>) Run(name <a href="/pkg/builtin/#string">string</a>, f func(b *<a href="#B">B</a>)) <a href="/pkg/builtin/#bool">bool</a></pre>
Run benchmarks f as a subbenchmark with the given name. It reports
whether there were any failures.

A subbenchmark is like any other benchmark. A benchmark that calls Run at
least once will not be measured itself and will be called once with N=1.




### <a id="B.RunParallel">func</a> (\*B) [RunParallel](https://golang.org/src/testing/benchmark.go?s=20296:20335#L700)
<pre>func (b *<a href="#B">B</a>) RunParallel(body func(*<a href="#PB">PB</a>))</pre>
RunParallel runs a benchmark in parallel.
It creates multiple goroutines and distributes b.N iterations among them.
The number of goroutines defaults to GOMAXPROCS. To increase parallelism for
non-CPU-bound benchmarks, call SetParallelism before RunParallel.
RunParallel is usually used with the go test -cpu flag.

The body function will be run in each goroutine. It should set up any
goroutine-local state and then iterate until pb.Next returns false.
It should not use the StartTimer, StopTimer, or ResetTimer functions,
because they have global effect. It should also not call Run.


<a id="example_B_RunParallel">Example</a>
```go
```

output:
```txt
```


### <a id="B.SetBytes">func</a> (\*B) [SetBytes](https://golang.org/src/testing/benchmark.go?s=4844:4873#L159)
<pre>func (b *<a href="#B">B</a>) SetBytes(n <a href="/pkg/builtin/#int64">int64</a>)</pre>
SetBytes records the number of bytes processed in a single operation.
If this is called, the benchmark will report ns/op and MB/s.




### <a id="B.SetParallelism">func</a> (\*B) [SetParallelism](https://golang.org/src/testing/benchmark.go?s=21503:21536#L744)
<pre>func (b *<a href="#B">B</a>) SetParallelism(p <a href="/pkg/builtin/#int">int</a>)</pre>
SetParallelism sets the number of goroutines used by RunParallel to p*GOMAXPROCS.
There is usually no need to call SetParallelism for CPU-bound benchmarks.
If p is less than 1, this call will have no effect.




### <a id="B.Skip">func</a> (\*B) [Skip](https://golang.org/src/testing/testing.go?s=25102:25144#L710)
<pre>func (c *<a href="#B">B</a>) Skip(args ...interface{})</pre>
Skip is equivalent to Log followed by SkipNow.




### <a id="B.SkipNow">func</a> (\*B) [SkipNow](https://golang.org/src/testing/testing.go?s=25820:25846#L729)
<pre>func (c *<a href="#B">B</a>) SkipNow()</pre>
SkipNow marks the test as having been skipped and stops its execution
by calling runtime.Goexit.
If a test fails (see Error, Errorf, Fail) and is then skipped,
it is still considered to have failed.
Execution will continue at the next test or benchmark. See also FailNow.
SkipNow must be called from the goroutine running the test, not from
other goroutines created during the test. Calling SkipNow does not stop
those other goroutines.




### <a id="B.Skipf">func</a> (\*B) [Skipf](https://golang.org/src/testing/testing.go?s=25245:25303#L716)
<pre>func (c *<a href="#B">B</a>) Skipf(format <a href="/pkg/builtin/#string">string</a>, args ...interface{})</pre>
Skipf is equivalent to Logf followed by SkipNow.




### <a id="B.Skipped">func</a> (\*B) [Skipped](https://golang.org/src/testing/testing.go?s=26029:26060#L742)
<pre>func (c *<a href="#B">B</a>) Skipped() <a href="/pkg/builtin/#bool">bool</a></pre>
Skipped reports whether the test was skipped.




### <a id="B.StartTimer">func</a> (\*B) [StartTimer](https://golang.org/src/testing/benchmark.go?s=3490:3514#L110)
<pre>func (b *<a href="#B">B</a>) StartTimer()</pre>
StartTimer starts timing a test. This function is called automatically
before a benchmark starts, but it can also be used to resume timing after
a call to StopTimer.




### <a id="B.StopTimer">func</a> (\*B) [StopTimer](https://golang.org/src/testing/benchmark.go?s=3836:3859#L123)
<pre>func (b *<a href="#B">B</a>) StopTimer()</pre>
StopTimer stops timing a test. This can be used to pause the timer
while performing complex initialization that you don't
want to measure.




## <a id="BenchmarkResult">type</a> [BenchmarkResult](https://golang.org/src/testing/benchmark.go?s=9848:10268#L336)
The results of a benchmark run.


<pre>type BenchmarkResult struct {
<span id="BenchmarkResult.N"></span>    N         <a href="/pkg/builtin/#int">int</a>           <span class="comment">// The number of iterations.</span>
<span id="BenchmarkResult.T"></span>    T         <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a> <span class="comment">// The total time taken.</span>
<span id="BenchmarkResult.Bytes"></span>    Bytes     <a href="/pkg/builtin/#int64">int64</a>         <span class="comment">// Bytes processed in one iteration.</span>
<span id="BenchmarkResult.MemAllocs"></span>    MemAllocs <a href="/pkg/builtin/#uint64">uint64</a>        <span class="comment">// The total number of memory allocations.</span>
<span id="BenchmarkResult.MemBytes"></span>    MemBytes  <a href="/pkg/builtin/#uint64">uint64</a>        <span class="comment">// The total number of bytes allocated.</span>

<span id="BenchmarkResult.Extra"></span>    <span class="comment">// Extra records additional metrics reported by ReportMetric.</span>
    Extra map[<a href="/pkg/builtin/#string">string</a>]<a href="/pkg/builtin/#float64">float64</a>
}
</pre>









### <a id="Benchmark">func</a> [Benchmark](https://golang.org/src/testing/benchmark.go?s=21994:22038#L758)
<pre>func Benchmark(f func(b *<a href="#B">B</a>)) <a href="#BenchmarkResult">BenchmarkResult</a></pre>
Benchmark benchmarks a single function. It is useful for creating
custom benchmarks that do not use the "go test" command.

If f depends on testing flags, then Init must be used to register
those flags before calling Benchmark and before calling flag.Parse.

If f calls Run, the result will be an estimate of running all its
subbenchmarks that don't call Run in sequence in a single benchmark.






### <a id="BenchmarkResult.AllocedBytesPerOp">func</a> (BenchmarkResult) [AllocedBytesPerOp](https://golang.org/src/testing/benchmark.go?s=11102:11152#L383)
<pre>func (r <a href="#BenchmarkResult">BenchmarkResult</a>) AllocedBytesPerOp() <a href="/pkg/builtin/#int64">int64</a></pre>
AllocedBytesPerOp returns the "B/op" metric,
which is calculated as r.MemBytes / r.N.




### <a id="BenchmarkResult.AllocsPerOp">func</a> (BenchmarkResult) [AllocsPerOp](https://golang.org/src/testing/benchmark.go?s=10830:10874#L371)
<pre>func (r <a href="#BenchmarkResult">BenchmarkResult</a>) AllocsPerOp() <a href="/pkg/builtin/#int64">int64</a></pre>
AllocsPerOp returns the "allocs/op" metric,
which is calculated as r.MemAllocs / r.N.




### <a id="BenchmarkResult.MemString">func</a> (BenchmarkResult) [MemString](https://golang.org/src/testing/benchmark.go?s=13057:13100#L461)
<pre>func (r <a href="#BenchmarkResult">BenchmarkResult</a>) MemString() <a href="/pkg/builtin/#string">string</a></pre>
MemString returns r.AllocedBytesPerOp and r.AllocsPerOp in the same format as 'go test'.




### <a id="BenchmarkResult.NsPerOp">func</a> (BenchmarkResult) [NsPerOp](https://golang.org/src/testing/benchmark.go?s=10309:10349#L348)
<pre>func (r <a href="#BenchmarkResult">BenchmarkResult</a>) NsPerOp() <a href="/pkg/builtin/#int64">int64</a></pre>
NsPerOp returns the "ns/op" metric.




### <a id="BenchmarkResult.String">func</a> (BenchmarkResult) [String](https://golang.org/src/testing/benchmark.go?s=11627:11667#L400)
<pre>func (r <a href="#BenchmarkResult">BenchmarkResult</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns a summary of the benchmark results.
It follows the benchmark result line format from
<a href="https://golang.org/design/14313-benchmark-format">https://golang.org/design/14313-benchmark-format</a>, not including the
benchmark name.
Extra metrics override built-in metrics of the same name.
String does not include allocs/op or B/op, since those are reported
by MemString.




## <a id="Cover">type</a> [Cover](https://golang.org/src/testing/cover.go?s=1090:1237#L24)
Cover records information about test coverage checking.
NOTE: This struct is internal to the testing infrastructure and may change.
It is not covered (yet) by the Go 1 compatibility guidelines.


<pre>type Cover struct {
<span id="Cover.Mode"></span>    Mode            <a href="/pkg/builtin/#string">string</a>
<span id="Cover.Counters"></span>    Counters        map[<a href="/pkg/builtin/#string">string</a>][]<a href="/pkg/builtin/#uint32">uint32</a>
<span id="Cover.Blocks"></span>    Blocks          map[<a href="/pkg/builtin/#string">string</a>][]<a href="#CoverBlock">CoverBlock</a>
<span id="Cover.CoveredPackages"></span>    CoveredPackages <a href="/pkg/builtin/#string">string</a>
}
</pre>











## <a id="CoverBlock">type</a> [CoverBlock](https://golang.org/src/testing/cover.go?s=596:868#L11)
CoverBlock records the coverage data for a single basic block.
The fields are 1-indexed, as in an editor: The opening line of
the file is number 1, for example. Columns are measured
in bytes.
NOTE: This struct is internal to the testing infrastructure and may change.
It is not covered (yet) by the Go 1 compatibility guidelines.


<pre>type CoverBlock struct {
<span id="CoverBlock.Line0"></span>    Line0 <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">// Line number for block start.</span>
<span id="CoverBlock.Col0"></span>    Col0  <a href="/pkg/builtin/#uint16">uint16</a> <span class="comment">// Column number for block start.</span>
<span id="CoverBlock.Line1"></span>    Line1 <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">// Line number for block end.</span>
<span id="CoverBlock.Col1"></span>    Col1  <a href="/pkg/builtin/#uint16">uint16</a> <span class="comment">// Column number for block end.</span>
<span id="CoverBlock.Stmts"></span>    Stmts <a href="/pkg/builtin/#uint16">uint16</a> <span class="comment">// Number of statements included in this block.</span>
}
</pre>











## <a id="InternalBenchmark">type</a> [InternalBenchmark](https://golang.org/src/testing/benchmark.go?s=1632:1695#L64)
An internal type but exported because it is cross-package; part of the implementation
of the "go test" command.


<pre>type InternalBenchmark struct {
<span id="InternalBenchmark.Name"></span>    Name <a href="/pkg/builtin/#string">string</a>
<span id="InternalBenchmark.F"></span>    F    func(b *<a href="#B">B</a>)
}
</pre>











## <a id="InternalExample">type</a> [InternalExample](https://golang.org/src/testing/example.go?s=229:330#L5)

<pre>type InternalExample struct {
<span id="InternalExample.Name"></span>    Name      <a href="/pkg/builtin/#string">string</a>
<span id="InternalExample.F"></span>    F         func()
<span id="InternalExample.Output"></span>    Output    <a href="/pkg/builtin/#string">string</a>
<span id="InternalExample.Unordered"></span>    Unordered <a href="/pkg/builtin/#bool">bool</a>
}
</pre>











## <a id="InternalTest">type</a> [InternalTest](https://golang.org/src/testing/testing.go?s=28447:28503#L823)
An internal type but exported because it is cross-package; part of the implementation
of the "go test" command.


<pre>type InternalTest struct {
<span id="InternalTest.Name"></span>    Name <a href="/pkg/builtin/#string">string</a>
<span id="InternalTest.F"></span>    F    func(*<a href="#T">T</a>)
}
</pre>











## <a id="M">type</a> [M](https://golang.org/src/testing/testing.go?s=35253:35438#L1038)
M is a type passed to a TestMain function to run the actual tests.


<pre>type M struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="MainStart">func</a> [MainStart](https://golang.org/src/testing/testing.go?s=36109:36223#L1067)
<pre>func MainStart(deps testDeps, tests []<a href="#InternalTest">InternalTest</a>, benchmarks []<a href="#InternalBenchmark">InternalBenchmark</a>, examples []<a href="#InternalExample">InternalExample</a>) *<a href="#M">M</a></pre>
MainStart is meant for use by tests generated by 'go test'.
It is not meant to be called directly and is not subject to the Go 1 compatibility document.
It may change signature from release to release.






### <a id="M.Run">func</a> (\*M) [Run](https://golang.org/src/testing/testing.go?s=36410:36431#L1078)
<pre>func (m *<a href="#M">M</a>) Run() <a href="/pkg/builtin/#int">int</a></pre>
Run runs the tests. It returns an exit code to pass to os.Exit.




## <a id="PB">type</a> [PB](https://golang.org/src/testing/benchmark.go?s=19070:19352#L667)
A PB is used by RunParallel for running parallel benchmarks.


<pre>type PB struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="PB.Next">func</a> (\*PB) [Next](https://golang.org/src/testing/benchmark.go?s=19416:19441#L675)
<pre>func (pb *<a href="#PB">PB</a>) Next() <a href="/pkg/builtin/#bool">bool</a></pre>
Next reports whether there are more iterations to execute.




## <a id="T">type</a> [T](https://golang.org/src/testing/testing.go?s=20490:20592#L568)
T is a type passed to Test functions to manage test state and support formatted test logs.
Logs are accumulated during execution and dumped to standard output when done.

A test ends when its Test function returns or calls any of the methods
FailNow, Fatal, Fatalf, SkipNow, Skip, or Skipf. Those methods, as well as
the Parallel method, must be called only from the goroutine running the
Test function.

The other reporting methods, such as the variations of Log and Error,
may be called simultaneously from multiple goroutines.


<pre>type T struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="T.Error">func</a> (\*T) [Error](https://golang.org/src/testing/testing.go?s=24490:24533#L686)
<pre>func (c *<a href="#T">T</a>) Error(args ...interface{})</pre>
Error is equivalent to Log followed by Fail.




### <a id="T.Errorf">func</a> (\*T) [Errorf](https://golang.org/src/testing/testing.go?s=24629:24688#L692)
<pre>func (c *<a href="#T">T</a>) Errorf(format <a href="/pkg/builtin/#string">string</a>, args ...interface{})</pre>
Errorf is equivalent to Logf followed by Fail.




### <a id="T.Fail">func</a> (\*T) [Fail](https://golang.org/src/testing/testing.go?s=20928:20951#L591)
<pre>func (c *<a href="#T">T</a>) Fail()</pre>
Fail marks the function as having failed but continues execution.




### <a id="T.FailNow">func</a> (\*T) [FailNow](https://golang.org/src/testing/testing.go?s=21819:21845#L620)
<pre>func (c *<a href="#T">T</a>) FailNow()</pre>
FailNow marks the function as having failed and stops its execution
by calling runtime.Goexit (which then runs all deferred calls in the
current goroutine).
Execution will continue at the next test or benchmark.
FailNow must be called from the goroutine running the
test or benchmark function, not from other goroutines
created during the test. Calling FailNow does not stop
those other goroutines.




### <a id="T.Failed">func</a> (\*T) [Failed](https://golang.org/src/testing/testing.go?s=21261:21291#L605)
<pre>func (c *<a href="#T">T</a>) Failed() <a href="/pkg/builtin/#bool">bool</a></pre>
Failed reports whether the function has failed.




### <a id="T.Fatal">func</a> (\*T) [Fatal](https://golang.org/src/testing/testing.go?s=24792:24835#L698)
<pre>func (c *<a href="#T">T</a>) Fatal(args ...interface{})</pre>
Fatal is equivalent to Log followed by FailNow.




### <a id="T.Fatalf">func</a> (\*T) [Fatalf](https://golang.org/src/testing/testing.go?s=24937:24996#L704)
<pre>func (c *<a href="#T">T</a>) Fatalf(format <a href="/pkg/builtin/#string">string</a>, args ...interface{})</pre>
Fatalf is equivalent to Logf followed by FailNow.




### <a id="T.Helper">func</a> (\*T) [Helper](https://golang.org/src/testing/testing.go?s=26324:26349#L751)
<pre>func (c *<a href="#T">T</a>) Helper()</pre>
Helper marks the calling function as a test helper function.
When printing file and line information, that function will be skipped.
Helper may be called simultaneously from multiple goroutines.




### <a id="T.Log">func</a> (\*T) [Log](https://golang.org/src/testing/testing.go?s=23899:23940#L676)
<pre>func (c *<a href="#T">T</a>) Log(args ...interface{})</pre>
Log formats its arguments using default formatting, analogous to Println,
and records the text in the error log. For tests, the text will be printed only if
the test fails or the -test.v flag is set. For benchmarks, the text is always
printed to avoid having performance depend on the value of the -test.v flag.




### <a id="T.Logf">func</a> (\*T) [Logf](https://golang.org/src/testing/testing.go?s=24343:24400#L683)
<pre>func (c *<a href="#T">T</a>) Logf(format <a href="/pkg/builtin/#string">string</a>, args ...interface{})</pre>
Logf formats its arguments according to the format, analogous to Printf, and
records the text in the error log. A final newline is added if not provided. For
tests, the text will be printed only if the test fails or the -test.v flag is
set. For benchmarks, the text is always printed to avoid having performance
depend on the value of the -test.v flag.




### <a id="T.Name">func</a> (\*T) [Name](https://golang.org/src/testing/testing.go?s=20684:20714#L577)
<pre>func (c *<a href="#T">T</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the name of the running test or benchmark.




### <a id="T.Parallel">func</a> (\*T) [Parallel](https://golang.org/src/testing/testing.go?s=27209:27231#L778)
<pre>func (t *<a href="#T">T</a>) Parallel()</pre>
Parallel signals that this test is to be run in parallel with (and only with)
other parallel tests. When a test is run multiple times due to use of
-test.count or -test.cpu, multiple instances of a single test never run in
parallel with each other.




### <a id="T.Run">func</a> (\*T) [Run](https://golang.org/src/testing/testing.go?s=30908:30955#L911)
<pre>func (t *<a href="#T">T</a>) Run(name <a href="/pkg/builtin/#string">string</a>, f func(t *<a href="#T">T</a>)) <a href="/pkg/builtin/#bool">bool</a></pre>
Run runs f as a subtest of t called name. It runs f in a separate goroutine
and blocks until f returns or calls t.Parallel to become a parallel test.
Run reports whether f succeeded (or at least did not fail before calling t.Parallel).

Run may be called simultaneously from multiple goroutines, but all such calls
must return before the outer test function for t returns.




### <a id="T.Skip">func</a> (\*T) [Skip](https://golang.org/src/testing/testing.go?s=25102:25144#L710)
<pre>func (c *<a href="#T">T</a>) Skip(args ...interface{})</pre>
Skip is equivalent to Log followed by SkipNow.




### <a id="T.SkipNow">func</a> (\*T) [SkipNow](https://golang.org/src/testing/testing.go?s=25820:25846#L729)
<pre>func (c *<a href="#T">T</a>) SkipNow()</pre>
SkipNow marks the test as having been skipped and stops its execution
by calling runtime.Goexit.
If a test fails (see Error, Errorf, Fail) and is then skipped,
it is still considered to have failed.
Execution will continue at the next test or benchmark. See also FailNow.
SkipNow must be called from the goroutine running the test, not from
other goroutines created during the test. Calling SkipNow does not stop
those other goroutines.




### <a id="T.Skipf">func</a> (\*T) [Skipf](https://golang.org/src/testing/testing.go?s=25245:25303#L716)
<pre>func (c *<a href="#T">T</a>) Skipf(format <a href="/pkg/builtin/#string">string</a>, args ...interface{})</pre>
Skipf is equivalent to Logf followed by SkipNow.




### <a id="T.Skipped">func</a> (\*T) [Skipped](https://golang.org/src/testing/testing.go?s=26029:26060#L742)
<pre>func (c *<a href="#T">T</a>) Skipped() <a href="/pkg/builtin/#bool">bool</a></pre>
Skipped reports whether the test was skipped.




## <a id="TB">type</a> [TB](https://golang.org/src/testing/testing.go?s=19346:19887#L532)
TB is the interface common to T and B.


<pre>type TB interface {
    Error(args ...interface{})
    Errorf(format <a href="/pkg/builtin/#string">string</a>, args ...interface{})
    Fail()
    FailNow()
    Failed() <a href="/pkg/builtin/#bool">bool</a>
    Fatal(args ...interface{})
    Fatalf(format <a href="/pkg/builtin/#string">string</a>, args ...interface{})
    Log(args ...interface{})
    Logf(format <a href="/pkg/builtin/#string">string</a>, args ...interface{})
    Name() <a href="/pkg/builtin/#string">string</a>
    Skip(args ...interface{})
    SkipNow()
    Skipf(format <a href="/pkg/builtin/#string">string</a>, args ...interface{})
    Skipped() <a href="/pkg/builtin/#bool">bool</a>
    Helper()
    <span class="comment">// contains filtered or unexported methods</span>
}</pre>














