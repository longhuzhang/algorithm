package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//param := []int{1, 1, 0, 6, 3}
	//param := []int{5, 2, 6, 3, 9, 1, 7, 3, 8, 4}
	param := []int{16, 3, 8, 3, 16, 5, 1, 11, 11, 6, 13, 5, 15, 5, 12, 0, 1, 18, 5, 0, 16, 13, 8}
	head := constructList(param)
	head = reverseEvenLengthGroups(head)
	PrintlnList(head)
	//[16,8,3,3,16,5,6,11,11,1,13,5,15,5,12,16,0,5,18,1,0,8,13]
}

var index = 0

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	index++
	if index%2 == 0 {
		start := head
		end := head
		var i int
		for i = 0; i < index; i++ {
			if end == nil {
				break
			}
			end = end.Next
		}
		if i%2 != 0 {
			return head
		}
		newHead := revertList(start, end)
		start.Next = reverseEvenLengthGroups(end)
		return newHead
	}
	cur := head
	length := 0
	for i := 0; i < index-1; i++ {
		if cur == nil {
			break
		}
		length++
		cur = cur.Next
	}

	//最后一组是偶数也要翻转
	if length != 0 && cur == nil && (length)%2 == 0 {
		return revertList(head, cur)
	}

	if cur == nil {
		return head
	}

	cur.Next = reverseEvenLengthGroups(cur.Next)
	index = 0
	return head
}

func revertList(start, end *ListNode) *ListNode {
	var pre *ListNode
	cur := start
	next := start
	for cur != end {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
