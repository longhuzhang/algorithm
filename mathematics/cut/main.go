package main

import "fmt"

func main() {
	fmt.Println(cuttingRope(10))
}
func cuttingRope(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	product := make([]int, n+1)
	product[0] = 0
	product[1] = 1
	product[2] = 2
	product[3] = 3

	var max int
	for i := 4; i <= n; i++ {
		max = 0
		for j := 1; j <= i/2; j++ {
			temp := (product[j] * product[i-j]) % 1000000007
			if max < temp {
				max = temp
			}
		}
		product[i] = max
	}
	return product[n]
}
