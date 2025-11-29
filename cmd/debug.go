package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func logDebug(level, msg string) {
       now := time.Now()
       ts := now.Format("0102 15:04:05.000000")
       var colorStart, colorEnd string
       switch level {
       case "ERROR":
	       colorStart = "\033[31m" // Red
       case "INFO ":
	       colorStart = "\033[33m" // Yellow
       default:
	       colorStart = ""
       }
       colorEnd = "\033[0m"
       fmt.Printf("%sI%s   %s [kubiq debug] %s%s\n", colorStart, ts, level, msg, colorEnd)
}

func printDebugInfo() {
	kubectlPath, err := exec.LookPath("kubectl")
	if err != nil {
		logDebug("INFO ", "kubectl not found in PATH")
	} else {
		logDebug("INFO ", fmt.Sprintf("kubectl path: %s", kubectlPath))
	}

	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		logDebug("INFO ", "KUBECONFIG not set (using default: ~/.kube/config)")
	} else {
		logDebug("INFO ", fmt.Sprintf("KUBECONFIG: %s", kubeconfig))
	}

	cmd := exec.Command("kubectl", "config", "current-context")
	cmd.Env = os.Environ()
	output, err := cmd.CombinedOutput()
	if err != nil {
		logDebug("ERROR", fmt.Sprintf("Error getting current context: %v", err))
	} else {
		logDebug("INFO ", fmt.Sprintf("kubectl current-context: %s", strings.TrimSpace(string(output))))
	}
}
