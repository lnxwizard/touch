package main

import (
	"fmt"
	"os"

	"github.com/AlperAkca79/touch/internal/touch"
)

func main() {
	// check length of command-line arguments, if not equals 2 print usage of touch else create file
	if len(os.Args) != 2 {
		fmt.Println("Usage: touch fileName")
	} else {
		touch.CreateFile(os.Args[1])
	}
}
