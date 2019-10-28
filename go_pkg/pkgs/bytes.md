

# bytes
`import "bytes"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package bytes implements functions for the manipulation of byte slices.
It is analogous to the facilities of the strings package.

bytes实现了操作byte切片的函数. 它与strings包的功能类似.


## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func Compare(a, b []byte) int](#Compare)
* [func Contains(b, subslice []byte) bool](#Contains)
* [func ContainsAny(b []byte, chars string) bool](#ContainsAny)
* [func ContainsRune(b []byte, r rune) bool](#ContainsRune)
* [func Count(s, sep []byte) int](#Count)
* [func Equal(a, b []byte) bool](#Equal)
* [func EqualFold(s, t []byte) bool](#EqualFold)
* [func Fields(s []byte) [][]byte](#Fields)
* [func FieldsFunc(s []byte, f func(rune) bool) [][]byte](#FieldsFunc)
* [func HasPrefix(s, prefix []byte) bool](#HasPrefix)
* [func HasSuffix(s, suffix []byte) bool](#HasSuffix)
* [func Index(s, sep []byte) int](#Index)
* [func IndexAny(s []byte, chars string) int](#IndexAny)
* [func IndexByte(b []byte, c byte) int](#IndexByte)
* [func IndexFunc(s []byte, f func(r rune) bool) int](#IndexFunc)
* [func IndexRune(s []byte, r rune) int](#IndexRune)
* [func Join(s [][]byte, sep []byte) []byte](#Join)
* [func LastIndex(s, sep []byte) int](#LastIndex)
* [func LastIndexAny(s []byte, chars string) int](#LastIndexAny)
* [func LastIndexByte(s []byte, c byte) int](#LastIndexByte)
* [func LastIndexFunc(s []byte, f func(r rune) bool) int](#LastIndexFunc)
* [func Map(mapping func(r rune) rune, s []byte) []byte](#Map)
* [func Repeat(b []byte, count int) []byte](#Repeat)
* [func Replace(s, old, new []byte, n int) []byte](#Replace)
* [func ReplaceAll(s, old, new []byte) []byte](#ReplaceAll)
* [func Runes(s []byte) []rune](#Runes)
* [func Split(s, sep []byte) [][]byte](#Split)
* [func SplitAfter(s, sep []byte) [][]byte](#SplitAfter)
* [func SplitAfterN(s, sep []byte, n int) [][]byte](#SplitAfterN)
* [func SplitN(s, sep []byte, n int) [][]byte](#SplitN)
* [func Title(s []byte) []byte](#Title)
* [func ToLower(s []byte) []byte](#ToLower)
* [func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte](#ToLowerSpecial)
* [func ToTitle(s []byte) []byte](#ToTitle)
* [func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte](#ToTitleSpecial)
* [func ToUpper(s []byte) []byte](#ToUpper)
* [func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte](#ToUpperSpecial)
* [func ToValidUTF8(s, replacement []byte) []byte](#ToValidUTF8)
* [func Trim(s []byte, cutset string) []byte](#Trim)
* [func TrimFunc(s []byte, f func(r rune) bool) []byte](#TrimFunc)
* [func TrimLeft(s []byte, cutset string) []byte](#TrimLeft)
* [func TrimLeftFunc(s []byte, f func(r rune) bool) []byte](#TrimLeftFunc)
* [func TrimPrefix(s, prefix []byte) []byte](#TrimPrefix)
* [func TrimRight(s []byte, cutset string) []byte](#TrimRight)
* [func TrimRightFunc(s []byte, f func(r rune) bool) []byte](#TrimRightFunc)
* [func TrimSpace(s []byte) []byte](#TrimSpace)
* [func TrimSuffix(s, suffix []byte) []byte](#TrimSuffix)
* [type Buffer](#Buffer)
  * [func NewBuffer(buf []byte) *Buffer](#NewBuffer)
  * [func NewBufferString(s string) *Buffer](#NewBufferString)
  * [func (b *Buffer) Bytes() []byte](#Buffer.Bytes)
  * [func (b *Buffer) Cap() int](#Buffer.Cap)
  * [func (b *Buffer) Grow(n int)](#Buffer.Grow)
  * [func (b *Buffer) Len() int](#Buffer.Len)
  * [func (b *Buffer) Next(n int) []byte](#Buffer.Next)
  * [func (b *Buffer) Read(p []byte) (n int, err error)](#Buffer.Read)
  * [func (b *Buffer) ReadByte() (byte, error)](#Buffer.ReadByte)
  * [func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)](#Buffer.ReadBytes)
  * [func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)](#Buffer.ReadFrom)
  * [func (b *Buffer) ReadRune() (r rune, size int, err error)](#Buffer.ReadRune)
  * [func (b *Buffer) ReadString(delim byte) (line string, err error)](#Buffer.ReadString)
  * [func (b *Buffer) Reset()](#Buffer.Reset)
  * [func (b *Buffer) String() string](#Buffer.String)
  * [func (b *Buffer) Truncate(n int)](#Buffer.Truncate)
  * [func (b *Buffer) UnreadByte() error](#Buffer.UnreadByte)
  * [func (b *Buffer) UnreadRune() error](#Buffer.UnreadRune)
  * [func (b *Buffer) Write(p []byte) (n int, err error)](#Buffer.Write)
  * [func (b *Buffer) WriteByte(c byte) error](#Buffer.WriteByte)
  * [func (b *Buffer) WriteRune(r rune) (n int, err error)](#Buffer.WriteRune)
  * [func (b *Buffer) WriteString(s string) (n int, err error)](#Buffer.WriteString)
  * [func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)](#Buffer.WriteTo)
* [type Reader](#Reader)
  * [func NewReader(b []byte) *Reader](#NewReader)
  * [func (r *Reader) Len() int](#Reader.Len)
  * [func (r *Reader) Read(b []byte) (n int, err error)](#Reader.Read)
  * [func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)](#Reader.ReadAt)
  * [func (r *Reader) ReadByte() (byte, error)](#Reader.ReadByte)
  * [func (r *Reader) ReadRune() (ch rune, size int, err error)](#Reader.ReadRune)
  * [func (r *Reader) Reset(b []byte)](#Reader.Reset)
  * [func (r *Reader) Seek(offset int64, whence int) (int64, error)](#Reader.Seek)
  * [func (r *Reader) Size() int64](#Reader.Size)
  * [func (r *Reader) UnreadByte() error](#Reader.UnreadByte)
  * [func (r *Reader) UnreadRune() error](#Reader.UnreadRune)
  * [func (r *Reader) WriteTo(w io.Writer) (n int64, err error)](#Reader.WriteTo)


#### <a id="pkg-examples">Examples</a>
* [Buffer](#example_Buffer)
* [Buffer.Grow](#example_Buffer_Grow)
* [Buffer.Len](#example_Buffer_Len)
* [Buffer (Reader)](#example_Buffer_reader)
* [Compare](#example_Compare)
* [Compare (Search)](#example_Compare_search)
* [Contains](#example_Contains)
* [ContainsAny](#example_ContainsAny)
* [ContainsRune](#example_ContainsRune)
* [Count](#example_Count)
* [Equal](#example_Equal)
* [EqualFold](#example_EqualFold)
* [Fields](#example_Fields)
* [FieldsFunc](#example_FieldsFunc)
* [HasPrefix](#example_HasPrefix)
* [HasSuffix](#example_HasSuffix)
* [Index](#example_Index)
* [IndexAny](#example_IndexAny)
* [IndexByte](#example_IndexByte)
* [IndexFunc](#example_IndexFunc)
* [IndexRune](#example_IndexRune)
* [Join](#example_Join)
* [LastIndex](#example_LastIndex)
* [LastIndexAny](#example_LastIndexAny)
* [LastIndexByte](#example_LastIndexByte)
* [LastIndexFunc](#example_LastIndexFunc)
* [Map](#example_Map)
* [Reader.Len](#example_Reader_Len)
* [Repeat](#example_Repeat)
* [Replace](#example_Replace)
* [ReplaceAll](#example_ReplaceAll)
* [Runes](#example_Runes)
* [Split](#example_Split)
* [SplitAfter](#example_SplitAfter)
* [SplitAfterN](#example_SplitAfterN)
* [SplitN](#example_SplitN)
* [Title](#example_Title)
* [ToLower](#example_ToLower)
* [ToLowerSpecial](#example_ToLowerSpecial)
* [ToTitle](#example_ToTitle)
* [ToTitleSpecial](#example_ToTitleSpecial)
* [ToUpper](#example_ToUpper)
* [ToUpperSpecial](#example_ToUpperSpecial)
* [Trim](#example_Trim)
* [TrimFunc](#example_TrimFunc)
* [TrimLeft](#example_TrimLeft)
* [TrimLeftFunc](#example_TrimLeftFunc)
* [TrimPrefix](#example_TrimPrefix)
* [TrimRight](#example_TrimRight)
* [TrimRightFunc](#example_TrimRightFunc)
* [TrimSpace](#example_TrimSpace)
* [TrimSuffix](#example_TrimSuffix)


#### <a id="pkg-files">Package files</a>
[buffer.go](https://golang.org/src/bytes/buffer.go) [bytes.go](https://golang.org/src/bytes/bytes.go) [reader.go](https://golang.org/src/bytes/reader.go) 


## <a id="pkg-constants">Constants</a>
MinRead is the minimum slice size passed to a Read call by
Buffer.ReadFrom. As long as the Buffer has at least MinRead bytes beyond
what is required to hold the contents of r, ReadFrom will not grow the
underlying buffer.

MinRead是Buffer.ReadFrom 调用 Read 时传递的最小尺寸. 只要Buffer至少有MinRead个字节的余量, 且余量足以容纳r中的内容, 那么ReadFrom就不会扩充底层的缓冲区.


<pre>const <span id="MinRead">MinRead</span> = 512</pre>

## <a id="pkg-variables">Variables</a>
ErrTooLarge is passed to panic if memory cannot be allocated to store data in a buffer.

ErrTooLarge是如果内存不能分配足够大的缓冲来保存数据时就会被传递给panic函数的error.

<pre>var <span id="ErrTooLarge">ErrTooLarge</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;bytes.Buffer: too large&#34;)</pre>

## <a id="Compare">func</a> [Compare](https://golang.org/src/bytes/bytes.go?s=833:862#L16)
<pre>func Compare(a, b []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#int">int</a></pre>
Compare returns an integer comparing two byte slices lexicographically.
The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
A nil argument is equivalent to an empty slice.

Compare返回一个整数表示两个byte切片按字典顺序比较的结果. 如果a==b返回0, a < b 返回-1,如果 a > b 返回+1. 传入nil则被视为空slice.

<a id="example_Compare">Example</a>
```go
package main

