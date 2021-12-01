package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	param
	ans
}

type param struct {
	nums   []int
	target int
}

type ans struct {
	one []int
}

var qs []question = []question{
	{
		param{[]int{1, 2, 3, 4, 5}, 8},
		ans{[]int{2, 4}},
	},
	{
		param{[]int{2, 5, 7, 10}, 12},
		ans{[]int{0, 3}},
	},
	{
		param{[]int{0, 1, 2, 5}, 5},
		ans{[]int{0, 3}},
	},
	{
		param{[]int{2, 3, 5}, 10},
		ans{[]int{}},
	},
	{
		param{[]int{0, 1}, 1},
		ans{[]int{0, 1}},
	},
}

func TestTwoSum(t *testing.T) {
	fmt.Println("nums, target")
	for _, q := range qs {
		res := twoSum2(q.nums, q.target)
		if !isEqual(res, q.one) {
			t.Error("return error")
		}
		fmt.Printf("%v, %v\n", q.nums, q.target)
		fmt.Printf("%v\n", res)
		fmt.Println("-----------------------")
	}
}

func isEqual(a, b []int) bool {
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
