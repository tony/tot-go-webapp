package main

import (
	"net/http"
	"os/exec"
	"strings"

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

	sessions_cmd := exec.Command("tmux", "list-sessions")
	out, _ := sessions_cmd.CombinedOutput()
	sessions := strings.TrimSpace(string(out))

	c.HTML(http.StatusOK, "hi.html", gin.H{
		"title":     "Main website",
		"tmux_path": tmux_path,
		"sessions":  sessions,
	})
}
