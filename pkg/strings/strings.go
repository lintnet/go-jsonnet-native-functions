package strings

import (
	"fmt"
	"strings"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func Contains(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"s", "substr"},
		Func: func(s []interface{}) (interface{}, error) {
			s0, ok := s[0].(string)
			if !ok {
				return []any{
					false, map[string]any{
						"message": fmt.Sprintf("s must be a string: %v", s[0]),
					},
				}, nil
			}
			substr, ok := s[1].(string)
			if !ok {
				return []any{
					false, map[string]any{
						"message": fmt.Sprintf("substr must be a string: %v", s[1]),
					},
				}, nil
			}
			return []any{strings.Contains(s0, substr), nil}, nil
		},
	}
}

func TrimSpace(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"s"},
		Func: func(s []interface{}) (interface{}, error) {
			s0, ok := s[0].(string)
			if !ok {
				return []any{
					"", map[string]any{
						"message": fmt.Sprintf("s must be a string: %v", s[0]),
					},
				}, nil
			}
			return []any{strings.TrimSpace(s0), nil}, nil
		},
	}
}

func TrimPrefix(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"s", "prefix"},
		Func: func(s []interface{}) (interface{}, error) {
			s0, ok := s[0].(string)
			if !ok {
				return []any{
					"", map[string]any{
						"message": fmt.Sprintf("s must be a string: %v", s[0]),
					},
				}, nil
			}
			prefix, ok := s[1].(string)
			if !ok {
				return []any{
					"", map[string]any{
						"message": fmt.Sprintf("prefix must be a string: %v", s[1]),
					},
				}, nil
			}
			return []any{strings.TrimPrefix(s0, prefix), nil}, nil
		},
	}
}
