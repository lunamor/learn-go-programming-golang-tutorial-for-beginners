package main

import (
	"fmt"
)

const myConst bool = false

//enumerated constants: counters, changes value with each iota
const c = iota

// in blocks
const (
	// iota is scoped to the block
	_              = iota // underscore: write only, we don't care about the value
	catSpecialist  = iota
	dogSpecialist  = iota
	birdSpecialist // infer the pattern
)

func main() {
	// can shadow
	fmt.Printf("%v, %T\n", myConst, myConst)
	// typed constant
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	// can't change the value

	// can have any primative type constants

	// can add with variables
	var b int = 27
	fmt.Printf("%v, %T\n", myConst+b, myConst+b)

	// infered type
	const a = 32
	fmt.Printf("%v, %T\n", a, a)

	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", catSpecialist, catSpecialist)
	fmt.Printf("%v, %T\n", dogSpecialist, dogSpecialist)
	fmt.Printf("%v, %T\n", birdSpecialist, birdSpecialist)
}

// naming convention: same as variables
