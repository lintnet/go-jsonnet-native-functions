package regexp

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func MatchString(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"pattern", "s", "opts"},
		Func: func(s []interface{}) (any, error) {
			noErr := false
			switch len(s) {
			case 0, 1:
				return nil, errors.New("pattern and s are required")
			case 2:
			case 3:
				opts, ok := s[2].(map[string]any)
				if !ok {
					return nil, errors.New("opts must be map[string]any")
				}
				if a, ok := opts["no_error"]; ok {
					b, ok := a.(bool)
					if !ok {
						return nil, errors.New("no_error must be bool")
					}
					noErr = b
				}
			default:
				return nil, errors.New("too many arguments")
			}
			pattern, ok := s[0].(string)
			if !ok {
				if noErr {
					return map[string]any{
						"error": fmt.Sprintf("pattern must be a string: %v", s[0]),
					}, nil
				}
				return nil, fmt.Errorf("pattern must be a string: %v", s[0])
			}
			s1, ok := s[1].(string)
			if !ok {
				if noErr {
					return map[string]any{
						"error": fmt.Sprintf("s must be a string: %v", s[0]),
					}, nil
				}
				return nil, fmt.Errorf("s must be a string: %v", s[1])
			}
			a, err := regexp.MatchString(pattern, s1)
			if err != nil {
				if noErr {
					return map[string]any{
						"error": err.Error(),
					}, nil
				}
				return a, err
			}
			if noErr {
				return map[string]any{
					"value": a,
				}, nil
			}
			return a, nil
		},
	}
}
