package main

import "fmt"

func main() {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	s1 := []int{0, 1, 2, 3, 4}
	e5 := s1[5] // 切片索引访问越界
	_ = e5
}
