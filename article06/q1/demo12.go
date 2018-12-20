package main

import "fmt"

var container = []string{"one", "two", "three"}

func main() {
	container := map[int]string{1:"one", 2:"two", 3:"three"}

	_, ok1 := interface{}(container).([]string)
	_, ok2 := interface{}(container).(map[int]string)

	if !(ok1 || ok2) {
		fmt.Printf("Error: unsupported cpmtainer type: %T\n", container)
		return
	}
	fmt.Printf("The element is %q. (container type: %T)\n",
			container[1], container)
 }