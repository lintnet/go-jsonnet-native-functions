# go-jsonnet-native-functions

[![PkgGoDev](https://pkg.go.dev/badge/github.com/lintnet/go-jsonnet-native-functions)](https://pkg.go.dev/github.com/lintnet/go-jsonnet-native-functions)

Go package porting several Go's Standard libraries to [go-jsonnet](https://github.com/google/go-jsonnet)'s [native functions](https://pkg.go.dev/github.com/google/go-jsonnet#NativeFunction)

## Document

https://pkg.go.dev/github.com/lintnet/go-jsonnet-native-functions

## How to use

This package has sub packages `pkg/{Stanard library package path}`, which port Standard libraries `{Standard library package path}`.
For example, [github.com/lintnet/go-jsonnet-native-functions/pkg/path/filepath](https://pkg.go.dev/github.com/lintnet/go-jsonnet-native-functions/pkg/path/filepath) ports [path/filepath](https://pkg.go.dev/path/filepath).
Each Package has generator functions to generate native functions that port same name's standard library's functions.
For example, [github.com/lintnet/go-jsonnet-native-functions/pkg/path/filepath](https://pkg.go.dev/github.com/lintnet/go-jsonnet-native-functions/pkg/path/filepath#Base) ports [path/filepath#Base](https://pkg.go.dev/path/filepath#Base).
Generator functions' signatures are consistent.

```go
func(name string) *jsonnet.NativeFunction
```

The argument `name` is a generated native function's name.

e.g.

```go
package main

import (
	"fmt"
	"log"

	"github.com/lintnet/go-jsonnet" // Fork google/go-jsonnet
	"github.com/lintnet/go-jsonnet-native-functions/pkg/strings"
)

func main() {
	vm := jsonnet.MakeVM()
	vm.NativeFunction(strings.TrimPrefix("trimPrefix"))
	code := `std.native("trimPrefix")("foo/v1.0.0", "foo/")`
	result, err := vm.EvaluateAnonymousSnippet("test.jsonnet", code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result) // ["v1.0.0", null]
}
```

## API design

Basically, native functions of this library don't raise errors.
Instead, they return an array to return multiple values and an error same as Go's functions.

For example, `strings.TrimPrefix` returns an array whose first element is a trimmed string and second element is an object representing an error.

If there is no error, the error object is `null`.

```go
vm.NativeFunction(strings.TrimPrefix("trimPrefix"))

// `["v1.0.0", null]`, nil
result, err := vm.EvaluateAnonymousSnippet("test.jsonnet", `std.native("trimPrefix")("foo/v1.0.0", "foo/")`)
```

If there is an error, the error object isn't `null`.

```go
vm.NativeFunction(strings.TrimPrefix("trimPrefix"))

// `["", {"message": "substr must be a string: true"}]`, nil
result2, err2 := vm.EvaluateAnonymousSnippet("test.jsonnet", `std.native("trimPrefix")("foo/v1.0.0", true)`)
```

An error object has a string field `message`.

```json
{"message": "substr must be a string: true"}
```

Even if Go functions don't return an error, native function's return value has an error object if argument types are invalid.
For example, arguments of `strings.TrimPrefix` must be strings but non string objects can be passed to the function in Jsonnet.

```go
result2, err2 := vm.EvaluateAnonymousSnippet("test.jsonnet", `std.native("trimPrefix")("foo/v1.0.0", true)`)
```

In that case, the return value has an error object.

```json
[
  "",
  {
    "message": "substr must be a string: true"
  }
]
```

### Why don't these functions raise an error?

Jsonnet doesn't support catching errors. [google/jsonnet#415](https://github.com/google/jsonnet/issues/415)
So if these functions raise an error, there is no way to handle the error in Jsonnet.
We think we should be able to handle errors ourselves, so we decided not to raise an error.

### Why do these functions return an array?

Jsonnet functions can't return multiple values, but Go functions can.
So if we want to return multiple values in Jsonnet, we need to embed them into a single object somehow.
There are several ways to achieve it, but we think returning an array is the most simplest way.

The exception is functions return not an array but a single error object if they need to return only an error.

## :warning: Some functions don't work with google/go-jsonnet

Some functions don't work with google/go-jsonnet due to bugs of google/go-jsonnet.
So please use [lintnet/go-jsonnet](https://github.com/lintnet/go-jsonnet) instead.
lintnet/go-jsonnet is a fork of google/go-jsonnet.
lintent/go-jsonnet is compatible with google/go-jsonnet, but includes some bug fixes.
Please see [lintnet/go-jsonnet#2](https://github.com/lintnet/go-jsonnet/issues/2).

## Does this library port third party Go libraries?

No. This library ports only Go's standard library.
This is because we don't want to depend on third party libraries other than github.com/google/go-jsonnet.
If we want to port third party Go libraries, we would create other repositories.

## LICENSE

[MIT](LICENSE)
