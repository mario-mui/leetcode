package main

import (
	"fmt"
	leetcode "lcmod/leetcode"
)

func main() {
	// leetcode 001
	fmt.Println("leetcode001: [2,5,9,12,6] 8 ===>", leetcode.TwoSum([]int{2, 5, 9, 12, 6}, 8))

	// leetcode 002
	l1 := leetcode.LinkNode{Value: 2}
	l1Node1 := leetcode.LinkNode{Value: 4}
	l1Node2 := leetcode.LinkNode{Value: 6}
	l1Node1.Next = &l1Node2
	l1.Next = &l1Node1

	l2 := leetcode.LinkNode{Value: 1}
	l2Node1 := leetcode.LinkNode{Value: 3}
	l2Node2 := leetcode.LinkNode{Value: 5}
	l2Node1.Next = &l2Node2
	l2.Next = &l2Node1

	l3 := leetcode.AddTwoNumbers(&l1, &l2)
	fmt.Println("leetcode002: 2->4->6 1->3->5 ===>", l3.Value, "->", l3.Next.Value, "->", l3.Next.Next.Value)

	// leetcode 003
	fmt.Println("leetcode003: ascfards ===>", leetcode.LengthOfLongestSubstring("ascfards"))

	// leetcode 004
	fmt.Println("leetcode003: {2, 5, 9, 12} {1, 3, 4, 8, 11}  ===>", leetcode.FindMedianSortedArrays([]int{2, 3, 9, 12}, []int{1, 3, 4, 8, 11}))
}