import (
	"bytes"
)

func main() {
	// Interpret Compare's result by comparing it to zero.
    // 通过和0比较来解释Compare的结果.
	var a, b []byte
	if bytes.Compare(a, b) < 0 {
		// a less b
	}
	if bytes.Compare(a, b) <= 0 {
		// a less or equal b
	}
	if bytes.Compare(a, b) > 0 {
		// a greater b
	}
	if bytes.Compare(a, b) >= 0 {
		// a greater or equal b
	}

	// Prefer Equal to Compare for equality comparisons.
    // 倾向使用Equal取代Compare来比较是否相等.
	if bytes.Equal(a, b) {
		// a equal b
	}
	if !bytes.Equal(a, b) {
		// a not equal b
	}
}
```
<a id="example_Compare_search">Example (Search)</a>
```go
package main

import (
	"bytes"
	"sort"
)

func main() {
	// Binary search to find a matching byte slice.
    // 通过二分查找来获取匹配的byte slice.
	var needle []byte
	var haystack [][]byte // Assume sorted
	i := sort.Search(len(haystack), func(i int) bool {
		// Return haystack[i] >= needle.
		return bytes.Compare(haystack[i], needle) >= 0
	})
	if i < len(haystack) && bytes.Equal(haystack[i], needle) {
		// Found it!
	}
}
```

## <a id="Contains">func</a> [Contains](https://golang.org/src/bytes/bytes.go?s=1857:1895#L65)
<pre>func Contains(b, subslice []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Contains reports whether subslice is within b.

Contains 是判断b是否包含subslice.

<a id="example_Contains">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("foo")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("bar")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("")))
	fmt.Println(bytes.Contains([]byte(""), []byte("")))
}
```

output:
```txt
true
false
true
true
```

## <a id="ContainsAny">func</a> [ContainsAny](https://golang.org/src/bytes/bytes.go?s=2025:2070#L70)
<pre>func ContainsAny(b []<a href="/pkg/builtin/#byte">byte</a>, chars <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
ContainsAny reports whether any of the UTF-8-encoded code points in chars are within b.

ContainsAny 是判断b是否包含chars中的任意一个UTF-8编码的Unicode码点(字符).

<a id="example_ContainsAny">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "fÄo!"))
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "去是伟大的."))
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), ""))
	fmt.Println(bytes.ContainsAny([]byte(""), ""))
}
```

output:
```txt
true
true
false
false
```

## <a id="ContainsRune">func</a> [ContainsRune](https://golang.org/src/bytes/bytes.go?s=2197:2237#L75)
<pre>func ContainsRune(b []<a href="/pkg/builtin/#byte">byte</a>, r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
ContainsRune reports whether the rune is contained in the UTF-8-encoded byte slice b.

ContainsRune 是判断b是否包含UTF-8编码的(字符)r.

<a id="example_ContainsRune">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.ContainsRune([]byte("I like seafood."), 'f'))
	fmt.Println(bytes.ContainsRune([]byte("I like seafood."), 'ö'))
	fmt.Println(bytes.ContainsRune([]byte("去是伟大的!"), '大'))
	fmt.Println(bytes.ContainsRune([]byte("去是伟大的!"), '!'))
	fmt.Println(bytes.ContainsRune([]byte(""), '@'))
}
```

output:
```txt
true
false
true
true
false
```

## <a id="Count">func</a> [Count](https://golang.org/src/bytes/bytes.go?s=1547:1576#L45)
<pre>func Count(s, sep []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#int">int</a></pre>
Count counts the number of non-overlapping instances of sep in s.
If sep is an empty slice, Count returns 1 + the number of UTF-8-encoded code points in s.

Count 是统计sep(不重叠的)在s中出现的次数. 如果sep为空slice,Count返回s中 UTF-8 码点的数量+1.

<a id="example_Count">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Count([]byte("cheese"), []byte("e")))
	fmt.Println(bytes.Count([]byte("five"), []byte(""))) // before & after each rune
}
```

output:
```txt
3
5
```

## <a id="Equal">func</a> [Equal](https://golang.org/src/bytes/bytes.go?s=505:533#L8)
<pre>func Equal(a, b []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Equal reports whether a and b
are the same length and contain the same bytes.
A nil argument is equivalent to an empty slice.

Equal 会返回一个布尔值判断a和b是否长度相等且内容相同. nil被视为空slice.

<a id="example_Equal">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Equal([]byte("Go"), []byte("Go")))
	fmt.Println(bytes.Equal([]byte("Go"), []byte("C++")))
}
```

output:
```txt
true
false
```
## <a id="EqualFold">func</a> [EqualFold](https://golang.org/src/bytes/bytes.go?s=24857:24889#L929)
<pre>func EqualFold(s, t []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
EqualFold reports whether s and t, interpreted as UTF-8 strings,
are equal under Unicode case-folding.

EqualFold 是判断s和t是否相等,判断时会先解析成UTF-8字符串再进行比较且忽略大小写.

<a id="example_EqualFold">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.EqualFold([]byte("Go"), []byte("go")))
}
```

output:
```txt
true
```
## <a id="Fields">func</a> [Fields](https://golang.org/src/bytes/bytes.go?s=8137:8167#L303)
<pre>func Fields(s []<a href="/pkg/builtin/#byte">byte</a>) [][]<a href="/pkg/builtin/#byte">byte</a></pre>
Fields interprets s as a sequence of UTF-8-encoded code points.
It splits the slice s around each instance of one or more consecutive white space
characters, as defined by unicode.IsSpace, returning a slice of subslices of s or an
empty slice if s contains only white space.

Fields会将s解析成UTF-8编码的码点, 再把一个或者连续多个空白符(由unicode.IsSpace定义)当作间隔来分割s并返回分割结果. 如果s中只包含空白符,则返回空slice.

<a id="example_Fields">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("Fields are: %q", bytes.Fields([]byte("  foo bar  baz   ")))
}
```

output:
```txt
Fields are: ["foo" "bar" "baz"]
```
## <a id="FieldsFunc">func</a> [FieldsFunc](https://golang.org/src/bytes/bytes.go?s=9597:9650#L359)
<pre>func FieldsFunc(s []<a href="/pkg/builtin/#byte">byte</a>, f func(<a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a>) [][]<a href="/pkg/builtin/#byte">byte</a></pre>
FieldsFunc interprets s as a sequence of UTF-8-encoded code points.
It splits the slice s at each run of code points c satisfying f(c) and
returns a slice of subslices of s. If all code points in s satisfy f(c), or
len(s) == 0, an empty slice is returned.
FieldsFunc makes no guarantees about the order in which it calls f(c).
If f does not return consistent results for a given c, FieldsFunc may crash.

FieldsFunc会将s解析为UTF-8编码的码点序列, 遍历所有码点判断是否满足f(c),依此将s切分成多个子切片(与Fields类似). 如果s中所有码点都满足f(c)或者len(s)==0,返回空slice. FieldsFunc不保证调用f(c)的顺序. 如果f没有根据给定的c返回一致(即相同)的结果,FieldsFunc可能会crash.

<a id="example_FieldsFunc">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", bytes.FieldsFunc([]byte("  foo1;bar2,baz3..."), f))
}
```

output:
```txt
Fields are: ["foo1" "bar2" "baz3"]
```
## <a id="HasPrefix">func</a> [HasPrefix](https://golang.org/src/bytes/bytes.go?s=11117:11154#L430)
<pre>func HasPrefix(s, prefix []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
HasPrefix tests whether the byte slice s begins with prefix.

HasPrefix 判断prefix是否是s的前缀.

<a id="example_HasPrefix">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("Go")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("C")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("")))
}
```

