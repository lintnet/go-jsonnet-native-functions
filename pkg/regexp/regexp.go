package regexp

import (
	"regexp"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
	"github.com/lintnet/go-jsonnet-native-functions/util"
)

func MatchString(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"pattern", "s"},
		Func: func(s []any) (any, error) {
			pattern, ok := s[0].(string)
			if !ok {
				return []any{
					false, util.NewErrorf("pattern must be a string: %v", s[0]),
				}, nil
			}
			s1, ok := s[1].(string)
			if !ok {
				return []any{
					false, util.NewErrorf("s must be a string: %v", s[1]),
				}, nil
			}
			a, err := regexp.MatchString(pattern, s1)
			if err != nil {
				return []any{
					a, util.NewError(err.Error()),
				}, nil
			}
			return []any{
				a, nil,
			}, nil
		},
	}
}
