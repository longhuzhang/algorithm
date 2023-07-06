package main

import "fmt"

func main() {

	// var s = "(()))"
	// var s = "())"
	// var s = "))())("
	var s = "(((((("
	s = "(()))(()))()())))"

	fmt.Println(minAddToMakeValid(s))
}

func minAddToMakeValid(s string) int {
	var need, res int
	for _, val := range s {
		if val == '(' {
			need += 2
			if need %2 != 0{
				res ++
				need --
			}
		} else {
			need--
			if need == -1 {
				need = 1
				res++
			}
		}
	}
	return need + res

}
