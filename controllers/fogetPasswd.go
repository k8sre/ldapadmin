package controllers

import (
	"github.com/astaxie/beego"
	"github.com/k8sre/ldapadmin/models"
)

type PasswdController struct {
	BaseController
}

func (c *PasswdController)Prepare() {
	c.TplName = "info.html"
	username := c.GetString("username")
	if username == ""{
		c.ServeError("username is empty in url path.")
		return
	}
	secret  := c.GetString("secret")
	if secret == ""{
		c.ServeAlert("secret is empty in url path.")
		return
	}
	models.ClientInstance.Connet()
	passwd,_,err := models.ClientInstance.GetMailAndPasswd(username)
	if err !=nil{
		c.ServeError(err.Error())
		return
	}
	if passwd == ""{
		c.ServeError("passwd is empty")
		return
	}
	beego.Info(models.Encode(username,passwd),username,passwd)
	beego.Info(secret)
	if models.Encode(username,passwd) != secret{
		c.ServeError("secret is out of data,plase get a new one")
		return
	}
}

func (c *PasswdController)Get(){
	c.Data["error_display"] = "none"
	c.Data["forgetPass"] = "active"
	c.TplName = "setPasswd.html"
}

func (c *PasswdController)Post(){
	username := c.GetString("username")
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
	err := models.ClientInstance.SetPasswd(username,u.NewPasswd)
	if err !=nil{
		c.ServeError(err.Error())
		return
	}
	c.ServeAlert("finished!")

}