

# subtle
`import "crypto/subtle"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package subtle implements functions that are often useful in cryptographic
code but require careful thought to use correctly.




## <a id="pkg-index">Index</a>
* [func ConstantTimeByteEq(x, y uint8) int](#ConstantTimeByteEq)
* [func ConstantTimeCompare(x, y []byte) int](#ConstantTimeCompare)
* [func ConstantTimeCopy(v int, x, y []byte)](#ConstantTimeCopy)
* [func ConstantTimeEq(x, y int32) int](#ConstantTimeEq)
* [func ConstantTimeLessOrEq(x, y int) int](#ConstantTimeLessOrEq)
* [func ConstantTimeSelect(v, x, y int) int](#ConstantTimeSelect)




#### <a id="pkg-files">Package files</a>
[constant_time.go](https://golang.org/src/crypto/subtle/constant_time.go) 






## <a id="ConstantTimeByteEq">func</a> [ConstantTimeByteEq](https://golang.org/src/crypto/subtle/constant_time.go?s=937:976#L21)
<pre>func ConstantTimeByteEq(x, y <a href="/pkg/builtin/#uint8">uint8</a>) <a href="/pkg/builtin/#int">int</a></pre>
ConstantTimeByteEq returns 1 if x == y and 0 otherwise.



## <a id="ConstantTimeCompare">func</a> [ConstantTimeCompare](https://golang.org/src/crypto/subtle/constant_time.go?s=505:546#L2)
<pre>func ConstantTimeCompare(x, y []<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#int">int</a></pre>
ConstantTimeCompare returns 1 if the two slices, x and y, have equal contents
and 0 otherwise. The time taken is a function of the length of the slices and
is independent of the contents.



## <a id="ConstantTimeCopy">func</a> [ConstantTimeCopy](https://golang.org/src/crypto/subtle/constant_time.go?s=1341:1382#L33)
<pre>func ConstantTimeCopy(v <a href="/pkg/builtin/#int">int</a>, x, y []<a href="/pkg/builtin/#byte">byte</a>)</pre>
ConstantTimeCopy copies the contents of y into x (a slice of equal length)
if v == 1. If v == 0, x is left unchanged. Its behavior is undefined if v
takes any other value.



## <a id="ConstantTimeEq">func</a> [ConstantTimeEq](https://golang.org/src/crypto/subtle/constant_time.go?s=1074:1109#L26)
<pre>func ConstantTimeEq(x, y <a href="/pkg/builtin/#int32">int32</a>) <a href="/pkg/builtin/#int">int</a></pre>
ConstantTimeEq returns 1 if x == y and 0 otherwise.



## <a id="ConstantTimeLessOrEq">func</a> [ConstantTimeLessOrEq](https://golang.org/src/crypto/subtle/constant_time.go?s=1707:1746#L47)
<pre>func ConstantTimeLessOrEq(x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
ConstantTimeLessOrEq returns 1 if x <= y and 0 otherwise.
Its behavior is undefined if x or y are negative or > 2**31 - 1.



## <a id="ConstantTimeSelect">func</a> [ConstantTimeSelect](https://golang.org/src/crypto/subtle/constant_time.go?s=806:846#L18)
<pre>func ConstantTimeSelect(v, x, y <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#int">int</a></pre>
ConstantTimeSelect returns x if v == 1 and y if v == 0.
Its behavior is undefined if v takes any other value.








