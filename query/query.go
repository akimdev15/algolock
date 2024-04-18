package query

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
)

const (
	URL = "https://leetcode.com/graphql/"
)

type QuestionResponse struct {
	Data QuestionData `json:"data"`
}

type QuestionData struct {
	Question Question `json:"question"`
}

type Question struct {
	QuestionID string `json:"questionId"`
	Title      string `json:"title"`
	Difficulty string `json:"difficulty"`
	Likes      int    `json:"likes"`
	Dislikes   int    `json:"dislikes"`
}

// ---------------------------------------------

type SubmissionResponse struct {
	Data SubmissionData `json:"data"`
}

type SubmissionData struct {
	Submission []Submission `json:"recentAcSubmissionList"`
}

type Submission struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	TitleSlug string `json:"titleSlug"`
	Timestamp string `json:"timestamp"`
}

// ---------------------------------------------

// GetQuestionDetails - get question id and title with question name
func GetQuestionDetails(questionName string) (Question, error) {

	query := `
		query questionTitle($titleSlug: String!) {
			question(titleSlug: $titleSlug) {
				questionId
				title
				difficulty
				likes
				dislikes
			}
		}
	`

	variables := map[string]interface{}{
		"titleSlug": questionName,
	}
	// Send the GraphQL request
	response, err := sendGraphQLRequest(URL, query, variables)
	if err != nil {
		fmt.Println("Error sending GraphQL request:", err)
		return Question{}, err
	}

	// Unmarshal JSON response into structs
	var responseData QuestionResponse
	if err := json.Unmarshal(response, &responseData); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return Question{}, err
	}

	return responseData.Data.Question, nil
}

func GetRecentSubmissions(limit int) ([]Submission, error) {
	username := getLeetcodeUsername()
	query := `
		query recentAcSubmissions($username: String!, $limit: Int!) {
		  recentAcSubmissionList(username: $username, limit: $limit) {
				id
				title
				titleSlug
				timestamp
		  }
		}
	`

	variables := map[string]interface{}{
		"username": username,
		"limit":    limit,
	}

	// Send the GraphQL request
	response, err := sendGraphQLRequest(URL, query, variables)
	if err != nil {
		fmt.Println("Error sending GraphQL request:", err)
		return nil, err
	}

	// Unmarshal JSON response into structs
	var responseData SubmissionResponse
	if err := json.Unmarshal(response, &responseData); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return nil, err
	}

	return responseData.Data.Submission, nil

}

func sendGraphQLRequest(url, query string, variables map[string]interface{}) ([]byte, error) {
	// Create GraphQL request body
	requestBody := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// Send POST request to GraphQL endpoint
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func getLeetcodeUsername() string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	username := os.Getenv("username")
	if username == "" {
		fmt.Println("Username not set. Please run the command setup.")
		os.Exit(1)
	}

	return username
}
