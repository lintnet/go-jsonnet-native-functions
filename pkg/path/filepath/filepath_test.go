package filepath_test

import (
	"fmt"
	"testing"

	"github.com/google/go-jsonnet"
	"github.com/lintnet/go-jsonnet-native-functions/pkg/path/filepath"
	"github.com/lintnet/go-jsonnet-native-functions/testutil"
)

func TestBase(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		path string
		exp  []any
	}{
		{
			name: "true",
			path: "foo/hello.txt",
			exp:  []any{"hello.txt", nil},
		},
	}
	vm := jsonnet.MakeVM()
	vm.NativeFunction(filepath.Base("path/filepath.base"))

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("path/filepath.base")("%s")`, d.path)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}
