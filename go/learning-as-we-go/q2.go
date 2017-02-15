package main


func main()  {
	for index := 1; index <= 100; index++ {
		switch {
		case index % 3 == 0 && index % 5 == 0:
			println("FizzBuzz")
		case index % 3 == 0:
			println("Fizz")
		case index % 5 == 0:
			println("Buzz")
		default:
			println(index)
		}
	}
}
