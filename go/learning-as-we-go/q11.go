package main

import (
	"fmt"
)

func main() {
	m := []int{1, 3, 4}
	f := func(i int) int {
		return i * i
	}
	fmt.Println(map1(f, m))

	m2 := []string{"a", "b", "1"}
	f2 := func(i string) string {
		return i + i
	}
	fmt.Println(map2(f2, m2))
}

func map1(f func(int) int, arr []int) []int {
	ret := make([]int, len(arr))

	for k, v := range arr {
		ret[k] = f(v)
	}

	return ret
}

func map2(f func(string) string, arr []string) []string {
	ret := make([]string, len(arr))

	for k, v := range arr {
		ret[k] = f(v)
	}

	return ret
}
