package main

import "fmt"

func main() {
	fmt.Println(f(7, 2))
	fmt.Println(f(2, 7))
}

func f(x int, y int) (a, b int) {
	if x > y {
		return y, x
	}

	return x, y
}
