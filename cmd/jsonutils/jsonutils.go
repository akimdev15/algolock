package jsonutils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Question struct {
	ID         string `json:"ID"`
	Name       string `json:"Name"`
	URL        string `json:"Url"`
	Solved     string `json:"Solved"`
	Difficulty string `json:"Difficulty"`
	UpdatedAt  string `json:"UpdatedAt"`
	Confidence string `json:"Confidence"`
}

func OpenJsonFile() ([]Question, error) {
	file, err := os.Open("sample/sample_questions.json")
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
