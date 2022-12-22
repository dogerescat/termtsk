package cmd

import (
	"fmt"
	"termtsk/ui/form"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create",
	Long:  `create`,
	Run: func(cmd *cobra.Command, args []string) {
		task := form.Run()
		fmt.Println(task)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
