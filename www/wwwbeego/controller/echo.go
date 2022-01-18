// Package controller
// Copyright (C) Vito
// By Vito on 2022/1/18 14:46

package controller

import (
	"fmt"
	"github.com/astaxie/beego"
)

type EchoResult struct {
	Code int32 `json:"code"`
	Info string `json:"info"`
	Data string `json:"data"`
}

type EchoController struct {
	beego.Controller
}

func (c *EchoController) Get() {
	name := c.Input().Get("name")
	out := c.Input().Get("out")
	echoResult := &EchoResult{Code: 200, Info: "ok", Data: fmt.Sprintf("Pong message: Hello %s", name)}

	switch out {
	case "xml":
		c.Data["xml"] = echoResult
		c.ServeXML()
	case "yml", "yaml":
		c.Data["yaml"] = echoResult
		c.ServeYAML()
	case "json":
		c.Data["json"] = echoResult
		c.ServeJSON()
	default:
		c.Data["json"] = echoResult
		c.ServeJSON()
	}
}