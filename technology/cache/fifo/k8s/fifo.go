package main

import (
	"fmt"
	"github.com/pkg/errors"
	fifoCache "k8s.io/client-go/tools/cache"
	"sync"
)

func testFifoObjectKeyFunc(obj interface{}) (string, error) {
	return obj.(testFifoObject).name, nil
}

type testFifoObject struct {
	name string
	val  interface{}
}

func mkFifoObj(name string, val interface{}) testFifoObject {
	return testFifoObject{name: name, val: val}
}

func fifo() error {
	wg := new(sync.WaitGroup)
	f := fifoCache.NewFIFO(testFifoObjectKeyFunc)
	const amount = 50

	wg.Add(2)
	go func() {
		for i := 0; i < amount; i++ {
			f.Add(mkFifoObj(string([]rune{'a', rune(i)}), i+1))
		}
		wg.Done()
	}()
	go func() {
		for u := uint64(0); u < amount; u++ {
			f.Add(mkFifoObj(string([]rune{'b', rune(u)}), u+1))
		}
		wg.Done()
	}()

	wg.Wait()
	lastInt := int(0)
	lastUint := uint64(0)
	for i := 0; i < amount*2; i++ {
		fmt.Printf("size of fifo : %d\n", len(f.ListKeys()))
		switch obj := fifoCache.Pop(f).(testFifoObject).val.(type) {
		case int:
			if obj <= lastInt {
				return errors.New(fmt.Sprintf("got %v (int) out of order, last was %v", obj, lastInt))
			}
			lastInt = obj
		case uint64:
			if obj <= lastUint {
				return errors.New(fmt.Sprintf("got %v (uint) out of order, last was %v", obj, lastUint))
			} else {
				lastUint = obj
			}
		default:
			return errors.New(fmt.Sprintf("unexpected type %#v", obj))
		}
	}

	return nil
}

func main() {
	if err := fifo(); err != nil {
		fmt.Errorf("fifo does not work perfectly!!")
	}
}
