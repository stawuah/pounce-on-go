package main

import "fmt"

// You've defined a new type, but you can also just use *int directly
// in the function signature. This is often simpler.
type egine_number *int

func main() {
	ten := 10
	fmt.Println("Value of 'ten' before update:", ten)

	// To pass the address of `ten`, we use the '&' operator.
	// The function `updateEngineNumber` now receives a pointer to 'ten'.
	updateEngineNumber(&ten)

	fmt.Println("Value of 'ten' after update:", ten)
}

// The function `updateEngineNumber` now accepts a parameter that is a pointer to an integer.
// This is the `*int` type, which you named `egine_number`.
// The parameter `p` is a pointer.
func updateEngineNumber(p egine_number) {
	// To modify the value at the memory address that `p` points to,
	// we must first dereference the pointer using the '*' operator.
	// The `*p` expression gives us the actual integer value that `p` points to.

	// We can then re-assign a new value to the dereferenced pointer.
	// For example, to increment the value by 1:
	*p = *p + 1

	// You can also use the shorthand assignment operator if you prefer:
	// *p++ is invalid syntax in Go.
	// *p += 1
}
