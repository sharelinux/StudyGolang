package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")
	caller()
	fmt.Println("Exit function main.")
}

func caller() {
	fmt.Println("Enter function Caller.")
	panic(errors.New("Something wrong")) // 正例
	panic(fmt.Println)                   // 反例
	fmt.Println("Exit function Caller.")
}
