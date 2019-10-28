

# strings
`import "strings"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package strings implements simple functions to manipulate UTF-8 encoded strings.

For information about UTF-8 strings in Go, see <a href="https://blog.golang.org/strings">https://blog.golang.org/strings</a>.

strings 实现简单的函数来操纵 UTF-8 编码的字符串.

有关 Go 中 UTF-8 字符串的信息，请参阅https://blog.golang.org/strings.


## <a id="pkg-index">Index</a>
* [func Compare(a, b string) int](#Compare)
* [func Contains(s, substr string) bool](#Contains)
* [func ContainsAny(s, chars string) bool](#ContainsAny)
* [func ContainsRune(s string, r rune) bool](#ContainsRune)
* [func Count(s, substr string) int](#Count)
* [func EqualFold(s, t string) bool](#EqualFold)
* [func Fields(s string) []string](#Fields)
* [func FieldsFunc(s string, f func(rune) bool) []string](#FieldsFunc)
* [func HasPrefix(s, prefix string) bool](#HasPrefix)
* [func HasSuffix(s, suffix string) bool](#HasSuffix)
* [func Index(s, substr string) int](#Index)
* [func IndexAny(s, chars string) int](#IndexAny)
* [func IndexByte(s string, c byte) int](#IndexByte)
* [func IndexFunc(s string, f func(rune) bool) int](#IndexFunc)
* [func IndexRune(s string, r rune) int](#IndexRune)
* [func Join(a []string, sep string) string](#Join)
* [func LastIndex(s, substr string) int](#LastIndex)
* [func LastIndexAny(s, chars string) int](#LastIndexAny)
* [func LastIndexByte(s string, c byte) int](#LastIndexByte)
* [func LastIndexFunc(s string, f func(rune) bool) int](#LastIndexFunc)
* [func Map(mapping func(rune) rune, s string) string](#Map)
* [func Repeat(s string, count int) string](#Repeat)
* [func Replace(s, old, new string, n int) string](#Replace)
* [func ReplaceAll(s, old, new string) string](#ReplaceAll)
* [func Split(s, sep string) []string](#Split)
* [func SplitAfter(s, sep string) []string](#SplitAfter)
* [func SplitAfterN(s, sep string, n int) []string](#SplitAfterN)
* [func SplitN(s, sep string, n int) []string](#SplitN)
* [func Title(s string) string](#Title)
* [func ToLower(s string) string](#ToLower)
* [func ToLowerSpecial(c unicode.SpecialCase, s string) string](#ToLowerSpecial)
* [func ToTitle(s string) string](#ToTitle)
* [func ToTitleSpecial(c unicode.SpecialCase, s string) string](#ToTitleSpecial)
* [func ToUpper(s string) string](#ToUpper)
* [func ToUpperSpecial(c unicode.SpecialCase, s string) string](#ToUpperSpecial)
* [func ToValidUTF8(s, replacement string) string](#ToValidUTF8)
* [func Trim(s string, cutset string) string](#Trim)
* [func TrimFunc(s string, f func(rune) bool) string](#TrimFunc)
* [func TrimLeft(s string, cutset string) string](#TrimLeft)
* [func TrimLeftFunc(s string, f func(rune) bool) string](#TrimLeftFunc)
* [func TrimPrefix(s, prefix string) string](#TrimPrefix)
* [func TrimRight(s string, cutset string) string](#TrimRight)
* [func TrimRightFunc(s string, f func(rune) bool) string](#TrimRightFunc)
* [func TrimSpace(s string) string](#TrimSpace)
* [func TrimSuffix(s, suffix string) string](#TrimSuffix)
* [type Builder](#Builder)
  * [func (b *Builder) Cap() int](#Builder.Cap)
  * [func (b *Builder) Grow(n int)](#Builder.Grow)
  * [func (b *Builder) Len() int](#Builder.Len)
  * [func (b *Builder) Reset()](#Builder.Reset)
  * [func (b *Builder) String() string](#Builder.String)
  * [func (b *Builder) Write(p []byte) (int, error)](#Builder.Write)
  * [func (b *Builder) WriteByte(c byte) error](#Builder.WriteByte)
  * [func (b *Builder) WriteRune(r rune) (int, error)](#Builder.WriteRune)
  * [func (b *Builder) WriteString(s string) (int, error)](#Builder.WriteString)
* [type Reader](#Reader)
  * [func NewReader(s string) *Reader](#NewReader)
  * [func (r *Reader) Len() int](#Reader.Len)
  * [func (r *Reader) Read(b []byte) (n int, err error)](#Reader.Read)
  * [func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)](#Reader.ReadAt)
  * [func (r *Reader) ReadByte() (byte, error)](#Reader.ReadByte)
  * [func (r *Reader) ReadRune() (ch rune, size int, err error)](#Reader.ReadRune)
  * [func (r *Reader) Reset(s string)](#Reader.Reset)
  * [func (r *Reader) Seek(offset int64, whence int) (int64, error)](#Reader.Seek)
  * [func (r *Reader) Size() int64](#Reader.Size)
  * [func (r *Reader) UnreadByte() error](#Reader.UnreadByte)
  * [func (r *Reader) UnreadRune() error](#Reader.UnreadRune)
  * [func (r *Reader) WriteTo(w io.Writer) (n int64, err error)](#Reader.WriteTo)
* [type Replacer](#Replacer)
  * [func NewReplacer(oldnew ...string) *Replacer](#NewReplacer)
  * [func (r *Replacer) Replace(s string) string](#Replacer.Replace)
  * [func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)](#Replacer.WriteString)


#### <a id="pkg-examples">Examples</a>
* [Builder](#example_Builder)
* [Compare](#example_Compare)
* [Contains](#example_Contains)
* [ContainsAny](#example_ContainsAny)
* [ContainsRune](#example_ContainsRune)
* [Count](#example_Count)
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
* [NewReplacer](#example_NewReplacer)
* [Repeat](#example_Repeat)
* [Replace](#example_Replace)
* [ReplaceAll](#example_ReplaceAll)
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
[builder.go](https://golang.org/src/strings/builder.go) [compare.go](https://golang.org/src/strings/compare.go) [reader.go](https://golang.org/src/strings/reader.go) [replace.go](https://golang.org/src/strings/replace.go) [search.go](https://golang.org/src/strings/search.go) [strings.go](https://golang.org/src/strings/strings.go) 






## <a id="Compare">func</a> [Compare](https://golang.org/src/strings/compare.go?s=490:519#L3)
<pre>func Compare(a, b <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#int">int</a></pre>
Compare returns an integer comparing two strings lexicographically.
The result will be 0 if a==b, -1 if a < b, and +1 if a > b.

Compare is included only for symmetry with package bytes.
It is usually clearer and always faster to use the built-in
string comparison operators ==, <, >, and so on.

Compare 根据字典顺序比较两个字符串 a 和 b, 并返回一个整数. a == b 时返回 0, a < b 时返回 -1, a > b 时返回 +1.

Compare 仅是为了和 bytes 包对应, 通常使用内置的比较运算符（==，<，>）更清晰且效率更高.

<a id="example_Compare">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.Compare("a", "b"))
  fmt.Println(strings.Compare("a", "a"))
  fmt.Println(strings.Compare("b", "a"))
}
```

