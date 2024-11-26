package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Main() {
	router := gin.Default()

	// Employee API Method
	router.GET("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "employee Get Method",
		})
	})

	router.POST("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "employee Get Method",
		})
	})

	router.PUT("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Employee PUT Method",
		})
	})

	router.DELETE("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Employee DELETE Method",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
