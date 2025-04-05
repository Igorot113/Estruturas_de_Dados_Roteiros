package main

import "fmt"

//Passagem por valor
func increment(val int) {
	val++
	fmt.Println("Dentro da função increment: ", val)
}

// Passagem por referencia

func incrementByPointer(val *int) {
	*val++
	fmt.Println("Dentro da função incrementByPointer: ", *val)
}
func main() {
	//passagem por valor
	x := 10
	fmt.Println("Fora da função increment: ", x)
	increment(x)

	// Passagem por referencia
	y := 11
	fmt.Println("Fora da função incrementByPointer: ", y)
	incrementByPointer(&y)
}
