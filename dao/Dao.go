package dao

import (
	"StarClub/conf"
	"StarClub/model"
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var CTX = context.Background()
var RedisDB = redis.NewClient(&redis.Options{
	Addr:     conf.Config.Redis.Addr,
	Password: conf.Config.Redis.Password, // no password set
	DB:       conf.Config.Redis.DB,       // use default DB
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

	dsn := conf.Config.MYSQL.Username + ":" +
		conf.Config.MYSQL.Password + "@tcp(" +
		conf.Config.MYSQL.Addr + ")/" +
		conf.Config.MYSQL.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) //这里用短变量声明会有歧义
	//DB, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//先创建表
	if err = DB.AutoMigrate(
		model.Activity{},
		model.UserInfo{},
		model.Comment{},
		model.Post{},
		model.ClubFollows{},
		model.MyActivity{},
		model.ClubList{},
	); err != nil {
		log.Panicln(err)
	}

}
