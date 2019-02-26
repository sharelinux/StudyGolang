package main

import "fmt"

func main() {
	// 示例1
	//value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//switch 1 + 3 { // 这条语句无法通过编译，由于case表达式的值类型必须相同或者能够都统一到switch表达式的结果类型
	//case value1[0], value1[1]:
	//	fmt.Println("0 or 1")
	//case value1[2], value1[3]:
	//	fmt.Println("2 or 3")
	//case value1[4], value1[5], value1[6]:
	//	fmt.Println("4 or 5 or 6")
	//}

	// 示例2
	value2 := [...]int{0, 1, 2, 3, 4, 5, 6}
	switch value2[4] {
	case 0, 1:
		fmt.Println("0 or 1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}
}
