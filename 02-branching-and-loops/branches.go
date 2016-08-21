package main

func main() {

	/*************/
	/** if/else **/
	/*************/

	foo := 2

	if foo == 1 {
		println("bar")
	} else {
		println("buz")
	}

	// OR

	if foo = 1; foo == 1 {
		println("bar")
	} else {
		println("buz")
	}

	/************/
	/** switch **/
	/************/

	switch foo = 2; foo {
	case 1:
		println(foo)
	case 2:
		println(foo + 1)
	}

	// OR

	foo = 5

	switch {
	case foo == 1:
		println(foo)
	case foo < 1:
		println(foo + 1)
	default:
		println(5)
	}

}
