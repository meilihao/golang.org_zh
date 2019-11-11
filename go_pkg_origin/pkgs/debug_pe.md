

# pe
`import "debug/pe"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package pe implements access to PE (Microsoft Windows Portable Executable) files.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [type COFFSymbol](#COFFSymbol)
  * [func (sym *COFFSymbol) FullName(st StringTable) (string, error)](#COFFSymbol.FullName)
* [type DataDirectory](#DataDirectory)
* [type File](#File)
  * [func NewFile(r io.ReaderAt) (*File, error)](#NewFile)
  * [func Open(name string) (*File, error)](#Open)
  * [func (f *File) Close() error](#File.Close)
  * [func (f *File) DWARF() (*dwarf.Data, error)](#File.DWARF)
  * [func (f *File) ImportedLibraries() ([]string, error)](#File.ImportedLibraries)
  * [func (f *File) ImportedSymbols() ([]string, error)](#File.ImportedSymbols)
  * [func (f *File) Section(name string) *Section](#File.Section)
* [type FileHeader](#FileHeader)
* [type FormatError](#FormatError)
  * [func (e *FormatError) Error() string](#FormatError.Error)
* [type ImportDirectory](#ImportDirectory)
* [type OptionalHeader32](#OptionalHeader32)
* [type OptionalHeader64](#OptionalHeader64)
* [type Reloc](#Reloc)
* [type Section](#Section)
  * [func (s *Section) Data() ([]byte, error)](#Section.Data)
  * [func (s *Section) Open() io.ReadSeeker](#Section.Open)
* [type SectionHeader](#SectionHeader)
* [type SectionHeader32](#SectionHeader32)
* [type StringTable](#StringTable)
  * [func (st StringTable) String(start uint32) (string, error)](#StringTable.String)
* [type Symbol](#Symbol)




#### <a id="pkg-files">Package files</a>
[file.go](https://golang.org/src/debug/pe/file.go) [pe.go](https://golang.org/src/debug/pe/pe.go) [section.go](https://golang.org/src/debug/pe/section.go) [string.go](https://golang.org/src/debug/pe/string.go) [symbol.go](https://golang.org/src/debug/pe/symbol.go) 


## <a id="pkg-constants">Constants</a>

<pre>const (
    <span id="IMAGE_FILE_MACHINE_UNKNOWN">IMAGE_FILE_MACHINE_UNKNOWN</span>   = 0x0
    <span id="IMAGE_FILE_MACHINE_AM33">IMAGE_FILE_MACHINE_AM33</span>      = 0x1d3
    <span id="IMAGE_FILE_MACHINE_AMD64">IMAGE_FILE_MACHINE_AMD64</span>     = 0x8664
    <span id="IMAGE_FILE_MACHINE_ARM">IMAGE_FILE_MACHINE_ARM</span>       = 0x1c0
    <span id="IMAGE_FILE_MACHINE_ARMNT">IMAGE_FILE_MACHINE_ARMNT</span>     = 0x1c4
    <span id="IMAGE_FILE_MACHINE_ARM64">IMAGE_FILE_MACHINE_ARM64</span>     = 0xaa64
    <span id="IMAGE_FILE_MACHINE_EBC">IMAGE_FILE_MACHINE_EBC</span>       = 0xebc
    <span id="IMAGE_FILE_MACHINE_I386">IMAGE_FILE_MACHINE_I386</span>      = 0x14c
    <span id="IMAGE_FILE_MACHINE_IA64">IMAGE_FILE_MACHINE_IA64</span>      = 0x200
    <span id="IMAGE_FILE_MACHINE_M32R">IMAGE_FILE_MACHINE_M32R</span>      = 0x9041
    <span id="IMAGE_FILE_MACHINE_MIPS16">IMAGE_FILE_MACHINE_MIPS16</span>    = 0x266
    <span id="IMAGE_FILE_MACHINE_MIPSFPU">IMAGE_FILE_MACHINE_MIPSFPU</span>   = 0x366
    <span id="IMAGE_FILE_MACHINE_MIPSFPU16">IMAGE_FILE_MACHINE_MIPSFPU16</span> = 0x466
    <span id="IMAGE_FILE_MACHINE_POWERPC">IMAGE_FILE_MACHINE_POWERPC</span>   = 0x1f0
    <span id="IMAGE_FILE_MACHINE_POWERPCFP">IMAGE_FILE_MACHINE_POWERPCFP</span> = 0x1f1
    <span id="IMAGE_FILE_MACHINE_R4000">IMAGE_FILE_MACHINE_R4000</span>     = 0x166
    <span id="IMAGE_FILE_MACHINE_SH3">IMAGE_FILE_MACHINE_SH3</span>       = 0x1a2
    <span id="IMAGE_FILE_MACHINE_SH3DSP">IMAGE_FILE_MACHINE_SH3DSP</span>    = 0x1a3
    <span id="IMAGE_FILE_MACHINE_SH4">IMAGE_FILE_MACHINE_SH4</span>       = 0x1a6
    <span id="IMAGE_FILE_MACHINE_SH5">IMAGE_FILE_MACHINE_SH5</span>       = 0x1a8
    <span id="IMAGE_FILE_MACHINE_THUMB">IMAGE_FILE_MACHINE_THUMB</span>     = 0x1c2
    <span id="IMAGE_FILE_MACHINE_WCEMIPSV2">IMAGE_FILE_MACHINE_WCEMIPSV2</span> = 0x169
)</pre>IMAGE_DIRECTORY_ENTRY constants


<pre>const (
    <span id="IMAGE_DIRECTORY_ENTRY_EXPORT">IMAGE_DIRECTORY_ENTRY_EXPORT</span>         = 0
    <span id="IMAGE_DIRECTORY_ENTRY_IMPORT">IMAGE_DIRECTORY_ENTRY_IMPORT</span>         = 1
    <span id="IMAGE_DIRECTORY_ENTRY_RESOURCE">IMAGE_DIRECTORY_ENTRY_RESOURCE</span>       = 2
    <span id="IMAGE_DIRECTORY_ENTRY_EXCEPTION">IMAGE_DIRECTORY_ENTRY_EXCEPTION</span>      = 3
    <span id="IMAGE_DIRECTORY_ENTRY_SECURITY">IMAGE_DIRECTORY_ENTRY_SECURITY</span>       = 4
    <span id="IMAGE_DIRECTORY_ENTRY_BASERELOC">IMAGE_DIRECTORY_ENTRY_BASERELOC</span>      = 5
    <span id="IMAGE_DIRECTORY_ENTRY_DEBUG">IMAGE_DIRECTORY_ENTRY_DEBUG</span>          = 6
    <span id="IMAGE_DIRECTORY_ENTRY_ARCHITECTURE">IMAGE_DIRECTORY_ENTRY_ARCHITECTURE</span>   = 7
    <span id="IMAGE_DIRECTORY_ENTRY_GLOBALPTR">IMAGE_DIRECTORY_ENTRY_GLOBALPTR</span>      = 8
    <span id="IMAGE_DIRECTORY_ENTRY_TLS">IMAGE_DIRECTORY_ENTRY_TLS</span>            = 9
    <span id="IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG">IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG</span>    = 10
    <span id="IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT">IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT</span>   = 11
    <span id="IMAGE_DIRECTORY_ENTRY_IAT">IMAGE_DIRECTORY_ENTRY_IAT</span>            = 12
    <span id="IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT">IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT</span>   = 13
    <span id="IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR">IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR</span> = 14
)</pre>
<pre>const <span id="COFFSymbolSize">COFFSymbolSize</span> = 18</pre>





## <a id="COFFSymbol">type</a> [COFFSymbol](https://golang.org/src/debug/pe/symbol.go?s=301:488#L6)
COFFSymbol represents single COFF symbol table record.


<pre>type COFFSymbol struct {
<span id="COFFSymbol.Name"></span>    Name               [8]<a href="/pkg/builtin/#uint8">uint8</a>
<span id="COFFSymbol.Value"></span>    Value              <a href="/pkg/builtin/#uint32">uint32</a>
<span id="COFFSymbol.SectionNumber"></span>    SectionNumber      <a href="/pkg/builtin/#int16">int16</a>
<span id="COFFSymbol.Type"></span>    Type               <a href="/pkg/builtin/#uint16">uint16</a>
<span id="COFFSymbol.StorageClass"></span>    StorageClass       <a href="/pkg/builtin/#uint8">uint8</a>
<span id="COFFSymbol.NumberOfAuxSymbols"></span>    NumberOfAuxSymbols <a href="/pkg/builtin/#uint8">uint8</a>
}
</pre>











### <a id="COFFSymbol.FullName">func</a> (\*COFFSymbol) [FullName](https://golang.org/src/debug/pe/symbol.go?s=1472:1535#L45)
<pre>func (sym *<a href="#COFFSymbol">COFFSymbol</a>) FullName(st <a href="#StringTable">StringTable</a>) (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
FullName finds real name of symbol sym. Normally name is stored
in sym.Name, but if it is longer then 8 characters, it is stored
in COFF string table st instead.




## <a id="DataDirectory">type</a> [DataDirectory](https://golang.org/src/debug/pe/pe.go?s=403:478#L7)

<pre>type DataDirectory struct {
<span id="DataDirectory.VirtualAddress"></span>    VirtualAddress <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DataDirectory.Size"></span>    Size           <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="File">type</a> [File](https://golang.org/src/debug/pe/file.go?s=497:850#L13)
A File represents an open PE file.


<pre>type File struct {
    <a href="#FileHeader">FileHeader</a>
<span id="File.OptionalHeader"></span>    OptionalHeader interface{} <span class="comment">// of type *OptionalHeader32 or *OptionalHeader64</span>
<span id="File.Sections"></span>    Sections       []*<a href="#Section">Section</a>
<span id="File.Symbols"></span>    Symbols        []*<a href="#Symbol">Symbol</a>    <span class="comment">// COFF symbols with auxiliary symbol records removed</span>
<span id="File.COFFSymbols"></span>    COFFSymbols    []<a href="#COFFSymbol">COFFSymbol</a> <span class="comment">// all COFF symbols (including auxiliary symbol records)</span>
<span id="File.StringTable"></span>    StringTable    <a href="#StringTable">StringTable</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewFile">func</a> [NewFile](https://golang.org/src/debug/pe/file.go?s=1735:1777#L59)
<pre>func NewFile(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReaderAt">ReaderAt</a>) (*<a href="#File">File</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewFile creates a new File for accessing a PE binary in an underlying reader.




### <a id="Open">func</a> [Open](https://golang.org/src/debug/pe/file.go?s=935:972#L25)
<pre>func Open(name <a href="/pkg/builtin/#string">string</a>) (*<a href="#File">File</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Open opens the named file using os.Open and prepares it for use as a PE binary.






### <a id="File.Close">func</a> (\*File) [Close](https://golang.org/src/debug/pe/file.go?s=1262:1290#L42)
<pre>func (f *<a href="#File">File</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the File.
If the File was created using NewFile directly instead of Open,
Close has no effect.




### <a id="File.DWARF">func</a> (\*File) [DWARF](https://golang.org/src/debug/pe/file.go?s=5927:5970#L212)
<pre>func (f *<a href="#File">File</a>) DWARF() (*<a href="/pkg/debug/dwarf/">dwarf</a>.<a href="/pkg/debug/dwarf/#Data">Data</a>, <a href="/pkg/builtin/#error">error</a>)</pre>



### <a id="File.ImportedLibraries">func</a> (\*File) [ImportedLibraries](https://golang.org/src/debug/pe/file.go?s=11647:11699#L432)
<pre>func (f *<a href="#File">File</a>) ImportedLibraries() ([]<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ImportedLibraries returns the names of all libraries
referred to by the binary f that are expected to be
linked with the binary at dynamic link time.




### <a id="File.ImportedSymbols">func</a> (\*File) [ImportedSymbols](https://golang.org/src/debug/pe/file.go?s=8253:8303#L316)
<pre>func (f *<a href="#File">File</a>) ImportedSymbols() ([]<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ImportedSymbols returns the names of all symbols
referred to by the binary f that are expected to be
satisfied by other libraries at dynamic load time.
It does not return weak symbols.




### <a id="File.Section">func</a> (\*File) [Section](https://golang.org/src/debug/pe/file.go?s=5792:5836#L203)
<pre>func (f *<a href="#File">File</a>) Section(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Section">Section</a></pre>
Section returns the first section with the given name, or nil if no such
section exists.




## <a id="FileHeader">type</a> [FileHeader](https://golang.org/src/debug/pe/pe.go?s=172:401#L1)

<pre>type FileHeader struct {
<span id="FileHeader.Machine"></span>    Machine              <a href="/pkg/builtin/#uint16">uint16</a>
<span id="FileHeader.NumberOfSections"></span>    NumberOfSections     <a href="/pkg/builtin/#uint16">uint16</a>
<span id="FileHeader.TimeDateStamp"></span>    TimeDateStamp        <a href="/pkg/builtin/#uint32">uint32</a>
<span id="FileHeader.PointerToSymbolTable"></span>    PointerToSymbolTable <a href="/pkg/builtin/#uint32">uint32</a>
<span id="FileHeader.NumberOfSymbols"></span>    NumberOfSymbols      <a href="/pkg/builtin/#uint32">uint32</a>
<span id="FileHeader.SizeOfOptionalHeader"></span>    SizeOfOptionalHeader <a href="/pkg/builtin/#uint16">uint16</a>
<span id="FileHeader.Characteristics"></span>    Characteristics      <a href="/pkg/builtin/#uint16">uint16</a>
}
</pre>











## <a id="FormatError">type</a> [FormatError](https://golang.org/src/debug/pe/file.go?s=11866:11893#L440)
FormatError is unused.
The type is retained for compatibility.


<pre>type FormatError struct {
}
</pre>











### <a id="FormatError.Error">func</a> (\*FormatError) [Error](https://golang.org/src/debug/pe/file.go?s=11895:11931#L443)
<pre>func (e *<a href="#FormatError">FormatError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="ImportDirectory">type</a> [ImportDirectory](https://golang.org/src/debug/pe/file.go?s=7875:8054#L302)

<pre>type ImportDirectory struct {
<span id="ImportDirectory.OriginalFirstThunk"></span>    OriginalFirstThunk <a href="/pkg/builtin/#uint32">uint32</a>
<span id="ImportDirectory.TimeDateStamp"></span>    TimeDateStamp      <a href="/pkg/builtin/#uint32">uint32</a>
<span id="ImportDirectory.ForwarderChain"></span>    ForwarderChain     <a href="/pkg/builtin/#uint32">uint32</a>
<span id="ImportDirectory.Name"></span>    Name               <a href="/pkg/builtin/#uint32">uint32</a>
<span id="ImportDirectory.FirstThunk"></span>    FirstThunk         <a href="/pkg/builtin/#uint32">uint32</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











## <a id="OptionalHeader32">type</a> [OptionalHeader32](https://golang.org/src/debug/pe/pe.go?s=480:1637#L12)

<pre>type OptionalHeader32 struct {
<span id="OptionalHeader32.Magic"></span>    Magic                       <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader32.MajorLinkerVersion"></span>    MajorLinkerVersion          <a href="/pkg/builtin/#uint8">uint8</a>
<span id="OptionalHeader32.MinorLinkerVersion"></span>    MinorLinkerVersion          <a href="/pkg/builtin/#uint8">uint8</a>
<span id="OptionalHeader32.SizeOfCode"></span>    SizeOfCode                  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.SizeOfInitializedData"></span>    SizeOfInitializedData       <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.SizeOfUninitializedData"></span>    SizeOfUninitializedData     <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.AddressOfEntryPoint"></span>    AddressOfEntryPoint         <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.BaseOfCode"></span>    BaseOfCode                  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.BaseOfData"></span>    BaseOfData                  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.ImageBase"></span>    ImageBase                   <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.SectionAlignment"></span>    SectionAlignment            <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.FileAlignment"></span>    FileAlignment               <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.MajorOperatingSystemVersion"></span>    MajorOperatingSystemVersion <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader32.MinorOperatingSystemVersion"></span>    MinorOperatingSystemVersion <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader32.MajorImageVersion"></span>    MajorImageVersion           <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader32.MinorImageVersion"></span>    MinorImageVersion           <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader32.MajorSubsystemVersion"></span>    MajorSubsystemVersion       <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader32.MinorSubsystemVersion"></span>    MinorSubsystemVersion       <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader32.Win32VersionValue"></span>    Win32VersionValue           <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.SizeOfImage"></span>    SizeOfImage                 <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.SizeOfHeaders"></span>    SizeOfHeaders               <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.CheckSum"></span>    CheckSum                    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.Subsystem"></span>    Subsystem                   <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader32.DllCharacteristics"></span>    DllCharacteristics          <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader32.SizeOfStackReserve"></span>    SizeOfStackReserve          <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.SizeOfStackCommit"></span>    SizeOfStackCommit           <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.SizeOfHeapReserve"></span>    SizeOfHeapReserve           <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.SizeOfHeapCommit"></span>    SizeOfHeapCommit            <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.LoaderFlags"></span>    LoaderFlags                 <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.NumberOfRvaAndSizes"></span>    NumberOfRvaAndSizes         <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader32.DataDirectory"></span>    DataDirectory               [16]<a href="#DataDirectory">DataDirectory</a>
}
</pre>











## <a id="OptionalHeader64">type</a> [OptionalHeader64](https://golang.org/src/debug/pe/pe.go?s=1639:2760#L46)

<pre>type OptionalHeader64 struct {
<span id="OptionalHeader64.Magic"></span>    Magic                       <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader64.MajorLinkerVersion"></span>    MajorLinkerVersion          <a href="/pkg/builtin/#uint8">uint8</a>
<span id="OptionalHeader64.MinorLinkerVersion"></span>    MinorLinkerVersion          <a href="/pkg/builtin/#uint8">uint8</a>
<span id="OptionalHeader64.SizeOfCode"></span>    SizeOfCode                  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader64.SizeOfInitializedData"></span>    SizeOfInitializedData       <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader64.SizeOfUninitializedData"></span>    SizeOfUninitializedData     <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader64.AddressOfEntryPoint"></span>    AddressOfEntryPoint         <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader64.BaseOfCode"></span>    BaseOfCode                  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader64.ImageBase"></span>    ImageBase                   <a href="/pkg/builtin/#uint64">uint64</a>
<span id="OptionalHeader64.SectionAlignment"></span>    SectionAlignment            <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader64.FileAlignment"></span>    FileAlignment               <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader64.MajorOperatingSystemVersion"></span>    MajorOperatingSystemVersion <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader64.MinorOperatingSystemVersion"></span>    MinorOperatingSystemVersion <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader64.MajorImageVersion"></span>    MajorImageVersion           <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader64.MinorImageVersion"></span>    MinorImageVersion           <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader64.MajorSubsystemVersion"></span>    MajorSubsystemVersion       <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader64.MinorSubsystemVersion"></span>    MinorSubsystemVersion       <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader64.Win32VersionValue"></span>    Win32VersionValue           <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader64.SizeOfImage"></span>    SizeOfImage                 <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader64.SizeOfHeaders"></span>    SizeOfHeaders               <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader64.CheckSum"></span>    CheckSum                    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader64.Subsystem"></span>    Subsystem                   <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader64.DllCharacteristics"></span>    DllCharacteristics          <a href="/pkg/builtin/#uint16">uint16</a>
<span id="OptionalHeader64.SizeOfStackReserve"></span>    SizeOfStackReserve          <a href="/pkg/builtin/#uint64">uint64</a>
<span id="OptionalHeader64.SizeOfStackCommit"></span>    SizeOfStackCommit           <a href="/pkg/builtin/#uint64">uint64</a>
<span id="OptionalHeader64.SizeOfHeapReserve"></span>    SizeOfHeapReserve           <a href="/pkg/builtin/#uint64">uint64</a>
<span id="OptionalHeader64.SizeOfHeapCommit"></span>    SizeOfHeapCommit            <a href="/pkg/builtin/#uint64">uint64</a>
<span id="OptionalHeader64.LoaderFlags"></span>    LoaderFlags                 <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader64.NumberOfRvaAndSizes"></span>    NumberOfRvaAndSizes         <a href="/pkg/builtin/#uint32">uint32</a>
<span id="OptionalHeader64.DataDirectory"></span>    DataDirectory               [16]<a href="#DataDirectory">DataDirectory</a>
}
</pre>











## <a id="Reloc">type</a> [Reloc](https://golang.org/src/debug/pe/section.go?s=1185:1281#L36)
Reloc represents a PE COFF relocation.
Each section contains its own relocation list.


<pre>type Reloc struct {
<span id="Reloc.VirtualAddress"></span>    VirtualAddress   <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Reloc.SymbolTableIndex"></span>    SymbolTableIndex <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Reloc.Type"></span>    Type             <a href="/pkg/builtin/#uint16">uint16</a>
}
</pre>











## <a id="Section">type</a> [Section](https://golang.org/src/debug/pe/section.go?s=2245:2570#L74)
Section provides access to PE COFF section.


<pre>type Section struct {
    <a href="#SectionHeader">SectionHeader</a>
<span id="Section.Relocs"></span>    Relocs []<a href="#Reloc">Reloc</a>

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











### <a id="Section.Data">func</a> (\*Section) [Data](https://golang.org/src/debug/pe/section.go?s=2632:2672#L89)
<pre>func (s *<a href="#Section">Section</a>) Data() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Data reads and returns the contents of the PE section s.




### <a id="Section.Open">func</a> (\*Section) [Open](https://golang.org/src/debug/pe/section.go?s=2859:2897#L99)
<pre>func (s *<a href="#Section">Section</a>) Open() <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadSeeker">ReadSeeker</a></pre>
Open returns a new ReadSeeker reading the PE section s.




## <a id="SectionHeader">type</a> [SectionHeader](https://golang.org/src/debug/pe/section.go?s=1877:2196#L60)
SectionHeader is similar to SectionHeader32 with Name
field replaced by Go string.


<pre>type SectionHeader struct {
<span id="SectionHeader.Name"></span>    Name                 <a href="/pkg/builtin/#string">string</a>
<span id="SectionHeader.VirtualSize"></span>    VirtualSize          <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader.VirtualAddress"></span>    VirtualAddress       <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader.Size"></span>    Size                 <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader.Offset"></span>    Offset               <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader.PointerToRelocations"></span>    PointerToRelocations <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader.PointerToLineNumbers"></span>    PointerToLineNumbers <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader.NumberOfRelocations"></span>    NumberOfRelocations  <a href="/pkg/builtin/#uint16">uint16</a>
<span id="SectionHeader.NumberOfLineNumbers"></span>    NumberOfLineNumbers  <a href="/pkg/builtin/#uint16">uint16</a>
<span id="SectionHeader.Characteristics"></span>    Characteristics      <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="SectionHeader32">type</a> [SectionHeader32](https://golang.org/src/debug/pe/section.go?s=286:609#L5)
SectionHeader32 represents real PE COFF section header.


<pre>type SectionHeader32 struct {
<span id="SectionHeader32.Name"></span>    Name                 [8]<a href="/pkg/builtin/#uint8">uint8</a>
<span id="SectionHeader32.VirtualSize"></span>    VirtualSize          <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader32.VirtualAddress"></span>    VirtualAddress       <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader32.SizeOfRawData"></span>    SizeOfRawData        <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader32.PointerToRawData"></span>    PointerToRawData     <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader32.PointerToRelocations"></span>    PointerToRelocations <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader32.PointerToLineNumbers"></span>    PointerToLineNumbers <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader32.NumberOfRelocations"></span>    NumberOfRelocations  <a href="/pkg/builtin/#uint16">uint16</a>
<span id="SectionHeader32.NumberOfLineNumbers"></span>    NumberOfLineNumbers  <a href="/pkg/builtin/#uint16">uint16</a>
<span id="SectionHeader32.Characteristics"></span>    Characteristics      <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="StringTable">type</a> [StringTable](https://golang.org/src/debug/pe/string.go?s=481:504#L15)
StringTable is a COFF string table.


<pre>type StringTable []<a href="/pkg/builtin/#byte">byte</a></pre>











### <a id="StringTable.String">func</a> (StringTable) [String](https://golang.org/src/debug/pe/string.go?s=1454:1512#L48)
<pre>func (st <a href="#StringTable">StringTable</a>) String(start <a href="/pkg/builtin/#uint32">uint32</a>) (<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
String extracts string from COFF string table st at offset start.




## <a id="Symbol">type</a> [Symbol](https://golang.org/src/debug/pe/symbol.go?s=2347:2477#L82)
Symbol is similar to COFFSymbol with Name field replaced
by Go string. Symbol also does not have NumberOfAuxSymbols.


<pre>type Symbol struct {
<span id="Symbol.Name"></span>    Name          <a href="/pkg/builtin/#string">string</a>
<span id="Symbol.Value"></span>    Value         <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Symbol.SectionNumber"></span>    SectionNumber <a href="/pkg/builtin/#int16">int16</a>
<span id="Symbol.Type"></span>    Type          <a href="/pkg/builtin/#uint16">uint16</a>
<span id="Symbol.StorageClass"></span>    StorageClass  <a href="/pkg/builtin/#uint8">uint8</a>
}
</pre>














