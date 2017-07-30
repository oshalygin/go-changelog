package main

import (
	"fmt"

	"github.com/fatih/color"
)

const directory string = "."

func main() {
	formatter := `commit:       %H%nAuthor:       %ae%nAuthor Date:  %cd%n%n%s%n%n%b`
	log, err := GitLog(directory, formatter)

	if err != nil {
		color.Red("Exiting...")
	}

	fmt.Println(log)
}
