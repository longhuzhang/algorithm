package list_conflict

//Definition for singly-linked list.type

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	fast := head.Next.Next
	last := head.Next

	for fast != last && fast.Next != nil && fast != nil {
		fast = fast.Next.Next
		last = last.Next
	}

	if fast != last || fast == nil || fast.Next == nil {
		return nil
	}

	fast = head
	for fast != last {
		fast = fast.Next
		last = last.Next
	}

	return fast
}
