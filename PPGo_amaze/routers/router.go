package routers

import (
	"PPGo_amaze/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 默认登录
	beego.Router("/", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")
	beego.Router("/no_auth", &controllers.LoginController{}, "*:NoAuth")
	beego.Router("/home", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")
	// 深指
	beego.Router("/stockindex399", &controllers.Stockindex{}, "*:Index399")
	// 上指
	beego.Router("/stockindex000001", &controllers.Stockindex{}, "*:Index000001")
	// 创业板
	beego.Router("/stockindex300", &controllers.Stockindex{}, "*:Index300")
	// 个股
	beego.Router("/stockindex", &controllers.Stockindex{}, "*:Index")

}
