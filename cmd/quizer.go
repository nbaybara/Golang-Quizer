package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var scores []int
var questions Questions

func askQuestion(question Questions) int {

	var totalScore = 0
	questions := question.Questions

	done := make(chan string)

	totalQuestions := len(questions)

	for i := 0; i < totalQuestions; i++ {

		explanation := questions[i].Explanation
		question := questions[i].Question
		answers := questions[i].Answers
		correctAnswer := questions[i].CorrectAnswers

		ans, err := eachQuestion(question, answers, correctAnswer, explanation, done)
		seperateLines()
		if err != nil && ans == -1 {
			return totalScore
		}
		totalScore += ans
	}
	return totalScore
}

func eachQuestion(Quest string, answers []string, correctAnswer string, explanation string, done <-chan string) (int, error) {

	var answer string
	yellowColor()
	fmt.Println("\nQuestion:\n ")

	fmt.Printf("%s: \n", Quest)

	cyanColor()
	fmt.Println("\nAnswers:\n ")

	for x := 0; x < len(answers); x++ {
		fmt.Println("-", answers[x])
	}

	whiteColor()
	fmt.Println("\nPlease Type your answer?:\n ")

	magentaColor()
	fmt.Scan(&answer)
	return getEachAnswer(answer, correctAnswer, explanation)
}

func getEachAnswer(answer string, correctAnswer string, explanation string) (int, error) {

	score := 0
	if checkAnswer(answer, correctAnswer) {

		rightAnswer()
		score = 1
		return score, nil
	}

	wrongAnswer(explanation)
	return 0, fmt.Errorf("Wrong Answer")
}

func rightAnswer() {
	greenColor()
	fmt.Printf("\nWell done! Yo go girl/boy? \n ")
}

func wrongAnswer(explanation string) {
	redColor()
	fmt.Printf("\nVuppsy dupsy!\n ")
	boldfont()
	greenColor()
	fmt.Printf("\nexplanation :  %s ", explanation)
}

var startQuizCmd = &cobra.Command{
	Use:   "startQuiz",
	Short: "A brief description of your command",
	Long:  `A long.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("startQuiz called")
	},
}

func init() {

	rootCmd.AddCommand(startQuizCmd)
}
