package nativefunc

import (
	"fmt"
	"path/filepath"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func filepathBase() *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   "path/filepath.base",
		Params: ast.Identifiers{"path"},
		Func: func(s []interface{}) (interface{}, error) {
			path, ok := s[0].(string)
			if !ok {
				return nil, fmt.Errorf("path must be a string: %v", s[0])
			}
			return filepath.Base(path), nil
		},
	}
}
