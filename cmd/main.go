package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

const kubiqVersion = "0.1.6-beta"

func main() {
	origArgs := os.Args[1:]
	debug := false
	var args []string
	for _, arg := range origArgs {
		if arg == "--debug" {
			debug = true
		} else {
			args = append(args, arg)
		}
	}

	if len(args) == 0 {
		fmt.Printf("kubiq version %s (beta)\n", kubiqVersion)
		fmt.Println("-- kubiq wraps kubectl and adds extensions --")
		cmd := exec.Command("kubectl", "help")
		var outBuf bytes.Buffer
		cmd.Stdout = &outBuf
		cmd.Stderr = &outBuf
		cmd.Env = os.Environ()
		err := cmd.Run()
		printWithGuidance(&outBuf)
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
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()
	err := cmd.Run()
	if outBuf.Len() > 0 {
		printWithGuidance(&outBuf)
	}
	if errBuf.Len() > 0 {
		printWithGuidance(&errBuf)
	}
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
		fmt.Fprintf(os.Stderr, "Error running kubectl: %v\n", err)
		os.Exit(1)
	}
}
