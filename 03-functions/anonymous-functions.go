package main

func main() {

	addFunc := func(terms ...int) (numTerms int, sum int) {
		for _, term := range terms {
			sum += term
		}

		numTerms = len(terms)

		return numTerms, sum
	}

	numTerms, sum := addFunc(2, 4, 5)

	println(numTerms)
	println(sum)

}
