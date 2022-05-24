package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

func main() {

	/* -------------------------------------------------------------------------- */
	/*                                    ARRAYS                                  */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "ARRAYS")

	fmt.Printf("\n******** %s ******\n", "Declaration")

	// way 1
	systems := [3]string{"Linux", "MacOS", "Windows"}
	fmt.Printf("Array 1: %v with len %d and cap %d \n", systems, len(systems), cap(systems))

	// way 2
	var numbers [5]int
	numbers[0] = 25
	numbers[1] = 12
	numbers[2] = 90
	numbers[3] = 100
	numbers[4] = 33
	fmt.Printf("Array 2: %v with len %d and cap %d \n", numbers, len(numbers), cap(numbers))

	fmt.Printf("\n******** %s ******\n", "Multidimensional Arrays")

	multidimensionalArr := [3][3]string{
		{"*", "*", "*"},
		{"*", "*", "*"},
		{"*", "*", "*"},
	}
	// print the multidimensional array with format
	for _, secondArr := range multidimensionalArr {
		for _, val := range secondArr {
			fmt.Print(val + " ")
		}
		fmt.Println("")
	}

	fmt.Printf("\n******** %s ******\n", "Array parameters")
	// array is a value type. See how we update the array inside the function without altering the original value
	updateAnArray(systems)
	fmt.Println("array outside updateAnArray function", systems)

	/* -------------------------------------------------------------------------- */
	/*                                  SLICES                                    */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "SLICES")

	fmt.Printf("\n******** %s ******\n", "Declaration")

	// way 1
	animals := []string{"Dog", "Cat", "Tigger", "Bird", "Spider", "Ant", "Elephant"}
	fmt.Printf("Slice: %v with len %d \n", animals, len(animals))

	// way 2
	cats := make([]string, 0, 20)
	fmt.Printf("Slice: %v with len %d and cap %d \n", cats, len(cats), cap(cats))

	fmt.Printf("\n******** %s ******\n", "Append")
	// adding elements to the slice
	cats = append(cats, "garfield")
	cats = append(cats, "silvestre")
	cats = append(cats, "tom")
	fmt.Printf("Slice: %v with len %d and cap %d \n", cats, len(cats), cap(cats))

	fmt.Printf("\n******** %s ******\n", "Multidimensional Slices")
	// declare an slice of slices
	data := make([][]string, 0)
	data = append(data, animals)
	data = append(data, cats)
	fmt.Println("Slice of Slices: ", data)

	fmt.Printf("\n******** %s ******\n", "Slice parameters")
	// the slice is the reference to a segment of an array, see how when updating the value of the slice inside the function, its original value is also altered
	updateASlice(cats)
	fmt.Println("slice outside updateASlice function", cats)

	fmt.Printf("\n******** %s ******\n", "Reslicing")
	// As we said the slice is a segment of an array, so we can make references to new segments and create new slices as follows
	newSlice := animals[2:6]
	fmt.Printf("Reslicing: %v with len %d and cap %d \n", newSlice, len(newSlice), cap(newSlice))

	/* -------------------------------------------------------------------------- */
	/*                                   MAPS                                     */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "MAPS")

	fmt.Printf("\n******** %s ******\n", "Declaration")

	// way 1
	map1 := map[string][]string{
		"animals": {"Dog", "Cat", "Tigger", "Bird", "Spider", "Ant", "Elephant"},
		"humans":  {"Juan", "Paul", "Maria", "Marcos", "Ana"},
	}

	// way 2
	map2 := make(map[string][]string, 2)
	map2["animals"] = []string{"Dog", "Cat", "Tigger", "Bird", "Spider", "Ant", "Elephant"}
	map2["humans"] = []string{"Juan", "Paul", "Maria", "Marcos", "Ana"}

	fmt.Println("map 1", map1)
	fmt.Println("map 2", map2)

	fmt.Printf("\n******** %s ******\n", "Delete from a map")

	delete(map1, "animals")
	fmt.Println("mapa 1 after delete", map1)

	fmt.Printf("\n******** %s ******\n", "Zero value of a map is nil")
	// uncomment the line below and see the nil reference error
	// Nil reference
	/*
		var mapNil map[int]bool
		mapNil[1] = true
	*/

	fmt.Printf("\n******** %s ******\n", "Map parameter")
	// a map is a reference, see how when updating the value of the map inside the function, its original value is also altered
	updateAMap(map1)
	fmt.Println("map1 after call updateAMap: ", map1)

	/* -------------------------------------------------------------------------- */
	/*                                  STRUCTS                                   */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "STRUCTS")

	// struct declaration
	type Hero struct {
		Name string
	}

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Hero Hero
	}

	// create new instance of a struct
	spiderman := Hero{
		Name: "spiderman",
	}

	// see how the structure person has another structure inside (Hero)
	// this is called composition and allows us to work with inheritance
	person := Person{
		Name: "Peter Parker",
		Age:  18,
		Hero: spiderman,
	}

	b, _ := json.Marshal(person)

	// The tags of a structure allow, among many of its functions, to identify the fields when working with json, xml, etc.
	// See how in the print below, the fields are identified with the tag key (`json:"key"`).
	// If the structure field doesn't have a tag, the key will be the same as the field name (as is the case with Hero).
	fmt.Println(string(b))

	fmt.Printf("\n******** %s ******\n", "Empty struct")

	// An empty struct doesn't consume storage.
	// look at the following zero values of some types and see what their memory space is
	fmt.Println(unsafe.Sizeof(struct{}{}))
	fmt.Println(unsafe.Sizeof(false))
	fmt.Println(unsafe.Sizeof(0))
	fmt.Println(unsafe.Sizeof(""))

	// this is quite useful when you care about sending a signal more than data (in channels for example which you will see later)
	// or if you want to create a map where you only care about the value of the keys for example
	// see the following example and note the difference in memory space
	fmt.Println()
	fmt.Println(unsafe.Sizeof([10000]int{}))
	fmt.Println(unsafe.Sizeof([10000]struct{}{}))
}

// updates an array
func updateAnArray(array [3]string) {
	array[0] = "x"
	array[1] = "y"
	array[2] = "z"
	fmt.Println("array inside the updateAnArray function", array)
}

// updates a slice
func updateASlice(slice []string) {
	slice[0] = "x"
	slice[1] = "y"
	slice[2] = "z"
	fmt.Println("slice inside the updateASlice function", slice)
}

// updates a map
func updateAMap(map1 map[string][]string) {
	map1["newKey"] = []string{"some", "new", "value"}
	fmt.Println("map1 inside updateAMap function: ", map1)
}
