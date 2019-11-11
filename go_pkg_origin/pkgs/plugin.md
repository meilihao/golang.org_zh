

# plugin
`import "plugin"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package plugin implements loading and symbol resolution of Go plugins.

A plugin is a Go main package with exported functions and variables that
has been built with:


	go build -buildmode=plugin

When a plugin is first opened, the init functions of all packages not
already part of the program are called. The main function is not run.
A plugin is only initialized once, and cannot be closed.

Currently plugins are only supported on Linux and macOS.
Please report any issues.




## <a id="pkg-index">Index</a>
* [type Plugin](#Plugin)
  * [func Open(path string) (*Plugin, error)](#Open)
  * [func (p *Plugin) Lookup(symName string) (Symbol, error)](#Plugin.Lookup)
* [type Symbol](#Symbol)




#### <a id="pkg-files">Package files</a>
[plugin.go](https://golang.org/src/plugin/plugin.go) [plugin_dlopen.go](https://golang.org/src/plugin/plugin_dlopen.go) 








## <a id="Plugin">type</a> [Plugin](https://golang.org/src/plugin/plugin.go?s=720:902#L11)
Plugin is a loaded Go plugin.


<pre>type Plugin struct {
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="Open">func</a> [Open](https://golang.org/src/plugin/plugin.go?s=1065:1104#L21)
<pre>func Open(path <a href="/pkg/builtin/#string">string</a>) (*<a href="#Plugin">Plugin</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Open opens a Go plugin.
If a path has already been opened, then the existing *Plugin is returned.
It is safe for concurrent use by multiple goroutines.






### <a id="Plugin.Lookup">func</a> (\*Plugin) [Lookup](https://golang.org/src/plugin/plugin.go?s=1346:1401#L29)
<pre>func (p *<a href="#Plugin">Plugin</a>) Lookup(symName <a href="/pkg/builtin/#string">string</a>) (<a href="#Symbol">Symbol</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Lookup searches for a symbol named symName in plugin p.
A symbol is any exported variable or function.
It reports an error if the symbol is not found.
It is safe for concurrent use by multiple goroutines.




## <a id="Symbol">type</a> [Symbol](https://golang.org/src/plugin/plugin.go?s=2020:2043#L62)
A Symbol is a pointer to a variable or function.

For example, a plugin defined as


	package main
	
	import "fmt"
	
	var V int
	
	func F() { fmt.Printf("Hello, number %d\n", V) }

may be loaded with the Open function and then the exported package
symbols V and F can be accessed


	p, err := plugin.Open("plugin_name.so")
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}
	*v.(*int) = 7
	f.(func())() // prints "Hello, number 7"


<pre>type Symbol interface{}</pre>














