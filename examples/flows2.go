package main

import "fmt"

func main() {
	fmt.Println("Calling f.")
	f()
	fmt.Println("Returned normally from f.")
}

func handlePanic(message string) {
	if r := recover(); r != nil {
		fmt.Println(message, r)
	}
}

func f() {
	defer handlePanic("Recovered in f")
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
		fmt.Println("This won't print")
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
