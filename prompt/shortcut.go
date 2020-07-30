package prompt

import (
	"fmt"
)

func print(x interface{}) {
	fmt.Printf("%v\r\n", x)
}

var printf = fmt.Printf
