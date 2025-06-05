// File: main.go
// This file contains the main application logic, including the input loop and command dispatch.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	// Import the cmd package from the current module
	// Assuming your module is named something like "gosh"
	// and the cmd package is in a subdirectory "cmd"
	// For local development, if main.go is in the root and cmd is a subdirectory,
	// you might run "go mod init <your_module_name>" in the root
	// and then use "<your_module_name>/cmd"
	// For simplicity in this example, we'll assume a local package structure
	// that Go can resolve. If you're building a real module, adjust the import path.
	// For now, let's assume `cmd` is a package in the same directory for this example,
	// or use a relative import if a proper module isn't set up.
	// However, direct relative imports like "./cmd" are discouraged in module mode.
	// The best practice is to set up a go.mod file.
	// For this snippet, we'll write it as if "cmd" is a package that Go can find.
	"gosh/cmd" // Replace "gosh" with your actual module name
)

func main() {
	// Initialize commands from the cmd package
	cmd.InitRegistry()

	outputWriter := os.Stdout // For normal output.
	errorWriter := os.Stderr  // For errors and prompts.
	prompt := "gosh> "

	fmt.Fprintln(outputWriter, "Welcome to My Custom CLI Shell!")
	fmt.Fprintln(outputWriter, "Type 'help' for available commands, or 'exit' to quit.")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(outputWriter, prompt) // Print the prompt.

		input, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) { // Handle Ctrl+D or client disconnect.
				fmt.Fprintln(outputWriter, "\nSession ended.")
				os.Exit(0) // Normal exit.
			}
			// For other critical read errors.
			fmt.Fprintf(errorWriter, "Critical error reading input: %v\n", err)
			os.Exit(1) // Exit with an error code.
		}

		input = strings.TrimSpace(input)
		if input == "" { // Skip empty input lines.
			continue
		}

		parts := strings.Fields(input) // Split input into command name and arguments.
		commandName := parts[0]
		var args []string
		if len(parts) > 1 {
			args = parts[1:]
		}

		// Look up the command in the registry.
		if command, exists := cmd.GetCommand(commandName); exists {
			// Execute the command.
			execErr := command.Execute(args)
			if execErr != nil {
				if errors.Is(execErr, cmd.ErrExitCLI) { // Use the error from the cmd package
					// The command (ExitCommand) signaled a clean exit.
					os.Exit(0)
				}
				// Print other execution errors to stderr.
				fmt.Fprintf(errorWriter, "Error executing command '%s': %v\n", commandName, execErr)
			}
		} else {
			fmt.Fprintf(errorWriter, "Unknown command: '%s'. Type 'help' for available commands.\n", commandName)
		}
	}
}