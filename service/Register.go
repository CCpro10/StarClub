package service

import (
	"StarClub/dao"
	"StarClub/model"
	"StarClub/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

//发送验证码
func SendVerifyCode(c *gin.Context) {
	var user model.UserRegister
	//发送的验证码
	var vcode string
	//发送时间
	var sendtime time.Time
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "邮箱参数绑定失败:" + err.Error()})
	}
	//检查邮箱格式是否合法
	if util.CheckEmail(user.Email) == false {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "邮箱格式不正确,请确保输入有效的邮箱"})
		return
	}

	//通过Email从Redis获取上一次发送的时间,如果时间间隔小于一分钟,则拒绝发送
	ObtainedValue, err := dao.RedisDB.Get(dao.CTX, user.Email).Result()
	//如果value查询成功

	if err == nil {
		//取出value里的发送时间sendtime
		parts := strings.SplitN(ObtainedValue, " ", 2)
		var sendtime time.Time
		_ = json.Unmarshal([]byte(parts[1]), &sendtime)

		//验证码在前一分钟以内发送过,则拒绝发送
		if sendtime.Add(1 * time.Minute).After(time.Now()) {
			c.JSON(http.StatusTooManyRequests, gin.H{"msg": "验证码发送过于频繁"})
			return
		}
	}

	//发送验证码邮件
	if vcode, sendtime, err = util.EmailVerify(user.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "邮箱发送错误,error=" + err.Error()})
		return
	}

	//封装value,value包括了用户的验证码和发送时间
	sendtime2, _ := json.Marshal(sendtime)
	value := vcode + " " + string(sendtime2)

	//通过Redis设置验证码有效时间为十分钟
	dao.RedisDB.Set(dao.CTX, user.Email, value, 10*time.Minute)
	c.JSON(http.StatusOK, gin.H{"msg": "验证码发送成功"})
	return
}

//注册路由
func Register(c *gin.Context) {
	//从请求中把数据取出
	var requestUser model.UserRegister
	err := c.ShouldBind(&requestUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户参数绑定失败:" + err.Error()})
	}

	//验证学号,邮件,密码的结构
	if len(requestUser.StudentId) != 10 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "学号格式错误"})
		return
	}
	if util.CheckEmail(requestUser.Email) == false {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "邮箱格式不正确,请确保输入有效的邮箱"})
		return
	}
	if len(requestUser.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "密码长度不能小于六位"})
		return
	}

	//判断在数据库中邮箱是否存在
	var user model.UserInfo
	//查询并赋值给user
	dao.DB.Where("email=?", requestUser.Email).First(&user)
	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "此邮箱已被注册"})
		return
	}

	//查询Redis,检测验证码是否正确,是否过期
	value, err := dao.RedisDB.Get(dao.CTX, requestUser.Email).Result()

	//如果查不到此邮箱对应验证码
	if err == redis.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "验证码错误或已经过期"})
		return
	} else if err != nil { //查询失败了
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "服务器查询验证码失败"})
		return
	}

	//取出value里的验证码vcode
	parts := strings.SplitN(value, " ", 2)
	vcode := parts[0]

	//通过邮箱查到的验证码和用户输入的验证码不一样
	if vcode != requestUser.VerifyCode {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "验证码错误或已经过期"})
		return
	}

	// 对密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(requestUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "用户密码加密失败"})
		return
	}

	//创建用户实例,存入注册信息
	var userinfo = model.UserInfo{
		Email:     requestUser.Email,
		StudentId: requestUser.StudentId,
		Password:  string(hashPassword),
	}
	err = dao.DB.Create(&userinfo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "注册信息存入数据库失败:" + err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "注册成功,请你重新登录",
			"date": userinfo,
		})
	}
}
