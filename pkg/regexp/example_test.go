package regexp_test

import (
	"encoding/json"
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
	var a []any
	if err := json.Unmarshal([]byte(result), &a); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v %v\n", a[0], a[1])
	// Output: true <nil>
}
