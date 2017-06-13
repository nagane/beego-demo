package routers

import (
	"beego-demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.PostController{}, "*:Index")
	beego.AutoRouter(&controllers.PostController{})
}
