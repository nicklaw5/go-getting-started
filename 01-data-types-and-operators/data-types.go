package main

func main() {

	// integers
	var myVar int
	myVar = 43
	println(myVar)

	// floating point numbers
	var myFloat32 float32 = 42.
	println(myFloat32)

	// strings
	myString := "Hello Go"
	println(myString)

	// complex numbers
	myComplex := complex(2, 3)
	println(myComplex)
	println(real(myComplex))
	println(imag(myComplex))
}
