package main

import (
	"fmt"
)

func main() {
	s := []int{5, 7, 8, 3, 1}
	sort(s)
	fmt.Println(s)
}

func sort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}
