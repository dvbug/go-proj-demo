// Package go_proj_demo
// Copyright (C) Vito
// By Vito on 2022/1/18 11:55

package main

import (
	"dvbug.com/scripts"
	"dvbug.com/wwwbeego"
)

func main() {
	scripts.Info("This is a info log")
	scripts.Main()
	//scripts.PcMain()
	wwwbeego.Main()
}
