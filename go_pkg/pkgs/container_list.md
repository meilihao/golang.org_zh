

# list
`import "container/list"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package list implements a doubly linked list.

To iterate over a list (where l is a *List):

list 是双向链表的实现.

迭代一个链表(l是*List):

	for e := l.Front(); e != nil; e = e.Next() {
		// do something with e.Value
	}


<a id="example_">Example</a>
```go
package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
```

output:
```txt
1
2
3
4
```


## <a id="pkg-index">Index</a>
* [type Element](#Element)
  * [func (e *Element) Next() *Element](#Element.Next)
  * [func (e *Element) Prev() *Element](#Element.Prev)
* [type List](#List)
  * [func New() *List](#New)
  * [func (l *List) Back() *Element](#List.Back)
  * [func (l *List) Front() *Element](#List.Front)
  * [func (l *List) Init() *List](#List.Init)
  * [func (l *List) InsertAfter(v interface{}, mark *Element) *Element](#List.InsertAfter)
  * [func (l *List) InsertBefore(v interface{}, mark *Element) *Element](#List.InsertBefore)
  * [func (l *List) Len() int](#List.Len)
  * [func (l *List) MoveAfter(e, mark *Element)](#List.MoveAfter)
  * [func (l *List) MoveBefore(e, mark *Element)](#List.MoveBefore)
  * [func (l *List) MoveToBack(e *Element)](#List.MoveToBack)
  * [func (l *List) MoveToFront(e *Element)](#List.MoveToFront)
  * [func (l *List) PushBack(v interface{}) *Element](#List.PushBack)
  * [func (l *List) PushBackList(other *List)](#List.PushBackList)
  * [func (l *List) PushFront(v interface{}) *Element](#List.PushFront)
  * [func (l *List) PushFrontList(other *List)](#List.PushFrontList)
  * [func (l *List) Remove(e *Element) interface{}](#List.Remove)


#### <a id="pkg-examples">Examples</a>
* [Package](#example_)


#### <a id="pkg-files">Package files</a>
[list.go](https://golang.org/src/container/list/list.go) 








## <a id="Element">type</a> [Element](https://golang.org/src/container/list/list.go?s=406:874#L5)
Element is an element of a linked list.

Element 是链表中的一个元素.

<pre>type Element struct {

    <span class="comment">// The value stored with this element. // 存储在元素中的值.</span>
<span id="Element.Value"></span>    Value interface{}
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="Element.Next">func</a> (\*Element) [Next](https://golang.org/src/container/list/list.go?s=922:955#L21)
<pre>func (e *<a href="#Element">Element</a>) Next() *<a href="#Element">Element</a></pre>
Next returns the next list element or nil.

Next 返回 下一个节点或nil.


### <a id="Element.Prev">func</a> (\*Element) [Prev](https://golang.org/src/container/list/list.go?s=1091:1124#L29)
<pre>func (e *<a href="#Element">Element</a>) Prev() *<a href="#Element">Element</a></pre>
Prev returns the previous list element or nil.


Prev 返回上一个节点或nil.

## <a id="List">type</a> [List](https://golang.org/src/container/list/list.go?s=1309:1486#L38)
List represents a doubly linked list.
The zero value for List is an empty list ready to use.

List 表示双向链表. 其零值表示一个空的链表, 以待使用.

<pre>type List struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="New">func</a> [New](https://golang.org/src/container/list/list.go?s=1662:1678#L52)
<pre>func New() *<a href="#List">List</a></pre>
New returns an initialized list.


New 返回一个初始化好的list.



### <a id="List.Back">func</a> (\*List) [Back](https://golang.org/src/container/list/list.go?s=2063:2093#L67)
<pre>func (l *<a href="#List">List</a>) Back() *<a href="#Element">Element</a></pre>
Back returns the last element of list l or nil if the list is empty.

Back 返回list的最后一个节点, 如果l为空则返回nil.


### <a id="List.Front">func</a> (\*List) [Front](https://golang.org/src/container/list/list.go?s=1901:1932#L59)
<pre>func (l *<a href="#List">List</a>) Front() *<a href="#Element">Element</a></pre>
Front returns the first element of list l or nil if the list is empty.

Front 返回list中的第一个节点, 如果l为空则返回nil.


### <a id="List.Init">func</a> (\*List) [Init](https://golang.org/src/container/list/list.go?s=1526:1553#L44)
<pre>func (l *<a href="#List">List</a>) Init() *<a href="#List">List</a></pre>
Init initializes or clears list l.

Init 初始化或清空list.


### <a id="List.InsertAfter">func</a> (\*List) [InsertAfter](https://golang.org/src/container/list/list.go?s=4491:4556#L164)
<pre>func (l *<a href="#List">List</a>) InsertAfter(v interface{}, mark *<a href="#Element">Element</a>) *<a href="#Element">Element</a></pre>
InsertAfter inserts a new element e with value v immediately after mark and returns e.
If mark is not an element of l, the list is not modified.
The mark must not be nil.

InsertAfter 在mark节点后面插入一个值为v的新节点e, 并返回该节点. 如果mark不是list的元素, 那么list不会被改动. mark节点必须不为nil.


### <a id="List.InsertBefore">func</a> (\*List) [InsertBefore](https://golang.org/src/container/list/list.go?s=4109:4175#L153)
<pre>func (l *<a href="#List">List</a>) InsertBefore(v interface{}, mark *<a href="#Element">Element</a>) *<a href="#Element">Element</a></pre>
InsertBefore inserts a new element e with value v immediately before mark and returns e.
If mark is not an element of l, the list is not modified.
The mark must not be nil.


InsertBefore 在mark节点后面插入一个值为v的新节点e, 并返回该节点. 如果mark不是list的元素, 那么list不会被改动. mark节点必须不为nil.

### <a id="List.Len">func</a> (\*List) [Len](https://golang.org/src/container/list/list.go?s=1784:1808#L56)
<pre>func (l *<a href="#List">List</a>) Len() <a href="/pkg/builtin/#int">int</a></pre>
Len returns the number of elements of list l.
The complexity is O(1).


Len 返回 list 的节点数量. 复杂度是O(1).

### <a id="List.MoveAfter">func</a> (\*List) [MoveAfter](https://golang.org/src/container/list/list.go?s=5818:5860#L207)
<pre>func (l *<a href="#List">List</a>) MoveAfter(e, mark *<a href="#Element">Element</a>)</pre>
MoveAfter moves element e to its new position after mark.
If e or mark is not an element of l, or e == mark, the list is not modified.
The element and mark must not be nil.


MoveAfter 将节点e移动到mark后面. 如果e或mark不是list的节点, 或者 e == mark, 那么list不会被改动. e和mark必须不为nil.

### <a id="List.MoveBefore">func</a> (\*List) [MoveBefore](https://golang.org/src/container/list/list.go?s=5504:5547#L197)
<pre>func (l *<a href="#List">List</a>) MoveBefore(e, mark *<a href="#Element">Element</a>)</pre>
MoveBefore moves element e to its new position before mark.
If e or mark is not an element of l, or e == mark, the list is not modified.
The element and mark must not be nil.




### <a id="List.MoveToBack">func</a> (\*List) [MoveToBack](https://golang.org/src/container/list/list.go?s=5146:5183#L186)
<pre>func (l *<a href="#List">List</a>) MoveToBack(e *<a href="#Element">Element</a>)</pre>
MoveToBack moves element e to the back of list l.
If e is not an element of l, the list is not modified.
The element must not be nil.


MoveToBack 将节点e移动到list的末尾. 如果e不是list的节点, 那么list不会被改动. e必须不为nil.


### <a id="List.MoveToFront">func</a> (\*List) [MoveToFront](https://golang.org/src/container/list/list.go?s=4832:4870#L175)
<pre>func (l *<a href="#List">List</a>) MoveToFront(e *<a href="#Element">Element</a>)</pre>
MoveToFront moves element e to the front of list l.
If e is not an element of l, the list is not modified.
The element must not be nil.

MoveToFront 将节点e移动到list的开头. 如果e不是list的节点, 那么list不会被改动. e必须不为nil.


### <a id="List.PushBack">func</a> (\*List) [PushBack](https://golang.org/src/container/list/list.go?s=3822:3869#L145)
<pre>func (l *<a href="#List">List</a>) PushBack(v interface{}) *<a href="#Element">Element</a></pre>
PushBack inserts a new element e with value v at the back of list l and returns e.


PushBack 在list尾部插入一个值为v的新节点e, 并返回该节点.

### <a id="List.PushBackList">func</a> (\*List) [PushBackList](https://golang.org/src/container/list/list.go?s=6079:6119#L216)
<pre>func (l *<a href="#List">List</a>) PushBackList(other *<a href="#List">List</a>)</pre>
PushBackList inserts a copy of an other list at the back of list l.
The lists l and other may be the same. They must not be nil.


PushBackList 在list尾部插入另一个list的copy. l和other可能相同. 它们必须不为nil.

### <a id="List.PushFront">func</a> (\*List) [PushFront](https://golang.org/src/container/list/list.go?s=3634:3682#L139)
<pre>func (l *<a href="#List">List</a>) PushFront(v interface{}) *<a href="#Element">Element</a></pre>
PushFront inserts a new element e with value v at the front of list l and returns e.

PushFront 在list开头插入一个值为v的新节点e, 并返回该节点.


### <a id="List.PushFrontList">func</a> (\*List) [PushFrontList](https://golang.org/src/container/list/list.go?s=6388:6429#L225)
<pre>func (l *<a href="#List">List</a>) PushFrontList(other *<a href="#List">List</a>)</pre>
PushFrontList inserts a copy of an other list at the front of list l.
The lists l and other may be the same. They must not be nil.


PushFrontList 在list头部插入另一个list的copy. l和other可能相同. 它们必须不为nil.

### <a id="List.Remove">func</a> (\*List) [Remove](https://golang.org/src/container/list/list.go?s=3306:3351#L129)
<pre>func (l *<a href="#List">List</a>) Remove(e *<a href="#Element">Element</a>) interface{}</pre>
Remove removes e from l if e is an element of list l.
It returns the element value e.Value.
The element must not be nil.


Remove 移除e, 前提是e是list的节点. 它会返回该节点的Value. 该节点必须不为nil.




