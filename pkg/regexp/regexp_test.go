package regexp_test

import (
	"fmt"
	"testing"

	"github.com/google/go-jsonnet"
	"github.com/lintnet/go-jsonnet-native-functions/pkg/regexp"
	"github.com/lintnet/go-jsonnet-native-functions/testutil"
)

func TestMatchString(t *testing.T) {
	t.Parallel()
	data := []struct {
		name    string
		pattern string
		s       string
		exp     []any
		isErr   bool
	}{
		{
			name:    "true",
			pattern: "l{2}",
			s:       "hello",
			exp:     []any{true, nil},
		},
		{
			name:    "false",
			pattern: "a{2}",
			s:       "hello",
			exp:     []any{false, nil},
		},
		{
			name:    "invalid regular expression",
			pattern: "*",
			s:       "hello",
			exp: []any{
				false, map[string]any{
					"message": "error parsing regexp: missing argument to repetition operator: `*`",
				},
			},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(regexp.MatchString("regexp.match"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("regexp.match")("%s", "%s")`, d.pattern, d.s)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}
