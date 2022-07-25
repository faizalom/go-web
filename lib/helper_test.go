package lib_test

import (
	"fmt"

	"github.com/faizalom/go-web/lib"
)

func ExampleHashPassword() {
	fmt.Println(lib.TestPassword("olleh"))
	// Output: olleh
}
