package nativefunc

import (
	"regexp"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func Match() *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   "match",
		Params: ast.Identifiers{"pattern", "s"},
		Func: func(s []interface{}) (interface{}, error) {
			return regexp.MatchString(s[0].(string), s[1].(string))
		},
	}
}
