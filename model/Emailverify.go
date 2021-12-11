package model

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"math/rand"
	"net/smtp"
	"regexp"
	"time"
)

//正则表达式检验邮箱地址的合法性
func CheckEmail(email string) (b bool) {
	m1, _ := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+", email)
	if m1 {
		return true
	}
	if string([]byte(email)[:4]) == "www." || string([]byte(email)[:4]) == "WWW." {
		m2, _ := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+", string([]byte(email)[4:]))
		if m2 {
			return true
		}
	}
	return false
}

//邮箱验证,给输入的邮箱发送验证码,并返回验证码和发送时间
func EmailVerify(emaddr string) (vcode string, sendtime time.Time, err error) {

	// 简单设置 log 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	em := email.NewEmail()

	// 设置 sender 发送方 的邮箱,此处可以填写自己的邮箱
	em.From = "星社官方 <1797249167@qq.com>"

	// 设置 receiver 接收方 的邮箱
	//把地址去除"www."发送,确保用户收到邮件
	if string([]byte(emaddr)[:4]) == "www." || string([]byte(emaddr)[:4]) == "WWW." {
		em.To = []string{string([]byte(emaddr)[4:])}
	} else {
		em.To = []string{emaddr}
	}

	// 设置主题
	em.Subject = "[星社]邮箱验证"

	//生成随机数
	rand1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	randcode := fmt.Sprintf("%06v", rand1.Int31n(1000000))

	//导入随机数
	em.HTML = []byte(fmt.Sprintf(`<div>
        <div>
            尊敬的用户，您好！
        </div> 
            <p>    你本次的验证码为: %s ,为了保证账号安全，验证码有效期为10分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>    
    </div>`, randcode))

	//设置服务器相关的配置
	//err = em.Send("smtp.qq.com:465", smtp.PlainAuth("", "1797249167@qq.com", "muckhyskaauhfidh", "smtp.qq.com"))
	err = em.Send("smtp.qq.com:587", smtp.PlainAuth("", "1797249167@qq.com", "muckhyskaauhfidh", "smtp.qq.com"))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("send successfully ... ")
	return randcode, time.Now(), err
}
