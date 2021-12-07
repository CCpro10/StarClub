package model
import "time"

//活动
type Activity struct {

	ID        uint   `form:"primary_key"`
	Userid    int    `form:"userid"`  //活动发布者的id
	Author    string `form:"author"` //发布者的呢名(社团名称)
	Article   string `form:"article"` //活动标题
	Address   string  `form:"address"`//活动地点
	Context   string `form:"context"` //发布的内容
	CreatedAt time.Time              //发布时间

}
