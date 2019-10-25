

# context
`import "context"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package context defines the Context type, which carries deadlines,
cancellation signals, and other request-scoped values across API boundaries
and between processes.

Incoming requests to a server should create a Context, and outgoing
calls to servers should accept a Context. The chain of function
calls between them must propagate the Context, optionally replacing
it with a derived Context created using WithCancel, WithDeadline,
WithTimeout, or WithValue. When a Context is canceled, all
Contexts derived from it are also canceled.

The WithCancel, WithDeadline, and WithTimeout functions take a
Context (the parent) and return a derived Context (the child) and a
CancelFunc. Calling the CancelFunc cancels the child and its
children, removes the parent's reference to the child, and stops
any associated timers. Failing to call the CancelFunc leaks the
child and its children until the parent is canceled or the timer
fires. The go vet tool checks that CancelFuncs are used on all
control-flow paths.

Programs that use Contexts should follow these rules to keep interfaces
consistent across packages and enable static analysis tools to check context
propagation:

Do not store Contexts inside a struct type; instead, pass a Context
explicitly to each function that needs it. The Context should be the first
parameter, typically named ctx:


	func DoSomething(ctx context.Context, arg Arg) error {
		// ... use ctx ...
	}

Do not pass a nil Context, even if a function permits it. Pass context.TODO
if you are unsure about which Context to use.

Use context Values only for request-scoped data that transits processes and
APIs, not for passing optional parameters to functions.

The same Context may be passed to functions running in different goroutines;
Contexts are safe for simultaneous use by multiple goroutines.

See <a href="https://blog.golang.org/context">https://blog.golang.org/context</a> for example code for a server that uses
Contexts.

context 定义了 Context 类型, 它可在 API 边界和进程之间传递截止时间, 取消信号, 以及其他请求的生命周期中的值.

请求到达server后需创建一个 Context 实例. 服务器的响应也应该接受 Context。在它们之间的函数调用链必须传递该Context, 也可以用 WithCancel, WithDeadline, WithTimeout 或 WithValue 衍生出的新 Context 实例代替它. 当取消一个 Context 实例时, 所有由它衍生出的 Context 实例都会被取消.

函数 WithCancel，WithDeadline，WithTimeout 接受一个父级 Context 实例, 并返回一个衍生的子代 Context 和一个 CancelFunc. 调用 CancelFunc 会取消该子代Context及其后代的Context, 删除父级与它的联系并停止所有关联的定时器. 不调用 CancelFunc 会使子代及其子代泄漏，直到父代被取消或计时器触发. `go vet` 工具可以检查是否所有的流程控制都使用了 CancelFunc.

为了各个包之间的接口保持一致并且能够使用静态工具检查context的传递，使用 Context 时需要遵循以下规则：

不要将 Context 储存在结构体中, 而是在函数中按需传递 Context. Context 应该作为函数的第一个参数, 一般叫 ctx：

```go
func DoSomething(ctx context.Context, arg Arg) error {
    // ... use ctx ...
}
```

即使函数允许也不要传递值为 nil 的 Context. 如果你不确定使用哪个 Context 则可以使用 context.TODO.

仅将context的Value用于传递处理过程和API的生命周期中的数据, 不将其作为函数的可选参数.

相同的 Context 可以传递给不同的 goroutines. Context 可以安全地被多个 goroutine 同时使用.

server端示例见 https://blog.golang.org/context.



## <a id="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func WithCancel(parent Context) (ctx Context, cancel CancelFunc)](#WithCancel)
* [func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)](#WithDeadline)
* [func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)](#WithTimeout)
* [type CancelFunc](#CancelFunc)
* [type Context](#Context)
  * [func Background() Context](#Background)
  * [func TODO() Context](#TODO)
  * [func WithValue(parent Context, key, val interface{}) Context](#WithValue)


#### <a id="pkg-examples">Examples</a>
* [WithCancel](#example_WithCancel)
* [WithDeadline](#example_WithDeadline)
* [WithTimeout](#example_WithTimeout)
* [WithValue](#example_WithValue)


#### <a id="pkg-files">Package files</a>
[context.go](https://golang.org/src/context/context.go) 




## <a id="pkg-variables">Variables</a>
Canceled is the error returned by Context.Err when the context is canceled.

Canceled 是 context 取消后, `Context.Err()`返回的error.

<pre>var <span id="Canceled">Canceled</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;context canceled&#34;)</pre>DeadlineExceeded is the error returned by Context.Err when the context's
deadline passes.

DeadlineExceeded 是 context 超过截止时间后, `Context.Err()`返回的error.

<pre>var <span id="DeadlineExceeded">DeadlineExceeded</span> <a href="/pkg/builtin/#error">error</a> = deadlineExceededError{}</pre>

## <a id="WithCancel">func</a> [WithCancel](https://golang.org/src/context/context.go?s=8304:8368#L219)
<pre>func WithCancel(parent <a href="#Context">Context</a>) (ctx <a href="#Context">Context</a>, cancel <a href="#CancelFunc">CancelFunc</a>)</pre>
WithCancel returns a copy of parent with a new Done channel. The returned
context's Done channel is closed when the returned cancel function is called
or when the parent context's Done channel is closed, whichever happens first.

Canceling this context releases resources associated with it, so code should
call cancel as soon as the operations running in this Context complete.

WithCancel 返回拥有 Done channel 的 parent 副本. 当 cancel 函数被调用后或者父context的Done channel被关闭后(以先发生的为准), 返回的子context的 Done channel会被close.

取消 context 会释放和它关联的资源，所以我们应该在 Context 完成后立即调用cancel.

<a id="example_WithCancel">Example</a>
<p>This example demonstrates the use of a cancelable context to prevent a
goroutine leak. By the end of the example function, the goroutine started
by gen will return without leaking.
</p>

此示例演示了在使用可取消的context时如何防止goroutine泄漏. 在示例函数结束时，由gen启动的goroutine将返回而不会泄漏.

```go
package main

