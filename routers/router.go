package routers

import (
	"LoanHouse/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/api/v1.0/areas", &controllers.AreaController{},"get:GetArea")

    beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{},"get:GetHouseIndex;delete:DeleteSessionData")

    beego.Router("/api/v1.0/session", &controllers.SesssionController{},"get:GetSessionData")
    beego.Router("/api/v1.0/users", &controllers.UserController{},"post:Reg")

}
