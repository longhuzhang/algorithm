package main

import (
	"bytes"
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestPalindrome("zhanglhonghuuh"))
}

func longestPalindrome(str string) string {
	tempArr := ProcessStr(str)
	countArr := make([]int, len(tempArr))
	var id, max int
	var maxIdx, maxLength int = 0, -1
	for i := 1; i < len(tempArr)-1; i++ {
		if i > max {
			countArr[i] = 1
		} else {
			countArr[i] = int(math.Min(float64(countArr[id*2-i]), float64(max-i)))
		}
		var left, right int
		for {
			left =   i - countArr[i]
			right = countArr[i] + i
			if tempArr[left] != tempArr[right] {
				break
			}
			countArr[i]++
		}

		//超出最大max要更新
		if countArr[i]+i > max {
			id = i
			max = countArr[i] + i
		}

		//更新最大值
		if maxLength < countArr[i]-1 {
			maxLength = countArr[i] - 1
			maxIdx = i
		}
	}

	start := (maxIdx - maxLength) / 2
	return str[start : start+maxLength]
}

func ProcessStr(str string) string {
	byteBuffer := bytes.NewBufferString("")
	byteBuffer.WriteString("^")
	for i := 0; i < len(str); i++ {
		byteBuffer.WriteString(fmt.Sprintf("#%s", string(str[i])))
	}
	byteBuffer.WriteString("#$")
	return byteBuffer.String()
}
