package main

import "fmt"

func AppendAhead() {
	var x []int
	for i := 2; i < 10; i += 2 {
		x = append(x, i)

	}
	reverseArray(x)

	fmt.Println(x)
}

func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
