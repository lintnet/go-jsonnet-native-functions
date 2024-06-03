package path

import (
	"path"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
	"github.com/lintnet/go-jsonnet-native-functions/util"
)

func Base(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"path"},
		Func: func(s []any) (any, error) {
			p, ok := s[0].(string)
			if !ok {
				return []any{
					"", util.NewErrorf("path must be a string: %v", s[0]),
				}, nil
			}
			return []any{path.Base(p), nil}, nil
		},
	}
}

func Clean(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"path"},
		Func: func(s []any) (any, error) {
			p, ok := s[0].(string)
			if !ok {
				return []any{
					"", util.NewErrorf("path must be a string: %v", s[0]),
				}, nil
			}
			return []any{path.Clean(p), nil}, nil
		},
	}
}

func Dir(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"path"},
		Func: func(s []any) (any, error) {
			p, ok := s[0].(string)
			if !ok {
				return []any{
					"", util.NewErrorf("path must be a string: %v", s[0]),
				}, nil
			}
			return []any{path.Dir(p), nil}, nil
		},
	}
}

func Ext(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"path"},
		Func: func(s []any) (any, error) {
			p, ok := s[0].(string)
			if !ok {
				return []any{
					"", util.NewErrorf("path must be a string: %v", s[0]),
				}, nil
			}
			return []any{path.Ext(p), nil}, nil
		},
	}
}

func IsAbs(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"path"},
		Func: func(s []any) (any, error) {
			p, ok := s[0].(string)
			if !ok {
				return []any{
					false, util.NewErrorf("path must be a string: %v", s[0]),
				}, nil
			}
			return []any{path.IsAbs(p), nil}, nil
		},
	}
}

func Match(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"pattern", "name"},
		Func: func(s []any) (any, error) {
			pattern, ok := s[0].(string)
			if !ok {
				return []any{
					false, util.NewErrorf("pattern must be a string: %v", s[0]),
				}, nil
			}
			name, ok := s[1].(string)
			if !ok {
				return []any{
					false, util.NewErrorf("name must be a string: %v", s[1]),
				}, nil
			}
			matched, err := path.Match(pattern, name)
			if err != nil {
				return []any{
					false, util.NewError(err.Error()),
				}, nil
			}
			return []any{matched, nil}, nil
		},
	}
}

func Split(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"path"},
		Func: func(s []any) (any, error) {
			p, ok := s[0].(string)
			if !ok {
				return []any{
					"", "", util.NewErrorf("path must be a string: %v", s[0]),
				}, nil
			}
			dir, file := path.Split(p)
			return []any{dir, file, nil}, nil
		},
	}
}
