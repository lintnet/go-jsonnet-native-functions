package regexp_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-jsonnet"
	"github.com/suzuki-shunsuke/go-jsonnet-native-functions/native/regexp"
)

func TestMatch(t *testing.T) {
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
		d := d
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
