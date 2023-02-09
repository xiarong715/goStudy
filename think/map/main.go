package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// make 分配空间定义map
	m := make(map[string]int, 8)
	m["hello"] = 1
	fmt.Println(m)
	fmt.Printf("len = %d\n", len(m))

	// 直接定义 map
	m = map[string]int{
		"tom": 2,
		"jim": 3,
		"dom": 4,
	}

	// 判断某个key是否存在
	value, ok := m["jim"] // 存在时，ok为true
	if ok {
		fmt.Println(value)
	}

	// 遍历map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 删除键值对
	delete(m, "dom") // 删除key为dom的键值对
	fmt.Println(m)

	// 按照加入的顺序遍历map
	rand.Seed(time.Now().UnixNano()) // 随机种子

	scoreMap := make(map[string]int, 100)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		scoreMap[key] = rand.Intn(100)
	}

	keys := make([]string, 0, 100)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, scoreMap[k])
	}
}

// 数值类型、布尔类型、字符串类型、组数类型、结构体类型、指针类型、接口类型、切片类型、Map类型、Chan类型
