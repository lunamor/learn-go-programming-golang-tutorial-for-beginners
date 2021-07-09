package main

import (
	"fmt"
)

func main() {

	// "Hello Go !" is the argument
	sayMessage("Hello Go!")
	for i := 0; i < 5; i++ {
		count("hello", i)
	}
	sayGreeting("Greetings", "friend")

	name := "Mike"
	changeName(name)
	fmt.Println(name)

	changeNameForReal(&name)
	fmt.Println(name)

	sum(1, 2, 3, 4, 5)

	sumReturn := returnSum(1, 2, 3, 4, 5, 6)
	fmt.Println("The sum is", sumReturn)

	sumReturnPointer := sumReturnPointer(1, 2, 3)
	fmt.Println("The sum is", *sumReturnPointer)

	fmt.Println("Sum is", sumNamedReturn(1, 2, 3, 4))

	d, err := divide(4.9, 0.0)
	if err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Println(d)

	// anonymous function
	func() {
		fmt.Println("Hello, Go!")
	}()

	// anonymous function with param
	for i := 0; i < 3; i++ {
		func(num int) {
			fmt.Println(num)
		}(i)
	}

	// function as variable
	f := func() {
		fmt.Println("Hey there")
	}
	f()

	// example with params and return values
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		}
		return a / b, nil
	}
	d2, err2 := divide(23.4, 2.5)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(d2)

	greeter := greeter{
		greeting: "Hello",
		name:     "Mike",
	}
	greeter.greet()
	fmt.Println(greeter.name)

	greeter.greetPolitely()
	fmt.Println(greeter.name)

	bitcoin := Bitcoin(3)
	bitcoin.toss()
}

// msg is a parameter
func sayMessage(msg string) {
	fmt.Println(msg)
}

// multiple params
func count(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
}

// the values passed to the params are copies
func changeName(name string) {
	name = "Dave"
	fmt.Println(name)
}

// we can use pointers to modify the values
func changeNameForReal(name *string) {
	*name = "Sam"
	fmt.Println(*name)
}

//variatic parameters (...) any amount of integers, values is a slice
// we can only have ONE variatic param, and it must be the last one
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is", result)
}

// function with return value
func returnSum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

// return a pointer
func sumReturnPointer(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

// named return value
func sumNamedReturn(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

// multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

// entry point of go application: package main, main function

// naming convention: same as everything else

// methods

type greeter struct {
	greeting string
	name     string
}

// value receiver: the receiver is the value greeter, which is a copy of the greeter object
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "No name"
}

// pointer receiver
func (g *greeter) greetPolitely() {
	fmt.Println(g.greeting, g.name, "are you well?")
	g.name = "Name Changed"
}

// method on a integer type
type Bitcoin int

func (b Bitcoin) toss() {
	fmt.Println("Tossing a Bitcoin")
}

// functions are types
// can be used as arguments, return values, ...
