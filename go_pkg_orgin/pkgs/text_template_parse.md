

# parse
`import "text/template/parse"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package parse builds parse trees for templates as defined by text/template
and html/template. Clients should use those packages to construct templates
rather than this one, which provides shared internal data structures not
intended for general use.




## <a id="pkg-index">Index</a>
* [func IsEmptyTree(n Node) bool](#IsEmptyTree)
* [func Parse(name, text, leftDelim, rightDelim string, funcs ...map[string]interface{}) (map[string]*Tree, error)](#Parse)
* [type ActionNode](#ActionNode)
  * [func (a *ActionNode) Copy() Node](#ActionNode.Copy)
  * [func (a *ActionNode) String() string](#ActionNode.String)
* [type BoolNode](#BoolNode)
  * [func (b *BoolNode) Copy() Node](#BoolNode.Copy)
  * [func (b *BoolNode) String() string](#BoolNode.String)
* [type BranchNode](#BranchNode)
  * [func (b *BranchNode) Copy() Node](#BranchNode.Copy)
  * [func (b *BranchNode) String() string](#BranchNode.String)
* [type ChainNode](#ChainNode)
  * [func (c *ChainNode) Add(field string)](#ChainNode.Add)
  * [func (c *ChainNode) Copy() Node](#ChainNode.Copy)
  * [func (c *ChainNode) String() string](#ChainNode.String)
* [type CommandNode](#CommandNode)
  * [func (c *CommandNode) Copy() Node](#CommandNode.Copy)
  * [func (c *CommandNode) String() string](#CommandNode.String)
* [type DotNode](#DotNode)
  * [func (d *DotNode) Copy() Node](#DotNode.Copy)
  * [func (d *DotNode) String() string](#DotNode.String)
  * [func (d *DotNode) Type() NodeType](#DotNode.Type)
* [type FieldNode](#FieldNode)
  * [func (f *FieldNode) Copy() Node](#FieldNode.Copy)
  * [func (f *FieldNode) String() string](#FieldNode.String)
* [type IdentifierNode](#IdentifierNode)
  * [func NewIdentifier(ident string) *IdentifierNode](#NewIdentifier)
  * [func (i *IdentifierNode) Copy() Node](#IdentifierNode.Copy)
  * [func (i *IdentifierNode) SetPos(pos Pos) *IdentifierNode](#IdentifierNode.SetPos)
  * [func (i *IdentifierNode) SetTree(t *Tree) *IdentifierNode](#IdentifierNode.SetTree)
  * [func (i *IdentifierNode) String() string](#IdentifierNode.String)
* [type IfNode](#IfNode)
  * [func (i *IfNode) Copy() Node](#IfNode.Copy)
* [type ListNode](#ListNode)
  * [func (l *ListNode) Copy() Node](#ListNode.Copy)
  * [func (l *ListNode) CopyList() *ListNode](#ListNode.CopyList)
  * [func (l *ListNode) String() string](#ListNode.String)
* [type NilNode](#NilNode)
  * [func (n *NilNode) Copy() Node](#NilNode.Copy)
  * [func (n *NilNode) String() string](#NilNode.String)
  * [func (n *NilNode) Type() NodeType](#NilNode.Type)
* [type Node](#Node)
* [type NodeType](#NodeType)
  * [func (t NodeType) Type() NodeType](#NodeType.Type)
* [type NumberNode](#NumberNode)
  * [func (n *NumberNode) Copy() Node](#NumberNode.Copy)
  * [func (n *NumberNode) String() string](#NumberNode.String)
* [type PipeNode](#PipeNode)
  * [func (p *PipeNode) Copy() Node](#PipeNode.Copy)
  * [func (p *PipeNode) CopyPipe() *PipeNode](#PipeNode.CopyPipe)
  * [func (p *PipeNode) String() string](#PipeNode.String)
* [type Pos](#Pos)
  * [func (p Pos) Position() Pos](#Pos.Position)
* [type RangeNode](#RangeNode)
  * [func (r *RangeNode) Copy() Node](#RangeNode.Copy)
* [type StringNode](#StringNode)
  * [func (s *StringNode) Copy() Node](#StringNode.Copy)
  * [func (s *StringNode) String() string](#StringNode.String)
* [type TemplateNode](#TemplateNode)
  * [func (t *TemplateNode) Copy() Node](#TemplateNode.Copy)
  * [func (t *TemplateNode) String() string](#TemplateNode.String)
* [type TextNode](#TextNode)
  * [func (t *TextNode) Copy() Node](#TextNode.Copy)
  * [func (t *TextNode) String() string](#TextNode.String)
* [type Tree](#Tree)
  * [func New(name string, funcs ...map[string]interface{}) *Tree](#New)
  * [func (t *Tree) Copy() *Tree](#Tree.Copy)
  * [func (t *Tree) ErrorContext(n Node) (location, context string)](#Tree.ErrorContext)
  * [func (t *Tree) Parse(text, leftDelim, rightDelim string, treeSet map[string]*Tree, funcs ...map[string]interface{}) (tree *Tree, err error)](#Tree.Parse)
* [type VariableNode](#VariableNode)
  * [func (v *VariableNode) Copy() Node](#VariableNode.Copy)
  * [func (v *VariableNode) String() string](#VariableNode.String)
* [type WithNode](#WithNode)
  * [func (w *WithNode) Copy() Node](#WithNode.Copy)




#### <a id="pkg-files">Package files</a>
[lex.go](https://golang.org/src/text/template/parse/lex.go) [node.go](https://golang.org/src/text/template/parse/node.go) [parse.go](https://golang.org/src/text/template/parse/parse.go) 






## <a id="IsEmptyTree">func</a> [IsEmptyTree](https://golang.org/src/text/template/parse/parse.go?s=6689:6718#L239)
<pre>func IsEmptyTree(n <a href="#Node">Node</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
IsEmptyTree reports whether this tree (node) is empty of everything but space.



## <a id="Parse">func</a> [Parse](https://golang.org/src/text/template/parse/parse.go?s=1642:1753#L41)
<pre>func Parse(name, text, leftDelim, rightDelim <a href="/pkg/builtin/#string">string</a>, funcs ...map[<a href="/pkg/builtin/#string">string</a>]interface{}) (map[<a href="/pkg/builtin/#string">string</a>]*<a href="#Tree">Tree</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Parse returns a map from template name to parse.Tree, created by parsing the
templates described in the argument string. The top-level template will be
given the specified name. If an error is encountered, parsing stops and an
empty map is returned with the error.





## <a id="ActionNode">type</a> [ActionNode](https://golang.org/src/text/template/parse/node.go?s=5352:5537#L199)
ActionNode holds an action (something bounded by delimiters).
Control actions have their own nodes; ActionNode represents simple
ones such as field evaluations and parenthesized pipelines.


<pre>type ActionNode struct {
    <a href="#NodeType">NodeType</a>
    <a href="#Pos">Pos</a>

<span id="ActionNode.Line"></span>    Line <a href="/pkg/builtin/#int">int</a>       <span class="comment">// The line number in the input. Deprecated: Kept for compatibility.</span>
<span id="ActionNode.Pipe"></span>    Pipe *<a href="#PipeNode">PipeNode</a> <span class="comment">// The pipeline in the action.</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="ActionNode.Copy">func</a> (\*ActionNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=5832:5864#L220)
<pre>func (a *<a href="#ActionNode">ActionNode</a>) Copy() <a href="#Node">Node</a></pre>



### <a id="ActionNode.String">func</a> (\*ActionNode) [String](https://golang.org/src/text/template/parse/node.go?s=5699:5735#L211)
<pre>func (a *<a href="#ActionNode">ActionNode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="BoolNode">type</a> [BoolNode](https://golang.org/src/text/template/parse/node.go?s=11635:11735#L481)
BoolNode holds a boolean constant.


<pre>type BoolNode struct {
    <a href="#NodeType">NodeType</a>
    <a href="#Pos">Pos</a>

<span id="BoolNode.True"></span>    True <a href="/pkg/builtin/#bool">bool</a> <span class="comment">// The value of the boolean constant.</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="BoolNode.Copy">func</a> (\*BoolNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=12000:12030#L503)
<pre>func (b *<a href="#BoolNode">BoolNode</a>) Copy() <a href="#Node">Node</a></pre>



### <a id="BoolNode.String">func</a> (\*BoolNode) [String](https://golang.org/src/text/template/parse/node.go?s=11862:11896#L492)
<pre>func (b *<a href="#BoolNode">BoolNode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="BranchNode">type</a> [BranchNode](https://golang.org/src/text/template/parse/node.go?s=17560:17903#L720)
BranchNode is the common representation of if, range, and with.


<pre>type BranchNode struct {
    <a href="#NodeType">NodeType</a>
    <a href="#Pos">Pos</a>

<span id="BranchNode.Line"></span>    Line     <a href="/pkg/builtin/#int">int</a>       <span class="comment">// The line number in the input. Deprecated: Kept for compatibility.</span>
<span id="BranchNode.Pipe"></span>    Pipe     *<a href="#PipeNode">PipeNode</a> <span class="comment">// The pipeline to be evaluated.</span>
<span id="BranchNode.List"></span>    List     *<a href="#ListNode">ListNode</a> <span class="comment">// What to execute if the value is non-empty.</span>
<span id="BranchNode.ElseList"></span>    ElseList *<a href="#ListNode">ListNode</a> <span class="comment">// What to execute if the value is empty (nil if absent).</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="BranchNode.Copy">func</a> (\*BranchNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=18348:18380#L752)
<pre>func (b *<a href="#BranchNode">BranchNode</a>) Copy() <a href="#Node">Node</a></pre>



### <a id="BranchNode.String">func</a> (\*BranchNode) [String](https://golang.org/src/text/template/parse/node.go?s=17905:17941#L730)
<pre>func (b *<a href="#BranchNode">BranchNode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="ChainNode">type</a> [ChainNode](https://golang.org/src/text/template/parse/node.go?s=10634:10752#L437)
ChainNode holds a term followed by a chain of field accesses (identifier starting with '.').
The names may be chained ('.x.y').
The periods are dropped from each ident.


<pre>type ChainNode struct {
    <a href="#NodeType">NodeType</a>
    <a href="#Pos">Pos</a>

<span id="ChainNode.Node"></span>    Node  <a href="#Node">Node</a>
<span id="ChainNode.Field"></span>    Field []<a href="/pkg/builtin/#string">string</a> <span class="comment">// The identifiers in lexical order.</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="ChainNode.Add">func</a> (\*ChainNode) [Add](https://golang.org/src/text/template/parse/node.go?s=10971:11008#L450)
<pre>func (c *<a href="#ChainNode">ChainNode</a>) Add(field <a href="/pkg/builtin/#string">string</a>)</pre>
Add adds the named field (which should start with a period) to the end of the chain.




### <a id="ChainNode.Copy">func</a> (\*ChainNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=11445:11476#L476)
<pre>func (c *<a href="#ChainNode">ChainNode</a>) Copy() <a href="#Node">Node</a></pre>



### <a id="ChainNode.String">func</a> (\*ChainNode) [String](https://golang.org/src/text/template/parse/node.go?s=11205:11240#L461)
<pre>func (c *<a href="#ChainNode">ChainNode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="CommandNode">type</a> [CommandNode](https://golang.org/src/text/template/parse/node.go?s=6001:6131#L226)
CommandNode holds a command (a pipeline inside an evaluating action).


<pre>type CommandNode struct {
    <a href="#NodeType">NodeType</a>
    <a href="#Pos">Pos</a>

<span id="CommandNode.Args"></span>    Args []<a href="#Node">Node</a> <span class="comment">// Arguments in lexical order: Identifier, field, or constant.</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="CommandNode.Copy">func</a> (\*CommandNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=6605:6638#L260)
<pre>func (c *<a href="#CommandNode">CommandNode</a>) Copy() <a href="#Node">Node</a></pre>



### <a id="CommandNode.String">func</a> (\*CommandNode) [String](https://golang.org/src/text/template/parse/node.go?s=6321:6358#L241)
<pre>func (c *<a href="#CommandNode">CommandNode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="DotNode">type</a> [DotNode](https://golang.org/src/text/template/parse/node.go?s=8594:8642#L345)
DotNode holds the special identifier '.'.


<pre>type DotNode struct {
    <a href="#NodeType">NodeType</a>
    <a href="#Pos">Pos</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="DotNode.Copy">func</a> (\*DotNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=9053:9082#L370)
<pre>func (d *<a href="#DotNode">DotNode</a>) Copy() <a href="#Node">Node</a></pre>



### <a id="DotNode.String">func</a> (\*DotNode) [String](https://golang.org/src/text/template/parse/node.go?s=8953:8986#L362)
<pre>func (d *<a href="#DotNode">DotNode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="DotNode.Type">func</a> (\*DotNode) [Type](https://golang.org/src/text/template/parse/node.go?s=8742:8775#L355)
<pre>func (d *<a href="#DotNode">DotNode</a>) Type() <a href="#NodeType">NodeType</a></pre>



## <a id="FieldNode">type</a> [FieldNode](https://golang.org/src/text/template/parse/node.go?s=9861:9967#L407)
FieldNode holds a field (identifier starting with '.').
The names may be chained ('.x.y').
The period is dropped from each ident.


<pre>type FieldNode struct {
    <a href="#NodeType">NodeType</a>
    <a href="#Pos">Pos</a>

<span id="FieldNode.Ident"></span>    Ident []<a href="/pkg/builtin/#string">string</a> <span class="comment">// The identifiers in lexical order.</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="FieldNode.Copy">func</a> (\*FieldNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=10318:10349#L430)
<pre>func (f *<a href="#FieldNode">FieldNode</a>) Copy() <a href="#Node">Node</a></pre>



### <a id="FieldNode.String">func</a> (\*FieldNode) [String](https://golang.org/src/text/template/parse/node.go?s=10158:10193#L418)
<pre>func (f *<a href="#FieldNode">FieldNode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="IdentifierNode">type</a> [IdentifierNode](https://golang.org/src/text/template/parse/node.go?s=6803:6901#L272)
IdentifierNode holds an identifier.


<pre>type IdentifierNode struct {
    <a href="#NodeType">NodeType</a>
    <a href="#Pos">Pos</a>

<span id="IdentifierNode.Ident"></span>    Ident <a href="/pkg/builtin/#string">string</a> <span class="comment">// The identifier&#39;s name.</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="NewIdentifier">func</a> [NewIdentifier](https://golang.org/src/text/template/parse/node.go?s=6981:7029#L280)
<pre>func NewIdentifier(ident <a href="/pkg/builtin/#string">string</a>) *<a href="#IdentifierNode">IdentifierNode</a></pre>
NewIdentifier returns a new IdentifierNode with the given identifier name.






### <a id="IdentifierNode.Copy">func</a> (\*IdentifierNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=7694:7730#L308)
<pre>func (i *<a href="#IdentifierNode">IdentifierNode</a>) Copy() <a href="#Node">Node</a></pre>



### <a id="IdentifierNode.SetPos">func</a> (\*IdentifierNode) [SetPos](https://golang.org/src/text/template/parse/node.go?s=7245:7301#L287)
<pre>func (i *<a href="#IdentifierNode">IdentifierNode</a>) SetPos(pos <a href="#Pos">Pos</a>) *<a href="#IdentifierNode">IdentifierNode</a></pre>
SetPos sets the position. NewIdentifier is a public method so we can't modify its signature.
Chained for convenience.
TODO: fix one day?




### <a id="IdentifierNode.SetTree">func</a> (\*IdentifierNode) [SetTree](https://golang.org/src/text/template/parse/node.go?s=7493:7550#L295)
<pre>func (i *<a href="#IdentifierNode">IdentifierNode</a>) SetTree(t *<a href="#Tree">Tree</a>) *<a href="#IdentifierNode">IdentifierNode</a></pre>
SetTree sets the parent tree for the node. NewIdentifier is a public method so we can't modify its signature.
Chained for convenience.
TODO: fix one day?




### <a id="IdentifierNode.String">func</a> (\*IdentifierNode) [String](https://golang.org/src/text/template/parse/node.go?s=7576:7616#L300)
<pre>func (i *<a href="#IdentifierNode">IdentifierNode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="IfNode">type</a> [IfNode](https://golang.org/src/text/template/parse/node.go?s=18748:18782#L766)
IfNode represents an {{if}} action and its commands.


<pre>type IfNode struct {
    <a href="#BranchNode">BranchNode</a>
}
</pre>











### <a id="IfNode.Copy">func</a> (\*IfNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=18998:19026#L774)
<pre>func (i *<a href="#IfNode">IfNode</a>) Copy() <a href="#Node">Node</a></pre>



## <a id="ListNode">type</a> [ListNode](https://golang.org/src/text/template/parse/node.go?s=2597:2702#L67)
ListNode holds a sequence of nodes.


<pre>type ListNode struct {
    <a href="#NodeType">NodeType</a>
    <a href="#Pos">Pos</a>

<span id="ListNode.Nodes"></span>    Nodes []<a href="#Node">Node</a> <span class="comment">// The element nodes in lexical order.</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="ListNode.Copy">func</a> (\*ListNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=3228:3258#L105)
<pre>func (l *<a href="#ListNode">ListNode</a>) Copy() <a href="#Node">Node</a></pre>



### <a id="ListNode.CopyList">func</a> (\*ListNode) [CopyList](https://golang.org/src/text/template/parse/node.go?s=3059:3098#L94)
<pre>func (l *<a href="#ListNode">ListNode</a>) CopyList() *<a href="#ListNode">ListNode</a></pre>



### <a id="ListNode.String">func</a> (\*ListNode) [String](https://golang.org/src/text/template/parse/node.go?s=2925:2959#L86)
<pre>func (l *<a href="#ListNode">ListNode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="NilNode">type</a> [NilNode](https://golang.org/src/text/template/parse/node.go?s=9199:9247#L375)
NilNode holds the special identifier 'nil' representing an untyped nil constant.


<pre>type NilNode struct {
    <a href="#NodeType">NodeType</a>
    <a href="#Pos">Pos</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="NilNode.Copy">func</a> (\*NilNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=9660:9689#L400)
<pre>func (n *<a href="#NilNode">NilNode</a>) Copy() <a href="#Node">Node</a></pre>



### <a id="NilNode.String">func</a> (\*NilNode) [String](https://golang.org/src/text/template/parse/node.go?s=9558:9591#L392)
<pre>func (n *<a href="#NilNode">NilNode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



### <a id="NilNode.Type">func</a> (\*NilNode) [Type](https://golang.org/src/text/template/parse/node.go?s=9347:9380#L385)
<pre>func (n *<a href="#NilNode">NilNode</a>) Type() <a href="#NodeType">NodeType</a></pre>



## <a id="Node">type</a> [Node](https://golang.org/src/text/template/parse/node.go?s=496:942#L11)
A Node is an element in the parse tree. The interface is trivial.
The interface contains an unexported method so that only
types local to this package can satisfy it.


<pre>type Node interface {
    Type() <a href="#NodeType">NodeType</a>
    String() <a href="/pkg/builtin/#string">string</a>
    <span class="comment">// Copy does a deep copy of the Node and all its components.</span>
    <span class="comment">// To avoid type assertions, some XxxNodes also have specialized</span>
    <span class="comment">// CopyXxx methods that return *XxxNode.</span>
    Copy() <a href="#Node">Node</a>
    Position() <a href="#Pos">Pos</a> <span class="comment">// byte position of start of node in full original input string</span>
    <span class="comment">// contains filtered or unexported methods</span>
}</pre>











## <a id="NodeType">type</a> [NodeType](https://golang.org/src/text/template/parse/node.go?s=998:1015#L25)
NodeType identifies the type of a parse tree node.


<pre>type NodeType <a href="/pkg/builtin/#int">int</a></pre>



<pre>const (
    <span id="NodeText">NodeText</span>    <a href="#NodeType">NodeType</a> = <a href="/pkg/builtin/#iota">iota</a> <span class="comment">// Plain text.</span>
    <span id="NodeAction">NodeAction</span>                  <span class="comment">// A non-control action such as a field evaluation.</span>
    <span id="NodeBool">NodeBool</span>                    <span class="comment">// A boolean constant.</span>
    <span id="NodeChain">NodeChain</span>                   <span class="comment">// A sequence of field accesses.</span>
    <span id="NodeCommand">NodeCommand</span>                 <span class="comment">// An element of a pipeline.</span>
    <span id="NodeDot">NodeDot</span>                     <span class="comment">// The cursor, dot.</span>

    <span id="NodeField">NodeField</span>      <span class="comment">// A field or method name.</span>
    <span id="NodeIdentifier">NodeIdentifier</span> <span class="comment">// An identifier; always a function name.</span>
    <span id="NodeIf">NodeIf</span>         <span class="comment">// An if action.</span>
    <span id="NodeList">NodeList</span>       <span class="comment">// A list of Nodes.</span>
    <span id="NodeNil">NodeNil</span>        <span class="comment">// An untyped nil constant.</span>
    <span id="NodeNumber">NodeNumber</span>     <span class="comment">// A numerical constant.</span>
    <span id="NodePipe">NodePipe</span>       <span class="comment">// A pipeline of commands.</span>
    <span id="NodeRange">NodeRange</span>      <span class="comment">// A range action.</span>
    <span id="NodeString">NodeString</span>     <span class="comment">// A string constant.</span>
    <span id="NodeTemplate">NodeTemplate</span>   <span class="comment">// A template invocation action.</span>
    <span id="NodeVariable">NodeVariable</span>   <span class="comment">// A $ variable.</span>
    <span id="NodeWith">NodeWith</span>       <span class="comment">// A with action.</span>
)</pre>









### <a id="NodeType.Type">func</a> (NodeType) [Type](https://golang.org/src/text/template/parse/node.go?s=1305:1338#L37)
<pre>func (t <a href="#NodeType">NodeType</a>) Type() <a href="#NodeType">NodeType</a></pre>
Type returns itself and provides an easy default implementation
for embedding in a Node. Embedded in all non-trivial Nodes.




## <a id="NumberNode">type</a> [NumberNode](https://golang.org/src/text/template/parse/node.go?s=12315:12882#L510)
NumberNode holds a number: signed or unsigned integer, float, or complex.
The value is parsed and stored under all the types that can represent the value.
This simulates in a small amount of code the behavior of Go's ideal constants.


<pre>type NumberNode struct {
    <a href="#NodeType">NodeType</a>
    <a href="#Pos">Pos</a>

<span id="NumberNode.IsInt"></span>    IsInt      <a href="/pkg/builtin/#bool">bool</a>       <span class="comment">// Number has an integral value.</span>
<span id="NumberNode.IsUint"></span>    IsUint     <a href="/pkg/builtin/#bool">bool</a>       <span class="comment">// Number has an unsigned integral value.</span>
<span id="NumberNode.IsFloat"></span>    IsFloat    <a href="/pkg/builtin/#bool">bool</a>       <span class="comment">// Number has a floating-point value.</span>
<span id="NumberNode.IsComplex"></span>    IsComplex  <a href="/pkg/builtin/#bool">bool</a>       <span class="comment">// Number is complex.</span>
<span id="NumberNode.Int64"></span>    Int64      <a href="/pkg/builtin/#int64">int64</a>      <span class="comment">// The signed integer value.</span>
<span id="NumberNode.Uint64"></span>    Uint64     <a href="/pkg/builtin/#uint64">uint64</a>     <span class="comment">// The unsigned integer value.</span>
<span id="NumberNode.Float64"></span>    Float64    <a href="/pkg/builtin/#float64">float64</a>    <span class="comment">// The floating-point value.</span>
<span id="NumberNode.Complex128"></span>    Complex128 <a href="/pkg/builtin/#complex128">complex128</a> <span class="comment">// The complex value.</span>
<span id="NumberNode.Text"></span>    Text       <a href="/pkg/builtin/#string">string</a>     <span class="comment">// The original textual representation from the input.</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="NumberNode.Copy">func</a> (\*NumberNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=15814:15846#L636)
<pre>func (n *<a href="#NumberNode">NumberNode</a>) Copy() <a href="#Node">Node</a></pre>



### <a id="NumberNode.String">func</a> (\*NumberNode) [String](https://golang.org/src/text/template/parse/node.go?s=15705:15741#L628)
<pre>func (n *<a href="#NumberNode">NumberNode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="PipeNode">type</a> [PipeNode](https://golang.org/src/text/template/parse/node.go?s=3864:4208#L134)
PipeNode holds a pipeline with optional declaration


<pre>type PipeNode struct {
    <a href="#NodeType">NodeType</a>
    <a href="#Pos">Pos</a>

<span id="PipeNode.Line"></span>    Line     <a href="/pkg/builtin/#int">int</a>             <span class="comment">// The line number in the input. Deprecated: Kept for compatibility.</span>
<span id="PipeNode.IsAssign"></span>    IsAssign <a href="/pkg/builtin/#bool">bool</a>            <span class="comment">// The variables are being assigned, not declared.</span>
<span id="PipeNode.Decl"></span>    Decl     []*<a href="#VariableNode">VariableNode</a> <span class="comment">// Variables in lexical order.</span>
<span id="PipeNode.Cmds"></span>    Cmds     []*<a href="#CommandNode">CommandNode</a>  <span class="comment">// The commands in lexical order.</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="PipeNode.Copy">func</a> (\*PipeNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=5097:5127#L192)
<pre>func (p *<a href="#PipeNode">PipeNode</a>) Copy() <a href="#Node">Node</a></pre>



### <a id="PipeNode.CopyPipe">func</a> (\*PipeNode) [CopyPipe](https://golang.org/src/text/template/parse/node.go?s=4772:4811#L176)
<pre>func (p *<a href="#PipeNode">PipeNode</a>) CopyPipe() *<a href="#PipeNode">PipeNode</a></pre>



### <a id="PipeNode.String">func</a> (\*PipeNode) [String](https://golang.org/src/text/template/parse/node.go?s=4459:4493#L152)
<pre>func (p *<a href="#PipeNode">PipeNode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Pos">type</a> [Pos](https://golang.org/src/text/template/parse/node.go?s=1118:1130#L29)
Pos represents a byte position in the original input text from which
this template was parsed.


<pre>type Pos <a href="/pkg/builtin/#int">int</a></pre>











### <a id="Pos.Position">func</a> (Pos) [Position](https://golang.org/src/text/template/parse/node.go?s=1132:1159#L31)
<pre>func (p <a href="#Pos">Pos</a>) Position() <a href="#Pos">Pos</a></pre>



## <a id="RangeNode">type</a> [RangeNode](https://golang.org/src/text/template/parse/node.go?s=19188:19225#L779)
RangeNode represents a {{range}} action and its commands.


<pre>type RangeNode struct {
    <a href="#BranchNode">BranchNode</a>
}
</pre>











### <a id="RangeNode.Copy">func</a> (\*RangeNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=19453:19484#L787)
<pre>func (r *<a href="#RangeNode">RangeNode</a>) Copy() <a href="#Node">Node</a></pre>



## <a id="StringNode">type</a> [StringNode](https://golang.org/src/text/template/parse/node.go?s=15990:16163#L643)
StringNode holds a string constant. The value has been "unquoted".


<pre>type StringNode struct {
    <a href="#NodeType">NodeType</a>
    <a href="#Pos">Pos</a>

<span id="StringNode.Quoted"></span>    Quoted <a href="/pkg/builtin/#string">string</a> <span class="comment">// The original text of the string, with quotes.</span>
<span id="StringNode.Text"></span>    Text   <a href="/pkg/builtin/#string">string</a> <span class="comment">// The string, after quote processing.</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="StringNode.Copy">func</a> (\*StringNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=16431:16463#L663)
<pre>func (s *<a href="#StringNode">StringNode</a>) Copy() <a href="#Node">Node</a></pre>



### <a id="StringNode.String">func</a> (\*StringNode) [String](https://golang.org/src/text/template/parse/node.go?s=16320:16356#L655)
<pre>func (s *<a href="#StringNode">StringNode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="TemplateNode">type</a> [TemplateNode](https://golang.org/src/text/template/parse/node.go?s=20090:20354#L805)
TemplateNode represents a {{template}} action.


<pre>type TemplateNode struct {
    <a href="#NodeType">NodeType</a>
    <a href="#Pos">Pos</a>

<span id="TemplateNode.Line"></span>    Line <a href="/pkg/builtin/#int">int</a>       <span class="comment">// The line number in the input. Deprecated: Kept for compatibility.</span>
<span id="TemplateNode.Name"></span>    Name <a href="/pkg/builtin/#string">string</a>    <span class="comment">// The name of the template (unquoted).</span>
<span id="TemplateNode.Pipe"></span>    Pipe *<a href="#PipeNode">PipeNode</a> <span class="comment">// The command to evaluate as dot for the template.</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="TemplateNode.Copy">func</a> (\*TemplateNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=20776:20810#L829)
<pre>func (t *<a href="#TemplateNode">TemplateNode</a>) Copy() <a href="#Node">Node</a></pre>



### <a id="TemplateNode.String">func</a> (\*TemplateNode) [String](https://golang.org/src/text/template/parse/node.go?s=20549:20587#L818)
<pre>func (t *<a href="#TemplateNode">TemplateNode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="TextNode">type</a> [TextNode](https://golang.org/src/text/template/parse/node.go?s=3315:3411#L110)
TextNode holds plain text.


<pre>type TextNode struct {
    <a href="#NodeType">NodeType</a>
    <a href="#Pos">Pos</a>

<span id="TextNode.Text"></span>    Text []<a href="/pkg/builtin/#byte">byte</a> <span class="comment">// The text; may span newlines.</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="TextNode.Copy">func</a> (\*TextNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=3678:3708#L129)
<pre>func (t *<a href="#TextNode">TextNode</a>) Copy() <a href="#Node">Node</a></pre>



### <a id="TextNode.String">func</a> (\*TextNode) [String](https://golang.org/src/text/template/parse/node.go?s=3548:3582#L121)
<pre>func (t *<a href="#TextNode">TextNode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Tree">type</a> [Tree](https://golang.org/src/text/template/parse/parse.go?s=557:1113#L10)
Tree is the representation of a single parsed template.


<pre>type Tree struct {
<span id="Tree.Name"></span>    Name      <a href="/pkg/builtin/#string">string</a>    <span class="comment">// name of the template represented by the tree.</span>
<span id="Tree.ParseName"></span>    ParseName <a href="/pkg/builtin/#string">string</a>    <span class="comment">// name of the top-level template during parsing, for error messages.</span>
<span id="Tree.Root"></span>    Root      *<a href="#ListNode">ListNode</a> <span class="comment">// top-level root of the tree.</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>









### <a id="New">func</a> [New](https://golang.org/src/text/template/parse/parse.go?s=3234:3294#L115)
<pre>func New(name <a href="/pkg/builtin/#string">string</a>, funcs ...map[<a href="/pkg/builtin/#string">string</a>]interface{}) *<a href="#Tree">Tree</a></pre>
New allocates a new parse tree with the given name.






### <a id="Tree.Copy">func</a> (\*Tree) [Copy](https://golang.org/src/text/template/parse/parse.go?s=1183:1210#L25)
<pre>func (t *<a href="#Tree">Tree</a>) Copy() *<a href="#Tree">Tree</a></pre>
Copy returns a copy of the Tree. Any parsing state is discarded.




### <a id="Tree.ErrorContext">func</a> (\*Tree) [ErrorContext](https://golang.org/src/text/template/parse/parse.go?s=3564:3626#L125)
<pre>func (t *<a href="#Tree">Tree</a>) ErrorContext(n <a href="#Node">Node</a>) (location, context <a href="/pkg/builtin/#string">string</a>)</pre>
ErrorContext returns a textual representation of the location of the node in the input text.
The receiver is only used when the node does not have a pointer to the tree inside,
which can occur in old code.




### <a id="Tree.Parse">func</a> (\*Tree) [Parse](https://golang.org/src/text/template/parse/parse.go?s=6019:6158#L215)
<pre>func (t *<a href="#Tree">Tree</a>) Parse(text, leftDelim, rightDelim <a href="/pkg/builtin/#string">string</a>, treeSet map[<a href="/pkg/builtin/#string">string</a>]*<a href="#Tree">Tree</a>, funcs ...map[<a href="/pkg/builtin/#string">string</a>]interface{}) (tree *<a href="#Tree">Tree</a>, err <a href="/pkg/builtin/#error">error</a>)</pre>
Parse parses the template definition string to construct a representation of
the template for execution. If either action delimiter string is empty, the
default ("{{" or "}}") is used. Embedded template definitions are added to
the treeSet map.




## <a id="VariableNode">type</a> [VariableNode](https://golang.org/src/text/template/parse/node.go?s=7927:8045#L314)
AssignNode holds a list of variable names, possibly with chained field
accesses. The dollar sign is part of the (first) name.


<pre>type VariableNode struct {
    <a href="#NodeType">NodeType</a>
    <a href="#Pos">Pos</a>

<span id="VariableNode.Ident"></span>    Ident []<a href="/pkg/builtin/#string">string</a> <span class="comment">// Variable name and fields in lexical order.</span>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="VariableNode.Copy">func</a> (\*VariableNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=8402:8436#L340)
<pre>func (v *<a href="#VariableNode">VariableNode</a>) Copy() <a href="#Node">Node</a></pre>



### <a id="VariableNode.String">func</a> (\*VariableNode) [String](https://golang.org/src/text/template/parse/node.go?s=8213:8251#L325)
<pre>func (v *<a href="#VariableNode">VariableNode</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="WithNode">type</a> [WithNode](https://golang.org/src/text/template/parse/node.go?s=19647:19683#L792)
WithNode represents a {{with}} action and its commands.


<pre>type WithNode struct {
    <a href="#BranchNode">BranchNode</a>
}
</pre>











### <a id="WithNode.Copy">func</a> (\*WithNode) [Copy](https://golang.org/src/text/template/parse/node.go?s=19907:19937#L800)
<pre>func (w *<a href="#WithNode">WithNode</a>) Copy() <a href="#Node">Node</a></pre>






