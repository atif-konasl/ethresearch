package main

func fanInString(input1, input2 <-chan string) <- chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

//func main() {
//	//For multiplexing pattern
//	joe := boring("Joe")
//	alice := boring("Alice")
//	anyone := fanInString(joe, alice)
//
//	for i := 0; i < 5; i++ {
//		fmt.Println(<-anyone)
//	}
//}