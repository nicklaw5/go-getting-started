package main

import "fmt"

func main() {

	/************/
	/** ARRAYS **/
	/************/

	// explicitly decalre an array by size
	myArray := [3]int{}
	myArray[0] = 42
	myArray[1] = 43
	myArray[2] = 44
	fmt.Println(myArray)

	// declare an array without a size
	myArray2 := [...]int{45, 46, 47, 48, 49}
	fmt.Println(myArray2)

	// get array size
	fmt.Println(len(myArray))
	fmt.Println(len(myArray2))

	/************/
	/** SLICES **/
	/************/

	// create a new slice
	mySlice := []float32{12., 15., 18.}
	fmt.Println("mySlice:", mySlice)
	fmt.Println("mySlice Lenght:", len(mySlice))

	// desclare a slice from an array
	mySlice2 := myArray2[:]
	fmt.Println(mySlice2)
	fmt.Println(mySlice2[2])

	// append to slice
	mySlice2 = append(mySlice2, 99, 89, 877)
	fmt.Println(mySlice2)

	// creating a slice from make() function
	mySlice3 := make([]float32, 100)
	mySlice3[0] = 56
	mySlice3[2] = 98
	mySlice3[5] = 97
	fmt.Println(mySlice3)

	/**********/
	/** MAPS **/
	/**********/
	myMap := make(map[int]string)
	myMap[5] = "Five"
	myMap[8] = "Eight"
	fmt.Println(myMap)
	fmt.Println(myMap[10]) // return the "zero value" of a string, which is an empty string ("")

}
