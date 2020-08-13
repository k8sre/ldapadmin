package controllers

import (
	"github.com/astaxie/beego"
	"github.com/k8sre/ldapadmin/models"
	"net/url"
	"fmt"
)

type MailController struct {
	BaseController
}

func (c *MailController)Get(){
	c.Data["forgetPass"] = "active"
	c.Data["error_display"] = "none"
	c.TplName = "forgetPasswd.html"
}

func (c *MailController)Post(){
	c.TplName = "forgetPasswd.html"
	u := user{}
	if err := c.ParseForm(&u); err != nil {
		beego.Error(err)
		c.ServeError(err.Error())
		return
	}
	username := u.Name
	passwd,mail,err := models.ClientInstance.GetMailAndPasswd(username)
	if err !=nil{
		c.ServeError(err.Error())
		return
	}
	secret := models.Encode(username,passwd)
	param := url.Values{}
	param.Add("secret",secret)
	param.Add("username",username)
	host := beego.AppConfig.String("host")
	url := host + "/set_passwd?" + param.Encode()

	content := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
	<title>index</title>
	<!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
	<link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

	<!-- 可选的 Bootstrap 主题文件（一般不用引入） -->
	<link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">

	<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
	<script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
	</head>
	<body>
	<div class="container">
	<div class="jumbotron">
	<h1>Hello, world!</h1>
	<p>You received this mail because you are changing you password. 
       Click this blue link to set you new password </p>
	<p><a class="btn btn-primary btn-lg" href="%v" role="button">Set password</a></p>
	</div>
	</div>

	</body>
	</html>
    `,url)
	err =models.Mail(content,mail)
	if err !=nil{
		beego.Error(err)
		c.ServeError(err.Error())
		return
	}
	c.ServeAlert("mail had send to "+mail)

}