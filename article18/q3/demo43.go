package main

import "fmt"

func main() {
	// 示例1
	//value3 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//switch value3[4] { // 这条语句无法通过编译，由于case表达式的条件重复
	//case 0, 1, 2:
	//	fmt.Println("0 or 1 or 2")
	//case 2, 3, 4:
	//	fmt.Println("2 or 3 or 4")
	//case 4, 5, 6:
	//	fmt.Println("4 or 5 ot 6")
	//}

	// 示例2
	value4 := [...]int{0, 1, 2, 3, 4, 5, 6}
	switch value4[4] {
	case value4[0], value4[1], value4[2]:
		fmt.Println("0 or 1 or 2")
	case value4[2], value4[3], value4[4]:
		fmt.Println("2 or 3 or 4")
	case value4[4], value4[5], value4[6]:
		fmt.Println("4 or 5 ot 6")
	}

	// 示例3
	//value5 := interface{}(byte(127))
	//switch t := value5.(type) { // 这条语句无法编译通过。 因为byte是unit8的别名类型，case语句条件重复
	//case uint8, uint16:
	//	fmt.Println("uint8 or unint16")
	//case byte:
	//	fmt.Println("byte")
	//default:
	//	fmt.Printf("unsupported type: %T", t)
	//}
}