output:
```txt
-1
0
1
```

## <a id="Contains">func</a> [Contains](https://golang.org/src/strings/strings.go?s=2326:2362#L88)
<pre>func Contains(s, substr <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Contains reports whether substr is within s.

Contains 判断 s 是否包含子串 substr.

<a id="example_Contains">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.Contains("seafood", "foo"))
  fmt.Println(strings.Contains("seafood", "bar"))
  fmt.Println(strings.Contains("seafood", ""))
  fmt.Println(strings.Contains("", ""))
}
```

output:
```txt
true
false
true
true
```

## <a id="ContainsAny">func</a> [ContainsAny](https://golang.org/src/strings/strings.go?s=2476:2514#L93)
<pre>func ContainsAny(s, chars <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
ContainsAny reports whether any Unicode code points in chars are within s.

ContainsAny 判断 s 中是否包含 chars 的任意Unicode码点.

<a id="example_ContainsAny">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.ContainsAny("team", "i"))
  fmt.Println(strings.ContainsAny("fail", "ui"))
  fmt.Println(strings.ContainsAny("ure", "ui"))
  fmt.Println(strings.ContainsAny("failure", "ui"))
  fmt.Println(strings.ContainsAny("foo", ""))
  fmt.Println(strings.ContainsAny("", ""))
}
```

output:
```txt
false
true
true
true
false
false
```

## <a id="ContainsRune">func</a> [ContainsRune](https://golang.org/src/strings/strings.go?s=2622:2662#L98)
<pre>func ContainsRune(s <a href="/pkg/builtin/#string">string</a>, r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
ContainsRune reports whether the Unicode code point r is within s.

ContainsRune 判断 s 是否包含 Unicode码点r.

<a id="example_ContainsRune">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  // Finds whether a string contains a particular Unicode code point.
  // The code point for the lowercase letter "a", for example, is 97.
  fmt.Println(strings.ContainsRune("aardvark", 97))
  fmt.Println(strings.ContainsRune("timeout", 97))
}
```

output:
```txt
true
false
```

## <a id="Count">func</a> [Count](https://golang.org/src/strings/strings.go?s=1986:2018#L68)
<pre>func Count(s, substr <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#int">int</a></pre>
Count counts the number of non-overlapping instances of substr in s.
If substr is an empty string, Count returns 1 + the number of Unicode code points in s.

Count 返回在 s 中非重叠出现 substr 的次数. 如果 substr 为空，则返回`s 中的 Unicode码点数量+1`.

<a id="example_Count">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.Count("cheese", "e"))
  fmt.Println(strings.Count("five", "")) // before & after each rune
}
```

output:
```txt
3
5
```

## <a id="EqualFold">func</a> [EqualFold](https://golang.org/src/strings/strings.go?s=24676:24708#L963)
<pre>func EqualFold(s, t <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
EqualFold reports whether s and t, interpreted as UTF-8 strings,
are equal under Unicode case-folding.

EqualFold 判断 s 和 t (两者均被视为UTF-8字符串)在忽略大小写的情况下是否相等.

<a id="example_EqualFold">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.EqualFold("Go", "go"))
}
```

output:
```txt
true
```

## <a id="Fields">func</a> [Fields](https://golang.org/src/strings/strings.go?s=8396:8426#L319)
<pre>func Fields(s <a href="/pkg/builtin/#string">string</a>) []<a href="/pkg/builtin/#string">string</a></pre>
Fields splits the string s around each instance of one or more consecutive white space
characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an
empty slice if s contains only white space.

Fields 根据空白字符（Go 中通过 unicode.IsSpace 判断是否为空白字符）分割字符串并将结果以切片形式返回. 当 s 只包含空白字符时返回值为空的slice.

<a id="example_Fields">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
}
```

output:
```txt
Fields are: ["foo" "bar" "baz"]
```

## <a id="FieldsFunc">func</a> [FieldsFunc](https://golang.org/src/strings/strings.go?s=9792:9845#L373)
<pre>func FieldsFunc(s <a href="/pkg/builtin/#string">string</a>, f func(<a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a>) []<a href="/pkg/builtin/#string">string</a></pre>
FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c)
and returns an array of slices of s. If all code points in s satisfy f(c) or the
string is empty, an empty slice is returned.
FieldsFunc makes no guarantees about the order in which it calls f(c).
If f does not return consistent results for a given c, FieldsFunc may crash.

FieldsFunc 会遍历所有码点判断是否满足f(c),依此将s分割并返回一个slice. 如果s中所有码点都满足f(c)或者len(s)==0,返回空slice. FieldsFunc不保证调用f(c)的顺序. 如果f没有根据给定的c返回一致(即相同)的结果, FieldsFunc可能会crash.

<a id="example_FieldsFunc">Example</a>
```go
package main

import (
  "fmt"
  "strings"
  "unicode"
)

func main() {
  f := func(c rune) bool {
    return !unicode.IsLetter(c) && !unicode.IsNumber(c)
  }
  fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f))
}
```

