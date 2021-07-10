package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	// must use make() to create a channel
	// channels are strongly typed: they can only send and receive msgs of that type
	ch := make(chan int)
	wg.Add(2)
	go func() {
		// will receive data from the channel (using <-)
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		// will send data to the channel (using <-)
		ch <- 42
		wg.Done()
	}()
	wg.Wait()

	fmt.Println()

	// multiple goroutines
	ch = make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			// this line will pause the execution of this goroutine until there is space in the channel, only one msg can be in the channel at one time
			ch <- 42
			wg.Done()
		}()
	}
	wg.Wait()

	// both goroutines are both sending and receiving
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		ch <- 27
		wg.Done()
	}()
	go func() {
		ch <- 15
		fmt.Println(<-ch)
		wg.Done()
	}()
	wg.Wait()

	wg.Add(2)
	// receive only channel
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	// send only channel
	go func(ch chan<- int) {
		ch <- 21
		wg.Done()
	}(ch)
	wg.Wait()

	// buffered channel
	// internal data store that can store 50 integers
	ch2 := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)
		wg.Done()
	}(ch2)
	go func(ch chan<- int) {
		ch <- 32
		ch <- 33
		wg.Done()
	}(ch2)
	wg.Wait()

	// using for range loop to receive all the data
	wg.Add(2)
	go func(ch <-chan int) {
		// note: different syntax, only pull out value, not the key
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch2)
	go func(ch chan<- int) {
		ch <- 10
		ch <- 11
		// we need to close the channel to signal that no more msgs are coming, channel is closed
		close(ch2)
		wg.Done()
	}(ch2)
	wg.Wait()

	ch3 := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {
		for {
			// get value and boolean ok (true if open, false if closed)
			// works the same as the previous example (for range)
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch3)
	go func(ch chan<- int) {
		ch <- 55
		ch <- 54
		close(ch3)
		wg.Done()
	}(ch3)
	wg.Wait()

	// used to monitor log message coming in channel
	// go logger()
	go loggerSelect()
	// we can add a defer
	// defer func() {
	// 	close(logCh)
	// }()
	// logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	// logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	logCh2 <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh2 <- logEntry{time.Now(), logInfo, "App is shutting down"}
	// application shut down when the main() finishes, logger() goroutine is torn down forcefully, generally don't want this, can leak resourse
	time.Sleep(100 * time.Millisecond)

	// send msg to doneCh2 when we want the logger to shut down
	doneCh2 <- struct{}{}
}

// channels can pass data between different goroutines in a way that is safe and prevent issues such as race and memory sharing problems

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)

func logger() {
	for entry := range logCh {
		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	}
}

var logCh2 = make(chan logEntry, 50)

// struct with no field require 0 memory allocation
// can't send any data through it, except whether a msg has been sent or receive, it is called a signal only channel
var doneCh2 = make(chan struct{})

func loggerSelect() {
	for {
		// select: entire statement will block until a msg is receive on one channel it's listening for
		select {
		case entry := <-logCh2:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh2:
			break
		}
	}
}

// we can also have a default for the select to make a non blocking select

// select is like a switch for channels to monitor channels
