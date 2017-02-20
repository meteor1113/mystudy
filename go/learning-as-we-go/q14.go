package main

import "fmt"

func main() {
	p := plusTwo()
	fmt.Println(p(2))

	p1 := plusX(3)
	fmt.Println(p1(2))
}

func plusTwo() func(int) int {
	return func(i int) int {
		return i + 2
	}
}

func plusX(x int) func(int) int {
	return func(i int) int {
		return i + x
	}
}
