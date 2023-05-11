package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
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

func AddWithError() {
	var n int32 = 0
	//e := make(chan error)

	startTime := time.Now()

	var g errgroup.Group
	for i := 0; i < 10; i++ {
		d := i
		g.Go(func() error {
			// if d == 6 {
			// 	return errors.New("mock error")
			// }

			fmt.Println("I is", d)

			atomic.AddInt32(&n, 1)
			return nil
		})
	}

	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Println(err.Error())
	}

	// wg.Wait()
	fmt.Println(n)
	endTime := time.Now()
	//wg.Wait()
	fmt.Println(endTime.Sub(startTime))
	fmt.Println("AddWithError end")
}
