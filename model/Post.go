package model

import "time"

//社团发布的讨论帖
type Post struct {
	ID        uint   `gorm:"primary_key" `
	Userid    int    `form:"userid"`  //社团的id(主键)
	Article   string `form:"article"` //标题
	Context   string `form:"context"` //内容
	Author    string `form:"author"`  //作者,就是社团的名字
	CreatedAt time.Time
}
