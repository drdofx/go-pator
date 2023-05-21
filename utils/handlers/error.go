package handlers

import "github.com/gin-gonic/gin"

func HandleError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{
		"status": "error",
		"data":   nil,
		"error":  err,
	})
}
