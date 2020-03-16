package common

import (
	"fmt"
	"strings"
	"time"
)

func ReadCommonInputs() (bool, int) {
	// shuffle questions
	fmt.Print("Shuffle problems? (Y/N) ")
	var shuffleResponse string
	fmt.Scan(&shuffleResponse)
	var shouldShuffle bool = strings.ToLower(strings.TrimSpace(shuffleResponse)) == "y"

	// no of questions
	fmt.Print("No of problems = ")
	var numProblems int
	fmt.Scan(&numProblems)

	return shouldShuffle, numProblems
}

func RunQuiz(numProblems int) {
	// show score
	defer ShowScore()

	// accept answers
	for idx := 0; idx < numProblems; idx++ {
		// get question
		var question Problem = GetProblem(idx)
		// show question
		ShowProblem(question)
		// start timer
		nowTime := time.Now()
		// wait for and read input
		var ans int
		fmt.Scan(&ans)
		// stop timer
		duration := time.Since(nowTime).Seconds()
		// parse response
		response := Response{
			QNo:  idx,
			Ans:  ans,
			Time: float32(duration),
		}
		// update score
		UpdateScore(response)
	}
}
