package main

func CalculateYears(years int) (result [3]int) {
	yearsArray := [3]int{years}
	yearsArray[1] = calculateCatYears(years)
	yearsArray[2] = calculateDogYears(years)
	return yearsArray
}

func calculateCatYears(years int) int {

	if years == 1 {
		return 15
	}
	if years == 2 {
		return 15 + 9
	}

	// depois de 2 anos, cada ano equivale a 4 anos de um gato (+24 dos dois primeiros)
	return 4*(years-2) + 24
}

func calculateDogYears(years int) int {
	if years == 1 {
		return 15
	}
	if years == 2 {
		return 15 + 9
	}

	// depois de 2 anos, cada ano equivale a 5 anos de um cão (+24 dos dois primeiros)
	return 5*(years-2) + 24

}
