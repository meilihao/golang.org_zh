

# os
`import "os"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package os provides a platform-independent interface to operating system
functionality. The design is Unix-like, although the error handling is
Go-like; failing calls return values of type error rather than error numbers.
Often, more information is available within the error. For example,
if a call that takes a file name fails, such as Open or Stat, the error
will include the failing file name when printed and will be of type
*PathError, which may be unpacked for more information.

The os interface is intended to be uniform across all operating systems.
Features not generally available appear in the system-specific package syscall.

Here is a simple example, opening a file and reading some of it.


	file, err := os.Open("file.go") // For read access.
	if err != nil {
		log.Fatal(err)
	}

If the open fails, the error string will be self-explanatory, like


	open file.go: no such file or directory

The file's data can then be read into a slice of bytes. Read and
Write take their byte counts from the length of the argument slice.


	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])

Note: The maximum number of concurrent operations on a File may be limited by
the OS or the system. The number should be high, but exceeding it may degrade
performance or cause other issues.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func Chdir(dir string) error](#Chdir)
* [func Chmod(name string, mode FileMode) error](#Chmod)
* [func Chown(name string, uid, gid int) error](#Chown)
* [func Chtimes(name string, atime time.Time, mtime time.Time) error](#Chtimes)
* [func Clearenv()](#Clearenv)
* [func Environ() []string](#Environ)
* [func Executable() (string, error)](#Executable)
* [func Exit(code int)](#Exit)
* [func Expand(s string, mapping func(string) string) string](#Expand)
* [func ExpandEnv(s string) string](#ExpandEnv)
* [func Getegid() int](#Getegid)
* [func Getenv(key string) string](#Getenv)
* [func Geteuid() int](#Geteuid)
* [func Getgid() int](#Getgid)
* [func Getgroups() ([]int, error)](#Getgroups)
* [func Getpagesize() int](#Getpagesize)
* [func Getpid() int](#Getpid)
* [func Getppid() int](#Getppid)
* [func Getuid() int](#Getuid)
* [func Getwd() (dir string, err error)](#Getwd)
* [func Hostname() (name string, err error)](#Hostname)
* [func IsExist(err error) bool](#IsExist)
* [func IsNotExist(err error) bool](#IsNotExist)
* [func IsPathSeparator(c uint8) bool](#IsPathSeparator)
* [func IsPermission(err error) bool](#IsPermission)
* [func IsTimeout(err error) bool](#IsTimeout)
* [func Lchown(name string, uid, gid int) error](#Lchown)
* [func Link(oldname, newname string) error](#Link)
* [func LookupEnv(key string) (string, bool)](#LookupEnv)
* [func Mkdir(name string, perm FileMode) error](#Mkdir)
* [func MkdirAll(path string, perm FileMode) error](#MkdirAll)
* [func NewSyscallError(syscall string, err error) error](#NewSyscallError)
* [func Pipe() (r *File, w *File, err error)](#Pipe)
* [func Readlink(name string) (string, error)](#Readlink)
* [func Remove(name string) error](#Remove)
* [func RemoveAll(path string) error](#RemoveAll)
* [func Rename(oldpath, newpath string) error](#Rename)
* [func SameFile(fi1, fi2 FileInfo) bool](#SameFile)
* [func Setenv(key, value string) error](#Setenv)
* [func Symlink(oldname, newname string) error](#Symlink)
* [func TempDir() string](#TempDir)
* [func Truncate(name string, size int64) error](#Truncate)
* [func Unsetenv(key string) error](#Unsetenv)
* [func UserCacheDir() (string, error)](#UserCacheDir)
* [func UserConfigDir() (string, error)](#UserConfigDir)
* [func UserHomeDir() (string, error)](#UserHomeDir)
* [type File](#File)
  * [func Create(name string) (*File, error)](#Create)
  * [func NewFile(fd uintptr, name string) *File](#NewFile)
  * [func Open(name string) (*File, error)](#Open)
  * [func OpenFile(name string, flag int, perm FileMode) (*File, error)](#OpenFile)
  * [func (f *File) Chdir() error](#File.Chdir)
  * [func (f *File) Chmod(mode FileMode) error](#File.Chmod)
  * [func (f *File) Chown(uid, gid int) error](#File.Chown)
  * [func (f *File) Close() error](#File.Close)
  * [func (f *File) Fd() uintptr](#File.Fd)
  * [func (f *File) Name() string](#File.Name)
  * [func (f *File) Read(b []byte) (n int, err error)](#File.Read)
  * [func (f *File) ReadAt(b []byte, off int64) (n int, err error)](#File.ReadAt)
  * [func (f *File) Readdir(n int) ([]FileInfo, error)](#File.Readdir)
  * [func (f *File) Readdirnames(n int) (names []string, err error)](#File.Readdirnames)
  * [func (f *File) Seek(offset int64, whence int) (ret int64, err error)](#File.Seek)
  * [func (f *File) SetDeadline(t time.Time) error](#File.SetDeadline)
  * [func (f *File) SetReadDeadline(t time.Time) error](#File.SetReadDeadline)
  * [func (f *File) SetWriteDeadline(t time.Time) error](#File.SetWriteDeadline)
  * [func (f *File) Stat() (FileInfo, error)](#File.Stat)
  * [func (f *File) Sync() error](#File.Sync)
  * [func (f *File) SyscallConn() (syscall.RawConn, error)](#File.SyscallConn)
  * [func (f *File) Truncate(size int64) error](#File.Truncate)
  * [func (f *File) Write(b []byte) (n int, err error)](#File.Write)
  * [func (f *File) WriteAt(b []byte, off int64) (n int, err error)](#File.WriteAt)
  * [func (f *File) WriteString(s string) (n int, err error)](#File.WriteString)
* [type FileInfo](#FileInfo)
  * [func Lstat(name string) (FileInfo, error)](#Lstat)
  * [func Stat(name string) (FileInfo, error)](#Stat)
* [type FileMode](#FileMode)
  * [func (m FileMode) IsDir() bool](#FileMode.IsDir)
  * [func (m FileMode) IsRegular() bool](#FileMode.IsRegular)
  * [func (m FileMode) Perm() FileMode](#FileMode.Perm)
  * [func (m FileMode) String() string](#FileMode.String)
* [type LinkError](#LinkError)
  * [func (e *LinkError) Error() string](#LinkError.Error)
  * [func (e *LinkError) Unwrap() error](#LinkError.Unwrap)
* [type PathError](#PathError)
  * [func (e *PathError) Error() string](#PathError.Error)
  * [func (e *PathError) Timeout() bool](#PathError.Timeout)
  * [func (e *PathError) Unwrap() error](#PathError.Unwrap)
* [type ProcAttr](#ProcAttr)
* [type Process](#Process)
  * [func FindProcess(pid int) (*Process, error)](#FindProcess)
  * [func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)](#StartProcess)
  * [func (p *Process) Kill() error](#Process.Kill)
  * [func (p *Process) Release() error](#Process.Release)
  * [func (p *Process) Signal(sig Signal) error](#Process.Signal)
  * [func (p *Process) Wait() (*ProcessState, error)](#Process.Wait)
* [type ProcessState](#ProcessState)
  * [func (p *ProcessState) ExitCode() int](#ProcessState.ExitCode)
  * [func (p *ProcessState) Exited() bool](#ProcessState.Exited)
  * [func (p *ProcessState) Pid() int](#ProcessState.Pid)
  * [func (p *ProcessState) String() string](#ProcessState.String)
  * [func (p *ProcessState) Success() bool](#ProcessState.Success)
  * [func (p *ProcessState) Sys() interface{}](#ProcessState.Sys)
  * [func (p *ProcessState) SysUsage() interface{}](#ProcessState.SysUsage)
  * [func (p *ProcessState) SystemTime() time.Duration](#ProcessState.SystemTime)
  * [func (p *ProcessState) UserTime() time.Duration](#ProcessState.UserTime)
* [type Signal](#Signal)
* [type SyscallError](#SyscallError)
  * [func (e *SyscallError) Error() string](#SyscallError.Error)
  * [func (e *SyscallError) Timeout() bool](#SyscallError.Timeout)
  * [func (e *SyscallError) Unwrap() error](#SyscallError.Unwrap)


#### <a id="pkg-examples">Examples</a>
* [Chmod](#example_Chmod)
* [Chtimes](#example_Chtimes)
* [Expand](#example_Expand)
* [ExpandEnv](#example_ExpandEnv)
* [FileMode](#example_FileMode)
* [Getenv](#example_Getenv)
* [IsNotExist](#example_IsNotExist)
* [LookupEnv](#example_LookupEnv)
* [OpenFile](#example_OpenFile)
* [OpenFile (Append)](#example_OpenFile_append)
* [Unsetenv](#example_Unsetenv)


#### <a id="pkg-files">Package files</a>
[dir.go](https://golang.org/src/os/dir.go) [dir_unix.go](https://golang.org/src/os/dir_unix.go) [env.go](https://golang.org/src/os/env.go) [env_default.go](https://golang.org/src/os/env_default.go) [error.go](https://golang.org/src/os/error.go) [error_errno.go](https://golang.org/src/os/error_errno.go) [error_posix.go](https://golang.org/src/os/error_posix.go) [exec.go](https://golang.org/src/os/exec.go) [exec_posix.go](https://golang.org/src/os/exec_posix.go) [exec_unix.go](https://golang.org/src/os/exec_unix.go) [executable.go](https://golang.org/src/os/executable.go) [executable_procfs.go](https://golang.org/src/os/executable_procfs.go) [file.go](https://golang.org/src/os/file.go) [file_posix.go](https://golang.org/src/os/file_posix.go) [file_unix.go](https://golang.org/src/os/file_unix.go) [getwd.go](https://golang.org/src/os/getwd.go) [path.go](https://golang.org/src/os/path.go) [path_unix.go](https://golang.org/src/os/path_unix.go) [pipe_linux.go](https://golang.org/src/os/pipe_linux.go) [proc.go](https://golang.org/src/os/proc.go) [rawconn.go](https://golang.org/src/os/rawconn.go) [removeall_at.go](https://golang.org/src/os/removeall_at.go) [stat.go](https://golang.org/src/os/stat.go) [stat_linux.go](https://golang.org/src/os/stat_linux.go) [stat_unix.go](https://golang.org/src/os/stat_unix.go) [sticky_notbsd.go](https://golang.org/src/os/sticky_notbsd.go) [str.go](https://golang.org/src/os/str.go) [sys.go](https://golang.org/src/os/sys.go) [sys_linux.go](https://golang.org/src/os/sys_linux.go) [sys_unix.go](https://golang.org/src/os/sys_unix.go) [types.go](https://golang.org/src/os/types.go) [types_unix.go](https://golang.org/src/os/types_unix.go) [wait_waitid.go](https://golang.org/src/os/wait_waitid.go) 


## <a id="pkg-constants">Constants</a>
Flags to OpenFile wrapping those of the underlying system. Not all
flags may be implemented on a given system.


<pre>const (
    <span class="comment">// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.</span>
    <span id="O_RDONLY">O_RDONLY</span> <a href="/pkg/builtin/#int">int</a> = <a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#O_RDONLY">O_RDONLY</a> <span class="comment">// open the file read-only.</span>
    <span id="O_WRONLY">O_WRONLY</span> <a href="/pkg/builtin/#int">int</a> = <a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#O_WRONLY">O_WRONLY</a> <span class="comment">// open the file write-only.</span>
    <span id="O_RDWR">O_RDWR</span>   <a href="/pkg/builtin/#int">int</a> = <a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#O_RDWR">O_RDWR</a>   <span class="comment">// open the file read-write.</span>
    <span class="comment">// The remaining values may be or&#39;ed in to control behavior.</span>
    <span id="O_APPEND">O_APPEND</span> <a href="/pkg/builtin/#int">int</a> = <a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#O_APPEND">O_APPEND</a> <span class="comment">// append data to the file when writing.</span>
    <span id="O_CREATE">O_CREATE</span> <a href="/pkg/builtin/#int">int</a> = <a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#O_CREAT">O_CREAT</a>  <span class="comment">// create a new file if none exists.</span>
    <span id="O_EXCL">O_EXCL</span>   <a href="/pkg/builtin/#int">int</a> = <a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#O_EXCL">O_EXCL</a>   <span class="comment">// used with O_CREATE, file must not exist.</span>
    <span id="O_SYNC">O_SYNC</span>   <a href="/pkg/builtin/#int">int</a> = <a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#O_SYNC">O_SYNC</a>   <span class="comment">// open for synchronous I/O.</span>
    <span id="O_TRUNC">O_TRUNC</span>  <a href="/pkg/builtin/#int">int</a> = <a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#O_TRUNC">O_TRUNC</a>  <span class="comment">// truncate regular writable file when opened.</span>
)</pre>Seek whence values.

Deprecated: Use io.SeekStart, io.SeekCurrent, and io.SeekEnd.


<pre>const (
    <span id="SEEK_SET">SEEK_SET</span> <a href="/pkg/builtin/#int">int</a> = 0 <span class="comment">// seek relative to the origin of the file</span>
    <span id="SEEK_CUR">SEEK_CUR</span> <a href="/pkg/builtin/#int">int</a> = 1 <span class="comment">// seek relative to the current offset</span>
    <span id="SEEK_END">SEEK_END</span> <a href="/pkg/builtin/#int">int</a> = 2 <span class="comment">// seek relative to the end</span>
)</pre>
<pre>const (
    <span id="PathSeparator">PathSeparator</span>     = &#39;/&#39; <span class="comment">// OS-specific path separator</span>
    <span id="PathListSeparator">PathListSeparator</span> = &#39;:&#39; <span class="comment">// OS-specific path list separator</span>
)</pre>DevNull is the name of the operating system's ``null device.''
On Unix-like systems, it is "/dev/null"; on Windows, "NUL".


<pre>const <span id="DevNull">DevNull</span> = &#34;/dev/null&#34;</pre>

## <a id="pkg-variables">Variables</a>
Portable analogs of some common system call errors.

Errors returned from this package may be tested against these errors
with errors.Is.


<pre>var (
    <span class="comment">// ErrInvalid indicates an invalid argument.</span>
    <span class="comment">// Methods on File will return this error when the receiver is nil.</span>
    <span id="ErrInvalid">ErrInvalid</span> = errInvalid() <span class="comment">// &#34;invalid argument&#34;</span>

    <span id="ErrPermission">ErrPermission</span> = errPermission() <span class="comment">// &#34;permission denied&#34;</span>
    <span id="ErrExist">ErrExist</span>      = errExist()      <span class="comment">// &#34;file already exists&#34;</span>
    <span id="ErrNotExist">ErrNotExist</span>   = errNotExist()   <span class="comment">// &#34;file does not exist&#34;</span>
    <span id="ErrClosed">ErrClosed</span>     = errClosed()     <span class="comment">// &#34;file already closed&#34;</span>
    <span id="ErrNoDeadline">ErrNoDeadline</span> = errNoDeadline() <span class="comment">// &#34;file type does not support deadline&#34;</span>
)</pre>Stdin, Stdout, and Stderr are open Files pointing to the standard input,
standard output, and standard error file descriptors.

Note that the Go runtime writes to standard error for panics and crashes;
closing Stderr may cause those messages to go elsewhere, perhaps
to a file opened later.


<pre>var (
    <span id="Stdin">Stdin</span>  = <a href="#NewFile">NewFile</a>(<a href="/pkg/builtin/#uintptr">uintptr</a>(<a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#Stdin">Stdin</a>), &#34;/dev/stdin&#34;)
    <span id="Stdout">Stdout</span> = <a href="#NewFile">NewFile</a>(<a href="/pkg/builtin/#uintptr">uintptr</a>(<a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#Stdout">Stdout</a>), &#34;/dev/stdout&#34;)
    <span id="Stderr">Stderr</span> = <a href="#NewFile">NewFile</a>(<a href="/pkg/builtin/#uintptr">uintptr</a>(<a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#Stderr">Stderr</a>), &#34;/dev/stderr&#34;)
)</pre>Args hold the command-line arguments, starting with the program name.


<pre>var <span id="Args">Args</span> []<a href="/pkg/builtin/#string">string</a></pre>

## <a id="Chdir">func</a> [Chdir](https://golang.org/src/os/file.go?s=7660:7688#L251)
<pre>func Chdir(dir <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Chdir changes the current working directory to the named directory.
If there is an error, it will be of type *PathError.



## <a id="Chmod">func</a> [Chmod](https://golang.org/src/os/file.go?s=15028:15072#L495)
<pre>func Chmod(name <a href="/pkg/builtin/#string">string</a>, mode <a href="#FileMode">FileMode</a>) <a href="/pkg/builtin/#error">error</a></pre>
Chmod changes the mode of the named file to mode.
If the file is a symbolic link, it changes the mode of the link's target.
If there is an error, it will be of type *PathError.

A different subset of the mode bits are used, depending on the
operating system.

On Unix, the mode's permission bits, ModeSetuid, ModeSetgid, and
ModeSticky are used.

On Windows, only the 0200 bit (owner writable) of mode is used; it
controls whether the file's read-only attribute is set or cleared.
The other bits are currently unused. For compatibility with Go 1.12
and earlier, use a non-zero mode. Use mode 0400 for a read-only
file and 0600 for a readable+writable file.

On Plan 9, the mode's permission bits, ModeAppend, ModeExclusive,
and ModeTemporary are used.


<a id="example_Chmod">Example</a>
```go
```

output:
```txt
```

## <a id="Chown">func</a> [Chown](https://golang.org/src/os/file_posix.go?s=1525:1568#L48)
<pre>func Chown(name <a href="/pkg/builtin/#string">string</a>, uid, gid <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#error">error</a></pre>
Chown changes the numeric uid and gid of the named file.
If the file is a symbolic link, it changes the uid and gid of the link's target.
A uid or gid of -1 means to not change that value.
If there is an error, it will be of type *PathError.

On Windows or Plan 9, Chown always returns the syscall.EWINDOWS or
EPLAN9 error, wrapped in *PathError.



## <a id="Chtimes">func</a> [Chtimes](https://golang.org/src/os/file_posix.go?s=3495:3560#L115)
<pre>func Chtimes(name <a href="/pkg/builtin/#string">string</a>, atime <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>, mtime <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
Chtimes changes the access and modification times of the named
file, similar to the Unix utime() or utimes() functions.

The underlying filesystem may truncate or round the values to a
less precise time unit.
If there is an error, it will be of type *PathError.


<a id="example_Chtimes">Example</a>
```go
```

output:
```txt
```

## <a id="Clearenv">func</a> [Clearenv](https://golang.org/src/os/env.go?s=3747:3762#L123)
<pre>func Clearenv()</pre>
Clearenv deletes all environment variables.



## <a id="Environ">func</a> [Environ](https://golang.org/src/os/env.go?s=3883:3906#L129)
<pre>func Environ() []<a href="/pkg/builtin/#string">string</a></pre>
Environ returns a copy of strings representing the environment,
in the form "key=value".



## <a id="Executable">func</a> [Executable](https://golang.org/src/os/executable.go?s=758:791#L10)
<pre>func Executable() (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Executable returns the path name for the executable that started
the current process. There is no guarantee that the path is still
pointing to the correct executable. If a symlink was used to start
the process, depending on the operating system, the result might
be the symlink or the path it pointed to. If a stable result is
needed, path/filepath.EvalSymlinks might help.

Executable returns an absolute path unless an error occurred.

The main use case is finding resources located relative to an
executable.

Executable is not supported on nacl.



## <a id="Exit">func</a> [Exit](https://golang.org/src/os/proc.go?s=1631:1650#L51)
<pre>func Exit(code <a href="/pkg/builtin/#int">int</a>)</pre>
Exit causes the current program to exit with the given status code.
Conventionally, code zero indicates success, non-zero an error.
The program terminates immediately; deferred functions are not run.

For portability, the status code should be in the range [0, 125].



## <a id="Expand">func</a> [Expand](https://golang.org/src/os/env.go?s=403:460#L6)
<pre>func Expand(s <a href="/pkg/builtin/#string">string</a>, mapping func(<a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Expand replaces ${var} or $var in the string based on the mapping function.
For example, os.ExpandEnv(s) is equivalent to os.Expand(s, os.Getenv).


<a id="example_Expand">Example</a>
```go
```

output:
```txt
```

## <a id="ExpandEnv">func</a> [ExpandEnv](https://golang.org/src/os/env.go?s=1321:1352#L40)
<pre>func ExpandEnv(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
ExpandEnv replaces ${var} or $var in the string according to the values
of the current environment variables. References to undefined
variables are replaced by the empty string.


<a id="example_ExpandEnv">Example</a>
```go
```

output:
```txt
```

## <a id="Getegid">func</a> [Getegid](https://golang.org/src/os/proc.go?s=999:1017#L35)
<pre>func Getegid() <a href="/pkg/builtin/#int">int</a></pre>
Getegid returns the numeric effective group id of the caller.

On Windows, it returns -1.



## <a id="Getenv">func</a> [Getenv](https://golang.org/src/os/env.go?s=2860:2890#L91)
<pre>func Getenv(key <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Getenv retrieves the value of the environment variable named by the key.
It returns the value, which will be empty if the variable is not present.
To distinguish between an empty value and an unset value, use LookupEnv.


<a id="example_Getenv">Example</a>
```go
```

output:
```txt
```

## <a id="Geteuid">func</a> [Geteuid](https://golang.org/src/os/proc.go?s=718:736#L25)
<pre>func Geteuid() <a href="/pkg/builtin/#int">int</a></pre>
Geteuid returns the numeric effective user id of the caller.

On Windows, it returns -1.



## <a id="Getgid">func</a> [Getgid](https://golang.org/src/os/proc.go?s=854:871#L30)
<pre>func Getgid() <a href="/pkg/builtin/#int">int</a></pre>
Getgid returns the numeric group id of the caller.

On Windows, it returns -1.



## <a id="Getgroups">func</a> [Getgroups](https://golang.org/src/os/proc.go?s=1235:1266#L41)
<pre>func Getgroups() ([]<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Getgroups returns a list of the numeric ids of groups that the caller belongs to.

On Windows, it returns syscall.EWINDOWS. See the os/user package
for a possible alternative.



## <a id="Getpagesize">func</a> [Getpagesize](https://golang.org/src/os/types.go?s=268:290#L3)
<pre>func Getpagesize() <a href="/pkg/builtin/#int">int</a></pre>
Getpagesize returns the underlying system's memory page size.



## <a id="Getpid">func</a> [Getpid](https://golang.org/src/os/exec.go?s=2156:2173#L61)
<pre>func Getpid() <a href="/pkg/builtin/#int">int</a></pre>
Getpid returns the process id of the caller.



## <a id="Getppid">func</a> [Getppid](https://golang.org/src/os/exec.go?s=2261:2279#L64)
<pre>func Getppid() <a href="/pkg/builtin/#int">int</a></pre>
Getppid returns the process id of the caller's parent.



## <a id="Getuid">func</a> [Getuid](https://golang.org/src/os/proc.go?s=574:591#L20)
<pre>func Getuid() <a href="/pkg/builtin/#int">int</a></pre>
Getuid returns the numeric user id of the caller.

On Windows, it returns -1.



## <a id="Getwd">func</a> [Getwd](https://golang.org/src/os/getwd.go?s=620:656#L16)
<pre>func Getwd() (dir <a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Getwd returns a rooted path name corresponding to the
current directory. If the current directory can be
reached via multiple paths (due to symbolic links),
Getwd may return any one of them.



## <a id="Hostname">func</a> [Hostname](https://golang.org/src/os/sys.go?s=230:270#L1)
<pre>func Hostname() (name <a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Hostname returns the host name reported by the kernel.



## <a id="IsExist">func</a> [IsExist](https://golang.org/src/os/error.go?s=2582:2610#L75)
<pre>func IsExist(err <a href="/pkg/builtin/#error">error</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsExist returns a boolean indicating whether the error is known to report
that a file or directory already exists. It is satisfied by ErrExist as
well as some syscall errors.



## <a id="IsNotExist">func</a> [IsNotExist](https://golang.org/src/os/error.go?s=2847:2878#L82)
<pre>func IsNotExist(err <a href="/pkg/builtin/#error">error</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsNotExist returns a boolean indicating whether the error is known to
report that a file or directory does not exist. It is satisfied by
ErrNotExist as well as some syscall errors.


<a id="example_IsNotExist">Example</a>
```go
```

output:
```txt
```

## <a id="IsPathSeparator">func</a> [IsPathSeparator](https://golang.org/src/os/path_unix.go?s=453:487#L5)
<pre>func IsPathSeparator(c <a href="/pkg/builtin/#uint8">uint8</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsPathSeparator reports whether c is a directory separator character.



## <a id="IsPermission">func</a> [IsPermission](https://golang.org/src/os/error.go?s=3108:3141#L89)
<pre>func IsPermission(err <a href="/pkg/builtin/#error">error</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsPermission returns a boolean indicating whether the error is known to
report that permission is denied. It is satisfied by ErrPermission as well
as some syscall errors.



## <a id="IsTimeout">func</a> [IsTimeout](https://golang.org/src/os/error.go?s=3300:3330#L95)
<pre>func IsTimeout(err <a href="/pkg/builtin/#error">error</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsTimeout returns a boolean indicating whether the error is known
to report that a timeout occurred.



## <a id="Lchown">func</a> [Lchown](https://golang.org/src/os/file_posix.go?s=1967:2011#L61)
<pre>func Lchown(name <a href="/pkg/builtin/#string">string</a>, uid, gid <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#error">error</a></pre>
Lchown changes the numeric uid and gid of the named file.
If the file is a symbolic link, it changes the uid and gid of the link itself.
If there is an error, it will be of type *PathError.

On Windows, it always returns the syscall.EWINDOWS error, wrapped
in *PathError.



## <a id="Link">func</a> [Link](https://golang.org/src/os/file_unix.go?s=10546:10586#L344)
<pre>func Link(oldname, newname <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Link creates newname as a hard link to the oldname file.
If there is an error, it will be of type *LinkError.



## <a id="LookupEnv">func</a> [LookupEnv](https://golang.org/src/os/env.go?s=3235:3276#L102)
<pre>func LookupEnv(key <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#bool">bool</a>)</pre>
LookupEnv retrieves the value of the environment variable named
by the key. If the variable is present in the environment the
value (which may be empty) is returned and the boolean is true.
Otherwise the returned value will be empty and the boolean will
be false.


<a id="example_LookupEnv">Example</a>
```go
```

output:
```txt
```

## <a id="Mkdir">func</a> [Mkdir](https://golang.org/src/os/file.go?s=6935:6979#L220)
<pre>func Mkdir(name <a href="/pkg/builtin/#string">string</a>, perm <a href="#FileMode">FileMode</a>) <a href="/pkg/builtin/#error">error</a></pre>
Mkdir creates a new directory with the specified name and permission
bits (before umask).
If there is an error, it will be of type *PathError.



## <a id="MkdirAll">func</a> [MkdirAll](https://golang.org/src/os/path.go?s=497:544#L8)
<pre>func MkdirAll(path <a href="/pkg/builtin/#string">string</a>, perm <a href="#FileMode">FileMode</a>) <a href="/pkg/builtin/#error">error</a></pre>
MkdirAll creates a directory named path,
along with any necessary parents, and returns nil,
or else returns an error.
The permission bits perm (before umask) are used for all
directories that MkdirAll creates.
If path is already a directory, MkdirAll does nothing
and returns nil.



## <a id="NewSyscallError">func</a> [NewSyscallError](https://golang.org/src/os/error.go?s=2270:2323#L65)
<pre>func NewSyscallError(syscall <a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>) <a href="/pkg/builtin/#error">error</a></pre>
NewSyscallError returns, as an error, a new SyscallError
with the given system call name and error details.
As a convenience, if err is nil, NewSyscallError returns nil.



## <a id="Pipe">func</a> [Pipe](https://golang.org/src/os/pipe_linux.go?s=319:360#L1)
<pre>func Pipe() (r *<a href="#File">File</a>, w *<a href="#File">File</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Pipe returns a connected pair of Files; reads from r return bytes written to w.
It returns the files and an error, if any.



## <a id="Readlink">func</a> [Readlink](https://golang.org/src/os/file_unix.go?s=11748:11790#L391)
<pre>func Readlink(name <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Readlink returns the destination of the named symbolic link.
If there is an error, it will be of type *PathError.



## <a id="Remove">func</a> [Remove](https://golang.org/src/os/file_unix.go?s=9421:9451#L301)
<pre>func Remove(name <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Remove removes the named file or (empty) directory.
If there is an error, it will be of type *PathError.



## <a id="RemoveAll">func</a> [RemoveAll](https://golang.org/src/os/path.go?s=1823:1856#L59)
<pre>func RemoveAll(path <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
RemoveAll removes path and any children it contains.
It removes everything it can but returns the first error
it encounters. If the path does not exist, RemoveAll
returns nil (no error).
If there is an error, it will be of type *PathError.



## <a id="Rename">func</a> [Rename](https://golang.org/src/os/file.go?s=9626:9668#L306)
<pre>func Rename(oldpath, newpath <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Rename renames (moves) oldpath to newpath.
If newpath already exists and is not a directory, Rename replaces it.
OS-specific restrictions may apply when oldpath and newpath are in different directories.
If there is an error, it will be of type *LinkError.



## <a id="SameFile">func</a> [SameFile](https://golang.org/src/os/types.go?s=4113:4150#L107)
<pre>func SameFile(fi1, fi2 <a href="#FileInfo">FileInfo</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
SameFile reports whether fi1 and fi2 describe the same file.
For example, on Unix this means that the device and inode fields
of the two underlying structures are identical; on other systems
the decision may be based on the path names.
SameFile only applies to results returned by this package's Stat.
It returns false in other cases.



## <a id="Setenv">func</a> [Setenv](https://golang.org/src/os/env.go?s=3434:3470#L109)
<pre>func Setenv(key, value <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Setenv sets the value of the environment variable named by the key.
It returns an error, if any.



## <a id="Symlink">func</a> [Symlink](https://golang.org/src/os/file_unix.go?s=10822:10865#L354)
<pre>func Symlink(oldname, newname <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Symlink creates newname as a symbolic link to oldname.
If there is an error, it will be of type *LinkError.



## <a id="TempDir">func</a> [TempDir](https://golang.org/src/os/file.go?s=10724:10745#L341)
<pre>func TempDir() <a href="/pkg/builtin/#string">string</a></pre>
TempDir returns the default directory to use for temporary files.

On Unix systems, it returns $TMPDIR if non-empty, else /tmp.
On Windows, it uses GetTempPath, returning the first non-empty
value from %TMP%, %TEMP%, %USERPROFILE%, or the Windows directory.
On Plan 9, it returns /tmp.

The directory is neither guaranteed to exist nor have accessible
permissions.



## <a id="Truncate">func</a> [Truncate](https://golang.org/src/os/file_unix.go?s=9154:9198#L292)
<pre>func Truncate(name <a href="/pkg/builtin/#string">string</a>, size <a href="/pkg/builtin/#int64">int64</a>) <a href="/pkg/builtin/#error">error</a></pre>
Truncate changes the size of the named file.
If the file is a symbolic link, it changes the size of the link's target.
If there is an error, it will be of type *PathError.



## <a id="Unsetenv">func</a> [Unsetenv](https://golang.org/src/os/env.go?s=3633:3664#L118)
<pre>func Unsetenv(key <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Unsetenv unsets a single environment variable.


<a id="example_Unsetenv">Example</a>
```go
```

output:
```txt
```

## <a id="UserCacheDir">func</a> [UserCacheDir](https://golang.org/src/os/file.go?s=11379:11414#L358)
<pre>func UserCacheDir() (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
UserCacheDir returns the default root directory to use for user-specific
cached data. Users should create their own application-specific subdirectory
within this one and use that.

On Unix systems, it returns $XDG_CACHE_HOME as specified by
<a href="https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html">https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html</a> if
non-empty, else $HOME/.cache.
On Darwin, it returns $HOME/Library/Caches.
On Windows, it returns %LocalAppData%.
On Plan 9, it returns $home/lib/cache.

If the location cannot be determined (for example, $HOME is not defined),
then it will return an error.



## <a id="UserConfigDir">func</a> [UserConfigDir](https://golang.org/src/os/file.go?s=12713:12749#L409)
<pre>func UserConfigDir() (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
UserConfigDir returns the default root directory to use for user-specific
configuration data. Users should create their own application-specific
subdirectory within this one and use that.

On Unix systems, it returns $XDG_CONFIG_HOME as specified by
<a href="https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html">https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html</a> if
non-empty, else $HOME/.config.
On Darwin, it returns $HOME/Library/Application Support.
On Windows, it returns %AppData%.
On Plan 9, it returns $home/lib.

If the location cannot be determined (for example, $HOME is not defined),
then it will return an error.



## <a id="UserHomeDir">func</a> [UserHomeDir](https://golang.org/src/os/file.go?s=13657:13691#L452)
<pre>func UserHomeDir() (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
UserHomeDir returns the current user's home directory.

On Unix, including macOS, it returns the $HOME environment variable.
On Windows, it returns %USERPROFILE%.
On Plan 9, it returns the $home environment variable.





## <a id="File">type</a> [File](https://golang.org/src/os/types.go?s=369:411#L6)
File represents an open file descriptor.


<pre>type File struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="Create">func</a> [Create](https://golang.org/src/os/file.go?s=8597:8636#L278)
<pre>func Create(name <a href="/pkg/builtin/#string">string</a>) (*<a href="#File">File</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Create creates or truncates the named file. If the file already exists,
it is truncated. If the file does not exist, it is created with mode 0666
(before umask). If successful, methods on the returned File can
be used for I/O; the associated file descriptor has mode O_RDWR.
If there is an error, it will be of type *PathError.




### <a id="NewFile">func</a> [NewFile](https://golang.org/src/os/file_unix.go?s=2693:2736#L73)
<pre>func NewFile(fd <a href="/pkg/builtin/#uintptr">uintptr</a>, name <a href="/pkg/builtin/#string">string</a>) *<a href="#File">File</a></pre>
NewFile returns a new File with the given file descriptor and
name. The returned value will be nil if fd is not a valid file
descriptor. On Unix systems, if the file descriptor is in
non-blocking mode, NewFile will attempt to return a pollable File
(one for which the SetDeadline methods work).




### <a id="Open">func</a> [Open](https://golang.org/src/os/file.go?s=8175:8212#L269)
<pre>func Open(name <a href="/pkg/builtin/#string">string</a>) (*<a href="#File">File</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Open opens the named file for reading. If successful, methods on
the returned file can be used for reading; the associated file
descriptor has mode O_RDONLY.
If there is an error, it will be of type *PathError.




### <a id="OpenFile">func</a> [OpenFile](https://golang.org/src/os/file.go?s=9082:9148#L288)
<pre>func OpenFile(name <a href="/pkg/builtin/#string">string</a>, flag <a href="/pkg/builtin/#int">int</a>, perm <a href="#FileMode">FileMode</a>) (*<a href="#File">File</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
OpenFile is the generalized open call; most users will use Open
or Create instead. It opens the named file with specified flag
(O_RDONLY etc.). If the file does not exist, and the O_CREATE flag
is passed, it is created with mode perm (before umask). If successful,
methods on the returned File can be used for I/O.
If there is an error, it will be of type *PathError.


<a id="example_OpenFile">Example</a>
```go
```

output:
```txt
```
<a id="example_OpenFile_append">Example (Append)</a>
```go
```

output:
```txt
```




### <a id="File.Chdir">func</a> (\*File) [Chdir](https://golang.org/src/os/file_posix.go?s=3978:4006#L128)
<pre>func (f *<a href="#File">File</a>) Chdir() <a href="/pkg/builtin/#error">error</a></pre>
Chdir changes the current working directory to the file,
which must be a directory.
If there is an error, it will be of type *PathError.




### <a id="File.Chmod">func</a> (\*File) [Chmod](https://golang.org/src/os/file.go?s=15206:15247#L499)
<pre>func (f *<a href="#File">File</a>) Chmod(mode <a href="#FileMode">FileMode</a>) <a href="/pkg/builtin/#error">error</a></pre>
Chmod changes the mode of the file to mode.
If there is an error, it will be of type *PathError.




### <a id="File.Chown">func</a> (\*File) [Chown](https://golang.org/src/os/file_posix.go?s=2329:2369#L73)
<pre>func (f *<a href="#File">File</a>) Chown(uid, gid <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#error">error</a></pre>
Chown changes the numeric uid and gid of the named file.
If there is an error, it will be of type *PathError.

On Windows, it always returns the syscall.EWINDOWS error, wrapped
in *PathError.




### <a id="File.Close">func</a> (\*File) [Close](https://golang.org/src/os/file_unix.go?s=7043:7071#L219)
<pre>func (f *<a href="#File">File</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the File, rendering it unusable for I/O.
On files that support SetDeadline, any pending I/O operations will
be canceled and return immediately with an error.
Close will return an error if it has already been called.




### <a id="File.Fd">func</a> (\*File) [Fd](https://golang.org/src/os/file_unix.go?s=1949:1976#L51)
<pre>func (f *<a href="#File">File</a>) Fd() <a href="/pkg/builtin/#uintptr">uintptr</a></pre>
Fd returns the integer Unix file descriptor referencing the open file.
The file descriptor is valid only until f.Close is called or f is garbage collected.
On Unix systems this will cause the SetDeadline methods to stop working.




### <a id="File.Name">func</a> (\*File) [Name](https://golang.org/src/os/file.go?s=1800:1828#L44)
<pre>func (f *<a href="#File">File</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the name of the file as presented to Open.




### <a id="File.Read">func</a> (\*File) [Read](https://golang.org/src/os/file.go?s=3908:3956#L102)
<pre>func (f *<a href="#File">File</a>) Read(b []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Read reads up to len(b) bytes from the File.
It returns the number of bytes read and any error encountered.
At end of file, Read returns 0, io.EOF.




### <a id="File.ReadAt">func</a> (\*File) [ReadAt](https://golang.org/src/os/file.go?s=4311:4372#L114)
<pre>func (f *<a href="#File">File</a>) ReadAt(b []<a href="/pkg/builtin/#byte">byte</a>, off <a href="/pkg/builtin/#int64">int64</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadAt reads len(b) bytes from the File starting at byte offset off.
It returns the number of bytes read and the error, if any.
ReadAt always returns a non-nil error when n < len(b).
At end of file, that error is io.EOF.




### <a id="File.Readdir">func</a> (\*File) [Readdir](https://golang.org/src/os/dir.go?s=980:1029#L12)
<pre>func (f *<a href="#File">File</a>) Readdir(n <a href="/pkg/builtin/#int">int</a>) ([]<a href="#FileInfo">FileInfo</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Readdir reads the contents of the directory associated with file and
returns a slice of up to n FileInfo values, as would be returned
by Lstat, in directory order. Subsequent calls on the same file will yield
further FileInfos.

If n > 0, Readdir returns at most n FileInfo structures. In this case, if
Readdir returns an empty slice, it will return a non-nil error
explaining why. At the end of a directory, the error is io.EOF.

If n <= 0, Readdir returns all the FileInfo from the directory in
a single slice. In this case, if Readdir succeeds (reads all
the way to the end of the directory), it returns the slice and a
nil error. If it encounters an error before the end of the
directory, Readdir returns the FileInfo read until that point
and a non-nil error.




### <a id="File.Readdirnames">func</a> (\*File) [Readdirnames](https://golang.org/src/os/dir.go?s=1898:1960#L34)
<pre>func (f *<a href="#File">File</a>) Readdirnames(n <a href="/pkg/builtin/#int">int</a>) (names []<a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Readdirnames reads the contents of the directory associated with file
and returns a slice of up to n names of files in the directory,
in directory order. Subsequent calls on the same file will yield
further names.

If n > 0, Readdirnames returns at most n names. In this case, if
Readdirnames returns an empty slice, it will return a non-nil error
explaining why. At the end of a directory, the error is io.EOF.

If n <= 0, Readdirnames returns all the names from the directory in
a single slice. In this case, if Readdirnames succeeds (reads all
the way to the end of the directory), it returns the slice and a
nil error. If it encounters an error before the end of the
directory, Readdirnames returns the names read until that point and
a non-nil error.




### <a id="File.Seek">func</a> (\*File) [Seek](https://golang.org/src/os/file.go?s=6290:6358#L197)
<pre>func (f *<a href="#File">File</a>) Seek(offset <a href="/pkg/builtin/#int64">int64</a>, whence <a href="/pkg/builtin/#int">int</a>) (ret <a href="/pkg/builtin/#int64">int64</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Seek sets the offset for the next Read or Write on file to offset, interpreted
according to whence: 0 means relative to the origin of the file, 1 means
relative to the current offset, and 2 means relative to the end.
It returns the new offset and an error, if any.
The behavior of Seek on a file opened with O_APPEND is not specified.




### <a id="File.SetDeadline">func</a> (\*File) [SetDeadline](https://golang.org/src/os/file.go?s=16427:16472#L523)
<pre>func (f *<a href="#File">File</a>) SetDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetDeadline sets the read and write deadlines for a File.
It is equivalent to calling both SetReadDeadline and SetWriteDeadline.

Only some kinds of files support setting a deadline. Calls to SetDeadline
for files that do not support deadlines will return ErrNoDeadline.
On most systems ordinary files do not support deadlines, but pipes do.

A deadline is an absolute time after which I/O operations fail with an
error instead of blocking. The deadline applies to all future and pending
I/O, not just the immediately following call to Read or Write.
After a deadline has been exceeded, the connection can be refreshed
by setting a deadline in the future.

An error returned after a timeout fails will implement the
Timeout method, and calling the Timeout method will return true.
The PathError and SyscallError types implement the Timeout method.
In general, call IsTimeout to test whether an error indicates a timeout.

An idle timeout can be implemented by repeatedly extending
the deadline after successful Read or Write calls.

A zero value for t means I/O operations will not time out.




### <a id="File.SetReadDeadline">func</a> (\*File) [SetReadDeadline](https://golang.org/src/os/file.go?s=16715:16764#L531)
<pre>func (f *<a href="#File">File</a>) SetReadDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetReadDeadline sets the deadline for future Read calls and any
currently-blocked Read call.
A zero value for t means Read will not time out.
Not all files support setting deadlines; see SetDeadline.




### <a id="File.SetWriteDeadline">func</a> (\*File) [SetWriteDeadline](https://golang.org/src/os/file.go?s=17130:17180#L541)
<pre>func (f *<a href="#File">File</a>) SetWriteDeadline(t <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) <a href="/pkg/builtin/#error">error</a></pre>
SetWriteDeadline sets the deadline for any future Write calls and any
currently-blocked Write call.
Even if Write times out, it may return n > 0, indicating that
some of the data was successfully written.
A zero value for t means Write will not time out.
Not all files support setting deadlines; see SetDeadline.




### <a id="File.Stat">func</a> (\*File) [Stat](https://golang.org/src/os/stat_unix.go?s=389:428#L5)
<pre>func (f *<a href="#File">File</a>) Stat() (<a href="#FileInfo">FileInfo</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Stat returns the FileInfo structure describing file.
If there is an error, it will be of type *PathError.




### <a id="File.Sync">func</a> (\*File) [Sync](https://golang.org/src/os/file_posix.go?s=3041:3068#L99)
<pre>func (f *<a href="#File">File</a>) Sync() <a href="/pkg/builtin/#error">error</a></pre>
Sync commits the current contents of the file to stable storage.
Typically, this means flushing the file system's in-memory copy
of recently written data to disk.




### <a id="File.SyscallConn">func</a> (\*File) [SyscallConn](https://golang.org/src/os/file.go?s=17298:17351#L547)
<pre>func (f *<a href="#File">File</a>) SyscallConn() (<a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#RawConn">RawConn</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
SyscallConn returns a raw file.
This implements the syscall.Conn interface.




### <a id="File.Truncate">func</a> (\*File) [Truncate](https://golang.org/src/os/file_posix.go?s=2664:2705#L86)
<pre>func (f *<a href="#File">File</a>) Truncate(size <a href="/pkg/builtin/#int64">int64</a>) <a href="/pkg/builtin/#error">error</a></pre>
Truncate changes the size of the file.
It does not change the I/O offset.
If there is an error, it will be of type *PathError.




### <a id="File.Write">func</a> (\*File) [Write](https://golang.org/src/os/file.go?s=4844:4893#L139)
<pre>func (f *<a href="#File">File</a>) Write(b []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Write writes len(b) bytes to the File.
It returns the number of bytes written and an error, if any.
Write returns a non-nil error when n != len(b).




### <a id="File.WriteAt">func</a> (\*File) [WriteAt](https://golang.org/src/os/file.go?s=5502:5564#L167)
<pre>func (f *<a href="#File">File</a>) WriteAt(b []<a href="/pkg/builtin/#byte">byte</a>, off <a href="/pkg/builtin/#int64">int64</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
WriteAt writes len(b) bytes to the File starting at byte offset off.
It returns the number of bytes written and an error, if any.
WriteAt returns a non-nil error when n != len(b).

If file was opened with the O_APPEND flag, WriteAt returns an error.




### <a id="File.WriteString">func</a> (\*File) [WriteString](https://golang.org/src/os/file.go?s=6695:6750#L213)
<pre>func (f *<a href="#File">File</a>) WriteString(s <a href="/pkg/builtin/#string">string</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
WriteString is like Write, but writes the contents of string s rather than
a slice of bytes.




## <a id="FileInfo">type</a> [FileInfo](https://golang.org/src/os/types.go?s=479:840#L11)
A FileInfo describes a file and is returned by Stat and Lstat.


<pre>type FileInfo interface {
    Name() <a href="/pkg/builtin/#string">string</a>       <span class="comment">// base name of the file</span>
    Size() <a href="/pkg/builtin/#int64">int64</a>        <span class="comment">// length in bytes for regular files; system-dependent for others</span>
    Mode() <a href="#FileMode">FileMode</a>     <span class="comment">// file mode bits</span>
    ModTime() <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a> <span class="comment">// modification time</span>
    IsDir() <a href="/pkg/builtin/#bool">bool</a>        <span class="comment">// abbreviation for Mode().IsDir()</span>
    Sys() interface{}   <span class="comment">// underlying data source (can return nil)</span>
}</pre>









### <a id="Lstat">func</a> [Lstat](https://golang.org/src/os/stat.go?s=642:683#L10)
<pre>func Lstat(name <a href="/pkg/builtin/#string">string</a>) (<a href="#FileInfo">FileInfo</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Lstat returns a FileInfo describing the named file.
If the file is a symbolic link, the returned FileInfo
describes the symbolic link. Lstat makes no attempt to follow the link.
If there is an error, it will be of type *PathError.




### <a id="Stat">func</a> [Stat](https://golang.org/src/os/stat.go?s=309:349#L1)
<pre>func Stat(name <a href="/pkg/builtin/#string">string</a>) (<a href="#FileInfo">FileInfo</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Stat returns a FileInfo describing the named file.
If there is an error, it will be of type *PathError.






## <a id="FileMode">type</a> [FileMode](https://golang.org/src/os/types.go?s=1131:1151#L25)
A FileMode represents a file's mode and permission bits.
The bits have the same definition on all systems, so that
information about files can be moved from one system
to another portably. Not all bits apply to all systems.
The only required bit is ModeDir for directories.


<pre>type FileMode <a href="/pkg/builtin/#uint32">uint32</a></pre>


The defined file mode bits are the most significant bits of the FileMode.
The nine least-significant bits are the standard Unix rwxrwxrwx permissions.
The values of these bits should be considered part of the public API and
may be used in wire protocols or disk representations: they must not be
changed, although new bits might be added.


<pre>const (
    <span class="comment">// The single letters are the abbreviations</span>
    <span class="comment">// used by the String method&#39;s formatting.</span>
    <span id="ModeDir">ModeDir</span>        <a href="#FileMode">FileMode</a> = 1 &lt;&lt; (32 - 1 - <a href="/pkg/builtin/#iota">iota</a>) <span class="comment">// d: is a directory</span>
    <span id="ModeAppend">ModeAppend</span>                                     <span class="comment">// a: append-only</span>
    <span id="ModeExclusive">ModeExclusive</span>                                  <span class="comment">// l: exclusive use</span>
    <span id="ModeTemporary">ModeTemporary</span>                                  <span class="comment">// T: temporary file; Plan 9 only</span>
    <span id="ModeSymlink">ModeSymlink</span>                                    <span class="comment">// L: symbolic link</span>
    <span id="ModeDevice">ModeDevice</span>                                     <span class="comment">// D: device file</span>
    <span id="ModeNamedPipe">ModeNamedPipe</span>                                  <span class="comment">// p: named pipe (FIFO)</span>
    <span id="ModeSocket">ModeSocket</span>                                     <span class="comment">// S: Unix domain socket</span>
    <span id="ModeSetuid">ModeSetuid</span>                                     <span class="comment">// u: setuid</span>
    <span id="ModeSetgid">ModeSetgid</span>                                     <span class="comment">// g: setgid</span>
    <span id="ModeCharDevice">ModeCharDevice</span>                                 <span class="comment">// c: Unix character device, when ModeDevice is set</span>
    <span id="ModeSticky">ModeSticky</span>                                     <span class="comment">// t: sticky</span>
    <span id="ModeIrregular">ModeIrregular</span>                                  <span class="comment">// ?: non-regular file; nothing else is known about this file</span>

    <span class="comment">// Mask for the type bits. For regular files, none will be set.</span>
    <span id="ModeType">ModeType</span> = <a href="#ModeDir">ModeDir</a> | <a href="#ModeSymlink">ModeSymlink</a> | <a href="#ModeNamedPipe">ModeNamedPipe</a> | <a href="#ModeSocket">ModeSocket</a> | <a href="#ModeDevice">ModeDevice</a> | <a href="#ModeCharDevice">ModeCharDevice</a> | <a href="#ModeIrregular">ModeIrregular</a>

    <span id="ModePerm">ModePerm</span> <a href="#FileMode">FileMode</a> = 0777 <span class="comment">// Unix permission bits</span>
)</pre>



<a id="example_FileMode">Example</a>
```go
```

output:
```txt
```






### <a id="FileMode.IsDir">func</a> (FileMode) [IsDir](https://golang.org/src/os/types.go?s=3303:3333#L83)
<pre>func (m <a href="#FileMode">FileMode</a>) IsDir() <a href="/pkg/builtin/#bool">bool</a></pre>
IsDir reports whether m describes a directory.
That is, it tests for the ModeDir bit being set in m.




### <a id="FileMode.IsRegular">func</a> (FileMode) [IsRegular](https://golang.org/src/os/types.go?s=3472:3506#L89)
<pre>func (m <a href="#FileMode">FileMode</a>) IsRegular() <a href="/pkg/builtin/#bool">bool</a></pre>
IsRegular reports whether m describes a regular file.
That is, it tests that no mode type bits are set.




### <a id="FileMode.Perm">func</a> (FileMode) [Perm](https://golang.org/src/os/types.go?s=3583:3616#L94)
<pre>func (m <a href="#FileMode">FileMode</a>) Perm() <a href="#FileMode">FileMode</a></pre>
Perm returns the Unix permission bits in m.




### <a id="FileMode.String">func</a> (FileMode) [String](https://golang.org/src/os/types.go?s=2790:2823#L55)
<pre>func (m <a href="#FileMode">FileMode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="LinkError">type</a> [LinkError](https://golang.org/src/os/file.go?s=3519:3591#L84)
LinkError records an error during a link or symlink or rename
system call and the paths that caused it.


<pre>type LinkError struct {
<span id="LinkError.Op"></span>    Op  <a href="/pkg/builtin/#string">string</a>
<span id="LinkError.Old"></span>    Old <a href="/pkg/builtin/#string">string</a>
<span id="LinkError.New"></span>    New <a href="/pkg/builtin/#string">string</a>
<span id="LinkError.Err"></span>    Err <a href="/pkg/builtin/#error">error</a>
}
</pre>











### <a id="LinkError.Error">func</a> (\*LinkError) [Error](https://golang.org/src/os/file.go?s=3593:3627#L91)
<pre>func (e *<a href="#LinkError">LinkError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="LinkError.Unwrap">func</a> (\*LinkError) [Unwrap](https://golang.org/src/os/file.go?s=3697:3731#L95)
<pre>func (e *<a href="#LinkError">LinkError</a>) Unwrap() <a href="/pkg/builtin/#error">error</a></pre>



## <a id="PathError">type</a> [PathError](https://golang.org/src/os/error.go?s=1314:1377#L30)
PathError records an error and the operation and file path that caused it.


<pre>type PathError struct {
<span id="PathError.Op"></span>    Op   <a href="/pkg/builtin/#string">string</a>
<span id="PathError.Path"></span>    Path <a href="/pkg/builtin/#string">string</a>
<span id="PathError.Err"></span>    Err  <a href="/pkg/builtin/#error">error</a>
}
</pre>











### <a id="PathError.Error">func</a> (\*PathError) [Error](https://golang.org/src/os/error.go?s=1379:1413#L36)
<pre>func (e *<a href="#PathError">PathError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="PathError.Timeout">func</a> (\*PathError) [Timeout](https://golang.org/src/os/error.go?s=1582:1616#L41)
<pre>func (e *<a href="#PathError">PathError</a>) Timeout() <a href="/pkg/builtin/#bool">bool</a></pre>
Timeout reports whether this error represents a timeout.




### <a id="PathError.Unwrap">func</a> (\*PathError) [Unwrap](https://golang.org/src/os/error.go?s=1469:1503#L38)
<pre>func (e *<a href="#PathError">PathError</a>) Unwrap() <a href="/pkg/builtin/#error">error</a></pre>



## <a id="ProcAttr">type</a> [ProcAttr](https://golang.org/src/os/exec.go?s=968:1859#L30)
ProcAttr holds the attributes that will be applied to a new process
started by StartProcess.


<pre>type ProcAttr struct {
    <span class="comment">// If Dir is non-empty, the child changes into the directory before</span>
    <span class="comment">// creating the process.</span>
<span id="ProcAttr.Dir"></span>    Dir <a href="/pkg/builtin/#string">string</a>
    <span class="comment">// If Env is non-nil, it gives the environment variables for the</span>
    <span class="comment">// new process in the form returned by Environ.</span>
    <span class="comment">// If it is nil, the result of Environ will be used.</span>
<span id="ProcAttr.Env"></span>    Env []<a href="/pkg/builtin/#string">string</a>
<span id="ProcAttr.Files"></span>    <span class="comment">// Files specifies the open files inherited by the new process. The</span>
    <span class="comment">// first three entries correspond to standard input, standard output, and</span>
    <span class="comment">// standard error. An implementation may support additional entries,</span>
    <span class="comment">// depending on the underlying operating system. A nil entry corresponds</span>
    <span class="comment">// to that file being closed when the process starts.</span>
    Files []*<a href="#File">File</a>

    <span class="comment">// Operating system-specific process creation attributes.</span>
    <span class="comment">// Note that setting this field means that your program</span>
    <span class="comment">// may not execute properly or even compile on some</span>
    <span class="comment">// operating systems.</span>
<span id="ProcAttr.Sys"></span>    Sys *<a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#SysProcAttr">SysProcAttr</a>
}
</pre>











## <a id="Process">type</a> [Process](https://golang.org/src/os/exec.go?s=332:573#L7)
Process stores the information about a process created by StartProcess.


<pre>type Process struct {
<span id="Process.Pid"></span>    Pid <a href="/pkg/builtin/#int">int</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="FindProcess">func</a> [FindProcess](https://golang.org/src/os/exec.go?s=2615:2658#L73)
<pre>func FindProcess(pid <a href="/pkg/builtin/#int">int</a>) (*<a href="#Process">Process</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
FindProcess looks for a running process by its pid.

The Process it returns can be used to obtain information
about the underlying operating system process.

On Unix systems, FindProcess always succeeds and returns a Process
for the given pid, regardless of whether the process exists.




### <a id="StartProcess">func</a> [StartProcess](https://golang.org/src/os/exec.go?s=3326:3405#L90)
<pre>func StartProcess(name <a href="/pkg/builtin/#string">string</a>, argv []<a href="/pkg/builtin/#string">string</a>, attr *<a href="#ProcAttr">ProcAttr</a>) (*<a href="#Process">Process</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
StartProcess starts a new process with the program, arguments and attributes
specified by name, argv and attr. The argv slice will become os.Args in the
new process, so it normally starts with the program name.

If the calling goroutine has locked the operating system thread
with runtime.LockOSThread and modified any inheritable OS-level
thread state (for example, Linux or Plan 9 name spaces), the new
process will inherit the caller's thread state.

StartProcess is a low-level interface. The os/exec package provides
higher-level interfaces.

If there is an error, it will be of type *PathError.






### <a id="Process.Kill">func</a> (\*Process) [Kill](https://golang.org/src/os/exec.go?s=3878:3908#L105)
<pre>func (p *<a href="#Process">Process</a>) Kill() <a href="/pkg/builtin/#error">error</a></pre>
Kill causes the Process to exit immediately. Kill does not wait until
the Process has actually exited. This only kills the Process itself,
not any other processes it may have started.




### <a id="Process.Release">func</a> (\*Process) [Release](https://golang.org/src/os/exec.go?s=3626:3659#L98)
<pre>func (p *<a href="#Process">Process</a>) Release() <a href="/pkg/builtin/#error">error</a></pre>
Release releases any resources associated with the Process p,
rendering it unusable in the future.
Release only needs to be called if Wait is not.




### <a id="Process.Signal">func</a> (\*Process) [Signal](https://golang.org/src/os/exec.go?s=4386:4428#L120)
<pre>func (p *<a href="#Process">Process</a>) Signal(sig <a href="#Signal">Signal</a>) <a href="/pkg/builtin/#error">error</a></pre>
Signal sends a signal to the Process.
Sending Interrupt on Windows is not implemented.




### <a id="Process.Wait">func</a> (\*Process) [Wait](https://golang.org/src/os/exec.go?s=4223:4270#L114)
<pre>func (p *<a href="#Process">Process</a>) Wait() (*<a href="#ProcessState">ProcessState</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Wait waits for the Process to exit, and then returns a
ProcessState describing its status and an error, if any.
Wait releases any resources associated with the Process.
On most operating systems, the Process must be a child
of the current process or an error will be returned.




## <a id="ProcessState">type</a> [ProcessState](https://golang.org/src/os/exec_posix.go?s=1764:1924#L53)
ProcessState stores information about a process, as reported by Wait.


<pre>type ProcessState struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="ProcessState.ExitCode">func</a> (\*ProcessState) [ExitCode](https://golang.org/src/os/exec_posix.go?s=3042:3079#L107)
<pre>func (p *<a href="#ProcessState">ProcessState</a>) ExitCode() <a href="/pkg/builtin/#int">int</a></pre>
ExitCode returns the exit code of the exited process, or -1
if the process hasn't exited or was terminated by a signal.




### <a id="ProcessState.Exited">func</a> (\*ProcessState) [Exited](https://golang.org/src/os/exec.go?s=4818:4854#L135)
<pre>func (p *<a href="#ProcessState">ProcessState</a>) Exited() <a href="/pkg/builtin/#bool">bool</a></pre>
Exited reports whether the program has exited.




### <a id="ProcessState.Pid">func</a> (\*ProcessState) [Pid](https://golang.org/src/os/exec_posix.go?s=1979:2011#L60)
<pre>func (p *<a href="#ProcessState">ProcessState</a>) Pid() <a href="/pkg/builtin/#int">int</a></pre>
Pid returns the process id of the exited process.




### <a id="ProcessState.String">func</a> (\*ProcessState) [String](https://golang.org/src/os/exec_posix.go?s=2308:2346#L80)
<pre>func (p *<a href="#ProcessState">ProcessState</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="ProcessState.Success">func</a> (\*ProcessState) [Success](https://golang.org/src/os/exec.go?s=4978:5015#L141)
<pre>func (p *<a href="#ProcessState">ProcessState</a>) Success() <a href="/pkg/builtin/#bool">bool</a></pre>
Success reports whether the program exited successfully,
such as with exit status 0 on Unix.




### <a id="ProcessState.Sys">func</a> (\*ProcessState) [Sys](https://golang.org/src/os/exec.go?s=5222:5262#L148)
<pre>func (p *<a href="#ProcessState">ProcessState</a>) Sys() interface{}</pre>
Sys returns system-dependent exit information about
the process. Convert it to the appropriate underlying
type, such as syscall.WaitStatus on Unix, to access its contents.




### <a id="ProcessState.SysUsage">func</a> (\*ProcessState) [SysUsage](https://golang.org/src/os/exec.go?s=5583:5628#L157)
<pre>func (p *<a href="#ProcessState">ProcessState</a>) SysUsage() interface{}</pre>
SysUsage returns system-dependent resource usage information about
the exited process. Convert it to the appropriate underlying
type, such as *syscall.Rusage on Unix, to access its contents.
(On Unix, *syscall.Rusage matches struct rusage as defined in the
getrusage(2) manual page.)




### <a id="ProcessState.SystemTime">func</a> (\*ProcessState) [SystemTime](https://golang.org/src/os/exec.go?s=4690:4739#L130)
<pre>func (p *<a href="#ProcessState">ProcessState</a>) SystemTime() <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a></pre>
SystemTime returns the system CPU time of the exited process and its children.




### <a id="ProcessState.UserTime">func</a> (\*ProcessState) [UserTime](https://golang.org/src/os/exec.go?s=4534:4581#L125)
<pre>func (p *<a href="#ProcessState">ProcessState</a>) UserTime() <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a></pre>
UserTime returns the user CPU time of the exited process and its children.




## <a id="Signal">type</a> [Signal](https://golang.org/src/os/exec.go?s=2015:2106#L55)
A Signal represents an operating system signal.
The usual underlying implementation is operating system-dependent:
on Unix it is syscall.Signal.


<pre>type Signal interface {
    String() <a href="/pkg/builtin/#string">string</a>
    Signal() <span class="comment">// to distinguish from other Stringers</span>
}</pre>




The only signal values guaranteed to be present in the os package on all
systems are os.Interrupt (send the process an interrupt) and os.Kill (force
the process to exit). On Windows, sending os.Interrupt to a process with
os.Process.Signal is not implemented; it will return an error instead of
sending a signal.


<pre>var (
    <span id="Interrupt">Interrupt</span> <a href="#Signal">Signal</a> = <a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#SIGINT">SIGINT</a>
    <span id="Kill">Kill</span>      <a href="#Signal">Signal</a> = <a href="/pkg/syscall/">syscall</a>.<a href="/pkg/syscall/#SIGKILL">SIGKILL</a>
)</pre>







## <a id="SyscallError">type</a> [SyscallError](https://golang.org/src/os/error.go?s=1736:1795#L47)
SyscallError records an error from a specific system call.


<pre>type SyscallError struct {
<span id="SyscallError.Syscall"></span>    Syscall <a href="/pkg/builtin/#string">string</a>
<span id="SyscallError.Err"></span>    Err     <a href="/pkg/builtin/#error">error</a>
}
</pre>











### <a id="SyscallError.Error">func</a> (\*SyscallError) [Error](https://golang.org/src/os/error.go?s=1797:1834#L52)
<pre>func (e *<a href="#SyscallError">SyscallError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="SyscallError.Timeout">func</a> (\*SyscallError) [Timeout](https://golang.org/src/os/error.go?s=1996:2033#L57)
<pre>func (e *<a href="#SyscallError">SyscallError</a>) Timeout() <a href="/pkg/builtin/#bool">bool</a></pre>
Timeout reports whether this error represents a timeout.




### <a id="SyscallError.Unwrap">func</a> (\*SyscallError) [Unwrap](https://golang.org/src/os/error.go?s=1880:1917#L54)
<pre>func (e *<a href="#SyscallError">SyscallError</a>) Unwrap() <a href="/pkg/builtin/#error">error</a></pre>






