package routers

import (
	"BeegoProject123/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/register", &controllers.MainController{})
}

func INIT (){
	beego.Router("/",&controllers.From{})
}