output:
```txt
true
false
true
```
## <a id="HasSuffix">func</a> [HasSuffix](https://golang.org/src/bytes/bytes.go?s=11287:11324#L435)
<pre>func HasSuffix(s, suffix []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
HasSuffix tests whether the byte slice s ends with suffix.

 HasSuffix 判断suffix是否是s的后缀.

<a id="example_HasSuffix">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("go")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("O")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("Ami")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("")))
}
```

output:
```txt
true
false
false
true
```
## <a id="Index">func</a> [Index](https://golang.org/src/bytes/bytes.go?s=26022:26051#L983)
<pre>func Index(s, sep []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#int">int</a></pre>
Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.

Index返回sep在s中第一次出现的位置,不存在则返回-1.

<a id="example_Index">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Index([]byte("chicken"), []byte("ken")))
	fmt.Println(bytes.Index([]byte("chicken"), []byte("dmr")))
}
```

output:
```txt
4
-1
```
## <a id="IndexAny">func</a> [IndexAny](https://golang.org/src/bytes/bytes.go?s=4507:4548#L171)
<pre>func IndexAny(s []<a href="/pkg/builtin/#byte">byte</a>, chars <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#int">int</a></pre>
IndexAny interprets s as a sequence of UTF-8-encoded Unicode code points.
It returns the byte index of the first occurrence in s of any of the Unicode
code points in chars. It returns -1 if chars is empty or if there is no code
point in common.

IndexAny会将s解析成UTF-8编码的码点序列,返回字符串chars中的任意一个Unicode码点在s中第一次出现的位置. 如果chars为空字符串或者不存在共同的码点则返回-1.

<a id="example_IndexAny">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.IndexAny([]byte("chicken"), "aeiouy"))
	fmt.Println(bytes.IndexAny([]byte("crwth"), "aeiouy"))
}
```

output:
```txt
2
-1
```
## <a id="IndexByte">func</a> [IndexByte](https://golang.org/src/bytes/bytes.go?s=2368:2404#L80)
<pre>func IndexByte(b []<a href="/pkg/builtin/#byte">byte</a>, c <a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#int">int</a></pre>
IndexByte returns the index of the first instance of c in b, or -1 if c is not present in b.

IndexByte 返回c在s中第一次出现的位置，不存在则返回-1.

<a id="example_IndexByte">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.IndexByte([]byte("chicken"), byte('k')))
	fmt.Println(bytes.IndexByte([]byte("chicken"), byte('g')))
}
```

output:
```txt
4
-1
```
## <a id="IndexFunc">func</a> [IndexFunc](https://golang.org/src/bytes/bytes.go?s=18897:18946#L710)
<pre>func IndexFunc(s []<a href="/pkg/builtin/#byte">byte</a>, f func(r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a>) <a href="/pkg/builtin/#int">int</a></pre>
IndexFunc interprets s as a sequence of UTF-8-encoded code points.
It returns the byte index in s of the first Unicode
code point satisfying f(c), or -1 if none do.

IndexFunc将s解析成UTF-8编码的码点序列, 返回s中第一个满足f(c)的Unicode码点的位置, 不存在则返回-1.

<a id="example_IndexFunc">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(bytes.IndexFunc([]byte("Hello, 世界"), f))
	fmt.Println(bytes.IndexFunc([]byte("Hello, world"), f))
}
```

output:
```txt
7
-1
```
## <a id="IndexRune">func</a> [IndexRune](https://golang.org/src/bytes/bytes.go?s=3835:3871#L145)
<pre>func IndexRune(s []<a href="/pkg/builtin/#byte">byte</a>, r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#int">int</a></pre>
IndexRune interprets s as a sequence of UTF-8-encoded code points.
It returns the byte index of the first occurrence in s of the given rune.
It returns -1 if rune is not present in s.
If r is utf8.RuneError, it returns the first instance of any
invalid UTF-8 byte sequence.

IndexRune将s解析成UTF-8编码的码点序列,返回r在s中第一次出现的位置;如果不存在则返回-1. 如果r是utf8.RuneError,它会返回第一个无效的UTF-8 byte序列的位置.

<a id="example_IndexRune">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.IndexRune([]byte("chicken"), 'k'))
	fmt.Println(bytes.IndexRune([]byte("chicken"), 'd'))
}
```

output:
```txt
4
-1
```
## <a id="Join">func</a> [Join](https://golang.org/src/bytes/bytes.go?s=10680:10720#L407)
<pre>func Join(s [][]<a href="/pkg/builtin/#byte">byte</a>, sep []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
Join concatenates the elements of s to create a new byte slice. The separator
sep is placed between elements in the resulting slice.

Join 会将一系列byte切片连接成一个新的byte切片,它们之间用sep来分隔.

<a id="example_Join">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	fmt.Printf("%s", bytes.Join(s, []byte(", ")))
}
```

output:
```txt
foo, bar, baz
```
## <a id="LastIndex">func</a> [LastIndex](https://golang.org/src/bytes/bytes.go?s=2658:2691#L94)
<pre>func LastIndex(s, sep []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#int">int</a></pre>
LastIndex returns the index of the last instance of sep in s, or -1 if sep is not present in s.

LastIndex 会返回切片sep在s中最后一次出现的位置, 不存在则返回-1.

<a id="example_LastIndex">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Index([]byte("go gopher"), []byte("go")))
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("go")))
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("rodent")))
}
```

output:
```txt
0
3
-1
```
## <a id="LastIndexAny">func</a> [LastIndexAny](https://golang.org/src/bytes/bytes.go?s=5281:5326#L207)
<pre>func LastIndexAny(s []<a href="/pkg/builtin/#byte">byte</a>, chars <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#int">int</a></pre>
LastIndexAny interprets s as a sequence of UTF-8-encoded Unicode code
points. It returns the byte index of the last occurrence in s of any of
the Unicode code points in chars. It returns -1 if chars is empty or if
there is no code point in common.

LastIndexAny 会将s解析成UTF-8编码的unicode码点序列, 返回字符串chars中的任意一个Unicode码点在s中最后一次出现的位置. 如果chars为空字符串或者不存在则返回-1.

<a id="example_LastIndexAny">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.LastIndexAny([]byte("go gopher"), "MüQp"))
	fmt.Println(bytes.LastIndexAny([]byte("go 地鼠"), "地大"))
	fmt.Println(bytes.LastIndexAny([]byte("go gopher"), "z,!."))
}
```

output:
```txt
5
3
-1
```
## <a id="LastIndexByte">func</a> [LastIndexByte](https://golang.org/src/bytes/bytes.go?s=3417:3457#L131)
<pre>func LastIndexByte(s []<a href="/pkg/builtin/#byte">byte</a>, c <a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#int">int</a></pre>
LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.

LastIndexByte 会返回字符c在s中最后一次出现的位置，不存在则返回-1.

<a id="example_LastIndexByte">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.LastIndexByte([]byte("go gopher"), byte('g')))
	fmt.Println(bytes.LastIndexByte([]byte("go gopher"), byte('r')))
	fmt.Println(bytes.LastIndexByte([]byte("go gopher"), byte('z')))
}
```

output:
```txt
3
8
-1
```
## <a id="LastIndexFunc">func</a> [LastIndexFunc](https://golang.org/src/bytes/bytes.go?s=19159:19212#L717)
<pre>func LastIndexFunc(s []<a href="/pkg/builtin/#byte">byte</a>, f func(r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a>) <a href="/pkg/builtin/#int">int</a></pre>
LastIndexFunc interprets s as a sequence of UTF-8-encoded code points.
It returns the byte index in s of the last Unicode
code point satisfying f(c), or -1 if none do.

LastIndexFunc 会将s解析成UTF-8编码的码点序列, 返回s中最后一个满足f(c)的码点的位置, 不存在则返回-1.

<a id="example_LastIndexFunc">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsLetter))
	fmt.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsPunct))
	fmt.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsNumber))
}
```

output:
```txt
8
9
-1
```
## <a id="Map">func</a> [Map](https://golang.org/src/bytes/bytes.go?s=11704:11756#L443)
<pre>func Map(mapping func(r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#rune">rune</a>, s []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
Map returns a copy of the byte slice s with all its characters modified
according to the mapping function. If mapping returns a negative value, the character is
dropped from the byte slice with no replacement. The characters in s and the
output are interpreted as UTF-8-encoded code points.

Map 返回s的副本, 且所有字符已根据mapping函数进行了修改. 如果mapping返回一个负值, 将会丢弃该码值且不会被替换.s和输出结果均被视为UTF-8编码的码点序列.

<a id="example_Map">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Printf("%s", bytes.Map(rot13, []byte("'Twas brillig and the slithy gopher...")))
}
```

