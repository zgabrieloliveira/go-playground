package main

// recebe um número e soma todos os números de 1 até ele
func AcumSum(num int) int {
	var sum int
	for i := 1; i < num; i++ {
		sum += i
	}
	return sum
}
