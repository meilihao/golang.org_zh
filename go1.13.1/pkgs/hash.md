

# hash
`import "hash"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package hash provides interfaces for hash functions.


<a id="example__binaryMarshaler">Example (BinaryMarshaler)</a>


## <a id="pkg-index">Index</a>
* [type Hash](#Hash)
* [type Hash32](#Hash32)
* [type Hash64](#Hash64)


#### <a id="pkg-examples">Examples</a>
* [Package (BinaryMarshaler)](#example__binaryMarshaler)


#### <a id="pkg-files">Package files</a>
[hash.go](https://golang.org/src/hash/hash.go) 








## <a id="Hash">type</a> [Hash](https://golang.org/src/hash/hash.go?s=1238:1887#L16)
Hash is the common interface implemented by all hash functions.

Hash implementations in the standard library (e.g. hash/crc32 and
crypto/sha256) implement the encoding.BinaryMarshaler and
encoding.BinaryUnmarshaler interfaces. Marshaling a hash implementation
allows its internal state to be saved and used for additional processing
later, without having to re-write the data previously written to the hash.
The hash state may contain portions of the input in its original form,
which users are expected to handle for any possible security implications.

Compatibility: Any future changes to hash or crypto packages will endeavor
to maintain compatibility with state encoded using previous versions.
That is, any released versions of the packages should be able to
decode data written with any previously released version,
subject to issues such as security fixes.
See the Go compatibility document for background: <a href="https://golang.org/doc/go1compat">https://golang.org/doc/go1compat</a>


<pre>type Hash interface {
    <span class="comment">// Write (via the embedded io.Writer interface) adds more data to the running hash.</span>
    <span class="comment">// It never returns an error.</span>
    <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>

    <span class="comment">// Sum appends the current hash to b and returns the resulting slice.</span>
    <span class="comment">// It does not change the underlying hash state.</span>
    Sum(b []<a href="/pkg/builtin/#byte">byte</a>) []<a href="/pkg/builtin/#byte">byte</a>

    <span class="comment">// Reset resets the Hash to its initial state.</span>
    Reset()

    <span class="comment">// Size returns the number of bytes Sum will return.</span>
    Size() <a href="/pkg/builtin/#int">int</a>

    <span class="comment">// BlockSize returns the hash&#39;s underlying block size.</span>
    <span class="comment">// The Write method must be able to accept any amount</span>
    <span class="comment">// of data, but it may operate more efficiently if all writes</span>
    <span class="comment">// are a multiple of the block size.</span>
    BlockSize() <a href="/pkg/builtin/#int">int</a>
}</pre>











## <a id="Hash32">type</a> [Hash32](https://golang.org/src/hash/hash.go?s=1965:2012#L39)
Hash32 is the common interface implemented by all 32-bit hash functions.


<pre>type Hash32 interface {
    <a href="#Hash">Hash</a>
    Sum32() <a href="/pkg/builtin/#uint32">uint32</a>
}</pre>











## <a id="Hash64">type</a> [Hash64](https://golang.org/src/hash/hash.go?s=2090:2137#L45)
Hash64 is the common interface implemented by all 64-bit hash functions.


<pre>type Hash64 interface {
    <a href="#Hash">Hash</a>
    Sum64() <a href="/pkg/builtin/#uint64">uint64</a>
}</pre>















