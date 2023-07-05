package main

import "fmt"

type UE struct {
	count  int
	Parent []int
	Size   []int
}

func NewUE(n int) UE {
	source := new(UE)
	source.count = n
	source.Parent = make([]int, n)
	source.Size = make([]int, n)

	for i := 0; i < n; i++ {
		source.Size[i] = 1
		source.Parent[i] = i
	}
	return *source
}

func (u UE) isConnect(p, q int) bool {
	rootp := u.find(p)
	rootq := u.find(q)
	return rootp == rootq
}

func (u UE) find(point int) int {
	temp := point
	for temp != u.Parent[temp] {
		temp = u.Parent[u.Parent[temp]]
	}
	return temp
}

func (u *UE) union(p, q int) {
	rootp := u.find(p)
	rootq := u.find(q)

	if rootp == rootq {
		return
	}

	if u.Size[rootp] > u.Size[rootq] {
		u.Parent[rootq] = u.Parent[rootp]
		u.Size[rootp] += u.Size[rootq]
	} else {
		u.Parent[rootp] = u.Parent[rootq]
		u.Size[rootq] = u.Size[rootp]
	}
	u.count--
}

func equationsPossible(equations []string) bool {

	ue := NewUE(26)
	var temp string
	for i := 0; i < len(equations); i++ {
		temp = equations[i]
		if temp[1] == '=' {
			ue.union(int(temp[0]-'a'), int(temp[3]-'a'))
		}
	}

	for i := 0; i < len(equations); i++ {
		temp = equations[i]
		if temp[1] == '!' {
			if ue.isConnect(int(temp[0]-'a'), int(temp[3]-'a')) {
				return false
			}
		}
	}

	return true
}

func main() {
	arr := []string{"c==c", "b==d", "x!=z"}

	fmt.Println(equationsPossible(arr))
}