package main

import "fmt"

// Why Do You Need to Be Careful With Loop Variable in Go
func main() {
	list := make([]int, 0)
	for i := 0; i < 10; i++ {
		list = append(list, i)
	}

	ch := make(chan int)
	// Here the address of loop variable, i just muted its value and not change its address.
	for i := range list {

		go func(x int) {
			fmt.Printf("addr of i: %x and value of i: %d\n", &x, x)
			ch <- x
		}(i)

		// Race conditions when taking reference of loop variable and passing it to another goroutine:

		//go func() {
		//	fmt.Printf("addr of i: %x and value of i: %d\n", &i, i)
		//	ch <- i
		//}()
	}

	for i := 0; i < 10; i++ {
		<- ch
	}
}


