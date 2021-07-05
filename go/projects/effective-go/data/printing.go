package main

import "fmt"

type T struct {
	a int
	b float64
	c string
}

func printBytesToHex() {
	bytes := []byte("printing bytes to hexadecimal")
	fmt.Printf("%v\n", bytes)
	fmt.Printf("%x\n", bytes)
}

func printType() {
	x := 12
	fmt.Printf("%T\n", x)
}

func (t *T) String() string {
	return fmt.Sprintf("%d____%g____%q", t.a, t.b, t.c)
}

func modifiedFormat() {
	t := &T{ 7, -2.35, "abc" }
	fmt.Printf("%v\n", t)
}

func appendSliceToSlice() {
	x := []int{1,2,3}
	y := []int{4,5,6}
	x = append(x, y...)
	fmt.Println(x)
}

func printController() {
	printBytesToHex()
	printType()
	modifiedFormat()
	appendSliceToSlice()
}