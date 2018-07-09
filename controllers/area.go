package controllers

import (
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	"LoanHouse/models"
)

type AreaController struct {
	beego.Controller
}

//
type Area struct {
	errno  int
	errmsg string
}

func (this *AreaController) RetData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *AreaController) GetArea() {
	beego.Info(" connect success ")

	 resp := make(map[string]interface{})
	defer c.RetData(resp)
	//从session拿数据

	//从mysql中拿到erea数据
	o := orm.NewOrm()
	var areas []models.Area

	num,error := o.QueryTable("area").All(&areas)
	if error != nil {
		beego.Error(" 查询数据错误 ")
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	if(num == 0){
		beego.Error(" 查询成功，数据为空 ")
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	}

	resp["errno"] = 0
	resp["errmsg"] = "ok"
	resp["data"] = areas
}
