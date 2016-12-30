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

func get_sessions(tmuxbin string) []string {
	_cmd := exec.Command(tmuxbin, "list-sessions")
	out, _ := _cmd.CombinedOutput()
	return strings.Split(strings.TrimSpace(string(out)), "\n")
}

func get_windows(tmuxbin string) []string {
	_cmd := exec.Command(tmuxbin, "list-windows")
	out, _ := _cmd.CombinedOutput()
	return strings.Split(strings.TrimSpace(string(out)), "\n")
}

func get_panes(tmuxbin string) []string {
	_cmd := exec.Command(tmuxbin, "list-panes")
	out, _ := _cmd.CombinedOutput()
	return strings.Split(strings.TrimSpace(string(out)), "\n")
}

func get_clients(tmuxbin string) []string {
	_cmd := exec.Command(tmuxbin, "list-clients")
	out, _ := _cmd.CombinedOutput()
	return strings.Split(strings.TrimSpace(string(out)), "\n")
}

func index(c *gin.Context) {
	tmux_path := c.Query("tmux_path")

	if tmux_path == "" {
		tmux_path, _ = exec.LookPath("tmux")
	}

	sessions := get_sessions(tmux_path)
	windows := get_windows(tmux_path)
	panes := get_panes(tmux_path)
	clients := get_clients(tmux_path)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":     "tmux control panel",
		"tmux_path": tmux_path,
		"sessions":  sessions,
		"windows":   windows,
		"panes":     panes,
		"clients":   clients,
	})
}
