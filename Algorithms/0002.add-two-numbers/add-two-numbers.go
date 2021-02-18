package problem0002

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// ListNode defines for singly-linked list.
//  type ListNode struct {
//      Val int
//      Next *ListNode
//  }
type ListNode = kit.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resPre := &ListNode{}
	cur := resPre
	carry := 0
        
	// Iterate through lists until all of the list is nil
	for l1 != nil || l2 != nil || carry > 0 {
		// Initalize sum with previous carry
		sum := carry

		// Add l1 val to sum and move l1
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		// Add l2 val to sum and move l2 to next
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// get carry if sum is greater than 10
		carry = sum / 10

		// store mod 10 val in curr node next
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return resPre.Next
}
