

# macho
`import "debug/macho"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package macho implements access to Mach-O object files.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [type Cpu](#Cpu)
  * [func (i Cpu) GoString() string](#Cpu.GoString)
  * [func (i Cpu) String() string](#Cpu.String)
* [type Dylib](#Dylib)
* [type DylibCmd](#DylibCmd)
* [type Dysymtab](#Dysymtab)
* [type DysymtabCmd](#DysymtabCmd)
* [type FatArch](#FatArch)
* [type FatArchHeader](#FatArchHeader)
* [type FatFile](#FatFile)
  * [func NewFatFile(r io.ReaderAt) (*FatFile, error)](#NewFatFile)
  * [func OpenFat(name string) (*FatFile, error)](#OpenFat)
  * [func (ff *FatFile) Close() error](#FatFile.Close)
* [type File](#File)
  * [func NewFile(r io.ReaderAt) (*File, error)](#NewFile)
  * [func Open(name string) (*File, error)](#Open)
  * [func (f *File) Close() error](#File.Close)
  * [func (f *File) DWARF() (*dwarf.Data, error)](#File.DWARF)
  * [func (f *File) ImportedLibraries() ([]string, error)](#File.ImportedLibraries)
  * [func (f *File) ImportedSymbols() ([]string, error)](#File.ImportedSymbols)
  * [func (f *File) Section(name string) *Section](#File.Section)
  * [func (f *File) Segment(name string) *Segment](#File.Segment)
* [type FileHeader](#FileHeader)
* [type FormatError](#FormatError)
  * [func (e *FormatError) Error() string](#FormatError.Error)
* [type Load](#Load)
* [type LoadBytes](#LoadBytes)
  * [func (b LoadBytes) Raw() []byte](#LoadBytes.Raw)
* [type LoadCmd](#LoadCmd)
  * [func (i LoadCmd) GoString() string](#LoadCmd.GoString)
  * [func (i LoadCmd) String() string](#LoadCmd.String)
* [type Nlist32](#Nlist32)
* [type Nlist64](#Nlist64)
* [type Regs386](#Regs386)
* [type RegsAMD64](#RegsAMD64)
* [type Reloc](#Reloc)
* [type RelocTypeARM](#RelocTypeARM)
  * [func (r RelocTypeARM) GoString() string](#RelocTypeARM.GoString)
  * [func (i RelocTypeARM) String() string](#RelocTypeARM.String)
* [type RelocTypeARM64](#RelocTypeARM64)
  * [func (r RelocTypeARM64) GoString() string](#RelocTypeARM64.GoString)
  * [func (i RelocTypeARM64) String() string](#RelocTypeARM64.String)
* [type RelocTypeGeneric](#RelocTypeGeneric)
  * [func (r RelocTypeGeneric) GoString() string](#RelocTypeGeneric.GoString)
  * [func (i RelocTypeGeneric) String() string](#RelocTypeGeneric.String)
* [type RelocTypeX86_64](#RelocTypeX86_64)
  * [func (r RelocTypeX86_64) GoString() string](#RelocTypeX86_64.GoString)
  * [func (i RelocTypeX86_64) String() string](#RelocTypeX86_64.String)
* [type Rpath](#Rpath)
* [type RpathCmd](#RpathCmd)
* [type Section](#Section)
  * [func (s *Section) Data() ([]byte, error)](#Section.Data)
  * [func (s *Section) Open() io.ReadSeeker](#Section.Open)
* [type Section32](#Section32)
* [type Section64](#Section64)
* [type SectionHeader](#SectionHeader)
* [type Segment](#Segment)
  * [func (s *Segment) Data() ([]byte, error)](#Segment.Data)
  * [func (s *Segment) Open() io.ReadSeeker](#Segment.Open)
* [type Segment32](#Segment32)
* [type Segment64](#Segment64)
* [type SegmentHeader](#SegmentHeader)
* [type Symbol](#Symbol)
* [type Symtab](#Symtab)
* [type SymtabCmd](#SymtabCmd)
* [type Thread](#Thread)
* [type Type](#Type)
  * [func (t Type) GoString() string](#Type.GoString)
  * [func (t Type) String() string](#Type.String)




#### <a id="pkg-files">Package files</a>
[fat.go](https://golang.org/src/debug/macho/fat.go) [file.go](https://golang.org/src/debug/macho/file.go) [macho.go](https://golang.org/src/debug/macho/macho.go) [reloctype.go](https://golang.org/src/debug/macho/reloctype.go) [reloctype_string.go](https://golang.org/src/debug/macho/reloctype_string.go) 


## <a id="pkg-constants">Constants</a>

<pre>const (
    <span id="Magic32">Magic32</span>  <a href="/pkg/builtin/#uint32">uint32</a> = 0xfeedface
    <span id="Magic64">Magic64</span>  <a href="/pkg/builtin/#uint32">uint32</a> = 0xfeedfacf
    <span id="MagicFat">MagicFat</span> <a href="/pkg/builtin/#uint32">uint32</a> = 0xcafebabe
)</pre>
<pre>const (
    <span id="FlagNoUndefs">FlagNoUndefs</span>              <a href="/pkg/builtin/#uint32">uint32</a> = 0x1
    <span id="FlagIncrLink">FlagIncrLink</span>              <a href="/pkg/builtin/#uint32">uint32</a> = 0x2
    <span id="FlagDyldLink">FlagDyldLink</span>              <a href="/pkg/builtin/#uint32">uint32</a> = 0x4
    <span id="FlagBindAtLoad">FlagBindAtLoad</span>            <a href="/pkg/builtin/#uint32">uint32</a> = 0x8
    <span id="FlagPrebound">FlagPrebound</span>              <a href="/pkg/builtin/#uint32">uint32</a> = 0x10
    <span id="FlagSplitSegs">FlagSplitSegs</span>             <a href="/pkg/builtin/#uint32">uint32</a> = 0x20
    <span id="FlagLazyInit">FlagLazyInit</span>              <a href="/pkg/builtin/#uint32">uint32</a> = 0x40
    <span id="FlagTwoLevel">FlagTwoLevel</span>              <a href="/pkg/builtin/#uint32">uint32</a> = 0x80
    <span id="FlagForceFlat">FlagForceFlat</span>             <a href="/pkg/builtin/#uint32">uint32</a> = 0x100
    <span id="FlagNoMultiDefs">FlagNoMultiDefs</span>           <a href="/pkg/builtin/#uint32">uint32</a> = 0x200
    <span id="FlagNoFixPrebinding">FlagNoFixPrebinding</span>       <a href="/pkg/builtin/#uint32">uint32</a> = 0x400
    <span id="FlagPrebindable">FlagPrebindable</span>           <a href="/pkg/builtin/#uint32">uint32</a> = 0x800
    <span id="FlagAllModsBound">FlagAllModsBound</span>          <a href="/pkg/builtin/#uint32">uint32</a> = 0x1000
    <span id="FlagSubsectionsViaSymbols">FlagSubsectionsViaSymbols</span> <a href="/pkg/builtin/#uint32">uint32</a> = 0x2000
    <span id="FlagCanonical">FlagCanonical</span>             <a href="/pkg/builtin/#uint32">uint32</a> = 0x4000
    <span id="FlagWeakDefines">FlagWeakDefines</span>           <a href="/pkg/builtin/#uint32">uint32</a> = 0x8000
    <span id="FlagBindsToWeak">FlagBindsToWeak</span>           <a href="/pkg/builtin/#uint32">uint32</a> = 0x10000
    <span id="FlagAllowStackExecution">FlagAllowStackExecution</span>   <a href="/pkg/builtin/#uint32">uint32</a> = 0x20000
    <span id="FlagRootSafe">FlagRootSafe</span>              <a href="/pkg/builtin/#uint32">uint32</a> = 0x40000
    <span id="FlagSetuidSafe">FlagSetuidSafe</span>            <a href="/pkg/builtin/#uint32">uint32</a> = 0x80000
    <span id="FlagNoReexportedDylibs">FlagNoReexportedDylibs</span>    <a href="/pkg/builtin/#uint32">uint32</a> = 0x100000
    <span id="FlagPIE">FlagPIE</span>                   <a href="/pkg/builtin/#uint32">uint32</a> = 0x200000
    <span id="FlagDeadStrippableDylib">FlagDeadStrippableDylib</span>   <a href="/pkg/builtin/#uint32">uint32</a> = 0x400000
    <span id="FlagHasTLVDescriptors">FlagHasTLVDescriptors</span>     <a href="/pkg/builtin/#uint32">uint32</a> = 0x800000
    <span id="FlagNoHeapExecution">FlagNoHeapExecution</span>       <a href="/pkg/builtin/#uint32">uint32</a> = 0x1000000
    <span id="FlagAppExtensionSafe">FlagAppExtensionSafe</span>      <a href="/pkg/builtin/#uint32">uint32</a> = 0x2000000
)</pre>

## <a id="pkg-variables">Variables</a>
ErrNotFat is returned from NewFatFile or OpenFat when the file is not a
universal binary but may be a thin binary, based on its magic number.


<pre>var <span id="ErrNotFat">ErrNotFat</span> = &amp;<a href="#FormatError">FormatError</a>{0, &#34;not a fat Mach-O file&#34;, <a href="/pkg/builtin/#nil">nil</a>}</pre>



## <a id="Cpu">type</a> [Cpu](https://golang.org/src/debug/macho/macho.go?s=1536:1551#L50)
A Cpu is a Mach-O cpu type.


<pre>type Cpu <a href="/pkg/builtin/#uint32">uint32</a></pre>



<pre>const (
    <span id="Cpu386">Cpu386</span>   <a href="#Cpu">Cpu</a> = 7
    <span id="CpuAmd64">CpuAmd64</span> <a href="#Cpu">Cpu</a> = <a href="#Cpu386">Cpu386</a> | cpuArch64
    <span id="CpuArm">CpuArm</span>   <a href="#Cpu">Cpu</a> = 12
    <span id="CpuArm64">CpuArm64</span> <a href="#Cpu">Cpu</a> = <a href="#CpuArm">CpuArm</a> | cpuArch64
    <span id="CpuPpc">CpuPpc</span>   <a href="#Cpu">Cpu</a> = 18
    <span id="CpuPpc64">CpuPpc64</span> <a href="#Cpu">Cpu</a> = <a href="#CpuPpc">CpuPpc</a> | cpuArch64
)</pre>









### <a id="Cpu.GoString">func</a> (Cpu) [GoString](https://golang.org/src/debug/macho/macho.go?s=2055:2085#L73)
<pre>func (i <a href="#Cpu">Cpu</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Cpu.String">func</a> (Cpu) [String](https://golang.org/src/debug/macho/macho.go?s=1972:2000#L72)
<pre>func (i <a href="#Cpu">Cpu</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Dylib">type</a> [Dylib](https://golang.org/src/debug/macho/file.go?s=3266:3390#L131)
A Dylib represents a Mach-O load dynamic library command.


<pre>type Dylib struct {
    <a href="#LoadBytes">LoadBytes</a>
<span id="Dylib.Name"></span>    Name           <a href="/pkg/builtin/#string">string</a>
<span id="Dylib.Time"></span>    Time           <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Dylib.CurrentVersion"></span>    CurrentVersion <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Dylib.CompatVersion"></span>    CompatVersion  <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="DylibCmd">type</a> [DylibCmd](https://golang.org/src/debug/macho/macho.go?s=4429:4594#L168)
A DylibCmd is a Mach-O load dynamic library command.


<pre>type DylibCmd struct {
<span id="DylibCmd.Cmd"></span>    Cmd            <a href="#LoadCmd">LoadCmd</a>
<span id="DylibCmd.Len"></span>    Len            <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DylibCmd.Name"></span>    Name           <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DylibCmd.Time"></span>    Time           <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DylibCmd.CurrentVersion"></span>    CurrentVersion <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DylibCmd.CompatVersion"></span>    CompatVersion  <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="Dysymtab">type</a> [Dysymtab](https://golang.org/src/debug/macho/file.go?s=3571:3670#L147)
A Dysymtab represents a Mach-O dynamic symbol table command.


<pre>type Dysymtab struct {
    <a href="#LoadBytes">LoadBytes</a>
    <a href="#DysymtabCmd">DysymtabCmd</a>
<span id="Dysymtab.IndirectSyms"></span>    IndirectSyms []<a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">// indices into Symtab.Syms</span>
}
</pre>











## <a id="DysymtabCmd">type</a> [DysymtabCmd](https://golang.org/src/debug/macho/macho.go?s=3865:4369#L144)
A DysymtabCmd is a Mach-O dynamic symbol table command.


<pre>type DysymtabCmd struct {
<span id="DysymtabCmd.Cmd"></span>    Cmd            <a href="#LoadCmd">LoadCmd</a>
<span id="DysymtabCmd.Len"></span>    Len            <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Ilocalsym"></span>    Ilocalsym      <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Nlocalsym"></span>    Nlocalsym      <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Iextdefsym"></span>    Iextdefsym     <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Nextdefsym"></span>    Nextdefsym     <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Iundefsym"></span>    Iundefsym      <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Nundefsym"></span>    Nundefsym      <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Tocoffset"></span>    Tocoffset      <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Ntoc"></span>    Ntoc           <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Modtaboff"></span>    Modtaboff      <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Nmodtab"></span>    Nmodtab        <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Extrefsymoff"></span>    Extrefsymoff   <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Nextrefsyms"></span>    Nextrefsyms    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Indirectsymoff"></span>    Indirectsymoff <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Nindirectsyms"></span>    Nindirectsyms  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Extreloff"></span>    Extreloff      <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Nextrel"></span>    Nextrel        <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Locreloff"></span>    Locreloff      <a href="/pkg/builtin/#uint32">uint32</a>
<span id="DysymtabCmd.Nlocrel"></span>    Nlocrel        <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="FatArch">type</a> [FatArch](https://golang.org/src/debug/macho/fat.go?s=646:691#L23)
A FatArch is a Mach-O File inside a FatFile.


<pre>type FatArch struct {
    <a href="#FatArchHeader">FatArchHeader</a>
    *<a href="#File">File</a>
}
</pre>











## <a id="FatArchHeader">type</a> [FatArchHeader](https://golang.org/src/debug/macho/fat.go?s=462:563#L12)
A FatArchHeader represents a fat header for a specific image architecture.


<pre>type FatArchHeader struct {
<span id="FatArchHeader.Cpu"></span>    Cpu    <a href="#Cpu">Cpu</a>
<span id="FatArchHeader.SubCpu"></span>    SubCpu <a href="/pkg/builtin/#uint32">uint32</a>
<span id="FatArchHeader.Offset"></span>    Offset <a href="/pkg/builtin/#uint32">uint32</a>
<span id="FatArchHeader.Size"></span>    Size   <a href="/pkg/builtin/#uint32">uint32</a>
<span id="FatArchHeader.Align"></span>    Align  <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="FatFile">type</a> [FatFile](https://golang.org/src/debug/macho/fat.go?s=308:382#L5)
A FatFile is a Mach-O universal binary that contains at least one architecture.


<pre>type FatFile struct {
<span id="FatFile.Magic"></span>    Magic  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="FatFile.Arches"></span>    Arches []<a href="#FatArch">FatArch</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewFatFile">func</a> [NewFatFile](https://golang.org/src/debug/macho/fat.go?s=1075:1123#L35)
<pre>func NewFatFile(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReaderAt">ReaderAt</a>) (*<a href="#FatFile">FatFile</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewFatFile creates a new FatFile for accessing all the Mach-O images in a
universal binary. The Mach-O binary is expected to start at position 0 in
the ReaderAt.




### <a id="OpenFat">func</a> [OpenFat](https://golang.org/src/debug/macho/fat.go?s=3687:3730#L115)
<pre>func OpenFat(name <a href="/pkg/builtin/#string">string</a>) (*<a href="#FatFile">FatFile</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
OpenFat opens the named file using os.Open and prepares it for use as a Mach-O
universal binary.






### <a id="FatFile.Close">func</a> (\*FatFile) [Close](https://golang.org/src/debug/macho/fat.go?s=3906:3938#L129)
<pre>func (ff *<a href="#FatFile">FatFile</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>



## <a id="File">type</a> [File](https://golang.org/src/debug/macho/file.go?s=430:588#L12)
A File represents an open Mach-O file.


<pre>type File struct {
    <a href="#FileHeader">FileHeader</a>
<span id="File.ByteOrder"></span>    ByteOrder <a href="/pkg/encoding/binary/">binary</a>.<a href="/pkg/encoding/binary/#ByteOrder">ByteOrder</a>
<span id="File.Loads"></span>    Loads     []<a href="#Load">Load</a>
<span id="File.Sections"></span>    Sections  []*<a href="#Section">Section</a>

<span id="File.Symtab"></span>    Symtab   *<a href="#Symtab">Symtab</a>
<span id="File.Dysymtab"></span>    Dysymtab *<a href="#Dysymtab">Dysymtab</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewFile">func</a> [NewFile](https://golang.org/src/debug/macho/file.go?s=5006:5048#L218)
<pre>func NewFile(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReaderAt">ReaderAt</a>) (*<a href="#File">File</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewFile creates a new File for accessing a Mach-O binary in an underlying reader.
The Mach-O binary is expected to start at position 0 in the ReaderAt.




### <a id="Open">func</a> [Open](https://golang.org/src/debug/macho/file.go?s=4393:4430#L190)
<pre>func Open(name <a href="/pkg/builtin/#string">string</a>) (*<a href="#File">File</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Open opens the named file using os.Open and prepares it for use as a Mach-O binary.






### <a id="File.Close">func</a> (\*File) [Close](https://golang.org/src/debug/macho/file.go?s=4720:4748#L207)
<pre>func (f *<a href="#File">File</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the File.
If the File was created using NewFile directly instead of Open,
Close has no effect.




### <a id="File.DWARF">func</a> (\*File) [DWARF](https://golang.org/src/debug/macho/file.go?s=13936:13979#L569)
<pre>func (f *<a href="#File">File</a>) DWARF() (*<a href="/pkg/debug/dwarf/">dwarf</a>.<a href="/pkg/debug/dwarf/#Data">Data</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DWARF returns the DWARF debug information for the Mach-O file.




### <a id="File.ImportedLibraries">func</a> (\*File) [ImportedLibraries](https://golang.org/src/debug/macho/file.go?s=16319:16371#L670)
<pre>func (f *<a href="#File">File</a>) ImportedLibraries() ([]<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ImportedLibraries returns the paths of all libraries
referred to by the binary f that are expected to be
linked with the binary at dynamic link time.




### <a id="File.ImportedSymbols">func</a> (\*File) [ImportedSymbols](https://golang.org/src/debug/macho/file.go?s=15827:15877#L653)
<pre>func (f *<a href="#File">File</a>) ImportedSymbols() ([]<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ImportedSymbols returns the names of all symbols
referred to by the binary f that are expected to be
satisfied by other libraries at dynamic load time.




### <a id="File.Section">func</a> (\*File) [Section](https://golang.org/src/debug/macho/file.go?s=13735:13779#L559)
<pre>func (f *<a href="#File">File</a>) Section(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Section">Section</a></pre>
Section returns the first section with the given name, or nil if no such
section exists.




### <a id="File.Segment">func</a> (\*File) [Segment](https://golang.org/src/debug/macho/file.go?s=13479:13523#L548)
<pre>func (f *<a href="#File">File</a>) Segment(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Segment">Segment</a></pre>
Segment returns the first Segment with the given name, or nil if no such segment exists.




## <a id="FileHeader">type</a> [FileHeader](https://golang.org/src/debug/macho/macho.go?s=690:816#L8)
A FileHeader represents a Mach-O file header.


<pre>type FileHeader struct {
<span id="FileHeader.Magic"></span>    Magic  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="FileHeader.Cpu"></span>    Cpu    <a href="#Cpu">Cpu</a>
<span id="FileHeader.SubCpu"></span>    SubCpu <a href="/pkg/builtin/#uint32">uint32</a>
<span id="FileHeader.Type"></span>    Type   <a href="#Type">Type</a>
<span id="FileHeader.Ncmd"></span>    Ncmd   <a href="/pkg/builtin/#uint32">uint32</a>
<span id="FileHeader.Cmdsz"></span>    Cmdsz  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="FileHeader.Flags"></span>    Flags  <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="FormatError">type</a> [FormatError](https://golang.org/src/debug/macho/file.go?s=4057:4124#L174)
FormatError is returned by some operations if the data does
not have the correct format for an object file.


<pre>type FormatError struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="FormatError.Error">func</a> (\*FormatError) [Error](https://golang.org/src/debug/macho/file.go?s=4126:4162#L180)
<pre>func (e *<a href="#FormatError">FormatError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Load">type</a> [Load](https://golang.org/src/debug/macho/file.go?s=636:673#L25)
A Load represents any Mach-O load command.


<pre>type Load interface {
    Raw() []<a href="/pkg/builtin/#byte">byte</a>
}</pre>











## <a id="LoadBytes">type</a> [LoadBytes](https://golang.org/src/debug/macho/file.go?s=743:764#L30)
A LoadBytes is the uninterpreted bytes of a Mach-O load command.


<pre>type LoadBytes []<a href="/pkg/builtin/#byte">byte</a></pre>











### <a id="LoadBytes.Raw">func</a> (LoadBytes) [Raw](https://golang.org/src/debug/macho/file.go?s=766:797#L32)
<pre>func (b <a href="#LoadBytes">LoadBytes</a>) Raw() []<a href="/pkg/builtin/#byte">byte</a></pre>



## <a id="LoadCmd">type</a> [LoadCmd](https://golang.org/src/debug/macho/macho.go?s=2177:2196#L76)
A LoadCmd is a Mach-O load command.


<pre>type LoadCmd <a href="/pkg/builtin/#uint32">uint32</a></pre>



<pre>const (
    <span id="LoadCmdSegment">LoadCmdSegment</span>    <a href="#LoadCmd">LoadCmd</a> = 0x1
    <span id="LoadCmdSymtab">LoadCmdSymtab</span>     <a href="#LoadCmd">LoadCmd</a> = 0x2
    <span id="LoadCmdThread">LoadCmdThread</span>     <a href="#LoadCmd">LoadCmd</a> = 0x4
    <span id="LoadCmdUnixThread">LoadCmdUnixThread</span> <a href="#LoadCmd">LoadCmd</a> = 0x5 <span class="comment">// thread+stack</span>
    <span id="LoadCmdDysymtab">LoadCmdDysymtab</span>   <a href="#LoadCmd">LoadCmd</a> = 0xb
    <span id="LoadCmdDylib">LoadCmdDylib</span>      <a href="#LoadCmd">LoadCmd</a> = 0xc <span class="comment">// load dylib command</span>
    <span id="LoadCmdDylinker">LoadCmdDylinker</span>   <a href="#LoadCmd">LoadCmd</a> = 0xf <span class="comment">// id dylinker command (not load dylinker command)</span>
    <span id="LoadCmdSegment64">LoadCmdSegment64</span>  <a href="#LoadCmd">LoadCmd</a> = 0x19
    <span id="LoadCmdRpath">LoadCmdRpath</span>      <a href="#LoadCmd">LoadCmd</a> = 0x8000001c
)</pre>









### <a id="LoadCmd.GoString">func</a> (LoadCmd) [GoString](https://golang.org/src/debug/macho/macho.go?s=2991:3025#L100)
<pre>func (i <a href="#LoadCmd">LoadCmd</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="LoadCmd.String">func</a> (LoadCmd) [String](https://golang.org/src/debug/macho/macho.go?s=2904:2936#L99)
<pre>func (i <a href="#LoadCmd">LoadCmd</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Nlist32">type</a> [Nlist32](https://golang.org/src/debug/macho/macho.go?s=6562:6653#L254)
An Nlist32 is a Mach-O 32-bit symbol table entry.


<pre>type Nlist32 struct {
<span id="Nlist32.Name"></span>    Name  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Nlist32.Type"></span>    Type  <a href="/pkg/builtin/#uint8">uint8</a>
<span id="Nlist32.Sect"></span>    Sect  <a href="/pkg/builtin/#uint8">uint8</a>
<span id="Nlist32.Desc"></span>    Desc  <a href="/pkg/builtin/#uint16">uint16</a>
<span id="Nlist32.Value"></span>    Value <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="Nlist64">type</a> [Nlist64](https://golang.org/src/debug/macho/macho.go?s=6708:6799#L263)
An Nlist64 is a Mach-O 64-bit symbol table entry.


<pre>type Nlist64 struct {
<span id="Nlist64.Name"></span>    Name  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Nlist64.Type"></span>    Type  <a href="/pkg/builtin/#uint8">uint8</a>
<span id="Nlist64.Sect"></span>    Sect  <a href="/pkg/builtin/#uint8">uint8</a>
<span id="Nlist64.Desc"></span>    Desc  <a href="/pkg/builtin/#uint16">uint16</a>
<span id="Nlist64.Value"></span>    Value <a href="/pkg/builtin/#uint64">uint64</a>
}
</pre>











## <a id="Regs386">type</a> [Regs386](https://golang.org/src/debug/macho/macho.go?s=6850:7097#L272)
Regs386 is the Mach-O 386 register structure.


<pre>type Regs386 struct {
<span id="Regs386.AX"></span>    AX    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Regs386.BX"></span>    BX    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Regs386.CX"></span>    CX    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Regs386.DX"></span>    DX    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Regs386.DI"></span>    DI    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Regs386.SI"></span>    SI    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Regs386.BP"></span>    BP    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Regs386.SP"></span>    SP    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Regs386.SS"></span>    SS    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Regs386.FLAGS"></span>    FLAGS <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Regs386.IP"></span>    IP    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Regs386.CS"></span>    CS    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Regs386.DS"></span>    DS    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Regs386.ES"></span>    ES    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Regs386.FS"></span>    FS    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Regs386.GS"></span>    GS    <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="RegsAMD64">type</a> [RegsAMD64](https://golang.org/src/debug/macho/macho.go?s=7152:7471#L292)
RegsAMD64 is the Mach-O AMD64 register structure.


<pre>type RegsAMD64 struct {
<span id="RegsAMD64.AX"></span>    AX    <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.BX"></span>    BX    <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.CX"></span>    CX    <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.DX"></span>    DX    <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.DI"></span>    DI    <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.SI"></span>    SI    <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.BP"></span>    BP    <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.SP"></span>    SP    <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.R8"></span>    R8    <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.R9"></span>    R9    <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.R10"></span>    R10   <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.R11"></span>    R11   <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.R12"></span>    R12   <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.R13"></span>    R13   <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.R14"></span>    R14   <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.R15"></span>    R15   <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.IP"></span>    IP    <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.FLAGS"></span>    FLAGS <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.CS"></span>    CS    <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.FS"></span>    FS    <a href="/pkg/builtin/#uint64">uint64</a>
<span id="RegsAMD64.GS"></span>    GS    <a href="/pkg/builtin/#uint64">uint64</a>
}
</pre>











## <a id="Reloc">type</a> [Reloc](https://golang.org/src/debug/macho/file.go?s=2074:2496#L90)
A Reloc represents a Mach-O relocation.


<pre>type Reloc struct {
<span id="Reloc.Addr"></span>    Addr  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Reloc.Value"></span>    Value <a href="/pkg/builtin/#uint32">uint32</a>
    <span class="comment">// when Scattered == false &amp;&amp; Extern == true, Value is the symbol number.</span>
    <span class="comment">// when Scattered == false &amp;&amp; Extern == false, Value is the section number.</span>
    <span class="comment">// when Scattered == true, Value is the value that this reloc refers to.</span>
<span id="Reloc.Type"></span>    Type      <a href="/pkg/builtin/#uint8">uint8</a>
<span id="Reloc.Len"></span>    Len       <a href="/pkg/builtin/#uint8">uint8</a> <span class="comment">// 0=byte, 1=word, 2=long, 3=quad</span>
<span id="Reloc.Pcrel"></span>    Pcrel     <a href="/pkg/builtin/#bool">bool</a>
<span id="Reloc.Extern"></span>    Extern    <a href="/pkg/builtin/#bool">bool</a> <span class="comment">// valid if Scattered == false</span>
<span id="Reloc.Scattered"></span>    Scattered <a href="/pkg/builtin/#bool">bool</a>
}
</pre>











## <a id="RelocTypeARM">type</a> [RelocTypeARM](https://golang.org/src/debug/macho/reloctype.go?s=1280:1301#L29)

<pre>type RelocTypeARM <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="ARM_RELOC_VANILLA">ARM_RELOC_VANILLA</span>        <a href="#RelocTypeARM">RelocTypeARM</a> = 0
    <span id="ARM_RELOC_PAIR">ARM_RELOC_PAIR</span>           <a href="#RelocTypeARM">RelocTypeARM</a> = 1
    <span id="ARM_RELOC_SECTDIFF">ARM_RELOC_SECTDIFF</span>       <a href="#RelocTypeARM">RelocTypeARM</a> = 2
    <span id="ARM_RELOC_LOCAL_SECTDIFF">ARM_RELOC_LOCAL_SECTDIFF</span> <a href="#RelocTypeARM">RelocTypeARM</a> = 3
    <span id="ARM_RELOC_PB_LA_PTR">ARM_RELOC_PB_LA_PTR</span>      <a href="#RelocTypeARM">RelocTypeARM</a> = 4
    <span id="ARM_RELOC_BR24">ARM_RELOC_BR24</span>           <a href="#RelocTypeARM">RelocTypeARM</a> = 5
    <span id="ARM_THUMB_RELOC_BR22">ARM_THUMB_RELOC_BR22</span>     <a href="#RelocTypeARM">RelocTypeARM</a> = 6
    <span id="ARM_THUMB_32BIT_BRANCH">ARM_THUMB_32BIT_BRANCH</span>   <a href="#RelocTypeARM">RelocTypeARM</a> = 7
    <span id="ARM_RELOC_HALF">ARM_RELOC_HALF</span>           <a href="#RelocTypeARM">RelocTypeARM</a> = 8
    <span id="ARM_RELOC_HALF_SECTDIFF">ARM_RELOC_HALF_SECTDIFF</span>  <a href="#RelocTypeARM">RelocTypeARM</a> = 9
)</pre>









### <a id="RelocTypeARM.GoString">func</a> (RelocTypeARM) [GoString](https://golang.org/src/debug/macho/reloctype.go?s=1744:1783#L44)
<pre>func (r <a href="#RelocTypeARM">RelocTypeARM</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="RelocTypeARM.String">func</a> (RelocTypeARM) [String](https://golang.org/src/debug/macho/reloctype_string.go?s=1578:1615#L23)
<pre>func (i <a href="#RelocTypeARM">RelocTypeARM</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="RelocTypeARM64">type</a> [RelocTypeARM64](https://golang.org/src/debug/macho/reloctype.go?s=1818:1841#L46)

<pre>type RelocTypeARM64 <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="ARM64_RELOC_UNSIGNED">ARM64_RELOC_UNSIGNED</span>            <a href="#RelocTypeARM64">RelocTypeARM64</a> = 0
    <span id="ARM64_RELOC_SUBTRACTOR">ARM64_RELOC_SUBTRACTOR</span>          <a href="#RelocTypeARM64">RelocTypeARM64</a> = 1
    <span id="ARM64_RELOC_BRANCH26">ARM64_RELOC_BRANCH26</span>            <a href="#RelocTypeARM64">RelocTypeARM64</a> = 2
    <span id="ARM64_RELOC_PAGE21">ARM64_RELOC_PAGE21</span>              <a href="#RelocTypeARM64">RelocTypeARM64</a> = 3
    <span id="ARM64_RELOC_PAGEOFF12">ARM64_RELOC_PAGEOFF12</span>           <a href="#RelocTypeARM64">RelocTypeARM64</a> = 4
    <span id="ARM64_RELOC_GOT_LOAD_PAGE21">ARM64_RELOC_GOT_LOAD_PAGE21</span>     <a href="#RelocTypeARM64">RelocTypeARM64</a> = 5
    <span id="ARM64_RELOC_GOT_LOAD_PAGEOFF12">ARM64_RELOC_GOT_LOAD_PAGEOFF12</span>  <a href="#RelocTypeARM64">RelocTypeARM64</a> = 6
    <span id="ARM64_RELOC_POINTER_TO_GOT">ARM64_RELOC_POINTER_TO_GOT</span>      <a href="#RelocTypeARM64">RelocTypeARM64</a> = 7
    <span id="ARM64_RELOC_TLVP_LOAD_PAGE21">ARM64_RELOC_TLVP_LOAD_PAGE21</span>    <a href="#RelocTypeARM64">RelocTypeARM64</a> = 8
    <span id="ARM64_RELOC_TLVP_LOAD_PAGEOFF12">ARM64_RELOC_TLVP_LOAD_PAGEOFF12</span> <a href="#RelocTypeARM64">RelocTypeARM64</a> = 9
    <span id="ARM64_RELOC_ADDEND">ARM64_RELOC_ADDEND</span>              <a href="#RelocTypeARM64">RelocTypeARM64</a> = 10
)</pre>









### <a id="RelocTypeARM64.GoString">func</a> (RelocTypeARM64) [GoString](https://golang.org/src/debug/macho/reloctype.go?s=2427:2468#L62)
<pre>func (r <a href="#RelocTypeARM64">RelocTypeARM64</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="RelocTypeARM64.String">func</a> (RelocTypeARM64) [String](https://golang.org/src/debug/macho/reloctype_string.go?s=2215:2254#L34)
<pre>func (i <a href="#RelocTypeARM64">RelocTypeARM64</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="RelocTypeGeneric">type</a> [RelocTypeGeneric](https://golang.org/src/debug/macho/reloctype.go?s=294:319#L1)

<pre>type RelocTypeGeneric <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="GENERIC_RELOC_VANILLA">GENERIC_RELOC_VANILLA</span>        <a href="#RelocTypeGeneric">RelocTypeGeneric</a> = 0
    <span id="GENERIC_RELOC_PAIR">GENERIC_RELOC_PAIR</span>           <a href="#RelocTypeGeneric">RelocTypeGeneric</a> = 1
    <span id="GENERIC_RELOC_SECTDIFF">GENERIC_RELOC_SECTDIFF</span>       <a href="#RelocTypeGeneric">RelocTypeGeneric</a> = 2
    <span id="GENERIC_RELOC_PB_LA_PTR">GENERIC_RELOC_PB_LA_PTR</span>      <a href="#RelocTypeGeneric">RelocTypeGeneric</a> = 3
    <span id="GENERIC_RELOC_LOCAL_SECTDIFF">GENERIC_RELOC_LOCAL_SECTDIFF</span> <a href="#RelocTypeGeneric">RelocTypeGeneric</a> = 4
    <span id="GENERIC_RELOC_TLV">GENERIC_RELOC_TLV</span>            <a href="#RelocTypeGeneric">RelocTypeGeneric</a> = 5
)</pre>









### <a id="RelocTypeGeneric.GoString">func</a> (RelocTypeGeneric) [GoString](https://golang.org/src/debug/macho/reloctype.go?s=638:681#L10)
<pre>func (r <a href="#RelocTypeGeneric">RelocTypeGeneric</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="RelocTypeGeneric.String">func</a> (RelocTypeGeneric) [String](https://golang.org/src/debug/macho/reloctype_string.go?s=410:451#L1)
<pre>func (i <a href="#RelocTypeGeneric">RelocTypeGeneric</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="RelocTypeX86_64">type</a> [RelocTypeX86_64](https://golang.org/src/debug/macho/reloctype.go?s=716:740#L12)

<pre>type RelocTypeX86_64 <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="X86_64_RELOC_UNSIGNED">X86_64_RELOC_UNSIGNED</span>   <a href="#RelocTypeX86_64">RelocTypeX86_64</a> = 0
    <span id="X86_64_RELOC_SIGNED">X86_64_RELOC_SIGNED</span>     <a href="#RelocTypeX86_64">RelocTypeX86_64</a> = 1
    <span id="X86_64_RELOC_BRANCH">X86_64_RELOC_BRANCH</span>     <a href="#RelocTypeX86_64">RelocTypeX86_64</a> = 2
    <span id="X86_64_RELOC_GOT_LOAD">X86_64_RELOC_GOT_LOAD</span>   <a href="#RelocTypeX86_64">RelocTypeX86_64</a> = 3
    <span id="X86_64_RELOC_GOT">X86_64_RELOC_GOT</span>        <a href="#RelocTypeX86_64">RelocTypeX86_64</a> = 4
    <span id="X86_64_RELOC_SUBTRACTOR">X86_64_RELOC_SUBTRACTOR</span> <a href="#RelocTypeX86_64">RelocTypeX86_64</a> = 5
    <span id="X86_64_RELOC_SIGNED_1">X86_64_RELOC_SIGNED_1</span>   <a href="#RelocTypeX86_64">RelocTypeX86_64</a> = 6
    <span id="X86_64_RELOC_SIGNED_2">X86_64_RELOC_SIGNED_2</span>   <a href="#RelocTypeX86_64">RelocTypeX86_64</a> = 7
    <span id="X86_64_RELOC_SIGNED_4">X86_64_RELOC_SIGNED_4</span>   <a href="#RelocTypeX86_64">RelocTypeX86_64</a> = 8
    <span id="X86_64_RELOC_TLV">X86_64_RELOC_TLV</span>        <a href="#RelocTypeX86_64">RelocTypeX86_64</a> = 9
)</pre>









### <a id="RelocTypeX86_64.GoString">func</a> (RelocTypeX86\_64) [GoString](https://golang.org/src/debug/macho/reloctype.go?s=1203:1245#L27)
<pre>func (r <a href="#RelocTypeX86_64">RelocTypeX86_64</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="RelocTypeX86_64.String">func</a> (RelocTypeX86\_64) [String](https://golang.org/src/debug/macho/reloctype_string.go?s=1007:1047#L12)
<pre>func (i <a href="#RelocTypeX86_64">RelocTypeX86_64</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Rpath">type</a> [Rpath](https://golang.org/src/debug/macho/file.go?s=3718:3763#L154)
A Rpath represents a Mach-O rpath command.


<pre>type Rpath struct {
    <a href="#LoadBytes">LoadBytes</a>
<span id="Rpath.Path"></span>    Path <a href="/pkg/builtin/#string">string</a>
}
</pre>











## <a id="RpathCmd">type</a> [RpathCmd](https://golang.org/src/debug/macho/macho.go?s=4639:4702#L178)
A RpathCmd is a Mach-O rpath command.


<pre>type RpathCmd struct {
<span id="RpathCmd.Cmd"></span>    Cmd  <a href="#LoadCmd">LoadCmd</a>
<span id="RpathCmd.Len"></span>    Len  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="RpathCmd.Path"></span>    Path <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="Section">type</a> [Section](https://golang.org/src/debug/macho/file.go?s=2498:2823#L103)

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











### <a id="Section.Data">func</a> (\*Section) [Data](https://golang.org/src/debug/macho/file.go?s=2887:2927#L118)
<pre>func (s *<a href="#Section">Section</a>) Data() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Data reads and returns the contents of the Mach-O section.




### <a id="Section.Open">func</a> (\*Section) [Open](https://golang.org/src/debug/macho/file.go?s=3116:3154#L128)
<pre>func (s *<a href="#Section">Section</a>) Open() <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadSeeker">ReadSeeker</a></pre>
Open returns a new ReadSeeker reading the Mach-O section.




## <a id="Section32">type</a> [Section32](https://golang.org/src/debug/macho/macho.go?s=6006:6222#L223)
A Section32 is a 32-bit Mach-O section header.


<pre>type Section32 struct {
<span id="Section32.Name"></span>    Name     [16]<a href="/pkg/builtin/#byte">byte</a>
<span id="Section32.Seg"></span>    Seg      [16]<a href="/pkg/builtin/#byte">byte</a>
<span id="Section32.Addr"></span>    Addr     <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Section32.Size"></span>    Size     <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Section32.Offset"></span>    Offset   <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Section32.Align"></span>    Align    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Section32.Reloff"></span>    Reloff   <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Section32.Nreloc"></span>    Nreloc   <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Section32.Flags"></span>    Flags    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Section32.Reserve1"></span>    Reserve1 <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Section32.Reserve2"></span>    Reserve2 <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="Section64">type</a> [Section64](https://golang.org/src/debug/macho/macho.go?s=6274:6507#L238)
A Section64 is a 64-bit Mach-O section header.


<pre>type Section64 struct {
<span id="Section64.Name"></span>    Name     [16]<a href="/pkg/builtin/#byte">byte</a>
<span id="Section64.Seg"></span>    Seg      [16]<a href="/pkg/builtin/#byte">byte</a>
<span id="Section64.Addr"></span>    Addr     <a href="/pkg/builtin/#uint64">uint64</a>
<span id="Section64.Size"></span>    Size     <a href="/pkg/builtin/#uint64">uint64</a>
<span id="Section64.Offset"></span>    Offset   <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Section64.Align"></span>    Align    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Section64.Reloff"></span>    Reloff   <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Section64.Nreloc"></span>    Nreloc   <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Section64.Flags"></span>    Flags    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Section64.Reserve1"></span>    Reserve1 <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Section64.Reserve2"></span>    Reserve2 <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Section64.Reserve3"></span>    Reserve3 <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="SectionHeader">type</a> [SectionHeader](https://golang.org/src/debug/macho/file.go?s=1865:2029#L77)

<pre>type SectionHeader struct {
<span id="SectionHeader.Name"></span>    Name   <a href="/pkg/builtin/#string">string</a>
<span id="SectionHeader.Seg"></span>    Seg    <a href="/pkg/builtin/#string">string</a>
<span id="SectionHeader.Addr"></span>    Addr   <a href="/pkg/builtin/#uint64">uint64</a>
<span id="SectionHeader.Size"></span>    Size   <a href="/pkg/builtin/#uint64">uint64</a>
<span id="SectionHeader.Offset"></span>    Offset <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader.Align"></span>    Align  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader.Reloff"></span>    Reloff <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader.Nreloc"></span>    Nreloc <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader.Flags"></span>    Flags  <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="Segment">type</a> [Segment](https://golang.org/src/debug/macho/file.go?s=1177:1497#L50)
A Segment represents a Mach-O 32-bit or 64-bit load segment command.


<pre>type Segment struct {
    <a href="#LoadBytes">LoadBytes</a>
    <a href="#SegmentHeader">SegmentHeader</a>

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











### <a id="Segment.Data">func</a> (\*Segment) [Data](https://golang.org/src/debug/macho/file.go?s=1554:1594#L65)
<pre>func (s *<a href="#Segment">Segment</a>) Data() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Data reads and returns the contents of the segment.




### <a id="Segment.Open">func</a> (\*Segment) [Open](https://golang.org/src/debug/macho/file.go?s=1776:1814#L75)
<pre>func (s *<a href="#Segment">Segment</a>) Open() <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadSeeker">ReadSeeker</a></pre>
Open returns a new ReadSeeker reading the segment.




## <a id="Segment32">type</a> [Segment32](https://golang.org/src/debug/macho/macho.go?s=3143:3354#L104)
A Segment32 is a 32-bit Mach-O segment load command.


<pre>type Segment32 struct {
<span id="Segment32.Cmd"></span>    Cmd     <a href="#LoadCmd">LoadCmd</a>
<span id="Segment32.Len"></span>    Len     <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Segment32.Name"></span>    Name    [16]<a href="/pkg/builtin/#byte">byte</a>
<span id="Segment32.Addr"></span>    Addr    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Segment32.Memsz"></span>    Memsz   <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Segment32.Offset"></span>    Offset  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Segment32.Filesz"></span>    Filesz  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Segment32.Maxprot"></span>    Maxprot <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Segment32.Prot"></span>    Prot    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Segment32.Nsect"></span>    Nsect   <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Segment32.Flag"></span>    Flag    <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="Segment64">type</a> [Segment64](https://golang.org/src/debug/macho/macho.go?s=3414:3625#L119)
A Segment64 is a 64-bit Mach-O segment load command.


<pre>type Segment64 struct {
<span id="Segment64.Cmd"></span>    Cmd     <a href="#LoadCmd">LoadCmd</a>
<span id="Segment64.Len"></span>    Len     <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Segment64.Name"></span>    Name    [16]<a href="/pkg/builtin/#byte">byte</a>
<span id="Segment64.Addr"></span>    Addr    <a href="/pkg/builtin/#uint64">uint64</a>
<span id="Segment64.Memsz"></span>    Memsz   <a href="/pkg/builtin/#uint64">uint64</a>
<span id="Segment64.Offset"></span>    Offset  <a href="/pkg/builtin/#uint64">uint64</a>
<span id="Segment64.Filesz"></span>    Filesz  <a href="/pkg/builtin/#uint64">uint64</a>
<span id="Segment64.Maxprot"></span>    Maxprot <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Segment64.Prot"></span>    Prot    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Segment64.Nsect"></span>    Nsect   <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Segment64.Flag"></span>    Flag    <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="SegmentHeader">type</a> [SegmentHeader](https://golang.org/src/debug/macho/file.go?s=897:1103#L35)
A SegmentHeader is the header for a Mach-O 32-bit or 64-bit load segment command.


<pre>type SegmentHeader struct {
<span id="SegmentHeader.Cmd"></span>    Cmd     <a href="#LoadCmd">LoadCmd</a>
<span id="SegmentHeader.Len"></span>    Len     <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SegmentHeader.Name"></span>    Name    <a href="/pkg/builtin/#string">string</a>
<span id="SegmentHeader.Addr"></span>    Addr    <a href="/pkg/builtin/#uint64">uint64</a>
<span id="SegmentHeader.Memsz"></span>    Memsz   <a href="/pkg/builtin/#uint64">uint64</a>
<span id="SegmentHeader.Offset"></span>    Offset  <a href="/pkg/builtin/#uint64">uint64</a>
<span id="SegmentHeader.Filesz"></span>    Filesz  <a href="/pkg/builtin/#uint64">uint64</a>
<span id="SegmentHeader.Maxprot"></span>    Maxprot <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SegmentHeader.Prot"></span>    Prot    <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SegmentHeader.Nsect"></span>    Nsect   <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SegmentHeader.Flag"></span>    Flag    <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="Symbol">type</a> [Symbol](https://golang.org/src/debug/macho/file.go?s=3826:3916#L160)
A Symbol is a Mach-O 32-bit or 64-bit symbol table entry.


<pre>type Symbol struct {
<span id="Symbol.Name"></span>    Name  <a href="/pkg/builtin/#string">string</a>
<span id="Symbol.Type"></span>    Type  <a href="/pkg/builtin/#uint8">uint8</a>
<span id="Symbol.Sect"></span>    Sect  <a href="/pkg/builtin/#uint8">uint8</a>
<span id="Symbol.Desc"></span>    Desc  <a href="/pkg/builtin/#uint16">uint16</a>
<span id="Symbol.Value"></span>    Value <a href="/pkg/builtin/#uint64">uint64</a>
}
</pre>











## <a id="Symtab">type</a> [Symtab](https://golang.org/src/debug/macho/file.go?s=3446:3505#L140)
A Symtab represents a Mach-O symbol table command.


<pre>type Symtab struct {
    <a href="#LoadBytes">LoadBytes</a>
    <a href="#SymtabCmd">SymtabCmd</a>
<span id="Symtab.Syms"></span>    Syms []<a href="#Symbol">Symbol</a>
}
</pre>











## <a id="SymtabCmd">type</a> [SymtabCmd](https://golang.org/src/debug/macho/macho.go?s=3678:3802#L134)
A SymtabCmd is a Mach-O symbol table command.


<pre>type SymtabCmd struct {
<span id="SymtabCmd.Cmd"></span>    Cmd     <a href="#LoadCmd">LoadCmd</a>
<span id="SymtabCmd.Len"></span>    Len     <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SymtabCmd.Symoff"></span>    Symoff  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SymtabCmd.Nsyms"></span>    Nsyms   <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SymtabCmd.Stroff"></span>    Stroff  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SymtabCmd.Strsize"></span>    Strsize <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="Thread">type</a> [Thread](https://golang.org/src/debug/macho/macho.go?s=4752:4829#L185)
A Thread is a Mach-O thread state command.


<pre>type Thread struct {
<span id="Thread.Cmd"></span>    Cmd  <a href="#LoadCmd">LoadCmd</a>
<span id="Thread.Len"></span>    Len  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Thread.Type"></span>    Type <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Thread.Data"></span>    Data []<a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="Type">type</a> [Type](https://golang.org/src/debug/macho/macho.go?s=1070:1086#L30)
A Type is the Mach-O file type, e.g. an object file, executable, or dynamic library.


<pre>type Type <a href="/pkg/builtin/#uint32">uint32</a></pre>



<pre>const (
    <span id="TypeObj">TypeObj</span>    <a href="#Type">Type</a> = 1
    <span id="TypeExec">TypeExec</span>   <a href="#Type">Type</a> = 2
    <span id="TypeDylib">TypeDylib</span>  <a href="#Type">Type</a> = 6
    <span id="TypeBundle">TypeBundle</span> <a href="#Type">Type</a> = 8
)</pre>









### <a id="Type.GoString">func</a> (Type) [GoString](https://golang.org/src/debug/macho/macho.go?s=1420:1451#L47)
<pre>func (t <a href="#Type">Type</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Type.String">func</a> (Type) [String](https://golang.org/src/debug/macho/macho.go?s=1335:1364#L46)
<pre>func (t <a href="#Type">Type</a>) String() <a href="/pkg/builtin/#string">string</a></pre>







