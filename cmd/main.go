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
	printKubiqHelp := func(r io.Reader, doReplace bool) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r)
		out := buf.String()
		if doReplace {
			out = strings.ReplaceAll(out, "kubectl", "kubiq")
		}
		fmt.Print(out)
		if doReplace && (strings.Contains(strings.ToLower(out), "usage:") || strings.Contains(strings.ToLower(out), "help")) {
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
		printKubiqHelp(&outBuf, true)
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
	// Detect if --v=N or -v=N is present (for kubectl verbose logs)
	doReplace := true
	for _, a := range args {
		if strings.HasPrefix(a, "--v=") || strings.HasPrefix(a, "-v=") {
			doReplace = false
			break
		}
	}
	err := cmd.Run()
	// Print kubectl output, replacing 'kubectl' with 'kubiq' and adding guidance if help/usage detected
	if outBuf.Len() > 0 {
		printKubiqHelp(&outBuf, doReplace)
	}
	if errBuf.Len() > 0 {
		printKubiqHelp(&errBuf, doReplace)
	}
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
		fmt.Fprintf(os.Stderr, "Error running kubectl: %v\n", err)
		os.Exit(1)
	}
}
