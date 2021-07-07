package main

import (
	"fmt"
	"strconv"
)

// at the package level, can't use :=
var a int = 99

// lowercase: visible package
var b int = 98

// Uppercase: visible to the world
var D int = 97

// declare variables in a block
var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

func main() {
	// declare but initialize later
	var i int
	i = 43

	// if go can't figure out the type
	var j float32 = 42

	// visible in block
	k := 41

	// shadowing: the variable in the innermost scope take precedence
	fmt.Println(a)
	var a int = 40

	//convertion
	var l = float32(a)

	// to convert strings we need the strconv package
	var m string = strconv.Itoa(i)

	fmt.Println(i, j, k, a, l, m)
}

// variables must always be used

// short variable name in small scopes (eg. i)
// longer names for longer lifespan

// acronyme should be all uppercase (eg. theURL, theHTTPRequest)
