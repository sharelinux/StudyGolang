package main

import (
	"errors"
	"fmt"
)

func echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("Empty request")
		return
	}
	response = fmt.Sprintf("echo: %s", request)
	return
}

func main() {
	// 示例1
	for _, req := range []string{"", "Hello"} {
		fmt.Printf("request: %s\n", req)
		resp, err := echo(req)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		fmt.Printf("response: %s\n", resp)
	}
	fmt.Println()

	// 示例2
	err1 := errors.New(fmt.Sprintf("Invalid contents: %s", "#$%"))
	err2 := fmt.Errorf("Invalid contents: %s", "#$%")
	if err1.Error() == err2.Error() {
		fmt.Printf("The error messages in err1 and err2 are the same.")
	}
}
