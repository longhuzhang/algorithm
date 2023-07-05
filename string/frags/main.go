package main

import (
	"fmt"
	"math"
)

func minNumberOfFrogs(croakOfFrogs string) int {

	if len(croakOfFrogs)%5 != 0 {
		return -1
	}

	var c, r, o, a, k = 0, 0, 0, 0, 0
	count := 0
	for i := 0; i < len(croakOfFrogs); i++ {
		switch croakOfFrogs[i] {
		case 'c':
			c++
			count = int(math.Max(float64(count), float64(c - k)))
		case 'r':
			r ++
		case 'o':
			o++
		case'a':
			a++
		case'k':
			k++
		}
		if c<r || r<o|| o<a||a<k{
			return -1
		}
	}

	if c == r&& r==o&&o==a&& a==k{
		return count
	} 
	return -1

}

func main() {
	num := minNumberOfFrogs("croakcroak")
	fmt.Println(num)
}