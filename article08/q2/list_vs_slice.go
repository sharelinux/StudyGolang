package q2

import (
	"container/list"
	"fmt"
	"strconv"
	"time"
)

func main() {
	t0 := time.Now()

	// Version 1: use container list.
	for i := 0; i < 10000000; i++ {
		// New list.
		l := list.New()
		l.PushBack(strconv.Itoa(i))
	}

	t1 := time.Now()

	// Version 2: use slice
	for i := 0; i < 10000000; i++ {
		s := []string{}
		s = append(s, strconv.Itoa(i))
	}

	t2 := time.Now()
	// Results
	fmt.Printf("Contarner time: %s\n", t1.Sub(t0))
	fmt.Printf("Slice time: %s\n", t2.Sub(t1))
}
