package main

import (
	"fmt"
)

func main() {
	// boolean
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	m := 1 == 2
	fmt.Printf("%v, %T\n", m, m)

	// zero value for bool is false

	// numeric types
	// zero value is 0

	// int8, int16, int32, int64
	a := 43
	fmt.Printf("%v, %T\n", a, a)

	// unsigned int: uint8, uint16, uint32
	var b uint16 = 42
	fmt.Printf("%v, %T\n", b, b)

	// a byte is a uint8

	// operations
	x := 10
	y := 3
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	// we can only operate on integers of the same type
	fmt.Println(uint16(a) + b)

	// bit operators
	fmt.Println("bit operation")
	fmt.Println(x & y)  // 1010 & 0011 -> 0010 -> 2
	fmt.Println(x | y)  // 1010 | 0011 -> 1011 -> 11
	fmt.Println(x ^ y)  // (xor) -> 1001 -> 9
	fmt.Println(x &^ y) // (and not) -> 0100 -> 8

	// bit shifting
	z := 8 // 1000
	// shift to the left: multiply by 2
	fmt.Println(z << 3) // 0100 0000 -> 64
	// shift to the right: divide by 2
	fmt.Println(z >> 3) // 0001 -> 1

	// float32, float64
	var w float64 = 3.14
	w = 13.7e72 // exponential notation
	w = 2.1e14  // exponential notation
	fmt.Printf("%v, %T\n", w, w)

	// arithmetic operations: + - * /

	// complex numbers (complex64, complex128)
	var v complex64 = 1 + 2i
	var u complex64 = 2 + 5.2i
	fmt.Printf("%v, %T\n", v, v)

	// operations
	fmt.Println(v + u)
	fmt.Println(v - u)
	fmt.Println(v * u)
	fmt.Println(v / u)

	// zero value: 0

	// get the real part or imaginary part
	fmt.Printf("%v, %T\n", real(v), real(v))
	fmt.Printf("%v, %T\n", imag(v), imag(v))

	// make a complex number
	var t complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", t, t)

	// zero value: (0+0i)

	// text types

	// strings: collections UTF-8 characters
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)
	// can treat them like arrays
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	// strings are immutable (can't do s[2] = "s")

	// concatenation
	s2 := "another string"
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	// can create slices of bytes
	slice := []byte(s)
	fmt.Printf("%v, %T\n", slice, slice)

	// rune: UTF-32 character - it is a int32 (uses single quotes)
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)

	// zero value: "" or ''
}
