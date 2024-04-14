package utils

import (
	"errors"
	"fmt"
	"math/rand"
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
