package main

import (
	"fmt"
)

func main()  {
	// f1()
	// f2()
	f3()
}

func f1()  {
	for index := 0; index < 10; index++ {
		fmt.Println(index)
	}
}

func f2()  {
	i := 0
w:
	i++
	if i < 10 {
		fmt.Println(i)
		goto w
	}
}

func f3()  {
	arr := []int { 0, 1, 2, 3, 4, 5, 6, 7 }
	fmt.Printf("%v", arr)
	// for _, v := range arr {
	// 	fmt.Println(v)
	// }
}
