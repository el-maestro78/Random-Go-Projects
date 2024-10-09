package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func getPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", index)
	r.GET("/ping", getPing)
	err := r.Run(":8070")
	if err != nil {
		return
	}
}
