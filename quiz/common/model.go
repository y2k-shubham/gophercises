package common

import "fmt"

var questions []Problem = ReadQuestions(QuizPath() + "questions.txt")
var score Score = Score{
	CorrectQNo:   []int{},
	IncorrectQNo: []int{},
	QTime:        []float32{},
}

func GetQuestion(qNo int) Problem {
	return questions[qNo]
}

func UpdateScore(response Response) {
	score.Update(GetQuestion(response.QNo), response)
}

func ShowScore() {
	fmt.Println(score.ToString())
}
