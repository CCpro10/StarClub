package model

import (
	"github.com/jinzhu/gorm"
)

type UserInfo struct {
	gorm.Model
	Email 		string `form:"email"json:"email"`	   //电子邮箱
	Studentid   string  `form:"studentid"json:"studentid"` //真实学号
	Username    string `form:"username"json:"username"`   //用户的呢名
	Password    string `form:"password"json:"password"`
	Name        string `form:"name"json:"name"`   //用户真实姓名
}

type UserRegister struct {
	Email 		string `form:"email"json:"email"`	   //电子邮箱
	Studentid   string  `form:"studentid"json:"studentid"` //真实学号
	Password    string `form:"password"json:"password"`   //密码
	Vcode       string `form:"vcode"json:"vcode"`
}




