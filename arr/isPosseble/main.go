package main

import "fmt"

func main() {
	fmt.Println(isPossible([]int{1,2,3,6,8,9,10}))
}

func isPossible(nums []int) bool {
	freq := make(map[int]int, len(nums))
	need := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}

	for _, val := range nums {

		if freq[val] == 0 {
			continue
		}
		if need[val] > 0 {
			freq[val]--
			need[val]--
			need[val+1]++
		} else if freq[val] > 0 && freq[val+1] > 0 && freq[val+2] > 0 {
			freq[val]--
			freq[val+1]--
			freq[val+2]--
			need[val+3]++
		} else {
			return false
		}

	}
	return true

}
