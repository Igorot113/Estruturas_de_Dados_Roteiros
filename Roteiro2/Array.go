package main

import "fmt"

func main() {

	var arr [5]int
	fmt.Print(arr)

	fmt.Println()

	arr2 := [3]int{10, 20, 30}
	fmt.Print(arr2)

	fmt.Println()

	arr3 := [3]int{5, 10, 15}
	arr3[1] = 20
	fmt.Println(arr3[1])

	arr4 := [3]int{10, 20, 30}

	for i := 0; i < len(arr4); i++ {
		fmt.Printf("%d ", arr4[i])
	}

	fmt.Println()

	for i, v := range arr2 {
		fmt.Printf("Índice: %d, Valor: %d\n", i, v)
	}

	fmt.Print("Valor: ")
	for _, v := range arr2 {
		fmt.Printf("%d ", v)
	}
}
