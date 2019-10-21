

# sort
`import "sort"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package sort provides primitives for sorting slices and user-defined
collections.


<a id="example_">Example</a>
```go
```

output:
```txt
```
<a id="example__sortKeys">Example (SortKeys)</a>
```go
```

output:
```txt
```
<p>ExampleSortKeys demonstrates a technique for sorting a struct type using programmable sort criteria.
</p><a id="example__sortMultiKeys">Example (SortMultiKeys)</a>
```go
```

output:
```txt
```
<p>ExampleMultiKeys demonstrates a technique for sorting a struct type using different
sets of multiple fields in the comparison. We chain together &#34;Less&#34; functions, each of
which compares a single field.
</p><a id="example__sortWrapper">Example (SortWrapper)</a>
```go
```

output:
```txt
```


## <a id="pkg-index">Index</a>
* [func Float64s(a []float64)](#Float64s)
* [func Float64sAreSorted(a []float64) bool](#Float64sAreSorted)
* [func Ints(a []int)](#Ints)
* [func IntsAreSorted(a []int) bool](#IntsAreSorted)
* [func IsSorted(data Interface) bool](#IsSorted)
* [func Search(n int, f func(int) bool) int](#Search)
* [func SearchFloat64s(a []float64, x float64) int](#SearchFloat64s)
* [func SearchInts(a []int, x int) int](#SearchInts)
* [func SearchStrings(a []string, x string) int](#SearchStrings)
* [func Slice(slice interface{}, less func(i, j int) bool)](#Slice)
* [func SliceIsSorted(slice interface{}, less func(i, j int) bool) bool](#SliceIsSorted)
* [func SliceStable(slice interface{}, less func(i, j int) bool)](#SliceStable)
* [func Sort(data Interface)](#Sort)
* [func Stable(data Interface)](#Stable)
* [func Strings(a []string)](#Strings)
* [func StringsAreSorted(a []string) bool](#StringsAreSorted)
* [type Float64Slice](#Float64Slice)
  * [func (p Float64Slice) Len() int](#Float64Slice.Len)
  * [func (p Float64Slice) Less(i, j int) bool](#Float64Slice.Less)
  * [func (p Float64Slice) Search(x float64) int](#Float64Slice.Search)
  * [func (p Float64Slice) Sort()](#Float64Slice.Sort)
  * [func (p Float64Slice) Swap(i, j int)](#Float64Slice.Swap)
* [type IntSlice](#IntSlice)
  * [func (p IntSlice) Len() int](#IntSlice.Len)
  * [func (p IntSlice) Less(i, j int) bool](#IntSlice.Less)
  * [func (p IntSlice) Search(x int) int](#IntSlice.Search)
  * [func (p IntSlice) Sort()](#IntSlice.Sort)
  * [func (p IntSlice) Swap(i, j int)](#IntSlice.Swap)
* [type Interface](#Interface)
  * [func Reverse(data Interface) Interface](#Reverse)
* [type StringSlice](#StringSlice)
  * [func (p StringSlice) Len() int](#StringSlice.Len)
  * [func (p StringSlice) Less(i, j int) bool](#StringSlice.Less)
  * [func (p StringSlice) Search(x string) int](#StringSlice.Search)
  * [func (p StringSlice) Sort()](#StringSlice.Sort)
  * [func (p StringSlice) Swap(i, j int)](#StringSlice.Swap)


#### <a id="pkg-examples">Examples</a>
* [Package](#example_)
* [Float64s](#example_Float64s)
* [Float64sAreSorted](#example_Float64sAreSorted)
* [Ints](#example_Ints)
* [IntsAreSorted](#example_IntsAreSorted)
* [Reverse](#example_Reverse)
* [Search](#example_Search)
* [Search (DescendingOrder)](#example_Search_descendingOrder)
* [Slice](#example_Slice)
* [SliceStable](#example_SliceStable)
* [Strings](#example_Strings)
* [Package (SortKeys)](#example__sortKeys)
* [Package (SortMultiKeys)](#example__sortMultiKeys)
* [Package (SortWrapper)](#example__sortWrapper)


#### <a id="pkg-files">Package files</a>
[search.go](https://golang.org/src/sort/search.go) [slice.go](https://golang.org/src/sort/slice.go) [slice_go113.go](https://golang.org/src/sort/slice_go113.go) [sort.go](https://golang.org/src/sort/sort.go) [zfuncversion.go](https://golang.org/src/sort/zfuncversion.go) 






## <a id="Float64s">func</a> [Float64s](https://golang.org/src/sort/sort.go?s=8174:8200#L301)
<pre>func Float64s(a []<a href="/pkg/builtin/#float64">float64</a>)</pre>
Float64s sorts a slice of float64s in increasing order
(not-a-number values are treated as less than other values).


<a id="example_Float64s">Example</a>
```go
```

output:
```txt
```

## <a id="Float64sAreSorted">func</a> [Float64sAreSorted](https://golang.org/src/sort/sort.go?s=8630:8670#L311)
<pre>func Float64sAreSorted(a []<a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Float64sAreSorted tests whether a slice of float64s is sorted in increasing order
(not-a-number values are treated as less than other values).


<a id="example_Float64sAreSorted">Example</a>
```go
```

output:
```txt
```

## <a id="Ints">func</a> [Ints](https://golang.org/src/sort/sort.go?s=8010:8028#L297)
<pre>func Ints(a []<a href="/pkg/builtin/#int">int</a>)</pre>
Ints sorts a slice of ints in increasing order.


<a id="example_Ints">Example</a>
```go
```

output:
```txt
```

## <a id="IntsAreSorted">func</a> [IntsAreSorted](https://golang.org/src/sort/sort.go?s=8414:8446#L307)
<pre>func IntsAreSorted(a []<a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IntsAreSorted tests whether a slice of ints is sorted in increasing order.


<a id="example_IntsAreSorted">Example</a>
```go
```

output:
```txt
```

## <a id="IsSorted">func</a> [IsSorted](https://golang.org/src/sort/sort.go?s=6374:6408#L246)
<pre>func IsSorted(data <a href="#Interface">Interface</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsSorted reports whether data is sorted.



## <a id="Search">func</a> [Search](https://golang.org/src/sort/search.go?s=2246:2286#L49)
<pre>func Search(n <a href="/pkg/builtin/#int">int</a>, f func(<a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#bool">bool</a>) <a href="/pkg/builtin/#int">int</a></pre>
Search uses binary search to find and return the smallest index i
in [0, n) at which f(i) is true, assuming that on the range [0, n),
f(i) == true implies f(i+1) == true. That is, Search requires that
f is false for some (possibly empty) prefix of the input range [0, n)
and then true for the (possibly empty) remainder; Search returns
the first true index. If there is no such index, Search returns n.
(Note that the "not found" return value is not -1 as in, for instance,
strings.Index.)
Search calls f(i) only for i in the range [0, n).

A common use of Search is to find the index i for a value x in
a sorted, indexable data structure such as an array or slice.
In this case, the argument f, typically a closure, captures the value
to be searched for, and how the data structure is indexed and
ordered.

For instance, given a slice data sorted in ascending order,
the call Search(len(data), func(i int) bool { return data[i] >= 23 })
returns the smallest index i such that data[i] >= 23. If the caller
wants to find whether 23 is in the slice, it must test data[i] == 23
separately.

Searching data sorted in descending order would use the <=
operator instead of the >= operator.

To complete the example above, the following code tries to find the value
x in an integer slice data sorted in ascending order:


	x := 23
	i := sort.Search(len(data), func(i int) bool { return data[i] >= x })
	if i < len(data) && data[i] == x {
		// x is present at data[i]
	} else {
		// x is not present in data,
		// but i is the index where it would be inserted.
	}

As a more whimsical example, this program guesses your number:


	func GuessingGame() {
		var s string
		fmt.Printf("Pick an integer from 0 to 100.\n")
		answer := sort.Search(100, func(i int) bool {
			fmt.Printf("Is your number <= %d? ", i)
			fmt.Scanf("%s", &s)
			return s != "" && s[0] == 'y'
		})
		fmt.Printf("Your number is %d.\n", answer)
	}


<a id="example_Search">Example</a>
```go
```

output:
```txt
```
<p>This example demonstrates searching a list sorted in ascending order.
</p><a id="example_Search_descendingOrder">Example (DescendingOrder)</a>
```go
```

output:
```txt
```
<p>This example demonstrates searching a list sorted in descending order.
The approach is the same as searching a list in ascending order,
but with the condition inverted.
</p>
## <a id="SearchFloat64s">func</a> [SearchFloat64s](https://golang.org/src/sort/search.go?s=3317:3364#L82)
<pre>func SearchFloat64s(a []<a href="/pkg/builtin/#float64">float64</a>, x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#int">int</a></pre>
SearchFloat64s searches for x in a sorted slice of float64s and returns the index
as specified by Search. The return value is the index to insert x if x is not
present (it could be len(a)).
The slice must be sorted in ascending order.



## <a id="SearchInts">func</a> [SearchInts](https://golang.org/src/sort/search.go?s=2964:2999#L73)
<pre>func SearchInts(a []<a href="/pkg/builtin/#int">int</a>, x <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
SearchInts searches for x in a sorted slice of ints and returns the index
as specified by Search. The return value is the index to insert x if x is
not present (it could be len(a)).
The slice must be sorted in ascending order.



## <a id="SearchStrings">func</a> [SearchStrings](https://golang.org/src/sort/search.go?s=3680:3724#L91)
<pre>func SearchStrings(a []<a href="/pkg/builtin/#string">string</a>, x <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#int">int</a></pre>
SearchStrings searches for x in a sorted slice of strings and returns the index
as specified by Search. The return value is the index to insert x if x is not
present (it could be len(a)).
The slice must be sorted in ascending order.



## <a id="Slice">func</a> [Slice](https://golang.org/src/sort/slice.go?s=396:451#L3)
<pre>func Slice(slice interface{}, less func(i, j <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#bool">bool</a>)</pre>
Slice sorts the provided slice given the provided less function.

The sort is not guaranteed to be stable. For a stable sort, use
SliceStable.

The function panics if the provided interface is not a slice.


<a id="example_Slice">Example</a>
```go
```

output:
```txt
```

## <a id="SliceIsSorted">func</a> [SliceIsSorted](https://golang.org/src/sort/slice.go?s=1090:1158#L23)
<pre>func SliceIsSorted(slice interface{}, less func(i, j <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#bool">bool</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
SliceIsSorted tests whether a slice is sorted.

The function panics if the provided interface is not a slice.



## <a id="SliceStable">func</a> [SliceStable](https://golang.org/src/sort/slice.go?s=800:861#L14)
<pre>func SliceStable(slice interface{}, less func(i, j <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#bool">bool</a>)</pre>
SliceStable sorts the provided slice given the provided less
function while keeping the original order of equal elements.

The function panics if the provided interface is not a slice.


<a id="example_SliceStable">Example</a>
```go
```

output:
```txt
```

## <a id="Sort">func</a> [Sort](https://golang.org/src/sort/sort.go?s=5416:5441#L206)
<pre>func Sort(data <a href="#Interface">Interface</a>)</pre>
Sort sorts data.
It makes one call to data.Len to determine n, and O(n*log(n)) calls to
data.Less and data.Swap. The sort is not guaranteed to be stable.



## <a id="Stable">func</a> [Stable](https://golang.org/src/sort/sort.go?s=10596:10623#L346)
<pre>func Stable(data <a href="#Interface">Interface</a>)</pre>
Stable sorts data while keeping the original order of equal elements.

It makes one call to data.Len to determine n, O(n*log(n)) calls to
data.Less and O(n*log(n)*log(n)) calls to data.Swap.



## <a id="Strings">func</a> [Strings](https://golang.org/src/sort/sort.go?s=8285:8309#L304)
<pre>func Strings(a []<a href="/pkg/builtin/#string">string</a>)</pre>
Strings sorts a slice of strings in increasing order.


<a id="example_Strings">Example</a>
```go
```

output:
```txt
```

## <a id="StringsAreSorted">func</a> [StringsAreSorted](https://golang.org/src/sort/sort.go?s=8793:8831#L314)
<pre>func StringsAreSorted(a []<a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
StringsAreSorted tests whether a slice of strings is sorted in increasing order.





## <a id="Float64Slice">type</a> [Float64Slice](https://golang.org/src/sort/sort.go?s=7078:7105#L270)
Float64Slice attaches the methods of Interface to []float64, sorting in increasing order
(not-a-number values are treated as less than other values).


<pre>type Float64Slice []<a href="/pkg/builtin/#float64">float64</a></pre>











### <a id="Float64Slice.Len">func</a> (Float64Slice) [Len](https://golang.org/src/sort/sort.go?s=7107:7138#L272)
<pre>func (p <a href="#Float64Slice">Float64Slice</a>) Len() <a href="/pkg/builtin/#int">int</a></pre>



### <a id="Float64Slice.Less">func</a> (Float64Slice) [Less](https://golang.org/src/sort/sort.go?s=7167:7208#L273)
<pre>func (p <a href="#Float64Slice">Float64Slice</a>) Less(i, j <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#bool">bool</a></pre>



### <a id="Float64Slice.Search">func</a> (Float64Slice) [Search](https://golang.org/src/sort/search.go?s=4011:4054#L99)
<pre>func (p <a href="#Float64Slice">Float64Slice</a>) Search(x <a href="/pkg/builtin/#float64">float64</a>) <a href="/pkg/builtin/#int">int</a></pre>
Search returns the result of applying SearchFloat64s to the receiver and x.




### <a id="Float64Slice.Sort">func</a> (Float64Slice) [Sort](https://golang.org/src/sort/sort.go?s=7490:7518#L282)
<pre>func (p <a href="#Float64Slice">Float64Slice</a>) Sort()</pre>
Sort is a convenience method.




### <a id="Float64Slice.Swap">func</a> (Float64Slice) [Swap](https://golang.org/src/sort/sort.go?s=7263:7299#L274)
<pre>func (p <a href="#Float64Slice">Float64Slice</a>) Swap(i, j <a href="/pkg/builtin/#int">int</a>)</pre>



## <a id="IntSlice">type</a> [IntSlice](https://golang.org/src/sort/sort.go?s=6646:6665#L259)
IntSlice attaches the methods of Interface to []int, sorting in increasing order.


<pre>type IntSlice []<a href="/pkg/builtin/#int">int</a></pre>











### <a id="IntSlice.Len">func</a> (IntSlice) [Len](https://golang.org/src/sort/sort.go?s=6667:6694#L261)
<pre>func (p <a href="#IntSlice">IntSlice</a>) Len() <a href="/pkg/builtin/#int">int</a></pre>



### <a id="IntSlice.Less">func</a> (IntSlice) [Less](https://golang.org/src/sort/sort.go?s=6723:6760#L262)
<pre>func (p <a href="#IntSlice">IntSlice</a>) Less(i, j <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#bool">bool</a></pre>



### <a id="IntSlice.Search">func</a> (IntSlice) [Search](https://golang.org/src/sort/search.go?s=3867:3902#L96)
<pre>func (p <a href="#IntSlice">IntSlice</a>) Search(x <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
Search returns the result of applying SearchInts to the receiver and x.




### <a id="IntSlice.Sort">func</a> (IntSlice) [Sort](https://golang.org/src/sort/sort.go?s=6884:6908#L266)
<pre>func (p <a href="#IntSlice">IntSlice</a>) Sort()</pre>
Sort is a convenience method.




### <a id="IntSlice.Swap">func</a> (IntSlice) [Swap](https://golang.org/src/sort/sort.go?s=6784:6816#L263)
<pre>func (p <a href="#IntSlice">IntSlice</a>) Swap(i, j <a href="/pkg/builtin/#int">int</a>)</pre>



## <a id="Interface">type</a> [Interface](https://golang.org/src/sort/sort.go?s=505:783#L4)
A type, typically a collection, that satisfies sort.Interface can be
sorted by the routines in this package. The methods require that the
elements of the collection be enumerated by an integer index.


<pre>type Interface interface {
    <span class="comment">// Len is the number of elements in the collection.</span>
    Len() <a href="/pkg/builtin/#int">int</a>
    <span class="comment">// Less reports whether the element with</span>
    <span class="comment">// index i should sort before the element with index j.</span>
    Less(i, j <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#bool">bool</a>
    <span class="comment">// Swap swaps the elements with indexes i and j.</span>
    Swap(i, j <a href="/pkg/builtin/#int">int</a>)
}</pre>









### <a id="Reverse">func</a> [Reverse](https://golang.org/src/sort/sort.go?s=6263:6301#L241)
<pre>func Reverse(data <a href="#Interface">Interface</a>) <a href="#Interface">Interface</a></pre>
Reverse returns the reverse order for data.


<a id="example_Reverse">Example</a>
```go
```

output:
```txt
```




## <a id="StringSlice">type</a> [StringSlice](https://golang.org/src/sort/sort.go?s=7623:7648#L285)
StringSlice attaches the methods of Interface to []string, sorting in increasing order.


<pre>type StringSlice []<a href="/pkg/builtin/#string">string</a></pre>











### <a id="StringSlice.Len">func</a> (StringSlice) [Len](https://golang.org/src/sort/sort.go?s=7650:7680#L287)
<pre>func (p <a href="#StringSlice">StringSlice</a>) Len() <a href="/pkg/builtin/#int">int</a></pre>



### <a id="StringSlice.Less">func</a> (StringSlice) [Less](https://golang.org/src/sort/sort.go?s=7709:7749#L288)
<pre>func (p <a href="#StringSlice">StringSlice</a>) Less(i, j <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#bool">bool</a></pre>



### <a id="StringSlice.Search">func</a> (StringSlice) [Search](https://golang.org/src/sort/search.go?s=4166:4207#L102)
<pre>func (p <a href="#StringSlice">StringSlice</a>) Search(x <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#int">int</a></pre>
Search returns the result of applying SearchStrings to the receiver and x.




### <a id="StringSlice.Sort">func</a> (StringSlice) [Sort](https://golang.org/src/sort/sort.go?s=7876:7903#L292)
<pre>func (p <a href="#StringSlice">StringSlice</a>) Sort()</pre>
Sort is a convenience method.




### <a id="StringSlice.Swap">func</a> (StringSlice) [Swap](https://golang.org/src/sort/sort.go?s=7773:7808#L289)
<pre>func (p <a href="#StringSlice">StringSlice</a>) Swap(i, j <a href="/pkg/builtin/#int">int</a>)</pre>






