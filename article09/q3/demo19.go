package main

import "fmt"

func main() {
	var m map[string]int
	//fmt.Printf("%#v\n", m)

	key := "two"
	elem, ok := m[key]
	fmt.Printf("The element paired with key %q in nil map: %d (%v)\n",
		key, elem, ok) // 在nil map上读取不会发生panic;

	fmt.Printf("The length of nil map: %d\n", len(m)) // nil map的长度为0

	fmt.Printf("Delete the key-element pair by key %q...\n",
		key)
	delete(m, key) //在nil map上删除不会发生panic;

	//elem = 2
	//fmt.Printf("Add a key-element pair to a nil map...")
	//m[key] = elem // 在nil map生新增键值对会发生panic;

	fmt.Println()

	m1 := make(map[string]int)
	m1["one"] = 1 // map 增加键值对
	m1["two"] = 2
	m1["three"] = 3

	for k, v := range m1 {
		fmt.Printf("%s=%d\n", k, v) // 遍历map，map是无序的
	}

	fmt.Println()
	m1["two"] = 22 // 对map 进行某个键的更新操作
	fmt.Printf("%+v\n", m1)

	delete(m1, "tree")
	fmt.Printf("%#v\n", m1) // 对map进行删除某个键操作
}
