package main

func main() {
	emp := newPrinter()
	emp.message = "foo"

	emp.printMessage()
}

// main struct
type messagePrinter struct {
	message string
}

// inherit the compostion of messagePrinter
type enhancedMessagePrinter struct {
	messagePrinter
}

// method of messagePrinter
func (mp *messagePrinter) printMessage() {
	println(mp.message)
}

// constructor
func newPrinter() *enhancedMessagePrinter {
	result := enhancedMessagePrinter{}
	// result.myMap = map[string]string{}
	return &result
}
