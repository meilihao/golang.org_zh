

# ioutil
`import "io/ioutil"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package ioutil implements some I/O utility functions.




## <a id="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func NopCloser(r io.Reader) io.ReadCloser](#NopCloser)
* [func ReadAll(r io.Reader) ([]byte, error)](#ReadAll)
* [func ReadDir(dirname string) ([]os.FileInfo, error)](#ReadDir)
* [func ReadFile(filename string) ([]byte, error)](#ReadFile)
* [func TempDir(dir, prefix string) (name string, err error)](#TempDir)
* [func TempFile(dir, pattern string) (f *os.File, err error)](#TempFile)
* [func WriteFile(filename string, data []byte, perm os.FileMode) error](#WriteFile)


#### <a id="pkg-examples">Examples</a>
* [ReadAll](#example_ReadAll)
* [ReadDir](#example_ReadDir)
* [ReadFile](#example_ReadFile)
* [TempDir](#example_TempDir)
* [TempFile](#example_TempFile)
* [TempFile (Suffix)](#example_TempFile_suffix)
* [WriteFile](#example_WriteFile)


#### <a id="pkg-files">Package files</a>
[ioutil.go](https://golang.org/src/io/ioutil/ioutil.go) [tempfile.go](https://golang.org/src/io/ioutil/tempfile.go) 




## <a id="pkg-variables">Variables</a>
Discard is an io.Writer on which all Write calls succeed
without doing anything.


<pre>var <span id="Discard">Discard</span> <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a> = devNull(0)</pre>

## <a id="NopCloser">func</a> [NopCloser](https://golang.org/src/io/ioutil/ioutil.go?s=3458:3499#L108)
<pre>func NopCloser(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadCloser">ReadCloser</a></pre>
NopCloser returns a ReadCloser with a no-op Close method wrapping
the provided Reader r.



## <a id="ReadAll">func</a> [ReadAll](https://golang.org/src/io/ioutil/ioutil.go?s=1186:1227#L34)
<pre>func ReadAll(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadAll reads from r until an error or EOF and returns the data it read.
A successful call returns err == nil, not err == EOF. Because ReadAll is
defined to read from src until EOF, it does not treat an EOF from Read
as an error to be reported.


<a id="example_ReadAll">Example</a>

## <a id="ReadDir">func</a> [ReadDir](https://golang.org/src/io/ioutil/ioutil.go?s=2978:3029#L86)
<pre>func ReadDir(dirname <a href="/pkg/builtin/#string">string</a>) ([]<a href="/pkg/os/">os</a>.<a href="/pkg/os/#FileInfo">FileInfo</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadDir reads the directory named by dirname and returns
a list of directory entries sorted by filename.


<a id="example_ReadDir">Example</a>

## <a id="ReadFile">func</a> [ReadFile](https://golang.org/src/io/ioutil/ioutil.go?s=1503:1549#L42)
<pre>func ReadFile(filename <a href="/pkg/builtin/#string">string</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadFile reads the file named by filename and returns the contents.
A successful call returns err == nil, not err == EOF. Because ReadFile
reads the whole file, it does not treat an EOF from Read as an error
to be reported.


<a id="example_ReadFile">Example</a>

## <a id="TempDir">func</a> [TempDir](https://golang.org/src/io/ioutil/tempfile.go?s=2446:2503#L76)
<pre>func TempDir(dir, prefix <a href="/pkg/builtin/#string">string</a>) (name <a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
TempDir creates a new temporary directory in the directory dir
with a name beginning with prefix and returns the path of the
new directory. If dir is the empty string, TempDir uses the
default directory for temporary files (see os.TempDir).
Multiple programs calling TempDir simultaneously
will not choose the same directory. It is the caller's responsibility
to remove the directory when no longer needed.


<a id="example_TempDir">Example</a>

## <a id="TempFile">func</a> [TempFile](https://golang.org/src/io/ioutil/tempfile.go?s=1419:1477#L40)
<pre>func TempFile(dir, pattern <a href="/pkg/builtin/#string">string</a>) (f *<a href="/pkg/os/">os</a>.<a href="/pkg/os/#File">File</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
TempFile creates a new temporary file in the directory dir,
opens the file for reading and writing, and returns the resulting *os.File.
The filename is generated by taking pattern and adding a random
string to the end. If pattern includes a "*", the random string
replaces the last "*".
If dir is the empty string, TempFile uses the default directory
for temporary files (see os.TempDir).
Multiple programs calling TempFile simultaneously
will not choose the same file. The caller can use f.Name()
to find the pathname of the file. It is the caller's responsibility
to remove the file when no longer needed.


<a id="example_TempFile">Example</a>
<a id="example_TempFile_suffix">Example (Suffix)</a>

## <a id="WriteFile">func</a> [WriteFile](https://golang.org/src/io/ioutil/ioutil.go?s=2534:2602#L69)
<pre>func WriteFile(filename <a href="/pkg/builtin/#string">string</a>, data []<a href="/pkg/builtin/#byte">byte</a>, perm <a href="/pkg/os/">os</a>.<a href="/pkg/os/#FileMode">FileMode</a>) <a href="/pkg/builtin/#error">error</a></pre>
WriteFile writes data to a file named by filename.
If the file does not exist, WriteFile creates it with permissions perm;
otherwise WriteFile truncates it before writing.


<a id="example_WriteFile">Example</a>






