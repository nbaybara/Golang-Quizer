package cmd

import (
	"fmt"
	"os"
)

func calculateAverage(numbers []int) int {

	totalLength := len(numbers)
	totalScore := 0

	for i := 0; i < totalLength; i++ {
		totalScore += numbers[i]
	}
	return totalScore / totalLength
}

func calculatePercentageOfScore(userScore int, scores []int) int {

	worstQuizer := 0
	scoreEliminated := removeDuplicate(scores)
	totalLength := len(scoreEliminated)

	if totalLength == 1 {
		return 100
	}

	fmt.Fprintf(os.Stdout, " totalLength (%d )\n", userScore)

	for i := 0; i < totalLength; i++ {

		if userScore > scores[i] {
			worstQuizer++
		}
	}

	return worstQuizer * 100 / totalLength
}

func removeDuplicate(scores []int) []int {
	m := make(map[int]int)
	for _, x := range scores {
		m[x] = x
	}
	var ClearedArr []int
	for x, _ := range m {
		ClearedArr = append(ClearedArr, x)
	}
	return ClearedArr
}
