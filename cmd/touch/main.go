package main

import (
	"fmt"
	"os"

	"github.com/AlperAkca79/touch/internal/touch"
)

func main() {
	// check length of command-line arguments, if not equals 2 print usage of touch else create file
	if len(os.Args) < 2 {
		fmt.Println("Usage: touch <filename>")
	} else {
		// multiple file creating support
		for i := 1; i < len(os.Args); i++ {
			touch.CreateFile(os.Args[i])
		}
	}
}
