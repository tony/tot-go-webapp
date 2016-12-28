package main

import (
	"net/http"
	"os/exec"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", hi)
	router.GET("/index", hi)
	router.Run("localhost:8080")
}

func hi(c *gin.Context) {
	tmux_path, _ := c.Get("bin")
	if tmux_path == nil {
		tmux_path, _ = exec.LookPath("tmux")
	}

	c.HTML(http.StatusOK, "hi.html", gin.H{
		"title":     "Main website",
		"tmux_path": tmux_path,
	})
}
