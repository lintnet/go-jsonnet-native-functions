package strings_test

import (
	"fmt"
	"testing"

	"github.com/google/go-jsonnet"
	nstrings "github.com/lintnet/go-jsonnet-native-functions/pkg/strings"
	"github.com/lintnet/go-jsonnet-native-functions/testutil"
	"github.com/lintnet/go-jsonnet-native-functions/util"
)

func TestContains(t *testing.T) {
	t.Parallel()
	data := []struct {
		name   string
		s      any
		substr any
		exp    []any
	}{
		{
			name:   "true",
			s:      `"hello"`,
			substr: `"el"`,
			exp:    []any{true, nil},
		},
		{
			name:   "false",
			s:      `"hello"`,
			substr: `"no"`,
			exp:    []any{false, nil},
		},
		{
			name:   "s must be a string",
			s:      0,
			substr: `"no"`,
			exp:    []any{false, util.NewError("s must be a string: 0")},
		},
		{
			name:   "substr must be a string",
			s:      `"hello"`,
			substr: 0,
			exp:    []any{false, util.NewError("substr must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.Contains("strings.contains"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.contains")(%v, %v)`, d.s, d.substr)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestTrimPrefix(t *testing.T) {
	t.Parallel()
	data := []struct {
		name   string
		s      any
		prefix any
		exp    []any
	}{
		{
			name:   "true",
			s:      `"foo/v1.0.0"`,
			prefix: `"foo/"`,
			exp:    []any{"v1.0.0", nil},
		},
		{
			name:   "false",
			s:      `"foo/v1.0.0"`,
			prefix: `"bar/"`,
			exp:    []any{"foo/v1.0.0", nil},
		},
		{
			name:   "s must be a string",
			s:      0,
			prefix: `"bar/"`,
			exp:    []any{"", util.NewError("s must be a string: 0")},
		},
		{
			name:   "prefix must be a string",
			s:      `"foo/v1.0.0"`,
			prefix: 0,
			exp:    []any{"", util.NewError("prefix must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.TrimPrefix("strings.trimPrefix"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.trimPrefix")(%v, %v)`, d.s, d.prefix)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestTrimSpace(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		s    any
		exp  []any
	}{
		{
			name: "normal",
			s:    `" hello  "`,
			exp:  []any{"hello", nil},
		},
		{
			name: "normal",
			s:    0,
			exp:  []any{"", util.NewError("s must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.TrimSpace("strings.trimSpace"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.trimSpace")(%v)`, d.s)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}
