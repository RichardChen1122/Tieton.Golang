package main

import "fmt"

func printGenericAny[T any](foo, bar T) {
	fmt.Println(foo)
	fmt.Println(bar)
}

// any is an alias for interface{}. Spec: Interface types:
// For convenience, the predeclared type any is an alias for the empty interface.

func printGeneric[T interface{}](foo, bar T) {
	fmt.Println(foo)
	fmt.Println(bar)
}

func printInterface(foo, bar any) {
	fmt.Println(foo)
	fmt.Println(bar)
}

func RunGeneric() {
	printInterface(40, 0.1)
	//printGeneric(100, 0.1)   //compile error,  T inferred to int, "0.1" not assignable to int
	printGeneric[interface{}](100, 0.1)
	printGeneric[any](100, 0.1)

	var x interface{} = 42
	var y any = 42

	printGenericAny(x, y) // Error: x is not of type int
	printGenericAny(y, x) // OK: y is of type int
	printInterface(x, y)  // OK: x is assignable to interface{}
	printInterface(y, x)  // OK: y is assignable to interface{}
}
