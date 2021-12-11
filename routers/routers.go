package routers

import (
	"StarClub/middleware"
	"StarClub/service"
	"github.com/gin-gonic/gin"
)

func BeganRouters(r *gin.Engine) {
	r.Use(middleware.Cors())
	r.POST("/sendverifycode", service.SendVerifyCode)
	r.POST("/register", service.Register)
	r.POST("/login", service.Login)
	Authgroup := r.Group("/auth")
	Authgroup.Use(middleware.JWTAuthMiddleware)
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

}
