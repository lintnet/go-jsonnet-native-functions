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
				return nil, fmt.Errorf("s must be a string: %v", s[0])
			}
			substr, ok := s[1].(string)
			if !ok {
				return nil, fmt.Errorf("substr must be a string: %v", s[1])
			}
			return strings.Contains(s0, substr), nil
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
				return nil, fmt.Errorf("s must be a string: %v", s[0])
			}
			return strings.TrimSpace(s0), nil
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
				return nil, fmt.Errorf("s must be a string: %v", s[0])
			}
			prefix, ok := s[1].(string)
			if !ok {
				return nil, fmt.Errorf("prefix must be a string: %v", s[1])
			}
			return strings.TrimPrefix(s0, prefix), nil
		},
	}
}
