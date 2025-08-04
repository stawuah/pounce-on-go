package main

import "fmt"

type bite *[32]byte

func main() {
	var name *string

	// var newBite bite

	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // last element
	fmt.Println(name)
	// fmt.Println(newBite)

	// Create an actual array to work with
	var byteArray [32]byte
	// Fill it with some sample data
	for i := 0; i < 32; i++ {
		byteArray[i] = byte(i)
	}

	// Set newBite to point to our array
	newBite := &byteArray

	// Call the function
	processBite(newBite)

	// Print the type information
	fmt.Printf("Type of newBite: %T\n", newBite)
}

// Function that takes a bite (which is *[32]byte) as parameter
func processBite(b bite) {
	fmt.Println("Processing bite...")

	// Check if the pointer is nil
	if b == nil {
		fmt.Println("bite is nil, cannot loop")
		return
	}

	// Loop over the array that the pointer points to
	for i, value := range *b { // Dereference the pointer to get the array
		fmt.Printf("Index %d: %d\n", i, value)
	}

	// Alternative loop using traditional for loop
	fmt.Println("\nUsing traditional for loop:")
	for i := 0; i < len(*b); i++ {
		fmt.Printf("Index %d: %d\n", i, (*b)[i])
	}
}
