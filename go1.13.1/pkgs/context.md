

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


<pre>var <span id="Canceled">Canceled</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;context canceled&#34;)</pre>DeadlineExceeded is the error returned by Context.Err when the context's
deadline passes.


<pre>var <span id="DeadlineExceeded">DeadlineExceeded</span> <a href="/pkg/builtin/#error">error</a> = deadlineExceededError{}</pre>

## <a id="WithCancel">func</a> [WithCancel](https://golang.org/src/context/context.go?s=8304:8368#L219)
<pre>func WithCancel(parent <a href="#Context">Context</a>) (ctx <a href="#Context">Context</a>, cancel <a href="#CancelFunc">CancelFunc</a>)</pre>
WithCancel returns a copy of parent with a new Done channel. The returned
context's Done channel is closed when the returned cancel function is called
or when the parent context's Done channel is closed, whichever happens first.

Canceling this context releases resources associated with it, so code should
call cancel as soon as the operations running in this Context complete.


<a id="example_WithCancel">Example</a>
<p>This example demonstrates the use of a cancelable context to prevent a
goroutine leak. By the end of the example function, the goroutine started
by gen will return without leaking.
</p>
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


<a id="example_WithDeadline">Example</a>
<p>This example passes a context with an arbitrary deadline to tell a blocking
function that it should abandon its work as soon as it gets to it.
</p>
## <a id="WithTimeout">func</a> [WithTimeout](https://golang.org/src/context/context.go?s=14448:14525#L453)
<pre>func WithTimeout(parent <a href="#Context">Context</a>, timeout <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>) (<a href="#Context">Context</a>, <a href="#CancelFunc">CancelFunc</a>)</pre>
WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).

Canceling this context releases resources associated with it, so code should
call cancel as soon as the operations running in this Context complete:


	func slowOperationWithTimeout(ctx context.Context) (Result, error) {
		ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
		defer cancel()  // releases resources if slowOperation completes before timeout elapses
		return slowOperation(ctx)
	}


<a id="example_WithTimeout">Example</a>
<p>This example passes a context with a timeout to tell a blocking function that
it should abandon its work after the timeout elapses.
</p>


## <a id="CancelFunc">type</a> [CancelFunc](https://golang.org/src/context/context.go?s=7884:7906#L211)
A CancelFunc tells an operation to abandon its work.
A CancelFunc does not wait for the work to stop.
A CancelFunc may be called by multiple goroutines simultaneously.
After the first call, subsequent calls to a CancelFunc do nothing.


<pre>type CancelFunc func()</pre>











## <a id="Context">type</a> [Context](https://golang.org/src/context/context.go?s=2445:5904#L51)
A Context carries a deadline, a cancellation signal, and other values across
API boundaries.

Context's methods may be called by multiple goroutines simultaneously.


