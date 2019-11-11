

# gosym
`import "debug/gosym"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package gosym implements access to the Go symbol
and line number tables embedded in Go binaries generated
by the gc compilers.




## <a id="pkg-index">Index</a>
* [type DecodingError](#DecodingError)
  * [func (e *DecodingError) Error() string](#DecodingError.Error)
* [type Func](#Func)
* [type LineTable](#LineTable)
  * [func NewLineTable(data []byte, text uint64) *LineTable](#NewLineTable)
  * [func (t *LineTable) LineToPC(line int, maxpc uint64) uint64](#LineTable.LineToPC)
  * [func (t *LineTable) PCToLine(pc uint64) int](#LineTable.PCToLine)
* [type Obj](#Obj)
* [type Sym](#Sym)
  * [func (s *Sym) BaseName() string](#Sym.BaseName)
  * [func (s *Sym) PackageName() string](#Sym.PackageName)
  * [func (s *Sym) ReceiverName() string](#Sym.ReceiverName)
  * [func (s *Sym) Static() bool](#Sym.Static)
* [type Table](#Table)
  * [func NewTable(symtab []byte, pcln *LineTable) (*Table, error)](#NewTable)
  * [func (t *Table) LineToPC(file string, line int) (pc uint64, fn *Func, err error)](#Table.LineToPC)
  * [func (t *Table) LookupFunc(name string) *Func](#Table.LookupFunc)
  * [func (t *Table) LookupSym(name string) *Sym](#Table.LookupSym)
  * [func (t *Table) PCToFunc(pc uint64) *Func](#Table.PCToFunc)
  * [func (t *Table) PCToLine(pc uint64) (file string, line int, fn *Func)](#Table.PCToLine)
  * [func (t *Table) SymByAddr(addr uint64) *Sym](#Table.SymByAddr)
* [type UnknownFileError](#UnknownFileError)
  * [func (e UnknownFileError) Error() string](#UnknownFileError.Error)
* [type UnknownLineError](#UnknownLineError)
  * [func (e *UnknownLineError) Error() string](#UnknownLineError.Error)




#### <a id="pkg-files">Package files</a>
[pclntab.go](https://golang.org/src/debug/gosym/pclntab.go) [symtab.go](https://golang.org/src/debug/gosym/symtab.go) 








## <a id="DecodingError">type</a> [DecodingError](https://golang.org/src/debug/gosym/symtab.go?s=16029:16096#L700)
DecodingError represents an error during the decoding of
the symbol table.


<pre>type DecodingError struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="DecodingError.Error">func</a> (\*DecodingError) [Error](https://golang.org/src/debug/gosym/symtab.go?s=16098:16136#L706)
<pre>func (e *<a href="#DecodingError">DecodingError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Func">type</a> [Func](https://golang.org/src/debug/gosym/symtab.go?s=1986:2207#L71)
A Func collects information about a single function.


<pre>type Func struct {
<span id="Func.Entry"></span>    Entry <a href="/pkg/builtin/#uint64">uint64</a>
    *<a href="#Sym">Sym</a>
<span id="Func.End"></span>    End       <a href="/pkg/builtin/#uint64">uint64</a>
<span id="Func.Params"></span>    Params    []*<a href="#Sym">Sym</a> <span class="comment">// nil for Go 1.3 and later binaries</span>
<span id="Func.Locals"></span>    Locals    []*<a href="#Sym">Sym</a> <span class="comment">// nil for Go 1.3 and later binaries</span>
<span id="Func.FrameSize"></span>    FrameSize <a href="/pkg/builtin/#int">int</a>
<span id="Func.LineTable"></span>    LineTable *<a href="#LineTable">LineTable</a>
<span id="Func.Obj"></span>    Obj       *<a href="#Obj">Obj</a>
}
</pre>











## <a id="LineTable">type</a> [LineTable](https://golang.org/src/debug/gosym/pclntab.go?s=1010:1410#L20)
A LineTable is a data structure mapping program counters to line numbers.

In Go 1.1 and earlier, each function (represented by a Func) had its own LineTable,
and the line number corresponded to a numbering of all source lines in the
program, across all files. That absolute line number would then have to be
converted separately to a file name and line number within the file.

In Go 1.2, the format of the data changed so that there is a single LineTable
for the entire program, shared by all Funcs, and there are no absolute line
numbers, just line numbers within specific files.

For the most part, LineTable's methods should be treated as an internal
detail of the package; callers should use the methods on Table instead.


<pre>type LineTable struct {
<span id="LineTable.Data"></span>    Data []<a href="/pkg/builtin/#byte">byte</a>
<span id="LineTable.PC"></span>    PC   <a href="/pkg/builtin/#uint64">uint64</a>
<span id="LineTable.Line"></span>    Line <a href="/pkg/builtin/#int">int</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewLineTable">func</a> [NewLineTable](https://golang.org/src/debug/gosym/pclntab.go?s=3569:3623#L116)
<pre>func NewLineTable(data []<a href="/pkg/builtin/#byte">byte</a>, text <a href="/pkg/builtin/#uint64">uint64</a>) *<a href="#LineTable">LineTable</a></pre>
NewLineTable returns a new PC/line table
corresponding to the encoded data.
Text must be the start address of the
corresponding text segment.






### <a id="LineTable.LineToPC">func</a> (\*LineTable) [LineToPC](https://golang.org/src/debug/gosym/pclntab.go?s=3159:3218#L100)
<pre>func (t *<a href="#LineTable">LineTable</a>) LineToPC(line <a href="/pkg/builtin/#int">int</a>, maxpc <a href="/pkg/builtin/#uint64">uint64</a>) <a href="/pkg/builtin/#uint64">uint64</a></pre>
LineToPC returns the program counter for the given line number,
considering only program counters before maxpc.

Deprecated: Use Table's LineToPC method instead.




### <a id="LineTable.PCToLine">func</a> (\*LineTable) [PCToLine](https://golang.org/src/debug/gosym/pclntab.go?s=2845:2888#L88)
<pre>func (t *<a href="#LineTable">LineTable</a>) PCToLine(pc <a href="/pkg/builtin/#uint64">uint64</a>) <a href="/pkg/builtin/#int">int</a></pre>
PCToLine returns the line number for the given program counter.

Deprecated: Use Table's PCToLine method instead.




## <a id="Obj">type</a> [Obj](https://golang.org/src/debug/gosym/symtab.go?s=2637:2952#L93)
An Obj represents a collection of functions in a symbol table.

The exact method of division of a binary into separate Objs is an internal detail
of the symbol table format.

In early versions of Go each source file became a different Obj.

In Go 1 and Go 1.1, each package produced one Obj for all Go sources
and one Obj per C source file.

In Go 1.2, there is a single Obj for the entire program.


<pre>type Obj struct {
<span id="Obj.Funcs"></span>    <span class="comment">// Funcs is a list of functions in the Obj.</span>
    Funcs []<a href="#Func">Func</a>

    <span class="comment">// In Go 1.1 and earlier, Paths is a list of symbols corresponding</span>
    <span class="comment">// to the source file names that produced the Obj.</span>
    <span class="comment">// In Go 1.2, Paths is nil.</span>
    <span class="comment">// Use the keys of Table.Files to obtain a list of source files.</span>
<span id="Obj.Paths"></span>    Paths []<a href="#Sym">Sym</a> <span class="comment">// meta</span>
}
</pre>











## <a id="Sym">type</a> [Sym](https://golang.org/src/debug/gosym/symtab.go?s=448:601#L13)
A Sym represents a single symbol table entry.


<pre>type Sym struct {
<span id="Sym.Value"></span>    Value  <a href="/pkg/builtin/#uint64">uint64</a>
<span id="Sym.Type"></span>    Type   <a href="/pkg/builtin/#byte">byte</a>
<span id="Sym.Name"></span>    Name   <a href="/pkg/builtin/#string">string</a>
<span id="Sym.GoType"></span>    GoType <a href="/pkg/builtin/#uint64">uint64</a>
    <span class="comment">// If this symbol is a function symbol, the corresponding Func</span>
<span id="Sym.Func"></span>    Func *<a href="#Func">Func</a>
}
</pre>











### <a id="Sym.BaseName">func</a> (\*Sym) [BaseName](https://golang.org/src/debug/gosym/symtab.go?s=1802:1833#L63)
<pre>func (s *<a href="#Sym">Sym</a>) BaseName() <a href="/pkg/builtin/#string">string</a></pre>
BaseName returns the symbol name without the package or receiver name.




### <a id="Sym.PackageName">func</a> (\*Sym) [PackageName](https://golang.org/src/debug/gosym/symtab.go?s=838:872#L27)
<pre>func (s *<a href="#Sym">Sym</a>) PackageName() <a href="/pkg/builtin/#string">string</a></pre>
PackageName returns the package part of the symbol name,
or the empty string if there is none.




### <a id="Sym.ReceiverName">func</a> (\*Sym) [ReceiverName](https://golang.org/src/debug/gosym/symtab.go?s=1429:1464#L49)
<pre>func (s *<a href="#Sym">Sym</a>) ReceiverName() <a href="/pkg/builtin/#string">string</a></pre>
ReceiverName returns the receiver type name of this symbol,
or the empty string if there is none.




### <a id="Sym.Static">func</a> (\*Sym) [Static](https://golang.org/src/debug/gosym/symtab.go?s=683:710#L23)
<pre>func (s *<a href="#Sym">Sym</a>) Static() <a href="/pkg/builtin/#bool">bool</a></pre>
Static reports whether this symbol is static (not visible outside its file).




## <a id="Table">type</a> [Table](https://golang.org/src/debug/gosym/symtab.go?s=3151:3406#L111)
Table represents a Go symbol table. It stores all of the
symbols decoded from the program and provides methods to translate
between symbols, names, and addresses.


<pre>type Table struct {
<span id="Table.Syms"></span>    Syms  []<a href="#Sym">Sym</a> <span class="comment">// nil for Go 1.3 and later binaries</span>
<span id="Table.Funcs"></span>    Funcs []<a href="#Func">Func</a>
<span id="Table.Files"></span>    Files map[<a href="/pkg/builtin/#string">string</a>]*<a href="#Obj">Obj</a> <span class="comment">// nil for Go 1.2 and later binaries</span>
<span id="Table.Objs"></span>    Objs  []<a href="#Obj">Obj</a>           <span class="comment">// nil for Go 1.2 and later binaries</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewTable">func</a> [NewTable](https://golang.org/src/debug/gosym/symtab.go?s=6911:6972#L276)
<pre>func NewTable(symtab []<a href="/pkg/builtin/#byte">byte</a>, pcln *<a href="#LineTable">LineTable</a>) (*<a href="#Table">Table</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewTable decodes the Go symbol table (the ".gosymtab" section in ELF),
returning an in-memory representation.
Starting with Go 1.3, the Go symbol table no longer includes symbol data.






### <a id="Table.LineToPC">func</a> (\*Table) [LineToPC](https://golang.org/src/debug/gosym/symtab.go?s=12054:12134#L511)
<pre>func (t *<a href="#Table">Table</a>) LineToPC(file <a href="/pkg/builtin/#string">string</a>, line <a href="/pkg/builtin/#int">int</a>) (pc <a href="/pkg/builtin/#uint64">uint64</a>, fn *<a href="#Func">Func</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
LineToPC looks up the first program counter on the given line in
the named file. It returns UnknownPathError or UnknownLineError if
there is an error looking up this line.




### <a id="Table.LookupFunc">func</a> (\*Table) [LookupFunc](https://golang.org/src/debug/gosym/symtab.go?s=13115:13160#L557)
<pre>func (t *<a href="#Table">Table</a>) LookupFunc(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Func">Func</a></pre>
LookupFunc returns the text, data, or bss symbol with the given name,
or nil if no such symbol is found.




### <a id="Table.LookupSym">func</a> (\*Table) [LookupSym](https://golang.org/src/debug/gosym/symtab.go?s=12753:12796#L541)
<pre>func (t *<a href="#Table">Table</a>) LookupSym(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Sym">Sym</a></pre>
LookupSym returns the text, data, or bss symbol with the given name,
or nil if no such symbol is found.




### <a id="Table.PCToFunc">func</a> (\*Table) [PCToFunc](https://golang.org/src/debug/gosym/symtab.go?s=11160:11201#L476)
<pre>func (t *<a href="#Table">Table</a>) PCToFunc(pc <a href="/pkg/builtin/#uint64">uint64</a>) *<a href="#Func">Func</a></pre>
PCToFunc returns the function containing the program counter pc,
or nil if there is no such function.




### <a id="Table.PCToLine">func</a> (\*Table) [PCToLine](https://golang.org/src/debug/gosym/symtab.go?s=11567:11636#L495)
<pre>func (t *<a href="#Table">Table</a>) PCToLine(pc <a href="/pkg/builtin/#uint64">uint64</a>) (file <a href="/pkg/builtin/#string">string</a>, line <a href="/pkg/builtin/#int">int</a>, fn *<a href="#Func">Func</a>)</pre>
PCToLine looks up line number information for a program counter.
If there is no information, it returns fn == nil.




### <a id="Table.SymByAddr">func</a> (\*Table) [SymByAddr](https://golang.org/src/debug/gosym/symtab.go?s=13350:13393#L568)
<pre>func (t *<a href="#Table">Table</a>) SymByAddr(addr <a href="/pkg/builtin/#uint64">uint64</a>) *<a href="#Sym">Sym</a></pre>
SymByAddr returns the text, data, or bss symbol starting at the given address.




## <a id="UnknownFileError">type</a> [UnknownFileError](https://golang.org/src/debug/gosym/symtab.go?s=15485:15513#L682)
UnknownFileError represents a failure to find the specific file in
the symbol table.


<pre>type UnknownFileError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="UnknownFileError.Error">func</a> (UnknownFileError) [Error](https://golang.org/src/debug/gosym/symtab.go?s=15515:15555#L684)
<pre>func (e <a href="#UnknownFileError">UnknownFileError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="UnknownLineError">type</a> [UnknownLineError](https://golang.org/src/debug/gosym/symtab.go?s=15784:15839#L689)
UnknownLineError represents a failure to map a line to a program
counter, either because the line is beyond the bounds of the file
or because there is no code on the given line.


<pre>type UnknownLineError struct {
<span id="UnknownLineError.File"></span>    File <a href="/pkg/builtin/#string">string</a>
<span id="UnknownLineError.Line"></span>    Line <a href="/pkg/builtin/#int">int</a>
}
</pre>











### <a id="UnknownLineError.Error">func</a> (\*UnknownLineError) [Error](https://golang.org/src/debug/gosym/symtab.go?s=15841:15882#L694)
<pre>func (e *<a href="#UnknownLineError">UnknownLineError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>






