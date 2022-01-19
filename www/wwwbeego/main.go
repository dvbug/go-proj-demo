// Package beego
// Copyright (C) Vito
// By Vito on 2022/1/18 14:37

package wwwbeego

import (
	"fmt"
	"github.com/astaxie/beego"
	cl "go-proj-demo/www/wwwbeego/controller"
)

func Main() {
	err := beego.LoadAppConfig("ini", "www/wwwbeego/conf/app.conf")
	if err != nil {
		fmt.Println(fmt.Sprintf("Invalid app.conf: %s", err))
		return
	}
	beego.Router("/echo", &cl.EchoController{})
	beego.Run()
}
