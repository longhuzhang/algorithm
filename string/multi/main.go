package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(multiply("3", "8"))
}

func multiply(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)

	res := make([]int, l1+l2)

	var  data int
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			x := num1[i] - '0'
			y := num2[j] - '0'
			data = int(x*y) + res[i+j + 1] 
			res[i + j] += data / 10
			res[i+j + 1] = data % 10 
		}
	}

	var arrByte = bytes.NewBuffer(nil)
	var triped bool
	for i := 0; i <len(res); i++{
		if res[i] == 0&& !triped{
			continue
		}
		triped = true
		temp := byte(res[i] + '0')
		arrByte.Write([]byte{temp})
	}
	if arrByte.Len() == 0{
		return ""
	}
	return arrByte.String()
}