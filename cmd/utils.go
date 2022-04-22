package cmd

import (
	"fmt"
	"log"
	"strings"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func seperateLines() {
	whiteColor()
	fmt.Printf("\n * * * * * * * * * * \n")

}

func checkAnswer(ans string, expectedAnswer string) bool {
	if strings.Compare(strings.Trim(strings.ToLower(ans), "\n"), strings.ToLower(expectedAnswer)) == 0 {
		return true
	}
	return false
}
