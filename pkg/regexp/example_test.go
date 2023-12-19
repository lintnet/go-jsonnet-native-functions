package regexp_test

import (
	"fmt"
	"log"

	"github.com/google/go-jsonnet"
	"github.com/lintnet/go-jsonnet-native-functions/pkg/regexp"
)

func Example() {
	vm := jsonnet.MakeVM()
	vm.NativeFunction(regexp.MatchString("regexp.match"))
	code := `std.native("regexp.match")("l{2}", "hello")`
	result, err := vm.EvaluateAnonymousSnippet("test.jsonnet", code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	// Output: true
}
