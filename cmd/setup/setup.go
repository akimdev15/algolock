/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package setup

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// Setup represents the setup command
var Setup = &cobra.Command{
	Use:   "setup",
	Short: "Initialize the project",
	Long:  `Setup command will setup the dev env to get everything necessary to run the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get user's leetcode id
		leetcodeUsername := getLeetcodeID()
		if leetcodeUsername == "" {
			fmt.Println("please input leetcodeUsername")
			return
		}

		createEnvFile(leetcodeUsername)
	},
}

func init() {
}

func getLeetcodeID() string {
	fmt.Print("Enter your leetcode id: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	leetcodeID := strings.TrimSpace(scanner.Text())
	return leetcodeID
}

// createEnvFile - creates an env file and saves username
func createEnvFile(username string) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		file, err := os.Create(".env")
		if err != nil {
			fmt.Println("Error creating .env file")
			os.Exit(1)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("%s=%s\n", "username", username))
		if err != nil {
			fmt.Println("Error writing to .env file. Please try to setup again...")
			os.Exit(1)
		}
	} else if err != nil {
		fmt.Println("Error reading .env file. Please try to setup again...")
		os.Exit(1)
	} else {
		fmt.Println("You already have a .env file. If you would like to update the username, please edit the .env file")
	}
}
