package strings_test

import (
	"fmt"
	"testing"

	"github.com/google/go-jsonnet"
	nstrings "github.com/lintnet/go-jsonnet-native-functions/pkg/strings"
	"github.com/lintnet/go-jsonnet-native-functions/testutil"
)

func TestContains(t *testing.T) {
	t.Parallel()
	data := []struct {
		name   string
		s      string
		substr string
		exp    []any
	}{
		{
			name:   "true",
			s:      "hello",
			substr: "el",
			exp:    []any{true, nil},
		},
		{
			name:   "false",
			s:      "hello",
			substr: "no",
			exp:    []any{false, nil},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.Contains("strings.contains"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.contains")("%s", "%s")`, d.s, d.substr)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestTrimPrefix(t *testing.T) {
	t.Parallel()
	data := []struct {
		name   string
		s      string
		prefix string
		exp    []any
	}{
		{
			name:   "true",
			s:      "foo/v1.0.0",
			prefix: "foo/",
			exp:    []any{"v1.0.0", nil},
		},
		{
			name:   "false",
			s:      "foo/v1.0.0",
			prefix: "bar/",
			exp:    []any{"foo/v1.0.0", nil},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.TrimPrefix("strings.trimPrefix"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.trimPrefix")("%s", "%s")`, d.s, d.prefix)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestTrimSpace(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		s    string
		exp  []any
	}{
		{
			name: "normal",
			s:    " hello  ",
			exp:  []any{"hello", nil},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.TrimSpace("strings.trimSpace"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.trimSpace")("%s")`, d.s)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}
