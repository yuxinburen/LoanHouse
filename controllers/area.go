package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"LoanHouse/models"
	//_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/cache"
	"time"
)

type AreaController struct {
	beego.Controller
}

//
type Area struct {
	errno  int
	errmsg string
}

func (this *AreaController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *AreaController) GetArea() {
	beego.Info(" connect success ")

	resp := make(map[string]interface{})
	defer c.RetData(resp)
	//从缓存中拿数据。恒定不变的内容，静态不会变化的内容，一般会放在redis缓存中
	bm, err := cache.NewCache("redis", `{"key":"collectionName","conn":":6379","dbNum":"0"}`)
	errCache := bm.Put("aaa", "bbss", time.Second*3600)
	if errCache != nil {
		beego.Error("  redis cache wrong ")
		return
	}
	beego.Info("cache_conn.aa=", bm.Get("aaa"))

	//从mysql中拿到erea数据
	o := orm.NewOrm()
	var areas []models.Area

	num, err := o.QueryTable("area").All(&areas)
	if err != nil {
		beego.Error(" 查询数据错误 ")
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	if (num == 0) {
		beego.Error(" 查询成功，数据为空 ")
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	}

	resp["errno"] = 0
	resp["errmsg"] = "ok"
	resp["data"] = areas
}
