package nativefunc

import (
	"errors"
	"fmt"

	"github.com/google/go-jsonnet"
)

var ErrUnknownFunction = errors.New("unknown function")

var funcList = []func() *jsonnet.NativeFunction{ //nolint:gochecknoglobals
	regexMatch,
	contains,
	trimPrefix,
	trimSpace,
}

// SetAll sets all Native Functions of this package to VM.
func SetAll(vm *jsonnet.VM) {
	for _, fc := range funcList {
		vm.NativeFunction(fc())
	}
}

// Set sets given Native Functions to VM.
// names is a list of Native Function names such as "strings.contains".
func Set(vm *jsonnet.VM, names ...string) error {
	allFuncs := make(map[string]*jsonnet.NativeFunction, len(funcList))
	for _, fc := range funcList {
		nf := fc()
		allFuncs[nf.Name] = nf
	}
	for _, name := range names {
		f, ok := allFuncs[name]
		if !ok {
			return fmt.Errorf("%w: %s", ErrUnknownFunction, name)
		}
		vm.NativeFunction(f)
	}
	return nil
}
