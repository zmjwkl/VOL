package handler

import (
	"VOL/k8s"
	"github.com/gin-gonic/gin"
	"net/http"
)

type K8sCommand struct {
	Action    string `json:"action" binding:"required"`
	Resource  string `json:"resource" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Namespace string `json:"namespace"`
}

func HandleK8sCommand(c *gin.Context) {
	var cmd K8sCommand

	// 解析 JSON
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 执行 Kubernetes 命令
	output, err := k8s.ExecuteCommand(cmd.Action, cmd.Resource, cmd.Name, cmd.Namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "output": output})
		return
	}

	c.JSON(http.StatusOK, gin.H{"output": output})
}

func Handler_get_node(c *gin.Context) {
	username := c.DefaultQuery("username", "all")
	if username == "all" {
		// 执行 Kubernetes 命令
		output, err := k8s.ExecuteCommand_getnodes()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "output": output})
			return
		}
		c.String(http.StatusOK, output)
	} else {
		output, err := k8s.ExecuteCommand_getnode(username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "output": output})
			return
		}
		c.String(http.StatusOK, output)
	}

}