output:
```txt
Fields are: ["foo1" "bar2" "baz3"]
```

## <a id="HasPrefix">func</a> [HasPrefix](https://golang.org/src/strings/strings.go?s=11153:11190#L438)
<pre>func HasPrefix(s, prefix <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
HasPrefix tests whether the string s begins with prefix.

HasPrefix 判断 s 是否以前缀 prefix 开头.

<a id="example_HasPrefix">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.HasPrefix("Gopher", "Go"))
  fmt.Println(strings.HasPrefix("Gopher", "C"))
  fmt.Println(strings.HasPrefix("Gopher", ""))
}
```

output:
```txt
true
false
true
```

## <a id="HasSuffix">func</a> [HasSuffix](https://golang.org/src/strings/strings.go?s=11314:11351#L443)
<pre>func HasSuffix(s, suffix <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
HasSuffix tests whether the string s ends with suffix.

HasSuffix 判断 s 是否以后缀 suffix 结尾.

<a id="example_HasSuffix">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.HasSuffix("Amigo", "go"))
  fmt.Println(strings.HasSuffix("Amigo", "O"))
  fmt.Println(strings.HasSuffix("Amigo", "Ami"))
  fmt.Println(strings.HasSuffix("Amigo", ""))
}
```

output:
```txt
true
false
false
true
```

## <a id="Index">func</a> [Index](https://golang.org/src/strings/strings.go?s=25852:25884#L1017)
<pre>func Index(s, substr <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#int">int</a></pre>
Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.

Index 返回 substr 第一次出现在 s 中的位置, 不存在则返回-1.

<a id="example_Index">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.Index("chicken", "ken"))
  fmt.Println(strings.Index("chicken", "dmr"))
}
```

output:
```txt
4
-1
```

## <a id="IndexAny">func</a> [IndexAny](https://golang.org/src/strings/strings.go?s=4328:4362#L168)
<pre>func IndexAny(s, chars <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#int">int</a></pre>
IndexAny returns the index of the first instance of any Unicode code point
from chars in s, or -1 if no Unicode code point from chars is present in s.

IndexAny 返回字符串chars中的任意一个Unicode码点在s中第一次出现的位置. 如果不存在共同的码点则返回-1.

<a id="example_IndexAny">Example</a>
```gopackage main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.IndexAny("chicken", "aeiouy"))
  fmt.Println(strings.IndexAny("crwth", "aeiouy"))
}
```

output:
```txt
2
-1
```

## <a id="IndexByte">func</a> [IndexByte](https://golang.org/src/strings/strings.go?s=3564:3600#L140)
<pre>func IndexByte(s <a href="/pkg/builtin/#string">string</a>, c <a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#int">int</a></pre>
IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.

IndexByte 返回字符c在s中第一次出现的位置，不存在则返回-1.

<a id="example_IndexByte">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.IndexByte("golang", 'g'))
  fmt.Println(strings.IndexByte("gophers", 'h'))
  fmt.Println(strings.IndexByte("golang", 'x'))
}
```

output:
```txt
0
3
-1
```

## <a id="IndexFunc">func</a> [IndexFunc](https://golang.org/src/strings/strings.go?s=18737:18784#L752)
<pre>func IndexFunc(s <a href="/pkg/builtin/#string">string</a>, f func(<a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a>) <a href="/pkg/builtin/#int">int</a></pre>
IndexFunc returns the index into s of the first Unicode
code point satisfying f(c), or -1 if none do.

IndexFunc 返回s中第一个满足f(c)的Unicode码点的位置, 不存在则返回-1.

<a id="example_IndexFunc">Example</a>
```go
package main

import (
  "fmt"
  "strings"
  "unicode"
)

func main() {
  f := func(c rune) bool {
    return unicode.Is(unicode.Han, c)
  }
  fmt.Println(strings.IndexFunc("Hello, 世界", f))
  fmt.Println(strings.IndexFunc("Hello, world", f))
}
```

output:
```txt
7
-1
```

## <a id="IndexRune">func</a> [IndexRune](https://golang.org/src/strings/strings.go?s=3860:3896#L148)
<pre>func IndexRune(s <a href="/pkg/builtin/#string">string</a>, r <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#int">int</a></pre>
IndexRune returns the index of the first instance of the Unicode code point
r, or -1 if rune is not present in s.
If r is utf8.RuneError, it returns the first instance of any
invalid UTF-8 byte sequence.



IndexRune 返回r在s中第一次出现的位置;如果不存在则返回-1. 如果r是utf8.RuneError,它会返回第一个无效的UTF-8 byte序列的位置.

<a id="example_IndexRune">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.IndexRune("chicken", 'k'))
  fmt.Println(strings.IndexRune("chicken", 'd'))
}
```

output:
```txt
4
-1
```

## <a id="Join">func</a> [Join](https://golang.org/src/strings/strings.go?s=10765:10805#L415)
<pre>func Join(a []<a href="/pkg/builtin/#string">string</a>, sep <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Join concatenates the elements of a to create a single string. The separator string
sep is placed between elements in the resulting string.

Join 会将一系列string切片连接成一个新的string,它们之间用sep来分隔.

<a id="example_Join">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  s := []string{"foo", "bar", "baz"}
  fmt.Println(strings.Join(s, ", "))
}
```

output:
```txt
foo, bar, baz
```

## <a id="LastIndex">func</a> [LastIndex](https://golang.org/src/strings/strings.go?s=2802:2838#L103)
<pre>func LastIndex(s, substr <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#int">int</a></pre>
LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.

LastIndex 会返回substr在s中最后一次出现的位置, 不存在则返回-1.

<a id="example_LastIndex">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.Index("go gopher", "go"))
  fmt.Println(strings.LastIndex("go gopher", "go"))
  fmt.Println(strings.LastIndex("go gopher", "rodent"))
}
```

output:
```txt
0
3
-1
```

## <a id="LastIndexAny">func</a> [LastIndexAny](https://golang.org/src/strings/strings.go?s=4869:4907#L196)
<pre>func LastIndexAny(s, chars <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#int">int</a></pre>
LastIndexAny returns the index of the last instance of any Unicode code
point from chars in s, or -1 if no Unicode code point from chars is
present in s.

LastIndexAny 返回字符串chars中的任意一个Unicode码点在s中最后一次出现的位置, 不存在则返回-1.

<a id="example_LastIndexAny">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.LastIndexAny("go gopher", "go"))
  fmt.Println(strings.LastIndexAny("go gopher", "rodent"))
  fmt.Println(strings.LastIndexAny("go gopher", "fail"))
}
```

