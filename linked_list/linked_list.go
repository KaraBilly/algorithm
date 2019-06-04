package linked_list


/**
 单链表反转
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
思路 ：先暂存上一个node ，后一个node指向前一个node
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
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

