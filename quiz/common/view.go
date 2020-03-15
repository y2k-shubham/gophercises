package common

import (
	"fmt"
	"regexp"
	"strconv"
)

type Question struct {
	Op1 int
	Op2 int
	Ans int
}

func (q Question) Parse(str string) Question {
	digitRegex := regexp.MustCompile("\\d+")
	numbers := digitRegex.FindAllString(str, -1)
	Op1, _ := strconv.Atoi(numbers[0])
	Op2, _ := strconv.Atoi(numbers[1])
	Ans, _ := strconv.Atoi(numbers[2])
	return Question{
		Op1: Op1,
		Op2: Op2,
		Ans: Ans,
	}
}

func (q Question) ToString(maskAnswer bool) string {
	if maskAnswer {
		return fmt.Sprintf("%d + %d = ?", q.Op1, q.Op2)
	} else {
		return fmt.Sprintf("%d + %d = %d", q.Op1, q.Op2, q.Ans)
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

func (s *Score) Update(question Question, response Response) {
	//fmt.Println("before update, score was:-\n" + s.ToString())
	adjustedQNo := response.QNo + 1
	if response.Ans == question.Ans {
		s.CorrectQNo = append(s.CorrectQNo, adjustedQNo)
	} else {
		s.IncorrectQNo = append(s.IncorrectQNo, adjustedQNo)
	}
	s.QTime = append(s.QTime, response.Time)
	//fmt.Println("after update, score was:-\n" + s.ToString())
}

func ShowQuestion(q Question) {
	fmt.Println(q.ToString(true))
}
