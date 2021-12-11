package main

import (
	"StarClub/util"
	"log"
)

//测试发邮件
func main() {
	em := "www.Co15770778807@126.com"
	flag := util.CheckEmail(em)
	vcode, time, err := util.EmailVerify(em)
	log.Println(flag, vcode, time, err)
}
