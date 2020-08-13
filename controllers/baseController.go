package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}


func (c *BaseController)ServeError(err string){
	c.Data["error_display"] = "block"
	c.Data["error_form"] = "alert-danger"
	c.Data["info"] = err
	c.Render()
	c.StopRun()
}

func (c *BaseController)ServeAlert(alert string){
	c.Data["error_display"] = "block"
	c.Data["error_form"] = "alert-success"
	c.Data["info"] = alert
	c.Render()
	c.StopRun()
}