output:
```txt
4
8
-1
```

## <a id="LastIndexByte">func</a> [LastIndexByte](https://golang.org/src/strings/strings.go?s=5419:5459#L224)
<pre>func LastIndexByte(s <a href="/pkg/builtin/#string">string</a>, c <a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#int">int</a></pre>
LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.

LastIndexByte 返回字符c在s中最后一次出现的位置, 不存在则返回-1.

<a id="example_LastIndexByte">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.LastIndexByte("Hello, world", 'l'))
  fmt.Println(strings.LastIndexByte("Hello, world", 'o'))
  fmt.Println(strings.LastIndexByte("Hello, world", 'x'))
}
```

output:
```txt
10
8
-1
```

## <a id="LastIndexFunc">func</a> [LastIndexFunc](https://golang.org/src/strings/strings.go?s=18931:18982#L758)
<pre>func LastIndexFunc(s <a href="/pkg/builtin/#string">string</a>, f func(<a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a>) <a href="/pkg/builtin/#int">int</a></pre>
LastIndexFunc returns the index into s of the last
Unicode code point satisfying f(c), or -1 if none do.

LastIndexFunc 返回s中最后一个满足f(c)的码点的位置, 不存在则返回-1.

<a id="example_LastIndexFunc">Example</a>
```go
package main

import (
  "fmt"
  "strings"
  "unicode"
)

func main() {
  fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber))
  fmt.Println(strings.LastIndexFunc("123 go", unicode.IsNumber))
  fmt.Println(strings.LastIndexFunc("go", unicode.IsNumber))
}
```

output:
```txt
5
2
-1
```

## <a id="Map">func</a> [Map](https://golang.org/src/strings/strings.go?s=11634:11684#L450)
<pre>func Map(mapping func(<a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#rune">rune</a>, s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Map returns a copy of the string s with all its characters modified
according to the mapping function. If mapping returns a negative value, the character is
dropped from the string with no replacement.

Map 返回s的副本, 且所有字符已根据mapping函数进行了修改. 如果mapping返回一个负值, 将会丢弃该码值且不会被替换.

<a id="example_Map">Example</a>
```go
package main

import (
  "fmt"
  "strings"
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
  fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
}
```

output:
```txt
'Gjnf oevyyvt naq gur fyvgul tbcure...
```

## <a id="Repeat">func</a> [Repeat](https://golang.org/src/strings/strings.go?s=12987:13026#L513)
<pre>func Repeat(s <a href="/pkg/builtin/#string">string</a>, count <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#string">string</a></pre>
Repeat returns a new string consisting of count copies of the string s.

It panics if count is negative or if
the result of (len(s) * count) overflows.

Repeat 返回count个s串联而成的新string.

如果count是负数或(len(s) * count)的值溢出, 它会panic.

<a id="example_Repeat">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println("ba" + strings.Repeat("na", 2))
}
```

output:
```txt
banana
```

## <a id="Replace">func</a> [Replace](https://golang.org/src/strings/strings.go?s=23551:23597#L918)
<pre>func Replace(s, old, new <a href="/pkg/builtin/#string">string</a>, n <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#string">string</a></pre>
Replace returns a copy of the string s with the first n
non-overlapping instances of old replaced by new.
If old is empty, it matches at the beginning of the string
and after each UTF-8 sequence, yielding up to k+1 replacements
for a k-rune string.
If n < 0, there is no limit on the number of replacements.

Replace 返回s的一个副本,且将s中前n个不重叠的old替换成了new. 如果old为空,则在每个utf8字符前面和后面都插入一个new, k个字符需要k+1次替换. 如果n<0则不会限制替换次数.

<a id="example_Replace">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
  fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
}
```

output:
```txt
oinky oinky oink
moo moo moo
```

## <a id="ReplaceAll">func</a> [ReplaceAll](https://golang.org/src/strings/strings.go?s=24486:24528#L957)
<pre>func ReplaceAll(s, old, new <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
ReplaceAll returns a copy of the string s with all
non-overlapping instances of old replaced by new.
If old is empty, it matches at the beginning of the string
and after each UTF-8 sequence, yielding up to k+1 replacements
for a k-rune string.

ReplaceAll 返回s的一个副本,且将s中所有不重叠的old切片替换成了new. 如果old为空,则在每个utf8字符前面和后面都插入一个new, k个字符需要k+1次替换.

<a id="example_ReplaceAll">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))
}
```

output:
```txt
moo moo moo
```

## <a id="Split">func</a> [Split](https://golang.org/src/strings/strings.go?s=7505:7539#L298)
<pre>func Split(s, sep <a href="/pkg/builtin/#string">string</a>) []<a href="/pkg/builtin/#string">string</a></pre>
Split slices s into all substrings separated by sep and returns a slice of
the substrings between those separators.

If s does not contain sep and sep is not empty, Split returns a
slice of length 1 whose only element is s.

If sep is empty, Split splits after each UTF-8 sequence. If both s
and sep are empty, Split returns an empty slice.

It is equivalent to SplitN with a count of -1.

Split 根据sep将s分割并返回分割结果.

如果s不包含sep且sep不为空，则 Split 返回长度为1的slice, 其唯一元素为s.

如果sep为空,Split按每个utf8编码切分. 如果s和sep都为空, Split返回空slice.

它等效于参数为-1的SplitN方法.