output:
```txt
'Gjnf oevyyvt naq gur fyvgul tbcure...
```
## <a id="Repeat">func</a> [Repeat](https://golang.org/src/bytes/bytes.go?s=12715:12754#L480)
<pre>func Repeat(b []<a href="/pkg/builtin/#byte">byte</a>, count <a href="/pkg/builtin/#int">int</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
Repeat returns a new byte slice consisting of count copies of b.

It panics if count is negative or if
the result of (len(b) * count) overflows.

Repeat 返回count个b串联而成的新slice.

如果count是负数或(len(b) * count)的值溢出, 它会panic.

<a id="example_Repeat">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("ba%s", bytes.Repeat([]byte("na"), 2))
}
```

output:
```txt
banana
```
## <a id="Replace">func</a> [Replace](https://golang.org/src/bytes/bytes.go?s=23763:23809#L882)
<pre>func Replace(s, old, new []<a href="/pkg/builtin/#byte">byte</a>, n <a href="/pkg/builtin/#int">int</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
Replace returns a copy of the slice s with the first n
non-overlapping instances of old replaced by new.
If old is empty, it matches at the beginning of the slice
and after each UTF-8 sequence, yielding up to k+1 replacements
for a k-rune slice.
If n < 0, there is no limit on the number of replacements.

Replace 返回s的一个副本,且将s中前n个不重叠的old切片替换成了new. 如果old为空,则在每个utf8字符前面和后面都插入一个new, k个字符需要k+1次替换. 如果n<0则不会限制替换次数.

<a id="example_Replace">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("k"), []byte("ky"), 2))
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("oink"), []byte("moo"), -1))
}
```

output:
```txt
oinky oinky oink
moo moo moo
```
## <a id="ReplaceAll">func</a> [ReplaceAll](https://golang.org/src/bytes/bytes.go?s=24667:24709#L923)
<pre>func ReplaceAll(s, old, new []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
ReplaceAll returns a copy of the slice s with all
non-overlapping instances of old replaced by new.
If old is empty, it matches at the beginning of the slice
and after each UTF-8 sequence, yielding up to k+1 replacements
for a k-rune slice.

ReplaceAll 返回s的一个副本,且将s中所有不重叠的old切片替换成了new. 如果old为空,则在每个utf8字符前面和后面都插入一个new, k个字符需要k+1次替换.

<a id="example_ReplaceAll">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", bytes.ReplaceAll([]byte("oink oink oink"), []byte("oink"), []byte("moo")))
}
```

output:
```txt
moo moo moo
```
## <a id="Runes">func</a> [Runes](https://golang.org/src/bytes/bytes.go?s=23272:23299#L864)
<pre>func Runes(s []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#rune">rune</a></pre>
Runes interprets s as a sequence of UTF-8-encoded code points.
It returns a slice of runes (Unicode code points) equivalent to s.

Runes 将s解析成UTF-8编码的码点序列, 并返回和s等价的rune slice(Unicode编码). 

<a id="example_Runes">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	rs := bytes.Runes([]byte("go gopher"))
	for _, r := range rs {
		fmt.Printf("%#U\n", r)
	}
}
```

output:
```txt
U+0067 'g'
U+006F 'o'
U+0020 ' '
U+0067 'g'
U+006F 'o'
U+0070 'p'
U+0068 'h'
U+0065 'e'
U+0072 'r'
```
## <a id="Split">func</a> [Split](https://golang.org/src/bytes/bytes.go?s=7381:7415#L287)
<pre>func Split(s, sep []<a href="/pkg/builtin/#byte">byte</a>) [][]<a href="/pkg/builtin/#byte">byte</a></pre>
Split slices s into all subslices separated by sep and returns a slice of
the subslices between those separators.
If sep is empty, Split splits after each UTF-8 sequence.
It is equivalent to SplitN with a count of -1.

Split 根据sep将s分割并返回分割结果. 如果sep为空,Split按每个utf8编码切分, 它等效于参数为-1的SplitN方法.

<a id="example_Split">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte(",")))
	fmt.Printf("%q\n", bytes.Split([]byte("a man a plan a canal panama"), []byte("a ")))
	fmt.Printf("%q\n", bytes.Split([]byte(" xyz "), []byte("")))
	fmt.Printf("%q\n", bytes.Split([]byte(""), []byte("Bernardo O'Higgins")))
}
```

output:
```txt
["a" "b" "c"]
["" "man " "plan " "canal panama"]
[" " "x" "y" "z" " "]
[""]
```
## <a id="SplitAfter">func</a> [SplitAfter](https://golang.org/src/bytes/bytes.go?s=7684:7723#L293)
<pre>func SplitAfter(s, sep []<a href="/pkg/builtin/#byte">byte</a>) [][]<a href="/pkg/builtin/#byte">byte</a></pre>
SplitAfter slices s into all subslices after each instance of sep and
returns a slice of those subslices.
If sep is empty, SplitAfter splits after each UTF-8 sequence.
It is equivalent to SplitAfterN with a count of -1.


SplitAfter 从 s 中的每个 sep 之后的位置分割 s 并返回分割结果. 如sep为空,SplitAfter会按每个utf8编码切分. 它等效于参数为-1的SplitAfterN方法.

<a id="example_SplitAfter">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%q\n", bytes.SplitAfter([]byte("a,b,c"), []byte(",")))
}
```

output:
```txt
["a," "b," "c"]
```
## <a id="SplitAfterN">func</a> [SplitAfterN](https://golang.org/src/bytes/bytes.go?s=7060:7107#L279)
<pre>func SplitAfterN(s, sep []<a href="/pkg/builtin/#byte">byte</a>, n <a href="/pkg/builtin/#int">int</a>) [][]<a href="/pkg/builtin/#byte">byte</a></pre>
SplitAfterN slices s into subslices after each instance of sep and
returns a slice of those subslices.
If sep is empty, SplitAfterN splits after each UTF-8 sequence.
The count determines the number of subslices to return:


	n > 0: at most n subslices; the last subslice will be the unsplit remainder. // 至多 n 个元素; 其中最后一个元素不会再分割.
	n == 0: the result is nil (zero subslices) // 返回nil
	n < 0: all subslices // 所有的子切片

SplitAfterN 从 s 中的每个 sep 之后的位置分割 s 并返回分割结果. 如sep为空,SplitAfterN会按每个utf8编码切分. 参数n决定返回的切片的数目:

<a id="example_SplitAfterN">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%q\n", bytes.SplitAfterN([]byte("a,b,c"), []byte(","), 2))
}
```

output:
```txt
["a," "b,c"]
```
## <a id="SplitN">func</a> [SplitN](https://golang.org/src/bytes/bytes.go?s=6592:6634#L270)
<pre>func SplitN(s, sep []<a href="/pkg/builtin/#byte">byte</a>, n <a href="/pkg/builtin/#int">int</a>) [][]<a href="/pkg/builtin/#byte">byte</a></pre>
SplitN slices s into subslices separated by sep and returns a slice of
the subslices between those separators.
If sep is empty, SplitN splits after each UTF-8 sequence.
The count determines the number of subslices to return:


	n > 0: at most n subslices; the last subslice will be the unsplit remainder. 至多 n 个元素; 其中最后一个元素不会再分割.
	n == 0: the result is nil (zero subslices) // 返回nil
	n < 0: all subslices // 所有的子切片

SplitN 根据sep将s分割并返回分割结果. 如sep为空,SplitN会按每个utf8编码切分. 参数n决定返回的切片的数目:

<a id="example_SplitN">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%q\n", bytes.SplitN([]byte("a,b,c"), []byte(","), 2))
	z := bytes.SplitN([]byte("a,b,c"), []byte(","), 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)
}
```

output:
```txt
["a" "b,c"]
[] (nil = true)
```
## <a id="Title">func</a> [Title](https://golang.org/src/bytes/bytes.go?s=17058:17085#L643)
<pre>func Title(s []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
Title treats s as UTF-8-encoded bytes and returns a copy with all Unicode letters that begin
words mapped to their title case.

BUG(rsc): The rule Title uses for word boundaries does not handle Unicode punctuation properly.

Title 将s视为UTF-8编码的字节序列, 它会返回s的一个副本，并把s中每个单词的首字母改为其Unicode的Title(大写)形式.

BUG(rsc): Title方法中用于分词的规则不能正确地处理Unicode标点符号.


<a id="example_Title">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s", bytes.Title([]byte("her royal highness")))
}
```

output:
```txt
Her Royal Highness
```
## <a id="ToLower">func</a> [ToLower](https://golang.org/src/bytes/bytes.go?s=13971:14000#L536)
<pre>func ToLower(s []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
ToLower returns a copy of the byte slice s with all Unicode letters mapped to
their lower case.

ToLower 返回s的一个副本，并把其中所有的Unicode字母转换为小写.

<a id="example_ToLower">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s", bytes.ToLower([]byte("Gopher")))
}
```

