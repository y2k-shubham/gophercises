package common

import (
	"fmt"
	"math/rand"
	"time"
)

var problems []Problem = ReadProblems(QuizPath() + "questions.txt")
var score Score = Score{
	CorrectQNo:   []int{},
	IncorrectQNo: []int{},
	QTime:        []float32{},
}

func ShuffleProblems() {
	// https://stackoverflow.com/a/46185753/3679900
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})
}

func GetProblem(qNo int) Problem {
	return problems[qNo]
}

func UpdateScore(response Response) {
	score.Update(GetProblem(response.QNo), response)
}

func ShowScore() {
	fmt.Println(score.ToString())
}
