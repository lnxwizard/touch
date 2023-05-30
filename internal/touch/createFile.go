package touch

import (
	"os"
	"path/filepath"

	"github.com/AlperAkca79/touch/pkg/utils"
)

func CreateFile(fileName string) {
	// get current working directory
	var cwd = utils.Getwd()
	var filePath = filepath.Join(cwd, fileName)

	// Validation
	utils.FileValidation(filePath, cwd)

	// create file
	createFile, err := os.Create(filePath)
	if err != nil {
		utils.ErrorMessage(fileName, err)
		os.Exit(1)
	}

	// close file
	defer createFile.Close()

	// Message success create
	utils.TouchingMessage(fileName)
}
