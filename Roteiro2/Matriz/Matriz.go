package main

import "fmt"

func main() {
	var matriz [2][3]int
	fmt.Println(matriz)

	matriz2 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(matriz2)

	matriz3 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(matriz3[1][0])
	matriz3[0][1] = 9
	fmt.Println(matriz3)

	for i := 0; i < len(matriz2); i++ {
		for j := 0; j < len(matriz2[i]); j++ {
			fmt.Printf("%d ", matriz2[i][j])
		}
		fmt.Println()
	}

	for a, linha := range matriz2 {
		for b, valor := range linha {
			fmt.Printf("Matriz[%d][%d] = %d\n", a, b, valor)
		}
	}
}
