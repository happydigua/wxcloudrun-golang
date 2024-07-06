package msg

import (
	"github.com/gin-gonic/gin"
)

func LetSendMsg(c *gin.Context) {

	// 打印路径参数
	for key, value := range c.Params {
		c.String(200, "Path Param: %s = %s\n", key, value)
	}
}
