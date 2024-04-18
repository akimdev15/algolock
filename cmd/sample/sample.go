package sample

import (
	"fmt"
	"github.com/akimdev15/leetcode/cmd/jsonutils"
	"github.com/spf13/cobra"
)

// SampleCmd represents the sample command
var SampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		addSampleQuestionsToDatabase()
	},
}

func init() {
}

func addSampleQuestionsToDatabase() {
	questions, err := jsonutils.OpenJsonFile()
	if err != nil {
		fmt.Println("jsonutils.OpenJsonFile() err: ", err)
	}
	for _, question := range questions {
		fmt.Println("Question: ", question.Name)
	}
}
