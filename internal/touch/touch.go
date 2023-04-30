package touch

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func CreateFile(fileName string) {
	// create file
	os.Create(fileName)

	// set foreground green
	color.Set(color.FgGreen)

	fmt.Print("Touching ")

	// unset foreground color
	color.Unset()

	fmt.Println(fileName)
}
