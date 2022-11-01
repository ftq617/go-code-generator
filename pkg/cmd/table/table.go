package table

import "github.com/spf13/cobra"

func NewCmdTable() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "table",
		Short: "Manage table",
	}

	// Register child command
	cmd.AddCommand(NewCmdList())
	cmd.AddCommand(NewCmdShow())

	return cmd
}
