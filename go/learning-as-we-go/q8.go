package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s stack
	s.push(25)
	fmt.Printf("stack %v\n", s)
	s.push(10)
	fmt.Printf("stack %v\n", s)
}

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	if s.i+1 > 9 {
		return
	}

	s.data[s.i] = k
	s.i++
}

func (s stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
