package controllers

import (
	"github.com/astaxie/beego"
	"LoanHouse/models"
	"encoding/json"
	"github.com/astaxie/beego/orm"
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


//登陆
func (this *SesssionController) Login(){
	//得到用户数据
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	json.Unmarshal(this.Ctx.Input.RequestBody,&resp)

	beego.Info("====name:",resp["mobile"],"=======password:",resp["password"])

	//合法性检测
	if resp["mobile"] == nil || resp["password"] == nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		beego.Error(" login error ")
		return
	}

	//与数据库匹配判断账号密码是否正确
	o := orm.NewOrm()
	user := models.User{Name:resp["mobile"].(string)}
	qs := o.QueryTable("user")
	err := qs.Filter("mobile",resp["mobile"]).One(&user)
	if err != nil{
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	if user.Password_hash!= resp["password"]{
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}

	//设置session
	this.SetSession("name",resp["mobile"])
	this.SetSession("mobile",resp["mobile"])
	this.SetSession("user_id",user.Id)

	//返回json数据给前端
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = user

}
