package linked_list

import "testing"

func init(){
	node3.Val = 3
	node3.Next = nil
	node2.Val = 2
	node2.Next = &node3
	node1.Val = 1
	node1.Next = &node2
}

func TestReverseList(t *testing.T){
	//_ = ListNode{1, &node2}
	lst := node1.GetSlice()
	t.Log(lst)
	reverseHead := reverseList(&node1)
	//reverseLst := reverseHead.GetSlice()
	t.Log(reverseHead.GetSlice())
}
