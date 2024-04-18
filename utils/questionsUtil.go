package utils

import (
	"errors"
	"fmt"
	"github.com/akimdev15/leetcode/internal/database"
	"github.com/jedib0t/go-pretty/v6/table"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func PickRandomIdx(sliceLen int) (int, error) {
	if sliceLen == 0 {
		return 0, errors.New("provided length is 0")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIdx := r.Intn(sliceLen)
	return randomIdx, nil
}

func OpenURL(url string) error {
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

// PrintQuestions - prints all the questions using the table library
func PrintQuestions(questions []database.Question) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Difficulty", "Solved", "Confidence"})
	for _, question := range questions {
		t.AppendRow([]interface{}{question.ID, question.Name, question.Difficulty, question.Solved, question.Confidence})
		t.AppendSeparator()
	}
	t.Render()
}
