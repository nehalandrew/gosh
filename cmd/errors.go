// ==========================================================================================
// File: cmd/errors.go
// This file defines custom errors used by the commands package.
package cmd

import "errors"

// ErrExitCLI is a special error used to signal that the CLI should terminate.
var ErrExitCLI = errors.New("user requested exit")
