package main

import "fmt"

func constructList(arr []int) *ListNode {
	dummyNode := new(ListNode)
	ahead := dummyNode
	for i := range arr {
		dummyNode.Next = &ListNode{
			Val: arr[i],
		}
		dummyNode = dummyNode.Next
	}
	return ahead.Next
}

func PrintlnList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" ")
		}
		head = head.Next
	}
	fmt.Println()
}
