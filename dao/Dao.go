package dao

import (
	"StarClub/model"
	"StarClub/util"
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//获取配置文件
var Conf = util.GetConf()

var CTX = context.Background()
var RedisDB = redis.NewClient(&redis.Options{
	Addr:     Conf.Redis.Addr,
	Password: Conf.Redis.Password, // no password set
	DB:       Conf.Redis.DB,       // use default DB
})

//Redis相关全局变量
func InitRedis() {
	_, err := RedisDB.Ping(CTX).Result()
	if err != nil {
		panic(err)
	}
}

//mysql,gorm配置
var DB *gorm.DB

func InitMySQL() {

	dsn := Conf.MYSQL.Username + ":" +
		Conf.MYSQL.Password + "@tcp(" +
		Conf.MYSQL.Addr + ")/" +
		Conf.MYSQL.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open("mysql", dsn)
	//DB, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//先创建表
	DB.AutoMigrate(
		model.Activity{},
		model.UserInfo{},
		model.Comment{},
		model.Post{},
		model.ClubFollows{},
		model.MyActivity{},
		model.ClubList{},
	)

}
