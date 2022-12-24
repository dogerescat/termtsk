package cmd

import (
	"fmt"
	"log"
	"termtsk/ui/form"
	"termtsk/ui/list"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Long:  "list",
	Run: func(cmd *cobra.Command, args []string) {
		taskList := database.GetAll()
		n := list.Run(taskList)
		fmt.Println(n)
		if n >= 0 {
			task := form.Run(taskList[n])
			if task != nil {
				err := database.Update(task)
				if err != nil {
					log.Println(err.Error())
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
