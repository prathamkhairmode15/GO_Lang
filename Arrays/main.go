package main

import "fmt"

//fixed size then arrays can be used
//memory can be optimized using the arrays
//constant time access to the elements
func main() {
	var nums [5]int
	//array length
	fmt.Println(len(nums))

	//adding element to the array
	nums[0] = 1
	fmt.Println(nums[0])

	fmt.Println(nums)

	//boolean array
	var vals [5]bool
	vals[3] = true
	fmt.Println(vals)

	//string array
	var names [4]string
	names[0] = "Pratham"
	fmt.Println(names)

	//2d arrays
	nums2d := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums2d)
}
