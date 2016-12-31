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

func getList(tmuxbin string, cmd string) []string {
	_cmd := exec.Command(tmuxbin, cmd)
	out, err := _cmd.Output()
	if err == nil {
		return strings.Split(strings.TrimSpace(string(out)), "\n")
	} else {
		return []string{}
	}
}

func getSessions(tmuxbin string) []string {
	return getList(tmuxbin, "list-sessions")
}

func getWindows(tmuxbin string) []string {
	return getList(tmuxbin, "list-windows")
}

func getPanes(tmuxbin string) []string {
	return getList(tmuxbin, "list-panes")
}

func getClients(tmuxbin string) []string {
	return getList(tmuxbin, "list-clients")
}

func index(c *gin.Context) {
	tmuxPath := c.Query("tmux_path")

	if tmuxPath == "" {
		tmuxPath, _ = exec.LookPath("tmux")
	}

	sessions := getSessions(tmuxPath)
	windows := getWindows(tmuxPath)
	fmt.Printf("windows: %v", windows)
	println(len(windows))
	panes := getPanes(tmuxPath)
	clients := getClients(tmuxPath)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":     "tmux control panel",
		"tmux_path": tmuxPath,
		"sessions":  sessions,
		"windows":   windows,
		"panes":     panes,
		"clients":   clients,
	})
}
