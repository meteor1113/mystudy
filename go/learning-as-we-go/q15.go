package stack

import "strconv"

type stack struct {
	i    int
	data [10]int
}

func (s *stack) Push(k int) {
	if s.i+1 > 9 {
		return
	}

	s.data[s.i] = k
	s.i++
}

func (s *stack) Pop() int {
	s.i--
	return s.data[s.i]
}

func (s stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
