package main

import (
	"github.com/gin-gonic/gin"
	"wxcloudrun-golang/msg"
)

func main() {

	r := gin.Default()
	r.GET("/send", msg.LetSendMsg)
}
