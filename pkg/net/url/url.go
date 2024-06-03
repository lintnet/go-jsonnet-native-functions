package url

import (
	"net/url"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
	"github.com/lintnet/go-jsonnet-native-functions/util"
)

func Parse(name string) *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"rawURL"},
		Func: func(s []any) (any, error) {
			rawURL, ok := s[0].(string)
			if !ok {
				return []any{
					nil, util.NewErrorf("rawURL must be a string: %v", s[0]),
				}, nil
			}
			u, err := url.Parse(rawURL)
			if err != nil {
				return []any{
					nil, util.NewError(err.Error()),
				}, nil
			}
			q := u.Query()
			query := make(map[string]any, len(q))
			for k, v := range q {
				a := make([]any, len(v))
				for i, b := range v {
					a[i] = b
				}
				query[k] = a
			}
			return []any{map[string]any{
				"Scheme":      u.Scheme,
				"Opaque":      u.Opaque,
				"Host":        u.Host,
				"Path":        u.Path,
				"RawPath":     u.RawPath,
				"OmitHost":    u.OmitHost,
				"ForceQuery":  u.ForceQuery,
				"RawQuery":    u.RawQuery,
				"Fragment":    u.Fragment,
				"RawFragment": u.RawFragment,
				"Query":       query,
			}, nil}, nil
		},
	}
}
