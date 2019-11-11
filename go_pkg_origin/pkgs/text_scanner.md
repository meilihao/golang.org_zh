

# scanner
`import "text/scanner"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package scanner provides a scanner and tokenizer for UTF-8-encoded text.
It takes an io.Reader providing the source, which then can be tokenized
through repeated calls to the Scan function. For compatibility with
existing tools, the NUL character is not allowed. If the first character
in the source is a UTF-8 encoded byte order mark (BOM), it is discarded.

By default, a Scanner skips white space and Go comments and recognizes all
literals as defined by the Go language specification. It may be
customized to recognize only a subset of those literals and to recognize
different identifier and white space characters.



<a id="example_">Example</a>


```go
```

output:
```txt
```

<a id="example__isIdentRune">Example (IsIdentRune)</a>


```go
```

output:
```txt
```

<a id="example__mode">Example (Mode)</a>


```go
```

output:
```txt
```

<a id="example__whitespace">Example (Whitespace)</a>


```go
```

output:
```txt
```


## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func TokenString(tok rune) string](#TokenString)
* [type Position](#Position)
  * [func (pos *Position) IsValid() bool](#Position.IsValid)
  * [func (pos Position) String() string](#Position.String)
* [type Scanner](#Scanner)
  * [func (s *Scanner) Init(src io.Reader) *Scanner](#Scanner.Init)
  * [func (s *Scanner) Next() rune](#Scanner.Next)
  * [func (s *Scanner) Peek() rune](#Scanner.Peek)
  * [func (s *Scanner) Pos() (pos Position)](#Scanner.Pos)
  * [func (s *Scanner) Scan() rune](#Scanner.Scan)
  * [func (s *Scanner) TokenText() string](#Scanner.TokenText)


#### <a id="pkg-examples">Examples</a>
* [Package](#example_)
* [Package (IsIdentRune)](#example__isIdentRune)
* [Package (Mode)](#example__mode)
* [Package (Whitespace)](#example__whitespace)


#### <a id="pkg-files">Package files</a>
[scanner.go](https://golang.org/src/text/scanner/scanner.go) 


## <a id="pkg-constants">Constants</a>
Predefined mode bits to control recognition of tokens. For instance,
to configure a Scanner such that it only recognizes (Go) identifiers,
integers, and skips comments, set the Scanner's Mode field to:


	ScanIdents | ScanInts | SkipComments

With the exceptions of comments, which are skipped if SkipComments is
set, unrecognized tokens are not ignored. Instead, the scanner simply
returns the respective individual characters (or possibly sub-tokens).
For instance, if the mode is ScanIdents (not ScanStrings), the string
"foo" is scanned as the token sequence '"' Ident '"'.

Use GoTokens to configure the Scanner such that it accepts all Go
literal tokens including Go identifiers. Comments will be skipped.


<pre>const (
    <span id="ScanIdents">ScanIdents</span>     = 1 &lt;&lt; -<a href="#Ident">Ident</a>
    <span id="ScanInts">ScanInts</span>       = 1 &lt;&lt; -<a href="#Int">Int</a>
    <span id="ScanFloats">ScanFloats</span>     = 1 &lt;&lt; -<a href="#Float">Float</a> <span class="comment">// includes Ints and hexadecimal floats</span>
    <span id="ScanChars">ScanChars</span>      = 1 &lt;&lt; -<a href="#Char">Char</a>
    <span id="ScanStrings">ScanStrings</span>    = 1 &lt;&lt; -<a href="#String">String</a>
    <span id="ScanRawStrings">ScanRawStrings</span> = 1 &lt;&lt; -<a href="#RawString">RawString</a>
    <span id="ScanComments">ScanComments</span>   = 1 &lt;&lt; -<a href="#Comment">Comment</a>
    <span id="SkipComments">SkipComments</span>   = 1 &lt;&lt; -skipComment <span class="comment">// if set with ScanComments, comments become white space</span>
    <span id="GoTokens">GoTokens</span>       = <a href="#ScanIdents">ScanIdents</a> | <a href="#ScanFloats">ScanFloats</a> | <a href="#ScanChars">ScanChars</a> | <a href="#ScanStrings">ScanStrings</a> | <a href="#ScanRawStrings">ScanRawStrings</a> | <a href="#ScanComments">ScanComments</a> | <a href="#SkipComments">SkipComments</a>
)</pre>The result of Scan is one of these tokens or a Unicode character.


<pre>const (
    <span id="EOF">EOF</span> = -(<a href="/pkg/builtin/#iota">iota</a> + 1)
    <span id="Ident">Ident</span>
    <span id="Int">Int</span>
    <span id="Float">Float</span>
    <span id="Char">Char</span>
    <span id="String">String</span>
    <span id="RawString">RawString</span>
    <span id="Comment">Comment</span>
)</pre>GoWhitespace is the default value for the Scanner's Whitespace field.
Its value selects Go's white space characters.


<pre>const <span id="GoWhitespace">GoWhitespace</span> = 1&lt;&lt;&#39;\t&#39; | 1&lt;&lt;&#39;\n&#39; | 1&lt;&lt;&#39;\r&#39; | 1&lt;&lt;&#39; &#39;</pre>



## <a id="TokenString">func</a> [TokenString](https://golang.org/src/text/scanner/scanner.go?s=3205:3238#L93)
<pre>func TokenString(tok <a href="/pkg/builtin/#rune">rune</a>) <a href="/pkg/builtin/#string">string</a></pre>
TokenString returns a printable string for a token or Unicode character.





## <a id="Position">type</a> [Position](https://golang.org/src/text/scanner/scanner.go?s=987:1218#L18)
A source position is represented by a Position value.
A position is valid if Line > 0.


<pre>type Position struct {
<span id="Position.Filename"></span>    Filename <a href="/pkg/builtin/#string">string</a> <span class="comment">// filename, if any</span>
<span id="Position.Offset"></span>    Offset   <a href="/pkg/builtin/#int">int</a>    <span class="comment">// byte offset, starting at 0</span>
<span id="Position.Line"></span>    Line     <a href="/pkg/builtin/#int">int</a>    <span class="comment">// line number, starting at 1</span>
<span id="Position.Column"></span>    Column   <a href="/pkg/builtin/#int">int</a>    <span class="comment">// column number, starting at 1 (character count per line)</span>
}
</pre>











### <a id="Position.IsValid">func</a> (\*Position) [IsValid](https://golang.org/src/text/scanner/scanner.go?s=1270:1305#L26)
<pre>func (pos *<a href="#Position">Position</a>) IsValid() <a href="/pkg/builtin/#bool">bool</a></pre>
IsValid reports whether the position is valid.




### <a id="Position.String">func</a> (Position) [String](https://golang.org/src/text/scanner/scanner.go?s=1331:1366#L28)
<pre>func (pos <a href="#Position">Position</a>) String() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="Scanner">type</a> [Scanner](https://golang.org/src/text/scanner/scanner.go?s=3650:6185#L107)
A Scanner implements reading of Unicode characters and tokens from an io.Reader.


<pre>type Scanner struct {

<span id="Scanner.Error"></span>    <span class="comment">// Error is called for each error encountered. If no Error</span>
    <span class="comment">// function is set, the error is reported to os.Stderr.</span>
    Error func(s *<a href="#Scanner">Scanner</a>, msg <a href="/pkg/builtin/#string">string</a>)

<span id="Scanner.ErrorCount"></span>    <span class="comment">// ErrorCount is incremented by one for each error encountered.</span>
    ErrorCount <a href="/pkg/builtin/#int">int</a>

    <span class="comment">// The Mode field controls which tokens are recognized. For instance,</span>
    <span class="comment">// to recognize Ints, set the ScanInts bit in Mode. The field may be</span>
    <span class="comment">// changed at any time.</span>
<span id="Scanner.Mode"></span>    Mode <a href="/pkg/builtin/#uint">uint</a>

    <span class="comment">// The Whitespace field controls which characters are recognized</span>
    <span class="comment">// as white space. To recognize a character ch &lt;= &#39; &#39; as white space,</span>
    <span class="comment">// set the ch&#39;th bit in Whitespace (the Scanner&#39;s behavior is undefined</span>
    <span class="comment">// for values ch &gt; &#39; &#39;). The field may be changed at any time.</span>
<span id="Scanner.Whitespace"></span>    Whitespace <a href="/pkg/builtin/#uint64">uint64</a>

<span id="Scanner.IsIdentRune"></span>    <span class="comment">// IsIdentRune is a predicate controlling the characters accepted</span>
    <span class="comment">// as the ith rune in an identifier. The set of valid characters</span>
    <span class="comment">// must not intersect with the set of white space characters.</span>
    <span class="comment">// If no IsIdentRune function is set, regular Go identifiers are</span>
    <span class="comment">// accepted instead. The field may be changed at any time.</span>
    IsIdentRune func(ch <a href="/pkg/builtin/#rune">rune</a>, i <a href="/pkg/builtin/#int">int</a>) <a href="/pkg/builtin/#bool">bool</a>

    <span class="comment">// Start position of most recently scanned token; set by Scan.</span>
    <span class="comment">// Calling Init or Next invalidates the position (Line == 0).</span>
    <span class="comment">// The Filename field is always left untouched by the Scanner.</span>
    <span class="comment">// If an error is reported (via Error) and Position is invalid,</span>
    <span class="comment">// the scanner is not inside a token. Call Pos to obtain an error</span>
    <span class="comment">// position in that case, or to obtain the position immediately</span>
    <span class="comment">// after the most recently scanned token.</span>
    <a href="#Position">Position</a>
    <span class="comment">// contains filtered or unexported fields</span>
}
</pre>











### <a id="Scanner.Init">func</a> (\*Scanner) [Init](https://golang.org/src/text/scanner/scanner.go?s=6365:6411#L172)
<pre>func (s *<a href="#Scanner">Scanner</a>) Init(src <a href="/pkg/io/">io</a>.<a href="/pkg/io/#Reader">Reader</a>) *<a href="#Scanner">Scanner</a></pre>
Init initializes a Scanner with a new source and returns s.
Error is set to nil, ErrorCount is set to 0, Mode is set to GoTokens,
and Whitespace is set to GoWhitespace.




### <a id="Scanner.Next">func</a> (\*Scanner) [Next](https://golang.org/src/text/scanner/scanner.go?s=9676:9705#L295)
<pre>func (s *<a href="#Scanner">Scanner</a>) Next() <a href="/pkg/builtin/#rune">rune</a></pre>
Next reads and returns the next Unicode character.
It returns EOF at the end of the source. It reports
a read error by calling s.Error, if not nil; otherwise
it prints an error message to os.Stderr. Next does not
update the Scanner's Position field; use Pos() to
get the current position.




### <a id="Scanner.Peek">func</a> (\*Scanner) [Peek](https://golang.org/src/text/scanner/scanner.go?s=10037:10066#L308)
<pre>func (s *<a href="#Scanner">Scanner</a>) Peek() <a href="/pkg/builtin/#rune">rune</a></pre>
Peek returns the next Unicode character in the source without advancing
the scanner. It returns EOF if the scanner's position is at the last
character of the source.




### <a id="Scanner.Pos">func</a> (\*Scanner) [Pos](https://golang.org/src/text/scanner/scanner.go?s=19451:19489#L739)
<pre>func (s *<a href="#Scanner">Scanner</a>) Pos() (pos <a href="#Position">Position</a>)</pre>
Pos returns the position of the character immediately after
the character or token returned by the last call to Next or Scan.
Use the Scanner's Position field for the start position of the most
recently scanned token.




### <a id="Scanner.Scan">func</a> (\*Scanner) [Scan](https://golang.org/src/text/scanner/scanner.go?s=17337:17366#L637)
<pre>func (s *<a href="#Scanner">Scanner</a>) Scan() <a href="/pkg/builtin/#rune">rune</a></pre>
Scan reads the next token or Unicode character from source and returns it.
It only recognizes tokens t for which the respective Mode bit (1<<-t) is set.
It returns EOF at the end of the source. It reports scanner errors (read and
token errors) by calling s.Error, if not nil; otherwise it prints an error
message to os.Stderr.




### <a id="Scanner.TokenText">func</a> (\*Scanner) [TokenText](https://golang.org/src/text/scanner/scanner.go?s=20039:20075#L761)
<pre>func (s *<a href="#Scanner">Scanner</a>) TokenText() <a href="/pkg/builtin/#string">string</a></pre>
TokenText returns the string corresponding to the most recently scanned token.
Valid after calling Scan and in calls of Scanner.Error.







