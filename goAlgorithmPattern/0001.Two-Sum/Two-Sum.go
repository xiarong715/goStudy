package leetcode

func twoSum1(num []int, target int) []int {
	len := len(num)
	for i := 0; i <= len/2; i++ {
		for j := i + 1; j < len; j++ {
			if num[i]+num[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum2(num []int, target int) []int {
	m := make(map[int]int)
	for k, v := range num {
		if _, ok := m[target-v]; !ok {
			m[target-v] = k
		} else {
			return []int{m[target-v], k}
		}
	}
	return []int{}
}

// [1,2,3,4,5]   8
