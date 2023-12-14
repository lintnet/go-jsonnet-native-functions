package nativefunc

import (
	"fmt"
	"regexp"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func Match() *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   "regexp.match",
		Params: ast.Identifiers{"pattern", "s"},
		Func: func(s []interface{}) (interface{}, error) {
			pattern, ok := s[0].(string)
			if !ok {
				return nil, fmt.Errorf("pattern must be a string: %v", s[0])
			}
			s1, ok := s[1].(string)
			if !ok {
				return nil, fmt.Errorf("s must be a string: %v", s[1])
			}
			return regexp.MatchString(pattern, s1) //nolint:wrapcheck
		},
	}
}
