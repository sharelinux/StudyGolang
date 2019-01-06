package main

import (
	"container/list"
	"fmt"
)

func main() {
	// New list.
	l := list.New()
	l.PushBack("a") // 在列表l后面插入value值
	l.PushBack("b")
	l.PushBack("c")

	printList(l)

	e4 := l.PushFront("d")
	printList(l)
	l.MoveToBack(e4)  // 把给定的元素移动到链表的最后端
	l.MoveToFront(e4) // 把给定的元素移动到链表的最前端
	l.Remove(e4)      // 从链表l删除
	printList(l)

	e5 := l.PushBack("e")
	l.InsertBefore("d", e5)
	printList(l)

	e6 := l.PushBack("g")
	e7 := l.PushBack("f")

	printList(l)
	l.MoveAfter(e6, e7)  //把给定的元素移动到另一个元素的后面
	l.MoveBefore(e7, e6) //把给定的元素移动到另一个元素的前面
	printList(l)
}

// 正序打印链表
func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}

// 逆序打印链表
func printListR(l *list.List) {
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}
