package service

import (
	"StarClub/dao"
	"StarClub/model"
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
)

//按活动时间降序展示所有活动
func ShowActivities(c *gin.Context) {
	offsetnum := c.Query("offsetnum")
	var activities []model.Activity
	if err := dao.DB.Order("activity_time").Offset(offsetnum).Limit(10).Find(&activities); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询访问失败"})
		log.Println(err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": activities, "msg": "查询成功"})
	}
}

//显示单个活动详情
func ShowActivity(c *gin.Context) {
	activityid := c.Query("activityid")
	var activity model.Activity
	dao.DB.Where("ID=?", activityid).Find(&activity)
	if activity.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "此活动不存在"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": activity, "msg": "查询成功"})
		return
	}

}

//按活动时间降序展示所有活动
func ShowMyActivities(c *gin.Context) {
	offsetnum := c.Query("offsetnum")
	// 从token中获取登录用户的id
	userid, ok := c.Get("userid")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "获取用户信息失败"})
		return
	}

	userid = userid.(uint)
	var activities []model.Activity

	dao.DB.Where("userid=?",
		dao.DB.Table("club_follows").Select("club_id").Where("user_id=?", userid)).
		Order("activity_time").Find(&activities)

	//if err:=dao.DB.Where().Order("activity_time").Offset(offsetnum).Limit(10).Find(&activities);err!=nil{
	//	c.JSON(http.StatusInternalServerError,gin.H{"msg":"数据库查询访问失败"})
	//	log.Println(err)
	//	return
	//}else {
	//	c.JSON(http.StatusOK,gin.H{"data":activities,"msg":"查询成功"})
	//}
}
