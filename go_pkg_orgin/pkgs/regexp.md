

# regexp
`import "regexp"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package regexp implements regular expression search.

The syntax of the regular expressions accepted is the same
general syntax used by Perl, Python, and other languages.
More precisely, it is the syntax accepted by RE2 and described at
<a href="https://golang.org/s/re2syntax">https://golang.org/s/re2syntax</a>, except for \C.
For an overview of the syntax, run


	go doc regexp/syntax

The regexp implementation provided by this package is
guaranteed to run in time linear in the size of the input.
(This is a property not guaranteed by most open source
implementations of regular expressions.) For more information
about this property, see


	<a href="https://swtch.com/~rsc/regexp/regexp1.html">https://swtch.com/~rsc/regexp/regexp1.html</a>

or any book about automata theory.

All characters are UTF-8-encoded code points.

There are 16 methods of Regexp that match a regular expression and identify
the matched text. Their names are matched by this regular expression:


	Find(All)?(String)?(Submatch)?(Index)?

If 'All' is present, the routine matches successive non-overlapping
matches of the entire expression. Empty matches abutting a preceding
match are ignored. The return value is a slice containing the successive
return values of the corresponding non-'All' routine. These routines take
an extra integer argument, n. If n >= 0, the function returns at most n
matches/submatches; otherwise, it returns all of them.

If 'String' is present, the argument is a string; otherwise it is a slice
of bytes; return values are adjusted as appropriate.

If 'Submatch' is present, the return value is a slice identifying the
successive submatches of the expression. Submatches are matches of
parenthesized subexpressions (also known as capturing groups) within the
regular expression, numbered from left to right in order of opening
parenthesis. Submatch 0 is the match of the entire expression, submatch 1
the match of the first parenthesized subexpression, and so on.

If 'Index' is present, matches and submatches are identified by byte index
pairs within the input string: result[2*n:2*n+1] identifies the indexes of
the nth submatch. The pair for n==0 identifies the match of the entire
expression. If 'Index' is not present, the match is identified by the text
of the match/submatch. If an index is negative or text is nil, it means that
subexpression did not match any string in the input. For 'String' versions
an empty string means either no match or an empty match.

There is also a subset of the methods that can be applied to text read
from a RuneReader:


	MatchReader, FindReaderIndex, FindReaderSubmatchIndex

This set may grow. Note that regular expression matches may need to
examine text beyond the text returned by a match, so the methods that
match text from a RuneReader may read arbitrarily far into the input
before returning.

(There are a few other methods that do not match this pattern.)


<a id="example_">Example</a>
```go
```

output:
```txt
```


## <a id="pkg-index">Index</a>
* [func Match(pattern string, b []byte) (matched bool, err error)](#Match)
* [func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)](#MatchReader)
* [func MatchString(pattern string, s string) (matched bool, err error)](#MatchString)
* [func QuoteMeta(s string) string](#QuoteMeta)
* [type Regexp](#Regexp)
  * [func Compile(expr string) (*Regexp, error)](#Compile)
  * [func CompilePOSIX(expr string) (*Regexp, error)](#CompilePOSIX)
  * [func MustCompile(str string) *Regexp](#MustCompile)
  * [func MustCompilePOSIX(str string) *Regexp](#MustCompilePOSIX)
  * [func (re *Regexp) Copy() *Regexp](#Regexp.Copy)
  * [func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte](#Regexp.Expand)
  * [func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte](#Regexp.ExpandString)
  * [func (re *Regexp) Find(b []byte) []byte](#Regexp.Find)
  * [func (re *Regexp) FindAll(b []byte, n int) [][]byte](#Regexp.FindAll)
  * [func (re *Regexp) FindAllIndex(b []byte, n int) [][]int](#Regexp.FindAllIndex)
  * [func (re *Regexp) FindAllString(s string, n int) []string](#Regexp.FindAllString)
  * [func (re *Regexp) FindAllStringIndex(s string, n int) [][]int](#Regexp.FindAllStringIndex)
  * [func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string](#Regexp.FindAllStringSubmatch)
  * [func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int](#Regexp.FindAllStringSubmatchIndex)
  * [func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte](#Regexp.FindAllSubmatch)
  * [func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int](#Regexp.FindAllSubmatchIndex)
  * [func (re *Regexp) FindIndex(b []byte) (loc []int)](#Regexp.FindIndex)
  * [func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int)](#Regexp.FindReaderIndex)
  * [func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int](#Regexp.FindReaderSubmatchIndex)
  * [func (re *Regexp) FindString(s string) string](#Regexp.FindString)
  * [func (re *Regexp) FindStringIndex(s string) (loc []int)](#Regexp.FindStringIndex)
  * [func (re *Regexp) FindStringSubmatch(s string) []string](#Regexp.FindStringSubmatch)
  * [func (re *Regexp) FindStringSubmatchIndex(s string) []int](#Regexp.FindStringSubmatchIndex)
  * [func (re *Regexp) FindSubmatch(b []byte) [][]byte](#Regexp.FindSubmatch)
  * [func (re *Regexp) FindSubmatchIndex(b []byte) []int](#Regexp.FindSubmatchIndex)
  * [func (re *Regexp) LiteralPrefix() (prefix string, complete bool)](#Regexp.LiteralPrefix)
  * [func (re *Regexp) Longest()](#Regexp.Longest)
  * [func (re *Regexp) Match(b []byte) bool](#Regexp.Match)
  * [func (re *Regexp) MatchReader(r io.RuneReader) bool](#Regexp.MatchReader)
  * [func (re *Regexp) MatchString(s string) bool](#Regexp.MatchString)
  * [func (re *Regexp) NumSubexp() int](#Regexp.NumSubexp)
  * [func (re *Regexp) ReplaceAll(src, repl []byte) []byte](#Regexp.ReplaceAll)
  * [func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte](#Regexp.ReplaceAllFunc)
  * [func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte](#Regexp.ReplaceAllLiteral)
  * [func (re *Regexp) ReplaceAllLiteralString(src, repl string) string](#Regexp.ReplaceAllLiteralString)
  * [func (re *Regexp) ReplaceAllString(src, repl string) string](#Regexp.ReplaceAllString)
  * [func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string](#Regexp.ReplaceAllStringFunc)
  * [func (re *Regexp) Split(s string, n int) []string](#Regexp.Split)
  * [func (re *Regexp) String() string](#Regexp.String)
  * [func (re *Regexp) SubexpNames() []string](#Regexp.SubexpNames)


#### <a id="pkg-examples">Examples</a>
* [Package](#example_)
* [Match](#example_Match)
* [MatchString](#example_MatchString)
* [QuoteMeta](#example_QuoteMeta)
* [Regexp.Expand](#example_Regexp_Expand)
* [Regexp.ExpandString](#example_Regexp_ExpandString)
* [Regexp.Find](#example_Regexp_Find)
* [Regexp.FindAll](#example_Regexp_FindAll)
* [Regexp.FindAllString](#example_Regexp_FindAllString)
* [Regexp.FindAllStringSubmatch](#example_Regexp_FindAllStringSubmatch)
* [Regexp.FindAllStringSubmatchIndex](#example_Regexp_FindAllStringSubmatchIndex)
* [Regexp.FindAllSubmatch](#example_Regexp_FindAllSubmatch)
* [Regexp.FindAllSubmatchIndex](#example_Regexp_FindAllSubmatchIndex)
* [Regexp.FindIndex](#example_Regexp_FindIndex)
* [Regexp.FindString](#example_Regexp_FindString)
* [Regexp.FindStringIndex](#example_Regexp_FindStringIndex)
* [Regexp.FindStringSubmatch](#example_Regexp_FindStringSubmatch)
* [Regexp.FindSubmatch](#example_Regexp_FindSubmatch)
* [Regexp.Match](#example_Regexp_Match)
* [Regexp.MatchString](#example_Regexp_MatchString)
* [Regexp.ReplaceAllLiteralString](#example_Regexp_ReplaceAllLiteralString)
* [Regexp.ReplaceAllString](#example_Regexp_ReplaceAllString)
* [Regexp.ReplaceAllStringFunc](#example_Regexp_ReplaceAllStringFunc)
* [Regexp.Split](#example_Regexp_Split)
* [Regexp.SubexpNames](#example_Regexp_SubexpNames)


#### <a id="pkg-files">Package files</a>
[backtrack.go](https://golang.org/src/regexp/backtrack.go) [exec.go](https://golang.org/src/regexp/exec.go) [onepass.go](https://golang.org/src/regexp/onepass.go) [regexp.go](https://golang.org/src/regexp/regexp.go) 






## <a id="Match">func</a> [Match](https://golang.org/src/regexp/regexp.go?s=16735:16797#L531)
<pre>func Match(pattern <a href="/pkg/builtin/#string">string</a>, b []<a href="/pkg/builtin/#byte">byte</a>) (matched <a href="/pkg/builtin/#bool">bool</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Match reports whether the byte slice b
contains any match of the regular expression pattern.
More complicated queries need to use Compile and the full Regexp interface.


<a id="example_Match">Example</a>
```go
```

output:
```txt
```

## <a id="MatchReader">func</a> [MatchReader](https://golang.org/src/regexp/regexp.go?s=16022:16097#L509)
<pre>func MatchReader(pattern <a href="/pkg/builtin/#string">string</a>, r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#RuneReader">RuneReader</a>) (matched <a href="/pkg/builtin/#bool">bool</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
MatchReader reports whether the text returned by the RuneReader
contains any match of the regular expression pattern.
More complicated queries need to use Compile and the full Regexp interface.



## <a id="MatchString">func</a> [MatchString](https://golang.org/src/regexp/regexp.go?s=16383:16451#L520)
<pre>func MatchString(pattern <a href="/pkg/builtin/#string">string</a>, s <a href="/pkg/builtin/#string">string</a>) (matched <a href="/pkg/builtin/#bool">bool</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
MatchString reports whether the string s
contains any match of the regular expression pattern.
More complicated queries need to use Compile and the full Regexp interface.


<a id="example_MatchString">Example</a>
```go
```

output:
```txt
```

## <a id="QuoteMeta">func</a> [QuoteMeta](https://golang.org/src/regexp/regexp.go?s=21972:22003#L692)
<pre>func QuoteMeta(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
QuoteMeta returns a string that escapes all regular expression metacharacters
inside the argument text; the returned string is a regular expression matching
the literal text.


<a id="example_QuoteMeta">Example</a>
```go
```

output:
```txt
```



## <a id="Regexp">type</a> [Regexp](https://golang.org/src/regexp/regexp.go?s=3439:4412#L72)
Regexp is the representation of a compiled regular expression.
A Regexp is safe for concurrent use by multiple goroutines,
except for configuration methods, such as Longest.


<pre>type Regexp struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="Compile">func</a> [Compile](https://golang.org/src/regexp/regexp.go?s=5651:5693#L122)
<pre>func Compile(expr <a href="/pkg/builtin/#string">string</a>) (*<a href="#Regexp">Regexp</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Compile parses a regular expression and returns, if successful,
a Regexp object that can be used to match against text.

When matching against text, the regexp returns a match that
begins as early as possible in the input (leftmost), and among those
it chooses the one that a backtracking search would have found first.
This so-called leftmost-first matching is the same semantics
that Perl, Python, and other implementations use, although this
package implements it without the expense of backtracking.
For POSIX leftmost-longest matching, see CompilePOSIX.




### <a id="CompilePOSIX">func</a> [CompilePOSIX](https://golang.org/src/regexp/regexp.go?s=6813:6860#L145)
<pre>func CompilePOSIX(expr <a href="/pkg/builtin/#string">string</a>) (*<a href="#Regexp">Regexp</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
CompilePOSIX is like Compile but restricts the regular expression
to POSIX ERE (egrep) syntax and changes the match semantics to
leftmost-longest.

That is, when matching against text, the regexp returns a match that
begins as early as possible in the input (leftmost), and among those
it chooses a match that is as long as possible.
This so-called leftmost-longest matching is the same semantics
that early regular expression implementations used and that POSIX
specifies.

However, there can be multiple leftmost-longest matches, with different
submatch choices, and here this package diverges from POSIX.
Among the possible leftmost-longest matches, this package chooses
the one that a backtracking search would have found first, while POSIX
specifies that the match be chosen to maximize the length of the first
subexpression, then the second, and so on from left to right.
The POSIX rule is computationally prohibitive and not even well-defined.
See <a href="https://swtch.com/~rsc/regexp/regexp2.html#posix">https://swtch.com/~rsc/regexp/regexp2.html#posix</a> for details.




### <a id="MustCompile">func</a> [MustCompile](https://golang.org/src/regexp/regexp.go?s=10768:10804#L298)
<pre>func MustCompile(str <a href="/pkg/builtin/#string">string</a>) *<a href="#Regexp">Regexp</a></pre>
MustCompile is like Compile but panics if the expression cannot be parsed.
It simplifies safe initialization of global variables holding compiled regular
expressions.




### <a id="MustCompilePOSIX">func</a> [MustCompilePOSIX](https://golang.org/src/regexp/regexp.go?s=11123:11164#L309)
<pre>func MustCompilePOSIX(str <a href="/pkg/builtin/#string">string</a>) *<a href="#Regexp">Regexp</a></pre>
MustCompilePOSIX is like CompilePOSIX but panics if the expression cannot be parsed.
It simplifies safe initialization of global variables holding compiled regular
expressions.






### <a id="Regexp.Copy">func</a> (\*Regexp) [Copy](https://golang.org/src/regexp/regexp.go?s=5000:5032#L107)
<pre>func (re *<a href="#Regexp">Regexp</a>) Copy() *<a href="#Regexp">Regexp</a></pre>
Copy returns a new Regexp object copied from re.
Calling Longest on one copy does not affect another.

Deprecated: In earlier releases, when using a Regexp in multiple goroutines,
giving each goroutine its own copy helped to avoid lock contention.
As of Go 1.12, using Copy is no longer necessary to avoid lock contention.
Copy may still be appropriate if the reason for its use is to make
two copies with different Longest settings.




### <a id="Regexp.Expand">func</a> (\*Regexp) [Expand](https://golang.org/src/regexp/regexp.go?s=27515:27600#L884)
<pre>func (re *<a href="#Regexp">Regexp</a>) Expand(dst []<a href="/pkg/builtin/#byte">byte</a>, template []<a href="/pkg/builtin/#byte">byte</a>, src []<a href="/pkg/builtin/#byte">byte</a>, match []<a href="/pkg/builtin/#int">int</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
Expand appends template to dst and returns the result; during the
append, Expand replaces variables in the template with corresponding
matches drawn from src. The match slice should have been returned by
FindSubmatchIndex.

In the template, a variable is denoted by a substring of the form
$name or ${name}, where name is a non-empty sequence of letters,
digits, and underscores. A purely numeric name like $1 refers to
the submatch with the corresponding index; other names refer to
capturing parentheses named with the (?P<name>...) syntax. A
reference to an out of range or unmatched index or a name that is not
present in the regular expression is replaced with an empty slice.

In the $name form, name is taken to be as long as possible: $1x is
equivalent to ${1x}, not ${1}x, and, $10 is equivalent to ${10}, not ${1}0.

To insert a literal $ in the output, use $$ in the template.


<a id="example_Regexp_Expand">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.ExpandString">func</a> (\*Regexp) [ExpandString](https://golang.org/src/regexp/regexp.go?s=27839:27930#L891)
<pre>func (re *<a href="#Regexp">Regexp</a>) ExpandString(dst []<a href="/pkg/builtin/#byte">byte</a>, template <a href="/pkg/builtin/#string">string</a>, src <a href="/pkg/builtin/#string">string</a>, match []<a href="/pkg/builtin/#int">int</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
ExpandString is like Expand but the template and source are strings.
It appends to and returns a byte slice in order to give the calling
code control over allocation.


<a id="example_Regexp_ExpandString">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.Find">func</a> (\*Regexp) [Find](https://golang.org/src/regexp/regexp.go?s=24085:24124#L787)
<pre>func (re *<a href="#Regexp">Regexp</a>) Find(b []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
Find returns a slice holding the text of the leftmost match in b of the regular expression.
A return value of nil indicates no match.


<a id="example_Regexp_Find">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.FindAll">func</a> (\*Regexp) [FindAll](https://golang.org/src/regexp/regexp.go?s=32389:32440#L1048)
<pre>func (re *<a href="#Regexp">Regexp</a>) FindAll(b []<a href="/pkg/builtin/#byte">byte</a>, n <a href="/pkg/builtin/#int">int</a>) [][]<a href="/pkg/builtin/#byte">byte</a></pre>
FindAll is the 'All' version of Find; it returns a slice of all successive
matches of the expression, as defined by the 'All' description in the
package comment.
A return value of nil indicates no match.


<a id="example_Regexp_FindAll">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.FindAllIndex">func</a> (\*Regexp) [FindAllIndex](https://golang.org/src/regexp/regexp.go?s=32912:32967#L1066)
<pre>func (re *<a href="#Regexp">Regexp</a>) FindAllIndex(b []<a href="/pkg/builtin/#byte">byte</a>, n <a href="/pkg/builtin/#int">int</a>) [][]<a href="/pkg/builtin/#int">int</a></pre>
FindAllIndex is the 'All' version of FindIndex; it returns a slice of all
successive matches of the expression, as defined by the 'All' description
in the package comment.
A return value of nil indicates no match.




### <a id="Regexp.FindAllString">func</a> (\*Regexp) [FindAllString](https://golang.org/src/regexp/regexp.go?s=33420:33477#L1084)
<pre>func (re *<a href="#Regexp">Regexp</a>) FindAllString(s <a href="/pkg/builtin/#string">string</a>, n <a href="/pkg/builtin/#int">int</a>) []<a href="/pkg/builtin/#string">string</a></pre>
FindAllString is the 'All' version of FindString; it returns a slice of all
successive matches of the expression, as defined by the 'All' description
in the package comment.
A return value of nil indicates no match.


<a id="example_Regexp_FindAllString">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.FindAllStringIndex">func</a> (\*Regexp) [FindAllStringIndex](https://golang.org/src/regexp/regexp.go?s=33953:34014#L1102)
<pre>func (re *<a href="#Regexp">Regexp</a>) FindAllStringIndex(s <a href="/pkg/builtin/#string">string</a>, n <a href="/pkg/builtin/#int">int</a>) [][]<a href="/pkg/builtin/#int">int</a></pre>
FindAllStringIndex is the 'All' version of FindStringIndex; it returns a
slice of all successive matches of the expression, as defined by the 'All'
description in the package comment.
A return value of nil indicates no match.




### <a id="Regexp.FindAllStringSubmatch">func</a> (\*Regexp) [FindAllStringSubmatch](https://golang.org/src/regexp/regexp.go?s=35681:35748#L1162)
<pre>func (re *<a href="#Regexp">Regexp</a>) FindAllStringSubmatch(s <a href="/pkg/builtin/#string">string</a>, n <a href="/pkg/builtin/#int">int</a>) [][]<a href="/pkg/builtin/#string">string</a></pre>
FindAllStringSubmatch is the 'All' version of FindStringSubmatch; it
returns a slice of all successive matches of the expression, as defined by
the 'All' description in the package comment.
A return value of nil indicates no match.


<a id="example_Regexp_FindAllStringSubmatch">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.FindAllStringSubmatchIndex">func</a> (\*Regexp) [FindAllStringSubmatchIndex](https://golang.org/src/regexp/regexp.go?s=36372:36441#L1187)
<pre>func (re *<a href="#Regexp">Regexp</a>) FindAllStringSubmatchIndex(s <a href="/pkg/builtin/#string">string</a>, n <a href="/pkg/builtin/#int">int</a>) [][]<a href="/pkg/builtin/#int">int</a></pre>
FindAllStringSubmatchIndex is the 'All' version of
FindStringSubmatchIndex; it returns a slice of all successive matches of
the expression, as defined by the 'All' description in the package
comment.
A return value of nil indicates no match.


<a id="example_Regexp_FindAllStringSubmatchIndex">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.FindAllSubmatch">func</a> (\*Regexp) [FindAllSubmatch](https://golang.org/src/regexp/regexp.go?s=34472:34533#L1120)
<pre>func (re *<a href="#Regexp">Regexp</a>) FindAllSubmatch(b []<a href="/pkg/builtin/#byte">byte</a>, n <a href="/pkg/builtin/#int">int</a>) [][][]<a href="/pkg/builtin/#byte">byte</a></pre>
FindAllSubmatch is the 'All' version of FindSubmatch; it returns a slice
of all successive matches of the expression, as defined by the 'All'
description in the package comment.
A return value of nil indicates no match.


<a id="example_Regexp_FindAllSubmatch">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.FindAllSubmatchIndex">func</a> (\*Regexp) [FindAllSubmatchIndex](https://golang.org/src/regexp/regexp.go?s=35154:35217#L1144)
<pre>func (re *<a href="#Regexp">Regexp</a>) FindAllSubmatchIndex(b []<a href="/pkg/builtin/#byte">byte</a>, n <a href="/pkg/builtin/#int">int</a>) [][]<a href="/pkg/builtin/#int">int</a></pre>
FindAllSubmatchIndex is the 'All' version of FindSubmatchIndex; it returns
a slice of all successive matches of the expression, as defined by the
'All' description in the package comment.
A return value of nil indicates no match.


<a id="example_Regexp_FindAllSubmatchIndex">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.FindIndex">func</a> (\*Regexp) [FindIndex](https://golang.org/src/regexp/regexp.go?s=24476:24525#L800)
<pre>func (re *<a href="#Regexp">Regexp</a>) FindIndex(b []<a href="/pkg/builtin/#byte">byte</a>) (loc []<a href="/pkg/builtin/#int">int</a>)</pre>
FindIndex returns a two-element slice of integers defining the location of
the leftmost match in b of the regular expression. The match itself is at
b[loc[0]:loc[1]].
A return value of nil indicates no match.


<a id="example_Regexp_FindIndex">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.FindReaderIndex">func</a> (\*Regexp) [FindReaderIndex](https://golang.org/src/regexp/regexp.go?s=25821:25883#L839)
<pre>func (re *<a href="#Regexp">Regexp</a>) FindReaderIndex(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#RuneReader">RuneReader</a>) (loc []<a href="/pkg/builtin/#int">int</a>)</pre>
FindReaderIndex returns a two-element slice of integers defining the
location of the leftmost match of the regular expression in text read from
the RuneReader. The match text was found in the input stream at
byte offset loc[0] through loc[1]-1.
A return value of nil indicates no match.




### <a id="Regexp.FindReaderSubmatchIndex">func</a> (\*Regexp) [FindReaderSubmatchIndex](https://golang.org/src/regexp/regexp.go?s=31955:32019#L1038)
<pre>func (re *<a href="#Regexp">Regexp</a>) FindReaderSubmatchIndex(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#RuneReader">RuneReader</a>) []<a href="/pkg/builtin/#int">int</a></pre>
FindReaderSubmatchIndex returns a slice holding the index pairs
identifying the leftmost match of the regular expression of text read by
the RuneReader, and the matches, if any, of its subexpressions, as defined
by the 'Submatch' and 'Index' descriptions in the package comment. A
return value of nil indicates no match.




### <a id="Regexp.FindString">func</a> (\*Regexp) [FindString](https://golang.org/src/regexp/regexp.go?s=24971:25016#L813)
<pre>func (re *<a href="#Regexp">Regexp</a>) FindString(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
FindString returns a string holding the text of the leftmost match in s of the regular
expression. If there is no match, the return value is an empty string,
but it will also be empty if the regular expression successfully matches
an empty string. Use FindStringIndex or FindStringSubmatch if it is
necessary to distinguish these cases.


<a id="example_Regexp_FindString">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.FindStringIndex">func</a> (\*Regexp) [FindStringIndex](https://golang.org/src/regexp/regexp.go?s=25369:25424#L826)
<pre>func (re *<a href="#Regexp">Regexp</a>) FindStringIndex(s <a href="/pkg/builtin/#string">string</a>) (loc []<a href="/pkg/builtin/#int">int</a>)</pre>
FindStringIndex returns a two-element slice of integers defining the
location of the leftmost match in s of the regular expression. The match
itself is at s[loc[0]:loc[1]].
A return value of nil indicates no match.


<a id="example_Regexp_FindStringIndex">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.FindStringSubmatch">func</a> (\*Regexp) [FindStringSubmatch](https://golang.org/src/regexp/regexp.go?s=30861:30916#L1009)
<pre>func (re *<a href="#Regexp">Regexp</a>) FindStringSubmatch(s <a href="/pkg/builtin/#string">string</a>) []<a href="/pkg/builtin/#string">string</a></pre>
FindStringSubmatch returns a slice of strings holding the text of the
leftmost match of the regular expression in s and the matches, if any, of
its subexpressions, as defined by the 'Submatch' description in the
package comment.
A return value of nil indicates no match.


<a id="example_Regexp_FindStringSubmatch">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.FindStringSubmatchIndex">func</a> (\*Regexp) [FindStringSubmatchIndex](https://golang.org/src/regexp/regexp.go?s=31490:31547#L1029)
<pre>func (re *<a href="#Regexp">Regexp</a>) FindStringSubmatchIndex(s <a href="/pkg/builtin/#string">string</a>) []<a href="/pkg/builtin/#int">int</a></pre>
FindStringSubmatchIndex returns a slice holding the index pairs
identifying the leftmost match of the regular expression in s and the
matches, if any, of its subexpressions, as defined by the 'Submatch' and
'Index' descriptions in the package comment.
A return value of nil indicates no match.




### <a id="Regexp.FindSubmatch">func</a> (\*Regexp) [FindSubmatch](https://golang.org/src/regexp/regexp.go?s=26257:26306#L852)
<pre>func (re *<a href="#Regexp">Regexp</a>) FindSubmatch(b []<a href="/pkg/builtin/#byte">byte</a>) [][]<a href="/pkg/builtin/#byte">byte</a></pre>
FindSubmatch returns a slice of slices holding the text of the leftmost
match of the regular expression in b and the matches, if any, of its
subexpressions, as defined by the 'Submatch' descriptions in the package
comment.
A return value of nil indicates no match.


<a id="example_Regexp_FindSubmatch">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.FindSubmatchIndex">func</a> (\*Regexp) [FindSubmatchIndex](https://golang.org/src/regexp/regexp.go?s=30453:30504#L1000)
<pre>func (re *<a href="#Regexp">Regexp</a>) FindSubmatchIndex(b []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#int">int</a></pre>
FindSubmatchIndex returns a slice holding the index pairs identifying the
leftmost match of the regular expression in b and the matches, if any, of
its subexpressions, as defined by the 'Submatch' and 'Index' descriptions
in the package comment.
A return value of nil indicates no match.




### <a id="Regexp.LiteralPrefix">func</a> (\*Regexp) [LiteralPrefix](https://golang.org/src/regexp/regexp.go?s=15158:15222#L484)
<pre>func (re *<a href="#Regexp">Regexp</a>) LiteralPrefix() (prefix <a href="/pkg/builtin/#string">string</a>, complete <a href="/pkg/builtin/#bool">bool</a>)</pre>
LiteralPrefix returns a literal string that must begin any match
of the regular expression re. It returns the boolean true if the
literal string comprises the entire regular expression.




### <a id="Regexp.Longest">func</a> (\*Regexp) [Longest](https://golang.org/src/regexp/regexp.go?s=7268:7295#L155)
<pre>func (re *<a href="#Regexp">Regexp</a>) Longest()</pre>
Longest makes future searches prefer the leftmost-longest match.
That is, when matching against text, the regexp returns a match that
begins as early as possible in the input (leftmost), and among those
it chooses a match that is as long as possible.
This method modifies the Regexp and may not be called concurrently
with any other methods.




### <a id="Regexp.Match">func</a> (\*Regexp) [Match](https://golang.org/src/regexp/regexp.go?s=15744:15782#L502)
<pre>func (re *<a href="#Regexp">Regexp</a>) Match(b []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Match reports whether the byte slice b
contains any match of the regular expression re.


<a id="example_Regexp_Match">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.MatchReader">func</a> (\*Regexp) [MatchReader](https://golang.org/src/regexp/regexp.go?s=15384:15435#L490)
<pre>func (re *<a href="#Regexp">Regexp</a>) MatchReader(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#RuneReader">RuneReader</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
MatchReader reports whether the text returned by the RuneReader
contains any match of the regular expression re.




### <a id="Regexp.MatchString">func</a> (\*Regexp) [MatchString](https://golang.org/src/regexp/regexp.go?s=15568:15612#L496)
<pre>func (re *<a href="#Regexp">Regexp</a>) MatchString(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
MatchString reports whether the string s
contains any match of the regular expression re.


<a id="example_Regexp_MatchString">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.NumSubexp">func</a> (\*Regexp) [NumSubexp](https://golang.org/src/regexp/regexp.go?s=11501:11534#L325)
<pre>func (re *<a href="#Regexp">Regexp</a>) NumSubexp() <a href="/pkg/builtin/#int">int</a></pre>
NumSubexp returns the number of parenthesized subexpressions in this Regexp.




### <a id="Regexp.ReplaceAll">func</a> (\*Regexp) [ReplaceAll](https://golang.org/src/regexp/regexp.go?s=20246:20299#L641)
<pre>func (re *<a href="#Regexp">Regexp</a>) ReplaceAll(src, repl []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
ReplaceAll returns a copy of src, replacing matches of the Regexp
with the replacement text repl. Inside repl, $ signs are interpreted as
in Expand, so for instance $1 represents the text of the first submatch.




### <a id="Regexp.ReplaceAllFunc">func</a> (\*Regexp) [ReplaceAllFunc](https://golang.org/src/regexp/regexp.go?s=21198:21275#L669)
<pre>func (re *<a href="#Regexp">Regexp</a>) ReplaceAllFunc(src []<a href="/pkg/builtin/#byte">byte</a>, repl func([]<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
ReplaceAllFunc returns a copy of src in which all matches of the
Regexp have been replaced by the return value of function repl applied
to the matched byte slice. The replacement returned by repl is substituted
directly, without using Expand.




### <a id="Regexp.ReplaceAllLiteral">func</a> (\*Regexp) [ReplaceAllLiteral](https://golang.org/src/regexp/regexp.go?s=20770:20830#L659)
<pre>func (re *<a href="#Regexp">Regexp</a>) ReplaceAllLiteral(src, repl []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
ReplaceAllLiteral returns a copy of src, replacing matches of the Regexp
with the replacement bytes repl. The replacement repl is substituted directly,
without using Expand.




### <a id="Regexp.ReplaceAllLiteralString">func</a> (\*Regexp) [ReplaceAllLiteralString](https://golang.org/src/regexp/regexp.go?s=17595:17661#L556)
<pre>func (re *<a href="#Regexp">Regexp</a>) ReplaceAllLiteralString(src, repl <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
ReplaceAllLiteralString returns a copy of src, replacing matches of the Regexp
with the replacement string repl. The replacement repl is substituted directly,
without using Expand.


<a id="example_Regexp_ReplaceAllLiteralString">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.ReplaceAllString">func</a> (\*Regexp) [ReplaceAllString](https://golang.org/src/regexp/regexp.go?s=17125:17184#L542)
<pre>func (re *<a href="#Regexp">Regexp</a>) ReplaceAllString(src, repl <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
ReplaceAllString returns a copy of src, replacing matches of the Regexp
with the replacement string repl. Inside repl, $ signs are interpreted as
in Expand, so for instance $1 represents the text of the first submatch.


<a id="example_Regexp_ReplaceAllString">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.ReplaceAllStringFunc">func</a> (\*Regexp) [ReplaceAllStringFunc](https://golang.org/src/regexp/regexp.go?s=18043:18126#L566)
<pre>func (re *<a href="#Regexp">Regexp</a>) ReplaceAllStringFunc(src <a href="/pkg/builtin/#string">string</a>, repl func(<a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
ReplaceAllStringFunc returns a copy of src in which all matches of the
Regexp have been replaced by the return value of function repl applied
to the matched substring. The replacement returned by repl is substituted
directly, without using Expand.


<a id="example_Regexp_ReplaceAllStringFunc">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.Split">func</a> (\*Regexp) [Split](https://golang.org/src/regexp/regexp.go?s=37375:37424#L1216)
<pre>func (re *<a href="#Regexp">Regexp</a>) Split(s <a href="/pkg/builtin/#string">string</a>, n <a href="/pkg/builtin/#int">int</a>) []<a href="/pkg/builtin/#string">string</a></pre>
Split slices s into substrings separated by the expression and returns a slice of
the substrings between those expression matches.

The slice returned by this method consists of all the substrings of s
not contained in the slice returned by FindAllString. When called on an expression
that contains no metacharacters, it is equivalent to strings.SplitN.

Example:


	s := regexp.MustCompile("a*").Split("abaabaccadaaae", 5)
	// s: ["", "b", "b", "c", "cadaaae"]

The count determines the number of substrings to return:


	n > 0: at most n substrings; the last substring will be the unsplit remainder.
	n == 0: the result is nil (zero substrings)
	n < 0: all substrings


<a id="example_Regexp_Split">Example</a>
```go
```

output:
```txt
```


### <a id="Regexp.String">func</a> (\*Regexp) [String](https://golang.org/src/regexp/regexp.go?s=4488:4521#L95)
<pre>func (re *<a href="#Regexp">Regexp</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns the source text used to compile the regular expression.




### <a id="Regexp.SubexpNames">func</a> (\*Regexp) [SubexpNames](https://golang.org/src/regexp/regexp.go?s=11895:11935#L334)
<pre>func (re *<a href="#Regexp">Regexp</a>) SubexpNames() []<a href="/pkg/builtin/#string">string</a></pre>
SubexpNames returns the names of the parenthesized subexpressions
in this Regexp. The name for the first sub-expression is names[1],
so that if m is a match slice, the name for m[i] is SubexpNames()[i].
Since the Regexp as a whole cannot be named, names[0] is always
the empty string. The slice should not be modified.


<a id="example_Regexp_SubexpNames">Example</a>
```go
```

output:
```txt
```





