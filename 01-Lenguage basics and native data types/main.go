package main

import (
	"fmt"
	"strings"
)

func main() {
	/* -------------------------------------------------------------------------- */
	/*                          VARIABLES AND CONSTANTS                           */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "VARIABLES AND CONSTANTS")

	// ways to declare a variable
	var age1 int = 20                   // with var keyword
	age2 := 20                          // short variable declaration
	var position, age = "Developer", 28 // multiple declaration

	fmt.Println("Age 1: ", age1)
	fmt.Println("Age 2: ", age2)
	fmt.Println("I'm a "+position+" and my age is ", age)

	// constant declaration
	const limit = 100
	const minimun = 200

	fmt.Println("Limit is a constant: ", limit)
	fmt.Println("Minimun is a constant: ", minimun)

	// A constant is a value that cannot be changed at runtime.
	// Uncomment the line below and run the program again to see the error
	// limit = 300

	/* -------------------------------------------------------------------------- */
	/*                     BASIC TYPES OPERATORS AND EXPRESSIONS                  */
	/* -------------------------------------------------------------------------- */
	fmt.Printf("\n--------------------- %s -----------------------------\n", "BASIC TYPES OPERATORS AND EXPRESSIONS")

	fmt.Printf("\n******** %s ******\n", "Operations")
	num1 := 10
	num2 := 3

	// addition
	addition := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num2, addition)

	// subtraction
	subtraction := num1 - num2
	fmt.Printf("%d - %d = %d\n", num1, num2, subtraction)

	// multiplication
	multiplication := num1 * num2
	fmt.Printf("%d * %d = %d\n", num1, num2, multiplication)

	// division
	division := num1 / num2
	fmt.Printf("%d / %d = %d\n", num1, num2, division)

	// modulo operator
	modulo := num1 % num2
	fmt.Println(num1, " % ", num2, " = ", modulo)

	// float division
	var a float32 = 5.0
	var b float32 = 3.0
	floatDivision := a / b
	fmt.Printf("%f / %f = %f\n", a, b, floatDivision)

	num := 5
	// increment of num by 1
	num++
	fmt.Printf("num++ = %d\n", num)

	// decrement of num by 1
	num--
	fmt.Printf("num-- = %d\n", num)

	fmt.Printf("\n******** %s ******\n", "Compound Assignment Operators")
	// Compound Assignment Operators
	number := 2
	number += 6
	// Here, += is additional assignment operator. It first adds 6 to the value of number (2) and assigns final result (8) to number.
	fmt.Println("number is: ", number)

	fmt.Printf("\n******** %s ******\n", "Relational operators")
	// Relational operators
	string1 := "mundo"
	string2 := "world"
	value1 := 55
	value2 := 33

	// == (equal to) returns true if value1 and value1 are equal
	fmt.Printf("%s == %s ? %t\n", string1, string2, string1 == string2)

	// != (equal to) returns true if value1 and value1 are not equal
	fmt.Printf("%s != %s ? %t\n", string1, string2, string1 != string2)

	// > (greater than) returns true if value1 is greater than value2
	fmt.Printf("%d > %d ? %t\n", value1, value2, value1 > value2)

	// < (less than) returns true if value1 is less than value2
	fmt.Printf("%d < %d ? %t\n", value2, value1, value2 < value1)

	// >= (greater than or equal to) returns true if value1 is either greater than or equal to value2
	fmt.Printf("%d >= %d ? %t\n", value1, value2, value1 >= value2)

	// <= (less than or equal to) returns true if value1 is either less than or equal to value2
	fmt.Printf("%d <= %d ? %t\n", value2, value1, value2 <= value1)

	fmt.Printf("\n******** %s ******\n", "Logical operators")
	// Logical operators
	bool1 := true
	bool2 := false

	// && (Logical AND) returns true if both expressions exp1 and exp2 are true
	fmt.Printf("%t && %t ? %t\n", bool1, bool2, bool1 && bool2)

	// || (Logical OR) returns true if any one of the expressions is true
	fmt.Printf("%t || %t ? %t\n", bool1, bool2, bool1 || bool2)

	// ! (Logical NOT) returns true if exp is false and returns false if exp is true
	fmt.Printf("%t && !%t ? %t\n", bool1, bool2, bool1 && !bool2)

	/* -------------------------------------------------------------------------- */
	/*                                  STRINGS                                   */
	/* -------------------------------------------------------------------------- */
	/*
		A string is a slice of bytes in Go. Strings can be created by enclosing a set of characters inside " " or ` `.
	*/
	fmt.Printf("\n--------------------- %s -----------------------------\n", "STRINGS")

	s1 := "Hello world"
	s2 := `Hola mundo`

	fmt.Println("string 1 is: ", s1)
	fmt.Println("string 2 is: ", s2)

	// a string can be concatenated
	fmt.Println("s1 + s2 = ", s1+s2)

	// to know the length of a string
	fmt.Println("s1 length = ", len(s1))

	// strings package implements functions to facilitate working with strings
	// to convert to upperCase
	fmt.Println("s1 ToUpper = ", strings.ToUpper(s1))
	// to convert to lowerCase
	fmt.Println("s2 ToLower = ", strings.ToLower(s2))

	// visit https://pkg.go.dev/strings to learn more about strings

	/* -------------------------------------------------------------------------- */
	/*                           POINTER AND ADDRESSES                            */
	/* -------------------------------------------------------------------------- */
	fmt.Printf("\n--------------------- %s -----------------------------\n", "POINTER AND ADDRESSES")

	// pointer declaration
	var ptr *string

	// the init valur of a pointer in NIL (always keep this in mind to avoid a nil pointer reference error)
	fmt.Println("init valur of ptr", ptr)

	text := "this is an init text"
	ptr = &text

	// dereferencing (* syntax) a pointer means accessing the value of the variable to which the pointer points.
	fmt.Println("address of ptr is: ", ptr)
	fmt.Println("value of ptr is: ", *ptr)

	// update the value of a pointer will take effect on all variables that point to the same memory address
	*ptr = "new text from pointer"
	fmt.Println("value of ptr is: ", *ptr)
	// see how the text variable also changed
	fmt.Println("value of text var is: ", text)
}
