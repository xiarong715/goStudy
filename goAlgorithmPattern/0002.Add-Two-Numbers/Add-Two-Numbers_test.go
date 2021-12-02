package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoNums(t *testing.T) {
	list1 := createLinkedListWithNum(235)
	list2 := createLinkedListWithNum(111)
	fmt.Print("list1: ")
	printLinkedList(list1)
	fmt.Print("list2: ")
	printLinkedList(list2)
	list := AddTwoNums(list1, list2)
	fmt.Print("list: ")
	printLinkedList(list)
	if !isEqual(list, createLinkedListWithNum(346)) {
		t.Error("return error")
	}
}

func printLinkedList(head *LinkedList) {
	temp := head
	for temp.next != nil {
		fmt.Printf("%d ", temp.next.value)
		temp = temp.next
	}
	fmt.Println()
}

func isEqual(a, b *LinkedList) bool {
	list1 := a.next
	list2 := b.next
	for list1 != nil {
		if list1.value != list2.value {
			return false
		}
		list1 = list1.next
		list2 = list2.next
	}
	return true
}
