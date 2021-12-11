package service

import (
	"StarClub/dao"
	"StarClub/model"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"

	"log"
	"net/http"
)

//关注,传入 clubid
func FollowClub(c *gin.Context) {
	//绑定参数
	var clubfollers = model.ClubFollows{}
	if err := c.ShouldBind(&clubfollers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数绑定失败"})
		return
	}
	//获取用户信息
	userid, ok := c.Get("userid")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "获取用户信息失败"})
		return
	}

	//查询有没有关注此社团
	var clubfoller model.ClubFollows

	dao.DB.Where("club_id=? AND user_id=?", clubfollers.ClubId, userid).First(&clubfoller)
	if clubfoller.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "您已关注此社团"})
		return
	} else { //创建关注关系
		clubfollers.UserId = userid.(uint)
		if err := dao.DB.Create(&clubfollers).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "创建失败"})
			log.Println(err.Error())
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": "关注成功"})
			return
		}
	}
}

//取消关注,传入 clubid
func Unsubscribe(c *gin.Context) {
	//绑定参数
	var clubfollers = model.ClubFollows{}
	if err := c.ShouldBind(&clubfollers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数绑定失败"})
		return
	}
	//获取用户信息
	userid, ok := c.Get("userid")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "获取用户信息失败"})
		return
	}

	//查询有没有关注此社团
	var clubfoller model.ClubFollows
	dao.DB.Where("club_id=? AND user_id=?", clubfollers.ClubId, userid).First(&clubfoller)
	if clubfoller.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "已取消关注此社团"})
		return
	} else { //创建关注关系
		if err := dao.DB.Delete(&clubfollers, clubfoller.ID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "取消关注失败"})
			log.Println(err.Error())
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": "取消关注成功"})
			return
		}
	}
}
