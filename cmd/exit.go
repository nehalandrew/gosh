// ==========================================================================================
// File: cmd/exit.go
// This file contains the implementation for the 'exit' command.
package cmd

import "fmt"

// ExitCommand implements the Commander interface for the "exit" command.
type ExitCommand struct{}

// Execute prints a goodbye message and returns ErrExitCLI to signal termination.
func (c *ExitCommand) Execute(args []string) error {
	fmt.Println("Goodbye!")
	return ErrExitCLI // Signal to the main loop to terminate.
}

// Description returns the description for the exit command.
func (c *ExitCommand) Description() string { return "Exit this shell." }

// Usage returns the usage string for the exit command.
func (c *ExitCommand) Usage() string { return "exit" }
