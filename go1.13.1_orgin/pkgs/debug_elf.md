

# elf
`import "debug/elf"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package elf implements access to ELF object files.




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func R_INFO(sym, typ uint32) uint64](#R_INFO)
* [func R_INFO32(sym, typ uint32) uint32](#R_INFO32)
* [func R_SYM32(info uint32) uint32](#R_SYM32)
* [func R_SYM64(info uint64) uint32](#R_SYM64)
* [func R_TYPE32(info uint32) uint32](#R_TYPE32)
* [func R_TYPE64(info uint64) uint32](#R_TYPE64)
* [func ST_INFO(bind SymBind, typ SymType) uint8](#ST_INFO)
* [type Chdr32](#Chdr32)
* [type Chdr64](#Chdr64)
* [type Class](#Class)
  * [func (i Class) GoString() string](#Class.GoString)
  * [func (i Class) String() string](#Class.String)
* [type CompressionType](#CompressionType)
  * [func (i CompressionType) GoString() string](#CompressionType.GoString)
  * [func (i CompressionType) String() string](#CompressionType.String)
* [type Data](#Data)
  * [func (i Data) GoString() string](#Data.GoString)
  * [func (i Data) String() string](#Data.String)
* [type Dyn32](#Dyn32)
* [type Dyn64](#Dyn64)
* [type DynFlag](#DynFlag)
  * [func (i DynFlag) GoString() string](#DynFlag.GoString)
  * [func (i DynFlag) String() string](#DynFlag.String)
* [type DynTag](#DynTag)
  * [func (i DynTag) GoString() string](#DynTag.GoString)
  * [func (i DynTag) String() string](#DynTag.String)
* [type File](#File)
  * [func NewFile(r io.ReaderAt) (*File, error)](#NewFile)
  * [func Open(name string) (*File, error)](#Open)
  * [func (f *File) Close() error](#File.Close)
  * [func (f *File) DWARF() (*dwarf.Data, error)](#File.DWARF)
  * [func (f *File) DynString(tag DynTag) ([]string, error)](#File.DynString)
  * [func (f *File) DynamicSymbols() ([]Symbol, error)](#File.DynamicSymbols)
  * [func (f *File) ImportedLibraries() ([]string, error)](#File.ImportedLibraries)
  * [func (f *File) ImportedSymbols() ([]ImportedSymbol, error)](#File.ImportedSymbols)
  * [func (f *File) Section(name string) *Section](#File.Section)
  * [func (f *File) SectionByType(typ SectionType) *Section](#File.SectionByType)
  * [func (f *File) Symbols() ([]Symbol, error)](#File.Symbols)
* [type FileHeader](#FileHeader)
* [type FormatError](#FormatError)
  * [func (e *FormatError) Error() string](#FormatError.Error)
* [type Header32](#Header32)
* [type Header64](#Header64)
* [type ImportedSymbol](#ImportedSymbol)
* [type Machine](#Machine)
  * [func (i Machine) GoString() string](#Machine.GoString)
  * [func (i Machine) String() string](#Machine.String)
* [type NType](#NType)
  * [func (i NType) GoString() string](#NType.GoString)
  * [func (i NType) String() string](#NType.String)
* [type OSABI](#OSABI)
  * [func (i OSABI) GoString() string](#OSABI.GoString)
  * [func (i OSABI) String() string](#OSABI.String)
* [type Prog](#Prog)
  * [func (p *Prog) Open() io.ReadSeeker](#Prog.Open)
* [type Prog32](#Prog32)
* [type Prog64](#Prog64)
* [type ProgFlag](#ProgFlag)
  * [func (i ProgFlag) GoString() string](#ProgFlag.GoString)
  * [func (i ProgFlag) String() string](#ProgFlag.String)
* [type ProgHeader](#ProgHeader)
* [type ProgType](#ProgType)
  * [func (i ProgType) GoString() string](#ProgType.GoString)
  * [func (i ProgType) String() string](#ProgType.String)
* [type R_386](#R_386)
  * [func (i R_386) GoString() string](#R_386.GoString)
  * [func (i R_386) String() string](#R_386.String)
* [type R_390](#R_390)
  * [func (i R_390) GoString() string](#R_390.GoString)
  * [func (i R_390) String() string](#R_390.String)
* [type R_AARCH64](#R_AARCH64)
  * [func (i R_AARCH64) GoString() string](#R_AARCH64.GoString)
  * [func (i R_AARCH64) String() string](#R_AARCH64.String)
* [type R_ALPHA](#R_ALPHA)
  * [func (i R_ALPHA) GoString() string](#R_ALPHA.GoString)
  * [func (i R_ALPHA) String() string](#R_ALPHA.String)
* [type R_ARM](#R_ARM)
  * [func (i R_ARM) GoString() string](#R_ARM.GoString)
  * [func (i R_ARM) String() string](#R_ARM.String)
* [type R_MIPS](#R_MIPS)
  * [func (i R_MIPS) GoString() string](#R_MIPS.GoString)
  * [func (i R_MIPS) String() string](#R_MIPS.String)
* [type R_PPC](#R_PPC)
  * [func (i R_PPC) GoString() string](#R_PPC.GoString)
  * [func (i R_PPC) String() string](#R_PPC.String)
* [type R_PPC64](#R_PPC64)
  * [func (i R_PPC64) GoString() string](#R_PPC64.GoString)
  * [func (i R_PPC64) String() string](#R_PPC64.String)
* [type R_RISCV](#R_RISCV)
  * [func (i R_RISCV) GoString() string](#R_RISCV.GoString)
  * [func (i R_RISCV) String() string](#R_RISCV.String)
* [type R_SPARC](#R_SPARC)
  * [func (i R_SPARC) GoString() string](#R_SPARC.GoString)
  * [func (i R_SPARC) String() string](#R_SPARC.String)
* [type R_X86_64](#R_X86_64)
  * [func (i R_X86_64) GoString() string](#R_X86_64.GoString)
  * [func (i R_X86_64) String() string](#R_X86_64.String)
* [type Rel32](#Rel32)
* [type Rel64](#Rel64)
* [type Rela32](#Rela32)
* [type Rela64](#Rela64)
* [type Section](#Section)
  * [func (s *Section) Data() ([]byte, error)](#Section.Data)
  * [func (s *Section) Open() io.ReadSeeker](#Section.Open)
* [type Section32](#Section32)
* [type Section64](#Section64)
* [type SectionFlag](#SectionFlag)
  * [func (i SectionFlag) GoString() string](#SectionFlag.GoString)
  * [func (i SectionFlag) String() string](#SectionFlag.String)
* [type SectionHeader](#SectionHeader)
* [type SectionIndex](#SectionIndex)
  * [func (i SectionIndex) GoString() string](#SectionIndex.GoString)
  * [func (i SectionIndex) String() string](#SectionIndex.String)
* [type SectionType](#SectionType)
  * [func (i SectionType) GoString() string](#SectionType.GoString)
  * [func (i SectionType) String() string](#SectionType.String)
* [type Sym32](#Sym32)
* [type Sym64](#Sym64)
* [type SymBind](#SymBind)
  * [func ST_BIND(info uint8) SymBind](#ST_BIND)
  * [func (i SymBind) GoString() string](#SymBind.GoString)
  * [func (i SymBind) String() string](#SymBind.String)
* [type SymType](#SymType)
  * [func ST_TYPE(info uint8) SymType](#ST_TYPE)
  * [func (i SymType) GoString() string](#SymType.GoString)
  * [func (i SymType) String() string](#SymType.String)
* [type SymVis](#SymVis)
  * [func ST_VISIBILITY(other uint8) SymVis](#ST_VISIBILITY)
  * [func (i SymVis) GoString() string](#SymVis.GoString)
  * [func (i SymVis) String() string](#SymVis.String)
* [type Symbol](#Symbol)
* [type Type](#Type)
  * [func (i Type) GoString() string](#Type.GoString)
  * [func (i Type) String() string](#Type.String)
* [type Version](#Version)
  * [func (i Version) GoString() string](#Version.GoString)
  * [func (i Version) String() string](#Version.String)




#### <a id="pkg-files">Package files</a>
[elf.go](https://golang.org/src/debug/elf/elf.go) [file.go](https://golang.org/src/debug/elf/file.go) [reader.go](https://golang.org/src/debug/elf/reader.go) 


## <a id="pkg-constants">Constants</a>
Indexes into the Header.Ident array.


<pre>const (
    <span id="EI_CLASS">EI_CLASS</span>      = 4  <span class="comment">/* Class of machine. */</span>
    <span id="EI_DATA">EI_DATA</span>       = 5  <span class="comment">/* Data format. */</span>
    <span id="EI_VERSION">EI_VERSION</span>    = 6  <span class="comment">/* ELF format version. */</span>
    <span id="EI_OSABI">EI_OSABI</span>      = 7  <span class="comment">/* Operating system / ABI identification */</span>
    <span id="EI_ABIVERSION">EI_ABIVERSION</span> = 8  <span class="comment">/* ABI version */</span>
    <span id="EI_PAD">EI_PAD</span>        = 9  <span class="comment">/* Start of padding (per SVR4 ABI). */</span>
    <span id="EI_NIDENT">EI_NIDENT</span>     = 16 <span class="comment">/* Size of e_ident array. */</span>
)</pre>Magic number for the elf trampoline, chosen wisely to be an immediate value.


<pre>const <span id="ARM_MAGIC_TRAMP_NUMBER">ARM_MAGIC_TRAMP_NUMBER</span> = 0x5c000003</pre>Initial magic number for ELF files.


<pre>const <span id="ELFMAG">ELFMAG</span> = &#34;\177ELF&#34;</pre>
<pre>const <span id="Sym32Size">Sym32Size</span> = 16</pre>
<pre>const <span id="Sym64Size">Sym64Size</span> = 24</pre>

## <a id="pkg-variables">Variables</a>
ErrNoSymbols is returned by File.Symbols and File.DynamicSymbols
if there is no such section in the File.


<pre>var <span id="ErrNoSymbols">ErrNoSymbols</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;no symbol section&#34;)</pre>

## <a id="R_INFO">func</a> [R_INFO](https://golang.org/src/debug/elf/elf.go?s=112031:112066#L2919)
<pre>func R_INFO(sym, typ <a href="/pkg/builtin/#uint32">uint32</a>) <a href="/pkg/builtin/#uint64">uint64</a></pre>


## <a id="R_INFO32">func</a> [R_INFO32](https://golang.org/src/debug/elf/elf.go?s=108834:108871#L2817)
<pre>func R_INFO32(sym, typ <a href="/pkg/builtin/#uint32">uint32</a>) <a href="/pkg/builtin/#uint32">uint32</a></pre>


## <a id="R_SYM32">func</a> [R_SYM32](https://golang.org/src/debug/elf/elf.go?s=108714:108746#L2815)
<pre>func R_SYM32(info <a href="/pkg/builtin/#uint32">uint32</a>) <a href="/pkg/builtin/#uint32">uint32</a></pre>


## <a id="R_SYM64">func</a> [R_SYM64](https://golang.org/src/debug/elf/elf.go?s=111905:111937#L2917)
<pre>func R_SYM64(info <a href="/pkg/builtin/#uint64">uint64</a>) <a href="/pkg/builtin/#uint32">uint32</a></pre>


## <a id="R_TYPE32">func</a> [R_TYPE32](https://golang.org/src/debug/elf/elf.go?s=108773:108806#L2816)
<pre>func R_TYPE32(info <a href="/pkg/builtin/#uint32">uint32</a>) <a href="/pkg/builtin/#uint32">uint32</a></pre>


## <a id="R_TYPE64">func</a> [R_TYPE64](https://golang.org/src/debug/elf/elf.go?s=111971:112004#L2918)
<pre>func R_TYPE64(info <a href="/pkg/builtin/#uint64">uint64</a>) <a href="/pkg/builtin/#uint32">uint32</a></pre>


## <a id="ST_INFO">func</a> [ST_INFO](https://golang.org/src/debug/elf/elf.go?s=109168:109213#L2833)
<pre>func ST_INFO(bind <a href="#SymBind">SymBind</a>, typ <a href="#SymType">SymType</a>) <a href="/pkg/builtin/#uint8">uint8</a></pre>




## <a id="Chdr32">type</a> [Chdr32](https://golang.org/src/debug/elf/elf.go?s=108227:108303#L2792)
ELF32 Compression header.


<pre>type Chdr32 struct {
<span id="Chdr32.Type"></span>    Type      <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Chdr32.Size"></span>    Size      <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Chdr32.Addralign"></span>    Addralign <a href="/pkg/builtin/#uint32">uint32</a>
}
</pre>











## <a id="Chdr64">type</a> [Chdr64](https://golang.org/src/debug/elf/elf.go?s=111378:111488#L2893)
ELF64 Compression header.


<pre>type Chdr64 struct {
<span id="Chdr64.Type"></span>    Type <a href="/pkg/builtin/#uint32">uint32</a>

<span id="Chdr64.Size"></span>    Size      <a href="/pkg/builtin/#uint64">uint64</a>
<span id="Chdr64.Addralign"></span>    Addralign <a href="/pkg/builtin/#uint64">uint64</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











## <a id="Class">type</a> [Class](https://golang.org/src/debug/elf/elf.go?s=3527:3542#L74)
Class is found in Header.Ident[EI_CLASS] and Header.Class.


<pre>type Class <a href="/pkg/builtin/#byte">byte</a></pre>



<pre>const (
    <span id="ELFCLASSNONE">ELFCLASSNONE</span> <a href="#Class">Class</a> = 0 <span class="comment">/* Unknown class. */</span>
    <span id="ELFCLASS32">ELFCLASS32</span>   <a href="#Class">Class</a> = 1 <span class="comment">/* 32-bit architecture. */</span>
    <span id="ELFCLASS64">ELFCLASS64</span>   <a href="#Class">Class</a> = 2 <span class="comment">/* 64-bit architecture. */</span>
)</pre>









### <a id="Class.GoString">func</a> (Class) [GoString](https://golang.org/src/debug/elf/elf.go?s=3884:3916#L89)
<pre>func (i <a href="#Class">Class</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Class.String">func</a> (Class) [String](https://golang.org/src/debug/elf/elf.go?s=3797:3827#L88)
<pre>func (i <a href="#Class">Class</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="CompressionType">type</a> [CompressionType](https://golang.org/src/debug/elf/elf.go?s=29996:30020#L713)
Section compression type.


<pre>type CompressionType <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="COMPRESS_ZLIB">COMPRESS_ZLIB</span>   <a href="#CompressionType">CompressionType</a> = 1          <span class="comment">/* ZLIB compression. */</span>
    <span id="COMPRESS_LOOS">COMPRESS_LOOS</span>   <a href="#CompressionType">CompressionType</a> = 0x60000000 <span class="comment">/* First OS-specific. */</span>
    <span id="COMPRESS_HIOS">COMPRESS_HIOS</span>   <a href="#CompressionType">CompressionType</a> = 0x6fffffff <span class="comment">/* Last OS-specific. */</span>
    <span id="COMPRESS_LOPROC">COMPRESS_LOPROC</span> <a href="#CompressionType">CompressionType</a> = 0x70000000 <span class="comment">/* First processor-specific type. */</span>
    <span id="COMPRESS_HIPROC">COMPRESS_HIPROC</span> <a href="#CompressionType">CompressionType</a> = 0x7fffffff <span class="comment">/* Last processor-specific type. */</span>
)</pre>









### <a id="CompressionType.GoString">func</a> (CompressionType) [GoString](https://golang.org/src/debug/elf/elf.go?s=30706:30748#L732)
<pre>func (i <a href="#CompressionType">CompressionType</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="CompressionType.String">func</a> (CompressionType) [String](https://golang.org/src/debug/elf/elf.go?s=30603:30643#L731)
<pre>func (i <a href="#CompressionType">CompressionType</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Data">type</a> [Data](https://golang.org/src/debug/elf/elf.go?s=4030:4044#L92)
Data is found in Header.Ident[EI_DATA] and Header.Data.


<pre>type Data <a href="/pkg/builtin/#byte">byte</a></pre>



<pre>const (
    <span id="ELFDATANONE">ELFDATANONE</span> <a href="#Data">Data</a> = 0 <span class="comment">/* Unknown data format. */</span>
    <span id="ELFDATA2LSB">ELFDATA2LSB</span> <a href="#Data">Data</a> = 1 <span class="comment">/* 2&#39;s complement little-endian. */</span>
    <span id="ELFDATA2MSB">ELFDATA2MSB</span> <a href="#Data">Data</a> = 2 <span class="comment">/* 2&#39;s complement big-endian. */</span>
)</pre>









### <a id="Data.GoString">func</a> (Data) [GoString](https://golang.org/src/debug/elf/elf.go?s=4399:4430#L107)
<pre>func (i <a href="#Data">Data</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Data.String">func</a> (Data) [String](https://golang.org/src/debug/elf/elf.go?s=4314:4343#L106)
<pre>func (i <a href="#Data">Data</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Dyn32">type</a> [Dyn32](https://golang.org/src/debug/elf/elf.go?s=108104:108196#L2786)
ELF32 Dynamic structure. The ".dynamic" section contains an array of them.


<pre>type Dyn32 struct {
<span id="Dyn32.Tag"></span>    Tag <a href="/pkg/builtin/#int32">int32</a>  <span class="comment">/* Entry type. */</span>
<span id="Dyn32.Val"></span>    Val <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Integer/Address value. */</span>
}
</pre>











## <a id="Dyn64">type</a> [Dyn64](https://golang.org/src/debug/elf/elf.go?s=111256:111347#L2887)
ELF64 Dynamic structure. The ".dynamic" section contains an array of them.


<pre>type Dyn64 struct {
<span id="Dyn64.Tag"></span>    Tag <a href="/pkg/builtin/#int64">int64</a>  <span class="comment">/* Entry type. */</span>
<span id="Dyn64.Val"></span>    Val <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Integer/address value */</span>
}
</pre>











## <a id="DynFlag">type</a> [DynFlag](https://golang.org/src/debug/elf/elf.go?s=36914:36930#L889)
DT_FLAGS values.


<pre>type DynFlag <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="DF_ORIGIN">DF_ORIGIN</span> <a href="#DynFlag">DynFlag</a> = 0x0001 <span class="comment">/* Indicates that the object being loaded may
       make reference to the
       $ORIGIN substitution string */</span>
    <span id="DF_SYMBOLIC">DF_SYMBOLIC</span> <a href="#DynFlag">DynFlag</a> = 0x0002 <span class="comment">/* Indicates &#34;symbolic&#34; linking. */</span>
    <span id="DF_TEXTREL">DF_TEXTREL</span>  <a href="#DynFlag">DynFlag</a> = 0x0004 <span class="comment">/* Indicates there may be relocations in non-writable segments. */</span>
    <span id="DF_BIND_NOW">DF_BIND_NOW</span> <a href="#DynFlag">DynFlag</a> = 0x0008 <span class="comment">/* Indicates that the dynamic linker should
       process all relocations for the object
       containing this entry before transferring
       control to the program. */</span>
    <span id="DF_STATIC_TLS">DF_STATIC_TLS</span> <a href="#DynFlag">DynFlag</a> = 0x0010 <span class="comment">/* Indicates that the shared object or
       executable contains code using a static
       thread-local storage scheme. */</span>
)</pre>









### <a id="DynFlag.GoString">func</a> (DynFlag) [GoString](https://golang.org/src/debug/elf/elf.go?s=37835:37869#L915)
<pre>func (i <a href="#DynFlag">DynFlag</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="DynFlag.String">func</a> (DynFlag) [String](https://golang.org/src/debug/elf/elf.go?s=37748:37780#L914)
<pre>func (i <a href="#DynFlag">DynFlag</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="DynTag">type</a> [DynTag](https://golang.org/src/debug/elf/elf.go?s=32674:32689#L791)
Dyn.Tag


<pre>type DynTag <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="DT_NULL">DT_NULL</span>         <a href="#DynTag">DynTag</a> = 0  <span class="comment">/* Terminating entry. */</span>
    <span id="DT_NEEDED">DT_NEEDED</span>       <a href="#DynTag">DynTag</a> = 1  <span class="comment">/* String table offset of a needed shared library. */</span>
    <span id="DT_PLTRELSZ">DT_PLTRELSZ</span>     <a href="#DynTag">DynTag</a> = 2  <span class="comment">/* Total size in bytes of PLT relocations. */</span>
    <span id="DT_PLTGOT">DT_PLTGOT</span>       <a href="#DynTag">DynTag</a> = 3  <span class="comment">/* Processor-dependent address. */</span>
    <span id="DT_HASH">DT_HASH</span>         <a href="#DynTag">DynTag</a> = 4  <span class="comment">/* Address of symbol hash table. */</span>
    <span id="DT_STRTAB">DT_STRTAB</span>       <a href="#DynTag">DynTag</a> = 5  <span class="comment">/* Address of string table. */</span>
    <span id="DT_SYMTAB">DT_SYMTAB</span>       <a href="#DynTag">DynTag</a> = 6  <span class="comment">/* Address of symbol table. */</span>
    <span id="DT_RELA">DT_RELA</span>         <a href="#DynTag">DynTag</a> = 7  <span class="comment">/* Address of ElfNN_Rela relocations. */</span>
    <span id="DT_RELASZ">DT_RELASZ</span>       <a href="#DynTag">DynTag</a> = 8  <span class="comment">/* Total size of ElfNN_Rela relocations. */</span>
    <span id="DT_RELAENT">DT_RELAENT</span>      <a href="#DynTag">DynTag</a> = 9  <span class="comment">/* Size of each ElfNN_Rela relocation entry. */</span>
    <span id="DT_STRSZ">DT_STRSZ</span>        <a href="#DynTag">DynTag</a> = 10 <span class="comment">/* Size of string table. */</span>
    <span id="DT_SYMENT">DT_SYMENT</span>       <a href="#DynTag">DynTag</a> = 11 <span class="comment">/* Size of each symbol table entry. */</span>
    <span id="DT_INIT">DT_INIT</span>         <a href="#DynTag">DynTag</a> = 12 <span class="comment">/* Address of initialization function. */</span>
    <span id="DT_FINI">DT_FINI</span>         <a href="#DynTag">DynTag</a> = 13 <span class="comment">/* Address of finalization function. */</span>
    <span id="DT_SONAME">DT_SONAME</span>       <a href="#DynTag">DynTag</a> = 14 <span class="comment">/* String table offset of shared object name. */</span>
    <span id="DT_RPATH">DT_RPATH</span>        <a href="#DynTag">DynTag</a> = 15 <span class="comment">/* String table offset of library path. [sup] */</span>
    <span id="DT_SYMBOLIC">DT_SYMBOLIC</span>     <a href="#DynTag">DynTag</a> = 16 <span class="comment">/* Indicates &#34;symbolic&#34; linking. [sup] */</span>
    <span id="DT_REL">DT_REL</span>          <a href="#DynTag">DynTag</a> = 17 <span class="comment">/* Address of ElfNN_Rel relocations. */</span>
    <span id="DT_RELSZ">DT_RELSZ</span>        <a href="#DynTag">DynTag</a> = 18 <span class="comment">/* Total size of ElfNN_Rel relocations. */</span>
    <span id="DT_RELENT">DT_RELENT</span>       <a href="#DynTag">DynTag</a> = 19 <span class="comment">/* Size of each ElfNN_Rel relocation. */</span>
    <span id="DT_PLTREL">DT_PLTREL</span>       <a href="#DynTag">DynTag</a> = 20 <span class="comment">/* Type of relocation used for PLT. */</span>
    <span id="DT_DEBUG">DT_DEBUG</span>        <a href="#DynTag">DynTag</a> = 21 <span class="comment">/* Reserved (not used). */</span>
    <span id="DT_TEXTREL">DT_TEXTREL</span>      <a href="#DynTag">DynTag</a> = 22 <span class="comment">/* Indicates there may be relocations in non-writable segments. [sup] */</span>
    <span id="DT_JMPREL">DT_JMPREL</span>       <a href="#DynTag">DynTag</a> = 23 <span class="comment">/* Address of PLT relocations. */</span>
    <span id="DT_BIND_NOW">DT_BIND_NOW</span>     <a href="#DynTag">DynTag</a> = 24 <span class="comment">/* [sup] */</span>
    <span id="DT_INIT_ARRAY">DT_INIT_ARRAY</span>   <a href="#DynTag">DynTag</a> = 25 <span class="comment">/* Address of the array of pointers to initialization functions */</span>
    <span id="DT_FINI_ARRAY">DT_FINI_ARRAY</span>   <a href="#DynTag">DynTag</a> = 26 <span class="comment">/* Address of the array of pointers to termination functions */</span>
    <span id="DT_INIT_ARRAYSZ">DT_INIT_ARRAYSZ</span> <a href="#DynTag">DynTag</a> = 27 <span class="comment">/* Size in bytes of the array of initialization functions. */</span>
    <span id="DT_FINI_ARRAYSZ">DT_FINI_ARRAYSZ</span> <a href="#DynTag">DynTag</a> = 28 <span class="comment">/* Size in bytes of the array of termination functions. */</span>
    <span id="DT_RUNPATH">DT_RUNPATH</span>      <a href="#DynTag">DynTag</a> = 29 <span class="comment">/* String table offset of a null-terminated library search path string. */</span>
    <span id="DT_FLAGS">DT_FLAGS</span>        <a href="#DynTag">DynTag</a> = 30 <span class="comment">/* Object specific flag values. */</span>
    <span id="DT_ENCODING">DT_ENCODING</span>     <a href="#DynTag">DynTag</a> = 32 <span class="comment">/* Values greater than or equal to DT_ENCODING
       and less than DT_LOOS follow the rules for
       the interpretation of the d_un union
       as follows: even == &#39;d_ptr&#39;, even == &#39;d_val&#39;
       or none */</span>
    <span id="DT_PREINIT_ARRAY">DT_PREINIT_ARRAY</span>   <a href="#DynTag">DynTag</a> = 32         <span class="comment">/* Address of the array of pointers to pre-initialization functions. */</span>
    <span id="DT_PREINIT_ARRAYSZ">DT_PREINIT_ARRAYSZ</span> <a href="#DynTag">DynTag</a> = 33         <span class="comment">/* Size in bytes of the array of pre-initialization functions. */</span>
    <span id="DT_LOOS">DT_LOOS</span>            <a href="#DynTag">DynTag</a> = 0x6000000d <span class="comment">/* First OS-specific */</span>
    <span id="DT_HIOS">DT_HIOS</span>            <a href="#DynTag">DynTag</a> = 0x6ffff000 <span class="comment">/* Last OS-specific */</span>
    <span id="DT_VERSYM">DT_VERSYM</span>          <a href="#DynTag">DynTag</a> = 0x6ffffff0
    <span id="DT_VERNEED">DT_VERNEED</span>         <a href="#DynTag">DynTag</a> = 0x6ffffffe
    <span id="DT_VERNEEDNUM">DT_VERNEEDNUM</span>      <a href="#DynTag">DynTag</a> = 0x6fffffff
    <span id="DT_LOPROC">DT_LOPROC</span>          <a href="#DynTag">DynTag</a> = 0x70000000 <span class="comment">/* First processor-specific type. */</span>
    <span id="DT_HIPROC">DT_HIPROC</span>          <a href="#DynTag">DynTag</a> = 0x7fffffff <span class="comment">/* Last processor-specific type. */</span>
)</pre>









### <a id="DynTag.GoString">func</a> (DynTag) [GoString](https://golang.org/src/debug/elf/elf.go?s=36809:36842#L886)
<pre>func (i <a href="#DynTag">DynTag</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="DynTag.String">func</a> (DynTag) [String](https://golang.org/src/debug/elf/elf.go?s=36724:36755#L885)
<pre>func (i <a href="#DynTag">DynTag</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="File">type</a> [File](https://golang.org/src/debug/elf/file.go?s=994:1127#L41)
A File represents an open ELF file.


<pre>type File struct {
    <a href="#FileHeader">FileHeader</a>
<span id="File.Sections"></span>    Sections []*<a href="#Section">Section</a>
<span id="File.Progs"></span>    Progs    []*<a href="#Prog">Prog</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewFile">func</a> [NewFile](https://golang.org/src/debug/elf/file.go?s=5710:5752#L230)
<pre>func NewFile(r <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReaderAt">ReaderAt</a>) (*<a href="#File">File</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
NewFile creates a new File for accessing an ELF binary in an underlying reader.
The ELF binary is expected to start at position 0 in the ReaderAt.




### <a id="Open">func</a> [Open](https://golang.org/src/debug/elf/file.go?s=4850:4887#L191)
<pre>func Open(name <a href="/pkg/builtin/#string">string</a>) (*<a href="#File">File</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Open opens the named file using os.Open and prepares it for use as an ELF binary.






### <a id="File.Close">func</a> (\*File) [Close](https://golang.org/src/debug/elf/file.go?s=5177:5205#L208)
<pre>func (f *<a href="#File">File</a>) Close() <a href="/pkg/builtin/#error">error</a></pre>
Close closes the File.
If the File was created using NewFile directly instead of Open,
Close has no effect.




### <a id="File.DWARF">func</a> (\*File) [DWARF](https://golang.org/src/debug/elf/file.go?s=27155:27198#L1112)
<pre>func (f *<a href="#File">File</a>) DWARF() (*<a href="/pkg/debug/dwarf/">dwarf</a>.<a href="/pkg/debug/dwarf/#Data">Data</a>, <a href="/pkg/builtin/#error">error</a>)</pre>



### <a id="File.DynString">func</a> (\*File) [DynString](https://golang.org/src/debug/elf/file.go?s=33526:33580#L1380)
<pre>func (f *<a href="#File">File</a>) DynString(tag <a href="#DynTag">DynTag</a>) ([]<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DynString returns the strings listed for the given tag in the file's dynamic
section.

The tag must be one that takes string values: DT_NEEDED, DT_SONAME, DT_RPATH, or
DT_RUNPATH.




### <a id="File.DynamicSymbols">func</a> (\*File) [DynamicSymbols](https://golang.org/src/debug/elf/file.go?s=30133:30182#L1232)
<pre>func (f *<a href="#File">File</a>) DynamicSymbols() ([]<a href="#Symbol">Symbol</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
DynamicSymbols returns the dynamic symbol table for f. The symbols
will be listed in the order they appear in f.

If f has a symbol version table, the returned Symbols will have
initialized Version and Library fields.

For compatibility with Symbols, DynamicSymbols omits the null symbol at index 0.
After retrieving the symbols as symtab, an externally supplied index x
corresponds to symtab[x-1], not symtab[x].




### <a id="File.ImportedLibraries">func</a> (\*File) [ImportedLibraries](https://golang.org/src/debug/elf/file.go?s=33243:33295#L1371)
<pre>func (f *<a href="#File">File</a>) ImportedLibraries() ([]<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ImportedLibraries returns the names of all libraries
referred to by the binary f that are expected to be
linked with the binary at dynamic link time.




### <a id="File.ImportedSymbols">func</a> (\*File) [ImportedSymbols](https://golang.org/src/debug/elf/file.go?s=30673:30731#L1255)
<pre>func (f *<a href="#File">File</a>) ImportedSymbols() ([]<a href="#ImportedSymbol">ImportedSymbol</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ImportedSymbols returns the names of all symbols
referred to by the binary f that are expected to be
satisfied by other libraries at dynamic load time.
It does not return weak symbols.




### <a id="File.Section">func</a> (\*File) [Section](https://golang.org/src/debug/elf/file.go?s=14737:14781#L581)
<pre>func (f *<a href="#File">File</a>) Section(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Section">Section</a></pre>
Section returns a section with the given name, or nil if no such
section exists.




### <a id="File.SectionByType">func</a> (\*File) [SectionByType](https://golang.org/src/debug/elf/file.go?s=5413:5467#L219)
<pre>func (f *<a href="#File">File</a>) SectionByType(typ <a href="#SectionType">SectionType</a>) *<a href="#Section">Section</a></pre>
SectionByType returns the first section in f with the
given type, or nil if there is no such section.




### <a id="File.Symbols">func</a> (\*File) [Symbols](https://golang.org/src/debug/elf/file.go?s=29588:29630#L1218)
<pre>func (f *<a href="#File">File</a>) Symbols() ([]<a href="#Symbol">Symbol</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Symbols returns the symbol table for f. The symbols will be listed in the order
they appear in f.

For compatibility with Go 1.0, Symbols omits the null symbol at index 0.
After retrieving the symbols as symtab, an externally supplied index x
corresponds to symtab[x-1], not symtab[x].




## <a id="FileHeader">type</a> [FileHeader](https://golang.org/src/debug/elf/file.go?s=751:953#L28)
A FileHeader represents an ELF file header.


<pre>type FileHeader struct {
<span id="FileHeader.Class"></span>    Class      <a href="#Class">Class</a>
<span id="FileHeader.Data"></span>    Data       <a href="#Data">Data</a>
<span id="FileHeader.Version"></span>    Version    <a href="#Version">Version</a>
<span id="FileHeader.OSABI"></span>    OSABI      <a href="#OSABI">OSABI</a>
<span id="FileHeader.ABIVersion"></span>    ABIVersion <a href="/pkg/builtin/#uint8">uint8</a>
<span id="FileHeader.ByteOrder"></span>    ByteOrder  <a href="/pkg/encoding/binary/">binary</a>.<a href="/pkg/encoding/binary/#ByteOrder">ByteOrder</a>
<span id="FileHeader.Type"></span>    Type       <a href="#Type">Type</a>
<span id="FileHeader.Machine"></span>    Machine    <a href="#Machine">Machine</a>
<span id="FileHeader.Entry"></span>    Entry      <a href="/pkg/builtin/#uint64">uint64</a>
}
</pre>











## <a id="FormatError">type</a> [FormatError](https://golang.org/src/debug/elf/file.go?s=4516:4583#L175)

<pre>type FormatError struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="FormatError.Error">func</a> (\*FormatError) [Error](https://golang.org/src/debug/elf/file.go?s=4585:4621#L181)
<pre>func (e *<a href="#FormatError">FormatError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Header32">type</a> [Header32](https://golang.org/src/debug/elf/elf.go?s=106214:107059#L2742)
ELF32 File header.


<pre>type Header32 struct {
<span id="Header32.Ident"></span>    Ident     [<a href="#EI_NIDENT">EI_NIDENT</a>]<a href="/pkg/builtin/#byte">byte</a> <span class="comment">/* File identification. */</span>
<span id="Header32.Type"></span>    Type      <a href="/pkg/builtin/#uint16">uint16</a>          <span class="comment">/* File type. */</span>
<span id="Header32.Machine"></span>    Machine   <a href="/pkg/builtin/#uint16">uint16</a>          <span class="comment">/* Machine architecture. */</span>
<span id="Header32.Version"></span>    Version   <a href="/pkg/builtin/#uint32">uint32</a>          <span class="comment">/* ELF format version. */</span>
<span id="Header32.Entry"></span>    Entry     <a href="/pkg/builtin/#uint32">uint32</a>          <span class="comment">/* Entry point. */</span>
<span id="Header32.Phoff"></span>    Phoff     <a href="/pkg/builtin/#uint32">uint32</a>          <span class="comment">/* Program header file offset. */</span>
<span id="Header32.Shoff"></span>    Shoff     <a href="/pkg/builtin/#uint32">uint32</a>          <span class="comment">/* Section header file offset. */</span>
<span id="Header32.Flags"></span>    Flags     <a href="/pkg/builtin/#uint32">uint32</a>          <span class="comment">/* Architecture-specific flags. */</span>
<span id="Header32.Ehsize"></span>    Ehsize    <a href="/pkg/builtin/#uint16">uint16</a>          <span class="comment">/* Size of ELF header in bytes. */</span>
<span id="Header32.Phentsize"></span>    Phentsize <a href="/pkg/builtin/#uint16">uint16</a>          <span class="comment">/* Size of program header entry. */</span>
<span id="Header32.Phnum"></span>    Phnum     <a href="/pkg/builtin/#uint16">uint16</a>          <span class="comment">/* Number of program header entries. */</span>
<span id="Header32.Shentsize"></span>    Shentsize <a href="/pkg/builtin/#uint16">uint16</a>          <span class="comment">/* Size of section header entry. */</span>
<span id="Header32.Shnum"></span>    Shnum     <a href="/pkg/builtin/#uint16">uint16</a>          <span class="comment">/* Number of section header entries. */</span>
<span id="Header32.Shstrndx"></span>    Shstrndx  <a href="/pkg/builtin/#uint16">uint16</a>          <span class="comment">/* Section name strings section. */</span>
}
</pre>











## <a id="Header64">type</a> [Header64](https://golang.org/src/debug/elf/elf.go?s=109366:110211#L2843)
ELF64 file header.


<pre>type Header64 struct {
<span id="Header64.Ident"></span>    Ident     [<a href="#EI_NIDENT">EI_NIDENT</a>]<a href="/pkg/builtin/#byte">byte</a> <span class="comment">/* File identification. */</span>
<span id="Header64.Type"></span>    Type      <a href="/pkg/builtin/#uint16">uint16</a>          <span class="comment">/* File type. */</span>
<span id="Header64.Machine"></span>    Machine   <a href="/pkg/builtin/#uint16">uint16</a>          <span class="comment">/* Machine architecture. */</span>
<span id="Header64.Version"></span>    Version   <a href="/pkg/builtin/#uint32">uint32</a>          <span class="comment">/* ELF format version. */</span>
<span id="Header64.Entry"></span>    Entry     <a href="/pkg/builtin/#uint64">uint64</a>          <span class="comment">/* Entry point. */</span>
<span id="Header64.Phoff"></span>    Phoff     <a href="/pkg/builtin/#uint64">uint64</a>          <span class="comment">/* Program header file offset. */</span>
<span id="Header64.Shoff"></span>    Shoff     <a href="/pkg/builtin/#uint64">uint64</a>          <span class="comment">/* Section header file offset. */</span>
<span id="Header64.Flags"></span>    Flags     <a href="/pkg/builtin/#uint32">uint32</a>          <span class="comment">/* Architecture-specific flags. */</span>
<span id="Header64.Ehsize"></span>    Ehsize    <a href="/pkg/builtin/#uint16">uint16</a>          <span class="comment">/* Size of ELF header in bytes. */</span>
<span id="Header64.Phentsize"></span>    Phentsize <a href="/pkg/builtin/#uint16">uint16</a>          <span class="comment">/* Size of program header entry. */</span>
<span id="Header64.Phnum"></span>    Phnum     <a href="/pkg/builtin/#uint16">uint16</a>          <span class="comment">/* Number of program header entries. */</span>
<span id="Header64.Shentsize"></span>    Shentsize <a href="/pkg/builtin/#uint16">uint16</a>          <span class="comment">/* Size of section header entry. */</span>
<span id="Header64.Shnum"></span>    Shnum     <a href="/pkg/builtin/#uint16">uint16</a>          <span class="comment">/* Number of section header entries. */</span>
<span id="Header64.Shstrndx"></span>    Shstrndx  <a href="/pkg/builtin/#uint16">uint16</a>          <span class="comment">/* Section name strings section. */</span>
}
</pre>











## <a id="ImportedSymbol">type</a> [ImportedSymbol](https://golang.org/src/debug/elf/file.go?s=30396:30474#L1245)

<pre>type ImportedSymbol struct {
<span id="ImportedSymbol.Name"></span>    Name    <a href="/pkg/builtin/#string">string</a>
<span id="ImportedSymbol.Version"></span>    Version <a href="/pkg/builtin/#string">string</a>
<span id="ImportedSymbol.Library"></span>    Library <a href="/pkg/builtin/#string">string</a>
}
</pre>











## <a id="Machine">type</a> [Machine](https://golang.org/src/debug/elf/elf.go?s=7290:7309#L192)
Machine is found in Header.Machine.


<pre>type Machine <a href="/pkg/builtin/#uint16">uint16</a></pre>



<pre>const (
    <span id="EM_NONE">EM_NONE</span>          <a href="#Machine">Machine</a> = 0   <span class="comment">/* Unknown machine. */</span>
    <span id="EM_M32">EM_M32</span>           <a href="#Machine">Machine</a> = 1   <span class="comment">/* AT&amp;T WE32100. */</span>
    <span id="EM_SPARC">EM_SPARC</span>         <a href="#Machine">Machine</a> = 2   <span class="comment">/* Sun SPARC. */</span>
    <span id="EM_386">EM_386</span>           <a href="#Machine">Machine</a> = 3   <span class="comment">/* Intel i386. */</span>
    <span id="EM_68K">EM_68K</span>           <a href="#Machine">Machine</a> = 4   <span class="comment">/* Motorola 68000. */</span>
    <span id="EM_88K">EM_88K</span>           <a href="#Machine">Machine</a> = 5   <span class="comment">/* Motorola 88000. */</span>
    <span id="EM_860">EM_860</span>           <a href="#Machine">Machine</a> = 7   <span class="comment">/* Intel i860. */</span>
    <span id="EM_MIPS">EM_MIPS</span>          <a href="#Machine">Machine</a> = 8   <span class="comment">/* MIPS R3000 Big-Endian only. */</span>
    <span id="EM_S370">EM_S370</span>          <a href="#Machine">Machine</a> = 9   <span class="comment">/* IBM System/370. */</span>
    <span id="EM_MIPS_RS3_LE">EM_MIPS_RS3_LE</span>   <a href="#Machine">Machine</a> = 10  <span class="comment">/* MIPS R3000 Little-Endian. */</span>
    <span id="EM_PARISC">EM_PARISC</span>        <a href="#Machine">Machine</a> = 15  <span class="comment">/* HP PA-RISC. */</span>
    <span id="EM_VPP500">EM_VPP500</span>        <a href="#Machine">Machine</a> = 17  <span class="comment">/* Fujitsu VPP500. */</span>
    <span id="EM_SPARC32PLUS">EM_SPARC32PLUS</span>   <a href="#Machine">Machine</a> = 18  <span class="comment">/* SPARC v8plus. */</span>
    <span id="EM_960">EM_960</span>           <a href="#Machine">Machine</a> = 19  <span class="comment">/* Intel 80960. */</span>
    <span id="EM_PPC">EM_PPC</span>           <a href="#Machine">Machine</a> = 20  <span class="comment">/* PowerPC 32-bit. */</span>
    <span id="EM_PPC64">EM_PPC64</span>         <a href="#Machine">Machine</a> = 21  <span class="comment">/* PowerPC 64-bit. */</span>
    <span id="EM_S390">EM_S390</span>          <a href="#Machine">Machine</a> = 22  <span class="comment">/* IBM System/390. */</span>
    <span id="EM_V800">EM_V800</span>          <a href="#Machine">Machine</a> = 36  <span class="comment">/* NEC V800. */</span>
    <span id="EM_FR20">EM_FR20</span>          <a href="#Machine">Machine</a> = 37  <span class="comment">/* Fujitsu FR20. */</span>
    <span id="EM_RH32">EM_RH32</span>          <a href="#Machine">Machine</a> = 38  <span class="comment">/* TRW RH-32. */</span>
    <span id="EM_RCE">EM_RCE</span>           <a href="#Machine">Machine</a> = 39  <span class="comment">/* Motorola RCE. */</span>
    <span id="EM_ARM">EM_ARM</span>           <a href="#Machine">Machine</a> = 40  <span class="comment">/* ARM. */</span>
    <span id="EM_SH">EM_SH</span>            <a href="#Machine">Machine</a> = 42  <span class="comment">/* Hitachi SH. */</span>
    <span id="EM_SPARCV9">EM_SPARCV9</span>       <a href="#Machine">Machine</a> = 43  <span class="comment">/* SPARC v9 64-bit. */</span>
    <span id="EM_TRICORE">EM_TRICORE</span>       <a href="#Machine">Machine</a> = 44  <span class="comment">/* Siemens TriCore embedded processor. */</span>
    <span id="EM_ARC">EM_ARC</span>           <a href="#Machine">Machine</a> = 45  <span class="comment">/* Argonaut RISC Core. */</span>
    <span id="EM_H8_300">EM_H8_300</span>        <a href="#Machine">Machine</a> = 46  <span class="comment">/* Hitachi H8/300. */</span>
    <span id="EM_H8_300H">EM_H8_300H</span>       <a href="#Machine">Machine</a> = 47  <span class="comment">/* Hitachi H8/300H. */</span>
    <span id="EM_H8S">EM_H8S</span>           <a href="#Machine">Machine</a> = 48  <span class="comment">/* Hitachi H8S. */</span>
    <span id="EM_H8_500">EM_H8_500</span>        <a href="#Machine">Machine</a> = 49  <span class="comment">/* Hitachi H8/500. */</span>
    <span id="EM_IA_64">EM_IA_64</span>         <a href="#Machine">Machine</a> = 50  <span class="comment">/* Intel IA-64 Processor. */</span>
    <span id="EM_MIPS_X">EM_MIPS_X</span>        <a href="#Machine">Machine</a> = 51  <span class="comment">/* Stanford MIPS-X. */</span>
    <span id="EM_COLDFIRE">EM_COLDFIRE</span>      <a href="#Machine">Machine</a> = 52  <span class="comment">/* Motorola ColdFire. */</span>
    <span id="EM_68HC12">EM_68HC12</span>        <a href="#Machine">Machine</a> = 53  <span class="comment">/* Motorola M68HC12. */</span>
    <span id="EM_MMA">EM_MMA</span>           <a href="#Machine">Machine</a> = 54  <span class="comment">/* Fujitsu MMA. */</span>
    <span id="EM_PCP">EM_PCP</span>           <a href="#Machine">Machine</a> = 55  <span class="comment">/* Siemens PCP. */</span>
    <span id="EM_NCPU">EM_NCPU</span>          <a href="#Machine">Machine</a> = 56  <span class="comment">/* Sony nCPU. */</span>
    <span id="EM_NDR1">EM_NDR1</span>          <a href="#Machine">Machine</a> = 57  <span class="comment">/* Denso NDR1 microprocessor. */</span>
    <span id="EM_STARCORE">EM_STARCORE</span>      <a href="#Machine">Machine</a> = 58  <span class="comment">/* Motorola Star*Core processor. */</span>
    <span id="EM_ME16">EM_ME16</span>          <a href="#Machine">Machine</a> = 59  <span class="comment">/* Toyota ME16 processor. */</span>
    <span id="EM_ST100">EM_ST100</span>         <a href="#Machine">Machine</a> = 60  <span class="comment">/* STMicroelectronics ST100 processor. */</span>
    <span id="EM_TINYJ">EM_TINYJ</span>         <a href="#Machine">Machine</a> = 61  <span class="comment">/* Advanced Logic Corp. TinyJ processor. */</span>
    <span id="EM_X86_64">EM_X86_64</span>        <a href="#Machine">Machine</a> = 62  <span class="comment">/* Advanced Micro Devices x86-64 */</span>
    <span id="EM_PDSP">EM_PDSP</span>          <a href="#Machine">Machine</a> = 63  <span class="comment">/* Sony DSP Processor */</span>
    <span id="EM_PDP10">EM_PDP10</span>         <a href="#Machine">Machine</a> = 64  <span class="comment">/* Digital Equipment Corp. PDP-10 */</span>
    <span id="EM_PDP11">EM_PDP11</span>         <a href="#Machine">Machine</a> = 65  <span class="comment">/* Digital Equipment Corp. PDP-11 */</span>
    <span id="EM_FX66">EM_FX66</span>          <a href="#Machine">Machine</a> = 66  <span class="comment">/* Siemens FX66 microcontroller */</span>
    <span id="EM_ST9PLUS">EM_ST9PLUS</span>       <a href="#Machine">Machine</a> = 67  <span class="comment">/* STMicroelectronics ST9+ 8/16 bit microcontroller */</span>
    <span id="EM_ST7">EM_ST7</span>           <a href="#Machine">Machine</a> = 68  <span class="comment">/* STMicroelectronics ST7 8-bit microcontroller */</span>
    <span id="EM_68HC16">EM_68HC16</span>        <a href="#Machine">Machine</a> = 69  <span class="comment">/* Motorola MC68HC16 Microcontroller */</span>
    <span id="EM_68HC11">EM_68HC11</span>        <a href="#Machine">Machine</a> = 70  <span class="comment">/* Motorola MC68HC11 Microcontroller */</span>
    <span id="EM_68HC08">EM_68HC08</span>        <a href="#Machine">Machine</a> = 71  <span class="comment">/* Motorola MC68HC08 Microcontroller */</span>
    <span id="EM_68HC05">EM_68HC05</span>        <a href="#Machine">Machine</a> = 72  <span class="comment">/* Motorola MC68HC05 Microcontroller */</span>
    <span id="EM_SVX">EM_SVX</span>           <a href="#Machine">Machine</a> = 73  <span class="comment">/* Silicon Graphics SVx */</span>
    <span id="EM_ST19">EM_ST19</span>          <a href="#Machine">Machine</a> = 74  <span class="comment">/* STMicroelectronics ST19 8-bit microcontroller */</span>
    <span id="EM_VAX">EM_VAX</span>           <a href="#Machine">Machine</a> = 75  <span class="comment">/* Digital VAX */</span>
    <span id="EM_CRIS">EM_CRIS</span>          <a href="#Machine">Machine</a> = 76  <span class="comment">/* Axis Communications 32-bit embedded processor */</span>
    <span id="EM_JAVELIN">EM_JAVELIN</span>       <a href="#Machine">Machine</a> = 77  <span class="comment">/* Infineon Technologies 32-bit embedded processor */</span>
    <span id="EM_FIREPATH">EM_FIREPATH</span>      <a href="#Machine">Machine</a> = 78  <span class="comment">/* Element 14 64-bit DSP Processor */</span>
    <span id="EM_ZSP">EM_ZSP</span>           <a href="#Machine">Machine</a> = 79  <span class="comment">/* LSI Logic 16-bit DSP Processor */</span>
    <span id="EM_MMIX">EM_MMIX</span>          <a href="#Machine">Machine</a> = 80  <span class="comment">/* Donald Knuth&#39;s educational 64-bit processor */</span>
    <span id="EM_HUANY">EM_HUANY</span>         <a href="#Machine">Machine</a> = 81  <span class="comment">/* Harvard University machine-independent object files */</span>
    <span id="EM_PRISM">EM_PRISM</span>         <a href="#Machine">Machine</a> = 82  <span class="comment">/* SiTera Prism */</span>
    <span id="EM_AVR">EM_AVR</span>           <a href="#Machine">Machine</a> = 83  <span class="comment">/* Atmel AVR 8-bit microcontroller */</span>
    <span id="EM_FR30">EM_FR30</span>          <a href="#Machine">Machine</a> = 84  <span class="comment">/* Fujitsu FR30 */</span>
    <span id="EM_D10V">EM_D10V</span>          <a href="#Machine">Machine</a> = 85  <span class="comment">/* Mitsubishi D10V */</span>
    <span id="EM_D30V">EM_D30V</span>          <a href="#Machine">Machine</a> = 86  <span class="comment">/* Mitsubishi D30V */</span>
    <span id="EM_V850">EM_V850</span>          <a href="#Machine">Machine</a> = 87  <span class="comment">/* NEC v850 */</span>
    <span id="EM_M32R">EM_M32R</span>          <a href="#Machine">Machine</a> = 88  <span class="comment">/* Mitsubishi M32R */</span>
    <span id="EM_MN10300">EM_MN10300</span>       <a href="#Machine">Machine</a> = 89  <span class="comment">/* Matsushita MN10300 */</span>
    <span id="EM_MN10200">EM_MN10200</span>       <a href="#Machine">Machine</a> = 90  <span class="comment">/* Matsushita MN10200 */</span>
    <span id="EM_PJ">EM_PJ</span>            <a href="#Machine">Machine</a> = 91  <span class="comment">/* picoJava */</span>
    <span id="EM_OPENRISC">EM_OPENRISC</span>      <a href="#Machine">Machine</a> = 92  <span class="comment">/* OpenRISC 32-bit embedded processor */</span>
    <span id="EM_ARC_COMPACT">EM_ARC_COMPACT</span>   <a href="#Machine">Machine</a> = 93  <span class="comment">/* ARC International ARCompact processor (old spelling/synonym: EM_ARC_A5) */</span>
    <span id="EM_XTENSA">EM_XTENSA</span>        <a href="#Machine">Machine</a> = 94  <span class="comment">/* Tensilica Xtensa Architecture */</span>
    <span id="EM_VIDEOCORE">EM_VIDEOCORE</span>     <a href="#Machine">Machine</a> = 95  <span class="comment">/* Alphamosaic VideoCore processor */</span>
    <span id="EM_TMM_GPP">EM_TMM_GPP</span>       <a href="#Machine">Machine</a> = 96  <span class="comment">/* Thompson Multimedia General Purpose Processor */</span>
    <span id="EM_NS32K">EM_NS32K</span>         <a href="#Machine">Machine</a> = 97  <span class="comment">/* National Semiconductor 32000 series */</span>
    <span id="EM_TPC">EM_TPC</span>           <a href="#Machine">Machine</a> = 98  <span class="comment">/* Tenor Network TPC processor */</span>
    <span id="EM_SNP1K">EM_SNP1K</span>         <a href="#Machine">Machine</a> = 99  <span class="comment">/* Trebia SNP 1000 processor */</span>
    <span id="EM_ST200">EM_ST200</span>         <a href="#Machine">Machine</a> = 100 <span class="comment">/* STMicroelectronics (www.st.com) ST200 microcontroller */</span>
    <span id="EM_IP2K">EM_IP2K</span>          <a href="#Machine">Machine</a> = 101 <span class="comment">/* Ubicom IP2xxx microcontroller family */</span>
    <span id="EM_MAX">EM_MAX</span>           <a href="#Machine">Machine</a> = 102 <span class="comment">/* MAX Processor */</span>
    <span id="EM_CR">EM_CR</span>            <a href="#Machine">Machine</a> = 103 <span class="comment">/* National Semiconductor CompactRISC microprocessor */</span>
    <span id="EM_F2MC16">EM_F2MC16</span>        <a href="#Machine">Machine</a> = 104 <span class="comment">/* Fujitsu F2MC16 */</span>
    <span id="EM_MSP430">EM_MSP430</span>        <a href="#Machine">Machine</a> = 105 <span class="comment">/* Texas Instruments embedded microcontroller msp430 */</span>
    <span id="EM_BLACKFIN">EM_BLACKFIN</span>      <a href="#Machine">Machine</a> = 106 <span class="comment">/* Analog Devices Blackfin (DSP) processor */</span>
    <span id="EM_SE_C33">EM_SE_C33</span>        <a href="#Machine">Machine</a> = 107 <span class="comment">/* S1C33 Family of Seiko Epson processors */</span>
    <span id="EM_SEP">EM_SEP</span>           <a href="#Machine">Machine</a> = 108 <span class="comment">/* Sharp embedded microprocessor */</span>
    <span id="EM_ARCA">EM_ARCA</span>          <a href="#Machine">Machine</a> = 109 <span class="comment">/* Arca RISC Microprocessor */</span>
    <span id="EM_UNICORE">EM_UNICORE</span>       <a href="#Machine">Machine</a> = 110 <span class="comment">/* Microprocessor series from PKU-Unity Ltd. and MPRC of Peking University */</span>
    <span id="EM_EXCESS">EM_EXCESS</span>        <a href="#Machine">Machine</a> = 111 <span class="comment">/* eXcess: 16/32/64-bit configurable embedded CPU */</span>
    <span id="EM_DXP">EM_DXP</span>           <a href="#Machine">Machine</a> = 112 <span class="comment">/* Icera Semiconductor Inc. Deep Execution Processor */</span>
    <span id="EM_ALTERA_NIOS2">EM_ALTERA_NIOS2</span>  <a href="#Machine">Machine</a> = 113 <span class="comment">/* Altera Nios II soft-core processor */</span>
    <span id="EM_CRX">EM_CRX</span>           <a href="#Machine">Machine</a> = 114 <span class="comment">/* National Semiconductor CompactRISC CRX microprocessor */</span>
    <span id="EM_XGATE">EM_XGATE</span>         <a href="#Machine">Machine</a> = 115 <span class="comment">/* Motorola XGATE embedded processor */</span>
    <span id="EM_C166">EM_C166</span>          <a href="#Machine">Machine</a> = 116 <span class="comment">/* Infineon C16x/XC16x processor */</span>
    <span id="EM_M16C">EM_M16C</span>          <a href="#Machine">Machine</a> = 117 <span class="comment">/* Renesas M16C series microprocessors */</span>
    <span id="EM_DSPIC30F">EM_DSPIC30F</span>      <a href="#Machine">Machine</a> = 118 <span class="comment">/* Microchip Technology dsPIC30F Digital Signal Controller */</span>
    <span id="EM_CE">EM_CE</span>            <a href="#Machine">Machine</a> = 119 <span class="comment">/* Freescale Communication Engine RISC core */</span>
    <span id="EM_M32C">EM_M32C</span>          <a href="#Machine">Machine</a> = 120 <span class="comment">/* Renesas M32C series microprocessors */</span>
    <span id="EM_TSK3000">EM_TSK3000</span>       <a href="#Machine">Machine</a> = 131 <span class="comment">/* Altium TSK3000 core */</span>
    <span id="EM_RS08">EM_RS08</span>          <a href="#Machine">Machine</a> = 132 <span class="comment">/* Freescale RS08 embedded processor */</span>
    <span id="EM_SHARC">EM_SHARC</span>         <a href="#Machine">Machine</a> = 133 <span class="comment">/* Analog Devices SHARC family of 32-bit DSP processors */</span>
    <span id="EM_ECOG2">EM_ECOG2</span>         <a href="#Machine">Machine</a> = 134 <span class="comment">/* Cyan Technology eCOG2 microprocessor */</span>
    <span id="EM_SCORE7">EM_SCORE7</span>        <a href="#Machine">Machine</a> = 135 <span class="comment">/* Sunplus S+core7 RISC processor */</span>
    <span id="EM_DSP24">EM_DSP24</span>         <a href="#Machine">Machine</a> = 136 <span class="comment">/* New Japan Radio (NJR) 24-bit DSP Processor */</span>
    <span id="EM_VIDEOCORE3">EM_VIDEOCORE3</span>    <a href="#Machine">Machine</a> = 137 <span class="comment">/* Broadcom VideoCore III processor */</span>
    <span id="EM_LATTICEMICO32">EM_LATTICEMICO32</span> <a href="#Machine">Machine</a> = 138 <span class="comment">/* RISC processor for Lattice FPGA architecture */</span>
    <span id="EM_SE_C17">EM_SE_C17</span>        <a href="#Machine">Machine</a> = 139 <span class="comment">/* Seiko Epson C17 family */</span>
    <span id="EM_TI_C6000">EM_TI_C6000</span>      <a href="#Machine">Machine</a> = 140 <span class="comment">/* The Texas Instruments TMS320C6000 DSP family */</span>
    <span id="EM_TI_C2000">EM_TI_C2000</span>      <a href="#Machine">Machine</a> = 141 <span class="comment">/* The Texas Instruments TMS320C2000 DSP family */</span>
    <span id="EM_TI_C5500">EM_TI_C5500</span>      <a href="#Machine">Machine</a> = 142 <span class="comment">/* The Texas Instruments TMS320C55x DSP family */</span>
    <span id="EM_TI_ARP32">EM_TI_ARP32</span>      <a href="#Machine">Machine</a> = 143 <span class="comment">/* Texas Instruments Application Specific RISC Processor, 32bit fetch */</span>
    <span id="EM_TI_PRU">EM_TI_PRU</span>        <a href="#Machine">Machine</a> = 144 <span class="comment">/* Texas Instruments Programmable Realtime Unit */</span>
    <span id="EM_MMDSP_PLUS">EM_MMDSP_PLUS</span>    <a href="#Machine">Machine</a> = 160 <span class="comment">/* STMicroelectronics 64bit VLIW Data Signal Processor */</span>
    <span id="EM_CYPRESS_M8C">EM_CYPRESS_M8C</span>   <a href="#Machine">Machine</a> = 161 <span class="comment">/* Cypress M8C microprocessor */</span>
    <span id="EM_R32C">EM_R32C</span>          <a href="#Machine">Machine</a> = 162 <span class="comment">/* Renesas R32C series microprocessors */</span>
    <span id="EM_TRIMEDIA">EM_TRIMEDIA</span>      <a href="#Machine">Machine</a> = 163 <span class="comment">/* NXP Semiconductors TriMedia architecture family */</span>
    <span id="EM_QDSP6">EM_QDSP6</span>         <a href="#Machine">Machine</a> = 164 <span class="comment">/* QUALCOMM DSP6 Processor */</span>
    <span id="EM_8051">EM_8051</span>          <a href="#Machine">Machine</a> = 165 <span class="comment">/* Intel 8051 and variants */</span>
    <span id="EM_STXP7X">EM_STXP7X</span>        <a href="#Machine">Machine</a> = 166 <span class="comment">/* STMicroelectronics STxP7x family of configurable and extensible RISC processors */</span>
    <span id="EM_NDS32">EM_NDS32</span>         <a href="#Machine">Machine</a> = 167 <span class="comment">/* Andes Technology compact code size embedded RISC processor family */</span>
    <span id="EM_ECOG1">EM_ECOG1</span>         <a href="#Machine">Machine</a> = 168 <span class="comment">/* Cyan Technology eCOG1X family */</span>
    <span id="EM_ECOG1X">EM_ECOG1X</span>        <a href="#Machine">Machine</a> = 168 <span class="comment">/* Cyan Technology eCOG1X family */</span>
    <span id="EM_MAXQ30">EM_MAXQ30</span>        <a href="#Machine">Machine</a> = 169 <span class="comment">/* Dallas Semiconductor MAXQ30 Core Micro-controllers */</span>
    <span id="EM_XIMO16">EM_XIMO16</span>        <a href="#Machine">Machine</a> = 170 <span class="comment">/* New Japan Radio (NJR) 16-bit DSP Processor */</span>
    <span id="EM_MANIK">EM_MANIK</span>         <a href="#Machine">Machine</a> = 171 <span class="comment">/* M2000 Reconfigurable RISC Microprocessor */</span>
    <span id="EM_CRAYNV2">EM_CRAYNV2</span>       <a href="#Machine">Machine</a> = 172 <span class="comment">/* Cray Inc. NV2 vector architecture */</span>
    <span id="EM_RX">EM_RX</span>            <a href="#Machine">Machine</a> = 173 <span class="comment">/* Renesas RX family */</span>
    <span id="EM_METAG">EM_METAG</span>         <a href="#Machine">Machine</a> = 174 <span class="comment">/* Imagination Technologies META processor architecture */</span>
    <span id="EM_MCST_ELBRUS">EM_MCST_ELBRUS</span>   <a href="#Machine">Machine</a> = 175 <span class="comment">/* MCST Elbrus general purpose hardware architecture */</span>
    <span id="EM_ECOG16">EM_ECOG16</span>        <a href="#Machine">Machine</a> = 176 <span class="comment">/* Cyan Technology eCOG16 family */</span>
    <span id="EM_CR16">EM_CR16</span>          <a href="#Machine">Machine</a> = 177 <span class="comment">/* National Semiconductor CompactRISC CR16 16-bit microprocessor */</span>
    <span id="EM_ETPU">EM_ETPU</span>          <a href="#Machine">Machine</a> = 178 <span class="comment">/* Freescale Extended Time Processing Unit */</span>
    <span id="EM_SLE9X">EM_SLE9X</span>         <a href="#Machine">Machine</a> = 179 <span class="comment">/* Infineon Technologies SLE9X core */</span>
    <span id="EM_L10M">EM_L10M</span>          <a href="#Machine">Machine</a> = 180 <span class="comment">/* Intel L10M */</span>
    <span id="EM_K10M">EM_K10M</span>          <a href="#Machine">Machine</a> = 181 <span class="comment">/* Intel K10M */</span>
    <span id="EM_AARCH64">EM_AARCH64</span>       <a href="#Machine">Machine</a> = 183 <span class="comment">/* ARM 64-bit Architecture (AArch64) */</span>
    <span id="EM_AVR32">EM_AVR32</span>         <a href="#Machine">Machine</a> = 185 <span class="comment">/* Atmel Corporation 32-bit microprocessor family */</span>
    <span id="EM_STM8">EM_STM8</span>          <a href="#Machine">Machine</a> = 186 <span class="comment">/* STMicroeletronics STM8 8-bit microcontroller */</span>
    <span id="EM_TILE64">EM_TILE64</span>        <a href="#Machine">Machine</a> = 187 <span class="comment">/* Tilera TILE64 multicore architecture family */</span>
    <span id="EM_TILEPRO">EM_TILEPRO</span>       <a href="#Machine">Machine</a> = 188 <span class="comment">/* Tilera TILEPro multicore architecture family */</span>
    <span id="EM_MICROBLAZE">EM_MICROBLAZE</span>    <a href="#Machine">Machine</a> = 189 <span class="comment">/* Xilinx MicroBlaze 32-bit RISC soft processor core */</span>
    <span id="EM_CUDA">EM_CUDA</span>          <a href="#Machine">Machine</a> = 190 <span class="comment">/* NVIDIA CUDA architecture */</span>
    <span id="EM_TILEGX">EM_TILEGX</span>        <a href="#Machine">Machine</a> = 191 <span class="comment">/* Tilera TILE-Gx multicore architecture family */</span>
    <span id="EM_CLOUDSHIELD">EM_CLOUDSHIELD</span>   <a href="#Machine">Machine</a> = 192 <span class="comment">/* CloudShield architecture family */</span>
    <span id="EM_COREA_1ST">EM_COREA_1ST</span>     <a href="#Machine">Machine</a> = 193 <span class="comment">/* KIPO-KAIST Core-A 1st generation processor family */</span>
    <span id="EM_COREA_2ND">EM_COREA_2ND</span>     <a href="#Machine">Machine</a> = 194 <span class="comment">/* KIPO-KAIST Core-A 2nd generation processor family */</span>
    <span id="EM_ARC_COMPACT2">EM_ARC_COMPACT2</span>  <a href="#Machine">Machine</a> = 195 <span class="comment">/* Synopsys ARCompact V2 */</span>
    <span id="EM_OPEN8">EM_OPEN8</span>         <a href="#Machine">Machine</a> = 196 <span class="comment">/* Open8 8-bit RISC soft processor core */</span>
    <span id="EM_RL78">EM_RL78</span>          <a href="#Machine">Machine</a> = 197 <span class="comment">/* Renesas RL78 family */</span>
    <span id="EM_VIDEOCORE5">EM_VIDEOCORE5</span>    <a href="#Machine">Machine</a> = 198 <span class="comment">/* Broadcom VideoCore V processor */</span>
    <span id="EM_78KOR">EM_78KOR</span>         <a href="#Machine">Machine</a> = 199 <span class="comment">/* Renesas 78KOR family */</span>
    <span id="EM_56800EX">EM_56800EX</span>       <a href="#Machine">Machine</a> = 200 <span class="comment">/* Freescale 56800EX Digital Signal Controller (DSC) */</span>
    <span id="EM_BA1">EM_BA1</span>           <a href="#Machine">Machine</a> = 201 <span class="comment">/* Beyond BA1 CPU architecture */</span>
    <span id="EM_BA2">EM_BA2</span>           <a href="#Machine">Machine</a> = 202 <span class="comment">/* Beyond BA2 CPU architecture */</span>
    <span id="EM_XCORE">EM_XCORE</span>         <a href="#Machine">Machine</a> = 203 <span class="comment">/* XMOS xCORE processor family */</span>
    <span id="EM_MCHP_PIC">EM_MCHP_PIC</span>      <a href="#Machine">Machine</a> = 204 <span class="comment">/* Microchip 8-bit PIC(r) family */</span>
    <span id="EM_INTEL205">EM_INTEL205</span>      <a href="#Machine">Machine</a> = 205 <span class="comment">/* Reserved by Intel */</span>
    <span id="EM_INTEL206">EM_INTEL206</span>      <a href="#Machine">Machine</a> = 206 <span class="comment">/* Reserved by Intel */</span>
    <span id="EM_INTEL207">EM_INTEL207</span>      <a href="#Machine">Machine</a> = 207 <span class="comment">/* Reserved by Intel */</span>
    <span id="EM_INTEL208">EM_INTEL208</span>      <a href="#Machine">Machine</a> = 208 <span class="comment">/* Reserved by Intel */</span>
    <span id="EM_INTEL209">EM_INTEL209</span>      <a href="#Machine">Machine</a> = 209 <span class="comment">/* Reserved by Intel */</span>
    <span id="EM_KM32">EM_KM32</span>          <a href="#Machine">Machine</a> = 210 <span class="comment">/* KM211 KM32 32-bit processor */</span>
    <span id="EM_KMX32">EM_KMX32</span>         <a href="#Machine">Machine</a> = 211 <span class="comment">/* KM211 KMX32 32-bit processor */</span>
    <span id="EM_KMX16">EM_KMX16</span>         <a href="#Machine">Machine</a> = 212 <span class="comment">/* KM211 KMX16 16-bit processor */</span>
    <span id="EM_KMX8">EM_KMX8</span>          <a href="#Machine">Machine</a> = 213 <span class="comment">/* KM211 KMX8 8-bit processor */</span>
    <span id="EM_KVARC">EM_KVARC</span>         <a href="#Machine">Machine</a> = 214 <span class="comment">/* KM211 KVARC processor */</span>
    <span id="EM_CDP">EM_CDP</span>           <a href="#Machine">Machine</a> = 215 <span class="comment">/* Paneve CDP architecture family */</span>
    <span id="EM_COGE">EM_COGE</span>          <a href="#Machine">Machine</a> = 216 <span class="comment">/* Cognitive Smart Memory Processor */</span>
    <span id="EM_COOL">EM_COOL</span>          <a href="#Machine">Machine</a> = 217 <span class="comment">/* Bluechip Systems CoolEngine */</span>
    <span id="EM_NORC">EM_NORC</span>          <a href="#Machine">Machine</a> = 218 <span class="comment">/* Nanoradio Optimized RISC */</span>
    <span id="EM_CSR_KALIMBA">EM_CSR_KALIMBA</span>   <a href="#Machine">Machine</a> = 219 <span class="comment">/* CSR Kalimba architecture family */</span>
    <span id="EM_Z80">EM_Z80</span>           <a href="#Machine">Machine</a> = 220 <span class="comment">/* Zilog Z80 */</span>
    <span id="EM_VISIUM">EM_VISIUM</span>        <a href="#Machine">Machine</a> = 221 <span class="comment">/* Controls and Data Services VISIUMcore processor */</span>
    <span id="EM_FT32">EM_FT32</span>          <a href="#Machine">Machine</a> = 222 <span class="comment">/* FTDI Chip FT32 high performance 32-bit RISC architecture */</span>
    <span id="EM_MOXIE">EM_MOXIE</span>         <a href="#Machine">Machine</a> = 223 <span class="comment">/* Moxie processor family */</span>
    <span id="EM_AMDGPU">EM_AMDGPU</span>        <a href="#Machine">Machine</a> = 224 <span class="comment">/* AMD GPU architecture */</span>
    <span id="EM_RISCV">EM_RISCV</span>         <a href="#Machine">Machine</a> = 243 <span class="comment">/* RISC-V */</span>
    <span id="EM_LANAI">EM_LANAI</span>         <a href="#Machine">Machine</a> = 244 <span class="comment">/* Lanai 32-bit processor */</span>
    <span id="EM_BPF">EM_BPF</span>           <a href="#Machine">Machine</a> = 247 <span class="comment">/* Linux BPF – in-kernel virtual machine */</span>

    <span class="comment">/* Non-standard or deprecated. */</span>
    <span id="EM_486">EM_486</span>         <a href="#Machine">Machine</a> = 6      <span class="comment">/* Intel i486. */</span>
    <span id="EM_MIPS_RS4_BE">EM_MIPS_RS4_BE</span> <a href="#Machine">Machine</a> = 10     <span class="comment">/* MIPS R4000 Big-Endian */</span>
    <span id="EM_ALPHA_STD">EM_ALPHA_STD</span>   <a href="#Machine">Machine</a> = 41     <span class="comment">/* Digital Alpha (standard value). */</span>
    <span id="EM_ALPHA">EM_ALPHA</span>       <a href="#Machine">Machine</a> = 0x9026 <span class="comment">/* Alpha (written in the absence of an ABI) */</span>
)</pre>









### <a id="Machine.GoString">func</a> (Machine) [GoString](https://golang.org/src/debug/elf/elf.go?s=24010:24044#L577)
<pre>func (i <a href="#Machine">Machine</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Machine.String">func</a> (Machine) [String](https://golang.org/src/debug/elf/elf.go?s=23919:23951#L576)
<pre>func (i <a href="#Machine">Machine</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="NType">type</a> [NType](https://golang.org/src/debug/elf/elf.go?s=37959:37973#L918)
NType values; used in core files.


<pre>type NType <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="NT_PRSTATUS">NT_PRSTATUS</span> <a href="#NType">NType</a> = 1 <span class="comment">/* Process status. */</span>
    <span id="NT_FPREGSET">NT_FPREGSET</span> <a href="#NType">NType</a> = 2 <span class="comment">/* Floating point registers. */</span>
    <span id="NT_PRPSINFO">NT_PRPSINFO</span> <a href="#NType">NType</a> = 3 <span class="comment">/* Process state info. */</span>
)</pre>









### <a id="NType.GoString">func</a> (NType) [GoString](https://golang.org/src/debug/elf/elf.go?s=38318:38350#L933)
<pre>func (i <a href="#NType">NType</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="NType.String">func</a> (NType) [String](https://golang.org/src/debug/elf/elf.go?s=38231:38261#L932)
<pre>func (i <a href="#NType">NType</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="OSABI">type</a> [OSABI](https://golang.org/src/debug/elf/elf.go?s=4546:4561#L110)
OSABI is found in Header.Ident[EI_OSABI] and Header.OSABI.


<pre>type OSABI <a href="/pkg/builtin/#byte">byte</a></pre>



<pre>const (
    <span id="ELFOSABI_NONE">ELFOSABI_NONE</span>       <a href="#OSABI">OSABI</a> = 0   <span class="comment">/* UNIX System V ABI */</span>
    <span id="ELFOSABI_HPUX">ELFOSABI_HPUX</span>       <a href="#OSABI">OSABI</a> = 1   <span class="comment">/* HP-UX operating system */</span>
    <span id="ELFOSABI_NETBSD">ELFOSABI_NETBSD</span>     <a href="#OSABI">OSABI</a> = 2   <span class="comment">/* NetBSD */</span>
    <span id="ELFOSABI_LINUX">ELFOSABI_LINUX</span>      <a href="#OSABI">OSABI</a> = 3   <span class="comment">/* GNU/Linux */</span>
    <span id="ELFOSABI_HURD">ELFOSABI_HURD</span>       <a href="#OSABI">OSABI</a> = 4   <span class="comment">/* GNU/Hurd */</span>
    <span id="ELFOSABI_86OPEN">ELFOSABI_86OPEN</span>     <a href="#OSABI">OSABI</a> = 5   <span class="comment">/* 86Open common IA32 ABI */</span>
    <span id="ELFOSABI_SOLARIS">ELFOSABI_SOLARIS</span>    <a href="#OSABI">OSABI</a> = 6   <span class="comment">/* Solaris */</span>
    <span id="ELFOSABI_AIX">ELFOSABI_AIX</span>        <a href="#OSABI">OSABI</a> = 7   <span class="comment">/* AIX */</span>
    <span id="ELFOSABI_IRIX">ELFOSABI_IRIX</span>       <a href="#OSABI">OSABI</a> = 8   <span class="comment">/* IRIX */</span>
    <span id="ELFOSABI_FREEBSD">ELFOSABI_FREEBSD</span>    <a href="#OSABI">OSABI</a> = 9   <span class="comment">/* FreeBSD */</span>
    <span id="ELFOSABI_TRU64">ELFOSABI_TRU64</span>      <a href="#OSABI">OSABI</a> = 10  <span class="comment">/* TRU64 UNIX */</span>
    <span id="ELFOSABI_MODESTO">ELFOSABI_MODESTO</span>    <a href="#OSABI">OSABI</a> = 11  <span class="comment">/* Novell Modesto */</span>
    <span id="ELFOSABI_OPENBSD">ELFOSABI_OPENBSD</span>    <a href="#OSABI">OSABI</a> = 12  <span class="comment">/* OpenBSD */</span>
    <span id="ELFOSABI_OPENVMS">ELFOSABI_OPENVMS</span>    <a href="#OSABI">OSABI</a> = 13  <span class="comment">/* Open VMS */</span>
    <span id="ELFOSABI_NSK">ELFOSABI_NSK</span>        <a href="#OSABI">OSABI</a> = 14  <span class="comment">/* HP Non-Stop Kernel */</span>
    <span id="ELFOSABI_AROS">ELFOSABI_AROS</span>       <a href="#OSABI">OSABI</a> = 15  <span class="comment">/* Amiga Research OS */</span>
    <span id="ELFOSABI_FENIXOS">ELFOSABI_FENIXOS</span>    <a href="#OSABI">OSABI</a> = 16  <span class="comment">/* The FenixOS highly scalable multi-core OS */</span>
    <span id="ELFOSABI_CLOUDABI">ELFOSABI_CLOUDABI</span>   <a href="#OSABI">OSABI</a> = 17  <span class="comment">/* Nuxi CloudABI */</span>
    <span id="ELFOSABI_ARM">ELFOSABI_ARM</span>        <a href="#OSABI">OSABI</a> = 97  <span class="comment">/* ARM */</span>
    <span id="ELFOSABI_STANDALONE">ELFOSABI_STANDALONE</span> <a href="#OSABI">OSABI</a> = 255 <span class="comment">/* Standalone (embedded) application */</span>
)</pre>









### <a id="OSABI.GoString">func</a> (OSABI) [GoString](https://golang.org/src/debug/elf/elf.go?s=6265:6297#L159)
<pre>func (i <a href="#OSABI">OSABI</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="OSABI.String">func</a> (OSABI) [String](https://golang.org/src/debug/elf/elf.go?s=6178:6208#L158)
<pre>func (i <a href="#OSABI">OSABI</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Prog">type</a> [Prog](https://golang.org/src/debug/elf/file.go?s=3760:4063#L142)
A Prog represents a single ELF program header in an ELF binary.


<pre>type Prog struct {
    <a href="#ProgHeader">ProgHeader</a>

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











### <a id="Prog.Open">func</a> (\*Prog) [Open](https://golang.org/src/debug/elf/file.go?s=4128:4163#L156)
<pre>func (p *<a href="#Prog">Prog</a>) Open() <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadSeeker">ReadSeeker</a></pre>
Open returns a new ReadSeeker reading the ELF program body.




## <a id="Prog32">type</a> [Prog32](https://golang.org/src/debug/elf/elf.go?s=107626:108024#L2774)
ELF32 Program header.


<pre>type Prog32 struct {
<span id="Prog32.Type"></span>    Type   <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Entry type. */</span>
<span id="Prog32.Off"></span>    Off    <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* File offset of contents. */</span>
<span id="Prog32.Vaddr"></span>    Vaddr  <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Virtual address in memory image. */</span>
<span id="Prog32.Paddr"></span>    Paddr  <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Physical address (not used). */</span>
<span id="Prog32.Filesz"></span>    Filesz <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Size of contents in file. */</span>
<span id="Prog32.Memsz"></span>    Memsz  <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Size of contents in memory. */</span>
<span id="Prog32.Flags"></span>    Flags  <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Access permission flags. */</span>
<span id="Prog32.Align"></span>    Align  <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Alignment in memory and file. */</span>
}
</pre>











## <a id="Prog64">type</a> [Prog64](https://golang.org/src/debug/elf/elf.go?s=110778:111176#L2875)
ELF64 Program header.


<pre>type Prog64 struct {
<span id="Prog64.Type"></span>    Type   <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Entry type. */</span>
<span id="Prog64.Flags"></span>    Flags  <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Access permission flags. */</span>
<span id="Prog64.Off"></span>    Off    <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* File offset of contents. */</span>
<span id="Prog64.Vaddr"></span>    Vaddr  <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Virtual address in memory image. */</span>
<span id="Prog64.Paddr"></span>    Paddr  <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Physical address (not used). */</span>
<span id="Prog64.Filesz"></span>    Filesz <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Size of contents in file. */</span>
<span id="Prog64.Memsz"></span>    Memsz  <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Size of contents in memory. */</span>
<span id="Prog64.Align"></span>    Align  <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Alignment in memory and file. */</span>
}
</pre>











## <a id="ProgFlag">type</a> [ProgFlag](https://golang.org/src/debug/elf/elf.go?s=32098:32118#L771)
Prog.Flag


<pre>type ProgFlag <a href="/pkg/builtin/#uint32">uint32</a></pre>



<pre>const (
    <span id="PF_X">PF_X</span>        <a href="#ProgFlag">ProgFlag</a> = 0x1        <span class="comment">/* Executable. */</span>
    <span id="PF_W">PF_W</span>        <a href="#ProgFlag">ProgFlag</a> = 0x2        <span class="comment">/* Writable. */</span>
    <span id="PF_R">PF_R</span>        <a href="#ProgFlag">ProgFlag</a> = 0x4        <span class="comment">/* Readable. */</span>
    <span id="PF_MASKOS">PF_MASKOS</span>   <a href="#ProgFlag">ProgFlag</a> = 0x0ff00000 <span class="comment">/* Operating system-specific. */</span>
    <span id="PF_MASKPROC">PF_MASKPROC</span> <a href="#ProgFlag">ProgFlag</a> = 0xf0000000 <span class="comment">/* Processor-specific. */</span>
)</pre>









### <a id="ProgFlag.GoString">func</a> (ProgFlag) [GoString](https://golang.org/src/debug/elf/elf.go?s=32578:32613#L788)
<pre>func (i <a href="#ProgFlag">ProgFlag</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="ProgFlag.String">func</a> (ProgFlag) [String](https://golang.org/src/debug/elf/elf.go?s=32493:32526#L787)
<pre>func (i <a href="#ProgFlag">ProgFlag</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="ProgHeader">type</a> [ProgHeader](https://golang.org/src/debug/elf/file.go?s=3541:3691#L130)
A ProgHeader represents a single ELF program header.


<pre>type ProgHeader struct {
<span id="ProgHeader.Type"></span>    Type   <a href="#ProgType">ProgType</a>
<span id="ProgHeader.Flags"></span>    Flags  <a href="#ProgFlag">ProgFlag</a>
<span id="ProgHeader.Off"></span>    Off    <a href="/pkg/builtin/#uint64">uint64</a>
<span id="ProgHeader.Vaddr"></span>    Vaddr  <a href="/pkg/builtin/#uint64">uint64</a>
<span id="ProgHeader.Paddr"></span>    Paddr  <a href="/pkg/builtin/#uint64">uint64</a>
<span id="ProgHeader.Filesz"></span>    Filesz <a href="/pkg/builtin/#uint64">uint64</a>
<span id="ProgHeader.Memsz"></span>    Memsz  <a href="/pkg/builtin/#uint64">uint64</a>
<span id="ProgHeader.Align"></span>    Align  <a href="/pkg/builtin/#uint64">uint64</a>
}
</pre>











## <a id="ProgType">type</a> [ProgType](https://golang.org/src/debug/elf/elf.go?s=30822:30839#L735)
Prog.Type


<pre>type ProgType <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="PT_NULL">PT_NULL</span>    <a href="#ProgType">ProgType</a> = 0          <span class="comment">/* Unused entry. */</span>
    <span id="PT_LOAD">PT_LOAD</span>    <a href="#ProgType">ProgType</a> = 1          <span class="comment">/* Loadable segment. */</span>
    <span id="PT_DYNAMIC">PT_DYNAMIC</span> <a href="#ProgType">ProgType</a> = 2          <span class="comment">/* Dynamic linking information segment. */</span>
    <span id="PT_INTERP">PT_INTERP</span>  <a href="#ProgType">ProgType</a> = 3          <span class="comment">/* Pathname of interpreter. */</span>
    <span id="PT_NOTE">PT_NOTE</span>    <a href="#ProgType">ProgType</a> = 4          <span class="comment">/* Auxiliary information. */</span>
    <span id="PT_SHLIB">PT_SHLIB</span>   <a href="#ProgType">ProgType</a> = 5          <span class="comment">/* Reserved (not used). */</span>
    <span id="PT_PHDR">PT_PHDR</span>    <a href="#ProgType">ProgType</a> = 6          <span class="comment">/* Location of program header itself. */</span>
    <span id="PT_TLS">PT_TLS</span>     <a href="#ProgType">ProgType</a> = 7          <span class="comment">/* Thread local storage segment */</span>
    <span id="PT_LOOS">PT_LOOS</span>    <a href="#ProgType">ProgType</a> = 0x60000000 <span class="comment">/* First OS-specific. */</span>
    <span id="PT_HIOS">PT_HIOS</span>    <a href="#ProgType">ProgType</a> = 0x6fffffff <span class="comment">/* Last OS-specific. */</span>
    <span id="PT_LOPROC">PT_LOPROC</span>  <a href="#ProgType">ProgType</a> = 0x70000000 <span class="comment">/* First processor-specific type. */</span>
    <span id="PT_HIPROC">PT_HIPROC</span>  <a href="#ProgType">ProgType</a> = 0x7fffffff <span class="comment">/* Last processor-specific type. */</span>
)</pre>









### <a id="ProgType.GoString">func</a> (ProgType) [GoString](https://golang.org/src/debug/elf/elf.go?s=31998:32033#L768)
<pre>func (i <a href="#ProgType">ProgType</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="ProgType.String">func</a> (ProgType) [String](https://golang.org/src/debug/elf/elf.go?s=31911:31944#L767)
<pre>func (i <a href="#ProgType">ProgType</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="R_386">type</a> [R_386](https://golang.org/src/debug/elf/elf.go?s=70292:70306#L1758)
Relocation types for 386.


<pre>type R_386 <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="R_386_NONE">R_386_NONE</span>          <a href="#R_386">R_386</a> = 0  <span class="comment">/* No relocation. */</span>
    <span id="R_386_32">R_386_32</span>            <a href="#R_386">R_386</a> = 1  <span class="comment">/* Add symbol value. */</span>
    <span id="R_386_PC32">R_386_PC32</span>          <a href="#R_386">R_386</a> = 2  <span class="comment">/* Add PC-relative symbol value. */</span>
    <span id="R_386_GOT32">R_386_GOT32</span>         <a href="#R_386">R_386</a> = 3  <span class="comment">/* Add PC-relative GOT offset. */</span>
    <span id="R_386_PLT32">R_386_PLT32</span>         <a href="#R_386">R_386</a> = 4  <span class="comment">/* Add PC-relative PLT offset. */</span>
    <span id="R_386_COPY">R_386_COPY</span>          <a href="#R_386">R_386</a> = 5  <span class="comment">/* Copy data from shared object. */</span>
    <span id="R_386_GLOB_DAT">R_386_GLOB_DAT</span>      <a href="#R_386">R_386</a> = 6  <span class="comment">/* Set GOT entry to data address. */</span>
    <span id="R_386_JMP_SLOT">R_386_JMP_SLOT</span>      <a href="#R_386">R_386</a> = 7  <span class="comment">/* Set GOT entry to code address. */</span>
    <span id="R_386_RELATIVE">R_386_RELATIVE</span>      <a href="#R_386">R_386</a> = 8  <span class="comment">/* Add load address of shared object. */</span>
    <span id="R_386_GOTOFF">R_386_GOTOFF</span>        <a href="#R_386">R_386</a> = 9  <span class="comment">/* Add GOT-relative symbol address. */</span>
    <span id="R_386_GOTPC">R_386_GOTPC</span>         <a href="#R_386">R_386</a> = 10 <span class="comment">/* Add PC-relative GOT table address. */</span>
    <span id="R_386_32PLT">R_386_32PLT</span>         <a href="#R_386">R_386</a> = 11
    <span id="R_386_TLS_TPOFF">R_386_TLS_TPOFF</span>     <a href="#R_386">R_386</a> = 14 <span class="comment">/* Negative offset in static TLS block */</span>
    <span id="R_386_TLS_IE">R_386_TLS_IE</span>        <a href="#R_386">R_386</a> = 15 <span class="comment">/* Absolute address of GOT for -ve static TLS */</span>
    <span id="R_386_TLS_GOTIE">R_386_TLS_GOTIE</span>     <a href="#R_386">R_386</a> = 16 <span class="comment">/* GOT entry for negative static TLS block */</span>
    <span id="R_386_TLS_LE">R_386_TLS_LE</span>        <a href="#R_386">R_386</a> = 17 <span class="comment">/* Negative offset relative to static TLS */</span>
    <span id="R_386_TLS_GD">R_386_TLS_GD</span>        <a href="#R_386">R_386</a> = 18 <span class="comment">/* 32 bit offset to GOT (index,off) pair */</span>
    <span id="R_386_TLS_LDM">R_386_TLS_LDM</span>       <a href="#R_386">R_386</a> = 19 <span class="comment">/* 32 bit offset to GOT (index,zero) pair */</span>
    <span id="R_386_16">R_386_16</span>            <a href="#R_386">R_386</a> = 20
    <span id="R_386_PC16">R_386_PC16</span>          <a href="#R_386">R_386</a> = 21
    <span id="R_386_8">R_386_8</span>             <a href="#R_386">R_386</a> = 22
    <span id="R_386_PC8">R_386_PC8</span>           <a href="#R_386">R_386</a> = 23
    <span id="R_386_TLS_GD_32">R_386_TLS_GD_32</span>     <a href="#R_386">R_386</a> = 24 <span class="comment">/* 32 bit offset to GOT (index,off) pair */</span>
    <span id="R_386_TLS_GD_PUSH">R_386_TLS_GD_PUSH</span>   <a href="#R_386">R_386</a> = 25 <span class="comment">/* pushl instruction for Sun ABI GD sequence */</span>
    <span id="R_386_TLS_GD_CALL">R_386_TLS_GD_CALL</span>   <a href="#R_386">R_386</a> = 26 <span class="comment">/* call instruction for Sun ABI GD sequence */</span>
    <span id="R_386_TLS_GD_POP">R_386_TLS_GD_POP</span>    <a href="#R_386">R_386</a> = 27 <span class="comment">/* popl instruction for Sun ABI GD sequence */</span>
    <span id="R_386_TLS_LDM_32">R_386_TLS_LDM_32</span>    <a href="#R_386">R_386</a> = 28 <span class="comment">/* 32 bit offset to GOT (index,zero) pair */</span>
    <span id="R_386_TLS_LDM_PUSH">R_386_TLS_LDM_PUSH</span>  <a href="#R_386">R_386</a> = 29 <span class="comment">/* pushl instruction for Sun ABI LD sequence */</span>
    <span id="R_386_TLS_LDM_CALL">R_386_TLS_LDM_CALL</span>  <a href="#R_386">R_386</a> = 30 <span class="comment">/* call instruction for Sun ABI LD sequence */</span>
    <span id="R_386_TLS_LDM_POP">R_386_TLS_LDM_POP</span>   <a href="#R_386">R_386</a> = 31 <span class="comment">/* popl instruction for Sun ABI LD sequence */</span>
    <span id="R_386_TLS_LDO_32">R_386_TLS_LDO_32</span>    <a href="#R_386">R_386</a> = 32 <span class="comment">/* 32 bit offset from start of TLS block */</span>
    <span id="R_386_TLS_IE_32">R_386_TLS_IE_32</span>     <a href="#R_386">R_386</a> = 33 <span class="comment">/* 32 bit offset to GOT static TLS offset entry */</span>
    <span id="R_386_TLS_LE_32">R_386_TLS_LE_32</span>     <a href="#R_386">R_386</a> = 34 <span class="comment">/* 32 bit offset within static TLS block */</span>
    <span id="R_386_TLS_DTPMOD32">R_386_TLS_DTPMOD32</span>  <a href="#R_386">R_386</a> = 35 <span class="comment">/* GOT entry containing TLS index */</span>
    <span id="R_386_TLS_DTPOFF32">R_386_TLS_DTPOFF32</span>  <a href="#R_386">R_386</a> = 36 <span class="comment">/* GOT entry containing TLS offset */</span>
    <span id="R_386_TLS_TPOFF32">R_386_TLS_TPOFF32</span>   <a href="#R_386">R_386</a> = 37 <span class="comment">/* GOT entry of -ve static TLS offset */</span>
    <span id="R_386_SIZE32">R_386_SIZE32</span>        <a href="#R_386">R_386</a> = 38
    <span id="R_386_TLS_GOTDESC">R_386_TLS_GOTDESC</span>   <a href="#R_386">R_386</a> = 39
    <span id="R_386_TLS_DESC_CALL">R_386_TLS_DESC_CALL</span> <a href="#R_386">R_386</a> = 40
    <span id="R_386_TLS_DESC">R_386_TLS_DESC</span>      <a href="#R_386">R_386</a> = 41
    <span id="R_386_IRELATIVE">R_386_IRELATIVE</span>     <a href="#R_386">R_386</a> = 42
    <span id="R_386_GOT32X">R_386_GOT32X</span>        <a href="#R_386">R_386</a> = 43
)</pre>









### <a id="R_386.GoString">func</a> (R\_386) [GoString](https://golang.org/src/debug/elf/elf.go?s=74083:74115#L1851)
<pre>func (i <a href="#R_386">R_386</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="R_386.String">func</a> (R\_386) [String](https://golang.org/src/debug/elf/elf.go?s=73997:74027#L1850)
<pre>func (i <a href="#R_386">R_386</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="R_390">type</a> [R_390](https://golang.org/src/debug/elf/elf.go?s=99232:99246#L2481)
Relocation types for s390x processors.


<pre>type R_390 <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="R_390_NONE">R_390_NONE</span>        <a href="#R_390">R_390</a> = 0
    <span id="R_390_8">R_390_8</span>           <a href="#R_390">R_390</a> = 1
    <span id="R_390_12">R_390_12</span>          <a href="#R_390">R_390</a> = 2
    <span id="R_390_16">R_390_16</span>          <a href="#R_390">R_390</a> = 3
    <span id="R_390_32">R_390_32</span>          <a href="#R_390">R_390</a> = 4
    <span id="R_390_PC32">R_390_PC32</span>        <a href="#R_390">R_390</a> = 5
    <span id="R_390_GOT12">R_390_GOT12</span>       <a href="#R_390">R_390</a> = 6
    <span id="R_390_GOT32">R_390_GOT32</span>       <a href="#R_390">R_390</a> = 7
    <span id="R_390_PLT32">R_390_PLT32</span>       <a href="#R_390">R_390</a> = 8
    <span id="R_390_COPY">R_390_COPY</span>        <a href="#R_390">R_390</a> = 9
    <span id="R_390_GLOB_DAT">R_390_GLOB_DAT</span>    <a href="#R_390">R_390</a> = 10
    <span id="R_390_JMP_SLOT">R_390_JMP_SLOT</span>    <a href="#R_390">R_390</a> = 11
    <span id="R_390_RELATIVE">R_390_RELATIVE</span>    <a href="#R_390">R_390</a> = 12
    <span id="R_390_GOTOFF">R_390_GOTOFF</span>      <a href="#R_390">R_390</a> = 13
    <span id="R_390_GOTPC">R_390_GOTPC</span>       <a href="#R_390">R_390</a> = 14
    <span id="R_390_GOT16">R_390_GOT16</span>       <a href="#R_390">R_390</a> = 15
    <span id="R_390_PC16">R_390_PC16</span>        <a href="#R_390">R_390</a> = 16
    <span id="R_390_PC16DBL">R_390_PC16DBL</span>     <a href="#R_390">R_390</a> = 17
    <span id="R_390_PLT16DBL">R_390_PLT16DBL</span>    <a href="#R_390">R_390</a> = 18
    <span id="R_390_PC32DBL">R_390_PC32DBL</span>     <a href="#R_390">R_390</a> = 19
    <span id="R_390_PLT32DBL">R_390_PLT32DBL</span>    <a href="#R_390">R_390</a> = 20
    <span id="R_390_GOTPCDBL">R_390_GOTPCDBL</span>    <a href="#R_390">R_390</a> = 21
    <span id="R_390_64">R_390_64</span>          <a href="#R_390">R_390</a> = 22
    <span id="R_390_PC64">R_390_PC64</span>        <a href="#R_390">R_390</a> = 23
    <span id="R_390_GOT64">R_390_GOT64</span>       <a href="#R_390">R_390</a> = 24
    <span id="R_390_PLT64">R_390_PLT64</span>       <a href="#R_390">R_390</a> = 25
    <span id="R_390_GOTENT">R_390_GOTENT</span>      <a href="#R_390">R_390</a> = 26
    <span id="R_390_GOTOFF16">R_390_GOTOFF16</span>    <a href="#R_390">R_390</a> = 27
    <span id="R_390_GOTOFF64">R_390_GOTOFF64</span>    <a href="#R_390">R_390</a> = 28
    <span id="R_390_GOTPLT12">R_390_GOTPLT12</span>    <a href="#R_390">R_390</a> = 29
    <span id="R_390_GOTPLT16">R_390_GOTPLT16</span>    <a href="#R_390">R_390</a> = 30
    <span id="R_390_GOTPLT32">R_390_GOTPLT32</span>    <a href="#R_390">R_390</a> = 31
    <span id="R_390_GOTPLT64">R_390_GOTPLT64</span>    <a href="#R_390">R_390</a> = 32
    <span id="R_390_GOTPLTENT">R_390_GOTPLTENT</span>   <a href="#R_390">R_390</a> = 33
    <span id="R_390_GOTPLTOFF16">R_390_GOTPLTOFF16</span> <a href="#R_390">R_390</a> = 34
    <span id="R_390_GOTPLTOFF32">R_390_GOTPLTOFF32</span> <a href="#R_390">R_390</a> = 35
    <span id="R_390_GOTPLTOFF64">R_390_GOTPLTOFF64</span> <a href="#R_390">R_390</a> = 36
    <span id="R_390_TLS_LOAD">R_390_TLS_LOAD</span>    <a href="#R_390">R_390</a> = 37
    <span id="R_390_TLS_GDCALL">R_390_TLS_GDCALL</span>  <a href="#R_390">R_390</a> = 38
    <span id="R_390_TLS_LDCALL">R_390_TLS_LDCALL</span>  <a href="#R_390">R_390</a> = 39
    <span id="R_390_TLS_GD32">R_390_TLS_GD32</span>    <a href="#R_390">R_390</a> = 40
    <span id="R_390_TLS_GD64">R_390_TLS_GD64</span>    <a href="#R_390">R_390</a> = 41
    <span id="R_390_TLS_GOTIE12">R_390_TLS_GOTIE12</span> <a href="#R_390">R_390</a> = 42
    <span id="R_390_TLS_GOTIE32">R_390_TLS_GOTIE32</span> <a href="#R_390">R_390</a> = 43
    <span id="R_390_TLS_GOTIE64">R_390_TLS_GOTIE64</span> <a href="#R_390">R_390</a> = 44
    <span id="R_390_TLS_LDM32">R_390_TLS_LDM32</span>   <a href="#R_390">R_390</a> = 45
    <span id="R_390_TLS_LDM64">R_390_TLS_LDM64</span>   <a href="#R_390">R_390</a> = 46
    <span id="R_390_TLS_IE32">R_390_TLS_IE32</span>    <a href="#R_390">R_390</a> = 47
    <span id="R_390_TLS_IE64">R_390_TLS_IE64</span>    <a href="#R_390">R_390</a> = 48
    <span id="R_390_TLS_IEENT">R_390_TLS_IEENT</span>   <a href="#R_390">R_390</a> = 49
    <span id="R_390_TLS_LE32">R_390_TLS_LE32</span>    <a href="#R_390">R_390</a> = 50
    <span id="R_390_TLS_LE64">R_390_TLS_LE64</span>    <a href="#R_390">R_390</a> = 51
    <span id="R_390_TLS_LDO32">R_390_TLS_LDO32</span>   <a href="#R_390">R_390</a> = 52
    <span id="R_390_TLS_LDO64">R_390_TLS_LDO64</span>   <a href="#R_390">R_390</a> = 53
    <span id="R_390_TLS_DTPMOD">R_390_TLS_DTPMOD</span>  <a href="#R_390">R_390</a> = 54
    <span id="R_390_TLS_DTPOFF">R_390_TLS_DTPOFF</span>  <a href="#R_390">R_390</a> = 55
    <span id="R_390_TLS_TPOFF">R_390_TLS_TPOFF</span>   <a href="#R_390">R_390</a> = 56
    <span id="R_390_20">R_390_20</span>          <a href="#R_390">R_390</a> = 57
    <span id="R_390_GOT20">R_390_GOT20</span>       <a href="#R_390">R_390</a> = 58
    <span id="R_390_GOTPLT20">R_390_GOTPLT20</span>    <a href="#R_390">R_390</a> = 59
    <span id="R_390_TLS_GOTIE20">R_390_TLS_GOTIE20</span> <a href="#R_390">R_390</a> = 60
)</pre>









### <a id="R_390.GoString">func</a> (R\_390) [GoString](https://golang.org/src/debug/elf/elf.go?s=102661:102693#L2612)
<pre>func (i <a href="#R_390">R_390</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="R_390.String">func</a> (R\_390) [String](https://golang.org/src/debug/elf/elf.go?s=102575:102605#L2611)
<pre>func (i <a href="#R_390">R_390</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="R_AARCH64">type</a> [R_AARCH64](https://golang.org/src/debug/elf/elf.go?s=45095:45113#L1118)
Relocation types for AArch64 (aka arm64)


<pre>type R_AARCH64 <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="R_AARCH64_NONE">R_AARCH64_NONE</span>                            <a href="#R_AARCH64">R_AARCH64</a> = 0
    <span id="R_AARCH64_P32_ABS32">R_AARCH64_P32_ABS32</span>                       <a href="#R_AARCH64">R_AARCH64</a> = 1
    <span id="R_AARCH64_P32_ABS16">R_AARCH64_P32_ABS16</span>                       <a href="#R_AARCH64">R_AARCH64</a> = 2
    <span id="R_AARCH64_P32_PREL32">R_AARCH64_P32_PREL32</span>                      <a href="#R_AARCH64">R_AARCH64</a> = 3
    <span id="R_AARCH64_P32_PREL16">R_AARCH64_P32_PREL16</span>                      <a href="#R_AARCH64">R_AARCH64</a> = 4
    <span id="R_AARCH64_P32_MOVW_UABS_G0">R_AARCH64_P32_MOVW_UABS_G0</span>                <a href="#R_AARCH64">R_AARCH64</a> = 5
    <span id="R_AARCH64_P32_MOVW_UABS_G0_NC">R_AARCH64_P32_MOVW_UABS_G0_NC</span>             <a href="#R_AARCH64">R_AARCH64</a> = 6
    <span id="R_AARCH64_P32_MOVW_UABS_G1">R_AARCH64_P32_MOVW_UABS_G1</span>                <a href="#R_AARCH64">R_AARCH64</a> = 7
    <span id="R_AARCH64_P32_MOVW_SABS_G0">R_AARCH64_P32_MOVW_SABS_G0</span>                <a href="#R_AARCH64">R_AARCH64</a> = 8
    <span id="R_AARCH64_P32_LD_PREL_LO19">R_AARCH64_P32_LD_PREL_LO19</span>                <a href="#R_AARCH64">R_AARCH64</a> = 9
    <span id="R_AARCH64_P32_ADR_PREL_LO21">R_AARCH64_P32_ADR_PREL_LO21</span>               <a href="#R_AARCH64">R_AARCH64</a> = 10
    <span id="R_AARCH64_P32_ADR_PREL_PG_HI21">R_AARCH64_P32_ADR_PREL_PG_HI21</span>            <a href="#R_AARCH64">R_AARCH64</a> = 11
    <span id="R_AARCH64_P32_ADD_ABS_LO12_NC">R_AARCH64_P32_ADD_ABS_LO12_NC</span>             <a href="#R_AARCH64">R_AARCH64</a> = 12
    <span id="R_AARCH64_P32_LDST8_ABS_LO12_NC">R_AARCH64_P32_LDST8_ABS_LO12_NC</span>           <a href="#R_AARCH64">R_AARCH64</a> = 13
    <span id="R_AARCH64_P32_LDST16_ABS_LO12_NC">R_AARCH64_P32_LDST16_ABS_LO12_NC</span>          <a href="#R_AARCH64">R_AARCH64</a> = 14
    <span id="R_AARCH64_P32_LDST32_ABS_LO12_NC">R_AARCH64_P32_LDST32_ABS_LO12_NC</span>          <a href="#R_AARCH64">R_AARCH64</a> = 15
    <span id="R_AARCH64_P32_LDST64_ABS_LO12_NC">R_AARCH64_P32_LDST64_ABS_LO12_NC</span>          <a href="#R_AARCH64">R_AARCH64</a> = 16
    <span id="R_AARCH64_P32_LDST128_ABS_LO12_NC">R_AARCH64_P32_LDST128_ABS_LO12_NC</span>         <a href="#R_AARCH64">R_AARCH64</a> = 17
    <span id="R_AARCH64_P32_TSTBR14">R_AARCH64_P32_TSTBR14</span>                     <a href="#R_AARCH64">R_AARCH64</a> = 18
    <span id="R_AARCH64_P32_CONDBR19">R_AARCH64_P32_CONDBR19</span>                    <a href="#R_AARCH64">R_AARCH64</a> = 19
    <span id="R_AARCH64_P32_JUMP26">R_AARCH64_P32_JUMP26</span>                      <a href="#R_AARCH64">R_AARCH64</a> = 20
    <span id="R_AARCH64_P32_CALL26">R_AARCH64_P32_CALL26</span>                      <a href="#R_AARCH64">R_AARCH64</a> = 21
    <span id="R_AARCH64_P32_GOT_LD_PREL19">R_AARCH64_P32_GOT_LD_PREL19</span>               <a href="#R_AARCH64">R_AARCH64</a> = 25
    <span id="R_AARCH64_P32_ADR_GOT_PAGE">R_AARCH64_P32_ADR_GOT_PAGE</span>                <a href="#R_AARCH64">R_AARCH64</a> = 26
    <span id="R_AARCH64_P32_LD32_GOT_LO12_NC">R_AARCH64_P32_LD32_GOT_LO12_NC</span>            <a href="#R_AARCH64">R_AARCH64</a> = 27
    <span id="R_AARCH64_P32_TLSGD_ADR_PAGE21">R_AARCH64_P32_TLSGD_ADR_PAGE21</span>            <a href="#R_AARCH64">R_AARCH64</a> = 81
    <span id="R_AARCH64_P32_TLSGD_ADD_LO12_NC">R_AARCH64_P32_TLSGD_ADD_LO12_NC</span>           <a href="#R_AARCH64">R_AARCH64</a> = 82
    <span id="R_AARCH64_P32_TLSIE_ADR_GOTTPREL_PAGE21">R_AARCH64_P32_TLSIE_ADR_GOTTPREL_PAGE21</span>   <a href="#R_AARCH64">R_AARCH64</a> = 103
    <span id="R_AARCH64_P32_TLSIE_LD32_GOTTPREL_LO12_NC">R_AARCH64_P32_TLSIE_LD32_GOTTPREL_LO12_NC</span> <a href="#R_AARCH64">R_AARCH64</a> = 104
    <span id="R_AARCH64_P32_TLSIE_LD_GOTTPREL_PREL19">R_AARCH64_P32_TLSIE_LD_GOTTPREL_PREL19</span>    <a href="#R_AARCH64">R_AARCH64</a> = 105
    <span id="R_AARCH64_P32_TLSLE_MOVW_TPREL_G1">R_AARCH64_P32_TLSLE_MOVW_TPREL_G1</span>         <a href="#R_AARCH64">R_AARCH64</a> = 106
    <span id="R_AARCH64_P32_TLSLE_MOVW_TPREL_G0">R_AARCH64_P32_TLSLE_MOVW_TPREL_G0</span>         <a href="#R_AARCH64">R_AARCH64</a> = 107
    <span id="R_AARCH64_P32_TLSLE_MOVW_TPREL_G0_NC">R_AARCH64_P32_TLSLE_MOVW_TPREL_G0_NC</span>      <a href="#R_AARCH64">R_AARCH64</a> = 108
    <span id="R_AARCH64_P32_TLSLE_ADD_TPREL_HI12">R_AARCH64_P32_TLSLE_ADD_TPREL_HI12</span>        <a href="#R_AARCH64">R_AARCH64</a> = 109
    <span id="R_AARCH64_P32_TLSLE_ADD_TPREL_LO12">R_AARCH64_P32_TLSLE_ADD_TPREL_LO12</span>        <a href="#R_AARCH64">R_AARCH64</a> = 110
    <span id="R_AARCH64_P32_TLSLE_ADD_TPREL_LO12_NC">R_AARCH64_P32_TLSLE_ADD_TPREL_LO12_NC</span>     <a href="#R_AARCH64">R_AARCH64</a> = 111
    <span id="R_AARCH64_P32_TLSDESC_LD_PREL19">R_AARCH64_P32_TLSDESC_LD_PREL19</span>           <a href="#R_AARCH64">R_AARCH64</a> = 122
    <span id="R_AARCH64_P32_TLSDESC_ADR_PREL21">R_AARCH64_P32_TLSDESC_ADR_PREL21</span>          <a href="#R_AARCH64">R_AARCH64</a> = 123
    <span id="R_AARCH64_P32_TLSDESC_ADR_PAGE21">R_AARCH64_P32_TLSDESC_ADR_PAGE21</span>          <a href="#R_AARCH64">R_AARCH64</a> = 124
    <span id="R_AARCH64_P32_TLSDESC_LD32_LO12_NC">R_AARCH64_P32_TLSDESC_LD32_LO12_NC</span>        <a href="#R_AARCH64">R_AARCH64</a> = 125
    <span id="R_AARCH64_P32_TLSDESC_ADD_LO12_NC">R_AARCH64_P32_TLSDESC_ADD_LO12_NC</span>         <a href="#R_AARCH64">R_AARCH64</a> = 126
    <span id="R_AARCH64_P32_TLSDESC_CALL">R_AARCH64_P32_TLSDESC_CALL</span>                <a href="#R_AARCH64">R_AARCH64</a> = 127
    <span id="R_AARCH64_P32_COPY">R_AARCH64_P32_COPY</span>                        <a href="#R_AARCH64">R_AARCH64</a> = 180
    <span id="R_AARCH64_P32_GLOB_DAT">R_AARCH64_P32_GLOB_DAT</span>                    <a href="#R_AARCH64">R_AARCH64</a> = 181
    <span id="R_AARCH64_P32_JUMP_SLOT">R_AARCH64_P32_JUMP_SLOT</span>                   <a href="#R_AARCH64">R_AARCH64</a> = 182
    <span id="R_AARCH64_P32_RELATIVE">R_AARCH64_P32_RELATIVE</span>                    <a href="#R_AARCH64">R_AARCH64</a> = 183
    <span id="R_AARCH64_P32_TLS_DTPMOD">R_AARCH64_P32_TLS_DTPMOD</span>                  <a href="#R_AARCH64">R_AARCH64</a> = 184
    <span id="R_AARCH64_P32_TLS_DTPREL">R_AARCH64_P32_TLS_DTPREL</span>                  <a href="#R_AARCH64">R_AARCH64</a> = 185
    <span id="R_AARCH64_P32_TLS_TPREL">R_AARCH64_P32_TLS_TPREL</span>                   <a href="#R_AARCH64">R_AARCH64</a> = 186
    <span id="R_AARCH64_P32_TLSDESC">R_AARCH64_P32_TLSDESC</span>                     <a href="#R_AARCH64">R_AARCH64</a> = 187
    <span id="R_AARCH64_P32_IRELATIVE">R_AARCH64_P32_IRELATIVE</span>                   <a href="#R_AARCH64">R_AARCH64</a> = 188
    <span id="R_AARCH64_NULL">R_AARCH64_NULL</span>                            <a href="#R_AARCH64">R_AARCH64</a> = 256
    <span id="R_AARCH64_ABS64">R_AARCH64_ABS64</span>                           <a href="#R_AARCH64">R_AARCH64</a> = 257
    <span id="R_AARCH64_ABS32">R_AARCH64_ABS32</span>                           <a href="#R_AARCH64">R_AARCH64</a> = 258
    <span id="R_AARCH64_ABS16">R_AARCH64_ABS16</span>                           <a href="#R_AARCH64">R_AARCH64</a> = 259
    <span id="R_AARCH64_PREL64">R_AARCH64_PREL64</span>                          <a href="#R_AARCH64">R_AARCH64</a> = 260
    <span id="R_AARCH64_PREL32">R_AARCH64_PREL32</span>                          <a href="#R_AARCH64">R_AARCH64</a> = 261
    <span id="R_AARCH64_PREL16">R_AARCH64_PREL16</span>                          <a href="#R_AARCH64">R_AARCH64</a> = 262
    <span id="R_AARCH64_MOVW_UABS_G0">R_AARCH64_MOVW_UABS_G0</span>                    <a href="#R_AARCH64">R_AARCH64</a> = 263
    <span id="R_AARCH64_MOVW_UABS_G0_NC">R_AARCH64_MOVW_UABS_G0_NC</span>                 <a href="#R_AARCH64">R_AARCH64</a> = 264
    <span id="R_AARCH64_MOVW_UABS_G1">R_AARCH64_MOVW_UABS_G1</span>                    <a href="#R_AARCH64">R_AARCH64</a> = 265
    <span id="R_AARCH64_MOVW_UABS_G1_NC">R_AARCH64_MOVW_UABS_G1_NC</span>                 <a href="#R_AARCH64">R_AARCH64</a> = 266
    <span id="R_AARCH64_MOVW_UABS_G2">R_AARCH64_MOVW_UABS_G2</span>                    <a href="#R_AARCH64">R_AARCH64</a> = 267
    <span id="R_AARCH64_MOVW_UABS_G2_NC">R_AARCH64_MOVW_UABS_G2_NC</span>                 <a href="#R_AARCH64">R_AARCH64</a> = 268
    <span id="R_AARCH64_MOVW_UABS_G3">R_AARCH64_MOVW_UABS_G3</span>                    <a href="#R_AARCH64">R_AARCH64</a> = 269
    <span id="R_AARCH64_MOVW_SABS_G0">R_AARCH64_MOVW_SABS_G0</span>                    <a href="#R_AARCH64">R_AARCH64</a> = 270
    <span id="R_AARCH64_MOVW_SABS_G1">R_AARCH64_MOVW_SABS_G1</span>                    <a href="#R_AARCH64">R_AARCH64</a> = 271
    <span id="R_AARCH64_MOVW_SABS_G2">R_AARCH64_MOVW_SABS_G2</span>                    <a href="#R_AARCH64">R_AARCH64</a> = 272
    <span id="R_AARCH64_LD_PREL_LO19">R_AARCH64_LD_PREL_LO19</span>                    <a href="#R_AARCH64">R_AARCH64</a> = 273
    <span id="R_AARCH64_ADR_PREL_LO21">R_AARCH64_ADR_PREL_LO21</span>                   <a href="#R_AARCH64">R_AARCH64</a> = 274
    <span id="R_AARCH64_ADR_PREL_PG_HI21">R_AARCH64_ADR_PREL_PG_HI21</span>                <a href="#R_AARCH64">R_AARCH64</a> = 275
    <span id="R_AARCH64_ADR_PREL_PG_HI21_NC">R_AARCH64_ADR_PREL_PG_HI21_NC</span>             <a href="#R_AARCH64">R_AARCH64</a> = 276
    <span id="R_AARCH64_ADD_ABS_LO12_NC">R_AARCH64_ADD_ABS_LO12_NC</span>                 <a href="#R_AARCH64">R_AARCH64</a> = 277
    <span id="R_AARCH64_LDST8_ABS_LO12_NC">R_AARCH64_LDST8_ABS_LO12_NC</span>               <a href="#R_AARCH64">R_AARCH64</a> = 278
    <span id="R_AARCH64_TSTBR14">R_AARCH64_TSTBR14</span>                         <a href="#R_AARCH64">R_AARCH64</a> = 279
    <span id="R_AARCH64_CONDBR19">R_AARCH64_CONDBR19</span>                        <a href="#R_AARCH64">R_AARCH64</a> = 280
    <span id="R_AARCH64_JUMP26">R_AARCH64_JUMP26</span>                          <a href="#R_AARCH64">R_AARCH64</a> = 282
    <span id="R_AARCH64_CALL26">R_AARCH64_CALL26</span>                          <a href="#R_AARCH64">R_AARCH64</a> = 283
    <span id="R_AARCH64_LDST16_ABS_LO12_NC">R_AARCH64_LDST16_ABS_LO12_NC</span>              <a href="#R_AARCH64">R_AARCH64</a> = 284
    <span id="R_AARCH64_LDST32_ABS_LO12_NC">R_AARCH64_LDST32_ABS_LO12_NC</span>              <a href="#R_AARCH64">R_AARCH64</a> = 285
    <span id="R_AARCH64_LDST64_ABS_LO12_NC">R_AARCH64_LDST64_ABS_LO12_NC</span>              <a href="#R_AARCH64">R_AARCH64</a> = 286
    <span id="R_AARCH64_LDST128_ABS_LO12_NC">R_AARCH64_LDST128_ABS_LO12_NC</span>             <a href="#R_AARCH64">R_AARCH64</a> = 299
    <span id="R_AARCH64_GOT_LD_PREL19">R_AARCH64_GOT_LD_PREL19</span>                   <a href="#R_AARCH64">R_AARCH64</a> = 309
    <span id="R_AARCH64_LD64_GOTOFF_LO15">R_AARCH64_LD64_GOTOFF_LO15</span>                <a href="#R_AARCH64">R_AARCH64</a> = 310
    <span id="R_AARCH64_ADR_GOT_PAGE">R_AARCH64_ADR_GOT_PAGE</span>                    <a href="#R_AARCH64">R_AARCH64</a> = 311
    <span id="R_AARCH64_LD64_GOT_LO12_NC">R_AARCH64_LD64_GOT_LO12_NC</span>                <a href="#R_AARCH64">R_AARCH64</a> = 312
    <span id="R_AARCH64_LD64_GOTPAGE_LO15">R_AARCH64_LD64_GOTPAGE_LO15</span>               <a href="#R_AARCH64">R_AARCH64</a> = 313
    <span id="R_AARCH64_TLSGD_ADR_PREL21">R_AARCH64_TLSGD_ADR_PREL21</span>                <a href="#R_AARCH64">R_AARCH64</a> = 512
    <span id="R_AARCH64_TLSGD_ADR_PAGE21">R_AARCH64_TLSGD_ADR_PAGE21</span>                <a href="#R_AARCH64">R_AARCH64</a> = 513
    <span id="R_AARCH64_TLSGD_ADD_LO12_NC">R_AARCH64_TLSGD_ADD_LO12_NC</span>               <a href="#R_AARCH64">R_AARCH64</a> = 514
    <span id="R_AARCH64_TLSGD_MOVW_G1">R_AARCH64_TLSGD_MOVW_G1</span>                   <a href="#R_AARCH64">R_AARCH64</a> = 515
    <span id="R_AARCH64_TLSGD_MOVW_G0_NC">R_AARCH64_TLSGD_MOVW_G0_NC</span>                <a href="#R_AARCH64">R_AARCH64</a> = 516
    <span id="R_AARCH64_TLSLD_ADR_PREL21">R_AARCH64_TLSLD_ADR_PREL21</span>                <a href="#R_AARCH64">R_AARCH64</a> = 517
    <span id="R_AARCH64_TLSLD_ADR_PAGE21">R_AARCH64_TLSLD_ADR_PAGE21</span>                <a href="#R_AARCH64">R_AARCH64</a> = 518
    <span id="R_AARCH64_TLSIE_MOVW_GOTTPREL_G1">R_AARCH64_TLSIE_MOVW_GOTTPREL_G1</span>          <a href="#R_AARCH64">R_AARCH64</a> = 539
    <span id="R_AARCH64_TLSIE_MOVW_GOTTPREL_G0_NC">R_AARCH64_TLSIE_MOVW_GOTTPREL_G0_NC</span>       <a href="#R_AARCH64">R_AARCH64</a> = 540
    <span id="R_AARCH64_TLSIE_ADR_GOTTPREL_PAGE21">R_AARCH64_TLSIE_ADR_GOTTPREL_PAGE21</span>       <a href="#R_AARCH64">R_AARCH64</a> = 541
    <span id="R_AARCH64_TLSIE_LD64_GOTTPREL_LO12_NC">R_AARCH64_TLSIE_LD64_GOTTPREL_LO12_NC</span>     <a href="#R_AARCH64">R_AARCH64</a> = 542
    <span id="R_AARCH64_TLSIE_LD_GOTTPREL_PREL19">R_AARCH64_TLSIE_LD_GOTTPREL_PREL19</span>        <a href="#R_AARCH64">R_AARCH64</a> = 543
    <span id="R_AARCH64_TLSLE_MOVW_TPREL_G2">R_AARCH64_TLSLE_MOVW_TPREL_G2</span>             <a href="#R_AARCH64">R_AARCH64</a> = 544
    <span id="R_AARCH64_TLSLE_MOVW_TPREL_G1">R_AARCH64_TLSLE_MOVW_TPREL_G1</span>             <a href="#R_AARCH64">R_AARCH64</a> = 545
    <span id="R_AARCH64_TLSLE_MOVW_TPREL_G1_NC">R_AARCH64_TLSLE_MOVW_TPREL_G1_NC</span>          <a href="#R_AARCH64">R_AARCH64</a> = 546
    <span id="R_AARCH64_TLSLE_MOVW_TPREL_G0">R_AARCH64_TLSLE_MOVW_TPREL_G0</span>             <a href="#R_AARCH64">R_AARCH64</a> = 547
    <span id="R_AARCH64_TLSLE_MOVW_TPREL_G0_NC">R_AARCH64_TLSLE_MOVW_TPREL_G0_NC</span>          <a href="#R_AARCH64">R_AARCH64</a> = 548
    <span id="R_AARCH64_TLSLE_ADD_TPREL_HI12">R_AARCH64_TLSLE_ADD_TPREL_HI12</span>            <a href="#R_AARCH64">R_AARCH64</a> = 549
    <span id="R_AARCH64_TLSLE_ADD_TPREL_LO12">R_AARCH64_TLSLE_ADD_TPREL_LO12</span>            <a href="#R_AARCH64">R_AARCH64</a> = 550
    <span id="R_AARCH64_TLSLE_ADD_TPREL_LO12_NC">R_AARCH64_TLSLE_ADD_TPREL_LO12_NC</span>         <a href="#R_AARCH64">R_AARCH64</a> = 551
    <span id="R_AARCH64_TLSDESC_LD_PREL19">R_AARCH64_TLSDESC_LD_PREL19</span>               <a href="#R_AARCH64">R_AARCH64</a> = 560
    <span id="R_AARCH64_TLSDESC_ADR_PREL21">R_AARCH64_TLSDESC_ADR_PREL21</span>              <a href="#R_AARCH64">R_AARCH64</a> = 561
    <span id="R_AARCH64_TLSDESC_ADR_PAGE21">R_AARCH64_TLSDESC_ADR_PAGE21</span>              <a href="#R_AARCH64">R_AARCH64</a> = 562
    <span id="R_AARCH64_TLSDESC_LD64_LO12_NC">R_AARCH64_TLSDESC_LD64_LO12_NC</span>            <a href="#R_AARCH64">R_AARCH64</a> = 563
    <span id="R_AARCH64_TLSDESC_ADD_LO12_NC">R_AARCH64_TLSDESC_ADD_LO12_NC</span>             <a href="#R_AARCH64">R_AARCH64</a> = 564
    <span id="R_AARCH64_TLSDESC_OFF_G1">R_AARCH64_TLSDESC_OFF_G1</span>                  <a href="#R_AARCH64">R_AARCH64</a> = 565
    <span id="R_AARCH64_TLSDESC_OFF_G0_NC">R_AARCH64_TLSDESC_OFF_G0_NC</span>               <a href="#R_AARCH64">R_AARCH64</a> = 566
    <span id="R_AARCH64_TLSDESC_LDR">R_AARCH64_TLSDESC_LDR</span>                     <a href="#R_AARCH64">R_AARCH64</a> = 567
    <span id="R_AARCH64_TLSDESC_ADD">R_AARCH64_TLSDESC_ADD</span>                     <a href="#R_AARCH64">R_AARCH64</a> = 568
    <span id="R_AARCH64_TLSDESC_CALL">R_AARCH64_TLSDESC_CALL</span>                    <a href="#R_AARCH64">R_AARCH64</a> = 569
    <span id="R_AARCH64_TLSLE_LDST128_TPREL_LO12">R_AARCH64_TLSLE_LDST128_TPREL_LO12</span>        <a href="#R_AARCH64">R_AARCH64</a> = 570
    <span id="R_AARCH64_TLSLE_LDST128_TPREL_LO12_NC">R_AARCH64_TLSLE_LDST128_TPREL_LO12_NC</span>     <a href="#R_AARCH64">R_AARCH64</a> = 571
    <span id="R_AARCH64_TLSLD_LDST128_DTPREL_LO12">R_AARCH64_TLSLD_LDST128_DTPREL_LO12</span>       <a href="#R_AARCH64">R_AARCH64</a> = 572
    <span id="R_AARCH64_TLSLD_LDST128_DTPREL_LO12_NC">R_AARCH64_TLSLD_LDST128_DTPREL_LO12_NC</span>    <a href="#R_AARCH64">R_AARCH64</a> = 573
    <span id="R_AARCH64_COPY">R_AARCH64_COPY</span>                            <a href="#R_AARCH64">R_AARCH64</a> = 1024
    <span id="R_AARCH64_GLOB_DAT">R_AARCH64_GLOB_DAT</span>                        <a href="#R_AARCH64">R_AARCH64</a> = 1025
    <span id="R_AARCH64_JUMP_SLOT">R_AARCH64_JUMP_SLOT</span>                       <a href="#R_AARCH64">R_AARCH64</a> = 1026
    <span id="R_AARCH64_RELATIVE">R_AARCH64_RELATIVE</span>                        <a href="#R_AARCH64">R_AARCH64</a> = 1027
    <span id="R_AARCH64_TLS_DTPMOD64">R_AARCH64_TLS_DTPMOD64</span>                    <a href="#R_AARCH64">R_AARCH64</a> = 1028
    <span id="R_AARCH64_TLS_DTPREL64">R_AARCH64_TLS_DTPREL64</span>                    <a href="#R_AARCH64">R_AARCH64</a> = 1029
    <span id="R_AARCH64_TLS_TPREL64">R_AARCH64_TLS_TPREL64</span>                     <a href="#R_AARCH64">R_AARCH64</a> = 1030
    <span id="R_AARCH64_TLSDESC">R_AARCH64_TLSDESC</span>                         <a href="#R_AARCH64">R_AARCH64</a> = 1031
    <span id="R_AARCH64_IRELATIVE">R_AARCH64_IRELATIVE</span>                       <a href="#R_AARCH64">R_AARCH64</a> = 1032
)</pre>









### <a id="R_AARCH64.GoString">func</a> (R\_AARCH64) [GoString](https://golang.org/src/debug/elf/elf.go?s=57801:57837#L1387)
<pre>func (i <a href="#R_AARCH64">R_AARCH64</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="R_AARCH64.String">func</a> (R\_AARCH64) [String](https://golang.org/src/debug/elf/elf.go?s=57707:57741#L1386)
<pre>func (i <a href="#R_AARCH64">R_AARCH64</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="R_ALPHA">type</a> [R_ALPHA](https://golang.org/src/debug/elf/elf.go?s=57926:57942#L1390)
Relocation types for Alpha.


<pre>type R_ALPHA <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="R_ALPHA_NONE">R_ALPHA_NONE</span>           <a href="#R_ALPHA">R_ALPHA</a> = 0  <span class="comment">/* No reloc */</span>
    <span id="R_ALPHA_REFLONG">R_ALPHA_REFLONG</span>        <a href="#R_ALPHA">R_ALPHA</a> = 1  <span class="comment">/* Direct 32 bit */</span>
    <span id="R_ALPHA_REFQUAD">R_ALPHA_REFQUAD</span>        <a href="#R_ALPHA">R_ALPHA</a> = 2  <span class="comment">/* Direct 64 bit */</span>
    <span id="R_ALPHA_GPREL32">R_ALPHA_GPREL32</span>        <a href="#R_ALPHA">R_ALPHA</a> = 3  <span class="comment">/* GP relative 32 bit */</span>
    <span id="R_ALPHA_LITERAL">R_ALPHA_LITERAL</span>        <a href="#R_ALPHA">R_ALPHA</a> = 4  <span class="comment">/* GP relative 16 bit w/optimization */</span>
    <span id="R_ALPHA_LITUSE">R_ALPHA_LITUSE</span>         <a href="#R_ALPHA">R_ALPHA</a> = 5  <span class="comment">/* Optimization hint for LITERAL */</span>
    <span id="R_ALPHA_GPDISP">R_ALPHA_GPDISP</span>         <a href="#R_ALPHA">R_ALPHA</a> = 6  <span class="comment">/* Add displacement to GP */</span>
    <span id="R_ALPHA_BRADDR">R_ALPHA_BRADDR</span>         <a href="#R_ALPHA">R_ALPHA</a> = 7  <span class="comment">/* PC+4 relative 23 bit shifted */</span>
    <span id="R_ALPHA_HINT">R_ALPHA_HINT</span>           <a href="#R_ALPHA">R_ALPHA</a> = 8  <span class="comment">/* PC+4 relative 16 bit shifted */</span>
    <span id="R_ALPHA_SREL16">R_ALPHA_SREL16</span>         <a href="#R_ALPHA">R_ALPHA</a> = 9  <span class="comment">/* PC relative 16 bit */</span>
    <span id="R_ALPHA_SREL32">R_ALPHA_SREL32</span>         <a href="#R_ALPHA">R_ALPHA</a> = 10 <span class="comment">/* PC relative 32 bit */</span>
    <span id="R_ALPHA_SREL64">R_ALPHA_SREL64</span>         <a href="#R_ALPHA">R_ALPHA</a> = 11 <span class="comment">/* PC relative 64 bit */</span>
    <span id="R_ALPHA_OP_PUSH">R_ALPHA_OP_PUSH</span>        <a href="#R_ALPHA">R_ALPHA</a> = 12 <span class="comment">/* OP stack push */</span>
    <span id="R_ALPHA_OP_STORE">R_ALPHA_OP_STORE</span>       <a href="#R_ALPHA">R_ALPHA</a> = 13 <span class="comment">/* OP stack pop and store */</span>
    <span id="R_ALPHA_OP_PSUB">R_ALPHA_OP_PSUB</span>        <a href="#R_ALPHA">R_ALPHA</a> = 14 <span class="comment">/* OP stack subtract */</span>
    <span id="R_ALPHA_OP_PRSHIFT">R_ALPHA_OP_PRSHIFT</span>     <a href="#R_ALPHA">R_ALPHA</a> = 15 <span class="comment">/* OP stack right shift */</span>
    <span id="R_ALPHA_GPVALUE">R_ALPHA_GPVALUE</span>        <a href="#R_ALPHA">R_ALPHA</a> = 16
    <span id="R_ALPHA_GPRELHIGH">R_ALPHA_GPRELHIGH</span>      <a href="#R_ALPHA">R_ALPHA</a> = 17
    <span id="R_ALPHA_GPRELLOW">R_ALPHA_GPRELLOW</span>       <a href="#R_ALPHA">R_ALPHA</a> = 18
    <span id="R_ALPHA_IMMED_GP_16">R_ALPHA_IMMED_GP_16</span>    <a href="#R_ALPHA">R_ALPHA</a> = 19
    <span id="R_ALPHA_IMMED_GP_HI32">R_ALPHA_IMMED_GP_HI32</span>  <a href="#R_ALPHA">R_ALPHA</a> = 20
    <span id="R_ALPHA_IMMED_SCN_HI32">R_ALPHA_IMMED_SCN_HI32</span> <a href="#R_ALPHA">R_ALPHA</a> = 21
    <span id="R_ALPHA_IMMED_BR_HI32">R_ALPHA_IMMED_BR_HI32</span>  <a href="#R_ALPHA">R_ALPHA</a> = 22
    <span id="R_ALPHA_IMMED_LO32">R_ALPHA_IMMED_LO32</span>     <a href="#R_ALPHA">R_ALPHA</a> = 23
    <span id="R_ALPHA_COPY">R_ALPHA_COPY</span>           <a href="#R_ALPHA">R_ALPHA</a> = 24 <span class="comment">/* Copy symbol at runtime */</span>
    <span id="R_ALPHA_GLOB_DAT">R_ALPHA_GLOB_DAT</span>       <a href="#R_ALPHA">R_ALPHA</a> = 25 <span class="comment">/* Create GOT entry */</span>
    <span id="R_ALPHA_JMP_SLOT">R_ALPHA_JMP_SLOT</span>       <a href="#R_ALPHA">R_ALPHA</a> = 26 <span class="comment">/* Create PLT entry */</span>
    <span id="R_ALPHA_RELATIVE">R_ALPHA_RELATIVE</span>       <a href="#R_ALPHA">R_ALPHA</a> = 27 <span class="comment">/* Adjust by program base */</span>
)</pre>









### <a id="R_ALPHA.GoString">func</a> (R\_ALPHA) [GoString](https://golang.org/src/debug/elf/elf.go?s=60388:60422#L1455)
<pre>func (i <a href="#R_ALPHA">R_ALPHA</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="R_ALPHA.String">func</a> (R\_ALPHA) [String](https://golang.org/src/debug/elf/elf.go?s=60298:60330#L1454)
<pre>func (i <a href="#R_ALPHA">R_ALPHA</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="R_ARM">type</a> [R_ARM](https://golang.org/src/debug/elf/elf.go?s=60507:60521#L1458)
Relocation types for ARM.


<pre>type R_ARM <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="R_ARM_NONE">R_ARM_NONE</span>               <a href="#R_ARM">R_ARM</a> = 0 <span class="comment">/* No relocation. */</span>
    <span id="R_ARM_PC24">R_ARM_PC24</span>               <a href="#R_ARM">R_ARM</a> = 1
    <span id="R_ARM_ABS32">R_ARM_ABS32</span>              <a href="#R_ARM">R_ARM</a> = 2
    <span id="R_ARM_REL32">R_ARM_REL32</span>              <a href="#R_ARM">R_ARM</a> = 3
    <span id="R_ARM_PC13">R_ARM_PC13</span>               <a href="#R_ARM">R_ARM</a> = 4
    <span id="R_ARM_ABS16">R_ARM_ABS16</span>              <a href="#R_ARM">R_ARM</a> = 5
    <span id="R_ARM_ABS12">R_ARM_ABS12</span>              <a href="#R_ARM">R_ARM</a> = 6
    <span id="R_ARM_THM_ABS5">R_ARM_THM_ABS5</span>           <a href="#R_ARM">R_ARM</a> = 7
    <span id="R_ARM_ABS8">R_ARM_ABS8</span>               <a href="#R_ARM">R_ARM</a> = 8
    <span id="R_ARM_SBREL32">R_ARM_SBREL32</span>            <a href="#R_ARM">R_ARM</a> = 9
    <span id="R_ARM_THM_PC22">R_ARM_THM_PC22</span>           <a href="#R_ARM">R_ARM</a> = 10
    <span id="R_ARM_THM_PC8">R_ARM_THM_PC8</span>            <a href="#R_ARM">R_ARM</a> = 11
    <span id="R_ARM_AMP_VCALL9">R_ARM_AMP_VCALL9</span>         <a href="#R_ARM">R_ARM</a> = 12
    <span id="R_ARM_SWI24">R_ARM_SWI24</span>              <a href="#R_ARM">R_ARM</a> = 13
    <span id="R_ARM_THM_SWI8">R_ARM_THM_SWI8</span>           <a href="#R_ARM">R_ARM</a> = 14
    <span id="R_ARM_XPC25">R_ARM_XPC25</span>              <a href="#R_ARM">R_ARM</a> = 15
    <span id="R_ARM_THM_XPC22">R_ARM_THM_XPC22</span>          <a href="#R_ARM">R_ARM</a> = 16
    <span id="R_ARM_TLS_DTPMOD32">R_ARM_TLS_DTPMOD32</span>       <a href="#R_ARM">R_ARM</a> = 17
    <span id="R_ARM_TLS_DTPOFF32">R_ARM_TLS_DTPOFF32</span>       <a href="#R_ARM">R_ARM</a> = 18
    <span id="R_ARM_TLS_TPOFF32">R_ARM_TLS_TPOFF32</span>        <a href="#R_ARM">R_ARM</a> = 19
    <span id="R_ARM_COPY">R_ARM_COPY</span>               <a href="#R_ARM">R_ARM</a> = 20 <span class="comment">/* Copy data from shared object. */</span>
    <span id="R_ARM_GLOB_DAT">R_ARM_GLOB_DAT</span>           <a href="#R_ARM">R_ARM</a> = 21 <span class="comment">/* Set GOT entry to data address. */</span>
    <span id="R_ARM_JUMP_SLOT">R_ARM_JUMP_SLOT</span>          <a href="#R_ARM">R_ARM</a> = 22 <span class="comment">/* Set GOT entry to code address. */</span>
    <span id="R_ARM_RELATIVE">R_ARM_RELATIVE</span>           <a href="#R_ARM">R_ARM</a> = 23 <span class="comment">/* Add load address of shared object. */</span>
    <span id="R_ARM_GOTOFF">R_ARM_GOTOFF</span>             <a href="#R_ARM">R_ARM</a> = 24 <span class="comment">/* Add GOT-relative symbol address. */</span>
    <span id="R_ARM_GOTPC">R_ARM_GOTPC</span>              <a href="#R_ARM">R_ARM</a> = 25 <span class="comment">/* Add PC-relative GOT table address. */</span>
    <span id="R_ARM_GOT32">R_ARM_GOT32</span>              <a href="#R_ARM">R_ARM</a> = 26 <span class="comment">/* Add PC-relative GOT offset. */</span>
    <span id="R_ARM_PLT32">R_ARM_PLT32</span>              <a href="#R_ARM">R_ARM</a> = 27 <span class="comment">/* Add PC-relative PLT offset. */</span>
    <span id="R_ARM_CALL">R_ARM_CALL</span>               <a href="#R_ARM">R_ARM</a> = 28
    <span id="R_ARM_JUMP24">R_ARM_JUMP24</span>             <a href="#R_ARM">R_ARM</a> = 29
    <span id="R_ARM_THM_JUMP24">R_ARM_THM_JUMP24</span>         <a href="#R_ARM">R_ARM</a> = 30
    <span id="R_ARM_BASE_ABS">R_ARM_BASE_ABS</span>           <a href="#R_ARM">R_ARM</a> = 31
    <span id="R_ARM_ALU_PCREL_7_0">R_ARM_ALU_PCREL_7_0</span>      <a href="#R_ARM">R_ARM</a> = 32
    <span id="R_ARM_ALU_PCREL_15_8">R_ARM_ALU_PCREL_15_8</span>     <a href="#R_ARM">R_ARM</a> = 33
    <span id="R_ARM_ALU_PCREL_23_15">R_ARM_ALU_PCREL_23_15</span>    <a href="#R_ARM">R_ARM</a> = 34
    <span id="R_ARM_LDR_SBREL_11_10_NC">R_ARM_LDR_SBREL_11_10_NC</span> <a href="#R_ARM">R_ARM</a> = 35
    <span id="R_ARM_ALU_SBREL_19_12_NC">R_ARM_ALU_SBREL_19_12_NC</span> <a href="#R_ARM">R_ARM</a> = 36
    <span id="R_ARM_ALU_SBREL_27_20_CK">R_ARM_ALU_SBREL_27_20_CK</span> <a href="#R_ARM">R_ARM</a> = 37
    <span id="R_ARM_TARGET1">R_ARM_TARGET1</span>            <a href="#R_ARM">R_ARM</a> = 38
    <span id="R_ARM_SBREL31">R_ARM_SBREL31</span>            <a href="#R_ARM">R_ARM</a> = 39
    <span id="R_ARM_V4BX">R_ARM_V4BX</span>               <a href="#R_ARM">R_ARM</a> = 40
    <span id="R_ARM_TARGET2">R_ARM_TARGET2</span>            <a href="#R_ARM">R_ARM</a> = 41
    <span id="R_ARM_PREL31">R_ARM_PREL31</span>             <a href="#R_ARM">R_ARM</a> = 42
    <span id="R_ARM_MOVW_ABS_NC">R_ARM_MOVW_ABS_NC</span>        <a href="#R_ARM">R_ARM</a> = 43
    <span id="R_ARM_MOVT_ABS">R_ARM_MOVT_ABS</span>           <a href="#R_ARM">R_ARM</a> = 44
    <span id="R_ARM_MOVW_PREL_NC">R_ARM_MOVW_PREL_NC</span>       <a href="#R_ARM">R_ARM</a> = 45
    <span id="R_ARM_MOVT_PREL">R_ARM_MOVT_PREL</span>          <a href="#R_ARM">R_ARM</a> = 46
    <span id="R_ARM_THM_MOVW_ABS_NC">R_ARM_THM_MOVW_ABS_NC</span>    <a href="#R_ARM">R_ARM</a> = 47
    <span id="R_ARM_THM_MOVT_ABS">R_ARM_THM_MOVT_ABS</span>       <a href="#R_ARM">R_ARM</a> = 48
    <span id="R_ARM_THM_MOVW_PREL_NC">R_ARM_THM_MOVW_PREL_NC</span>   <a href="#R_ARM">R_ARM</a> = 49
    <span id="R_ARM_THM_MOVT_PREL">R_ARM_THM_MOVT_PREL</span>      <a href="#R_ARM">R_ARM</a> = 50
    <span id="R_ARM_THM_JUMP19">R_ARM_THM_JUMP19</span>         <a href="#R_ARM">R_ARM</a> = 51
    <span id="R_ARM_THM_JUMP6">R_ARM_THM_JUMP6</span>          <a href="#R_ARM">R_ARM</a> = 52
    <span id="R_ARM_THM_ALU_PREL_11_0">R_ARM_THM_ALU_PREL_11_0</span>  <a href="#R_ARM">R_ARM</a> = 53
    <span id="R_ARM_THM_PC12">R_ARM_THM_PC12</span>           <a href="#R_ARM">R_ARM</a> = 54
    <span id="R_ARM_ABS32_NOI">R_ARM_ABS32_NOI</span>          <a href="#R_ARM">R_ARM</a> = 55
    <span id="R_ARM_REL32_NOI">R_ARM_REL32_NOI</span>          <a href="#R_ARM">R_ARM</a> = 56
    <span id="R_ARM_ALU_PC_G0_NC">R_ARM_ALU_PC_G0_NC</span>       <a href="#R_ARM">R_ARM</a> = 57
    <span id="R_ARM_ALU_PC_G0">R_ARM_ALU_PC_G0</span>          <a href="#R_ARM">R_ARM</a> = 58
    <span id="R_ARM_ALU_PC_G1_NC">R_ARM_ALU_PC_G1_NC</span>       <a href="#R_ARM">R_ARM</a> = 59
    <span id="R_ARM_ALU_PC_G1">R_ARM_ALU_PC_G1</span>          <a href="#R_ARM">R_ARM</a> = 60
    <span id="R_ARM_ALU_PC_G2">R_ARM_ALU_PC_G2</span>          <a href="#R_ARM">R_ARM</a> = 61
    <span id="R_ARM_LDR_PC_G1">R_ARM_LDR_PC_G1</span>          <a href="#R_ARM">R_ARM</a> = 62
    <span id="R_ARM_LDR_PC_G2">R_ARM_LDR_PC_G2</span>          <a href="#R_ARM">R_ARM</a> = 63
    <span id="R_ARM_LDRS_PC_G0">R_ARM_LDRS_PC_G0</span>         <a href="#R_ARM">R_ARM</a> = 64
    <span id="R_ARM_LDRS_PC_G1">R_ARM_LDRS_PC_G1</span>         <a href="#R_ARM">R_ARM</a> = 65
    <span id="R_ARM_LDRS_PC_G2">R_ARM_LDRS_PC_G2</span>         <a href="#R_ARM">R_ARM</a> = 66
    <span id="R_ARM_LDC_PC_G0">R_ARM_LDC_PC_G0</span>          <a href="#R_ARM">R_ARM</a> = 67
    <span id="R_ARM_LDC_PC_G1">R_ARM_LDC_PC_G1</span>          <a href="#R_ARM">R_ARM</a> = 68
    <span id="R_ARM_LDC_PC_G2">R_ARM_LDC_PC_G2</span>          <a href="#R_ARM">R_ARM</a> = 69
    <span id="R_ARM_ALU_SB_G0_NC">R_ARM_ALU_SB_G0_NC</span>       <a href="#R_ARM">R_ARM</a> = 70
    <span id="R_ARM_ALU_SB_G0">R_ARM_ALU_SB_G0</span>          <a href="#R_ARM">R_ARM</a> = 71
    <span id="R_ARM_ALU_SB_G1_NC">R_ARM_ALU_SB_G1_NC</span>       <a href="#R_ARM">R_ARM</a> = 72
    <span id="R_ARM_ALU_SB_G1">R_ARM_ALU_SB_G1</span>          <a href="#R_ARM">R_ARM</a> = 73
    <span id="R_ARM_ALU_SB_G2">R_ARM_ALU_SB_G2</span>          <a href="#R_ARM">R_ARM</a> = 74
    <span id="R_ARM_LDR_SB_G0">R_ARM_LDR_SB_G0</span>          <a href="#R_ARM">R_ARM</a> = 75
    <span id="R_ARM_LDR_SB_G1">R_ARM_LDR_SB_G1</span>          <a href="#R_ARM">R_ARM</a> = 76
    <span id="R_ARM_LDR_SB_G2">R_ARM_LDR_SB_G2</span>          <a href="#R_ARM">R_ARM</a> = 77
    <span id="R_ARM_LDRS_SB_G0">R_ARM_LDRS_SB_G0</span>         <a href="#R_ARM">R_ARM</a> = 78
    <span id="R_ARM_LDRS_SB_G1">R_ARM_LDRS_SB_G1</span>         <a href="#R_ARM">R_ARM</a> = 79
    <span id="R_ARM_LDRS_SB_G2">R_ARM_LDRS_SB_G2</span>         <a href="#R_ARM">R_ARM</a> = 80
    <span id="R_ARM_LDC_SB_G0">R_ARM_LDC_SB_G0</span>          <a href="#R_ARM">R_ARM</a> = 81
    <span id="R_ARM_LDC_SB_G1">R_ARM_LDC_SB_G1</span>          <a href="#R_ARM">R_ARM</a> = 82
    <span id="R_ARM_LDC_SB_G2">R_ARM_LDC_SB_G2</span>          <a href="#R_ARM">R_ARM</a> = 83
    <span id="R_ARM_MOVW_BREL_NC">R_ARM_MOVW_BREL_NC</span>       <a href="#R_ARM">R_ARM</a> = 84
    <span id="R_ARM_MOVT_BREL">R_ARM_MOVT_BREL</span>          <a href="#R_ARM">R_ARM</a> = 85
    <span id="R_ARM_MOVW_BREL">R_ARM_MOVW_BREL</span>          <a href="#R_ARM">R_ARM</a> = 86
    <span id="R_ARM_THM_MOVW_BREL_NC">R_ARM_THM_MOVW_BREL_NC</span>   <a href="#R_ARM">R_ARM</a> = 87
    <span id="R_ARM_THM_MOVT_BREL">R_ARM_THM_MOVT_BREL</span>      <a href="#R_ARM">R_ARM</a> = 88
    <span id="R_ARM_THM_MOVW_BREL">R_ARM_THM_MOVW_BREL</span>      <a href="#R_ARM">R_ARM</a> = 89
    <span id="R_ARM_TLS_GOTDESC">R_ARM_TLS_GOTDESC</span>        <a href="#R_ARM">R_ARM</a> = 90
    <span id="R_ARM_TLS_CALL">R_ARM_TLS_CALL</span>           <a href="#R_ARM">R_ARM</a> = 91
    <span id="R_ARM_TLS_DESCSEQ">R_ARM_TLS_DESCSEQ</span>        <a href="#R_ARM">R_ARM</a> = 92
    <span id="R_ARM_THM_TLS_CALL">R_ARM_THM_TLS_CALL</span>       <a href="#R_ARM">R_ARM</a> = 93
    <span id="R_ARM_PLT32_ABS">R_ARM_PLT32_ABS</span>          <a href="#R_ARM">R_ARM</a> = 94
    <span id="R_ARM_GOT_ABS">R_ARM_GOT_ABS</span>            <a href="#R_ARM">R_ARM</a> = 95
    <span id="R_ARM_GOT_PREL">R_ARM_GOT_PREL</span>           <a href="#R_ARM">R_ARM</a> = 96
    <span id="R_ARM_GOT_BREL12">R_ARM_GOT_BREL12</span>         <a href="#R_ARM">R_ARM</a> = 97
    <span id="R_ARM_GOTOFF12">R_ARM_GOTOFF12</span>           <a href="#R_ARM">R_ARM</a> = 98
    <span id="R_ARM_GOTRELAX">R_ARM_GOTRELAX</span>           <a href="#R_ARM">R_ARM</a> = 99
    <span id="R_ARM_GNU_VTENTRY">R_ARM_GNU_VTENTRY</span>        <a href="#R_ARM">R_ARM</a> = 100
    <span id="R_ARM_GNU_VTINHERIT">R_ARM_GNU_VTINHERIT</span>      <a href="#R_ARM">R_ARM</a> = 101
    <span id="R_ARM_THM_JUMP11">R_ARM_THM_JUMP11</span>         <a href="#R_ARM">R_ARM</a> = 102
    <span id="R_ARM_THM_JUMP8">R_ARM_THM_JUMP8</span>          <a href="#R_ARM">R_ARM</a> = 103
    <span id="R_ARM_TLS_GD32">R_ARM_TLS_GD32</span>           <a href="#R_ARM">R_ARM</a> = 104
    <span id="R_ARM_TLS_LDM32">R_ARM_TLS_LDM32</span>          <a href="#R_ARM">R_ARM</a> = 105
    <span id="R_ARM_TLS_LDO32">R_ARM_TLS_LDO32</span>          <a href="#R_ARM">R_ARM</a> = 106
    <span id="R_ARM_TLS_IE32">R_ARM_TLS_IE32</span>           <a href="#R_ARM">R_ARM</a> = 107
    <span id="R_ARM_TLS_LE32">R_ARM_TLS_LE32</span>           <a href="#R_ARM">R_ARM</a> = 108
    <span id="R_ARM_TLS_LDO12">R_ARM_TLS_LDO12</span>          <a href="#R_ARM">R_ARM</a> = 109
    <span id="R_ARM_TLS_LE12">R_ARM_TLS_LE12</span>           <a href="#R_ARM">R_ARM</a> = 110
    <span id="R_ARM_TLS_IE12GP">R_ARM_TLS_IE12GP</span>         <a href="#R_ARM">R_ARM</a> = 111
    <span id="R_ARM_PRIVATE_0">R_ARM_PRIVATE_0</span>          <a href="#R_ARM">R_ARM</a> = 112
    <span id="R_ARM_PRIVATE_1">R_ARM_PRIVATE_1</span>          <a href="#R_ARM">R_ARM</a> = 113
    <span id="R_ARM_PRIVATE_2">R_ARM_PRIVATE_2</span>          <a href="#R_ARM">R_ARM</a> = 114
    <span id="R_ARM_PRIVATE_3">R_ARM_PRIVATE_3</span>          <a href="#R_ARM">R_ARM</a> = 115
    <span id="R_ARM_PRIVATE_4">R_ARM_PRIVATE_4</span>          <a href="#R_ARM">R_ARM</a> = 116
    <span id="R_ARM_PRIVATE_5">R_ARM_PRIVATE_5</span>          <a href="#R_ARM">R_ARM</a> = 117
    <span id="R_ARM_PRIVATE_6">R_ARM_PRIVATE_6</span>          <a href="#R_ARM">R_ARM</a> = 118
    <span id="R_ARM_PRIVATE_7">R_ARM_PRIVATE_7</span>          <a href="#R_ARM">R_ARM</a> = 119
    <span id="R_ARM_PRIVATE_8">R_ARM_PRIVATE_8</span>          <a href="#R_ARM">R_ARM</a> = 120
    <span id="R_ARM_PRIVATE_9">R_ARM_PRIVATE_9</span>          <a href="#R_ARM">R_ARM</a> = 121
    <span id="R_ARM_PRIVATE_10">R_ARM_PRIVATE_10</span>         <a href="#R_ARM">R_ARM</a> = 122
    <span id="R_ARM_PRIVATE_11">R_ARM_PRIVATE_11</span>         <a href="#R_ARM">R_ARM</a> = 123
    <span id="R_ARM_PRIVATE_12">R_ARM_PRIVATE_12</span>         <a href="#R_ARM">R_ARM</a> = 124
    <span id="R_ARM_PRIVATE_13">R_ARM_PRIVATE_13</span>         <a href="#R_ARM">R_ARM</a> = 125
    <span id="R_ARM_PRIVATE_14">R_ARM_PRIVATE_14</span>         <a href="#R_ARM">R_ARM</a> = 126
    <span id="R_ARM_PRIVATE_15">R_ARM_PRIVATE_15</span>         <a href="#R_ARM">R_ARM</a> = 127
    <span id="R_ARM_ME_TOO">R_ARM_ME_TOO</span>             <a href="#R_ARM">R_ARM</a> = 128
    <span id="R_ARM_THM_TLS_DESCSEQ16">R_ARM_THM_TLS_DESCSEQ16</span>  <a href="#R_ARM">R_ARM</a> = 129
    <span id="R_ARM_THM_TLS_DESCSEQ32">R_ARM_THM_TLS_DESCSEQ32</span>  <a href="#R_ARM">R_ARM</a> = 130
    <span id="R_ARM_THM_GOT_BREL12">R_ARM_THM_GOT_BREL12</span>     <a href="#R_ARM">R_ARM</a> = 131
    <span id="R_ARM_THM_ALU_ABS_G0_NC">R_ARM_THM_ALU_ABS_G0_NC</span>  <a href="#R_ARM">R_ARM</a> = 132
    <span id="R_ARM_THM_ALU_ABS_G1_NC">R_ARM_THM_ALU_ABS_G1_NC</span>  <a href="#R_ARM">R_ARM</a> = 133
    <span id="R_ARM_THM_ALU_ABS_G2_NC">R_ARM_THM_ALU_ABS_G2_NC</span>  <a href="#R_ARM">R_ARM</a> = 134
    <span id="R_ARM_THM_ALU_ABS_G3">R_ARM_THM_ALU_ABS_G3</span>     <a href="#R_ARM">R_ARM</a> = 135
    <span id="R_ARM_IRELATIVE">R_ARM_IRELATIVE</span>          <a href="#R_ARM">R_ARM</a> = 160
    <span id="R_ARM_RXPC25">R_ARM_RXPC25</span>             <a href="#R_ARM">R_ARM</a> = 249
    <span id="R_ARM_RSBREL32">R_ARM_RSBREL32</span>           <a href="#R_ARM">R_ARM</a> = 250
    <span id="R_ARM_THM_RPC22">R_ARM_THM_RPC22</span>          <a href="#R_ARM">R_ARM</a> = 251
    <span id="R_ARM_RREL32">R_ARM_RREL32</span>             <a href="#R_ARM">R_ARM</a> = 252
    <span id="R_ARM_RABS32">R_ARM_RABS32</span>             <a href="#R_ARM">R_ARM</a> = 253
    <span id="R_ARM_RPC24">R_ARM_RPC24</span>              <a href="#R_ARM">R_ARM</a> = 254
    <span id="R_ARM_RBASE">R_ARM_RBASE</span>              <a href="#R_ARM">R_ARM</a> = 255
)</pre>









### <a id="R_ARM.GoString">func</a> (R\_ARM) [GoString](https://golang.org/src/debug/elf/elf.go?s=70177:70209#L1755)
<pre>func (i <a href="#R_ARM">R_ARM</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="R_ARM.String">func</a> (R\_ARM) [String](https://golang.org/src/debug/elf/elf.go?s=70091:70121#L1754)
<pre>func (i <a href="#R_ARM">R_ARM</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="R_MIPS">type</a> [R_MIPS](https://golang.org/src/debug/elf/elf.go?s=74199:74214#L1854)
Relocation types for MIPS.


<pre>type R_MIPS <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="R_MIPS_NONE">R_MIPS_NONE</span>          <a href="#R_MIPS">R_MIPS</a> = 0
    <span id="R_MIPS_16">R_MIPS_16</span>            <a href="#R_MIPS">R_MIPS</a> = 1
    <span id="R_MIPS_32">R_MIPS_32</span>            <a href="#R_MIPS">R_MIPS</a> = 2
    <span id="R_MIPS_REL32">R_MIPS_REL32</span>         <a href="#R_MIPS">R_MIPS</a> = 3
    <span id="R_MIPS_26">R_MIPS_26</span>            <a href="#R_MIPS">R_MIPS</a> = 4
    <span id="R_MIPS_HI16">R_MIPS_HI16</span>          <a href="#R_MIPS">R_MIPS</a> = 5  <span class="comment">/* high 16 bits of symbol value */</span>
    <span id="R_MIPS_LO16">R_MIPS_LO16</span>          <a href="#R_MIPS">R_MIPS</a> = 6  <span class="comment">/* low 16 bits of symbol value */</span>
    <span id="R_MIPS_GPREL16">R_MIPS_GPREL16</span>       <a href="#R_MIPS">R_MIPS</a> = 7  <span class="comment">/* GP-relative reference  */</span>
    <span id="R_MIPS_LITERAL">R_MIPS_LITERAL</span>       <a href="#R_MIPS">R_MIPS</a> = 8  <span class="comment">/* Reference to literal section  */</span>
    <span id="R_MIPS_GOT16">R_MIPS_GOT16</span>         <a href="#R_MIPS">R_MIPS</a> = 9  <span class="comment">/* Reference to global offset table */</span>
    <span id="R_MIPS_PC16">R_MIPS_PC16</span>          <a href="#R_MIPS">R_MIPS</a> = 10 <span class="comment">/* 16 bit PC relative reference */</span>
    <span id="R_MIPS_CALL16">R_MIPS_CALL16</span>        <a href="#R_MIPS">R_MIPS</a> = 11 <span class="comment">/* 16 bit call through glbl offset tbl */</span>
    <span id="R_MIPS_GPREL32">R_MIPS_GPREL32</span>       <a href="#R_MIPS">R_MIPS</a> = 12
    <span id="R_MIPS_SHIFT5">R_MIPS_SHIFT5</span>        <a href="#R_MIPS">R_MIPS</a> = 16
    <span id="R_MIPS_SHIFT6">R_MIPS_SHIFT6</span>        <a href="#R_MIPS">R_MIPS</a> = 17
    <span id="R_MIPS_64">R_MIPS_64</span>            <a href="#R_MIPS">R_MIPS</a> = 18
    <span id="R_MIPS_GOT_DISP">R_MIPS_GOT_DISP</span>      <a href="#R_MIPS">R_MIPS</a> = 19
    <span id="R_MIPS_GOT_PAGE">R_MIPS_GOT_PAGE</span>      <a href="#R_MIPS">R_MIPS</a> = 20
    <span id="R_MIPS_GOT_OFST">R_MIPS_GOT_OFST</span>      <a href="#R_MIPS">R_MIPS</a> = 21
    <span id="R_MIPS_GOT_HI16">R_MIPS_GOT_HI16</span>      <a href="#R_MIPS">R_MIPS</a> = 22
    <span id="R_MIPS_GOT_LO16">R_MIPS_GOT_LO16</span>      <a href="#R_MIPS">R_MIPS</a> = 23
    <span id="R_MIPS_SUB">R_MIPS_SUB</span>           <a href="#R_MIPS">R_MIPS</a> = 24
    <span id="R_MIPS_INSERT_A">R_MIPS_INSERT_A</span>      <a href="#R_MIPS">R_MIPS</a> = 25
    <span id="R_MIPS_INSERT_B">R_MIPS_INSERT_B</span>      <a href="#R_MIPS">R_MIPS</a> = 26
    <span id="R_MIPS_DELETE">R_MIPS_DELETE</span>        <a href="#R_MIPS">R_MIPS</a> = 27
    <span id="R_MIPS_HIGHER">R_MIPS_HIGHER</span>        <a href="#R_MIPS">R_MIPS</a> = 28
    <span id="R_MIPS_HIGHEST">R_MIPS_HIGHEST</span>       <a href="#R_MIPS">R_MIPS</a> = 29
    <span id="R_MIPS_CALL_HI16">R_MIPS_CALL_HI16</span>     <a href="#R_MIPS">R_MIPS</a> = 30
    <span id="R_MIPS_CALL_LO16">R_MIPS_CALL_LO16</span>     <a href="#R_MIPS">R_MIPS</a> = 31
    <span id="R_MIPS_SCN_DISP">R_MIPS_SCN_DISP</span>      <a href="#R_MIPS">R_MIPS</a> = 32
    <span id="R_MIPS_REL16">R_MIPS_REL16</span>         <a href="#R_MIPS">R_MIPS</a> = 33
    <span id="R_MIPS_ADD_IMMEDIATE">R_MIPS_ADD_IMMEDIATE</span> <a href="#R_MIPS">R_MIPS</a> = 34
    <span id="R_MIPS_PJUMP">R_MIPS_PJUMP</span>         <a href="#R_MIPS">R_MIPS</a> = 35
    <span id="R_MIPS_RELGOT">R_MIPS_RELGOT</span>        <a href="#R_MIPS">R_MIPS</a> = 36
    <span id="R_MIPS_JALR">R_MIPS_JALR</span>          <a href="#R_MIPS">R_MIPS</a> = 37

    <span id="R_MIPS_TLS_DTPMOD32">R_MIPS_TLS_DTPMOD32</span>    <a href="#R_MIPS">R_MIPS</a> = 38 <span class="comment">/* Module number 32 bit */</span>
    <span id="R_MIPS_TLS_DTPREL32">R_MIPS_TLS_DTPREL32</span>    <a href="#R_MIPS">R_MIPS</a> = 39 <span class="comment">/* Module-relative offset 32 bit */</span>
    <span id="R_MIPS_TLS_DTPMOD64">R_MIPS_TLS_DTPMOD64</span>    <a href="#R_MIPS">R_MIPS</a> = 40 <span class="comment">/* Module number 64 bit */</span>
    <span id="R_MIPS_TLS_DTPREL64">R_MIPS_TLS_DTPREL64</span>    <a href="#R_MIPS">R_MIPS</a> = 41 <span class="comment">/* Module-relative offset 64 bit */</span>
    <span id="R_MIPS_TLS_GD">R_MIPS_TLS_GD</span>          <a href="#R_MIPS">R_MIPS</a> = 42 <span class="comment">/* 16 bit GOT offset for GD */</span>
    <span id="R_MIPS_TLS_LDM">R_MIPS_TLS_LDM</span>         <a href="#R_MIPS">R_MIPS</a> = 43 <span class="comment">/* 16 bit GOT offset for LDM */</span>
    <span id="R_MIPS_TLS_DTPREL_HI16">R_MIPS_TLS_DTPREL_HI16</span> <a href="#R_MIPS">R_MIPS</a> = 44 <span class="comment">/* Module-relative offset, high 16 bits */</span>
    <span id="R_MIPS_TLS_DTPREL_LO16">R_MIPS_TLS_DTPREL_LO16</span> <a href="#R_MIPS">R_MIPS</a> = 45 <span class="comment">/* Module-relative offset, low 16 bits */</span>
    <span id="R_MIPS_TLS_GOTTPREL">R_MIPS_TLS_GOTTPREL</span>    <a href="#R_MIPS">R_MIPS</a> = 46 <span class="comment">/* 16 bit GOT offset for IE */</span>
    <span id="R_MIPS_TLS_TPREL32">R_MIPS_TLS_TPREL32</span>     <a href="#R_MIPS">R_MIPS</a> = 47 <span class="comment">/* TP-relative offset, 32 bit */</span>
    <span id="R_MIPS_TLS_TPREL64">R_MIPS_TLS_TPREL64</span>     <a href="#R_MIPS">R_MIPS</a> = 48 <span class="comment">/* TP-relative offset, 64 bit */</span>
    <span id="R_MIPS_TLS_TPREL_HI16">R_MIPS_TLS_TPREL_HI16</span>  <a href="#R_MIPS">R_MIPS</a> = 49 <span class="comment">/* TP-relative offset, high 16 bits */</span>
    <span id="R_MIPS_TLS_TPREL_LO16">R_MIPS_TLS_TPREL_LO16</span>  <a href="#R_MIPS">R_MIPS</a> = 50 <span class="comment">/* TP-relative offset, low 16 bits */</span>
)</pre>









### <a id="R_MIPS.GoString">func</a> (R\_MIPS) [GoString](https://golang.org/src/debug/elf/elf.go?s=77917:77950#L1960)
<pre>func (i <a href="#R_MIPS">R_MIPS</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="R_MIPS.String">func</a> (R\_MIPS) [String](https://golang.org/src/debug/elf/elf.go?s=77829:77860#L1959)
<pre>func (i <a href="#R_MIPS">R_MIPS</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="R_PPC">type</a> [R_PPC](https://golang.org/src/debug/elf/elf.go?s=78296:78310#L1968)
Relocation types for PowerPC.

Values that are shared by both R_PPC and R_PPC64 are prefixed with
R_POWERPC_ in the ELF standard. For the R_PPC type, the relevant
shared relocations have been renamed with the prefix R_PPC_.
The original name follows the value in a comment.


<pre>type R_PPC <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="R_PPC_NONE">R_PPC_NONE</span>            <a href="#R_PPC">R_PPC</a> = 0  <span class="comment">// R_POWERPC_NONE</span>
    <span id="R_PPC_ADDR32">R_PPC_ADDR32</span>          <a href="#R_PPC">R_PPC</a> = 1  <span class="comment">// R_POWERPC_ADDR32</span>
    <span id="R_PPC_ADDR24">R_PPC_ADDR24</span>          <a href="#R_PPC">R_PPC</a> = 2  <span class="comment">// R_POWERPC_ADDR24</span>
    <span id="R_PPC_ADDR16">R_PPC_ADDR16</span>          <a href="#R_PPC">R_PPC</a> = 3  <span class="comment">// R_POWERPC_ADDR16</span>
    <span id="R_PPC_ADDR16_LO">R_PPC_ADDR16_LO</span>       <a href="#R_PPC">R_PPC</a> = 4  <span class="comment">// R_POWERPC_ADDR16_LO</span>
    <span id="R_PPC_ADDR16_HI">R_PPC_ADDR16_HI</span>       <a href="#R_PPC">R_PPC</a> = 5  <span class="comment">// R_POWERPC_ADDR16_HI</span>
    <span id="R_PPC_ADDR16_HA">R_PPC_ADDR16_HA</span>       <a href="#R_PPC">R_PPC</a> = 6  <span class="comment">// R_POWERPC_ADDR16_HA</span>
    <span id="R_PPC_ADDR14">R_PPC_ADDR14</span>          <a href="#R_PPC">R_PPC</a> = 7  <span class="comment">// R_POWERPC_ADDR14</span>
    <span id="R_PPC_ADDR14_BRTAKEN">R_PPC_ADDR14_BRTAKEN</span>  <a href="#R_PPC">R_PPC</a> = 8  <span class="comment">// R_POWERPC_ADDR14_BRTAKEN</span>
    <span id="R_PPC_ADDR14_BRNTAKEN">R_PPC_ADDR14_BRNTAKEN</span> <a href="#R_PPC">R_PPC</a> = 9  <span class="comment">// R_POWERPC_ADDR14_BRNTAKEN</span>
    <span id="R_PPC_REL24">R_PPC_REL24</span>           <a href="#R_PPC">R_PPC</a> = 10 <span class="comment">// R_POWERPC_REL24</span>
    <span id="R_PPC_REL14">R_PPC_REL14</span>           <a href="#R_PPC">R_PPC</a> = 11 <span class="comment">// R_POWERPC_REL14</span>
    <span id="R_PPC_REL14_BRTAKEN">R_PPC_REL14_BRTAKEN</span>   <a href="#R_PPC">R_PPC</a> = 12 <span class="comment">// R_POWERPC_REL14_BRTAKEN</span>
    <span id="R_PPC_REL14_BRNTAKEN">R_PPC_REL14_BRNTAKEN</span>  <a href="#R_PPC">R_PPC</a> = 13 <span class="comment">// R_POWERPC_REL14_BRNTAKEN</span>
    <span id="R_PPC_GOT16">R_PPC_GOT16</span>           <a href="#R_PPC">R_PPC</a> = 14 <span class="comment">// R_POWERPC_GOT16</span>
    <span id="R_PPC_GOT16_LO">R_PPC_GOT16_LO</span>        <a href="#R_PPC">R_PPC</a> = 15 <span class="comment">// R_POWERPC_GOT16_LO</span>
    <span id="R_PPC_GOT16_HI">R_PPC_GOT16_HI</span>        <a href="#R_PPC">R_PPC</a> = 16 <span class="comment">// R_POWERPC_GOT16_HI</span>
    <span id="R_PPC_GOT16_HA">R_PPC_GOT16_HA</span>        <a href="#R_PPC">R_PPC</a> = 17 <span class="comment">// R_POWERPC_GOT16_HA</span>
    <span id="R_PPC_PLTREL24">R_PPC_PLTREL24</span>        <a href="#R_PPC">R_PPC</a> = 18
    <span id="R_PPC_COPY">R_PPC_COPY</span>            <a href="#R_PPC">R_PPC</a> = 19 <span class="comment">// R_POWERPC_COPY</span>
    <span id="R_PPC_GLOB_DAT">R_PPC_GLOB_DAT</span>        <a href="#R_PPC">R_PPC</a> = 20 <span class="comment">// R_POWERPC_GLOB_DAT</span>
    <span id="R_PPC_JMP_SLOT">R_PPC_JMP_SLOT</span>        <a href="#R_PPC">R_PPC</a> = 21 <span class="comment">// R_POWERPC_JMP_SLOT</span>
    <span id="R_PPC_RELATIVE">R_PPC_RELATIVE</span>        <a href="#R_PPC">R_PPC</a> = 22 <span class="comment">// R_POWERPC_RELATIVE</span>
    <span id="R_PPC_LOCAL24PC">R_PPC_LOCAL24PC</span>       <a href="#R_PPC">R_PPC</a> = 23
    <span id="R_PPC_UADDR32">R_PPC_UADDR32</span>         <a href="#R_PPC">R_PPC</a> = 24 <span class="comment">// R_POWERPC_UADDR32</span>
    <span id="R_PPC_UADDR16">R_PPC_UADDR16</span>         <a href="#R_PPC">R_PPC</a> = 25 <span class="comment">// R_POWERPC_UADDR16</span>
    <span id="R_PPC_REL32">R_PPC_REL32</span>           <a href="#R_PPC">R_PPC</a> = 26 <span class="comment">// R_POWERPC_REL32</span>
    <span id="R_PPC_PLT32">R_PPC_PLT32</span>           <a href="#R_PPC">R_PPC</a> = 27 <span class="comment">// R_POWERPC_PLT32</span>
    <span id="R_PPC_PLTREL32">R_PPC_PLTREL32</span>        <a href="#R_PPC">R_PPC</a> = 28 <span class="comment">// R_POWERPC_PLTREL32</span>
    <span id="R_PPC_PLT16_LO">R_PPC_PLT16_LO</span>        <a href="#R_PPC">R_PPC</a> = 29 <span class="comment">// R_POWERPC_PLT16_LO</span>
    <span id="R_PPC_PLT16_HI">R_PPC_PLT16_HI</span>        <a href="#R_PPC">R_PPC</a> = 30 <span class="comment">// R_POWERPC_PLT16_HI</span>
    <span id="R_PPC_PLT16_HA">R_PPC_PLT16_HA</span>        <a href="#R_PPC">R_PPC</a> = 31 <span class="comment">// R_POWERPC_PLT16_HA</span>
    <span id="R_PPC_SDAREL16">R_PPC_SDAREL16</span>        <a href="#R_PPC">R_PPC</a> = 32
    <span id="R_PPC_SECTOFF">R_PPC_SECTOFF</span>         <a href="#R_PPC">R_PPC</a> = 33 <span class="comment">// R_POWERPC_SECTOFF</span>
    <span id="R_PPC_SECTOFF_LO">R_PPC_SECTOFF_LO</span>      <a href="#R_PPC">R_PPC</a> = 34 <span class="comment">// R_POWERPC_SECTOFF_LO</span>
    <span id="R_PPC_SECTOFF_HI">R_PPC_SECTOFF_HI</span>      <a href="#R_PPC">R_PPC</a> = 35 <span class="comment">// R_POWERPC_SECTOFF_HI</span>
    <span id="R_PPC_SECTOFF_HA">R_PPC_SECTOFF_HA</span>      <a href="#R_PPC">R_PPC</a> = 36 <span class="comment">// R_POWERPC_SECTOFF_HA</span>
    <span id="R_PPC_TLS">R_PPC_TLS</span>             <a href="#R_PPC">R_PPC</a> = 67 <span class="comment">// R_POWERPC_TLS</span>
    <span id="R_PPC_DTPMOD32">R_PPC_DTPMOD32</span>        <a href="#R_PPC">R_PPC</a> = 68 <span class="comment">// R_POWERPC_DTPMOD32</span>
    <span id="R_PPC_TPREL16">R_PPC_TPREL16</span>         <a href="#R_PPC">R_PPC</a> = 69 <span class="comment">// R_POWERPC_TPREL16</span>
    <span id="R_PPC_TPREL16_LO">R_PPC_TPREL16_LO</span>      <a href="#R_PPC">R_PPC</a> = 70 <span class="comment">// R_POWERPC_TPREL16_LO</span>
    <span id="R_PPC_TPREL16_HI">R_PPC_TPREL16_HI</span>      <a href="#R_PPC">R_PPC</a> = 71 <span class="comment">// R_POWERPC_TPREL16_HI</span>
    <span id="R_PPC_TPREL16_HA">R_PPC_TPREL16_HA</span>      <a href="#R_PPC">R_PPC</a> = 72 <span class="comment">// R_POWERPC_TPREL16_HA</span>
    <span id="R_PPC_TPREL32">R_PPC_TPREL32</span>         <a href="#R_PPC">R_PPC</a> = 73 <span class="comment">// R_POWERPC_TPREL32</span>
    <span id="R_PPC_DTPREL16">R_PPC_DTPREL16</span>        <a href="#R_PPC">R_PPC</a> = 74 <span class="comment">// R_POWERPC_DTPREL16</span>
    <span id="R_PPC_DTPREL16_LO">R_PPC_DTPREL16_LO</span>     <a href="#R_PPC">R_PPC</a> = 75 <span class="comment">// R_POWERPC_DTPREL16_LO</span>
    <span id="R_PPC_DTPREL16_HI">R_PPC_DTPREL16_HI</span>     <a href="#R_PPC">R_PPC</a> = 76 <span class="comment">// R_POWERPC_DTPREL16_HI</span>
    <span id="R_PPC_DTPREL16_HA">R_PPC_DTPREL16_HA</span>     <a href="#R_PPC">R_PPC</a> = 77 <span class="comment">// R_POWERPC_DTPREL16_HA</span>
    <span id="R_PPC_DTPREL32">R_PPC_DTPREL32</span>        <a href="#R_PPC">R_PPC</a> = 78 <span class="comment">// R_POWERPC_DTPREL32</span>
    <span id="R_PPC_GOT_TLSGD16">R_PPC_GOT_TLSGD16</span>     <a href="#R_PPC">R_PPC</a> = 79 <span class="comment">// R_POWERPC_GOT_TLSGD16</span>
    <span id="R_PPC_GOT_TLSGD16_LO">R_PPC_GOT_TLSGD16_LO</span>  <a href="#R_PPC">R_PPC</a> = 80 <span class="comment">// R_POWERPC_GOT_TLSGD16_LO</span>
    <span id="R_PPC_GOT_TLSGD16_HI">R_PPC_GOT_TLSGD16_HI</span>  <a href="#R_PPC">R_PPC</a> = 81 <span class="comment">// R_POWERPC_GOT_TLSGD16_HI</span>
    <span id="R_PPC_GOT_TLSGD16_HA">R_PPC_GOT_TLSGD16_HA</span>  <a href="#R_PPC">R_PPC</a> = 82 <span class="comment">// R_POWERPC_GOT_TLSGD16_HA</span>
    <span id="R_PPC_GOT_TLSLD16">R_PPC_GOT_TLSLD16</span>     <a href="#R_PPC">R_PPC</a> = 83 <span class="comment">// R_POWERPC_GOT_TLSLD16</span>
    <span id="R_PPC_GOT_TLSLD16_LO">R_PPC_GOT_TLSLD16_LO</span>  <a href="#R_PPC">R_PPC</a> = 84 <span class="comment">// R_POWERPC_GOT_TLSLD16_LO</span>
    <span id="R_PPC_GOT_TLSLD16_HI">R_PPC_GOT_TLSLD16_HI</span>  <a href="#R_PPC">R_PPC</a> = 85 <span class="comment">// R_POWERPC_GOT_TLSLD16_HI</span>
    <span id="R_PPC_GOT_TLSLD16_HA">R_PPC_GOT_TLSLD16_HA</span>  <a href="#R_PPC">R_PPC</a> = 86 <span class="comment">// R_POWERPC_GOT_TLSLD16_HA</span>
    <span id="R_PPC_GOT_TPREL16">R_PPC_GOT_TPREL16</span>     <a href="#R_PPC">R_PPC</a> = 87 <span class="comment">// R_POWERPC_GOT_TPREL16</span>
    <span id="R_PPC_GOT_TPREL16_LO">R_PPC_GOT_TPREL16_LO</span>  <a href="#R_PPC">R_PPC</a> = 88 <span class="comment">// R_POWERPC_GOT_TPREL16_LO</span>
    <span id="R_PPC_GOT_TPREL16_HI">R_PPC_GOT_TPREL16_HI</span>  <a href="#R_PPC">R_PPC</a> = 89 <span class="comment">// R_POWERPC_GOT_TPREL16_HI</span>
    <span id="R_PPC_GOT_TPREL16_HA">R_PPC_GOT_TPREL16_HA</span>  <a href="#R_PPC">R_PPC</a> = 90 <span class="comment">// R_POWERPC_GOT_TPREL16_HA</span>
    <span id="R_PPC_EMB_NADDR32">R_PPC_EMB_NADDR32</span>     <a href="#R_PPC">R_PPC</a> = 101
    <span id="R_PPC_EMB_NADDR16">R_PPC_EMB_NADDR16</span>     <a href="#R_PPC">R_PPC</a> = 102
    <span id="R_PPC_EMB_NADDR16_LO">R_PPC_EMB_NADDR16_LO</span>  <a href="#R_PPC">R_PPC</a> = 103
    <span id="R_PPC_EMB_NADDR16_HI">R_PPC_EMB_NADDR16_HI</span>  <a href="#R_PPC">R_PPC</a> = 104
    <span id="R_PPC_EMB_NADDR16_HA">R_PPC_EMB_NADDR16_HA</span>  <a href="#R_PPC">R_PPC</a> = 105
    <span id="R_PPC_EMB_SDAI16">R_PPC_EMB_SDAI16</span>      <a href="#R_PPC">R_PPC</a> = 106
    <span id="R_PPC_EMB_SDA2I16">R_PPC_EMB_SDA2I16</span>     <a href="#R_PPC">R_PPC</a> = 107
    <span id="R_PPC_EMB_SDA2REL">R_PPC_EMB_SDA2REL</span>     <a href="#R_PPC">R_PPC</a> = 108
    <span id="R_PPC_EMB_SDA21">R_PPC_EMB_SDA21</span>       <a href="#R_PPC">R_PPC</a> = 109
    <span id="R_PPC_EMB_MRKREF">R_PPC_EMB_MRKREF</span>      <a href="#R_PPC">R_PPC</a> = 110
    <span id="R_PPC_EMB_RELSEC16">R_PPC_EMB_RELSEC16</span>    <a href="#R_PPC">R_PPC</a> = 111
    <span id="R_PPC_EMB_RELST_LO">R_PPC_EMB_RELST_LO</span>    <a href="#R_PPC">R_PPC</a> = 112
    <span id="R_PPC_EMB_RELST_HI">R_PPC_EMB_RELST_HI</span>    <a href="#R_PPC">R_PPC</a> = 113
    <span id="R_PPC_EMB_RELST_HA">R_PPC_EMB_RELST_HA</span>    <a href="#R_PPC">R_PPC</a> = 114
    <span id="R_PPC_EMB_BIT_FLD">R_PPC_EMB_BIT_FLD</span>     <a href="#R_PPC">R_PPC</a> = 115
    <span id="R_PPC_EMB_RELSDA">R_PPC_EMB_RELSDA</span>      <a href="#R_PPC">R_PPC</a> = 116
)</pre>









### <a id="R_PPC.GoString">func</a> (R\_PPC) [GoString](https://golang.org/src/debug/elf/elf.go?s=84479:84511#L2131)
<pre>func (i <a href="#R_PPC">R_PPC</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="R_PPC.String">func</a> (R\_PPC) [String](https://golang.org/src/debug/elf/elf.go?s=84393:84423#L2130)
<pre>func (i <a href="#R_PPC">R_PPC</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="R_PPC64">type</a> [R_PPC64](https://golang.org/src/debug/elf/elf.go?s=84900:84916#L2139)
Relocation types for 64-bit PowerPC or Power Architecture processors.

Values that are shared by both R_PPC and R_PPC64 are prefixed with
R_POWERPC_ in the ELF standard. For the R_PPC64 type, the relevant
shared relocations have been renamed with the prefix R_PPC64_.
The original name follows the value in a comment.


<pre>type R_PPC64 <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="R_PPC64_NONE">R_PPC64_NONE</span>               <a href="#R_PPC64">R_PPC64</a> = 0  <span class="comment">// R_POWERPC_NONE</span>
    <span id="R_PPC64_ADDR32">R_PPC64_ADDR32</span>             <a href="#R_PPC64">R_PPC64</a> = 1  <span class="comment">// R_POWERPC_ADDR32</span>
    <span id="R_PPC64_ADDR24">R_PPC64_ADDR24</span>             <a href="#R_PPC64">R_PPC64</a> = 2  <span class="comment">// R_POWERPC_ADDR24</span>
    <span id="R_PPC64_ADDR16">R_PPC64_ADDR16</span>             <a href="#R_PPC64">R_PPC64</a> = 3  <span class="comment">// R_POWERPC_ADDR16</span>
    <span id="R_PPC64_ADDR16_LO">R_PPC64_ADDR16_LO</span>          <a href="#R_PPC64">R_PPC64</a> = 4  <span class="comment">// R_POWERPC_ADDR16_LO</span>
    <span id="R_PPC64_ADDR16_HI">R_PPC64_ADDR16_HI</span>          <a href="#R_PPC64">R_PPC64</a> = 5  <span class="comment">// R_POWERPC_ADDR16_HI</span>
    <span id="R_PPC64_ADDR16_HA">R_PPC64_ADDR16_HA</span>          <a href="#R_PPC64">R_PPC64</a> = 6  <span class="comment">// R_POWERPC_ADDR16_HA</span>
    <span id="R_PPC64_ADDR14">R_PPC64_ADDR14</span>             <a href="#R_PPC64">R_PPC64</a> = 7  <span class="comment">// R_POWERPC_ADDR14</span>
    <span id="R_PPC64_ADDR14_BRTAKEN">R_PPC64_ADDR14_BRTAKEN</span>     <a href="#R_PPC64">R_PPC64</a> = 8  <span class="comment">// R_POWERPC_ADDR14_BRTAKEN</span>
    <span id="R_PPC64_ADDR14_BRNTAKEN">R_PPC64_ADDR14_BRNTAKEN</span>    <a href="#R_PPC64">R_PPC64</a> = 9  <span class="comment">// R_POWERPC_ADDR14_BRNTAKEN</span>
    <span id="R_PPC64_REL24">R_PPC64_REL24</span>              <a href="#R_PPC64">R_PPC64</a> = 10 <span class="comment">// R_POWERPC_REL24</span>
    <span id="R_PPC64_REL14">R_PPC64_REL14</span>              <a href="#R_PPC64">R_PPC64</a> = 11 <span class="comment">// R_POWERPC_REL14</span>
    <span id="R_PPC64_REL14_BRTAKEN">R_PPC64_REL14_BRTAKEN</span>      <a href="#R_PPC64">R_PPC64</a> = 12 <span class="comment">// R_POWERPC_REL14_BRTAKEN</span>
    <span id="R_PPC64_REL14_BRNTAKEN">R_PPC64_REL14_BRNTAKEN</span>     <a href="#R_PPC64">R_PPC64</a> = 13 <span class="comment">// R_POWERPC_REL14_BRNTAKEN</span>
    <span id="R_PPC64_GOT16">R_PPC64_GOT16</span>              <a href="#R_PPC64">R_PPC64</a> = 14 <span class="comment">// R_POWERPC_GOT16</span>
    <span id="R_PPC64_GOT16_LO">R_PPC64_GOT16_LO</span>           <a href="#R_PPC64">R_PPC64</a> = 15 <span class="comment">// R_POWERPC_GOT16_LO</span>
    <span id="R_PPC64_GOT16_HI">R_PPC64_GOT16_HI</span>           <a href="#R_PPC64">R_PPC64</a> = 16 <span class="comment">// R_POWERPC_GOT16_HI</span>
    <span id="R_PPC64_GOT16_HA">R_PPC64_GOT16_HA</span>           <a href="#R_PPC64">R_PPC64</a> = 17 <span class="comment">// R_POWERPC_GOT16_HA</span>
    <span id="R_PPC64_JMP_SLOT">R_PPC64_JMP_SLOT</span>           <a href="#R_PPC64">R_PPC64</a> = 21 <span class="comment">// R_POWERPC_JMP_SLOT</span>
    <span id="R_PPC64_REL32">R_PPC64_REL32</span>              <a href="#R_PPC64">R_PPC64</a> = 26 <span class="comment">// R_POWERPC_REL32</span>
    <span id="R_PPC64_ADDR64">R_PPC64_ADDR64</span>             <a href="#R_PPC64">R_PPC64</a> = 38
    <span id="R_PPC64_ADDR16_HIGHER">R_PPC64_ADDR16_HIGHER</span>      <a href="#R_PPC64">R_PPC64</a> = 39
    <span id="R_PPC64_ADDR16_HIGHERA">R_PPC64_ADDR16_HIGHERA</span>     <a href="#R_PPC64">R_PPC64</a> = 40
    <span id="R_PPC64_ADDR16_HIGHEST">R_PPC64_ADDR16_HIGHEST</span>     <a href="#R_PPC64">R_PPC64</a> = 41
    <span id="R_PPC64_ADDR16_HIGHESTA">R_PPC64_ADDR16_HIGHESTA</span>    <a href="#R_PPC64">R_PPC64</a> = 42
    <span id="R_PPC64_REL64">R_PPC64_REL64</span>              <a href="#R_PPC64">R_PPC64</a> = 44
    <span id="R_PPC64_TOC16">R_PPC64_TOC16</span>              <a href="#R_PPC64">R_PPC64</a> = 47
    <span id="R_PPC64_TOC16_LO">R_PPC64_TOC16_LO</span>           <a href="#R_PPC64">R_PPC64</a> = 48
    <span id="R_PPC64_TOC16_HI">R_PPC64_TOC16_HI</span>           <a href="#R_PPC64">R_PPC64</a> = 49
    <span id="R_PPC64_TOC16_HA">R_PPC64_TOC16_HA</span>           <a href="#R_PPC64">R_PPC64</a> = 50
    <span id="R_PPC64_TOC">R_PPC64_TOC</span>                <a href="#R_PPC64">R_PPC64</a> = 51
    <span id="R_PPC64_PLTGOT16">R_PPC64_PLTGOT16</span>           <a href="#R_PPC64">R_PPC64</a> = 52
    <span id="R_PPC64_PLTGOT16_LO">R_PPC64_PLTGOT16_LO</span>        <a href="#R_PPC64">R_PPC64</a> = 53
    <span id="R_PPC64_PLTGOT16_HI">R_PPC64_PLTGOT16_HI</span>        <a href="#R_PPC64">R_PPC64</a> = 54
    <span id="R_PPC64_PLTGOT16_HA">R_PPC64_PLTGOT16_HA</span>        <a href="#R_PPC64">R_PPC64</a> = 55
    <span id="R_PPC64_ADDR16_DS">R_PPC64_ADDR16_DS</span>          <a href="#R_PPC64">R_PPC64</a> = 56
    <span id="R_PPC64_ADDR16_LO_DS">R_PPC64_ADDR16_LO_DS</span>       <a href="#R_PPC64">R_PPC64</a> = 57
    <span id="R_PPC64_GOT16_DS">R_PPC64_GOT16_DS</span>           <a href="#R_PPC64">R_PPC64</a> = 58
    <span id="R_PPC64_GOT16_LO_DS">R_PPC64_GOT16_LO_DS</span>        <a href="#R_PPC64">R_PPC64</a> = 59
    <span id="R_PPC64_PLT16_LO_DS">R_PPC64_PLT16_LO_DS</span>        <a href="#R_PPC64">R_PPC64</a> = 60
    <span id="R_PPC64_SECTOFF_DS">R_PPC64_SECTOFF_DS</span>         <a href="#R_PPC64">R_PPC64</a> = 61
    <span id="R_PPC64_SECTOFF_LO_DS">R_PPC64_SECTOFF_LO_DS</span>      <a href="#R_PPC64">R_PPC64</a> = 61
    <span id="R_PPC64_TOC16_DS">R_PPC64_TOC16_DS</span>           <a href="#R_PPC64">R_PPC64</a> = 63
    <span id="R_PPC64_TOC16_LO_DS">R_PPC64_TOC16_LO_DS</span>        <a href="#R_PPC64">R_PPC64</a> = 64
    <span id="R_PPC64_PLTGOT16_DS">R_PPC64_PLTGOT16_DS</span>        <a href="#R_PPC64">R_PPC64</a> = 65
    <span id="R_PPC64_PLTGOT_LO_DS">R_PPC64_PLTGOT_LO_DS</span>       <a href="#R_PPC64">R_PPC64</a> = 66
    <span id="R_PPC64_TLS">R_PPC64_TLS</span>                <a href="#R_PPC64">R_PPC64</a> = 67 <span class="comment">// R_POWERPC_TLS</span>
    <span id="R_PPC64_DTPMOD64">R_PPC64_DTPMOD64</span>           <a href="#R_PPC64">R_PPC64</a> = 68 <span class="comment">// R_POWERPC_DTPMOD64</span>
    <span id="R_PPC64_TPREL16">R_PPC64_TPREL16</span>            <a href="#R_PPC64">R_PPC64</a> = 69 <span class="comment">// R_POWERPC_TPREL16</span>
    <span id="R_PPC64_TPREL16_LO">R_PPC64_TPREL16_LO</span>         <a href="#R_PPC64">R_PPC64</a> = 70 <span class="comment">// R_POWERPC_TPREL16_LO</span>
    <span id="R_PPC64_TPREL16_HI">R_PPC64_TPREL16_HI</span>         <a href="#R_PPC64">R_PPC64</a> = 71 <span class="comment">// R_POWERPC_TPREL16_HI</span>
    <span id="R_PPC64_TPREL16_HA">R_PPC64_TPREL16_HA</span>         <a href="#R_PPC64">R_PPC64</a> = 72 <span class="comment">// R_POWERPC_TPREL16_HA</span>
    <span id="R_PPC64_TPREL64">R_PPC64_TPREL64</span>            <a href="#R_PPC64">R_PPC64</a> = 73 <span class="comment">// R_POWERPC_TPREL64</span>
    <span id="R_PPC64_DTPREL16">R_PPC64_DTPREL16</span>           <a href="#R_PPC64">R_PPC64</a> = 74 <span class="comment">// R_POWERPC_DTPREL16</span>
    <span id="R_PPC64_DTPREL16_LO">R_PPC64_DTPREL16_LO</span>        <a href="#R_PPC64">R_PPC64</a> = 75 <span class="comment">// R_POWERPC_DTPREL16_LO</span>
    <span id="R_PPC64_DTPREL16_HI">R_PPC64_DTPREL16_HI</span>        <a href="#R_PPC64">R_PPC64</a> = 76 <span class="comment">// R_POWERPC_DTPREL16_HI</span>
    <span id="R_PPC64_DTPREL16_HA">R_PPC64_DTPREL16_HA</span>        <a href="#R_PPC64">R_PPC64</a> = 77 <span class="comment">// R_POWERPC_DTPREL16_HA</span>
    <span id="R_PPC64_DTPREL64">R_PPC64_DTPREL64</span>           <a href="#R_PPC64">R_PPC64</a> = 78 <span class="comment">// R_POWERPC_DTPREL64</span>
    <span id="R_PPC64_GOT_TLSGD16">R_PPC64_GOT_TLSGD16</span>        <a href="#R_PPC64">R_PPC64</a> = 79 <span class="comment">// R_POWERPC_GOT_TLSGD16</span>
    <span id="R_PPC64_GOT_TLSGD16_LO">R_PPC64_GOT_TLSGD16_LO</span>     <a href="#R_PPC64">R_PPC64</a> = 80 <span class="comment">// R_POWERPC_GOT_TLSGD16_LO</span>
    <span id="R_PPC64_GOT_TLSGD16_HI">R_PPC64_GOT_TLSGD16_HI</span>     <a href="#R_PPC64">R_PPC64</a> = 81 <span class="comment">// R_POWERPC_GOT_TLSGD16_HI</span>
    <span id="R_PPC64_GOT_TLSGD16_HA">R_PPC64_GOT_TLSGD16_HA</span>     <a href="#R_PPC64">R_PPC64</a> = 82 <span class="comment">// R_POWERPC_GOT_TLSGD16_HA</span>
    <span id="R_PPC64_GOT_TLSLD16">R_PPC64_GOT_TLSLD16</span>        <a href="#R_PPC64">R_PPC64</a> = 83 <span class="comment">// R_POWERPC_GOT_TLSLD16</span>
    <span id="R_PPC64_GOT_TLSLD16_LO">R_PPC64_GOT_TLSLD16_LO</span>     <a href="#R_PPC64">R_PPC64</a> = 84 <span class="comment">// R_POWERPC_GOT_TLSLD16_LO</span>
    <span id="R_PPC64_GOT_TLSLD16_HI">R_PPC64_GOT_TLSLD16_HI</span>     <a href="#R_PPC64">R_PPC64</a> = 85 <span class="comment">// R_POWERPC_GOT_TLSLD16_HI</span>
    <span id="R_PPC64_GOT_TLSLD16_HA">R_PPC64_GOT_TLSLD16_HA</span>     <a href="#R_PPC64">R_PPC64</a> = 86 <span class="comment">// R_POWERPC_GOT_TLSLD16_HA</span>
    <span id="R_PPC64_GOT_TPREL16_DS">R_PPC64_GOT_TPREL16_DS</span>     <a href="#R_PPC64">R_PPC64</a> = 87 <span class="comment">// R_POWERPC_GOT_TPREL16_DS</span>
    <span id="R_PPC64_GOT_TPREL16_LO_DS">R_PPC64_GOT_TPREL16_LO_DS</span>  <a href="#R_PPC64">R_PPC64</a> = 88 <span class="comment">// R_POWERPC_GOT_TPREL16_LO_DS</span>
    <span id="R_PPC64_GOT_TPREL16_HI">R_PPC64_GOT_TPREL16_HI</span>     <a href="#R_PPC64">R_PPC64</a> = 89 <span class="comment">// R_POWERPC_GOT_TPREL16_HI</span>
    <span id="R_PPC64_GOT_TPREL16_HA">R_PPC64_GOT_TPREL16_HA</span>     <a href="#R_PPC64">R_PPC64</a> = 90 <span class="comment">// R_POWERPC_GOT_TPREL16_HA</span>
    <span id="R_PPC64_GOT_DTPREL16_DS">R_PPC64_GOT_DTPREL16_DS</span>    <a href="#R_PPC64">R_PPC64</a> = 91 <span class="comment">// R_POWERPC_GOT_DTPREL16_DS</span>
    <span id="R_PPC64_GOT_DTPREL16_LO_DS">R_PPC64_GOT_DTPREL16_LO_DS</span> <a href="#R_PPC64">R_PPC64</a> = 92 <span class="comment">// R_POWERPC_GOT_DTPREL16_LO_DS</span>
    <span id="R_PPC64_GOT_DTPREL16_HI">R_PPC64_GOT_DTPREL16_HI</span>    <a href="#R_PPC64">R_PPC64</a> = 93 <span class="comment">// R_POWERPC_GOT_DTPREL16_HI</span>
    <span id="R_PPC64_GOT_DTPREL16_HA">R_PPC64_GOT_DTPREL16_HA</span>    <a href="#R_PPC64">R_PPC64</a> = 94 <span class="comment">// R_POWERPC_GOT_DTPREL16_HA</span>
    <span id="R_PPC64_TPREL16_DS">R_PPC64_TPREL16_DS</span>         <a href="#R_PPC64">R_PPC64</a> = 95
    <span id="R_PPC64_TPREL16_LO_DS">R_PPC64_TPREL16_LO_DS</span>      <a href="#R_PPC64">R_PPC64</a> = 96
    <span id="R_PPC64_TPREL16_HIGHER">R_PPC64_TPREL16_HIGHER</span>     <a href="#R_PPC64">R_PPC64</a> = 97
    <span id="R_PPC64_TPREL16_HIGHERA">R_PPC64_TPREL16_HIGHERA</span>    <a href="#R_PPC64">R_PPC64</a> = 98
    <span id="R_PPC64_TPREL16_HIGHEST">R_PPC64_TPREL16_HIGHEST</span>    <a href="#R_PPC64">R_PPC64</a> = 99
    <span id="R_PPC64_TPREL16_HIGHESTA">R_PPC64_TPREL16_HIGHESTA</span>   <a href="#R_PPC64">R_PPC64</a> = 100
    <span id="R_PPC64_DTPREL16_DS">R_PPC64_DTPREL16_DS</span>        <a href="#R_PPC64">R_PPC64</a> = 101
    <span id="R_PPC64_DTPREL16_LO_DS">R_PPC64_DTPREL16_LO_DS</span>     <a href="#R_PPC64">R_PPC64</a> = 102
    <span id="R_PPC64_DTPREL16_HIGHER">R_PPC64_DTPREL16_HIGHER</span>    <a href="#R_PPC64">R_PPC64</a> = 103
    <span id="R_PPC64_DTPREL16_HIGHERA">R_PPC64_DTPREL16_HIGHERA</span>   <a href="#R_PPC64">R_PPC64</a> = 104
    <span id="R_PPC64_DTPREL16_HIGHEST">R_PPC64_DTPREL16_HIGHEST</span>   <a href="#R_PPC64">R_PPC64</a> = 105
    <span id="R_PPC64_DTPREL16_HIGHESTA">R_PPC64_DTPREL16_HIGHESTA</span>  <a href="#R_PPC64">R_PPC64</a> = 106
    <span id="R_PPC64_TLSGD">R_PPC64_TLSGD</span>              <a href="#R_PPC64">R_PPC64</a> = 107
    <span id="R_PPC64_TLSLD">R_PPC64_TLSLD</span>              <a href="#R_PPC64">R_PPC64</a> = 108
    <span id="R_PPC64_TOCSAVE">R_PPC64_TOCSAVE</span>            <a href="#R_PPC64">R_PPC64</a> = 109
    <span id="R_PPC64_ADDR16_HIGH">R_PPC64_ADDR16_HIGH</span>        <a href="#R_PPC64">R_PPC64</a> = 110
    <span id="R_PPC64_ADDR16_HIGHA">R_PPC64_ADDR16_HIGHA</span>       <a href="#R_PPC64">R_PPC64</a> = 111
    <span id="R_PPC64_TPREL16_HIGH">R_PPC64_TPREL16_HIGH</span>       <a href="#R_PPC64">R_PPC64</a> = 112
    <span id="R_PPC64_TPREL16_HIGHA">R_PPC64_TPREL16_HIGHA</span>      <a href="#R_PPC64">R_PPC64</a> = 113
    <span id="R_PPC64_DTPREL16_HIGH">R_PPC64_DTPREL16_HIGH</span>      <a href="#R_PPC64">R_PPC64</a> = 114
    <span id="R_PPC64_DTPREL16_HIGHA">R_PPC64_DTPREL16_HIGHA</span>     <a href="#R_PPC64">R_PPC64</a> = 115
    <span id="R_PPC64_REL24_NOTOC">R_PPC64_REL24_NOTOC</span>        <a href="#R_PPC64">R_PPC64</a> = 116
    <span id="R_PPC64_ADDR64_LOCAL">R_PPC64_ADDR64_LOCAL</span>       <a href="#R_PPC64">R_PPC64</a> = 117
    <span id="R_PPC64_ENTRY">R_PPC64_ENTRY</span>              <a href="#R_PPC64">R_PPC64</a> = 118
    <span id="R_PPC64_REL16DX_HA">R_PPC64_REL16DX_HA</span>         <a href="#R_PPC64">R_PPC64</a> = 246 <span class="comment">// R_POWERPC_REL16DX_HA</span>
    <span id="R_PPC64_JMP_IREL">R_PPC64_JMP_IREL</span>           <a href="#R_PPC64">R_PPC64</a> = 247
    <span id="R_PPC64_IRELATIVE">R_PPC64_IRELATIVE</span>          <a href="#R_PPC64">R_PPC64</a> = 248 <span class="comment">// R_POWERPC_IRELATIVE</span>
    <span id="R_PPC64_REL16">R_PPC64_REL16</span>              <a href="#R_PPC64">R_PPC64</a> = 249 <span class="comment">// R_POWERPC_REL16</span>
    <span id="R_PPC64_REL16_LO">R_PPC64_REL16_LO</span>           <a href="#R_PPC64">R_PPC64</a> = 250 <span class="comment">// R_POWERPC_REL16_LO</span>
    <span id="R_PPC64_REL16_HI">R_PPC64_REL16_HI</span>           <a href="#R_PPC64">R_PPC64</a> = 251 <span class="comment">// R_POWERPC_REL16_HI</span>
    <span id="R_PPC64_REL16_HA">R_PPC64_REL16_HA</span>           <a href="#R_PPC64">R_PPC64</a> = 252 <span class="comment">// R_POWERPC_REL16_HA</span>
)</pre>









### <a id="R_PPC64.GoString">func</a> (R\_PPC64) [GoString](https://golang.org/src/debug/elf/elf.go?s=93793:93827#L2358)
<pre>func (i <a href="#R_PPC64">R_PPC64</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="R_PPC64.String">func</a> (R\_PPC64) [String](https://golang.org/src/debug/elf/elf.go?s=93703:93735#L2357)
<pre>func (i <a href="#R_PPC64">R_PPC64</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="R_RISCV">type</a> [R_RISCV](https://golang.org/src/debug/elf/elf.go?s=93926:93942#L2361)
Relocation types for RISC-V processors.


<pre>type R_RISCV <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="R_RISCV_NONE">R_RISCV_NONE</span>          <a href="#R_RISCV">R_RISCV</a> = 0  <span class="comment">/* No relocation. */</span>
    <span id="R_RISCV_32">R_RISCV_32</span>            <a href="#R_RISCV">R_RISCV</a> = 1  <span class="comment">/* Add 32 bit zero extended symbol value */</span>
    <span id="R_RISCV_64">R_RISCV_64</span>            <a href="#R_RISCV">R_RISCV</a> = 2  <span class="comment">/* Add 64 bit symbol value. */</span>
    <span id="R_RISCV_RELATIVE">R_RISCV_RELATIVE</span>      <a href="#R_RISCV">R_RISCV</a> = 3  <span class="comment">/* Add load address of shared object. */</span>
    <span id="R_RISCV_COPY">R_RISCV_COPY</span>          <a href="#R_RISCV">R_RISCV</a> = 4  <span class="comment">/* Copy data from shared object. */</span>
    <span id="R_RISCV_JUMP_SLOT">R_RISCV_JUMP_SLOT</span>     <a href="#R_RISCV">R_RISCV</a> = 5  <span class="comment">/* Set GOT entry to code address. */</span>
    <span id="R_RISCV_TLS_DTPMOD32">R_RISCV_TLS_DTPMOD32</span>  <a href="#R_RISCV">R_RISCV</a> = 6  <span class="comment">/* 32 bit ID of module containing symbol */</span>
    <span id="R_RISCV_TLS_DTPMOD64">R_RISCV_TLS_DTPMOD64</span>  <a href="#R_RISCV">R_RISCV</a> = 7  <span class="comment">/* ID of module containing symbol */</span>
    <span id="R_RISCV_TLS_DTPREL32">R_RISCV_TLS_DTPREL32</span>  <a href="#R_RISCV">R_RISCV</a> = 8  <span class="comment">/* 32 bit relative offset in TLS block */</span>
    <span id="R_RISCV_TLS_DTPREL64">R_RISCV_TLS_DTPREL64</span>  <a href="#R_RISCV">R_RISCV</a> = 9  <span class="comment">/* Relative offset in TLS block */</span>
    <span id="R_RISCV_TLS_TPREL32">R_RISCV_TLS_TPREL32</span>   <a href="#R_RISCV">R_RISCV</a> = 10 <span class="comment">/* 32 bit relative offset in static TLS block */</span>
    <span id="R_RISCV_TLS_TPREL64">R_RISCV_TLS_TPREL64</span>   <a href="#R_RISCV">R_RISCV</a> = 11 <span class="comment">/* Relative offset in static TLS block */</span>
    <span id="R_RISCV_BRANCH">R_RISCV_BRANCH</span>        <a href="#R_RISCV">R_RISCV</a> = 16 <span class="comment">/* PC-relative branch */</span>
    <span id="R_RISCV_JAL">R_RISCV_JAL</span>           <a href="#R_RISCV">R_RISCV</a> = 17 <span class="comment">/* PC-relative jump */</span>
    <span id="R_RISCV_CALL">R_RISCV_CALL</span>          <a href="#R_RISCV">R_RISCV</a> = 18 <span class="comment">/* PC-relative call */</span>
    <span id="R_RISCV_CALL_PLT">R_RISCV_CALL_PLT</span>      <a href="#R_RISCV">R_RISCV</a> = 19 <span class="comment">/* PC-relative call (PLT) */</span>
    <span id="R_RISCV_GOT_HI20">R_RISCV_GOT_HI20</span>      <a href="#R_RISCV">R_RISCV</a> = 20 <span class="comment">/* PC-relative GOT reference */</span>
    <span id="R_RISCV_TLS_GOT_HI20">R_RISCV_TLS_GOT_HI20</span>  <a href="#R_RISCV">R_RISCV</a> = 21 <span class="comment">/* PC-relative TLS IE GOT offset */</span>
    <span id="R_RISCV_TLS_GD_HI20">R_RISCV_TLS_GD_HI20</span>   <a href="#R_RISCV">R_RISCV</a> = 22 <span class="comment">/* PC-relative TLS GD reference */</span>
    <span id="R_RISCV_PCREL_HI20">R_RISCV_PCREL_HI20</span>    <a href="#R_RISCV">R_RISCV</a> = 23 <span class="comment">/* PC-relative reference */</span>
    <span id="R_RISCV_PCREL_LO12_I">R_RISCV_PCREL_LO12_I</span>  <a href="#R_RISCV">R_RISCV</a> = 24 <span class="comment">/* PC-relative reference */</span>
    <span id="R_RISCV_PCREL_LO12_S">R_RISCV_PCREL_LO12_S</span>  <a href="#R_RISCV">R_RISCV</a> = 25 <span class="comment">/* PC-relative reference */</span>
    <span id="R_RISCV_HI20">R_RISCV_HI20</span>          <a href="#R_RISCV">R_RISCV</a> = 26 <span class="comment">/* Absolute address */</span>
    <span id="R_RISCV_LO12_I">R_RISCV_LO12_I</span>        <a href="#R_RISCV">R_RISCV</a> = 27 <span class="comment">/* Absolute address */</span>
    <span id="R_RISCV_LO12_S">R_RISCV_LO12_S</span>        <a href="#R_RISCV">R_RISCV</a> = 28 <span class="comment">/* Absolute address */</span>
    <span id="R_RISCV_TPREL_HI20">R_RISCV_TPREL_HI20</span>    <a href="#R_RISCV">R_RISCV</a> = 29 <span class="comment">/* TLS LE thread offset */</span>
    <span id="R_RISCV_TPREL_LO12_I">R_RISCV_TPREL_LO12_I</span>  <a href="#R_RISCV">R_RISCV</a> = 30 <span class="comment">/* TLS LE thread offset */</span>
    <span id="R_RISCV_TPREL_LO12_S">R_RISCV_TPREL_LO12_S</span>  <a href="#R_RISCV">R_RISCV</a> = 31 <span class="comment">/* TLS LE thread offset */</span>
    <span id="R_RISCV_TPREL_ADD">R_RISCV_TPREL_ADD</span>     <a href="#R_RISCV">R_RISCV</a> = 32 <span class="comment">/* TLS LE thread usage */</span>
    <span id="R_RISCV_ADD8">R_RISCV_ADD8</span>          <a href="#R_RISCV">R_RISCV</a> = 33 <span class="comment">/* 8-bit label addition */</span>
    <span id="R_RISCV_ADD16">R_RISCV_ADD16</span>         <a href="#R_RISCV">R_RISCV</a> = 34 <span class="comment">/* 16-bit label addition */</span>
    <span id="R_RISCV_ADD32">R_RISCV_ADD32</span>         <a href="#R_RISCV">R_RISCV</a> = 35 <span class="comment">/* 32-bit label addition */</span>
    <span id="R_RISCV_ADD64">R_RISCV_ADD64</span>         <a href="#R_RISCV">R_RISCV</a> = 36 <span class="comment">/* 64-bit label addition */</span>
    <span id="R_RISCV_SUB8">R_RISCV_SUB8</span>          <a href="#R_RISCV">R_RISCV</a> = 37 <span class="comment">/* 8-bit label subtraction */</span>
    <span id="R_RISCV_SUB16">R_RISCV_SUB16</span>         <a href="#R_RISCV">R_RISCV</a> = 38 <span class="comment">/* 16-bit label subtraction */</span>
    <span id="R_RISCV_SUB32">R_RISCV_SUB32</span>         <a href="#R_RISCV">R_RISCV</a> = 39 <span class="comment">/* 32-bit label subtraction */</span>
    <span id="R_RISCV_SUB64">R_RISCV_SUB64</span>         <a href="#R_RISCV">R_RISCV</a> = 40 <span class="comment">/* 64-bit label subtraction */</span>
    <span id="R_RISCV_GNU_VTINHERIT">R_RISCV_GNU_VTINHERIT</span> <a href="#R_RISCV">R_RISCV</a> = 41 <span class="comment">/* GNU C++ vtable hierarchy */</span>
    <span id="R_RISCV_GNU_VTENTRY">R_RISCV_GNU_VTENTRY</span>   <a href="#R_RISCV">R_RISCV</a> = 42 <span class="comment">/* GNU C++ vtable member usage */</span>
    <span id="R_RISCV_ALIGN">R_RISCV_ALIGN</span>         <a href="#R_RISCV">R_RISCV</a> = 43 <span class="comment">/* Alignment statement */</span>
    <span id="R_RISCV_RVC_BRANCH">R_RISCV_RVC_BRANCH</span>    <a href="#R_RISCV">R_RISCV</a> = 44 <span class="comment">/* PC-relative branch offset */</span>
    <span id="R_RISCV_RVC_JUMP">R_RISCV_RVC_JUMP</span>      <a href="#R_RISCV">R_RISCV</a> = 45 <span class="comment">/* PC-relative jump offset */</span>
    <span id="R_RISCV_RVC_LUI">R_RISCV_RVC_LUI</span>       <a href="#R_RISCV">R_RISCV</a> = 46 <span class="comment">/* Absolute address */</span>
    <span id="R_RISCV_GPREL_I">R_RISCV_GPREL_I</span>       <a href="#R_RISCV">R_RISCV</a> = 47 <span class="comment">/* GP-relative reference */</span>
    <span id="R_RISCV_GPREL_S">R_RISCV_GPREL_S</span>       <a href="#R_RISCV">R_RISCV</a> = 48 <span class="comment">/* GP-relative reference */</span>
    <span id="R_RISCV_TPREL_I">R_RISCV_TPREL_I</span>       <a href="#R_RISCV">R_RISCV</a> = 49 <span class="comment">/* TP-relative TLS LE load */</span>
    <span id="R_RISCV_TPREL_S">R_RISCV_TPREL_S</span>       <a href="#R_RISCV">R_RISCV</a> = 50 <span class="comment">/* TP-relative TLS LE store */</span>
    <span id="R_RISCV_RELAX">R_RISCV_RELAX</span>         <a href="#R_RISCV">R_RISCV</a> = 51 <span class="comment">/* Instruction pair can be relaxed */</span>
    <span id="R_RISCV_SUB6">R_RISCV_SUB6</span>          <a href="#R_RISCV">R_RISCV</a> = 52 <span class="comment">/* Local label subtraction */</span>
    <span id="R_RISCV_SET6">R_RISCV_SET6</span>          <a href="#R_RISCV">R_RISCV</a> = 53 <span class="comment">/* Local label subtraction */</span>
    <span id="R_RISCV_SET8">R_RISCV_SET8</span>          <a href="#R_RISCV">R_RISCV</a> = 54 <span class="comment">/* Local label subtraction */</span>
    <span id="R_RISCV_SET16">R_RISCV_SET16</span>         <a href="#R_RISCV">R_RISCV</a> = 55 <span class="comment">/* Local label subtraction */</span>
    <span id="R_RISCV_SET32">R_RISCV_SET32</span>         <a href="#R_RISCV">R_RISCV</a> = 56 <span class="comment">/* Local label subtraction */</span>
    <span id="R_RISCV_32_PCREL">R_RISCV_32_PCREL</span>      <a href="#R_RISCV">R_RISCV</a> = 57 <span class="comment">/* 32-bit PC relative */</span>
)</pre>









### <a id="R_RISCV.GoString">func</a> (R\_RISCV) [GoString](https://golang.org/src/debug/elf/elf.go?s=99100:99134#L2478)
<pre>func (i <a href="#R_RISCV">R_RISCV</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="R_RISCV.String">func</a> (R\_RISCV) [String](https://golang.org/src/debug/elf/elf.go?s=99010:99042#L2477)
<pre>func (i <a href="#R_RISCV">R_RISCV</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="R_SPARC">type</a> [R_SPARC](https://golang.org/src/debug/elf/elf.go?s=102778:102794#L2615)
Relocation types for SPARC.


<pre>type R_SPARC <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="R_SPARC_NONE">R_SPARC_NONE</span>     <a href="#R_SPARC">R_SPARC</a> = 0
    <span id="R_SPARC_8">R_SPARC_8</span>        <a href="#R_SPARC">R_SPARC</a> = 1
    <span id="R_SPARC_16">R_SPARC_16</span>       <a href="#R_SPARC">R_SPARC</a> = 2
    <span id="R_SPARC_32">R_SPARC_32</span>       <a href="#R_SPARC">R_SPARC</a> = 3
    <span id="R_SPARC_DISP8">R_SPARC_DISP8</span>    <a href="#R_SPARC">R_SPARC</a> = 4
    <span id="R_SPARC_DISP16">R_SPARC_DISP16</span>   <a href="#R_SPARC">R_SPARC</a> = 5
    <span id="R_SPARC_DISP32">R_SPARC_DISP32</span>   <a href="#R_SPARC">R_SPARC</a> = 6
    <span id="R_SPARC_WDISP30">R_SPARC_WDISP30</span>  <a href="#R_SPARC">R_SPARC</a> = 7
    <span id="R_SPARC_WDISP22">R_SPARC_WDISP22</span>  <a href="#R_SPARC">R_SPARC</a> = 8
    <span id="R_SPARC_HI22">R_SPARC_HI22</span>     <a href="#R_SPARC">R_SPARC</a> = 9
    <span id="R_SPARC_22">R_SPARC_22</span>       <a href="#R_SPARC">R_SPARC</a> = 10
    <span id="R_SPARC_13">R_SPARC_13</span>       <a href="#R_SPARC">R_SPARC</a> = 11
    <span id="R_SPARC_LO10">R_SPARC_LO10</span>     <a href="#R_SPARC">R_SPARC</a> = 12
    <span id="R_SPARC_GOT10">R_SPARC_GOT10</span>    <a href="#R_SPARC">R_SPARC</a> = 13
    <span id="R_SPARC_GOT13">R_SPARC_GOT13</span>    <a href="#R_SPARC">R_SPARC</a> = 14
    <span id="R_SPARC_GOT22">R_SPARC_GOT22</span>    <a href="#R_SPARC">R_SPARC</a> = 15
    <span id="R_SPARC_PC10">R_SPARC_PC10</span>     <a href="#R_SPARC">R_SPARC</a> = 16
    <span id="R_SPARC_PC22">R_SPARC_PC22</span>     <a href="#R_SPARC">R_SPARC</a> = 17
    <span id="R_SPARC_WPLT30">R_SPARC_WPLT30</span>   <a href="#R_SPARC">R_SPARC</a> = 18
    <span id="R_SPARC_COPY">R_SPARC_COPY</span>     <a href="#R_SPARC">R_SPARC</a> = 19
    <span id="R_SPARC_GLOB_DAT">R_SPARC_GLOB_DAT</span> <a href="#R_SPARC">R_SPARC</a> = 20
    <span id="R_SPARC_JMP_SLOT">R_SPARC_JMP_SLOT</span> <a href="#R_SPARC">R_SPARC</a> = 21
    <span id="R_SPARC_RELATIVE">R_SPARC_RELATIVE</span> <a href="#R_SPARC">R_SPARC</a> = 22
    <span id="R_SPARC_UA32">R_SPARC_UA32</span>     <a href="#R_SPARC">R_SPARC</a> = 23
    <span id="R_SPARC_PLT32">R_SPARC_PLT32</span>    <a href="#R_SPARC">R_SPARC</a> = 24
    <span id="R_SPARC_HIPLT22">R_SPARC_HIPLT22</span>  <a href="#R_SPARC">R_SPARC</a> = 25
    <span id="R_SPARC_LOPLT10">R_SPARC_LOPLT10</span>  <a href="#R_SPARC">R_SPARC</a> = 26
    <span id="R_SPARC_PCPLT32">R_SPARC_PCPLT32</span>  <a href="#R_SPARC">R_SPARC</a> = 27
    <span id="R_SPARC_PCPLT22">R_SPARC_PCPLT22</span>  <a href="#R_SPARC">R_SPARC</a> = 28
    <span id="R_SPARC_PCPLT10">R_SPARC_PCPLT10</span>  <a href="#R_SPARC">R_SPARC</a> = 29
    <span id="R_SPARC_10">R_SPARC_10</span>       <a href="#R_SPARC">R_SPARC</a> = 30
    <span id="R_SPARC_11">R_SPARC_11</span>       <a href="#R_SPARC">R_SPARC</a> = 31
    <span id="R_SPARC_64">R_SPARC_64</span>       <a href="#R_SPARC">R_SPARC</a> = 32
    <span id="R_SPARC_OLO10">R_SPARC_OLO10</span>    <a href="#R_SPARC">R_SPARC</a> = 33
    <span id="R_SPARC_HH22">R_SPARC_HH22</span>     <a href="#R_SPARC">R_SPARC</a> = 34
    <span id="R_SPARC_HM10">R_SPARC_HM10</span>     <a href="#R_SPARC">R_SPARC</a> = 35
    <span id="R_SPARC_LM22">R_SPARC_LM22</span>     <a href="#R_SPARC">R_SPARC</a> = 36
    <span id="R_SPARC_PC_HH22">R_SPARC_PC_HH22</span>  <a href="#R_SPARC">R_SPARC</a> = 37
    <span id="R_SPARC_PC_HM10">R_SPARC_PC_HM10</span>  <a href="#R_SPARC">R_SPARC</a> = 38
    <span id="R_SPARC_PC_LM22">R_SPARC_PC_LM22</span>  <a href="#R_SPARC">R_SPARC</a> = 39
    <span id="R_SPARC_WDISP16">R_SPARC_WDISP16</span>  <a href="#R_SPARC">R_SPARC</a> = 40
    <span id="R_SPARC_WDISP19">R_SPARC_WDISP19</span>  <a href="#R_SPARC">R_SPARC</a> = 41
    <span id="R_SPARC_GLOB_JMP">R_SPARC_GLOB_JMP</span> <a href="#R_SPARC">R_SPARC</a> = 42
    <span id="R_SPARC_7">R_SPARC_7</span>        <a href="#R_SPARC">R_SPARC</a> = 43
    <span id="R_SPARC_5">R_SPARC_5</span>        <a href="#R_SPARC">R_SPARC</a> = 44
    <span id="R_SPARC_6">R_SPARC_6</span>        <a href="#R_SPARC">R_SPARC</a> = 45
    <span id="R_SPARC_DISP64">R_SPARC_DISP64</span>   <a href="#R_SPARC">R_SPARC</a> = 46
    <span id="R_SPARC_PLT64">R_SPARC_PLT64</span>    <a href="#R_SPARC">R_SPARC</a> = 47
    <span id="R_SPARC_HIX22">R_SPARC_HIX22</span>    <a href="#R_SPARC">R_SPARC</a> = 48
    <span id="R_SPARC_LOX10">R_SPARC_LOX10</span>    <a href="#R_SPARC">R_SPARC</a> = 49
    <span id="R_SPARC_H44">R_SPARC_H44</span>      <a href="#R_SPARC">R_SPARC</a> = 50
    <span id="R_SPARC_M44">R_SPARC_M44</span>      <a href="#R_SPARC">R_SPARC</a> = 51
    <span id="R_SPARC_L44">R_SPARC_L44</span>      <a href="#R_SPARC">R_SPARC</a> = 52
    <span id="R_SPARC_REGISTER">R_SPARC_REGISTER</span> <a href="#R_SPARC">R_SPARC</a> = 53
    <span id="R_SPARC_UA64">R_SPARC_UA64</span>     <a href="#R_SPARC">R_SPARC</a> = 54
    <span id="R_SPARC_UA16">R_SPARC_UA16</span>     <a href="#R_SPARC">R_SPARC</a> = 55
)</pre>









### <a id="R_SPARC.GoString">func</a> (R\_SPARC) [GoString](https://golang.org/src/debug/elf/elf.go?s=105979:106013#L2736)
<pre>func (i <a href="#R_SPARC">R_SPARC</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="R_SPARC.String">func</a> (R\_SPARC) [String](https://golang.org/src/debug/elf/elf.go?s=105889:105921#L2735)
<pre>func (i <a href="#R_SPARC">R_SPARC</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="R_X86_64">type</a> [R_X86_64](https://golang.org/src/debug/elf/elf.go?s=40958:40975#L1020)
Relocation types for x86-64.


<pre>type R_X86_64 <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="R_X86_64_NONE">R_X86_64_NONE</span>            <a href="#R_X86_64">R_X86_64</a> = 0  <span class="comment">/* No relocation. */</span>
    <span id="R_X86_64_64">R_X86_64_64</span>              <a href="#R_X86_64">R_X86_64</a> = 1  <span class="comment">/* Add 64 bit symbol value. */</span>
    <span id="R_X86_64_PC32">R_X86_64_PC32</span>            <a href="#R_X86_64">R_X86_64</a> = 2  <span class="comment">/* PC-relative 32 bit signed sym value. */</span>
    <span id="R_X86_64_GOT32">R_X86_64_GOT32</span>           <a href="#R_X86_64">R_X86_64</a> = 3  <span class="comment">/* PC-relative 32 bit GOT offset. */</span>
    <span id="R_X86_64_PLT32">R_X86_64_PLT32</span>           <a href="#R_X86_64">R_X86_64</a> = 4  <span class="comment">/* PC-relative 32 bit PLT offset. */</span>
    <span id="R_X86_64_COPY">R_X86_64_COPY</span>            <a href="#R_X86_64">R_X86_64</a> = 5  <span class="comment">/* Copy data from shared object. */</span>
    <span id="R_X86_64_GLOB_DAT">R_X86_64_GLOB_DAT</span>        <a href="#R_X86_64">R_X86_64</a> = 6  <span class="comment">/* Set GOT entry to data address. */</span>
    <span id="R_X86_64_JMP_SLOT">R_X86_64_JMP_SLOT</span>        <a href="#R_X86_64">R_X86_64</a> = 7  <span class="comment">/* Set GOT entry to code address. */</span>
    <span id="R_X86_64_RELATIVE">R_X86_64_RELATIVE</span>        <a href="#R_X86_64">R_X86_64</a> = 8  <span class="comment">/* Add load address of shared object. */</span>
    <span id="R_X86_64_GOTPCREL">R_X86_64_GOTPCREL</span>        <a href="#R_X86_64">R_X86_64</a> = 9  <span class="comment">/* Add 32 bit signed pcrel offset to GOT. */</span>
    <span id="R_X86_64_32">R_X86_64_32</span>              <a href="#R_X86_64">R_X86_64</a> = 10 <span class="comment">/* Add 32 bit zero extended symbol value */</span>
    <span id="R_X86_64_32S">R_X86_64_32S</span>             <a href="#R_X86_64">R_X86_64</a> = 11 <span class="comment">/* Add 32 bit sign extended symbol value */</span>
    <span id="R_X86_64_16">R_X86_64_16</span>              <a href="#R_X86_64">R_X86_64</a> = 12 <span class="comment">/* Add 16 bit zero extended symbol value */</span>
    <span id="R_X86_64_PC16">R_X86_64_PC16</span>            <a href="#R_X86_64">R_X86_64</a> = 13 <span class="comment">/* Add 16 bit signed extended pc relative symbol value */</span>
    <span id="R_X86_64_8">R_X86_64_8</span>               <a href="#R_X86_64">R_X86_64</a> = 14 <span class="comment">/* Add 8 bit zero extended symbol value */</span>
    <span id="R_X86_64_PC8">R_X86_64_PC8</span>             <a href="#R_X86_64">R_X86_64</a> = 15 <span class="comment">/* Add 8 bit signed extended pc relative symbol value */</span>
    <span id="R_X86_64_DTPMOD64">R_X86_64_DTPMOD64</span>        <a href="#R_X86_64">R_X86_64</a> = 16 <span class="comment">/* ID of module containing symbol */</span>
    <span id="R_X86_64_DTPOFF64">R_X86_64_DTPOFF64</span>        <a href="#R_X86_64">R_X86_64</a> = 17 <span class="comment">/* Offset in TLS block */</span>
    <span id="R_X86_64_TPOFF64">R_X86_64_TPOFF64</span>         <a href="#R_X86_64">R_X86_64</a> = 18 <span class="comment">/* Offset in static TLS block */</span>
    <span id="R_X86_64_TLSGD">R_X86_64_TLSGD</span>           <a href="#R_X86_64">R_X86_64</a> = 19 <span class="comment">/* PC relative offset to GD GOT entry */</span>
    <span id="R_X86_64_TLSLD">R_X86_64_TLSLD</span>           <a href="#R_X86_64">R_X86_64</a> = 20 <span class="comment">/* PC relative offset to LD GOT entry */</span>
    <span id="R_X86_64_DTPOFF32">R_X86_64_DTPOFF32</span>        <a href="#R_X86_64">R_X86_64</a> = 21 <span class="comment">/* Offset in TLS block */</span>
    <span id="R_X86_64_GOTTPOFF">R_X86_64_GOTTPOFF</span>        <a href="#R_X86_64">R_X86_64</a> = 22 <span class="comment">/* PC relative offset to IE GOT entry */</span>
    <span id="R_X86_64_TPOFF32">R_X86_64_TPOFF32</span>         <a href="#R_X86_64">R_X86_64</a> = 23 <span class="comment">/* Offset in static TLS block */</span>
    <span id="R_X86_64_PC64">R_X86_64_PC64</span>            <a href="#R_X86_64">R_X86_64</a> = 24 <span class="comment">/* PC relative 64-bit sign extended symbol value. */</span>
    <span id="R_X86_64_GOTOFF64">R_X86_64_GOTOFF64</span>        <a href="#R_X86_64">R_X86_64</a> = 25
    <span id="R_X86_64_GOTPC32">R_X86_64_GOTPC32</span>         <a href="#R_X86_64">R_X86_64</a> = 26
    <span id="R_X86_64_GOT64">R_X86_64_GOT64</span>           <a href="#R_X86_64">R_X86_64</a> = 27
    <span id="R_X86_64_GOTPCREL64">R_X86_64_GOTPCREL64</span>      <a href="#R_X86_64">R_X86_64</a> = 28
    <span id="R_X86_64_GOTPC64">R_X86_64_GOTPC64</span>         <a href="#R_X86_64">R_X86_64</a> = 29
    <span id="R_X86_64_GOTPLT64">R_X86_64_GOTPLT64</span>        <a href="#R_X86_64">R_X86_64</a> = 30
    <span id="R_X86_64_PLTOFF64">R_X86_64_PLTOFF64</span>        <a href="#R_X86_64">R_X86_64</a> = 31
    <span id="R_X86_64_SIZE32">R_X86_64_SIZE32</span>          <a href="#R_X86_64">R_X86_64</a> = 32
    <span id="R_X86_64_SIZE64">R_X86_64_SIZE64</span>          <a href="#R_X86_64">R_X86_64</a> = 33
    <span id="R_X86_64_GOTPC32_TLSDESC">R_X86_64_GOTPC32_TLSDESC</span> <a href="#R_X86_64">R_X86_64</a> = 34
    <span id="R_X86_64_TLSDESC_CALL">R_X86_64_TLSDESC_CALL</span>    <a href="#R_X86_64">R_X86_64</a> = 35
    <span id="R_X86_64_TLSDESC">R_X86_64_TLSDESC</span>         <a href="#R_X86_64">R_X86_64</a> = 36
    <span id="R_X86_64_IRELATIVE">R_X86_64_IRELATIVE</span>       <a href="#R_X86_64">R_X86_64</a> = 37
    <span id="R_X86_64_RELATIVE64">R_X86_64_RELATIVE64</span>      <a href="#R_X86_64">R_X86_64</a> = 38
    <span id="R_X86_64_PC32_BND">R_X86_64_PC32_BND</span>        <a href="#R_X86_64">R_X86_64</a> = 39
    <span id="R_X86_64_PLT32_BND">R_X86_64_PLT32_BND</span>       <a href="#R_X86_64">R_X86_64</a> = 40
    <span id="R_X86_64_GOTPCRELX">R_X86_64_GOTPCRELX</span>       <a href="#R_X86_64">R_X86_64</a> = 41
    <span id="R_X86_64_REX_GOTPCRELX">R_X86_64_REX_GOTPCRELX</span>   <a href="#R_X86_64">R_X86_64</a> = 42
)</pre>









### <a id="R_X86_64.GoString">func</a> (R\_X86\_64) [GoString](https://golang.org/src/debug/elf/elf.go?s=44959:44994#L1115)
<pre>func (i <a href="#R_X86_64">R_X86_64</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="R_X86_64.String">func</a> (R\_X86\_64) [String](https://golang.org/src/debug/elf/elf.go?s=44867:44900#L1114)
<pre>func (i <a href="#R_X86_64">R_X86_64</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Rel32">type</a> [Rel32](https://golang.org/src/debug/elf/elf.go?s=108390:108509#L2803)
ELF32 Relocations that don't need an addend field.


<pre>type Rel32 struct {
<span id="Rel32.Off"></span>    Off  <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Location to be relocated. */</span>
<span id="Rel32.Info"></span>    Info <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Relocation type and symbol index. */</span>
}
</pre>











## <a id="Rel64">type</a> [Rel64](https://golang.org/src/debug/elf/elf.go?s=111578:111697#L2905)
ELF64 relocations that don't need an addend field.


<pre>type Rel64 struct {
<span id="Rel64.Off"></span>    Off  <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Location to be relocated. */</span>
<span id="Rel64.Info"></span>    Info <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Relocation type and symbol index. */</span>
}
</pre>











## <a id="Rela32">type</a> [Rela32](https://golang.org/src/debug/elf/elf.go?s=108559:108712#L2809)
ELF32 Relocations that need an addend field.


<pre>type Rela32 struct {
<span id="Rela32.Off"></span>    Off    <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Location to be relocated. */</span>
<span id="Rela32.Info"></span>    Info   <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Relocation type and symbol index. */</span>
<span id="Rela32.Addend"></span>    Addend <a href="/pkg/builtin/#int32">int32</a>  <span class="comment">/* Addend. */</span>
}
</pre>











## <a id="Rela64">type</a> [Rela64](https://golang.org/src/debug/elf/elf.go?s=111750:111903#L2911)
ELF64 relocations that need an addend field.


<pre>type Rela64 struct {
<span id="Rela64.Off"></span>    Off    <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Location to be relocated. */</span>
<span id="Rela64.Info"></span>    Info   <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Relocation type and symbol index. */</span>
<span id="Rela64.Addend"></span>    Addend <a href="/pkg/builtin/#int64">int64</a>  <span class="comment">/* Addend. */</span>
}
</pre>











## <a id="Section">type</a> [Section](https://golang.org/src/debug/elf/file.go?s=1689:2218#L71)
A Section represents a single section in an ELF file.


<pre>type Section struct {
    <a href="#SectionHeader">SectionHeader</a>

    <span class="comment">// Embed ReaderAt for ReadAt method.</span>
    <span class="comment">// Do not embed SectionReader directly</span>
    <span class="comment">// to avoid having Read and Seek.</span>
    <span class="comment">// If a client wants Read and Seek it must use</span>
    <span class="comment">// Open() to avoid fighting over the seek offset</span>
    <span class="comment">// with other clients.</span>
    <span class="comment">//</span>
    <span class="comment">// ReaderAt may be nil if the section is not easily available</span>
    <span class="comment">// in a random-access form. For example, a compressed section</span>
    <span class="comment">// may have a nil ReaderAt.</span>
    <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReaderAt">ReaderAt</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="Section.Data">func</a> (\*Section) [Data](https://golang.org/src/debug/elf/file.go?s=2375:2415#L94)
<pre>func (s *<a href="#Section">Section</a>) Data() ([]<a href="/pkg/builtin/#byte">byte</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Data reads and returns the contents of the ELF section.
Even if the section is stored compressed in the ELF file,
Data returns uncompressed data.




### <a id="Section.Open">func</a> (\*Section) [Open](https://golang.org/src/debug/elf/file.go?s=2973:3011#L112)
<pre>func (s *<a href="#Section">Section</a>) Open() <a href="/pkg/io/">io</a>.<a href="/pkg/io/#ReadSeeker">ReadSeeker</a></pre>
Open returns a new ReadSeeker reading the ELF section.
Even if the section is stored compressed in the ELF file,
the ReadSeeker reads uncompressed data.




## <a id="Section32">type</a> [Section32](https://golang.org/src/debug/elf/elf.go?s=107086:107599#L2760)
ELF32 Section header.


<pre>type Section32 struct {
<span id="Section32.Name"></span>    Name      <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Section name (index into the section header string table). */</span>
<span id="Section32.Type"></span>    Type      <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Section type. */</span>
<span id="Section32.Flags"></span>    Flags     <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Section flags. */</span>
<span id="Section32.Addr"></span>    Addr      <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Address in memory image. */</span>
<span id="Section32.Off"></span>    Off       <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Offset in file. */</span>
<span id="Section32.Size"></span>    Size      <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Size in bytes. */</span>
<span id="Section32.Link"></span>    Link      <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Index of a related section. */</span>
<span id="Section32.Info"></span>    Info      <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Depends on section type. */</span>
<span id="Section32.Addralign"></span>    Addralign <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Alignment in bytes. */</span>
<span id="Section32.Entsize"></span>    Entsize   <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Size of each entry in section. */</span>
}
</pre>











## <a id="Section64">type</a> [Section64](https://golang.org/src/debug/elf/elf.go?s=110238:110751#L2861)
ELF64 Section header.


<pre>type Section64 struct {
<span id="Section64.Name"></span>    Name      <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Section name (index into the section header string table). */</span>
<span id="Section64.Type"></span>    Type      <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Section type. */</span>
<span id="Section64.Flags"></span>    Flags     <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Section flags. */</span>
<span id="Section64.Addr"></span>    Addr      <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Address in memory image. */</span>
<span id="Section64.Off"></span>    Off       <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Offset in file. */</span>
<span id="Section64.Size"></span>    Size      <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Size in bytes. */</span>
<span id="Section64.Link"></span>    Link      <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Index of a related section. */</span>
<span id="Section64.Info"></span>    Info      <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* Depends on section type. */</span>
<span id="Section64.Addralign"></span>    Addralign <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Alignment in bytes. */</span>
<span id="Section64.Entsize"></span>    Entsize   <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Size of each entry in section. */</span>
}
</pre>











## <a id="SectionFlag">type</a> [SectionFlag](https://golang.org/src/debug/elf/elf.go?s=28403:28426#L677)
Section flags.


<pre>type SectionFlag <a href="/pkg/builtin/#uint32">uint32</a></pre>



<pre>const (
    <span id="SHF_WRITE">SHF_WRITE</span>            <a href="#SectionFlag">SectionFlag</a> = 0x1        <span class="comment">/* Section contains writable data. */</span>
    <span id="SHF_ALLOC">SHF_ALLOC</span>            <a href="#SectionFlag">SectionFlag</a> = 0x2        <span class="comment">/* Section occupies memory. */</span>
    <span id="SHF_EXECINSTR">SHF_EXECINSTR</span>        <a href="#SectionFlag">SectionFlag</a> = 0x4        <span class="comment">/* Section contains instructions. */</span>
    <span id="SHF_MERGE">SHF_MERGE</span>            <a href="#SectionFlag">SectionFlag</a> = 0x10       <span class="comment">/* Section may be merged. */</span>
    <span id="SHF_STRINGS">SHF_STRINGS</span>          <a href="#SectionFlag">SectionFlag</a> = 0x20       <span class="comment">/* Section contains strings. */</span>
    <span id="SHF_INFO_LINK">SHF_INFO_LINK</span>        <a href="#SectionFlag">SectionFlag</a> = 0x40       <span class="comment">/* sh_info holds section index. */</span>
    <span id="SHF_LINK_ORDER">SHF_LINK_ORDER</span>       <a href="#SectionFlag">SectionFlag</a> = 0x80       <span class="comment">/* Special ordering requirements. */</span>
    <span id="SHF_OS_NONCONFORMING">SHF_OS_NONCONFORMING</span> <a href="#SectionFlag">SectionFlag</a> = 0x100      <span class="comment">/* OS-specific processing required. */</span>
    <span id="SHF_GROUP">SHF_GROUP</span>            <a href="#SectionFlag">SectionFlag</a> = 0x200      <span class="comment">/* Member of section group. */</span>
    <span id="SHF_TLS">SHF_TLS</span>              <a href="#SectionFlag">SectionFlag</a> = 0x400      <span class="comment">/* Section contains TLS data. */</span>
    <span id="SHF_COMPRESSED">SHF_COMPRESSED</span>       <a href="#SectionFlag">SectionFlag</a> = 0x800      <span class="comment">/* Section is compressed. */</span>
    <span id="SHF_MASKOS">SHF_MASKOS</span>           <a href="#SectionFlag">SectionFlag</a> = 0x0ff00000 <span class="comment">/* OS-specific semantics. */</span>
    <span id="SHF_MASKPROC">SHF_MASKPROC</span>         <a href="#SectionFlag">SectionFlag</a> = 0xf0000000 <span class="comment">/* Processor-specific semantics. */</span>
)</pre>









### <a id="SectionFlag.GoString">func</a> (SectionFlag) [GoString](https://golang.org/src/debug/elf/elf.go?s=29878:29916#L710)
<pre>func (i <a href="#SectionFlag">SectionFlag</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="SectionFlag.String">func</a> (SectionFlag) [String](https://golang.org/src/debug/elf/elf.go?s=29789:29825#L709)
<pre>func (i <a href="#SectionFlag">SectionFlag</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="SectionHeader">type</a> [SectionHeader](https://golang.org/src/debug/elf/file.go?s=1188:1630#L51)
A SectionHeader represents a single ELF section header.


<pre>type SectionHeader struct {
<span id="SectionHeader.Name"></span>    Name      <a href="/pkg/builtin/#string">string</a>
<span id="SectionHeader.Type"></span>    Type      <a href="#SectionType">SectionType</a>
<span id="SectionHeader.Flags"></span>    Flags     <a href="#SectionFlag">SectionFlag</a>
<span id="SectionHeader.Addr"></span>    Addr      <a href="/pkg/builtin/#uint64">uint64</a>
<span id="SectionHeader.Offset"></span>    Offset    <a href="/pkg/builtin/#uint64">uint64</a>
<span id="SectionHeader.Size"></span>    Size      <a href="/pkg/builtin/#uint64">uint64</a>
<span id="SectionHeader.Link"></span>    Link      <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader.Info"></span>    Info      <a href="/pkg/builtin/#uint32">uint32</a>
<span id="SectionHeader.Addralign"></span>    Addralign <a href="/pkg/builtin/#uint64">uint64</a>
<span id="SectionHeader.Entsize"></span>    Entsize   <a href="/pkg/builtin/#uint64">uint64</a>

<span id="SectionHeader.FileSize"></span>    <span class="comment">// FileSize is the size of this section in the file in bytes.</span>
    <span class="comment">// If a section is compressed, FileSize is the size of the</span>
    <span class="comment">// compressed data, while Size (above) is the size of the</span>
    <span class="comment">// uncompressed data.</span>
    FileSize <a href="/pkg/builtin/#uint64">uint64</a>
}
</pre>











## <a id="SectionIndex">type</a> [SectionIndex](https://golang.org/src/debug/elf/elf.go?s=24129:24150#L580)
Special section indices.


<pre>type SectionIndex <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="SHN_UNDEF">SHN_UNDEF</span>     <a href="#SectionIndex">SectionIndex</a> = 0      <span class="comment">/* Undefined, missing, irrelevant. */</span>
    <span id="SHN_LORESERVE">SHN_LORESERVE</span> <a href="#SectionIndex">SectionIndex</a> = 0xff00 <span class="comment">/* First of reserved range. */</span>
    <span id="SHN_LOPROC">SHN_LOPROC</span>    <a href="#SectionIndex">SectionIndex</a> = 0xff00 <span class="comment">/* First processor-specific. */</span>
    <span id="SHN_HIPROC">SHN_HIPROC</span>    <a href="#SectionIndex">SectionIndex</a> = 0xff1f <span class="comment">/* Last processor-specific. */</span>
    <span id="SHN_LOOS">SHN_LOOS</span>      <a href="#SectionIndex">SectionIndex</a> = 0xff20 <span class="comment">/* First operating system-specific. */</span>
    <span id="SHN_HIOS">SHN_HIOS</span>      <a href="#SectionIndex">SectionIndex</a> = 0xff3f <span class="comment">/* Last operating system-specific. */</span>
    <span id="SHN_ABS">SHN_ABS</span>       <a href="#SectionIndex">SectionIndex</a> = 0xfff1 <span class="comment">/* Absolute values. */</span>
    <span id="SHN_COMMON">SHN_COMMON</span>    <a href="#SectionIndex">SectionIndex</a> = 0xfff2 <span class="comment">/* Common data. */</span>
    <span id="SHN_XINDEX">SHN_XINDEX</span>    <a href="#SectionIndex">SectionIndex</a> = 0xffff <span class="comment">/* Escape; index stored elsewhere. */</span>
    <span id="SHN_HIRESERVE">SHN_HIRESERVE</span> <a href="#SectionIndex">SectionIndex</a> = 0xffff <span class="comment">/* Last of reserved range. */</span>
)</pre>









### <a id="SectionIndex.GoString">func</a> (SectionIndex) [GoString](https://golang.org/src/debug/elf/elf.go?s=25114:25153#L605)
<pre>func (i <a href="#SectionIndex">SectionIndex</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="SectionIndex.String">func</a> (SectionIndex) [String](https://golang.org/src/debug/elf/elf.go?s=25022:25059#L604)
<pre>func (i <a href="#SectionIndex">SectionIndex</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="SectionType">type</a> [SectionType](https://golang.org/src/debug/elf/elf.go?s=25223:25246#L608)
Section type.


<pre>type SectionType <a href="/pkg/builtin/#uint32">uint32</a></pre>



<pre>const (
    <span id="SHT_NULL">SHT_NULL</span>           <a href="#SectionType">SectionType</a> = 0          <span class="comment">/* inactive */</span>
    <span id="SHT_PROGBITS">SHT_PROGBITS</span>       <a href="#SectionType">SectionType</a> = 1          <span class="comment">/* program defined information */</span>
    <span id="SHT_SYMTAB">SHT_SYMTAB</span>         <a href="#SectionType">SectionType</a> = 2          <span class="comment">/* symbol table section */</span>
    <span id="SHT_STRTAB">SHT_STRTAB</span>         <a href="#SectionType">SectionType</a> = 3          <span class="comment">/* string table section */</span>
    <span id="SHT_RELA">SHT_RELA</span>           <a href="#SectionType">SectionType</a> = 4          <span class="comment">/* relocation section with addends */</span>
    <span id="SHT_HASH">SHT_HASH</span>           <a href="#SectionType">SectionType</a> = 5          <span class="comment">/* symbol hash table section */</span>
    <span id="SHT_DYNAMIC">SHT_DYNAMIC</span>        <a href="#SectionType">SectionType</a> = 6          <span class="comment">/* dynamic section */</span>
    <span id="SHT_NOTE">SHT_NOTE</span>           <a href="#SectionType">SectionType</a> = 7          <span class="comment">/* note section */</span>
    <span id="SHT_NOBITS">SHT_NOBITS</span>         <a href="#SectionType">SectionType</a> = 8          <span class="comment">/* no space section */</span>
    <span id="SHT_REL">SHT_REL</span>            <a href="#SectionType">SectionType</a> = 9          <span class="comment">/* relocation section - no addends */</span>
    <span id="SHT_SHLIB">SHT_SHLIB</span>          <a href="#SectionType">SectionType</a> = 10         <span class="comment">/* reserved - purpose unknown */</span>
    <span id="SHT_DYNSYM">SHT_DYNSYM</span>         <a href="#SectionType">SectionType</a> = 11         <span class="comment">/* dynamic symbol table section */</span>
    <span id="SHT_INIT_ARRAY">SHT_INIT_ARRAY</span>     <a href="#SectionType">SectionType</a> = 14         <span class="comment">/* Initialization function pointers. */</span>
    <span id="SHT_FINI_ARRAY">SHT_FINI_ARRAY</span>     <a href="#SectionType">SectionType</a> = 15         <span class="comment">/* Termination function pointers. */</span>
    <span id="SHT_PREINIT_ARRAY">SHT_PREINIT_ARRAY</span>  <a href="#SectionType">SectionType</a> = 16         <span class="comment">/* Pre-initialization function ptrs. */</span>
    <span id="SHT_GROUP">SHT_GROUP</span>          <a href="#SectionType">SectionType</a> = 17         <span class="comment">/* Section group. */</span>
    <span id="SHT_SYMTAB_SHNDX">SHT_SYMTAB_SHNDX</span>   <a href="#SectionType">SectionType</a> = 18         <span class="comment">/* Section indexes (see SHN_XINDEX). */</span>
    <span id="SHT_LOOS">SHT_LOOS</span>           <a href="#SectionType">SectionType</a> = 0x60000000 <span class="comment">/* First of OS specific semantics */</span>
    <span id="SHT_GNU_ATTRIBUTES">SHT_GNU_ATTRIBUTES</span> <a href="#SectionType">SectionType</a> = 0x6ffffff5 <span class="comment">/* GNU object attributes */</span>
    <span id="SHT_GNU_HASH">SHT_GNU_HASH</span>       <a href="#SectionType">SectionType</a> = 0x6ffffff6 <span class="comment">/* GNU hash table */</span>
    <span id="SHT_GNU_LIBLIST">SHT_GNU_LIBLIST</span>    <a href="#SectionType">SectionType</a> = 0x6ffffff7 <span class="comment">/* GNU prelink library list */</span>
    <span id="SHT_GNU_VERDEF">SHT_GNU_VERDEF</span>     <a href="#SectionType">SectionType</a> = 0x6ffffffd <span class="comment">/* GNU version definition section */</span>
    <span id="SHT_GNU_VERNEED">SHT_GNU_VERNEED</span>    <a href="#SectionType">SectionType</a> = 0x6ffffffe <span class="comment">/* GNU version needs section */</span>
    <span id="SHT_GNU_VERSYM">SHT_GNU_VERSYM</span>     <a href="#SectionType">SectionType</a> = 0x6fffffff <span class="comment">/* GNU version symbol table */</span>
    <span id="SHT_HIOS">SHT_HIOS</span>           <a href="#SectionType">SectionType</a> = 0x6fffffff <span class="comment">/* Last of OS specific semantics */</span>
    <span id="SHT_LOPROC">SHT_LOPROC</span>         <a href="#SectionType">SectionType</a> = 0x70000000 <span class="comment">/* reserved range for processor */</span>
    <span id="SHT_HIPROC">SHT_HIPROC</span>         <a href="#SectionType">SectionType</a> = 0x7fffffff <span class="comment">/* specific section header types */</span>
    <span id="SHT_LOUSER">SHT_LOUSER</span>         <a href="#SectionType">SectionType</a> = 0x80000000 <span class="comment">/* reserved range for application */</span>
    <span id="SHT_HIUSER">SHT_HIUSER</span>         <a href="#SectionType">SectionType</a> = 0xffffffff <span class="comment">/* specific indexes */</span>
)</pre>









### <a id="SectionType.GoString">func</a> (SectionType) [GoString](https://golang.org/src/debug/elf/elf.go?s=28294:28332#L674)
<pre>func (i <a href="#SectionType">SectionType</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="SectionType.String">func</a> (SectionType) [String](https://golang.org/src/debug/elf/elf.go?s=28203:28239#L673)
<pre>func (i <a href="#SectionType">SectionType</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Sym32">type</a> [Sym32](https://golang.org/src/debug/elf/elf.go?s=108914:109017#L2820)
ELF32 Symbol.


<pre>type Sym32 struct {
<span id="Sym32.Name"></span>    Name  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Sym32.Value"></span>    Value <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Sym32.Size"></span>    Size  <a href="/pkg/builtin/#uint32">uint32</a>
<span id="Sym32.Info"></span>    Info  <a href="/pkg/builtin/#uint8">uint8</a>
<span id="Sym32.Other"></span>    Other <a href="/pkg/builtin/#uint8">uint8</a>
<span id="Sym32.Shndx"></span>    Shndx <a href="/pkg/builtin/#uint16">uint16</a>
}
</pre>











## <a id="Sym64">type</a> [Sym64](https://golang.org/src/debug/elf/elf.go?s=112140:112426#L2922)
ELF64 symbol table entries.


<pre>type Sym64 struct {
<span id="Sym64.Name"></span>    Name  <a href="/pkg/builtin/#uint32">uint32</a> <span class="comment">/* String table index of name. */</span>
<span id="Sym64.Info"></span>    Info  <a href="/pkg/builtin/#uint8">uint8</a>  <span class="comment">/* Type and binding information. */</span>
<span id="Sym64.Other"></span>    Other <a href="/pkg/builtin/#uint8">uint8</a>  <span class="comment">/* Reserved (not used). */</span>
<span id="Sym64.Shndx"></span>    Shndx <a href="/pkg/builtin/#uint16">uint16</a> <span class="comment">/* Section index of symbol. */</span>
<span id="Sym64.Value"></span>    Value <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Symbol value. */</span>
<span id="Sym64.Size"></span>    Size  <a href="/pkg/builtin/#uint64">uint64</a> <span class="comment">/* Size of associated object. */</span>
}
</pre>











## <a id="SymBind">type</a> [SymBind](https://golang.org/src/debug/elf/elf.go?s=38452:38468#L936)
Symbol Binding - ELFNN_ST_BIND - st_info


<pre>type SymBind <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="STB_LOCAL">STB_LOCAL</span>  <a href="#SymBind">SymBind</a> = 0  <span class="comment">/* Local symbol */</span>
    <span id="STB_GLOBAL">STB_GLOBAL</span> <a href="#SymBind">SymBind</a> = 1  <span class="comment">/* Global symbol */</span>
    <span id="STB_WEAK">STB_WEAK</span>   <a href="#SymBind">SymBind</a> = 2  <span class="comment">/* like global - lower precedence */</span>
    <span id="STB_LOOS">STB_LOOS</span>   <a href="#SymBind">SymBind</a> = 10 <span class="comment">/* Reserved range for operating system */</span>
    <span id="STB_HIOS">STB_HIOS</span>   <a href="#SymBind">SymBind</a> = 12 <span class="comment">/*   specific semantics. */</span>
    <span id="STB_LOPROC">STB_LOPROC</span> <a href="#SymBind">SymBind</a> = 13 <span class="comment">/* reserved range for processor */</span>
    <span id="STB_HIPROC">STB_HIPROC</span> <a href="#SymBind">SymBind</a> = 15 <span class="comment">/*   specific semantics. */</span>
)</pre>







### <a id="ST_BIND">func</a> [ST_BIND](https://golang.org/src/debug/elf/elf.go?s=109041:109073#L2831)
<pre>func ST_BIND(info <a href="/pkg/builtin/#uint8">uint8</a>) <a href="#SymBind">SymBind</a></pre>





### <a id="SymBind.GoString">func</a> (SymBind) [GoString](https://golang.org/src/debug/elf/elf.go?s=39120:39154#L959)
<pre>func (i <a href="#SymBind">SymBind</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="SymBind.String">func</a> (SymBind) [String](https://golang.org/src/debug/elf/elf.go?s=39033:39065#L958)
<pre>func (i <a href="#SymBind">SymBind</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="SymType">type</a> [SymType](https://golang.org/src/debug/elf/elf.go?s=39251:39267#L962)
Symbol type - ELFNN_ST_TYPE - st_info


<pre>type SymType <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="STT_NOTYPE">STT_NOTYPE</span>  <a href="#SymType">SymType</a> = 0  <span class="comment">/* Unspecified type. */</span>
    <span id="STT_OBJECT">STT_OBJECT</span>  <a href="#SymType">SymType</a> = 1  <span class="comment">/* Data object. */</span>
    <span id="STT_FUNC">STT_FUNC</span>    <a href="#SymType">SymType</a> = 2  <span class="comment">/* Function. */</span>
    <span id="STT_SECTION">STT_SECTION</span> <a href="#SymType">SymType</a> = 3  <span class="comment">/* Section. */</span>
    <span id="STT_FILE">STT_FILE</span>    <a href="#SymType">SymType</a> = 4  <span class="comment">/* Source file. */</span>
    <span id="STT_COMMON">STT_COMMON</span>  <a href="#SymType">SymType</a> = 5  <span class="comment">/* Uninitialized common block. */</span>
    <span id="STT_TLS">STT_TLS</span>     <a href="#SymType">SymType</a> = 6  <span class="comment">/* TLS object. */</span>
    <span id="STT_LOOS">STT_LOOS</span>    <a href="#SymType">SymType</a> = 10 <span class="comment">/* Reserved range for operating system */</span>
    <span id="STT_HIOS">STT_HIOS</span>    <a href="#SymType">SymType</a> = 12 <span class="comment">/*   specific semantics. */</span>
    <span id="STT_LOPROC">STT_LOPROC</span>  <a href="#SymType">SymType</a> = 13 <span class="comment">/* reserved range for processor */</span>
    <span id="STT_HIPROC">STT_HIPROC</span>  <a href="#SymType">SymType</a> = 15 <span class="comment">/*   specific semantics. */</span>
)</pre>







### <a id="ST_TYPE">func</a> [ST_TYPE](https://golang.org/src/debug/elf/elf.go?s=109104:109136#L2832)
<pre>func ST_TYPE(info <a href="/pkg/builtin/#uint8">uint8</a>) <a href="#SymType">SymType</a></pre>





### <a id="SymType.GoString">func</a> (SymType) [GoString](https://golang.org/src/debug/elf/elf.go?s=40176:40210#L993)
<pre>func (i <a href="#SymType">SymType</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="SymType.String">func</a> (SymType) [String](https://golang.org/src/debug/elf/elf.go?s=40089:40121#L992)
<pre>func (i <a href="#SymType">SymType</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="SymVis">type</a> [SymVis](https://golang.org/src/debug/elf/elf.go?s=40320:40335#L996)
Symbol visibility - ELFNN_ST_VISIBILITY - st_other


<pre>type SymVis <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="STV_DEFAULT">STV_DEFAULT</span>   <a href="#SymVis">SymVis</a> = 0x0 <span class="comment">/* Default visibility (see binding). */</span>
    <span id="STV_INTERNAL">STV_INTERNAL</span>  <a href="#SymVis">SymVis</a> = 0x1 <span class="comment">/* Special meaning in relocatable objects. */</span>
    <span id="STV_HIDDEN">STV_HIDDEN</span>    <a href="#SymVis">SymVis</a> = 0x2 <span class="comment">/* Not visible. */</span>
    <span id="STV_PROTECTED">STV_PROTECTED</span> <a href="#SymVis">SymVis</a> = 0x3 <span class="comment">/* Visible but not preemptible. */</span>
)</pre>







### <a id="ST_VISIBILITY">func</a> [ST_VISIBILITY](https://golang.org/src/debug/elf/elf.go?s=109258:109296#L2836)
<pre>func ST_VISIBILITY(other <a href="/pkg/builtin/#uint8">uint8</a>) <a href="#SymVis">SymVis</a></pre>





### <a id="SymVis.GoString">func</a> (SymVis) [GoString](https://golang.org/src/debug/elf/elf.go?s=40811:40844#L1013)
<pre>func (i <a href="#SymVis">SymVis</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="SymVis.String">func</a> (SymVis) [String](https://golang.org/src/debug/elf/elf.go?s=40725:40756#L1012)
<pre>func (i <a href="#SymVis">SymVis</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Symbol">type</a> [Symbol](https://golang.org/src/debug/elf/file.go?s=4278:4492#L159)
A Symbol represents an entry in an ELF symbol table section.


<pre>type Symbol struct {
<span id="Symbol.Name"></span>    Name        <a href="/pkg/builtin/#string">string</a>
<span id="Symbol.Info"></span>    Info, Other <a href="/pkg/builtin/#byte">byte</a>
<span id="Symbol.Section"></span>    Section     <a href="#SectionIndex">SectionIndex</a>
<span id="Symbol.Value"></span>    Value, Size <a href="/pkg/builtin/#uint64">uint64</a>

<span id="Symbol.Version"></span>    <span class="comment">// Version and Library are present only for the dynamic symbol</span>
    <span class="comment">// table.</span>
    Version <a href="/pkg/builtin/#string">string</a>
<span id="Symbol.Library"></span>    Library <a href="/pkg/builtin/#string">string</a>
}
</pre>











## <a id="Type">type</a> [Type](https://golang.org/src/debug/elf/elf.go?s=6385:6401#L162)
Type is found in Header.Type.


<pre>type Type <a href="/pkg/builtin/#uint16">uint16</a></pre>



<pre>const (
    <span id="ET_NONE">ET_NONE</span>   <a href="#Type">Type</a> = 0      <span class="comment">/* Unknown type. */</span>
    <span id="ET_REL">ET_REL</span>    <a href="#Type">Type</a> = 1      <span class="comment">/* Relocatable. */</span>
    <span id="ET_EXEC">ET_EXEC</span>   <a href="#Type">Type</a> = 2      <span class="comment">/* Executable. */</span>
    <span id="ET_DYN">ET_DYN</span>    <a href="#Type">Type</a> = 3      <span class="comment">/* Shared object. */</span>
    <span id="ET_CORE">ET_CORE</span>   <a href="#Type">Type</a> = 4      <span class="comment">/* Core file. */</span>
    <span id="ET_LOOS">ET_LOOS</span>   <a href="#Type">Type</a> = 0xfe00 <span class="comment">/* First operating system specific. */</span>
    <span id="ET_HIOS">ET_HIOS</span>   <a href="#Type">Type</a> = 0xfeff <span class="comment">/* Last operating system-specific. */</span>
    <span id="ET_LOPROC">ET_LOPROC</span> <a href="#Type">Type</a> = 0xff00 <span class="comment">/* First processor-specific. */</span>
    <span id="ET_HIPROC">ET_HIPROC</span> <a href="#Type">Type</a> = 0xffff <span class="comment">/* Last processor-specific. */</span>
)</pre>









### <a id="Type.GoString">func</a> (Type) [GoString](https://golang.org/src/debug/elf/elf.go?s=7166:7197#L189)
<pre>func (i <a href="#Type">Type</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Type.String">func</a> (Type) [String](https://golang.org/src/debug/elf/elf.go?s=7081:7110#L188)
<pre>func (i <a href="#Type">Type</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Version">type</a> [Version](https://golang.org/src/debug/elf/elf.go?s=3133:3150#L58)
Version is found in Header.Ident[EI_VERSION] and Header.Version.


<pre>type Version <a href="/pkg/builtin/#byte">byte</a></pre>



<pre>const (
    <span id="EV_NONE">EV_NONE</span>    <a href="#Version">Version</a> = 0
    <span id="EV_CURRENT">EV_CURRENT</span> <a href="#Version">Version</a> = 1
)</pre>









### <a id="Version.GoString">func</a> (Version) [GoString](https://golang.org/src/debug/elf/elf.go?s=3374:3408#L71)
<pre>func (i <a href="#Version">Version</a>) GoString() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="Version.String">func</a> (Version) [String](https://golang.org/src/debug/elf/elf.go?s=3283:3315#L70)
<pre>func (i <a href="#Version">Version</a>) String() <a href="/pkg/builtin/#string">string</a></pre>







