package regexp

import (
	"fmt"
	"regexp"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func toErr(s string) map[string]any {
	return map[string]any{
		"message": s,
	}
}

func MatchString(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"pattern", "s"},
		Func: func(s []interface{}) (interface{}, error) {
			pattern, ok := s[0].(string)
			if !ok {
				return []any{
					false, toErr(fmt.Sprintf("pattern must be a string: %v", s[0])),
				}, nil
			}
			s1, ok := s[1].(string)
			if !ok {
				return []any{
					false, toErr(fmt.Sprintf("s must be a string: %v", s[1])),
				}, nil
			}
			a, err := regexp.MatchString(pattern, s1)
			if err != nil {
				return []any{
					a, toErr(err.Error()),
				}, nil
			}
			return []any{
				a, nil,
			}, nil
		},
	}
}
