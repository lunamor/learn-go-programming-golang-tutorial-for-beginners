package main

import (
	"fmt"
	"reflect"
)

func main() {

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)

	// using make()
	map2 := make(map[string]int)
	fmt.Println(map2)

	// get value
	fmt.Println(statePopulations["Ohio"])

	// add value
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations)

	// remove value
	delete(statePopulations, "California")
	fmt.Println(statePopulations)
	// California now has value 0 ***Caution***
	fmt.Println(statePopulations["California"])

	// 2nd return value, ok
	pop, ok := statePopulations["California"]
	fmt.Println(pop, ok) // ok is false because not in map

	// length
	fmt.Println(len(statePopulations))

	// maps are passed by reference, modifying can have side effects
	nums := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	nums2 := nums
	delete(nums2, 2)
	fmt.Println(nums)
	fmt.Println(nums2)

	// struct
	aDoctor := Doctor{
		number:    3,
		actorName: "Joh Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)

	// get field
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])

	// we can omit the field names, but it is not recommanded
	anotherDoc := Doctor{
		2,
		"Mike",
		[]string{
			"Ella",
			"Anna",
			"Lucy",
		},
	}
	fmt.Println(anotherDoc)

	// anonymous struct
	aPatient := struct{ name string }{name: "Patient 0"}
	fmt.Println(aPatient)

	// passed by copies
	anotherPatient := aPatient
	anotherPatient.name = "Tom"
	fmt.Println(aPatient)
	fmt.Println(anotherPatient)

	// we can use the address operator
	originalPatient := &aPatient
	originalPatient.name = "Sam"
	fmt.Println(aPatient)
	fmt.Println(originalPatient)

	// embedding
	b := Bird{}
	b.name = "Emu"
	b.origin = "Australia"
	b.speedKPH = 48
	b.canFly = false
	fmt.Println(b)

	anotherBird := Bird{
		Animal:   Animal{name: "Chicken", origin: "Earth"},
		speedKPH: 5,
		canFly:   false,
	}
	fmt.Println(anotherBird)

	//get the tag(just a string of text, something else will have to figure out what to do with it)
	s := reflect.TypeOf(Student{})
	field, _ := s.FieldByName("name")
	fmt.Println(field.Tag)
}

// must be able to check keys for equality

// return order of a map is not guaranteed

// struct
// collects any types of data
type Doctor struct {
	number     int // fields
	actorName  string
	companions []string
}

//naming convention: same as everything else

// embedding / composition
type Animal struct {
	name   string
	origin string
}

// Bird is not an Animal, it only has the characteristics
type Bird struct {
	Animal   // embedding
	speedKPH float32
	canFly   bool
}

// tags: space delimited subtags, key:"value"
// need reflection to get it
type Student struct {
	name string `required max:"100"`
	age  int
}
