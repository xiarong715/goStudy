package main

import "fmt"

var (
	coins        = 50
	users        = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	distribution = make(map[string]int, len(users))
)

// 边遍历边统计
func dispatchCoin() int {
	sum := 0
	for _, name := range users {
		count := 0
		for _, v := range name {
			if v == 'e' || v == 'E' {
				count++
			} else if v == 'i' || v == 'I' {
				count += 2
			} else if v == 'o' || v == 'O' {
				count += 3
			} else if v == 'u' || v == 'U' {
				count += 4
			}
		}
		distribution[name] = count
		sum += count
	}
	for name, coins := range distribution {
		fmt.Printf("%s get coins: %d\n", name, coins)
	}
	return 50 - sum
}
