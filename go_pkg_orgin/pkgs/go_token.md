

# token
`import "go/token"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package token defines constants representing the lexical tokens of the Go
programming language and basic operations on tokens (printing, predicates).


<a id="example__retrievePositionInfo">Example (RetrievePositionInfo)</a>
```go
```

output:
```txt
```


## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func IsExported(name string) bool](#IsExported)
* [func IsIdentifier(name string) bool](#IsIdentifier)
* [func IsKeyword(name string) bool](#IsKeyword)
* [type File](#File)
  * [func (f *File) AddLine(offset int)](#File.AddLine)
  * [func (f *File) AddLineColumnInfo(offset int, filename string, line, column int)](#File.AddLineColumnInfo)
  * [func (f *File) AddLineInfo(offset int, filename string, line int)](#File.AddLineInfo)
  * [func (f *File) Base() int](#File.Base)
  * [func (f *File) Line(p Pos) int](#File.Line)
  * [func (f *File) LineCount() int](#File.LineCount)
  * [func (f *File) LineStart(line int) Pos](#File.LineStart)
  * [func (f *File) MergeLine(line int)](#File.MergeLine)
  * [func (f *File) Name() string](#File.Name)
  * [func (f *File) Offset(p Pos) int](#File.Offset)
  * [func (f *File) Pos(offset int) Pos](#File.Pos)
  * [func (f *File) Position(p Pos) (pos Position)](#File.Position)
  * [func (f *File) PositionFor(p Pos, adjusted bool) (pos Position)](#File.PositionFor)
  * [func (f *File) SetLines(lines []int) bool](#File.SetLines)
  * [func (f *File) SetLinesForContent(content []byte)](#File.SetLinesForContent)
  * [func (f *File) Size() int](#File.Size)
* [type FileSet](#FileSet)
  * [func NewFileSet() *FileSet](#NewFileSet)
  * [func (s *FileSet) AddFile(filename string, base, size int) *File](#FileSet.AddFile)
  * [func (s *FileSet) Base() int](#FileSet.Base)
  * [func (s *FileSet) File(p Pos) (f *File)](#FileSet.File)
  * [func (s *FileSet) Iterate(f func(*File) bool)](#FileSet.Iterate)
  * [func (s *FileSet) Position(p Pos) (pos Position)](#FileSet.Position)
  * [func (s *FileSet) PositionFor(p Pos, adjusted bool) (pos Position)](#FileSet.PositionFor)
  * [func (s *FileSet) Read(decode func(interface{}) error) error](#FileSet.Read)
  * [func (s *FileSet) Write(encode func(interface{}) error) error](#FileSet.Write)
* [type Pos](#Pos)
  * [func (p Pos) IsValid() bool](#Pos.IsValid)
* [type Position](#Position)
  * [func (pos *Position) IsValid() bool](#Position.IsValid)
  * [func (pos Position) String() string](#Position.String)
* [type Token](#Token)
  * [func Lookup(ident string) Token](#Lookup)
  * [func (tok Token) IsKeyword() bool](#Token.IsKeyword)
  * [func (tok Token) IsLiteral() bool](#Token.IsLiteral)
  * [func (tok Token) IsOperator() bool](#Token.IsOperator)
  * [func (op Token) Precedence() int](#Token.Precedence)
  * [func (tok Token) String() string](#Token.String)


#### <a id="pkg-examples">Examples</a>
* [Package (RetrievePositionInfo)](#example__retrievePositionInfo)


#### <a id="pkg-files">Package files</a>
[position.go](https://golang.org/src/go/token/position.go) [serialize.go](https://golang.org/src/go/token/serialize.go) [token.go](https://golang.org/src/go/token/token.go) 


## <a id="pkg-constants">Constants</a>
A set of constants for precedence-based expression parsing.
Non-operators have lowest precedence, followed by operators
starting with precedence 1 up to unary operators. The highest
precedence serves as "catch-all" precedence for selector,
indexing, and other operator and delimiter tokens.


<pre>const (
    <span id="LowestPrec">LowestPrec</span>  = 0 <span class="comment">// non-operators</span>
    <span id="UnaryPrec">UnaryPrec</span>   = 6
    <span id="HighestPrec">HighestPrec</span> = 7
)</pre>



## <a id="IsExported">func</a> [IsExported](https://golang.org/src/go/token/token.go?s=5620:5653#L306)
<pre>func IsExported(name <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsExported reports whether name starts with an upper-case letter.



## <a id="IsIdentifier">func</a> [IsIdentifier](https://golang.org/src/go/token/token.go?s=6162:6197#L323)
<pre>func IsIdentifier(name <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsIdentifier reports whether name is a Go identifier, that is, a non-empty
string made up of letters, digits, and underscores, where the first character
is not a digit. Keywords are not identifiers.



## <a id="IsKeyword">func</a> [IsKeyword](https://golang.org/src/go/token/token.go?s=5809:5841#L313)
<pre>func IsKeyword(name <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsKeyword reports whether name is a Go keyword, such as "func" or "return".





## <a id="File">type</a> [File](https://golang.org/src/go/token/position.go?s=3120:3510#L86)
A File is a handle for a file belonging to a FileSet.
A File has a name, size, and line offset table.


<pre>type File struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="File.AddLine">func</a> (\*File) [AddLine](https://golang.org/src/go/token/position.go?s=4201:4235#L125)
<pre>func (f *<a href="#File">File</a>) AddLine(offset <a href="/pkg/builtin/#int">int</a>)</pre>
AddLine adds the line offset for a new line.
The line offset must be larger than the offset for the previous line
and smaller than the file size; otherwise the line offset is ignored.




### <a id="File.AddLineColumnInfo">func</a> (\*File) [AddLineColumnInfo](https://golang.org/src/go/token/position.go?s=7958:8037#L243)
<pre>func (f *<a href="#File">File</a>) AddLineColumnInfo(offset <a href="/pkg/builtin/#int">int</a>, filename <a href="/pkg/builtin/#string">string</a>, line, column <a href="/pkg/builtin/#int">int</a>)</pre>
AddLineColumnInfo adds alternative file, line, and column number
information for a given file offset. The offset must be larger
than the offset for the previously added alternative line info
and smaller than the file size; otherwise the information is
ignored.

AddLineColumnInfo is typically used to register alternative position
information for line directives such as //line filename:line:column.




### <a id="File.AddLineInfo">func</a> (\*File) [AddLineInfo](https://golang.org/src/go/token/position.go?s=7413:7478#L230)
<pre>func (f *<a href="#File">File</a>) AddLineInfo(offset <a href="/pkg/builtin/#int">int</a>, filename <a href="/pkg/builtin/#string">string</a>, line <a href="/pkg/builtin/#int">int</a>)</pre>
AddLineInfo is like AddLineColumnInfo with a column = 1 argument.
It is here for backward-compatibility for code prior to Go 1.11.




### <a id="File.Base">func</a> (\*File) [Base](https://golang.org/src/go/token/position.go?s=3699:3724#L104)
<pre>func (f *<a href="#File">File</a>) Base() <a href="/pkg/builtin/#int">int</a></pre>
Base returns the base offset of file f as registered with AddFile.




### <a id="File.Line">func</a> (\*File) [Line](https://golang.org/src/go/token/position.go?s=8882:8912#L276)
<pre>func (f *<a href="#File">File</a>) Line(p <a href="#Pos">Pos</a>) <a href="/pkg/builtin/#int">int</a></pre>
Line returns the line number for the given file position p;
p must be a Pos value in that file or NoPos.




### <a id="File.LineCount">func</a> (\*File) [LineCount](https://golang.org/src/go/token/position.go?s=3906:3936#L114)
<pre>func (f *<a href="#File">File</a>) LineCount() <a href="/pkg/builtin/#int">int</a></pre>
LineCount returns the number of lines in file f.




### <a id="File.LineStart">func</a> (\*File) [LineStart](https://golang.org/src/go/token/position.go?s=6717:6755#L205)
<pre>func (f *<a href="#File">File</a>) LineStart(line <a href="/pkg/builtin/#int">int</a>) <a href="#Pos">Pos</a></pre>
LineStart returns the Pos value of the start of the specified line.
It ignores any alternative positions set using AddLineColumnInfo.
LineStart panics if the 1-based line number is invalid.




### <a id="File.MergeLine">func</a> (\*File) [MergeLine](https://golang.org/src/go/token/position.go?s=4686:4720#L138)
<pre>func (f *<a href="#File">File</a>) MergeLine(line <a href="/pkg/builtin/#int">int</a>)</pre>
MergeLine merges a line with the following line. It is akin to replacing
the newline character at the end of the line with a space (to not change the
remaining offsets). To obtain the line number, consult e.g. Position.Line.
MergeLine will panic if given an invalid line number.




### <a id="File.Name">func</a> (\*File) [Name](https://golang.org/src/go/token/position.go?s=3580:3608#L99)
<pre>func (f *<a href="#File">File</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the file name of file f as registered with AddFile.




### <a id="File.Offset">func</a> (\*File) [Offset](https://golang.org/src/go/token/position.go?s=8626:8658#L266)
<pre>func (f *<a href="#File">File</a>) Offset(p <a href="#Pos">Pos</a>) <a href="/pkg/builtin/#int">int</a></pre>
Offset returns the offset for the given file position p;
p must be a valid Pos value in that file.
f.Offset(f.Pos(offset)) == offset.




### <a id="File.Pos">func</a> (\*File) [Pos](https://golang.org/src/go/token/position.go?s=8355:8389#L255)
<pre>func (f *<a href="#File">File</a>) Pos(offset <a href="/pkg/builtin/#int">int</a>) <a href="#Pos">Pos</a></pre>
Pos returns the Pos value for the given file offset;
the offset must be <= f.Size().
f.Pos(f.Offset(p)) == p.




### <a id="File.Position">func</a> (\*File) [Position](https://golang.org/src/go/token/position.go?s=11220:11265#L346)
<pre>func (f *<a href="#File">File</a>) Position(p <a href="#Pos">Pos</a>) (pos <a href="#Position">Position</a>)</pre>
Position returns the Position value for the given file position p.
Calling f.Position(p) is equivalent to calling f.PositionFor(p, true).




### <a id="File.PositionFor">func</a> (\*File) [PositionFor](https://golang.org/src/go/token/position.go?s=10861:10924#L333)
<pre>func (f *<a href="#File">File</a>) PositionFor(p <a href="#Pos">Pos</a>, adjusted <a href="/pkg/builtin/#bool">bool</a>) (pos <a href="#Position">Position</a>)</pre>
PositionFor returns the Position value for the given file position p.
If adjusted is set, the position may be adjusted by position-altering
//line comments; otherwise those comments are ignored.
p must be a Pos value in f or NoPos.




### <a id="File.SetLines">func</a> (\*File) [SetLines](https://golang.org/src/go/token/position.go?s=5806:5847#L165)
<pre>func (f *<a href="#File">File</a>) SetLines(lines []<a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
SetLines sets the line offsets for a file and reports whether it succeeded.
The line offsets are the offsets of the first character of each line;
for instance for the content "ab\nc\n" the line offsets are {0, 3}.
An empty file has an empty line offset table.
Each line offset must be larger than the offset for the previous line
and smaller than the file size; otherwise SetLines fails and returns
false.
Callers must not mutate the provided slice after SetLines returns.




### <a id="File.SetLinesForContent">func</a> (\*File) [SetLinesForContent](https://golang.org/src/go/token/position.go?s=6220:6269#L183)
<pre>func (f *<a href="#File">File</a>) SetLinesForContent(content []<a href="/pkg/builtin/#byte">byte</a>)</pre>
SetLinesForContent sets the line offsets for the given file content.
It ignores position-altering //line comments.




### <a id="File.Size">func</a> (\*File) [Size](https://golang.org/src/go/token/position.go?s=3808:3833#L109)
<pre>func (f *<a href="#File">File</a>) Size() <a href="/pkg/builtin/#int">int</a></pre>
Size returns the size of file f as registered with AddFile.




## <a id="FileSet">type</a> [FileSet](https://golang.org/src/go/token/position.go?s=11540:11780#L357)
A FileSet represents a set of source files.
Methods of file sets are synchronized; multiple goroutines
may invoke them concurrently.


<pre>type FileSet struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewFileSet">func</a> [NewFileSet](https://golang.org/src/go/token/position.go?s=11820:11846#L365)
<pre>func NewFileSet() *<a href="#FileSet">FileSet</a></pre>
NewFileSet creates a new file set.






### <a id="FileSet.AddFile">func</a> (\*FileSet) [AddFile](https://golang.org/src/go/token/position.go?s=12901:12965#L398)
<pre>func (s *<a href="#FileSet">FileSet</a>) AddFile(filename <a href="/pkg/builtin/#string">string</a>, base, size <a href="/pkg/builtin/#int">int</a>) *<a href="#File">File</a></pre>
AddFile adds a new file with a given filename, base offset, and file size
to the file set s and returns the file. Multiple files may have the same
name. The base offset must not be smaller than the FileSet's Base(), and
size must not be negative. As a special case, if a negative base is provided,
the current value of the FileSet's Base() is used instead.

Adding the file will set the file set's Base() value to base + size + 1
as the minimum base value for the next file. The following relationship
exists between a Pos value p for a given file offset offs:


	int(p) = base + offs

with offs in the range [0, size] and thus p in the range [base, base+size].
For convenience, File.Pos may be used to create file-specific position
values from a file offset.




### <a id="FileSet.Base">func</a> (\*FileSet) [Base](https://golang.org/src/go/token/position.go?s=12004:12032#L374)
<pre>func (s *<a href="#FileSet">FileSet</a>) Base() <a href="/pkg/builtin/#int">int</a></pre>
Base returns the minimum base offset that must be provided to
AddFile when adding the next file.




### <a id="FileSet.File">func</a> (\*FileSet) [File](https://golang.org/src/go/token/position.go?s=14611:14650#L468)
<pre>func (s *<a href="#FileSet">FileSet</a>) File(p <a href="#Pos">Pos</a>) (f *<a href="#File">File</a>)</pre>
File returns the file that contains the position p.
If no such file is found (for instance for p == NoPos),
the result is nil.




### <a id="FileSet.Iterate">func</a> (\*FileSet) [Iterate](https://golang.org/src/go/token/position.go?s=13573:13618#L423)
<pre>func (s *<a href="#FileSet">FileSet</a>) Iterate(f func(*<a href="#File">File</a>) <a href="/pkg/builtin/#bool">bool</a>)</pre>
Iterate calls f for the files in the file set in the order they were added
until f returns false.




### <a id="FileSet.Position">func</a> (\*FileSet) [Position](https://golang.org/src/go/token/position.go?s=15258:15306#L492)
<pre>func (s *<a href="#FileSet">FileSet</a>) Position(p <a href="#Pos">Pos</a>) (pos <a href="#Position">Position</a>)</pre>
Position converts a Pos p in the fileset into a Position value.
Calling s.Position(p) is equivalent to calling s.PositionFor(p, true).




### <a id="FileSet.PositionFor">func</a> (\*FileSet) [PositionFor](https://golang.org/src/go/token/position.go?s=14944:15010#L480)
<pre>func (s *<a href="#FileSet">FileSet</a>) PositionFor(p <a href="#Pos">Pos</a>, adjusted <a href="/pkg/builtin/#bool">bool</a>) (pos <a href="#Position">Position</a>)</pre>
PositionFor converts a Pos p in the fileset into a Position value.
If adjusted is set, the position may be adjusted by position-altering
//line comments; otherwise those comments are ignored.
p must be a Pos value in s or NoPos.




### <a id="FileSet.Read">func</a> (\*FileSet) [Read](https://golang.org/src/go/token/serialize.go?s=490:550#L12)
<pre>func (s *<a href="#FileSet">FileSet</a>) Read(decode func(interface{}) <a href="/pkg/builtin/#error">error</a>) <a href="/pkg/builtin/#error">error</a></pre>
Read calls decode to deserialize a file set into s; s must not be nil.




### <a id="FileSet.Write">func</a> (\*FileSet) [Write](https://golang.org/src/go/token/serialize.go?s=1012:1073#L40)
<pre>func (s *<a href="#FileSet">FileSet</a>) Write(encode func(interface{}) <a href="/pkg/builtin/#error">error</a>) <a href="/pkg/builtin/#error">error</a></pre>
Write calls encode to serialize the file set s.




## <a id="Pos">type</a> [Pos](https://golang.org/src/go/token/position.go?s=2521:2533#L66)
Pos is a compact encoding of a source position within a file set.
It can be converted into a Position for a more convenient, but much
larger, representation.

The Pos value for a given file is a number in the range [base, base+size],
where base and size are specified when adding the file to the file set via
AddFile.

To create the Pos value for a specific source offset (measured in bytes),
first add the respective file to the current file set using FileSet.AddFile
and then call File.Pos(offset) for that file. Given a Pos value p
for a specific file set fset, the corresponding Position value is
obtained by calling fset.Position(p).

Pos values can be compared directly with the usual comparison operators:
If two Pos values p and q are in the same file, comparing p and q is
equivalent to comparing the respective source file offsets. If p and q
are in different files, p < q is true if the file implied by p was added
to the respective file set before the file implied by q.


<pre>type Pos <a href="/pkg/builtin/#int">int</a></pre>


The zero value for Pos is NoPos; there is no file and line information
associated with it, and NoPos.IsValid() is false. NoPos is always
smaller than any other Pos value. The corresponding Position value
for NoPos is the zero value for Position.


<pre>const <span id="NoPos">NoPos</span> <a href="#Pos">Pos</a> = 0</pre>









### <a id="Pos.IsValid">func</a> (Pos) [IsValid](https://golang.org/src/go/token/position.go?s=2867:2894#L76)
<pre>func (p <a href="#Pos">Pos</a>) IsValid() <a href="/pkg/builtin/#bool">bool</a></pre>
IsValid reports whether the position is valid.




## <a id="Position">type</a> [Position](https://golang.org/src/go/token/position.go?s=459:671#L10)
Position describes an arbitrary source position
including the file, line, and column location.
A Position is valid if the line number is > 0.


<pre>type Position struct {
<span id="Position.Filename"></span>    Filename <a href="/pkg/builtin/#string">string</a> <span class="comment">// filename, if any</span>
<span id="Position.Offset"></span>    Offset   <a href="/pkg/builtin/#int">int</a>    <span class="comment">// offset, starting at 0</span>
<span id="Position.Line"></span>    Line     <a href="/pkg/builtin/#int">int</a>    <span class="comment">// line number, starting at 1</span>
<span id="Position.Column"></span>    Column   <a href="/pkg/builtin/#int">int</a>    <span class="comment">// column number, starting at 1 (byte count)</span>
}
</pre>











### <a id="Position.IsValid">func</a> (\*Position) [IsValid](https://golang.org/src/go/token/position.go?s=723:758#L18)
<pre>func (pos *<a href="#Position">Position</a>) IsValid() <a href="/pkg/builtin/#bool">bool</a></pre>
IsValid reports whether the position is valid.




### <a id="Position.String">func</a> (Position) [String](https://golang.org/src/go/token/position.go?s=1229:1264#L29)
<pre>func (pos <a href="#Position">Position</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns a string in one of several forms:


	file:line:column    valid position with file name
	file:line           valid position with file name but no column (column == 0)
	line:column         valid position without file name
	line                valid position without file name and no column (column == 0)
	file                invalid position with file name
	-                   invalid position without file name




## <a id="Token">type</a> [Token](https://golang.org/src/go/token/token.go?s=454:468#L7)
Token is the set of lexical tokens of the Go programming language.


<pre>type Token <a href="/pkg/builtin/#int">int</a></pre>


The list of tokens.


<pre>const (
    <span class="comment">// Special tokens</span>
    <span id="ILLEGAL">ILLEGAL</span> <a href="#Token">Token</a> = <a href="/pkg/builtin/#iota">iota</a>
    <span id="EOF">EOF</span>
    <span id="COMMENT">COMMENT</span>

    <span class="comment">// Identifiers and basic type literals</span>
    <span class="comment">// (these tokens stand for classes of literals)</span>
    <span id="IDENT">IDENT</span>  <span class="comment">// main</span>
    <span id="INT">INT</span>    <span class="comment">// 12345</span>
    <span id="FLOAT">FLOAT</span>  <span class="comment">// 123.45</span>
    <span id="IMAG">IMAG</span>   <span class="comment">// 123.45i</span>
    <span id="CHAR">CHAR</span>   <span class="comment">// &#39;a&#39;</span>
    <span id="STRING">STRING</span> <span class="comment">// &#34;abc&#34;</span>

    <span class="comment">// Operators and delimiters</span>
    <span id="ADD">ADD</span> <span class="comment">// +</span>
    <span id="SUB">SUB</span> <span class="comment">// -</span>
    <span id="MUL">MUL</span> <span class="comment">// *</span>
    <span id="QUO">QUO</span> <span class="comment">// /</span>
    <span id="REM">REM</span> <span class="comment">// %</span>

    <span id="AND">AND</span>     <span class="comment">// &amp;</span>
    <span id="OR">OR</span>      <span class="comment">// |</span>
    <span id="XOR">XOR</span>     <span class="comment">// ^</span>
    <span id="SHL">SHL</span>     <span class="comment">// &lt;&lt;</span>
    <span id="SHR">SHR</span>     <span class="comment">// &gt;&gt;</span>
    <span id="AND_NOT">AND_NOT</span> <span class="comment">// &amp;^</span>

    <span id="ADD_ASSIGN">ADD_ASSIGN</span> <span class="comment">// +=</span>
    <span id="SUB_ASSIGN">SUB_ASSIGN</span> <span class="comment">// -=</span>
    <span id="MUL_ASSIGN">MUL_ASSIGN</span> <span class="comment">// *=</span>
    <span id="QUO_ASSIGN">QUO_ASSIGN</span> <span class="comment">// /=</span>
    <span id="REM_ASSIGN">REM_ASSIGN</span> <span class="comment">// %=</span>

    <span id="AND_ASSIGN">AND_ASSIGN</span>     <span class="comment">// &amp;=</span>
    <span id="OR_ASSIGN">OR_ASSIGN</span>      <span class="comment">// |=</span>
    <span id="XOR_ASSIGN">XOR_ASSIGN</span>     <span class="comment">// ^=</span>
    <span id="SHL_ASSIGN">SHL_ASSIGN</span>     <span class="comment">// &lt;&lt;=</span>
    <span id="SHR_ASSIGN">SHR_ASSIGN</span>     <span class="comment">// &gt;&gt;=</span>
    <span id="AND_NOT_ASSIGN">AND_NOT_ASSIGN</span> <span class="comment">// &amp;^=</span>

    <span id="LAND">LAND</span>  <span class="comment">// &amp;&amp;</span>
    <span id="LOR">LOR</span>   <span class="comment">// ||</span>
    <span id="ARROW">ARROW</span> <span class="comment">// &lt;-</span>
    <span id="INC">INC</span>   <span class="comment">// ++</span>
    <span id="DEC">DEC</span>   <span class="comment">// --</span>

    <span id="EQL">EQL</span>    <span class="comment">// ==</span>
    <span id="LSS">LSS</span>    <span class="comment">// &lt;</span>
    <span id="GTR">GTR</span>    <span class="comment">// &gt;</span>
    <span id="ASSIGN">ASSIGN</span> <span class="comment">// =</span>
    <span id="NOT">NOT</span>    <span class="comment">// !</span>

    <span id="NEQ">NEQ</span>      <span class="comment">// !=</span>
    <span id="LEQ">LEQ</span>      <span class="comment">// &lt;=</span>
    <span id="GEQ">GEQ</span>      <span class="comment">// &gt;=</span>
    <span id="DEFINE">DEFINE</span>   <span class="comment">// :=</span>
    <span id="ELLIPSIS">ELLIPSIS</span> <span class="comment">// ...</span>

    <span id="LPAREN">LPAREN</span> <span class="comment">// (</span>
    <span id="LBRACK">LBRACK</span> <span class="comment">// [</span>
    <span id="LBRACE">LBRACE</span> <span class="comment">// {</span>
    <span id="COMMA">COMMA</span>  <span class="comment">// ,</span>
    <span id="PERIOD">PERIOD</span> <span class="comment">// .</span>

    <span id="RPAREN">RPAREN</span>    <span class="comment">// )</span>
    <span id="RBRACK">RBRACK</span>    <span class="comment">// ]</span>
    <span id="RBRACE">RBRACE</span>    <span class="comment">// }</span>
    <span id="SEMICOLON">SEMICOLON</span> <span class="comment">// ;</span>
    <span id="COLON">COLON</span>     <span class="comment">// :</span>

    <span class="comment">// Keywords</span>
    <span id="BREAK">BREAK</span>
    <span id="CASE">CASE</span>
    <span id="CHAN">CHAN</span>
    <span id="CONST">CONST</span>
    <span id="CONTINUE">CONTINUE</span>

    <span id="DEFAULT">DEFAULT</span>
    <span id="DEFER">DEFER</span>
    <span id="ELSE">ELSE</span>
    <span id="FALLTHROUGH">FALLTHROUGH</span>
    <span id="FOR">FOR</span>

    <span id="FUNC">FUNC</span>
    <span id="GO">GO</span>
    <span id="GOTO">GOTO</span>
    <span id="IF">IF</span>
    <span id="IMPORT">IMPORT</span>

    <span id="INTERFACE">INTERFACE</span>
    <span id="MAP">MAP</span>
    <span id="PACKAGE">PACKAGE</span>
    <span id="RANGE">RANGE</span>
    <span id="RETURN">RETURN</span>

    <span id="SELECT">SELECT</span>
    <span id="STRUCT">STRUCT</span>
    <span id="SWITCH">SWITCH</span>
    <span id="TYPE">TYPE</span>
    <span id="VAR">VAR</span>
)</pre>







### <a id="Lookup">func</a> [Lookup](https://golang.org/src/go/token/token.go?s=4817:4848#L280)
<pre>func Lookup(ident <a href="/pkg/builtin/#string">string</a>) <a href="#Token">Token</a></pre>
Lookup maps an identifier to its keyword token or IDENT (if not a keyword).






### <a id="Token.IsKeyword">func</a> (Token) [IsKeyword](https://golang.org/src/go/token/token.go?s=5463:5496#L302)
<pre>func (tok <a href="#Token">Token</a>) IsKeyword() <a href="/pkg/builtin/#bool">bool</a></pre>
IsKeyword returns true for tokens corresponding to keywords;
it returns false otherwise.




### <a id="Token.IsLiteral">func</a> (Token) [IsLiteral](https://golang.org/src/go/token/token.go?s=5077:5110#L292)
<pre>func (tok <a href="#Token">Token</a>) IsLiteral() <a href="/pkg/builtin/#bool">bool</a></pre>
IsLiteral returns true for tokens corresponding to identifiers
and basic type literals; it returns false otherwise.




### <a id="Token.IsOperator">func</a> (Token) [IsOperator](https://golang.org/src/go/token/token.go?s=5277:5311#L297)
<pre>func (tok <a href="#Token">Token</a>) IsOperator() <a href="/pkg/builtin/#bool">bool</a></pre>
IsOperator returns true for tokens corresponding to operators and
delimiters; it returns false otherwise.




### <a id="Token.Precedence">func</a> (Token) [Precedence](https://golang.org/src/go/token/token.go?s=4316:4348#L253)
<pre>func (op <a href="#Token">Token</a>) Precedence() <a href="/pkg/builtin/#int">int</a></pre>
Precedence returns the operator precedence of the binary
operator op. If op is not a binary operator, the result
is LowestPrecedence.




### <a id="Token.String">func</a> (Token) [String](https://golang.org/src/go/token/token.go?s=3598:3630#L226)
<pre>func (tok <a href="#Token">Token</a>) String() <a href="/pkg/builtin/#string">string</a></pre>
String returns the string corresponding to the token tok.
For operators, delimiters, and keywords the string is the actual
token character sequence (e.g., for the token ADD, the string is
"+"). For all other tokens the string corresponds to the token
constant name (e.g. for the token IDENT, the string is "IDENT").







