package main

import "fmt"

func verificarPar(x int) string {
	if x%2 == 0 {
		return "Ele é par"
	} else {
		return "Ele é impar"
	}
}

func main() {
	var numero1 int = 2
	numero2 := 3

	fmt.Print(verificarPar(numero1), "\n", verificarPar(numero2))
}
