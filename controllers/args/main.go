package args

import (
	"fmt"
	"github.com/MultiMx/QPT/global"
	log "github.com/sirupsen/logrus"
	"os"
)

func Run() {
	if len(os.Args) == 1 {
		log.Fatalln("参数缺失")
	}
	var length int
	var f func()
	switch os.Args[1] {
	case "install":
		length = 3
		f = installers
	case "update":
		length = 3
		f = updater
	case "version":
		f = func() {
			fmt.Println(global.VERSION)
		}
	case "config":
		length = 4
		f = configure
	default:
		log.Fatalf("命令 %s 不存在\n", os.Args[1])
	}

	if len(os.Args) < length {
		log.Fatalln("参数不足")
	}
	f()
}
