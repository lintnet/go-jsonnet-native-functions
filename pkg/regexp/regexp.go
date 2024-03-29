package regexp

import (
	"fmt"
	"regexp"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func MatchString(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
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
			return regexp.MatchString(pattern, s1)
		},
	}
}
