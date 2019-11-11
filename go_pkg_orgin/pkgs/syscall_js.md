

# js
`import "syscall/js"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package js gives access to the WebAssembly host environment when using the js/wasm architecture.
Its API is based on JavaScript semantics.

This package is EXPERIMENTAL. Its current scope is only to allow tests to run, but not yet to provide a
comprehensive API for users. It is exempt from the Go compatibility promise.




## <a id="pkg-index">Index</a>
* [func CopyBytesToGo(dst []byte, src Value) int](#CopyBytesToGo)
* [func CopyBytesToJS(dst Value, src []byte) int](#CopyBytesToJS)
* [type Error](#Error)
  * [func (e Error) Error() string](#Error.Error)
* [type Func](#Func)
  * [func FuncOf(fn func(this Value, args []Value) interface{}) Func](#FuncOf)
  * [func (c Func) Release()](#Func.Release)
* [type Type](#Type)
  * [func (t Type) String() string](#Type.String)
* [type Value](#Value)
  * [func Global() Value](#Global)
  * [func Null() Value](#Null)
  * [func Undefined() Value](#Undefined)
  * [func ValueOf(x interface{}) Value](#ValueOf)
  * [func (v Value) Bool() bool](#Value.Bool)
  * [func (v Value) Call(m string, args ...interface{}) Value](#Value.Call)
  * [func (v Value) Float() float64](#Value.Float)
  * [func (v Value) Get(p string) Value](#Value.Get)
  * [func (v Value) Index(i int) Value](#Value.Index)
  * [func (v Value) InstanceOf(t Value) bool](#Value.InstanceOf)
  * [func (v Value) Int() int](#Value.Int)
  * [func (v Value) Invoke(args ...interface{}) Value](#Value.Invoke)
  * [func (v Value) JSValue() Value](#Value.JSValue)
  * [func (v Value) Length() int](#Value.Length)
  * [func (v Value) New(args ...interface{}) Value](#Value.New)
  * [func (v Value) Set(p string, x interface{})](#Value.Set)
  * [func (v Value) SetIndex(i int, x interface{})](#Value.SetIndex)
  * [func (v Value) String() string](#Value.String)
  * [func (v Value) Truthy() bool](#Value.Truthy)
  * [func (v Value) Type() Type](#Value.Type)
* [type ValueError](#ValueError)
  * [func (e *ValueError) Error() string](#ValueError.Error)
* [type Wrapper](#Wrapper)


#### <a id="pkg-examples">Examples</a>
* [FuncOf](#example_FuncOf)


#### <a id="pkg-files">Package files</a>
[func.go](https://golang.org/src/syscall/js/func.go) [js.go](https://golang.org/src/syscall/js/js.go) 






## <a id="CopyBytesToGo">func</a> [CopyBytesToGo](https://golang.org/src/syscall/js/js.go?s=12861:12906#L473)
<pre>func CopyBytesToGo(dst []<a href="/pkg/builtin/#byte">byte</a>, src <a href="#Value">Value</a>) <a href="/pkg/builtin/#int">int</a></pre>
CopyBytesToGo copies bytes from the Uint8Array src to dst.
It returns the number of bytes copied, which will be the minimum of the lengths of src and dst.
CopyBytesToGo panics if src is not an Uint8Array.



## <a id="CopyBytesToJS">func</a> [CopyBytesToJS](https://golang.org/src/syscall/js/js.go?s=13311:13356#L486)
<pre>func CopyBytesToJS(dst <a href="#Value">Value</a>, src []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#int">int</a></pre>
CopyBytesToJS copies bytes from src to the Uint8Array dst.
It returns the number of bytes copied, which will be the minimum of the lengths of src and dst.
CopyBytesToJS panics if dst is not an Uint8Array.





## <a id="Error">type</a> [Error](https://golang.org/src/syscall/js/js.go?s=1912:1992#L54)
Error wraps a JavaScript error.


<pre>type Error struct {
    <span class="comment">// Value is the underlying JavaScript error value.</span>
    <a href="#Value">Value</a>
}
</pre>











### <a id="Error.Error">func</a> (Error) [Error](https://golang.org/src/syscall/js/js.go?s=2035:2064#L60)
<pre>func (e <a href="#Error">Error</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>
Error implements the error interface.




## <a id="Func">type</a> [Func](https://golang.org/src/syscall/js/func.go?s=448:545#L10)
Func is a wrapped Go function to be called by JavaScript.


<pre>type Func struct {
    <a href="#Value">Value</a> <span class="comment">// the JavaScript function that invokes the Go function</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="FuncOf">func</a> [FuncOf](https://golang.org/src/syscall/js/func.go?s=1420:1483#L28)
<pre>func FuncOf(fn func(this <a href="#Value">Value</a>, args []<a href="#Value">Value</a>) interface{}) <a href="#Func">Func</a></pre>
FuncOf returns a wrapped function.

Invoking the JavaScript function will synchronously call the Go function fn with the value of JavaScript's
"this" keyword and the arguments of the invocation.
The return value of the invocation is the result of the Go function mapped back to JavaScript according to ValueOf.

A wrapped function triggered during a call from Go to JavaScript gets executed on the same goroutine.
A wrapped function triggered by JavaScript's event loop gets executed on an extra goroutine.
Blocking operations in the wrapped function will block the event loop.
As a consequence, if one wrapped function blocks, other wrapped funcs will not be processed.
A blocking function should therefore explicitly start a new goroutine.

Func.Release must be called to free up resources when the function will not be used any more.



<a id="example_FuncOf">Example</a>


```go
```

output:
```txt
```




### <a id="Func.Release">func</a> (Func) [Release](https://golang.org/src/syscall/js/func.go?s=1762:1785#L42)
<pre>func (c <a href="#Func">Func</a>) Release()</pre>
Release frees up resources allocated for the function.
The function must not be invoked after calling Release.




## <a id="Type">type</a> [Type](https://golang.org/src/syscall/js/js.go?s=4831:4844#L171)
Type represents the JavaScript type of a Value.


<pre>type Type <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="TypeUndefined">TypeUndefined</span> <a href="#Type">Type</a> = <a href="/pkg/builtin/#iota">iota</a>
    <span id="TypeNull">TypeNull</span>
    <span id="TypeBoolean">TypeBoolean</span>
    <span id="TypeNumber">TypeNumber</span>
    <span id="TypeString">TypeString</span>
    <span id="TypeSymbol">TypeSymbol</span>
    <span id="TypeObject">TypeObject</span>
    <span id="TypeFunction">TypeFunction</span>
)</pre>









### <a id="Type.String">func</a> (Type) [String](https://golang.org/src/syscall/js/js.go?s=4969:4998#L184)
<pre>func (t <a href="#Type">Type</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Value">type</a> [Value](https://golang.org/src/syscall/js/js.go?s=1467:1497#L26)
Value represents a JavaScript value. The zero value is the JavaScript value "undefined".


<pre>type Value struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="Global">func</a> [Global](https://golang.org/src/syscall/js/js.go?s=2805:2824#L89)
<pre>func Global() <a href="#Value">Value</a></pre>
Global returns the JavaScript global object, usually "window" or "global".




### <a id="Null">func</a> [Null](https://golang.org/src/syscall/js/js.go?s=2686:2703#L84)
<pre>func Null() <a href="#Value">Value</a></pre>
Null returns the JavaScript value "null".




### <a id="Undefined">func</a> [Undefined](https://golang.org/src/syscall/js/js.go?s=2590:2612#L79)
<pre>func Undefined() <a href="#Value">Value</a></pre>
Undefined returns the JavaScript value "undefined".




### <a id="ValueOf">func</a> [ValueOf](https://golang.org/src/syscall/js/js.go?s=3509:3542#L107)
<pre>func ValueOf(x interface{}) <a href="#Value">Value</a></pre>
ValueOf returns x as a JavaScript value:


	| Go                     | JavaScript             |
	| ---------------------- | ---------------------- |
	| js.Value               | [its value]            |
	| js.Func                | function               |
	| nil                    | null                   |
	| bool                   | boolean                |
	| integers and floats    | number                 |
	| string                 | string                 |
	| []interface{}          | new array              |
	| map[string]interface{} | new object             |

Panics if x is not one of the expected types.






### <a id="Value.Bool">func</a> (Value) [Bool](https://golang.org/src/syscall/js/js.go?s=10343:10369#L382)
<pre>func (v <a href="#Value">Value</a>) Bool() <a href="/pkg/builtin/#bool">bool</a></pre>
Bool returns the value v as a bool.
It panics if v is not a JavaScript boolean.




### <a id="Value.Call">func</a> (Value) [Call](https://golang.org/src/syscall/js/js.go?s=7901:7957#L304)
<pre>func (v <a href="#Value">Value</a>) Call(m <a href="/pkg/builtin/#string">string</a>, args ...interface{}) <a href="#Value">Value</a></pre>
Call does a JavaScript call to the method m of value v with the given arguments.
It panics if v has no method m.
The arguments get mapped to JavaScript values according to the ValueOf function.




### <a id="Value.Float">func</a> (Value) [Float](https://golang.org/src/syscall/js/js.go?s=10032:10062#L370)
<pre>func (v <a href="#Value">Value</a>) Float() <a href="/pkg/builtin/#float64">float64</a></pre>
Float returns the value v as a float64.
It panics if v is not a JavaScript number.




### <a id="Value.Get">func</a> (Value) [Get](https://golang.org/src/syscall/js/js.go?s=6099:6133#L240)
<pre>func (v <a href="#Value">Value</a>) Get(p <a href="/pkg/builtin/#string">string</a>) <a href="#Value">Value</a></pre>
Get returns the JavaScript property p of value v.
It panics if v is not a JavaScript object.




### <a id="Value.Index">func</a> (Value) [Index](https://golang.org/src/syscall/js/js.go?s=6715:6748#L262)
<pre>func (v <a href="#Value">Value</a>) Index(i <a href="/pkg/builtin/#int">int</a>) <a href="#Value">Value</a></pre>
Index returns JavaScript index i of value v.
It panics if v is not a JavaScript object.




### <a id="Value.InstanceOf">func</a> (Value) [InstanceOf](https://golang.org/src/syscall/js/js.go?s=12199:12238#L452)
<pre>func (v <a href="#Value">Value</a>) InstanceOf(t <a href="#Value">Value</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
InstanceOf reports whether v is an instance of type t according to JavaScript's instanceof operator.




### <a id="Value.Int">func</a> (Value) [Int](https://golang.org/src/syscall/js/js.go?s=10193:10217#L376)
<pre>func (v <a href="#Value">Value</a>) Int() <a href="/pkg/builtin/#int">int</a></pre>
Int returns the value v truncated to an int.
It panics if v is not a JavaScript number.




### <a id="Value.Invoke">func</a> (Value) [Invoke](https://golang.org/src/syscall/js/js.go?s=8644:8692#L323)
<pre>func (v <a href="#Value">Value</a>) Invoke(args ...interface{}) <a href="#Value">Value</a></pre>
Invoke does a JavaScript call of the value v with the given arguments.
It panics if v is not a JavaScript function.
The arguments get mapped to JavaScript values according to the ValueOf function.




### <a id="Value.JSValue">func</a> (Value) [JSValue](https://golang.org/src/syscall/js/js.go?s=1540:1570#L31)
<pre>func (v <a href="#Value">Value</a>) JSValue() <a href="#Value">Value</a></pre>
JSValue implements Wrapper interface.




### <a id="Value.Length">func</a> (Value) [Length](https://golang.org/src/syscall/js/js.go?s=7517:7544#L292)
<pre>func (v <a href="#Value">Value</a>) Length() <a href="/pkg/builtin/#int">int</a></pre>
Length returns the JavaScript property "length" of v.
It panics if v is not a JavaScript object.




### <a id="Value.New">func</a> (Value) [New](https://golang.org/src/syscall/js/js.go?s=9231:9276#L339)
<pre>func (v <a href="#Value">Value</a>) New(args ...interface{}) <a href="#Value">Value</a></pre>
New uses JavaScript's "new" operator with value v as constructor and the given arguments.
It panics if v is not a JavaScript function.
The arguments get mapped to JavaScript values according to the ValueOf function.




### <a id="Value.Set">func</a> (Value) [Set](https://golang.org/src/syscall/js/js.go?s=6410:6453#L251)
<pre>func (v <a href="#Value">Value</a>) Set(p <a href="/pkg/builtin/#string">string</a>, x interface{})</pre>
Set sets the JavaScript property p of value v to ValueOf(x).
It panics if v is not a JavaScript object.




### <a id="Value.SetIndex">func</a> (Value) [SetIndex](https://golang.org/src/syscall/js/js.go?s=7030:7075#L273)
<pre>func (v <a href="#Value">Value</a>) SetIndex(i <a href="/pkg/builtin/#int">int</a>, x interface{})</pre>
SetIndex sets the JavaScript index i of value v to ValueOf(x).
It panics if v is not a JavaScript object.




### <a id="Value.String">func</a> (Value) [String](https://golang.org/src/syscall/js/js.go?s=11417:11447#L417)
<pre>func (v <a href="#Value">Value</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns the value v as a string.
String is a special case because of Go's String method convention. Unlike the other getters,
it does not panic if v's Type is not TypeString. Instead, it returns a string of the form "<T>"
or "<T: V>" where T is v's type and V is a string representation of v's value.




### <a id="Value.Truthy">func</a> (Value) [Truthy](https://golang.org/src/syscall/js/js.go?s=10751:10779#L396)
<pre>func (v <a href="#Value">Value</a>) Truthy() <a href="/pkg/builtin/#bool">bool</a></pre>
Truthy returns the JavaScript "truthiness" of the value v. In JavaScript,
false, 0, "", null, undefined, and NaN are "falsy", and everything else is
"truthy". See <a href="https://developer.mozilla.org/en-US/docs/Glossary/Truthy">https://developer.mozilla.org/en-US/docs/Glossary/Truthy</a>.




### <a id="Value.Type">func</a> (Value) [Type](https://golang.org/src/syscall/js/js.go?s=5591:5617#L213)
<pre>func (v <a href="#Value">Value</a>) Type() <a href="#Type">Type</a></pre>
Type returns the JavaScript type of the value v. It is similar to JavaScript's typeof operator,
except that it returns TypeNull instead of TypeObject for null.




## <a id="ValueError">type</a> [ValueError](https://golang.org/src/syscall/js/js.go?s=12481:12535#L461)
A ValueError occurs when a Value method is invoked on
a Value that does not support it. Such cases are documented
in the description of each method.


<pre>type ValueError struct {
<span id="ValueError.Method"></span>    Method <a href="/pkg/builtin/#string">string</a>
<span id="ValueError.Type"></span>    Type   <a href="#Type">Type</a>
}
</pre>











### <a id="ValueError.Error">func</a> (\*ValueError) [Error](https://golang.org/src/syscall/js/js.go?s=12537:12572#L466)
<pre>func (e *<a href="#ValueError">ValueError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Wrapper">type</a> [Wrapper](https://golang.org/src/syscall/js/js.go?s=1264:1373#L20)
Wrapper is implemented by types that are backed by a JavaScript value.


<pre>type Wrapper interface {
    <span class="comment">// JSValue returns a JavaScript value associated with an object.</span>
    JSValue() <a href="#Value">Value</a>
}</pre>














