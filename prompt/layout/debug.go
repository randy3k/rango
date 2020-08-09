package layout

import (
	"fmt"
	"io"
	"os"
)

func debugPrintln(x ...interface{}) error {
	file, err := os.Create("/tmp/rango")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.WriteString(file, fmt.Sprint(x...) + "\n")
	if err != nil {
		panic(err)
	}
	return file.Sync()
}
