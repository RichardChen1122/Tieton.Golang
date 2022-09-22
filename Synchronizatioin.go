package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func Add() {
	var n int32 = 0

	startTime := time.Now()

	var mutex *sync.Mutex = new(sync.Mutex)
	var wg sync.WaitGroup
	wg.Add(10000000)
	for i := 0; i < 10000000; i++ {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			defer wg.Done()
			n += 1
		}()
	}
	wg.Wait()
	fmt.Println(n)
	endTime := time.Now()
	//wg.Wait()
	fmt.Println(endTime.Sub(startTime))
	fmt.Println("end")
}

func Add2() {
	var n int32 = 0

	startTime := time.Now()

	var wg sync.WaitGroup
	wg.Add(10000000)
	for i := 0; i < 10000000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&n, 1)
		}()
	}
	wg.Wait()
	fmt.Println(n)
	endTime := time.Now()
	//wg.Wait()
	fmt.Println(endTime.Sub(startTime))
	fmt.Println("end")
}
