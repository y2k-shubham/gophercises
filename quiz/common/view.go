package common

type Question struct {
	Op1 int
	Op2 int
	Ans int
}

type Score struct {
	CorrectQNo    []int
	IncorrectQNo  []int
	UnansweredQNo []int
	QTime         []float32
}
