

# path
`import "path"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package path implements utility routines for manipulating slash-separated
paths.

The path package should only be used for paths separated by forward
slashes, such as the paths in URLs. This package does not deal with
Windows paths with drive letters or backslashes; to manipulate
operating system paths, use the path/filepath package.




## <a id="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func Base(path string) string](#Base)
* [func Clean(path string) string](#Clean)
* [func Dir(path string) string](#Dir)
* [func Ext(path string) string](#Ext)
* [func IsAbs(path string) bool](#IsAbs)
* [func Join(elem ...string) string](#Join)
* [func Match(pattern, name string) (matched bool, err error)](#Match)
* [func Split(path string) (dir, file string)](#Split)


#### <a id="pkg-examples">Examples</a>
* [Base](#example_Base)
* [Clean](#example_Clean)
* [Dir](#example_Dir)
* [Ext](#example_Ext)
* [IsAbs](#example_IsAbs)
* [Join](#example_Join)
* [Match](#example_Match)
* [Split](#example_Split)


#### <a id="pkg-files">Package files</a>
[match.go](https://golang.org/src/path/match.go) [path.go](https://golang.org/src/path/path.go) 




## <a id="pkg-variables">Variables</a>
ErrBadPattern indicates a pattern was malformed.


<pre>var <span id="ErrBadPattern">ErrBadPattern</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;syntax error in pattern&#34;)</pre>

## <a id="Base">func</a> [Base](https://golang.org/src/path/path.go?s=4737:4766#L171)
<pre>func Base(path <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Base returns the last element of path.
Trailing slashes are removed before extracting the last element.
If the path is empty, Base returns ".".
If the path consists entirely of slashes, Base returns "/".


<a id="example_Base">Example</a>
```go
```

output:
```txt
```

## <a id="Clean">func</a> [Clean](https://golang.org/src/path/path.go?s=2075:2105#L64)
<pre>func Clean(path <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Clean returns the shortest path name equivalent to path
by purely lexical processing. It applies the following rules
iteratively until no further processing can be done:


	1. Replace multiple slashes with a single slash.
	2. Eliminate each . path name element (the current directory).
	3. Eliminate each inner .. path name element (the parent directory)
	   along with the non-.. element that precedes it.
	4. Eliminate .. elements that begin a rooted path:
	   that is, replace "/.." by "/" at the beginning of a path.

The returned path ends in a slash only if it is the root "/".

If the result of this process is an empty string, Clean
returns the string ".".

See also Rob Pike, ``Lexical File Names in Plan 9 or
Getting Dot-Dot Right,''
<a href="https://9p.io/sys/doc/lexnames.html">https://9p.io/sys/doc/lexnames.html</a>


<a id="example_Clean">Example</a>
```go
```

output:
```txt
```

## <a id="Dir">func</a> [Dir](https://golang.org/src/path/path.go?s=5617:5645#L202)
<pre>func Dir(path <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Dir returns all but the last element of path, typically the path's directory.
After dropping the final element using Split, the path is Cleaned and trailing
slashes are removed.
If the path is empty, Dir returns ".".
If the path consists entirely of slashes followed by non-slash bytes, Dir
returns a single slash. In any other case, the returned path does not end in a
slash.


<a id="example_Dir">Example</a>
```go
```

output:
```txt
```

## <a id="Ext">func</a> [Ext](https://golang.org/src/path/path.go?s=4371:4399#L158)
<pre>func Ext(path <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Ext returns the file name extension used by path.
The extension is the suffix beginning at the final dot
in the final slash-separated element of path;
it is empty if there is no dot.


<a id="example_Ext">Example</a>
```go
```

output:
```txt
```

## <a id="IsAbs">func</a> [IsAbs](https://golang.org/src/path/path.go?s=5145:5173#L191)
<pre>func IsAbs(path <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsAbs reports whether the path is absolute.


<a id="example_IsAbs">Example</a>
```go
```

output:
```txt
```

## <a id="Join">func</a> [Join](https://golang.org/src/path/path.go?s=4034:4066#L145)
<pre>func Join(elem ...<a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Join joins any number of path elements into a single path, adding a
separating slash if necessary. The result is Cleaned; in particular,
all empty strings are ignored.


<a id="example_Join">Example</a>
```go
```

output:
```txt
```

## <a id="Match">func</a> [Match](https://golang.org/src/path/match.go?s=1084:1142#L28)
<pre>func Match(pattern, name <a href="/pkg/builtin/#string">string</a>) (matched <a href="/pkg/builtin/#bool">bool</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Match reports whether name matches the shell pattern.
The pattern syntax is:


	pattern:
		{ term }
	term:
		'*'         matches any sequence of non-/ characters
		'?'         matches any single non-/ character
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


<a id="example_Match">Example</a>
```go
```

output:
```txt
```

## <a id="Split">func</a> [Split](https://golang.org/src/path/path.go?s=3743:3785#L137)
<pre>func Split(path <a href="/pkg/builtin/#string">string</a>) (dir, file <a href="/pkg/builtin/#string">string</a>)</pre>
Split splits path immediately following the final slash,
separating it into a directory and file name component.
If there is no slash in path, Split returns an empty dir and
file set to path.
The returned values have the property that path = dir+file.


<a id="example_Split">Example</a>
```go
```

output:
```txt
```






