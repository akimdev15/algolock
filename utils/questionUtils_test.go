package utils

import (
	"github.com/akimdev15/leetcode/cmd/jsonutils"
	"testing"
)

func TestPickRandomQuestion(t *testing.T) {
	var emptyQuestions []jsonutils.Question
	_, err := PickRandomIdx(len(emptyQuestions))
	if err == nil || err.Error() != "provided length is 0" {
		t.Error("Expected error for empty questions slice, got nil")
	}

	nonEmptyQuestions := []jsonutils.Question{
		{ID: "001", Name: "Question 1", URL: "https://example.com/question1", Solved: 0},
		{ID: "002", Name: "Question 2", URL: "https://example.com/question2", Solved: 0},
		{ID: "003", Name: "Question 3", URL: "https://example.com/question3", Solved: 0},
	}

	randomIdx, err := PickRandomIdx(len(nonEmptyQuestions))
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if nonEmptyQuestions[randomIdx].ID == "" {
		t.Errorf("Expected non empty question ID, got %s", nonEmptyQuestions[randomIdx].ID)
	}

}
