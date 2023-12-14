package nativefunc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-jsonnet"
	"github.com/suzuki-shunsuke/go-jsonnet-native-functions/nativefunc"
)

func TestBase(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		path string
		exp  string
	}{
		{
			name: "true",
			path: "foo/hello.txt",
			exp:  `"hello.txt"`,
		},
	}
	vm := jsonnet.MakeVM()
	nativefunc.SetAll(vm)
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("path/filepath.base")("%s")`, d.path)
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
