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

func TestContainsAny(t *testing.T) {
	t.Parallel()
	data := []struct {
		name  string
		s     any
		chars any
		exp   []any
	}{
		{
			name:  "true",
			s:     `"hello"`,
			chars: `"ol"`,
			exp:   []any{true, nil},
		},
		{
			name:  "false",
			s:     `"hello"`,
			chars: `"na"`,
			exp:   []any{false, nil},
		},
		{
			name:  "s must be a string",
			s:     0,
			chars: `"no"`,
			exp:   []any{false, util.NewError("s must be a string: 0")},
		},
		{
			name:  "chars must be a string",
			s:     `"hello"`,
			chars: 0,
			exp:   []any{false, util.NewError("chars must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.ContainsAny("strings.containsAny"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.containsAny")(%v, %v)`, d.s, d.chars)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestCount(t *testing.T) {
	t.Parallel()
	data := []struct {
		name   string
		s      any
		substr any
		exp    []any
	}{
		{
			name:   "normal",
			s:      `"hello"`,
			substr: `"l"`,
			exp:    []any{2.0, nil},
		},
		{
			name:   "0",
			s:      `"hello"`,
			substr: `"x"`,
			exp:    []any{0.0, nil},
		},
		{
			name:   "empty",
			s:      `"hello"`,
			substr: `""`,
			exp:    []any{6.0, nil},
		},
		{
			name:   "s must be a string",
			s:      0,
			substr: `"no"`,
			exp:    []any{0.0, util.NewError("s must be a string: 0")},
		},
		{
			name:   "substr must be a string",
			s:      `"hello"`,
			substr: 0,
			exp:    []any{0.0, util.NewError("substr must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.Count("strings.count"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.count")(%v, %v)`, d.s, d.substr)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestCut(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		s    any
		sep  any
		exp  []any
	}{
		{
			name: "normal",
			s:    `"hello,world"`,
			sep:  `","`,
			exp:  []any{"hello", "world", true, nil},
		},
		{
			name: "not found",
			s:    `"hello,world"`,
			sep:  `"x"`,
			exp:  []any{"hello,world", "", false, nil},
		},
		{
			name: "s must be a string",
			s:    0,
			sep:  `"no"`,
			exp:  []any{"", "", false, util.NewError("s must be a string: 0")},
		},
		{
			name: "sep must be a string",
			s:    `"hello"`,
			sep:  0,
			exp:  []any{"", "", false, util.NewError("sep must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.Cut("strings.cut"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.cut")(%v, %v)`, d.s, d.sep)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestCutPrefix(t *testing.T) {
	t.Parallel()
	data := []struct {
		name   string
		s      any
		prefix any
		exp    []any
	}{
		{
			name:   "normal",
			s:      `"foo/v1.0.0"`,
			prefix: `"foo/"`,
			exp:    []any{"v1.0.0", true, nil},
		},
		{
			name:   "not found",
			s:      `"hello,world"`,
			prefix: `"llo"`,
			exp:    []any{"hello,world", false, nil},
		},
		{
			name:   "s must be a string",
			s:      0,
			prefix: `"no"`,
			exp:    []any{"", false, util.NewError("s must be a string: 0")},
		},
		{
			name:   "prefix must be a string",
			s:      `"hello"`,
			prefix: 0,
			exp:    []any{"", false, util.NewError("prefix must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.CutPrefix("strings.cutPrefix"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.cutPrefix")(%v, %v)`, d.s, d.prefix)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestCutSuffix(t *testing.T) {
	t.Parallel()
	data := []struct {
		name   string
		s      any
		suffix any
		exp    []any
	}{
		{
			name:   "normal",
			s:      `"foo/v1.0.0"`,
			suffix: `"/v1.0.0"`,
			exp:    []any{"foo", true, nil},
		},
		{
			name:   "not found",
			s:      `"hello,world"`,
			suffix: `"llo"`,
			exp:    []any{"hello,world", false, nil},
		},
		{
			name:   "s must be a string",
			s:      0,
			suffix: `"no"`,
			exp:    []any{"", false, util.NewError("s must be a string: 0")},
		},
		{
			name:   "suffix must be a string",
			s:      `"hello"`,
			suffix: 0,
			exp:    []any{"", false, util.NewError("suffix must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.CutSuffix("strings.cutSuffix"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.cutSuffix")(%v, %v)`, d.s, d.suffix)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestEqualFold(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		s    any
		t    any
		exp  []any
	}{
		{
			name: "true",
			s:    `"Go"`,
			t:    `"go"`,
			exp:  []any{true, nil},
		},
		{
			name: "false",
			s:    `"ß"`,
			t:    `"ss"`,
			exp:  []any{false, nil},
		},
		{
			name: "s must be a string",
			s:    0,
			t:    `"no"`,
			exp:  []any{false, util.NewError("s must be a string: 0")},
		},
		{
			name: "t must be a string",
			s:    `"hello"`,
			t:    0,
			exp:  []any{false, util.NewError("t must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.EqualFold("strings.equalFold"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.equalFold")(%v, %v)`, d.s, d.t)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestFields(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		s    any
		exp  []any
	}{
		{
			name: "normal",
			s:    `"  foo bar  baz   "`,
			exp:  []any{[]any{"foo", "bar", "baz"}, nil},
		},
		{
			name: "s must be a string",
			s:    0,
			exp:  []any{nil, util.NewError("s must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.Fields("strings.fields"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.fields")(%v)`, d.s)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestLastIndex(t *testing.T) {
	t.Parallel()
	data := []struct {
		name   string
		s      any
		substr any
		exp    []any
	}{
		{
			name:   "normal",
			s:      `"go gopher"`,
			substr: `"go"`,
			exp:    []any{3.0, nil},
		},
		{
			name:   "0",
			s:      `"hello"`,
			substr: `"x"`,
			exp:    []any{-1.0, nil},
		},
		{
			name:   "s must be a string",
			s:      0,
			substr: `"no"`,
			exp:    []any{0.0, util.NewError("s must be a string: 0")},
		},
		{
			name:   "substr must be a string",
			s:      `"hello"`,
			substr: 0,
			exp:    []any{0.0, util.NewError("substr must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.LastIndex("strings.lastIndex"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.lastIndex")(%v, %v)`, d.s, d.substr)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestLastIndexAny(t *testing.T) {
	t.Parallel()
	data := []struct {
		name  string
		s     any
		chars any
		exp   []any
	}{
		{
			name:  "normal",
			s:     `"go gopher"`,
			chars: `"go"`,
			exp:   []any{4.0, nil},
		},
		{
			name:  "0",
			s:     `"hello"`,
			chars: `"x"`,
			exp:   []any{-1.0, nil},
		},
		{
			name:  "s must be a string",
			s:     0,
			chars: `"no"`,
			exp:   []any{0.0, util.NewError("s must be a string: 0")},
		},
		{
			name:  "chars must be a string",
			s:     `"hello"`,
			chars: 0,
			exp:   []any{0.0, util.NewError("chars must be a string: 0")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.LastIndexAny("strings.lastIndexAny"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.lastIndexAny")(%v, %v)`, d.s, d.chars)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestRepeat(t *testing.T) {
	t.Parallel()
	data := []struct {
		name  string
		s     any
		count any
		exp   []any
	}{
		{
			name:  "normal",
			s:     `"na"`,
			count: 2,
			exp:   []any{"nana", nil},
		},
		{
			name:  "s must be a string",
			s:     0,
			count: 2,
			exp:   []any{"", util.NewError("s must be a string: 0")},
		},
		{
			name:  "count must be an int",
			s:     `"hello"`,
			count: `"foo"`,
			exp:   []any{"", util.NewError("count is invalid: the value must be an int: foo")},
		},
		{
			name:  "count must not be a float64",
			s:     `"na"`,
			count: 2.2,
			exp:   []any{"", util.NewError("count is invalid: the value must not be a float64: 2.2")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.Repeat("strings.repeat"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.repeat")(%v, %v)`, d.s, d.count)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}

func TestReplace(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		s    any
		old  any
		new  any
		n    any
		exp  []any
	}{
		{
			name: "normal",
			s:    `"oink oink oink"`,
			old:  `"k"`,
			new:  `"ky"`,
			n:    2,
			exp:  []any{"oinky oinky oink", nil},
		},
		{
			name: "s must be a string",
			s:    0,
			old:  `"k"`,
			new:  `"ky"`,
			n:    2,
			exp:  []any{"", util.NewError("s must be a string: 0")},
		},
		{
			name: "old must be a string",
			s:    `"oink oink oink"`,
			old:  0,
			new:  `"ky"`,
			n:    2,
			exp:  []any{"", util.NewError("old must be a string: 0")},
		},
		{
			name: "new must be a string",
			s:    `"oink oink oink"`,
			old:  `"k"`,
			new:  0,
			n:    2,
			exp:  []any{"", util.NewError("new must be a string: 0")},
		},
		{
			name: "n must be an int",
			s:    `"oink oink oink"`,
			old:  `"k"`,
			new:  `"ky"`,
			n:    `"foo"`,
			exp:  []any{"", util.NewError("n is invalid: the value must be an int: foo")},
		},
		{
			name: "n must not be a float64",
			s:    `"oink oink oink"`,
			old:  `"k"`,
			new:  `"ky"`,
			n:    2.2,
			exp:  []any{"", util.NewError("n is invalid: the value must not be a float64: 2.2")},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.Replace("strings.replace"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.replace")(%v, %v, %v, %v)`, d.s, d.old, d.new, d.n)
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
