package main

func main() {

}

func max(s []int) int {
	max := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}

func min(s []int) (min int) {
	min = s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return
}
