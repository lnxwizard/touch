package touch

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func CreateFile(fileName string) {
	// get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// check if the file exists in the current directory
	fileInfo, err := os.Stat(filepath.Join(cwd, fileName))
	if err != nil {
		if os.IsNotExist(err) {
			// create file
			createFile, err := os.Create(fileName)
			if err != nil {
				fmt.Println("Error while creating ", fileName, ":", err)
				os.Exit(1)
			}
			// close file
			defer createFile.Close()

			// set foreground green
			color.Set(color.FgGreen)

			fmt.Print("Touching ")

			// unset foreground color
			color.Unset()

			fmt.Println(fileName)
		} else {
			fmt.Println(err)
		}
		return
	}

	// The file exists.
	fmt.Println(fileInfo.Name(), "already exists in path:", cwd)
}
