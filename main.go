package main

import (
	_ "LoanHouse/routers"
	"github.com/astaxie/beego"
	"strings"
	"net/http"
	"github.com/astaxie/beego/context"
	_ "LoanHouse/models"
)

func main() {

	//1.再main中集成session环境  2.在配置文件中设置session
	beego.BConfig.WebConfig.Session.SessionOn = true

	ignoreStaticPath()
	beego.Run()
}

func ignoreStaticPath()  {
	//透明static
	beego.InsertFilter("/",beego.BeforeRouter,TransparentStatic)
	beego.InsertFilter("/*",beego.BeforeRouter,TransparentStatic)
}
func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ", orpath)
	//如果请求url还有api字段，说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
}
