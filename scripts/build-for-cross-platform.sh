#!/bin/bash
# build-for-cross-platform.sh
# Build kubiq for Windows, Linux, and macOS from any platform (requires Go installed)
# See related documentation for more details

set -e

# Set output directory
OUTDIR=build
mkdir -p "$OUTDIR"

# Build for Linux (x86_64)
echo "Building kubiq for Linux..."
GOOS=linux GOARCH=amd64 go build -o "$OUTDIR/kubiq" ./cmd

# Build for Windows (x86_64)
echo "Building kubiq for Windows..."
GOOS=windows GOARCH=amd64 go build -o "$OUTDIR/kubiq.exe" ./cmd

# Build for macOS (x86_64)
echo "Building kubiq for macOS..."
GOOS=darwin GOARCH=amd64 go build -o "$OUTDIR/kubiq-macos" ./cmd

echo "Builds complete! Binaries are in the $OUTDIR directory."

echo "\nQuick usage:"
echo "  Linux:   ./build/kubiq"
echo "  Windows: ./build/kubiq.exe"
echo "  macOS:   ./build/kubiq-macos"

echo "\nFor install and usage instructions, see related documentation."