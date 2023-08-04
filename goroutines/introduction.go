//A Goroutine is a function or method which executes independently and simultaneously
//in connection with any other Goroutines present in your program. Or in other words,
//every concurrently executing activity in Go language is known as a Goroutines.
//Every program contains at least a single Goroutine and that Goroutine is known as
//the main Goroutine. All the Goroutines are working under the main Goroutines if the
//main Goroutine terminated, then all the goroutine present in the program also terminated.

package main

import "fmt"

func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}
