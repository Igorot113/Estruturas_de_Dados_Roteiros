package main

import "fmt"

func main() {
	var nota int = 0
	fmt.Println(Classificacao(nota))
}

func Classificacao(nota int) string {
	switch nota {
	case 0, 1, 2:
		return "Muito Ruim"
	case 3, 4:
		return "Ruim"
	case 5, 6:
		return "Regular"
	case 7, 8:
		return "Bom"
	case 9, 10:
		return "Excelente"
	default:
		return "Não Existe"
	}

}
