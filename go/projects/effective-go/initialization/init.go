package main

import (
	"fmt"
	"os"
)

var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

func init() {
	fmt.Printf("%v\n", home)
	fmt.Printf("%v\n", user)
	fmt.Printf("%v\n", gopath)
}

func initController() {
	fmt.Println("should print this line after init function")
}