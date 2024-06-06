/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package review

import (
	"fmt"
	"github.com/akimdev15/leetcode/query"
	"github.com/akimdev15/leetcode/utils"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"os"
)

// ReviewCmd represents the review command
var ReviewCmd = &cobra.Command{
	Use:   "review",
	Short: "review recently solved questions",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		openRandomReviewQuestion()
	},
}

func init() {
}

func openRandomReviewQuestion() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	username := os.Getenv("username")
	if username == "" {
		fmt.Println("username is empty")
		return
	}

	fmt.Println("username is", username)
	submissions, err := query.GetRecentSubmissions(5)
	if err != nil {
		fmt.Println("get recently solved submissions err:", err)
		return
	}

	randomIdx, err := utils.PickRandomIdx(len(submissions))
	if err != nil {
		fmt.Println("get random idx err:", err)
		return
	}

	questionUrl := utils.ConstructLeetcodeURL(submissions[randomIdx].TitleSlug)

	err = utils.OpenURL(questionUrl)
	if err != nil {
		fmt.Println("open url err:", err)
		return
	}
}
