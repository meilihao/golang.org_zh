

# filepath
`import "path/filepath"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package filepath implements utility routines for manipulating filename paths
in a way compatible with the target operating system-defined file paths.

The filepath package uses either forward slashes or backslashes,
depending on the operating system. To process paths such as URLs
that always use forward slashes regardless of the operating
system, see the path package.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func Abs(path string) (string, error)](#Abs)
* [func Base(path string) string](#Base)
* [func Clean(path string) string](#Clean)
* [func Dir(path string) string](#Dir)
* [func EvalSymlinks(path string) (string, error)](#EvalSymlinks)
* [func Ext(path string) string](#Ext)
* [func FromSlash(path string) string](#FromSlash)
* [func Glob(pattern string) (matches []string, err error)](#Glob)
* [func HasPrefix(p, prefix string) bool](#HasPrefix)
* [func IsAbs(path string) bool](#IsAbs)
* [func Join(elem ...string) string](#Join)
* [func Match(pattern, name string) (matched bool, err error)](#Match)
* [func Rel(basepath, targpath string) (string, error)](#Rel)
* [func Split(path string) (dir, file string)](#Split)
* [func SplitList(path string) []string](#SplitList)
* [func ToSlash(path string) string](#ToSlash)
* [func VolumeName(path string) string](#VolumeName)
* [func Walk(root string, walkFn WalkFunc) error](#Walk)
* [type WalkFunc](#WalkFunc)


#### <a id="pkg-examples">Examples</a>
* [Base](#example_Base)
* [Dir](#example_Dir)
* [Ext](#example_Ext)
* [IsAbs](#example_IsAbs)
* [Join](#example_Join)
* [Match](#example_Match)
* [Rel](#example_Rel)
* [Split](#example_Split)
* [SplitList](#example_SplitList)
* [Walk](#example_Walk)


#### <a id="pkg-files">Package files</a>
[match.go](https://golang.org/src/path/filepath/match.go) [path.go](https://golang.org/src/path/filepath/path.go) [path_unix.go](https://golang.org/src/path/filepath/path_unix.go) [symlink.go](https://golang.org/src/path/filepath/symlink.go) [symlink_unix.go](https://golang.org/src/path/filepath/symlink_unix.go) 


## <a id="pkg-constants">Constants</a>

<pre>const (
    <span id="Separator">Separator</span>     = <a href="/pkg/os/">os</a>.<a href="/pkg/os/#PathSeparator">PathSeparator</a>
    <span id="ListSeparator">ListSeparator</span> = <a href="/pkg/os/">os</a>.<a href="/pkg/os/#PathListSeparator">PathListSeparator</a>
)</pre>

## <a id="pkg-variables">Variables</a>
ErrBadPattern indicates a pattern was malformed.


<pre>var <span id="ErrBadPattern">ErrBadPattern</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;syntax error in pattern&#34;)</pre>SkipDir is used as a return value from WalkFuncs to indicate that
the directory named in the call is to be skipped. It is not returned
as an error by any function.


<pre>var <span id="SkipDir">SkipDir</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;skip this directory&#34;)</pre>

## <a id="Abs">func</a> [Abs](https://golang.org/src/path/filepath/path.go?s=6920:6957#L230)
<pre>func Abs(path <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Abs returns an absolute representation of path.
If the path is not absolute it will be joined with the current
working directory to turn it into an absolute path. The absolute
path name for a given file is not guaranteed to be unique.
Abs calls Clean on the result.



## <a id="Base">func</a> [Base](https://golang.org/src/path/filepath/path.go?s=12939:12968#L422)
<pre>func Base(path <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Base returns the last element of path.
Trailing path separators are removed before extracting the last element.
If the path is empty, Base returns ".".
If the path consists entirely of separators, Base returns a single separator.


<a id="example_Base">Example</a>
```go
```

output:
```txt
```

## <a id="Clean">func</a> [Clean](https://golang.org/src/path/filepath/path.go?s=2501:2531#L78)
<pre>func Clean(path <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Clean returns the shortest path name equivalent to path
by purely lexical processing. It applies the following rules
iteratively until no further processing can be done:


	1. Replace multiple Separator elements with a single one.
	2. Eliminate each . path name element (the current directory).
	3. Eliminate each inner .. path name element (the parent directory)
	   along with the non-.. element that precedes it.
	4. Eliminate .. elements that begin a rooted path:
	   that is, replace "/.." by "/" at the beginning of a path,
	   assuming Separator is '/'.

The returned path ends in a slash only if it represents a root directory,
such as "/" on Unix or `C:\` on Windows.

Finally, any occurrences of slash are replaced by Separator.

If the result of this process is an empty string, Clean
returns the string ".".

See also Rob Pike, ``Lexical File Names in Plan 9 or
Getting Dot-Dot Right,''
<a href="https://9p.io/sys/doc/lexnames.html">https://9p.io/sys/doc/lexnames.html</a>



## <a id="Dir">func</a> [Dir](https://golang.org/src/path/filepath/path.go?s=13818:13846#L453)
<pre>func Dir(path <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Dir returns all but the last element of path, typically the path's directory.
After dropping the final element, Dir calls Clean on the path and trailing
slashes are removed.
If the path is empty, Dir returns ".".
If the path consists entirely of separators, Dir returns a single separator.
The returned path does not end in a separator unless it is the root directory.


<a id="example_Dir">Example</a>
```go
```

output:
```txt
```

## <a id="EvalSymlinks">func</a> [EvalSymlinks](https://golang.org/src/path/filepath/path.go?s=6560:6606#L221)
<pre>func EvalSymlinks(path <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
EvalSymlinks returns the path name after the evaluation of any symbolic
links.
If path is relative the result will be relative to the current directory,
unless one of the components is an absolute symbolic link.
EvalSymlinks calls Clean on the result.



## <a id="Ext">func</a> [Ext](https://golang.org/src/path/filepath/path.go?s=6129:6157#L207)
<pre>func Ext(path <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Ext returns the file name extension used by path.
The extension is the suffix beginning at the final dot
in the final element of path; it is empty if there is
no dot.


<a id="example_Ext">Example</a>
```go
```

output:
```txt
```

## <a id="FromSlash">func</a> [FromSlash](https://golang.org/src/path/filepath/path.go?s=4729:4763#L165)
<pre>func FromSlash(path <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
FromSlash returns the result of replacing each slash ('/') character
in path with a separator character. Multiple slashes are replaced
by multiple separators.



## <a id="Glob">func</a> [Glob](https://golang.org/src/path/filepath/match.go?s=5600:5655#L224)
<pre>func Glob(pattern <a href="/pkg/builtin/#string">string</a>) (matches []<a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Glob returns the names of all files matching pattern or nil
if there is no matching file. The syntax of patterns is the same
as in Match. The pattern may describe hierarchical names such as
/usr/*/bin/ed (assuming the Separator is '/').

Glob ignores file system errors such as I/O errors reading directories.
The only possible returned error is ErrBadPattern, when pattern
is malformed.



## <a id="HasPrefix">func</a> [HasPrefix](https://golang.org/src/path/filepath/path_unix.go?s=722:759#L16)
<pre>func HasPrefix(p, prefix <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
HasPrefix exists for historical compatibility and should not be used.

Deprecated: HasPrefix does not respect path boundaries and
does not ignore case when required.



## <a id="IsAbs">func</a> [IsAbs](https://golang.org/src/path/filepath/path_unix.go?s=325:353#L2)
<pre>func IsAbs(path <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsAbs reports whether the path is absolute.


<a id="example_IsAbs">Example</a>
```go
```

output:
```txt
```

## <a id="Join">func</a> [Join](https://golang.org/src/path/filepath/path.go?s=5893:5925#L199)
<pre>func Join(elem ...<a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Join joins any number of path elements into a single path, adding
a Separator if necessary. Join calls Clean on the result; in particular,
all empty strings are ignored.
On Windows, the result is a UNC path if and only if the first path
element is a UNC path.


<a id="example_Join">Example</a>
```go
```

output:
```txt
```

## <a id="Match">func</a> [Match](https://golang.org/src/path/filepath/match.go?s=1226:1284#L34)
<pre>func Match(pattern, name <a href="/pkg/builtin/#string">string</a>) (matched <a href="/pkg/builtin/#bool">bool</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Match reports whether name matches the shell file name pattern.
The pattern syntax is:


	pattern:
		{ term }
	term:
		'*'         matches any sequence of non-Separator characters
		'?'         matches any single non-Separator character
		'[' [ '^' ] { character-range } ']'
		            character class (must be non-empty)
		c           matches character c (c != '*', '?', '\\', '[')
		'\\' c      matches character c
	
	character-range:
		c           matches character c (c != '\\', '-', ']')
		'\\' c      matches character c
		lo '-' hi   matches character c for lo <= c <= hi

Match requires pattern to match all of name, not just a substring.
The only possible returned error is ErrBadPattern, when pattern
is malformed.

On Windows, escaping is disabled. Instead, '\\' is treated as
path separator.


<a id="example_Match">Example</a>
```go
```

output:
```txt
```

## <a id="Rel">func</a> [Rel](https://golang.org/src/path/filepath/path.go?s=7687:7738#L253)
<pre>func Rel(basepath, targpath <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Rel returns a relative path that is lexically equivalent to targpath when
joined to basepath with an intervening separator. That is,
Join(basepath, Rel(basepath, targpath)) is equivalent to targpath itself.
On success, the returned path will always be relative to basepath,
even if basepath and targpath share no elements.
An error is returned if targpath can't be made relative to basepath or if
knowing the current working directory would be necessary to compute it.
Rel calls Clean on the result.


<a id="example_Rel">Example</a>
```go
```

output:
```txt
```

## <a id="Split">func</a> [Split](https://golang.org/src/path/filepath/path.go?s=5432:5474#L185)
<pre>func Split(path <a href="/pkg/builtin/#string">string</a>) (dir, file <a href="/pkg/builtin/#string">string</a>)</pre>
Split splits path immediately following the final Separator,
separating it into a directory and file name component.
If there is no Separator in path, Split returns an empty dir
and file set to path.
The returned values have the property that path = dir+file.


<a id="example_Split">Example</a>
```go
```

output:
```txt
```

## <a id="SplitList">func</a> [SplitList](https://golang.org/src/path/filepath/path.go?s=5091:5127#L176)
<pre>func SplitList(path <a href="/pkg/builtin/#string">string</a>) []<a href="/pkg/builtin/#string">string</a></pre>
SplitList splits a list of paths joined by the OS-specific ListSeparator,
usually found in PATH or GOPATH environment variables.
Unlike strings.Split, SplitList returns an empty slice when passed an empty
string.


<a id="example_SplitList">Example</a>
```go
```

output:
```txt
```

## <a id="ToSlash">func</a> [ToSlash](https://golang.org/src/path/filepath/path.go?s=4426:4458#L155)
<pre>func ToSlash(path <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
ToSlash returns the result of replacing each separator character
in path with a slash ('/') character. Multiple separators are
replaced by multiple slashes.



## <a id="VolumeName">func</a> [VolumeName](https://golang.org/src/path/filepath/path.go?s=14264:14299#L471)
<pre>func VolumeName(path <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
VolumeName returns leading volume name.
Given "C:\foo\bar" it returns "C:" on Windows.
Given "\\host\share\foo" it returns "\\host\share".
On other platforms it returns "".



## <a id="Walk">func</a> [Walk](https://golang.org/src/path/filepath/path.go?s=12124:12169#L389)
<pre>func Walk(root <a href="/pkg/builtin/#string">string</a>, walkFn <a href="#WalkFunc">WalkFunc</a>) <a href="/pkg/builtin/#error">error</a></pre>
Walk walks the file tree rooted at root, calling walkFn for each file or
directory in the tree, including root. All errors that arise visiting files
and directories are filtered by walkFn. The files are walked in lexical
order, which makes the output deterministic but means that for very
large directories Walk can be inefficient.
Walk does not follow symbolic links.


<a id="example_Walk">Example</a>
```go
```

output:
```txt
```



## <a id="WalkFunc">type</a> [WalkFunc](https://golang.org/src/path/filepath/path.go?s=10511:10577#L341)
WalkFunc is the type of the function called for each file or directory
visited by Walk. The path argument contains the argument to Walk as a
prefix; that is, if Walk is called with "dir", which is a directory
containing the file "a", the walk function will be called with argument
"dir/a". The info argument is the os.FileInfo for the named path.

If there was a problem walking to the file or directory named by path, the
incoming error will describe the problem and the function can decide how
to handle that error (and Walk will not descend into that directory). In the
case of an error, the info argument will be nil. If an error is returned,
processing stops. The sole exception is when the function returns the special
value SkipDir. If the function returns SkipDir when invoked on a directory,
Walk skips the directory's contents entirely. If the function returns SkipDir
when invoked on a non-directory file, Walk skips the remaining files in the
containing directory.


<pre>type WalkFunc func(path <a href="/pkg/builtin/#string">string</a>, info <a href="/pkg/os/">os</a>.<a href="/pkg/os/#FileInfo">FileInfo</a>, err <a href="/pkg/builtin/#error">error</a>) <a href="/pkg/builtin/#error">error</a></pre>














