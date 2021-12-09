package main

import (
	"StarClub/model"
	"log"
)

func main() {
	em := "aab.Co15770778807@126.com"
	flag := model.CheckEmail(em)
	vcode, time, err := model.EmailVerify(em)
	log.Println(flag, vcode, time, err)
}
