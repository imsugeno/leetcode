package leetcode

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var binary string
	bs := getBinary(head, binary)
	result, _ := strconv.ParseInt(bs, 2, 64)

	return int(result)
}

func getBinary(head *ListNode, binary string) string {
	binary += strconv.Itoa(head.Val)
	if head.Next == nil {
		return binary
	}

	return getBinary(head.Next, binary)
}
