package util

import (
	"github.com/briandowns/spinner"
	"os"
	"time"
)

func Spinner() *spinner.Spinner {
	return spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithWriter(os.Stderr))
}
