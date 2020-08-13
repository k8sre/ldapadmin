package models

import (
	"time"
	"strconv"
	"crypto/md5"
	"fmt"
	"bytes"
	"encoding/base64"
	"crypto/sha512"
	"crypto/sha1"
)


func Encode(username,passwd string)string{
	encodeTime := strconv.FormatInt(time.Now().Unix()/7200*7200,10)
	b := bytes.NewBufferString(username)
	b.WriteString(passwd)
	b.WriteString(encodeTime)
	data := b.Bytes()
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}

func EntrySha512(passwd string)string{
	h := sha512.New()
	h.Write([]byte(passwd)) // 需要加密的字符串为
	//fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	p := string(base64.StdEncoding.EncodeToString(h.Sum(nil)))
	return "{SHA512}"+p
}

func EntrySha1(passwd string)string{
	salt := []byte("smart")
	h := sha1.New()
	h.Write([]byte(passwd))
	h.Write(salt)
	bs := h.Sum(nil)
	a := append(bs,salt...)
	bsa := base64.StdEncoding.EncodeToString(a)
	return "{SSHA}"+bsa
}