package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou/falcon-plus/modules/falcon2p8s/g"
)

func CheckIsScraping() gin.HandlerFunc {
	return func(c *gin.Context) {
		g.IsScraping = true
		c.Next()
		g.IsScraping = false
	}
}
