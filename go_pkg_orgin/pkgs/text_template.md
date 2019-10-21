

# template
`import "text/template"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a id="pkg-overview">Overview</a>
Package template implements data-driven templates for generating textual output.

To generate HTML output, see package html/template, which has the same interface
as this package but automatically secures HTML output against certain attacks.

Templates are executed by applying them to a data structure. Annotations in the
template refer to elements of the data structure (typically a field of a struct
or a key in a map) to control execution and derive values to be displayed.
Execution of the template walks the structure and sets the cursor, represented
by a period '.' and called "dot", to the value at the current location in the
structure as execution proceeds.

The input text for a template is UTF-8-encoded text in any format.
"Actions"--data evaluations or control structures--are delimited by
"{{" and "}}"; all text outside actions is copied to the output unchanged.
Except for raw strings, actions may not span newlines, although comments can.

Once parsed, a template may be executed safely in parallel, although if parallel
executions share a Writer the output may be interleaved.

Here is a trivial example that prints "17 items are made of wool".


	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil { panic(err) }

More intricate examples appear below.

### Text and spaces
By default, all text between actions is copied verbatim when the template is
executed. For example, the string " items are made of " in the example above appears
on standard output when the program is run.

However, to aid in formatting template source code, if an action's left delimiter
(by default "{{") is followed immediately by a minus sign and ASCII space character
("{{- "), all trailing white space is trimmed from the immediately preceding text.
Similarly, if the right delimiter ("}}") is preceded by a space and minus sign
(" -}}"), all leading white space is trimmed from the immediately following text.
In these trim markers, the ASCII space must be present; "{{-3}}" parses as an
action containing the number -3.

For instance, when executing the template whose source is


	"{{23 -}} < {{- 45}}"

the generated output would be


	"23<45"

For this trimming, the definition of white space characters is the same as in Go:
space, horizontal tab, carriage return, and newline.

### Actions
Here is the list of actions. "Arguments" and "pipelines" are evaluations of
data, defined in detail in the corresponding sections that follow.


	{{/* a comment */}}
	{{- /* a comment with white space trimmed from preceding and following text */ -}}
		A comment; discarded. May contain newlines.
		Comments do not nest and must start and end at the
		delimiters, as shown here.
	
	{{pipeline}}
		The default textual representation (the same as would be
		printed by fmt.Print) of the value of the pipeline is copied
		to the output.
	
	{{if pipeline}} T1 {{end}}
		If the value of the pipeline is empty, no output is generated;
		otherwise, T1 is executed. The empty values are false, 0, any
		nil pointer or interface value, and any array, slice, map, or
		string of length zero.
		Dot is unaffected.
	
	{{if pipeline}} T1 {{else}} T0 {{end}}
		If the value of the pipeline is empty, T0 is executed;
		otherwise, T1 is executed. Dot is unaffected.
	
	{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
		To simplify the appearance of if-else chains, the else action
		of an if may include another if directly; the effect is exactly
		the same as writing
			{{if pipeline}} T1 {{else}}{{if pipeline}} T0 {{end}}{{end}}
	
	{{range pipeline}} T1 {{end}}
		The value of the pipeline must be an array, slice, map, or channel.
		If the value of the pipeline has length zero, nothing is output;
		otherwise, dot is set to the successive elements of the array,
		slice, or map and T1 is executed. If the value is a map and the
		keys are of basic type with a defined order ("comparable"), the
		elements will be visited in sorted key order.
	
	{{range pipeline}} T1 {{else}} T0 {{end}}
		The value of the pipeline must be an array, slice, map, or channel.
		If the value of the pipeline has length zero, dot is unaffected and
		T0 is executed; otherwise, dot is set to the successive elements
		of the array, slice, or map and T1 is executed.
	
	{{template "name"}}
		The template with the specified name is executed with nil data.
	
	{{template "name" pipeline}}
		The template with the specified name is executed with dot set
		to the value of the pipeline.
	
	{{block "name" pipeline}} T1 {{end}}
		A block is shorthand for defining a template
			{{define "name"}} T1 {{end}}
		and then executing it in place
			{{template "name" pipeline}}
		The typical use is to define a set of root templates that are
		then customized by redefining the block templates within.
	
	{{with pipeline}} T1 {{end}}
		If the value of the pipeline is empty, no output is generated;
		otherwise, dot is set to the value of the pipeline and T1 is
		executed.
	
	{{with pipeline}} T1 {{else}} T0 {{end}}
		If the value of the pipeline is empty, dot is unaffected and T0
		is executed; otherwise, dot is set to the value of the pipeline
		and T1 is executed.

### Arguments
An argument is a simple value, denoted by one of the following.


	- A boolean, string, character, integer, floating-point, imaginary
	  or complex constant in Go syntax. These behave like Go's untyped
	  constants. Note that, as in Go, whether a large integer constant
	  overflows when assigned or passed to a function can depend on whether
	  the host machine's ints are 32 or 64 bits.
	- The keyword nil, representing an untyped Go nil.
	- The character '.' (period):
		.
	  The result is the value of dot.
	- A variable name, which is a (possibly empty) alphanumeric string
	  preceded by a dollar sign, such as
		$piOver2
	  or
		$
	  The result is the value of the variable.
	  Variables are described below.
	- The name of a field of the data, which must be a struct, preceded
	  by a period, such as
		.Field
	  The result is the value of the field. Field invocations may be
	  chained:
	    .Field1.Field2
	  Fields can also be evaluated on variables, including chaining:
	    $x.Field1.Field2
	- The name of a key of the data, which must be a map, preceded
	  by a period, such as
		.Key
	  The result is the map element value indexed by the key.
	  Key invocations may be chained and combined with fields to any
	  depth:
	    .Field1.Key1.Field2.Key2
	  Although the key must be an alphanumeric identifier, unlike with
	  field names they do not need to start with an upper case letter.
	  Keys can also be evaluated on variables, including chaining:
	    $x.key1.key2
	- The name of a niladic method of the data, preceded by a period,
	  such as
		.Method
	  The result is the value of invoking the method with dot as the
	  receiver, dot.Method(). Such a method must have one return value (of
	  any type) or two return values, the second of which is an error.
	  If it has two and the returned error is non-nil, execution terminates
	  and an error is returned to the caller as the value of Execute.
	  Method invocations may be chained and combined with fields and keys
	  to any depth:
	    .Field1.Key1.Method1.Field2.Key2.Method2
	  Methods can also be evaluated on variables, including chaining:
	    $x.Method1.Field
	- The name of a niladic function, such as
		fun
	  The result is the value of invoking the function, fun(). The return
	  types and values behave as in methods. Functions and function
	  names are described below.
	- A parenthesized instance of one the above, for grouping. The result
	  may be accessed by a field or map key invocation.
		print (.F1 arg1) (.F2 arg2)
		(.StructValuedMethod "arg").Field

Arguments may evaluate to any type; if they are pointers the implementation
automatically indirects to the base type when required.
If an evaluation yields a function value, such as a function-valued
field of a struct, the function is not invoked automatically, but it
can be used as a truth value for an if action and the like. To invoke
it, use the call function, defined below.

### Pipelines
A pipeline is a possibly chained sequence of "commands". A command is a simple
value (argument) or a function or method call, possibly with multiple arguments:


	Argument
		The result is the value of evaluating the argument.
	.Method [Argument...]
		The method can be alone or the last element of a chain but,
		unlike methods in the middle of a chain, it can take arguments.
		The result is the value of calling the method with the
		arguments:
			dot.Method(Argument1, etc.)
	functionName [Argument...]
		The result is the value of calling the function associated
		with the name:
			function(Argument1, etc.)
		Functions and function names are described below.

A pipeline may be "chained" by separating a sequence of commands with pipeline
characters '|'. In a chained pipeline, the result of each command is
passed as the last argument of the following command. The output of the final
command in the pipeline is the value of the pipeline.

The output of a command will be either one value or two values, the second of
which has type error. If that second value is present and evaluates to
non-nil, execution terminates and the error is returned to the caller of
Execute.

### Variables
A pipeline inside an action may initialize a variable to capture the result.
The initialization has syntax


	$variable := pipeline

where $variable is the name of the variable. An action that declares a
variable produces no output.

Variables previously declared can also be assigned, using the syntax


	$variable = pipeline

If a "range" action initializes a variable, the variable is set to the
successive elements of the iteration. Also, a "range" may declare two
variables, separated by a comma:


	range $index, $element := pipeline

in which case $index and $element are set to the successive values of the
array/slice index or map key and element, respectively. Note that if there is
only one variable, it is assigned the element; this is opposite to the
convention in Go range clauses.

A variable's scope extends to the "end" action of the control structure ("if",
"with", or "range") in which it is declared, or to the end of the template if
there is no such control structure. A template invocation does not inherit
variables from the point of its invocation.

When execution begins, $ is set to the data argument passed to Execute, that is,
to the starting value of dot.

### Examples
Here are some example one-line templates demonstrating pipelines and variables.
All produce the quoted word "output":


	{{"\"output\""}}
		A string constant.
	{{`"output"`}}
		A raw string constant.
	{{printf "%q" "output"}}
		A function call.
	{{"output" | printf "%q"}}
		A function call whose final argument comes from the previous
		command.
	{{printf "%q" (print "out" "put")}}
		A parenthesized argument.
	{{"put" | printf "%s%s" "out" | printf "%q"}}
		A more elaborate call.
	{{"output" | printf "%s" | printf "%q"}}
		A longer chain.
	{{with "output"}}{{printf "%q" .}}{{end}}
		A with action using dot.
	{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
		A with action that creates and uses a variable.
	{{with $x := "output"}}{{printf "%q" $x}}{{end}}
		A with action that uses the variable in another action.
	{{with $x := "output"}}{{$x | printf "%q"}}{{end}}
		The same, but pipelined.

### Functions
During execution functions are found in two function maps: first in the
template, then in the global function map. By default, no functions are defined
in the template but the Funcs method can be used to add them.

Predefined global functions are named as follows.


	and
		Returns the boolean AND of its arguments by returning the
		first empty argument or the last argument, that is,
		"and x y" behaves as "if x then y else x". All the
		arguments are evaluated.
	call
		Returns the result of calling the first argument, which
		must be a function, with the remaining arguments as parameters.
		Thus "call .X.Y 1 2" is, in Go notation, dot.X.Y(1, 2) where
		Y is a func-valued field, map entry, or the like.
		The first argument must be the result of an evaluation
		that yields a value of function type (as distinct from
		a predefined function such as print). The function must
		return either one or two result values, the second of which
		is of type error. If the arguments don't match the function
		or the returned error value is non-nil, execution stops.
	html
		Returns the escaped HTML equivalent of the textual
		representation of its arguments. This function is unavailable
		in html/template, with a few exceptions.
	index
		Returns the result of indexing its first argument by the
		following arguments. Thus "index x 1 2 3" is, in Go syntax,
		x[1][2][3]. Each indexed item must be a map, slice, or array.
	slice
		slice returns the result of slicing its first argument by the
		remaining arguments. Thus "slice x 1 2" is, in Go syntax, x[1:2],
		while "slice x" is x[:], "slice x 1" is x[1:], and "slice x 1 2 3"
		is x[1:2:3]. The first argument must be a string, slice, or array.
	js
		Returns the escaped JavaScript equivalent of the textual
		representation of its arguments.
	len
		Returns the integer length of its argument.
	not
		Returns the boolean negation of its single argument.
	or
		Returns the boolean OR of its arguments by returning the
		first non-empty argument or the last argument, that is,
		"or x y" behaves as "if x then x else y". All the
		arguments are evaluated.
	print
		An alias for fmt.Sprint
	printf
		An alias for fmt.Sprintf
	println
		An alias for fmt.Sprintln
	urlquery
		Returns the escaped value of the textual representation of
		its arguments in a form suitable for embedding in a URL query.
		This function is unavailable in html/template, with a few
		exceptions.

The boolean functions take any zero value to be false and a non-zero
value to be true.

There is also a set of binary comparison operators defined as
functions:


	eq
		Returns the boolean truth of arg1 == arg2
	ne
		Returns the boolean truth of arg1 != arg2
	lt
		Returns the boolean truth of arg1 < arg2
	le
		Returns the boolean truth of arg1 <= arg2
	gt
		Returns the boolean truth of arg1 > arg2
	ge
		Returns the boolean truth of arg1 >= arg2

For simpler multi-way equality tests, eq (only) accepts two or more
arguments and compares the second and subsequent to the first,
returning in effect


	arg1==arg2 || arg1==arg3 || arg1==arg4 ...

(Unlike with || in Go, however, eq is a function call and all the
arguments will be evaluated.)

The comparison functions work on basic types only (or named basic
types, such as "type Celsius float32"). They implement the Go rules
for comparison of values, except that size and exact type are
ignored, so any integer value, signed or unsigned, may be compared
with any other integer value. (The arithmetic value is compared,
not the bit pattern, so all negative integers are less than all
unsigned integers.) However, as usual, one may not compare an int
with a float32 and so on.

### Associated templates
Each template is named by a string specified when it is created. Also, each
template is associated with zero or more other templates that it may invoke by
name; such associations are transitive and form a name space of templates.

A template may use a template invocation to instantiate another associated
template; see the explanation of the "template" action above. The name must be
that of a template associated with the template that contains the invocation.

### Nested template definitions
When parsing a template, another template may be defined and associated with the
template being parsed. Template definitions must appear at the top level of the
template, much like global variables in a Go program.

The syntax of such definitions is to surround each template declaration with a
"define" and "end" action.

The define action names the template being created by providing a string
constant. Here is a simple example:


	`{{define "T1"}}ONE{{end}}
	{{define "T2"}}TWO{{end}}
	{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}
	{{template "T3"}}`

This defines two templates, T1 and T2, and a third T3 that invokes the other two
when it is executed. Finally it invokes T3. If executed this template will
produce the text


	ONE TWO

By construction, a template may reside in only one association. If it's
necessary to have a template addressable from multiple associations, the
template definition must be parsed multiple times to create distinct *Template
values, or must be copied with the Clone or AddParseTree method.

Parse may be called multiple times to assemble the various associated templates;
see the ParseFiles and ParseGlob functions and methods for simple ways to parse
related templates stored in files.

A template may be executed directly or through ExecuteTemplate, which executes
an associated template identified by name. To invoke our example above, we
might write,


	err := tmpl.Execute(os.Stdout, "no data needed")
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}

or to invoke a particular template explicitly by name,


	err := tmpl.ExecuteTemplate(os.Stdout, "T2", "no data needed")
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}




## <a id="pkg-index">Index</a>
* [func HTMLEscape(w io.Writer, b []byte)](#HTMLEscape)
* [func HTMLEscapeString(s string) string](#HTMLEscapeString)
* [func HTMLEscaper(args ...interface{}) string](#HTMLEscaper)
* [func IsTrue(val interface{}) (truth, ok bool)](#IsTrue)
* [func JSEscape(w io.Writer, b []byte)](#JSEscape)
* [func JSEscapeString(s string) string](#JSEscapeString)
* [func JSEscaper(args ...interface{}) string](#JSEscaper)
* [func URLQueryEscaper(args ...interface{}) string](#URLQueryEscaper)
* [type ExecError](#ExecError)
  * [func (e ExecError) Error() string](#ExecError.Error)
  * [func (e ExecError) Unwrap() error](#ExecError.Unwrap)
* [type FuncMap](#FuncMap)
* [type Template](#Template)
  * [func Must(t *Template, err error) *Template](#Must)
  * [func New(name string) *Template](#New)
  * [func ParseFiles(filenames ...string) (*Template, error)](#ParseFiles)
  * [func ParseGlob(pattern string) (*Template, error)](#ParseGlob)
  * [func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error)](#Template.AddParseTree)
  * [func (t *Template) Clone() (*Template, error)](#Template.Clone)
  * [func (t *Template) DefinedTemplates() string](#Template.DefinedTemplates)
  * [func (t *Template) Delims(left, right string) *Template](#Template.Delims)
  * [func (t *Template) Execute(wr io.Writer, data interface{}) error](#Template.Execute)
  * [func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error](#Template.ExecuteTemplate)
  * [func (t *Template) Funcs(funcMap FuncMap) *Template](#Template.Funcs)
  * [func (t *Template) Lookup(name string) *Template](#Template.Lookup)
  * [func (t *Template) Name() string](#Template.Name)
  * [func (t *Template) New(name string) *Template](#Template.New)
  * [func (t *Template) Option(opt ...string) *Template](#Template.Option)
  * [func (t *Template) Parse(text string) (*Template, error)](#Template.Parse)
  * [func (t *Template) ParseFiles(filenames ...string) (*Template, error)](#Template.ParseFiles)
  * [func (t *Template) ParseGlob(pattern string) (*Template, error)](#Template.ParseGlob)
  * [func (t *Template) Templates() []*Template](#Template.Templates)


#### <a id="pkg-examples">Examples</a>
* [Template](#example_Template)
* [Template (Block)](#example_Template_block)
* [Template (Func)](#example_Template_func)
* [Template (Glob)](#example_Template_glob)
* [Template (Helpers)](#example_Template_helpers)
* [Template (Share)](#example_Template_share)


#### <a id="pkg-files">Package files</a>
[doc.go](https://golang.org/src/text/template/doc.go) [exec.go](https://golang.org/src/text/template/exec.go) [funcs.go](https://golang.org/src/text/template/funcs.go) [helper.go](https://golang.org/src/text/template/helper.go) [option.go](https://golang.org/src/text/template/option.go) [template.go](https://golang.org/src/text/template/template.go) 






## <a id="HTMLEscape">func</a> [HTMLEscape](https://golang.org/src/text/template/funcs.go?s=15725:15763#L574)
<pre>func HTMLEscape(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, b []<a href="/pkg/builtin/#byte">byte</a>)</pre>
HTMLEscape writes to w the escaped HTML equivalent of the plain text data b.



## <a id="HTMLEscapeString">func</a> [HTMLEscapeString](https://golang.org/src/text/template/funcs.go?s=16202:16240#L602)
<pre>func HTMLEscapeString(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
HTMLEscapeString returns the escaped HTML equivalent of the plain text data s.



## <a id="HTMLEscaper">func</a> [HTMLEscaper](https://golang.org/src/text/template/funcs.go?s=16504:16548#L614)
<pre>func HTMLEscaper(args ...interface{}) <a href="/pkg/builtin/#string">string</a></pre>
HTMLEscaper returns the escaped HTML equivalent of the textual
representation of its arguments.



## <a id="IsTrue">func</a> [IsTrue](https://golang.org/src/text/template/exec.go?s=8705:8750#L296)
<pre>func IsTrue(val interface{}) (truth, ok <a href="/pkg/builtin/#bool">bool</a>)</pre>
IsTrue reports whether the value is 'true', in the sense of not the zero of its type,
and whether the value has a meaningful truth value. This is the definition of
truth used by if and other such actions.



## <a id="JSEscape">func</a> [JSEscape](https://golang.org/src/text/template/funcs.go?s=16924:16960#L632)
<pre>func JSEscape(w <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, b []<a href="/pkg/builtin/#byte">byte</a>)</pre>
JSEscape writes to w the escaped JavaScript equivalent of the plain text data b.



## <a id="JSEscapeString">func</a> [JSEscapeString](https://golang.org/src/text/template/funcs.go?s=17859:17895#L679)
<pre>func JSEscapeString(s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
JSEscapeString returns the escaped JavaScript equivalent of the plain text data s.



## <a id="JSEscaper">func</a> [JSEscaper](https://golang.org/src/text/template/funcs.go?s=18296:18338#L699)
<pre>func JSEscaper(args ...interface{}) <a href="/pkg/builtin/#string">string</a></pre>
JSEscaper returns the escaped JavaScript equivalent of the textual
representation of its arguments.



## <a id="URLQueryEscaper">func</a> [URLQueryEscaper](https://golang.org/src/text/template/funcs.go?s=18527:18575#L705)
<pre>func URLQueryEscaper(args ...interface{}) <a href="/pkg/builtin/#string">string</a></pre>
URLQueryEscaper returns the escaped value of the textual representation of
its arguments in a form suitable for embedding in a URL query.





## <a id="ExecError">type</a> [ExecError](https://golang.org/src/text/template/exec.go?s=3110:3206#L105)
ExecError is the custom error type returned when Execute has an
error evaluating its template. (If a write error occurs, the actual
error is returned; it will not be of type ExecError.)


<pre>type ExecError struct {
<span id="ExecError.Name"></span>    Name <a href="/pkg/builtin/#string">string</a> <span class="comment">// Name of template.</span>
<span id="ExecError.Err"></span>    Err  <a href="/pkg/builtin/#error">error</a>  <span class="comment">// Pre-formatted error.</span>
}
</pre>











### <a id="ExecError.Error">func</a> (ExecError) [Error](https://golang.org/src/text/template/exec.go?s=3208:3241#L110)
<pre>func (e <a href="#ExecError">ExecError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="ExecError.Unwrap">func</a> (ExecError) [Unwrap](https://golang.org/src/text/template/exec.go?s=3269:3302#L114)
<pre>func (e <a href="#ExecError">ExecError</a>) Unwrap() <a href="/pkg/builtin/#error">error</a></pre>



## <a id="FuncMap">type</a> [FuncMap](https://golang.org/src/text/template/funcs.go?s=1000:1035#L20)
FuncMap is the type of the map defining the mapping from names to functions.
Each function must have either a single return value, or two return values of
which the second has type error. In that case, if the second (error)
return value evaluates to non-nil during execution, execution terminates and
Execute returns that error.

When template execution invokes a function with an argument list, that list
must be assignable to the function's parameter types. Functions meant to
apply to arguments of arbitrary type can use parameters of type interface{} or
of type reflect.Value. Similarly, functions meant to return a result of arbitrary
type can return interface{} or reflect.Value.


<pre>type FuncMap map[<a href="/pkg/builtin/#string">string</a>]interface{}</pre>











## <a id="Template">type</a> [Template](https://golang.org/src/text/template/template.go?s=859:956#L18)
Template is the representation of a parsed template. The *parse.Tree
field is exported only for use by html/template and should be treated
as unexported by all other clients.


<pre>type Template struct {
    *<a href="/pkg/text/template/parse/">parse</a>.<a href="/pkg/text/template/parse/#Tree">Tree</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>





<a id="example_Template">Example</a>
```go
```

output:
```txt
```
<a id="example_Template_block">Example (Block)</a>
```go
```

output:
```txt
```
<a id="example_Template_func">Example (Func)</a>
```go
```

output:
```txt
```
<p>This example demonstrates a custom function to process template text.
It installs the strings.Title function and uses it to
Make Title Text Look Good In Our Template&#39;s Output.
</p><a id="example_Template_glob">Example (Glob)</a>
```go
```

output:
```txt
```
<p>Here we demonstrate loading a set of templates from a directory.
</p><a id="example_Template_helpers">Example (Helpers)</a>
```go
```

output:
```txt
```
<p>This example demonstrates one way to share some templates
and use them in different contexts. In this variant we add multiple driver
templates by hand to an existing bundle of templates.
</p><a id="example_Template_share">Example (Share)</a>
```go
```

output:
```txt
```
<p>This example demonstrates how to use one group of driver
templates with distinct sets of helper templates.
</p>



### <a id="Must">func</a> [Must](https://golang.org/src/text/template/helper.go?s=576:619#L11)
<pre>func Must(t *<a href="#Template">Template</a>, err <a href="/pkg/builtin/#error">error</a>) *<a href="#Template">Template</a></pre>
Must is a helper that wraps a call to a function returning (*Template, error)
and panics if the error is non-nil. It is intended for use in variable
initializations such as


	var t = template.Must(template.New("name").Parse("text"))




### <a id="New">func</a> [New](https://golang.org/src/text/template/template.go?s=1022:1053#L27)
<pre>func New(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Template">Template</a></pre>
New allocates a new, undefined template with the given name.




### <a id="ParseFiles">func</a> [ParseFiles](https://golang.org/src/text/template/helper.go?s=1224:1279#L27)
<pre>func ParseFiles(filenames ...<a href="/pkg/builtin/#string">string</a>) (*<a href="#Template">Template</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseFiles creates a new Template and parses the template definitions from
the named files. The returned template's name will have the base name and
parsed contents of the first file. There must be at least one file.
If an error occurs, parsing stops and the returned *Template is nil.

When parsing multiple files with the same name in different directories,
the last one mentioned will be the one that results.
For instance, ParseFiles("a/foo", "b/foo") stores "b/foo" as the template
named "foo", while "a/foo" is unavailable.




### <a id="ParseGlob">func</a> [ParseGlob](https://golang.org/src/text/template/helper.go?s=3809:3858#L93)
<pre>func ParseGlob(pattern <a href="/pkg/builtin/#string">string</a>) (*<a href="#Template">Template</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseGlob creates a new Template and parses the template definitions from
the files identified by the pattern. The files are matched according to the
semantics of filepath.Match, and the pattern must match at least one file.
The returned template will have the (base) name and (parsed) contents of the
first file matched by the pattern. ParseGlob is equivalent to calling
ParseFiles with the list of files matched by the pattern.

When parsing multiple files with the same name in different directories,
the last one mentioned will be the one that results.






### <a id="Template.AddParseTree">func</a> (\*Template) [AddParseTree](https://golang.org/src/text/template/template.go?s=3446:3527#L114)
<pre>func (t *<a href="#Template">Template</a>) AddParseTree(name <a href="/pkg/builtin/#string">string</a>, tree *<a href="/pkg/text/template/parse/">parse</a>.<a href="/pkg/text/template/parse/#Tree">Tree</a>) (*<a href="#Template">Template</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
AddParseTree adds parse tree for template with given name and associates it with t.
If the template does not already exist, it will create a new one.
If the template does exist, it will be replaced.




### <a id="Template.Clone">func</a> (\*Template) [Clone](https://golang.org/src/text/template/template.go?s=2497:2542#L75)
<pre>func (t *<a href="#Template">Template</a>) Clone() (*<a href="#Template">Template</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Clone returns a duplicate of the template, including all associated
templates. The actual representation is not copied, but the name space of
associated templates is, so further calls to Parse in the copy will add
templates to the copy but not to the original. Clone can be used to prepare
common templates and use them with variant definitions for other templates
by adding the variants after the clone is made.




### <a id="Template.DefinedTemplates">func</a> (\*Template) [DefinedTemplates](https://golang.org/src/text/template/exec.go?s=6556:6600#L219)
<pre>func (t *<a href="#Template">Template</a>) DefinedTemplates() <a href="/pkg/builtin/#string">string</a></pre>
DefinedTemplates returns a string listing the defined templates,
prefixed by the string "; defined templates are: ". If there are none,
it returns the empty string. For generating an error message here
and in html/template.




### <a id="Template.Delims">func</a> (\*Template) [Delims](https://golang.org/src/text/template/template.go?s=4440:4495#L146)
<pre>func (t *<a href="#Template">Template</a>) Delims(left, right <a href="/pkg/builtin/#string">string</a>) *<a href="#Template">Template</a></pre>
Delims sets the action delimiters to the specified strings, to be used in
subsequent calls to Parse, ParseFiles, or ParseGlob. Nested template
definitions will inherit the settings. An empty delimiter stands for the
corresponding default: {{ or }}.
The return value is the template, so calls can be chained.




### <a id="Template.Execute">func</a> (\*Template) [Execute](https://golang.org/src/text/template/exec.go?s=5823:5887#L193)
<pre>func (t *<a href="#Template">Template</a>) Execute(wr <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, data interface{}) <a href="/pkg/builtin/#error">error</a></pre>
Execute applies a parsed template to the specified data object,
and writes the output to wr.
If an error occurs executing the template or writing its output,
execution stops, but partial results may already have been written to
the output writer.
A template may be executed safely in parallel, although if parallel
executions share a Writer the output may be interleaved.

If data is a reflect.Value, the template applies to the concrete
value that the reflect.Value holds, as in fmt.Print.




### <a id="Template.ExecuteTemplate">func</a> (\*Template) [ExecuteTemplate](https://golang.org/src/text/template/exec.go?s=5003:5088#L172)
<pre>func (t *<a href="#Template">Template</a>) ExecuteTemplate(wr <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Writer">Writer</a>, name <a href="/pkg/builtin/#string">string</a>, data interface{}) <a href="/pkg/builtin/#error">error</a></pre>
ExecuteTemplate applies the template associated with t that has the given name
to the specified data object and writes the output to wr.
If an error occurs executing the template or writing its output,
execution stops, but partial results may already have been written to
the output writer.
A template may be executed safely in parallel, although if parallel
executions share a Writer the output may be interleaved.




### <a id="Template.Funcs">func</a> (\*Template) [Funcs](https://golang.org/src/text/template/template.go?s=4963:5014#L159)
<pre>func (t *<a href="#Template">Template</a>) Funcs(funcMap <a href="#FuncMap">FuncMap</a>) *<a href="#Template">Template</a></pre>
Funcs adds the elements of the argument map to the template's function map.
It must be called before the template is parsed.
It panics if a value in the map is not a function with appropriate return
type or if the name cannot be used syntactically as a function in a template.
It is legal to overwrite elements of the map. The return value is the template,
so calls can be chained.




### <a id="Template.Lookup">func</a> (\*Template) [Lookup](https://golang.org/src/text/template/template.go?s=5314:5362#L170)
<pre>func (t *<a href="#Template">Template</a>) Lookup(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Template">Template</a></pre>
Lookup returns the template with the given name that is associated with t.
It returns nil if there is no such template or the template has no definition.




### <a id="Template.Name">func</a> (\*Template) [Name](https://golang.org/src/text/template/template.go?s=1155:1187#L36)
<pre>func (t *<a href="#Template">Template</a>) Name() <a href="/pkg/builtin/#string">string</a></pre>
Name returns the name of the template.




### <a id="Template.New">func</a> (\*Template) [New](https://golang.org/src/text/template/template.go?s=1612:1657#L47)
<pre>func (t *<a href="#Template">Template</a>) New(name <a href="/pkg/builtin/#string">string</a>) *<a href="#Template">Template</a></pre>
New allocates a new, undefined template associated with the given one and with the same
delimiters. The association, which is transitive, allows one template to
invoke another with a {{template}} action.

Because associated templates share underlying data, template construction
cannot be done safely in parallel. Once the templates are constructed, they
can be executed in parallel.




### <a id="Template.Option">func</a> (\*Template) [Option](https://golang.org/src/text/template/option.go?s=1400:1450#L32)
<pre>func (t *<a href="#Template">Template</a>) Option(opt ...<a href="/pkg/builtin/#string">string</a>) *<a href="#Template">Template</a></pre>
Option sets options for the template. Options are described by
strings, either a simple string or "key=value". There can be at
most one equals sign in an option string. If the option string
is unrecognized or otherwise invalid, Option panics.

Known options:

missingkey: Control the behavior during execution if a map is
indexed with a key that is not present in the map.


	"missingkey=default" or "missingkey=invalid"
		The default behavior: Do nothing and continue execution.
		If printed, the result of the index operation is the string
		"<no value>".
	"missingkey=zero"
		The operation returns the zero value for the map type's element.
	"missingkey=error"
		Execution stops immediately with an error.




### <a id="Template.Parse">func</a> (\*Template) [Parse](https://golang.org/src/text/template/template.go?s=5984:6040#L187)
<pre>func (t *<a href="#Template">Template</a>) Parse(text <a href="/pkg/builtin/#string">string</a>) (*<a href="#Template">Template</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Parse parses text as a template body for t.
Named template definitions ({{define ...}} or {{block ...}} statements) in text
define additional templates associated with t and are removed from the
definition of t itself.

Templates can be redefined in successive calls to Parse.
A template definition with a body containing only white space and comments
is considered empty and will not replace an existing template's body.
This allows using Parse to add new named template definitions without
overwriting the main template body.




### <a id="Template.ParseFiles">func</a> (\*Template) [ParseFiles](https://golang.org/src/text/template/helper.go?s=2001:2070#L42)
<pre>func (t *<a href="#Template">Template</a>) ParseFiles(filenames ...<a href="/pkg/builtin/#string">string</a>) (*<a href="#Template">Template</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseFiles parses the named files and associates the resulting templates with
t. If an error occurs, parsing stops and the returned template is nil;
otherwise it is t. There must be at least one file.
Since the templates created by ParseFiles are named by the base
names of the argument files, t should usually have the name of one
of the (base) names of the files. If it does not, depending on t's
contents before calling ParseFiles, t.Execute may fail. In that
case use t.ExecuteTemplate to execute a valid template.

When parsing multiple files with the same name in different directories,
the last one mentioned will be the one that results.




### <a id="Template.ParseGlob">func</a> (\*Template) [ParseGlob](https://golang.org/src/text/template/helper.go?s=4383:4446#L105)
<pre>func (t *<a href="#Template">Template</a>) ParseGlob(pattern <a href="/pkg/builtin/#string">string</a>) (*<a href="#Template">Template</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
ParseGlob parses the template definitions in the files identified by the
pattern and associates the resulting templates with t. The files are matched
according to the semantics of filepath.Match, and the pattern must match at
least one file. ParseGlob is equivalent to calling t.ParseFiles with the
list of files matched by the pattern.

When parsing multiple files with the same name in different directories,
the last one mentioned will be the one that results.




### <a id="Template.Templates">func</a> (\*Template) [Templates](https://golang.org/src/text/template/template.go?s=3884:3926#L129)
<pre>func (t *<a href="#Template">Template</a>) Templates() []*<a href="#Template">Template</a></pre>
Templates returns a slice of defined templates associated with t.







