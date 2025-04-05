package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	p := Pessoa{
		Nome:  "Diogo",
		Idade: 14,
	}

	fmt.Println("Ola meu nome é: "+p.Nome+" e minha idade é:", p.Idade, "anos")
}
