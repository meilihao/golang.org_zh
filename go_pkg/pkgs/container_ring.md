

# ring
`import "container/ring"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package ring implements operations on circular lists.

ring 实现了环形(或循环)链表的操作.


## <a id="pkg-index">Index</a>
* [type Ring](#Ring)
  * [func New(n int) *Ring](#New)
  * [func (r *Ring) Do(f func(interface{}))](#Ring.Do)
  * [func (r *Ring) Len() int](#Ring.Len)
  * [func (r *Ring) Link(s *Ring) *Ring](#Ring.Link)
  * [func (r *Ring) Move(n int) *Ring](#Ring.Move)
  * [func (r *Ring) Next() *Ring](#Ring.Next)
  * [func (r *Ring) Prev() *Ring](#Ring.Prev)
  * [func (r *Ring) Unlink(n int) *Ring](#Ring.Unlink)


#### <a id="pkg-examples">Examples</a>
* [Ring.Do](#example_Ring_Do)
* [Ring.Len](#example_Ring_Len)
* [Ring.Link](#example_Ring_Link)
* [Ring.Move](#example_Ring_Move)
* [Ring.Next](#example_Ring_Next)
* [Ring.Prev](#example_Ring_Prev)
* [Ring.Unlink](#example_Ring_Unlink)


#### <a id="pkg-files">Package files</a>
[ring.go](https://golang.org/src/container/ring/ring.go) 








## <a id="Ring">type</a> [Ring](https://golang.org/src/container/ring/ring.go?s=523:633#L4)
A Ring is an element of a circular list, or ring.
Rings do not have a beginning or end; a pointer to any ring element
serves as reference to the entire ring. Empty rings are represented
as nil Ring pointers. The zero value for a Ring is a one-element
ring with a nil Value.


<pre>type Ring struct {
<span id="Ring.Value"></span>    Value interface{} <span class="comment">// for use by client; untouched by this library // 由用户使用. 本库不会涉及该值.</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="New">func</a> [New](https://golang.org/src/container/ring/ring.go?s=1389:1410#L52)
<pre>func New(n <a href="/pkg/builtin/#int">int</a>) *<a href="#Ring">Ring</a></pre>
New creates a ring of n elements.






### <a id="Ring.Do">func</a> (\*Ring) [Do](https://golang.org/src/container/ring/ring.go?s=3118:3156#L124)
<pre>func (r *<a href="#Ring">Ring</a>) Do(f func(interface{}))</pre>
Do calls function f on each element of the ring, in forward order.
The behavior of Do is undefined if f changes *r.


<a id="example_Ring_Do">Example</a>
```go
```

output:
```txt
```


### <a id="Ring.Len">func</a> (\*Ring) [Len](https://golang.org/src/container/ring/ring.go?s=2869:2893#L111)
<pre>func (r *<a href="#Ring">Ring</a>) Len() <a href="/pkg/builtin/#int">int</a></pre>
Len computes the number of elements in ring r.
It executes in time proportional to the number of elements.


<a id="example_Ring_Len">Example</a>
```go
```

output:
```txt
```


### <a id="Ring.Link">func</a> (\*Ring) [Link](https://golang.org/src/container/ring/ring.go?s=2221:2255#L83)
<pre>func (r *<a href="#Ring">Ring</a>) Link(s *<a href="#Ring">Ring</a>) *<a href="#Ring">Ring</a></pre>
Link connects ring r with ring s such that r.Next()
becomes s and returns the original value for r.Next().
r must not be empty.

If r and s point to the same ring, linking
them removes the elements between r and s from the ring.
The removed elements form a subring and the result is a
reference to that subring (if no elements were removed,
the result is still the original value for r.Next(),
and not nil).

If r and s point to different rings, linking
them creates a single ring with the elements of s inserted
after r. The result points to the element following the
last element of s after insertion.


<a id="example_Ring_Link">Example</a>
```go
```

output:
```txt
```


### <a id="Ring.Move">func</a> (\*Ring) [Move](https://golang.org/src/container/ring/ring.go?s=1146:1178#L34)
<pre>func (r *<a href="#Ring">Ring</a>) Move(n <a href="/pkg/builtin/#int">int</a>) *<a href="#Ring">Ring</a></pre>
Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0)
in the ring and returns that ring element. r must not be empty.


<a id="example_Ring_Move">Example</a>
```go
```

output:
```txt
```


### <a id="Ring.Next">func</a> (\*Ring) [Next](https://golang.org/src/container/ring/ring.go?s=762:789#L16)
<pre>func (r *<a href="#Ring">Ring</a>) Next() *<a href="#Ring">Ring</a></pre>
Next returns the next ring element. r must not be empty.


<a id="example_Ring_Next">Example</a>
```go
```

output:
```txt
```


### <a id="Ring.Prev">func</a> (\*Ring) [Prev](https://golang.org/src/container/ring/ring.go?s=915:942#L24)
<pre>func (r *<a href="#Ring">Ring</a>) Prev() *<a href="#Ring">Ring</a></pre>
Prev returns the previous ring element. r must not be empty.


<a id="example_Ring_Prev">Example</a>
```go
```

output:
```txt
```


### <a id="Ring.Unlink">func</a> (\*Ring) [Unlink](https://golang.org/src/container/ring/ring.go?s=2654:2688#L101)
<pre>func (r *<a href="#Ring">Ring</a>) Unlink(n <a href="/pkg/builtin/#int">int</a>) *<a href="#Ring">Ring</a></pre>
Unlink removes n % r.Len() elements from the ring r, starting
at r.Next(). If n % r.Len() == 0, r remains unchanged.
The result is the removed subring. r must not be empty.


<a id="example_Ring_Unlink">Example</a>
```go
```

output:
```txt
```





