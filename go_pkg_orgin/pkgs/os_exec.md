

# exec
`import "os/exec"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package exec runs external commands. It wraps os.StartProcess to make it
easier to remap stdin and stdout, connect I/O with pipes, and do other
adjustments.

Unlike the "system" library call from C and other languages, the
os/exec package intentionally does not invoke the system shell and
does not expand any glob patterns or handle other expansions,
pipelines, or redirections typically done by shells. The package
behaves more like C's "exec" family of functions. To expand glob
patterns, either call the shell directly, taking care to escape any
dangerous input, or use the path/filepath package's Glob function.
To expand environment variables, use package os's ExpandEnv.

Note that the examples in this package assume a Unix system.
They may not run on Windows, and they do not run in the Go Playground
used by golang.org and godoc.org.




## <a id="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func LookPath(file string) (string, error)](#LookPath)
* [type Cmd](#Cmd)
  * [func Command(name string, arg ...string) *Cmd](#Command)
  * [func CommandContext(ctx context.Context, name string, arg ...string) *Cmd](#CommandContext)
  * [func (c *Cmd) CombinedOutput() ([]byte, error)](#Cmd.CombinedOutput)
  * [func (c *Cmd) Output() ([]byte, error)](#Cmd.Output)
  * [func (c *Cmd) Run() error](#Cmd.Run)
  * [func (c *Cmd) Start() error](#Cmd.Start)
  * [func (c *Cmd) StderrPipe() (io.ReadCloser, error)](#Cmd.StderrPipe)
  * [func (c *Cmd) StdinPipe() (io.WriteCloser, error)](#Cmd.StdinPipe)
  * [func (c *Cmd) StdoutPipe() (io.ReadCloser, error)](#Cmd.StdoutPipe)
  * [func (c *Cmd) String() string](#Cmd.String)
  * [func (c *Cmd) Wait() error](#Cmd.Wait)
* [type Error](#Error)
  * [func (e *Error) Error() string](#Error.Error)
  * [func (e *Error) Unwrap() error](#Error.Unwrap)
* [type ExitError](#ExitError)
  * [func (e *ExitError) Error() string](#ExitError.Error)


#### <a id="pkg-examples">Examples</a>
* [Cmd.CombinedOutput](#example_Cmd_CombinedOutput)
* [Cmd.Output](#example_Cmd_Output)
* [Cmd.Run](#example_Cmd_Run)
* [Cmd.Start](#example_Cmd_Start)
* [Cmd.StderrPipe](#example_Cmd_StderrPipe)
* [Cmd.StdinPipe](#example_Cmd_StdinPipe)
* [Cmd.StdoutPipe](#example_Cmd_StdoutPipe)
* [Command](#example_Command)
* [CommandContext](#example_CommandContext)
* [Command (Environment)](#example_Command_environment)
* [LookPath](#example_LookPath)


#### <a id="pkg-files">Package files</a>
[exec.go](https://golang.org/src/os/exec/exec.go) [exec_unix.go](https://golang.org/src/os/exec/exec_unix.go) [lp_unix.go](https://golang.org/src/os/exec/lp_unix.go) 




## <a id="pkg-variables">Variables</a>
ErrNotFound is the error resulting if a path search failed to find an executable file.


<pre>var <span id="ErrNotFound">ErrNotFound</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;executable file not found in $PATH&#34;)</pre>

## <a id="LookPath">func</a> [LookPath](https://golang.org/src/os/exec/lp_unix.go?s=928:970#L24)
<pre>func LookPath(file <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
LookPath searches for an executable named file in the
directories named by the PATH environment variable.
If file contains a slash, it is tried directly and the PATH is not consulted.
The result may be an absolute path or a path relative to the current directory.



<a id="example_LookPath">Example</a>


```go
```

output:
```txt
```



## <a id="Cmd">type</a> [Cmd](https://golang.org/src/os/exec/exec.go?s=1709:5119#L46)
Cmd represents an external command being prepared or run.

A Cmd cannot be reused after calling its Run, Output or CombinedOutput
methods.


<pre>type Cmd struct {
<span id="Cmd.Path"></span>    <span class="comment">// Path is the path of the command to run.</span>
    <span class="comment">//</span>
    <span class="comment">// This is the only field that must be set to a non-zero</span>
    <span class="comment">// value. If Path is relative, it is evaluated relative</span>
    <span class="comment">// to Dir.</span>
    Path <a href="/pkg/builtin/#string">string</a>

<span id="Cmd.Args"></span>    <span class="comment">// Args holds command line arguments, including the command as Args[0].</span>
    <span class="comment">// If the Args field is empty or nil, Run uses {Path}.</span>
    <span class="comment">//</span>
    <span class="comment">// In typical use, both Path and Args are set by calling Command.</span>
    Args []<a href="/pkg/builtin/#string">string</a>

<span id="Cmd.Env"></span>    <span class="comment">// Env specifies the environment of the process.</span>
    <span class="comment">// Each entry is of the form &#34;key=value&#34;.</span>
    <span class="comment">// If Env is nil, the new process uses the current process&#39;s</span>
    <span class="comment">// environment.</span>
    <span class="comment">// If Env contains duplicate environment keys, only the last</span>
    <span class="comment">// value in the slice for each duplicate key is used.</span>
    <span class="comment">// As a special case on Windows, SYSTEMROOT is always added if</span>
    <span class="comment">// missing and not explicitly set to the empty string.</span>
    Env []<a href="/pkg/builtin/#string">string</a>

<span id="Cmd.Dir"></span>    <span class="comment">// Dir specifies the working directory of the command.</span>
    <span class="comment">// If Dir is the empty string, Run runs the command in the</span>
    <span class="comment">// calling process&#39;s current directory.</span>
    Dir <a href="/pkg/builtin/#string">string</a>

<span id="Cmd.Stdin"></span>    <span class="comment">// Stdin specifies the process&#39;s standard input.</span>
    <span class="comment">//</span>
    <span class="comment">// If Stdin is nil, the process reads from the null device (os.DevNull).</span>
    <span class="comment">//</span>
    <span class="comment">// If Stdin is an *os.File, the process&#39;s standard input is connected</span>
    <span class="comment">// directly to that file.</span>
    <span class="comment">//</span>
    <span class="comment">// Otherwise, during the execution of the command a separate</span>
    <span class="comment">// goroutine reads from Stdin and delivers that data to the command</span>
    <span class="comment">// over a pipe. In this case, Wait does not complete until the goroutine</span>
    <span class="comment">// stops copying, either because it has reached the end of Stdin</span>
    <span class="comment">// (EOF or a read error) or because writing to the pipe returned an error.</span>
    Stdin <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>

<span id="Cmd.Stdout"></span>    <span class="comment">// Stdout and Stderr specify the process&#39;s standard output and error.</span>
    <span class="comment">//</span>
    <span class="comment">// If either is nil, Run connects the corresponding file descriptor</span>
    <span class="comment">// to the null device (os.DevNull).</span>
    <span class="comment">//</span>
    <span class="comment">// If either is an *os.File, the corresponding output from the process</span>
    <span class="comment">// is connected directly to that file.</span>
    <span class="comment">//</span>
    <span class="comment">// Otherwise, during the execution of the command a separate goroutine</span>
    <span class="comment">// reads from the process over a pipe and delivers that data to the</span>
    <span class="comment">// corresponding Writer. In this case, Wait does not complete until the</span>
    <span class="comment">// goroutine reaches EOF or encounters an error.</span>
    <span class="comment">//</span>
    <span class="comment">// If Stdout and Stderr are the same writer, and have a type that can</span>
    <span class="comment">// be compared with ==, at most one goroutine at a time will call Write.</span>
    Stdout <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>
<span id="Cmd.Stderr"></span>    Stderr <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>

<span id="Cmd.ExtraFiles"></span>    <span class="comment">// ExtraFiles specifies additional open files to be inherited by the</span>
    <span class="comment">// new process. It does not include standard input, standard output, or</span>
    <span class="comment">// standard error. If non-nil, entry i becomes file descriptor 3+i.</span>
    <span class="comment">//</span>
    <span class="comment">// ExtraFiles is not supported on Windows.</span>
    ExtraFiles []*<a href="/pkg/os/">os</a>.<a href="/pkg/os/#File">File</a>

<span id="Cmd.SysProcAttr"></span>    <span class="comment">// SysProcAttr holds optional, operating system-specific attributes.</span>
    <span class="comment">// Run passes it to os.StartProcess as the os.ProcAttr&#39;s Sys field.</span>
    SysProcAttr *<a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#SysProcAttr">SysProcAttr</a>

<span id="Cmd.Process"></span>    <span class="comment">// Process is the underlying process, once started.</span>
    Process *<a href="/pkg/os/">os</a>.<a href="/pkg/os/#Process">Process</a>

<span id="Cmd.ProcessState"></span>    <span class="comment">// ProcessState contains information about an exited process,</span>
    <span class="comment">// available after a call to Wait or Run.</span>
    ProcessState *<a href="/pkg/os/">os</a>.<a href="/pkg/os/#ProcessState">ProcessState</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="Command">func</a> [Command](https://golang.org/src/os/exec/exec.go?s=6250:6295#L158)
<pre>func Command(name <a href="/pkg/builtin/#string">string</a>, arg ...<a href="/pkg/builtin/#string">string</a>) *<a href="#Cmd">Cmd</a></pre>
Command returns the Cmd struct to execute the named program with
the given arguments.

It sets only the Path and Args in the returned structure.

If name contains no path separators, Command uses LookPath to
resolve name to a complete path if possible. Otherwise it uses name
directly as Path.

The returned Cmd's Args field is constructed from the command name
followed by the elements of arg, so arg should not include the
command name itself. For example, Command("echo", "hello").
Args[0] is always name, not the possibly resolved Path.

On Windows, processes receive the whole command line as a single string
and do their own parsing. Command combines and quotes Args into a command
line string with an algorithm compatible with applications using
CommandLineToArgvW (which is the most common way). Notable exceptions are
msiexec.exe and cmd.exe (and thus, all batch files), which have a different
unquoting algorithm. In these or other similar cases, you can do the
quoting yourself and provide the full command line in SysProcAttr.CmdLine,
leaving Args empty.



<a id="example_Command">Example</a>


```go
```

output:
```txt
```

<a id="example_Command_environment">Example (Environment)</a>


```go
```

output:
```txt
```


### <a id="CommandContext">func</a> [CommandContext](https://golang.org/src/os/exec/exec.go?s=6740:6813#L178)
<pre>func CommandContext(ctx <a href="/pkg/context/">context</a>.<a href="/pkg/context/#Context">Context</a>, name <a href="/pkg/builtin/#string">string</a>, arg ...<a href="/pkg/builtin/#string">string</a>) *<a href="#Cmd">Cmd</a></pre>
CommandContext is like Command but includes a context.

The provided context is used to kill the process (by calling
os.Process.Kill) if the context becomes done before the command
completes on its own.



<a id="example_CommandContext">Example</a>


```go
```

output:
```txt
```




### <a id="Cmd.CombinedOutput">func</a> (\*Cmd) [CombinedOutput](https://golang.org/src/os/exec/exec.go?s=15900:15946#L541)
<pre>func (c *<a href="#Cmd">Cmd</a>) CombinedOutput() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
CombinedOutput runs the command and returns its combined standard
output and standard error.



<a id="example_Cmd_CombinedOutput">Example</a>


```go
```

output:
```txt
```


### <a id="Cmd.Output">func</a> (\*Cmd) [Output](https://golang.org/src/os/exec/exec.go?s=15363:15401#L518)
<pre>func (c *<a href="#Cmd">Cmd</a>) Output() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Output runs the command and returns its standard output.
Any returned error will usually be of type *ExitError.
If c.Stderr was nil, Output populates ExitError.Stderr.



<a id="example_Cmd_Output">Example</a>


```go
```

output:
```txt
```


### <a id="Cmd.Run">func</a> (\*Cmd) [Run](https://golang.org/src/os/exec/exec.go?s=10431:10456#L327)
<pre>func (c *<a href="#Cmd">Cmd</a>) Run() <a href="/pkg/builtin/#error">error</a></pre>
Run starts the specified command and waits for it to complete.

The returned error is nil if the command runs, has no problems
copying stdin, stdout, and stderr, and exits with a zero exit
status.

If the command starts but does not complete successfully, the error is of
type *ExitError. Other error types may be returned for other situations.

If the calling goroutine has locked the operating system thread
with runtime.LockOSThread and modified any inheritable OS-level
thread state (for example, Linux or Plan 9 name spaces), the new
process will inherit the caller's thread state.



<a id="example_Cmd_Run">Example</a>


```go
```

output:
```txt
```


### <a id="Cmd.Start">func</a> (\*Cmd) [Start](https://golang.org/src/os/exec/exec.go?s=11462:11489#L364)
<pre>func (c *<a href="#Cmd">Cmd</a>) Start() <a href="/pkg/builtin/#error">error</a></pre>
Start starts the specified command but does not wait for it to complete.

The Wait method will return the exit code and release associated resources
once the command exits.



<a id="example_Cmd_Start">Example</a>


```go
```

output:
```txt
```


### <a id="Cmd.StderrPipe">func</a> (\*Cmd) [StderrPipe](https://golang.org/src/os/exec/exec.go?s=18556:18605#L628)
<pre>func (c *<a href="#Cmd">Cmd</a>) StderrPipe() (<a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadCloser">ReadCloser</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
StderrPipe returns a pipe that will be connected to the command's
standard error when the command starts.

Wait will close the pipe after seeing the command exit, so most callers
need not close the pipe themselves; however, an implication is that
it is incorrect to call Wait before all reads from the pipe have completed.
For the same reason, it is incorrect to use Run when using StderrPipe.
See the StdoutPipe example for idiomatic usage.



<a id="example_Cmd_StderrPipe">Example</a>


```go
```

output:
```txt
```


### <a id="Cmd.StdinPipe">func</a> (\*Cmd) [StdinPipe](https://golang.org/src/os/exec/exec.go?s=16573:16622#L561)
<pre>func (c *<a href="#Cmd">Cmd</a>) StdinPipe() (<a href="/pkg/io/">io</a>.<a href="/pkg/io/#WriteCloser">WriteCloser</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
StdinPipe returns a pipe that will be connected to the command's
standard input when the command starts.
The pipe will be closed automatically after Wait sees the command exit.
A caller need only call Close to force the pipe to close sooner.
For example, if the command being run will not exit until standard input
is closed, the caller must close the pipe.



<a id="example_Cmd_StdinPipe">Example</a>


```go
```

output:
```txt
```


### <a id="Cmd.StdoutPipe">func</a> (\*Cmd) [StdoutPipe](https://golang.org/src/os/exec/exec.go?s=17670:17719#L603)
<pre>func (c *<a href="#Cmd">Cmd</a>) StdoutPipe() (<a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadCloser">ReadCloser</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
StdoutPipe returns a pipe that will be connected to the command's
standard output when the command starts.

Wait will close the pipe after seeing the command exit, so most callers
need not close the pipe themselves; however, an implication is that
it is incorrect to call Wait before all reads from the pipe have completed.
For the same reason, it is incorrect to call Run when using StdoutPipe.
See the example for idiomatic usage.



<a id="example_Cmd_StdoutPipe">Example</a>


```go
```

output:
```txt
```


### <a id="Cmd.String">func</a> (\*Cmd) [String](https://golang.org/src/os/exec/exec.go?s=7129:7158#L191)
<pre>func (c *<a href="#Cmd">Cmd</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns a human-readable description of c.
It is intended only for debugging.
In particular, it is not suitable for use as input to a shell.
The output of String may vary across Go releases.




### <a id="Cmd.Wait">func</a> (\*Cmd) [Wait](https://golang.org/src/os/exec/exec.go?s=14599:14625#L482)
<pre>func (c *<a href="#Cmd">Cmd</a>) Wait() <a href="/pkg/builtin/#error">error</a></pre>
Wait waits for the command to exit and waits for any copying to
stdin or copying from stdout or stderr to complete.

The command must have been started by Start.

The returned error is nil if the command runs, has no problems
copying stdin, stdout, and stderr, and exits with a zero exit
status.

If the command fails to run or doesn't complete successfully, the
error is of type *ExitError. Other error types may be
returned for I/O problems.

If any of c.Stdin, c.Stdout or c.Stderr are not an *os.File, Wait also waits
for the respective I/O loop copying to or from the process to complete.

Wait releases any resources associated with the Cmd.




## <a id="Error">type</a> [Error](https://golang.org/src/os/exec/exec.go?s=1274:1408#L29)
Error is returned by LookPath when it fails to classify a file as an
executable.


<pre>type Error struct {
<span id="Error.Name"></span>    <span class="comment">// Name is the file name for which the error occurred.</span>
    Name <a href="/pkg/builtin/#string">string</a>
<span id="Error.Err"></span>    <span class="comment">// Err is the underlying error.</span>
    Err <a href="/pkg/builtin/#error">error</a>
}
</pre>











### <a id="Error.Error">func</a> (\*Error) [Error](https://golang.org/src/os/exec/exec.go?s=1410:1440#L36)
<pre>func (e *<a href="#Error">Error</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Error.Unwrap">func</a> (\*Error) [Unwrap](https://golang.org/src/os/exec/exec.go?s=1510:1540#L40)
<pre>func (e *<a href="#Error">Error</a>) Unwrap() <a href="/pkg/builtin/#error">error</a></pre>



## <a id="ExitError">type</a> [ExitError](https://golang.org/src/os/exec/exec.go?s=13315:13831#L445)
An ExitError reports an unsuccessful exit by a command.


<pre>type ExitError struct {
    *<a href="/pkg/os/">os</a>.<a href="/pkg/os/#ProcessState">ProcessState</a>

<span id="ExitError.Stderr"></span>    <span class="comment">// Stderr holds a subset of the standard error output from the</span>
    <span class="comment">// Cmd.Output method if standard error was not otherwise being</span>
    <span class="comment">// collected.</span>
    <span class="comment">//</span>
    <span class="comment">// If the error output is long, Stderr may contain only a prefix</span>
    <span class="comment">// and suffix of the output, with the middle replaced with</span>
    <span class="comment">// text about the number of omitted bytes.</span>
    <span class="comment">//</span>
    <span class="comment">// Stderr is provided for debugging, for inclusion in error messages.</span>
    <span class="comment">// Users with other needs should redirect Cmd.Stderr as needed.</span>
    Stderr []<a href="/pkg/builtin/#byte">byte</a>
}
</pre>











### <a id="ExitError.Error">func</a> (\*ExitError) [Error](https://golang.org/src/os/exec/exec.go?s=13833:13867#L461)
<pre>func (e *<a href="#ExitError">ExitError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>






