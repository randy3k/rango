package cli

import (
	"bufio"
	"fmt"
	"github.com/randy3k/rango"
	"os"
	"runtime"
)

func Run() {
	runtime.LockOSThread()
	ok, err := rango.Initialize(rango.GetRhome(), rango.DefaultArgs())
	if err != nil {
		panic(err)
	} else if !ok {
		panic("R was not initialized")
	}
	rango.Callbacks.ReadConsole = readConsole
	rango.Callbacks.WriteConsole = writeConsole
	rango.SetCallbacks()
	rango.RunREPL()
	runtime.UnlockOSThread()
}

func readConsole(p string, add_history bool) string {
	fmt.Print(p)
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	if len(text) > 0 {
		text = text[0:(len(text) - 1)]
	}
	return text
}

func writeConsole(p string, otype int) {
	fmt.Print(p)
}
