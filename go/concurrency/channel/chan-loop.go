package main

import (
	"fmt"
	"sync"
)

type work struct {
	loopNum   int
}
var wg sync.WaitGroup

func loop1(wc chan *work, exitCh chan *struct{}) {
	fmt.Println("started loop1")
	for {
		select {
		case b := <-wc:
			fmt.Printf("got notification in loop1 no=%d\n", b.loopNum)
		case _, ok := <-exitCh:
			fmt.Printf("aborting loop1.. ok=%t\n", ok)
			wg.Done()
			return
		}
	}
}

func loop2(wc chan *work, exitCh chan *struct{}) {
	fmt.Println("started loop2")
	for {
		select {
		case b := <-wc:
			fmt.Printf("got notification in loop2 no=%d\n", b.loopNum)
		case _, ok := <-exitCh:
			fmt.Printf("aborting loop2.. ok=%t\n", ok)
			wg.Done()
			return
		}
	}
}

func Controller(count int) {
	wc := make(chan *work)
	exitCh := make(chan *struct{})
	wg.Add(2)

	go loop1(wc, exitCh)
	go loop2(wc, exitCh)

	fmt.Println("start sending notifications...")
	for i := 0; i < count; i++ {
		if i % 2 == 0 {
			wc <- &work{
				loopNum: i,
			}
		} else {
			wc <- &work{
				loopNum: i,
			}
		}
	}
	close(exitCh)
	wg.Wait()
	fmt.Println("aborting from main...")
}