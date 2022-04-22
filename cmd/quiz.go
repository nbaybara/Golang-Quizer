/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

func welcoming() {

	boldfont()
	blueColor()
	fmt.Printf("Welcome to Quiz, Quizer :) Are you ready to challenge ? Let's start \n")
}

func startQuiz() {

	resetColor()
	italicFont()
	totalScore := askQuestion(questions)
	finishQuiz(totalScore)
}

func giveBriefOfQuiz(userScore int, totalQuestions int) {

	whiteColor()
	scores = append(scores, userScore)
	average := calculatePercentageOfScore(userScore, scores)

	fmt.Fprintf(os.Stdout, " \nYou answered %d questions correctly (%d / %d)\n", userScore,
		userScore, totalQuestions)

	magentaColor()
	textUnderline()
	fmt.Fprintf(os.Stdout, " \n Your score is better than %s%d \n", "%", average)

}

func finishQuiz(totalScore int) {

	giveBriefOfQuiz(totalScore, len(questions.Questions))
	restartQuiz()
}

func restartQuiz() {

	var answer string
	blueColor()
	boldfont()

	fmt.Printf("Do u want to start test?: yes | no ?\n")
	whiteColor()

	fmt.Scan(&answer)

	if checkAnswer(answer, "yes") {

		startQuiz()
		return
	}
	redColor()
	fmt.Println("Thank you for testing my project. Feel free to contribute it :)")
	os.Stdin.Close()
}

func init() {

	/*	We dont need any backend services
		for now but in case we need to listen server....
		handleRequests() */

	readJson()
	welcoming()
	startQuiz()
}
