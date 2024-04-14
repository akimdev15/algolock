/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package random

import (
	"context"
	"fmt"
	"github.com/akimdev15/algolock/sql"
	"github.com/akimdev15/algolock/utils"
	"github.com/spf13/cobra"
)

// RandomCmd represents the random command
var RandomCmd = &cobra.Command{
	Use:   "random",
	Short: "picks a random question from the questions bank",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		openRandomQuestion()
	},
}

func init() {
}

func openRandomQuestion() {
	dbStruct, err := sql.InitDB()
	if err != nil {
		fmt.Println("Error initializing database. Error: ", err)
		return
	}
	defer dbStruct.DB.Close()

	ctx := context.Background()
	// Get question name from the url
	questionURL, err := dbStruct.Queries.GetRandomQuestionURL(ctx)
	if err != nil {
		fmt.Println("Error getting question URL. Error: ", err)
		return
	}

	// TODO - open the question on a link
	err = utils.OpenURL(questionURL)
	if err != nil {
		fmt.Println("error opening URL")
	}

}
