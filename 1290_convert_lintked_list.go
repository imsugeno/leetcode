package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	result := 0
	for node := head; node != nil; node = node.Next {
		result = result<<1 | node.Val
	}
	return result
}
