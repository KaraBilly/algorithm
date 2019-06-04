package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var node3 ListNode
var node2 ListNode
var node1 ListNode

func init(){
	node3.Val = 3
	node3.Next = nil
	node2.Val = 2
	node2.Next = &node3
	node1.Val = 1
	node1.Next = &node2
}

func TestGetLength(t *testing.T){
	var lstNode *ListNode = nil
	//nil
	length :=lstNode.GetLength()
	assert.Equal(t,0,length)

	node1.Next = nil
	length = node1.GetLength()
	assert.Equal(t,1,length)
	node1.Next = &node2

	node2.Next = nil
	length = node1.GetLength()
	node2.Next = &node3

	length = node1.GetLength()
	assert.Equal(t,3,length)
}

func TestGetSlice(t *testing.T){
	lst := make([]interface{},20)
	node1.Next = nil
	lst =append(lst, node1.Val)
	lst1:=node1.GetSlice()
	assert.Equal(t,lst,lst1)

	node1.Next = &node2
	node2.Next = nil
	lst =append(lst, node2.Val)
	lst2:=node1.GetSlice()
	assert.Equal(t,lst,lst2)

	node2.Next = &node3
	lst =append(lst, node3.Val)
	lst3:=node1.GetSlice()
	assert.Equal(t,lst,lst3)
}

func TestAppend(t *testing.T){
	lst := make([]interface{},20)
	lst = append(lst, 1,4,2,3)
	node1.Append(4)
	lstTest:=node1.GetSlice()
	assert.Equal(t,lst,lstTest)

}
