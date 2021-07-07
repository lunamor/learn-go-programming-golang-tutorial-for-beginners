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
	_             = iota // underscore: write only, we don't care about the value
	catSpecialist = iota
	// we can operate
	dogSpecialist  = iota + 2
	birdSpecialist // infer the pattern
)

// use iota for units
const (
	_  = iota // ignore the first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
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

	fileSize := 4000000000.
	fmt.Printf("%0.2fGB\n", fileSize/GB)

	// store info efficiently
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)
}

// naming convention: same as variables

// immutable

// values must be calculable at compile time
