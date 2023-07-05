package main

import "fmt"

func main() {

	str := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"

	num := longestPalindrome(str)
	fmt.Println(num)
}

// @lc code=start
func longestPalindrome(s string) int {

	length := len(s)
	hashChar := make(map[byte]int, length)
	for i := 0;i<length;i++{
		hashChar[s[i]] ++
	}

	var hasOdd bool
	var maxLen int
	for _, v := range hashChar{
		if v%2 == 0{
			maxLen += v
		}else {
			if !hasOdd {
				hasOdd = true
			}
			maxLen += (v - 1)
		}
	}
	if hasOdd{
		maxLen +=1
	}

	return maxLen
}

// @lc code=end


