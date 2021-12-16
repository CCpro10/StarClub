package model

import "time"

//活动
type Activity struct {
	ID           uint   `form:"primary_key"`
	UserId       uint   `form:"userid"`                          //活动发布者(社团)的id
	Author       string `form:"author"`                          //发布者的呢名(社团名称)
	Article      string `form:"article"`                         //活动标题
	Address      string `form:"address"`                         //活动地点
	Content      string `form:"content"`                         //发布的内容
	ActivityTime int64  `form:"activitytime"json:"activitytime"` //活动时间,毫秒时间戳

	CollectionNumbers int       `form:"collectionnumbers"json:"collectionnumbers"gorm:"default:0"` //活动收藏人数
	CreatedAt         time.Time //发布时间

}
