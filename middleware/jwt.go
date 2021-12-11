package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

const TokenExpireDuration = time.Hour * 24 * 30

var MySecret = []byte("天王盖地虎")

// MyClaims自定义声明结构体并内嵌jwt.StandardClaims
// 我们这里需要额外记录一个userid字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserId uint //表示用户ID
	jwt.StandardClaims
}

//生成令牌,传入userid,生成JWTstring和err
func GenToken(UserId uint) (string, error) {
	// 创建一个我们自己的声明/请求
	c := MyClaims{
		UserId, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "my-project",                               // 签发人
			Subject:   "user token",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的秘钥签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析tokenstring，返回一个包含信息的用户声明
func ParseToken(tokenString string) (*MyClaims, error) {
	// 通过(tokenstring,请求结构,返回秘钥的一个回调函数)这三个参数,返回一个token结构体,token包含了请求结构
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}

	// 校验token,token有效则返回claims请求，nil
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	//token无效，返回错误
	return nil, errors.New("invalid token")
}

//基于JWT的认证中间件
func JWTAuthMiddleware(c *gin.Context) {
	{
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在URL的query中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "请求头中auth为空,请先登录"})
			c.Abort()
			return
		}

		// 按空格分割,在第一个空格后分割成两部分
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "请求头中auth格式有误,请重新登录"})
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "此Token无效或已过期,请重新登录"})
			c.Abort()
			return
		}

		// 将当前请求的userid信息保存到请求的上下文c上
		//后续的处理函数可以用过c.Get("userid")来获取当前请求的用户信息
		c.Set("userid", mc.UserId)
		c.Next()
	}
}
