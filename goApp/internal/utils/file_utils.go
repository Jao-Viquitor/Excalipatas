package utils

import (
	"bufio"
	"os"
)

func ReadQuestions(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var questions []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		questions = append(questions, scanner.Text())
	}
	return questions, scanner.Err()
}
