/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package review

import (
	"fmt"
	"github.com/akimdev15/leetcode/leetcode"
	"github.com/akimdev15/leetcode/utils"
	"github.com/spf13/cobra"
)

const QuestionUrl = "https://leetcode.com/problems/%s/"

// ReviewCmd represents the review command
var ReviewCmd = &cobra.Command{
	Use:   "review",
	Short: "review recently solved questions",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
			return
		}

		openRandomReviewQuestion(args[0])
	},
}

func init() {
}

func openRandomReviewQuestion(username string) {
	if username == "" {
		fmt.Println("username is empty")
		return
	}

	questions, err := leetcode.GetRecentSubmissions(5)
	if err != nil {
		fmt.Println("get recently solved questions err:", err)
		return
	}

	randomIdx, err := utils.PickRandomIdx(len(questions))
	if err != nil {
		fmt.Println("get random idx err:", err)
		return
	}

	questionUrl := fmt.Sprintf(QuestionUrl, questions[randomIdx].TitleSlug)

	err = utils.OpenURL(questionUrl)
	if err != nil {
		fmt.Println("open url err:", err)
		return
	}
}
