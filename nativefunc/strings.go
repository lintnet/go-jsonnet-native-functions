package nativefunc

import (
	"strings"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func Contains() *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   "contains",
		Params: ast.Identifiers{"s", "substr"},
		Func: func(s []interface{}) (interface{}, error) {
			return strings.Contains(s[0].(string), s[1].(string)), nil
		},
	}
}

func TrimSpace() *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   "trimSpace",
		Params: ast.Identifiers{"s"},
		Func: func(s []interface{}) (interface{}, error) {
			return strings.TrimSpace(s[0].(string)), nil
		},
	}
}

func TrimPrefix() *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   "trimPrefix",
		Params: ast.Identifiers{"s", "prefix"},
		Func: func(s []interface{}) (interface{}, error) {
			return strings.TrimPrefix(s[0].(string), s[1].(string)), nil
		},
	}
}
