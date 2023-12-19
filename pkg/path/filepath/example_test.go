package filepath_test

import (
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
	fmt.Println(result)
	// Output: "hello.txt"
}
