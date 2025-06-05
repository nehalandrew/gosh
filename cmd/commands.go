// ==========================================================================================
// File: cmd/commands.go
// This file defines the Commander interface and manages the command registry.
package cmd

// Commander interface defines the contract for all executable commands.
// Each command is a stand-alone object.
type Commander interface {
	Execute(args []string) error // Executes the command's action.
	Description() string         // Returns a brief description of the command.
	Usage() string               // Returns the usage syntax for the command.
}

// commandRegistry stores all available commands, mapping their names to their Commander objects.
// It is kept private to the cmd package to enforce usage of GetCommand.
var commandRegistry map[string]Commander

// InitRegistry populates the commandRegistry with instances of all concrete commands.
// This function acts as the "Client" in the Command pattern, creating and configuring commands.
func InitRegistry() {
	commandRegistry = make(map[string]Commander) // Initialize the global map.

	// HelpCommand needs a reference to the registry map itself.
	// Note: HelpCommand is defined in its own file (cmd/help.go) but registered here.
	helpCmd := NewHelpCommand(&commandRegistry) // Assuming NewHelpCommand is public in this package
	commandRegistry["help"] = helpCmd

	commandRegistry["exit"] = &ExitCommand{}   // Assuming ExitCommand is public
	commandRegistry["echo"] = &EchoCommand{}   // Assuming EchoCommand is public
	commandRegistry["whoami"] = &WhoamiCommand{} // Assuming WhoamiCommand is public
}

// GetCommand retrieves a command by its name from the registry.
// It returns the command and a boolean indicating if the command was found.
func GetCommand(name string) (Commander, bool) {
	cmd, exists := commandRegistry[name]
	return cmd, exists
}