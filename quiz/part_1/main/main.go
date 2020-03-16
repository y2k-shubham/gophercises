package main

import (
	"github.com/y2k-shubham/gophercises-y2k/quiz/common"
)

func main() {
	shouldShuffleProblems, numProblems := common.ReadCommonInputs()
	if shouldShuffleProblems {
		common.ShuffleProblems()
	}

	common.RunQuiz(numProblems)
}
