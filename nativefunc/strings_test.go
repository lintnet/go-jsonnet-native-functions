package nativefunc_test

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/google/go-jsonnet"
	"github.com/suzuki-shunsuke/go-jsonnet-native-functions/nativefunc"
)

func Example() {
	vm := jsonnet.MakeVM()
	vm.NativeFunction(nativefunc.TrimPrefix())
	code := `std.native("trimPrefix")("foo/v1.0.0", "foo/")`
	result, err := vm.EvaluateAnonymousSnippet("test.jsonnet", code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	// Output: "v1.0.0"
}

func TestContains(t *testing.T) {
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
	vm.NativeFunction(nativefunc.Contains())
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("contains")("%s", "%s")`, d.s, d.substr)
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

func TestTrimPrefix(t *testing.T) {
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
	vm.NativeFunction(nativefunc.TrimPrefix())
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("trimPrefix")("%s", "%s")`, d.s, d.prefix)
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
	vm.NativeFunction(nativefunc.TrimSpace())
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("trimSpace")("%s")`, d.s)
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
