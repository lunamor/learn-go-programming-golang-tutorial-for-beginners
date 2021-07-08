package main

import (
	"fmt"
)

func main() {
	// for loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// multiple initializers
	for i, j := 0, 5; i < 5; i, j = i+1, j*5 {
		fmt.Println(i, j)
	}

	// we don't need all 3 statements

	// missing 1st statement
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	// 1st missing 3nd statement
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// infinite loop, missing condition
	for {
		fmt.Println(i)
		i++
		if i == 20 {
			break // leaves the loop
		}
	}

	// continue: exit this iteration of the loop
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	// nested loops
	for i := 1; i < 6; i++ {
		for j := 1; j < 6; j++ {
			fmt.Println(i * j)
		}
	}

	// labels
outer:
	for i := 1; i < 6; i++ {
		for j := 1; j < 6; j++ {
			fmt.Println(i * j)
			if i*j >= 15 {
				break outer
			}
		}
		fmt.Println("outer")
	}

	// work with collections
	s := []int{1, 2, 3}
	fmt.Println(s)

	// for range loop, returns key and value
	for k, v := range s {
		fmt.Println(k, v)
	}

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}

	// _ don't care about the key
	str := "hello"
	for _, v := range str {
		fmt.Println(string(v))
	}
}
