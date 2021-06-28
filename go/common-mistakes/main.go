package main

import "fmt"

func main() {
	ptrTest := definePointer()
	fmt.Printf("num: %d\nstr: %s", ptrTest.number, ptrTest.str)
}
