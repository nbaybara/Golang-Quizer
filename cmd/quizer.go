package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var scores []int
var questions Questions

func startQuiz() {

	color.Set(color.Italic)
	totalScore := askQuestion(questions)
	giveBriefOfQuiz(totalScore, len(questions.Questions))
	finishQuizer()
	color.Set(color.Reset)
}

func finishQuizer() {

	done := make(chan string)
	go getInput(done)

	fmt.Println("Do u want to tested again?: yes - no ??:")

	for {
		select {
		case ans := <-done:
			if checkAnswer(ans, "yes") {
				startQuiz()
			} else {
				color.Set(color.FgRed)
				fmt.Println("Thank you for testing my project. Feel free to contribute it :)")
				os.Stdin.Close()
			}

		}

	}
}

func checkAnswer(ans string, expectedAnswer string) bool {
	if strings.Compare(strings.Trim(strings.ToLower(ans), "\n"), strings.ToLower(expectedAnswer)) == 0 {
		return true
	}
	return false
}

func giveBriefOfQuiz(userScore int, totalQuestions int) {

	color.Set(color.Underline)
	color.Set(color.FgWhite)
	scores = append(scores, userScore)
	average := calculatePercentageOfScore(userScore, scores)

	fmt.Fprintf(os.Stdout, " \nYou answered %d questions correctly (%d / %d)\n", userScore,
		userScore, totalQuestions)
	fmt.Fprintf(os.Stdout, " \n Your score is better than %s%d \n", "%", average)

}

func askQuestion(question Questions) int {

	var totalScore = 0
	questions := question.Questions

	done := make(chan string)
	go getInput(done)

	totalQuestions := len(questions)

	for i := 0; i < totalQuestions; i++ {
		ans, err := eachQuestion(questions[i].Question, questions[i].Answers, questions[i].CorrectAnswers, done)
		if err != nil && ans == -1 {
			return totalScore
		}
		totalScore += ans
	}
	return totalScore
}

func eachQuestion(Quest string, answers []string, correctAnswer string, done <-chan string) (int, error) {

	color.Set(color.FgHiWhite)
	fmt.Println("\nQuestion:\n ")
	fmt.Printf("%s: \n", Quest)

	color.Set(color.FgHiCyan)
	fmt.Println("\nAnswers:\n ")

	for x := 0; x < len(answers); x++ {
		fmt.Println("-", answers[x])
	}

	color.Set(color.FgHiYellow)
	fmt.Println("\nPlease Type your answer?:\n ")
	color.Set(color.FgHiWhite)

	for {
		select {

		case ans := <-done:
			score := 0

			if checkAnswer(ans, correctAnswer) {
				color.Set(color.FgHiGreen)
				fmt.Printf("\nWell done! Yo go girl/boy? \n ")
				score = 1

			} else {
				color.Set(color.FgHiRed)
				fmt.Printf("\nVuppsy dupsy! Try better new time ;) \n ")
				return 0, fmt.Errorf("Wrong Answer")
			}
			seperateLines()
			return score, nil
		}

	}
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

func getInput(input chan string) {

	for {
		in := bufio.NewReader(os.Stdin)
		result, err := in.ReadString('\n')
		handleError(err)
		input <- result
	}
}
