package random

import (
	"github.com/akimdev15/algolock/cmd/jsonutils"
	"testing"
)

func TestPickRandomQuestion(t *testing.T) {
	var emptyQuestions []jsonutils.Question
	_, err := pickRandomQuestion(emptyQuestions)
	if err == nil || err.Error() != "no questions found" {
		t.Error("Expected error for empty questions slice, got nil")
	}

	nonEmptyQuestions := []jsonutils.Question{
		{ID: "001", Name: "Question 1", URL: "https://example.com/question1", Solved: 0},
		{ID: "002", Name: "Question 2", URL: "https://example.com/question2", Solved: 0},
		{ID: "003", Name: "Question 3", URL: "https://example.com/question3", Solved: 0},
	}

	randomQuestion, err := pickRandomQuestion(nonEmptyQuestions)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	found := false
	for _, question := range nonEmptyQuestions {
		if question == randomQuestion {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Question %v not found", randomQuestion)
	}
}
