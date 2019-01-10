package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Printf("%d + %d = %d\n", x, y, x+y)
}

func sum(s []int, c chan<- int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
