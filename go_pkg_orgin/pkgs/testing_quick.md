

# quick
`import "testing/quick"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package quick implements utility functions to help with black box testing.

The testing/quick package is frozen and is not accepting new features.




## <a id="pkg-index">Index</a>
* [func Check(f interface{}, config *Config) error](#Check)
* [func CheckEqual(f, g interface{}, config *Config) error](#CheckEqual)
* [func Value(t reflect.Type, rand *rand.Rand) (value reflect.Value, ok bool)](#Value)
* [type CheckEqualError](#CheckEqualError)
  * [func (s *CheckEqualError) Error() string](#CheckEqualError.Error)
* [type CheckError](#CheckError)
  * [func (s *CheckError) Error() string](#CheckError.Error)
* [type Config](#Config)
* [type Generator](#Generator)
* [type SetupError](#SetupError)
  * [func (s SetupError) Error() string](#SetupError.Error)




#### <a id="pkg-files">Package files</a>
[quick.go](https://golang.org/src/testing/quick/quick.go) 






## <a id="Check">func</a> [Check](https://golang.org/src/testing/quick/quick.go?s=7499:7546#L253)
<pre>func Check(f interface{}, config *<a href="#Config">Config</a>) <a href="/pkg/builtin/#error">error</a></pre>
Check looks for an input to f, any function that returns bool,
such that f returns false. It calls f repeatedly, with arbitrary
values for each argument. If f returns false on a given input,
Check returns that input as a *CheckError.
For example:


	func TestOddMultipleOfThree(t *testing.T) {
		f := func(x int) bool {
			y := OddMultipleOfThree(x)
			return y%2 == 1 && y%3 == 0
		}
		if err := quick.Check(f, nil); err != nil {
			t.Error(err)
		}
	}



## <a id="CheckEqual">func</a> [CheckEqual](https://golang.org/src/testing/quick/quick.go?s=8512:8567#L292)
<pre>func CheckEqual(f, g interface{}, config *<a href="#Config">Config</a>) <a href="/pkg/builtin/#error">error</a></pre>
CheckEqual looks for an input on which f and g return different results.
It calls f and g repeatedly with arbitrary values for each argument.
If f and g return different answers, CheckEqual returns a *CheckEqualError
describing the input and the outputs.



## <a id="Value">func</a> [Value](https://golang.org/src/testing/quick/quick.go?s=1618:1692#L49)
<pre>func Value(t <a href="/pkg/reflect/">reflect</a>.<a href="/pkg/reflect/#Type">Type</a>, rand *<a href="/pkg/math/rand/">rand</a>.<a href="/pkg/math/rand/#Rand">Rand</a>) (value <a href="/pkg/reflect/">reflect</a>.<a href="/pkg/reflect/#Value">Value</a>, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
Value returns an arbitrary value of the given type.
If the type implements the Generator interface, that will be used.
Note: To create arbitrary values for structs, all the fields must be exported.





## <a id="CheckEqualError">type</a> [CheckEqualError](https://golang.org/src/testing/quick/quick.go?s=6735:6818#L228)
A CheckEqualError is the result CheckEqual finding an error.


<pre>type CheckEqualError struct {
    <a href="#CheckError">CheckError</a>
<span id="CheckEqualError.Out1"></span>    Out1 []interface{}
<span id="CheckEqualError.Out2"></span>    Out2 []interface{}
}
</pre>











### <a id="CheckEqualError.Error">func</a> (\*CheckEqualError) [Error](https://golang.org/src/testing/quick/quick.go?s=6820:6860#L234)
<pre>func (s *<a href="#CheckEqualError">CheckEqualError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="CheckError">type</a> [CheckError](https://golang.org/src/testing/quick/quick.go?s=6498:6556#L218)
A CheckError is the result of Check finding an error.


<pre>type CheckError struct {
<span id="CheckError.Count"></span>    Count <a href="/pkg/builtin/#int">int</a>
<span id="CheckError.In"></span>    In    []interface{}
}
</pre>











### <a id="CheckError.Error">func</a> (\*CheckError) [Error](https://golang.org/src/testing/quick/quick.go?s=6558:6593#L223)
<pre>func (s *<a href="#CheckError">CheckError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Config">type</a> [Config](https://golang.org/src/testing/quick/quick.go?s=4954:5679#L167)
A Config structure contains options for running a test.


<pre>type Config struct {
<span id="Config.MaxCount"></span>    <span class="comment">// MaxCount sets the maximum number of iterations.</span>
    <span class="comment">// If zero, MaxCountScale is used.</span>
    MaxCount <a href="/pkg/builtin/#int">int</a>
<span id="Config.MaxCountScale"></span>    <span class="comment">// MaxCountScale is a non-negative scale factor applied to the</span>
    <span class="comment">// default maximum.</span>
    <span class="comment">// A count of zero implies the default, which is usually 100</span>
    <span class="comment">// but can be set by the -quickchecks flag.</span>
    MaxCountScale <a href="/pkg/builtin/#float64">float64</a>
<span id="Config.Rand"></span>    <span class="comment">// Rand specifies a source of random numbers.</span>
    <span class="comment">// If nil, a default pseudo-random source will be used.</span>
    Rand *<a href="/pkg/math/rand/">rand</a>.<a href="/pkg/math/rand/#Rand">Rand</a>
<span id="Config.Values"></span>    <span class="comment">// Values specifies a function to generate a slice of</span>
    <span class="comment">// arbitrary reflect.Values that are congruent with the</span>
    <span class="comment">// arguments to the function being tested.</span>
    <span class="comment">// If nil, the top-level Value function is used to generate them.</span>
    Values func([]<a href="/pkg/reflect/">reflect</a>.<a href="/pkg/reflect/#Value">Value</a>, *<a href="/pkg/math/rand/">rand</a>.<a href="/pkg/math/rand/#Rand">Rand</a>)
}
</pre>











## <a id="Generator">type</a> [Generator](https://golang.org/src/testing/quick/quick.go?s=575:764#L13)
A Generator can generate random values of its own type.


<pre>type Generator interface {
    <span class="comment">// Generate returns a random instance of the type on which it is a</span>
    <span class="comment">// method using the size as a size hint.</span>
    Generate(rand *<a href="/pkg/math/rand/">rand</a>.<a href="/pkg/math/rand/#Rand">Rand</a>, size <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/reflect/">reflect</a>.<a href="/pkg/reflect/#Value">Value</a>
}</pre>











## <a id="SetupError">type</a> [SetupError](https://golang.org/src/testing/quick/quick.go?s=6360:6382#L213)
A SetupError is the result of an error in the way that check is being
used, independent of the functions being tested.


<pre>type SetupError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="SetupError.Error">func</a> (SetupError) [Error](https://golang.org/src/testing/quick/quick.go?s=6384:6418#L215)
<pre>func (s <a href="#SetupError">SetupError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>






