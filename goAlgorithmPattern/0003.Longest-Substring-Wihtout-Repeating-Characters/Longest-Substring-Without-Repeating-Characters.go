package leetcode

// abcabdefg

func findLongestSubStr(arr string) int {
	maxLongest := 0
	tempMap := make(map[byte]int)
	leftPointer := 0
	rightPointer := 0
	length := len(arr)
	for range arr {
		if _, ok := tempMap[arr[rightPointer]]; ok {
			leftPointer = tempMap[arr[rightPointer]] + 1
			delete(tempMap, arr[rightPointer])
			if leftPointer > length {
				leftPointer = length
			}
		}
		tempMap[arr[rightPointer]] = rightPointer
		tempMax := rightPointer - leftPointer + 1
		if tempMax > maxLongest {
			maxLongest = tempMax
		}
		rightPointer++
	}
	return maxLongest
}
