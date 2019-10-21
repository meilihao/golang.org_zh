

# suffixarray
`import "index/suffixarray"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package suffixarray implements substring search in logarithmic time using
an in-memory suffix array.

Example use:


	// create index for some data
	index := suffixarray.New(data)
	
	// lookup byte slice s
	offsets1 := index.Lookup(s, -1) // the list of all indices where s occurs in data
	offsets2 := index.Lookup(s, 3)  // the list of at most 3 indices where s occurs in data




## <a id="pkg-index">Index</a>
* [type Index](#Index)
  * [func New(data []byte) *Index](#New)
  * [func (x *Index) Bytes() []byte](#Index.Bytes)
  * [func (x *Index) FindAllIndex(r *regexp.Regexp, n int) (result [][]int)](#Index.FindAllIndex)
  * [func (x *Index) Lookup(s []byte, n int) (result []int)](#Index.Lookup)
  * [func (x *Index) Read(r io.Reader) error](#Index.Read)
  * [func (x *Index) Write(w io.Writer) error](#Index.Write)


#### <a id="pkg-examples">Examples</a>
* [Index.Lookup](#example_Index_Lookup)


#### <a id="pkg-files">Package files</a>
[sais.go](https://golang.org/src/index/suffixarray/sais.go) [sais2.go](https://golang.org/src/index/suffixarray/sais2.go) [suffixarray.go](https://golang.org/src/index/suffixarray/suffixarray.go) 








## <a id="Index">type</a> [Index](https://golang.org/src/index/suffixarray/suffixarray.go?s=827:920#L25)
Index implements a suffix array for fast substring search.


<pre>type Index struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="New">func</a> [New](https://golang.org/src/index/suffixarray/suffixarray.go?s=1639:1667#L66)
<pre>func New(data []<a href="/pkg/builtin/#byte">byte</a>) *<a href="#Index">Index</a></pre>
New creates a new Index for data.
Index creation time is O(N) for N = len(data).






### <a id="Index.Bytes">func</a> (\*Index) [Bytes](https://golang.org/src/index/suffixarray/suffixarray.go?s=5282:5312#L224)
<pre>func (x *<a href="#Index">Index</a>) Bytes() []<a href="/pkg/builtin/#byte">byte</a></pre>
Bytes returns the data over which the index was created.
It must not be modified.




### <a id="Index.FindAllIndex">func</a> (\*Index) [FindAllIndex](https://golang.org/src/index/suffixarray/suffixarray.go?s=7062:7132#L280)
<pre>func (x *<a href="#Index">Index</a>) FindAllIndex(r *<a href="/pkg/regexp/">regexp</a>.<a href="/pkg/regexp/#Regexp">Regexp</a>, n <a href="/pkg/builtin/#int">int</a>) (result [][]<a href="/pkg/builtin/#int">int</a>)</pre>
FindAllIndex returns a sorted list of non-overlapping matches of the
regular expression r, where a match is a pair of indices specifying
the matched slice of x.Bytes(). If n < 0, all matches are returned
in successive order. Otherwise, at most n matches are returned and
they may not be successive. The result is nil if there are no matches,
or if n == 0.




### <a id="Index.Lookup">func</a> (\*Index) [Lookup](https://golang.org/src/index/suffixarray/suffixarray.go?s=6234:6288#L249)
<pre>func (x *<a href="#Index">Index</a>) Lookup(s []<a href="/pkg/builtin/#byte">byte</a>, n <a href="/pkg/builtin/#int">int</a>) (result []<a href="/pkg/builtin/#int">int</a>)</pre>
Lookup returns an unsorted list of at most n indices where the byte string s
occurs in the indexed data. If n < 0, all occurrences are returned.
The result is nil if s is empty, s is not found, or n == 0.
Lookup time is O(log(N)*len(s) + len(result)) where N is the
size of the indexed data.


<a id="example_Index_Lookup">Example</a>
```go
```

output:
```txt
```


### <a id="Index.Read">func</a> (\*Index) [Read](https://golang.org/src/index/suffixarray/suffixarray.go?s=3726:3765#L145)
<pre>func (x *<a href="#Index">Index</a>) Read(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) <a href="/pkg/builtin/#error">error</a></pre>
Read reads the index from r into x; x must not be nil.




### <a id="Index.Write">func</a> (\*Index) [Write](https://golang.org/src/index/suffixarray/suffixarray.go?s=4760:4800#L195)
<pre>func (x *<a href="#Index">Index</a>) Write(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>) <a href="/pkg/builtin/#error">error</a></pre>
Write writes the index x to w.







