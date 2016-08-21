package main

import "strings"

func main() {
	message := "Hello"
	sayHello(message)
	sayHello2(&message)
	sayHello3("Hi", "there", "Go", "!")
	println(message)
}

/** passing by value **/

func sayHello(message string) {
	println(message)
}

/** passing by reference **/

func sayHello2(message *string) {
	println(*message)
	*message = "Hello Go"
}

/** variatic functions **/

func sayHello3(messages ...string) {
	var str string
	for _, msg := range messages {
		str += " " + msg
	}

	println(strings.Trim(str, " "))
}
