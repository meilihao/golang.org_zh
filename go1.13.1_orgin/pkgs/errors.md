

# errors
`import "errors"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Package errors implements functions to manipulate errors.

The New function creates errors whose only content is a text message.

The Unwrap, Is and As functions work on errors that may wrap other errors.
An error wraps another error if its type has the method


	Unwrap() error

If e.Unwrap() returns a non-nil error w, then we say that e wraps w.

A simple way to create wrapped errors is to call fmt.Errorf and apply the %w verb
to the error argument:


	fmt.Errorf("... %w ...", ..., err, ...).Unwrap()

returns err.

Unwrap unpacks wrapped errors. If its argument's type has an
Unwrap method, it calls the method once. Otherwise, it returns nil.

Is unwraps its first argument sequentially looking for an error that matches the
second. It reports whether it finds a match. It should be used in preference to
simple equality checks:


	if errors.Is(err, os.ErrExist)

is preferable to


	if err == os.ErrExist

because the former will succeed if err wraps os.ErrExist.

As unwraps its first argument sequentially looking for an error that can be
assigned to its second argument, which must be a pointer. If it succeeds, it
performs the assignment and returns true. Otherwise, it returns false. The form


	var perr *os.PathError
	if errors.As(err, &perr) {
		fmt.Println(perr.Path)
	}

is preferable to


	if perr, ok := err.(*os.PathError); ok {
		fmt.Println(perr.Path)
	}

because the former will succeed if err wraps an *os.PathError.


<a id="example_">Example</a>


## <a id="pkg-index">Index</a>
* [func As(err error, target interface{}) bool](#As)
* [func Is(err, target error) bool](#Is)
* [func New(text string) error](#New)
* [func Unwrap(err error) error](#Unwrap)


#### <a id="pkg-examples">Examples</a>
* [Package](#example_)
* [As](#example_As)
* [New](#example_New)
* [New (Errorf)](#example_New_errorf)


#### <a id="pkg-files">Package files</a>
[errors.go](https://golang.org/src/errors/errors.go) [wrap.go](https://golang.org/src/errors/wrap.go) 






## <a id="As">func</a> [As](https://golang.org/src/errors/wrap.go?s=2058:2101#L56)
<pre>func As(err <a href="/pkg/builtin/#error">error</a>, target interface{}) <a href="/pkg/builtin/#bool">bool</a></pre>
As finds the first error in err's chain that matches target, and if so, sets
target to that error value and returns true.

The chain consists of err itself followed by the sequence of errors obtained by
repeatedly calling Unwrap.

An error matches target if the error's concrete value is assignable to the value
pointed to by target, or if the error has a method As(interface{}) bool such that
As(target) returns true. In the latter case, the As method is responsible for
setting target.

As will panic if target is not a non-nil pointer to either a type that implements
error, or to any interface type. As returns false if err is nil.


<a id="example_As">Example</a>

## <a id="Is">func</a> [Is](https://golang.org/src/errors/wrap.go?s=837:868#L21)
<pre>func Is(err, target <a href="/pkg/builtin/#error">error</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Is reports whether any error in err's chain matches target.

The chain consists of err itself followed by the sequence of errors obtained by
repeatedly calling Unwrap.

An error is considered to match a target if it is equal to that target or if
it implements a method Is(error) bool such that Is(target) returns true.



## <a id="New">func</a> [New](https://golang.org/src/errors/errors.go?s=1869:1896#L48)
<pre>func New(text <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
New returns an error that formats as the given text.
Each call to New returns a distinct error value even if the text is identical.


<a id="example_New">Example</a>
<a id="example_New_errorf">Example (Errorf)</a>
<p>The fmt package&#39;s Errorf function lets us use the package&#39;s formatting
features to create descriptive error messages.
</p>
## <a id="Unwrap">func</a> [Unwrap](https://golang.org/src/errors/wrap.go?s=372:400#L4)
<pre>func Unwrap(err <a href="/pkg/builtin/#error">error</a>) <a href="/pkg/builtin/#error">error</a></pre>
Unwrap returns the result of calling the Unwrap method on err, if err's
type contains an Unwrap method returning error.
Otherwise, Unwrap returns nil.









