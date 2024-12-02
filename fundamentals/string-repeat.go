package main

// import "strings"

// repete uma string um número de vezes através do parâmetro
// o resultado é uma string da concatenação dessas repetições
func RepeatStr(repetitions int, value string) string {
	var repeats string
	for i := 0; i < repetitions; i++ {
		repeats += value
	}
	return repeats
	// return strings.Repeat(value, repetitions)
}
