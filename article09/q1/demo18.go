package main

import "fmt"

func main() {
	//// 示例1
	//var badMap1 = map[[]int]int{} // 字典的键类型不可以是函数类型、字典类型和切片类型
	//_ = badMap1

	//// 示例2
	//var badMap2 = map[interface{}]int{
	//	"1":      1,
	//	[]int{2}: 2, //如果键类型是接口，那么键实际类型不能是函数类型、字典类型、切片类型；否则会引起panic
	//	3:        3,
	//}
	//_ = badMap2

	//// 示例3
	//var badMap3 map[[1][]string]int // 如果键的类型是数组类型，那么还要确保该类型的元素类型不是函数类型、字典类型或切片类型，这里会引发编译错误。
	//_ = badMap3

	//var t = [3][]string{} // 声明长度为3的字符串切片的数组
	//t[0] = []string{"a"}
	//t[1] = []string{"b"}
	//t[2] = []string{"c"}
	//fmt.Printf("%#v\n", t)

	//// 示例4
	//type BadKey1 struct {
	//	slice []string
	//}
	//
	//var badMap4 map[BadKey1]int // 如果键的类型是结构体类型，那么还要保证其中字段的类型的合法性。不能是函数类型、字典类型或切片类型
	//_ = badMap4

	//// 示例5
	//var badmap5 map[[1][2][3][]string]int // 如果键的类型是数组类型，那么还要确保该类型的元素类型不是函数类型、字典类型或切片类型; 无论埋的多深，编译器都会找出来
	//_ = badmap5

	//// 示例6
	//type BadKey2Field1 struct {
	//	slice []string
	//}
	//type BadKey2 struct {
	//	field BadKey2Field1
	//}
	//var badMap6 map[BadKey2]int // 如果键的类型是结构体类型，那么还要保证其中字段的类型的合法性。不能是函数类型、字典类型或切片类型, 无论埋藏多深;
	//_ = badMap6

	goodMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	// 遍历map
	for k, v := range goodMap {
		fmt.Printf("%s=%d\n", k, v)
	}
}
