package regexp_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-jsonnet"
	"github.com/lintnet/go-jsonnet-native-functions/pkg/regexp"
)

func TestMatchString(t *testing.T) {
	t.Parallel()
	data := []struct {
		name    string
		pattern string
		s       string
		exp     string
		isErr   bool
	}{
		{
			name:    "true",
			pattern: "l{2}",
			s:       "hello",
			exp:     "true",
		},
		{
			name:    "false",
			pattern: "a{2}",
			s:       "hello",
			exp:     "false",
		},
		{
			name:    "invalid regular expression",
			pattern: "*",
			s:       "hello",
			isErr:   true,
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(regexp.MatchString("regexp.match"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("regexp.match")("%s", "%s")`, d.pattern, d.s)
			result, err := vm.EvaluateAnonymousSnippet("test.jsonnet", code)
			if err != nil {
				if d.isErr {
					return
				}
				t.Fatal(err)
			}
			if d.isErr {
				t.Fatal("error must be returned")
			}
			trimmedResult := strings.TrimSpace(result)
			if trimmedResult != d.exp {
				t.Fatalf(`wanted "%s", got "%s"`, d.exp, trimmedResult)
			}
		})
	}
}

func TestMatchStringByArray(t *testing.T) {
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
			exp:     []any{false, "error parsing regexp: missing argument to repetition operator: `*`"},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(regexp.MatchStringReturnArray("regexp.match"))
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("regexp.match")("%s", "%s")`, d.pattern, d.s)
			result, err := vm.EvaluateAnonymousSnippet("test.jsonnet", code)
			if err != nil {
				if d.isErr {
					return
				}
				t.Fatal(err)
			}
			if d.isErr {
				t.Fatal("error must be returned")
			}
			var a any
			if err := json.Unmarshal([]byte(result), &a); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(a, d.exp) {
				t.Fatalf(`wanted "%v", got "%v"`, d.exp, a)
			}
		})
	}
}
