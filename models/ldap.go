package models

import (
	"gopkg.in/ldap.v2"
	"net"
	"github.com/astaxie/beego"
	"fmt"
	"errors"
	"encoding/json"
)
type Client struct {
	Address string
	Conn *ldap.Conn
	adminDN string
	adminPasswd string
	BaseDn string
}
var (
	ClientInstance *Client


)
func init() {
	ClientInstance = &Client{
		Address: beego.AppConfig.String("address"),
		adminDN: beego.AppConfig.String("adminDN"),
		adminPasswd: beego.AppConfig.String("adminPasswd"),
		BaseDn:beego.AppConfig.String("baseDn"),
	}
	ClientInstance.Connet()


}

func (client *Client)Connet()error{
	beego.Info("connect to ",client.Address)
	c,err := net.Dial("tcp",client.Address)
	if err !=nil{
		panic(err)
	}
	ClientInstance.Conn = ldap.NewConn(c, false)
	ClientInstance.Conn.Start()

	//admin login
	beego.Info("admin login ....")
	err = ClientInstance.Conn.Bind(client.adminDN,client.adminPasswd)
	if err !=nil{
		beego.Error("admin login error:",err)
		return err
	}
	return nil
}


func (client *Client)VerfiyAuth(username,passwd string) bool{
	username = "cn="+username
	err := client.Conn.Bind(username,passwd)
	if err !=nil{
		beego.Error("error:",err)
		return false
	}
	return true
}
func (client *Client)SearchByCN(username string)(*ldap.SearchResult,error){
	sreq := &ldap.SearchRequest{
		client.BaseDn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(cn=%v)",username),
		[]string{"dn","userpassword","mail"},
		nil,
	}
	sr,err := client.Conn.Search(sreq)
	if err !=nil{
		return nil,err
	}
	return sr,nil
}



func (client *Client)GetMailAndPasswd(username string)(passwd string,mail string,err error){
	//search user
	rs,err := client.SearchByCN(username)
	if err !=nil{
		beego.Error(err)
		return "","",err
	}
	if len(rs.Entries) < 1 {
		beego.Error("User does not exist")
		return "","",errors.New("User does not exist")
	}

	if len(rs.Entries) > 1 {
		beego.Error("Too many entries returned")
		return "","",errors.New("Too many entries returned")
	}
	jrs,_ := json.Marshal(rs)
	beego.Info(string(jrs))
	attrs := rs.Entries[0].Attributes
	for _,v := range attrs {
		if v.Name == "userPassword"{
			passwd = v.Values[0]
		}
		if v.Name == "mail"{
			mail = v.Values[0]
		}
	}
	return passwd,mail,nil
}

func (client *Client)ModifyPasswd(user,oldPasswd,newPasswd string)error{
	l := client.Conn



	//search user
	rs,err := client.SearchByCN(user)
	if err !=nil{
		beego.Error(err)
		return err
	}
	if len(rs.Entries) < 1 {
		beego.Error("User does not exist")
		return errors.New("User does not exist")
	}

	if len(rs.Entries) > 1 {
		beego.Error("Too many entries returned")
		return errors.New("Too many entries returned")
	}
	userdn := rs.Entries[0].DN



	//change passwd
	preq := &ldap.PasswordModifyRequest{
		UserIdentity:userdn,
		OldPassword:oldPasswd,
		NewPassword:newPasswd,
	}
	_,err = l.PasswordModify(preq)
	if err !=nil{
		beego.Error("modify passwd error:",err)
		return errors.New("modify passwd error:"+err.Error())
	}
	return nil
}

func (client *Client)Close(){
	client.Conn.Close()
}

func (client *Client)SetPasswd(user,newPasswd string)error{
	l := client.Conn

	//search user
	rs,err := client.SearchByCN(user)
	if err !=nil{
		beego.Error(err)
		return err
	}
	if len(rs.Entries) < 1 {
		beego.Error("User does not exist")
		return errors.New("User does not exist")
	}

	if len(rs.Entries) > 1 {
		beego.Error("Too many entries returned")
		return errors.New("Too many entries returned")
	}
	userdn := rs.Entries[0].DN

	p := ldap.PartialAttribute{
		Type: "userpassword",
		Vals:[]string{EntrySha1(newPasswd),},
	}
	//set passwd
	mreq := &ldap.ModifyRequest{
		DN:userdn,
		ReplaceAttributes:[]ldap.PartialAttribute{p},
	}

	err = l.Modify(mreq)
	if err !=nil{
		beego.Error(err)
		return err
	}
	return nil
}
