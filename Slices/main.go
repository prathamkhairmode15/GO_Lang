package main

import (
	"fmt"
	"slices"
)

// slices are the dynamic arrays
// they ae the most used constructs in go
// we have useful method in the slices
func main() {
	//unintialized slice is nil
	//var nums []int
	//fmt.Println(nums)
	//fmt.Println(len(nums))
	var nums = make([]int, 5, 20)
	//var nums = make([]int, size,initial capacity)
	//cap > capacity is maxinum number of elements can fit
	fmt.Println(cap(nums))
	// we use the function append to add the elements
	nums = append(nums, 10)
	//appending using the loop
	for i := 0; i < 5; i++ {
		nums = append(nums, i)
	}
	nums[4] = 3
	fmt.Println(nums)

	//we have copy() function to copy the elements from one slice to another
	var nums2 = make([]int, len(nums))
	nums = append(nums, 15)
	copy(nums2, nums)
	fmt.Println(nums, nums2)

	//slice operator :
	var n = []int{1, 2, 3, 4, 5}
	fmt.Println(n[0:3])

	//compare slices
	var a = []int{1, 2, 4}
	var b = []int{1, 2, 3}
	fmt.Println(slices.Equal(a, b))
}
