package add

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/akimdev15/algolock/cmd/jsonutils"
	"github.com/akimdev15/algolock/internal/database"
	"github.com/akimdev15/algolock/sql"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"net/url"
	"os"
	"strings"
	"time"
)

// AddCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new leetcode question",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
			return
		}

		addNewQuestion(args[0])
	},
}

func init() {
}

// addNewQuestion - adds new Leetcode question to the jsonutils file
func addNewQuestion(urlStr string) {
	queries, err := sql.InitDB()
	if err != nil {
		fmt.Println("Error initializing database. Error: ", err)
		return
	}

	ctx := context.Background()
	// Get question name from the url
	questionName, err := getQuestionNameFromURL(urlStr)
	if err != nil {
		return
	}
	fmt.Println("Question name: " + questionName)

	question, err := queries.CreateQuestion(ctx, database.CreateQuestionParams{
		ID:        uuid.New().String(),
		Name:      questionName,
		Url:       urlStr,
		Solved:    "0",
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		fmt.Println("Error saving question: ", err)
		return
	}

	fmt.Println("Successfully saved question: ", question)
}

func getQuestionNameFromURL(urlStr string) (string, error) {
	// Parse the url
	u, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("parse url error:", err)
		return "", err
	}

	// Split URL by "/"
	pathParts := strings.Split(u.Path, "/")

	// Grab the last words = name of the question
	var lastPart string
	for i := len(pathParts) - 1; i >= 0; i-- {
		if pathParts[i] != "" {
			lastPart = pathParts[i]
			break
		}
	}
	return lastPart, nil
}

func addQuestionToJsonFile(questionName string, urlStr string) error {
	questions, err := jsonutils.OpenJsonFile()
	if err != nil {
		return err
	}
	newQuestion := jsonutils.Question{
		ID:     "005",
		Name:   questionName,
		URL:    urlStr,
		Solved: 0,
	}
	questions = append(questions, newQuestion)

	// update the json
	updatedJSON, err := json.MarshalIndent(questions, "", "    ")
	if err != nil {
		fmt.Println("marshal error:", err)
		return err
	}
	fmt.Println("Successfully added a new question")

	if err := os.WriteFile("cmd/questions.json", updatedJSON, 0644); err != nil {
		fmt.Println("write file error:", err)
		return err
	}

	return nil
}
