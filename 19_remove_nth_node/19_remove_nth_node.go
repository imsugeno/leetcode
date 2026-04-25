package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	totalNodes := 0
	for node := head; node != nil; node = node.Next {
		totalNodes++
	}

	count := 0
	for node := head; node != nil; node = node.Next {
		if totalNodes-n-1 < 0 {
			return head.Next
		}
		if count == totalNodes-n-1 {
			node.Next = node.Next.Next
		}
		count++
	}

	return head
}

func main() {}
