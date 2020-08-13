package models
import (
"strings"
"net/smtp"
	"github.com/astaxie/beego"
)

var(
	err error
	user string
	passwd string
	host string
	subject string
	mailtype string
)

func init() {
	user = beego.AppConfig.String("email_user")
	passwd = beego.AppConfig.String("email_passwd")
	host = beego.AppConfig.String("email_host")
	mailtype = beego.AppConfig.String("mail_type")
	subject = beego.AppConfig.String("subject")
}



func sendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	from := "k8sre"
	msg := []byte("To: " + to + "\r\nFrom: " + from + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}


func Mail(content,to string ) (err error)  {
	err = sendMail(user,passwd,host,to,subject,content,mailtype)
	return err
}