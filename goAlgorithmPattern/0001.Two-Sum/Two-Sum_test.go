package leetcode

import (
	"fmt"
	"testing"
)

type question1 struct {
	param1
	ans1
}

type param1 struct {
	nums   []int
	target int
}

type ans1 struct {
	one []int
}

type question2 struct {
	param2
	ans2
}

type param2 struct {
	nums   []int
	target int
}

type ans2 struct {
	one [][]int
}

var qs1 []question1 = []question1{
	{
		param1{[]int{1, 2, 3, 4, 5}, 8},
		ans1{[]int{2, 4}},
	},
	{
		param1{[]int{2, 3, 7, 10}, 12},
		ans1{[]int{0, 3}},
	},
	{
		param1{[]int{0, 1, 2, 5}, 5},
		ans1{[]int{0, 3}},
	},
	{
		param1{[]int{2, 3, 5}, 10},
		ans1{[]int{}},
	},
	{
		param1{[]int{0, 1}, 1},
		ans1{[]int{0, 1}},
	},
}

var qs2 []question2 = []question2{
	{
		param2{[]int{1, 2, 3, 5, 7}, 8},
		ans2{[][]int{{0, 4}, {2, 3}}},
	},
	{
		param2{[]int{0, 1, 2, 3}, 3},
		ans2{[][]int{{0, 3}, {1, 2}}},
	},
	{
		param2{[]int{5, 6, 7, 8, 9, 10}, 15},
		ans2{[][]int{{0, 5}, {1, 4}, {2, 3}}},
	},
}

func TestTwoSum1(t *testing.T) {
	fmt.Println("nums, target")
	for _, q := range qs1 {
		res := twoSum1(q.nums, q.target)
		if !isEqual1(res, q.ans1.one) {
			t.Error("return error")
		}
		fmt.Printf("%v, %v\n", q.nums, q.target)
		fmt.Printf("%v\n", res)
		fmt.Println("-----------------------")
	}
}

func TestTwoSum2(t *testing.T) {
	fmt.Println("nums, target")
	for _, q := range qs1 {
		res := twoSum2(q.nums, q.target)
		if !isEqual1(res, q.ans1.one) {
			t.Error("return error")
		}
		fmt.Printf("%v, %v\n", q.nums, q.target)
		fmt.Printf("%v\n", res)
		fmt.Println("-----------------------")
	}
}

func TestTwoSum3(t *testing.T) {
	fmt.Println("nums, target")
	for _, q := range qs2 {
		res := twoSum3(q.nums, q.target)
		if !isEqual2(res, q.ans2.one) {
			t.Error("return error")
		}
		fmt.Printf("%v, %v\n", q.nums, q.target)
		fmt.Printf("%v\n", res)
		fmt.Println("-----------------------")
	}
}

func isEqual1(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, _ := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

func isEqual2(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	count := 0
	for k1, _ := range a {
		for k2, _ := range b {
			if isEqual1(a[k1], b[k2]) {
				count++
				break
			}
		}
	}
	return count == len(a)
}
