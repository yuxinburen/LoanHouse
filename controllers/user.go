package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"LoanHouse/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

//注册方法
func (c *UserController) Reg() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	json.Unmarshal(c.Ctx.Input.RequestBody, &resp)

	beego.Info(`resp["mobile"]=`, resp["mobile"])
	beego.Info(`resp["password"]=`, resp["password"])
	beego.Info(`resp["sms_code"]=`, resp["sms_code"])

	o := orm.NewOrm()
	user := models.User{}
	user.Password_hash = resp["password"].(string)
	user.Name = resp["mobile"].(string)
	user.Mobile = resp["mobile"].(string)
	_, err := o.Insert(&user)
	if (err != nil) {
		beego.Error(err)
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = "注册失败"
		return
	}
	beego.Info("注册成功")
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	c.SetSession("name",user.Name)
}
