package main

func main() {
	total := getTotal(1, 3, 5, 9)
	println(total)

	total1, total2 := getTotals(1, 3, 5, 9)
	println(total1)
	println(total2)

	total3, total4 := getTotalsWithNames(1, 3, 5, 9)
	println(total3)
	println(total4)
}

/** simple return function **/

func getTotal(terms ...int) int {
	var total int
	for _, term := range terms {
		total += term
	}

	return total
}

/** multiple return vales **/

func getTotals(terms ...int) (int, int) {
	var total int
	for _, term := range terms {
		total += term
	}

	return total, total + 1
}

/** multiple return value with name params **/

func getTotalsWithNames(terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}

	numTerms = len(terms)

	return numTerms, sum
}
