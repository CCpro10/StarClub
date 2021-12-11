package main

import (
	"StarClub/dao"
	"StarClub/routers"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	//连接数据库
	dao.InitMySQL()
	//连接Redis
	dao.InitRedis()
}

func main() {
	r := gin.Default()
	routers.BeganRouters(r)
	if err := r.Run(":9999"); err != nil {
		panic(err)
	}
}
