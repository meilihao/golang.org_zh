

# binary
`import "encoding/binary"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package binary implements simple translation between numbers and byte
sequences and encoding and decoding of varints.

Numbers are translated by reading and writing fixed-size values.
A fixed-size value is either a fixed-size arithmetic
type (bool, int8, uint8, int16, float32, complex64, ...)
or an array or struct containing only fixed-size values.

The varint functions encode and decode single integer values using
a variable-length encoding; smaller values require fewer bytes.
For a specification, see
<a href="https://developers.google.com/protocol-buffers/docs/encoding">https://developers.google.com/protocol-buffers/docs/encoding</a>.

This package favors simplicity over efficiency. Clients that require
high-performance serialization, especially for large data structures,
should look at more advanced solutions such as the encoding/gob
package or protocol buffers.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func PutUvarint(buf []byte, x uint64) int](#PutUvarint)
* [func PutVarint(buf []byte, x int64) int](#PutVarint)
* [func Read(r io.Reader, order ByteOrder, data interface{}) error](#Read)
* [func ReadUvarint(r io.ByteReader) (uint64, error)](#ReadUvarint)
* [func ReadVarint(r io.ByteReader) (int64, error)](#ReadVarint)
* [func Size(v interface{}) int](#Size)
* [func Uvarint(buf []byte) (uint64, int)](#Uvarint)
* [func Varint(buf []byte) (int64, int)](#Varint)
* [func Write(w io.Writer, order ByteOrder, data interface{}) error](#Write)
* [type ByteOrder](#ByteOrder)


#### <a id="pkg-examples">Examples</a>
* [ByteOrder (Get)](#example_ByteOrder_get)
* [ByteOrder (Put)](#example_ByteOrder_put)
* [PutUvarint](#example_PutUvarint)
* [PutVarint](#example_PutVarint)
* [Read](#example_Read)
* [Read (Multi)](#example_Read_multi)
* [Uvarint](#example_Uvarint)
* [Varint](#example_Varint)
* [Write](#example_Write)
* [Write (Multi)](#example_Write_multi)


#### <a id="pkg-files">Package files</a>
[binary.go](https://golang.org/src/encoding/binary/binary.go) [varint.go](https://golang.org/src/encoding/binary/varint.go) 


## <a id="pkg-constants">Constants</a>
MaxVarintLenN is the maximum length of a varint-encoded N-bit integer.


<pre>const (
    <span id="MaxVarintLen16">MaxVarintLen16</span> = 3
    <span id="MaxVarintLen32">MaxVarintLen32</span> = 5
    <span id="MaxVarintLen64">MaxVarintLen64</span> = 10
)</pre>

## <a id="pkg-variables">Variables</a>
BigEndian is the big-endian implementation of ByteOrder.


<pre>var <span id="BigEndian">BigEndian</span> bigEndian</pre>LittleEndian is the little-endian implementation of ByteOrder.


<pre>var <span id="LittleEndian">LittleEndian</span> littleEndian</pre>

## <a id="PutUvarint">func</a> [PutUvarint](https://golang.org/src/encoding/binary/varint.go?s=1611:1652#L31)
<pre>func PutUvarint(buf []<a href="/pkg/builtin/#byte">byte</a>, x <a href="/pkg/builtin/#uint64">uint64</a>) <a href="/pkg/builtin/#int">int</a></pre>
PutUvarint encodes a uint64 into buf and returns the number of bytes written.
If the buffer is too small, PutUvarint will panic.


<a id="example_PutUvarint">Example</a>

## <a id="PutVarint">func</a> [PutVarint](https://golang.org/src/encoding/binary/varint.go?s=2477:2516#L68)
<pre>func PutVarint(buf []<a href="/pkg/builtin/#byte">byte</a>, x <a href="/pkg/builtin/#int64">int64</a>) <a href="/pkg/builtin/#int">int</a></pre>
PutVarint encodes an int64 into buf and returns the number of bytes written.
If the buffer is too small, PutVarint will panic.


<a id="example_PutVarint">Example</a>

## <a id="Read">func</a> [Read](https://golang.org/src/encoding/binary/binary.go?s=5180:5243#L151)
<pre>func Read(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>, order <a href="#ByteOrder">ByteOrder</a>, data interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Read reads structured binary data from r into data.
Data must be a pointer to a fixed-size value or a slice
of fixed-size values.
Bytes read from r are decoded using the specified byte order
and written to successive fields of the data.
When decoding boolean values, a zero byte is decoded as false, and
any other non-zero byte is decoded as true.
When reading into structs, the field data for fields with
blank (_) field names is skipped; i.e., blank field names
may be used for padding.
When reading into a struct, all non-blank fields must be exported
or Read may panic.

The error is EOF only if no bytes were read.
If an EOF happens after reading some but not all the bytes,
Read returns ErrUnexpectedEOF.


<a id="example_Read">Example</a>
<a id="example_Read_multi">Example (Multi)</a>

## <a id="ReadUvarint">func</a> [ReadUvarint](https://golang.org/src/encoding/binary/varint.go?s=3248:3297#L96)
<pre>func ReadUvarint(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ByteReader">ByteReader</a>) (<a href="/pkg/builtin/#uint64">uint64</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadUvarint reads an encoded unsigned integer from r and returns it as a uint64.



## <a id="ReadVarint">func</a> [ReadVarint](https://golang.org/src/encoding/binary/varint.go?s=3647:3694#L116)
<pre>func ReadVarint(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ByteReader">ByteReader</a>) (<a href="/pkg/builtin/#int64">int64</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ReadVarint reads an encoded signed integer from r and returns it as an int64.



## <a id="Size">func</a> [Size](https://golang.org/src/encoding/binary/binary.go?s=9811:9839#L352)
<pre>func Size(v interface{}) <a href="/pkg/builtin/#int">int</a></pre>
Size returns how many bytes Write would generate to encode the value v, which
must be a fixed-size value or a slice of fixed-size values, or a pointer to such data.
If v is neither of these, Size returns -1.



## <a id="Uvarint">func</a> [Uvarint](https://golang.org/src/encoding/binary/varint.go?s=2070:2108#L50)
<pre>func Uvarint(buf []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#uint64">uint64</a>, <a href="/pkg/builtin/#int">int</a>)</pre>
Uvarint decodes a uint64 from buf and returns that value and the
number of bytes read (> 0). If an error occurred, the value is 0
and the number of bytes n is <= 0 meaning:


	n == 0: buf too small
	n  < 0: value larger than 64 bits (overflow)
	        and -n is the number of bytes read


<a id="example_Uvarint">Example</a>

## <a id="Varint">func</a> [Varint](https://golang.org/src/encoding/binary/varint.go?s=2926:2962#L84)
<pre>func Varint(buf []<a href="/pkg/builtin/#byte">byte</a>) (<a href="/pkg/builtin/#int64">int64</a>, <a href="/pkg/builtin/#int">int</a>)</pre>
Varint decodes an int64 from buf and returns that value and the
number of bytes read (> 0). If an error occurred, the value is 0
and the number of bytes n is <= 0 with the following meaning:


	n == 0: buf too small
	n  < 0: value larger than 64 bits (overflow)
	        and -n is the number of bytes read


<a id="example_Varint">Example</a>

## <a id="Write">func</a> [Write](https://golang.org/src/encoding/binary/binary.go?s=7469:7533#L244)
<pre>func Write(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, order <a href="#ByteOrder">ByteOrder</a>, data interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Write writes the binary representation of data into w.
Data must be a fixed-size value or a slice of fixed-size
values, or a pointer to such data.
Boolean values encode as one byte: 1 for true, and 0 for false.
Bytes written to w are encoded using the specified byte order
and read from successive fields of the data.
When writing structs, zero values are written for fields
with blank (_) field names.


<a id="example_Write">Example</a>
<a id="example_Write_multi">Example (Multi)</a>



## <a id="ByteOrder">type</a> [ByteOrder](https://golang.org/src/encoding/binary/binary.go?s=1176:1371#L23)
A ByteOrder specifies how to convert byte sequences into
16-, 32-, or 64-bit unsigned integers.


<pre>type ByteOrder interface {
    Uint16([]<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#uint16">uint16</a>
    Uint32([]<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#uint32">uint32</a>
    Uint64([]<a href="/pkg/builtin/#byte">byte</a>) <a href="/pkg/builtin/#uint64">uint64</a>
    PutUint16([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#uint16">uint16</a>)
    PutUint32([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#uint32">uint32</a>)
    PutUint64([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#uint64">uint64</a>)
    String() <a href="/pkg/builtin/#string">string</a>
}</pre>





<a id="example_ByteOrder_get">Example (Get)</a>
<a id="example_ByteOrder_put">Example (Put)</a>










