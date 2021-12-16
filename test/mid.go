package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()
	router.Use(FirstMiddleware(), SecondMiddleware())
	router.GET("/", func(c *gin.Context) {
		fmt.Println("process get request")
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello",
		})
	}) // (1)
	router.Run(":9001")
}

func FirstMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("first middleware before next()")
		isAbort := c.Query("isAbort")
		bAbort, err := strconv.ParseBool(isAbort)
		if err != nil {
			fmt.Printf("is abort value err, value %s\n", isAbort)
			c.Next() // (2)
		}
		if bAbort {
			fmt.Println("first middleware abort") //(3)
			//c.Abort()
			//c.AbortWithStatusJSON(http.StatusOK, "abort is true")
			return
		} else {
			fmt.Println("first middleware doesnot abort") //(4)
			return
		}

		fmt.Println("first middleware after next()")
	}
}

func SecondMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("current inside of second middleware")
	}
}
