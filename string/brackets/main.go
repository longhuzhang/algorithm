package main

import "fmt"

func main() {
	valid := isValid("{[]}")
	fmt.Println(valid)
}
func isValid(s string) bool {
	str := make([]string, 0)
	// fmt.Println(str,"length is ", len(str))
	for _, ch := range s{
		isInclude := false
		index := 0
		for in,st := range str{
			fmt.Println("st",st,"ch",string(ch), string(ch - 2), string(ch -1))
			if st == string(ch-2) || st == string(ch-1){
				isInclude =true
				index = in
			}
		}
		if isInclude ==true {
			// fmt.Println(len(str), index+1)
			if len(str)!= index+1 {
				return false
			}
			str = str[:index]
		}else {
			str = append(str,string(ch))
		}
		// fmt.Println("temp str",str)
	}
	// fmt.Println(str,"length is ", len(str))
	return len(str) == 0
}