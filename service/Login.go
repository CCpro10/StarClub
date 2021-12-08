package service

import (
	"StarClub/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

//登录,接收学号/邮箱 和密码,如果正确的话返回一个tokenstring
func Login(c *gin.Context) {

	var requestUser model.UserLogin
	err := c.ShouldBind(&requestUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "参数绑定失败",
		})
		return
	}
	fmt.Println(requestUser)

	// 校验密码格式是否正确
	if len(requestUser.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "密码长度不能小于六位"})
		return
	}

	var user =model.UserInfo{}
	//判断是学号还是邮箱
	studentid, err := strconv.Atoi(requestUser.EmailOrId)
	//如果输入的是学号
	if err==nil{
		if len(string(studentid)) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "学号格式错误"})
			return
		}
		model.DB.Where("studdent_id = ?", requestUser.EmailOrId).First(&user)
		//账号是否存在
		if user.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "此用户不存在，请先注册或检查是否输入有误"})
			return
		}
		//判断密码是否正确
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestUser.Password)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "账号或密码输入错误"})
			return
		}
		// 生成Token
		tokenString,err:= model.GenToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "用户token生成失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "登录成功",
			"data": gin.H{"token": tokenString},
		})
		return

		//输入的是邮箱
	}else {
		//先用正则表达式判断邮箱格式有没有错误
		ok:=model.CheckEmail(requestUser.EmailOrId)
		if !ok{
			c.JSON(http.StatusBadRequest, gin.H{"msg": "邮箱格式错误"})
			return
		}
		model.DB.Where("email = ?", requestUser.EmailOrId).First(&user)
		//查询邮箱是否存在
		if user.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "此用户不存在，请先注册或检查是否输入有误"})
			return
		}
		//判断密码是否正确
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestUser.Password)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "账号或密码输入错误"})
			return
		}
		// 生成Token
		tokenString,err:= model.GenToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "用户token生成失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "登录成功",
			"data": gin.H{"token": tokenString},
		})
		return
	}

}
