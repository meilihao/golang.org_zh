

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

Unwrap unpacks wrapped errors. If its argument's type has an
Unwrap method, it calls the method once. Otherwise, it returns nil.

A simple way to create wrapped errors is to call fmt.Errorf and apply the %w verb
to the error argument:


	errors.Unwrap(fmt.Errorf("... %w ...", ..., err, ...))

returns err.

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

errors包实现了用于处理错误的函数.

`New`函数会创建仅包含文本信息的`error`.

`Unwrap`,`Is`以及`As`函数用于处理`error`可能封装了其他`error`的情况. 如果它的类型有下面的方法, 那么该`error`就封装了另一个`error`:
```go
Unwrap() error
```

如果`e.Unwrap()`返回了一个不为nil的`error w`, 那么我们就说`e`封装了`w`.

Unwrap 可解出封装的error. 如果参数类型有 Unwrap 方法,调用一次即可得到error, 否则它会返回nil.

通过fmt.Errorf并使用`%w`格式即可创建一个封装过的error, 它会返回一个error.

```go
	errors.Unwrap(fmt.Errorf("... %w ...", ..., err, ...))
```

Is 会依次解封第一个参数以查找与第二个参数匹配的error. 它会判断是否找到. 它优于使用简单地比较:
```go
	if errors.Is(err, os.ErrExist)
```
优于:
```go
	if err == os.ErrExist
```
因为如果err封装了os.ErrExist, 前者能正常判断.

As 会依次解封第一个参数以查找能断言给第二个参数的error, 第二个参数必须是指针. 如果它能赋值成功就返回true, 否则返回false.


```go
	var perr *os.PathError
	if errors.As(err, &perr) {
		fmt.Println(perr.Path)
	}
```
优于:
```go
	if perr, ok := err.(*os.PathError); ok {
		fmt.Println(perr.Path)
	}
```

因为如果err封装了*os.PathError, 前者能正常判断.

<a id="example_">Example</a>
```go
package main

import (
	"fmt"
	"time"
)

// MyError is an error implementation that includes a time and message.
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
```

```txt
1989-03-15 22:30:00 +0000 UTC: the file system has gone away
```

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

As 会找到error链中第一个匹配到target的error, 那么它会让target指向该error的value, 然后返回true.

该链由error本身组成, 可通过反复调用 Unwrap 展开.

如果error的实际值能赋值给target指向的值或该error有`As(interface{}) bool`方法并返回true, 那么就认为参数err与target匹配.

如果target不是指向`实现了error方法的类型`或`interface{}`的非nil指针, As会panic. 如果参数err是nil, As 会返回false.


<a id="example_As">Example</a>
```go
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}

}
```

```txt
Failed at path: non-existing
```

## <a id="Is">func</a> [Is](https://golang.org/src/errors/wrap.go?s=837:868#L21)
<pre>func Is(err, target <a href="/pkg/builtin/#error">error</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
Is reports whether any error in err's chain matches target.

The chain consists of err itself followed by the sequence of errors obtained by
repeatedly calling Unwrap.

An error is considered to match a target if it is equal to that target or if
it implements a method Is(error) bool such that Is(target) returns true.

Is 判断 err的error链能否匹配到target.

该链由error本身组成, 可通过反复调用 Unwrap 展开.

如果一个error==target或它实现了`Is(error) bool`方法并返回true, 那么就认为参数err与target匹配.



## <a id="New">func</a> [New](https://golang.org/src/errors/errors.go?s=1875:1902#L48)
<pre>func New(text <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#error">error</a></pre>
New returns an error that formats as the given text.
Each call to New returns a distinct error value even if the text is identical.

New 返回包含给定text的error. 即使text相同, 每次调用都会返回不同的error.

<a id="example_New">Example</a>
```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
}
```

```txt
emit macho dwarf: elf header corrupted
```

<a id="example_New_errorf">Example (Errorf)</a>

The fmt package's Errorf function lets us use the package's formatting features to create descriptive error messages. 

fmt.Errorf可让我们根据fmt的format功能来创建带描述信息的error.
```go
package main

import (
	"fmt"
)

func main() {
	const name, id = "bimmler", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	if err != nil {
		fmt.Print(err)
	}
}
```

```txt
user "bimmler" (id 17) not found
```

## <a id="Unwrap">func</a> [Unwrap](https://golang.org/src/errors/wrap.go?s=372:400#L4)
<pre>func Unwrap(err <a href="/pkg/builtin/#error">error</a>) <a href="/pkg/builtin/#error">error</a></pre>
Unwrap returns the result of calling the Unwrap method on err, if err's
type contains an Unwrap method returning error.
Otherwise, Unwrap returns nil.

Unwrap 会调用err参数的 Unwrap 方法返回其结果, 如果该err有 Unwrap 方法会返回被封装的error, 否则返回nil.








