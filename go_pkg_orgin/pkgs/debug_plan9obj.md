

# plan9obj
`import "debug/plan9obj"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package plan9obj implements access to Plan 9 a.out object files.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [type File](#File)
  * [func NewFile(r io.ReaderAt) (*File, error)](#NewFile)
  * [func Open(name string) (*File, error)](#Open)
  * [func (f *File) Close() error](#File.Close)
  * [func (f *File) Section(name string) *Section](#File.Section)
  * [func (f *File) Symbols() ([]Sym, error)](#File.Symbols)
* [type FileHeader](#FileHeader)
* [type Section](#Section)
  * [func (s *Section) Data() ([]byte, error)](#Section.Data)
  * [func (s *Section) Open() io.ReadSeeker](#Section.Open)
* [type SectionHeader](#SectionHeader)
* [type Sym](#Sym)




#### <a id="pkg-files">Package files</a>
[file.go](https://golang.org/src/debug/plan9obj/file.go) [plan9obj.go](https://golang.org/src/debug/plan9obj/plan9obj.go) 


## <a id="pkg-constants">Constants</a>

<pre>const (
    <span id="Magic64">Magic64</span> = 0x8000 <span class="comment">// 64-bit expanded header</span>

    <span id="Magic386">Magic386</span>   = (4*11+0)*11 + 7
    <span id="MagicAMD64">MagicAMD64</span> = (4*26+0)*26 + 7 + <a href="#Magic64">Magic64</a>
    <span id="MagicARM">MagicARM</span>   = (4*20+0)*20 + 7
)</pre>





## <a id="File">type</a> [File](https://golang.org/src/debug/plan9obj/file.go?s=554:627#L17)
A File represents an open Plan 9 a.out file.


<pre>type File struct {
    <a href="#FileHeader">FileHeader</a>
<span id="File.Sections"></span>    Sections []*<a href="#Section">Section</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewFile">func</a> [NewFile](https://golang.org/src/debug/plan9obj/file.go?s=3073:3115#L125)
<pre>func NewFile(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReaderAt">ReaderAt</a>) (*<a href="#File">File</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewFile creates a new File for accessing a Plan 9 binary in an underlying reader.
The Plan 9 binary is expected to start at position 0 in the ReaderAt.




### <a id="Open">func</a> [Open](https://golang.org/src/debug/plan9obj/file.go?s=2249:2286#L88)
<pre>func Open(name <a href="/pkg/builtin/#string">string</a>) (*<a href="#File">File</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Open opens the named file using os.Open and prepares it for use as a Plan 9 a.out binary.






### <a id="File.Close">func</a> (\*File) [Close](https://golang.org/src/debug/plan9obj/file.go?s=2576:2604#L105)
<pre>func (f *<a href="#File">File</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the File.
If the File was created using NewFile directly instead of Open,
Close has no effect.




### <a id="File.Section">func</a> (\*File) [Section](https://golang.org/src/debug/plan9obj/file.go?s=6726:6770#L311)
<pre>func (f *<a href="#File">File</a>) Section(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Section">Section</a></pre>
Section returns a section with the given name, or nil if no such
section exists.




### <a id="File.Symbols">func</a> (\*File) [Symbols](https://golang.org/src/debug/plan9obj/file.go?s=6332:6371#L295)
<pre>func (f *<a href="#File">File</a>) Symbols() ([]<a href="#Sym">Sym</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Symbols returns the symbol table for f.




## <a id="FileHeader">type</a> [FileHeader](https://golang.org/src/debug/plan9obj/file.go?s=361:504#L7)
A FileHeader represents a Plan 9 a.out file header.


<pre>type FileHeader struct {
<span id="FileHeader.Magic"></span>    Magic       <a href="/pkg/builtin/#uint32">uint32</a>
<span id="FileHeader.Bss"></span>    Bss         <a href="/pkg/builtin/#uint32">uint32</a>
<span id="FileHeader.Entry"></span>    Entry       <a href="/pkg/builtin/#uint64">uint64</a>
<span id="FileHeader.PtrSize"></span>    PtrSize     <a href="/pkg/builtin/#int">int</a>
<span id="FileHeader.LoadAddress"></span>    LoadAddress <a href="/pkg/builtin/#uint64">uint64</a>
<span id="FileHeader.HdrSize"></span>    HdrSize     <a href="/pkg/builtin/#uint64">uint64</a>
}
</pre>











## <a id="Section">type</a> [Section](https://golang.org/src/debug/plan9obj/file.go?s=928:1237#L33)
A Section represents a single section in a Plan 9 a.out file.


<pre>type Section struct {
    <a href="#SectionHeader">SectionHeader</a>

    <span class="comment">// Embed ReaderAt for ReadAt method.</span>
    <span class="comment">// Do not embed SectionReader directly</span>
    <span class="comment">// to avoid having Read and Seek.</span>
    <span class="comment">// If a client wants Read and Seek it must use</span>
    <span class="comment">// Open() to avoid fighting over the seek offset</span>
    <span class="comment">// with other clients.</span>
    <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReaderAt">ReaderAt</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="Section.Data">func</a> (\*Section) [Data](https://golang.org/src/debug/plan9obj/file.go?s=1307:1347#L47)
<pre>func (s *<a href="#Section">Section</a>) Data() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Data reads and returns the contents of the Plan 9 a.out section.




### <a id="Section.Open">func</a> (\*Section) [Open](https://golang.org/src/debug/plan9obj/file.go?s=1542:1580#L57)
<pre>func (s *<a href="#Section">Section</a>) Open() <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadSeeker">ReadSeeker</a></pre>
Open returns a new ReadSeeker reading the Plan 9 a.out section.




## <a id="SectionHeader">type</a> [SectionHeader](https://golang.org/src/debug/plan9obj/file.go?s=787:861#L26)
A SectionHeader represents a single Plan 9 a.out section header.
This structure doesn't exist on-disk, but eases navigation
through the object file.


<pre>type SectionHeader struct {
<span id="SectionHeader.Name"></span>    Name   <a href="/pkg/builtin/#string">string</a>
<span id="SectionHeader.Size"></span>    Size   <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader.Offset"></span>    Offset <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="Sym">type</a> [Sym](https://golang.org/src/debug/plan9obj/file.go?s=1703:1762#L60)
A Symbol represents an entry in a Plan 9 a.out symbol table section.


<pre>type Sym struct {
<span id="Sym.Value"></span>    Value <a href="/pkg/builtin/#uint64">uint64</a>
<span id="Sym.Type"></span>    Type  <a href="/pkg/builtin/#rune">rune</a>
<span id="Sym.Name"></span>    Name  <a href="/pkg/builtin/#string">string</a>
}
</pre>















