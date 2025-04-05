package main

import "fmt"

func main() {
	var b float64 = 10.10
	var p2 *float64 = &b

	fmt.Println("Valor de b antes: ", b)
	*p2 = 11.11
	fmt.Println("Valor de b depois: ", b)
}
