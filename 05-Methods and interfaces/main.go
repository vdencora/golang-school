package main

import "fmt"

type Hero struct {
	Name string
	Age  int
	City string
}

// Figure is an interface which has two method signatures
type Figure interface {
	area() float64
	perimeter() float64
}

// square and rectangle will implement the Figure interface just making use of area and perimeter methods
type Square struct {
	Side float64
}

type Rectangle struct {
	Base   float64
	Height float64
}

func main() {
	/* -------------------------------------------------------------------------- */
	/*                                  METHODS                                   */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "METHODS")

	s := Hero{
		Name: "Spiderman",
		Age:  18,
		City: "New York",
	}

	// calling the GetProfile method
	profile := s.GetProfile()
	fmt.Println("Profile Method: ", profile)

	fmt.Printf("\n******** %s ******\n", "Method pointer receivers vs value receivers")

	// look that SetCity has a value receiver and SetAge a pointer receiver.
	// See the effect these methods have on the original caller

	// With value receiver
	fmt.Println("hero before change city: ", s)
	s.SetCity("California")
	fmt.Println("hero after change city:: ", s)

	// With pointer receiver
	fmt.Println("hero before change age: ", s)
	s.SetAge(28)
	fmt.Println("hero after change age:: ", s)

	/* -------------------------------------------------------------------------- */
	/*                                INTERFACES                                  */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "INTERFACES")

	// create new instances of square and rectangle (line 18)
	square := Square{Side: 5}
	rectangle := Rectangle{Base: 4, Height: 8}

	fmt.Println("Square area: ", square.area())
	fmt.Println("Square perimeter: ", square.perimeter())

	fmt.Println("Rectangle area: ", rectangle.area())
	fmt.Println("Rectangle perimeter: ", rectangle.perimeter())

	/* -------------------------------------------------------------------------- */
	/*                            Type Assertions                                */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "Type Assertions")

	figureName := ""

	// Feel free to assign to `f` the variable `square`, since they both implement the Figure interface, everything should work fine.
	var f Figure = rectangle

	// make an assertion to know the type of the Figure; here to know more about assertions -> https://go.dev/tour/methods/15
	switch f.(type) {
	case Square:
		figureName = "Square"
	case Rectangle:
		figureName = "Rectangle"
	default:
		figureName = "Unknow"
	}

	fmt.Println("Figure is: ", figureName)

	/* -------------------------------------------------------------------------- */
	/*                             Empty Interface                                */
	/* -------------------------------------------------------------------------- */

	fmt.Printf("\n--------------------- %s -----------------------------\n", "Empty Interface")

	// An interface that has zero methods is called an empty interface.
	// It is represented as interface{}. Since the empty interface has zero methods, all types implement the empty interface.
	var i interface{}

	// see how we can assign different types to the same variable of type interface
	i = "Hola, soy un string"
	fmt.Println("interface value: ", i)

	i = 10
	fmt.Println("interface value: ", i)

	i = true
	fmt.Printf("interface value: %v \n\n", i)

	// describeType recives an empty interface as argument, so we could send any type as parameter
	describeType("hola mundo")
	describeType(33)
	describeType(true)
	describeType(4.5)
}

/* -------------------------------------------------------------------------- */
/*                                  METHODS                                   */
/* -------------------------------------------------------------------------- */

// A method is just a function with a special receiver type between the func keyword and the method name. The receiver can either be a struct type or non-struct type.
func (h Hero) GetProfile() string {
	return fmt.Sprintf("I'm %s and I'm from %s", h.Name, h.City)
}

func (h *Hero) SetAge(age int) {
	h.Age = age
}

func (h Hero) SetCity(city string) {
	h.City = city
}

/* -------------------------------------------------------------------------- */
/*                            INTERFACE METHODS                               */
/* -------------------------------------------------------------------------- */

func (s Square) area() float64 {
	return s.Side * s.Side
}

func (s Square) perimeter() float64 {
	return s.Side * 4
}

func (r Rectangle) area() float64 {
	return r.Base * r.Height
}

func (r Rectangle) perimeter() float64 {
	return (r.Base * float64(2)) + (r.Height * float64(2))
}

func describeType(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}
