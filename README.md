# go-jsonnet-native-functions

[![PkgGoDev](https://pkg.go.dev/badge/github.com/suzuki-shunsuke/go-jsonnet-native-functions/nativefunc)](https://pkg.go.dev/github.com/suzuki-shunsuke/go-jsonnet-native-functions/nativefunc)

Go package porting several Go's Standard libraries to [go-jsonnet](https://github.com/google/go-jsonnet)'s [native functions](https://pkg.go.dev/github.com/google/go-jsonnet#NativeFunction)

## Document

https://pkg.go.dev/github.com/suzuki-shunsuke/go-jsonnet-native-functions

## How to use

This package has sub packages `pkg/{Stanard library package path}`, which port Standard libraries `<Standard library package path>`.
For example, [github.com/suzuki-shunsuke/go-jsonnet-native-functions/pkg/path/filepath](https://pkg.go.dev/github.com/suzuki-shunsuke/go-jsonnet-native-functions/pkg/path/filepath) ports [path/filepath](https://pkg.go.dev/path/filepath).
Each Package has generator functions to generate native functions that port same name's standard library's functions.
For example, [github.com/suzuki-shunsuke/go-jsonnet-native-functions/pkg/path/filepath](https://pkg.go.dev/github.com/suzuki-shunsuke/go-jsonnet-native-functions/pkg/path/filepath#Base) ports [path/filepath#Base](https://pkg.go.dev/path/filepath#Base).
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

	"github.com/google/go-jsonnet"
	"github.com/suzuki-shunsuke/go-jsonnet-native-functions/pkg/strings"
)

func main() {
	vm := jsonnet.MakeVM()
	vm.NativeFunction(strings.TrimPrefix("trimPrefix"))
	code := `std.native("trimPrefix")("foo/v1.0.0", "foo/")`
	result, err := vm.EvaluateAnonymousSnippet("test.jsonnet", code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	// Output: "v1.0.0"
}
```

## LICENSE

[MIT](LICENSE)
