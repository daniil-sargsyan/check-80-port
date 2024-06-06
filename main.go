package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/", handle)
	err := engine.Run(":80")
	if err != nil {
		log.Println(err)
	}
}

func handle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": ":80"})
}
