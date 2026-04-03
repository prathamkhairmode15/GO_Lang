package main

import (
	"fmt"
	"maps"
)

// maps > hash,objects,dictionary
func main() {
	//creating map
	var m = make(map[string]string)
	//setting an elements
	m["name"] = "Pratham"
	m["Dept"] = "CSE"
	//get element
	fmt.Println(m["name"], m["Dept"])
	//fmt.Println(m["phone"])
	//length of the map
	fmt.Println(len(m))

	//delete function to delete in the map
	delete(m, "Dept")
	fmt.Println(len(m))

	var n = map[string]int{"Age": 20, "Roll": 101}
	fmt.Println(n)

	//compare maps
	var m1 = map[string]string{"name": "Pratham", "Dept": "CSE"}
	var m2 = map[string]string{"name": "Pratham", "Dept": "CSE"}
	fmt.Println(maps.Equal(m1, m2)) //maps cannot be compared
}
