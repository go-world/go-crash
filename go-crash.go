package main

import "fmt"

func main() {
	beyondHello()
}

func beyondHello() {
	var x int
	x = 3 // Two line variable declaration and assignment

	y := 4 // Infer the type

	sum, prod := learnMultiple(x, y) // Functions can return two values.
	fmt.Println("sum:", sum, "prod", prod)
	learnTypes()
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {
	str := "Learning Go!!"

	s2 := `Testing multiple
  line`

	g := 'Σ' // Go source is UTF-8. This converts to rune data type which is an alias for int32

	f := 3.14195 // float 64
	c := 3 + 4i  // complex128, which internally is two for float64s

	var u uint = 7 // unsigned int
	var pi float32 = 22. / 7

	n := byte('\n') // byte is an alias for uint8

	// Arrays have fixed size at compile time.
	var a4 [4]int           // An array of 4 ints, intitalized to 0s
	a3 := [...]int{3, 1, 5} //Array intitalized with 3,1,5

	// Slices have dynamic sizes
	s := []int{1, 2, 3}    // Slice of length three.
	s = append(s, 4, 5, 6) // Append elements to existing slice.
	fmt.Println(s)

	s = append(s, []int{7, 8, 9}...)
	fmt.Println(s)

	p, q := learnMemory()
	fmt.Println(*p, *q)

	m := map[string]int{"three": 3, "four": 4}
	fmt.Println(m)
	m["one"] = 1

	// Unused variables result in error in GO. Use underscore to discard the values.
	_, _, _, _, _, _, _, _, _, _ = str, s2, g, f, u, pi, n, a3, c, a4

	learnFlowControl() // loops

}

// This is cool. GO allows named returns in the signature of the function.
// This way we can return from anywhere within the function
func learnNamedReturns(x, y int) (z int) {
	z = x * y
	return
}

func learnMemory() (p, q *int) {
	// Uptil now p and q are nil pointers.
	p = new(int)         // new is a function that allocates memory. p is no longer a nil pointer
	s := make([]int, 20) // Allocate 20 ints in a single block of memory
	s[3] = 7
	r := -2
	return &s[3], &r
}

func learnFlowControl() {
	if true {
		fmt.Println("if statements require curly braces")
	}

	x := 42

	// for loop
	for x := 0; x <= 3; x++ {
		fmt.Println("looping x", x)
	}

	for { // Infinite loop
		break
	}

	for key, value := range map[string]int{"one": 1, "two": 2} {
		fmt.Println(key, value)
	}

	// Closures
	xBig := func() bool {
		return x > 1000
	}

	fmt.Println(xBig())

	x = 1.3e3 // This makes x == 1300

	// Function literals can be used inline
	fmt.Println("Adding two numbers", func(a, b int) int {
		return a + b
	}(10, 2))

	goto: love
	love:
		learnFunctionFactory()
		learnDefer()
		learnInterfaces()
}
