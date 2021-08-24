package main

import (
	"fmt"
)

func main() {
	c1 := &ListNode{3, nil}
	b1 := &ListNode{4, c1}
	a1 := &ListNode{2, b1}
	list1 := a1

	c2 := &ListNode{4, nil}
	b2 := &ListNode{6, c2}
	a2 := &ListNode{5, b2}
	list2 := a2

	fmt.Println(addTwoNumbers(list1, list2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return "nil"
	}
	if l.Next == nil {
		return fmt.Sprintf("%d", l.Val)
	}
	return fmt.Sprintf("%d -> %s", l.Val, l.Next)
}

// Elementary Math Solution
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	start := new(ListNode)
	curr := start
	carry := 0
	for l1 != nil || l2 != nil {
		var num1, num2 int

		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}

		sum := num1 + num2 + carry
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next

		if carry > 0 {
			curr.Next = &ListNode{Val: carry}
		}
	}
	return start.Next
}