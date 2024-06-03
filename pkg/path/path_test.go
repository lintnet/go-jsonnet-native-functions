package path_test

import (
	"fmt"
	"testing"

	"github.com/google/go-jsonnet"
	"github.com/lintnet/go-jsonnet-native-functions/pkg/path"
	"github.com/lintnet/go-jsonnet-native-functions/testutil"
	"github.com/lintnet/go-jsonnet-native-functions/util"
)

func TestBase(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		path any
		exp  []any
	}{
		{
			name: "normal",
			path: `"foo/hello.txt"`,
			exp:  []any{"hello.txt", nil},
		},
		{
			name: "path must be a string",
			path: 0,
			exp:  []any{"", util.NewError("path must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(path.Base("path.base"))

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("path.base")(%v)`, d.path)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestClean(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		path any
		exp  []any
	}{
		{
			name: "normal",
			path: `"foo/../hello.txt"`,
			exp:  []any{"hello.txt", nil},
		},
		{
			name: "path must be a string",
			path: 0,
			exp:  []any{"", util.NewError("path must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(path.Clean("path.clean"))

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("path.clean")(%v)`, d.path)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestDir(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		path any
		exp  []any
	}{
		{
			name: "normal",
			path: `"foo/bar/hello.txt"`,
			exp:  []any{"foo/bar", nil},
		},
		{
			name: "path must be a string",
			path: 0,
			exp:  []any{"", util.NewError("path must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(path.Dir("path.Dir"))

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("path.Dir")(%v)`, d.path)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestExt(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		path any
		exp  []any
	}{
		{
			name: "normal",
			path: `"foo/bar/hello.txt"`,
			exp:  []any{".txt", nil},
		},
		{
			name: "path must be a string",
			path: 0,
			exp:  []any{"", util.NewError("path must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(path.Ext("path.Ext"))

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("path.Ext")(%v)`, d.path)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestIsAbs(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		path any
		exp  []any
	}{
		{
			name: "true",
			path: `"/foo/bar/hello.txt"`,
			exp:  []any{true, nil},
		},
		{
			name: "false",
			path: `"foo/bar/hello.txt"`,
			exp:  []any{false, nil},
		},
		{
			name: "path must be a string",
			path: 0,
			exp:  []any{false, util.NewError("path must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(path.IsAbs("path.IsAbs"))

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("path.IsAbs")(%v)`, d.path)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestMatch(t *testing.T) {
	t.Parallel()
	data := []struct {
		name    string
		pattern any
		n       any
		exp     []any
	}{
		{
			name:    "true",
			pattern: `"a*"`,
			n:       `"abs"`,
			exp:     []any{true, nil},
		},
		{
			name:    "false",
			pattern: `"a*/b"`,
			n:       `"a/c/b"`,
			exp:     []any{false, nil},
		},
		{
			name:    "pattern must be a string",
			pattern: 0,
			n:       `"abs"`,
			exp:     []any{false, util.NewError("pattern must be a string: 0")},
		},
		{
			name:    "name must be a string",
			pattern: `"a*"`,
			n:       0,
			exp:     []any{false, util.NewError("name must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(path.Match("path.Match"))

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("path.Match")(%v, %v)`, d.pattern, d.n)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestSplit(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		path any
		exp  []any
	}{
		{
			name: "normal",
			path: `"static/myfile.css"`,
			exp:  []any{"static/", "myfile.css", nil},
		},
		{
			name: "path must be a string",
			path: 0,
			exp:  []any{"", "", util.NewError("path must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(path.Split("path.Split"))

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("path.Split")(%v)`, d.path)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}
