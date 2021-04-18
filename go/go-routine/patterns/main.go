package main

import "fmt"

func main() {

	// For generator pattern
	joe := boring("Joe")
	alice := boring("Alice")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-alice)
	}

	// For multiplexing pattern
	//joe := boring("Joe")
	//alice := boring("Alice")
	//anyone := fanInString(joe, alice)
	//
	//for i := 0; i < 5; i++ {
	//	fmt.Println(<-anyone)
	//}

	// For sequencing pattern
	//anyone := fanInMsg(boringMsg("Joe"), boringMsg("Alice"))
	//for i := 0; i < 50; i++ {
	//	msg1 := <-anyone; fmt.Println(msg1.str)
	//	msg2 := <-anyone; fmt.Println(msg2.str)
	//
	//	msg1.wait <- true
	//	msg2.wait <- true
	//}

	fmt.Println("You're both boring; I'm leaving!!")
}
