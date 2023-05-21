package handlers

import "github.com/gin-gonic/gin"

func HandleSuccess(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"status": "success",
		"data":   data,
		"error":  nil,
	})
}
