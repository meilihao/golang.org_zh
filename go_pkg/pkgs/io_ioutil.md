

# ioutil
`import "io/ioutil"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package ioutil implements some I/O utility functions.

ioutil 包实现了一些实用的 I/O 函数.


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

Discard 是一个不做任何操作且永远返回成功的 io.Writer 实现.

<pre>var <span id="Discard">Discard</span> <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a> = devNull(0)</pre>

## <a id="NopCloser">func</a> [NopCloser](https://golang.org/src/io/ioutil/ioutil.go?s=3458:3499#L108)
<pre>func NopCloser(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadCloser">ReadCloser</a></pre>
NopCloser returns a ReadCloser with a no-op Close method wrapping
the provided Reader r.

NopCloser 返回一个ReadCloser对象，它带有no-op Close方法，并封装了给定的Reader对象 r.

## <a id="ReadAll">func</a> [ReadAll](https://golang.org/src/io/ioutil/ioutil.go?s=1186:1227#L34)
<pre>func ReadAll(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadAll reads from r until an error or EOF and returns the data it read.
A successful call returns err == nil, not err == EOF. Because ReadAll is
defined to read from src until EOF, it does not treat an EOF from Read
as an error to be reported.

ReadAll 从r读取直到遇到error或EOF并返回已读取的数据. 如果调用成功，返回的err为nil, 而不是EOF. 因为ReadAll定义为从源读取数据直到EOF, 因此它不会把EOF视为报错.

<a id="example_ReadAll">Example</a>
```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

}
```

output:
```txt
Go is a general-purpose language designed with systems programming in mind.
```

## <a id="ReadDir">func</a> [ReadDir](https://golang.org/src/io/ioutil/ioutil.go?s=2978:3029#L86)
<pre>func ReadDir(dirname <a href="/pkg/builtin/#string">string</a>) ([]<a href="/pkg/os/">os</a>.<a href="/pkg/os/#FileInfo">FileInfo</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadDir reads the directory named by dirname and returns
a list of directory entries sorted by filename.

ReadDir 根据指定的目录dirname，并返回已排序(根据文件名)的文件列表.

<a id="example_ReadDir">Example</a>
```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
```

output:
```txt
dev
etc
tmp
usr
```
## <a id="ReadFile">func</a> [ReadFile](https://golang.org/src/io/ioutil/ioutil.go?s=1503:1549#L42)
<pre>func ReadFile(filename <a href="/pkg/builtin/#string">string</a>) ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadFile reads the file named by filename and returns the contents.
A successful call returns err == nil, not err == EOF. Because ReadFile
reads the whole file, it does not treat an EOF from Read as an error
to be reported.

ReadFile 会从指定的文件filename中读取数据并返回文件的内容. 若调用成功, 返回的err为nil， 而不是EOF. 因为ReadFile定义为从源读取数据直到EOF, 它不会把EOF视为报错.

<a id="example_ReadFile">Example</a>
```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("testdata/hello")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)

}
```

output:
```txt
2009/11/10 23:00:00 open testdata/hello: No such file or directory
```
## <a id="TempDir">func</a> [TempDir](https://golang.org/src/io/ioutil/tempfile.go?s=2446:2503#L76)
<pre>func TempDir(dir, prefix <a href="/pkg/builtin/#string">string</a>) (name <a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
TempDir creates a new temporary directory in the directory dir
with a name beginning with prefix and returns the path of the
new directory. If dir is the empty string, TempDir uses the
default directory for temporary files (see os.TempDir).
Multiple programs calling TempDir simultaneously
will not choose the same directory. It is the caller's responsibility
to remove the directory when no longer needed.

TempDir 会在指定文件夹dir里创建一个新的,使用prfix作为前缀的临时文件夹，并返回该文件夹的路径. 如果dir是空字符串，TempDir使用默认的临时文件夹(参考os.TempDir方法). 如果多个程序同时调用该函数的话，它不会选中相同的文件夹(即并发安全). 调用者有责任在不需要时删除它.

<a id="example_TempDir">Example</a>
```go
package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}
}
```

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


TempFile 会在dir目录下创建一个新的,使用prefix为前缀的临时文件, 并打开该文件用于读写并返回其*os.File指针. 文件名通过pattern生成, 并以一个随机字符串结尾. 如果pattern包含"*", 会使用随机字符串来替换最后的"*". 如果如果dir是空字符串，TempFile使用默认的临时文件夹(参考os.TempDir方法). 如果多个程序调用该函数的话，它不会选中相同的文件(即并发安全). 调用者可以使用f.Name()函数来找到该文件的路径. 调用者有责任在不需要时删除它.

<a id="example_TempFile">Example</a>
```go
package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
```

<a id="example_TempFile_suffix">Example (Suffix)</a>
```go
package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example.*.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		tmpfile.Close()
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
```

## <a id="WriteFile">func</a> [WriteFile](https://golang.org/src/io/ioutil/ioutil.go?s=2534:2602#L69)
<pre>func WriteFile(filename <a href="/pkg/builtin/#string">string</a>, data []<a href="/pkg/builtin/#byte">byte</a>, perm <a href="/pkg/os/">os</a>.<a href="/pkg/os/#FileMode">FileMode</a>) <a href="/pkg/builtin/#error">error</a></pre>

WriteFile writes data to a file named by filename.
If the file does not exist, WriteFile creates it with permissions perm;
otherwise WriteFile truncates it before writing.

WriteFile 会向指定的文件filename写入数据. 如果该文件不存在, 则创建带有指定权限信息的文件; 否则会在写入前先截断.

<a id="example_WriteFile">Example</a>
```go
package main

import (
	"io/ioutil"
	"log"
)

func main() {
	message := []byte("Hello, Gophers!")
	err := ioutil.WriteFile("testdata/hello", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
```

output:
```txt
2009/11/10 23:00:00 open testdata/hello: No such file or directory
```






