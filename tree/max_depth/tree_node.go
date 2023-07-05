package main

import(
	"golang.org/x/time/rate"
)

type TreeNode struct {
    Val int
     Left *TreeNode
     Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}

	right  := maxDepth(root.Right)
	left := maxDepth(root.Left)
	if right > left{
		right++
		return right
	}
	left++
	return left
}

func main()  {
	rate.NewLimiter(44,3)
}