<pre>type Context interface {
    <span class="comment">// Deadline returns the time when work done on behalf of this context</span>
    <span class="comment">// should be canceled. Deadline returns ok==false when no deadline is</span>
    <span class="comment">// set. Successive calls to Deadline return the same results.</span>
    Deadline() (deadline <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Time">Time</a>, ok <a href="/pkg/builtin/#bool">bool</a>)

    <span class="comment">// Done returns a channel that&#39;s closed when work done on behalf of this</span>
    <span class="comment">// context should be canceled. Done may return nil if this context can</span>
    <span class="comment">// never be canceled. Successive calls to Done return the same value.</span>
    <span class="comment">//</span>
    <span class="comment">// WithCancel arranges for Done to be closed when cancel is called;</span>
    <span class="comment">// WithDeadline arranges for Done to be closed when the deadline</span>
    <span class="comment">// expires; WithTimeout arranges for Done to be closed when the timeout</span>
    <span class="comment">// elapses.</span>
    <span class="comment">//</span>
    <span class="comment">// Done is provided for use in select statements:</span>
    <span class="comment">//</span>
    <span class="comment">//  // Stream generates values with DoSomething and sends them to out</span>
    <span class="comment">//  // until DoSomething returns an error or ctx.Done is closed.</span>
    <span class="comment">//  func Stream(ctx context.Context, out chan&lt;- Value) error {</span>
    <span class="comment">//  	for {</span>
    <span class="comment">//  		v, err := DoSomething(ctx)</span>
    <span class="comment">//  		if err != nil {</span>
    <span class="comment">//  			return err</span>
    <span class="comment">//  		}</span>
    <span class="comment">//  		select {</span>
    <span class="comment">//  		case &lt;-ctx.Done():</span>
    <span class="comment">//  			return ctx.Err()</span>
    <span class="comment">//  		case out &lt;- v:</span>
    <span class="comment">//  		}</span>
    <span class="comment">//  	}</span>
    <span class="comment">//  }</span>
    <span class="comment">//</span>
    <span class="comment">// See https://blog.golang.org/pipelines for more examples of how to use</span>
    <span class="comment">// a Done channel for cancellation.</span>
    Done() &lt;-chan struct{}

    <span class="comment">// If Done is not yet closed, Err returns nil.</span>
    <span class="comment">// If Done is closed, Err returns a non-nil error explaining why:</span>
    <span class="comment">// Canceled if the context was canceled</span>
    <span class="comment">// or DeadlineExceeded if the context&#39;s deadline passed.</span>
    <span class="comment">// After Err returns a non-nil error, successive calls to Err return the same error.</span>
    Err() <a href="/pkg/builtin/#error">error</a>

    <span class="comment">// Value returns the value associated with this context for key, or nil</span>
    <span class="comment">// if no value is associated with key. Successive calls to Value with</span>
    <span class="comment">// the same key returns the same result.</span>
    <span class="comment">//</span>
    <span class="comment">// Use context values only for request-scoped data that transits</span>
    <span class="comment">// processes and API boundaries, not for passing optional parameters to</span>
    <span class="comment">// functions.</span>
    <span class="comment">//</span>
    <span class="comment">// A key identifies a specific value in a Context. Functions that wish</span>
    <span class="comment">// to store values in Context typically allocate a key in a global</span>
    <span class="comment">// variable then use that key as the argument to context.WithValue and</span>
    <span class="comment">// Context.Value. A key can be any type that supports equality;</span>
    <span class="comment">// packages should define keys as an unexported type to avoid</span>
    <span class="comment">// collisions.</span>
    <span class="comment">//</span>
    <span class="comment">// Packages that define a Context key should provide type-safe accessors</span>
    <span class="comment">// for the values stored using that key:</span>
    <span class="comment">//</span>
    <span class="comment">// 	// Package user defines a User type that&#39;s stored in Contexts.</span>
    <span class="comment">// 	package user</span>
    <span class="comment">//</span>
    <span class="comment">// 	import &#34;context&#34;</span>
    <span class="comment">//</span>
    <span class="comment">// 	// User is the type of value stored in the Contexts.</span>
    <span class="comment">// 	type User struct {...}</span>
    <span class="comment">//</span>
    <span class="comment">// 	// key is an unexported type for keys defined in this package.</span>
    <span class="comment">// 	// This prevents collisions with keys defined in other packages.</span>
    <span class="comment">// 	type key int</span>
    <span class="comment">//</span>
    <span class="comment">// 	// userKey is the key for user.User values in Contexts. It is</span>
    <span class="comment">// 	// unexported; clients use user.NewContext and user.FromContext</span>
    <span class="comment">// 	// instead of using this key directly.</span>
    <span class="comment">// 	var userKey key</span>
    <span class="comment">//</span>
    <span class="comment">// 	// NewContext returns a new Context that carries value u.</span>
    <span class="comment">// 	func NewContext(ctx context.Context, u *User) context.Context {</span>
    <span class="comment">// 		return context.WithValue(ctx, userKey, u)</span>
    <span class="comment">// 	}</span>
    <span class="comment">//</span>
    <span class="comment">// 	// FromContext returns the User value stored in ctx, if any.</span>
    <span class="comment">// 	func FromContext(ctx context.Context) (*User, bool) {</span>
    <span class="comment">// 		u, ok := ctx.Value(userKey).(*User)</span>
    <span class="comment">// 		return u, ok</span>
    <span class="comment">// 	}</span>
    Value(key interface{}) interface{}
}</pre>









### <a id="Background">func</a> [Background](https://golang.org/src/context/context.go?s=7311:7336#L195)
<pre>func Background() <a href="#Context">Context</a></pre>
Background returns a non-nil, empty Context. It is never canceled, has no
values, and has no deadline. It is typically used by the main function,
initialization, and tests, and as the top-level Context for incoming
requests.




### <a id="TODO">func</a> [TODO](https://golang.org/src/context/context.go?s=7599:7618#L203)
<pre>func TODO() <a href="#Context">Context</a></pre>
TODO returns a non-nil, empty Context. Code should use context.TODO when
it's unclear which Context to use or it is not yet available (because the
surrounding function has not yet been extended to accept a Context
parameter).




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


<a id="example_WithValue">Example</a>
<p>This example demonstrates how a value can be passed to the context
and also how to retrieve it if it exists.
</p>







