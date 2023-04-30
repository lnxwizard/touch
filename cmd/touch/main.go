package main

import (
	"fmt"
	"os"

	"github.com/AlperAkca79/touch/internal/touch"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: touch fileName")
	} else {
		touch.CreateFile(os.Args[1])
	}
}
