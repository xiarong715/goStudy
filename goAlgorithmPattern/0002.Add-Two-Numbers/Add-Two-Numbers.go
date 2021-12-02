package leetcode

import "fmt"

type LinkedList struct {
	value int
	next  *LinkedList
}

func createLinkedListWithNum(num int) *LinkedList {
	var head LinkedList
	fmt.Println(head)
	var temp *LinkedList = nil
	for num != 0 {
		digit := num % 10
		num = num / 10
		fmt.Printf("%d ", digit)

	}
	fmt.Println()
	return &head
}

func createNode(digit int) *LinkedList {
	return &LinkedList{value: digit}
}

func AddTwoNums(a, b *LinkedList) *LinkedList {
	list1 := a
	list2 := b
	carry := false
	digit := 0
	var head *LinkedList = nil
	temp := head
	for list1 != nil {
		digit = list1.value + list2.value
		if digit/10 != 0 {
			carry = true
			digit = digit % 10
		} else {
			carry = false
		}
		if temp == nil {
			temp = new(LinkedList)
			if carry {
				temp.value = digit + 1
				carry = false
			} else {
				temp.value = digit
			}
			if head == nil {
				head = temp
			}
			temp = temp.next
		}

		list1 = list1.next
		list2 = list2.next
	}
	return head
}
