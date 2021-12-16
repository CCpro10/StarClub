package service

import (
	"StarClub/dao"
	"StarClub/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//关注,传入 activityid
func AddActivity(c *gin.Context) {
	//绑定传入的参数
	var myactivity model.MyActivity
	if err := c.ShouldBind(&myactivity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数绑定失败"})
		return
	}
	//获取用户信息
	userid, ok := c.Get("userid")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "获取用户信息失败"})
		return
	}
	//查询有没有添加活动
	var myactivities model.MyActivity
	dao.DB.Where("activity_id=? AND user_id=?", myactivity.ActivityId, userid).First(&myactivities)
	if myactivities.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "您已添加此活动"})
		return
	} else { //添加活动
		//事务操作
		tx := dao.DB.Begin()
		myactivity.UserId = userid.(uint)
		if err := tx.Create(&myactivity).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "活动添加失败"})
			log.Println(err.Error())
			return
		} else {
			if _, err := dao.RedisDB.Incr(dao.CTX, strconv.Itoa(int(myactivity.ActivityId))).Result(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "活动添加失败"})
				//redis自增失败,回滚
				tx.Rollback()
				log.Println(err.Error())
				return
			} else {
				tx.Commit()
				c.JSON(http.StatusOK, gin.H{"msg": "添加成功"})
				return
			}
		}
	}
}

//取消关注,传入 activityid
func CancelActivity(c *gin.Context) {
	//绑定参数
	var myactivity = model.MyActivity{}
	if err := c.ShouldBind(&myactivity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数绑定失败"})
		return
	}
	//获取用户信息
	userid, ok := c.Get("userid")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "获取用户信息失败"})
		return
	}

	//查询有没有添加活动
	var myactivities model.MyActivity
	dao.DB.Where("activity_id=? AND user_id=?", myactivity.ActivityId, userid).First(&myactivities)
	if myactivities.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "您还未关注此活动"})
		return
	} else { //取消活动
		//mysql事务
		tx := dao.DB.Begin()
		myactivity.UserId = userid.(uint)
		if err := tx.Delete(&myactivity).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "活动取消失败"})
			log.Println(err.Error())
			return
		} else {
			//redis活动关注数
			if _, err := dao.RedisDB.Decr(dao.CTX, strconv.Itoa(int(myactivity.ActivityId))).Result(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "活动取消失败"})
				//redis自减失败,回滚
				tx.Rollback()
				log.Println(err.Error())
				return
			} else {
				tx.Commit()
				c.JSON(http.StatusOK, gin.H{"msg": "取消成功"})
				return
			}
		}
	}
}
