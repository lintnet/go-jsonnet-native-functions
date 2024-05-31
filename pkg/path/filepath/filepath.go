package filepath

import (
	"fmt"
	"path/filepath"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func Base(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"path"},
		Func: func(s []interface{}) (interface{}, error) {
			path, ok := s[0].(string)
			if !ok {
				return []any{"", map[string]any{
					"message": fmt.Sprintf("path must be a string: %v", s[0]),
				}}, nil
			}
			return []any{filepath.Base(path), nil}, nil
		},
	}
}