import (
	"context"
	"fmt"
)

func main() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
    // the internal goroutine started by gen.
    //
    // gen在单独的goroutine中生成整数, 然后将它们发送到channel上.
    // gen的调用者需要在消费完生成的整数后需要立即取消context, 避免gen生成的内部goroutine泄露
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
```

output
```txt
1
2
3
4
5
```

## <a id="WithDeadline">func</a> [WithDeadline](https://golang.org/src/context/context.go?s=12430:12498#L384)
<pre>func WithDeadline(parent <a href="#Context">Context</a>, d <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>) (<a href="#Context">Context</a>, <a href="#CancelFunc">CancelFunc</a>)</pre>
WithDeadline returns a copy of the parent context with the deadline adjusted
to be no later than d. If the parent's deadline is already earlier than d,
WithDeadline(parent, d) is semantically equivalent to parent. The returned
context's Done channel is closed when the deadline expires, when the returned
cancel function is called, or when the parent context's Done channel is
closed, whichever happens first.

Canceling this context releases resources associated with it, so code should
call cancel as soon as the operations running in this Context complete.


WithDeadline 返回拥有截止时间的 parent 副本, 并将截止时间调整为不晚于d. 如果 parent 的截止时间早于 d, 那么`WithDeadline(parent, d)`在语义上等同于使用 parent 的截止时间. 当截止时间到来, 或返回的cancel函数被调用, 或父context的Done channel被关闭(以先发生为准), 返回的context的Done channel会被关闭.

取消 context 会释放和他关联的资源，所以我们应该在 Context 完成后立即立即调用cancel.

<a id="example_WithDeadline">Example</a>
<p>This example passes a context with an arbitrary deadline to tell a blocking
function that it should abandon its work as soon as it gets to it.
</p>

这个例子传递了一个带有截止时间的context来告诉阻塞函数: 时间到后应该立即放弃继续执行.

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
    //
    // 即使 ctx 将要过期，在最后调用取消函数也是很好的做法. 不这么做会导致 context 和它的父级在不必要的时候继续存在.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
```

output
```txt
context deadline exceeded
```

