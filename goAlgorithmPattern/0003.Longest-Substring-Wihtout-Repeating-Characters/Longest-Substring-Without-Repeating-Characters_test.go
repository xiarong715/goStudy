package leetcode

import (
	"testing"
)

type question struct {
	params
	ans
}

type params struct {
	arr string
}

type ans struct {
	num int
}

var qs []question = []question{
	{
		params{"abcabdefg"},
		ans{7},
	},
	{
		params{"abcabcbb"},
		ans{3},
	},
	{
		params{"bbbbb"},
		ans{1},
	},
	{
		params{"pwwkew"},
		ans{3},
	},
}

func TestFindLongestSubstr(t *testing.T) {
	for _, item := range qs {
		res := findLongestSubStr(item.arr)
		if res != item.num {
			t.Error("return error")
		}
	}
}
