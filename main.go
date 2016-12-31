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
	router.GET("/tmux_partial", tmuxPartial)

	router.Run("localhost:8080")
}

func getList(tmuxPath string, cmd string) []string {
	_cmd := exec.Command(tmuxPath, cmd)
	out, err := _cmd.Output()
	if err == nil {
		return strings.Split(strings.TrimSpace(string(out)), "\n")
	}
	return []string{}
}

type tmuxData struct {
	sessions []string
	windows  []string
	panes    []string
	clients  []string
}

func getSessions(tmuxPath string) []string {
	return getList(tmuxPath, "list-sessions")
}

func getWindows(tmuxPath string) []string {
	return getList(tmuxPath, "list-windows")
}

func getPanes(tmuxPath string) []string {
	return getList(tmuxPath, "list-panes")
}

func getClients(tmuxPath string) []string {
	return getList(tmuxPath, "list-clients")
}

func getTmuxData(tmuxPath string) tmuxData {
	return tmuxData{
		getSessions(tmuxPath),
		getWindows(tmuxPath),
		getPanes(tmuxPath),
		getClients(tmuxPath),
	}
}

func tmuxPartial(c *gin.Context) {
	tmuxPath := c.Query("tmux_path")

	if tmuxPath == "" {
		tmuxPath, _ = exec.LookPath("tmux")
	}

	tmuxData := getTmuxData(tmuxPath)

	c.HTML(http.StatusOK, "content.html", gin.H{
		"tmux_path": tmuxPath,
		"sessions":  tmuxData.sessions,
		"windows":   tmuxData.windows,
		"panes":     tmuxData.panes,
		"clients":   tmuxData.clients,
	})
}

func index(c *gin.Context) {
	tmuxPath := c.Query("tmux_path")

	if tmuxPath == "" {
		tmuxPath, _ = exec.LookPath("tmux")
	}

	tmuxData := getTmuxData(tmuxPath)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":     "tmux control panel",
		"tmux_path": tmuxPath,
		"sessions":  tmuxData.sessions,
		"windows":   tmuxData.windows,
		"panes":     tmuxData.panes,
		"clients":   tmuxData.clients,
	})
}
