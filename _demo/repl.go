package main

import (
	"github.com/randy3k/rango"
)

func main() {
	ok, err := rango.Initialize(rango.GetRhome(), rango.DefaultArgs())
	if (err != nil) {
		panic(err)
	} else if (!ok) {
		panic("R was not initialized")
	}
	rango.SetCallbacks()
	rango.RunREPL()
}
