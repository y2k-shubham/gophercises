package common

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func QuizPath() string {
	_, path, _, _ := runtime.Caller(0)
	return strings.TrimSuffix(path, "common/utils.go")
}

func ReadProblems(filePath string) []Problem {
	problems := []Problem{}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create csv reader
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.TrimLeadingSpace = true

	parsedLines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, parsedLine := range parsedLines {
		answer, _ := strconv.Atoi(parsedLine[1])
		parsedProblem := Problem{
			Question: parsedLine[0],
			Answer:   answer,
		}
		problems = append(problems, parsedProblem)
	}

	return problems
}

func ShowProblems(problems []Problem) {
	fmt.Println("Problems are:-")
	for _, question := range problems {
		fmt.Println(question.ToString(false))
	}
	//fmt.Printf("%v", problems)
}
