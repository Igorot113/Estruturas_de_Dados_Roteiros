package main

import "fmt"

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func main() {
	product := Produto{
		Nome:       "Maçã",
		Preco:      12.50,
		Quantidade: 10,
	}

	fmt.Println("O nome do produto é:", product.Nome, "o preço é:", product.Preco, "a quantidade é:", product.Quantidade)
}
