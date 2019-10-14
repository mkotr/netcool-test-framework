package routers

import (
	"git02.ae.sda.corp.telstra.com/scm/wian/netcool-test-automation/nc-test-webapp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/probes", &controllers.ProbeController{})

	//api
	beego.Router("/api/probes", &controllers.ProbeController{}, "get:GetProbes")
}
