package system

import (
	"github.com/gin-gonic/gin"
)

// Ping Response
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
