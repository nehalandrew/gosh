// ==========================================================================================
// File: cmd/whoami.go
// This file contains the implementation for the 'whoami' command.
package cmd

import (
	"fmt"
	"os"
	"os/user"
)

// WhoamiCommand implements the Commander interface for the "whoami" command.
type WhoamiCommand struct{}

// Execute attempts to print the username of the current user.
func (c *WhoamiCommand) Execute(args []string) error {
	currentUser, err := user.Current()
	if err != nil {
		// Log the error to stderr for better diagnostics.
		fmt.Fprintf(os.Stderr, "Error retrieving current user: %v\n", err)
		// Return a more user-friendly error message, wrapping the original error.
		return fmt.Errorf("could not determine current user: %w", err)
	}
	fmt.Println(currentUser.Username)
	return nil
}

// Description returns the description for the whoami command.
func (c *WhoamiCommand) Description() string { return "Print the effective user name." }

// Usage returns the usage string for the whoami command.
func (c *WhoamiCommand) Usage() string { return "whoami" }
