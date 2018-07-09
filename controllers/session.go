package controllers

import (
	"github.com/astaxie/beego"
	"LoanHouse/models"
)

type SesssionController struct {
	beego.Controller
}

func (this *SesssionController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *SesssionController) GetSessionData() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	user := models.User{}
	//user.Name = "hongweiyu"

	resp["errno"] = models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)

	name := c.GetSession("name")
	beego.Info("session:", name)
	if name != nil {
		user.Name = name.(string)
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"] = user
	}
}

//删除session数据
func (this *SesssionController) DeleteSessionData(){
	resp := make(map[string]interface{})
	defer this.RetData(resp)

	this.DelSession("name")
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}
