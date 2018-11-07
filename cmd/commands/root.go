package commands

import (
	"os"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ifr",
		Short: "IFR (Instrument Flight Rules) validates resource definitions",
		Long: `Resource definition validator
Complete documentation is available at http://github.com/waterborne-labs/ifr`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}

	return rootCmd
}
