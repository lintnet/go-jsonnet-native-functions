package regexp_test

import (
	"fmt"
	"testing"

	"github.com/google/go-jsonnet"
	"github.com/lintnet/go-jsonnet-native-functions/pkg/regexp"
	"github.com/lintnet/go-jsonnet-native-functions/testutil"
	"github.com/lintnet/go-jsonnet-native-functions/util"
)

func TestMatchString(t *testing.T) {
	t.Parallel()
	data := []struct {
		name    string
		pattern any
		s       any
		exp     []any
		isErr   bool
	}{
		{
			name:    "true",
			pattern: `"l{2}"`,
			s:       `"hello"`,
			exp:     []any{true, nil},
		},
		{
			name:    "false",
			pattern: `"a{2}"`,
			s:       `"hello"`,
			exp:     []any{false, nil},
		},
		{
			name:    "invalid regular expression",
			pattern: `"*"`,
			s:       `"hello"`,
			exp: []any{
				false, util.NewError("error parsing regexp: missing argument to repetition operator: `*`"),
			},
		},
		{
			name:    "pattern must be a string",
			pattern: `0`,
			s:       `"hello"`,
			exp: []any{
				false, util.NewError("pattern must be a string: 0"),
			},
		},
		{
			name:    "s must be a string",
			pattern: `"a{2}"`,
			s:       `0`,
			exp: []any{
				false, util.NewError("s must be a string: 0"),
			},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(regexp.MatchString("regexp.match"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("regexp.match")(%v, %v)`, d.pattern, d.s)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}
