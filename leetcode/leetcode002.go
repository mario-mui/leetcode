package leetcode

type LinkNode struct {
	Value int
	Next  *LinkNode
}

func AddTwoNumbers(l1 *LinkNode, l2 *LinkNode) LinkNode {
	head := &LinkNode{Value: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Value
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Value
			l2 = l2.Next
		}
		current.Next = &LinkNode{Value: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}

	return *head.Next
}
