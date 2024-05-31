package filepath_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/go-jsonnet"
	"github.com/lintnet/go-jsonnet-native-functions/pkg/path/filepath"
)

func Example() {
	vm := jsonnet.MakeVM()
	vm.NativeFunction(filepath.Base("filepath.base"))
	code := `std.native("filepath.base")("foo/hello.txt")`
	result, err := vm.EvaluateAnonymousSnippet("test.jsonnet", code)
	if err != nil {
		log.Fatal(err)
	}
	var a []any
	if err := json.Unmarshal([]byte(result), &a); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v %v\n", a[0], a[1])
	// Output: hello.txt <nil>
}
