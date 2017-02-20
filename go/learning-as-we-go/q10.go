package main

import (
	"fmt"
)

func main() {
	fmt.Println(fibonacci(4))
	fmt.Println(fibonacci(6))
}

func fibonacci(num int) []int {
	arr := make([]int, num)
	arr[0] = 1
	arr[1] = 1

	for i := 2; i < num; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr
}
