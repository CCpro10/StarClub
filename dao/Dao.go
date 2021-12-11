package dao

import (
	"StarClub/model"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Redis相关全局变量
var CTX = context.Background()

var RedisDB = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func InitRedis() {
	_, err := RedisDB.Ping(CTX).Result()
	if err != nil {
		panic(err)
	}
}

//mysql,gorm配置
var (
	DB *gorm.DB
)

func InitMySQL() {
	//dsn := "debian-sys-maint:YW6xCg7iemGYaPGe@tcp(127.0.0.1:3306)/dbstarclub?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "chenchao:111111@tcp(127.0.0.1:3306)/db2?charset=utf8mb4&parseTime=True&loc=Local"
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