output:
```txt
gopher
```
## <a id="ToLowerSpecial">func</a> [ToLowerSpecial](https://golang.org/src/bytes/bytes.go?s=15115:15174#L575)
<pre>func ToLowerSpecial(c <a href="/pkg/unicode/">unicode</a>.<a href="/pkg/unicode/#SpecialCase">SpecialCase</a>, s []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
ToLowerSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their
lower case, giving priority to the special casing rules.

ToLowerSpecial 将s视为UTF-8编码的字节序列并返回s的一个副本, 已把其中所有的Unicode字母转换为小写,且优先使用指定的特殊转换规则.

<a id="example_ToLowerSpecial">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	str := []byte("AHOJ VÝVOJÁRİ GOLANG")
	totitle := bytes.ToLowerSpecial(unicode.AzeriCase, str)
	fmt.Println("Original : " + string(str))
	fmt.Println("ToLower : " + string(totitle))
}
```

output:
```txt
Original : AHOJ VÝVOJÁRİ GOLANG
ToLower : ahoj vývojári golang
```
## <a id="ToTitle">func</a> [ToTitle](https://golang.org/src/bytes/bytes.go?s=14610:14639#L565)
<pre>func ToTitle(s []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
ToTitle treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their title case.

ToTitle 将s视为UTF-8编码的字节序列并返回s的一个副本，已把其中所有的Unicode字母转换为title形式.

<a id="example_ToTitle">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", bytes.ToTitle([]byte("loud noises")))
	fmt.Printf("%s\n", bytes.ToTitle([]byte("хлеб")))
}
```

output:
```txt
LOUD NOISES
ХЛЕБ
```
## <a id="ToTitleSpecial">func</a> [ToTitleSpecial](https://golang.org/src/bytes/bytes.go?s=15380:15439#L581)
<pre>func ToTitleSpecial(c <a href="/pkg/unicode/">unicode</a>.<a href="/pkg/unicode/#SpecialCase">SpecialCase</a>, s []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
ToTitleSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their
title case, giving priority to the special casing rules.

ToTitleSpecial 将s视为UTF-8编码的字节序列并返回s的一个副本, 已把其中所有的Unicode字母转换为Title形式,且优先使用指定的特殊转换规则.

<a id="example_ToTitleSpecial">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	str := []byte("ahoj vývojári golang")
	totitle := bytes.ToTitleSpecial(unicode.AzeriCase, str)
	fmt.Println("Original : " + string(str))
	fmt.Println("ToTitle : " + string(totitle))
}
```

output:
```txt
Original : ahoj vývojári golang
ToTitle : AHOJ VÝVOJÁRİ GOLANG
```
## <a id="ToUpper">func</a> [ToUpper](https://golang.org/src/bytes/bytes.go?s=13323:13352#L505)
<pre>func ToUpper(s []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
ToUpper returns a copy of the byte slice s with all Unicode letters mapped to
their upper case.

ToUpper 返回s的一个副本，并把其中所有的Unicode字母转换为大写.

<a id="example_ToUpper">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s", bytes.ToUpper([]byte("Gopher")))
}
```

output:
```txt
GOPHER
```
## <a id="ToUpperSpecial">func</a> [ToUpperSpecial](https://golang.org/src/bytes/bytes.go?s=14850:14909#L569)
<pre>func ToUpperSpecial(c <a href="/pkg/unicode/">unicode</a>.<a href="/pkg/unicode/#SpecialCase">SpecialCase</a>, s []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
ToUpperSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their
upper case, giving priority to the special casing rules.

ToUpperSpecial 将s视为UTF-8编码的字节序列并返回s的一个副本, 已把其中所有的Unicode字母转换为大写形式,且优先使用指定的特殊转换规则.

<a id="example_ToUpperSpecial">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	str := []byte("ahoj vývojári golang")
	totitle := bytes.ToUpperSpecial(unicode.AzeriCase, str)
	fmt.Println("Original : " + string(str))
	fmt.Println("ToUpper : " + string(totitle))
}

```

output:
```txt
Original : ahoj vývojári golang
ToUpper : AHOJ VÝVOJÁRİ GOLANG
```
## <a id="ToValidUTF8">func</a> [ToValidUTF8](https://golang.org/src/bytes/bytes.go?s=15650:15696#L587)
<pre>func ToValidUTF8(s, replacement []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
ToValidUTF8 treats s as UTF-8-encoded bytes and returns a copy with each run of bytes
representing invalid UTF-8 replaced with the bytes in replacement, which may be empty.

ToValidUTF8 将s视为UTF-8编码的字节序列并返回一个副本，其中已将每个无效的UTF-8替换为replacement，replacement可以为空.

## <a id="Trim">func</a> [Trim](https://golang.org/src/bytes/bytes.go?s=21520:21561#L806)
<pre>func Trim(s []<a href="/pkg/builtin/#byte">byte</a>, cutset <a href="/pkg/builtin/#string">string</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
Trim returns a subslice of s by slicing off all leading and
trailing UTF-8-encoded code points contained in cutset.

Trim 返回将s前后端所有cutset中包含的utf8码点都已去掉的子切片.

<a id="example_Trim">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("[%q]", bytes.Trim([]byte(" !!! Achtung! Achtung! !!! "), "! "))
}
```

