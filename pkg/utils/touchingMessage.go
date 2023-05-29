package utils

import (
	"fmt"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

func TouchingMessage(fileName string) {
	// complete message
	completeMessage := "Touching " + fileName

	// marking word "Touching"
	emphasized := marker.Mark(completeMessage, marker.MatchAll("Touching"), color.New(color.FgGreen))

	// Printing complete message
	fmt.Println(emphasized)
}
