package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoNums(t *testing.T) {
	list1 := createLinkedListWithNum(235)
	list2 := createLinkedListWithNum(111)
	fmt.Println("list1:")
	temp := list1
	for temp != nil {
		fmt.Printf("%v ", temp.value)
		temp = temp.next
	}
	fmt.Println()
	fmt.Println("list2:")
	temp = list2
	for temp != nil {
		fmt.Printf("%v ", temp.value)
		temp = temp.next
	}
	fmt.Println()
	// list := AddTwoNums(list1, list2)
	// if isEqual(list, createLinkedListWithNum(346)) {
	// 	t.Error("return error")
	// }
}

func isEqual(a, b *LinkedList) bool {
	list1 := a
	list2 := b
	for {
		if list1 == nil && list2 != nil || list1 != nil && list2 == nil {
			return false
		}
		if list1 == nil || list2 == nil {
			return true
		}
		if list1.value != list2.value {
			return false
		}
		list1 = list1.next
		list2 = list2.next
	}
}
