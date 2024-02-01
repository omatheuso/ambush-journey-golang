package main

import "fmt"

func main() {
	fmt.Println("Calling f.")
	f()
	fmt.Println("Returned normally from f.")
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func f() {
	defer handlePanic()
	i := 2
	if i > 3 {
		panic("Panicking")
	}
	fmt.Println("Return normally from f.")
}
