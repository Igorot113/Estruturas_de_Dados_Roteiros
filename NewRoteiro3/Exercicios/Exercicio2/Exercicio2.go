package main

import "fmt"

func swap(x *int, y *int) {
	x, y = y, x
	fmt.Println("Depois do swap: ")
	fmt.Println("x = ", *x)
	fmt.Println("y = ", *y)
}

func main() {
	x := 10
	y := 20
	fmt.Println("Antes do swap: ")
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	swap(&x, &y)

}
