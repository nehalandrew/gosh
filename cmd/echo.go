// ==========================================================================================
// File: cmd/echo.go
// This file contains the implementation for the 'echo' command.
package cmd

import (
	"fmt"
	"strings"
)

// EchoCommand implements the Commander interface for the "echo" command.
type EchoCommand struct{}

// Execute prints back the arguments provided to it.
func (c *EchoCommand) Execute(args []string) error {
	if len(args) == 0 {
		fmt.Println("Nothing to echo. Please provide some text.")
		fmt.Printf("Usage: %s\n", c.Usage()) // Refer to its own Usage method.
		return nil                          // Not an error, just misuse.
	}
	fmt.Println(strings.Join(args, " "))
	return nil
}

// Description returns the description for the echo command.
func (c *EchoCommand) Description() string { return "Display a line of text provided as arguments." }

// Usage returns the usage string for the echo command.
func (c *EchoCommand) Usage() string { return "echo [text to display ...]" }
