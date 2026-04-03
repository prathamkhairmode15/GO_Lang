package main

import "fmt"

func main() {
	var a = []int{10, 20, 35, 67, 40}

	for i, num := range a {
		fmt.Printf("index: %d, value: %d\n", i, num)
	}

	sum := 0
	for _, s := range a {
		sum += s
	}
	fmt.Printf("Sum is : %d\n", sum)

	//range and map
	var m = map[string]int{"age": 20, "class": 2}
	for k, v := range m {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
	//we can also ignore value only print key
	for k := range m {
		fmt.Printf("key: %s\n", k)
	}
}
