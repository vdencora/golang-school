package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/* -------------------------------------------------------------------------- */
	/*                                    FUNCTIONS                               */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "FUNCTIONS")

	hi("Encora")
	r1, r2, r3 := returnsManyValues(3, "hello")
	fmt.Println("returnsManyValues functions returned: ", r1, r2, r3)

	/* -------------------------------------------------------------------------- */
	/*                    CALL FUNCTION BY VALUE AND REFERENCE                    */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "CALL FUNCTION BY VALUE AND REFERENCE")

	initValue := "This is the init value"
	callByValue(initValue)
	fmt.Println("value after call callByValue function: ", initValue)
	callByReference(&initValue)
	fmt.Println("value after call callByReference function: ", initValue)

	/* -------------------------------------------------------------------------- */
	/*                          NAMED RETURN VARIABLES                            */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "NAMED RETURN VARIABLES")
	x, y := returnNamedVariables(30)
	fmt.Println("return values: ", x, y)

	/* -------------------------------------------------------------------------- */
	/*                              BLANK IDENTIFIER                              */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "BLANK IDENTIFIER")
	// suppose you only want the 2do return value, you can ignore the other two with _
	_, wanted, _ := returnsManyValues(1, "I only want the value 2")
	fmt.Println("wanted: ", wanted)

	/* -------------------------------------------------------------------------- */
	/*                                   DEFER                                    */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "DEFER")
	printBeforeExit()

	/* -------------------------------------------------------------------------- */
	/*                       RECURSIVE FUNCTIONS                                  */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "RECURSIVE FUNCTIONS")
	factorialNum := 5
	fmt.Printf("factorial of %d: %d\n", factorialNum, factorial(factorialNum))

	/* -------------------------------------------------------------------------- */
	/*                        FUNCTIONS AS PARAMETERS                             */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "FUNCTIONS AS PARAMETERS")
	generateRandom := func() int {
		x1 := rand.NewSource(time.Now().UnixNano())
		y1 := rand.New(x1)

		return y1.Intn(200)
	}
	fmt.Println("multiplyRandomNumberBy3() = ", multiplyRandomNumberBy3(generateRandom))

	/* -------------------------------------------------------------------------- */
	/*                           VARIADIC FUNCTIONS                               */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "VARIADIC FUNCTIONS")
	// Feel free to add more parameters to see how the function works
	sayHiToEveryone("Carlos", "Jose", "Maria", "Jessica")
}

// function that doesn't return a value (ie. void in java)
func hi(name string) {
	fmt.Println("Hi ", name)
}

// function that receives many params and returns many values
func returnsManyValues(param1 int, param2 string) (string, int, bool) {
	fmt.Println("values received: ", param1, param2)
	return "hola", 3, true
}

// this function receives the value so, any modification here will only be in the scope of the function
func callByValue(data string) {
	// assign a new value to the function
	data = "new value"
	fmt.Println("value within callByValue function: ", data)
}

// this function receives a pointer so, any modification here will reflected outside the scope of this function
func callByReference(data *string) {
	// assign a new value to the function
	*data = "this was updated within callByReference function"
}

// returns named variables
func returnNamedVariables(num int) (x, y int) {
	x = num + 10
	y = num + 20
	return
}

// defer example
func printBeforeExit() {
	fmt.Println("counting ...")

	for i := 0; i < 10; i++ {
		// defer functions are put on a stack (their execution is in LIFO order)
		defer fmt.Println(i)
	}

	fmt.Println("finish")
}

// recursive function, calculates the factorial of a number
func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

// receives a function as parameter
func multiplyRandomNumberBy3(generateRandom func() int) int {
	return generateRandom() * 3
}

// variadic function
func sayHiToEveryone(names ...string) {
	for _, name := range names {
		fmt.Println("Hi ", name)
	}
}
