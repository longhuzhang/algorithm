package main

import "fmt"

func main() {
	fmt.Println(pancakeSort([]int{2,1,3}))

	//[2,3,2,2,1,1,3]
}


var resp  = make([]int, 0)

func pancakeSort(arr []int) []int {
	 pancake(arr, len(arr))
	 return resp
}

func pancake(arr []int, n int) []int {
	if n == 1 {
		return arr
	}

	var maxIndex, maxVal int
	var i = 0
	for ; i < n; i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
			maxIndex = i
		}
	}

	revert(arr[:maxIndex+ 1])
	resp = append(resp, maxIndex + 1)
	revert(arr[:n])
	resp = append(resp, n)
	return pancake(arr, n-1)
}

func revert(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
