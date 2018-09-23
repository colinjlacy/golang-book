package main

import "fmt"

type ByteCounter int

func main() {
	var counter ByteCounter

	counter.Write([]byte("hello"))
	fmt.Println(counter)

	counter = 0

	name := "Dolly"

	fmt.Fprintf(&counter, "hello, %s", name)
	fmt.Println(counter)
}

func (byteCounter *ByteCounter) Write(p []byte) (int, error) {
	*byteCounter += ByteCounter(len(p))
	return len(p), nil
}
