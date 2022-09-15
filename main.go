package main

import (
	"github.com/MultiMx/QPT/controllers/args"
	"github.com/MultiMx/QPT/modules"
)

func init() {
	modules.Init()
}

func main() {
	args.Run()
}
