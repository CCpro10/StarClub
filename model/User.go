package model

import (
	"github.com/jinzhu/gorm"
)

//用户信息
type UserInfo struct {
	gorm.Model
	//头像
	//关注的社团
	//收藏的活动
	Signature string `form:"signature"json:"signature"` //个性签名
	Email     string `form:"email"json:"email"`         //电子邮箱
	StudentId string `form:"studentid"json:"studentid"` //真实学号
	Username  string `form:"username"json:"username"`   //用户的呢名
	Password  string `form:"password"json:"password"`   //密码
	Name      string `form:"name"json:"name"`           //用户真实姓名
}

//用户注册信息
type UserRegister struct {
	Email      string `form:"email"json:"email"`         //电子邮箱
	StudentId  string `form:"studentid"json:"studentid"` //真实学号
	Password   string `form:"password"json:"password"`   //密码
	VerifyCode string `form:"vcode"json:"vcode"`         //验证码
}

//用户登录信息
type UserLogin struct {
	EmailOrId   string `form:"emailorid"json:"emailorid"`  //电子邮箱或学号
	Password   string `form:"password"json:"password"`   //密码
}




