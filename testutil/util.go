package testutil

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/google/go-jsonnet"
)

func Check(t *testing.T, vm *jsonnet.VM, code string, exp []any) {
	t.Helper()
	result, err := vm.EvaluateAnonymousSnippet("test.jsonnet", code)
	if err != nil {
		t.Fatal(err)
	}
	var a []any
	if err := json.Unmarshal([]byte(result), &a); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(a, exp) {
		t.Fatalf(`wanted "%v", got "%v"`, exp, a)
	}
}
