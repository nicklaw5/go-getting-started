package main

const (
	first = iota
	second
)

const (
	third = iota
)

const (
	fourth = 1 << (10 * iota)
	fifth
	sixth
)

func main() {
	println(first)
	println(second)
	println(third)
	println(fourth)
	println(fifth)
	println(sixth)
}
