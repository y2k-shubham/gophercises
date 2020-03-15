package common

import (
	"fmt"
	"strconv"
)

type Problem struct {
	Question string
	Answer   int
}

func (q Problem) Parse(parsedLine []string) Problem {
	Question := parsedLine[0]
	Answer, _ := strconv.Atoi(parsedLine[1])
	return Problem{
		Question: Question,
		Answer:   Answer,
	}
}

func (q Problem) ToString(maskAnswer bool) string {
	if maskAnswer {
		return fmt.Sprintf("%s = ?", q.Question)
	} else {
		return fmt.Sprintf("%s = %d", q.Question, q.Answer)
	}
}

type Response struct {
	QNo  int
	Ans  int
	Time float32
}

type Score struct {
	CorrectQNo   []int
	IncorrectQNo []int
	QTime        []float32
}

func (s Score) AvgTime() float32 {
	var sum float32 = 0.0
	for _, ele := range s.QTime {
		sum += ele
	}
	return (sum / float32(len(s.QTime)))
}

func (s Score) ToString() string {
	return fmt.Sprintf("CorrectQNo:\t%v\nIncorrectQNo:\t%v\nTimes:\t\t%v\nAvgTime:\t%f", s.CorrectQNo, s.IncorrectQNo, s.QTime, s.AvgTime())
}

func (s *Score) Update(question Problem, response Response) {
	//fmt.Println("before update, score was:-\n" + s.ToString())
	adjustedQNo := response.QNo + 1
	if response.Ans == question.Answer {
		s.CorrectQNo = append(s.CorrectQNo, adjustedQNo)
	} else {
		s.IncorrectQNo = append(s.IncorrectQNo, adjustedQNo)
	}
	s.QTime = append(s.QTime, response.Time)
	//fmt.Println("after update, score was:-\n" + s.ToString())
}

func ShowProblem(q Problem) {
	fmt.Println(q.ToString(true))
}
