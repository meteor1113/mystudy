package main

func main()  {
	println(ave([]float64 { 3.5, 5.6, 7 }))
}

func ave(arr []float64) (ave float64) {
	sum := 0.0
	for _, v := range arr {
		sum += v
	}

	if (len(arr) == 0) {
		ave = 0.0
	} else {
		ave = sum / float64(len(arr))
	}

	return
}
