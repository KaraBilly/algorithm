package linked_list

type Node struct {
	Val  interface{}
	Next *Node
}

/*
  Append Node behind giving Node
*/
func (head *ListNode) Append(i int) {
	if head == nil {
		head.Val = i
	}
	var cur ListNode
	cur.Val = i
	cur.Next = head.Next
	head.Next = &cur
}

func (head *ListNode) GetLength() int {
	if head == nil {
		return 0
	}

	iterator := head
	var length int = 1
	for iterator.Next != nil {
		length++
		iterator = iterator.Next
	}
	return length

}

func (head *ListNode) GetSlice() []interface{} {
	if head == nil {
		return nil
	}
	slice := make([]interface{}, 20)
	iterator := head
	for iterator.Next != nil {
		slice = append(slice, iterator.Val)
		iterator = iterator.Next
	}
	slice = append(slice, iterator.Val)
	return slice
}
