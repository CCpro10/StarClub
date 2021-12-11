package service

import (
	"StarClub/dao"
	"StarClub/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//按活动时间降序展示所有活动,只需要query中的offsetnum
func ShowActivities(c *gin.Context) {
	offset := c.Query("offsetnum")
	//检查offsetnum格式
	offsetnum, err := strconv.Atoi(offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "分页格式错误"})
		log.Println(err)
		return
	}
	var activities []model.Activity
	dao.DB.Offset(offsetnum).Limit(10).Order("activity_time DESC").Find(&activities)
	//通过redis获取活动收藏数
	if err := dao.GetActivitiesCollectionNumbers(activities); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"data": activities, "msg": "活动收藏数获取失败"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": activities, "msg": "访问成功"})
		return
	}
}

//显示单个活动详情,query传入activityid
func ShowActivity(c *gin.Context) {
	activityid := c.Query("activityid")
	var activity model.Activity
	dao.DB.Where("ID=?", activityid).Find(&activity)
	if activity.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "此活动不存在"})
		return
	} else {
		//通过redis获取活动收藏数
		if err := dao.GetActivityCollectionNumbers(&activity); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "活动收藏数获取失败"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"data": activity, "msg": "访问成功"})
			return
		}
	}
}

//按活动时间降序展示我关注的社团的所有活动,需要query中的offsetnum
func ShowMyClubActivities(c *gin.Context) {
	offset := c.Query("offsetnum")
	//检查offsetnum格式
	offsetnum, err := strconv.Atoi(offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "分页格式错误"})
		log.Println(err)
		return
	}
	// 从token中获取登录用户的id
	userid, ok := c.Get("userid")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "获取用户信息失败"})
		return
	}
	userid = userid.(uint)
	//先找用户关注的社团ID,再找社团id对应的活动
	var clubfollows []model.ClubFollows

	dao.DB.Table("club_follows").Select("club_id").Where("user_id=?", userid).Scan(&clubfollows)
	//找到我关注的社团id列表
	clubids := make([]uint, len(clubfollows))
	for k, v := range clubfollows {
		clubids[k] = v.ClubId
	}
	var activities []model.Activity
	dao.DB.Where("user_id in (?)", clubids).Order("activity_time DESC").Offset(offsetnum).Limit(10).Find(&activities)
	//通过redis获取活动收藏数
	if err := dao.GetActivitiesCollectionNumbers(activities); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"data": activities, "msg": "活动收藏数获取失败"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": activities, "msg": "访问成功"})
		return
	}
}

//展示我添加的社团的活动,需要query中的offsetnum
func ShowMyActivities(c *gin.Context) {
	offset := c.Query("offsetnum")
	//检查offsetnum格式
	offsetnum, err := strconv.Atoi(offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "分页格式错误"})
		log.Println(err)
		return
	}
	// 从token中获取登录用户的id
	userid, ok := c.Get("userid")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "获取用户信息失败"})
		return
	}
	userid = userid.(uint)
	//找用户添加的活动
	var myactivities []model.MyActivity

	dao.DB.Table("my_activities").Select("activity_id").Where("user_id=?", userid).Scan(&myactivities)
	//找到我关注的社团id列表
	activityids := make([]uint, len(myactivities))
	for k, v := range myactivities {
		activityids[k] = v.ActivityId
	}
	var activities []model.Activity
	dao.DB.Where("id in (?)", activityids).Order("activity_time DESC").Offset(offsetnum).Limit(10).Find(&activities)
	//通过redis获取活动收藏数
	if err := dao.GetActivitiesCollectionNumbers(activities); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"data": activities, "msg": "活动收藏数获取失败"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": activities, "msg": "访问成功"})
		return
	}
}
