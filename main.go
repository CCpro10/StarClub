package main

import (
	"StarClub/dao"
	"StarClub/model"
	"StarClub/service"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {

	//连接数据库
	dao.InitMySQL()

	//连接Redis
	dao.InitRedis()

	//defer dao.Close()  // 程序退出关闭数据库连接
	r := gin.Default()

	r.POST("/sendverifycode", service.SendVerifyCode)
	r.POST("/register", service.Register)
	r.POST("/login", service.Login)

	Authgroup := r.Group("/auth")
	Authgroup.Use(model.JWTAuthMiddleware)
	{
		//查看单个活动
		Authgroup.GET("/activity", service.ShowActivity)
		//查看所有活动
		Authgroup.GET("/activities", service.ShowActivities)
		//查看我关注的社团的活动
		Authgroup.GET("/myclubactivities", service.ShowMyClubActivities)
		//查看我的活动
		Authgroup.GET("/myactivities", service.ShowMyActivities)

		//社团发布活动
		Authgroup.POST("/activity", service.PostActivity)

		//关注社团
		Authgroup.POST("/subscribe", service.FollowClub)
		//取消关注
		Authgroup.DELETE("/subscribe", service.Unsubscribe)

		//添加我的活动
		Authgroup.POST("/myactivity", service.AddActivity)
		//取消添加
		Authgroup.DELETE("/myactivity", service.CancelActivity)

	}

	if err := r.Run(":9999"); err != nil {
		panic(err)
	}
}
