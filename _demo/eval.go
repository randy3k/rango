package main

import (
	"rango"
	"fmt"
)

func main() {
	ok, err := rango.Initialize(rango.GetRhome(), rango.DefaultArgs())
	if (err != nil) {
		panic(err)
	} else if (!ok) {
		panic("R was not initialized")
	}
	fmt.Println("hello")

}
