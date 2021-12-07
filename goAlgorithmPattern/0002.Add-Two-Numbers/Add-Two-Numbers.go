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

// func createLinkedListWithStr(num string) *LinkedList {
// 	var head LinkedList = LinkedList{value: -1}
// 	length := len(num)
// 	for i := 0; i < length; i++ {
// 		digit, _ := strconv.Atoi(string(num[i]))
// 		insertNode(&head, &LinkedList{value: digit})
// 	}
// 	return &head
// }

func insertNode(head *LinkedList, node *LinkedList) {
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = node
}

// func AddTwoNums(a, b *LinkedList) *LinkedList {
// 	list1 := a.next
// 	list2 := b.next
// 	digit := 0
// 	carry := false
// 	var head LinkedList = LinkedList{value: -1}
// 	for list1 != nil && list2 != nil {
// 		digit = list1.value + list2.value
// 		if carry {
// 			digit += 1
// 			carry = false
// 		}
// 		if digit/10 != 0 {
// 			digit = digit % 10
// 			carry = true
// 		} else {
// 			carry = false
// 		}
// 		insertNode(&head, &LinkedList{value: digit})
// 		list1 = list1.next
// 		list2 = list2.next
// 	}

// 	for list1 != nil {
// 		digit = list1.value
// 		if carry {
// 			digit += 1
// 		}
// 		if digit/10 != 0 {
// 			digit = digit % 10
// 			carry = true
// 		} else {
// 			carry = false
// 		}
// 		insertNode(&head, &LinkedList{value: digit})
// 		list1 = list1.next
// 	}
// 	for list2 != nil {
// 		digit = list2.value
// 		if carry {
// 			digit += 1
// 		}
// 		if digit/10 != 0 {
// 			digit = digit % 10
// 			carry = true
// 		} else {
// 			carry = false
// 		}
// 		insertNode(&head, &LinkedList{value: digit})
// 		list2 = list2.next
// 	}
// 	if carry {
// 		digit = 1
// 		insertNode(&head, &LinkedList{value: digit})
// 	}

// 	return &head
// }
func AddTwoNums(a, b *LinkedList) *LinkedList {
	list1 := a.next
	list2 := b.next
	digit := 0
	carry := false
	var head LinkedList = LinkedList{value: -1}
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			digit = list1.value + list2.value
		} else if list2 == nil {
			digit = list1.value
		} else {
			digit = list2.value
		}
		if carry {
			digit += 1
			carry = false
		}
		if digit/10 != 0 {
			digit = digit % 10
			carry = true
		} else {
			carry = false
		}
		insertNode(&head, &LinkedList{value: digit})
		if list1 != nil {
			list1 = list1.next
		}
		if list2 != nil {
			list2 = list2.next
		}
	}

	return &head
}
