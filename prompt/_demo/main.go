package main

import (
	"fmt"
	// "github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/randy3k/rango/prompt"
)

func main() {
	// p := prompt.NewPrompt()
	// for i := 0; i < 100; i++ {
	// p.Show("> ")
	// }

	s := "# Solve the quadratic equation ax**2 + bx + c = 0\n\n\n# import complex math module\nimport cmath\n\na = 1\nb = 5\nc = 6\n\n# calculate the discriminant\nd = (b**2) - (4*a*c)\n\n# find two solutions\nsol1 = (-b-cmath.sqrt(d))/(2*a)\nsol2 = (-b+cmath.sqrt(d))/(2*a)\n\nprint('The solution are {0} and {1}'.format(sol1,sol2))"
	lexer := lexers.Get("python")
	style := styles.Get("xcode")
	buffer := prompt.NewBuffer(lexer, style)
	screen := prompt.NewScreen(15, 45)
	container := prompt.NewContainer(buffer)

	buffer.SetText(s)

	container.WriteToScreen(screen)

	// fmt.Println(buffer.GetLines())

	// fmt.Printf("%v\n", style.Get(chroma.Operator))
	fmt.Println(screen)
	// fmt.Println(screen.Dirty)
}
