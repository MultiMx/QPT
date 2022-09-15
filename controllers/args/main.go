package args

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func Run() {
	if len(os.Args) < 3 {
		log.Fatalln("参数缺失")
	}
	switch os.Args[1] {
	case "install":
		installers()
	case "update":
		updater()
	default:
		log.Fatalf("命令 %s 不存在\n", os.Args[1])
	}
}
