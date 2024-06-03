package strings

import (
	"strings"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
	"github.com/lintnet/go-jsonnet-native-functions/util"
)

func Contains(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"s", "substr"},
		Func: func(s []any) (any, error) {
			s0, ok := s[0].(string)
			if !ok {
				return []any{
					false, util.NewErrorf("s must be a string: %v", s[0]),
				}, nil
			}
			substr, ok := s[1].(string)
			if !ok {
				return []any{
					false, util.NewErrorf("substr must be a string: %v", s[1]),
				}, nil
			}
			return []any{strings.Contains(s0, substr), nil}, nil
		},
	}
}

func ContainsAny(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"s", "chars"},
		Func: func(s []any) (any, error) {
			s0, ok := s[0].(string)
			if !ok {
				return []any{
					false, util.NewErrorf("s must be a string: %v", s[0]),
				}, nil
			}
			chars, ok := s[1].(string)
			if !ok {
				return []any{
					false, util.NewErrorf("chars must be a string: %v", s[1]),
				}, nil
			}
			return []any{strings.ContainsAny(s0, chars), nil}, nil
		},
	}
}

func Count(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"s", "substr"},
		Func: func(s []any) (any, error) {
			s0, ok := s[0].(string)
			if !ok {
				return []any{
					0, util.NewErrorf("s must be a string: %v", s[0]),
				}, nil
			}
			substr, ok := s[1].(string)
			if !ok {
				return []any{
					0, util.NewErrorf("substr must be a string: %v", s[1]),
				}, nil
			}
			return []any{strings.Count(s0, substr), nil}, nil
		},
	}
}

func Cut(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"s", "sep"},
		Func: func(s []any) (any, error) {
			s0, ok := s[0].(string)
			if !ok {
				return []any{
					"", "", false, util.NewErrorf("s must be a string: %v", s[0]),
				}, nil
			}
			sep, ok := s[1].(string)
			if !ok {
				return []any{
					"", "", false, util.NewErrorf("sep must be a string: %v", s[1]),
				}, nil
			}
			before, after, found := strings.Cut(s0, sep)
			return []any{before, after, found, nil}, nil
		},
	}
}

func CutPrefix(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"s", "prefix"},
		Func: func(s []any) (any, error) {
			s0, ok := s[0].(string)
			if !ok {
				return []any{
					"", false, util.NewErrorf("s must be a string: %v", s[0]),
				}, nil
			}
			prefix, ok := s[1].(string)
			if !ok {
				return []any{
					"", false, util.NewErrorf("prefix must be a string: %v", s[1]),
				}, nil
			}
			after, found := strings.CutPrefix(s0, prefix)
			return []any{after, found, nil}, nil
		},
	}
}

func CutSuffix(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"s", "suffix"},
		Func: func(s []any) (any, error) {
			s0, ok := s[0].(string)
			if !ok {
				return []any{
					"", false, util.NewErrorf("s must be a string: %v", s[0]),
				}, nil
			}
			suffix, ok := s[1].(string)
			if !ok {
				return []any{
					"", false, util.NewErrorf("suffix must be a string: %v", s[1]),
				}, nil
			}
			before, found := strings.CutSuffix(s0, suffix)
			return []any{before, found, nil}, nil
		},
	}
}

func EqualFold(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"s", "t"},
		Func: func(s []any) (any, error) {
			s0, ok := s[0].(string)
			if !ok {
				return []any{
					false, util.NewErrorf("s must be a string: %v", s[0]),
				}, nil
			}
			t, ok := s[1].(string)
			if !ok {
				return []any{
					false, util.NewErrorf("t must be a string: %v", s[1]),
				}, nil
			}
			return []any{strings.EqualFold(s0, t), nil}, nil
		},
	}
}

func Fields(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"s"},
		Func: func(s []any) (any, error) {
			s0, ok := s[0].(string)
			if !ok {
				return []any{
					nil, util.NewErrorf("s must be a string: %v", s[0]),
				}, nil
			}

			// We need to convert the []string to []any.
			// https://github.com/google/go-jsonnet/blob/fed90cd9cd733a87f9fb27cfb32a3e08a7695603/interpreter.go#L972
			a := strings.Fields(s0)
			b := make([]any, len(a))
			for i, v := range a {
				b[i] = v
			}

			return []any{b, nil}, nil
		},
	}
}

func LastIndex(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"s", "substr"},
		Func: func(s []any) (any, error) {
			s0, ok := s[0].(string)
			if !ok {
				return []any{
					0, util.NewErrorf("s must be a string: %v", s[0]),
				}, nil
			}
			substr, ok := s[1].(string)
			if !ok {
				return []any{
					0, util.NewErrorf("substr must be a string: %v", s[1]),
				}, nil
			}
			return []any{strings.LastIndex(s0, substr), nil}, nil
		},
	}
}

func LastIndexAny(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"s", "chars"},
		Func: func(s []any) (any, error) {
			s0, ok := s[0].(string)
			if !ok {
				return []any{
					0, util.NewErrorf("s must be a string: %v", s[0]),
				}, nil
			}
			chars, ok := s[1].(string)
			if !ok {
				return []any{
					0, util.NewErrorf("chars must be a string: %v", s[1]),
				}, nil
			}
			return []any{strings.LastIndexAny(s0, chars), nil}, nil
		},
	}
}

func Repeat(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"s", "count"},
		Func: func(s []any) (any, error) {
			s0, ok := s[0].(string)
			if !ok {
				return []any{
					"", util.NewErrorf("s must be a string: %v", s[0]),
				}, nil
			}
			count, err := util.ConvertToInt(s[1])
			if err != nil {
				return []any{ //nolint:nilerr
					"", util.NewError("count is invalid: " + err.Error()),
				}, nil
			}
			return []any{strings.Repeat(s0, count), nil}, nil
		},
	}
}

func Replace(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"s", "old", "new", "n"},
		Func: func(s []any) (any, error) {
			s0, ok := s[0].(string)
			if !ok {
				return []any{
					"", util.NewErrorf("s must be a string: %v", s[0]),
				}, nil
			}
			old, ok := s[1].(string)
			if !ok {
				return []any{
					"", util.NewErrorf("old must be a string: %v", s[1]),
				}, nil
			}
			newS, ok := s[2].(string)
			if !ok {
				return []any{
					"", util.NewErrorf("new must be a string: %v", s[2]),
				}, nil
			}
			n, err := util.ConvertToInt(s[3])
			if err != nil {
				return []any{ //nolint:nilerr
					"", util.NewError("n is invalid: " + err.Error()),
				}, nil
			}
			return []any{strings.Replace(s0, old, newS, n), nil}, nil
		},
	}
}

// Deprecated: Use std.trim(str) instead.
func TrimSpace(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"s"},
		Func: func(s []any) (any, error) {
			s0, ok := s[0].(string)
			if !ok {
				return []any{
					"", util.NewErrorf("s must be a string: %v", s[0]),
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
		Func: func(s []any) (any, error) {
			s0, ok := s[0].(string)
			if !ok {
				return []any{
					"", util.NewErrorf("s must be a string: %v", s[0]),
				}, nil
			}
			prefix, ok := s[1].(string)
			if !ok {
				return []any{
					"", util.NewErrorf("prefix must be a string: %v", s[1]),
				}, nil
			}
			return []any{strings.TrimPrefix(s0, prefix), nil}, nil
		},
	}
}
