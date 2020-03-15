package common

import (
	"bufio"
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

func ReadQuestions(filePath string) []Question {
	questions := []Question{}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var question Question
		question = question.Parse(line)
		questions = append(questions, question)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return questions
}

func ShowQuestions(questions []Question) {
	fmt.Println("Questions are:-")
	for _, question := range questions {
		fmt.Println(question.ToString(false))
	}
	//fmt.Printf("%v", questions)
}
