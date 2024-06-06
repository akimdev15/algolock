/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package open

import (
	"context"
	"fmt"
	"github.com/akimdev15/leetcode/sql"
	"github.com/akimdev15/leetcode/utils"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var questionNumber string

// OpenCmd represents the open command
var OpenCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a question when provided with the question number",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if questionNumber == "" {
			fmt.Println("Please provide a valid question number")
			os.Exit(1)
		}
		_, err := strconv.Atoi(questionNumber)
		if err != nil {
			fmt.Println("Please provide a valid question number")
			os.Exit(1)
		}

		openQuestion()
	},
}

func openQuestion() {
	dbStruct, err := sql.InitDB()
	if err != nil {
		fmt.Println("Error initializing database. Error: ", err)
		os.Exit(1)
	}
	defer dbStruct.DB.Close()

	ctx := context.Background()

	url, err := dbStruct.Queries.GetQuestionUrlByQuestionNum(ctx, questionNumber)
	if err != nil {
		fmt.Println("Error fetching questions. Error: ", err)
		os.Exit(1)
	}

	err = utils.OpenURL(url)
	if err != nil {
		fmt.Println("Error opening URL. Error: ", err)
		return
	}
}

func init() {
	OpenCmd.Flags().StringVarP(&questionNumber, "qNum", "q", "", "Specify the question number")
}
