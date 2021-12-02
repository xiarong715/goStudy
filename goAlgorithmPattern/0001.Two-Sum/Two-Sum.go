package leetcode

// 最初的想法
func twoSum1(nums []int, target int) []int {
	len := len(nums)
	for i := 0; i <= len/2; i++ {
		for j := i + 1; j < len; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 效率更高
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if _, ok := m[v]; !ok {
			m[target-v] = k
		} else {
			return []int{m[v], k}
		}
	}
	return []int{}
}

// 返回所有的组合
func twoSum3(nums []int, target int) [][]int {
	res := [][]int{}
	m := make(map[int]int)
	for k, v := range nums {
		if _, ok := m[v]; !ok {
			m[target-v] = k
		} else {
			res = append(res, []int{m[v], k})
		}
	}
	return res
}

// [1,2,3,4,5]   8
