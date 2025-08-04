package main

// var credentials string

// const number int = 42
// const (
// 	hieght int32   = 56
// 	length float32 = 78.89
// 	okay   bool    = true
// )

func AddfirstAndLastElements(list_numbers []int, target int) int {
	// for i := 0; i < len(list_numbers); i++ {
	// 	println(i)
	// }

	// i[0] then i[-1]
	// add  i[0] then i[-1]
	// key word is element not index

	total := 0

	for i, v := range list_numbers {
		println(i, v)

		first_element := list_numbers[0]
		last_element := list_numbers[len(list_numbers)-1]

		// A[0] = first
		// A[-1] = last element

		sum := first_element + last_element

		total = sum

	}
	println(total)
	return total

}

//  loop through the array and find two elements fiirst
// if the found elements sums to the tagert
// push the index of the element to a new array and return the indexes

func ReturnTwoSum(list_numbers []int, target int) []int {

	index_found := []int{}

	for index_of_value1, value_1 := range list_numbers {

		for index_of_value2, value_2 := range list_numbers {
			// if  value_1 + value_2 = target

			if target != 0 {
				found_indices := value_1+value_2 == target

				println(found_indices, value_2, value_1, "found elements and values of added found elements")

				complement := int(value_2+value_1) == int(value_1+value_2) && int(value_2+value_1) == target

				println(complement, "this is the complement found!!!")

				if complement {
					return append(index_found, index_of_value1, index_of_value2)
					// println(element_found, "here we are push to the element found array and see the indexes of the elements found")
					// return element_found
				}

			}
		}
		println("this is the index of the element found: ", index_found)
	}

	return index_found

}

func main() {

	array := []int{2, 7, 11, 15}

	result := ReturnTwoSum(array, 9)
	println("Final result:", result[0], result[1])
}
