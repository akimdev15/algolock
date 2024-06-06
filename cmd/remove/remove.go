/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package remove

import (
	"context"
	"fmt"
	"github.com/akimdev15/leetcode/sql"
	"github.com/spf13/cobra"
	"strings"
)

const BaseURL = "https://leetcode.com/problems/"

// RemoveCmd represents the remove command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a question from solved questions",
	Long:  `Provide either a url or the ID`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please provide an ID or an url")
			return
		} else if len(args) == 1 {
			isValidLeetcodeUrl(args[0])
		}
	},
}

func init() {
}

// removeQuestion - removes a question from the database
func removeQuestion() {
	dbStruct, err := sql.InitDB()
	if err != nil {
		fmt.Println("Error initializing database. Error: ", err)
		return
	}
	defer dbStruct.DB.Close()

	ctx := context.Background()

	dbStruct.Queries.DeleteQuestionById(ctx, "")
}

func isValidLeetcodeUrl(urlStr string) bool {
	if strings.HasPrefix(urlStr, BaseURL) {
		fmt.Println("HERE")
		newUrl := strings.Replace(urlStr, "/description/", "/", 1)
		fmt.Println(newUrl)
	}

	// 1. Check if the base url matches

	return true
}
