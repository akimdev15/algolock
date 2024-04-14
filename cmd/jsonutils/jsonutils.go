package jsonutils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Question struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	Solved int    `json:"solved"`
}

func OpenJsonFile() ([]Question, error) {
	file, err := os.Open("cmd/questions.json")
	if err != nil {
		fmt.Println("open questions.json error:", err)
		return nil, err
	}
	defer file.Close()

	// decode
	decoder := json.NewDecoder(file)

	var questions []Question
	if err := decoder.Decode(&questions); err != nil {
		fmt.Println("parse questions.json error:", err)
		return nil, err
	}

	return questions, nil
}
