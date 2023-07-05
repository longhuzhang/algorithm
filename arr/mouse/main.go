package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(miceAndCheese([]int{2, 1}, []int{1, 2}, 1))
}

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	l := len(reward1)
	var subArr = make([]int, l)
	var count int
	for i := 0; i < l; i++ {
		count += reward2[i]
		subArr[i] = reward1[i] - reward2[i]
	}

	sort.Slice(subArr, func(x int, y int) bool {
		return subArr[x] > subArr[y]
	})

	for i := 0;i<k;i++{
		count += subArr[i]
	}
	return count
}
