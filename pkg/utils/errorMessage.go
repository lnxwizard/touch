package utils

import (
	"fmt"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

func ErrorMessage(fileName string, err error) {
	// error message
	errorMessage := "Error while creating " + fileName + ":" + err.Error()

	// marking word "Error"
	emphasized := marker.Mark(errorMessage, marker.MatchAll("Error"), color.New(color.FgRed))

	// Printing error message
	fmt.Println(emphasized)
}
