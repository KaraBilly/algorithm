package linked_list

/**
Leetcode #206
 单链表反转
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
思路 ：先暂存上一个node ，后一个node指向前一个node
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	var prev *ListNode = nil
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev
}

/**
Leetcode #141
链表中 环的检测
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
快慢指针
快指针 为nil了还未相遇则不为环形链表
*/
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil && slow != fast {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow == fast
}

/**
Leetcode #21
递归
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}
}

/**
Leetcode #19
快指针比慢指针快n
首node特殊处理
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	slow, fast := head, head
	slow.Next = head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	//If delete node is head node
	if fast == nil {
		head = head.Next
		return head
	}

	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}

	slow.Next = slow.Next.Next
	return head
}

/**
Leetcode #876
快慢指针， 快 two step 慢 one step
*/
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}
