package main

import (
	"fmt"
	"github.com/y2k-shubham/gophercises-y2k/quiz/common"
	"time"
)

func main() {
	// show score
	defer common.ShowScore()

	shouldShuffleProblems, numProblems := common.ReadCommonInputs()
	if shouldShuffleProblems {
		common.ShuffleProblems()
	}

	// time limit
	fmt.Print("Time limit (seconds) = ")
	var timeLimitSecs int
	fmt.Scan(&timeLimitSecs)

	// create channel to allow interrupt
	stopchan := make(chan struct{})
	// run quiz and return when iterrupted
	go func() {
		for {
			select {
			default:
				common.RunQuiz(numProblems)
			case <-stopchan:
				return
			}
		}
	}()

	// sleep for timeout duration and then interrupt
	time.Sleep(time.Second * time.Duration(timeLimitSecs))
	close(stopchan)
}
