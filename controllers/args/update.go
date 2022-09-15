package args

import (
	"github.com/MultiMx/QPT/controllers/update/golang"
	"os"
)

func updater() {
	switch os.Args[2] {
	case "go":
		golang.MakeUpdate()
	}
}
