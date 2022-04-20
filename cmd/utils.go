package cmd

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func seperateLines() {
	color.Set(color.FgHiWhite)
	fmt.Printf("\n * * * * * * * * * * \n")

}
