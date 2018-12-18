package lib

import (
	"os"

	in "StudyGolang/article03/q3/lib/internal"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}