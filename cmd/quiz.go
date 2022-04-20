/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"net/http"

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

type Answer struct {
	Answer string `json: "answer"`
}

type Answers struct {
	Answers []Answer `json:"answers"`
}

type CorrectAnswer struct {
	Answer string `json: "answer"`
}

type Question struct {
	Question       string   `json: "question"`
	Answers        []string `json: "answer"`
	CorrectAnswers string   `json: "correctAnswer"`
}

type Questions struct {
	Questions []Question `json:"questions"`
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", readJsonFile).Methods("GET")
	log.Fatal(http.ListenAndServe(":8083", router))
}

func init() {

	readJson()
	startQuiz()
	getQuestions(url)
	handleRequests()
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
