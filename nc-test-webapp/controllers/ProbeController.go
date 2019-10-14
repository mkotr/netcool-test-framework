package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type ProbeController struct {
	beego.Controller
}

type Probe struct{
	Name string `json:"name"`
	Desc string `json:"desc"`
	Hostname string `json:"hostname"`
	Port string `json:"port"`
}

func (p *ProbeController) Get() {
	p.TplName = "probes/index.html"
}

func (p *ProbeController) GetProbes() {
	probe := Probe{
		Name: "name",
		Desc: "desc",
		Hostname: "hostname",
		Port: "port",
	}
	fmt.Println("hello")
	p.Data["json"] = &probe
	p.ServeJSON()
}