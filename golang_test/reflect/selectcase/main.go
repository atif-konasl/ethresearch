package main

import (
"log"
"reflect"
"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan string)
	c3 := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(50 * time.Millisecond)
			c1 <- i
			log.Println("c2 <- ", <-c2)
			c3 <- true
		}
	}()

	cases := []reflect.SelectCase{
		{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(c1),
		},
		{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(c2),
			Send: reflect.ValueOf("Hello"),
		},
		{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(c3),
		},
		{
			Dir: reflect.SelectDefault,
		},
	}
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		switch index, value, recvOK := reflect.Select(cases); index {
		case 0:
			// c1 was selected, recv is the value received
			// we get recvOK whether we need it or not
			log.Println("case 0:", value, recvOK)
		case 1:
			// c2 was selected, therefore "hello" was sent
			// recv and recvOK are garbage
			log.Println("case 1: There was a read on the channel ready")
		case 2:
			// c3 was selected, value is good if recvOK == true
			// recvOK is false if c3 is closed
			log.Println("case 2:", value, recvOK)
		case 3:
			// default case
			// recv and recvOK are useless here
			log.Println("case 3: reflect.SelectDefault,")
		}
	}
}
