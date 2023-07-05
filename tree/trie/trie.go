package main

import "fmt"

type Trie struct {
	//str   [26]*Trie
	hasStr  map[int]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		//str:   [26]*Trie{},
		hasStr: make(map[int]*Trie),
		isEnd: false,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	temp := this
	//var index int
	for _, v := range word {
		//index = int(v) - 'a'
		//if temp.str[index] == nil {
		//	temp.str[index] = new(Trie)
		//}
		if temp.hasStr[int(v)] == nil{
			temp.hasStr[int(v)] = &Trie{
				hasStr: make(map[int]*Trie),
			}
		}
		temp = temp.hasStr[int(v)]
	}
	temp.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this
	//var index int
	for _, ch := range word {
		//index = int(ch) - 'a'
		//fmt.Println(index, string(ch))
		//if root.str[index] == nil {
		//	return false
		//}
		//root = root.str[index]
		if root.hasStr[int(ch)] == nil{
			return false
		}
		root = root.hasStr[int(ch)]
	}
	return root.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this
	//var index int
	for _, ch := range prefix {
		//index = int(ch) - 'a'
		//if root.str[index] == nil {
		//	return false
		//}
		//root = root.str[index]
		if root.hasStr[int(ch)] == nil{
			return false
		}
		root = root.hasStr[int(ch)]
	}
	return true
}

func (this *Trie) Print() {
	temp := this
	//for i := 0; i < 26; i++ {
		//fmt.Println(temp.str[i], temp.str[i] != nil)
		//if temp.str[i] != nil {
		//	fmt.Print(string(i + 'a'))
		//	temp = temp.str[i]
		//	i = 0
		//	continue
		//}

	//}
	for k,v :=range temp.hasStr{
			fmt.Println(string(k))
			temp = v
	}
	fmt.Println()
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := Constructor()
	trie.Insert("apple")
	trie.Print()
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))

}
