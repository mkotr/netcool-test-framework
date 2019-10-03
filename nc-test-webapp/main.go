package main

import (
	_ "git02.ae.sda.corp.telstra.com/scm/wian/netcool-test-automation/nc-test-webapp/routers"
	"github.com/astaxie/beego"
)

func main() {
	//beego.BConfig.WebConfig.StaticDir["/static"] = "nc-test-web/dist"
	beego.SetStaticPath("/_nuxt", "web/dist/_nuxt")
	//beego.SetStaticPath("/js", "web/dist/_nuxt")
	//beego.SetStaticPath("/img", "nc-test-web/dist/_nuxt")

	beego.Run()
}
