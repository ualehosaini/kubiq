# kubiq

A cross-platform command-line tool that wraps `kubectl` and provides powerful, user-friendly extensions for Kubernetes management.

## Features

- **100% kubectl compatibility:** All arguments are passed directly to `kubectl` by default.
- **Aliases:** Use `kubiq`, `kbq`, or `kb` as the command. Users can also create their own aliases for `kubiq`.
- **Cross-platform:** Works on Windows, macOS, and Linux.
- **Extensions:** Adds advanced features, such as wildcard deletion:
  - Example: `kubiq delete pods *abc* -n *aaa*` deletes all pods containing `abc` in namespaces containing `aaa`.
- **Developer-friendly:** Designed to simplify debugging and scripting, reducing the need for complex shell commands or manual filtering.

## How It Works

1. **No Arguments:**
   - Shows the standard `kubectl` help plus help for kubiq extensions.
2. **With Arguments:**
   - Passes all arguments to `kubectl` and runs the command.
   - If `kubectl` returns an error (e.g., unsupported command), kubiq checks for extension handlers.
   - Example: For advanced delete commands, kubiq will:
     - Parse wildcards.
     - Use `kubectl get` to list resources.
     - Filter results.
     - Run `kubectl delete` on matches.

## Getting Started

1. **Install kubectl:**
   - kubiq requires `kubectl` to be installed and available in your PATH.
2. **Download kubiq:**
   - (Instructions to be added for downloading binaries or building from source)
3. **Usage:**
   - Use `kubiq` just like `kubectl`, or with its extended features.
   - Example:
     ```sh
     kubiq get pods
     kubiq delete pods *test* -n *dev*
     ```

## Aliases

- You can use `kbq` or `kb` as shortcuts for `kubiq`.
- User-defined aliases are supported for `kubiq` (not for `kubectl`).

## Contributing

Contributions are welcome! Please see the [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the terms of the MIT License. See [LICENSE](LICENSE) for details.

---

For more details, see the [kubiq-architecture.md](kubiq-architecture.md) and project documentation.
