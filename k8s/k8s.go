package k8s

import (
	"fmt"
	"os/exec"
)

// ExecuteCommand 执行 k8s 命令
func ExecuteCommand(action, resource, name, namespace string) (string, error) {
	// 示例：使用 kubectl 执行命令
	cmd := exec.Command("kubectl", action, resource, name, "-n", namespace)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v, output: %s", err, string(output)+" "+cmd.String()+"\n")
	}
	return string(output), nil
}
