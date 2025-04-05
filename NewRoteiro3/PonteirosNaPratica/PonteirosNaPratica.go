package main

import "fmt"

func main() {
	var a int = 42
	var p *int = &a
	fmt.Println("Valor de a: ", a)
	fmt.Println("Endereço de a: ", &a)
	fmt.Println("Valor de p(endereço de a): ", p)
	fmt.Println("Valor apontado por p: ", *p)

	//Modificando Valores atraves de Ponteiros

	fmt.Println("Valor de a antes da modificação: ", a)
	*p = 100
	fmt.Println("Valor de a depois da modificação: ", a)
}
