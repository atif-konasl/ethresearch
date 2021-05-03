package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
)

// This `ping` function only accepts a channel for sending
// values. It would be a compile-time error to try to
// receive on this channel.
func ping(pings chan<- string, msg string) {
	chanval := reflect.ValueOf(pings)
	chantyp := chanval.Type()
	log.WithField("chanValue", chanval).WithField(
		"chanDir", chantyp.ChanDir()).WithField(
		"chanName", chantyp.Name()).Info("channel info")
	//chanInfo(ping)
	pings <- msg
}

// The `pong` function accepts one channel for receives
// (`pings`) and a second for sends (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
	chanval := reflect.ValueOf(pings)
	chantyp := chanval.Type()
	log.WithField("chanValue", chanval).WithField(
		"chanDir", chantyp.ChanDir()).WithField(
		"chanName", chantyp.Name()).Info("channel info")
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

