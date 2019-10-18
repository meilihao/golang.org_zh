

# fnv
`import "hash/fnv"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package fnv implements FNV-1 and FNV-1a, non-cryptographic hash functions
created by Glenn Fowler, Landon Curt Noll, and Phong Vo.
See
<a href="https://en.wikipedia.org/wiki/Fowler-Noll-Vo_hash_function">https://en.wikipedia.org/wiki/Fowler-Noll-Vo_hash_function</a>.

All the hash.Hash implementations returned by this package also
implement encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
marshal and unmarshal the internal state of the hash.




## <a id="pkg-index">Index</a>
* [func New128() hash.Hash](#New128)
* [func New128a() hash.Hash](#New128a)
* [func New32() hash.Hash32](#New32)
* [func New32a() hash.Hash32](#New32a)
* [func New64() hash.Hash64](#New64)
* [func New64a() hash.Hash64](#New64a)




#### <a id="pkg-files">Package files</a>
[fnv.go](https://golang.org/src/hash/fnv/fnv.go) 






## <a id="New128">func</a> [New128](https://golang.org/src/hash/fnv/fnv.go?s=1839:1862#L61)
<pre>func New128() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a></pre>
New128 returns a new 128-bit FNV-1 hash.Hash.
Its Sum method will lay the value out in big-endian byte order.



## <a id="New128a">func</a> [New128a](https://golang.org/src/hash/fnv/fnv.go?s=2058:2082#L70)
<pre>func New128a() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash">Hash</a></pre>
New128a returns a new 128-bit FNV-1a hash.Hash.
Its Sum method will lay the value out in big-endian byte order.



## <a id="New32">func</a> [New32](https://golang.org/src/hash/fnv/fnv.go?s=1113:1137#L33)
<pre>func New32() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash32">Hash32</a></pre>
New32 returns a new 32-bit FNV-1 hash.Hash.
Its Sum method will lay the value out in big-endian byte order.



## <a id="New32a">func</a> [New32a](https://golang.org/src/hash/fnv/fnv.go?s=1294:1319#L40)
<pre>func New32a() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash32">Hash32</a></pre>
New32a returns a new 32-bit FNV-1a hash.Hash.
Its Sum method will lay the value out in big-endian byte order.



## <a id="New64">func</a> [New64](https://golang.org/src/hash/fnv/fnv.go?s=1475:1499#L47)
<pre>func New64() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash64">Hash64</a></pre>
New64 returns a new 64-bit FNV-1 hash.Hash.
Its Sum method will lay the value out in big-endian byte order.



## <a id="New64a">func</a> [New64a](https://golang.org/src/hash/fnv/fnv.go?s=1656:1681#L54)
<pre>func New64a() <a href="/pkg/hash/">hash</a>.<a href="/pkg/hash/#Hash64">Hash64</a></pre>
New64a returns a new 64-bit FNV-1a hash.Hash.
Its Sum method will lay the value out in big-endian byte order.









