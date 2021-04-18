package main

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

type one int

func (o *one) Increment() {
	log.WithField("one", *o).Info("incrementing one")
	*o++
}

func run(once *sync.Once, o *one, c chan bool) {
	once.Do(func() { o.Increment() })
	if v := *o; v != 1 {
		log.Errorf("once failed inside run: %d is not 1", v)
		return
	}
	c <- true
}

func main() {
	o := new(one)
	once := new(sync.Once)
	c := make(chan bool)

	const N = 10
	for i := 0; i < N; i++ {
		go run(once, o, c)
	}

	for i := 0; i < N; i++ {
		<-c
	}

	if *o != 1 {
		log.Errorf("once failed outside run: %d is not 1", *o)
	}
}
