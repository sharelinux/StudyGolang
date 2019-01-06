package main

import (
	"container/list"
	"fmt"
)

func main() {
	// New list.
	l := list.New()
	l.PushBack("bird") // 在列表l后面插入value值
	l.PushBack("cat")
	l.PushFront("snake") // 在列表l前面插入value值

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
