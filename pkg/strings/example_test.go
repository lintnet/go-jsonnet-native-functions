package strings_test

import (
	"fmt"
	"log"

	"github.com/google/go-jsonnet"
	"github.com/lintnet/go-jsonnet-native-functions/pkg/strings"
)

func Example() {
	vm := jsonnet.MakeVM()
	vm.NativeFunction(strings.TrimPrefix("trimPrefix"))
	code := `std.native("trimPrefix")("foo/v1.0.0", "foo/")`
	result, err := vm.EvaluateAnonymousSnippet("test.jsonnet", code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	// Output: "v1.0.0"
}
