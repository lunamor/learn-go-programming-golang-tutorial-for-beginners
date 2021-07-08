package main

import (
	"fmt"
	"math"
)

func main() {
	// if statements
	if true {
		fmt.Println("The test is true")
	}

	// the initializer syntax
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	// "pop, ok := statePopulations["Florida"]" is the initializer
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	// generating booleans
	// we can use <, >, ==, <=, >=, !=
	// logical operators: || (or), && (and), ! (not)
	// || can be short circuited if the first value is true
	// && if the first value is false
	// the rest won't be evaluated

	// else, if else
	x := -3
	if x == 0 {
		fmt.Println("Zero")
	} else if x > 0 {
		fmt.Println("Strictly positive")
	} else {
		fmt.Println("Strictly negative")
	}

	//***CAUTION*** working with floating points can be imprecise
	num := 0.123
	if num == math.Pow(math.Sqrt(num), 2) {
		fmt.Println("same")
	} else {
		fmt.Println("different")
	}

	// we can compare to an error threshold ***NOT PERFECT***
	if math.Abs(num/math.Pow(math.Sqrt(num), 2)-1) < 0.00001 {
		fmt.Println("same")
	} else {
		fmt.Println("different")
	}

	// switch statement
	switch i := 2 + 2; i {
	case 1:
		fmt.Println("one")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("others")
	}

	// test cases can't be the same

	// we can have an initializer

	// tagless syntax
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less or equal to ten")
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than 20")
	}

	// the first case that evaluate to true will execute

	// break is implied

	// if we want case to fallthrough: use fallthrough keyword
	switch {
	case i <= 10:
		fmt.Println("less or equal to ten")
		fallthrough
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than 20")
	}

	// type switch
	var j interface{} = "hello"
	switch j.(type) {
	case int:
		fmt.Println("j is an int")
	case float64:
		fmt.Println("j is a float64")
	case string:
		fmt.Println("j is a string")
		// use break to exit
		if j == "hello" {
			break
		}
		fmt.Println("this won't print")
	default:
		fmt.Println("j is another type")
	}

	// the interface{} type can take any type
}
