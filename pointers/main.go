package main

import (
	"fmt"
)

func main() {
	var a int = 42
	// copy a and assign to b
	b := a
	a = 27
	fmt.Println(a)
	fmt.Println(b)
	// *  (pointer to int), to address of a
	// c holds the memory address that's holding the data for a
	var c *int = &a
	fmt.Println(c)
	fmt.Println(&a)

	// dereferencing operator
	// * (dereference), find the memory loc and get the value
	fmt.Println(*c)

	// both will change if we modify the value of a
	a = 10
	fmt.Println(a, *c)

	// we can also modify the dereferenced value
	*c = 15
	fmt.Println(a, *c)

	// there is no pointer arithmetic in go
	d := [3]int{1, 2, 3}
	e := &d[0]
	f := &d[1]
	fmt.Printf("%v %p %p\n", d, e, f) // e and f are 4 bytes apart
	// ***THIS IS NOT ALLOWED IN GO*** (though it is in the "unsafe" package )
	// f = &d[1] - 4

	var ms *myStruct
	ms = &myStruct{foo: 43}
	fmt.Println(ms)

	var ms2 *myStruct
	// new()
	ms2 = new(myStruct)
	fmt.Println(ms2)
	// to get the field
	(*ms2).foo = 13
	fmt.Println((*ms2).foo)
	// we don't actually have to do that (syntaxtic sugar)
	ms2.foo = 14
	fmt.Println(ms2.foo)

}

type myStruct struct {
	foo int
}

// uninitialize pointers are nil

// all assignment operations in go are copy operations

// slice is a pointer pointing the the 1st element in the underlying array
// same with map
