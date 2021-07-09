package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// main function is executing in a goroutine
func main() {
	// tell go to spin a green thread and run sayHello in the green thread
	go sayHello()
	// main ended as soon as it spun a new goroutine, so sayHello didn't have time to print out its message
	time.Sleep(100 * time.Millisecond)

	// goroutine with anonymous function
	var msg = "Hello"
	// the function has access to variables in the outer scope
	go func() {
		fmt.Println(msg)
	}()
	time.Sleep(100 * time.Millisecond)

	// ***CAUTION***
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye"
	// the msg variables was modified before the msg had a change to be printed out (race condition)
	time.Sleep(100 * time.Millisecond)

	// solution: set msg as a parameter
	msg = "Hello"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)

	// not the best practice to use time.Sleep() for something like this

	msg = "hello"
	// we want to sync the goroutine of the main() with the one of the anonymous func
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		// decrement by one
		wg.Done()
	}(msg)
	msg = "Goodbye"
	// have the WaitGroup wait
	wg.Wait()

	for i := 0; i < 10; i++ {
		// let WaitGroup know we have 2 more goroutines running
		wg2.Add(2)
		go sayHello2()
		go increment2()
	}
	// make sure the main fct doesn't exit too early
	wg2.Wait()

	// result is a mess
	// goroutines are racing against each other

	for i := 0; i < 10; i++ {
		wg3.Add(2)
		go sayHello3()
		go increment3()
	}
	wg3.Wait()

	// move the Lock() and RLock() out
	// mutex is locked before the goroutines execute, and unlock when the goroutine is done
	// change the number of threads
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg4.Add(2)
		m4.RLock()
		go sayHello4()
		m4.Lock()
		go increment4()
	}

	// by default, go will give you the number of OS threads equal to the number of cores available on the machine
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	// too few threads == slow performance
	// too many == too much overhead, scheduler must work harder, etc
}

// WaitGroup: designed to synchronize multiple goroutines together
var wg = sync.WaitGroup{}

func sayHello() {
	fmt.Println("Hello")
}

var wg2 = sync.WaitGroup{}
var counter = 0

func sayHello2() {
	fmt.Println("Hello", counter)
	wg2.Done()
}

func increment2() {
	counter++
	wg2.Done()
}

// mutex

var wg3 = sync.WaitGroup{}
var counter3 = 0

// simple mutex: has a lock, can lock and unlock
// read write mutex: as many as want it can read, but only one can write at a time. If anything is reading, then we can't write to it at all
var m = sync.RWMutex{}

// protect the counter3 variable from concurrent reading and writing
func sayHello3() {
	// read lock
	m.RLock()
	fmt.Printf("Hello #%v\n", counter3)
	m.RUnlock()
	wg3.Done()
}

func increment3() {
	m.Lock()
	counter3++
	m.Unlock()
	wg3.Done()
}

var wg4 = sync.WaitGroup{}
var counter4 = 0
var m4 = sync.RWMutex{}

func sayHello4() {
	fmt.Printf("Hello #%v\n", counter4)
	m4.RUnlock()
	wg4.Done()
}

func increment4() {
	counter4++
	m4.Unlock()
	wg4.Done()
}

// messes up if sayHello3 execute twice in a row
// need to lock the mutex outside the context of the goroutines

// best practices

// don't create goroutines in libraries

// when creating a goroutine, know how it will end: avoid subtle memory leaks

// check for race conditions at compile time using go run -race
