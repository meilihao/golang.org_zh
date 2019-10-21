

# flag
`import "flag"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package flag implements command-line flag parsing.

### Usage
Define flags using flag.String(), Bool(), Int(), etc.

This declares an integer flag, -flagname, stored in the pointer ip, with type *int.


	import "flag"
	var ip = flag.Int("flagname", 1234, "help message for flagname")

If you like, you can bind the flag to a variable using the Var() functions.


	var flagvar int
	func init() {
		flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
	}

Or you can create custom flags that satisfy the Value interface (with
pointer receivers) and couple them to flag parsing by


	flag.Var(&flagVal, "name", "help message for flagname")

For such flags, the default value is just the initial value of the variable.

After all flags are defined, call


	flag.Parse()

to parse the command line into the defined flags.

Flags may then be used directly. If you're using the flags themselves,
they are all pointers; if you bind to variables, they're values.


	fmt.Println("ip has value ", *ip)
	fmt.Println("flagvar has value ", flagvar)

After parsing, the arguments following the flags are available as the
slice flag.Args() or individually as flag.Arg(i).
The arguments are indexed from 0 through flag.NArg()-1.

### Command line flag syntax
The following forms are permitted:


	-flag
	-flag=x
	-flag x  // non-boolean flags only

One or two minus signs may be used; they are equivalent.
The last form is not permitted for boolean flags because the
meaning of the command


	cmd -x *

where * is a Unix shell wildcard, will change if there is a file
called 0, false, etc. You must use the -flag=false form to turn
off a boolean flag.

Flag parsing stops just before the first non-flag argument
("-" is a non-flag argument) or after the terminator "--".

Integer flags accept 1234, 0664, 0x1234 and may be negative.
Boolean flags may be:


	1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False

Duration flags accept any input valid for time.ParseDuration.

The default set of command-line flags is controlled by
top-level functions.  The FlagSet type allows one to define
independent sets of flags, such as to implement subcommands
in a command-line interface. The methods of FlagSet are
analogous to the top-level functions for the command-line
flag set.


<a id="example_">Example</a>
```go
```

output:
```txt
```


## <a id="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func Arg(i int) string](#Arg)
* [func Args() []string](#Args)
* [func Bool(name string, value bool, usage string) *bool](#Bool)
* [func BoolVar(p *bool, name string, value bool, usage string)](#BoolVar)
* [func Duration(name string, value time.Duration, usage string) *time.Duration](#Duration)
* [func DurationVar(p *time.Duration, name string, value time.Duration, usage string)](#DurationVar)
* [func Float64(name string, value float64, usage string) *float64](#Float64)
* [func Float64Var(p *float64, name string, value float64, usage string)](#Float64Var)
* [func Int(name string, value int, usage string) *int](#Int)
* [func Int64(name string, value int64, usage string) *int64](#Int64)
* [func Int64Var(p *int64, name string, value int64, usage string)](#Int64Var)
* [func IntVar(p *int, name string, value int, usage string)](#IntVar)
* [func NArg() int](#NArg)
* [func NFlag() int](#NFlag)
* [func Parse()](#Parse)
* [func Parsed() bool](#Parsed)
* [func PrintDefaults()](#PrintDefaults)
* [func Set(name, value string) error](#Set)
* [func String(name string, value string, usage string) *string](#String)
* [func StringVar(p *string, name string, value string, usage string)](#StringVar)
* [func Uint(name string, value uint, usage string) *uint](#Uint)
* [func Uint64(name string, value uint64, usage string) *uint64](#Uint64)
* [func Uint64Var(p *uint64, name string, value uint64, usage string)](#Uint64Var)
* [func UintVar(p *uint, name string, value uint, usage string)](#UintVar)
* [func UnquoteUsage(flag *Flag) (name string, usage string)](#UnquoteUsage)
* [func Var(value Value, name string, usage string)](#Var)
* [func Visit(fn func(*Flag))](#Visit)
* [func VisitAll(fn func(*Flag))](#VisitAll)
* [type ErrorHandling](#ErrorHandling)
* [type Flag](#Flag)
  * [func Lookup(name string) *Flag](#Lookup)
* [type FlagSet](#FlagSet)
  * [func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet](#NewFlagSet)
  * [func (f *FlagSet) Arg(i int) string](#FlagSet.Arg)
  * [func (f *FlagSet) Args() []string](#FlagSet.Args)
  * [func (f *FlagSet) Bool(name string, value bool, usage string) *bool](#FlagSet.Bool)
  * [func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string)](#FlagSet.BoolVar)
  * [func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration](#FlagSet.Duration)
  * [func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string)](#FlagSet.DurationVar)
  * [func (f *FlagSet) ErrorHandling() ErrorHandling](#FlagSet.ErrorHandling)
  * [func (f *FlagSet) Float64(name string, value float64, usage string) *float64](#FlagSet.Float64)
  * [func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string)](#FlagSet.Float64Var)
  * [func (f *FlagSet) Init(name string, errorHandling ErrorHandling)](#FlagSet.Init)
  * [func (f *FlagSet) Int(name string, value int, usage string) *int](#FlagSet.Int)
  * [func (f *FlagSet) Int64(name string, value int64, usage string) *int64](#FlagSet.Int64)
  * [func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string)](#FlagSet.Int64Var)
  * [func (f *FlagSet) IntVar(p *int, name string, value int, usage string)](#FlagSet.IntVar)
  * [func (f *FlagSet) Lookup(name string) *Flag](#FlagSet.Lookup)
  * [func (f *FlagSet) NArg() int](#FlagSet.NArg)
  * [func (f *FlagSet) NFlag() int](#FlagSet.NFlag)
  * [func (f *FlagSet) Name() string](#FlagSet.Name)
  * [func (f *FlagSet) Output() io.Writer](#FlagSet.Output)
  * [func (f *FlagSet) Parse(arguments []string) error](#FlagSet.Parse)
  * [func (f *FlagSet) Parsed() bool](#FlagSet.Parsed)
  * [func (f *FlagSet) PrintDefaults()](#FlagSet.PrintDefaults)
  * [func (f *FlagSet) Set(name, value string) error](#FlagSet.Set)
  * [func (f *FlagSet) SetOutput(output io.Writer)](#FlagSet.SetOutput)
  * [func (f *FlagSet) String(name string, value string, usage string) *string](#FlagSet.String)
  * [func (f *FlagSet) StringVar(p *string, name string, value string, usage string)](#FlagSet.StringVar)
  * [func (f *FlagSet) Uint(name string, value uint, usage string) *uint](#FlagSet.Uint)
  * [func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64](#FlagSet.Uint64)
  * [func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string)](#FlagSet.Uint64Var)
  * [func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string)](#FlagSet.UintVar)
  * [func (f *FlagSet) Var(value Value, name string, usage string)](#FlagSet.Var)
  * [func (f *FlagSet) Visit(fn func(*Flag))](#FlagSet.Visit)
  * [func (f *FlagSet) VisitAll(fn func(*Flag))](#FlagSet.VisitAll)
* [type Getter](#Getter)
* [type Value](#Value)


#### <a id="pkg-examples">Examples</a>
* [Package](#example_)
* [Value](#example_Value)


#### <a id="pkg-files">Package files</a>
[flag.go](https://golang.org/src/flag/flag.go) 




## <a id="pkg-variables">Variables</a>
CommandLine is the default set of command-line flags, parsed from os.Args.
The top-level functions such as BoolVar, Arg, and so on are wrappers for the
methods of CommandLine.


<pre>var <span id="CommandLine">CommandLine</span> = <a href="#NewFlagSet">NewFlagSet</a>(<a href="/pkg/os/">os</a>.<a href="/pkg/os/#Args">Args</a>[0], <a href="#ExitOnError">ExitOnError</a>)</pre>ErrHelp is the error returned if the -help or -h flag is invoked
but no such flag is defined.


<pre>var <span id="ErrHelp">ErrHelp</span> = <a href="/pkg/errors/">errors</a>.<a href="/pkg/errors/#New">New</a>(&#34;flag: help requested&#34;)</pre>Usage prints a usage message documenting all defined command-line flags
to CommandLine's output, which by default is os.Stderr.
It is called when an error occurs while parsing flags.
The function is a variable that may be changed to point to a custom function.
By default it prints a simple header and calls PrintDefaults; for details about the
format of the output and how to control it, see the documentation for PrintDefaults.
Custom usage functions may choose to exit the program; by default exiting
happens anyway as the command line's error handling strategy is set to
ExitOnError.


<pre>var <span id="Usage">Usage</span> = func() {
    <a href="/pkg/fmt/">fmt</a>.<a href="/pkg/fmt/#Fprintf">Fprintf</a>(<a href="#CommandLine">CommandLine</a>.<a href="#Output">Output</a>(), &#34;Usage of %s:\n&#34;, <a href="/pkg/os/">os</a>.<a href="/pkg/os/#Args">Args</a>[0])
    <a href="#PrintDefaults">PrintDefaults</a>()
}</pre>

## <a id="Arg">func</a> [Arg](https://golang.org/src/flag/flag.go?s=17616:17638#L592)
<pre>func Arg(i <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#string">string</a></pre>
Arg returns the i'th command-line argument. Arg(0) is the first remaining argument
after flags have been processed. Arg returns an empty string if the
requested element does not exist.



## <a id="Args">func</a> [Args](https://golang.org/src/flag/flag.go?s=18076:18096#L606)
<pre>func Args() []<a href="/pkg/builtin/#string">string</a></pre>
Args returns the non-flag command-line arguments.



## <a id="Bool">func</a> [Bool](https://golang.org/src/flag/flag.go?s=19184:19238#L630)
<pre>func Bool(name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#bool">bool</a>, usage <a href="/pkg/builtin/#string">string</a>) *<a href="/pkg/builtin/#bool">bool</a></pre>
Bool defines a bool flag with specified name, default value, and usage string.
The return value is the address of a bool variable that stores the value of the flag.



## <a id="BoolVar">func</a> [BoolVar](https://golang.org/src/flag/flag.go?s=18589:18649#L616)
<pre>func BoolVar(p *<a href="/pkg/builtin/#bool">bool</a>, name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#bool">bool</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
BoolVar defines a bool flag with specified name, default value, and usage string.
The argument p points to a bool variable in which to store the value of the flag.



## <a id="Duration">func</a> [Duration](https://golang.org/src/flag/flag.go?s=27967:28043#L816)
<pre>func Duration(name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>, usage <a href="/pkg/builtin/#string">string</a>) *<a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a></pre>
Duration defines a time.Duration flag with specified name, default value, and usage string.
The return value is the address of a time.Duration variable that stores the value of the flag.
The flag accepts a value acceptable to time.ParseDuration.



## <a id="DurationVar">func</a> [DurationVar](https://golang.org/src/flag/flag.go?s=27143:27225#L800)
<pre>func DurationVar(p *<a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>, name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
DurationVar defines a time.Duration flag with specified name, default value, and usage string.
The argument p points to a time.Duration variable in which to store the value of the flag.
The flag accepts a value acceptable to time.ParseDuration.



## <a id="Float64">func</a> [Float64](https://golang.org/src/flag/flag.go?s=26369:26432#L786)
<pre>func Float64(name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#float64">float64</a>, usage <a href="/pkg/builtin/#string">string</a>) *<a href="/pkg/builtin/#float64">float64</a></pre>
Float64 defines a float64 flag with specified name, default value, and usage string.
The return value is the address of a float64 variable that stores the value of the flag.



## <a id="Float64Var">func</a> [Float64Var](https://golang.org/src/flag/flag.go?s=25729:25798#L772)
<pre>func Float64Var(p *<a href="/pkg/builtin/#float64">float64</a>, name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#float64">float64</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
Float64Var defines a float64 flag with specified name, default value, and usage string.
The argument p points to a float64 variable in which to store the value of the flag.



## <a id="Int">func</a> [Int](https://golang.org/src/flag/flag.go?s=20330:20381#L656)
<pre>func Int(name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#int">int</a>, usage <a href="/pkg/builtin/#string">string</a>) *<a href="/pkg/builtin/#int">int</a></pre>
Int defines an int flag with specified name, default value, and usage string.
The return value is the address of an int variable that stores the value of the flag.



## <a id="Int64">func</a> [Int64](https://golang.org/src/flag/flag.go?s=21522:21579#L682)
<pre>func Int64(name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#int64">int64</a>, usage <a href="/pkg/builtin/#string">string</a>) *<a href="/pkg/builtin/#int64">int64</a></pre>
Int64 defines an int64 flag with specified name, default value, and usage string.
The return value is the address of an int64 variable that stores the value of the flag.



## <a id="Int64Var">func</a> [Int64Var](https://golang.org/src/flag/flag.go?s=20908:20971#L668)
<pre>func Int64Var(p *<a href="/pkg/builtin/#int64">int64</a>, name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#int64">int64</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
Int64Var defines an int64 flag with specified name, default value, and usage string.
The argument p points to an int64 variable in which to store the value of the flag.



## <a id="IntVar">func</a> [IntVar](https://golang.org/src/flag/flag.go?s=19746:19803#L642)
<pre>func IntVar(p *<a href="/pkg/builtin/#int">int</a>, name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#int">int</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
IntVar defines an int flag with specified name, default value, and usage string.
The argument p points to an int variable in which to store the value of the flag.



## <a id="NArg">func</a> [NArg](https://golang.org/src/flag/flag.go?s=17880:17895#L600)
<pre>func NArg() <a href="/pkg/builtin/#int">int</a></pre>
NArg is the number of arguments remaining after flags have been processed.



## <a id="NFlag">func</a> [NFlag](https://golang.org/src/flag/flag.go?s=17082:17098#L577)
<pre>func NFlag() <a href="/pkg/builtin/#int">int</a></pre>
NFlag returns the number of command-line flags that have been set.



## <a id="Parse">func</a> [Parse](https://golang.org/src/flag/flag.go?s=32954:32966#L984)
<pre>func Parse()</pre>
Parse parses the command-line flags from os.Args[1:]. Must be called
after all flags are defined and before flags are accessed by the program.



## <a id="Parsed">func</a> [Parsed](https://golang.org/src/flag/flag.go?s=33126:33144#L990)
<pre>func Parsed() <a href="/pkg/builtin/#bool">bool</a></pre>
Parsed reports whether the command-line flags have been parsed.



## <a id="PrintDefaults">func</a> [PrintDefaults](https://golang.org/src/flag/flag.go?s=15720:15740#L541)
<pre>func PrintDefaults()</pre>
PrintDefaults prints, to standard error unless configured otherwise,
a usage message showing the default settings of all defined
command-line flags.
For an integer valued flag x, the default output has the form


	-x int
		usage-message-for-x (default 7)

The usage message will appear on a separate line for anything but
a bool flag with a one-byte name. For bool flags, the type is
omitted and if the flag name is one byte the usage message appears
on the same line. The parenthetical default is omitted if the
default is the zero value for the type. The listed type, here int,
can be changed by placing a back-quoted name in the flag's usage
string; the first such item in the message is taken to be a parameter
name to show in the message and the back quotes are stripped from
the message when displayed. For instance, given


	flag.String("I", "", "search `directory` for include files")

the output will be


	-I directory
		search directory for include files.

To change the destination for flag messages, call CommandLine.SetOutput.



## <a id="Set">func</a> [Set](https://golang.org/src/flag/flag.go?s=11795:11829#L428)
<pre>func Set(name, value <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Set sets the value of the named command-line flag.



## <a id="String">func</a> [String](https://golang.org/src/flag/flag.go?s=25123:25183#L760)
<pre>func String(name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#string">string</a>, usage <a href="/pkg/builtin/#string">string</a>) *<a href="/pkg/builtin/#string">string</a></pre>
String defines a string flag with specified name, default value, and usage string.
The return value is the address of a string variable that stores the value of the flag.



## <a id="StringVar">func</a> [StringVar](https://golang.org/src/flag/flag.go?s=24498:24564#L746)
<pre>func StringVar(p *<a href="/pkg/builtin/#string">string</a>, name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#string">string</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
StringVar defines a string flag with specified name, default value, and usage string.
The argument p points to a string variable in which to store the value of the flag.



## <a id="Uint">func</a> [Uint](https://golang.org/src/flag/flag.go?s=22689:22743#L708)
<pre>func Uint(name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#uint">uint</a>, usage <a href="/pkg/builtin/#string">string</a>) *<a href="/pkg/builtin/#uint">uint</a></pre>
Uint defines a uint flag with specified name, default value, and usage string.
The return value is the address of a uint variable that stores the value of the flag.



## <a id="Uint64">func</a> [Uint64](https://golang.org/src/flag/flag.go?s=23902:23962#L734)
<pre>func Uint64(name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#uint64">uint64</a>, usage <a href="/pkg/builtin/#string">string</a>) *<a href="/pkg/builtin/#uint64">uint64</a></pre>
Uint64 defines a uint64 flag with specified name, default value, and usage string.
The return value is the address of a uint64 variable that stores the value of the flag.



## <a id="Uint64Var">func</a> [Uint64Var](https://golang.org/src/flag/flag.go?s=23277:23343#L720)
<pre>func Uint64Var(p *<a href="/pkg/builtin/#uint64">uint64</a>, name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#uint64">uint64</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
Uint64Var defines a uint64 flag with specified name, default value, and usage string.
The argument p points to a uint64 variable in which to store the value of the flag.



## <a id="UintVar">func</a> [UintVar](https://golang.org/src/flag/flag.go?s=22094:22154#L694)
<pre>func UintVar(p *<a href="/pkg/builtin/#uint">uint</a>, name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#uint">uint</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
UintVar defines a uint flag with specified name, default value, and usage string.
The argument p points to a uint variable in which to store the value of the flag.



## <a id="UnquoteUsage">func</a> [UnquoteUsage](https://golang.org/src/flag/flag.go?s=12741:12798#L453)
<pre>func UnquoteUsage(flag *<a href="#Flag">Flag</a>) (name <a href="/pkg/builtin/#string">string</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
UnquoteUsage extracts a back-quoted name from the usage
string for a flag and returns it and the un-quoted usage.
Given "a `name` to show" it returns ("name", "a name to show").
If there are no back quotes, the name is an educated guess of the
type of the flag's value, or the empty string if the flag is boolean.



## <a id="Var">func</a> [Var](https://golang.org/src/flag/flag.go?s=29585:29633#L852)
<pre>func Var(value <a href="#Value">Value</a>, name <a href="/pkg/builtin/#string">string</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
Var defines a flag with the specified name and usage string. The type and
value of the flag are represented by the first argument, of type Value, which
typically holds a user-defined implementation of Value. For instance, the
caller could create a flag that turns a comma-separated string into a slice
of strings by giving the slice the methods of Value; in particular, Set would
decompose the comma-separated string into the slice.



## <a id="Visit">func</a> [Visit](https://golang.org/src/flag/flag.go?s=11017:11043#L395)
<pre>func Visit(fn func(*<a href="#Flag">Flag</a>))</pre>
Visit visits the command-line flags in lexicographical order, calling fn
for each. It visits only those flags that have been set.



## <a id="VisitAll">func</a> [VisitAll](https://golang.org/src/flag/flag.go?s=10594:10623#L381)
<pre>func VisitAll(fn func(*<a href="#Flag">Flag</a>))</pre>
VisitAll visits the command-line flags in lexicographical order, calling
fn for each. It visits all flags, even those not set.





## <a id="ErrorHandling">type</a> [ErrorHandling](https://golang.org/src/flag/flag.go?s=7915:7937#L296)
ErrorHandling defines how FlagSet.Parse behaves if the parse fails.


<pre>type ErrorHandling <a href="/pkg/builtin/#int">int</a></pre>


These constants cause FlagSet.Parse to behave as described if the parse fails.


<pre>const (
    <span id="ContinueOnError">ContinueOnError</span> <a href="#ErrorHandling">ErrorHandling</a> = <a href="/pkg/builtin/#iota">iota</a> <span class="comment">// Return a descriptive error.</span>
    <span id="ExitOnError">ExitOnError</span>                          <span class="comment">// Call os.Exit(2).</span>
    <span id="PanicOnError">PanicOnError</span>                         <span class="comment">// Call panic with a descriptive error.</span>
)</pre>









## <a id="Flag">type</a> [Flag](https://golang.org/src/flag/flag.go?s=9055:9259#L325)
A Flag represents the state of a flag.


<pre>type Flag struct {
<span id="Flag.Name"></span>    Name     <a href="/pkg/builtin/#string">string</a> <span class="comment">// name as it appears on command line</span>
<span id="Flag.Usage"></span>    Usage    <a href="/pkg/builtin/#string">string</a> <span class="comment">// help message</span>
<span id="Flag.Value"></span>    Value    <a href="#Value">Value</a>  <span class="comment">// value as set</span>
<span id="Flag.DefValue"></span>    DefValue <a href="/pkg/builtin/#string">string</a> <span class="comment">// default value (as text); for usage message</span>
}
</pre>









### <a id="Lookup">func</a> [Lookup](https://golang.org/src/flag/flag.go?s=11332:11362#L406)
<pre>func Lookup(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Flag">Flag</a></pre>
Lookup returns the Flag structure of the named command-line flag,
returning nil if none exists.






## <a id="FlagSet">type</a> [FlagSet](https://golang.org/src/flag/flag.go?s=8368:9011#L307)
A FlagSet represents a set of defined flags. The zero value of a FlagSet
has no name and has ContinueOnError error handling.


<pre>type FlagSet struct {
<span id="FlagSet.Usage"></span>    <span class="comment">// Usage is the function called when an error occurs while parsing flags.</span>
    <span class="comment">// The field is a function (not a method) that may be changed to point to</span>
    <span class="comment">// a custom error handler. What happens after Usage is called depends</span>
    <span class="comment">// on the ErrorHandling setting; for the command line, this defaults</span>
    <span class="comment">// to ExitOnError, which exits the program after calling Usage.</span>
    Usage func()
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewFlagSet">func</a> [NewFlagSet](https://golang.org/src/flag/flag.go?s=33952:34018#L1014)
<pre>func NewFlagSet(name <a href="/pkg/builtin/#string">string</a>, errorHandling <a href="#ErrorHandling">ErrorHandling</a>) *<a href="#FlagSet">FlagSet</a></pre>
NewFlagSet returns a new, empty flag set with the specified name and
error handling property. If the name is not empty, it will be printed
in the default usage message and in error messages.






### <a id="FlagSet.Arg">func</a> (\*FlagSet) [Arg](https://golang.org/src/flag/flag.go?s=17316:17351#L582)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Arg(i <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#string">string</a></pre>
Arg returns the i'th argument. Arg(0) is the first remaining argument
after flags have been processed. Arg returns an empty string if the
requested element does not exist.




### <a id="FlagSet.Args">func</a> (\*FlagSet) [Args](https://golang.org/src/flag/flag.go?s=17970:18003#L603)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Args() []<a href="/pkg/builtin/#string">string</a></pre>
Args returns the non-flag arguments.




### <a id="FlagSet.Bool">func</a> (\*FlagSet) [Bool](https://golang.org/src/flag/flag.go?s=18880:18947#L622)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Bool(name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#bool">bool</a>, usage <a href="/pkg/builtin/#string">string</a>) *<a href="/pkg/builtin/#bool">bool</a></pre>
Bool defines a bool flag with specified name, default value, and usage string.
The return value is the address of a bool variable that stores the value of the flag.




### <a id="FlagSet.BoolVar">func</a> (\*FlagSet) [BoolVar](https://golang.org/src/flag/flag.go?s=18296:18369#L610)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) BoolVar(p *<a href="/pkg/builtin/#bool">bool</a>, name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#bool">bool</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
BoolVar defines a bool flag with specified name, default value, and usage string.
The argument p points to a bool variable in which to store the value of the flag.




### <a id="FlagSet.Duration">func</a> (\*FlagSet) [Duration](https://golang.org/src/flag/flag.go?s=27544:27633#L807)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Duration(name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>, usage <a href="/pkg/builtin/#string">string</a>) *<a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a></pre>
Duration defines a time.Duration flag with specified name, default value, and usage string.
The return value is the address of a time.Duration variable that stores the value of the flag.
The flag accepts a value acceptable to time.ParseDuration.




### <a id="FlagSet.DurationVar">func</a> (\*FlagSet) [DurationVar](https://golang.org/src/flag/flag.go?s=26740:26835#L793)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) DurationVar(p *<a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>, name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/time/">time</a>.<a href="/pkg/time/#Duration">Duration</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
DurationVar defines a time.Duration flag with specified name, default value, and usage string.
The argument p points to a time.Duration variable in which to store the value of the flag.
The flag accepts a value acceptable to time.ParseDuration.




### <a id="FlagSet.ErrorHandling">func</a> (\*FlagSet) [ErrorHandling](https://golang.org/src/flag/flag.go?s=9984:10031#L361)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) ErrorHandling() <a href="#ErrorHandling">ErrorHandling</a></pre>
ErrorHandling returns the error handling behavior of the flag set.




### <a id="FlagSet.Float64">func</a> (\*FlagSet) [Float64](https://golang.org/src/flag/flag.go?s=26041:26117#L778)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Float64(name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#float64">float64</a>, usage <a href="/pkg/builtin/#string">string</a>) *<a href="/pkg/builtin/#float64">float64</a></pre>
Float64 defines a float64 flag with specified name, default value, and usage string.
The return value is the address of a float64 variable that stores the value of the flag.




### <a id="FlagSet.Float64Var">func</a> (\*FlagSet) [Float64Var](https://golang.org/src/flag/flag.go?s=25415:25497#L766)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Float64Var(p *<a href="/pkg/builtin/#float64">float64</a>, name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#float64">float64</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
Float64Var defines a float64 flag with specified name, default value, and usage string.
The argument p points to a float64 variable in which to store the value of the flag.




### <a id="FlagSet.Init">func</a> (\*FlagSet) [Init](https://golang.org/src/flag/flag.go?s=34301:34365#L1026)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Init(name <a href="/pkg/builtin/#string">string</a>, errorHandling <a href="#ErrorHandling">ErrorHandling</a>)</pre>
Init sets the name and error handling property for a flag set.
By default, the zero FlagSet uses an empty name and the
ContinueOnError error handling policy.




### <a id="FlagSet.Int">func</a> (\*FlagSet) [Int](https://golang.org/src/flag/flag.go?s=20032:20096#L648)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Int(name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#int">int</a>, usage <a href="/pkg/builtin/#string">string</a>) *<a href="/pkg/builtin/#int">int</a></pre>
Int defines an int flag with specified name, default value, and usage string.
The return value is the address of an int variable that stores the value of the flag.




### <a id="FlagSet.Int64">func</a> (\*FlagSet) [Int64](https://golang.org/src/flag/flag.go?s=21208:21278#L674)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Int64(name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#int64">int64</a>, usage <a href="/pkg/builtin/#string">string</a>) *<a href="/pkg/builtin/#int64">int64</a></pre>
Int64 defines an int64 flag with specified name, default value, and usage string.
The return value is the address of an int64 variable that stores the value of the flag.




### <a id="FlagSet.Int64Var">func</a> (\*FlagSet) [Int64Var](https://golang.org/src/flag/flag.go?s=20606:20682#L662)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Int64Var(p *<a href="/pkg/builtin/#int64">int64</a>, name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#int64">int64</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
Int64Var defines an int64 flag with specified name, default value, and usage string.
The argument p points to an int64 variable in which to store the value of the flag.




### <a id="FlagSet.IntVar">func</a> (\*FlagSet) [IntVar](https://golang.org/src/flag/flag.go?s=19458:19528#L636)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) IntVar(p *<a href="/pkg/builtin/#int">int</a>, name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#int">int</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
IntVar defines an int flag with specified name, default value, and usage string.
The argument p points to an int variable in which to store the value of the flag.




### <a id="FlagSet.Lookup">func</a> (\*FlagSet) [Lookup](https://golang.org/src/flag/flag.go?s=11158:11201#L400)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Lookup(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Flag">Flag</a></pre>
Lookup returns the Flag structure of the named flag, returning nil if none exists.




### <a id="FlagSet.NArg">func</a> (\*FlagSet) [NArg](https://golang.org/src/flag/flag.go?s=17749:17777#L597)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) NArg() <a href="/pkg/builtin/#int">int</a></pre>
NArg is the number of arguments remaining after flags have been processed.




### <a id="FlagSet.NFlag">func</a> (\*FlagSet) [NFlag](https://golang.org/src/flag/flag.go?s=16956:16985#L574)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) NFlag() <a href="/pkg/builtin/#int">int</a></pre>
NFlag returns the number of flags that have been set.




### <a id="FlagSet.Name">func</a> (\*FlagSet) [Name](https://golang.org/src/flag/flag.go?s=9862:9893#L356)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the name of the flag set.




### <a id="FlagSet.Output">func</a> (\*FlagSet) [Output](https://golang.org/src/flag/flag.go?s=9717:9753#L348)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Output() <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a></pre>
Output returns the destination for usage and error messages. os.Stderr is returned if
output was not set or was set to nil.




### <a id="FlagSet.Parse">func</a> (\*FlagSet) [Parse](https://golang.org/src/flag/flag.go?s=32361:32410#L954)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Parse(arguments []<a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Parse parses flag definitions from the argument list, which should not
include the command name. Must be called after all flags in the FlagSet
are defined and before flags are accessed by the program.
The return value will be ErrHelp if -help or -h were set but not defined.




### <a id="FlagSet.Parsed">func</a> (\*FlagSet) [Parsed](https://golang.org/src/flag/flag.go?s=32751:32782#L978)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Parsed() <a href="/pkg/builtin/#bool">bool</a></pre>
Parsed reports whether f.Parse has been called.




### <a id="FlagSet.PrintDefaults">func</a> (\*FlagSet) [PrintDefaults](https://golang.org/src/flag/flag.go?s=13746:13779#L490)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) PrintDefaults()</pre>
PrintDefaults prints, to standard error unless configured otherwise, the
default values of all defined command-line flags in the set. See the
documentation for the global function PrintDefaults for more information.




### <a id="FlagSet.Set">func</a> (\*FlagSet) [Set](https://golang.org/src/flag/flag.go?s=11442:11489#L411)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Set(name, value <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
Set sets the value of the named flag.




### <a id="FlagSet.SetOutput">func</a> (\*FlagSet) [SetOutput](https://golang.org/src/flag/flag.go?s=10165:10210#L367)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) SetOutput(output <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>)</pre>
SetOutput sets the destination for usage and error messages.
If output is nil, os.Stderr is used.




### <a id="FlagSet.String">func</a> (\*FlagSet) [String](https://golang.org/src/flag/flag.go?s=24803:24876#L752)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) String(name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#string">string</a>, usage <a href="/pkg/builtin/#string">string</a>) *<a href="/pkg/builtin/#string">string</a></pre>
String defines a string flag with specified name, default value, and usage string.
The return value is the address of a string variable that stores the value of the flag.




### <a id="FlagSet.StringVar">func</a> (\*FlagSet) [StringVar](https://golang.org/src/flag/flag.go?s=24191:24270#L740)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) StringVar(p *<a href="/pkg/builtin/#string">string</a>, name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#string">string</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
StringVar defines a string flag with specified name, default value, and usage string.
The argument p points to a string variable in which to store the value of the flag.




### <a id="FlagSet.Uint">func</a> (\*FlagSet) [Uint](https://golang.org/src/flag/flag.go?s=22385:22452#L700)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Uint(name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#uint">uint</a>, usage <a href="/pkg/builtin/#string">string</a>) *<a href="/pkg/builtin/#uint">uint</a></pre>
Uint defines a uint flag with specified name, default value, and usage string.
The return value is the address of a uint variable that stores the value of the flag.




### <a id="FlagSet.Uint64">func</a> (\*FlagSet) [Uint64](https://golang.org/src/flag/flag.go?s=23582:23655#L726)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Uint64(name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#uint64">uint64</a>, usage <a href="/pkg/builtin/#string">string</a>) *<a href="/pkg/builtin/#uint64">uint64</a></pre>
Uint64 defines a uint64 flag with specified name, default value, and usage string.
The return value is the address of a uint64 variable that stores the value of the flag.




### <a id="FlagSet.Uint64Var">func</a> (\*FlagSet) [Uint64Var](https://golang.org/src/flag/flag.go?s=22970:23049#L714)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Uint64Var(p *<a href="/pkg/builtin/#uint64">uint64</a>, name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#uint64">uint64</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
Uint64Var defines a uint64 flag with specified name, default value, and usage string.
The argument p points to a uint64 variable in which to store the value of the flag.




### <a id="FlagSet.UintVar">func</a> (\*FlagSet) [UintVar](https://golang.org/src/flag/flag.go?s=21801:21874#L688)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) UintVar(p *<a href="/pkg/builtin/#uint">uint</a>, name <a href="/pkg/builtin/#string">string</a>, value <a href="/pkg/builtin/#uint">uint</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
UintVar defines a uint flag with specified name, default value, and usage string.
The argument p points to a uint variable in which to store the value of the flag.




### <a id="FlagSet.Var">func</a> (\*FlagSet) [Var](https://golang.org/src/flag/flag.go?s=28549:28610#L826)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Var(value <a href="#Value">Value</a>, name <a href="/pkg/builtin/#string">string</a>, usage <a href="/pkg/builtin/#string">string</a>)</pre>
Var defines a flag with the specified name and usage string. The type and
value of the flag are represented by the first argument, of type Value, which
typically holds a user-defined implementation of Value. For instance, the
caller could create a flag that turns a comma-separated string into a slice
of strings by giving the slice the methods of Value; in particular, Set would
decompose the comma-separated string into the slice.




### <a id="FlagSet.Visit">func</a> (\*FlagSet) [Visit](https://golang.org/src/flag/flag.go?s=10778:10817#L387)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) Visit(fn func(*<a href="#Flag">Flag</a>))</pre>
Visit visits the flags in lexicographical order, calling fn for each.
It visits only those flags that have been set.




### <a id="FlagSet.VisitAll">func</a> (\*FlagSet) [VisitAll](https://golang.org/src/flag/flag.go?s=10355:10397#L373)
<pre>func (f *<a href="#FlagSet">FlagSet</a>) VisitAll(fn func(*<a href="#Flag">Flag</a>))</pre>
VisitAll visits the flags in lexicographical order, calling fn for each.
It visits all flags, even those not set.




## <a id="Getter">type</a> [Getter](https://golang.org/src/flag/flag.go?s=7791:7842#L290)
Getter is an interface that allows the contents of a Value to be retrieved.
It wraps the Value interface, rather than being part of it, because it
appeared after Go 1 and its compatibility rules. All Value types provided
by this package satisfy the Getter interface.


<pre>type Getter interface {
    <a href="#Value">Value</a>
    Get() interface{}
}</pre>











## <a id="Value">type</a> [Value](https://golang.org/src/flag/flag.go?s=7450:7510#L281)
Value is the interface to the dynamic value stored in a flag.
(The default value is represented as a string.)

If a Value has an IsBoolFlag() bool method returning true,
the command-line parser makes -name equivalent to -name=true
rather than using the next command-line argument.

Set is called once, in command line order, for each flag present.
The flag package may call the String method with a zero-valued receiver,
such as a nil pointer.


<pre>type Value interface {
    String() <a href="/pkg/builtin/#string">string</a>
    Set(<a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a>
}</pre>





<a id="example_Value">Example</a>
```go
```

output:
```txt
```









