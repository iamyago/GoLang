package main

import (
	"fmt"
)

// := declaração curta de variável so pode ser usada dentro de funções/code blocks
// var declaração normal de variável pode ser usada em qualquer lugar do código

func main() {
	x := 10
	y := "apenas"
	fmt.Printf("x: %v, y: %v", x, x)
	fmt.Printf("y: %v, x: %v", y, y)
	// %v : para valor
	// %T : para tipo da variável
	// %d : para números decimais
	// %b : para números binários
	// %#u : para unicode
	// %#x : para números hexadecimal
}
