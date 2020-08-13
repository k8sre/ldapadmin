package main

import (
	_ "github.com/k8sre/ldapadmin/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

