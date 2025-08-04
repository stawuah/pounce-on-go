package main

import "fmt"

func main() {

	var x = 5
	var a = 10
	var b = 11

	addOne(&x)

	fmt.Printf("value of x is %+v\n", x)

	fmt.Printf("value of a is %v and value of b is %+v \n", a, b)

	swap(&a, &b)

	fmt.Printf("value of a after sawp %v and value of b after swap is %+v \n", a, b)
}

func addOne(num *int) {
	*num = *num + 1
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
