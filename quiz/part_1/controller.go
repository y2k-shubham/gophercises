package part_1

import (
	"fmt"
	"github.com/y2k-shubham/gophercises-y2k/quiz/common"
	_ "github.com/y2k-shubham/gophercises-y2k/quiz/common"
	"time"
)

func Drive() {
	// no of questions
	fmt.Print("No of questions = ")
	var numQuestions int
	fmt.Scan(&numQuestions)
	// accept answers
	for idx := 0; idx < numQuestions; idx++ {
		// get question
		var question common.Question = common.GetQuestion(idx)
		// show question
		common.ShowQuestion(question)
		// start timer
		nowTime := time.Now()
		// wait for and read input
		var ans int
		fmt.Scan(&ans)
		// stop timer
		duration := time.Since(nowTime).Seconds()
		// parse response
		response := common.Response{
			QNo:  idx,
			Ans:  ans,
			Time: float32(duration),
		}
		// update score
		common.UpdateScore(response)
	}
	// show score
	common.ShowScore()
}
