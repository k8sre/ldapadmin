package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/k8sre/ldapadmin/models"
	"github.com/astaxie/beego"
)
type user struct {
	Name  string `form:"username"`
	OldPasswd   string         `form:"oldPasswd"`
	NewPasswd   string          `form:"newPasswd"`
	VerifyPasswd string         `form:"verifyPasswd"`
}

func init() {
	logs.SetLogger("console")
}

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["changePass"] = "active"
	c.Data["error_display"] = "none"
	c.TplName = "index.html"
}

func (c *MainController)Post(){
	c.TplName = "index.html"
	u := user{}
	if err := c.ParseForm(&u); err != nil {
		beego.Error(err)
		c.ServeError(err.Error())
		return
	}
	if u.NewPasswd != u.VerifyPasswd{
		c.ServeError("new password and verify password are not same")
		return
	}


	if err := models.ClientInstance.ModifyPasswd(u.Name,u.OldPasswd,u.NewPasswd);err ==nil {
		c.ServeAlert("change passwd ok"+"\n")
		return
	}else{
		c.ServeError("change passwd error!\n" + err.Error())
		return
	}
}
