package main

import "fmt"

type Counter struct {
	count int
}

type Rectangle struct {
	hieght float32
	width  float32
}

func (a Rectangle) Area() float32 {
	total_area := a.hieght * a.width

	return float32(total_area)
}

func (a *Rectangle) Scaler(factor float32) (float32, float32) {
	a.hieght *= factor
	a.width *= factor

	return a.hieght, a.width
}

// Increment modifies the original Counter struct.
// It uses a pointer receiver to change the 'count' field directly.
func (c *Counter) Increment() {
	c.count++ // This is shorthand for c.count = c.count + 1
}

// Value returns the current count without modifying the struct.
// It uses a value receiver, so it operates on a copy.
func (c Counter) Value() int {
	return c.count
}

func main() {
	var x = 5
	var a = 10
	var b = 11

	// Call addOne to modify x in place.
	addOne(&x)
	fmt.Printf("Value of x is now: %v\n", x)

	// Print values before swap.
	fmt.Printf("Value of a is %v and value of b is %v \n", a, b)

	// Call swap to modify a and b in place.
	swap(&a, &b)

	// Print values after swap.
	fmt.Printf("Value of a after swap is %v and value of b after swap is %v \n", a, b)

	// Create and print a new Counter.
	new_counter := Counter{1}
	fmt.Printf("Value of count before Increment: %v\n", new_counter.Value())

	// Call Increment on the pointer to the counter.
	// This modifies the original 'new_counter' struct.
	new_counter_ptr := &new_counter
	new_counter_ptr.Increment()

	// Print the value after Increment.
	fmt.Printf("Value of count after Increment: %v\n", new_counter.Value())

	// Calling the Increment method again will change the value again
	new_counter_ptr.Increment()
	fmt.Printf("Value of count after a second Increment: %v\n", new_counter.Value())

	myRectangle := &Rectangle{hieght: 10, width: 5}

	fmt.Printf("Original Area: %.2f\n", myRectangle.Area())

	newHieght, newWidth := myRectangle.Scaler(2)
	fmt.Printf("Scaled Rectangle: Height = %.2f, Width = %.2f\n", newHieght, newWidth)

	fmt.Printf("New Area: %.2f\n", myRectangle.Area())
}

// addOne uses a pointer to modify the integer in place.
func addOne(num *int) {
	*num = *num + 1
}

// swap uses pointers to swap the values of two integers.
func swap(a, b *int) {
	*a, *b = *b, *a
}
