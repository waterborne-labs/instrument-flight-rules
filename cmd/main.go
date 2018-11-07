package main

import (
	commandsRoot "github.com/waterborne-labs/instrument-flight-rules/cmd/commands"
	"github.com/gobuffalo/packr"
	"github.com/waterborne-labs/instrument-flight-rules/cmd/commands/validate"
	"fmt"
	"os"
)

func main() {
	box := packr.NewBox("../schemas")

	rootCmd := commandsRoot.NewRootCmd()
	rootCmd.AddCommand(validate.NewValidateCmd(box))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
