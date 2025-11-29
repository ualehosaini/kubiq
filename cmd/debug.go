package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func printDebugInfo() {
	kubectlPath, err := exec.LookPath("kubectl")
	if err != nil {
		fmt.Println("[kubiq debug] kubectl not found in PATH")
	} else {
		fmt.Printf("[kubiq debug] kubectl path: %s\n", kubectlPath)
	}

	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		fmt.Println("[kubiq debug] KUBECONFIG not set (using default: ~/.kube/config)")
	} else {
		fmt.Printf("[kubiq debug] KUBECONFIG: %s\n", kubeconfig)
	}

	cmd := exec.Command("kubectl", "config", "current-context")
	cmd.Env = os.Environ()
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("[kubiq debug] Error getting current context: %v\n", err)
	} else {
		fmt.Printf("[kubiq debug] kubectl current-context: %s\n", strings.TrimSpace(string(output)))
	}
}
