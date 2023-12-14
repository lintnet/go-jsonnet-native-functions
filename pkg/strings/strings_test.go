package strings_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-jsonnet"
	nstrings "github.com/suzuki-shunsuke/go-jsonnet-native-functions/pkg/strings"
)

func TestContains(t *testing.T) { //nolint:dupl
	t.Parallel()
	data := []struct {
		name   string
		s      string
		substr string
		exp    string
	}{
		{
			name:   "true",
			s:      "hello",
			substr: "el",
			exp:    "true",
		},
		{
			name:   "false",
			s:      "hello",
			substr: "no",
			exp:    "false",
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.Contains("strings.contains"))
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.contains")("%s", "%s")`, d.s, d.substr)
			result, err := vm.EvaluateAnonymousSnippet("test.jsonnet", code)
			if err != nil {
				t.Fatal(err)
			}
			trimmedResult := strings.TrimSpace(result)
			if trimmedResult != d.exp {
				t.Fatalf(`wanted "%s", got "%s"`, d.exp, trimmedResult)
			}
		})
	}
}

func TestTrimPrefix(t *testing.T) { //nolint:dupl
	t.Parallel()
	data := []struct {
		name   string
		s      string
		prefix string
		exp    string
	}{
		{
			name:   "true",
			s:      "foo/v1.0.0",
			prefix: "foo/",
			exp:    `"v1.0.0"`,
		},
		{
			name:   "false",
			s:      "foo/v1.0.0",
			prefix: "bar/",
			exp:    `"foo/v1.0.0"`,
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.TrimPrefix("strings.trimPrefix"))
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.trimPrefix")("%s", "%s")`, d.s, d.prefix)
			result, err := vm.EvaluateAnonymousSnippet("test.jsonnet", code)
			if err != nil {
				t.Fatal(err)
			}
			trimmedResult := strings.TrimSpace(result)
			if trimmedResult != d.exp {
				t.Fatalf(`wanted "%s", got "%s"`, d.exp, trimmedResult)
			}
		})
	}
}

func TestTrimSpace(t *testing.T) {
	t.Parallel()
	data := []struct {
		name   string
		s      string
		substr string
		exp    string
	}{
		{
			name: "normal",
			s:    " hello  ",
			exp: `"hello"
`,
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nstrings.TrimSpace("strings.trimSpace"))
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("strings.trimSpace")("%s")`, d.s)
			result, err := vm.EvaluateAnonymousSnippet("test.jsonnet", code)
			if err != nil {
				t.Fatal(err)
			}
			if result != d.exp {
				t.Fatalf(`wanted "%s", got "%s"`, d.exp, result)
			}
		})
	}
}
