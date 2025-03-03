package daiteapcli

import (
	"github.com/spf13/cobra"
)

var cloudcredentialCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "cloud-credential",
    Aliases: []string{"ccred"},
    Short:  "Command to interact with cloud credentials from current tenant",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
            printHelpAndExit(cmd)
        }
        return
    },
}

func init() {
    rootCmd.AddCommand(cloudcredentialCmd)
}