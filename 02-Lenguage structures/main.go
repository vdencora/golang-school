package main

import (
	"fmt"
	"time"
)

func main() {

	/* -------------------------------------------------------------------------- */
	/*                             IF-ELSE SENTENCE                               */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "IF-ELSE SENTENCE")
	num1 := 10
	num2 := 5

	if num1%num2 == 0 {
		fmt.Printf("%d is a multiple of %d \n", num1, num2)
	} else {
		fmt.Printf("%d is not a multiple of %d \n", num1, num2)
	}

	/* -------------------------------------------------------------------------- */
	/*                             SWITCH STATEMENT                               */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "SWITCH STATEMENT")

	// switch with condition
	fmt.Println("When is saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("is today!")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("the day after tomorrow")
	default:
		fmt.Println("it's so far")
	}

	// switch without condition
	fmt.Println("What are you according to your age?")
	age := 29
	switch {
	case age <= 10:
		fmt.Println("you are a boy")
	case age > 10 && age <= 18:
		fmt.Println("you're a teenager")
	case age > 18 && age <= 60:
		fmt.Println("you're an adult")
	default:
		fmt.Println("you're and old man")
	}

	/* -------------------------------------------------------------------------- */
	/*                              FOR STATEMENTS                                */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "FOR STATEMENTS")

	fmt.Printf("\n******** %s ******\n", "Counter-controlled iterations")
	for i := 1; i <= 10; i++ {
		fmt.Println("counting ", i)
		time.Sleep(300 * time.Millisecond)
	}

	fmt.Printf("\n******** %s ******\n", "Condition-controlled iterations")
	// For as a While
	i := 0
	for i != 5 {
		fmt.Println("for as a while", i)
		i++
	}
	// For as a Do-While
	i = 0
	for {
		fmt.Println("for as a do-while", i)
		i++
		if i >= 5 {
			break
		}
	}

	fmt.Printf("\n******** %s ******\n", "Range over data structs")
	// Range over a slice
	names := []string{"Marta", "Miguel", "Carlos", "Maria", "Mateo"}
	for indice, value := range names {
		fmt.Println("index: ", indice, "value: ", value)
	}

	// Range over a map
	mapaNombreEdad := map[string]int{
		"Marta":  20,
		"Miguel": 30,
		"Carlos": 18,
		"Maria":  35,
		"Mateo":  25,
	}
	for nombre, edad := range mapaNombreEdad {
		fmt.Println(nombre, " is ", edad, "years old")
	}

	fmt.Printf("\n******** %s ******\n", "Using continue")
	// skip no-multiples of two
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
