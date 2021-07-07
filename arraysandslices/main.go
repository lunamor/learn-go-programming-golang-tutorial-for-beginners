package main

import (
	"fmt"
)

func main() {
	// accessing is fast since data stored contiguously
	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	// we can use ...
	grades2 := [...]int{32, 76, 85}
	fmt.Printf("Grades: %v\n", grades2)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("Students: %v\n", students)

	// access element
	fmt.Printf("Student #1: %v\n", students[1])

	// size
	fmt.Printf("Number of Students: %v\n", len(students))

	// array of arrays
	var identitymatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}

	fmt.Printf("Identity matrix: %v\n", identitymatrix)

	// arrays are copied
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	// pass an address
	c := &a
	c[2] = 9
	fmt.Println(a)
	fmt.Println(c)

	// slices
	l := []int{1, 2, 3}
	fmt.Println(l)

	// length
	fmt.Printf("Length: %v\n", len(l))
	// capacity
	fmt.Printf("Capacity: %v\n", cap(l))

	// slices are reference types
	m := l
	m[1] = 4
	fmt.Printf("%v\n", l)
	fmt.Printf("%v\n", m)

	// we can slice slices and arrays
	d := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	e := d[:]   // slice of all elements
	f := e[3:]  // slice from 4th element to end
	g := d[:6]  // slice first 6 elements
	h := d[3:6] // slice the 4th, 5th, and 6th elements
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	j := make([]int, 3, 100) // type, length, capacity
	fmt.Println(j)
	fmt.Printf("Length: %v\n", len(j))
	fmt.Printf("Capacity: %v\n", cap(j))

	// append to slice
	j = append(j, 1, 4, 7, 9)
	fmt.Println(j)
	fmt.Printf("Length: %v\n", len(j))
	fmt.Printf("Capacity: %v\n", cap(j))

	// spread the slice out into individual arguments(using ...)
	j = append(j, []int{-1, -3, -6}...)
	fmt.Println(j)

	// removing elements
	x := []int{1, 2, 3, 4, 5}
	// remove first
	fmt.Println(x[1:])
	// remove last
	fmt.Println(x[:len(x)-1])
	// remove from middle
	fmt.Println(append(x[:2], x[3:]...))
	// the underlying array is modified ***BEWARE***
	fmt.Println(x)
}

// arrays have a fixed size, slices do not

// elements must be the same type

// slices are backed by arrays
