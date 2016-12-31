package main

import (
	"fmt"
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

func get_list(tmuxbin string, cmd string) []string {
	_cmd := exec.Command(tmuxbin, cmd)
	out, err := _cmd.Output()
	if err == nil {
		return strings.Split(strings.TrimSpace(string(out)), "\n")
	} else {
		return []string{}
	}
}

func get_sessions(tmuxbin string) []string {
	return get_list(tmuxbin, "list-sessions")
}

func get_windows(tmuxbin string) []string {
	return get_list(tmuxbin, "list-windows")
}

func get_panes(tmuxbin string) []string {
	return get_list(tmuxbin, "list-panes")
}

func get_clients(tmuxbin string) []string {
	return get_list(tmuxbin, "list-clients")
}

func index(c *gin.Context) {
	tmux_path := c.Query("tmux_path")

	if tmux_path == "" {
		tmux_path, _ = exec.LookPath("tmux")
	}

	sessions := get_sessions(tmux_path)
	windows := get_windows(tmux_path)
	fmt.Printf("windows: %v", windows)
	println(len(windows))
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
