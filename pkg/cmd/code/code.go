package code

import "github.com/spf13/cobra"

func NewCmdCode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "code",
		Short: "Manage code",
	}

	// Register child command
	cmd.AddCommand(NewCmdCreate())

	return cmd
}
