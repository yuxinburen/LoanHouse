package controllers

import (
	"github.com/astaxie/beego"
	"LoanHouse/models"
)

type HouseIndexController struct {
	beego.Controller
}

func (c *HouseIndexController) RetData(resp map[string]interface{}){
	c.Data["json"] = resp
	c.ServeJSON()
}

//路由方法
func (c *HouseIndexController) GetHouseIndex() {
	resp := make(map[string]interface{})
	c.RetData(resp)

	resp["errno"] = models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)

}
