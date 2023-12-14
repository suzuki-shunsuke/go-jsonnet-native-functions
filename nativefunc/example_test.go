package nativefunc_test

import (
	"fmt"
	"log"

	"github.com/google/go-jsonnet"
	"github.com/suzuki-shunsuke/go-jsonnet-native-functions/nativefunc"
)

func Example() {
	vm := jsonnet.MakeVM()
	nativefunc.SetAll(vm)
	code := `std.native("strings.trimPrefix")("foo/v1.0.0", "foo/")`
	result, err := vm.EvaluateAnonymousSnippet("test.jsonnet", code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	// Output: "v1.0.0"
}
