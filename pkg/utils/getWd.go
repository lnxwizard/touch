package utils

import (
	"fmt"
	"os"
)

func Getwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return cwd
}
