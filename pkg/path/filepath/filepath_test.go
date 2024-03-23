package filepath_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-jsonnet"
	"github.com/lintnet/go-jsonnet-native-functions/pkg/path/filepath"
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
	vm.NativeFunction(filepath.Base("path/filepath.base"))

	for _, d := range data {
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
