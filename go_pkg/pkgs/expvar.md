

# expvar
`import "expvar"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package expvar provides a standardized interface to public variables, such
as operation counters in servers. It exposes these variables via HTTP at
/debug/vars in JSON format.

Operations to set or modify these public variables are atomic.

In addition to adding the HTTP handler, this package registers the
following variables:


	cmdline   os.Args
	memstats  runtime.Memstats

The package is sometimes only imported for the side effect of
registering its HTTP handler and the above variables. To use it
this way, link this package into your program:


	import _ "expvar"




## <a id="pkg-index">Index</a>
* [func Do(f func(KeyValue))](#Do)
* [func Handler() http.Handler](#Handler)
* [func Publish(name string, v Var)](#Publish)
* [type Float](#Float)
  * [func NewFloat(name string) *Float](#NewFloat)
  * [func (v *Float) Add(delta float64)](#Float.Add)
  * [func (v *Float) Set(value float64)](#Float.Set)
  * [func (v *Float) String() string](#Float.String)
  * [func (v *Float) Value() float64](#Float.Value)
* [type Func](#Func)
  * [func (f Func) String() string](#Func.String)
  * [func (f Func) Value() interface{}](#Func.Value)
* [type Int](#Int)
  * [func NewInt(name string) *Int](#NewInt)
  * [func (v *Int) Add(delta int64)](#Int.Add)
  * [func (v *Int) Set(value int64)](#Int.Set)
  * [func (v *Int) String() string](#Int.String)
  * [func (v *Int) Value() int64](#Int.Value)
* [type KeyValue](#KeyValue)
* [type Map](#Map)
  * [func NewMap(name string) *Map](#NewMap)
  * [func (v *Map) Add(key string, delta int64)](#Map.Add)
  * [func (v *Map) AddFloat(key string, delta float64)](#Map.AddFloat)
  * [func (v *Map) Delete(key string)](#Map.Delete)
  * [func (v *Map) Do(f func(KeyValue))](#Map.Do)
  * [func (v *Map) Get(key string) Var](#Map.Get)
  * [func (v *Map) Init() *Map](#Map.Init)
  * [func (v *Map) Set(key string, av Var)](#Map.Set)
  * [func (v *Map) String() string](#Map.String)
* [type String](#String)
  * [func NewString(name string) *String](#NewString)
  * [func (v *String) Set(value string)](#String.Set)
  * [func (v *String) String() string](#String.String)
  * [func (v *String) Value() string](#String.Value)
* [type Var](#Var)
  * [func Get(name string) Var](#Get)




#### <a id="pkg-files">Package files</a>
[expvar.go](https://golang.org/src/expvar/expvar.go) 






## <a id="Do">func</a> [Do](https://golang.org/src/expvar/expvar.go?s=7126:7151#L313)
<pre>func Do(f func(<a href="#KeyValue">KeyValue</a>))</pre>
Do calls f for each exported variable.
The global variable map is locked during the iteration,
but existing entries may be concurrently updated.



## <a id="Handler">func</a> [Handler](https://golang.org/src/expvar/expvar.go?s=7735:7762#L339)
<pre>func Handler() <a href="/pkg/net/http/">http</a>.<a href="/pkg/net/http/#Handler">Handler</a></pre>
Handler returns the expvar HTTP Handler.

This is only needed to install the handler in a non-standard location.



## <a id="Publish">func</a> [Publish](https://golang.org/src/expvar/expvar.go?s=6154:6186#L266)
<pre>func Publish(name <a href="/pkg/builtin/#string">string</a>, v <a href="#Var">Var</a>)</pre>
Publish declares a named exported variable. This should be called from a
package's init function when it creates its Vars. If the name is already
registered then this will log.Panic.





## <a id="Float">type</a> [Float](https://golang.org/src/expvar/expvar.go?s=1643:1674#L59)
Float is a 64-bit float variable that satisfies the Var interface.


<pre>type Float struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewFloat">func</a> [NewFloat](https://golang.org/src/expvar/expvar.go?s=6716:6749#L292)
<pre>func NewFloat(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Float">Float</a></pre>





### <a id="Float.Add">func</a> (\*Float) [Add](https://golang.org/src/expvar/expvar.go?s=1919:1953#L73)
<pre>func (v *<a href="#Float">Float</a>) Add(delta <a href="/pkg/builtin/#float64">float64</a>)</pre>
Add adds delta to v.




### <a id="Float.Set">func</a> (\*Float) [Set](https://golang.org/src/expvar/expvar.go?s=2190:2224#L86)
<pre>func (v *<a href="#Float">Float</a>) Set(value <a href="/pkg/builtin/#float64">float64</a>)</pre>
Set sets v to value.




### <a id="Float.String">func</a> (\*Float) [String](https://golang.org/src/expvar/expvar.go?s=1767:1798#L67)
<pre>func (v *<a href="#Float">Float</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Float.Value">func</a> (\*Float) [Value](https://golang.org/src/expvar/expvar.go?s=1676:1707#L63)
<pre>func (v *<a href="#Float">Float</a>) Value() <a href="/pkg/builtin/#float64">float64</a></pre>



## <a id="Func">type</a> [Func](https://golang.org/src/expvar/expvar.go?s=5672:5700#L245)
Func implements Var by calling the function
and formatting the returned value using JSON.


<pre>type Func func() interface{}</pre>











### <a id="Func.String">func</a> (Func) [String](https://golang.org/src/expvar/expvar.go?s=5753:5782#L251)
<pre>func (f <a href="#Func">Func</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Func.Value">func</a> (Func) [Value](https://golang.org/src/expvar/expvar.go?s=5702:5735#L247)
<pre>func (f <a href="#Func">Func</a>) Value() interface{}</pre>



## <a id="Int">type</a> [Int](https://golang.org/src/expvar/expvar.go?s=1256:1284#L38)
Int is a 64-bit integer variable that satisfies the Var interface.


<pre>type Int struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewInt">func</a> [NewInt](https://golang.org/src/expvar/expvar.go?s=6638:6667#L286)
<pre>func NewInt(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Int">Int</a></pre>





### <a id="Int.Add">func</a> (\*Int) [Add](https://golang.org/src/expvar/expvar.go?s=1439:1469#L50)
<pre>func (v *<a href="#Int">Int</a>) Add(delta <a href="/pkg/builtin/#int64">int64</a>)</pre>



### <a id="Int.Set">func</a> (\*Int) [Set](https://golang.org/src/expvar/expvar.go?s=1505:1535#L54)
<pre>func (v *<a href="#Int">Int</a>) Set(value <a href="/pkg/builtin/#int64">int64</a>)</pre>



### <a id="Int.String">func</a> (\*Int) [String](https://golang.org/src/expvar/expvar.go?s=1350:1379#L46)
<pre>func (v *<a href="#Int">Int</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Int.Value">func</a> (\*Int) [Value](https://golang.org/src/expvar/expvar.go?s=1286:1313#L42)
<pre>func (v *<a href="#Int">Int</a>) Value() <a href="/pkg/builtin/#int64">int64</a></pre>



## <a id="KeyValue">type</a> [KeyValue](https://golang.org/src/expvar/expvar.go?s=2506:2555#L98)
KeyValue represents a single entry in a Map.


<pre>type KeyValue struct {
<span id="KeyValue.Key"></span>    Key   <a href="/pkg/builtin/#string">string</a>
<span id="KeyValue.Value"></span>    Value <a href="#Var">Var</a>
}
</pre>











## <a id="Map">type</a> [Map](https://golang.org/src/expvar/expvar.go?s=2354:2456#L91)
Map is a string-to-Var map variable that satisfies the Var interface.


<pre>type Map struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewMap">func</a> [NewMap](https://golang.org/src/expvar/expvar.go?s=6800:6829#L298)
<pre>func NewMap(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Map">Map</a></pre>





### <a id="Map.Add">func</a> (\*Map) [Add](https://golang.org/src/expvar/expvar.go?s=3971:4013#L165)
<pre>func (v *<a href="#Map">Map</a>) Add(key <a href="/pkg/builtin/#string">string</a>, delta <a href="/pkg/builtin/#int64">int64</a>)</pre>
Add adds delta to the *Int value stored under the given map key.




### <a id="Map.AddFloat">func</a> (\*Map) [AddFloat](https://golang.org/src/expvar/expvar.go?s=4304:4353#L182)
<pre>func (v *<a href="#Map">Map</a>) AddFloat(key <a href="/pkg/builtin/#string">string</a>, delta <a href="/pkg/builtin/#float64">float64</a>)</pre>
AddFloat adds delta to the *Float value stored under the given map key.




### <a id="Map.Delete">func</a> (\*Map) [Delete](https://golang.org/src/expvar/expvar.go?s=4621:4653#L199)
<pre>func (v *<a href="#Map">Map</a>) Delete(key <a href="/pkg/builtin/#string">string</a>)</pre>
Delete deletes the given key from the map.




### <a id="Map.Do">func</a> (\*Map) [Do](https://golang.org/src/expvar/expvar.go?s=4986:5020#L212)
<pre>func (v *<a href="#Map">Map</a>) Do(f func(<a href="#KeyValue">KeyValue</a>))</pre>
Do calls f for each entry in the map.
The map is locked during the iteration,
but existing entries may be concurrently updated.




### <a id="Map.Get">func</a> (\*Map) [Get](https://golang.org/src/expvar/expvar.go?s=3455:3488#L144)
<pre>func (v *<a href="#Map">Map</a>) Get(key <a href="/pkg/builtin/#string">string</a>) <a href="#Var">Var</a></pre>



### <a id="Map.Init">func</a> (\*Map) [Init](https://golang.org/src/expvar/expvar.go?s=2867:2892#L119)
<pre>func (v *<a href="#Map">Map</a>) Init() *<a href="#Map">Map</a></pre>
Init removes all keys from the map.




### <a id="Map.Set">func</a> (\*Map) [Set](https://golang.org/src/expvar/expvar.go?s=3546:3583#L150)
<pre>func (v *<a href="#Map">Map</a>) Set(key <a href="/pkg/builtin/#string">string</a>, av <a href="#Var">Var</a>)</pre>



### <a id="Map.String">func</a> (\*Map) [String](https://golang.org/src/expvar/expvar.go?s=2557:2586#L103)
<pre>func (v *<a href="#Map">Map</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="String">type</a> [String](https://golang.org/src/expvar/expvar.go?s=5214:5262#L222)
String is a string variable, and satisfies the Var interface.


<pre>type String struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewString">func</a> [NewString](https://golang.org/src/expvar/expvar.go?s=6885:6920#L304)
<pre>func NewString(name <a href="/pkg/builtin/#string">string</a>) *<a href="#String">String</a></pre>





### <a id="String.Set">func</a> (\*String) [Set](https://golang.org/src/expvar/expvar.go?s=5518:5552#L239)
<pre>func (v *<a href="#String">String</a>) Set(value <a href="/pkg/builtin/#string">string</a>)</pre>



### <a id="String.String">func</a> (\*String) [String](https://golang.org/src/expvar/expvar.go?s=5421:5453#L233)
<pre>func (v *<a href="#String">String</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String implements the Var interface. To get the unquoted string
use Value.




### <a id="String.Value">func</a> (\*String) [Value](https://golang.org/src/expvar/expvar.go?s=5264:5295#L226)
<pre>func (v *<a href="#String">String</a>) Value() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Var">type</a> [Var](https://golang.org/src/expvar/expvar.go?s=978:1184#L30)
Var is an abstract type for all exported variables.


<pre>type Var interface {
    <span class="comment">// String returns a valid JSON value for the variable.</span>
    <span class="comment">// Types with String methods that do not return valid JSON</span>
    <span class="comment">// (such as time.Time) must not be used as a Var.</span>
    String() <a href="/pkg/builtin/#string">string</a>
}</pre>









### <a id="Get">func</a> [Get](https://golang.org/src/expvar/expvar.go?s=6492:6517#L278)
<pre>func Get(name <a href="/pkg/builtin/#string">string</a>) <a href="#Var">Var</a></pre>
Get retrieves a named exported variable. It returns nil if the name has
not been registered.










