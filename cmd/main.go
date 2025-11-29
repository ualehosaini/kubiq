package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

const kubiqVersion = "0.1.5-beta"

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

	// Helper to print kubectl output and add guidance if help/usage detected
	printWithGuidance := func(r io.Reader) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r)
		out := buf.String()
		showGuidance := strings.Contains(strings.ToLower(out), "usage:") || strings.Contains(strings.ToLower(out), "help")
		yellow := "\033[33m"
		reset := "\033[0m"
		guidance := yellow + "[GUIDANCE] The following/above help text is from kubectl. You can use all the same commands with kubiq by simply replacing 'kubectl' with 'kubiq'. kubiq is a drop-in replacement for kubectl." + reset + "\n"
		if showGuidance {
			fmt.Print(guidance)
		}
		fmt.Print(out)
		if showGuidance {
			fmt.Print(guidance)
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
