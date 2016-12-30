package main

import (
	"net/http"
	"os/exec"
	"strings"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("resources/templates/*")
	router.Static("/static", "resources/static")

	router.GET("/", index)
	router.GET("/index", index)

	router.Run("localhost:8080")
}

func index(c *gin.Context) {
	tmux_path, _ := c.Get("bin")
	if tmux_path == nil {
		tmux_path, _ = exec.LookPath("tmux")
	}

	sessions_cmd := exec.Command("tmux", "list-sessions")
	out, _ := sessions_cmd.CombinedOutput()
	sessions := strings.Split(strings.TrimSpace(string(out)), "\n")

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":     "tmux control panel",
		"tmux_path": tmux_path,
		"sessions":  sessions,
	})
}
