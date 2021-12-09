package main

import "time"
import "fmt"

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Unix())                               //秒
	fmt.Println(time.Now().UnixNano() / int64(time.Millisecond)) //毫秒
	fmt.Println(time.Now().UnixNano() / 1000)                    //微秒
	fmt.Println(time.Now().UnixNano())                           //纳秒
}
