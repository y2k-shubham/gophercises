package common

import "fmt"

var problems []Problem = ReadProblems(QuizPath() + "questions.txt")
var score Score = Score{
	CorrectQNo:   []int{},
	IncorrectQNo: []int{},
	QTime:        []float32{},
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