output:
```txt
["Achtung! Achtung"]
```
## <a id="TrimFunc">func</a> [TrimFunc](https://golang.org/src/bytes/bytes.go?s=18137:18188#L685)
<pre>func TrimFunc(s []<a href="/pkg/builtin/#byte">byte</a>, f func(r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
TrimFunc returns a subslice of s by slicing off all leading and trailing
UTF-8-encoded code points c that satisfy f(c).

TrimFunc 返回将s前后端中所有满足f(c)的utf8码点都去掉的子切片.

<a id="example_TrimFunc">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(string(bytes.TrimFunc([]byte("go-gopher!"), unicode.IsLetter)))
	fmt.Println(string(bytes.TrimFunc([]byte("\"go-gopher!\""), unicode.IsLetter)))
	fmt.Println(string(bytes.TrimFunc([]byte("go-gopher!"), unicode.IsPunct)))
	fmt.Println(string(bytes.TrimFunc([]byte("1234go-gopher!567"), unicode.IsNumber)))
}
```

output:
```txt
-gopher!
"go-gopher!"
go-gopher
go-gopher!
```
## <a id="TrimLeft">func</a> [TrimLeft](https://golang.org/src/bytes/bytes.go?s=21724:21769#L812)
<pre>func TrimLeft(s []<a href="/pkg/builtin/#byte">byte</a>, cutset <a href="/pkg/builtin/#string">string</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
TrimLeft returns a subslice of s by slicing off all leading
UTF-8-encoded code points contained in cutset.

TrimLeft 返回将s前端中所有cutset包含的utf8码点都已去掉的子切片.

<a id="example_TrimLeft">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Print(string(bytes.TrimLeft([]byte("453gopher8257"), "0123456789")))
}
```

output:
```txt
gopher8257
```
## <a id="TrimLeftFunc">func</a> [TrimLeftFunc](https://golang.org/src/bytes/bytes.go?s=17545:17600#L662)
<pre>func TrimLeftFunc(s []<a href="/pkg/builtin/#byte">byte</a>, f func(r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
TrimLeftFunc treats s as UTF-8-encoded bytes and returns a subslice of s by slicing off
all leading UTF-8-encoded code points c that satisfy f(c).

TrimFunc 返回将s前端中所有满足f(c)的utf8码点都去掉的子切片.

<a id="example_TrimLeftFunc">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(string(bytes.TrimLeftFunc([]byte("go-gopher"), unicode.IsLetter)))
	fmt.Println(string(bytes.TrimLeftFunc([]byte("go-gopher!"), unicode.IsPunct)))
	fmt.Println(string(bytes.TrimLeftFunc([]byte("1234go-gopher!567"), unicode.IsNumber)))
}
```

output:
```txt
-gopher
go-gopher!
go-gopher!567
```
## <a id="TrimPrefix">func</a> [TrimPrefix](https://golang.org/src/bytes/bytes.go?s=18367:18407#L691)
<pre>func TrimPrefix(s, prefix []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
TrimPrefix returns s without the provided leading prefix string.
If s doesn't start with prefix, s is returned unchanged.

TrimPrefix 删除s头部的前缀prefix. 如果不存在前缀prefix则返回原始的s.

<a id="example_TrimPrefix">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b = []byte("Goodbye,, world!")
	b = bytes.TrimPrefix(b, []byte("Goodbye,"))
	b = bytes.TrimPrefix(b, []byte("See ya,"))
	fmt.Printf("Hello%s", b)
}
```

output:
```txt
Hello, world!
```
## <a id="TrimRight">func</a> [TrimRight](https://golang.org/src/bytes/bytes.go?s=21947:21993#L818)
<pre>func TrimRight(s []<a href="/pkg/builtin/#byte">byte</a>, cutset <a href="/pkg/builtin/#string">string</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
TrimRight returns a subslice of s by slicing off all trailing
UTF-8-encoded code points that are contained in cutset.

TrimRight 返回将s后端中所有cutset包含的utf8码点都已去掉的子切片.

<a id="example_TrimRight">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Print(string(bytes.TrimRight([]byte("453gopher8257"), "0123456789")))
}
```

output:
```txt
453gopher
```
## <a id="TrimRightFunc">func</a> [TrimRightFunc](https://golang.org/src/bytes/bytes.go?s=17798:17854#L672)
<pre>func TrimRightFunc(s []<a href="/pkg/builtin/#byte">byte</a>, f func(r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
TrimRightFunc returns a subslice of s by slicing off all trailing
UTF-8-encoded code points c that satisfy f(c).

TrimRightFunc 返回将s后端中所有满足f(c)的utf8码点都去掉的子切片.

<a id="example_TrimRightFunc">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(string(bytes.TrimRightFunc([]byte("go-gopher"), unicode.IsLetter)))
	fmt.Println(string(bytes.TrimRightFunc([]byte("go-gopher!"), unicode.IsPunct)))
	fmt.Println(string(bytes.TrimRightFunc([]byte("1234go-gopher!567"), unicode.IsNumber)))
}
```

output:
```txt
go-
go-gopher
1234go-gopher!
```
## <a id="TrimSpace">func</a> [TrimSpace](https://golang.org/src/bytes/bytes.go?s=22164:22195#L824)
<pre>func TrimSpace(s []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
TrimSpace returns a subslice of s by slicing off all leading and
trailing white space, as defined by Unicode.

TrimSpace 删除s前后端所有unicode中定义的空白符.

<a id="example_TrimSpace">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s", bytes.TrimSpace([]byte(" \t\n a lone gopher \n\t\r\n")))
}
```

output:
```txt
a lone gopher
```
## <a id="TrimSuffix">func</a> [TrimSuffix](https://golang.org/src/bytes/bytes.go?s=18605:18645#L700)
<pre>func TrimSuffix(s, suffix []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
TrimSuffix returns s without the provided trailing suffix string.
If s doesn't end with suffix, s is returned unchanged.

TrimPrefix 删除s尾部的后缀suffix. 如果不存在后缀suffix则返回原始的s.

<a id="example_TrimSuffix">Example</a>
```go
package main

import (
	"bytes"
	"os"
)

func main() {
	var b = []byte("Hello, goodbye, etc!")
	b = bytes.TrimSuffix(b, []byte("goodbye, etc!"))
	b = bytes.TrimSuffix(b, []byte("gopher"))
	b = append(b, bytes.TrimSuffix([]byte("world!"), []byte("x!"))...)
	os.Stdout.Write(b)
}
```

output:
```txt
Hello, world!
```


## <a id="Buffer">type</a> [Buffer](https://golang.org/src/bytes/buffer.go?s=492:717#L10)
A Buffer is a variable-sized buffer of bytes with Read and Write methods.
The zero value for Buffer is an empty buffer ready to use.

Buffer 是一个实现了Read和Write方法的带动态大小的字节buffer的缓冲器. 它的零值的缓冲区大小为0.

<pre>type Buffer struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Buffer">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer // A Buffer needs no initialization. // Buffer不需要初始化.
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "world!")
	b.WriteTo(os.Stdout)
}
```

output:
```txt
Hello world!
```
<a id="example_Buffer_reader">Example (Reader)</a>

```go
package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"os"
)

func main() {
    // A Buffer can turn a string or a []byte into an io.Reader.
    // Buffer能将字符串或[]byte转化为io.Reader.
	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder(base64.StdEncoding, buf)
	io.Copy(os.Stdout, dec)
}
```

output:
```txt
Gophers rule!
```


### <a id="NewBuffer">func</a> [NewBuffer](https://golang.org/src/bytes/buffer.go?s=14279:14313#L440)
<pre>func NewBuffer(buf []<a href="/pkg/builtin/#byte">byte</a>) *<a href="#Buffer">Buffer</a></pre>
NewBuffer creates and initializes a new Buffer using buf as its
initial contents. The new Buffer takes ownership of buf, and the
caller should not use buf after this call. NewBuffer is intended to
prepare a Buffer to read existing data. It can also be used to set
the initial size of the internal buffer for writing. To do that,
buf should have the desired capacity but a length of zero.

In most cases, new(Buffer) (or just declaring a Buffer variable) is
sufficient to initialize a Buffer.

NewBuffer 使用buf作为初始内容来创建并初始化一个Buffer. 新的Buffer拥有buf的所有权,使用本方法后不应该再使用buf. NewBuffer旨在准备一个缓冲区用于读取现有数据(即参数buf). 但它也可指定内部缓冲区的大小用于写, 如果要这么做，buf应是一个指定cap且长度为零的切片.

大多数情况下, new(Buffer)(或只是声明一个Buffer变量)就足以初始化一个Buffer了.


### <a id="NewBufferString">func</a> [NewBufferString](https://golang.org/src/bytes/buffer.go?s=14621:14659#L448)
<pre>func NewBufferString(s <a href="/pkg/builtin/#string">string</a>) *<a href="#Buffer">Buffer</a></pre>
NewBufferString creates and initializes a new Buffer using string s as its
initial contents. It is intended to prepare a buffer to read an existing
string.

In most cases, new(Buffer) (or just declaring a Buffer variable) is
sufficient to initialize a Buffer.

NewBufferString 使用s作为初始内容来创建并初始化一个Buffer. 它旨在准备一个缓冲区用于读取现有数据(即参数s).

大多数情况下, new(Buffer)(或只是声明一个Buffer变量)就足以初始化一个Buffer了.




### <a id="Buffer.Bytes">func</a> (\*Buffer) [Bytes](https://golang.org/src/bytes/buffer.go?s=2118:2149#L44)
<pre>func (b *<a href="#Buffer">Buffer</a>) Bytes() []<a href="/pkg/builtin/#byte">byte</a></pre>
Bytes returns a slice of length b.Len() holding the unread portion of the buffer.
The slice is valid for use only until the next buffer modification (that is,
only until the next call to a method like Read, Write, Reset, or Truncate).
The slice aliases the buffer content at least until the next buffer modification,
so immediate changes to the slice will affect the result of future reads.

Bytes b中未读数据的字节slice, 长度为b.Len(). 该切片只在下一次buffer改变前有效(即只有在其下一次调用Read, Write, Reset或Truncate方法前有效). 在下一次buffer改变前, 该切片是缓存内容的别名, 因此修改切片(的内容)将影响到未来的读取结果.


### <a id="Buffer.Cap">func</a> (\*Buffer) [Cap](https://golang.org/src/bytes/buffer.go?s=2943:2969#L67)
<pre>func (b *<a href="#Buffer">Buffer</a>) Cap() <a href="/pkg/builtin/#int">int</a></pre>
Cap returns the capacity of the buffer's underlying byte slice, that is, the
total space allocated for the buffer's data.

Cap 返回底层`[]byte`的cap. 它是分配给缓存的全部空间大小.


### <a id="Buffer.Grow">func</a> (\*Buffer) [Grow](https://golang.org/src/bytes/buffer.go?s=5301:5329#L147)
<pre>func (b *<a href="#Buffer">Buffer</a>) Grow(n <a href="/pkg/builtin/#int">int</a>)</pre>
Grow grows the buffer's capacity, if necessary, to guarantee space for
another n bytes. After Grow(n), at least n bytes can be written to the
buffer without another allocation.
If n is negative, Grow will panic.
If the buffer can't grow it will panic with ErrTooLarge.

Grow 如有必要会增加b中缓冲的容量, 确保有n个字节的剩余空间. 调用Grow(n)后,至少可以向缓冲中写入n字节数据且无需申请分配空间. 如果n是负数,Grow会panic.如果缓冲无法扩容,会返回ErrTooLarge.

<a id="example_Buffer_Grow">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	b.Grow(64)
	bb := b.Bytes()
	b.Write([]byte("64 bytes or fewer"))
	fmt.Printf("%q", bb[:b.Len()])
}
```

output:
```txt
"64 bytes or fewer"
```

### <a id="Buffer.Len">func</a> (\*Buffer) [Len](https://golang.org/src/bytes/buffer.go?s=2757:2783#L63)
<pre>func (b *<a href="#Buffer">Buffer</a>) Len() <a href="/pkg/builtin/#int">int</a></pre>
Len returns the number of bytes of the unread portion of the buffer;
b.Len() == len(b.Bytes()).

Len 返回缓冲中未读取部分的字节长度: b.Len() == len(b.Bytes()).

<a id="example_Buffer_Len">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	b.Grow(64)
	b.Write([]byte("abcde"))
	fmt.Printf("%d", b.Len())
}
```

output:
```txt
5
```

### <a id="Buffer.Next">func</a> (\*Buffer) [Next](https://golang.org/src/bytes/buffer.go?s=10051:10086#L308)
<pre>func (b *<a href="#Buffer">Buffer</a>) Next(n <a href="/pkg/builtin/#int">int</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
Next returns a slice containing the next n bytes from the buffer,
advancing the buffer as if the bytes had been returned by Read.
If there are fewer than n bytes in the buffer, Next returns the entire buffer.
The slice is only valid until the next call to a read or write method.

Next 返回一个slice包含了缓冲区中未读的n个字节, 并会步进缓冲区就像这些字节被Read方法读取了一样. 如果缓冲内数据的字节数<=n,它会返回整个缓冲区. 返回的slice只在下次调用read/write方法前有效.

### <a id="Buffer.Read">func</a> (\*Buffer) [Read](https://golang.org/src/bytes/buffer.go?s=9451:9501#L286)
<pre>func (b *<a href="#Buffer">Buffer</a>) Read(p []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Read reads the next len(p) bytes from the buffer or until the buffer
is drained. The return value n is the number of bytes read. If the
buffer has no data to return, err is io.EOF (unless len(p) is zero);
otherwise it is nil.

Read 从缓冲区(未读)中读取长度为len(p)的字节或直到缓冲区被读完. 返回值n是读取到的字节数. 如果缓冲区中没有数据可以返回,返回错误是io.EOF(除非len(p)是0);否则返回值err是nil.



### <a id="Buffer.ReadByte">func</a> (\*Buffer) [ReadByte](https://golang.org/src/bytes/buffer.go?s=10361:10402#L324)
<pre>func (b *<a href="#Buffer">Buffer</a>) ReadByte() (<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadByte reads and returns the next byte from the buffer.
If no byte is available, it returns error io.EOF.

ReadByte 从b中读取后续的一个字节. 如果没有数据可用,返回错误io.EOF.


### <a id="Buffer.ReadBytes">func</a> (\*Buffer) [ReadBytes](https://golang.org/src/bytes/buffer.go?s=12663:12726#L398)
<pre>func (b *<a href="#Buffer">Buffer</a>) ReadBytes(delim <a href="/pkg/builtin/#byte">byte</a>) (line []<a href="/pkg/builtin/#byte">byte</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadBytes reads until the first occurrence of delim in the input,
returning a slice containing the data up to and including the delimiter.
If ReadBytes encounters an error before finding a delimiter,
it returns the data read before the error and the error itself (often io.EOF).
ReadBytes returns err != nil if and only if the returned data does not end in
delim.

ReadBytes 读取直到第一次遇到delim, 返回一个包含已读取的数据和delim的切片. 如果ReadBytes在遇到delim之前就碰到了错误,它会返回在遇到错误之前已读取的数据和该错误(通常是io.EOF). 当且仅当ReadBytes方法返回的数据不以delim结尾时，返回的err != nil.


### <a id="Buffer.ReadFrom">func</a> (\*Buffer) [ReadFrom](https://golang.org/src/bytes/buffer.go?s=6793:6852#L189)
<pre>func (b *<a href="#Buffer">Buffer</a>) ReadFrom(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (n <a href="/pkg/builtin/#int64">int64</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadFrom reads data from r until EOF and appends it to the buffer, growing
the buffer as needed. The return value n is the number of bytes read. Any
error except io.EOF encountered during the read is also returned. If the
buffer becomes too large, ReadFrom will panic with ErrTooLarge.

ReadFrom 即读取直至遇到EOF(即将r中的数据全部读取到缓冲中), 如有必要会扩充缓冲. n为b读取到的字节数. 读取过程中io.EOF外的任何错误都会被返回. 如果缓冲变得太大, ReadFrom会panic并返回ErrTooLarge. 


### <a id="Buffer.ReadRune">func</a> (\*Buffer) [ReadRune](https://golang.org/src/bytes/buffer.go?s=10816:10873#L341)
<pre>func (b *<a href="#Buffer">Buffer</a>) ReadRune() (r <a href="/pkg/builtin/#rune">rune</a>, size <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadRune reads and returns the next UTF-8-encoded
Unicode code point from the buffer.
If no bytes are available, the error returned is io.EOF.
If the bytes are an erroneous UTF-8 encoding, it
consumes one byte and returns U+FFFD, 1.

ReadRune 从b中读取一个UTF-8编码的unicode码点. 如果没有数据可读, err会是io.EOF. 如果编码是无效的，它(读取位置)步进一个字节并且返回`U+FFFD(unicode.ReplacementChar),1(出参szie)`. 


### <a id="Buffer.ReadString">func</a> (\*Buffer) [ReadString](https://golang.org/src/bytes/buffer.go?s=13630:13694#L426)
<pre>func (b *<a href="#Buffer">Buffer</a>) ReadString(delim <a href="/pkg/builtin/#byte">byte</a>) (line <a href="/pkg/builtin/#string">string</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadString reads until the first occurrence of delim in the input,
returning a string containing the data up to and including the delimiter.
If ReadString encounters an error before finding a delimiter,
it returns the data read before the error and the error itself (often io.EOF).
ReadString returns err != nil if and only if the returned data does not end
in delim.

ReadString 会读取数据直到第一次遇到分隔符delim，返回一个包含已读取到数据的字符串(包括分隔符). 如果ReadString在遇到分隔符之前碰到了error, 它会返回遇到错误之前已读取的数据和这个错误(经常是io.EOF). 当且仅当ReadString方法返回的数据不以delim结尾时，返回的err != nil.


### <a id="Buffer.Reset">func</a> (\*Buffer) [Reset](https://golang.org/src/bytes/buffer.go?s=3534:3558#L87)
<pre>func (b *<a href="#Buffer">Buffer</a>) Reset()</pre>
Reset resets the buffer to be empty,
but it retains the underlying storage for use by future writes.
Reset is the same as Truncate(0).

Reset 重置缓冲,但会保留底层的数据存储空间用于将来的写操作. 其等价于Truncate(0).


### <a id="Buffer.String">func</a> (\*Buffer) [String](https://golang.org/src/bytes/buffer.go?s=2382:2414#L50)
<pre>func (b *<a href="#Buffer">Buffer</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns the contents of the unread portion of the buffer
as a string. If the Buffer is a nil pointer, it returns "<nil>".

To build strings more efficiently, see the strings.Builder type.

String 将缓冲中未读取的内容以字符串形式返回. 当b为nil时,返回"<nil>".

为了更有效地构建字符串可以使用 strings.Builder 类型.


### <a id="Buffer.Truncate">func</a> (\*Buffer) [Truncate](https://golang.org/src/bytes/buffer.go?s=3187:3219#L72)
<pre>func (b *<a href="#Buffer">Buffer</a>) Truncate(n <a href="/pkg/builtin/#int">int</a>)</pre>
Truncate discards all but the first n unread bytes from the buffer
but continues to use the same allocated storage.
It panics if n is negative or greater than the length of the buffer.

Truncate 只保留前n个未读字节的缓冲数据并清除其余数据, 且继续使用原有的缓冲. 如果n是负数或者大于缓冲长度将会panic.


### <a id="Buffer.UnreadByte">func</a> (\*Buffer) [UnreadByte](https://golang.org/src/bytes/buffer.go?s=12119:12154#L381)
<pre>func (b *<a href="#Buffer">Buffer</a>) UnreadByte() <a href="/pkg/builtin/#error">error</a></pre>
UnreadByte unreads the last byte returned by the most recent successful
read operation that read at least one byte. If a write has happened since
the last read, if the last read returned an error, or if the read read zero
bytes, UnreadByte returns an error.

UnreadByte 撤消最近读取成功(即至少读到一个字节)的最后一个字节. 如果最后一次读取后进行了写入操作,或最后一次读取返回了一个error或该操作读取到零个字节, 那么本方法都会返回一个error.


### <a id="Buffer.UnreadRune">func</a> (\*Buffer) [UnreadRune](https://golang.org/src/bytes/buffer.go?s=11474:11509#L364)
<pre>func (b *<a href="#Buffer">Buffer</a>) UnreadRune() <a href="/pkg/builtin/#error">error</a></pre>
UnreadRune unreads the last rune returned by ReadRune.
If the most recent read or write operation on the buffer was
not a successful ReadRune, UnreadRune returns an error.  (In this regard
it is stricter than UnreadByte, which will unread the last byte
from any read operation.)

UnreadRune 撤消最近ReadRune成功读取的一个rune. 如果最近一次读/写操作不是一次成功的ReadRune操作，它会返回一个error. (在这方面,它比UnreadByte(可以撤销任何读取操作的最后一个字节)严格.)


### <a id="Buffer.Write">func</a> (\*Buffer) [Write](https://golang.org/src/bytes/buffer.go?s=5642:5693#L158)
<pre>func (b *<a href="#Buffer">Buffer</a>) Write(p []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Write appends the contents of p to the buffer, growing the buffer as
needed. The return value n is the length of p; err is always nil. If the
buffer becomes too large, Write will panic with ErrTooLarge.

Write 将p的内容追加到缓冲中, 如有必要会扩充缓冲. 返回值n为len(p), 通常err为nil. 如果缓冲变得太大, 它会panic(ErrTooLarge).


### <a id="Buffer.WriteByte">func</a> (\*Buffer) [WriteByte](https://golang.org/src/bytes/buffer.go?s=8453:8493#L253)
<pre>func (b *<a href="#Buffer">Buffer</a>) WriteByte(c <a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
WriteByte appends the byte c to the buffer, growing the buffer as needed.
The returned error is always nil, but is included to match bufio.Writer's
WriteByte. If the buffer becomes too large, WriteByte will panic with
ErrTooLarge.

WriteByte 将c追加到缓冲中, 如有必要会扩充缓冲. 返回值永远都是nil, 目的是为了和bufio.Writer的WriteByte方法保持一致. 如果缓冲变得太大, 它会panic(ErrTooLarge).


### <a id="Buffer.WriteRune">func</a> (\*Buffer) [WriteRune](https://golang.org/src/bytes/buffer.go?s=8899:8952#L267)
<pre>func (b *<a href="#Buffer">Buffer</a>) WriteRune(r <a href="/pkg/builtin/#rune">rune</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
WriteRune appends the UTF-8 encoding of Unicode code point r to the
buffer, returning its length and an error, which is always nil but is
included to match bufio.Writer's WriteRune. The buffer is grown as needed;
if it becomes too large, WriteRune will panic with ErrTooLarge.

WriteRune 将utf8编码的unicode码值r追加到缓冲中, 如有必要会扩充缓冲, 并返回r的字节数和一个error. err永远都是nil, 目的是为了和bufio.Writer的WriteRune方法保持一致. 如果缓冲变得太大,它会panic(ErrTooLarge).


### <a id="Buffer.WriteString">func</a> (\*Buffer) [WriteString](https://golang.org/src/bytes/buffer.go?s=6050:6107#L170)
<pre>func (b *<a href="#Buffer">Buffer</a>) WriteString(s <a href="/pkg/builtin/#string">string</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
WriteString appends the contents of s to the buffer, growing the buffer as
needed. The return value n is the length of s; err is always nil. If the
buffer becomes too large, WriteString will panic with ErrTooLarge.

WriteString 将s的内容追加到缓冲中, 如有必要会扩充缓冲. 返回值n为len(p), err永远都是nil. 如果缓冲变得太大,它会panic(ErrTooLarge).


### <a id="Buffer.WriteTo">func</a> (\*Buffer) [WriteTo](https://golang.org/src/bytes/buffer.go?s=7711:7769#L226)
<pre>func (b *<a href="#Buffer">Buffer</a>) WriteTo(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) (n <a href="/pkg/builtin/#int64">int64</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
WriteTo writes data to w until the buffer is drained or an error occurs.
The return value n is the number of bytes written; it always fits into an
int, but it is int64 to match the io.WriterTo interface. Any error
encountered during the write is also returned.

WriteTo 将b的数据写入w直至没有数据或发生错误. 返回值n为写入的字节数; 它通常用int就可以存储, 但为了和io.WriterTo接口统一而使用int64类型. 在写时发生的任何错误都会被返回.


## <a id="Reader">type</a> [Reader](https://golang.org/src/bytes/reader.go?s=511:641#L8)
A Reader implements the io.Reader, io.ReaderAt, io.WriterTo, io.Seeker,
io.ByteScanner, and io.RuneScanner interfaces by reading from
a byte slice.
Unlike a Buffer, a Reader is read-only and supports seeking.
The zero value for Reader operates like a Reader of an empty slice.

Reader 从[]byte读取数据的读取器, 其实现了io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner和io.RuneScanner接口. Reader不像Buffer类型,它是只读的且支持Seek方法. Reader的零值的操作类似于使用空切片的Reader.

<pre>type Reader struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewReader">func</a> [NewReader](https://golang.org/src/bytes/reader.go?s=3877:3909#L150)
<pre>func NewReader(b []<a href="/pkg/builtin/#byte">byte</a>) *<a href="#Reader">Reader</a></pre>
NewReader returns a new Reader reading from b.

NewReader 创建一个从b中读取数据的Reader.




### <a id="Reader.Len">func</a> (\*Reader) [Len](https://golang.org/src/bytes/reader.go?s=717:743#L16)
<pre>func (r *<a href="#Reader">Reader</a>) Len() <a href="/pkg/builtin/#int">int</a></pre>
Len returns the number of bytes of the unread portion of the
slice.

Len 返回数据中未被读取部分的字节数.

<a id="example_Reader_Len">Example</a>
```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.NewReader([]byte("Hi!")).Len())
	fmt.Println(bytes.NewReader([]byte("こんにちは!")).Len())
}
```

output:
```txt
3
16
```

### <a id="Reader.Read">func</a> (\*Reader) [Read](https://golang.org/src/bytes/reader.go?s=1154:1204#L30)
<pre>func (r *<a href="#Reader">Reader</a>) Read(b []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Read implements the io.Reader interface.

Read 实现了io.Reader接口.


### <a id="Reader.ReadAt">func</a> (\*Reader) [ReadAt](https://golang.org/src/bytes/reader.go?s=1375:1438#L41)
<pre>func (r *<a href="#Reader">Reader</a>) ReadAt(b []<a href="/pkg/builtin/#byte">byte</a>, off <a href="/pkg/builtin/#int64">int64</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadAt implements the io.ReaderAt interface.

ReadAt 实现了io.ReaderAt接口.


### <a id="Reader.ReadByte">func</a> (\*Reader) [ReadByte](https://golang.org/src/bytes/reader.go?s=1736:1777#L57)
<pre>func (r *<a href="#Reader">Reader</a>) ReadByte() (<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadByte implements the io.ByteReader interface.

ReadByte 实现了io.ReadByte接口.


### <a id="Reader.ReadRune">func</a> (\*Reader) [ReadRune](https://golang.org/src/bytes/reader.go?s=2186:2244#L78)
<pre>func (r *<a href="#Reader">Reader</a>) ReadRune() (ch <a href="/pkg/builtin/#rune">rune</a>, size <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
ReadRune implements the io.RuneReader interface.

ReadRune 实现了io.ReadRune接口.


### <a id="Reader.Reset">func</a> (\*Reader) [Reset](https://golang.org/src/bytes/reader.go?s=3767:3799#L147)
<pre>func (r *<a href="#Reader">Reader</a>) Reset(b []<a href="/pkg/builtin/#byte">byte</a>)</pre>
Reset resets the Reader to be reading from b.

Reset 将Reader重置为从b中读取数据.


### <a id="Reader.Seek">func</a> (\*Reader) [Seek](https://golang.org/src/bytes/reader.go?s=2903:2965#L107)
<pre>func (r *<a href="#Reader">Reader</a>) Seek(offset <a href="/pkg/builtin/#int64">int64</a>, whence <a href="/pkg/builtin/#int">int</a>) (<a href="/pkg/builtin/#int64">int64</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Seek implements the io.Seeker interface.

Seek 实现了io.Seek接口.


### <a id="Reader.Size">func</a> (\*Reader) [Size](https://golang.org/src/bytes/reader.go?s=1052:1081#L27)
<pre>func (r *<a href="#Reader">Reader</a>) Size() <a href="/pkg/builtin/#int64">int64</a></pre>
Size returns the original length of the underlying byte slice.
Size is the number of bytes available for reading via ReadAt.
The returned value is always the same and is not affected by calls
to any other method.

Size 返回原始底层`[]byte`的长度. Size 是 通过 ReadAt 能够读取到的最大有效字节数. 它的返回值是固定的不受其他方法影响.


### <a id="Reader.UnreadByte">func</a> (\*Reader) [UnreadByte](https://golang.org/src/bytes/reader.go?s=1969:2004#L68)
<pre>func (r *<a href="#Reader">Reader</a>) UnreadByte() <a href="/pkg/builtin/#error">error</a></pre>
UnreadByte complements ReadByte in implementing the io.ByteScanner interface.

UnreadByte 是对实现了io.ByteScanner接口的ReadByte的补充.


### <a id="Reader.UnreadRune">func</a> (\*Reader) [UnreadRune](https://golang.org/src/bytes/reader.go?s=2568:2603#L94)
<pre>func (r *<a href="#Reader">Reader</a>) UnreadRune() <a href="/pkg/builtin/#error">error</a></pre>
UnreadRune complements ReadRune in implementing the io.RuneScanner interface.

UnreadRune 是对实现了io.RuneScanner接口的ReadRune的补充.


### <a id="Reader.WriteTo">func</a> (\*Reader) [WriteTo](https://golang.org/src/bytes/reader.go?s=3379:3437#L128)
<pre>func (r *<a href="#Reader">Reader</a>) WriteTo(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) (n <a href="/pkg/builtin/#int64">int64</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
WriteTo implements the io.WriterTo interface.


WriteTo 实现了io.WriteTo接口.