<a id="example_Split">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Printf("%q\n", strings.Split("a,b,c", ","))
  fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
  fmt.Printf("%q\n", strings.Split(" xyz ", ""))
  fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
}
```

output:
```txt
["a" "b" "c"]
["" "man " "plan " "canal panama"]
[" " "x" "y" "z" " "]
[""]
```

## <a id="SplitAfter">func</a> [SplitAfter](https://golang.org/src/strings/strings.go?s=8004:8043#L310)
<pre>func SplitAfter(s, sep <a href="/pkg/builtin/#string">string</a>) []<a href="/pkg/builtin/#string">string</a></pre>
SplitAfter slices s into all substrings after each instance of sep and
returns a slice of those substrings.

If s does not contain sep and sep is not empty, SplitAfter returns
a slice of length 1 whose only element is s.

If sep is empty, SplitAfter splits after each UTF-8 sequence. If
both s and sep are empty, SplitAfter returns an empty slice.

It is equivalent to SplitAfterN with a count of -1.

SplitAfter 从 s 中的每个 sep 之后的位置分割 s 并返回分割结果.

如果s不包含sep且sep不为空，则 SplitAfter 返回长度为1的slice, 其唯一元素为s.

如果sep为空, SplitAfter 按每个utf8编码切分. 如果s和sep都为空, SplitAfter 返回空slice.

它等效于参数为-1的SplitAfterN方法.

<a id="example_SplitAfter">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
}
```

output:
```txt
["a," "b," "c"]
```

## <a id="SplitAfterN">func</a> [SplitAfterN](https://golang.org/src/strings/strings.go?s=6998:7045#L284)
<pre>func SplitAfterN(s, sep <a href="/pkg/builtin/#string">string</a>, n <a href="/pkg/builtin/#int">int</a>) []<a href="/pkg/builtin/#string">string</a></pre>
SplitAfterN slices s into substrings after each instance of sep and
returns a slice of those substrings.

The count determines the number of substrings to return:


  n > 0: at most n substrings; the last substring will be the unsplit remainder.
  n == 0: the result is nil (zero substrings)
  n < 0: all substrings

Edge cases for s and sep (for example, empty strings) are handled
as described in the documentation for SplitAfter.

SplitAfterN slices s into substrings after each instance of sep and
returns a slice of those substrings.

SplitAfterN 从 s 中的每个 sep 之后的位置分割 s 并返回分割结果. 参数n决定返回的子字符串的数目:

  n > 0: at most n substrings; the last substring will be the unsplit remainder. // 至多 n 个元素; 其中最后一个元素不会再分割.
  n == 0: the result is nil (zero substrings) // 返回nil
  n < 0: all substrings // 所有的子切片

s和sep的边缘情况（例如，空字符串）的处理方法同SplitAfter文档中所述.


<a id="example_SplitAfterN">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))
}
```

output:
```txt
["a," "b,c"]
```

## <a id="SplitN">func</a> [SplitN](https://golang.org/src/strings/strings.go?s=6461:6503#L272)
<pre>func SplitN(s, sep <a href="/pkg/builtin/#string">string</a>, n <a href="/pkg/builtin/#int">int</a>) []<a href="/pkg/builtin/#string">string</a></pre>
SplitN slices s into substrings separated by sep and returns a slice of
the substrings between those separators.

The count determines the number of substrings to return:


  n > 0: at most n substrings; the last substring will be the unsplit remainder.
  n == 0: the result is nil (zero substrings)
  n < 0: all substrings

Edge cases for s and sep (for example, empty strings) are handled
as described in the documentation for Split.

Split 根据sep将s分割并返回分割结果. 参数n决定返回的子字符串的数目:

  n > 0: at most n substrings; the last substring will be the unsplit remainder. // 至多 n 个元素; 其中最后一个元素不会再分割.
  n == 0: the result is nil (zero substrings) // 返回nil
  n < 0: all substrings // 所有的子切片

s和sep的边缘情况（例如，空字符串）的处理方法同Split文档中所述.

<a id="example_SplitN">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
  z := strings.SplitN("a,b,c", ",", 0)
  fmt.Printf("%q (nil = %v)\n", z, z == nil)
}
```

output:
```txt
["a" "b,c"]
[] (nil = true)
```

## <a id="Title">func</a> [Title](https://golang.org/src/strings/strings.go?s=17489:17516#L704)
<pre>func Title(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Title returns a copy of the string s with all Unicode letters that begin words
mapped to their Unicode title case.

BUG(rsc): The rule Title uses for word boundaries does not handle Unicode punctuation properly.

Title 它会返回s的一个副本，并把s中每个单词的首字母改为其Unicode的Title(大写)形式.

BUG(rsc): Title方法中用于分词的规则不能正确地处理Unicode标点符号.

<a id="example_Title">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  // Compare this example to the ToTitle example.
  fmt.Println(strings.Title("her royal highness"))
  fmt.Println(strings.Title("loud noises"))
  fmt.Println(strings.Title("хлеб"))
}
```

output:
```txt
Her Royal Highness
Loud Noises
Хлеб
```

## <a id="ToLower">func</a> [ToLower](https://golang.org/src/strings/strings.go?s=14252:14281#L574)
<pre>func ToLower(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
ToLower returns s with all Unicode letters mapped to their lower case.

ToLower 返回s的一个副本，并把其中所有的Unicode字母转换为小写.

<a id="example_ToLower">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.ToLower("Gopher"))
}
```

output:
```txt
gopher
```

