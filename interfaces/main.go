package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// we can use ConsoleWriter as a Writer
	// polymorphic behavior
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello YouTube listeners, this is a test"))
	wc.Close()

	// type conversion, convert wc to *BufferedWriterCloser, and store in bwc
	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)

	// trying to conver to io.Reader
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

	// empty interface: interface defined on the fly that has no methods
	// everything can be cast to it, even primitives
	var myObj interface{} = NewBufferedWriterCloser()
	if wc2, ok2 := myObj.(WriterCloser); ok2 {
		wc2.Write([]byte("Hello YouTube listeners, this is a test"))
		wc2.Close()
	}
	r2, ok2 := myObj.(io.Reader)
	if ok2 {
		fmt.Println(r2)
	} else {
		fmt.Println("Conversion failed")
	}

	var myInterface interface{} = false
	// type switch
	switch myInterface.(type) {
	case int:
		fmt.Println("myInterface is an integer")
	case string:
		fmt.Println("myInterface is a string")
	default:
		fmt.Println("I don't know what myInterface is")
	}
}

// interface describes behavior
type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

// implecit implementation
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// interface for a integer alias
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// naming convention: should name it by what it does
// for single method interface: method name + er

type Closer interface {
	Close() error
}

// interface composed of other interfaces
type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}

	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// constructor function
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

// when implementing an interface, if we use a value type, the methods that implement the interface have to all have value receivers.
// if we implement the interface with a pointer, then we just have to have the methods there, regardless of the receiver type

// best practices

// use many small interfaces
// don't export interfaces for types that will be consumed
// do export interfaces for types that will be used by package
// design functions and methods to receive interfaces whenever possible
