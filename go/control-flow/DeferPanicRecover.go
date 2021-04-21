package main

import "fmt"

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f:", r)
		}
	}()

	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
		//return
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func panicfunc(i int) {
	if i > 5 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
		fmt.Println("After panicking")
	}
	fmt.Println("Printing in f", i)
}

func main() {
	//f()
	//fmt.Println("Returned normally from f.")
	panicfunc(5)
}