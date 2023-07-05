package main

import "fmt"

func main() {
	index = 0
	result := calculate(" 3+5 / 2 ")
	fmt.Println(result)
}

func calculate(s string) int {

	arrData := make([]int, 0)
	num := 0
	var sign byte = '+'
	for i := 0; i < len(s); i++ {


		if isDigital(s[i]) {
			num = num*10 + int(s[i]-'0')
		}

		if (!isDigital(s[i])&& s[i] != ' ') || i == len(s)-1 {
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

var index int

func recusionCalculate(str string) int {

	arrData := make([]int, 0)
	num := 0
	var sign byte = '+'
	for ; index < len(str); index++ {

		if isDigital(str[index]) {
			num = num*10 + int(str[index]-'0')
		}

		if str[index] == '(' {
			index ++
			num = recusionCalculate(str)
			index ++
		}

		if (!isDigital(str[index])&& str[index] != ' ') || index == len(str)-1 {
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

			sign = str[index]
			num = 0
		}
		if str[index] == ')' {
			break
		}
	}

	var count int
	for i := 0; i < len(arrData); i++ {
		count += arrData[i]
	}
	return count
}

func isDigital(b byte) bool {
	s := b - '0'
	if s < 0 || s > 9 {
		return false
	}
	return true
}
