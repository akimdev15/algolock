/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/akimdev15/leetcode/cmd/add"
	"github.com/akimdev15/leetcode/cmd/get"
	"github.com/akimdev15/leetcode/cmd/pull"
	"github.com/akimdev15/leetcode/cmd/random"
	"github.com/akimdev15/leetcode/cmd/review"
	"github.com/akimdev15/leetcode/cmd/sample"
	"github.com/akimdev15/leetcode/cmd/setup"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "leetcode",
	Short: "",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("root cmd")
	//},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// rootCmd.AddCommand()
	rootCmd.AddCommand(random.RandomCmd)
	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(review.ReviewCmd)
	rootCmd.AddCommand(get.GetCmd)
	rootCmd.AddCommand(setup.Setup)
	rootCmd.AddCommand(pull.PullCmd)
	rootCmd.AddCommand(sample.SampleCmd)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
