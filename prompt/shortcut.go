package prompt

import (
	"fmt"
	"io"
	"os"
)

func Print(x interface{}) {
	fmt.Printf("%v\r\n", x)
}

var Printf = fmt.Printf

func DebugPrint(x interface{}) error {
	file, err := os.Create("/tmp/rango")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.WriteString(file, fmt.Sprintf("%v\n", x))
	if err != nil {
		panic(err)
	}
	return file.Sync()
}