## <a id="WithTimeout">func</a> [WithTimeout](https://golang.org/src/context/context.go?s=14448:14525#L453)
<pre>func WithTimeout(parent <a href="#Context">Context</a>, timeout <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>) (<a href="#Context">Context</a>, <a href="#CancelFunc">CancelFunc</a>)</pre>
WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).

Canceling this context releases resources associated with it, so code should
call cancel as soon as the operations running in this Context complete:

WithTimeout 返回 WithDeadline(parent, time.Now().Add(timeout)).

取消 context 会释放和他关联的资源，所以我们应该在 Context 完成后立即立即调用cancel:


	func slowOperationWithTimeout(ctx context.Context) (Result, error) {
		ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
		defer cancel()  // releases resources if slowOperation completes before timeout elapses // 如果slowOperation在超时之前完成就立即释放资源
		return slowOperation(ctx)
	}


<a id="example_WithTimeout">Example</a>
<p>This example passes a context with a timeout to tell a blocking function that
it should abandon its work after the timeout elapses.
</p>

这个例子传递了一个带有超时时间的context来告诉阻塞函数: 超时后应该立即放弃继续执行.

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Pass a context with a timeout to tell a blocking function that it
    // should abandon its work after the timeout elapses.
    //
    // 通过一个传带超时的 context 可以通知阻塞函数应该在到达超时时间后放弃继续执行.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}

}
```

output
```txt
context deadline exceeded
```


## <a id="CancelFunc">type</a> [CancelFunc](https://golang.org/src/context/context.go?s=7884:7906#L211)
A CancelFunc tells an operation to abandon its work.
A CancelFunc does not wait for the work to stop.
A CancelFunc may be called by multiple goroutines simultaneously.
After the first call, subsequent calls to a CancelFunc do nothing.

CancelFunc 会通知一个操作终止其工作. CancelFunc 不会等待工作完成后再终止. 多个goroutine可以同时调用CancelFunc. 在首次调用 CancelFunc 后，再调用 CancelFunc 不做任何事.

<pre>type CancelFunc func()</pre>











## <a id="Context">type</a> [Context](https://golang.org/src/context/context.go?s=2445:5904#L51)
A Context carries a deadline, a cancellation signal, and other values across
API boundaries.

Context's methods may be called by multiple goroutines simultaneously.

Context可以携带截止时间, 取消信号和跨越 API 边界的其他值.

Context的方法可以同时被多个goroutine调用.

```go
type Context interface {
    // Deadline returns the time when work done on behalf of this context
    // should be canceled. Deadline returns ok==false when no deadline is
    // set. Successive calls to Deadline return the same results.
    //
    // Deadline 返回work的截止时间, 表示到时该context需要被取消. ok表示是否设置过deadline, ok==false 表示未设置. 连续调用 Deadline 返回同样的结果.
    Deadline() (deadline time.Time, ok bool)

    // Done returns a channel that's closed when work done on behalf of this
    // context should be canceled. Done may return nil if this context can
    // never be canceled. Successive calls to Done return the same value.
    //
    // WithCancel arranges for Done to be closed when cancel is called;
    // WithDeadline arranges for Done to be closed when the deadline
    // expires; WithTimeout arranges for Done to be closed when the timeout
    // elapses.
    //
    // Done is provided for use in select statements:
    //
    //  // Stream generates values with DoSomething and sends them to out
    //  // until DoSomething returns an error or ctx.Done is closed.
    //  func Stream(ctx context.Context, out chan<- Value) error {
    //  	for {
    //  		v, err := DoSomething(ctx)
    //  		if err != nil {
    //  			return err
    //  		}
    //  		select {
    //  		case <-ctx.Done():
    //  			return ctx.Err()
    //  		case out <- v:
    //  		}
    //  	}
    //  }
    //
    // See https://blog.golang.org/pipelines for more examples of how to use
    // a Done channel for cancellation.
    //
    // Done 返回 一个channel, 当它关闭了的话表示work已完成, 此时需要取消该context.
    // Done可能返回nil, 此时该context永远不会被取消. 连续调用 Done 返回相同的值.
    //
    // WithCancel 会在 cancel 调用后关闭 Done; WithDeadline 会在 deadline到期后关闭 Done; WithTimeout会在超时后关闭 Done.
    //
    // Done 和 select 搭配使用:
    //
    //  // Stream generates values with DoSomething and sends them to out
    //  // until DoSomething returns an error or ctx.Done is closed.
    //  func Stream(ctx context.Context, out chan<- Value) error {
    //      for {
    //          v, err := DoSomething(ctx)
    //          if err != nil {
    //              return err
    //          }
    //          select {
    //          case <-ctx.Done():
    //              return ctx.Err()
    //          case out <- v:
    //          }
    //      }
    //  }
    //
    // 参考 https://blog.golang.org/pipelines, 上面有更多的示例解释了如何使用Done channel来取消context.
    Done() <-chan struct{}

    // If Done is not yet closed, Err returns nil.
    // If Done is closed, Err returns a non-nil error explaining why:
    // Canceled if the context was canceled
    // or DeadlineExceeded if the context's deadline passed.
    // After Err returns a non-nil error, successive calls to Err return the same error.
    //
    // 如果 Done 还没有被关闭，Err 返回 nil.
    // 如果 Done 已经关闭，Err 回返回非 nil 值来解释原因:
    // Canceled : context 是被取消的
    // DeadlineExceeded: context 到达截止时间
    // 在 Err 返回非 nil 值后，连续调用 Err 都返回相同结果
    Err() error

    // Value returns the value associated with this context for key, or nil
    // if no value is associated with key. Successive calls to Value with
    // the same key returns the same result.
    //
    // Use context values only for request-scoped data that transits
    // processes and API boundaries, not for passing optional parameters to
    // functions.
    //
    // A key identifies a specific value in a Context. Functions that wish
    // to store values in Context typically allocate a key in a global
    // variable then use that key as the argument to context.WithValue and
    // Context.Value. A key can be any type that supports equality;
    // packages should define keys as an unexported type to avoid
    // collisions.
    //
    // Packages that define a Context key should provide type-safe accessors
    // for the values stored using that key:
    //
    // 	// Package user defines a User type that's stored in Contexts.
    // 	package user
    //
    // 	import "context"
    //
    // 	// User is the type of value stored in the Contexts.
    // 	type User struct {...}
    //
    // 	// key is an unexported type for keys defined in this package.
    // 	// This prevents collisions with keys defined in other packages.
    // 	type key int
    //
    // 	// userKey is the key for user.User values in Contexts. It is
    // 	// unexported; clients use user.NewContext and user.FromContext
    // 	// instead of using this key directly.
    // 	var userKey key
    //
    // 	// NewContext returns a new Context that carries value u.
    // 	func NewContext(ctx context.Context, u *User) context.Context {
    // 		return context.WithValue(ctx, userKey, u)
    // 	}
    //
    // 	// FromContext returns the User value stored in ctx, if any.
    // 	func FromContext(ctx context.Context) (*User, bool) {
    // 		u, ok := ctx.Value(userKey).(*User)
    // 		return u, ok
    // 	}
    //
    // Value 返回context中与该key关联的value, 如果该key没有关联值则会返回nil. 使用相同的key连续调用 Value 返回相同的值.
    //
    // 仅将context的Value用于传递处理过程和API的生命周期中的数据, 不将其作为函数的可选参数.
    //
    // Context 中使用key来标识特定的值.
    // 函数通常会希望分配一个全局变量来标识存储在 Context 中的值, 然后将该key作为参数传递给 context.WithValue 和 Context.Value.
    // key 可以是任意类型但必须支持相等比较, 通常需要将key定义成非导出类型已避免碰撞.
    //
    // 定义了一个 Context 中使用的key, 为相应的值提供了类型安全的访问器.
    // Packages that define a Context key should provide type-safe accessors
    // for the values stored using that key:
    //
    //  // Package user defines a User type that's stored in Contexts.
    //  // user 包定义了一个保存在 Context 中的 User 类型.
    //  package user
    //
    //  import "context"
    //
    //  // User is the type of value stored in the Contexts.
    //  // User 是保存在 Context 中的类型.
    //  type User struct {...}
    //
    //  // key is an unexported type for keys defined in this package.
    //  // This prevents collisions with keys defined in other packages.
    //  //
    //  // key 是未导出的类型, 用于表示该package中的keys.
    //  // 这样可以避免与其他包中定义的 key 类型发生冲突.
    //  type key int
    //
    //  // userKey 是 保存在Context中的user.User值的key.
    //  // 它是不可导出的; 用户可使用 user.NewContext 和 user.FromContext 以避免直接使用该key.
    //  var userKey key
    //
    //  // NewContext returns a new Context that carries value u.
    //  // NewContext 返回一个新的Context并携带了u.
    //  func NewContext(ctx context.Context, u *User) context.Context {
    //      return context.WithValue(ctx, userKey, u)
    //  }
    //
    //  // FromContext returns the User value stored in ctx, if any.
    //  // 如果存在, FromContext 返回 ctx 中 存储的 User 值.
    //  func FromContext(ctx context.Context) (*User, bool) {
    //      u, ok := ctx.Value(userKey).(*User)
    //      return u, ok
    //  }
    Value(key interface{}) interface{}
}
```


### <a id="Background">func</a> [Background](https://golang.org/src/context/context.go?s=7311:7336#L195)
<pre>func Background() <a href="#Context">Context</a></pre>
Background returns a non-nil, empty Context. It is never canceled, has no
values, and has no deadline. It is typically used by the main function,
initialization, and tests, and as the top-level Context for incoming
requests.

Background 返回一个非 nil 的空 Context. 它不能取消，不包含值，没有截止时间. 它的典型应该用场景包括：main 函数，初始化，测试，或者作为所有请求的顶级 Context.




### <a id="TODO">func</a> [TODO](https://golang.org/src/context/context.go?s=7599:7618#L203)
<pre>func TODO() <a href="#Context">Context</a></pre>
TODO returns a non-nil, empty Context. Code should use context.TODO when
it's unclear which Context to use or it is not yet available (because the
surrounding function has not yet been extended to accept a Context
parameter).

TODO 返回一个非 nil 的空 Context. 当不知道该使用哪种 Context 或者它尚未可用(因为周围的函数尚未扩展为可接受Context参数)时, 此时应该使用 context.TODO.



### <a id="WithValue">func</a> [WithValue](https://golang.org/src/context/context.go?s=15240:15300#L470)
<pre>func WithValue(parent <a href="#Context">Context</a>, key, val interface{}) <a href="#Context">Context</a></pre>
WithValue returns a copy of parent in which the value associated with key is
val.

Use context Values only for request-scoped data that transits processes and
APIs, not for passing optional parameters to functions.

The provided key must be comparable and should not be of type
string or any other built-in type to avoid collisions between
packages using context. Users of WithValue should define their own
types for keys. To avoid allocating when assigning to an
interface{}, context keys often have concrete type
struct{}. Alternatively, exported context key variables' static
type should be a pointer or interface.

WithValue 将 key = value 的键值对保存在 parent 的副本中并返回.

仅将context的Value用于传递处理过程和API的生命周期中的数据, 不将其作为函数的可选参数.

提供的key必须具有可比性, 并且不能使用字符串类型或其他任何的built-in类型以避免不同的package间使用context发生冲突. WithValue 的用户应该定义自己的 key 类型. 为了避免在分配时赋值给了`interface{}`，context的key通常是具有具体类型的`struct{}`. 另外，导出的contex 的key的静态类型应该是指针或接口.


<a id="example_WithValue">Example</a>
<p>This example demonstrates how a value can be passed to the context
and also how to retrieve it if it exists.
</p>

此示例演示如何将值传递给context以及如何将其取回(如果存在).

```go
package main

import (
	"context"
	"fmt"
)

func main() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))

}
```

output
```txt
found value: Go
key not found: color
```







