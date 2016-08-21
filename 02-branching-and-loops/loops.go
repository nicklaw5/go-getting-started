package main

func main() {

	// for loops

	for i := 0; i < 5; i++ {
		println(i)
	}

	// while loops

	i := 0
	for {
		i++
		println(i)

		if i > 5 {
			break
		}
	}

	// loop over collections (arrays, slices or maps)

	s := []string{"foo", "bar", "baz"}
	for id, v := range s {
		println(id, v)
	}

	// maps
	m := make(map[string]string)
	m["first"] = "foo"
	m["second"] = "bar"
	m["third"] = "baz"

	for k, v := range m {
		println(k, v)
	}
}
