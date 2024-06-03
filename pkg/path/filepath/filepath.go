package filepath

import (
	"path/filepath"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
	"github.com/lintnet/go-jsonnet-native-functions/util"
)

func Base(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"path"},
		Func: func(s []any) (any, error) {
			path, ok := s[0].(string)
			if !ok {
				return []any{
					"", util.NewErrorf("path must be a string: %v", s[0]),
				}, nil
			}
			return []any{filepath.Base(path), nil}, nil
		},
	}
}
