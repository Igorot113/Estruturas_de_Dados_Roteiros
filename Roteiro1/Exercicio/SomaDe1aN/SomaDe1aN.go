package main

import "fmt"

func fazerSoma(numero int) int {
	var iterador, soma int
	for iterador = 0; iterador < numero; iterador++ {
		soma += iterador
	}
	return soma
}
func main() {
	var numero1 int = 10
	fmt.Println(fazerSoma(numero1))
}
