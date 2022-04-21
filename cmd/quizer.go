package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var scores []int
var questions Questions

func checkAnswer(ans string, expectedAnswer string) bool {
	if strings.Compare(strings.Trim(strings.ToLower(ans), "\n"), strings.ToLower(expectedAnswer)) == 0 {
		return true
	}
	return false
}

func giveBriefOfQuiz(userScore int, totalQuestions int) {

	color.Set(color.FgWhite)
	scores = append(scores, userScore)
	average := calculatePercentageOfScore(userScore, scores)

	fmt.Fprintf(os.Stdout, " \nYou answered %d questions correctly (%d / %d)\n", userScore,
		userScore, totalQuestions)

	color.Set(color.FgMagenta)

	color.Set(color.Underline)
	fmt.Fprintf(os.Stdout, " \n Your score is better than %s%d \n", "%", average)

}

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

	color.Set(color.FgHiRed)
	fmt.Println("\nQuestion:\n ")

	color.Set(color.FgHiWhite)
	fmt.Printf("%s: \n", Quest)

	color.Set(color.FgHiCyan)
	fmt.Println("\nAnswers:\n ")

	for x := 0; x < len(answers); x++ {
		fmt.Println("-", answers[x])
	}

	color.Set(color.FgHiYellow)
	fmt.Println("\nPlease Type your answer?:\n ")

	color.Set(color.FgHiMagenta)
	var answer string
	fmt.Scan(&answer)

	return getEachAnswer(answer, correctAnswer, explanation)
}

func getEachAnswer(answer string, correctAnswer string, explanation string) (int, error) {

	score := 0
	if checkAnswer(answer, correctAnswer) {

		color.Set(color.FgHiGreen)
		fmt.Printf("\nWell done! Yo go girl/boy? \n ")
		score = 1
		return score, nil
	}

	color.Set(color.FgHiRed)
	fmt.Printf("\nVuppsy dupsy!\n ")

	color.Set(color.FgHiGreen)
	fmt.Printf("\nexplanation :  %s ", explanation)

	return 0, fmt.Errorf("Wrong Answer")
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
