// ==========================================================================================
// File: cmd/help.go
// This file contains the implementation for the 'help' command.
package cmd

import (
	"fmt"
	"sort"
)

// HelpCommand implements the Commander interface for the "help" command.
type HelpCommand struct {
	// registry is a pointer to the global commandRegistry.
	// This allows HelpCommand to iterate over all registered commands.
	registry *map[string]Commander // This refers to the package-level commandRegistry
}

// NewHelpCommand creates a new HelpCommand.
// It requires a pointer to the command registry to list available commands.
func NewHelpCommand(reg *map[string]Commander) *HelpCommand {
	return &HelpCommand{registry: reg}
}

// Execute prints the help information for all registered commands.
func (c *HelpCommand) Execute(args []string) error {
	fmt.Println("Available commands:")
	if c.registry == nil || len(*c.registry) == 0 {
		fmt.Println("  No commands registered.")
		return nil
	}

	// Create a slice of command names to sort them for consistent output.
	cmdNames := make([]string, 0, len(*c.registry))
	for name := range *c.registry {
		cmdNames = append(cmdNames, name)
	}
	sort.Strings(cmdNames) // Sort command names alphabetically.

	for _, name := range cmdNames {
		cmd := (*c.registry)[name] // Get the command from the map via the pointer.
		fmt.Printf("  %-12s - %s\n", name, cmd.Description())
		usage := cmd.Usage()
		// If Usage provides more details than just the command name, print it.
		if usage != "" && usage != name { // Avoid printing usage if it's just the command name.
			fmt.Printf("               Usage: %s\n", usage)
		}
	}
	return nil
}

// Description returns the description for the help command.
func (c *HelpCommand) Description() string { return "Show this help message." }

// Usage returns the usage string for the help command.
func (c *HelpCommand) Usage() string { return "help" }
