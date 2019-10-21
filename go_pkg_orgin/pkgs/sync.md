

# sync
`import "sync"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package sync provides basic synchronization primitives such as mutual
exclusion locks. Other than the Once and WaitGroup types, most are intended
for use by low-level library routines. Higher-level synchronization is
better done via channels and communication.

Values containing the types defined in this package should not be copied.




## <a id="pkg-index">Index</a>
* [type Cond](#Cond)
  * [func NewCond(l Locker) *Cond](#NewCond)
  * [func (c *Cond) Broadcast()](#Cond.Broadcast)
  * [func (c *Cond) Signal()](#Cond.Signal)
  * [func (c *Cond) Wait()](#Cond.Wait)
* [type Locker](#Locker)
* [type Map](#Map)
  * [func (m *Map) Delete(key interface{})](#Map.Delete)
  * [func (m *Map) Load(key interface{}) (value interface{}, ok bool)](#Map.Load)
  * [func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)](#Map.LoadOrStore)
  * [func (m *Map) Range(f func(key, value interface{}) bool)](#Map.Range)
  * [func (m *Map) Store(key, value interface{})](#Map.Store)
* [type Mutex](#Mutex)
  * [func (m *Mutex) Lock()](#Mutex.Lock)
  * [func (m *Mutex) Unlock()](#Mutex.Unlock)
* [type Once](#Once)
  * [func (o *Once) Do(f func())](#Once.Do)
* [type Pool](#Pool)
  * [func (p *Pool) Get() interface{}](#Pool.Get)
  * [func (p *Pool) Put(x interface{})](#Pool.Put)
* [type RWMutex](#RWMutex)
  * [func (rw *RWMutex) Lock()](#RWMutex.Lock)
  * [func (rw *RWMutex) RLock()](#RWMutex.RLock)
  * [func (rw *RWMutex) RLocker() Locker](#RWMutex.RLocker)
  * [func (rw *RWMutex) RUnlock()](#RWMutex.RUnlock)
  * [func (rw *RWMutex) Unlock()](#RWMutex.Unlock)
* [type WaitGroup](#WaitGroup)
  * [func (wg *WaitGroup) Add(delta int)](#WaitGroup.Add)
  * [func (wg *WaitGroup) Done()](#WaitGroup.Done)
  * [func (wg *WaitGroup) Wait()](#WaitGroup.Wait)


#### <a id="pkg-examples">Examples</a>
* [Once](#example_Once)
* [Pool](#example_Pool)
* [WaitGroup](#example_WaitGroup)


#### <a id="pkg-files">Package files</a>
[cond.go](https://golang.org/src/sync/cond.go) [map.go](https://golang.org/src/sync/map.go) [mutex.go](https://golang.org/src/sync/mutex.go) [once.go](https://golang.org/src/sync/once.go) [pool.go](https://golang.org/src/sync/pool.go) [poolqueue.go](https://golang.org/src/sync/poolqueue.go) [runtime.go](https://golang.org/src/sync/runtime.go) [rwmutex.go](https://golang.org/src/sync/rwmutex.go) [waitgroup.go](https://golang.org/src/sync/waitgroup.go) 








## <a id="Cond">type</a> [Cond](https://golang.org/src/sync/cond.go?s=555:699#L11)
Cond implements a condition variable, a rendezvous point
for goroutines waiting for or announcing the occurrence
of an event.

Each Cond has an associated Locker L (often a *Mutex or *RWMutex),
which must be held when changing the condition and
when calling the Wait method.

A Cond must not be copied after first use.


<pre>type Cond struct {

<span id="Cond.L"></span>    <span class="comment">// L is held while observing or changing the condition</span>
    L <a href="#Locker">Locker</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewCond">func</a> [NewCond](https://golang.org/src/sync/cond.go?s=746:774#L22)
<pre>func NewCond(l <a href="#Locker">Locker</a>) *<a href="#Cond">Cond</a></pre>
NewCond returns a new Cond with Locker l.






### <a id="Cond.Broadcast">func</a> (\*Cond) [Broadcast](https://golang.org/src/sync/cond.go?s=1867:1893#L63)
<pre>func (c *<a href="#Cond">Cond</a>) Broadcast()</pre>
Broadcast wakes all goroutines waiting on c.

It is allowed but not required for the caller to hold c.L
during the call.




### <a id="Cond.Signal">func</a> (\*Cond) [Signal](https://golang.org/src/sync/cond.go?s=1647:1670#L54)
<pre>func (c *<a href="#Cond">Cond</a>) Signal()</pre>
Signal wakes one goroutine waiting on c, if there is any.

It is allowed but not required for the caller to hold c.L
during the call.




### <a id="Cond.Wait">func</a> (\*Cond) [Wait](https://golang.org/src/sync/cond.go?s=1353:1374#L42)
<pre>func (c *<a href="#Cond">Cond</a>) Wait()</pre>
Wait atomically unlocks c.L and suspends execution
of the calling goroutine. After later resuming execution,
Wait locks c.L before returning. Unlike in other systems,
Wait cannot return unless awoken by Broadcast or Signal.

Because c.L is not locked when Wait first resumes, the caller
typically cannot assume that the condition is true when
Wait returns. Instead, the caller should Wait in a loop:


	c.L.Lock()
	for !condition() {
	    c.Wait()
	}
	... make use of condition ...
	c.L.Unlock()




## <a id="Locker">type</a> [Locker](https://golang.org/src/sync/mutex.go?s=881:924#L21)
A Locker represents an object that can be locked and unlocked.


<pre>type Locker interface {
    Lock()
    Unlock()
}</pre>











## <a id="Map">type</a> [Map](https://golang.org/src/sync/map.go?s=1149:2596#L17)
Map is like a Go map[interface{}]interface{} but is safe for concurrent use
by multiple goroutines without additional locking or coordination.
Loads, stores, and deletes run in amortized constant time.

The Map type is specialized. Most code should use a plain Go map instead,
with separate locking or coordination, for better type safety and to make it
easier to maintain other invariants along with the map content.

The Map type is optimized for two common use cases: (1) when the entry for a given
key is only ever written once but read many times, as in caches that only grow,
or (2) when multiple goroutines read, write, and overwrite entries for disjoint
sets of keys. In these two cases, use of a Map may significantly reduce lock
contention compared to a Go map paired with a separate Mutex or RWMutex.

The zero Map is empty and ready for use. A Map must not be copied after first use.


<pre>type Map struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="Map.Delete">func</a> (\*Map) [Delete](https://golang.org/src/sync/map.go?s=8705:8742#L257)
<pre>func (m *<a href="#Map">Map</a>) Delete(key interface{})</pre>
Delete deletes the value for a key.




### <a id="Map.Load">func</a> (\*Map) [Load](https://golang.org/src/sync/map.go?s=4102:4166#L92)
<pre>func (m *<a href="#Map">Map</a>) Load(key interface{}) (value interface{}, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
Load returns the value stored in the map for a key, or nil if no
value is present.
The ok result indicates whether value was found in the map.




### <a id="Map.LoadOrStore">func</a> (\*Map) [LoadOrStore](https://golang.org/src/sync/map.go?s=6882:6965#L189)
<pre>func (m *<a href="#Map">Map</a>) LoadOrStore(key, value interface{}) (actual interface{}, loaded <a href="/pkg/builtin/#bool">bool</a>)</pre>
LoadOrStore returns the existing value for the key if present.
Otherwise, it stores and returns the given value.
The loaded result is true if the value was loaded, false if stored.




### <a id="Map.Range">func</a> (\*Map) [Range](https://golang.org/src/sync/map.go?s=9749:9805#L296)
<pre>func (m *<a href="#Map">Map</a>) Range(f func(key, value interface{}) <a href="/pkg/builtin/#bool">bool</a>)</pre>
Range calls f sequentially for each key and value present in the map.
If f returns false, range stops the iteration.

Range does not necessarily correspond to any consistent snapshot of the Map's
contents: no key will be visited more than once, but if the value for any key
is stored or deleted concurrently, Range may reflect any mapping for that key
from any point during the Range call.

Range may be O(N) with the number of elements in the map even if f returns
false after a constant number of calls.




### <a id="Map.Store">func</a> (\*Map) [Store](https://golang.org/src/sync/map.go?s=5046:5089#L126)
<pre>func (m *<a href="#Map">Map</a>) Store(key, value interface{})</pre>
Store sets the value for a key.




## <a id="Mutex">type</a> [Mutex](https://golang.org/src/sync/mutex.go?s=765:813#L15)
A Mutex is a mutual exclusion lock.
The zero value for a Mutex is an unlocked mutex.

A Mutex must not be copied after first use.


<pre>type Mutex struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="Mutex.Lock">func</a> (\*Mutex) [Lock](https://golang.org/src/sync/mutex.go?s=2534:2556#L62)
<pre>func (m *<a href="#Mutex">Mutex</a>) Lock()</pre>
Lock locks m.
If the lock is already in use, the calling goroutine
blocks until the mutex is available.




### <a id="Mutex.Unlock">func</a> (\*Mutex) [Unlock](https://golang.org/src/sync/mutex.go?s=5882:5906#L169)
<pre>func (m *<a href="#Mutex">Mutex</a>) Unlock()</pre>
Unlock unlocks m.
It is a run-time error if m is not locked on entry to Unlock.

A locked Mutex is not associated with a particular goroutine.
It is allowed for one goroutine to lock a Mutex and then
arrange for another goroutine to unlock it.




## <a id="Once">type</a> [Once](https://golang.org/src/sync/once.go?s=260:641#L2)
Once is an object that will perform exactly one action.


<pre>type Once struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Once">Example</a>
```go
```

output:
```txt
```






### <a id="Once.Do">func</a> (\*Once) [Do](https://golang.org/src/sync/once.go?s=1473:1500#L30)
<pre>func (o *<a href="#Once">Once</a>) Do(f func())</pre>
Do calls the function f if and only if Do is being called for the
first time for this instance of Once. In other words, given


	var once Once

if once.Do(f) is called multiple times, only the first call will invoke f,
even if f has a different value in each invocation. A new instance of
Once is required for each function to execute.

Do is intended for initialization that must be run exactly once. Since f
is niladic, it may be necessary to use a function literal to capture the
arguments to a function to be invoked by Do:


	config.once.Do(func() { config.init(filename) })

Because no call to Do returns until the one call to f returns, if f causes
Do to be called, it will deadlock.

If f panics, Do considers it to have returned; future calls of Do return
without calling f.




## <a id="Pool">type</a> [Pool](https://golang.org/src/sync/pool.go?s=1633:2101#L34)
A Pool is a set of temporary objects that may be individually saved and
retrieved.

Any item stored in the Pool may be removed automatically at any time without
notification. If the Pool holds the only reference when this happens, the
item might be deallocated.

A Pool is safe for use by multiple goroutines simultaneously.

Pool's purpose is to cache allocated but unused items for later reuse,
relieving pressure on the garbage collector. That is, it makes it easy to
build efficient, thread-safe free lists. However, it is not suitable for all
free lists.

An appropriate use of a Pool is to manage a group of temporary items
silently shared among and potentially reused by concurrent independent
clients of a package. Pool provides a way to amortize allocation overhead
across many clients.

An example of good use of a Pool is in the fmt package, which maintains a
dynamically-sized store of temporary output buffers. The store scales under
load (when many goroutines are actively printing) and shrinks when
quiescent.

On the other hand, a free list maintained as part of a short-lived object is
not a suitable use for a Pool, since the overhead does not amortize well in
that scenario. It is more efficient to have such objects implement their own
free list.

A Pool must not be copied after first use.


<pre>type Pool struct {

<span id="Pool.New"></span>    <span class="comment">// New optionally specifies a function to generate</span>
    <span class="comment">// a value when Get would otherwise return nil.</span>
    <span class="comment">// It may not be changed concurrently with calls to Get.</span>
    New func() interface{}
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Pool">Example</a>
```go
```

output:
```txt
```






### <a id="Pool.Get">func</a> (\*Pool) [Get](https://golang.org/src/sync/pool.go?s=3925:3957#L114)
<pre>func (p *<a href="#Pool">Pool</a>) Get() interface{}</pre>
Get selects an arbitrary item from the Pool, removes it from the
Pool, and returns it to the caller.
Get may choose to ignore the pool and treat it as empty.
Callers should not assume any relation between values passed to Put and
the values returned by Get.

If Get would otherwise return nil and p.New is non-nil, Get returns
the result of calling p.New.




### <a id="Pool.Put">func</a> (\*Pool) [Put](https://golang.org/src/sync/pool.go?s=3164:3197#L80)
<pre>func (p *<a href="#Pool">Pool</a>) Put(x interface{})</pre>
Put adds x to the pool.




## <a id="RWMutex">type</a> [RWMutex](https://golang.org/src/sync/rwmutex.go?s=987:1319#L18)
A RWMutex is a reader/writer mutual exclusion lock.
The lock can be held by an arbitrary number of readers or a single writer.
The zero value for a RWMutex is an unlocked mutex.

A RWMutex must not be copied after first use.

If a goroutine holds a RWMutex for reading and another goroutine might
call Lock, no goroutine should expect to be able to acquire a read lock
until the initial read lock is released. In particular, this prohibits
recursive read locking. This is to ensure that the lock eventually becomes
available; a blocked Lock call excludes new readers from acquiring the
lock.


<pre>type RWMutex struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="RWMutex.Lock">func</a> (\*RWMutex) [Lock](https://golang.org/src/sync/rwmutex.go?s=2805:2830#L82)
<pre>func (rw *<a href="#RWMutex">RWMutex</a>) Lock()</pre>
Lock locks rw for writing.
If the lock is already locked for reading or writing,
Lock blocks until the lock is available.




### <a id="RWMutex.RLock">func</a> (\*RWMutex) [RLock](https://golang.org/src/sync/rwmutex.go?s=1558:1584#L33)
<pre>func (rw *<a href="#RWMutex">RWMutex</a>) RLock()</pre>
RLock locks rw for reading.

It should not be used for recursive read locking; a blocked Lock
call excludes new readers from acquiring the lock. See the
documentation on the RWMutex type.




### <a id="RWMutex.RLocker">func</a> (\*RWMutex) [RLocker](https://golang.org/src/sync/rwmutex.go?s=4328:4363#L134)
<pre>func (rw *<a href="#RWMutex">RWMutex</a>) RLocker() <a href="#Locker">Locker</a></pre>
RLocker returns a Locker interface that implements
the Lock and Unlock methods by calling rw.RLock and rw.RUnlock.




### <a id="RWMutex.RUnlock">func</a> (\*RWMutex) [RUnlock](https://golang.org/src/sync/rwmutex.go?s=2040:2068#L52)
<pre>func (rw *<a href="#RWMutex">RWMutex</a>) RUnlock()</pre>
RUnlock undoes a single RLock call;
it does not affect other simultaneous readers.
It is a run-time error if rw is not locked for reading
on entry to RUnlock.




### <a id="RWMutex.Unlock">func</a> (\*RWMutex) [Unlock](https://golang.org/src/sync/rwmutex.go?s=3664:3691#L108)
<pre>func (rw *<a href="#RWMutex">RWMutex</a>) Unlock()</pre>
Unlock unlocks rw for writing. It is a run-time error if rw is
not locked for writing on entry to Unlock.

As with Mutexes, a locked RWMutex is not associated with a particular
goroutine. One goroutine may RLock (Lock) a RWMutex and then
arrange for another goroutine to RUnlock (Unlock) it.




## <a id="WaitGroup">type</a> [WaitGroup](https://golang.org/src/sync/waitgroup.go?s=574:929#L10)
A WaitGroup waits for a collection of goroutines to finish.
The main goroutine calls Add to set the number of
goroutines to wait for. Then each of the goroutines
runs and calls Done when finished. At the same time,
Wait can be used to block until all goroutines have finished.

A WaitGroup must not be copied after first use.


<pre>type WaitGroup struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_WaitGroup">Example</a>
```go
```

output:
```txt
```
<p>This example fetches several URLs concurrently,
using a WaitGroup to block until all the fetches are complete.
</p>





### <a id="WaitGroup.Add">func</a> (\*WaitGroup) [Add](https://golang.org/src/sync/waitgroup.go?s=2022:2057#L43)
<pre>func (wg *<a href="#WaitGroup">WaitGroup</a>) Add(delta <a href="/pkg/builtin/#int">int</a>)</pre>
Add adds delta, which may be negative, to the WaitGroup counter.
If the counter becomes zero, all goroutines blocked on Wait are released.
If the counter goes negative, Add panics.

Note that calls with a positive delta that occur when the counter is zero
must happen before a Wait. Calls with a negative delta, or calls with a
positive delta that start when the counter is greater than zero, may happen
at any time.
Typically this means the calls to Add should execute before the statement
creating the goroutine or other event to be waited for.
If a WaitGroup is reused to wait for several independent sets of events,
new Add calls must happen after all previous Wait calls have returned.
See the WaitGroup example.




### <a id="WaitGroup.Done">func</a> (\*WaitGroup) [Done](https://golang.org/src/sync/waitgroup.go?s=3403:3430#L88)
<pre>func (wg *<a href="#WaitGroup">WaitGroup</a>) Done()</pre>
Done decrements the WaitGroup counter by one.




### <a id="WaitGroup.Wait">func</a> (\*WaitGroup) [Wait](https://golang.org/src/sync/waitgroup.go?s=3500:3527#L93)
<pre>func (wg *<a href="#WaitGroup">WaitGroup</a>) Wait()</pre>
Wait blocks until the WaitGroup counter is zero.







