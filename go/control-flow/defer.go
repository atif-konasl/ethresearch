package main

import "fmt"

func x() bool {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("defer in loop. i = %d\n", i)
	}
	fmt.Println("printing something before defer something...")
	return true
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func controller() {
	//x()
	b()
}
