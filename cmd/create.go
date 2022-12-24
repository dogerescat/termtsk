package cmd

import (
	"log"
	"termtsk/ui/form"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "new",
	Short: "create",
	Long:  "create",
	Run: func(cmd *cobra.Command, args []string) {
		task := &form.Task{}
		task = form.Run(task)
		if task != nil {
			err := database.Create(task)
			if err != nil {
				log.Println(err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
