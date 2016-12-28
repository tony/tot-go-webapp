package main

import "net/http"

import "gopkg.in/gin-gonic/gin.v1"

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", hi)
	router.GET("/index", hi)
	router.Run("localhost:8080")
}

func hi(c *gin.Context) {
	c.HTML(http.StatusOK, "hi.html", gin.H{
		"title": "Main website",
	})
}
