package url_test

import (
	"fmt"
	"testing"

	"github.com/google/go-jsonnet"
	"github.com/lintnet/go-jsonnet-native-functions/pkg/net/url"
	"github.com/lintnet/go-jsonnet-native-functions/testutil"
	"github.com/lintnet/go-jsonnet-native-functions/util"
)

func TestParse(t *testing.T) {
	t.Parallel()
	data := []struct {
		name   string
		rawURL any
		exp    []any
	}{
		{
			name:   "true",
			rawURL: `"http://example.com/foo/bar?lang=en&tag=go#top"`,
			exp: []any{map[string]any{
				"Scheme":      "http",
				"Opaque":      "",
				"Host":        "example.com",
				"Path":        "/foo/bar",
				"RawPath":     "",
				"OmitHost":    false,
				"ForceQuery":  false,
				"RawQuery":    "lang=en&tag=go",
				"Fragment":    "top",
				"RawFragment": "",
				"Query": map[string]any{
					"lang": []any{"en"},
					"tag":  []any{"go"},
				},
			}, nil},
		},
		{
			name:   "rawURL must be a string",
			rawURL: 0,
			exp:    []any{nil, util.NewError("rawURL must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(url.Parse("url.Parse"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("url.Parse")(%v)`, d.rawURL)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}
