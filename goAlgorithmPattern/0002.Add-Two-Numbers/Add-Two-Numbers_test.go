package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	params
	ans
}

type params struct {
	num1 int
	num2 int
}

type ans struct {
	num int
}

var qs []question = []question{
	{
		params{123, 234},
		ans{357},
	},
	{
		params{235, 111},
		ans{346},
	},
	{
		params{888, 111},
		ans{999},
	},
}

func TestAddTwoNums(t *testing.T) {
	for _, v := range qs {
		list1 := createLinkedListWithNum(v.params.num1)
		list2 := createLinkedListWithNum(v.params.num2)
		fmt.Print("list1: ")
		printLinkedList(list1)
		fmt.Print("list2: ")
		printLinkedList(list2)
		list := AddTwoNums(list1, list2)
		fmt.Print("list: ")
		printLinkedList(list)
		if !isEqual(list, createLinkedListWithNum(v.ans.num)) {
			t.Error("return error")
		}
		fmt.Println("------------------------")
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
