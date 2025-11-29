# kubiq CLI Architectural Design

## Overview
kubiq is a cross-platform command-line tool that wraps kubectl, providing user-friendly extensions and advanced features. It supports aliases (kbq, kb) and user-defined aliases for kubiq itself, and is designed for extensibility and developer-friendly debugging.

## High-Level Components

### 1. CLI Entry Point
- Parses command-line arguments.
- Handles alias resolution.
- Routes commands to kubectl passthrough or kubiq extension handlers.

### 2. Alias Manager
- Manages built-in and user-defined aliases for kubiq.
- Ensures aliases do not interfere with kubectl.

### 3. Command Router
- If no arguments: displays combined help (kubectl + kubiq extensions).
- If arguments: attempts to run via kubectl.
- On error/unsupported command: routes to kubiq extension handlers.

### 4. Kubectl Wrapper
- Executes kubectl with provided arguments.
- Captures output and errors for further processing.

### 5. Extension Handlers
- Implements custom logic for enhanced commands (e.g., advanced delete with wildcards).
- Example: `kubiq delete pods *abc* -n *aaa*` parses wildcards, lists resources, filters, and deletes matches.

### 6. Debug/Logging Module
- Provides verbose/debug output for developers.
- Logs command execution, errors, and handler activity.

### 7. Cross-Platform Support Layer
- Abstracts OS-specific operations (process spawning, path handling, etc.).
- Ensures consistent behavior on Windows, macOS, and Linux.

## Flow Diagram

```
User Input
   |
   v
[CLI Entry Point]
   |
   v
[Alias Manager] <--- (alias config)
   |
   v
[Command Router]
   |         \
   |          \
   v           v
[Kubectl Wrapper]   [Extension Handlers]
   |                   |
   v                   v
[Output/Errors]   [Custom Logic, then Kubectl Wrapper]
   |                   |
   v                   v
[Debug/Logging]   [Debug/Logging]
   |                   |
   v                   v
[User Output]
```

## Key Design Principles

- **100% kubectl compatibility**: Default behavior is passthrough.
- **Extensible**: New handlers/extensions can be added easily.
- **User-friendly**: Wildcard and batch operations, improved help, and error messages.
- **Cross-platform**: All features work on major OSes.
- **Debuggable**: Verbose logging and error tracing for developers.

---

This document provides a high-level architectural overview for the kubiq CLI tool. For further details or implementation plans, please specify the required section.
