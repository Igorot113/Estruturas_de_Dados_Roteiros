package main

import "fmt"

func main() {

	var matriz [3][3]string
	var valores string

	for i := 0; i < 3; i++ {
		fmt.Printf("Digite o Índice do aluno %d: ", i+1)
		fmt.Scan(&valores)
		matriz[i][0] = valores
		fmt.Printf("Digite o RA do aluno %d: ", i+1)
		fmt.Scan(&valores)
		matriz[i][1] = valores
		fmt.Printf("Digite o Nome do aluno %d: ", i+1)
		fmt.Scan(&valores)
		matriz[i][2] = valores
	}

	fmt.Println("\nDados do alunos:")
	fmt.Println("Índice  RA  Nome")
	for i := 0; i < 3; i++ {
		fmt.Println()
		for u := 0; u < 3; u++ {
			fmt.Print(matriz[i][u], " ")
		}
	}

}
