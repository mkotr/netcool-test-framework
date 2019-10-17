package routers

import (
	"github.com/mkotr/netcool-test-automation/nc-test-webapp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/probes", &controllers.ProbeController{})

	//api
	beego.Router("/api/probes", &controllers.ProbeController{}, "get:GetProbes")
	beego.Router("/api/fileParser", &controllers.UploadController{}, "post:ParseFile")
}
