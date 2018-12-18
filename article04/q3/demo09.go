package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var err error
	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n") // 对`err`进行重声明
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("%d bype(s) were written.\n", n)
}
