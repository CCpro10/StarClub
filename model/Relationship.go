package model

//表示用户和关注的社团
type ClubFollows struct {
	ID     uint `gorm:"primary_key"`
	ClubId uint `form:"clubid"json:"clubid"  ` //社团ID
	UserId uint `form:"userid"json:"userid"`   //用户ID
}

//表示 我的活动
type MyActivity struct {
	ID         uint `gorm:"primary_key"`
	ActivityId uint `form:"activityid"json:"activityid"` //活动ID
	UserId     uint `form:"userid"json:"userid"`         //用户ID
}

//所有社团列表
type ClubList struct {
	ID       uint   `gorm:"primary_key"`
	ClubId   uint   `form:"userid" `  //社团用户的id(主键)
	ClubName string `form:"clubname"` //社团的名称(username)
}
