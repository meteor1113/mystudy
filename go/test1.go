package main

import "fmt"
// import "unicode/utf8"
// import "os"

func main() {
    // foo()
	foo1()
    fmt.Printf("Hello, world ; or ´ϵ óϵ ; orこんにちは世界")
}

func foo()  {
	// var aa int
	const (
		aaa = iota
		bbb
	)

	a, b := 20, "adsa"
	fmt.Printf("%d\n", a)
	// fmt.Printf(b)
	fmt.Println(b)
	c64 := 5 + 5i
	fmt.Println(c64)
	if true {
		fmt.Println(4 + 6i)
	}
}

func foo1()  {
	arr := []int { 3, 5, 7, 8 }
	for k, v := range arr {
		fmt.Println(k, v)
	}
}

// func foo(i int) string {
//     // os.Chdir("")
//     return ""
// }

func aaa()  {
    // utf8.
    // fmt.Printf("")

}
