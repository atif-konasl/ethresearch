package main

import (
	"fmt"
	"math/rand"
)
// reverse method reverse the slice
func reverse(a []int) {
	len := len(a)
	for i := 0; i < len/2; i ++ {
		x := a[len-i-1]
		a[len-i-1] = a[i]
		a[i] = x
	}
}

// print method prints all ints of the slice
func print(a []int) {
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

// fillSliceWithRand method fills the slice with random ints
func fillSliceWithRand(a []int) {
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(100)
	}
}

func sliceController() {
	slice := make([]int, 10)
	fillSliceWithRand(slice)
	print(slice)
	reverse(slice)
	print(slice)
}
