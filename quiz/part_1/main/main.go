package main

import (
	"github.com/y2k-shubham/gophercises-y2k/quiz/common"
)

func main() {
	// show score
	defer common.ShowScore()

	shouldShuffleProblems, numProblems := common.ReadCommonInputs()
	if shouldShuffleProblems {
		common.ShuffleProblems()
	}

	common.RunQuiz(numProblems)
}
