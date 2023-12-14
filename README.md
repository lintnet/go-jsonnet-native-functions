# go-jsonnet-native-functions

[go-jsonnet](https://github.com/google/go-jsonnet)'s [Native functions](https://pkg.go.dev/github.com/google/go-jsonnet#NativeFunction)

## Motivation

Jsonnet's Standard Library lacks some useful functions, but go-jsonnet enables us to add functions by Native functions.
This library provides usefule Native functions.

## Examples

Please see test codes too.

```go
vm := jsonnet.MakeVM()
vm.NativeFunction(nativefunc.TrimPrefix())
code := `std.native("trimPrefix")("foo/v1.0.0", "foo")`
result, err := vm.EvaluateAnonymousSnippet("test.jsonnet", code)
```

## LICENSE

[MIT](LICENSE)