## <a id="ToLowerSpecial">func</a> [ToLowerSpecial](https://golang.org/src/strings/strings.go?s=15314:15373#L615)
<pre>func ToLowerSpecial(c <a href="/pkg/unicode/">unicode</a>.<a href="/pkg/unicode/#SpecialCase">SpecialCase</a>, s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their
lower case using the case mapping specified by c.

ToLowerSpecial 返回s的一个副本, 已把其中所有的Unicode字母转换为小写,且优先使用指定的特殊转换规则.

<a id="example_ToLowerSpecial">Example</a>
```go
package main

import (
  "fmt"
  "strings"
  "unicode"
)

func main() {
  fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))
}
```

output:
```txt
önnek iş
```

## <a id="ToTitle">func</a> [ToTitle](https://golang.org/src/strings/strings.go?s=14871:14900#L605)
<pre>func ToTitle(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
ToTitle returns a copy of the string s with all Unicode letters mapped to
their Unicode title case.

ToTitle 返回s的一个副本，已把其中所有的Unicode字母转换为title形式.

<a id="example_ToTitle">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  // Compare this example to the Title example.
  fmt.Println(strings.ToTitle("her royal highness"))
  fmt.Println(strings.ToTitle("loud noises"))
  fmt.Println(strings.ToTitle("хлеб"))
}
```

output:
```txt
HER ROYAL HIGHNESS
LOUD NOISES
ХЛЕБ
```

## <a id="ToTitleSpecial">func</a> [ToTitleSpecial](https://golang.org/src/strings/strings.go?s=15563:15622#L621)
<pre>func ToTitleSpecial(c <a href="/pkg/unicode/">unicode</a>.<a href="/pkg/unicode/#SpecialCase">SpecialCase</a>, s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their
Unicode title case, giving priority to the special casing rules.


ToTitleSpecial 返回s的一个副本, 已把其中所有的Unicode字母转换为Title形式, 且优先使用指定的特殊转换规则.

<a id="example_ToTitleSpecial">Example</a>
```go
package main

import (
  "fmt"
  "strings"
  "unicode"
)

func main() {
  fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))
}
```

output:
```txt
DÜNYANIN İLK BORSA YAPISI AİZONAİ KABUL EDİLİR
```

## <a id="ToUpper">func</a> [ToUpper](https://golang.org/src/strings/strings.go?s=13665:13694#L544)
<pre>func ToUpper(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
ToUpper returns s with all Unicode letters mapped to their upper case.

ToUpper 返回s的一个副本，并把其中所有的Unicode字母转换为大写.

<a id="example_ToUpper">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.ToUpper("Gopher"))
}
```

output:
```txt
GOPHER
```

