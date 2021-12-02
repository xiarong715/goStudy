package leetcode

type LinkedList struct {
	value int
	next  *LinkedList
}

func createLinkedListWithNum(num int) *LinkedList {
	var head LinkedList = LinkedList{value: -1}
	for num != 0 {
		digit := num % 10
		num = num / 10
		insertNode(&head, &LinkedList{value: digit})
	}
	return &head
}

func insertNode(head *LinkedList, node *LinkedList) {
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = node
}

func AddTwoNums(a, b *LinkedList) *LinkedList {
	list1 := a.next
	list2 := b.next
	carry := false
	digit := 0
	var head LinkedList = LinkedList{value: -1}
	for list1 != nil {
		digit = list1.value + list2.value
		if digit/10 != 0 {
			carry = true
			digit = digit % 10
		} else {
			carry = false
		}
		if carry {
			insertNode(&head, &LinkedList{value: digit + 1})
		} else {
			insertNode(&head, &LinkedList{value: digit})
		}
		list1 = list1.next
	}
	return &head
}
