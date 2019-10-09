

# heap
`import "container/heap"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package heap provides heap operations for any type that implements
heap.Interface. A heap is a tree with the property that each node is the
minimum-valued node in its subtree.

The minimum element in the tree is the root, at index 0.

A heap is a common way to implement a priority queue. To build a priority
queue, implement the Heap interface with the (negative) priority as the
ordering for the Less method, so Push adds items while Pop removes the
highest-priority item from the queue. The Examples include such an
implementation; the file example_pq_test.go has the complete source.


<a id="example__intHeap">Example (IntHeap)</a>
<p>This example inserts several ints into an IntHeap, checks the minimum,
and removes them in order of priority.
</p><a id="example__priorityQueue">Example (PriorityQueue)</a>
<p>This example creates a PriorityQueue with some items, adds and manipulates an item,
and then removes the items in priority order.
</p>

## <a id="pkg-index">Index</a>
* [func Fix(h Interface, i int)](#Fix)
* [func Init(h Interface)](#Init)
* [func Pop(h Interface) interface{}](#Pop)
* [func Push(h Interface, x interface{})](#Push)
* [func Remove(h Interface, i int) interface{}](#Remove)
* [type Interface](#Interface)


#### <a id="pkg-examples">Examples</a>
* [Package (IntHeap)](#example__intHeap)
* [Package (PriorityQueue)](#example__priorityQueue)


#### <a id="pkg-files">Package files</a>
[heap.go](https://golang.org/src/container/heap/heap.go) 






## <a id="Fix">func</a> [Fix](https://golang.org/src/container/heap/heap.go?s=2875:2903#L74)
<pre>func Fix(h <a href="#Interface">Interface</a>, i <a href="/pkg/builtin/#int">int</a>)</pre>
Fix re-establishes the heap ordering after the element at index i has changed its value.
Changing the value of the element at index i and then calling Fix is equivalent to,
but less expensive than, calling Remove(h, i) followed by a Push of the new value.
The complexity is O(log n) where n = h.Len().



## <a id="Init">func</a> [Init](https://golang.org/src/container/heap/heap.go?s=1750:1772#L32)
<pre>func Init(h <a href="#Interface">Interface</a>)</pre>
Init establishes the heap invariants required by the other routines in this package.
Init is idempotent with respect to the heap invariants
and may be called whenever the heap invariants may have been invalidated.
The complexity is O(n) where n = h.Len().



## <a id="Pop">func</a> [Pop](https://golang.org/src/container/heap/heap.go?s=2190:2223#L50)
<pre>func Pop(h <a href="#Interface">Interface</a>) interface{}</pre>
Pop removes and returns the minimum element (according to Less) from the heap.
The complexity is O(log n) where n = h.Len().
Pop is equivalent to Remove(h, 0).



## <a id="Push">func</a> [Push](https://golang.org/src/container/heap/heap.go?s=1949:1986#L42)
<pre>func Push(h <a href="#Interface">Interface</a>, x interface{})</pre>
Push pushes the element x onto the heap.
The complexity is O(log n) where n = h.Len().



## <a id="Remove">func</a> [Remove](https://golang.org/src/container/heap/heap.go?s=2409:2452#L59)
<pre>func Remove(h <a href="#Interface">Interface</a>, i <a href="/pkg/builtin/#int">int</a>) interface{}</pre>
Remove removes and returns the element at index i from the heap.
The complexity is O(log n) where n = h.Len().





## <a id="Interface">type</a> [Interface](https://golang.org/src/container/heap/heap.go?s=1328:1480#L22)
The Interface type describes the requirements
for a type using the routines in this package.
Any type that implements it may be used as a
min-heap with the following invariants (established after
Init has been called or if the data is empty or sorted):


	!h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()

Note that Push and Pop in this interface are for package heap's
implementation to call. To add and remove things from the heap,
use heap.Push and heap.Pop.


<pre>type Interface interface {
    <a href="/pkg/sort/">sort</a>.<a href="/pkg/sort/#Interface">Interface</a>
    Push(x interface{}) <span class="comment">// add x as element Len()</span>
    Pop() interface{}   <span class="comment">// remove and return element Len() - 1.</span>
}</pre>















