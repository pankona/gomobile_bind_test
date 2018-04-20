package lib

import "C"
import (
	"fmt"
)

//export goCallback
func goCallback() {
	fmt.Println("callbacked!")
}
