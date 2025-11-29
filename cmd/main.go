package main

import (
	"fmt"
	"os"
	"os/exec"
)

const kubiqVersion = "0.1.4-beta"

func main() {
	args := os.Args[1:]
	debug := false
	for _, arg := range args {
		if arg == "--debug" {
			debug = true
			break
		}
	}
	if len(args) == 0 {
		fmt.Printf("kubiq version %s (beta)\n", kubiqVersion)
		fmt.Println("-- kubiq wraps kubectl and adds extensions --")
		cmd := exec.Command("kubectl", "help")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Env = os.Environ()
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error running kubectl help: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if len(args) == 1 && (args[0] == "version" || args[0] == "--version" || args[0] == "-v") {
		fmt.Printf("kubiq version %s (beta)\n", kubiqVersion)
		return
	}

	if debug {
		printDebugInfo()
	}

	cmd := exec.Command("kubectl", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
		fmt.Fprintf(os.Stderr, "Error running kubectl: %v\n", err)
		os.Exit(1)
	}
}
