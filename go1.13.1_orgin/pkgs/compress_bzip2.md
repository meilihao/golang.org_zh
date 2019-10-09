

# bzip2
`import "compress/bzip2"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package bzip2 implements bzip2 decompression.




## <a id="pkg-index">Index</a>
* [func NewReader(r io.Reader) io.Reader](#NewReader)
* [type StructuralError](#StructuralError)
  * [func (s StructuralError) Error() string](#StructuralError.Error)




#### <a id="pkg-files">Package files</a>
[bit_reader.go](https://golang.org/src/compress/bzip2/bit_reader.go) [bzip2.go](https://golang.org/src/compress/bzip2/bzip2.go) [huffman.go](https://golang.org/src/compress/bzip2/huffman.go) [move_to_front.go](https://golang.org/src/compress/bzip2/move_to_front.go) 






## <a id="NewReader">func</a> [NewReader](https://golang.org/src/compress/bzip2/bzip2.go?s=1722:1759#L36)
<pre>func NewReader(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a></pre>
NewReader returns an io.Reader which decompresses bzip2 data from r.
If r does not also implement io.ByteReader,
the decompressor may read more data than necessary from r.





## <a id="StructuralError">type</a> [StructuralError](https://golang.org/src/compress/bzip2/bzip2.go?s=566:593#L7)
A StructuralError is returned when the bzip2 data is found to be
syntactically invalid.


<pre>type StructuralError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="StructuralError.Error">func</a> (StructuralError) [Error](https://golang.org/src/compress/bzip2/bzip2.go?s=595:634#L9)
<pre>func (s <a href="#StructuralError">StructuralError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>







