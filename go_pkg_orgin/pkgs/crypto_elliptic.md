

# elliptic
`import "crypto/elliptic"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package elliptic implements several standard elliptic curves over prime
fields.




## <a id="pkg-index">Index</a>
* [func GenerateKey(curve Curve, rand io.Reader) (priv []byte, x, y *big.Int, err error)](#GenerateKey)
* [func Marshal(curve Curve, x, y *big.Int) []byte](#Marshal)
* [func Unmarshal(curve Curve, data []byte) (x, y *big.Int)](#Unmarshal)
* [type Curve](#Curve)
  * [func P224() Curve](#P224)
  * [func P256() Curve](#P256)
  * [func P384() Curve](#P384)
  * [func P521() Curve](#P521)
* [type CurveParams](#CurveParams)
  * [func (curve *CurveParams) Add(x1, y1, x2, y2 *big.Int) (*big.Int, *big.Int)](#CurveParams.Add)
  * [func (curve *CurveParams) Double(x1, y1 *big.Int) (*big.Int, *big.Int)](#CurveParams.Double)
  * [func (curve *CurveParams) IsOnCurve(x, y *big.Int) bool](#CurveParams.IsOnCurve)
  * [func (curve *CurveParams) Params() *CurveParams](#CurveParams.Params)
  * [func (curve *CurveParams) ScalarBaseMult(k []byte) (*big.Int, *big.Int)](#CurveParams.ScalarBaseMult)
  * [func (curve *CurveParams) ScalarMult(Bx, By *big.Int, k []byte) (*big.Int, *big.Int)](#CurveParams.ScalarMult)




#### <a id="pkg-files">Package files</a>
[elliptic.go](https://golang.org/src/crypto/elliptic/elliptic.go) [p224.go](https://golang.org/src/crypto/elliptic/p224.go) [p256_asm.go](https://golang.org/src/crypto/elliptic/p256_asm.go) 






## <a id="GenerateKey">func</a> [GenerateKey](https://golang.org/src/crypto/elliptic/elliptic.go?s=7397:7482#L267)
<pre>func GenerateKey(curve <a href="#Curve">Curve</a>, rand <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) (priv []<a href="/pkg/builtin/#byte">byte</a>, x, y *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
GenerateKey returns a public/private key pair. The private key is
generated using the given reader, which must return random data.



## <a id="Marshal">func</a> [Marshal](https://golang.org/src/crypto/elliptic/elliptic.go?s=8258:8305#L296)
<pre>func Marshal(curve <a href="#Curve">Curve</a>, x, y *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>) []<a href="/pkg/builtin/#byte">byte</a></pre>
Marshal converts a point into the uncompressed form specified in section 4.3.6 of ANSI X9.62.



## <a id="Unmarshal">func</a> [Unmarshal](https://golang.org/src/crypto/elliptic/elliptic.go?s=8747:8803#L312)
<pre>func Unmarshal(curve <a href="#Curve">Curve</a>, data []<a href="/pkg/builtin/#byte">byte</a>) (x, y *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>)</pre>
Unmarshal converts a point, serialized by Marshal, into an x, y pair.
It is an error if the point is not in uncompressed form or is not on the curve.
On error, x = nil.





## <a id="Curve">type</a> [Curve](https://golang.org/src/crypto/elliptic/elliptic.go?s=872:1510#L14)
A Curve represents a short-form Weierstrass curve with a=-3.
See <a href="https://www.hyperelliptic.org/EFD/g1p/auto-shortw.html">https://www.hyperelliptic.org/EFD/g1p/auto-shortw.html</a>


<pre>type Curve interface {
    <span class="comment">// Params returns the parameters for the curve.</span>
    Params() *<a href="#CurveParams">CurveParams</a>
    <span class="comment">// IsOnCurve reports whether the given (x,y) lies on the curve.</span>
    IsOnCurve(x, y *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>) <a href="/pkg/builtin/#bool">bool</a>
    <span class="comment">// Add returns the sum of (x1,y1) and (x2,y2)</span>
    Add(x1, y1, x2, y2 *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>) (x, y *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>)
    <span class="comment">// Double returns 2*(x,y)</span>
    Double(x1, y1 *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>) (x, y *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>)
    <span class="comment">// ScalarMult returns k*(Bx,By) where k is a number in big-endian form.</span>
    ScalarMult(x1, y1 *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>, k []<a href="/pkg/builtin/#byte">byte</a>) (x, y *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>)
    <span class="comment">// ScalarBaseMult returns k*G, where G is the base point of the group</span>
    <span class="comment">// and k is an integer in big-endian form.</span>
    ScalarBaseMult(k []<a href="/pkg/builtin/#byte">byte</a>) (x, y *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>)
}</pre>









### <a id="P224">func</a> [P224](https://golang.org/src/crypto/elliptic/p224.go?s=1375:1392#L31)
<pre>func P224() <a href="#Curve">Curve</a></pre>
P224 returns a Curve which implements P-224 (see FIPS 186-3, section D.2.2).

The cryptographic operations are implemented using constant-time algorithms.




### <a id="P256">func</a> [P256](https://golang.org/src/crypto/elliptic/elliptic.go?s=11372:11389#L368)
<pre>func P256() <a href="#Curve">Curve</a></pre>
P256 returns a Curve which implements P-256 (see FIPS 186-3, section D.2.3)

The cryptographic operations are implemented using constant-time algorithms.




### <a id="P384">func</a> [P384](https://golang.org/src/crypto/elliptic/elliptic.go?s=11581:11598#L376)
<pre>func P384() <a href="#Curve">Curve</a></pre>
P384 returns a Curve which implements P-384 (see FIPS 186-3, section D.2.4)

The cryptographic operations do not use constant-time algorithms.




### <a id="P521">func</a> [P521](https://golang.org/src/crypto/elliptic/elliptic.go?s=11790:11807#L384)
<pre>func P521() <a href="#Curve">Curve</a></pre>
P521 returns a Curve which implements P-521 (see FIPS 186-3, section D.2.5)

The cryptographic operations do not use constant-time algorithms.






## <a id="CurveParams">type</a> [CurveParams](https://golang.org/src/crypto/elliptic/elliptic.go?s=1647:1986#L32)
CurveParams contains the parameters of an elliptic curve and also provides
a generic, non-constant time implementation of Curve.


<pre>type CurveParams struct {
<span id="CurveParams.P"></span>    P       *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a> <span class="comment">// the order of the underlying field</span>
<span id="CurveParams.N"></span>    N       *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a> <span class="comment">// the order of the base point</span>
<span id="CurveParams.B"></span>    B       *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a> <span class="comment">// the constant of the curve equation</span>
<span id="CurveParams.Gx"></span>    Gx, Gy  *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a> <span class="comment">// (x,y) of the base point</span>
<span id="CurveParams.BitSize"></span>    BitSize <a href="/pkg/builtin/#int">int</a>      <span class="comment">// the size of the underlying field</span>
<span id="CurveParams.Name"></span>    Name    <a href="/pkg/builtin/#string">string</a>   <span class="comment">// the canonical name of the curve</span>
}
</pre>











### <a id="CurveParams.Add">func</a> (\*CurveParams) [Add](https://golang.org/src/crypto/elliptic/elliptic.go?s=3250:3325#L92)
<pre>func (curve *<a href="#CurveParams">CurveParams</a>) Add(x1, y1, x2, y2 *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>) (*<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>, *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>)</pre>



### <a id="CurveParams.Double">func</a> (\*CurveParams) [Double](https://golang.org/src/crypto/elliptic/elliptic.go?s=5097:5167#L176)
<pre>func (curve *<a href="#CurveParams">CurveParams</a>) Double(x1, y1 *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>) (*<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>, *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>)</pre>



### <a id="CurveParams.IsOnCurve">func</a> (\*CurveParams) [IsOnCurve](https://golang.org/src/crypto/elliptic/elliptic.go?s=2055:2110#L45)
<pre>func (curve *<a href="#CurveParams">CurveParams</a>) IsOnCurve(x, y *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>) <a href="/pkg/builtin/#bool">bool</a></pre>



### <a id="CurveParams.Params">func</a> (\*CurveParams) [Params](https://golang.org/src/crypto/elliptic/elliptic.go?s=1988:2035#L41)
<pre>func (curve *<a href="#CurveParams">CurveParams</a>) Params() *<a href="#CurveParams">CurveParams</a></pre>



### <a id="CurveParams.ScalarBaseMult">func</a> (\*CurveParams) [ScalarBaseMult](https://golang.org/src/crypto/elliptic/elliptic.go?s=7072:7143#L259)
<pre>func (curve *<a href="#CurveParams">CurveParams</a>) ScalarBaseMult(k []<a href="/pkg/builtin/#byte">byte</a>) (*<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>, *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>)</pre>



### <a id="CurveParams.ScalarMult">func</a> (\*CurveParams) [ScalarMult](https://golang.org/src/crypto/elliptic/elliptic.go?s=6637:6721#L242)
<pre>func (curve *<a href="#CurveParams">CurveParams</a>) ScalarMult(Bx, By *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>, k []<a href="/pkg/builtin/#byte">byte</a>) (*<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>, *<a href="/pkg/math/big/">big</a>.<a href="/pkg/math/big/#Int">Int</a>)</pre>






