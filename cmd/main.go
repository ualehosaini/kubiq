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

	// Helper to process and print kubectl output, replacing 'kubectl' with 'kubiq' and adding guidance
	printKubiqHelp := func(r io.Reader) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r)
		out := buf.String()
		replaced := strings.ReplaceAll(out, "kubectl", "kubiq")
		fmt.Print(replaced)
		if strings.Contains(strings.ToLower(replaced), "usage:") || strings.Contains(strings.ToLower(replaced), "help") {
			fmt.Println("\n[GUIDANCE] You are using kubiq, a wrapper for kubectl. All kubectl commands and flags work the same way, but you can use 'kubiq' instead of 'kubectl' in all examples and help texts.")
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
		printKubiqHelp(&outBuf)
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
	// Print kubectl output, replacing 'kubectl' with 'kubiq' and adding guidance if help/usage detected
	if outBuf.Len() > 0 {
		printKubiqHelp(&outBuf)
	}
	if errBuf.Len() > 0 {
		printKubiqHelp(&errBuf)
	}
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
		fmt.Fprintf(os.Stderr, "Error running kubectl: %v\n", err)
		os.Exit(1)
	}
}
