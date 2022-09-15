package main

import (
	"fmt"
	"github.com/MultiMx/QPT/controllers/install/golang"
)

func main() {
	//args.Run()
	v, e := golang.GetLatestVersion()
	if e != nil {
		panic(e)
	}
	vl, e := golang.GetLocalVersion()
	if e != nil {
		panic(e)
	}
	if vl == v {
		fmt.Println(true)
	}
	fmt.Println(vl, v)
}
