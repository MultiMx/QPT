package args

import (
	"github.com/MultiMx/QPT/controllers/install/golang"
	"os"
)

func installers() {
	switch os.Args[2] {
	case "go":
		golang.MakeInstall()
	}
}
