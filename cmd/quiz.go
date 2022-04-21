/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "quizer",
	Short: "A brief description of your application",
	Long:  `.`,
}

func Execute() {
	/*	err := rootCmd.Execute()
		if err != nil {
			os.Exit(1)
		}*/
}

type Question struct {
	Question       string   `string: "question"`
	Answers        []string `json: "answer"`
	CorrectAnswers string   `string: "correctAnswer"`
	Explanation    string   `string: "explanation"`
}

type Questions struct {
	Questions []Question `json:"questions"`
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", readJsonFile).Methods("GET")
	log.Fatal(http.ListenAndServe(":8083", router))
}

func startQuiz() {

	color.Set(color.Reset)
	color.Set(color.Italic)
	totalScore := askQuestion(questions)
	finishQuiz(totalScore)
}

func finishQuiz(totalScore int) {
	giveBriefOfQuiz(totalScore, len(questions.Questions))
	restartQuiz()

}

func restartQuiz() {

	color.Set(color.FgHiMagenta)
	var answer string
	color.Set(color.FgHiMagenta)
	fmt.Printf("Do u want to tested again?: yes | no ?\n")
	fmt.Scan(&answer)

	if checkAnswer(answer, "yes") {

		startQuiz()
		return
	}

	color.Set(color.FgRed)
	fmt.Println("Thank you for testing my project. Feel free to contribute it :)")
	os.Stdin.Close()

}

func init() {

	/*	We dont need any backend services
		for now but in case we need to listen server....
		handleRequests() */

	readJson()
	startQuiz()
}
