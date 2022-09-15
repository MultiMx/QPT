package util

import (
	"github.com/briandowns/spinner"
	"os"
	"time"
)

func Spinner(pin string) *spinner.Spinner {
	spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithWriter(os.Stderr))
	spin.Suffix = " " + pin
	return spin
}
