package routers

import (
	"LoanHouse/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    //首页获取区域
    beego.Router("/api/v1.0/areas", &controllers.AreaController{},"get:GetArea")

    //首页
    beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{},"get:GetHouseIndex")

    beego.Router("/api/v1.0/session", &controllers.SesssionController{},"get:GetSessionData;delete:DeleteSessionData")

    //注册
    beego.Router("/api/v1.0/users", &controllers.UserController{},"post:Reg")

    //登陆接口
    beego.Router("/api/v1.0/sessions", &controllers.SesssionController{},"post:Login")
	
}
