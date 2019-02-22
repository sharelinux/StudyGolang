package main

import "fmt"

func main() {
	// 办法1
	/*
		for i := 0; i < 10; i++ {
			go func() {
				fmt.Println(i)
			}()
		}

		time.Sleep(time.Millisecond * 500)
	*/

	// 办法2
	num := 10
	sign := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}

	for j := 0; j < num; j++ {
		<-sign
	}
}
