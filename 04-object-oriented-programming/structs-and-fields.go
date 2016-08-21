package main

import "fmt"

func main() {
	var1 := myStruct{}
	var1.myField = "bar"
	println(var1.myField)

	var2 := myStruct{"foo"}
	println(var2.myField)

	var3 := myStruct{myField: "baz"}
	println(var3.myField)

	var4 := new(myStruct)
	var4.myField = "bug"
	println(var4.myField)

	var5 := myStruct{}
	var5.myField = "boo"
	fmt.Println(var5)

	var6 := &myStruct{"bat"}
	fmt.Println(var6)

}

type myStruct struct {
	myField string
}
