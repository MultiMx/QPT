package args

import (
	"github.com/MultiMx/QPT/controllers/update/golang"
	"github.com/MultiMx/QPT/controllers/update/qpt"
	"os"
)

func updater() {
	switch os.Args[2] {
	case "go":
		golang.MakeUpdate()
	case "qpt":
		qpt.MakeUpdate()
	}
}
