/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package random

import (
	"errors"
	"fmt"
	"github.com/akimdev15/algolock/cmd/jsonutils"
	"github.com/spf13/cobra"
	"math/rand"
	"os/exec"
	"runtime"
	"time"
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
	questions, err := jsonutils.OpenJsonFile()
	if err != nil {
		fmt.Println("error opening questions file")
		return
	}

	pickedQuestion, err := pickRandomQuestion(questions)
	if err != nil {
		fmt.Println("error picking question. error: ", err)
		return
	}

	fmt.Println("Picked Question: ", pickedQuestion)

	// TODO - open the question on a link
	err = openURL(pickedQuestion.URL)
	if err != nil {
		fmt.Println("error opening URL")
	}

}

func pickRandomQuestion(questions []jsonutils.Question) (jsonutils.Question, error) {
	if len(questions) == 0 {
		return jsonutils.Question{}, errors.New("no questions found")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIdx := r.Intn(len(questions))
	fmt.Println("Idx: ", randomIdx)
	return questions[randomIdx], nil
}

func openURL(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}
