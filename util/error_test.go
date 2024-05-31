package util_test

import (
	"reflect"
	"testing"

	"github.com/lintnet/go-jsonnet-native-functions/util"
)

func TestNewError(t *testing.T) {
	t.Parallel()
	act := util.NewError("test error")
	exp := map[string]any{"message": "test error"}
	if !reflect.DeepEqual(act, exp) {
		t.Fatalf("wanted %v, got %v ", exp, act)
	}
}

func TestNewErrorf(t *testing.T) {
	t.Parallel()
	act := util.NewErrorf("s must be a string: %v", 0)
	exp := map[string]any{"message": "s must be a string: 0"}
	if !reflect.DeepEqual(act, exp) {
		t.Fatalf("wanted %v, got %v ", exp, act)
	}
}
