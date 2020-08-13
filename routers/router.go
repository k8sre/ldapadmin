package routers

import (
	"github.com/k8sre/ldapadmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/foget_passwd", &controllers.MailController{})
	beego.Router("/set_passwd", &controllers.PasswdController{})

}
