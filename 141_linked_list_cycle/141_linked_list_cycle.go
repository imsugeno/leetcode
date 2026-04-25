package main

type ListNode struct {
     Val int
     Next *ListNode
}

func hasCycle(head *ListNode) bool {
    seen := make(map[*ListNode]bool)

	for node := head; node != nil; node = node.Next {
		if seen[node] {
			return true
		}

		seen[node] = true
	}

	return false
}

func main() {}
