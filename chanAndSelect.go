package main

import (
	"fmt"
	"time"
)

func Chann(ch chan int, stopCh chan bool) {

	for j := 0; j < 10; j++ {
		ch <- j
		time.Sleep(time.Second * 5)
	}
	stopCh <- true
}

func ChanAndSelect() {
	ch := make(chan int)
	i := 0
	c := 0
	stopCh := make(chan bool)

	go Chann(ch, stopCh)

	for {
		fmt.Println("Jump into loop", i)
		i++
		select {
		case c = <-ch:
			fmt.Println("Receive C", c)
			time.Sleep(time.Second * 10)
		case s := <-ch:
			fmt.Println("Receive S", s)
		case _ = <-stopCh:
			goto end
		}
	}
end:
}
