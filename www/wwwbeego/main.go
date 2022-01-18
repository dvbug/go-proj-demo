// Package beego
// Copyright (C) Beijing Xiaomi Co., Ltd
// By Vito on 2022/1/18 14:37

package wwwbeego

import (
	"dvbug.com/wwwbeego/controller"
	"fmt"
	"github.com/astaxie/beego"
)

func Main() {
	err := beego.LoadAppConfig("ini", "/home/vito/code/me/workspace-go/go-proj-demo/www/wwwbeego/conf/app.conf")
	if err != nil {
		fmt.Println(fmt.Sprintf("Invalid app.conf: %s", err))
		return
	}
	beego.Router("/echo", &controller.EchoController{})
	beego.Run()
}