## <a id="ToUpperSpecial">func</a> [ToUpperSpecial](https://golang.org/src/strings/strings.go?s=15080:15139#L609)
<pre>func ToUpperSpecial(c <a href="/pkg/unicode/">unicode</a>.<a href="/pkg/unicode/#SpecialCase">SpecialCase</a>, s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their
upper case using the case mapping specified by c.

ToUpperSpecial 返回s的一个副本, 已把其中所有的Unicode字母转换为大写形式, 且优先使用指定的特殊转换规则.

<a id="example_ToUpperSpecial">Example</a>
```go
package main

import (
  "fmt"
  "strings"
  "unicode"
)

func main() {
  fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))
}
```

output:
```txt
ÖRNEK İŞ
```

## <a id="ToValidUTF8">func</a> [ToValidUTF8](https://golang.org/src/strings/strings.go?s=15805:15851#L627)
<pre>func ToValidUTF8(s, replacement <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
ToValidUTF8 returns a copy of the string s with each run of invalid UTF-8 byte sequences
replaced by the replacement string, which may be empty.

ToValidUTF8 返回一个副本，其中已将每个无效的UTF-8替换为replacement，replacement可以为空.

## <a id="Trim">func</a> [Trim](https://golang.org/src/strings/strings.go?s=21049:21090#L830)
<pre>func Trim(s <a href="/pkg/builtin/#string">string</a>, cutset <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Trim returns a slice of the string s with all leading and
trailing Unicode code points contained in cutset removed.

Trim 返回将s前后端所有cutset中包含的utf8码点都已去掉的字符串.

<a id="example_Trim">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Print(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
}
```

output:
```txt
Hello, Gophers
```

## <a id="TrimFunc">func</a> [TrimFunc](https://golang.org/src/strings/strings.go?s=18529:18578#L746)
<pre>func TrimFunc(s <a href="/pkg/builtin/#string">string</a>, f func(<a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a>) <a href="/pkg/builtin/#string">string</a></pre>
TrimFunc returns a slice of the string s with all leading
and trailing Unicode code points c satisfying f(c) removed.

TrimFunc 返回将s前后端中所有满足f(c)的utf8码点都去掉的字符串.

<a id="example_TrimFunc">Example</a>
```go
package main

import (
  "fmt"
  "strings"
  "unicode"
)

func main() {
  fmt.Print(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
    return !unicode.IsLetter(r) && !unicode.IsNumber(r)
  }))
}
```

output:
```txt
Hello, Gophers
```

## <a id="TrimLeft">func</a> [TrimLeft](https://golang.org/src/strings/strings.go?s=21347:21392#L841)
<pre>func TrimLeft(s <a href="/pkg/builtin/#string">string</a>, cutset <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
TrimLeft returns a slice of the string s with all leading
Unicode code points contained in cutset removed.

To remove a prefix, use TrimPrefix instead.

TrimLeft 返回将s前端中所有cutset包含的utf8码点都已去掉的字符串.

如果是移除prefix, 请使用 TrimPrefix.

<a id="example_TrimLeft">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Print(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡"))
}
```

output:
```txt
Hello, Gophers!!!
```

## <a id="TrimLeftFunc">func</a> [TrimLeftFunc](https://golang.org/src/strings/strings.go?s=17938:17991#L723)
<pre>func TrimLeftFunc(s <a href="/pkg/builtin/#string">string</a>, f func(<a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a>) <a href="/pkg/builtin/#string">string</a></pre>
TrimLeftFunc returns a slice of the string s with all leading
Unicode code points c satisfying f(c) removed.

TrimFunc 返回将s前端中所有满足f(c)的utf8码点都去掉的字符串.

<a id="example_TrimLeftFunc">Example</a>
```go
package main

import (
  "fmt"
  "strings"
  "unicode"
)

func main() {
  fmt.Print(strings.TrimLeftFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
    return !unicode.IsLetter(r) && !unicode.IsNumber(r)
  }))
}
```

output:
```txt
Hello, Gophers!!!
```

## <a id="TrimPrefix">func</a> [TrimPrefix](https://golang.org/src/strings/strings.go?s=22869:22909#L896)
<pre>func TrimPrefix(s, prefix <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
TrimPrefix returns s without the provided leading prefix string.
If s doesn't start with prefix, s is returned unchanged.

TrimPrefix 删除s头部的前缀prefix. 如果不存在前缀prefix则返回原始的s.

<a id="example_TrimPrefix">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  var s = "¡¡¡Hello, Gophers!!!"
  s = strings.TrimPrefix(s, "¡¡¡Hello, ")
  s = strings.TrimPrefix(s, "¡¡¡Howdy, ")
  fmt.Print(s)
}
```

output:
```txt
Gophers!!!
```

## <a id="TrimRight">func</a> [TrimRight](https://golang.org/src/strings/strings.go?s=21656:21702#L852)
<pre>func TrimRight(s <a href="/pkg/builtin/#string">string</a>, cutset <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
TrimRight returns a slice of the string s, with all trailing
Unicode code points contained in cutset removed.

To remove a suffix, use TrimSuffix instead.

TrimRight 返回将s后端中所有cutset包含的utf8码点都已去掉的字符串.

<a id="example_TrimRight">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Print(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡"))
}
```

output:
```txt
¡¡¡Hello, Gophers
```

## <a id="TrimRightFunc">func</a> [TrimRightFunc](https://golang.org/src/strings/strings.go?s=18186:18240#L733)
<pre>func TrimRightFunc(s <a href="/pkg/builtin/#string">string</a>, f func(<a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#bool">bool</a>) <a href="/pkg/builtin/#string">string</a></pre>
TrimRightFunc returns a slice of the string s with all trailing
Unicode code points c satisfying f(c) removed.

TrimRightFunc 返回将s后端中所有满足f(c)的utf8码点都去掉的字符串.

<a id="example_TrimRightFunc">Example</a>
```go
package main

import (
  "fmt"
  "strings"
  "unicode"
)

func main() {
  fmt.Print(strings.TrimRightFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
    return !unicode.IsLetter(r) && !unicode.IsNumber(r)
  }))
}
```

output:
```txt
¡¡¡Hello, Gophers
```

## <a id="TrimSpace">func</a> [TrimSpace](https://golang.org/src/strings/strings.go?s=21924:21955#L861)
<pre>func TrimSpace(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
TrimSpace returns a slice of the string s, with all leading
and trailing white space removed, as defined by Unicode.

TrimSpace 删除s前后端所有unicode中定义的空白符.

<a id="example_TrimSpace">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
}
```

output:
```txt
Hello, Gophers
```

## <a id="TrimSuffix">func</a> [TrimSuffix](https://golang.org/src/strings/strings.go?s=23107:23147#L905)
<pre>func TrimSuffix(s, suffix <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
TrimSuffix returns s without the provided trailing suffix string.
If s doesn't end with suffix, s is returned unchanged.

TrimPrefix 删除s尾部的后缀suffix. 如果不存在后缀suffix则返回原始的s.

<a id="example_TrimSuffix">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  var s = "¡¡¡Hello, Gophers!!!"
  s = strings.TrimSuffix(s, ", Gophers!!!")
  s = strings.TrimSuffix(s, ", Marmots!!!")
  fmt.Print(s)
}
```

output:
```txt
¡¡¡Hello
```



## <a id="Builder">type</a> [Builder](https://golang.org/src/strings/builder.go?s=386:479#L5)
A Builder is used to efficiently build a string using Write methods.
It minimizes memory copying. The zero value is ready to use.
Do not copy a non-zero Builder.

Builder 可以使用 Write 高效地创建字符串. 它会把内存拷贝降到最低程度. 可以直接使用 Builder 的零值. 不要复制一个非零值的 Builder.

<pre>type Builder struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Builder">Example</a>
```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  var b strings.Builder
  for i := 3; i >= 1; i-- {
    fmt.Fprintf(&b, "%d...", i)
  }
  b.WriteString("ignition")
  fmt.Println(b.String())

}
```

output:
```txt
3...2...1...ignition
```






### <a id="Builder.Cap">func</a> (\*Builder) [Cap](https://golang.org/src/strings/builder.go?s=1761:1788#L46)
<pre>func (b *<a href="#Builder">Builder</a>) Cap() <a href="/pkg/builtin/#int">int</a></pre>
Cap returns the capacity of the builder's underlying byte slice. It is the
total space allocated for the string being built and includes any bytes
already written.




### <a id="Builder.Grow">func</a> (\*Builder) [Grow](https://golang.org/src/strings/builder.go?s=2344:2373#L65)
<pre>func (b *<a href="#Builder">Builder</a>) Grow(n <a href="/pkg/builtin/#int">int</a>)</pre>
Grow grows b's capacity, if necessary, to guarantee space for
another n bytes. After Grow(n), at least n bytes can be written to b
without another allocation. If n is negative, Grow panics.




### <a id="Builder.Len">func</a> (\*Builder) [Len](https://golang.org/src/strings/builder.go?s=1537:1564#L41)
<pre>func (b *<a href="#Builder">Builder</a>) Len() <a href="/pkg/builtin/#int">int</a></pre>
Len returns the number of accumulated bytes; b.Len() == len(b.String()).




### <a id="Builder.Reset">func</a> (\*Builder) [Reset](https://golang.org/src/strings/builder.go?s=1853:1878#L49)
<pre>func (b *<a href="#Builder">Builder</a>) Reset()</pre>
Reset resets the Builder to be empty.




### <a id="Builder.String">func</a> (\*Builder) [String](https://golang.org/src/strings/builder.go?s=1379:1412#L36)
<pre>func (b *<a href="#Builder">Builder</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns the accumulated string.




### <a id="Builder.Write">func</a> (\*Builder) [Write](https://golang.org/src/strings/builder.go?s=2591:2637#L77)
<pre>func (b *<a href="#Builder">Builder</a>) Write(p []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Write appends the contents of p to b's buffer.
Write always returns len(p), nil.




### <a id="Builder.WriteByte">func</a> (\*Builder) [WriteByte](https://golang.org/src/strings/builder.go?s=2791:2832#L85)
<pre>func (b *<a href="#Builder">Builder</a>) WriteByte(c <a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#error">error</a></pre>
WriteByte appends the byte c to b's buffer.
The returned error is always nil.




### <a id="Builder.WriteRune">func</a> (\*Builder) [WriteRune](https://golang.org/src/strings/builder.go?s=3017:3065#L93)
<pre>func (b *<a href="#Builder">Builder</a>) WriteRune(r <a href="/pkg/builtin/#rune">rune</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
WriteRune appends the UTF-8 encoding of Unicode code point r to b's buffer.
It returns the length of r and a nil error.




### <a id="Builder.WriteString">func</a> (\*Builder) [WriteString](https://golang.org/src/strings/builder.go?s=3425:3477#L110)
<pre>func (b *<a href="#Builder">Builder</a>) WriteString(s <a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#int">int</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
WriteString appends the contents of s to b's buffer.
It returns the length of s and a nil error.




## <a id="Reader">type</a> [Reader](https://golang.org/src/strings/reader.go?s=446:576#L7)
A Reader implements the io.Reader, io.ReaderAt, io.Seeker, io.WriterTo,
io.ByteScanner, and io.RuneScanner interfaces by reading
from a string.
The zero value for Reader operates like a Reader of an empty string.


<pre>type Reader struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewReader">func</a> [NewReader](https://golang.org/src/strings/reader.go?s=3567:3599#L144)
<pre>func NewReader(s <a href="/pkg/builtin/#string">string</a>) *<a href="#Reader">Reader</a></pre>
NewReader returns a new Reader reading from s.
It is similar to bytes.NewBufferString but more efficient and read-only.






### <a id="Reader.Len">func</a> (\*Reader) [Len](https://golang.org/src/strings/reader.go?s=653:679#L15)
<pre>func (r *<a href="#Reader">Reader</a>) Len() <a href="/pkg/builtin/#int">int</a></pre>
Len returns the number of bytes of the unread portion of the
string.




### <a id="Reader.Read">func</a> (\*Reader) [Read](https://golang.org/src/strings/reader.go?s=1042:1092#L28)
<pre>func (r *<a href="#Reader">Reader</a>) Read(b []<a href="/pkg/builtin/#byte">byte</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>



### <a id="Reader.ReadAt">func</a> (\*Reader) [ReadAt](https://golang.org/src/strings/reader.go?s=1215:1278#L38)
<pre>func (r *<a href="#Reader">Reader</a>) ReadAt(b []<a href="/pkg/builtin/#byte">byte</a>, off <a href="/pkg/builtin/#int64">int64</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>



### <a id="Reader.ReadByte">func</a> (\*Reader) [ReadByte](https://golang.org/src/strings/reader.go?s=1526:1567#L53)
<pre>func (r *<a href="#Reader">Reader</a>) ReadByte() (<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>



### <a id="Reader.ReadRune">func</a> (\*Reader) [ReadRune](https://golang.org/src/strings/reader.go?s=1846:1904#L72)
<pre>func (r *<a href="#Reader">Reader</a>) ReadRune() (ch <a href="/pkg/builtin/#rune">rune</a>, size <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>



### <a id="Reader.Reset">func</a> (\*Reader) [Reset](https://golang.org/src/strings/reader.go?s=3381:3413#L140)
<pre>func (r *<a href="#Reader">Reader</a>) Reset(s <a href="/pkg/builtin/#string">string</a>)</pre>
Reset resets the Reader to be reading from s.




### <a id="Reader.Seek">func</a> (\*Reader) [Seek](https://golang.org/src/strings/reader.go?s=2495:2557#L100)
<pre>func (r *<a href="#Reader">Reader</a>) Seek(offset <a href="/pkg/builtin/#int64">int64</a>, whence <a href="/pkg/builtin/#int">int</a>) (<a href="/pkg/builtin/#int64">int64</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Seek implements the io.Seeker interface.




### <a id="Reader.Size">func</a> (\*Reader) [Size](https://golang.org/src/strings/reader.go?s=984:1013#L26)
<pre>func (r *<a href="#Reader">Reader</a>) Size() <a href="/pkg/builtin/#int64">int64</a></pre>
Size returns the original length of the underlying string.
Size is the number of bytes available for reading via ReadAt.
The returned value is always the same and is not affected by calls
to any other method.




### <a id="Reader.UnreadByte">func</a> (\*Reader) [UnreadByte](https://golang.org/src/strings/reader.go?s=1678:1713#L63)
<pre>func (r *<a href="#Reader">Reader</a>) UnreadByte() <a href="/pkg/builtin/#error">error</a></pre>



### <a id="Reader.UnreadRune">func</a> (\*Reader) [UnreadRune](https://golang.org/src/strings/reader.go?s=2155:2190#L87)
<pre>func (r *<a href="#Reader">Reader</a>) UnreadRune() <a href="/pkg/builtin/#error">error</a></pre>



### <a id="Reader.WriteTo">func</a> (\*Reader) [WriteTo](https://golang.org/src/strings/reader.go?s=2975:3033#L121)
<pre>func (r *<a href="#Reader">Reader</a>) WriteTo(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) (n <a href="/pkg/builtin/#int64">int64</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
WriteTo implements the io.WriterTo interface.




## <a id="Replacer">type</a> [Replacer](https://golang.org/src/strings/replace.go?s=318:421#L4)
Replacer replaces a list of strings with replacements.
It is safe for concurrent use by multiple goroutines.


<pre>type Replacer struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewReplacer">func</a> [NewReplacer](https://golang.org/src/strings/replace.go?s=916:960#L22)
<pre>func NewReplacer(oldnew ...<a href="/pkg/builtin/#string">string</a>) *<a href="#Replacer">Replacer</a></pre>
NewReplacer returns a new Replacer from a list of old, new string
pairs. Replacements are performed in the order they appear in the
target string, without overlapping matches. The old string
comparisons are done in argument order.

NewReplacer panics if given an odd number of arguments.


<a id="example_NewReplacer">Example</a>
```go
```

output:
```txt
```




### <a id="Replacer.Replace">func</a> (\*Replacer) [Replace](https://golang.org/src/strings/replace.go?s=2497:2540#L85)
<pre>func (r *<a href="#Replacer">Replacer</a>) Replace(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
Replace returns a copy of s with all replacements performed.




### <a id="Replacer.WriteString">func</a> (\*Replacer) [WriteString](https://golang.org/src/strings/replace.go?s=2655:2727#L91)
<pre>func (r *<a href="#Replacer">Replacer</a>) WriteString(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, s <a href="/pkg/builtin/#string">string</a>) (n <a href="/pkg/builtin/#int">int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
WriteString writes s to w with all replacements performed.







