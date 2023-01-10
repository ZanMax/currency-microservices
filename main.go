package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "currency",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
