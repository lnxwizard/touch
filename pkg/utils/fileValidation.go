package utils

import (
	"fmt"
	"os"
)

func FileValidation(filePath, cwd string) {
	// check if the file exists in the current directory
	fileInfo, err := os.Stat(filePath)

	// The file exists.
	if err == nil {
		fmt.Println(fileInfo.Name(), "already exists in path:", cwd)
		os.Exit(0)
	}

	// Is other error
	if err != nil && !os.IsNotExist(err) {
		fmt.Println(err)
		os.Exit(1)
	}
}
