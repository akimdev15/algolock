package get

import (
	"context"
	"fmt"
	"github.com/akimdev15/leetcode/internal/database"
	"github.com/akimdev15/leetcode/query"
	"github.com/akimdev15/leetcode/sql"
	"github.com/akimdev15/leetcode/utils"
	"github.com/spf13/cobra"
	"os"
)

var (
	questionCount int
	all           bool
)

// GetCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get questions",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if all {
			questions := getAllQuestions()
			utils.PrintQuestions(questions)
		} else if questionCount > 0 {
			fmt.Printf("Getting recent %d questions\n", questionCount)
			questions := getRecentQuestions(questionCount)
			utils.PrintQuestions(questions)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	GetCmd.Flags().IntVarP(&questionCount, "recent", "r", 0, "Specify number of recently solved questions to fetch")
	GetCmd.Flags().BoolVarP(&all, "all", "a", false, "Get all saved questions")
}

func getAllQuestions() []database.Question {
	fmt.Println("Getting all saved questions...")
	dbStruct, err := sql.InitDB()
	if err != nil {
		fmt.Println("Error initializing database. Error: ", err)
		return nil
	}
	defer dbStruct.DB.Close()

	ctx := context.Background()

	questions, err := dbStruct.Queries.GetAllQuestions(ctx)
	if err != nil {
		fmt.Println("Error fetching questions. Error: ", err)
		return nil
	}

	return questions
}

func getRecentQuestions(questionCount int) []database.Question {
	dbStruct, err := sql.InitDB()
	if err != nil {
		fmt.Println("Error initializing database. Error: ", err)
		return nil
	}
	defer dbStruct.DB.Close()

	ctx := context.Background()

	recentQuestions := make([]database.Question, 0)

	// Get n questions solved from leetcode
	submissions, err := query.GetRecentSubmissions(questionCount)

	// Create a map for quick lookup of submission IDs
	submissionMap := make(map[string]struct{}, len(submissions))
	for _, submission := range submissions {
		submissionMap[submission.Title] = struct{}{}
	}

	allQuestions, err := dbStruct.Queries.GetAllQuestions(ctx)
	if err != nil {
		fmt.Println("Error fetching latest questions. Error: ", err)
		os.Exit(1)
	}

	for _, q := range allQuestions {
		_, exists := submissionMap[q.Name]
		if exists && len(recentQuestions) < questionCount {
			recentQuestions = append(recentQuestions, q)
		}
	}

	return recentQuestions
}
