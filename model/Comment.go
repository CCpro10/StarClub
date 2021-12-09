package model

import "time"

//帖子社区下面的评论
type Comment struct {
	ID        uint
	Postid    int    `form:"postid" `  //帖子的id
	Userid    int    `form:"userid" `  //后端根据token识别的用户id
	Author    string `form:"author" `  //用户呢名
	Context   string `form:"context" ` //用户输入的内容
	CreatedAt time.Time
}
