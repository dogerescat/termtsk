package cmd

import (
	"termtsk/ui/flex"

	"github.com/spf13/cobra"
)

type Flex struct {
}

var flexCmd = &cobra.Command{
	Use:   "flex",
	Short: "flex",
	Long:  "flex",
	Run: func(cmd *cobra.Command, args []string) {
		flex.Run()
	},
}

func init() {
	rootCmd.AddCommand(flexCmd)
}
