

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

heap 为实现了heap.Interface接口的任意类型提供了 heap 操作. heap 是一颗树, 且每个节点是其子树的最小值.

树中最小的元素是root, 索引是0.

堆是实现优先级队列的常用方法. 要创建一个优先级队列, 需要实现 Heap 接口, 并在 Less 方法中以(负的)优先作为排序方式, 因此 Push 会添加元素, 而Pop会删除队列中优先级最高的元素. 例子中包括了这样的一个实现, 即文件example_pq_test.go包含完整的源代码.

<a id="example__intHeap">Example (IntHeap)</a>
<p>This example inserts several ints into an IntHeap, checks the minimum,
and removes them in order of priority.
</p>

本示例将多个int插入IntHeap中，检查最小值， 并按优先顺序将其删除.
```go
// This example demonstrates an integer heap built using the heap interface.
// 这个例子演示了一个使用 heap 接口构建的整数堆.
package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
// IntHeap 是元素为整数的最小堆.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    //
    // Push 和 Pop 使用 pointer作为receiver是因为它们不仅需要修改slice内容还要修改slice的长度.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
//
// 本示例将多个int插入IntHeap中，检查最小值， 并按优先顺序将其删除.
func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
```

output:
```txt
minimum: 1
1 2 3 5
```
<a id="example__priorityQueue">Example (PriorityQueue)</a>
<p>This example creates a PriorityQueue with some items, adds and manipulates an item,
and then removes the items in priority order.
</p>

本示例创建了一个包含若干元素的PriorityQueue， 再添加和操作一个元素, 然后按优先级顺序删除该元素.
```go
// This example demonstrates a priority queue built using the heap interface.
// 这个例子演示了一个使用 heap 接口构建的优先级队列.
package main

import (
	"container/heap"
	"fmt"
)

// An Item is something we manage in a priority queue.
// Item 是优先队列需要管理的元素
type Item struct {
	value    string // The value of the item; arbitrary. // 元素的值, 任意内容
	priority int    // The priority of the item in the queue. // 元素在队列中的优先级
    // The index is needed by update and is maintained by the heap.Interface methods.
    // 索引由 heap.Interface 方法维护, 也被 update 方法使用.
	index int // The index of the item in the heap. // 元素在heap中的索引
}

// A PriorityQueue implements heap.Interface and holds Items.
// PriorityQueue 类型实现了 heap.Interface 并存储Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    // We want Pop to give us the highest, not lowest, priority so we use greater than here.
    // 我们想要 Pop 函数获得最高的(优先级)，而不是最低的，所以我们需在这里使用大于号.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
// update 方法修改队列里元素的优先级和值.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
//
// 本示例创建了一个包含若干元素的PriorityQueue， 再添加和操作一个元素, 然后按优先级顺序删除该元素.
func main() {
    // Some items and their priorities.
    // 若干带有优先级的元素.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
    // establish the priority queue (heap) invariants.
    //
    // 创建一个优先级队列，把 items 放进去，然后建立优先队列（堆）.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

    // Insert a new item and then modify its priority.
    // 插入一个新的元素，然后修改它的优先级.
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

    // Take the items out; they arrive in decreasing priority order.
    // 取出所有元素，它们会按照优先级递减的顺序被取出
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
}
```

output:
```txt
05:orange 04:pear 03:banana 02:apple
```
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


在索引i的元素更改其值之后，Fix会重建堆顺序. 更改索引i处元素的值，然后调用Fix等价于调用Remove（h，i）后再紧跟着 Push 新值，但是花费更小. 当 n = h.Len()时, 复杂度是O(log n).



## <a id="Init">func</a> [Init](https://golang.org/src/container/heap/heap.go?s=1750:1772#L32)
<pre>func Init(h <a href="#Interface">Interface</a>)</pre>
Init establishes the heap invariants required by the other routines in this package.
Init is idempotent with respect to the heap invariants
and may be called whenever the heap invariants may have been invalidated.
The complexity is O(n) where n = h.Len().

Init 会建立 heap需要遵循的规则???. Init 是幂等的, 遵循head 规则. 只要heap的规则被打破就可以调用它. 当 n = h.Len()时, 复杂度是O(n).


## <a id="Pop">func</a> [Pop](https://golang.org/src/container/heap/heap.go?s=2190:2223#L50)
<pre>func Pop(h <a href="#Interface">Interface</a>) interface{}</pre>
Pop removes and returns the minimum element (according to Less) from the heap.
The complexity is O(log n) where n = h.Len().
Pop is equivalent to Remove(h, 0).

Pop 会 移除并返回 heap中最小的元素(通过 Less 方法). 但n = h.Len()时, 复杂度是O(log n). Pop等价于Remove(h, 0).


## <a id="Push">func</a> [Push](https://golang.org/src/container/heap/heap.go?s=1949:1986#L42)
<pre>func Push(h <a href="#Interface">Interface</a>, x interface{})</pre>
Push pushes the element x onto the heap.
The complexity is O(log n) where n = h.Len().

Push 将元素x 放入heap. 当 n = h.Len() 时, 复杂度是O(log n).

## <a id="Remove">func</a> [Remove](https://golang.org/src/container/heap/heap.go?s=2409:2452#L59)
<pre>func Remove(h <a href="#Interface">Interface</a>, i <a href="/pkg/builtin/#int">int</a>) interface{}</pre>
Remove removes and returns the element at index i from the heap.
The complexity is O(log n) where n = h.Len().

Remove 移除heap中给定索引i的元素. 当 n = h.Len() 时, 复杂度是O(log n) .



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

Interface 描述了使用本包时需要满足的要求. 可以用作最小堆的任何类型实现需要遵循以下规则 (Init调用后会遵循(该规则)或数据是空的或已排序)


	!h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()


注意, 接口中的Push和Pop是heap的实现者调用的. 为了添加或移除元素, 请使用heap.Push 和 heap.Pop.

<pre>type Interface interface {
    <a href="/pkg/sort/">sort</a>.<a href="/pkg/sort/#Interface">Interface</a>
    Push(x interface{}) <span class="comment">// add x as element Len()</span>
    Pop() interface{}   <span class="comment">// remove and return element Len() - 1.</span>
}</pre>















