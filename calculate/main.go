package main

import "fmt"

func main() {

	result := calculate("8/4 + 9 * 2")
	fmt.Println(result)
}

func calculate(s string) int {

	arrData := make([]int, 0)
	num := 0
	var sign byte = '+'
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}

		if isDigital(s[i]) {
			num = num*10 + int(s[i]-'0')
		}



		if !isDigital(s[i]) || i == len(s)-1 {
			switch sign {
			case '+':
				arrData = append(arrData, num)
			case '-':
				arrData = append(arrData, -num)
			case '*':
				lastIndex := len(arrData) - 1
				arrData[lastIndex] = arrData[lastIndex] * num
			case '/':
				lastIndex := len(arrData) - 1
				arrData[lastIndex] = arrData[lastIndex] / num
			}

			sign = s[i]
			num = 0
		}
	}

	var count int
	for i := 0; i < len(arrData); i++ {
		count += arrData[i]
	}
	return count
}


func recusionCalculate(str string) int {
	
}

func isDigital(b byte) bool {
	s := b - '0'
	if s < 0 || s > 9 {
		return false
	}
	return true
}
