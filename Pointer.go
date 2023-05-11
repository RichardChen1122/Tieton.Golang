package main

import "fmt"

type DemoType struct {
	Name string
	Age  string
}

var demo = &DemoType{Name: "Richard", Age: "18"}

func PrintPtr() {
	println(demo)
	fmt.Println(demo)
	println(&demo)
}
