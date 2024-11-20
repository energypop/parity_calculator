package main

import (
	"fmt"
)

func main() {
	// Taking binary blocks as input
	var firstBlock string
	var secondBlock string

	fmt.Print("Introduce el primer bloque binario: ")
	fmt.Scanln(&firstBlock)

	fmt.Print("Introduce el segundo bloque binario: ")
	fmt.Scanln(&secondBlock)

	// Ensure both blocks have the same length
	if len(firstBlock) != len(secondBlock) {
		fmt.Println("Los bloques deben tener la misma longitud.")
		return
	}

	for i := 0; i < len(firstBlock); i++ {
		
		if firstBlock[i] == secondBlock[i] {
			fmt.Print("0")
		} else {
			fmt.Print("1")
		}
	}
}
