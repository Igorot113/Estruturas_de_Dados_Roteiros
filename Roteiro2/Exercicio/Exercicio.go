package main

import "fmt"

func main() {
	var vetor [6]int

	vetor[0] = 1
	vetor[1] = 2
	vetor[2] = 4
	vetor[3] = 7
	vetor[4] = 7
	vetor[5] = 5

	fmt.Print("Printando manualmente: ", vetor[0])
	fmt.Print(" ", vetor[1])
	fmt.Print(" ", vetor[2])
	fmt.Print(" ", vetor[3])
	fmt.Print(" ", vetor[4])
	fmt.Print(" ", vetor[5])
	fmt.Println()
	fmt.Print("Usando o for: ")
	for i := 0; i < len(vetor); i++ {
		fmt.Print(" ", vetor[i])
	}
	fmt.Println()
	fmt.Println("Usando o Range: ")
	for i, v := range vetor {
		fmt.Printf("Index: %d Valor: %d\n", i, v)
	}

}
