package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func printWithGuidance(r io.Reader) {
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
