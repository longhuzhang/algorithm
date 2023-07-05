package main

import "fmt"

func main() {
	water := trapDoublePoint([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	fmt.Println(water)
}

func trap(height []int) int {
	var count int
	length := len(height)
	lMax := make([]int, length)
	rMax := make([]int, length)
	lMax[0] = height[0]
	rMax[length - 1] = height[length - 1]

	for i := 1;i<length;i++{
		lMax[i] = Max(height[i], lMax[i - 1])
	}

	for i := length-2; i>=0;i--{
		rMax[i] = Max(height[i], rMax[i+1])
	}

	for i := 0; i < len(height); i++ {
		var l, r = 0, 0
		l = Max(lMax[i], height[i])
		r = Max(rMax[i], height[i])
		count = count + (Min(l, r) - height[i])
	}

	return count
}



func trapDoublePoint(height []int) int {
	var count, lMax, rMax, left, right int
	lMax = height[0]
	rMax = height[len(height) - 1]
	right = len(height) - 1
	
	for left <= right {

		lMax = Max(lMax, height[left])
		rMax = Max(rMax, height[right])

		if lMax < rMax{
			count += lMax - height[left]
			left ++
		}else {
			count += rMax - height[right]
			right --
		}
	}

	return count
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
