package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(acumSum(5))
}

// recebe um número e soma todos os números de 1 até ele
func acumSum(num int) int {
	if num <= 0 {
		return 1
	}
	return num + acumSum(num-1)
}
