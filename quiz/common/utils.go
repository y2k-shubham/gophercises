package common

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

func QuizPath() string {
	_, path, _, _ := runtime.Caller(0)
	return strings.TrimSuffix(path, "common/utils.go")
}

func ReadProblems(filePath string) []Problem {
	// open file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create csv reader
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.TrimLeadingSpace = true

	// read lines
	parsedLines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// parse lines
	problems := make([]Problem, len(parsedLines))
	for i, parsedLine := range parsedLines {
		problems[i] = Parse(parsedLine)
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
