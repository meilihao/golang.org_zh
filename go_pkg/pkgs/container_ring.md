

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

Ring 是 环形链表中的一个元素或环. 环形链表没有起点或终点; 指向任意一个环的指针都可用作整个环的引用. 使用值为 nil 的 Ring 指针来表示空环形链表. Ring的零值(它的Value是nil)是具有一个元素的环形链表.


<pre>type Ring struct {
<span id="Ring.Value"></span>    Value interface{} <span class="comment">// for use by client; untouched by this library // 由用户使用. 本库不会涉及该值.</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="New">func</a> [New](https://golang.org/src/container/ring/ring.go?s=1389:1410#L52)
<pre>func New(n <a href="/pkg/builtin/#int">int</a>) *<a href="#Ring">Ring</a></pre>
New creates a ring of n elements.


New 创建一个有n个元素的环形链表.



### <a id="Ring.Do">func</a> (\*Ring) [Do](https://golang.org/src/container/ring/ring.go?s=3118:3156#L124)
<pre>func (r *<a href="#Ring">Ring</a>) Do(f func(interface{}))</pre>
Do calls function f on each element of the ring, in forward order.
The behavior of Do is undefined if f changes *r.

Do 对环形链表的每个元素都调用函数 f. 如果函数 f 改变了环形链表，Do 方法的结果将无法预测.

<a id="example_Ring_Do">Example</a>
```go
ackage main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
    // 用若干整数来初始化环形链表
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring and print its contents
    // 迭代环形链表并打印内容
	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

}
```

output:
```txt
0
1
2
3
4
```


### <a id="Ring.Len">func</a> (\*Ring) [Len](https://golang.org/src/container/ring/ring.go?s=2869:2893#L111)
<pre>func (r *<a href="#Ring">Ring</a>) Len() <a href="/pkg/builtin/#int">int</a></pre>
Len computes the number of elements in ring r.
It executes in time proportional to the number of elements.

Len 返回 r 的元素个数. 函数的执行时间与元素个数成正比.

<a id="example_Ring_Len">Example</a>
```go
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 4
	r := ring.New(4)

	// Print out its length
	fmt.Println(r.Len())

}
```

output:
```txt
4
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

Link 在r后面连接s即r.Next()指向了s, 返回原有的r.Next()的值. r必须不为空.

如果 r和s指向相同的环形链表, 连接时会删除r和s之间的元素. 被移除的元素会形成一个子环形链表，函数返回的是对该子环形链表的引用(如果没有删除任何元素, 则结果是原有的 r.Next()的值且不为nil).

如果r和s指向不同的环形链表, 通过在r后面插入s形成新的环形链表. 返回插入s后的最后一个元素之后的元素.

<a id="example_Ring_Link">Example</a>
```go
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create two rings, r and s, of size 2
	r := ring.New(2)
	s := ring.New(2)

	// Get the length of the ring
	lr := r.Len()
	ls := s.Len()

	// Initialize r with 0s
	for i := 0; i < lr; i++ {
		r.Value = 0
		r = r.Next()
	}

	// Initialize s with 1s
	for j := 0; j < ls; j++ {
		s.Value = 1
		s = s.Next()
	}

	// Link ring r and ring s
	rs := r.Link(s)

	// Iterate through the combined ring and print its contents
	rs.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

}
```

output:
```txt
0
0
1
1
```


### <a id="Ring.Move">func</a> (\*Ring) [Move](https://golang.org/src/container/ring/ring.go?s=1146:1178#L34)
<pre>func (r *<a href="#Ring">Ring</a>) Move(n <a href="/pkg/builtin/#int">int</a>) *<a href="#Ring">Ring</a></pre>
Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0)
in the ring and returns that ring element. r must not be empty.

Move 把当前r的指针移动 n % r.Len() 次, n<0时向后移动, n>=0时向前移动.

r必须不为空.

<a id="example_Ring_Move">Example</a>
```go
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Move the pointer forward by three steps
	r = r.Move(3)

	// Iterate through the ring and print its contents
	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

}
```

output:
```txt
3
4
0
1
2
```


### <a id="Ring.Next">func</a> (\*Ring) [Next](https://golang.org/src/container/ring/ring.go?s=762:789#L16)
<pre>func (r *<a href="#Ring">Ring</a>) Next() *<a href="#Ring">Ring</a></pre>
Next returns the next ring element. r must not be empty.

Next r移动到下一个元素. r必须不为空.

<a id="example_Ring_Next">Example</a>
```go
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring and print its contents
	for j := 0; j < n; j++ {
		fmt.Println(r.Value)
		r = r.Next()
	}

}
```

output:
```txt
0
1
2
3
4
```


### <a id="Ring.Prev">func</a> (\*Ring) [Prev](https://golang.org/src/container/ring/ring.go?s=915:942#L24)
<pre>func (r *<a href="#Ring">Ring</a>) Prev() *<a href="#Ring">Ring</a></pre>
Prev returns the previous ring element. r must not be empty.

Prev r移动到前一个元素. r必须不为空.

<a id="example_Ring_Prev">Example</a>
```go
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring backwards and print its contents
	for j := 0; j < n; j++ {
		r = r.Prev()
		fmt.Println(r.Value)
	}

}
```

output:
```txt
4
3
2
1
0
```


### <a id="Ring.Unlink">func</a> (\*Ring) [Unlink](https://golang.org/src/container/ring/ring.go?s=2654:2688#L101)
<pre>func (r *<a href="#Ring">Ring</a>) Unlink(n <a href="/pkg/builtin/#int">int</a>) *<a href="#Ring">Ring</a></pre>
Unlink removes n % r.Len() elements from the ring r, starting
at r.Next(). If n % r.Len() == 0, r remains unchanged.
The result is the removed subring. r must not be empty.

Unlink 从r.Next()开始从环形链表中移除n % r.Len()个元素. 如果n % r.Len() ==0, 则不操作. 返回移除部分元素后的子环形链表. r必须不为空.

<a id="example_Ring_Unlink">Example</a>
```go
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 6
	r := ring.New(6)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Unlink three elements from r, starting from r.Next()
	r.Unlink(3)

	// Iterate through the remaining ring and print its contents
	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

}
```

output:
```txt
0
4
5
```





