package controllers

import (
	"fmt"

	"github.com/mkotr/netcool-test-automation/nc-test-webapp/service"
	"github.com/astaxie/beego"
)

//ProbeController - handler for probe api
type ProbeController struct {
	beego.Controller
	service service.ProbeService
}

//Get - the page for this route.
func (p *ProbeController) Get() {
	p.TplName = "probes/index.html"
}

//GetProbes - retrieves probe information from service.
func (p *ProbeController) GetProbes() {
	probes, err := p.service.GetAll()
	if err != nil {
		fmt.Println("error")
	}
	p.Data["json"] = &probes
	p.ServeJSON()
}
