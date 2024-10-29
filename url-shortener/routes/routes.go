package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener/handlers"
)

func GetUrl(c *gin.Context) {
	url, err := handlers.GetOriginalUrl() //TODO
	c.JSON(http.StatusOK, handlers.Shortener(url))
}

func ShortenUrl(c *gin.Context) {
	c.JSON(http.StatusOK, handlers.Shortener(url))
}
