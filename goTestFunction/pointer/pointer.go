package main

import "fmt"

type LinkedList struct {
	value int
	next  *LinkedList
}

func insertNode(head *LinkedList, node *LinkedList) {
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = node
}

// 有头结点
func createLinkedList() *LinkedList {
	var head LinkedList = LinkedList{value: -1}
	for i := 0; i < 6; i++ {
		insertNode(&head, &LinkedList{value: i})
	}
	return &head
}

func main() {
	list := createLinkedList()
	for list != nil {
		fmt.Printf("%d ", list.value)
		list = list.next
	}
	fmt.Println()
}
