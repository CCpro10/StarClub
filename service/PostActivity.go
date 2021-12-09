package service

import (
	"StarClub/dao"
	"StarClub/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//发布活动,接收标题,内容,地点,时间
func PostActivity(c *gin.Context) {

	var RequestActivity model.Activity
	// 数据绑定
	if err := c.ShouldBind(&RequestActivity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "活动参数绑定失败"})
		log.Printf(err.Error())
		return
	}

	// 从token中获取登录用户的id
	userid, ok := c.Get("userid")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "获取用户信息失败"})
		return
	}
	//查询是否为社团号,检查权限
	var user model.UserInfo
	if err := dao.DB.Where("ID=?", userid.(uint)).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "用户查询失败"})
		log.Println(err.Error())
		return
	}
	if !user.IsClub {
		c.JSON(http.StatusForbidden, gin.H{"msg": "只有社团才可以发布活动"})
		return
	}

	//设置活动的社团名称和社团id
	RequestActivity.Author = user.Username
	RequestActivity.UserId = userid.(uint)
	// 创建活动
	if err := dao.DB.Create(&RequestActivity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "活动创建失败"})
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"mes":   "创建成功",
		"创建的结构": RequestActivity,
	})
}
