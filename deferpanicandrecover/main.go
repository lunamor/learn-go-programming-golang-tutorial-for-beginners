package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// defer
	fmt.Println("start")
	fmt.Println("middle")
	// defer: will execute after the function is done, but before it return the result to the calling function
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	// executed in LIFO: last in first out
	fmt.Println("end")

	// example with http calls
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	// defer to close resourse, so not to forget
	defer res.Body.Close()
	//ReadAll takes in a stream and parse it out as a []byte
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", robots)

	// it takes the value at the time the line is reached
	a := "start"
	defer fmt.Println(a)
	a = "end"

	// panic
	// one, zero := 1, 0
	// ans := one / zero
	// fmt.Println(ans)

	// to panic
	// order of execution: main fct -> defer -> panic -> return value
	// defer fmt.Println("this was deferred")
	// panic("something bad happened")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Go!"))
	// })
	// err2 := http.ListenAndServe(":8080", nil)
	// if err2 != nil {
	// 	panic(err2.Error())
	// }

	// recover - similar to catch
	fmt.Println("start recover")
	panicker()
	// this will still execute
	fmt.Println("end recover")
}

// the defer statement takes a function call

// recover
func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			// if we don't want to recover, rethrow panic
			panic(err)
		}
	}()
	panic("Uh oh")
	fmt.Println("done panicking")
}

// panic: program will exit

// recover: only useful in deferred functions
// only higher fcts in call stack will continue
