package main

import (
	"fmt"
)

func main() {
	pfoo(1, 4, 5)
	pfoo(5, 6, 7, 8, 9, 3, 0)
}

func pfoo(nums ...int) {
	fmt.Printf("len: %d\n", len(nums))
	for _, i := range nums {
		fmt.Println(i)
	}
}
