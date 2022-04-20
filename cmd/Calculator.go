package cmd

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
	totalLength := len(scores)

	if totalLength == 1 {
		return 100
	}

	for i := 0; i < totalLength; i++ {

		if userScore < scores[i] {
			worstQuizer++
		}
	}

	return worstQuizer * 100 / totalLength
}
