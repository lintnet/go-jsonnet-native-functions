package filepath_test

import (
	"fmt"
	"testing"

	"github.com/google/go-jsonnet"
	"github.com/lintnet/go-jsonnet-native-functions/pkg/path/filepath"
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
	vm.NativeFunction(filepath.Base("path/filepath.base"))

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			code := fmt.Sprintf(`std.native("path/filepath.base")(%v)`, d.path)
			testutil.Check(t, vm, code, d.exp)
		})
	}
}
