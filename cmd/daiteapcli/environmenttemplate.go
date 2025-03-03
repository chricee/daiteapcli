package daiteapcli

import (
	"github.com/spf13/cobra"
)

var environmenttemplateCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "environment-template",
    Aliases: []string{"temp"},
    Short:  "Command to interact with environment templates from current tenant",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
            printHelpAndExit(cmd)
        }
        return
    },
}

func init() {
    rootCmd.AddCommand(environmenttemplateCmd)
}